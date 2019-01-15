package jdb

import (
	"fmt"

	scribble "github.com/nanobox-io/golang-scribble"
)

var JDB, _ = scribble.New("./gui/jdb/", nil)

var VCF InfConf

func init() {

	// ICF := InfConf{
	// 	Lang:  "en",
	// 	Deno:  "min",
	// 	Fiat:  "min",
	// 	Theme: "min",
	// 	CCSS:  "min",
	// 	Start: "min",
	// 	Tray:  true,
	// }
	// NCF := NetConf{
	// 	TLS:     true,
	// 	Network: "network",
	// 	RPC:     "rpc",
	// 	SRPC:    "srpc",
	// 	TLSpub:  "tlspub",
	// 	TLSpri:  "tlspri",
	// 	Proxy:   "rpc",
	// }
	// SCF := SecConf{
	// 	Network: "network",
	// }
	// MCF := MiningConf{
	// 	Algo:  "network",
	// 	CPU:   "network",
	// 	Cores: 6,
	// }
	// JDB.Write("conf", "interface", ICF)
	// JDB.Write("conf", "network", NCF)
	// JDB.Write("conf", "security", SCF)
	// JDB.Write("conf", "mining", MCF)

	if err := JDB.Read("conf", "interface", &VCF); err != nil {
		fmt.Println("Error", err)
	}
	// configuration, err := JDB.ReadAll("conf")
	// if err != nil {
	// 	fmt.Println("Error", err)
	// }

	// for _, f := range configuration {
	// 	var conf interface{}
	// 	if err := json.Unmarshal([]byte(f), &conf); err != nil {
	// 		fmt.Println("Error", err)
	// 	}
	// 	CNF = append(CNF, conf)
	// }

	// db.Write("data", "vlibs", VLB)
	// db.Write("data", "vpages", VPG)
	// db.Write("data", "vicons", VIC)
	// db.Write("data", "vimgs", VIM)

	// if err := db.Read("data", "vlibs", &VueLibs); err != nil {
	// 	fmt.Println("Error", err)
	// }
	// if err := db.Read("data", "vpages", &VuePages); err != nil {
	// 	fmt.Println("Error", err)
	// }
	// if err := db.Read("data", "vicons", &VueIcons); err != nil {
	// 	fmt.Println("Error", err)
	// }
	// if err := db.Read("data", "vimgs", &VueImgs); err != nil {
	// 	fmt.Println("Error", err)
	// }
	// if err := db.Read("data", "vicons", &VueIcons); err != nil {
	// 	fmt.Println("Error", err)
	// }

	// if err := db.Read("vicons", "vicons", &VIcons); err != nil {
	// 	fmt.Println("Error", err)
	// }
	// if err := db.Read("data", "settings", &Settings); err != nil {
	// 	fmt.Println("Error", err)
	// }

}
