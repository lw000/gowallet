package table

import "time"

type TTradeInfo struct {
	UserId      int       `db:"user_id"`      // 用户ID
	FwalletId   int       `db:"fwallet_id"`   // 钱包ID
	TwalletId   int       `db:"twallet_id"`   // 钱包ID
	Amount      float64   `db:"amount"`       // 交易金额
	TradeCode   string    `db:"trade_code"`   // 交易码
	TradeTime   time.Time `db:"trade_time"`   // 交易时间
	TradeStatus time.Time `db:"trade_status"` // 交易状态
}
