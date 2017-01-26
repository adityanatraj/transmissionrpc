package main

import (
	"fmt"

	"github.com/adityanatraj/transmission"
)

func main() {

	t := transmission.Client{
		Address: transmission.AddressFromHost("192.168.0.103"),
	}

	fmt.Println(t)

	// minka := transmission.WhichTorrents{SHAs: []string{"06c5ec4a92221b6d691122e816ebfd4975a51f66"}}
	// if err := t.Remove(minka, false); err != nil {
	// fmt.Println(err)
	// }

	// url := "https://passthepopcorn.me/torrents.php?action=download&id=469029&authkey=fa7b212371a0c99612ee928628bdd471&torrent_pass=onepnhzs0s8mqe11kbx2qwkdmjiex5nu"
	// added, err := t.AddFromURL(url, transmission.AddParams{
	// 	DownloadDir: "/media/anatraj/Monopoly/Media/HD_Movies/",
	// 	Paused:      true,
	// })

	// added, err := t.AddFromPath("minka.torrent", transmission.AddParams{
	// 	DownloadDir: "/media/anatraj/Monopoly/Media/Music/",
	// 	Paused:      true,
	// })

	beasts := transmission.WhichTorrents{
		SHAs: []string{"8cf6d7ea4f7b4ac79d1e6042352910de1b3cba22"},
	}

	err := t.SetLocation(beasts, "/media/anatraj/Monopoly/Media/HD_Movies/")
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}

	// fmt.Println(added)

	fmt.Println(t)

}
