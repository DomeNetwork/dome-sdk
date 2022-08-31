package sdk

import (
	"fmt"
	"math/big"

	"github.com/domenetwork/dome-lib/pkg/log"
	"github.com/domenetwork/dome-lib/pkg/wallet"
)

// Account will return the account for the provided coin.
func (sdk *SDK) Account(coin wallet.Coin) (addr *wallet.Account, err error) {
	log.D("sdk", "coins", "account", coin)
	addr = coin.GetAccount()
	if addr == nil {
		err = fmt.Errorf("unable to load base account for coin `%s`", coin.GetName())
	}
	return
}

// Accounts returns a list of all accounts tracked in the provided coin.
func (sdk *SDK) Accounts(coin wallet.Coin) (addrs []*wallet.Account, err error) {
	log.D("sdk", "coins", "accounts", coin)
	addrs = coin.GetAccounts()
	return
}

// AddCoins allows for the easy addition of coin types.
func (sdk *SDK) AddCoins(coins ...wallet.Coin) {
	log.D("sdk", "coins", "add coins", coins)
	sdk.wallet.AddCoins(coins...)
}

// Balance will return the balance of the provided coin.  Each coin manages its
// own definition of a balance.
func (sdk *SDK) Balance(coin wallet.Coin) (balance *big.Int, err error) {
	log.D("sdk", "coins", "balance", coin)
	balance, err = coin.GetBalance()
	return
}

// Coins will return a list of all the currently known coins.
func (sdk *SDK) Coins() (coins []wallet.Coin, err error) {
	log.D("sdk", "coins", "coins")
	coins = sdk.wallet.GetCoins()
	return
}

// Coin returns the coin type when provided its name.
func (sdk *SDK) Coin(name string) (coin wallet.Coin, err error) {
	log.D("sdk", "coins", "coin", name)
	coin = sdk.wallet.GetCoin(name)
	return
}

// Gas returns the suggested gas price for a transaction of the provided coin.
// Coins manage their own definition of what gas, if any, is.
func (sdk *SDK) Gas(coin wallet.Coin) (gas *big.Int, err error) {
	log.D("sdk", "coins", "gas", coin)
	gas, err = coin.GetGas()
	return
}

// Generate a new account for the provided coin.  Each coin handles this generation
// step and its documentation should be consulted for more details.
func (sdk *SDK) Generate(coin wallet.Coin) (acct *wallet.Account, err error) {
	log.D("sdk", "coins", "generate", coin)
	acct, err = coin.GetKey(sdk.wallet)
	return
}

// Send a transaction for the provided coin to the given address in the amount provided.  The
// data provided will be attached to the transaction.
func (sdk *SDK) Send(coin wallet.Coin, to string, amount *big.Int, data []byte) (tx string, err error) {
	log.D("sdk", "coins", "send", coin, to, amount, data)
	tx, err = coin.SendTX(to, amount, data)
	return
}
