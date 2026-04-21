package mlb

import (
	"context"
	"fmt"
	"log"
	"pick_and_go/database/sqlc"
	"time"
)

func (client *SportClient) GetAllPlayers() error {
	endpoint := "/api/v1/sports/1/players?season=2026"
	var players AllPlayersJSON
	if err := client.RequestAndDecode(endpoint, &players); err != nil {
		return err
	}

	for _, player := range players.People {
		params := sqlc.CreateNewPlayerParams{
			ID:              int32(player.ID),
			FirstName:       player.FirstName,
			LastName:        player.LastName,
			Age:             int32(player.CurrentAge),
			IsActive:        player.Active,
			TeamID:          int32(player.CurrentTeam.ID),
			PrimaryPosition: player.PrimaryPosition.Abbreviation,
			Batside:         player.BatSide.Code,
			Pitchhand:       player.PitchHand.Code,
		}
		if err := client.Queries.CreateNewPlayer(context.Background(), params); err != nil {
			return fmt.Errorf("Couldn't insert values into players table: %w", err)
		}
		if err := client.GetPitchingStats(int32(player.ID)); err != nil {
			return err
		}
		time.Sleep(200 * time.Millisecond)
	}
	return nil
}

func (client *SportClient) GetPitchingStats(playerID int32) error {
	endpoint := fmt.Sprintf("/api/v1/people/%d/stats?stats=gameLog&group=pitching&season=2026", playerID)
	var pitchingStats PitchingStatsJSON
	if err := client.RequestAndDecode(endpoint, &pitchingStats); err != nil {
		return err
	}
	for _, data := range pitchingStats.Stats {
		for _, game := range data.Splits {
			log.Printf("Looking up game_pk: %d", game.Game.GamePk)
			_, err := client.Queries.GetGameByGPK(context.Background(), int32(game.Game.GamePk))
			if err != nil {
				log.Printf("Game %d not found: %v", game.Game.GamePk, err)
				continue
			}
			params := sqlc.CreatePitchingEntryParams{
				PlayerID:           playerID,
				Gamepk:             int32(game.Game.GamePk),
				Era:                game.Stat.Era,
				Whip:               game.Stat.Whip,
				InningsPitched:     game.Stat.InningsPitched,
				Strikeouts:         int32(game.Stat.Strikeouts),
				Walks:              int32(game.Stat.Walks),
				HomeRuns:           int32(game.Stat.HomeRuns),
				EarnedRuns:         int32(game.Stat.EarnedRuns),
				Hits:               int32(game.Stat.Hits),
				Wins:               int32(game.Stat.Wins),
				Losses:             int32(game.Stat.Losses),
				GamesStarted:       int32(game.Stat.GamesStarted),
				Saves:              int32(game.Stat.Saves),
				BlownSaves:         int32(game.Stat.BlownSaves),
				StrikeoutsPer9:     game.Stat.StrikeoutsPer9,
				WalksPer9:          game.Stat.WalksPer9,
				StrikeoutWalkRatio: game.Stat.StrikeoutWalkRatio,
			}
			if err := client.Queries.CreatePitchingEntry(context.Background(), params); err != nil {
				return fmt.Errorf("Couldn't insert values into pitching table: %w", err)
			}
		}
	}
	return nil
}
