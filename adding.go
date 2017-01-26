package transmission

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

const (
	CMD_TORRENT_ADD = "torrent-add"
)

type AddParams struct {
	DownloadDir    string `json:"download-dir,omitempty"`
	FilesUnwanted  []int  `json:"files-unwanted",omitempty`
	FilesWanted    []int  `json:"files-wanted,omitempty"`
	Paused         bool   `json:"paused,omitempty"`
	PeerLimit      int    `json:"peer-limit,omitempty"`
	PriorityHigh   []int  `json:"priority-high,omitempty"`
	PriorityLow    []int  `json:"priority-low,omitempty"`
	PriorityNormal []int  `json:"priority-normal,omitempty"`
}

type AddResponse struct {
	Torrent AddedTorrent `json:"torrent-added"`
}

type AddedTorrent struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Hashstring string `json:"hashString"`
}

type addClientParams struct {
	DownloadDir    string `json:"download-dir,omitempty"`
	Filename       string `json:"filename,omitempty"`
	FilesUnwanted  []int  `json:"files-unwanted",omitempty`
	FilesWanted    []int  `json:"files-wanted,omitempty"`
	Metainfo       string `json:"metainfo,omitempty"`
	Paused         bool   `json:"paused,omitempty"`
	PeerLimit      int    `json:"peer-limit,omitempty"`
	PriorityHigh   []int  `json:"priority-high,omitempty"`
	PriorityLow    []int  `json:"priority-low,omitempty"`
	PriorityNormal []int  `json:"priority-normal,omitempty"`
}

func (t *Client) AddFromURL(url string, params AddParams) (AddResponse, error) {
	var resp AddResponse

	cmd, err := createAddCommand(addClientParams{
		DownloadDir:    params.DownloadDir,
		Filename:       url,
		FilesUnwanted:  params.FilesUnwanted,
		FilesWanted:    params.FilesWanted,
		Paused:         params.Paused,
		PeerLimit:      params.PeerLimit,
		PriorityHigh:   params.PriorityHigh,
		PriorityNormal: params.PriorityNormal,
		PriorityLow:    params.PriorityLow,
	})
	if err != nil {
		return resp, err
	}

	return t.sendAddCommand(cmd)
}

func (t *Client) AddFromPath(path string, params AddParams) (AddResponse, error) {
	var resp AddResponse

	fileData, err := ioutil.ReadFile(path)
	if err != nil {
		return resp, err
	}

	cmd, err := createAddCommand(addClientParams{
		DownloadDir:    params.DownloadDir,
		FilesUnwanted:  params.FilesUnwanted,
		FilesWanted:    params.FilesWanted,
		Metainfo:       base64.StdEncoding.EncodeToString(fileData),
		Paused:         params.Paused,
		PeerLimit:      params.PeerLimit,
		PriorityHigh:   params.PriorityHigh,
		PriorityNormal: params.PriorityNormal,
		PriorityLow:    params.PriorityLow,
	})
	if err != nil {
		return resp, err
	}

	return t.sendAddCommand(cmd)
}

func (t *Client) sendAddCommand(cmd *Command) (AddResponse, error) {
	body, err := t.execute(cmd)
	if err != nil {
		return AddResponse{}, err
	}
	return translateResponse(body)
}

func translateResponse(body []byte) (AddResponse, error) {
	var resp AddResponse
	fmt.Printf("translating: %v \n", string(body))
	if err := json.Unmarshal(body, &resp); err != nil {
		return AddResponse{}, err
	}
	return resp, nil
}

func createAddCommand(params addClientParams) (*Command, error) {

	return &Command{
		Method:    CMD_TORRENT_ADD,
		Arguments: params,
	}, nil

	return nil, nil
}
