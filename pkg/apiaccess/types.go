// Code generated by tutone: DO NOT EDIT
package apiaccess

// APIAccessIngestKeyErrorType - The type of error.
type APIAccessIngestKeyErrorType string

var APIAccessIngestKeyErrorTypeTypes = struct {
	// Occurs when the user issuing the mutation does not have sufficient permissions to perform the action for a key.
	FORBIDDEN APIAccessIngestKeyErrorType
	// Occurs when the action taken on a key did not successfully pass validation.
	INVALID APIAccessIngestKeyErrorType
	// Occurs when the requested key `id` was not found.
	NOT_FOUND APIAccessIngestKeyErrorType
}{
	// Occurs when the user issuing the mutation does not have sufficient permissions to perform the action for a key.
	FORBIDDEN: "FORBIDDEN",
	// Occurs when the action taken on a key did not successfully pass validation.
	INVALID: "INVALID",
	// Occurs when the requested key `id` was not found.
	NOT_FOUND: "NOT_FOUND",
}

// APIAccessIngestKeyType - The type of ingest key, which dictates what types of agents can use it to report.
type APIAccessIngestKeyType string

var APIAccessIngestKeyTypeTypes = struct {
	// Ingest keys of type `BROWSER` mean browser agents will use them to report data to New Relic.
	BROWSER APIAccessIngestKeyType
	// For ingest keys of type `LICENSE`: APM and Infrastructure agents use the key to report data to New Relic.
	LICENSE APIAccessIngestKeyType
}{
	// Ingest keys of type `BROWSER` mean browser agents will use them to report data to New Relic.
	BROWSER: "BROWSER",
	// For ingest keys of type `LICENSE`: APM and Infrastructure agents use the key to report data to New Relic.
	LICENSE: "LICENSE",
}

// APIAccessKeyType - The type of key.
type APIAccessKeyType string

var APIAccessKeyTypeTypes = struct {
	// An ingest key is used by New Relic agents to authenticate with New Relic and send data to the assigned account.
	INGEST APIAccessKeyType
	// A user key is used by New Relic users to authenticate with New Relic and to interact with the New Relic GraphQL API .
	USER APIAccessKeyType
}{
	// An ingest key is used by New Relic agents to authenticate with New Relic and send data to the assigned account.
	INGEST: "INGEST",
	// A user key is used by New Relic users to authenticate with New Relic and to interact with the New Relic GraphQL API .
	USER: "USER",
}

// APIAccessUserKeyErrorType - The type of error.
type APIAccessUserKeyErrorType string

var APIAccessUserKeyErrorTypeTypes = struct {
	// Occurs when the user issuing the mutation does not have sufficient permissions to perform the action for a key.
	FORBIDDEN APIAccessUserKeyErrorType
	// Occurs when the action taken on a key did not successfully pass validation.
	INVALID APIAccessUserKeyErrorType
	// Occurs when the requested key `id` was not found.
	NOT_FOUND APIAccessUserKeyErrorType
}{
	// Occurs when the user issuing the mutation does not have sufficient permissions to perform the action for a key.
	FORBIDDEN: "FORBIDDEN",
	// Occurs when the action taken on a key did not successfully pass validation.
	INVALID: "INVALID",
	// Occurs when the requested key `id` was not found.
	NOT_FOUND: "NOT_FOUND",
}

// APIAccessActorStitchedFields -
type APIAccessActorStitchedFields struct {
	// Fetch a single key by ID and type.
	Key APIAccessKey `json:"key"`
	// A list of keys scoped to the current actor and filter arguments.
	KeySearch APIAccessKeySearchResult `json:"keySearch"`
}

// APIAccessCreateIngestKeyInput - The input for any ingest keys you want to create. Each ingest key must have a type that communicates what kind of data it is for. You can optionally add a name or notes to your key, which can be updated later.
type APIAccessCreateIngestKeyInput struct {
	// The account ID indicating which account you want to make the key for. This cannot be updated once created.
	AccountID int `json:"accountId"`
	// The type of ingest key you want to create. This cannot be updated once created.
	IngestType APIAccessIngestKeyType `json:"ingestType"`
	// The name of the key. This can be updated later.
	Name string `json:"name"`
	// Any notes about this ingest key. This can be updated later.
	Notes string `json:"notes"`
}

// APIAccessCreateInput - The input object to create one or more keys.
type APIAccessCreateInput struct {
	// Ingest keys are used by agents to report data about your applications to New Relic. Each ingest key input entered here must have a type that communicates what kind of data it is for. You can optionally add a name or notes to your key, which can be updated later.
	Ingest []APIAccessCreateIngestKeyInput `json:"ingest"`
	// Create user keys. You can optionally add a name or notes to your key, which can be updated later.
	User []APIAccessCreateUserKeyInput `json:"user"`
}

// APIAccessCreateKeyResponse - The response of the create keys mutation.
type APIAccessCreateKeyResponse struct {
	// Lists all successfully created keys.
	CreatedKeys []APIAccessKey `json:"createdKeys"`
	// Lists all errors for keys that could not be created. Each error maps to a single key input.
	Errors []APIAccessKeyError `json:"errors"`
}

