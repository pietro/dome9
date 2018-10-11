package dome9

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

const testTrustID = "1337-trust"

func TestAccountTrusts_GetAssumableRoles(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/AccountTrust/assumable-roles", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `[
{
  "accountName": "string",
  "accountId": 0,
  "roles": [
    "string"
  ]
}
]`)
	})

	assumableRoles, _, err := client.AccountTrusts.GetAssumableRoles(ctx)
	if err != nil {
		t.Errorf("AccountTrusts.GetAssumableRoles returned error: %v", err)
	}

	expected := []AccountTrustAssumableRoles{{AccountName: "string", AccountID: 0, Roles: []string{"string"}}}

	if !reflect.DeepEqual(assumableRoles, expected) {
		t.Errorf("AccountTrusts.List\ngot=%#v\nwant=%#v", assumableRoles, expected)
	}
}

func TestAccountTrusts_List(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/AccountTrust", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `[
{
  "id": "00000000-0000-0000-0000-000000000000",
  "targetAccountName": "string",
  "sourceAccountName": "string",
  "sourceAccountId": "00000000-0000-0000-0000-000000000000",
  "description": "string",
  "restrictions": {
    "roles": [
      "string"
    ]
  }
}
]`)
	})

	accountTrusts, _, err := client.AccountTrusts.List(ctx, "MyAccountIsTarget")
	if err != nil {
		t.Errorf("AccountTrusts.List returned error: %v", err)
	}

	expected := []AccountTrust{{ID: "00000000-0000-0000-0000-000000000000", TargetAccountName: "string", SourceAccountName: "string", SourceAccountID: "00000000-0000-0000-0000-000000000000", Description: "string", Restrictions: &AccountTrustRestrictions{Roles: []string{"string"}}}}

	if !reflect.DeepEqual(accountTrusts, expected) {
		t.Errorf("AccountTrusts.List\ngot=%#v\nwant=%#v", accountTrusts, expected)
	}
}

func TestAccountTrusts_Delete(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/AccountTrust/"+testTrustID, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
		w.WriteHeader(http.StatusNoContent)
	})

	_, err := client.AccountTrusts.Delete(ctx, testTrustID)
	if err != nil {
		t.Errorf("AccountTrusts.Delete returned error: %v", err)
	}
}

func TestAccountTrusts_Create(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/AccountTrust", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprint(w, `{}`)
	})

	accountTrust := &AccountTrustCreateRequest{SourceAccountID: "0000", Description: "string", Restrictions: &AccountTrustRestrictions{Roles: []string{"string"}}}

	_, err := client.AccountTrusts.Create(ctx, accountTrust)
	if err != nil {
		t.Errorf("AzureCloudAccounts.Create returned error: %v", err)
	}
}

func TestAccountTrusts_Update(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/AccountTrust/"+testTrustID, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPut)
		fmt.Fprint(w, `{}`)
	})

	accountTrust := &AccountTrustUpdateRequest{Description: "string", Restrictions: &AccountTrustRestrictions{Roles: []string{"string"}}}

	_, err := client.AccountTrusts.Update(ctx, testTrustID, accountTrust)
	if err != nil {
		t.Errorf("AccountTrusts.Update returned error: %v", err)
	}
}
