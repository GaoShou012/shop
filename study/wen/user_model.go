package GORMDemo

import (
	"errors"
	"time"
)

type User struct {
	// 用户ID
	Id int `json:"id" gorm:"column:id;primary_key"`
	// 千音ID
	QChatId string `json:"qchat_id" gorm:"column:qchat_id;type:char(50);default:'';comment:'千音ID'"`
	// 密码
	Password string `json:"password" gorm:"column:password;type:char(200);default:'';comment:'密码'"`
	// 性别 0-未知 1-男 2-女
	Sex int `json:"sex" gorm:"column:sex;type:tinyint(4);default:0;comment:'性别 0-未知 1-男 2-女'"`
	// 昵称
	NickName string `json:"nick_name" gorm:"column:nick_name;type:varchar(20);default:'';comment:'用户昵称'"`
	// 邮箱
	Email string `json:"email" gorm:"column:email;type:char(200);default:'';comment:'邮箱'"`
	// 国际区号
	TelArea string `json:"tel_area" gorm:"column:tel_area;type:char(10);default:'';comment:'国际区号'"`
	// 手机号码
	Mobile string `json:"mobile" gorm:"column:mobile;type:char(20);default:'';comment:'手机号码'"`
	// 头像
	Photo string `json:"photo" gorm:"column:photo;type:char(200);default:'';comment:'头像'"`
	// 账户状态 0-正常 1-未激活 2-锁定 3-岳飞状态 4-已注销
	Status int `json:"status" gorm:"column:status;type:tinyint(4);default:1;comment:'帐户状态 0-正常 1-未激活 2-锁定 3-岳飞状态 4-已注销'"`
	// 账户类型 0-内部员工 1-外部用户
	Type int `json:"type" gorm:"column:type;type:tinyint(4);default:1;comment:'帐户类型 0-内部员工 1-外部用户'"`
	// 帐号是否系统创建 0-用户自创帐号(默认值) 1-系统创建
	IsSystem int `json:"is_system" gorm:"column:is_system;type:tinyint(4);default:0;comment:'帐号是否系统创建 0-用户自创建帐号(默认值) 1-系统创建'"`
	// 帐号创建时间
	CreateDate time.Time `json:"create_date"`
	// 帐号更新时间
	UpdateDate time.Time `json:"update_date"`
	// 是否默认的千音ID
	IsDefaultID int `json:"is_default_id" gorm:"column:is_default_id;type:tinyint(4);default:1;comment:'是否默认的千音ID'"`
}

func (u *User) TableName() string { return "user" }

func (u *User) IsExist(query string) (int, error) {
	var count int
	db := DB
	if query != "" {
		err := db.Model(&User{}).Where(query).Count(&count).Error
		if err != nil {
			return 0, err
		}
		return count, nil
	}

	return 0, errors.New("User.IsExist fail, query is empty")

}

func (u *User) Create() error {
	db := DB

	err := db.Create(u).Error
	if err != nil {
		return err
	}

	return nil
}

func (u *User) Update(fields map[string]interface{}) error {
	db := DB
	return db.Model(u).Update(fields).Error
}

func (u *User) Retrieve(conditions string) ([]User, error) {
	users := make([]User, 0)

	db := DB
	err := db.Where(conditions).Find(&users).Error

	return users, err
}

func (u *User) Delete(conditions string) error {
	db := DB
	err := db.Model(&User{}).Unscoped().Where(conditions).Delete(User{}).Error
	return err
}
