package transmission

import "errors"

const (
	CMD_TORRENT_REMOVE = "torrent-remove"
)

type RemoveTorrentParams struct {
	IDs        []interface{} `json:"ids"`
	DeleteData bool          `json:"delete-local-data"`
}

func (t *Client) Remove(which WhichTorrents, deleteData bool) error {

	cmd, err := createRemoveCommand(which, deleteData)
	if err != nil {
		return err
	}

	if _, err := t.execute(cmd); err != nil {
		return err
	}

	return nil
}

func createRemoveCommand(which WhichTorrents, deleteData bool) (*Command, error) {
	if which.Recent {
		return nil, errors.New(ErrRecentNotValid)
	}

	params := RemoveTorrentParams{
		DeleteData: deleteData,
	}

	if len(which.IDs) == 0 && len(which.SHAs) == 0 {
		params.IDs = []interface{}{}
		return &Command{
			Method:    CMD_TORRENT_REMOVE,
			Arguments: params,
		}, nil
	}

	params.IDs = mergeIDsAndSHAs(which)

	return &Command{
		Method:    CMD_TORRENT_REMOVE,
		Arguments: params,
	}, nil
}
