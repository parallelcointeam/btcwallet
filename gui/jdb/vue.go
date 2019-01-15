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
var ol, _ = ioutil.ReadFile("./gui/assets/vue/js/ol.js")
var umd, _ = ioutil.ReadFile("./gui/assets/vue/js/umd.js")
var easybar, _ = ioutil.ReadFile("./gui/assets/vue/js/easybar.js")

var set, _ = ioutil.ReadFile("./gui/assets/vue/components/settings.js")

var homePage, _ = ioutil.ReadFile("./gui/assets/vue/components/pages/home.js")
var sendPage, _ = ioutil.ReadFile("./gui/assets/vue/components/pages/send.js")
var receivePage, _ = ioutil.ReadFile("./gui/assets/vue/components/pages/receive.js")
var addressbookPage, _ = ioutil.ReadFile("./gui/assets/vue/components/pages/addressbook.js")
var historyPage, _ = ioutil.ReadFile("./gui/assets/vue/components/pages/history.js")
var peersPage, _ = ioutil.ReadFile("./gui/assets/vue/components/pages/peers.js")
var blocksPage, _ = ioutil.ReadFile("./gui/assets/vue/components/pages/blocks.js")
var helpPage, _ = ioutil.ReadFile("./gui/assets/vue/components/pages/help.js")

var buefycss, _ = ioutil.ReadFile("./gui/assets/vue/css/buefy.css")
var vuelayerscss, _ = ioutil.ReadFile("./gui/assets/vue/css/vuelayers.css")
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

	"vue":     vue,
	"easybar": easybar,
	"buefyjs": buefyjs,
	"ol":      ol,
	"umd":     umd,

	"settings":        set,
	"homepage":        homePage,
	"sendpage":        sendPage,
	"receivepage":     receivePage,
	"addressbookpage": addressbookPage,
	"historypage":     historyPage,
	"peerspage":       peersPage,
	"blockspage":      blocksPage,
	"helppage":        helpPage,
}
