package table

// TCustWallet 用户钱包表
type TCustWallet struct {
	UserId     int     `db:"user_id"`     // 用户ID
	WalletId   int     `db:"wallet_id"`   // 钱包ID
	WalletName string  `db:"wallet_name"` // 钱包名字
	Amount     float64 `db:"amount"`      // 钱包金额
	Frozen     float64 `db:"frozen"`      // 冻结金额
	Status     int     `db:"status"`      // 钱包状态
}
