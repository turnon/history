package main

import (
	"flag"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/turnon/history/db"
	"github.com/turnon/history/statistic"
)

func main() {
	// parse flag
	logSql := flag.Bool("sql", false, "log all sql")
	flag.Parse()

	dbPath := flag.Arg(0)
	data := db.Conn(dbPath)

	if *logSql {
		fmt.Printf("connecting: %s\n", dbPath)
		data.DebugMode()
	}

	// boot server
	r := gin.Default()
	r.GET("/", func(ctx *gin.Context) {
		cond := db.Condition{}
		ctx.ShouldBindQuery(&cond)
		visits := data.Visits(cond)

		st := statistic.NewVisitsPerDate(
			func(v *db.Visit) string { return v.Host },
			func(v *db.Visit) int { return 1 },
			cond.VisitTimeGte,
			cond.VisitTimeLte,
		)
		st.AddVisits(visits)
		st.Render(ctx.Writer)
	})
	r.Run()
}
