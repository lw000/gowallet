package wallet

import (
	"fmt"
	"gowallet/errors"
	"sync"
)

type Wallet struct {
	sync.RWMutex
	walletId   int     // 钱包ID
	walletName string  // 钱包名字
	amount     float64 // 钱包金额
	frozen     float64 // 冻结金额
	status     int     // 钱包状态
}

func (w *Wallet) Status() int {
	return w.status
}

func (w *Wallet) SetStatus(status int) {
	w.status = status
}

// 冻结金额
func (w *Wallet) Frozen() float64 {
	w.RLock()
	frozen := w.frozen
	w.RUnlock()
	return frozen
}

// 钱包金额
func (w *Wallet) Amount() float64 {
	w.RLock()
	amount := w.amount
	w.RUnlock()
	return amount
}

func (w *Wallet) TransferIn(amount float64) error {
	if amount < 0 {
		return errors.New(0, fmt.Sprintf("转入金额错误[%f]", amount), "")
	}

	w.Lock()
	w.amount += amount
	w.Unlock()
	return nil
}

func (w *Wallet) TransferOut(amount float64) error {
	if amount < 0 {
		return errors.New(0, fmt.Sprintf("转出金额错误[%f]", amount), "")
	}

	if amount > w.amount {
		return errors.New(0, fmt.Sprintf("账户金额不足，转出失败[%f]", amount), "")
	}

	w.Lock()
	w.amount -= amount
	w.Unlock()
	return nil
}

func (w *Wallet) WalletName() string {
	return w.walletName
}

func (w *Wallet) SetWalletName(walletName string) {
	w.walletName = walletName
}

func (w *Wallet) WalletId() int {
	return w.walletId
}

func (w *Wallet) SetWalletId(walletId int) {
	w.walletId = walletId
}
