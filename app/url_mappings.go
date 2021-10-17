package app

import (
	"github.com/rahul-kr/page-statistics/controllers/ping"
	"github.com/rahul-kr/page-statistics/controllers/statistics"
)

func mapUrls() {
	// filter result based on query parameters
	router.HandleFunc("/ping", ping.Ping).Methods("GET")
	// filter result based on query parameters
	router.HandleFunc("/counter/v1/statistics/article_id/{article_id}", statistics.ShowStatistics).Methods("GET")
	// Create sample data
	router.HandleFunc("/counter/v1/statistics", statistics.InsertStatistics).Methods("POST")

}
