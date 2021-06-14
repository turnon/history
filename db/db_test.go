package db

import "testing"

var db *Db

func TestMain(m *testing.M) {
	db = Conn("../tmp/History.db")
	m.Run()
}
func TestQuery(t *testing.T) {
	db.Query()
}
