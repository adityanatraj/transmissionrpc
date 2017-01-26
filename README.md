# Transmission RPC Client

[Transmission](transmissionbt.com) is a great bittorrent client to run as a headless appliance. 

This library implements the RPC API provided by transmission server. While there
are a few other libraries to do this ([go-transmission](https://github.com/lnguyen/go-transmission), [transmission](https://github.com/Tubbebubbe/transmission), etc), this one attempts
to provide a full interface to the API. 


## Quickstart

```golang
import "transmission"
import "fmt"

func main() {
  t := transmission.Client{
    Address:  transmission.AddressFromHost("localhost"),
  }

  // start ALL torrents
  if err := t.Start(transmission.WhichTorrents{}); err != nil {
    fmt.Println(err)
    return
  }

  // remove the torrent with ID: 11 AND delete the data
  deleteMe := transmission.WhichTorrents{IDs: []int{11}}
  if err := t.Remove(deleteMe, true); err != nil {
    fmt.Println(err)
    return
  }

}
```

## Yet Incomplete

01/25: currently fixing an issue with add and torrent-duplicate behavior.

Relative to the [available RPC spec](https://trac.transmissionbt.com/browser/branches/1.7x/doc/rpc-spec.txt?order=name), the current progress is:
- 3.1 Torrent Actions {Start, Stop, Verify, Reannounce} - done
- 3.2 Torrent Mutation {}
- 3.3 Torrent Access {}
- 3.4 Torrent Add {AddFromPath, AddFromURL} - done
- 3.5 Torrent Remove {Remove} - done
- 3.6 Torrent Move {Move, SetLocation} - done
- 4.1.1 Session Mutation {}
- 4.1.2 Session Access {}
- 4.2 Session Stats {}
- 4.3 Blocklist {}
- 4.4 Port Check {}

## License:

The MIT License (MIT)
Copyright (c) 2016 adityanatraj

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
