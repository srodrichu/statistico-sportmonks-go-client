package sportmonks

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

var prematchOddsResponse = `{
	"data": [
		{
			"id": 1,
			"fixture_id": 11867289,
			"market_id": 2,
			"bookmaker_id": 3,
			"label": "1X2",
			"value": "2.5",
			"name": "Home Win",
			"sort_order": 1,
			"market_description": "Match Odds",
			"probability": "40%",
			"dp3": "2.500",
			"fractional": "3/2",
			"american": "+150",
			"winning": false,
			"stopped": false,
			"total": null,
			"handicap": null,
			"participants": "Team A vs Team B",
			"created_at": "2023-10-01T12:00:00Z",
			"original_label": null,
			"latest_bookmaker_update": "2023-10-01T12:00:00Z",
			"fixture": {
				"id": 11867289,
				"starting_at": "2023-10-01T12:00:00Z",
				"state": {
					"id": 1,
					"name": "finished"
					
					}
			}
		}
	],
	"pagination": {
		"total": 1,
		"count": 1,
		"per_page": 10,
		"current_page": 1,
		"total_pages": 1
	},
	"subscription": [],
	"rate_limit": {
		"limit": 100,
		"remaining": 99,
		"reset": 1633024800
	},
	"timezone": "UTC"
}`

func TestAllPrematchOdds(t *testing.T) {
	url := defaultBaseURL + "/football/odds/pre-match?api_token=api-key&include=&page=1"

	t.Run("returns prematch odds struct slice", func(t *testing.T) {
		server := mockResponseServer(t, prematchOddsResponse, 200, url)

		client := newTestHTTPClient(server)

		odds, _, err := client.AllPrematchOdds(context.Background(), []string{}, map[string][]int{}, 1)

		if err != nil {
			t.Fatalf("Test failed, expected nil, got %s", err.Error())
		}

		assertPrematchOdds(t, &odds[0])
	})

	t.Run("returns bad status code error", func(t *testing.T) {
		server := mockResponseServer(t, errorResponse, 400, url)

		client := newTestHTTPClient(server)

		odds, _, err := client.AllPrematchOdds(context.Background(), []string{}, map[string][]int{}, 1)

		if odds != nil {
			t.Fatalf("Test failed, expected nil, got %+v", odds)
		}

		assertError(t, err)
	})
}

func TestPrematchOddsByFixtureID(t *testing.T) {
	url := defaultBaseURL + "/football/odds/pre-match/fixtures/11867289?api_token=api-key&include=&page=1"

	t.Run("returns prematch odds struct slice", func(t *testing.T) {
		server := mockResponseServer(t, prematchOddsResponse, 200, url)

		client := newTestHTTPClient(server)

		odds, _, err := client.PrematchOddsByFixtureID(context.Background(), 11867289, []string{}, map[string][]int{}, 1)

		if err != nil {
			t.Fatalf("Test failed, expected nil, got %s", err.Error())
		}

		assertPrematchOdds(t, &odds[0])
	})

	t.Run("returns bad status code error", func(t *testing.T) {
		server := mockResponseServer(t, errorResponse, 400, url)

		client := newTestHTTPClient(server)

		odds, _, err := client.PrematchOddsByFixtureID(context.Background(), 11867289, []string{}, map[string][]int{}, 1)

		if odds != nil {
			t.Fatalf("Test failed, expected nil, got %+v", odds)
		}

		assertError(t, err)
	})
}

func assertPrematchOdds(t *testing.T, odds *PrematchOdds) {
	assert.Equal(t, 1, odds.ID)
	assert.Equal(t, 11867289, odds.FixtureID)
	assert.Equal(t, 2, odds.MarketID)
	assert.Equal(t, 3, odds.BookmakerID)
	assert.Equal(t, "1X2", odds.Label)
	assert.Equal(t, "2.5", odds.Value)
	assert.Equal(t, "Home Win", odds.Name)
	assert.Equal(t, 1, *odds.SortOrder)
	assert.Equal(t, "Match Odds", odds.MarketDescription)
	assert.Equal(t, "40%", odds.Probability)
	assert.Equal(t, "2.500", odds.Dp3)
	assert.Equal(t, "3/2", odds.Fractional)
	assert.Equal(t, "+150", odds.American)
	assert.Equal(t, false, odds.Winning)
	assert.Equal(t, false, odds.Stopped)
	assert.Nil(t, odds.Total)
	assert.Nil(t, odds.Handicap)
	assert.Equal(t, "Team A vs Team B", odds.Participants)
	assert.Equal(t, "2023-10-01T12:00:00Z", odds.CreatedAt)
	assert.Nil(t, odds.OriginalLabel)
	assert.Equal(t, "2023-10-01T12:00:00Z", odds.LatestBookmakerUpdate)
	assert.Equal(t, 11867289, odds.Fixture.ID)
	assert.Equal(t, "2023-10-01T12:00:00Z", odds.Fixture.StartingAt)
	assert.Equal(t, "finished", odds.Fixture.FixtureState.Name)
}
