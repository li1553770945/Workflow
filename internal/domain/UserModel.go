package domain

import (
	"HomeWorkGo/infra/database"
	"crypto/md5"
	"encoding/hex"
	"time"
)

type UserModel struct {
	ID         int       `json:"id,omitempty" gorm:"primaryKey"`
	Username   string    `json:"username,omitempty"  validate:"required" gorm:"type:varchar(30);uniqueIndex"`
	Name       string    `json:"name,omitempty"  validate:"required"`
	Password   string    `json:"-"  validate:"-"`
	StudentNo  string    `json:"student_no"`
	Status     int       `json:"status,omitempty"`
	Validation string    `json:"-"  validate:"-"`
	ResetCnt   int       `json:"-"  validate:"-"`
	CreatedAt  time.Time `json:"created_at,omitempty" gorm:"autoCreateTime,omitempty"`
	LastLogin  time.Time `json:"last_login,omitempty"`
}

func GetUserByUserName(username string) (user *UserModel, err error) {
	user = new(UserModel)
	if err = database.DB.Where("username = ?", username).First(user).Error; err != nil {
		return nil, err
	}
	return
}

func GetUserById(id int) (user *UserModel, err error) {
	user = new(UserModel)
	if err = database.DB.Where("id=?", id).First(user).Error; err != nil {
		return nil, err
	}
	return
}
func Encrypt(passwd string) string {
	m := md5.New()
	m.Write([]byte(passwd))
	res := hex.EncodeToString(m.Sum(nil))
	return res
}
func SetUserPasswd(user *UserModel, passwd string) {
	user.Password = Encrypt(passwd)
}
func CheckUserPasswd(user *UserModel, passwd string) bool {
	return Encrypt(passwd) == user.Password
}
func UpdateUser(user *UserModel) (err error) {
	err = database.DB.Save(user).Error
	return err
}

func ChangePassword(user *UserModel, password string) (err error) {
	user.Password = Encrypt(password)
	err = database.DB.Save(user).Error
	return err
}
func ResetPassword(user *UserModel, password string) (err error) {
	user.Password = Encrypt(password)
	user.Validation = ""
	user.ResetCnt = 0
	err = database.DB.Save(user).Error
	return err
}
func ResetUsername(user *UserModel, username string) (err error) {
	user.Username = username
	user.Validation = ""
	user.ResetCnt = 0
	err = database.DB.Save(user).Error
	return err
}

type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}
