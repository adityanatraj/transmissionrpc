#!/usr/bin/env python
import json

torrent_type = {
  "activityDate":"number",
  "addedDate":"number",
  "announceResponse":"string",
  "announceURL":"string",
  "bandwidthPriority":"number",
  "comment":"string",
  "corruptEver":"number",
  "creator":"string",
  "dateCreated":"number",
  "desiredAvailable":"number",
  "doneDate":"number",
  "downloadDir":"string",
  "downloadedEver":"number",
  "downloaders":"number",
  "downloadLimit":"number",
  "downloadLimited":"boolean",
  "error":"number",
  "errorString":"string",
  "eta":"number",
  "files":[{
    "key": "string",
    "bytesCompleted": "number",
    "length": "number",
    "name": "string"
  }],
  "fileStats":[{
    "bytesCompleted": "number",
    "wanted": "boolean",
    "priority": "number"
  }],
  "hashString":"string",
  "haveUnchecked":"number",
  "haveValid":"number",
  "honorsSessionLimits":"boolean",
  "id":"number",
  "isPrivate":"boolean",
  "lastAnnounceTime":"number",
  "lastScrapeTime":"number",
  "leechers":"number",
  "leftUntilDone":"number",
  "manualAnnounceTime":"number",
  "maxConnectedPeers":"number",
  "name":"string",
  "nextAnnounceTime":"number",
  "nextScrapeTime":"number",
  "peer-limit":"number",
  "peers": [{
    "address":"string",
    "clientName":"string",
    "clientIsChoked":"boolean",
    "clientIsInterested":"boolean",
    "isDownloadingFrom":"boolean",
    "isEncrypted":"boolean",
    "isIncoming":"boolean",
    "isUploadingTo":"boolean",
    "peerIsChoked":"boolean",
    "peerIsInterested":"boolean",
    "port":"number",
    "progress":"double",
    "rateToClient":"number",
    "rateToPeer":"number",
  }], 
  "peersConnected":"number",
  "peersFrom":{
    "fromCache": "number",
    "fromIncoming": "number",
    "fromPex": "number",
    "fromTracker": "number"
  }, 
  "peersGettingFromUs":"number",
  "peersKnown":"number",
  "peersSendingToUs":"number",
  "percentDone":"double",
  "pieces":"string", 
  "pieceCount":"number",
  "pieceSize":"number",
  "priorities":["number"], 
  "rateDownload":"number",
  "rateUpload":"number",
  "recheckProgress":"double",
  "scrapeResponse":"string",
  "scrapeURL":"string",
  "seeders":"number",
  "seedRatioLimit":"double",
  "seedRatioMode":"number",
  "sizeWhenDone":"number",
  "startDate":"number",
  "status":"number",
  "swarmSpeed":"number",
  "timesCompleted":"number",
  "trackers":[{
    "announce": "string",
    "scrape": "string",
    "tier": "number"
  }], 
  "totalSize":"number",
  "torrentFile":"string",
  "uploadedEver":"number",
  "uploadLimit":"number",
  "uploadLimited":"boolean",
  "uploadRatio":"double",
  "wanted":["boolean"], 
  "webseeds":["string"], 
  "webseedsSendingToUs":"number"
}

set_torrent_type = {   
  "bandwidthPriority":"number",
  "downloadLimit":"number",
  "downloadLimited":"boolean",
  "files-wanted":["number"],
  "files-unwanted":["number"],
  "honorsSessionLimits":"boolean",
  "ids":["string"],
  "location":"string",
  "peer-limit":"number",
  "priority-high":["number"],
  "priority-low":["number"],
  "priority-normal":["number"],
  "seedRatioLimit":"double",
  "seedRatioMode":"number",
  "uploadLimit":"number",
  "uploadLimited":"boolean"
}

add_torrent_type = {
  "download-dir":"string",
  "filename":"string",
  "metainfo":"string",
  "paused":"boolean",
  "peer-limit":"number",
  "files-wanted":["number"],
  "files-unwanted":["number"],
  "priority-high":["number"],
  "priority-low":["number"],
  "priority-normal":["number"]
}

session_type = {
  "alt-speed-down":"number",
  "alt-speed-enabled":"boolean",
  "alt-speed-time-begin":"number",
  "alt-speed-time-enabled":"boolean",
  "alt-speed-time-end":"number",
  "alt-speed-time-day":"number",
  "alt-speed-up":"number",
  "blocklist-enabled":"boolean",
  "blocklist-size":"number",
  "dht-enabled":"boolean",
  "encryption":"string",
  "download-dir":"string",
  "peer-limit-global":"number",
  "peer-limit-per-torrent":"number",
  "pex-enabled":"boolean",
  "peer-port":"number",
  "peer-port-random-on-start":"boolean",
  "port-forwarding-enabled":"boolean",
  "rpc-version":"number",
  "rpc-version-minimum":"number",
  "seedRatioLimit":"double",
  "seedRatioLimited":"boolean",
  "speed-limit-down":"number",
  "speed-limit-down-enabled":"boolean",
  "speed-limit-up":"number",
  "speed-limit-up-enabled":"boolean",
  "version":"string"
}

session_stats_type = {
  "activeTorrentCount" :"number",
  "downloadSpeed"      :"number",
  "pausedTorrentCount" :"number",
  "torrentCount"       :"number",
  "uploadSpeed"        :"number",
  "cumulative-stats"   : {
     "uploadedBytes": "number",
     "downloadedBytes": "number",
     "filesAdded"     : "number",
     "sessionCount"   : "number",
     "secondsActive"  : "number",
  },
  "current-stats"      : {
     "uploadedBytes": "number",
     "downloadedBytes": "number",
     "filesAdded"     : "number",
     "sessionCount"   : "number",
     "secondsActive"  : "number",
  }
}


to_type = {
  "number": 2,
  "double": 2.2,
  "string": "two",
  "boolean": True,
}

def create_example(d):
  example = {}
  for key, value in d.iteritems():
    if isinstance(value, str):
      example[key] = to_type[value]
    elif isinstance(value, dict):
      example[key] = create_example(value)
    elif isinstance(value, list):
      if isinstance(value[0], str):
        example[key] = [to_type[value[0]]]
      elif isinstance(value[0], dict):
        example[key] = [create_example(value[0])]
      else:
        print('cant handle this yet')
  return example

# print(json.dumps(create_example(torrent_type)))
# print(json.dumps(create_example(set_torrent_type)))
# print(json.dumps(create_example(add_torrent_type)))
# print(json.dumps(create_example(session_type)))
print(json.dumps(create_example(session_stats_type)))

