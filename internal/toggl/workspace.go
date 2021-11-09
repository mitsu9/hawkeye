package toggl

import (
	"encoding/json"
	"errors"
	"log"
)

type Workspace struct {
	Id                          int    `json:"id"`
	Name                        string `json:"name"`
	Profile                     int    `json:"profile"`
	Premium                     bool   `json:"premium"`
	Admin                       bool   `json:"admin"`
	DefaultHourlyRate           int    `json:"default_hourly_rate"`
	DefaultCurrency             string `json:"default_currency"`
	OnlyAdminsMayCreateProjects bool   `json:"only_admins_may_create_projects"`
	OnlyAdminsSeeBillableRate   bool   `json:"only_admins_see_billable_rates"`
	OnlyAdminsSeeTeamDashboard  bool   `json:"only_admins_see_team_dashboard"`
	ProjectsBillableByDefault   bool   `json:"projects_billable_by_default"`
	Rounding                    int    `json:"rounding"`
	RoundingMinutes             int    `json:"rounding_minutes"`
	ApiToken                    string `json:"api_token"`
	CreatedAt                   string `json:"at"`
	IcalEnabled                 bool   `json:"ical_enabled"`
}

func (c Client) GetWorkspaces() []Workspace {
	url := "https://api.track.toggl.com/api/v8/workspaces"
	bytes := c.getRequest(url, nil)

	var workspaces []Workspace
	if err := json.Unmarshal(bytes, &workspaces); err != nil {
		log.Fatal(err)
	}

	return workspaces
}

func (c Client) GetWorkspace(name string) (*Workspace, error) {
	workspaces := c.GetWorkspaces()
	for _, ws := range workspaces {
		if ws.Name == name {
			return &ws, nil
		}
	}
	return nil, errors.New("Not found workspace (name: " + name + ")")
}
