package transmission

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

const (
	POST = "POST"

	ErrUninitialized = "not yet initialized"
)

func (t *Client) getSessionHeader() error {

	req, err := http.NewRequest(POST, t.Address, bytes.NewReader([]byte{}))
	if err != nil {
		return err
	}

	req.SetBasicAuth(t.Username, t.Password)

	if t.client == nil {
		t.client = &http.Client{}
	}

	res, err := t.client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	t.token = res.Header.Get(SESSION_HEADER)

	return nil
}

func (t *Client) authRequest(req *http.Request) error {
	if t.token == "" {
		if err := t.getSessionHeader(); err != nil {
			return err
		}
	}

	req.Header.Add(SESSION_HEADER, t.token)
	req.SetBasicAuth(t.Username, t.Password)

	return nil
}

func (t *Client) execute(cmd *Command) ([]byte, error) {
	if t == nil {
		return nil, errors.New(ErrUninitialized)
	}

	body, err := json.Marshal(cmd)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(POST, t.Address, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	if err := t.authRequest(req); err != nil {
		return nil, err
	}

	res, err := t.client.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode == http.StatusConflict {
		res.Body.Close()
		t.token = ""

		req, err := http.NewRequest(POST, t.Address, bytes.NewReader(body))
		if err := t.authRequest(req); err != nil {
			return nil, err
		}

		res, err = t.client.Do(req)
		if err != nil {
			return nil, err
		}
	}
	defer res.Body.Close()

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var resp Response
	if err = json.Unmarshal(resBody, &resp); err != nil {
		return nil, err
	}

	return resp.Arguments, nil
}
