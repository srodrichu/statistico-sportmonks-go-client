package sportmonks

import (
	"context"
	"net/url"
	"strconv"
	"strings"
)

// VideoHighlights provides a struct representation of a VideoHighlights resource.
type VideoHighlights struct {
	FixtureID   int         `json:"fixture_id"`
	EventID     *int        `json:"event_id"`
	Location    string      `json:"location"`
	Type        string      `json:"type"`
	CreatedAt   DateTime    `json:"created_at"`
	FixtureData FixtureData `json:"fixture,omitempty"`
}

// Fixture returns associated fixture data.
func (v *VideoHighlights) Fixture() *Fixture {
	return v.FixtureData.Data
}

// VideoHighlights fetches VideoHighlight resources. The endpoint used within this method is paginated,
// to select the required page use the 'page' method argument. Page information including current page and total
// page are included within the Meta response. Use the includes slice to enrich the response data.
func (c *HTTPClient) VideoHighlights(ctx context.Context, page int, includes []string) ([]VideoHighlights, *Meta, error) {
	values := url.Values{
		"page":    {strconv.Itoa(page)},
		"include": {strings.Join(includes, ",")},
	}

	response := struct {
		Data []VideoHighlights `json:"data"`
		Meta *Meta             `json:"meta"`
	}{}

	err := c.getResource(ctx, videoHighlightsURI, values, &response)

	if err != nil {
		return nil, nil, err
	}

	return response.Data, response.Meta, err
}
