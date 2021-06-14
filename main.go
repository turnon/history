package main

import (
	"flag"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/turnon/history/db"
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

		ctx.JSON(200, gin.H{
			"result": visits,
		})
	})
	r.Run()
}
