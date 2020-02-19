package table

// TWallet 钱包表
type TWallet struct {
	WalletId     int    `db:"wallet_id"`     // 钱包ID
	WalletName   string `db:"wallet_name"`   // 钱包名字
	WalletStatus int    `db:"wallet_status"` // 钱包状态
}
