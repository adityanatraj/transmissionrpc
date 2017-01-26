package transmission

type Torrent struct {
	ActivityDate      int    `json:"activityDate,omitempty"`
	AddedDate         int    `json:"addedDate,omitempty"`
	AnnounceResponse  string `json:"announceResponse,omitempty"`
	AnnounceURL       string `json:"announceURL,omitempty"`
	BandwidthPriority int    `json:"bandwidthPriority,omitempty"`
	Comment           string `json:"comment,omitempty"`
	CorruptEver       int    `json:"corruptEver,omitempty"`
	Creator           string `json:"creator,omitempty"`
	DateCreated       int    `json:"dateCreated,omitempty"`
	DesiredAvailable  int    `json:"desiredAvailable,omitempty"`
	DoneDate          int    `json:"doneDate,omitempty"`
	DownloadDir       string `json:"downloadDir,omitempty"`
	DownloadLimit     int    `json:"downloadLimit,omitempty"`
	DownloadLimited   bool   `json:"downloadLimited,omitempty"`
	DownloadedEver    int    `json:"downloadedEver,omitempty"`
	Downloaders       int    `json:"downloaders,omitempty"`
	Error             int    `json:"error,omitempty"`
	ErrorString       string `json:"errorString,omitempty"`
	Eta               int    `json:"eta,omitempty"`
	FileStats         []struct {
		BytesCompleted int  `json:"bytesCompleted,omitempty"`
		Priority       int  `json:"priority,omitempty"`
		Wanted         bool `json:"wanted,omitempty"`
	} `json:"fileStats"`
	Files []struct {
		BytesCompleted int    `json:"bytesCompleted,omitempty"`
		Key            string `json:"key,omitempty"`
		Length         int    `json:"length,omitempty"`
		Name           string `json:"name,omitempty"`
	} `json:"files"`
	HashString          string `json:"hashString,omitempty"`
	HaveUnchecked       int    `json:"haveUnchecked,omitempty"`
	HaveValid           int    `json:"haveValid,omitempty"`
	HonorsSessionLimits bool   `json:"honorsSessionLimits,omitempty"`
	ID                  int    `json:"id,omitempty"`
	IsPrivate           bool   `json:"isPrivate,omitempty"`
	LastAnnounceTime    int    `json:"lastAnnounceTime,omitempty"`
	LastScrapeTime      int    `json:"lastScrapeTime,omitempty"`
	Leechers            int    `json:"leechers,omitempty"`
	LeftUntilDone       int    `json:"leftUntilDone,omitempty"`
	ManualAnnounceTime  int    `json:"manualAnnounceTime,omitempty"`
	MaxConnectedPeers   int    `json:"maxConnectedPeers,omitempty"`
	Name                string `json:"name,omitempty"`
	NextAnnounceTime    int    `json:"nextAnnounceTime,omitempty"`
	NextScrapeTime      int    `json:"nextScrapeTime,omitempty"`
	PeerLimit           int    `json:"peer-limit,omitempty"`
	Peers               []struct {
		Address            string  `json:"address,omitempty"`
		ClientIsChoked     bool    `json:"clientIsChoked,omitempty"`
		ClientIsInterested bool    `json:"clientIsInterested,omitempty"`
		ClientName         string  `json:"clientName,omitempty"`
		IsDownloadingFrom  bool    `json:"isDownloadingFrom,omitempty"`
		IsEncrypted        bool    `json:"isEncrypted,omitempty"`
		IsIncoming         bool    `json:"isIncoming,omitempty"`
		IsUploadingTo      bool    `json:"isUploadingTo,omitempty"`
		PeerIsChoked       bool    `json:"peerIsChoked,omitempty"`
		PeerIsInterested   bool    `json:"peerIsInterested,omitempty"`
		Port               int     `json:"port,omitempty"`
		Progress           float64 `json:"progress,omitempty"`
		RateToClient       int     `json:"rateToClient,omitempty"`
		RateToPeer         int     `json:"rateToPeer,omitempty"`
	} `json:"peers"`
	PeersConnected int `json:"peersConnected,omitempty"`
	PeersFrom      struct {
		FromCache    int `json:"fromCache,omitempty"`
		FromIncoming int `json:"fromIncoming,omitempty"`
		FromPex      int `json:"fromPex,omitempty"`
		FromTracker  int `json:"fromTracker,omitempty"`
	} `json:"peersFrom"`
	PeersGettingFromUs int     `json:"peersGettingFromUs,omitempty"`
	PeersKnown         int     `json:"peersKnown,omitempty"`
	PeersSendingToUs   int     `json:"peersSendingToUs,omitempty"`
	PercentDone        float64 `json:"percentDone,omitempty"`
	PieceCount         int     `json:"pieceCount,omitempty"`
	PieceSize          int     `json:"pieceSize,omitempty"`
	// base64-encoded string since JSON can't hold binary strings
	Pieces          string  `json:"pieces,omitempty"`
	Priorities      []int   `json:"priorities,omitempty"`
	RateDownload    int     `json:"rateDownload,omitempty"`
	RateUpload      int     `json:"rateUpload,omitempty"`
	RecheckProgress float64 `json:"recheckProgress,omitempty"`
	ScrapeResponse  string  `json:"scrapeResponse,omitempty"`
	ScrapeURL       string  `json:"scrapeURL,omitempty"`
	SeedRatioLimit  float64 `json:"seedRatioLimit,omitempty"`
	SeedRatioMode   int     `json:"seedRatioMode,omitempty"`
	Seeders         int     `json:"seeders,omitempty"`
	SizeWhenDone    int     `json:"sizeWhenDone,omitempty"`
	StartDate       int     `json:"startDate,omitempty"`
	Status          int     `json:"status,omitempty"`
	SwarmSpeed      int     `json:"swarmSpeed,omitempty"`
	TimesCompleted  int     `json:"timesCompleted,omitempty"`
	TorrentFile     string  `json:"torrentFile,omitempty"`
	TotalSize       int     `json:"totalSize,omitempty"`
	Trackers        []struct {
		Announce string `json:"announce,omitempty"`
		Scrape   string `json:"scrape,omitempty"`
		Tier     int    `json:"tier,omitempty"`
	} `json:"trackers"`
	UploadLimit         int      `json:"uploadLimit,omitempty"`
	UploadLimited       bool     `json:"uploadLimited,omitempty"`
	UploadRatio         float64  `json:"uploadRatio,omitempty"`
	UploadedEver        int      `json:"uploadedEver,omitempty"`
	Wanted              []bool   `json:"wanted,omitempty"`
	Webseeds            []string `json:"webseeds,omitempty"`
	WebseedsSendingToUs int      `json:"webseedsSendingToUs,omitempty"`
}

