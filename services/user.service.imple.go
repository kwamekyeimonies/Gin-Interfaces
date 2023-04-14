package services

import (
	"context"

	"github.com/kwamekyeimonies/Gin-Interfaces/models"
	"go.mongodb.org/mongo-driver/mongo"
)



type UserServiceImpl struct {
	usercollection *mongo.Collection
	ctx context.Context
}

func NewService(usercollection *mongo.Collection,ctx context.Context) UserService{
	return &UserServiceImpl{
		usercollection: usercollection,
		ctx: ctx,
	}
}

func (u *UserServiceImpl) CreateUser(user *models.User) error{
	_,err := u.usercollection.InsertOne(u.ctx,user)
	return err
}

func (u *UserServiceImpl) GetUser(name *string) (*models.User,error){
	return nil, nil
}

func (u *UserServiceImpl) GetAllUsers() ([]*models.User,error){
	return nil,nil 
}

func (u *UserServiceImpl) UpdateUser(user *models.User) error{
	return nil
}

func (u *UserServiceImpl)DeleteUser(name *string) error{
	return nil
}
