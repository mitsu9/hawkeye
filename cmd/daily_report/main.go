package main

import (
	"fmt"
	"github.com/mitsu9/hawkeye/internal/toggl"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	apiToken := os.Getenv("TOGGL_API_TOKEN")
	workSpace := os.Getenv("TOGGL_WORK_SPACE")

	client := toggl.NewClient(apiToken)

	ws, err := client.GetWorkspace(workSpace)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("workspace:", ws.Id, ws.Name)

	now := time.Now()
	since := now.Format("2006-01-02")
	until := now.Add(24 * time.Hour).Format("2006-01-02")
	summary := client.GetSummary(ws, since, until)

	fmt.Printf("Reports (%s)\n", since)
	for _, item := range summary {
		hour := strconv.Itoa(item.TimeMinutes / 60)
		minutes := strconv.Itoa(item.TimeMinutes % 60)
		if item.TimeMinutes < 60 {
			fmt.Printf("- %s: %s (%smins)\n", item.Project, item.Title, minutes)
		} else {
			fmt.Printf("- %s: %s (%sh%smins)\n", item.Project, item.Title, hour, minutes)
		}
	}
}
