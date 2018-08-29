package dome9

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

const testAccountID = "1337-acct"

func TestAzureCloudAccounts_ListAzureCloudAccounts(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/AzureCloudAccount", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `[
  {
    "id": "00000000-0000-0000-0000-000000000000",
    "name": "string",
    "subscriptionId": "string",
    "tenantId": "string",
    "credentials": {
      "clientId": "string",
      "clientPassword": "string"
    },
    "operationMode": "Read",
    "error": "string",
    "creationDate": "2018-08-26T16:11:12Z"
  }
]`)
	})

	azureAccounts, _, err := client.AzureCloudAccounts.List()
	if err != nil {
		t.Errorf("AzureCloudAccounts.List returned error: %v", err)
	}

	expected := []AzureCloudAccount{{ID: "00000000-0000-0000-0000-000000000000", Name: "string", SubscriptionID: "string", TenantID: "string", Credentials: &AzureAccountCredentials{ClientID: "string", ClientPassword: "string"}, OperationMode: "Read", Error: "string", CreationDate: "2018-08-26T16:11:12Z"}}

	if !reflect.DeepEqual(azureAccounts, expected) {
		t.Errorf("AzureCloudAccounts.List\n got=%#v\nwant=%#v", azureAccounts, expected)
	}
}

func TestAzureCloudAccounts_DeleteAzureCloudAccount(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/AzureCloudAccount/"+testAccountID, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
		w.WriteHeader(http.StatusNoContent)
	})

	_, err := client.AzureCloudAccounts.Delete(testAccountID)
	if err != nil {
		t.Errorf("AzureCloudAccounts.Delete returned error: %v", err)
	}

}

func TestAzureCloudAccounts_CreateAzureCloudAccount(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/AzureCloudAccount", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprint(w, `{}`)
	})

	azureAccount := AzureCloudAccount{ID: "00000000-0000-0000-0000-000000000000", Name: "string", SubscriptionID: "string", TenantID: "string", Credentials: &AzureAccountCredentials{ClientID: "string", ClientPassword: "string"}, OperationMode: "Read", Error: "string", CreationDate: "2018-08-26T16:11:12Z"}

	_, err := client.AzureCloudAccounts.Create(azureAccount)
	if err != nil {
		t.Errorf("AzureCloudAccounts.Create returned error: %v", err)
	}
}

func TestAzureCloudAccounts_GetAzureCloudAccountsMissingPermissions(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/AzureCloudAccount/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{
  "id": "00000000-0000-0000-0000-000000000000",
  "actions": [
    {
      "type": "string",
      "subType": "string",
      "total": 0,
      "error": {
        "code": "string",
        "message": "string"
      }
    }
  ]
}`)
	})

	missingPerms, _, err := client.AzureCloudAccounts.GetMissingPermissions("00000000-0000-0000-0000-000000000000")
	if err != nil {
		t.Errorf("AzureCloudAccounts.GetMissingPermissions returned error: %v", err)
	}

	expected := &CloudAccountMissingPermissions{ID: "00000000-0000-0000-0000-000000000000", Actions: []CloudAccountExternalActionStatus{{Type: "string", SubType: "string", Total: 0, Error: &CloudAccountActionFailure{Code: "string", Message: "string"}}}}

	if !reflect.DeepEqual(missingPerms, expected) {
		t.Errorf("AzureCloudAccounts.GetMissingPermissions\n got=%#v\nwant=%#v", missingPerms, expected)
	}
}

