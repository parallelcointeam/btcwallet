package vue

import (
	"fmt"
	"strconv"

	"github.com/parallelcointeam/mod/waddrmgr"
	"github.com/parallelcointeam/mod/wallet"
	"github.com/parallelcointeam/mod/wallet/txrules"
	"github.com/parallelcointeam/pod/btcjson"
	"github.com/parallelcointeam/pod/btcutil"
	// "github.com/parallelcointeam/pod/rpcclient"
)

var WLT *wallet.Wallet

type BlockChain struct {
	GetInfo                 *btcjson.InfoWalletResult        `json:"getinfo"`
	ListTransactions        []btcjson.ListTransactionsResult `json:"listtransactions"`
	ListAllTransactions     []btcjson.ListTransactionsResult `json:"listalltransactions"`
	ListAllSendTransactions []btcjson.ListTransactionsResult `json:"listallsendtransactions"`
	Balance                 float64                          `json:"balance"`
	UnConfirmed             float64                          `json:"unconfirmed"`
	// GetInfo interface{} `json:"getinfo"`
}
type SendToAddress struct {
	Address string  `json:"address"`
	Label   string  `json:"label"`
	Amount  float64 `json:"amount"`
}

func (k *BlockChain) GetInfoData() {

	// List Transactions

	k.ListTransactions, _ = WLT.ListTransactions(0, 10)

	// List All Transactions

	k.ListAllTransactions, _ = WLT.ListAllTransactions()

	// List Send Transactions
	var listallsendtransactions []btcjson.ListTransactionsResult
	for _, sent := range k.ListAllTransactions {
		if sent.Category == "send" {
			listallsendtransactions = append(listallsendtransactions, sent)
		}
	}
	k.ListAllSendTransactions = listallsendtransactions

	// // Balance
	// var balance btcutil.Amount
	// accountName := "*"
	// if accountName == "*" {
	// 	balance, err = WLT.CalculateBalance(1)
	// 	if err != nil {
	// 	}
	// }
	// k.Balance = balance.ToDUO()

	// UnConfirmed
	acctName := "default"
	account, err := WLT.AccountNumber(waddrmgr.KeyScopeBIP0044, acctName)
	if err != nil {
	}
	bals, err := WLT.CalculateAccountBalances(account, 1)
	if err != nil {
	}
	unconfirmed := (bals.Total - bals.Spendable).ToDUO()

	k.UnConfirmed = unconfirmed

	// t, err := follower.New("/home/marcetin/.mod/logs/testnet/mod.log", follower.Config{
	// 	Whence: io.SeekEnd,
	// 	Offset: 0,
	// 	Reopen: true,
	// })

	// for line := range t.Lines() {
	// 	fmt.Println("ddddddddddddd", line)
	// }
	// blk := WLT.Manager.SyncedTo()

	// block := chain.GetBlock(blk.Hash.String())
	// fmt.Println("GetInfoGetInfoGetInfoGetInfoGetInfoGetInfoGetInfoGetInfoGetInfoGetInfoGetInfo", k.GetInfo)
	// fmt.Println("listtransactionslisttransactionslisttransactionslisttransactionslisttransactions", k.ListTransactions)
	// fmt.Println("tttttttttttttttttttttttttttttt", blk.Hash.String())
	// fmt.Println("BalanceBalanceBalanceBalanceBalance", k.Balance)

}
func (k *SendToAddress) SendDUO(vaddress string, vlabel string, vamount interface{}) {
	amount, err := strconv.ParseFloat(vamount.(string), 64)
	if err != nil {
	}
	amt, err := btcutil.NewAmount(amount)
	if err != nil {
	}

	pairs := map[string]btcutil.Amount{
		vaddress: amt,
	}

	sendPairs(WLT, pairs, waddrmgr.DefaultAccountNum, 1, txrules.DefaultRelayFeePerKb)

	fmt.Println("----------- address::", vaddress)
	fmt.Println("----------- label::::", vlabel)
	fmt.Println("----------- amount:::", vamount)

	// fmt.Println("----------- amount:::", &legacyrpc.RPCHandlers)

}
