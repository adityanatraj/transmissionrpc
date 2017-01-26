package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strings"

	"io"

	log "github.com/Sirupsen/logrus"
	"github.com/adityanatraj/transmission"
)

const (
	ErrCLIComplete = "received exit line"
)

func init() {
	log.SetLevel(log.DebugLevel)
}

// CLI is a single "session" of talking to transmission
type CLI struct {
	env          map[string]interface{}
	lineNumber   int
	transmission *transmission.Client
}

const (
	cmdExit    = "exit"
	cmdConnect = "connect"
	cmdState   = "state"

	cmdStart = "start"
	cmdStop  = "stop"
	cmdAdd   = "add"
	cmdMove  = "move"
)

func (cli *CLI) Connect(host string) {
	log.Debugf("connecting to: %s", host)
	cli.transmission = &transmission.Client{
		Address: transmission.AddressFromHost(host),
	}
}

func (cli CLI) isConnected() bool {
	return cli.transmission != nil
}

func (cli *CLI) Start(hash ...string) {
	log.Debugf("starting: %v", hash)
	fmt.Println(cli.transmission.Start(transmission.WhichTorrents{
		SHAs: hash,
	}))
}

func (cli *CLI) Stop(hash ...string) {
	log.Debugf("stopping: %v", hash)
	fmt.Println(cli.transmission.Stop(transmission.WhichTorrents{
		SHAs: hash,
	}))
}

func (cli *CLI) Add(url string) {
	log.Debugf("adding: %v", url)
	fmt.Println(cli.transmission.AddFromURL(url, transmission.AddParams{
		DownloadDir: "/media/anatraj/Monopoly/Downloads",
		Paused:      true,
	}))
}

func (cli *CLI) State() {
	fmt.Println(cli)
}

var exitExpression = regexp.MustCompile(cmdExit)

// ReadLines reads as many lines as it can until
func (cli *CLI) ReadLines(instream io.Reader) error {
	fmt.Print("$ ")
	scanner := bufio.NewScanner(instream)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")

		cmd, args := line[0], line[1:]
		log.Debugf("cmd: %s args: %v", cmd, args)
		switch cmd {
		case cmdExit:
			return errors.New(ErrCLIComplete)
		case cmdConnect:
			cli.Connect(args[0])
		default:
			if !cli.isConnected() {
				fmt.Println("must be connected. run `connect <host>`")
				continue
			}
			switch cmd {
			case cmdStart:
				cli.Start(args...)
			case cmdStop:
				cli.Stop(args...)
			case cmdAdd:
				if len(args) > 1 {
					fmt.Println("at least 1 and at most 1 url")
				}
				cli.Add(args[0])
			case cmdState:
				cli.State()
			default:
				fmt.Println("unknown command")
			}
		}
		fmt.Print("$ ")
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func main() {
	cli := CLI{
		transmission: &transmission.Client{
			Address: "thinkpad.local",
		},
	}
	cli.ReadLines(os.Stdin)
}
