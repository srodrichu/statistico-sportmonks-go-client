package sportmonks

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

var leaguesResponse = `{
	"data": [
		{
		  "id": 8,
		  "sport_id": 1,
		  "country_id": 462,
		  "name": "Premier League",
		  "active": true,
		  "short_code": "UK PL",
		  "image_path": "https://cdn.sportmonks.com/images/soccer/leagues/8/8.png",
		  "type": "league",
		  "sub_type": "domestic",
		  "last_played_at": "2024-09-22 15:30:00",
		  "category": 1,
		  "has_jerseys": false
		}
   ],
	"pagination": {
		"count": 2,
		"per_page": 25,
		"current_page": 1,
		"next_page": null,
		"has_more": false
	},
	"subscription": [
		{
			"meta": {
				"trial_ends_at": null,
				"ends_at": "2024-10-26 12:06:34",
				"current_timestamp": 1728372666
			},
			"plans": [
				{
					"plan": "Joe Sweeny Custom Plan",
					"sport": "Football",
					"category": "Advanced"
				}
			],
			"add_ons": [],
			"widgets": []
		}
	],
	"rate_limit": {
		"resets_in_seconds": 3386,
		"remaining": 2997,
		"requested_entity": "League"
	},
	"timezone": "UTC"
}`

var leaguesIncludesResponse = `{
	"data": [
		{
		  "id": 8,
		  "sport_id": 1,
		  "country_id": 462,
		  "name": "Premier League",
		  "active": true,
		  "short_code": "UK PL",
		  "image_path": "https://cdn.sportmonks.com/images/soccer/leagues/8/8.png",
		  "type": "league",
		  "sub_type": "domestic",
		  "last_played_at": "2024-09-22 15:30:00",
		  "category": 1,
		  "has_jerseys": false
		}
   ]
}`

var leagueResponse = `{
	"data": {
      "id": 8,
      "sport_id": 1,
      "country_id": 462,
      "name": "Premier League",
      "active": true,
      "short_code": "UK PL",
      "image_path": "https://cdn.sportmonks.com/images/soccer/leagues/8/8.png",
      "type": "league",
      "sub_type": "domestic",
      "last_played_at": "2024-09-22 15:30:00",
      "category": 1,
      "has_jerseys": false
    },
	"pagination": {
		"count": 2,
		"per_page": 25,
		"current_page": 1,
		"next_page": null,
		"has_more": false
	},
	"subscription": [
		{
			"meta": {
				"trial_ends_at": null,
				"ends_at": "2024-10-26 12:06:34",
				"current_timestamp": 1728372666
			},
			"plans": [
				{
					"plan": "Joe Sweeny Custom Plan",
					"sport": "Football",
					"category": "Advanced"
				}
			],
			"add_ons": [],
			"widgets": []
		}
	],
	"rate_limit": {
		"resets_in_seconds": 3386,
		"remaining": 2997,
		"requested_entity": "League"
	},
	"timezone": "UTC"
}`

var leagueIncludesResponse = `{
	"data": {
      "id": 8,
      "sport_id": 1,
      "country_id": 462,
      "name": "Premier League",
      "active": true,
      "short_code": "UK PL",
      "image_path": "https://cdn.sportmonks.com/images/soccer/leagues/8/8.png",
      "type": "league",
      "sub_type": "domestic",
      "last_played_at": "2024-09-22 15:30:00",
      "category": 1,
      "has_jerseys": false
    }
}`

