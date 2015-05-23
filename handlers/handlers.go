package handlers

////////////////////////////////////////////////////////////////////////////////////////////////
/// Imports
////////////////////////////////////////////////////////////////////////////////////////////////

import(
	"io"
    "fmt"
	"net/http"
	"time"
	"log"
	"sync"
	"github.com/gorilla/handlers"
	"github.com/PuerkitoBio/throttled"
	"github.com/PuerkitoBio/throttled/store"
	"github.com/helyx-io/commute-importer/session"
    "github.com/helyx-io/commute-importer/config"
    "github.com/garyburd/redigo/redis"
)


////////////////////////////////////////////////////////////////////////////////////////////////
/// Handlers
////////////////////////////////////////////////////////////////////////////////////////////////

func TimeoutHandler(h http.Handler) http.Handler {
	return http.TimeoutHandler(h, 1 * time.Second, "timed out")
}

func LoggedInHandler(h http.Handler) http.Handler  {
	return http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Has Token: ", session.HasToken(r))
		if !session.HasToken(r) {
			http.Redirect(w, r, "/login.html", 302)
		}
	}))
}

func setupRedis(redisInfos *config.RedisConfig) *redis.Pool {
	pool := &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 30 * time.Second,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", fmt.Sprintf("%s:%d", redisInfos.Host, redisInfos.Port))
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}

	return pool
}


func ThrottleHandler(h http.Handler, redisInfos *config.RedisConfig) http.Handler {

	var mu sync.Mutex
	var ok, ko int

//	start := time.Now()

	quota := throttled.Q{Requests: 100, Window: time.Minute}
	varyBy := throttled.VaryBy{Path: true}
	st := store.NewRedisStore(setupRedis(redisInfos), "throttled:", 0)

	t := throttled.RateLimit(quota, &varyBy, st)

	// Set its denied handler
	t.DeniedHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		log.Printf("KO: %s", time.Since(start))

		throttled.DefaultDeniedHandler.ServeHTTP(w, r)

		mu.Lock()
		defer mu.Unlock()
		ko++
	})

//	go func() {
//		for _ = range time.Tick(10 * time.Second) {
//			mu.Lock()
//			log.Printf("ok: %d, ko: %d", ok, ko)
//			mu.Unlock()
//		}
//	}()

	return t.Throttle(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		log.Printf("ok: %s", time.Since(start))

		h.ServeHTTP(w, r)

		mu.Lock()
		defer mu.Unlock()
		ok++
	}))
}

func LoggingHandler(out io.Writer) (func(h http.Handler) http.Handler) {
	return func(h http.Handler) http.Handler {
		return handlers.LoggingHandler(out, h)
	}
}
