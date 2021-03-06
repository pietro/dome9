package dome9

import (
	"context"
	"fmt"
	"net/http"
)

const assessmentHistoriesBasePath = "v2/AssessmentHistoryV2"

// AssessmentHistoriesService resource has methods to retrieve lists of previous
// compliance assessments, and specific assessment results. Assessments
// histories can be retrieved for a specified period of time (from, to), and can
// be filtered by assessment name. Each result has an id, which is used to then
// retrieve the specific assessment results.
// See: https://api-v2-docs.dome9.com/#Dome9-API-AssessmentHistoryV2
type AssessmentHistoriesService interface {
	GetBundleResults(context.Context, string, string, string, string, string) ([]AssessmentHistoryResult, *http.Response, error)
	GetAssessmentResult(context.Context, string) (*AssessmentHistoryResult, *http.Response, error)
	DeleteAssessmentResult(context.Context, string) (*http.Response, error)
}

// AssessmentHistoriesServiceOp handles communication with the AssessmentHistories
// related methods of the Dome9 API.
type AssessmentHistoriesServiceOp struct {
	client *Client
}

var _ AssessmentHistoriesService = &AssessmentHistoriesServiceOp{}

// AssessmentHistoryResult
type AssessmentHistoryResult struct {
	TriggeredBy      string                        `json:"triggeredBy"`
	Tests            []RuleTestResult              `json:"tests"`
	TestEntities     interface{}                   `json:"testEntities"`
	CreatedTime      string                        `json:"createdTime"`
	ID               int64                         `json:"id"`
	AssessmentPassed bool                          `json:"assessmentPassed"`
	HasErrors        bool                          `json:"hasErrors"`
	Stats            AssessmentHistoryStats        `json:"stats"`
	Request          AssessmentHistoryBundleResult `json:"request"`
}

// AssessmentHistoryStats
type AssessmentHistoryStats struct {
	Passed          int32 `json:"passed"`
	Failed          int32 `json:"failed"`
	Error           int32 `json:"error"`
	FailedTests     int32 `json:"failedTests"`
	LogicallyTested int32 `json:"logicallyTested"`
	FailedEntities  int32 `json:"failedEntities"`
}

// AssessmentHistoryBundleResult
type AssessmentHistoryBundleResult struct {
	Template               bool                 `json:"isTemplate"`
	ID                     int64                `json:"id"`
	Name                   string               `json:"name"`
	Description            string               `json:"description"`
	CFT                    AssessmentCFTRequest `json:"cft"`
	IsCFT                  bool                 `json:"isCft"`
	Dome9CloudAccountID    string               `json:"dome9CloudAccountId"`
	ExternalCloudAccountID string               `json:"externalCloudAccountId"`
	CloudAccountID         string               `json:"cloudAccountId"`
	Region                 string               `json:"region"`
	CloudNetwork           string               `json:"cloudNetwork"`
	CloudAccountType       string               `json:"cloudAccountType"`
	RequestID              string               `json:"requestId"`
}

// GetBundleResults
func (s *AssessmentHistoriesServiceOp) GetBundleResults(ctx context.Context, bundleID, cloudAccountIDs, fromTime, epsilonInMinutes, requestID string) ([]AssessmentHistoryResult, *http.Response, error) {
	path := fmt.Sprintf("%s?bundleId=%s&cloudAccountIds=%s&fromTime=%s&epsilonInMinutes=%s&requestId=%s", assessmentHistoriesBasePath, bundleID, cloudAccountIDs, fromTime, epsilonInMinutes, requestID)

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	var assessmentHistories []AssessmentHistoryResult
	resp, err := s.client.Do(ctx, req, &assessmentHistories)
	if err != nil {
		return nil, resp, err
	}

	return assessmentHistories, resp, err
}

// GetAssessmentResult
func (s *AssessmentHistoriesServiceOp) GetAssessmentResult(ctx context.Context, assessmentID string) (*AssessmentHistoryResult, *http.Response, error) {
	path := fmt.Sprintf("%s/%s", assessmentHistoriesBasePath, assessmentID)

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	assessmentHistoryResult := new(AssessmentHistoryResult)
	resp, err := s.client.Do(ctx, req, &assessmentHistoryResult)
	if err != nil {
		return nil, resp, err
	}

	return assessmentHistoryResult, resp, err
}

// DeleteAssessmentResult
func (s *AssessmentHistoriesServiceOp) DeleteAssessmentResult(ctx context.Context, assessmentID string) (*http.Response, error) {
	path := fmt.Sprintf("%s?historyId=%s", assessmentHistoriesBasePath, assessmentID)

	req, err := s.client.NewRequest(ctx, http.MethodDelete, path, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(ctx, req, nil)
	if err != nil {
		return resp, err
	}

	// Delete returns a 204 No Content.
	// Error on anything else.
	if resp.StatusCode != 204 {
		return resp, fmt.Errorf("Expected Status Code 204. Got: %v", resp.StatusCode)
	}

	return resp, err
}
