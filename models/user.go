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
func GetUsers()map[string]interface{}{
	return map[string]interface{}{
		"success":true,
		"users":Db,
	}
}

// Сгенерировать уникальный id
func HashId()string{
	now := time.Now()
	sec := now.Unix()
	hash := md5.Sum([]byte(strconv.FormatInt(sec, 10)))
	return hex.EncodeToString(hash[:])
}

// Получить юзера по id
func(u *User)GetUser()map[string]interface{}{
	if user, ok := Db[u.Id]; ok {
		return map[string]interface{}{
			"success":true,
			"user":user,
		}
	}else{
		return map[string]interface{}{
			"success":false,
			"msg":"User not found",
		}
	}
}

// Создать юзера
func(u *User)CreateUser()map[string]interface{}{
	u.Id = HashId()
	Db[u.Id] = u

	_ = refreshDb()
	
	return map[string]interface{}{
		"success":true,
		"msg":"User has been created",
		"user":u,
	}
}

// Обновить юзера
func(u *User)UpdateUser()map[string]interface{}{

	if _, ok := Db[u.Id]; !ok {
		return map[string]interface{}{
			"success":false,
			"msg":"User not found",
		}
	}

	Db[u.Id] = u

	_ = refreshDb()

	return map[string]interface{}{
		"success":true,
		"msg":"User has been updated",
		"user":u,
	}
}

// Удалить юзера
func(u *User)DeleteUser()map[string]interface{}{
	if _, ok := Db[u.Id]; !ok {
		return map[string]interface{}{
			"success":false,
			"msg":"User not found",
		}
	}

	delete(Db, u.Id)

	_ = refreshDb()

	return map[string]interface{}{
		"success":true,
		"msg":"User has been deleted",
		"user":u,
	}
}