package models

import(
	"encoding/json"
	"io/ioutil"
	"os"
	"sync"
	"time"
	"fmt"
)

// Объект юзеров для парсинга json файла
type Users struct{
	Users []*User `json:"users"`
}


// Глобальный массив с пользователями
type DataBase struct{
	Mutex *sync.Mutex
	users map[string]*User
	add chan *User
	update chan *User
	delete chan string
	newUsers []*User
}

var Db *DataBase

// Инициализация базы данных, парсинг json файла
func InitDB(){

	Db = &DataBase{
		Mutex: new(sync.Mutex),
		users: make(map[string]*User),
		add: make(chan *User),
		update: make(chan *User),
		delete: make(chan string),
		newUsers: make([]*User, 0),
	}

	var u Users

	jsonFile, err := os.Open("users.json")
	if err != nil {
	    fmt.Println(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(byteValue, &u)

	for _, v := range(u.Users){
		Db.users[v.Id] = v
	}

}

// Прослушивание каналов для работы с дб
func(db *DataBase)RunDBHub(){
	ticker := time.NewTicker(3 * time.Second)
	for{
		select{
		case u := <-db.add:	// Добавить юзера
			db.users[u.Id] = u
			db.newUsers = append(db.newUsers, u)
		case u := <-db.update: // Обновить юзера
			db.users[u.Id] = u
		case id := <- db.delete: // Удалить юзера
			delete(db.users, id)
		case <-ticker.C: // Обновлять js бд каждые 3 секунды
			if len(db.newUsers) > 0{

				var u Users

				for _, v := range(db.users){
					u.Users = append(u.Users, v)
				}

				file, _ := json.MarshalIndent(u, "", " ")
				err := ioutil.WriteFile("users.json", file, 0644)
				if err != nil {
				    fmt.Println(err)
				}else{
					db.newUsers = make([]*User, 0)
				}

				fmt.Println("New users have been added")
			}
		}
	}
}
