package mlb

import (
	"context"
	"encoding/json"
	"fmt"
	"pick_and_go/database/sqlc"
	"time"
)

func (client *SportClient) GetSeasonSchedule() error {
	endpoint := "/api/v1/schedule?sportId=1&season=2026&gameType=R"
	url := buildURL(endpoint)

	res, err := client.Get(url)
	if err != nil {
		return fmt.Errorf("Request to URL: %s failed.", url)
	}

	defer res.Body.Close()
	decoder := json.NewDecoder(res.Body)
	var allGames AllGamesJSON

	if err := decoder.Decode(&allGames); err != nil {
		return fmt.Errorf("Couldn't decode JSON into games struct: %w", err)
	}
	type Game struct {
		date       time.Time
		gamePk     int
		homeTeamID int
		homeScore  int
		awayTeamID int
		awayScore  int
	}
	for _, date := range allGames.Dates {
		for _, game := range date.Games {
			if game.Status.DetailedState != "Final" {
				continue
			}
			params := sqlc.CreateNewGameParams{
				DatePlayed: timeToPgDate(game.GameDate),
				Gamepk:     int32(game.GamePk),
				HomeTeamID: int32(game.Teams.Home.Team.ID),
				HomeScore:  int32(game.Teams.Home.Score),
				AwayTeamID: int32(game.Teams.Away.Team.ID),
				AwayScore:  int32(game.Teams.Away.Score)}
			if err := client.Queries.CreateNewGame(context.Background(), params); err != nil {
				return fmt.Errorf("Couldn't insert game into games table: %w", err)
			}
		}
	}
	return nil
}

func (client *SportClient) GetLineScore(gamePk int) error {
	endpoint := fmt.Sprintf("https://statsapi.mlb.com/api/v1/game/%d/linescore", gamePk)
	url := buildURL(endpoint)

	res, err := client.Get(url)
	if err != nil {
		return fmt.Errorf("Request to URL: %s failed.", url)
	}
	defer res.Body.Close()
	return nil
}
