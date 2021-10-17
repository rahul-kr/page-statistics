package statistics

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/rahul-kr/page-statistics/datasources/mysql/mta_db"
)

const articleType = "statistics_article_view_content"

type intervalCount struct {
	Reference string `json:"reference"`
	RowCount  int64  `json:"count"`
}

type attributes struct {
	Attributes []intervalCount `json:"count,omitempty"`
}
type statisticsResponseData struct {
	ArticleId   string     `json:"article_id,omitempty"`
	ArticleType string     `json:"type,omitempty"`
	Attributes  attributes `json:"attributes,omitempty"`
}

type statisticsAPIResponse struct {
	ResponseData statisticsResponseData `json:"data"`
}

type rowCount struct {
	Count int64
}

/*
* GET API
* endpoint /counter/v1/statistics/article_id/{article_id}
 */
func ShowStatistics(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	inputMapCount := params["article_id"]

	var statisticsAPIResponse statisticsAPIResponse
	var statisticsResponseData statisticsResponseData
	var result attributes
	var row intervalCount
	var count rowCount

	intervals := []string{"420 MINUTE", "10 HOUR", "2 DAY"}

	for _, v := range intervals {
		mta_db.Client.Raw("SELECT count(article_id) AS count FROM view_maps WHERE article_id = ? AND date_time >= NOW() - INTERVAL "+v+" ;", inputMapCount).Scan(&count)
		row.RowCount = count.Count
		row.Reference = v + " ago"
		result.Attributes = append(result.Attributes, row)
	}
	statisticsResponseData.ArticleId = inputMapCount
	statisticsResponseData.ArticleType = articleType
	statisticsResponseData.Attributes = result

	statisticsAPIResponse.ResponseData = statisticsResponseData
	json.NewEncoder(w).Encode(statisticsAPIResponse)
}

/*
* POST API
* endpoint /counter/v1/statistics
 */
func InsertStatistics(w http.ResponseWriter, r *http.Request) {
	var viewMap mta_db.ViewMap
	var datetime = time.Now()
	dt := datetime.Format(time.RFC3339)
	viewMap.DateTime = dt
	viewMap.Type = articleType
	json.NewDecoder(r.Body).Decode(&viewMap)
	// Insert into table
	mta_db.Client.Create(&viewMap)
	//prepare response
	var statisticsResponseData statisticsResponseData
	var statisticsAPIResponse statisticsAPIResponse
	w.Header().Set("Content-Type", "application/json")
	statisticsResponseData.ArticleId = viewMap.ArticleId
	statisticsAPIResponse.ResponseData = statisticsResponseData
	json.NewEncoder(w).Encode(statisticsAPIResponse)
}
