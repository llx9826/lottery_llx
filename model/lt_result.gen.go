// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameLtResult = "lt_result"

// LtResult mapped from table <lt_result>
type LtResult struct {
	ID         int32  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	GiftID     int32  `gorm:"column:gift_id;not null" json:"gift_id"`         // 奖品ID，关联lt_gift表
	GiftName   string `gorm:"column:gift_name;not null" json:"gift_name"`     // 奖品名称
	GiftType   int32  `gorm:"column:gift_type;not null" json:"gift_type"`     // 奖品类型，同lt_gift. gtype
	UID        int32  `gorm:"column:uid;not null" json:"uid"`                 // 用户ID
	Username   string `gorm:"column:username;not null" json:"username"`       // 用户名
	PrizeCode  int32  `gorm:"column:prize_code;not null" json:"prize_code"`   // 抽奖编号（4位的随机数）
	GiftData   string `gorm:"column:gift_data;not null" json:"gift_data"`     // 获奖信息
	SysCreated int32  `gorm:"column:sys_created;not null" json:"sys_created"` // 创建时间
	SysIP      string `gorm:"column:sys_ip;not null" json:"sys_ip"`           // 用户抽奖的IP
	SysStatus  int32  `gorm:"column:sys_status;not null" json:"sys_status"`   // 状态，0 正常，1删除，2作弊
}

// TableName LtResult's table name
func (*LtResult) TableName() string {
	return TableNameLtResult
}