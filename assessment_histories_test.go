package dome9

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestAssessmentHistories_GetBundleResults(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/AssessmentHistoryV2", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `[{
  "triggeredBy": "Unknown",
  "tests": [
    {
      "error": "string",
      "testedCount": 0,
      "relevantCount": 0,
      "nonComplyingCount": 0,
      "entityResults": [
        {
          "isRelevant": true,
          "isValid": true,
          "error": "string",
          "testObj": {}
        }
      ],
      "rule": {
        "name": "string",
        "severity": "Low",
        "logic": "string",
        "description": "string",
        "remediation": "string",
        "complianceTag": "string",
        "domain": "string",
        "priority": "string",
        "controlTitle": "string",
        "ruleId": "string",
        "logicHash": "string",
        "isDefault": true
      },
      "testPassed": true
    }
  ],
  "testEntities": {
    "notSupported": [
      {}
    ],
    "instance": [
      {}
    ]
  },
  "createdTime": "2018-08-26T16:11:12Z",
  "id": 0,
  "assessmentPassed": true,
  "hasErrors": true,
  "stats": {
    "passed": 0,
    "failed": 0,
    "error": 0,
    "failedTests": 0,
    "logicallyTested": 0,
    "failedEntities": 0
  },
  "request": {
    "isTemplate": true,
    "id": 0,
    "name": "string",
    "description": "string",
    "cft": {
      "rootName": "string",
      "params": [
        {
          "key": "string",
          "value": "string"
        }
      ],
      "files": [
        {
          "name": "string",
          "template": "string"
        }
      ]
    },
    "isCft": true,
    "dome9CloudAccountId": "00000000-0000-0000-0000-000000000000",
    "externalCloudAccountId": "string",
    "cloudAccountId": "string",
    "region": "string",
    "cloudNetwork": "string",
    "cloudAccountType": "Aws",
    "requestId": "00000000-0000-0000-0000-000000000000"
  }
}]`)
	})

	assessmentHistories, _, err := client.AssessmentHistories.GetBundleResults(ctx, "123", "abc,def", "2018-08-26T16:11:12Z", "1", "0000")
	if err != nil {
		t.Errorf("AssessmentHistories.GetBundleResults returned error: %v", err)
	}

	expected := []AssessmentHistoryResult{{
		TriggeredBy: "Unknown",
		Tests: []RuleTestResult{{
			Error:             "string",
			TestedCount:       0,
			RelevantCount:     0,
			NonComplyingCount: 0,
			EntityResults: []ValidationResult{{
				Relevant: true,
				Valid:    true,
				Error:    "string",
				TestObj:  map[string]interface{}{}}},
			Rule: &RuleEntity{
				Name:          "string",
				Severity:      "Low",
				Logic:         "string",
				Description:   "string",
				Remediation:   "string",
				ComplianceTag: "string",
				Domain:        "string",
				Priority:      "string",
				ControlTitle:  "string",
				RuleID:        "string",
				LogicHash:     "string",
				Default:       true},
			TestPassed: true}},
		TestEntities: map[string]interface{}{
			"notSupported": []interface{}{map[string]interface{}{}},
			"instance":     []interface{}{map[string]interface{}{}}},
		CreatedTime:      "2018-08-26T16:11:12Z",
		ID:               0,
		AssessmentPassed: true,
		HasErrors:        true,
		Stats: AssessmentHistoryStats{
			Passed:          0,
			Failed:          0,
			Error:           0,
			FailedTests:     0,
			LogicallyTested: 0,
			FailedEntities:  0},
		Request: AssessmentHistoryBundleResult{
			Template:    true,
			ID:          0,
			Name:        "string",
			Description: "string",
			CFT: AssessmentCFTRequest{
				RootName: "string",
				Params: []CFTParameterRequest{{
					Key:   "string",
					Value: "string"}},
				Files: []CFTFileRequest{{
					Name:     "string",
					Template: "string"}}},
			IsCFT:                  true,
			Dome9CloudAccountID:    "00000000-0000-0000-0000-000000000000",
			ExternalCloudAccountID: "string",
			CloudAccountID:         "string",
			Region:                 "string",
			CloudNetwork:           "string",
			CloudAccountType:       "Aws",
			RequestID:              "00000000-0000-0000-0000-000000000000"},
	}}

	if !reflect.DeepEqual(assessmentHistories, expected) {
		t.Errorf("AssessmentHistories.GetBundleResults\n got=%#v\nwant=%#v", assessmentHistories, expected)
	}
}

func TestAssessmentHistories_GetAssessmentResult(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/AssessmentHistoryV2/123", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{
  "triggeredBy": "Unknown",
  "tests": [
    {
      "error": "string",
      "testedCount": 0,
      "relevantCount": 0,
      "nonComplyingCount": 0,
      "entityResults": [
        {
          "isRelevant": true,
          "isValid": true,
          "error": "string",
          "testObj": {}
        }
      ],
      "rule": {
        "name": "string",
        "severity": "Low",
        "logic": "string",
        "description": "string",
        "remediation": "string",
        "complianceTag": "string",
        "domain": "string",
        "priority": "string",
        "controlTitle": "string",
        "ruleId": "string",
        "logicHash": "string",
        "isDefault": true
      },
      "testPassed": true
    }
  ],
  "testEntities": {
    "notSupported": [
      {}
    ],
    "instance": [
      {}
    ]
  },
  "createdTime": "2018-08-26T16:11:12Z",
  "id": 0,
  "assessmentPassed": true,
  "hasErrors": true,
  "stats": {
    "passed": 0,
    "failed": 0,
    "error": 0,
    "failedTests": 0,
    "logicallyTested": 0,
    "failedEntities": 0
  },
  "request": {
    "isTemplate": true,
    "id": 0,
    "name": "string",
    "description": "string",
    "cft": {
      "rootName": "string",
      "params": [
        {
          "key": "string",
          "value": "string"
        }
      ],
      "files": [
        {
          "name": "string",
          "template": "string"
        }
      ]
    },
    "isCft": true,
    "dome9CloudAccountId": "00000000-0000-0000-0000-000000000000",
    "externalCloudAccountId": "string",
    "cloudAccountId": "string",
    "region": "string",
    "cloudNetwork": "string",
    "cloudAccountType": "Aws",
    "requestId": "00000000-0000-0000-0000-000000000000"
  }
}`)
	})
	assessmentHistory, _, err := client.AssessmentHistories.GetAssessmentResult(ctx, "123")
	if err != nil {
		t.Errorf("AssessmentHistories.GetAssessmentResult returned error: %v", err)
	}

	expected := &AssessmentHistoryResult{
		TriggeredBy: "Unknown",
		Tests: []RuleTestResult{{
			Error:             "string",
			TestedCount:       0,
			RelevantCount:     0,
			NonComplyingCount: 0,
			EntityResults: []ValidationResult{{
				Relevant: true,
				Valid:    true,
				Error:    "string",
				TestObj:  map[string]interface{}{}}},
			Rule: &RuleEntity{
				Name:          "string",
				Severity:      "Low",
				Logic:         "string",
				Description:   "string",
				Remediation:   "string",
				ComplianceTag: "string",
				Domain:        "string",
				Priority:      "string",
				ControlTitle:  "string",
				RuleID:        "string",
				LogicHash:     "string",
				Default:       true},
			TestPassed: true}},
		TestEntities: map[string]interface{}{
			"notSupported": []interface{}{map[string]interface{}{}},
			"instance":     []interface{}{map[string]interface{}{}}},
		CreatedTime:      "2018-08-26T16:11:12Z",
		ID:               0,
		AssessmentPassed: true,
		HasErrors:        true,
		Stats: AssessmentHistoryStats{
			Passed:          0,
			Failed:          0,
			Error:           0,
			FailedTests:     0,
			LogicallyTested: 0,
			FailedEntities:  0},
		Request: AssessmentHistoryBundleResult{
			Template:    true,
			ID:          0,
			Name:        "string",
			Description: "string",
			CFT: AssessmentCFTRequest{
				RootName: "string",
				Params: []CFTParameterRequest{{
					Key:   "string",
					Value: "string"}},
				Files: []CFTFileRequest{{
					Name:     "string",
					Template: "string"}}},
			IsCFT:                  true,
			Dome9CloudAccountID:    "00000000-0000-0000-0000-000000000000",
			ExternalCloudAccountID: "string",
			CloudAccountID:         "string",
			Region:                 "string",
			CloudNetwork:           "string",
			CloudAccountType:       "Aws",
			RequestID:              "00000000-0000-0000-0000-000000000000"},
	}

	if !reflect.DeepEqual(assessmentHistory, expected) {
		t.Errorf("AssessmentHistories.GetAssessmentResult\n got=%#v\nwant=%#v", assessmentHistory, expected)
	}
}

func TestAssessmentHistories_DeleteAssessmentResult(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/AssessmentHistoryV2", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
		w.WriteHeader(http.StatusNoContent)
	})

	_, err := client.AssessmentHistories.DeleteAssessmentResult(ctx, "123")
	if err != nil {
		t.Errorf("AssessmentHistories.DeleteAssessmentResult returned error: %v", err)
	}
}
