package db

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func FindUserByUsername(username string) (user User, err error) {
	err = DB.Table("users").Where("username = ? AND password = ?", username).First(&user).Error
	return
}
