package c_internal

import (
	"encoding/base64"
	"fmt"
	"math/rand"
	"mygo/web/configs"
	"mygo/web/models/m_internal"
	"net/http"
	"strconv"
	"time"
	"unsafe"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

var validate *validator.Validate

type LoginParam struct {
	Email    string `json:"email" binding:"required,email" error_message:"Invalid email format"`
	Password string `json:"password" binding:"required,min=8" error_message:"Password must be at least 8 characters long"`
}

func init() {
	validate = validator.New()
	// rand.Seed(time.Now().UnixNano())
}

var user m_internal.User

// validation in laravel and go lang
// func validateName(name string) error {
//     if len(name) == 0 {
//         return errors.New("Name is required")
//     }
//     if len(name) > 255 {
//         return errors.New("Name is too long")
//     }
//     return nil
// }
// func createUserHandler(w http.ResponseWriter, r *http.Request) {
//     name := r.FormValue("name")
//     if err := validateName(name); err != nil {
//         http.Error(w, err.Error(), http.StatusBadRequest)
//         return
//     }

//     // Other validation checks for email, password, etc.

//	    // Create user if validation passes
//	}
func Login(c *gin.Context) {
	var loginParam LoginParam
	var generateToken string
	// c.ShouldBindJSON(&loginParam)
	// body, err := ioutil.ReadAll(c.Request.Body)
	// if err != nil {
	// 	c.JSON(400, gin.H{"error": err.Error()})
	// 	return
	// }

	// // Check if the request body is empty
	// if len(body) == 0 {
	// 	c.JSON(400, gin.H{"error": "Request body is empty"})
	// 	return
	// }

	if err := c.ShouldBindJSON(&loginParam); err != nil {
		// var validationErrors map[string]string
		validationErrors := make(map[string]string)
		for _, fieldErr := range err.(validator.ValidationErrors) {

			fieldName := fieldErr.Field()
			validationTag := fieldErr.Tag()
			fmt.Println("fieldName:", fieldName)
			fmt.Println("validationTag:", validationTag)

			// Customize error messages based on field and validation rule
			switch fieldName {
			case "Email":
				switch validationTag {
				case "required":
					validationErrors[fieldName] = "Email is required bro"
				case "email":
					validationErrors[fieldName] = "Invalid email format"
				}
			// case "username":
			// 	switch validationTag {
			// 	case "required":
			// 		validationErrors[fieldName] = "Username is required"
			// 	}
			case "password":
				switch validationTag {
				case "required":
					validationErrors[fieldName] = "Password is required"
				case "min":
					validationErrors[fieldName] = "Password must be at least 8 characters long"
				}
			}
		}

		fmt.Println("Error binding JSON:", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"errors": validationErrors})

		// c.JSON(400, gin.H{"errors": err.Error()})
		return
	}
	// fmt.Printf("Name: %s, Email: %s\n", loginParam.Identity, loginParam.Password)
	// result := configs.SQL.Table("internal.users").Where("username = ? OR email = ?", loginParam.Identity, loginParam.Identity).First(&user)
	result := configs.SQL.Table(`internal.users`).Where("email = ?", loginParam.Email).First(&user)
	if result.RowsAffected > 0 {
		if err := bcrypt.CompareHashAndPassword([]byte(*user.Password), []byte(loginParam.Password)); err == nil {
			if generateToken, err = GenerateApiToken(); err == nil {

				updateRes := configs.SQL.Table(`internal.users`).Model(&user).Update("Api_Token", generateToken)
				if updateRes.Error != nil {
					c.JSON(
						http.StatusBadRequest,
						gin.H{
							"ok":      false,
							"data":    nil,
							"message": "Please Login Again",
						},
					)
				} else {
					c.JSON(
						http.StatusOK,
						gin.H{
							"ok": true,
							"data": gin.H{
								"access_token": generateToken,
							},
							"message": "Success login",
						},
					)
				}
			} else {
				c.AbortWithStatusJSON(
					http.StatusBadRequest,
					gin.H{
						"ok":      false,
						"data":    nil,
						"message": err.Error(),
					},
				)
			}
		} else {
			c.AbortWithStatusJSON(
				http.StatusBadRequest,
				gin.H{
					"ok":      false,
					"data":    nil,
					"message": "Password salah",
				},
			)
		}
	} else {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{
				"ok":      false,
				"data":    nil,
				"message": "User not found",
			},
		)
	}
}

// func GenerateJWT(identity string) (string, error) {
// 	token := jwt.New(jwt.SigningMethodHS256)
// 	claims := token.Claims.(jwt.MapClaims)
// 	claims["exp"] = time.Now().Add(24 * 7 * time.Hour).Unix()
// 	claims["identity"] = identity
// 	tokenString, err := token.SignedString([]byte(viper.GetString("SECRET_KEY")))
// 	if err != nil {
// 		return "", err
// 	}
// 	return tokenString, nil
// }

func GenerateApiToken() (string, error) {
	token := RandStringBytesMaskImprSrcUnsafe(30) + strconv.FormatInt(time.Now().UTC().Add(24*7*time.Hour).UnixNano()/int64(time.Millisecond), 10) + base64.StdEncoding.EncodeToString([]byte(user.Email))
	// claims := token.Claims.(jwt.MapClaims)
	// claims["exp"] = time.Now().Add(24 * 7 * time.Hour).Unix()
	// claims["identity"] = identity
	// tokenString, err := token.SignedString([]byte(viper.GetString("SECRET_KEY")))
	// if err != nil {
	// 	return "", err
	// }
	return token, nil
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

var src = rand.NewSource(time.Now().UnixNano())

func RandStringBytesMaskImprSrcUnsafe(n int) string {
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return *(*string)(unsafe.Pointer(&b))
}
