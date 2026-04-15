package mlb

import (
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
		SpringLeague struct {
			ID           int    `json:"id"`
			Name         string `json:"name"`
			Link         string `json:"link"`
			Abbreviation string `json:"abbreviation"`
		} `json:"springLeague"`
		AllStarStatus string `json:"allStarStatus"`
		ID            int    `json:"id"`
		Name          string `json:"name"`
		Link          string `json:"link"`
		Season        int    `json:"season"`
		Venue         struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
			Link string `json:"link"`
		} `json:"venue"`
		SpringVenue struct {
			ID   int    `json:"id"`
			Link string `json:"link"`
		} `json:"springVenue"`
		TeamCode        string `json:"teamCode"`
		FileCode        string `json:"fileCode"`
		Abbreviation    string `json:"abbreviation"`
		TeamName        string `json:"teamName"`
		LocationName    string `json:"locationName"`
		FirstYearOfPlay string `json:"firstYearOfPlay"`
		League          struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
			Link string `json:"link"`
		} `json:"league"`
		Division struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
			Link string `json:"link"`
		} `json:"division"`
		Sport struct {
			ID   int    `json:"id"`
			Link string `json:"link"`
			Name string `json:"name"`
		} `json:"sport"`
		ShortName     string `json:"shortName"`
		FranchiseName string `json:"franchiseName"`
		ClubName      string `json:"clubName"`
		Active        bool   `json:"active"`
	} `json:"teams"`
}

// Records
type TeamRecordsJSON struct {
	Copyright string `json:"copyright"`
	Records   []struct {
		StandingsType string `json:"standingsType"`
		League        struct {
			ID   int    `json:"id"`
			Link string `json:"link"`
		} `json:"league"`
		Division struct {
			ID   int    `json:"id"`
			Link string `json:"link"`
		} `json:"division"`
		Sport struct {
			ID   int    `json:"id"`
			Link string `json:"link"`
		} `json:"sport"`
		RoundRobin struct {
			Status string `json:"status"`
		} `json:"roundRobin"`
		LastUpdated time.Time `json:"lastUpdated"`
		TeamRecords []struct {
			Team struct {
				ID   int    `json:"id"`
				Name string `json:"name"`
				Link string `json:"link"`
			} `json:"team"`
			Season string `json:"season"`
			Streak struct {
				StreakCode   string `json:"streakCode"`
				StreakType   string `json:"streakType"`
				StreakNumber int    `json:"streakNumber"`
			} `json:"streak"`
			DivisionRank          string `json:"divisionRank"`
			LeagueRank            string `json:"leagueRank"`
			SportRank             string `json:"sportRank"`
			GamesPlayed           int    `json:"gamesPlayed"`
			GamesBack             string `json:"gamesBack"`
			WildCardGamesBack     string `json:"wildCardGamesBack"`
			LeagueGamesBack       string `json:"leagueGamesBack"`
			SpringLeagueGamesBack string `json:"springLeagueGamesBack"`
			SportGamesBack        string `json:"sportGamesBack"`
			DivisionGamesBack     string `json:"divisionGamesBack"`
			ConferenceGamesBack   string `json:"conferenceGamesBack"`
			LeagueRecord          struct {
				Wins   int    `json:"wins"`
				Losses int    `json:"losses"`
				Ties   int    `json:"ties"`
				Pct    string `json:"pct"`
			} `json:"leagueRecord"`
			LastUpdated time.Time `json:"lastUpdated"`
			Records     struct {
				SplitRecords []struct {
					Wins   int    `json:"wins"`
					Losses int    `json:"losses"`
					Type   string `json:"type"`
					Pct    string `json:"pct"`
				} `json:"splitRecords"`
				DivisionRecords []struct {
					Wins     int    `json:"wins"`
					Losses   int    `json:"losses"`
					Pct      string `json:"pct"`
					Division struct {
						ID   int    `json:"id"`
						Name string `json:"name"`
						Link string `json:"link"`
					} `json:"division"`
				} `json:"divisionRecords"`
				OverallRecords []struct {
					Wins   int    `json:"wins"`
					Losses int    `json:"losses"`
					Type   string `json:"type"`
					Pct    string `json:"pct"`
				} `json:"overallRecords"`
				LeagueRecords []struct {
					Wins   int    `json:"wins"`
					Losses int    `json:"losses"`
					Pct    string `json:"pct"`
					League struct {
						ID   int    `json:"id"`
						Name string `json:"name"`
						Link string `json:"link"`
					} `json:"league"`
				} `json:"leagueRecords"`
				ExpectedRecords []struct {
					Wins   int    `json:"wins"`
					Losses int    `json:"losses"`
					Type   string `json:"type"`
					Pct    string `json:"pct"`
				} `json:"expectedRecords"`
			} `json:"records"`
			RunsAllowed                 int    `json:"runsAllowed"`
			RunsScored                  int    `json:"runsScored"`
			DivisionChamp               bool   `json:"divisionChamp"`
			DivisionLeader              bool   `json:"divisionLeader"`
			HasWildcard                 bool   `json:"hasWildcard"`
			Clinched                    bool   `json:"clinched"`
			EliminationNumber           string `json:"eliminationNumber"`
			EliminationNumberSport      string `json:"eliminationNumberSport"`
			EliminationNumberLeague     string `json:"eliminationNumberLeague"`
			EliminationNumberDivision   string `json:"eliminationNumberDivision"`
			EliminationNumberConference string `json:"eliminationNumberConference"`
			WildCardEliminationNumber   string `json:"wildCardEliminationNumber"`
			Wins                        int    `json:"wins"`
			Losses                      int    `json:"losses"`
			RunDifferential             int    `json:"runDifferential"`
			WinningPercentage           string `json:"winningPercentage"`
			WildCardRank                string `json:"wildCardRank,omitempty"`
			WildCardLeader              bool   `json:"wildCardLeader,omitempty"`
		} `json:"teamRecords"`
	} `json:"records"`
}

