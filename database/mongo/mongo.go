package mongo

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"github.com/akinsella/go-playground/models"
	"github.com/akinsella/go-playground/tasks"
)

type MongoStopTimesImportTask struct {
	tasks.ImportTask
}

func (m *MongoStopTimesImportTask) DoWork(workRoutine int) {
	m.InsertStopTimes(stopTimesInserter);
}

func stopTimesInserter(sts *models.StopTimes) (error)  {

	mSession := GetSession()

	defer mSession.Close()

	c := mSession.DB("gtfs").C("stop_times")

	bulk := c.Bulk()

	for _, st := range sts.Records {
		bulk.Insert(st)
	}

	_, err := bulk.Run()

	return err
}

var (
	mgoSession *mgo.Session
)

func GetSession() *mgo.Session {
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

func WithCollection(databaseName string, collectionName string, fn func(*mgo.Collection) error) {
	session := GetSession()
	defer session.Close()

	collection := session.DB(databaseName).C(collectionName)

	if err := fn(collection); err != nil {
		panic(err)
	}
}

func FindAll(databaseName string, collectionName string) (results[]interface{}) {
	session := GetSession()
	defer session.Close()

	collection := session.DB(databaseName).C(collectionName)
	err := collection.Find(bson.M{}).All(&results)
	if err != nil {
		panic(err)
	}

	fmt.Println("Results: ", results)

	return results
}
