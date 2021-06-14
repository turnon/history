package db

import (
	"github.com/turnon/history/epoch"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const epochFormat = "2006-01-02"

type Visit struct {
	ID            int
	Url           int
	Link          Url `gorm:"foreignKey:Url"`
	VisitTime     int64
	VisitDuration int
}

type Url struct {
	ID            int
	Url           string
	Title         string
	VisitCount    int
	LastVisitTime int64
}

type Condition struct {
	Limit        int
	Url          string
	Title        string
	VisitTimeGte string
	VisitTimeLte string
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
	joining := db.Order("visits.visit_time desc").Joins("Link")

	if cond.Limit != 0 {
		joining = joining.Limit(cond.Limit)
	}
	if cond.Url != "" {
		joining = joining.Where("Link.url like ?", "%"+cond.Url+"%")
	}
	if cond.Title != "" {
		joining = joining.Where("Link.title like ?", "%"+cond.Title+"%")
	}
	if cond.VisitTimeGte != "" {
		joining = joining.Where("visits.visit_time >= ?", epoch.To(cond.VisitTimeGte, epochFormat))
	}
	if cond.VisitTimeLte != "" {
		joining = joining.Where("visits.visit_time <= ?", epoch.To(cond.VisitTimeLte, epochFormat))
	}

	var visits []*Visit
	joining.Find(&visits)
	return visits
}

func (v *Visit) VisitTimeStr() string {
	return epoch.From(v.VisitTime, epochFormat)
}
