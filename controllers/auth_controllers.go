package controllers

import (
	"backend_services/models"
	"backend_services/models/reqresp"
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthController struct {
	DB *gorm.DB
}

func NewAuthController(db *gorm.DB) *AuthController {
	return &AuthController{
		DB: db,
	}
}

// Login godoc
// @Summary User login
// @Description Authenticate user and return tokens
// @Tags auth
// @Accept json
// @Produce json
// @Param login body models.LoginRequest true "Login credentials"
// @Success 200 {object} reqresp.TokenResponse
// @Failure 400 {object} reqresp.ErrorResponse
// @Failure 401 {object} reqresp.ErrorResponse
// @Router /auth/login [post]
func (ac *AuthController) Login(c *fiber.Ctx) error {
	var data models.LoginRequest

	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&reqresp.ErrorResponse{
			Message: "Invalid request payload",
			Status:  "error",
		})
	}

	var user models.User
	result := ac.DB.Where("email = ? ", data.Email).First(&user)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusUnauthorized).JSON(&reqresp.ErrorResponse{
				Message: "Invalid credentials",
				Status:  "error",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(&reqresp.ErrorResponse{
			Message: "Error while checking credentials",
			Status:  "error",
		})
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password)) != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(&reqresp.ErrorResponse{
			Message: "Invalid credentials",
			Status:  "error",
		})
	}

	tokenAgeHours, _ := strconv.Atoi(os.Getenv("TOKEN_AGE_HOUR"))
	secret := os.Getenv("TOKEN_SECRET")

	accessToken := jwt.New(jwt.SigningMethodHS256)
	claims := accessToken.Claims.(jwt.MapClaims)
	claims["id"] = user.ID
	claims["name"] = user.Name
	claims["email"] = user.Email
	claims["phone"] = user.Phone
	claims["kind"] = "access"
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(tokenAgeHours)).Unix()

	accessTokenString, err := accessToken.SignedString([]byte(secret))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&reqresp.ErrorResponse{
			Message: "Failed to generate access token",
			Status:  "error",
		})
	}

	// Create refresh token
	refreshTokenAgeHours, _ := strconv.Atoi(os.Getenv("REFRESH_TOKEN_AGE_HOUR"))
	refreshToken := jwt.New(jwt.SigningMethodHS256)
	refreshClaims := refreshToken.Claims.(jwt.MapClaims)
	refreshClaims["id"] = user.ID
	refreshClaims["name"] = user.Name
	refreshClaims["email"] = user.Email
	refreshClaims["phone"] = user.Phone
	refreshClaims["kind"] = "refresh"
	refreshClaims["exp"] = time.Now().Add(time.Hour * time.Duration(refreshTokenAgeHours)).Unix()

	refreshTokenString, err := refreshToken.SignedString([]byte(secret))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&reqresp.ErrorResponse{
			Message: "Failed to generate refresh token",
			Status:  "error",
		})
	}

	// Set the access token in the cookie
	c.Cookie(&fiber.Cookie{
		Name:     "access_token",
		Value:    accessTokenString,
		Expires:  time.Now().Add(time.Hour * time.Duration(tokenAgeHours)),
		HTTPOnly: true,
	})

	// Return the access and refresh tokens
	return c.JSON(&reqresp.TokenResponse{
		Message:      "Success",
		AccessToken:  accessTokenString,
		RefreshToken: refreshTokenString,
	})
}

// GetProfile godoc
// @Summary Get user profile
// @Description Retrieve user details from the access token
// @Tags auth
// @Produce json
// @Success 200 {object} models.User
// @Failure 401 {object} reqresp.ErrorResponse
// @Router /auth/profile [get]
func (ac *AuthController) GetProfile(c *fiber.Ctx) error {
	user := c.Locals("user").(models.User)

	return c.JSON(fiber.Map{
		"message": "Success",
		"user": fiber.Map{
			"id":    user.ID,
			"name":  user.Name,
			"email": user.Email,
			"phone": user.Phone,
		},
	})
}

// Register godoc
// @Summary User registration
// @Description Register a new user and return the created user data
// @Tags auth
// @Accept json
// @Produce json
// @Param customer body models.UserExample true "Customer data"
// @Success 201 {object} models.User
// @Failure 400 {object} reqresp.ErrorResponse
// @Failure 500 {object} reqresp.ErrorResponse
// @Router /auth/register [post]
func (ac *AuthController) Register(c *fiber.Ctx) error {
	var data models.User

	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&reqresp.ErrorResponse{
			Message: "Invalid request payload",
			Status:  "error",
		})
	}

	var existingUser models.User
	if err := ac.DB.Where("email = ?", data.Email).First(&existingUser).Error; err == nil {
		return c.Status(fiber.StatusBadRequest).JSON(&reqresp.ErrorResponse{
			Message: "Email already exists",
			Status:  "error",
		})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&reqresp.ErrorResponse{
			Message: "Error while hashing password",
			Status:  "error",
		})
	}

	data.Password = string(hashedPassword)

	if err := ac.DB.Create(&data).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&reqresp.ErrorResponse{
			Message: "Error while saving customer",
			Status:  "error",
		})
	}

	return c.JSON(&reqresp.SuccessResponse{
		Status:  "success",
		Message: "Success Create successfully",
		Data:    data,
	})
}

// Logout godoc
// @Summary User logout
// @Description Logout user by clearing the access token cookie
// @Tags auth
// @Produce json
// @Success 200 {object} reqresp.SuccessResponse
// @Router /auth/logout [post]
func (ac *AuthController) Logout(c *fiber.Ctx) error {
	c.Cookie(&fiber.Cookie{
		Name:     "access_token",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	})

	return c.JSON(&reqresp.SuccessResponse{
		Status:  "success",
		Message: "Logged out successfully",
	})
}
