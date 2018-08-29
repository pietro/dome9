package dome9

// CloudAccountMissingPermissions is the list of permissions Dome9 is missing for the Cloud Account.
type CloudAccountMissingPermissions struct {
	ID      string                             `json:"id"`
	Actions []CloudAccountExternalActionStatus `json:"actions"`
}

// CloudAccountExternalActionStatus is the status of an action Dome9 performed on a Cloud Account.
type CloudAccountExternalActionStatus struct {
	Type    string                     `json:"type"`
	SubType string                     `json:"subType"`
	Total   int64                      `json:"total"`
	Error   *CloudAccountActionFailure `json:"error"`
}

// CloudAccountActionFailure has the error code and detailed message from a failed action Dome9 performed on a Cloud Account.
type CloudAccountActionFailure struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// MissingPermission has the details of permissions Dome9 is missing for an Entity type and Sub type.
type MissingPermission struct {
	Srl               string                     `json:"srl"`
	ConsecutiveFails  int32                      `json:"consecutiveFails"`
	LastFail          string                     `json:"lastFail"`
	LastSuccess       string                     `json:"lastSuccess"`
	FirstFail         string                     `json:"firstFail"`
	LastFailErrorCode string                     `json:"lastFailErrorCode"`
	LastFailMessage   string                     `json:"lastFailMessage"`
	ID                string                     `json:"id"`
	RetryMetadata     *MissingPermissionMetadata `json:"retryMetadata"`
	CloudAccountID    string                     `json:"cloudAccountId"`
	Vendor            string                     `json:"vendor"`
}

// MissingPermissionMetadata is the metadata of permissions Dome9 is missing for an Entity type and Sub type.
type MissingPermissionMetadata struct {
	Permissions []string `json:"permissions"`
	EntityType  string   `json:"entityType"`
	SubType     string   `json:"subType"`
}
