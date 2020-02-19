package table

import (
	"time"
)

// TCustWalletPwd 用户钱包配置表
type TCustWalletPwd struct {
	UserId         int       `db:"user_id"`        // 用户ID
	Pwd            string    `db:"wallet_id"`      // 钱包密码
	Pwdtype        string    `db:"wallet_name"`    // 密码类型
	Lockstatus     int       `db:"lockstatus"`     // 钱包锁定状态
	Lockdate       time.Time `db:"lockdate"`       // 钱包锁定时间
	LockExpiredate time.Time `db:"lockexpiredate"` // 钱包锁定过期时间
}