// allDivsions
type AllDivisionsJSON struct {
	Copyright string `json:"copyright"`
	Divisions []struct {
		ID           int    `json:"id"`
		Name         string `json:"name"`
		Season       string `json:"season"`
		NameShort    string `json:"nameShort"`
		Link         string `json:"link"`
		Abbreviation string `json:"abbreviation"`
		League       struct {
			ID   int    `json:"id"`
			Link string `json:"link"`
		} `json:"league"`
		Sport struct {
			ID   int    `json:"id"`
			Link string `json:"link"`
		} `json:"sport"`
		HasWildcard     bool `json:"hasWildcard"`
		SortOrder       int  `json:"sortOrder"`
		Active          bool `json:"active"`
		NumPlayoffTeams int  `json:"numPlayoffTeams,omitempty"`
	} `json:"divisions"`
}

// Games
type AllGamesJSON struct {
	Copyright            string `json:"copyright"`
	TotalItems           int    `json:"totalItems"`
	TotalEvents          int    `json:"totalEvents"`
	TotalGames           int    `json:"totalGames"`
	TotalGamesInProgress int    `json:"totalGamesInProgress"`
	Dates                []struct {
		Date                 string `json:"date"`
		TotalItems           int    `json:"totalItems"`
		TotalEvents          int    `json:"totalEvents"`
		TotalGames           int    `json:"totalGames"`
		TotalGamesInProgress int    `json:"totalGamesInProgress"`
		Games                []struct {
			GamePk       int       `json:"gamePk"`
			GameGUID     string    `json:"gameGuid"`
			Link         string    `json:"link"`
			GameType     string    `json:"gameType"`
			Season       string    `json:"season"`
			GameDate     time.Time `json:"gameDate"`
			OfficialDate string    `json:"officialDate"`
			Status       struct {
				AbstractGameState string `json:"abstractGameState"`
				CodedGameState    string `json:"codedGameState"`
				DetailedState     string `json:"detailedState"`
				StatusCode        string `json:"statusCode"`
				StartTimeTBD      bool   `json:"startTimeTBD"`
				AbstractGameCode  string `json:"abstractGameCode"`
			} `json:"status"`
			Teams struct {
				Away struct {
					Team struct {
						ID   int    `json:"id"`
						Name string `json:"name"`
						Link string `json:"link"`
					} `json:"team"`
					LeagueRecord struct {
						Wins   int    `json:"wins"`
						Losses int    `json:"losses"`
						Pct    string `json:"pct"`
					} `json:"leagueRecord"`
					Score        int  `json:"score"`
					IsWinner     bool `json:"isWinner"`
					SplitSquad   bool `json:"splitSquad"`
					SeriesNumber int  `json:"seriesNumber"`
				} `json:"away"`
				Home struct {
					Team struct {
						ID   int    `json:"id"`
						Name string `json:"name"`
						Link string `json:"link"`
					} `json:"team"`
					LeagueRecord struct {
						Wins   int    `json:"wins"`
						Losses int    `json:"losses"`
						Pct    string `json:"pct"`
					} `json:"leagueRecord"`
					Score        int  `json:"score"`
					IsWinner     bool `json:"isWinner"`
					SplitSquad   bool `json:"splitSquad"`
					SeriesNumber int  `json:"seriesNumber"`
				} `json:"home"`
			} `json:"teams"`
			Venue struct {
				ID   int    `json:"id"`
				Name string `json:"name"`
				Link string `json:"link"`
			} `json:"venue"`
			Content struct {
				Link string `json:"link"`
			} `json:"content"`
			IsTie                  bool   `json:"isTie,omitempty"`
			GameNumber             int    `json:"gameNumber"`
			PublicFacing           bool   `json:"publicFacing"`
			DoubleHeader           string `json:"doubleHeader"`
			GamedayType            string `json:"gamedayType"`
			Tiebreaker             string `json:"tiebreaker"`
			CalendarEventID        string `json:"calendarEventID"`
			SeasonDisplay          string `json:"seasonDisplay"`
			DayNight               string `json:"dayNight"`
			ScheduledInnings       int    `json:"scheduledInnings"`
			ReverseHomeAwayStatus  bool   `json:"reverseHomeAwayStatus"`
			InningBreakLength      int    `json:"inningBreakLength"`
			GamesInSeries          int    `json:"gamesInSeries"`
			SeriesGameNumber       int    `json:"seriesGameNumber"`
			SeriesDescription      string `json:"seriesDescription"`
			RecordSource           string `json:"recordSource"`
			IfNecessary            string `json:"ifNecessary"`
			IfNecessaryDescription string `json:"ifNecessaryDescription"`
		} `json:"games"`
		Events []any `json:"events"`
	} `json:"dates"`
}

