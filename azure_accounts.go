package dome9

import (
	"fmt"
	"net/http"
)

const azureCloudAccountBasePath = "v2/AzureCloudAccount"

// AzureCloudAccountsService is an interface for interfacing with the
// AzureCloudAccount endpoints of the Dome9 API.
// See: https://api-v2-docs.dome9.com/#Dome9-API-AzureCloudAccount
type AzureCloudAccountsService interface {
	List() ([]AzureCloudAccount, *http.Response, error)
	Delete(string) (*http.Response, error)
	Create(AzureCloudAccount) (*http.Response, error)
	GetMissingPermissions(string) (*CloudAccountMissingPermissions, *http.Response, error)
	GetMissingPermissionsByEntityType(string, string, string) ([]MissingPermission, *http.Response, error)
	ResetMissingPermissions(string) (*http.Response, error)
	UpdateOperationMode(string, AzureAccountOperationMode) (*AzureCloudAccount, *http.Response, error)
	UpdateAccountName(string, AzureAccountNameMode) (*AzureCloudAccount, *http.Response, error)
}

// AzureCloudAccountsServiceOp handles communication with the AzureCloudAccount
// related methods of the Dome9 API.
type AzureCloudAccountsServiceOp struct {
	client *Client
}

var _ AzureCloudAccountsService = &AzureCloudAccountsServiceOp{}

// AzureAccountCredentials are the credentials used to access an Azure account.
type AzureAccountCredentials struct {
	ClientID       string `json:"clientId,omitempty"`
	ClientPassword string `json:"clientPassword,omitempty"`
}

// AzureCloudAccount are the details of an Azure account.
type AzureCloudAccount struct {
	ID             string                   `json:"id"`
	Name           string                   `json:"name"`
	SubscriptionID string                   `json:"subscriptionId"`
	TenantID       string                   `json:"tenantID"`
	Credentials    *AzureAccountCredentials `json:"credentials"`
	OperationMode  string                   `json:"operationMode"`
	Error          string                   `json:"error"`
	CreationDate   string                   `json:"creationDate"`
}

// AzureAccountOperationMode is the operations mode for an Azure account in Dome9. Modes can be Read-Only or Manage.
type AzureAccountOperationMode struct {
	OperationMode string `json:"operationMode"`
}

// AzureAccountNameMode is used to create the JSON object to update an Azure Account Name.
type AzureAccountNameMode struct {
	Name string `json:"name"`
}

type azureCloudAccountsRoot []AzureCloudAccount

type missingPermissionRoot []MissingPermission

// List all AzureCloudAccounts.
func (s *AzureCloudAccountsServiceOp) List() ([]AzureCloudAccount, *http.Response, error) {
	path := azureCloudAccountBasePath

	req, err := s.client.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	azureAccounts := new(azureCloudAccountsRoot)
	resp, err := s.client.Do(req, &azureAccounts)
	if err != nil {
		return nil, resp, err
	}

	return *azureAccounts, resp, err
}

// Delete an Azure account from a Dome9 account (the Azure account is not deleted from Azure).
func (s *AzureCloudAccountsServiceOp) Delete(accountID string) (*http.Response, error) {
	path := fmt.Sprintf("%s/%s", azureCloudAccountBasePath, accountID)

	req, err := s.client.NewRequest(http.MethodDelete, path, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req, nil)
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

// Create (onboard) an Azure account to the user's Dome9 account.
func (s *AzureCloudAccountsServiceOp) Create(azureAccount AzureCloudAccount) (*http.Response, error) {
	path := azureCloudAccountBasePath

	req, err := s.client.NewRequest(http.MethodPost, path, azureAccount)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req, nil)
	if err != nil {
		return resp, err
	}

	// Create returns a 200 with empty body.
	// Error on anything else.
	if resp.StatusCode != 200 {
		return resp, fmt.Errorf("Expected Status Code 200. Got: %v", resp.StatusCode)
	}

	return resp, err
}

// GetMissingPermissions lists missing permissions for an Azure account in a Dome9 account.
func (s *AzureCloudAccountsServiceOp) GetMissingPermissions(accountID string) (*CloudAccountMissingPermissions, *http.Response, error) {
	path := fmt.Sprintf("%s/%s/MissingPermissions", azureCloudAccountBasePath, accountID)

	req, err := s.client.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	missingPerms := new(CloudAccountMissingPermissions)
	resp, err := s.client.Do(req, &missingPerms)
	if err != nil {
		return nil, resp, err
	}

	return missingPerms, resp, err
}

// GetMissingPermissionsByEntityType lists missing permissions for a specific cloud entity type and Azure cloud account.
func (s *AzureCloudAccountsServiceOp) GetMissingPermissionsByEntityType(accountID, entityType, subType string) ([]MissingPermission, *http.Response, error) {
	path := fmt.Sprintf("%s/%s/MissingPermissions?entityType=%s&subType=%s", azureCloudAccountBasePath, accountID, entityType, subType)

	req, err := s.client.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	missingPerms := new(missingPermissionRoot)
	resp, err := s.client.Do(req, &missingPerms)
	if err != nil {
		return nil, resp, err
	}

	return *missingPerms, resp, err
}

// ResetMissingPermissions resets (re-validate) the missing permissions indication for an Azure account in Dome9.
func (s *AzureCloudAccountsServiceOp) ResetMissingPermissions(accountID string) (*http.Response, error) {
	path := fmt.Sprintf("%s/%s/MissingPermissions/Reset", azureCloudAccountBasePath, accountID)

	req, err := s.client.NewRequest(http.MethodPut, path, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req, nil)
	if err != nil {
		return resp, err
	}

	// ResetMissingPermissions returns a 204 No Content.
	// Error on anything else.
	if resp.StatusCode != 204 {
		return resp, fmt.Errorf("Expected Status Code 204. Got: %v", resp.StatusCode)
	}
	return resp, err
}

// UpdateOperationMode changes the operations mode for an Azure account in Dome9. Modes can be Read-Only or Manage.
func (s *AzureCloudAccountsServiceOp) UpdateOperationMode(accountID string, operationMode AzureAccountOperationMode) (*AzureCloudAccount, *http.Response, error) {
	path := fmt.Sprintf("%s/%s/OperationMode", azureCloudAccountBasePath, accountID)

	req, err := s.client.NewRequest(http.MethodPut, path, operationMode)
	if err != nil {
		return nil, nil, err
	}

	azureAccount := new(AzureCloudAccount)
	resp, err := s.client.Do(req, &azureAccount)
	if err != nil {
		return nil, resp, err
	}

	return azureAccount, resp, err
}

// UpdateAccountName changes the account name (in Dome9) for an Azure account.
func (s *AzureCloudAccountsServiceOp) UpdateAccountName(accountID string, accountName AzureAccountNameMode) (*AzureCloudAccount, *http.Response, error) {
	path := fmt.Sprintf("%s/%s/AccountName", azureCloudAccountBasePath, accountID)

	req, err := s.client.NewRequest(http.MethodPut, path, accountName)
	if err != nil {
		return nil, nil, err
	}

	azureAccount := new(AzureCloudAccount)
	resp, err := s.client.Do(req, &azureAccount)
	if err != nil {
		return nil, resp, err
	}

	return azureAccount, resp, err
}
