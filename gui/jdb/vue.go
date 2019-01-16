package jdb

import (
	"io/ioutil"
)

type VLibs map[string][]byte

var apphtml, _ = ioutil.ReadFile("./gui/assets/vue/app.html")
var appjs, _ = ioutil.ReadFile("./gui/assets/vue/app.js")
var appcss, _ = ioutil.ReadFile("./gui/assets/vue/app.css")

var vue, _ = ioutil.ReadFile("./gui/assets/vue/js/vue.js")
var buefyjs, _ = ioutil.ReadFile("./gui/assets/vue/js/buefy.js")
var terminal, _ = ioutil.ReadFile("./gui/assets/vue/js/terminal.js")
var ol, _ = ioutil.ReadFile("./gui/assets/vue/js/ol.js")
var umd, _ = ioutil.ReadFile("./gui/assets/vue/js/umd.js")
var scroller, _ = ioutil.ReadFile("./gui/assets/vue/js/scroller.js")

var commands, _ = ioutil.ReadFile("./gui/assets/vue/js/commands.js")
var tasks, _ = ioutil.ReadFile("./gui/assets/vue/js/tasks.js")

var set, _ = ioutil.ReadFile("./gui/assets/vue/components/settings.js")

var homePage, _ = ioutil.ReadFile("./gui/assets/vue/components/pages/home.js")
var sendPage, _ = ioutil.ReadFile("./gui/assets/vue/components/pages/send.js")
var receivePage, _ = ioutil.ReadFile("./gui/assets/vue/components/pages/receive.js")
var addressbookPage, _ = ioutil.ReadFile("./gui/assets/vue/components/pages/addressbook.js")
var historyPage, _ = ioutil.ReadFile("./gui/assets/vue/components/pages/history.js")
var peersPage, _ = ioutil.ReadFile("./gui/assets/vue/components/pages/peers.js")
var blocksPage, _ = ioutil.ReadFile("./gui/assets/vue/components/pages/blocks.js")
var helpPage, _ = ioutil.ReadFile("./gui/assets/vue/components/pages/help.js")
var consolePage, _ = ioutil.ReadFile("./gui/assets/vue/components/pages/console.js")

var buefycss, _ = ioutil.ReadFile("./gui/assets/vue/css/buefy.css")
var vuelayerscss, _ = ioutil.ReadFile("./gui/assets/vue/css/vuelayers.css")
var scrollercss, _ = ioutil.ReadFile("./gui/assets/vue/js/scroller.css")

var fontawesome, _ = ioutil.ReadFile("./gui/assets/vue/css/fontawesome.css")
var materialdesignicons, _ = ioutil.ReadFile("./gui/assets/vue/css/materialdesignicons.css")

var VLB VLibs = VLibs{
	"apphtml": apphtml,
	"appjs":   appjs,
	"appcss":  appcss,

	"fontawesome":         fontawesome,
	"materialdesignicons": materialdesignicons,
	"buefycss":            buefycss,
	"vuelayerscss":        vuelayerscss,
	"scrollercss":         scrollercss,

	"vue":      vue,
	"scroller": scroller,
	"terminal": terminal,
	"buefyjs":  buefyjs,
	"ol":       ol,
	"umd":      umd,

	"commands": commands,
	"tasks":    tasks,

	"settings":        set,
	"homepage":        homePage,
	"sendpage":        sendPage,
	"receivepage":     receivePage,
	"addressbookpage": addressbookPage,
	"historypage":     historyPage,
	"peerspage":       peersPage,
	"blockspage":      blocksPage,
	"helppage":        helpPage,
	"consolepage":     consolePage,
}
