package todoist

import (
	"bytes"
	"encoding/json"
	"math/rand"
	"net/http"
	"time"
)

const apiURL = "https://todoist.com/api/v8/sync"

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// Client is a client for sending HTTP requests,
// and has a required token and id at the time of request.
type Client struct {
	httpClient http.Client
	token      string
	id         string
}

// Project is an object returned as a response from Todoist.
type Project struct {
	IsArchived   int    `json:"is_archived,omitempty"`
	ParentID     int    `json:"parent_id"`
	Color        int    `json:"color,omitempty"`
	Shared       bool   `json:"shared,omitempty"`
	InboxProject bool   `json:"inbox_project,omitempty"`
	ID           int    `json:"id,omitempty"`
	Collapsed    int    `json:"collapsed,omitempty"`
	ItemOrder    int    `json:"item_order,omitempty"`
	Name         string `json:"name,omitempty"`
	IsDeleted    int    `json:"is_deleted,omitempty"`
	Indent       int    `json:"indent,omitempty"`
}

// Projects is a list of project returned as a response from Todoist.
type Projects []Project

// Due is an object returned as a response from Todoist.
type Due struct {
	Date        string      `json:"date"`
	Timezone    interface{} `json:"timezone"`
	IsRecurring bool        `json:"is_recurring"`
	String      string      `json:"string"`
	Lang        string      `json:"lang"`
}

// Item is an object returned as a response from Todoist.
type Item struct {
	Collapsed      int           `json:"collapsed,omitempty"`
	DateAdded      time.Time     `json:"date_added,omitempty"`
	ItemOrder      int           `json:"item_order,omitempty"`
	IsArchived     int           `json:"is_archived,omitempty"`
	Indent         int           `json:"indent,omitempty"`
	AllDay         bool          `json:"all_day,omitempty"`
	DayOrder       int           `json:"day_order,omitempty"`
	AssignedByUID  int           `json:"assigned_by_uid,omitempty"`
	ResponsibleUID interface{}   `json:"responsible_uid,omitempty"`
	SyncID         interface{}   `json:"sync_id,omitempty"`
	Checked        int           `json:"checked,omitempty"`
	UserID         int           `json:"user_id,omitempty"`
	Labels         []interface{} `json:"labels,omitempty"`
	IsDeleted      int           `json:"is_deleted,omitempty"`
	Due            Due           `json:"due,omitempty"`
	ProjectID      int           `json:"project_id,omitempty"`
	InHistory      int           `json:"in_history,omitempty"`
	Content        string        `json:"content,omitempty"`
	ID             int           `json:"id,omitempty"`
	Priority       int           `json:"priority,omitempty"`
}

// Items is a list of item returned as a response from Todoist.
type Items []Item

// MapStatus is an object returnd as a response from Todoist.
// This is returned when the request command is invalid.
// And return string when the request command is valid.
type MapStatus struct {
	ErrorTag   string `json:"error_tag,omitempty"`
	ErrorCode  int    `json:"error_code,omitempty"`
	HTTPCode   int    `json:"http_code,omitempty"`
	ErrorExtra struct {
	} `json:"error_extra,omitempty"`
	Error string `json:"error,omitempty"`
}

// APIResponse is an object returnd as a response from Todoist.
type APIResponse struct {
	SyncStatus    interface{}      `json:"sync_status,omitempty"`
	TempIDMapping map[string]int64 `json:"temp_id_mapping,omitempty"`
	FullSync      bool             `json:"full_sync,omitempty"`
	Projects      Projects         `json:"projects,omitempty"`
	Items         Items            `json:"items,omitempty"`
	SyncToken     string           `json:"sync_token,omitempty"`
}

// NewClient returns a Client, which has an HTTP client and a token set.
// Token value to be set is passed as an argument.
func NewClient(token string) Client {
	return Client{
		httpClient: http.Client{},
		token:      token,
	}
}

// SetNewID sets a random string as an ID for Client of the receiver.
func (c Client) SetNewID() Client {
	rand.Seed(time.Now().UnixNano())
	c.id = randStringRunes(99)
	return c
}

// Req is responsible for processing the request.
// Actually, based on the received arguments, create a command,
// create a payload, create a request, and call a method to make a request.
func (c Client) Req(commandType string, args map[string]interface{}) (APIResponse, error) {
	var command Command
	var payload string

	if args != nil {
		command = newCommand(commandType, c.id, commandArgs(args))
		payload = newPayload(c.token, payloadCommands(Commands{command}))
	} else {
		payload = newPayload(c.token)
	}

	req := newReq(bytes.NewBufferString(payload))
	resp, _ := c.httpClient.Do(req)
	decodeResp, err := responseUnmarshal(resp)
	return decodeResp, err
}

// newReq returns a Request, which has Content-Type header set.
func newReq(body *bytes.Buffer) *http.Request {
	req, _ := http.NewRequest("POST", apiURL, body)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	return req
}

// randStringRunes returns a string, which is random.
func randStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

// responseUnmarshal returns APIResponsem, which has decoded http response.
func responseUnmarshal(resp *http.Response) (APIResponse, error) {
	result := APIResponse{}
	err := json.NewDecoder(resp.Body).Decode(&result)
	switch convValue := result.SyncStatus.(type) {
	case string:
		result.SyncStatus = convValue
	case MapStatus:
		result.SyncStatus = convValue
	}
	return result, err
}
