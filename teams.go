package sportmonks

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

// Team provides a struct representation of a Team resource.
type Team struct {
	ID                         int                        `json:"id"`
	LegacyID                   int                        `json:"legacy_id"`
	Name                       string                     `json:"name"`
	ShortCode                  string                     `json:"short_code"`
	Twitter                    *string                    `json:"twitter"`
	CountryID                  int                        `json:"country_id"`
	NationalTeam               bool                       `json:"national_team"`
	Founded                    int                        `json:"founded"`
	LogoPath                   *string                    `json:"logo_path"`
	VenueID                    int                        `json:"venue_id"`
	CurrentSeasonID            int                        `json:"current_season_id"`
	AggregatedAssistScorerData AggregatedAssistScorerData `json:"aggregatedAssistscorers,omitempty"`
	AggregatedCardScorerData   AggregatedCardScorerData   `json:"aggregatedCardscorers,omitempty"`
	AggregatedGoalScorerData   AggregatedGoalScorerData   `json:"aggregatedGoalscorers,omitempty"`
	AssistScorerData           AssistScorerData           `json:"assistscorers,omitempty"`
	CardScorerData             CardScorerData             `json:"cardscorers,omitempty"`
	CoachData                  CoachData                  `json:"coach,omitempty"`
	CountryData                CountryData                `json:"country,omitempty"`
	FIFARankingData            RankingData                `json:"fifaranking"`
	GoalScorerData             GoalScorerData             `json:"goalscorers,omitempty"`
	Latest                     FixturesData               `json:"latest,omitempty"`
	LeagueData                 LeagueData                 `json:"league,omitempty"`
	LocalFixtureData           FixturesData               `json:"localFixtures,omitempty"`
	SidelinedData              SidelinedData              `json:"sidelined,omitempty"`
	SquadData                  SquadPlayerData            `json:"squad,omitempty"`
	StatsData                  TeamSeasonStatsData        `json:"stats,omitempty"`
	TransfersData              TransfersData              `json:"transfers,omitempty"`
	UEFARankingData            RankingData                `json:"uefaranking. omitempty"`
	Upcoming                   FixturesData               `json:"upcoming,omitempty"`
	VenueData                  VenueData                  `json:"venue,omitempty"`
	VisitorFixtureData         FixturesData               `json:"visitorFixtures,omitempty"`
	VisitorResultData          FixturesData               `json:"visitorResults,omitempty"`
}

// AggregatedAssistScorers returns aggregated assist scorer data.
func (t *Team) AggregatedAssistScorers() []AssistScorer {
	return t.AggregatedAssistScorerData.Data
}

// AggregatedCardScorers returns aggregated card scorer data.
func (t *Team) AggregatedCardScorers() []CardScorer {
	return t.AggregatedCardScorerData.Data
}

// AggregatedGoalScorers returns aggregated goal scorer data.
func (t *Team) AggregatedGoalScorers() []CardScorer {
	return t.AggregatedCardScorerData.Data
}

// AssistScorers returns assist scorer data.
func (t *Team) AssistScorers() []AssistScorer {
	return t.AssistScorerData.Data
}

// CardScorers returns card scorer data.
func (t *Team) CardScorers() []CardScorer {
	return t.CardScorerData.Data
}

// Coach returns the coach data.
func (t *Team) Coach() *Coach {
	return t.CoachData.Data
}

// Country returns the country data.
func (t *Team) Country() *Country {
	return t.CountryData.Data
}

// FIFARanking returns the current FIFA ranking.
func (t *Team) FIFARanking() *Ranking {
	return t.FIFARankingData.Data
}

// GoalScorers returns goal scorer data.
func (t *Team) GoalScorers() []GoalScorer {
	return t.GoalScorerData.Data
}

// LatestResults returns the latest completed fixture information.
func (t *Team) LatestResults() []Fixture {
	return t.Latest.Data
}

// League returns the current league data.
func (t *Team) League() *League {
	return t.LeagueData.Data
}

// LocalFixtures returns all fixture data.
func (t *Team) LocalFixtures() []Fixture {
	return t.LocalFixtureData.Data
}

// Sidelined returns player data for injured players.
func (t *Team) Sidelined() []Sidelined {
	return t.SidelinedData.Data
}

// Squad returns current squad data.
func (t *Team) Squad() []SquadPlayer {
	return t.SquadData.Data
}

// Stats returns the aggregated team stats for the current season
func (t *Team) Stats() *TeamSeasonStats {
	return t.StatsData.Data
}

// Transfers returns the teams transfer activity
func (t *Team) Transfers() []Transfer {
	return t.TransfersData.Data
}

// UEFARanking returns the current UEFA ranking.
func (t *Team) UEFARanking() *Ranking {
	return t.UEFARankingData.Data
}

// UpcomingFixtures returns all upcoming fixture data.
func (t *Team) UpcomingFixtures() []Fixture {
	return t.Upcoming.Data
}

// Venue returns venue data.
func (t *Team) Venue() *Venue {
	return t.VenueData.Data
}

// VisitorFixtures returns all upcoming fixture data for the opposition team.
func (t *Team) VisitorFixtures() []Fixture {
	return t.VisitorFixtureData.Data
}

// VisitorResults returns all completed fixture data for the opposition team.
func (t *Team) VisitorResults() []Fixture {
	return t.VisitorResultData.Data
}

// TeamByID fetches a Team resource by ID. Use the includes slice of string to enrich the response data.
func (c *HTTPClient) TeamByID(ctx context.Context, id int, includes []string, filters map[string][]int) (*Team, *Meta, error) {
	path := fmt.Sprintf(teamsURI+"/%d", id)

	values := url.Values{
		"include": {strings.Join(includes, ",")},
	}

	formatFilters(&values, filters)

	response := struct {
		Data *Team `json:"data"`
		Meta *Meta `json:"meta"`
	}{}

	err := c.getResource(ctx, path, values, &response)

	if err != nil {
		return nil, nil, err
	}

	return response.Data, response.Meta, err
}

// TeamsBySeasonID fetches Team resources associated to a season ID. The endpoint used within this method is paginated,
// to select the required page use the 'page' method argument. Page information including current page and total page
// are included within the Meta response. Use the includes slice of string to enrich the response data.
func (c *HTTPClient) TeamsBySeasonID(ctx context.Context, seasonID, page int, includes []string) ([]Team, *Meta, error) {
	path := fmt.Sprintf(teamsSeasonURI+"/%d", seasonID)

	values := url.Values{
		"include": {strings.Join(includes, ",")},
		"page":    {strconv.Itoa(page)},
	}

	response := struct {
		Data []Team `json:"data"`
		Meta *Meta  `json:"meta"`
	}{}

	err := c.getResource(ctx, path, values, &response)

	if err != nil {
		return nil, nil, err
	}

	return response.Data, response.Meta, err
}
