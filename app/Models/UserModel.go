package Models

type UserModel struct {
	Id        int64  `db:"ID" json:"id"`
	Username  string `db:"Username" json:"username"`
	Password  string `db:"Password" json:"password"`
	Firstname string `db:"Firstname" json:"firstname"`
	Lastname  string `db:"Lastname" json:"lastname"`
}
