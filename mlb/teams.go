package mlb

import (
	"context"
	"encoding/json"
	"fmt"
	"pick_and_go/database/sqlc"
)

func (client *SportClient) GetAllTeams() error {
	endpoint := "/api/v1/teams?sportId=1"
	url := buildURL(endpoint)

	res, err := client.Get(url)
	if err != nil {
		return fmt.Errorf("Request to URL: %s failed.", url)
	}

	defer res.Body.Close()
	decoder := json.NewDecoder(res.Body)

	var data AllTeamsJSON
	if err := decoder.Decode(&data); err != nil {
		return fmt.Errorf("Couldn't decode JSON into teams struct")
	}

	for _, team := range data.Teams {
		if err := client.Queries.CreateTeam(context.Background(), sqlc.CreateTeamParams{ID: int32(team.ID), TeamName: team.Name, Abbreviation: team.Abbreviation, DivisionID: int32(team.Division.ID)}); err != nil {
			return fmt.Errorf("Couldn't insert values into teams table: %w", err)
		}
	}
	return nil
}
