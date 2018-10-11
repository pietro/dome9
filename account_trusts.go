package dome9

import (
	"context"
	"fmt"
	"net/http"
)

const accountTrustsBasePath = "v2/AccountTrust"

// AccountTrustsService resource has methods to define a trust relationship
// between Dome9 accounts. This allows users in a one account to make changes
// and perform operations in another account.
// See: https://api-v2-docs.dome9.com/#Dome9-API-AccountTrust
type AccountTrustsService interface {
	GetAssumableRoles(context.Context) ([]AccountTrustAssumableRoles, *http.Response, error)
	List(context.Context, string) ([]AccountTrust, *http.Response, error)
	Create(context.Context, *AccountTrustCreateRequest) (*http.Response, error)
	Update(context.Context, string, *AccountTrustUpdateRequest) (*http.Response, error)
	Delete(context.Context, string) (*http.Response, error)
}

// AccountTrustsServiceOp handles communication with the AccountTrusts
// related methods of the Dome9 API.
type AccountTrustsServiceOp struct {
	client *Client
}

var _ AccountTrustsService = &AccountTrustsServiceOp{}

// AccountTrustAssumableRoles
type AccountTrustAssumableRoles struct {
	AccountName string
	AccountID   int64
	Roles       []string
}

// AccountTrust
type AccountTrust struct {
	ID                string
	TargetAccountName string
	SourceAccountName string
	SourceAccountID   string
	Description       string
	Restrictions      *AccountTrustRestrictions
}

// AccountTrustRestrictions
type AccountTrustRestrictions struct {
	Roles []string
}

// AccountTrustCreateRequest
type AccountTrustCreateRequest struct {
	SourceAccountID string
	Description     string
	Restrictions    *AccountTrustRestrictions
}

// AccountTrustUpdateRequest
type AccountTrustUpdateRequest struct {
	Description  string
	Restrictions *AccountTrustRestrictions
}

// GetAssumableRoles
func (s *AccountTrustsServiceOp) GetAssumableRoles(ctx context.Context) ([]AccountTrustAssumableRoles, *http.Response, error) {
	path := fmt.Sprintf("%s/assumable-roles", accountTrustsBasePath)

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	var assumableRoles []AccountTrustAssumableRoles
	resp, err := s.client.Do(ctx, req, &assumableRoles)
	if err != nil {
		return nil, resp, err
	}

	return assumableRoles, resp, err
}

// List of accounts which are trusted by or trust this account according to the given "trustDirection""
func (s *AccountTrustsServiceOp) List(ctx context.Context, trustDirection string) ([]AccountTrust, *http.Response, error) {
	path := fmt.Sprintf("%s?trustDirection=%s", accountTrustsBasePath, trustDirection)

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	var trusts []AccountTrust
	resp, err := s.client.Do(ctx, req, &trusts)
	if err != nil {
		return nil, resp, err
	}

	return trusts, resp, err
}

// Create
func (s *AccountTrustsServiceOp) Create(ctx context.Context, createRequest *AccountTrustCreateRequest) (*http.Response, error) {
	path := accountTrustsBasePath

	req, err := s.client.NewRequest(ctx, http.MethodPost, path, createRequest)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(ctx, req, nil)
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

// Update
func (s *AccountTrustsServiceOp) Update(ctx context.Context, trustID string, updateRequest *AccountTrustUpdateRequest) (*http.Response, error) {
	path := fmt.Sprintf("%s/%s", accountTrustsBasePath, trustID)

	req, err := s.client.NewRequest(ctx, http.MethodPut, path, updateRequest)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(ctx, req, nil)
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

// Delete
func (s *AccountTrustsServiceOp) Delete(ctx context.Context, trustID string) (*http.Response, error) {
	path := fmt.Sprintf("%s/%s", accountTrustsBasePath, trustID)

	req, err := s.client.NewRequest(ctx, http.MethodDelete, path, nil)
	if err != nil {
		return nil, nil
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
