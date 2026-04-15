package mlb

import (
	"context"
	"encoding/json"
	"fmt"
	"pick_and_go/database/sqlc"
)

func (client *SportClient) GetAllDivisions() error {
	endpoint := "/api/v1/divisions"
	url := fmt.Sprintf("%s%s", reqPrefix, endpoint)

	res, err := client.Get(url)
	if err != nil {
		return fmt.Errorf("Request to URL: %s failed.", url)
	}

	defer res.Body.Close()
	decoder := json.NewDecoder(res.Body)
	var allDivisions AllDivisionsJSON

	if err := decoder.Decode(&allDivisions); err != nil {
		return fmt.Errorf("Couldn't decode JSON into divisions struct")
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
