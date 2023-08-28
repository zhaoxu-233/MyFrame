package repository

import (
	"context"
	"exercise_code/webook/internal/domain"
	"exercise_code/webook/internal/repository/dao"
)

var (
	ErrUserDuplicateEmail = dao.ErrUserDuplicateEmail
	ErrUserNotFound       = dao.ErrUserNotFound
)

type UserRepository struct {
	dao *dao.UserDAO
}

func NewUserRepository(dao *dao.UserDAO) *UserRepository {
	return &UserRepository{
		dao: dao,
	}
}

func (r *UserRepository) Create(c context.Context, u domain.User) error {
	return r.dao.Insert(c, dao.User{
		Email:    u.Email,
		PassWord: u.PassWord,
	})
	//再这里操作缓存
}
func (r *UserRepository) FindById(int) {

}
func (r *UserRepository) FindByEmail(c context.Context, email string) (domain.User, error) {
	u, err := r.dao.FindByEmail(c, email)
	if err != nil {
		return domain.User{}, err
	}
	return domain.User{
		Id:       u.Id,
		Email:    u.Email,
		PassWord: u.PassWord,
		Name:     u.Name,
		Brithday: u.Brithday,
		Info:     u.Info,
	}, nil
}

func (r *UserRepository) EditUser(c context.Context, email string, user domain.User) error {
	_, err := r.dao.EditUser(c, email, dao.User{
		Name:     user.Name,
		Info:     user.Info,
		Brithday: user.Brithday,
	})
	if err != nil {
		return err
	}
	return err
}
