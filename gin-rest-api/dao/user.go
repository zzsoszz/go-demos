package  dao



import (
	"fmt"
	db "gin-rest-api/commons"
	"gin-rest-api/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
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