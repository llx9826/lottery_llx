// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameLtUser = "lt_user"

// LtUser mapped from table <lt_user>
type LtUser struct {
	ID         int32  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Username   string `gorm:"column:username;not null" json:"username"`       // 用户名
	Blacktime  int32  `gorm:"column:blacktime;not null" json:"blacktime"`     // 黑名单限制到期时间
	Realname   string `gorm:"column:realname;not null" json:"realname"`       // 联系人
	Mobile     string `gorm:"column:mobile;not null" json:"mobile"`           // 手机号
	Address    string `gorm:"column:address;not null" json:"address"`         // 联系地址
	SysCreated int32  `gorm:"column:sys_created;not null" json:"sys_created"` // 创建时间
	SysUpdated int32  `gorm:"column:sys_updated;not null" json:"sys_updated"` // 修改时间
	SysIP      string `gorm:"column:sys_ip;not null" json:"sys_ip"`           // IP地址
}

// TableName LtUser's table name
func (*LtUser) TableName() string {
	return TableNameLtUser
}
