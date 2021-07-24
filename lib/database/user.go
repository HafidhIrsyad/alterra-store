package database

import (
	"alterra_store/configs"
	"alterra_store/helpers"
	"alterra_store/models/users"
)

func RegisterUser(usersCreate users.UserCreate) (users.User, error) {
	hash, _ := helpers.HashPassword(usersCreate.Password)
	var userDB users.User

	userDB.Name = usersCreate.Name
	userDB.Address = usersCreate.Address
	userDB.Email = usersCreate.Email
	userDB.Password = hash

	err := configs.DB.Create(&userDB).Error
	if err != nil {
		return userDB, err
	}
	return userDB, nil
}

func GetDataUserAll() (dataResult []users.User, err error) {
	err = configs.DB.Find(&dataResult).Error
	if err != nil {
		return nil, err
	}
	return
}

func LoginUser(userLogin users.UserLogin) (users.User, error) {
	var userDB users.User

	err := configs.DB.Where("email = ?", userLogin.Email).First(&userDB).Error
	checkHash, _ := helpers.CheckPasswordHash(userLogin.Password, userDB.Password)

	if err != nil && !checkHash {
		return userDB, err
	}
	return userDB, nil
}

func GetUserDetail(userId int) (users.User, error) {
	var userDB users.User
	err := configs.DB.First(&userDB, userId).Error

	if err != nil {
		return userDB, err
	}
	return userDB, nil
}

func CheckHashPassword(confirmPassword string, idUser int) (verified bool, err error) {
	var userDB users.User

	err = configs.DB.Where("id = ?", idUser).First(&userDB).Error
	confirmedUser, _ := helpers.CheckPasswordHash(confirmPassword, userDB.Password)

	if !confirmedUser {
		return false, err
	}
	return true, err
}

func EditUser(userEdit users.UserEdit, idUser int) (users.User, error) {
	hash, _ := helpers.HashPassword(userEdit.NewPassword)
	var userDB users.User
	err := configs.DB.First(&userDB, idUser).Error

	userDB.Name = userEdit.Name
	userDB.Address = userEdit.Address
	userDB.Email = userEdit.Email
	userDB.Password = hash

	configs.DB.Save(&userDB)

	if err != nil {
		return userDB, err
	}
	return userDB, nil
}
