package agency

import (
	"github.com/gorilla/mux"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"log"
	"fmt"
	"net/http"
	"encoding/json"
)

type Agency struct {
	Id bson.ObjectId `bson:"_id" json:"id"`
	AgencyId string `bson:"agency_id" json:"agencyId"`
	Name string `bson:"agency_name" json:"name"`
	Url string `bson:"agency_url" json:"url"`
	Timezone string `bson:"agency_timezone" json:"timezone"`
	Lang string `bson:"agency_lang" json:"lang"`
	Key string `bson:"agency_key" json:"key"`
}

type AgenciesJSON struct {
	Agencies []Agency `json:"agencies"`
}

func Agencies(w http.ResponseWriter, r *http.Request) {

	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	agencies := []Agency{}
	c := session.DB("gtfs").C("agency")
	err = c.Find(bson.M{}).All(&agencies)
	if err != nil {
		panic(err)
	}

	fmt.Println("Agencies: ", agencies)

	bytes, err := json.MarshalIndent(agencies, "", "  ")
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(bytes)
}

func AgenciesByKey(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	agencyKey := params["agencyKey"]
	log.Printf("Agency Key: %s", agencyKey)
	w.Write([]byte("Agency Key: " + agencyKey))
}
