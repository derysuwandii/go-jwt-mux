package models

type TblUser struct {
	Id          int64  `json:"id"`
	NamaLengkap string `json:"nama_lengkap"`
	Email       string `json:"email"`
	Username    string `json:"username"`
	Password    string `json:"password"`
}