type SetTorrentParams struct {
	BandwidthPriority   int      `json:"bandwidthPriority,omitempty"`
	DownloadLimit       int      `json:"downloadLimit,omitempty"`
	DownloadLimited     bool     `json:"downloadLimited"`
	Files_unwanted      []int    `json:"files-unwanted,omitempty"`
	Files_wanted        []int    `json:"files-wanted,omitempty"`
	HonorsSessionLimits bool     `json:"honorsSessionLimits,omitempty"`
	IDs                 []string `json:"ids,omitempty"`
	Location            string   `json:"location,omitempty"`
	PeerLimit           int      `json:"peer-limit,omitempty"`
	PriorityHigh        []int    `json:"priority-high,omitempty"`
	PriorityLow         []int    `json:"priority-low,omitempty"`
	PriorityNormal      []int    `json:"priority-normal,omitempty"`
	SeedRatioLimit      float64  `json:"seedRatioLimit,omitempty"`
	SeedRatioMode       int      `json:"seedRatioMode,omitempty"`
	UploadLimit         int      `json:"uploadLimit,omitempty"`
	UploadLimited       bool     `json:"uploadLimited,omitempty"`
}

type GetTorrentParams struct {
	IDs    []string `json:"ids,omitempty"`
	Fields []string `json:"fields"`
}

