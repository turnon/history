package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
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

type Condition struct {
	Limit int
	Url   string
	Title string
}

type Db struct {
	*gorm.DB
}

func Conn(path string) *Db {
	db, err := gorm.Open(sqlite.Open(path), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return &Db{db}
}

func (db *Db) DebugMode() {
	db.DB = db.DB.Debug()
}

func (db *Db) Visits(cond *Condition) []*Visit {
	joining := db.Order("visits.id desc").Joins("Link")

	if cond.Limit != 0 {
		joining = joining.Limit(cond.Limit)
	}
	if cond.Url != "" {
		joining = joining.Where("Link.url like ?", "%"+cond.Url+"%")
	}
	if cond.Title != "" {
		joining = joining.Where("Link.title like ?", "%"+cond.Title+"%")
	}

	var visits []*Visit
	joining.Find(&visits)
	return visits
}