func TestAzureCloudAccounts_GetAzureCloudAccountsMissingPermissionsByEntityType(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/AzureCloudAccount/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `[
  {
    "srl": "string",
    "consecutiveFails": 0,
    "lastFail": "2018-08-26T16:11:12Z",
    "lastSuccess": "2018-08-26T16:11:12Z",
    "firstFail": "2018-08-26T16:11:12Z",
    "lastFailErrorCode": "string",
    "lastFailMessage": "string",
    "id": "00000000-0000-0000-0000-000000000000",
    "retryMetadata": {
      "permissions": [
        "string"
      ],
      "entityType": "string",
      "subType": "string"
    },
    "cloudAccountId": "00000000-0000-0000-0000-000000000000",
    "vendor": "aws"
  }
]`)
	})

	missingPerms, _, err := client.AzureCloudAccounts.GetMissingPermissionsByEntityType("00000000-0000-0000-0000-000000000000", "entType", "subType")
	if err != nil {
		t.Errorf("AzureCloudAccounts.GetMissingPermissionsByEntityType returned error: %v", err)
	}

	expected := []MissingPermission{{Srl: "string", ConsecutiveFails: 0, LastFail: "2018-08-26T16:11:12Z", LastSuccess: "2018-08-26T16:11:12Z", FirstFail: "2018-08-26T16:11:12Z", LastFailErrorCode: "string", LastFailMessage: "string", ID: "00000000-0000-0000-0000-000000000000", RetryMetadata: &MissingPermissionMetadata{Permissions: []string{"string"}, EntityType: "string", SubType: "string"}, CloudAccountID: "00000000-0000-0000-0000-000000000000", Vendor: "aws"}}

	if !reflect.DeepEqual(missingPerms, expected) {
		t.Errorf("AzureCloudAccounts.GetMissingPermissionsByEntityType\n got=%#v\nwant=%#v", missingPerms, expected)
	}
}

func TestAzureCloudAccounts_ResetMissingPermissions(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/AzureCloudAccount/"+testAccountID+"/MissingPermissions/Reset", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPut)
		w.WriteHeader(http.StatusNoContent)
	})

	_, err := client.AzureCloudAccounts.ResetMissingPermissions(testAccountID)
	if err != nil {
		t.Errorf("AzureCloudAccounts.ResetMissingPermissions returned error: %v", err)
	}
}

func TestAzureCloudAccounts_UpdateOperationMode(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/AzureCloudAccount/"+testAccountID+"/OperationMode", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPut)
		fmt.Fprint(w, `{
    "id": "00000000-0000-0000-0000-000000000000",
    "name": "string",
    "subscriptionId": "string",
    "tenantId": "string",
    "credentials": {
      "clientId": "string",
      "clientPassword": "string"
    },
    "operationMode": "Read",
    "error": "string",
    "creationDate": "2018-08-26T16:11:12Z"
}`)
	})

	opMode := AzureAccountOperationMode{OperationMode: "Read"}
	azureAccount, _, err := client.AzureCloudAccounts.UpdateOperationMode(testAccountID, opMode)
	if err != nil {
		t.Errorf("AzureCloudAccounts.UpdateOperationMode returned error: %v", err)
	}

	expected := &AzureCloudAccount{ID: "00000000-0000-0000-0000-000000000000", Name: "string", SubscriptionID: "string", TenantID: "string", Credentials: &AzureAccountCredentials{ClientID: "string", ClientPassword: "string"}, OperationMode: "Read", Error: "string", CreationDate: "2018-08-26T16:11:12Z"}

	if !reflect.DeepEqual(azureAccount, expected) {
		t.Errorf("AzureCloudAccounts.List\n got=%#v\nwant=%#v", azureAccount, expected)
	}
}

func TestAzureCloudAccounts_UpdateAccountName(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/AzureCloudAccount/"+testAccountID+"/AccountName", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPut)
		fmt.Fprint(w, `{
    "id": "00000000-0000-0000-0000-000000000000",
    "name": "string",
    "subscriptionId": "string",
    "tenantId": "string",
    "credentials": {
      "clientId": "string",
      "clientPassword": "string"
    },
    "operationMode": "Read",
    "error": "string",
    "creationDate": "2018-08-26T16:11:12Z"
}`)
	})

	acctName := AzureAccountNameMode{Name: "foobared"}
	azureAccount, _, err := client.AzureCloudAccounts.UpdateAccountName(testAccountID, acctName)
	if err != nil {
		t.Errorf("AzureCloudAccounts.UpdateOperationMode returned error: %v", err)
	}

	expected := &AzureCloudAccount{ID: "00000000-0000-0000-0000-000000000000", Name: "string", SubscriptionID: "string", TenantID: "string", Credentials: &AzureAccountCredentials{ClientID: "string", ClientPassword: "string"}, OperationMode: "Read", Error: "string", CreationDate: "2018-08-26T16:11:12Z"}

	if !reflect.DeepEqual(azureAccount, expected) {
		t.Errorf("AzureCloudAccounts.List\n got=%#v\nwant=%#v", azureAccount, expected)
	}
}
