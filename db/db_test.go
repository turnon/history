package db

import (
	"testing"
)

var db *Db

func TestMain(m *testing.M) {
	db = Conn("../tmp/History.db")
	db.DebugMode()
	m.Run()
}
func TestVisits(t *testing.T) {
	db.Visits(&Condition{
		Limit: 1,
		Url:   "x",
		Title: "y",
	})
}
