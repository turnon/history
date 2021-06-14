package dba

import (
	"encoding/json"
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Visit struct {
	ID            int
	Url           int
	Link          Url `gorm:"foreignKey:Url"`
	VisitTime     int
	VisitDuration int
}

type Url struct {
	ID            int
	Url           string
	Title         string
	VisitCount    int
	LastVisitTime int
}

func Q() {
	dbPath := "../tmp/History.db"
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	var visits []Visit
	var count int64
	db.Order("id desc").Limit(10).Preload(clause.Associations).Find(&visits)
	db.Model(&Visit{}).Count(&count)

	bytes, _ := json.MarshalIndent(visits, "", "  ")
	fmt.Println(string(bytes))
	fmt.Println(count)
}
