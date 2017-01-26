package transmission

import "errors"

const (
	ErrRecentIdMix = "invalid use of WhichTorrents. dont mix IDs\\SHAs & Recent."

	IDS             = "ids"
	RECENTLY_ACTIVE = "recently-active"

	CMD_TORRENT_START      = "torrent-start"
	CMD_TORRENT_STOP       = "torrent-stop"
	CMD_TORRENT_VERIFY     = "torrent-verify"
	CMD_TORRENT_REANNOUNCE = "torrent-reannounce"
)

func (t *Client) Start(which WhichTorrents) error {

	cmd, err := createIDOnlyCommand(CMD_TORRENT_START, which)
	if err != nil {
		return err
	}

	if _, err = t.execute(cmd); err != nil {
		return err
	}

	return nil
}

func (t *Client) Stop(which WhichTorrents) error {

	cmd, err := createIDOnlyCommand(CMD_TORRENT_STOP, which)
	if err != nil {
		return err
	}

	if _, err = t.execute(cmd); err != nil {
		return err
	}

	return nil
}

func (t *Client) Verify(which WhichTorrents) error {

	cmd, err := createIDOnlyCommand(CMD_TORRENT_VERIFY, which)
	if err != nil {
		return err
	}

	if _, err = t.execute(cmd); err != nil {
		return err
	}

	return nil
}

func (t *Client) Reannounce(which WhichTorrents) error {

	cmd, err := createIDOnlyCommand(CMD_TORRENT_REANNOUNCE, which)
	if err != nil {
		return err
	}

	if _, err = t.execute(cmd); err != nil {
		return err
	}

	return nil
}

func createIDOnlyCommand(command string, which WhichTorrents) (*Command, error) {
	if which.Recent && (len(which.IDs) > 0 || len(which.SHAs) > 0) {
		return nil, errors.New(ErrRecentIdMix)
	}

	if !which.Recent && len(which.IDs) == 0 && len(which.SHAs) == 0 {
		return &Command{
			Method:    command,
			Arguments: map[int]int{},
		}, nil
	}

	arguments := map[string]interface{}{}

	if which.Recent {
		arguments[IDS] = RECENTLY_ACTIVE
	} else {
		arguments[IDS] = mergeIDsAndSHAs(which)
	}

	return &Command{
		Method:    command,
		Arguments: arguments,
	}, nil
}
