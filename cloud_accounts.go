package dome9

// Missing permissions for a Cloud Account.
type CloudAccountMissingPermissions struct {
	Id      string                             `json:"id"`
	Actions []CloudAccountExternalActionStatus `json:"actions"`
}

type CloudAccountExternalActionStatus struct {
	Type    string                     `json:"type"`
	SubType string                     `json:"subType"`
	Total   int64                      `json:"total"`
	Error   *CloudAccountActionFailure `json:"error"`
}

type CloudAccountActionFailure struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type MissingPermission struct {
	Srl               string                     `json:"srl"`
	ConsecutiveFails  int32                      `json:"consecutiveFails"`
	LastFail          string                     `json:"lastFail"`
	LastSuccess       string                     `json:"lastSuccess"`
	FirstFail         string                     `json:"firstFail"`
	LastFailErrorCode string                     `json:"lastFailErrorCode"`
	LastFailMessage   string                     `json:"lastFailMessage"`
	Id                string                     `json:"id"`
	RetryMetadata     *MissingPermissionMetadata `json:"retryMetadata"`
	CloudAccountId    string                     `json:"cloudAccountId"`
	Vendor            string                     `json:"vendor"`
}

type MissingPermissionMetadata struct {
	Permissions []string `json:"permissions"`
	EntityType  string   `json:"entityType"`
	SubType     string   `json:"subType"`
}
