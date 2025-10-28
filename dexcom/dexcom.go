package dexcom

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type Dexcom struct {
	client    *http.Client
	baseURL   string
	appID     string
	username  string
	password  string
	accountID string
	sessionID string
}

func NewDexcom(username, password string, region Region) (*Dexcom, error) {
	baseURL, ok := DEXCOM_BASE_URLS[region]
	if !ok {
		return nil, ErrInvalidRegion
	}

	d := &Dexcom{
		client:   &http.Client{},
		baseURL:  baseURL,
		appID:    DEXCOM_APPLICATION_IDS[region],
		username: username,
		password: password,
	}

	if err := d.authenticate(); err != nil {
		return nil, err
	}
	return d, nil
}

func (d *Dexcom) post(endpoint string, params url.Values, body any) (json.RawMessage, error) {
	u := d.baseURL + endpoint
	if params != nil {
		u += "?" + params.Encode()
	}

	var bodyReader io.Reader
	if body != nil {
		data, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		bodyReader = bytes.NewReader(data)
	}

	req, err := http.NewRequest("POST", u, bodyReader)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := d.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result json.RawMessage
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	// Check for error response
	if resp.StatusCode >= 400 {
		var errResp struct {
			Code    string
			Message string
		}
		json.Unmarshal(result, &errResp)
		return nil, d.mapError(errResp.Code, errResp.Message)
	}

	return result, nil
}

func (d *Dexcom) mapError(code, message string) error {
	switch code {
	case "SessionIdNotFound", "SessionNotValid":
		return ErrSessionExpired
	case "AccountPasswordInvalid":
		return ErrAccountAuthFailed
	case "SSO_AuthenticateMaxAttemptsExceeded":
		return ErrAccountMaxAttempts
	default:
		return fmt.Errorf("%s: %s", code, message)
	}
}

func (d *Dexcom) authenticate() error {
	// Get account ID
	result, err := d.post(DEXCOM_AUTHENTICATE_ENDPOINT, nil, map[string]string{
		"accountName":   d.username,
		"password":      d.password,
		"applicationId": d.appID,
	})
	if err != nil {
		return err
	}
	json.Unmarshal(result, &d.accountID)
	d.accountID = strings.Trim(d.accountID, `"`) // Handle both string and object response

	// Get session ID
	result, err = d.post(DEXCOM_LOGIN_ID_ENDPOINT, nil, map[string]string{
		"accountId":     d.accountID,
		"password":      d.password,
		"applicationId": d.appID,
	})
	if err != nil {
		return err
	}
	json.Unmarshal(result, &d.sessionID)
	d.sessionID = strings.Trim(d.sessionID, `"`)

	return nil
}