func TestLeagues(t *testing.T) {
	t.Run("returns slice of League struct", func(t *testing.T) {
		url := defaultBaseURL + "/football/leagues?api_token=api-key&include=&page=1"

		server := mockResponseServer(t, leaguesResponse, 200, url)

		client := newTestHTTPClient(server)

		leagues, _, err := client.Leagues(context.Background(), 1, []string{})

		if err != nil {
			t.Fatalf("Test failed, expected nil, got %s", err.Error())
		}

		assertLeague(t, &leagues[0])
	})

	t.Run("returns slice of League struct with includes data", func(t *testing.T) {
		url := defaultBaseURL + "/football/leagues?api_token=api-key&include=country%3Bseason%3Bseasons&page=1"

		server := mockResponseServer(t, leaguesIncludesResponse, 200, url)

		client := newTestHTTPClient(server)

		leagues, _, err := client.Leagues(context.Background(), 1, []string{"country", "season", "seasons"})

		if err != nil {
			t.Fatalf("Test failed, expected nil, got %s", err.Error())
		}

		assertLeague(t, &leagues[0])
	})

	t.Run("returns bad status code error", func(t *testing.T) {
		url := defaultBaseURL + "/football/leagues?api_token=api-key&include=&page=1"

		server := mockResponseServer(t, errorResponse, 400, url)

		client := newTestHTTPClient(server)

		leagues, _, err := client.Leagues(context.Background(), 1, []string{})

		if leagues != nil {
			t.Fatalf("Test failed, expected nil, got %+v", leagues)
		}

		assertError(t, err)
	})

	t.Run("can handle response details", func(t *testing.T) {
		t.Helper()

		url := defaultBaseURL + "/football/leagues?api_token=api-key&include=country%3Bseason%3Bseasons&page=1"

		server := mockResponseServer(t, leaguesResponse, 200, url)

		client := newTestHTTPClient(server)

		_, details, _ := client.Leagues(context.Background(), 1, []string{"country", "season", "seasons"})

		assertResponseDetails(t, details, "League")
	})
}

func TestLeagueByID(t *testing.T) {
	t.Run("returns a single League struct", func(t *testing.T) {
		url := defaultBaseURL + "/football/leagues/82?api_token=api-key&include="

		server := mockResponseServer(t, leagueResponse, 200, url)

		client := newTestHTTPClient(server)

		league, _, err := client.LeagueByID(context.Background(), 82, []string{})

		if err != nil {
			t.Fatalf("Test failed, expected nil, got %s", err.Error())
		}

		assertLeague(t, league)
	})

	t.Run("returns a League struct with includes data", func(t *testing.T) {
		url := defaultBaseURL + "/football/leagues/82?api_token=api-key&include=country%3Bseason%3Bseasons"

		server := mockResponseServer(t, leagueIncludesResponse, 200, url)

		client := newTestHTTPClient(server)

		league, _, err := client.LeagueByID(context.Background(), 82, []string{"country", "season", "seasons"})

		if err != nil {
			t.Fatalf("Test failed, expected nil, got %s", err.Error())
		}

		assertLeague(t, league)
	})

	t.Run("returns bad status code error", func(t *testing.T) {
		url := defaultBaseURL + "/football/leagues/82?api_token=api-key&include="

		server := mockResponseServer(t, errorResponse, 400, url)

		client := newTestHTTPClient(server)

		league, _, err := client.LeagueByID(context.Background(), 82, []string{})

		if league != nil {
			t.Fatalf("Test failed, expected nil, got %+v", league)
		}

		assertError(t, err)
	})

	t.Run("can handle response details response", func(t *testing.T) {
		url := defaultBaseURL + "/football/leagues/82?api_token=api-key&include="

		server := mockResponseServer(t, leagueResponse, 200, url)

		client := newTestHTTPClient(server)

		_, details, _ := client.LeagueByID(context.Background(), 82, []string{})

		assertResponseDetails(t, details, "League")
	})
}

func assertLeague(t *testing.T, league *League) {
	assert.Equal(t, 8, league.ID)
	assert.Equal(t, 1, league.SportID)
	assert.Equal(t, 462, league.CountryID)
	assert.Equal(t, "Premier League", league.Name)
	assert.Equal(t, true, league.Active)
	assert.Equal(t, "UK PL", league.ShortCode)
	assert.Equal(t, "https://cdn.sportmonks.com/images/soccer/leagues/8/8.png", league.ImagePath)
	assert.Equal(t, "league", league.Type)
	assert.Equal(t, "domestic", league.SubType)
	assert.Equal(t, "2024-09-22 15:30:00", league.LastPlayedAt)
	assert.Equal(t, 1, league.Category)
	assert.Equal(t, false, league.HasJerseys)
}

func assertSubscription(t *testing.T, sub Subscription) {
	assert.Nil(t, sub.Meta[0].TrialEndsAt)
	assert.Equal(t, "2024-10-26 12:06:34", sub.Meta[0].EndsAt)
	assert.Equal(t, int64(1728317818), sub.Meta[0].CurrentTimestamp)

	assert.Equal(t, 1, len(sub.Plans))
	assert.Equal(t, "Joe Sweeny Custom Plan", sub.Plans[0].Plan)
	assert.Equal(t, "Football", sub.Plans[0].Sport)
	assert.Equal(t, "Advanced", sub.Plans[0].Category)
}