type LineScoreJSON struct {
	Copyright            string `json:"copyright"`
	CurrentInning        int    `json:"currentInning"`
	CurrentInningOrdinal string `json:"currentInningOrdinal"`
	InningState          string `json:"inningState"`
	InningHalf           string `json:"inningHalf"`
	IsTopInning          bool   `json:"isTopInning"`
	ScheduledInnings     int    `json:"scheduledInnings"`
	Innings              []struct {
		Num        int    `json:"num"`
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
	Teams struct {
		Home struct {
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
	} `json:"teams"`
	Defense struct {
		Pitcher struct {
			ID       int    `json:"id"`
			FullName string `json:"fullName"`
			Link     string `json:"link"`
		} `json:"pitcher"`
		Catcher struct {
			ID       int    `json:"id"`
			FullName string `json:"fullName"`
			Link     string `json:"link"`
		} `json:"catcher"`
		First struct {
			ID       int    `json:"id"`
			FullName string `json:"fullName"`
			Link     string `json:"link"`
		} `json:"first"`
		Second struct {
			ID       int    `json:"id"`
			FullName string `json:"fullName"`
			Link     string `json:"link"`
		} `json:"second"`
		Third struct {
			ID       int    `json:"id"`
			FullName string `json:"fullName"`
			Link     string `json:"link"`
		} `json:"third"`
		Shortstop struct {
			ID       int    `json:"id"`
			FullName string `json:"fullName"`
			Link     string `json:"link"`
		} `json:"shortstop"`
		Left struct {
			ID       int    `json:"id"`
			FullName string `json:"fullName"`
			Link     string `json:"link"`
		} `json:"left"`
		Center struct {
			ID       int    `json:"id"`
			FullName string `json:"fullName"`
			Link     string `json:"link"`
		} `json:"center"`
		Right struct {
			ID       int    `json:"id"`
			FullName string `json:"fullName"`
			Link     string `json:"link"`
		} `json:"right"`
		Batter struct {
			ID       int    `json:"id"`
			FullName string `json:"fullName"`
			Link     string `json:"link"`
		} `json:"batter"`
		OnDeck struct {
			ID       int    `json:"id"`
			FullName string `json:"fullName"`
			Link     string `json:"link"`
		} `json:"onDeck"`
		InHole struct {
			ID       int    `json:"id"`
			FullName string `json:"fullName"`
			Link     string `json:"link"`
		} `json:"inHole"`
		BattingOrder int `json:"battingOrder"`
		Team         struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
			Link string `json:"link"`
		} `json:"team"`
	} `json:"defense"`
	Offense struct {
		Batter struct {
			ID       int    `json:"id"`
			FullName string `json:"fullName"`
			Link     string `json:"link"`
		} `json:"batter"`
		OnDeck struct {
			ID       int    `json:"id"`
			FullName string `json:"fullName"`
			Link     string `json:"link"`
		} `json:"onDeck"`
		InHole struct {
			ID       int    `json:"id"`
			FullName string `json:"fullName"`
			Link     string `json:"link"`
		} `json:"inHole"`
		Pitcher struct {
			ID       int    `json:"id"`
			FullName string `json:"fullName"`
			Link     string `json:"link"`
		} `json:"pitcher"`
		BattingOrder int `json:"battingOrder"`
		Team         struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
			Link string `json:"link"`
		} `json:"team"`
	} `json:"offense"`
	Balls   int `json:"balls"`
	Strikes int `json:"strikes"`
	Outs    int `json:"outs"`
}

func timeToPgDate(time time.Time) pgtype.Date {
	date := pgtype.Date{
		Time:  time,
		Valid: true,
	}
	return date
}
