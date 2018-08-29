package dome9

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestAssessments_RunBundle(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/assessment/bundleV2", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprint(w, `{
  "request": {
    "dome9CloudAccountId": "00000000-0000-0000-0000-000000000000",
    "externalCloudAccountId": "string",
    "cloudAccountId": "string",
    "region": "string",
    "cloudNetwork": "string",
    "cloudAccountType": "Aws",
    "requestId": "00000000-0000-0000-0000-000000000000"
  },
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
  "locationMetadata": {
    "account": {
      "srl": "string",
      "name": "string",
      "id": "string",
      "externalId": "string"
    },
    "region": {
      "srl": "string",
      "name": "string",
      "id": "string",
      "externalId": "string"
    },
    "cloudNetwork": {
      "srl": "string",
      "name": "string",
      "id": "string",
      "externalId": "string"
    }
  },
  "testEntities": {
    "notSupported": [
      {}
    ],
    "instance": [
      {}
    ],
    "securityGroup": [
      {}
    ],
    "elb": [
      {}
    ],
    "rds": [
      {}
    ],
    "lambda": [
      {}
    ],
    "region": [
      {}
    ],
    "virtualMachine": [
      {}
    ],
    "networkSecurityGroup": [
      {}
    ],
    "cloudTrail": [
      {}
    ],
    "nacl": [
      {}
    ],
    "vpc": [
      {}
    ],
    "subnet": [
      {}
    ],
    "s3Bucket": [
      {}
    ],
    "applicationLoadBalancer": [
      {}
    ],
    "iamUser": [
      {}
    ],
    "iamRole": [
      {}
    ],
    "iam": [
      {}
    ],
    "redshift": [
      {}
    ],
    "kms": [
      {}
    ],
    "default": [
      {}
    ],
    "vmInstance": [
      {}
    ],
    "iamGroup": [
      {}
    ],
    "efs": [
      {}
    ],
    "network": [
      {}
    ],
    "elastiCache": [
      {}
    ],
    "loadBalancer": [
      {}
    ],
    "vNet": [
      {}
    ],
    "sqldb": [
      {}
    ],
    "redisCache": [
      {}
    ],
    "applicationGateway": [
      {}
    ],
    "resourceGroup": [
      {}
    ],
    "sqlServer": [
      {}
    ],
    "ecsCluster": [
      {}
    ],
    "keyVault": [
      {}
    ],
    "networkLoadBalancer": [
      {}
    ],
    "networkInterface": [
      {}
    ],
    "ecsTaskDefinition": [
      {}
    ],
    "iamPolicy": [
      {}
    ],
    "volume": [
      {}
    ],
    "cloudFront": [
      {}
    ],
    "kinesis": [
      {}
    ],
    "iamServerCertificate": [
      {}
    ],
    "route53HostedZone": [
      {}
    ],
    "route53RecordSetGroup": [
      {}
    ],
    "acmCertificate": [
      {}
    ],
    "route53Domain": [
      {}
    ],
    "storageAccount": [
      {}
    ],
    "dynamoDbTable": [
      {}
    ],
    "ami": [
      {}
    ],
    "vpnGateway": [
      {}
    ],
    "virtualMfaDevices": [
      {}
    ],
    "internetGateway": [
      {}
    ],
    "wafRegional": [
      {}
    ],
    "lock": [
      {}
    ],
    "vpnConnection": [
      {}
    ],
    "ecsTask": [
      {}
    ],
    "customerGateway": [
      {}
    ],
    "gcpSecurityGroup": [
      {}
    ],
    "elasticIP": [
      {}
    ],
    "iamInstanceProfile": [
      {}
    ]
  },
  "assessmentPassed": true,
  "hasErrors": true,
  "id": 0
}`)
	})

	bundle := &AssessmentBundleRequest{
		ID:          0,
		Name:        "string",
		Description: "string",
		CFT: &AssessmentCFTRequest{
			RootName: "string",
			Params:   []CFTParameterRequest{{Key: "string", Value: "string"}},
			Files:    []CFTFileRequest{{Name: "string", Template: "string"}}},
		IsCFT:                  true,
		Dome9CloudAccountID:    "00000000-0000-0000-0000-000000000000",
		ExternalCloudAccountID: "string",
		CloudAccountID:         "string",
		Region:                 "string",
		CloudNetwork:           "string",
		CloudAccountType:       "Aws",
		RequestID:              "00000000-0000-0000-0000-000000000000"}

	assessmentResult, _, err := client.Assessments.RunBundle(bundle)
	if err != nil {
		t.Errorf("Assessments.RunBundle returned error: %v", err)
	}

	expected := &AssessmentResult{
		Request: BaseAssessmentRequest{
			Dome9CloudAccountID:    "00000000-0000-0000-0000-000000000000",
			ExternalCloudAccountID: "string",
			CloudAccountID:         "string",
			Region:                 "string",
			CloudNetwork:           "string",
			CloudAccountType:       "Aws",
			RequestID:              "00000000-0000-0000-0000-000000000000"},
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
		LocationMetadata: &LocationMetadata{
			Account: &LocationConventionMetadata{
				SRL:        "string",
				Name:       "string",
				ID:         "string",
				ExternalID: "string"},
			Region: &LocationConventionMetadata{
				SRL:        "string",
				Name:       "string",
				ID:         "string",
				ExternalID: "string"},
			CloudNetwork: &LocationConventionMetadata{
				SRL:        "string",
				Name:       "string",
				ID:         "string",
				ExternalID: "string"}},
		TestEntities: map[string]interface{}{"sqldb": []interface{}{map[string]interface{}{}},
			"applicationGateway":      []interface{}{map[string]interface{}{}},
			"iamPolicy":               []interface{}{map[string]interface{}{}},
			"volume":                  []interface{}{map[string]interface{}{}},
			"nacl":                    []interface{}{map[string]interface{}{}},
			"subnet":                  []interface{}{map[string]interface{}{}},
			"loadBalancer":            []interface{}{map[string]interface{}{}},
			"kms":                     []interface{}{map[string]interface{}{}},
			"default":                 []interface{}{map[string]interface{}{}},
			"wafRegional":             []interface{}{map[string]interface{}{}},
			"storageAccount":          []interface{}{map[string]interface{}{}},
			"notSupported":            []interface{}{map[string]interface{}{}},
			"securityGroup":           []interface{}{map[string]interface{}{}},
			"vmInstance":              []interface{}{map[string]interface{}{}},
			"iamGroup":                []interface{}{map[string]interface{}{}},
			"rds":                     []interface{}{map[string]interface{}{}},
			"virtualMachine":          []interface{}{map[string]interface{}{}},
			"applicationLoadBalancer": []interface{}{map[string]interface{}{}},
			"vpnGateway":              []interface{}{map[string]interface{}{}},
			"vpnConnection":           []interface{}{map[string]interface{}{}},
			"region":                  []interface{}{map[string]interface{}{}},
			"sqlServer":               []interface{}{map[string]interface{}{}},
			"kinesis":                 []interface{}{map[string]interface{}{}},
			"elastiCache":             []interface{}{map[string]interface{}{}},
			"dynamoDbTable":           []interface{}{map[string]interface{}{}},
			"gcpSecurityGroup":        []interface{}{map[string]interface{}{}},
			"networkSecurityGroup":    []interface{}{map[string]interface{}{}},
			"vpc":                     []interface{}{map[string]interface{}{}},
			"iamUser":                 []interface{}{map[string]interface{}{}},
			"network":                 []interface{}{map[string]interface{}{}},
			"cloudFront":              []interface{}{map[string]interface{}{}},
			"ecsTask":                 []interface{}{map[string]interface{}{}},
			"elb":                     []interface{}{map[string]interface{}{}},
			"resourceGroup":           []interface{}{map[string]interface{}{}},
			"route53HostedZone":       []interface{}{map[string]interface{}{}},
			"redisCache":              []interface{}{map[string]interface{}{}},
			"networkLoadBalancer":     []interface{}{map[string]interface{}{}},
			"customerGateway":         []interface{}{map[string]interface{}{}},
			"iamRole":                 []interface{}{map[string]interface{}{}},
			"redshift":                []interface{}{map[string]interface{}{}},
			"networkInterface":        []interface{}{map[string]interface{}{}},
			"acmCertificate":          []interface{}{map[string]interface{}{}},
			"lock":                    []interface{}{map[string]interface{}{}},
			"iamInstanceProfile":      []interface{}{map[string]interface{}{}},
			"s3Bucket":                []interface{}{map[string]interface{}{}},
			"iam":                     []interface{}{map[string]interface{}{}},
			"vNet":                    []interface{}{map[string]interface{}{}},
			"internetGateway":         []interface{}{map[string]interface{}{}},
			"route53RecordSetGroup":   []interface{}{map[string]interface{}{}},
			"virtualMfaDevices":       []interface{}{map[string]interface{}{}},
			"ecsTaskDefinition":       []interface{}{map[string]interface{}{}},
			"route53Domain":           []interface{}{map[string]interface{}{}},
			"ami":                     []interface{}{map[string]interface{}{}},
			"iamServerCertificate":    []interface{}{map[string]interface{}{}},
			"elasticIP":               []interface{}{map[string]interface{}{}},
			"instance":                []interface{}{map[string]interface{}{}},
			"cloudTrail":              []interface{}{map[string]interface{}{}},
			"keyVault":                []interface{}{map[string]interface{}{}},
			"lambda":                  []interface{}{map[string]interface{}{}},
			"efs":                     []interface{}{map[string]interface{}{}},
			"ecsCluster":              []interface{}{map[string]interface{}{}}},
		AssessmentPassed: true,
		HasErrors:        true,
		ID:               0}

	if !reflect.DeepEqual(assessmentResult, expected) {
		t.Errorf("Assessments.RunBundle\n got=%#v\nwant=%#v", assessmentResult, expected)
	}
}
