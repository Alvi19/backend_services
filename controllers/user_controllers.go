package controllers

import (
	"backend_services/database"
	"backend_services/helper"
	"backend_services/models"
	"backend_services/models/reqresp"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserControllers struct {
	Validator *validator.Validate
	DB        *gorm.DB
}

func NewUserController(db *gorm.DB) *UserControllers {
	return &UserControllers{
		Validator: validator.New(),
		DB:        db,
	}
}

// GetAllUsers godoc
// @Summary Get all users
// @Description Get a list of all users
// @Tags users
// @Accept json
// @Produce json
// @Param page query int false "Page number"
// @Param page_size query int false "Page size"
// @Param vendor_id query int false "Filter by vendor ID"
// @Param role_id query int false "Filter by role ID"
// @Param name query string false "Filter by name"
// @Param email query string false "Filter by email"
// @Param phone query string false "Filter by phone"
// @Success 200 {array} models.UserIndex
// @Failure 400 {object} string
// @Router /users [get]
func (uc *UserControllers) GetAllUsers(c *fiber.Ctx) error {
	var events []models.User

	filters := []string{
		"vendor_id", "role_id", "name", "email", "phone",
	}

	filteredData, err := helper.GetFilteredDataWithPagination(&models.User{}, uc.DB, c, filters)
	if err != nil {
		return helper.ErrorResponse(c, "Failed to get user", err.Error())
	}

	result := filteredData.Query.
		Preload("Vendor").
		Find(&events)

	if result.Error != nil {
		return helper.ErrorResponse(c, "Failed to fetch data", result.Error.Error())
	}

	response := helper.PaginatedResponse{
		Total:    filteredData.Total,
		Page:     filteredData.Page,
		PageSize: filteredData.PageSize,
		Data:     events,
	}

	return helper.SuccessResponse(c, "Success get all user", response)
}

// GetUser godoc
// @Summary Get user by ID
// @Description Get a user by their unique ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} models.User
// @Failure 404 {object} string
// @Router /users/{id} [get]
func (uc *UserControllers) GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.User

	if err := uc.DB.Preload("Vendor").First(&user, "uuid = ?", id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Ticket not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(user)
}

// CreateUser godoc
// @Summary Create a new user
// @Description Create a new user record
// @Tags users
// @Accept json
// @Produce json
// @Param user body models.UserExample true "User Info"
// @Success 201 {object} models.User
// @Failure 400 {object} string
// @Router /users [post]
func (cc *UserControllers) CreateUser(c *fiber.Ctx) error {
	var usr models.User

	if err := c.BodyParser(&usr); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to parse request body",
		})
	}

	if err := cc.Validator.Struct(usr); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Validation failed",
		})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(usr.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to hash password",
		})
	}
	usr.Password = string(hashedPassword)

	if err := cc.DB.Create(&usr).Error; err != nil {
		return helper.ErrorResponse(c, "Failed to create Ticket", err.Error())
	}

	return helper.SuccessResponse(c, "Success create user", usr)
}

// UpdateUser godoc
// @Summary Update user by ID
// @Description Update an existing user's details by ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param user body models.UserExample true "Updated User Info"
// @Success 200 {object} models.User
// @Failure 400 {object} string
// @Failure 404 {object} string
// @Router /users/{id} [put]
func (uc *UserControllers) UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")

	var input map[string]interface{}

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&reqresp.ErrorResponse{
			Status:  "error",
			Message: "Invalid request body",
		})
	}

	var brand models.User

	if err := database.DB.First(&brand, id).Error; err != nil {
		return helper.ErrorResponse(c, "Failed to get user", err.Error())
	}

	if err := database.DB.Model(&brand).Updates(input).Error; err != nil {
		return helper.ErrorResponse(c, "Failed to update user", err.Error())
	}

	return helper.SuccessResponse(c, "Success update user", brand)
}

// DeleteUser godoc
// @Summary Delete user by ID
// @Description Soft delete a user record (update DeletedAt field)
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User UUID"
// @Success 204 {string} string "User deleted"
// @Failure 404 {object} string
// @Router /users/{id} [delete]
func (uc *UserControllers) DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")

	var record models.User

	if err := database.DB.First(&record, "uuid = ?", id).Error; err != nil {
		return helper.ErrorResponse(c, "Failed to get user", err.Error())
	}

	if err := database.DB.Model(&record).Update("deleted_at", time.Now()).Error; err != nil {
		return helper.ErrorResponse(c, "Failed to delete", err.Error())
	}

	return helper.SuccessResponse(c, "Success delete user", record)
}
