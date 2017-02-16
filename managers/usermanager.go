package managers

import "simple-json-restapi/models"

//Return list of users.
func GetUsers() []models.User {
	users := []models.User{
		models.User{Id: 1, Nick: "test1", Email: "test1@gmail.com", Firstname: "FirstName1", LastName: "LastName1"},
		models.User{Id: 2, Nick: "test2", Email: "test2@gmail.com", Firstname: "FirstName2", LastName: "LastName2"},
		models.User{Id: 3, Nick: "test3", Email: "test3@gmail.com", Firstname: "FirstName3", LastName: "LastName3"},
	}
	return users
}

// return user by ID
func GetUserById(id int) models.User {

	for _, item := range GetUsers() {
		if item.Id == id {
			return item
		}
	}

	return models.User{}
}
