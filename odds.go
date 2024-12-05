package sportmonks

import (
	"context"
	"net/url"
	"strconv"
	"strings"
)

type PrematchOdds struct {
	ID                    int     `json:"id"`
	FixtureID             int     `json:"fixture_id"`
	MarketID              int     `json:"market_id"`
	BookmakerID           int     `json:"bookmaker_id"`
	Label                 string  `json:"label"`
	Value                 string  `json:"value"`
	Name                  string  `json:"name"`
	SortOrder             *int    `json:"sort_order"`
	MarketDescription     string  `json:"market_description"`
	Probability           string  `json:"probability"`
	Dp3                   string  `json:"dp3"`
	Fractional            string  `json:"fractional"`
	American              string  `json:"american"`
	Winning               bool    `json:"winning"`
	Stopped               bool    `json:"stopped"`
	Total                 string  `json:"total"`
	Handicap              string  `json:"handicap"`
	Participants          string  `json:"participants"`
	CreatedAt             string  `json:"created_at"`
	OriginalLabel         *string `json:"original_label"`
	LatestBookmakerUpdate string  `json:"latest_bookmaker_update"`
	Fixture               Fixture `json:"fixture"`
}

func (c *HTTPClient) AllPrematchOdds(ctx context.Context, includes []string, filters map[string][]int, page int) ([]PrematchOdds, *ResponseDetails, error) {
	path := prematchOddsURI

	return multipleOddsResponse(ctx, c, path, includes, filters, page)
}

func (c *HTTPClient) PrematchOddsByFixtureID(ctx context.Context, id int, includes []string, filters map[string][]int, page int) ([]PrematchOdds, *ResponseDetails, error) {
	path := prematchOddsURIByFixtureID + "/" + strconv.Itoa(id)

	return multipleOddsResponse(ctx, c, path, includes, filters, page)
}

func (c *HTTPClient) PrematchOddsByFixtureIDAndBookmakerID(ctx context.Context, fixtureID, bookmakerID int, includes []string, filters map[string][]int, page int) ([]PrematchOdds, *ResponseDetails, error) {
	path := prematchOddsURIByFixtureID + "/" + strconv.Itoa(fixtureID) + "/bookmakers/" + strconv.Itoa(bookmakerID)

	return multipleOddsResponse(ctx, c, path, includes, filters, page)
}

func (c *HTTPClient) PrematchOddsByFixtureIDAndMarketID(ctx context.Context, fixtureID, marketID int, includes []string, filters map[string][]int, page int) ([]PrematchOdds, *ResponseDetails, error) {
	path := prematchOddsURIByFixtureID + "/" + strconv.Itoa(fixtureID) + "/markets/" + strconv.Itoa(marketID)

	return multipleOddsResponse(ctx, c, path, includes, filters, page)
}

func (c *HTTPClient) LatestOdds(ctx context.Context, includes []string, filters map[string][]int) ([]PrematchOdds, *ResponseDetails, error) {
	path := lastUpdatedOddsURI

	return multipleOddsResponse(ctx, c, path, includes, filters, 0)
}

func multipleOddsResponse(ctx context.Context, client *HTTPClient, path string, includes []string, filters map[string][]int, page int) ([]PrematchOdds, *ResponseDetails, error) {

	var values url.Values

	if page == 0 {
		values = url.Values{
			"include": {strings.Join(includes, ";")},
		}
	} else {
		values = url.Values{
			"include": {strings.Join(includes, ";")},
			"page":    {strconv.Itoa(page)},
		}
	}

	formatFilters(&values, filters)

	response := struct {
		Data         []PrematchOdds `json:"data"`
		Pagination   *Pagination    `json:"pagination"`
		Subscription []Subscription `json:"subscription"`
		RateLimit    RateLimit      `json:"rate_limit"`
		TimeZone     string         `json:"timezone"`
	}{}

	err := client.getResource(ctx, path, values, &response)

	if err != nil {
		return nil, nil, err
	}

	return response.Data, &ResponseDetails{
		Subscription: response.Subscription,
		RateLimit:    response.RateLimit,
		TimeZone:     response.TimeZone,
	}, nil
}
