package toggl

import (
	"encoding/json"
	"log"
	"strconv"
)

type SummaryItem struct {
	Project     string
	Title       string
	TimeMinutes int
}

type JsonSummary struct {
	TotalGrand    int               `json:"total_grand"`
	TotalBillable int               `json:"total_billable"`
	Data          []JsonSummaryData `json:"data"`
}

type JsonSummaryData struct {
	Id    int                   `json:"id"`
	Time  int                   `json:"time"`
	Items []JsonSummaryDataItem `json:"items"`
	Title struct {
		Project  string `json:"project"`
		Client   string `json:"client"`
		Color    string `json:"color"`
		HexColor string `json:"hex_color"`
	}
}

type JsonSummaryDataItem struct {
	Time     int    `json:"time"`
	Currency string `json:"cur"`
	Sum      int    `json:"sum"`
	Rate     int    `json:"rate"`
	Title    struct {
		TimeEntry string `json:"time_entry"`
	}
}

func (c Client) GetSummary(ws *Workspace, since string, until string) []SummaryItem {
	url := "https://api.track.toggl.com/reports/api/v2/summary"
	params := map[string]string{
		"workspace_id": strconv.Itoa(ws.Id),
		"since":        since,
		"until":        until,
		"user_agent":   "hawkeye",
	}
	bytes := c.getRequest(url, params)

	var jsonSummary JsonSummary
	if err := json.Unmarshal(bytes, &jsonSummary); err != nil {
		log.Fatal(err)
	}

	var summary []SummaryItem
	for _, data := range jsonSummary.Data {
		for _, item := range data.Items {
			time_minutes := item.Time / 1000 / 60 // item.Time is in milliseconds
			si := SummaryItem{
				Project:     data.Title.Project,
				Title:       item.Title.TimeEntry,
				TimeMinutes: time_minutes,
			}
			summary = append(summary, si)
		}
	}

	return summary
}