type AddTorrentParams struct {
	DownloadDir   string `json:"download-dir,omitempty"`
	Filename      string `json:"filename,omitempty"`
	FilesUnwanted []int  `json:"files-unwanted",omitempty`
	FilesWanted   []int  `json:"files-wanted,omitempty"`
	// this is base64-encoded .torrent content
	Metainfo       string `json:"metainfo,omitempty"`
	Paused         bool   `json:"paused,omitempty"`
	PeerLimit      int    `json:"peer-limit,omitempty"`
	PriorityHigh   []int  `json:"priority-high,omitempty"`
	PriorityLow    []int  `json:"priority-low,omitempty"`
	PriorityNormal []int  `json:"priority-normal,omitempty"`
}

type AddTorrentResponse struct {
	ID         string `json:"id,omitempty"`
	Name       string `json:"name,omitempty"`
	Hashstring string `json:"hashString,omitempty"`
}

type Session struct {
	AltSpeedDown          int     `json:"alt-speed-down,omitempty"`
	AltSpeedEnabled       bool    `json:"alt-speed-enabled,omitempty"`
	AltSpeedTimeBegin     int     `json:"alt-speed-time-begin,omitempty"`
	AltSpeedTimeDay       int     `json:"alt-speed-time-day,omitempty"`
	AltSpeedTimeEnabled   bool    `json:"alt-speed-time-enabled,omitempty"`
	AltSpeedTimeEnd       int     `json:"alt-speed-time-end,omitempty"`
	AltSpeedUp            int     `json:"alt-speed-up,omitempty"`
	BlocklistEnabled      bool    `json:"blocklist-enabled,omitempty"`
	BlocklistSize         int     `json:"blocklist-size,omitempty"`
	DHTEnabled            bool    `json:"dht-enabled,omitempty"`
	DownloadDir           string  `json:"download-dir,omitempty"`
	Encryption            string  `json:"encryption,omitempty"`
	PeerLimitGlobal       int     `json:"peer-limit-global,omitempty"`
	PeerLimitPerTorrent   int     `json:"peer-limit-per-torrent,omitempty"`
	PeerPort              int     `json:"peer-port,omitempty"`
	PeerPortRandomOnStart bool    `json:"peer-port-random-on-start,omitempty"`
	PEXEnabled            bool    `json:"pex-enabled,omitempty"`
	PortForwardingEnabled bool    `json:"port-forwarding-enabled,omitempty"`
	RPCVersion            int     `json:"rpc-version,omitempty"`
	RPCVersionMinimum     int     `json:"rpc-version-minimum,omitempty"`
	SeedRatioLimit        float64 `json:"seedRatioLimit,omitempty"`
	SeedRatioLimited      bool    `json:"seedRatioLimited,omitempty"`
	SpeedLimitDown        int     `json:"speed-limit-down,omitempty"`
	SpeedLimitDownEnabled bool    `json:"speed-limit-down-enabled,omitempty"`
	SpeedLimitUp          int     `json:"speed-limit-up,omitempty"`
	SpeedLimitUpEnabled   bool    `json:"speed-limit-up-enabled,omitempty"`
	Version               string  `json:"version,omitempty"`
}

type SessionStats struct {
	ActiveTorrentCount int `json:"activeTorrentCount,omitempty"`
	CumulativeStats    struct {
		DownloadedBytes int `json:"downloadedBytes,omitempty"`
		FilesAdded      int `json:"filesAdded,omitempty"`
		SecondsActive   int `json:"secondsActive,omitempty"`
		SessionCount    int `json:"sessionCount,omitempty"`
		UploadedBytes   int `json:"uploadedBytes,omitempty"`
	} `json:"cumulative-stats,omitempty"`
	CurrentStats struct {
		DownloadedBytes int `json:"downloadedBytes,omitempty"`
		FilesAdded      int `json:"filesAdded,omitempty"`
		SecondsActive   int `json:"secondsActive,omitempty"`
		SessionCount    int `json:"sessionCount,omitempty"`
		UploadedBytes   int `json:"uploadedBytes,omitempty"`
	} `json:"current-stats,omitempty"`
	DownloadSpeed      int `json:"downloadSpeed,omitempty"`
	PausedTorrentCount int `json:"pausedTorrentCount,omitempty"`
	TorrentCount       int `json:"torrentCount,omitempty"`
	UploadSpeed        int `json:"uploadSpeed,omitempty"`
}
