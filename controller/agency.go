package controller

import (
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"fmt"
	"net/http"
	"encoding/json"
)

type AgencyController struct { }

func (agencyController *AgencyController) Init(r *mux.Router) {
	r.HandleFunc("/", agencyController.Agencies)
	r.HandleFunc("/{id:[0-9]+}", agencyController.AgenciesByKey)
}

type Agency struct {
	Id bson.ObjectId `bson:"_id" json:"id"`
	AgencyId string `bson:"agency_id" json:"agencyId"`
	Name string `bson:"agency_name" json:"name"`
	Url string `bson:"agency_url" json:"url"`
	Timezone string `bson:"agency_timezone" json:"timezone"`
	Lang string `bson:"agency_lang" json:"lang"`
	Key string `bson:"agency_key" json:"key"`
}

var (
	mgoSession *mgo.Session
)

func getSession() *mgo.Session {
	if mgoSession == nil {
		var err error
		//	mgoSession, err := mgo.Dial("localhost:27017,localhost:27018,localhost:27019,localhost:27020")
		mgoSession, err = mgo.Dial("localhost:27017?maxPoolSize=500")
		if err != nil {
			panic(err)
		}
	}

	return mgoSession.Clone()
}

func withCollection(databaseName string, collectionName string, fn func(*mgo.Collection) error) {
	session := getSession()
	defer session.Close()

	collection := session.DB(databaseName).C(collectionName)

	if err := fn(collection); err != nil {
		panic(err)
	}
}

func findAll(databaseName string, collectionName string) (results[]interface{}) {
	session := getSession()
	defer session.Close()

	collection := session.DB(databaseName).C(collectionName)
	err := collection.Find(bson.M{}).All(&results)
	if err != nil {
		panic(err)
	}

	fmt.Println("Results: ", results)

	return results
}

func sendJson(w http.ResponseWriter, data interface{}) {
	bytes, err := json.MarshalIndent(data, "", "  ")
	if  err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(bytes)
}

func (ac *AgencyController) Agencies(w http.ResponseWriter, r *http.Request) {
	results := findAll("gtfs", "agency")
	sendJson(w, results)
}

func (ac *AgencyController) AgenciesByKey(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	agencyKey := params["agencyKey"]

	log.Printf("Agency Key: %s", agencyKey)
	w.Write([]byte("Agency Key: " + agencyKey))
}
