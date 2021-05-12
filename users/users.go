package users

type User struct {
  Login		string		`json:"login"`
  Name		string		`json:"name"`
  Surname	string		`json:"surname"`
}


func NewById(userId int) *User {
  user := User{"test","test","user"}
  return &user
}

