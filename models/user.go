package models

import(
	"time"
	"strconv"
	"crypto/md5"
	"encoding/hex"
)

// Объект юзер
type User struct{
	Id string `json:"id"`
	Name string `json:"name"`
}

// Получить все записи из бд
func GetUsers()(interface{}, error){
	return map[string]interface{}{
		"success":true,
		"user":Db.users,
	}, nil
}

// Сгенерировать уникальный id
func HashId()string{
	now := time.Now()
	sec := now.Unix()
	hash := md5.Sum([]byte(strconv.FormatInt(sec, 10)))
	return hex.EncodeToString(hash[:])
}

// Получить юзера по id
func(u *User)GetUser()(interface{}, error){
	if user, ok := Db.users[u.Id]; ok {
		return map[string]interface{}{
			"success":true,
			"user":user,
		}, nil
	}else{
		return map[string]interface{}{
			"success":false,
			"msg":"User not found",
		}, nil
	}
}

// Создать юзера
func(u *User)CreateUser(){
	Db.Mutex.Lock()
	defer Db.Mutex.Unlock()
	u.Id = HashId()
	Db.add <- u
}

// Обновить юзера
func(u *User)UpdateUser()(interface{}, error){

	if _, ok := Db.users[u.Id]; !ok {
		return map[string]interface{}{
			"success":false,
			"msg":"User not found",
		}, nil
	}

	Db.update <- u

	return map[string]interface{}{
		"success":true,
		"msg":"User has been updated",
		"user":u,
	}, nil
}

// Удалить юзера
func(u *User)DeleteUser()(interface{}, error){
	if _, ok := Db.users[u.Id]; !ok {
		return map[string]interface{}{
			"success":false,
			"msg":"User not found",
		}, nil
	}

	Db.delete <- u.Id

	return map[string]interface{}{
		"success":true,
		"msg":"User has been deleted",
		"user":u,
	}, nil
}