package transmission

import "errors"

const (
	ErrRecentNotValid = "command does not allow use of recently-active"

	CMD_TORRENT_SET_LOCATION = "torrent-set-location"
)

type SetLocationParams struct {
	IDs      []interface{} `json:"ids"`
	Location string        `json:"location"`
	Move     bool          `json:"move"`
}

func (t *Client) Move(which WhichTorrents, location string) error {
	cmd, err := createMoveCommand(which, location, true)
	if err != nil {
		return err
	}

	if _, err = t.execute(cmd); err != nil {
		return err
	}

	return nil
}

func (t *Client) SetLocation(which WhichTorrents, location string) error {
	cmd, err := createMoveCommand(which, location, false)
	if err != nil {
		return err
	}

	if _, err = t.execute(cmd); err != nil {
		return err
	}

	return nil
}

func createMoveCommand(which WhichTorrents, location string, move bool) (*Command, error) {

	if which.Recent {
		return nil, errors.New(ErrRecentNotValid)
	}

	params := SetLocationParams{
		Location: location,
		Move:     move,
	}

	if len(which.IDs) == 0 && len(which.SHAs) == 0 {
		params.IDs = []interface{}{}
		return &Command{
			Method:    CMD_TORRENT_SET_LOCATION,
			Arguments: params,
		}, nil
	}

	params.IDs = mergeIDsAndSHAs(which)

	return &Command{
		Method:    CMD_TORRENT_SET_LOCATION,
		Arguments: params,
	}, nil

}
