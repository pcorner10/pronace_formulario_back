package orm

import (
	"pronaces_back/internal/domain/entity"
	"pronaces_back/internal/infraestructure/database"
	"testing"
)

var db = database.GetDB()

func TestUserRepository_Create(t *testing.T) {

	// Set up the repository with the database connection
	repo := NewUserRepository(db)

	// Set up the user to be created
	user := entity.User{
		UserName: "John Doe",
		Email:    "johndoe@example.com",
	}

	// Call the function being tested
	err := repo.Create(user)
	if err != nil {
		t.Fatalf("Error creating user: %s", err)
	}

	// Check that the user was created
	createdUser := &entity.User{}
	err = db.First(createdUser, user.ID).Error
	if err != nil {
		t.Fatalf("Error getting created user: %s", err)
	}

	if createdUser.UserName != user.UserName || createdUser.Email != user.Email {
		t.Errorf("Created user does not match expected user. Got %+v, expected %+v", createdUser, user)
	}
}

func TestUserRepository_GetByEmail(t *testing.T) {
	// Set up the repository with the database connection
	repo := &userRepository{db: db}

	// Set up the user to be created
	user := &entity.User{
		UserName: "John Doe",
		Email:    "johndoe@example.com",
	}

	// Create the user in the database
	err := db.Create(user).Error
	if err != nil {
		t.Fatalf("Error creating user: %s", err)
	}

	// Call the function being tested
	retrievedUser, err := repo.GetByEmail(user.Email)
	if err != nil {
		t.Fatalf("Error getting user by email: %s", err)
	}

	// Check that the retrieved user matches the expected user
	if retrievedUser.ID != user.ID || retrievedUser.UserName != user.UserName || retrievedUser.Email != user.Email {
		t.Errorf("Retrieved user does not match expected user. Got %+v, expected %+v", retrievedUser, user)
	}
}
