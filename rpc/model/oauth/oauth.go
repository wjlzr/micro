package oauth

import (
	"github.com/tal-tech/go-zero/core/logx"
	"gorm.io/gorm"
	"time"
)

type (
	UserModel struct {
		eloquent  *gorm.DB
		tableName string
	}

	User struct {
		Id             int64     `db:"id"`               // 编号
		Name           string    `db:"name"`             // 用户名
		NickName       string    `db:"nick_name"`        // 昵称
		Avatar         string    `db:"avatar"`           // 头像
		Password       string    `db:"password"`         // 密码
		Salt           string    `db:"salt"`             // 加密盐
		Email          string    `db:"email"`            // 邮箱
		Mobile         string    `db:"mobile"`           // 手机号
		Status         int64     `db:"status"`           // 状态  0：禁用   1：正常
		DeptId         int64     `db:"dept_id"`          // 机构ID
		JobId          int64     `db:"job_id"`           // 岗位ID
		CreateBy       string    `db:"create_by"`        // 创建人
		CreateTime     time.Time `db:"create_time"`      // 创建时间
		LastUpdateBy   string    `db:"last_update_by"`   // 更新人
		LastUpdateTime time.Time `db:"last_update_time"` // 更新时间
		DelFlag        int64     `db:"del_flag"`         // 是否删除  -1：已删除  0：正常

	}
)

// NewSysUserModel Separate data service
func NewSysUserModel(eloquent *gorm.DB) *UserModel {
	return &UserModel{
		eloquent:  eloquent,
		tableName: "user",
	}
}

// 查询单个
func (m *UserModel) FindOne(condition map[string]interface{}) (*User, int64, error) {

	var user User

	result := m.eloquent.Table(m.tableName).Select("id", "name").Where(condition).FirstOrInit(&user)
	if result.Error != nil {
		logx.Errorf("user FindOne First Err：%s", result.Error)
		return nil, 0, result.Error
	}

	return &user, result.RowsAffected, nil
}
