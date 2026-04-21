package mlb

import (
	"context"
	"fmt"
	"pick_and_go/database/sqlc"
)

func (client *SportClient) GetAllDivisions() error {
	endpoint := "/api/v1/divisions"
	var allDivisions AllDivisionsJSON
	if err := client.RequestAndDecode(endpoint, &allDivisions); err != nil {
		return err
	}

	for _, division := range allDivisions.Divisions {
		if division.Sport.ID != 1 {
			continue
		}
		if err := client.Queries.CreateDivision(context.Background(), sqlc.CreateDivisionParams{ID: int32(division.ID), Name: division.NameShort}); err != nil {
			return fmt.Errorf("Couldn't insert values into divisions table: %w", err)
		}
	}
	return nil
}
