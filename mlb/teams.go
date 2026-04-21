package mlb

import (
	"context"
	"fmt"
	"pick_and_go/database/sqlc"
)

func (client *SportClient) GetAllTeams() error {
	endpoint := "/api/v1/teams?sportId=1"
	var teams AllTeamsJSON

	if err := client.RequestAndDecode(endpoint, &teams); err != nil {
		return err
	}

	for _, team := range teams.Teams {
		if err := client.Queries.CreateTeam(context.Background(), sqlc.CreateTeamParams{ID: int32(team.ID), TeamName: team.Name, Abbreviation: team.Abbreviation, DivisionID: int32(team.Division.ID)}); err != nil {
			return fmt.Errorf("Couldn't insert values into teams table: %w", err)
		}
	}
	return nil
}
