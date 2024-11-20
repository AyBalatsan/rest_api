package todo

type User struct {
	// Посмотреть потом есть ли расширение для описание структуры JSON
	Id       int    `json:"-" db:"id"`
	Name     string `json:"name" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
