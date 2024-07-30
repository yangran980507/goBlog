// Package user 用户模型
package user

import (
	"blog/pkg/encryption"
	"blog/pkg/mysql"
)

// User 用户模型
type User struct {
	// 用户编号
	ID uint `json:"id,omitempty" gorm:"column:id;primaryKey;autoIncrement"`
	// 登陆名
	LoginName string `json:"login_name" gorm:"index"`
	// 真实名
	TrueName string `json:"true_name"`
	// 密码
	PassWord string `json:"-"`
	// 地址
	Address string `json:"address"`
	// 邮编
	PostCode string `json:"post_code"`
	// 电话号码
	Phone string `json:"phone"`
	// 账号是否可用
	Freeze bool `json:"freeze"`
	// 用户身份 1：用户；2：管理
	IsManager bool `json:"-"`
	// 折扣等
	Grade int `json:"grade"`
	// 消费金额
	float64 `json:"amount"`
}

// Create 增加数据
func (userModel *User) Create() {
	mysql.DB.Create(&userModel)
}

// ComparePSW 密码比较
func (userModel *User) ComparePSW(psw string) bool {
	return encryption.BcryptCheck(psw, userModel.PassWord)
}

// MemberFreezeUpdate 更新用户 freeze 字段
func (userModel *User) MemberFreezeUpdate() (user User, row int64) {

	row = mysql.DB.Model(&User{}).Where("id = ?", userModel.ID).
		Update("freeze", userModel.Freeze).RowsAffected

	mysql.DB.Model(&User{}).
		Where("id = ?", userModel.ID).First(&user)

	return

}
