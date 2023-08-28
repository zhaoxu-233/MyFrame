package service

import (
	"context"
	"errors"
	"exercise_code/webook/internal/domain"
	"exercise_code/webook/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrUserDuplicateEmail    = repository.ErrUserDuplicateEmail
	ErrInvalidUserorPassWord = errors.New("账号/邮箱或密码错误")
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{
		repo,
	}
}

//加入cont参数是为了保持链路和超时控制
//参数不需要用指针
//service所需的领域对象，定义在domain中
//函数中不知道该返回什么时，一定要返回error信息
func (svc *UserService) SignUp(c context.Context, u domain.User) error {
	//加密放在哪里
	hash, err := bcrypt.GenerateFromPassword([]byte(u.PassWord), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.PassWord = string(hash)
	return svc.repo.Create(c, u)
}

func (svc *UserService) Login(c context.Context, email, password string) (domain.User, error) {
	//去查找用户
	u, err := svc.repo.FindByEmail(c, email)
	if err == repository.ErrUserNotFound {
		return domain.User{}, ErrInvalidUserorPassWord
	}
	if err != nil {
		return domain.User{}, err
	}
	//比较数据库查出来的密码u.password和传入的密码svc.password是否一致
	err = bcrypt.CompareHashAndPassword([]byte(u.PassWord), []byte(password))
	if err != nil {
		return domain.User{}, ErrInvalidUserorPassWord
	}
	return u, err

}

func (svc *UserService) Edit(c context.Context, email string, user domain.User) error {
	//使用唯一索引email去查找要修改的用户
	_, err := svc.repo.FindByEmail(c, email)
	if err == repository.ErrUserNotFound {
		return ErrInvalidUserorPassWord
	}
	if err != nil {
		return err
	}
	err = svc.repo.EditUser(c, email, user)
	return err
}

func (svc *UserService) Profile(c context.Context, email string) (domain.User, error) {
	//去查找用户
	u, err := svc.repo.FindByEmail(c, email)
	if err == repository.ErrUserNotFound {
		return domain.User{}, ErrInvalidUserorPassWord
	}
	if err != nil {
		return domain.User{}, err
	}
	return u, nil

}
