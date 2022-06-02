package databases

import (
	"happy/models"
	"log"
)

func AddUser(user models.User) (int64,error){
	// 将user存入数据库
	info, err :=  DBClient.Insert(user)
	/*if err != nil {
		log.Fatal("user插入数据库失败")
		return 0
	}*/
	return info,err
}

func DeleteUser(id int) (int64,error) {
	info, err := DBClient.Where("id = ?",id).Delete(models.User{})
	/*if err != nil {
		log.Fatal("user删除数据库失败")
		return 0
	}*/
	return info,err
}

func EditUser(user models.User) (int64,error) {

	info, err := DBClient.ID(user.Id).Update(user)
	if err != nil {
		log.Println("user更新数据库失败")
		return 0,err
	}
	return info,err
}

func FindAllUser() ([]models.User,error) {
	users := make([]models.User,0)
	err := DBClient.Find(&users)
/*	if err != nil {
		log.Fatal("user查询所有数据失败")
		return nil
	}*/
	return users,err
}

func FindUserById(id int) (models.User,error) {
	user := new(models.User)
	_,err := DBClient.Where("id = ? ",id).Get(user)
	if err != nil {
		log.Println("ser查询数据失败通过id",err.Error())
		//log.Fatal("user查询数据失败通过id")
		 return *user,err
	}
	return *user,nil

}

func FindUserByUserName(username string, password string) models.User {
	user := new(models.User)
	_,err := DBClient.Where("username = ? and password = ?",username,password).Get(user)
	if err != nil {
		log.Fatal("user查询数据失败通过id")
		// return nil
	}
	return *user
}

