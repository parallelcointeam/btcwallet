package gui

import (
	"net/url"

	"github.com/parallelcointeam/mod/gui/jdb"
	"github.com/parallelcointeam/mod/gui/vue"
	"github.com/parallelcointeam/mod/wallet"
	"github.com/zserge/webview"
)

type VDATA struct {
	Config jdb.InfConf       `json:"config"`
	Pages  map[string]string `json:"pages"`
	Icons  map[string]string `json:"icons"`
	Imgs   map[string][]byte `json:"imgs"`
}

func GUI(wlt *wallet.Wallet) {
	vue.WLT = wlt
	// libs := jdb.VueLibs
	// pages := jdb.VuePages
	w := webview.New(webview.Settings{
		Title:     "ParallelCoin - DUO - True Story",
		Width:     1400,
		Height:    800,
		URL:       `data:text/html,` + url.PathEscape(string(jdb.VLB["apphtml"])),
		Debug:     true,
		Resizable: false,
	})
	defer w.Exit()
	w.Dispatch(func() {
		// w.Bind("blockchaindata", []interface{}{(*btcjson.InfoWalletResult)(nil)})

		w.Bind("blockchaindata", &vue.BlockChain{})
		w.Bind("sendtoaddress", &vue.SendToAddress{})
		w.Bind("language", &vue.Language{})

		// w.Bind("rpchandlers", &vue.RPCHandlers{})
		// w.Bind("blockchaindata", &vue.BlockChain{})
		//w.Bind("icons", &icons)
		w.Bind("vuedata", &VDATA{
			Config: jdb.VCF,
			Pages:  jdb.VPG,
			Icons:  jdb.VIC,
			Imgs:   jdb.VIM,
		})

		w.Eval(string(jdb.VLB["appjs"]))
		w.InjectCSS(string(jdb.VLB["appcss"]))

		w.InjectCSS(string(jdb.VLB["buefycss"]))
		w.InjectCSS(string(jdb.VLB["fontawesome"]))
		w.InjectCSS(string(jdb.VLB["materialdesignicons"]))
		w.InjectCSS(string(jdb.VLB["vuelayerscss"]))

		w.Eval(string(jdb.VLB["vue"]))
		// w.Eval(string(jdb.VLB["easybar"]))
		// w.Eval(string(jdb.VLB["buefyjs"]))
		// w.Eval(string(jdb.VLB["ol"]))
		// w.Eval(string(jdb.VLB["umd"]))

		// w.Eval(string(jdb.VLB["settings"]))

		// w.Eval(string(jdb.VLB["homepage"]))
		// w.Eval(string(jdb.VLB["sendpage"]))
		// w.Eval(string(jdb.VLB["receivepage"]))
		// w.Eval(string(jdb.VLB["addressbookpage"]))
		// w.Eval(string(jdb.VLB["historypage"]))
		// w.Eval(string(jdb.VLB["peerspage"]))
		// w.Eval(string(jdb.VLB["blockspage"]))
		// w.Eval(string(jdb.VLB["helppage"]))

	})

	w.Run()
}
