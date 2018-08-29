package dome9

import (
	"fmt"
	"net/http"
)

const assessmentsBasePath = "v2/assessment"

// AsessmentsService resource has methods to run compliance assessments on cloud
// accounts, and view the results.
// See: https://api-v2-docs.dome9.com/#Dome9-API-Assessment
type AssessmentsService interface {
	RunBundle(*AssessmentBundleRequest) (*AssessmentResult, *http.Response, error)
}

// AssessmentsServiceOp handles communication with the Assessments
// related methods of the Dome9 API.
type AssessmentsServiceOp struct {
	client *Client
}

var _ AssessmentsService = &AssessmentsServiceOp{}

// AssessmentBundleRequest
type AssessmentBundleRequest struct {
	ID                     int64                 `json:"id"`
	Name                   string                `json:"name"`
	Description            string                `json:"description"`
	CFT                    *AssessmentCFTRequest `json:"cft"`
	IsCFT                  bool                  `json:"isCft"`
	Dome9CloudAccountID    string                `json:"dome9CloudAccountId"`
	ExternalCloudAccountID string                `json:"externalCloudAccountId"`
	CloudAccountID         string                `json:"cloudAccountId"`
	Region                 string                `json:"region"`
	CloudNetwork           string                `json:"cloudNetwork"`
	CloudAccountType       string                `json:"cloudAccountType"`
	RequestID              string                `json:"requestId"`
}

// AssessmentCFTRequest
type AssessmentCFTRequest struct {
	RootName string                `json:"rootName"`
	Params   []CFTParameterRequest `json:"params"`
	Files    []CFTFileRequest      `json:"files"`
}

// CFTParameterRequest
type CFTParameterRequest struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// CFTFileRequest
type CFTFileRequest struct {
	Name     string `json:"name"`
	Template string `json:"template"`
}

// AssessmentResult aggregates results from multiple tests.
type AssessmentResult struct {
	Request          BaseAssessmentRequest `json:"request"`
	Tests            []RuleTestResult      `json:"tests"`
	LocationMetadata *LocationMetadata     `json:"locationMetadata"`
	TestEntities     interface{}           `json:"testEntities"`
	AssessmentPassed bool                  `json:"assessmentPassed"`
	HasErrors        bool                  `json:"hasErrors"`
	ID               int64                 `json:"id"`
}

// BaseAssessmentRequest
type BaseAssessmentRequest struct {
	Dome9CloudAccountID    string `json:"dome9CloudAccountId"`
	ExternalCloudAccountID string `json:"externalCloudAccountId"`
	CloudAccountID         string `json:"cloudAccountId"`
	Region                 string `json:"region"`
	CloudNetwork           string `json:"cloudNetwork"`
	CloudAccountType       string `json:"cloudAccountType"`
	RequestID              string `json:"requestId"`
}

// RuleTestResult
type RuleTestResult struct {
	Error             string             `json:"error"`
	TestedCount       int32              `json:"testesCount"`
	RelevantCount     int32              `json:"relevantCount"`
	NonComplyingCount int32              `json:"nonComplyingCount"`
	EntityResults     []ValidationResult `json:"entityResults"`
	Rule              *RuleEntity        `json:"rule"`
	TestPassed        bool               `json:"testPassed"`
}

// ValidationResult
type ValidationResult struct {
	Relevant bool        `json:"isRelevant"`
	Valid    bool        `json:"isValid"`
	Error    string      `json:"error"`
	TestObj  interface{} `json:"testObj"`
}

// RuleEntity
type RuleEntity struct {
	Name          string `json:"name"`
	Severity      string `json:"severity"`
	Logic         string `json:"logic"`
	Description   string `json:"description"`
	Remediation   string `json:"remediation"`
	ComplianceTag string `json:"complianceTag"`
	Domain        string `json:"domain"`
	Priority      string `json:"priority"`
	ControlTitle  string `json:"controlTitle"`
	RuleID        string `json:"ruleId"`
	LogicHash     string `json:"logicHash"`
	Default       bool   `json:"isDefault"`
}

// LocationMetadata
type LocationMetadata struct {
	Account      *LocationConventionMetadata `json:"account"`
	Region       *LocationConventionMetadata `json:"region"`
	CloudNetwork *LocationConventionMetadata `json:"cloudNetwork"`
}

// LocationConventionMetadata
type LocationConventionMetadata struct {
	SRL        string `json:"srl"`
	Name       string `json:"name"`
	ID         string `json:"id"`
	ExternalID string `json:"externalId"`
}

// RunBundle runs an assessment on a cloud environment using a bundle.
func (s *AssessmentsServiceOp) RunBundle(bundleRequest *AssessmentBundleRequest) (*AssessmentResult, *http.Response, error) {
	path := fmt.Sprintf("%s/bundleV2", assessmentsBasePath)

	req, err := s.client.NewRequest(http.MethodPost, path, bundleRequest)
	if err != nil {
		return nil, nil, err
	}

	result := new(AssessmentResult)
	resp, err := s.client.Do(req, result)
	if err != nil {
		return nil, resp, err
	}

	return result, resp, err
}
