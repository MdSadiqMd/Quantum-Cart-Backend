package helpers

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/MdSadiqMd/Quantum-Cart-Backend/internal/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type Auth struct {
	Secret string
}

func NewAuth(secret string) Auth {
	return Auth{
		Secret: secret,
	}
}

func (auth Auth) CreateHashedPassowrd(password string) (string, error) {
	if len(password) == 0 {
		return "", errors.New("password is required")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", errors.New("failed to hash password")
	}
	return string(hashedPassword), nil
}

func (auth Auth) GenerateToken(id uint, email string, role string) (string, error) {
	if id == 0 || len(email) == 0 || len(role) == 0 {
		return "", errors.New("failed to generate token")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": id,
		"email":   email,
		"role":    role,
		"expiry":  time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(auth.Secret))
	if err != nil {
		return "", errors.New("unable to Sign the Token")
	}
	return tokenString, nil
}

func (auth Auth) VerifyPassword(password string, hashedPassword string) error {
	if len(password) == 0 || len(hashedPassword) == 0 {
		return errors.New("password is required")
	}

	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return errors.New("password does not match")
	}
	return nil
}

func (auth Auth) VerifyToken(token string) (models.User, error) {
	tokenHeader := strings.Split(token, " ")
	if len(tokenHeader) != 2 {
		return models.User{}, errors.New("invalid token")
	}

	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(tokenHeader[1], claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(auth.Secret), nil
	})
	if err != nil {
		return models.User{}, errors.New("invalid token")
	}

	if claims["expiry"].(float64) < float64(time.Now().Unix()) {
		return models.User{}, errors.New("token expired")
	}

	user := models.User{
		Id:       uint(claims["user_id"].(float64)),
		Email:    claims["email"].(string),
		UserType: claims["role"].(string),
	}
	return user, nil
}

func (auth Auth) CurrentUser(ctx *fiber.Ctx) error {
	authHeader := ctx.GetReqHeaders()["Authorization"][0]

	user, err := auth.VerifyToken(authHeader)
	if err == nil && user.Id > 0 {
		ctx.Locals("user", user)
		return ctx.Next()
	} else {
		return ctx.Status(http.StatusUnauthorized).JSON(&fiber.Map{
			"error": err,
		})
	}
}

func (auth Auth) GetCurrentUser(ctx *fiber.Ctx) models.User {
	user := ctx.Locals("user")
	return user.(models.User)
}

func (auth Auth) Authorize(ctx *fiber.Ctx) error {
	authHeader := ctx.GetReqHeaders()["Authorization"][0]
	user, err := auth.VerifyToken(authHeader)
	if err == nil && user.Id > 0 {
		ctx.Locals("user", user)
		return ctx.Next()
	} else {
		return ctx.Status(http.StatusUnauthorized).JSON(&fiber.Map{
			"error": err,
		})
	}
}

func (auth Auth) AuthorizeSeller(ctx *fiber.Ctx) error {
	authHeader := ctx.GetReqHeaders()["Authorization"][0]
	user, err := auth.VerifyToken(authHeader)

	if err != nil {
		return ctx.Status(http.StatusUnauthorized).JSON(&fiber.Map{
			"error": err,
			"message": "Unauthorized",
		})
	} else if user.Id > 0 && user.UserType == "seller" {
		ctx.Locals("user", user)
		return ctx.Next()
	} else {
		return ctx.Status(http.StatusUnauthorized).JSON(&fiber.Map{
			"error": err,
			"message": "Please login as a seller",
		})
	}
}

func (auth Auth) GenerateCode() (int, error) {
	return RandomNumbers(6)
}
