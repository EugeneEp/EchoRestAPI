package models

import(
	"encoding/json"
	"io/ioutil"
	"os"
	"fmt"
)

// Объект юзеров для парсинга json файла
type Users struct{
	Users []*User `json:"users"`
}


// Глобальный массив с пользователями
var Db map[string]*User


// Инициализация базы данных, парсинг json файла
func InitDB(){

	Db = make(map[string]*User)

	var u Users

	jsonFile, err := os.Open("users.json")
	if err != nil {
	    fmt.Println(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(byteValue, &u)

	for _, v := range(u.Users){
		Db[v.Id] = v
	}

}

// Обновить значения json файла базы данных
func refreshDb()error{
	var u Users

	for _, v := range(Db){
		u.Users = append(u.Users, v)
	}

	file, _ := json.MarshalIndent(u, "", " ")
	err := ioutil.WriteFile("users.json", file, 0644)
	if err != nil {
	    return err
	}

	return nil

}