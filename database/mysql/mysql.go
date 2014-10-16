package mysql

import (
	"fmt"
	"strings"
	"database/sql"
	"github.com/akinsella/go-playground/models"
	"github.com/akinsella/go-playground/tasks"
	_ "github.com/go-sql-driver/mysql"
)

type MySQL struct {}


type MySQLStopTimesImportTask struct {
	tasks.ImportTask
}

func (m *MySQLStopTimesImportTask) DoWork(workRoutine int) {
	m.InsertStopTimes(insertStopTimes);
}

func insertStopTimes(sts *models.StopTimes) (error)  {

	db, err := sql.Open("mysql", "gtfs:gtfs@/gtfs?charset=utf8mb4,utf8");

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	valueStrings := make([]string, 0, len(sts.Records))
	valueArgs := make([]interface{}, 0, len(sts.Records) * 9)

	for _, st := range sts.Records {
		valueStrings = append(valueStrings, "('RATP', ?, ?, ?, ?, ?, ?, ?, ?)")
		valueArgs = append(
			valueArgs,
			st.TripId,
			st.ArrivalTime,
			st.DepartureTime,
			st.StopId,
			st.StopSequence,
			st.StopHeadSign,
			st.PickupType,
			st.DropOffType,
		)
	}

	stmt := fmt.Sprintf("INSERT INTO stop_times (agency_key, trip_id, arrival_time, departure_time, stop_id, stop_sequence, stop_head_sign, pickup_type, drop_off_type) VALUES %s", strings.Join(valueStrings, ","))

	_, err = db.Exec(stmt, valueArgs...)

	return err
}
