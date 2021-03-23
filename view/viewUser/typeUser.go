package viewUser

type User struct {
	Id       string `json:"id"`
	MenuId   string `json:"menuId"`
	Admin    int    `json:"admin"`
	Name     string `json:"name"`
	NickName string `json:"nickname"`
	Mail     string `json:"email"`
	Password string `json:"-"`
}
