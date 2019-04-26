package  dao



import (
	"fmt"
	db "gin-rest-api-ioc/commons"
	"gin-rest-api-ioc/model"
	"github.com/jinzhu/gorm"
)



type UserDao struct {
	db *gorm.DB
}


func  New()  (service *UserDao) {
	return &UserDao{
		db:db.GetDb(),
	}
}

func (dao *UserDao) GetUserById(id string) (user *model.User,err error) {
	 var u model.User;
	 err=dao.db.First(&u,1).Error;
	 return &u,err
}

func init()  {

	fmt.Println("dao");
}