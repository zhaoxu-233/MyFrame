package dao

import (
	"context"
	"errors"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
	"time"
)

var (
	ErrUserDuplicateEmail = errors.New("邮箱冲突")
	ErrUserNotFound       = gorm.ErrRecordNotFound
)

type UserDAO struct {
	db *gorm.DB
}

func NewUserDAO(db *gorm.DB) *UserDAO {
	return &UserDAO{
		db: db,
	}
}
func (dao *UserDAO) Insert(c context.Context, u User) error {
	//存毫秒数
	now := time.Now().UnixMilli()
	u.UpdateTime = now
	u.CreateTime = now
	//保持context的链路调用
	//校验冲突，所有insert的时候需要拿到err
	err := dao.db.WithContext(c).Create(&u).Error
	//获取mysqlerr，并进行断言
	if mysqlErr, ok := err.(*mysql.MySQLError); ok {
		//mysqlerr的错误码
		const uniqueConflictsErrNo uint16 = 1062
		if mysqlErr.Number == uniqueConflictsErrNo {
			return ErrUserDuplicateEmail
		}
	}
	return err
}

func (dao *UserDAO) FindByEmail(c context.Context, email string) (User, error) {
	var u User
	//err := dao.db.WithContext(c).Where("email = ?", email).First(&u).Error
	err := dao.db.WithContext(c).First(&u, "email = ?", email).Error
	return u, err
}

func (dao *UserDAO) EditUser(c context.Context, email string, user User) (User, error) {
	var u User
	err := dao.db.WithContext(c).First(&u, "email = ?", email).Error
	u.Info = user.Info
	u.Brithday = user.Brithday
	u.Name = user.Name
	dao.db.Save(u)
	return u, err
}

//User直接对应数据库表结构
//别的叫法，entity，model，po（persistent object）
type User struct {
	Id int64 `gorm:"primaryKey,autoIncrement"`
	//每个用户都是单独的邮箱，所以要作为唯一索引
	Email    string `gorm:"unique"`
	PassWord string
	Name     string
	Brithday string
	Info     string
	//创建时间，毫秒数
	CreateTime int64
	//更新时间，毫秒数
	UpdateTime int64
}
