package repository

import (
	"errors"
	"fmt"

	"github.com/TechMaster/golang/08Fiber/Repository/model"
)

type UserRepo struct {
	Users  map[int64]*model.User
	autoID int64
}

var User UserRepo

func init() {
	User = UserRepo{autoID: 0}
	User.Users = make(map[int64]*model.User)
	User.InitData("sql:45312")
}

func (r *UserRepo) getAutoID() int64 {
	r.autoID += 1
	return r.autoID
}
func (r *UserRepo) CreateNewUser(user *model.User) int64 {
	nextID := r.getAutoID()
	user.Id = nextID
	r.Users[nextID] = user
	return nextID
}
func (r *UserRepo) InitData(connection string) {
	fmt.Println("Connect to ", connection)

	r.CreateNewUser(&model.User{
		FirstName:  "Administrator",
		LastName:   "dsadsad",
		Username:   "admin",
		Email:      "admin@gmail.com",
		Password:   "admin",
		Avatar:     "https://robohash.org/eaquequasincidunt.png?size=50x50&set=set1",
		Gender:     "Genderfluid",
		Phone:      "933-658-1213",
		Birthday:   "1994-03-23",
		Status:     true,
		CreatedAt:  1609483221000,
		ModifiedAt: 1609483221000})

	r.CreateNewUser(&model.User{
		FirstName:  "Client 1",
		LastName:   "dsad",
		Username:   "client1",
		Email:      "client1@gmail.com",
		Password:   "client",
		Avatar:     "https://robohash.org/accusantiumminimamagni.png?size=50x50&set=set1",
		Gender:     "Male",
		Phone:      "510-449-7332",
		Birthday:   "2002-03-11",
		Status:     false,
		CreatedAt:  1617440961000,
		ModifiedAt: 1618301961000})
}

func (r *UserRepo) GetAllUsers() map[int64]*model.User {
	return r.Users
}

func (r *UserRepo) FindUserById(Id int64) (*model.User, error) {
	if user, ok := r.Users[Id]; ok {
		return user, nil
	} else {
		return nil, errors.New("user not found")
	}
}

//del
func (r *UserRepo) DeleteUserById(id int64) error {
	if _, ok := r.Users[id]; ok {
		delete(r.Users, id)
		return nil
	} else {
		return errors.New("User not found")
	}
}

//create
func (r *UserRepo) CheckUserByUsername(user *model.User) error {
	for _, v := range r.Users {
		if v.Username == user.Username {
			return errors.New("User not found")
		}
	}
	return nil
}
func (r *UserRepo) UpdateUser(user *model.User) error {
	if _, ok := r.Users[user.Id]; ok {
		r.Users[user.Id] = user
		return nil //tìm được
	} else {
		return errors.New("user not found")
	}
}
