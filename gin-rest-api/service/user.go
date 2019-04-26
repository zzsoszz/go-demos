package user

import (
	"fmt"
	"gin-rest-api/dao"
	"gin-rest-api/model"
)


type UserService struct {
	dao *dao.UserDao
}


func New() (service *UserService)   {
	return &UserService{
		dao:dao.New(),
	}
}


func (s *UserService)  GetUserById(id string) (user *model.User,err error) {
	return s.dao.GetUserById(id);
}


func init()  {
	fmt.Println("service");
}