// APIAccessCreateUserKeyInput - The input for any ingest keys you want to create. Each ingest key must have a type that communicates what kind of data it is for. You can optionally add a name or notes to your key, which can be updated later.
type APIAccessCreateUserKeyInput struct {
	// The account ID indicating which account you want to make the key for. This cannot be updated once created.
	AccountID int `json:"accountId"`
	// The name of the key. This can be updated later.
	Name string `json:"name"`
	// Any notes about this ingest key. This can be updated later.
	Notes string `json:"notes"`
	// The user ID indicating which user you want to make the key for. This cannot be updated once created.
	UserID int `json:"userId"`
}

// APIAccessDeleteInput - The input to delete keys.
type APIAccessDeleteInput struct {
	// A list of the ingest key `id`s that you want to delete.
	IngestKeyIDs []string `json:"ingestKeyIds"`
	// A list of the user key `id`s that you want to delete.
	UserKeyIDs []string `json:"userKeyIds"`
}

// APIAccessDeleteKeyResponse - The response of the key delete mutation.
type APIAccessDeleteKeyResponse struct {
	// The `id`s of the successfully deleted ingest keys and any errors that occurred when deleting keys.
	DeletedKeys []APIAccessDeletedKey `json:"deletedKeys"`
	// Lists all errors for keys that could not be deleted. Each error maps to a single key input.
	Errors []APIAccessKeyError `json:"errors"`
}

// APIAccessDeletedKey - The deleted key response of the key delete mutation.
type APIAccessDeletedKey struct {
	// The `id` of the deleted key.
	ID string `json:"id"`
}

// APIAccessKey - A key for accessing New Relic APIs.
type APIAccessKey struct {
	// The ID of the key. This can be used to identify a key without revealing the key itself (used to update and delete).
	ID string `json:"id"`
	// The keystring of the key.
	Key string `json:"key"`
	// The name of the key. This can be used a short identifier for easy reference.
	Name string `json:"name"`
	// Any notes can be attached to a key. This is intended for more a more detailed description of the key use if desired.
	Notes string `json:"notes"`
	// The type of key, indicating what New Relic APIs it can be used to access.
	Type APIAccessKeyType `json:"type"`
}

func (x *APIAccessKey) ImplementsAPIAccessKey() {}

// APIAccessKeyError - A key error. Each error maps to a single key input.
type APIAccessKeyError struct {
	// A message about why the key creation failed.
	Message string `json:"message"`
	// The type of the key.
	Type APIAccessKeyType `json:"type"`
}

func (x *APIAccessKeyError) ImplementsAPIAccessKeyError() {}

// APIAccessKeySearchQuery - Parameters by which to filter the search.
type APIAccessKeySearchQuery struct {
	// Criteria by which to narrow the scope of keys to be returned.
	Scope APIAccessKeySearchScope `json:"scope"`
	// A list of key types to be included in the search. If no types are provided, all types will be returned by default.
	Types []APIAccessKeyType `json:"types"`
}

// APIAccessKeySearchResult - A list of all keys scoped to the current actor.
type APIAccessKeySearchResult struct {
	// A list of all keys scoped to the current actor.
	Keys []APIAccessKey `json:"keys"`
	// The next cursor, used for pagination. If a cursor is present, it means more keys can be fetched.
	NextCursor string `json:"nextCursor"`
}

// APIAccessKeySearchScope - The scope of keys to be returned. Note that some filters only apply to certain key types.
type APIAccessKeySearchScope struct {
	// A list of key account IDs.
	AccountIDs []int `json:"accountIds"`
	// The ingest type of the key. Only applies to ingest keys, and does not affect user key filtering.
	IngestTypes []APIAccessIngestKeyType `json:"ingestTypes"`
	// A list of key user ids. Only applies to user keys, and does not affect ingest key filtering.
	UserIDs []int `json:"userIds"`
}

// APIAccessUpdateIngestKeyInput - The `id` and data to update one or more keys.
type APIAccessUpdateIngestKeyInput struct {
	// The `id` of the key you want to update.
	KeyID string `json:"keyId"`
	// The name you want to assign to the key.
	Name string `json:"name"`
	// The notes you want to assign to the key.
	Notes string `json:"notes"`
}

// APIAccessUpdateInput - The `id` and data to update one or more keys.
type APIAccessUpdateInput struct {
	// A list of the configurations of each ingest key you want to update.
	Ingest []APIAccessUpdateIngestKeyInput `json:"ingest"`
	// A list of the configurations of each user key you want to update.
	User []APIAccessUpdateUserKeyInput `json:"user"`
}

// APIAccessUpdateKeyResponse - The response of the update keys mutation.
type APIAccessUpdateKeyResponse struct {
	// Lists all errors for keys that could not be updated. Each error maps to a single key input.
	Errors []APIAccessKeyError `json:"errors"`
	// Lists all successfully updated keys.
	UpdatedKeys []APIAccessKey `json:"updatedKeys"`
}

// APIAccessUpdateUserKeyInput - The `id` and data to update one or more keys.
type APIAccessUpdateUserKeyInput struct {
	// The `id` of the key you want to update.
	KeyID string `json:"keyId"`
	// The name you want to assign to the key.
	Name string `json:"name"`
	// The notes you want to assign to the key.
	Notes string `json:"notes"`
}

// APIAccessKeyError - A key error. Each error maps to a single key input.
type APIAccessKeyErrorInterface interface {
	ImplementsAPIAccessKeyErrorInterface()
}

// APIAccessKey - A key for accessing New Relic APIs.
type APIAccessKeyInterface interface {
	ImplementsAPIAccessKeyInterface()
}