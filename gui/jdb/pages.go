package jdb

import (
	"io/ioutil"
)

type VPages map[string]string

var app, _ = ioutil.ReadFile("./assets/pages/app.html")
var home, _ = ioutil.ReadFile("./assets/pages/home.html")
var send, _ = ioutil.ReadFile("./assets/pages/send.html")
var receive, _ = ioutil.ReadFile("./assets/pages/receive.html")
var history, _ = ioutil.ReadFile("./assets/pages/history.html")
var addressbook, _ = ioutil.ReadFile("./assets/pages/addressbook.html")
var settings, _ = ioutil.ReadFile("./assets/pages/settings.html")
var peers, _ = ioutil.ReadFile("./assets/pages/peers.html")
var blocks, _ = ioutil.ReadFile("./assets/pages/blocks.html")
var about, _ = ioutil.ReadFile("./assets/pages/about.html")
var help, _ = ioutil.ReadFile("./assets/pages/help.html")

var VPG VPages = VPages{
	"app":         string(app),
	"home":        string(home),
	"send":        string(send),
	"receive":     string(receive),
	"history":     string(history),
	"addressbook": string(addressbook),
	"settings":    string(settings),
	"peers":       string(peers),
	"blocks":      string(blocks),
	"about":       string(about),
	"help":        string(help),
}

// func init() {
// 	fmt.Println("daaaaaa", string(home))
// }
