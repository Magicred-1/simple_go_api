package routes

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

var DB *gorm.DB

type User struct {
	gorm.Model
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

func GetAllUsers(c *fiber.Ctx) error {
	if DB == nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Database connection error")
	}
	var users []User
	DB.Find(&users)
	return c.JSON(users)
}

func GetUserById(c *fiber.Ctx) error {
	// Get parameter value
	id := c.Params("id")

	// Get a database connection
	if DB == nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Database connection error")
	}

	// Get user from database
	var user User
	result := DB.First(&user, id)

	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(result.Error.Error())
	}

	return c.JSON(user)
}

func CreateUser(c *fiber.Ctx) error {
	if DB == nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Database connection error")
	}

	// JSON body parsed into struct
	user := new(User)

	if err := c.BodyParser(user); err != nil {
		return err
	}

	// Create the user
	result := DB.Create(user)

	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(result.Error.Error())
	}

	return c.JSON(user)
}

func DeleteUser(c *fiber.Ctx) error {
	// Get parameter value
	id := c.Params("id")

	if DB == nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Database connection error")
	}

	// Delete user
	result := DB.Delete(&User{}, id)

	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(result.Error.Error())
	}

	return c.SendString(fmt.Sprintf("User with ID %s deleted", id))
}

func UpdateUser(c *fiber.Ctx) error {
	// Get parameter value
	id := c.Params("id")

	if DB == nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Database connection error")
	}

	// JSON body parsed into struct
	user := new(User)

	if err := c.BodyParser(user); err != nil {
		return err
	}

	// Update user in database
	result := DB.Model(&user).Where("id = ?", id).Updates(User{Name: user.Name, Email: user.Email, Age: user.Age})

	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(result.Error.Error())
	}

	return c.JSON(user)
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
