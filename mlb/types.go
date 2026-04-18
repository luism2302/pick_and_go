package mlb

import (
	"context"
	"fmt"
	"net/http"
	"pick_and_go/database/sqlc"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

// general
const (
	reqPrefix = "https://statsapi.mlb.com"
)

type SportClient struct {
	http.Client
	Queries *sqlc.Queries
}

func NewSportClient(db sqlc.DBTX) *SportClient {
	return &SportClient{Queries: sqlc.New(db)}
}

func buildURL(apiEndpoint string) string {
	return fmt.Sprintf("%s%s", reqPrefix, apiEndpoint)
}

// Teams
type AllTeamsJSON struct {
	Copyright string `json:"copyright"`
	Teams     []struct {
		ID           int    `json:"id"`
		Name         string `json:"name"`
		Abbreviation string `json:"abbreviation"`
		Division     struct {
			ID int `json:"id"`
		} `json:"division"`
	}
}

// Records
type TeamRecordsJSON struct {
	Records []struct {
		TeamRecords []struct {
			Team struct {
				ID int `json:"id"`
			} `json:"team"`
			Streak struct {
				StreakCode string `json:"streakCode"`
			} `json:"streak"`
			LeagueRecord struct {
				Wins   int    `json:"wins"`
				Losses int    `json:"losses"`
				Ties   int    `json:"ties"`
				Pct    string `json:"pct"`
			} `json:"leagueRecord"`
			Records struct {
				SplitRecords []struct {
					Wins   int    `json:"wins"`
					Losses int    `json:"losses"`
					Type   string `json:"type"`
					Pct    string `json:"pct"`
				} `json:"splitRecords"`
			} `json:"records"`
			RunsAllowed int `json:"runsAllowed"`
			RunsScored  int `json:"runsScored"`
		} `json:"teamRecords"`
	} `json:"records"`
}

// allDivsions
type AllDivisionsJSON struct {
	Divisions []struct {
		ID        int    `json:"id"`
		NameShort string `json:"nameShort"`
		Sport     struct {
			ID int `json:"id"`
		} `json:"sport"`
	}
}

// Games
type AllGamesJSON struct {
	Dates []struct {
		Date  string `json:"date"`
		Games []struct {
			GamePk   int       `json:"gamePk"`
			GameDate time.Time `json:"gameDate"`
			Status   struct {
				DetailedState string `json:"detailedState"`
			} `json:"status"`
			Teams struct {
				Away struct {
					Team struct {
						ID int `json:"id"`
					} `json:"team"`
					Score int `json:"score"`
				} `json:"away"`
				Home struct {
					Team struct {
						ID int `json:"id"`
					} `json:"team"`
					Score int `json:"score"`
				} `json:"home"`
			} `json:"teams"`
		} `json:"games"`
	} `json:"dates"`
}

type LineScoreJSON struct {
	Innings []struct {
		OrdinalNum string `json:"ordinalNum"`
		Home       struct {
			Runs       int `json:"runs"`
			Hits       int `json:"hits"`
			Errors     int `json:"errors"`
			LeftOnBase int `json:"leftOnBase"`
		} `json:"home"`
		Away struct {
			Runs       int `json:"runs"`
			Hits       int `json:"hits"`
			Errors     int `json:"errors"`
			LeftOnBase int `json:"leftOnBase"`
		} `json:"away"`
	} `json:"innings"`
}

func timeToPgDate(time time.Time) pgtype.Date {
	date := pgtype.Date{
		Time:  time,
		Valid: true,
	}
	return date
}

func (client *SportClient) ResetResults() error {
	if err := client.Queries.ResetGames(context.Background()); err != nil {
		return fmt.Errorf("Couldn't reset games table: %w", err)
	}
	if err := client.Queries.ResetRecords(context.Background()); err != nil {
		return fmt.Errorf("Couldn't reset records table: %w", err)
	}
	return nil
}
