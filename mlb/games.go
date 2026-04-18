package mlb

import (
	"context"
	"encoding/json"
	"fmt"
	"pick_and_go/database/sqlc"
	"time"
)

func (client *SportClient) GetGameResults() error {
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
			if err := client.GetLineScore(game.GamePk); err != nil {
				return err
			}
			time.Sleep(200 * time.Millisecond)
		}
	}
	return nil
}

func (client *SportClient) GetLineScore(gamePk int) error {
	endpoint := fmt.Sprintf("/api/v1/game/%d/linescore", gamePk)
	url := buildURL(endpoint)

	res, err := client.Get(url)
	if err != nil {
		return fmt.Errorf("Request to URL: %s failed.", url)
	}
	defer res.Body.Close()
	decoder := json.NewDecoder(res.Body)

	var lineScore LineScoreJSON
	if err := decoder.Decode(&lineScore); err != nil {
		return fmt.Errorf("Couldn't decode JSON into linescore struct: %w", err)
	}

	params := sqlc.CreateNewInningParams{
		Gamepk: int32(gamePk),
	}

	for _, inning := range lineScore.Innings {
		params.InningName = inning.OrdinalNum
		params.HomeRuns = int32(inning.Home.Runs)
		params.HomeHits = int32(inning.Home.Hits)
		params.HomeErrors = int32(inning.Home.Errors)
		params.AwayRuns = int32(inning.Away.Runs)
		params.AwayHits = int32(inning.Away.Hits)
		params.AwayErrors = int32(inning.Away.Errors)
		if err := client.Queries.CreateNewInning(context.Background(), params); err != nil {
			return fmt.Errorf("Couldn't insert inning into innings table: %w", err)
		}
	}
	return nil
}
