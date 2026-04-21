package mlb

import (
	"context"
	"fmt"
	"pick_and_go/database/sqlc"
)

func (client *SportClient) GetTeamRecords() error {
	endpoint := "/api/v1/standings?leagueId=103,104&season=2026"
	var records TeamRecordsJSON

	if err := client.RequestAndDecode(endpoint, &records); err != nil {
		return err
	}

	for _, register := range records.Records {
		for _, leagueTeamRecords := range register.TeamRecords {
			parameters := sqlc.CreateTeamRecordParams{}
			parameters.TeamID = int32(leagueTeamRecords.Team.ID)
			parameters.Wins = int32(leagueTeamRecords.LeagueRecord.Wins)
			parameters.Losses = int32(leagueTeamRecords.LeagueRecord.Losses)
			parameters.Pct = leagueTeamRecords.LeagueRecord.Pct
			parameters.Streak = leagueTeamRecords.Streak.StreakCode
			parameters.RunsAgainst = int32(leagueTeamRecords.RunsAllowed)
			parameters.RunsScored = int32(leagueTeamRecords.RunsScored)

			splitRecords := leagueTeamRecords.Records.SplitRecords
			for _, record := range splitRecords {
				if record.Type == "home" {
					parameters.HomeWins = int32(record.Wins)
					parameters.HomeLosses = int32(record.Losses)
				} else if record.Type == "away" {
					parameters.AwayWins = int32(record.Wins)
					parameters.AwayLosses = int32(record.Losses)
				}
			}
			if err := client.Queries.CreateTeamRecord(context.Background(), parameters); err != nil {
				return fmt.Errorf("Couldn't insert record into records table: %w", err)
			}
		}
	}

	return nil
}
