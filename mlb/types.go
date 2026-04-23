package mlb

import (
	"context"
	"encoding/json"
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
	Client    *http.Client
	Queries   *sqlc.Queries
	GamesSeen map[int]bool
}

func NewSportClient(db sqlc.DBTX) *SportClient {
	client := &http.Client{Timeout: 30 * time.Second}
	return &SportClient{Client: client, Queries: sqlc.New(db), GamesSeen: make(map[int]bool)}
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

type AllPlayersJSON struct {
	People []Player
}

type Player struct {
	ID          int    `json:"id"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	CurrentAge  int    `json:"currentAge"`
	Active      bool   `json:"active"`
	CurrentTeam struct {
		ID int `json:"id"`
	} `json:"currentTeam"`
	PrimaryPosition struct {
		Abbreviation string `json:"abbreviation"`
	} `json:"primaryPosition"`
	BatSide struct {
		Code string `json:"code"`
	} `json:"batSide"`
	PitchHand struct {
		Code string `json:"code"`
	} `json:"pitchHand"`
}

type PitchingStatsJSON struct {
	Stats []struct {
		Group struct {
			DisplayName string `json:"displayName"`
		} `json:"group"`
		Splits []PitchingSplit `json:"splits"`
	} `json:"stats"`
}

type PitchingSplit struct {
	Season string `json:"season"`
	Game   struct {
		GamePk int `json:"gamePk"`
	} `json:"game"`
	Stat PitchingStat `json:"stat"`
}

type PitchingStat struct {
	GamesPlayed        int    `json:"gamesPlayed"`
	GamesStarted       int    `json:"gamesStarted"`
	InningsPitched     string `json:"inningsPitched"`
	Wins               int    `json:"wins"`
	Losses             int    `json:"losses"`
	Saves              int    `json:"saves"`
	Holds              int    `json:"holds"`
	BlownSaves         int    `json:"blownSaves"`
	Era                string `json:"era"`
	Whip               string `json:"whip"`
	Strikeouts         int    `json:"strikeouts"`
	Walks              int    `json:"baseOnBalls"`
	Hits               int    `json:"hits"`
	HomeRuns           int    `json:"homeRuns"`
	EarnedRuns         int    `json:"earnedRuns"`
	HitBatsmen         int    `json:"hitBatsmen"`
	StrikeoutsPer9     string `json:"strikeoutsPer9Inn"`
	WalksPer9          string `json:"walksPer9Inn"`
	HitsPer9           string `json:"hitsPer9Inn"`
	StrikeoutWalkRatio string `json:"strikeoutWalkRatio"`
}

type BattingStatsJSON struct {
	Stats []struct {
		Group struct {
			DisplayName string `json:"displayName"` // "hitting"
		} `json:"group"`
		Splits []BattingSplit `json:"splits"`
	} `json:"stats"`
}

type BattingSplit struct {
	Game struct {
		GamePk int `json:"gamePk"`
	} `json:"game"`
	Stat BattingStat `json:"stat"`
}

type BattingStat struct {
	AtBats         int    `json:"atBats"`
	Runs           int    `json:"runs"`
	Hits           int    `json:"hits"`
	Doubles        int    `json:"doubles"`
	Triples        int    `json:"triples"`
	HomeRuns       int    `json:"homeRuns"`
	RBI            int    `json:"rbi"`
	StolenBases    int    `json:"stolenBases"`
	CaughtStealing int    `json:"caughtStealing"`
	Walks          int    `json:"baseOnBalls"`
	Strikeouts     int    `json:"strikeOuts"`
	HitByPitch     int    `json:"hitByPitch"`
	Avg            string `json:"avg"`
	OBP            string `json:"obp"`
	Slugging       string `json:"slg"`
	OPS            string `json:"ops"`
	LeftOnBase     int    `json:"leftOnBase"`
	SacBunts       int    `json:"sacBunts"`
	SacFlies       int    `json:"sacFlies"`
}

func (client *SportClient) ResetResults() error {
	if err := client.Queries.ResetDivisions(context.Background()); err != nil {
		return fmt.Errorf("Couldn't reset divisions table: %w", err)
	}
	if err := client.Queries.ResetTeams(context.Background()); err != nil {
		return fmt.Errorf("Couldn't reset teams table: %w", err)
	}
	if err := client.Queries.ResetGames(context.Background()); err != nil {
		return fmt.Errorf("Couldn't reset games table: %w", err)
	}
	if err := client.Queries.ResetRecords(context.Background()); err != nil {
		return fmt.Errorf("Couldn't reset records table: %w", err)
	}
	if err := client.Queries.ResetInnings(context.Background()); err != nil {
		return fmt.Errorf("Couldn't reset innings table: %w", err)
	}
	if err := client.Queries.ResetPlayers(context.Background()); err != nil {
		return fmt.Errorf("Couldn't reset players table: %w", err)
	}
	if err := client.Queries.ResetPitching(context.Background()); err != nil {
		return fmt.Errorf("Couldn't reset pitching table: %w", err)
	}
	if err := client.Queries.ResetBatting(context.Background()); err != nil {
		return fmt.Errorf("Couldn't reset batting table: %w", err)
	}
	return nil
}

func (client *SportClient) UpdateResults() error {
	if err := client.GetAllDivisions(); err != nil {
		return err
	}
	time.Sleep(1 * time.Second)
	if err := client.GetAllTeams(); err != nil {
		return err
	}
	time.Sleep(1 * time.Second)
	if err := client.GetGameResults(); err != nil {
		return err
	}
	time.Sleep(1 * time.Second)
	if err := client.GetTeamRecords(); err != nil {
		return err
	}
	time.Sleep(1 * time.Second)
	if err := client.GetAllPlayers(); err != nil {
		return err
	}
	return nil
}

func (client *SportClient) RequestAndDecode(endpoint string, target any) error {
	url := buildURL(endpoint)

	const maxAttempts = 4
	for attempt := range maxAttempts {
		res, err := client.Client.Get(url)
		if err == nil {
			defer res.Body.Close()
			if err := json.NewDecoder(res.Body).Decode(target); err != nil {
				return fmt.Errorf("couldn't decode JSON into %T: %w", target, err)
			}
			return nil
		}

		if attempt == maxAttempts-1 {
			return fmt.Errorf("Request to %s failed after %d attempts: %w", url, maxAttempts, err)
		}
		wait := time.Duration(1<<attempt) * time.Second // 1s, 2s, 4s
		time.Sleep(wait)
	}
	return nil
}
