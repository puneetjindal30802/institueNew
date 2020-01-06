package controllers

import (
	"encoding/json"

	"instituteNew/config"
	"instituteNew/helpers"
	"instituteNew/models"
	"instituteNew/utils"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

/*
 *	Function to add new user
 *
 *	Return ResponseController
 */
func SignUpNewUser(c *gin.Context) {
	var userData models.UserSignup
	response := ResponseController{}
	userErr := json.NewDecoder(c.Request.Body).Decode(&userData)
	userData.FullName = userData.FirstName + " " + userData.LastName
	if userErr != nil {
		response := ResponseController{
			config.FailureCode,
			config.FailureFlag,
			"Data is not in the proper format.",
			nil,
		}
		GetResponse(c, response)
		return
	}
	validateErr := ValidateUser(userData)
	if validateErr != "" {
		response := ResponseController{
			config.FailureCode,
			config.FailureFlag,
			validateErr,
			nil,
		}
		GetResponse(c, response)
		return
	}
	// Check user is already registered or not.
	_, err := models.GetRegisteredUser(userData.Email)
	if err == nil {
		response := ResponseController{
			config.FailureCode,
			config.FailureFlag,
			"This email is already exists. Please registered with other one.",
			nil,
		}
		GetResponse(c, response)
		return
	}
	hashedPassword, hashError := bcrypt.GenerateFromPassword([]byte(userData.Password), bcrypt.DefaultCost)
	if hashError != nil {
		response = ResponseController{
			config.FailureCode,
			config.FailureFlag,
			config.FailureMsg,
			nil,
		}
		GetResponse(c, response)
		return
	}
	userData.Password = string(hashedPassword)
	userData.OTPCode = utils.RandomStringToken(6)
	userData.JoinedStatus = 0
	userData.TimeDate = time.Now().Unix()
	saveUserError := models.SaveUserData(userData)
	if saveUserError != nil {
		response = ResponseController{
			config.FailureCode,
			config.FailureFlag,
			"There is something wrong while saving user. Please try it later.",
			nil,
		}
	} else {
		go helpers.SendOtpMail(userData.FullName, userData.Email, userData.OTPCode)
		response = ResponseController{
			config.SuccessCode,
			config.SuccessFlag,
			"Email send successfully.",
			nil,
		}
	}
	GetResponse(c, response)
	return
}

/*
 *	Function to authorize the user
 *
 *	Returns ResponseController
 */
func LogInUser(c *gin.Context) {
	var currUser models.AuthorizeUser
	response := ResponseController{}
	var sessionData models.Session
	userErr := json.NewDecoder(c.Request.Body).Decode(&currUser)
	if userErr != nil {
		response = ResponseController{
			config.FailureCode,
			config.FailureFlag,
			"Data is not in the proper format.",
			nil,
		}
		GetResponse(c, response)
		return
	}
	authorizeUser, authErr := models.GetUserLoginInfo(currUser.Email, currUser.Password)
	if authErr != nil {
		response = ResponseController{
			config.FailureCode,
			config.FailureFlag,
			"Wrong email and password.",
			nil,
		}
		GetResponse(c, response)
		return
	}
	PassErr := bcrypt.CompareHashAndPassword([]byte(authorizeUser.Password), []byte(currUser.Password))
	if PassErr != nil {
		response = ResponseController{
			config.FailureCode,
			config.FailureFlag,
			"Password entered is wrong.",
			nil,
		}
		GetResponse(c, response)
		return
	}
	sessionData.UserId = authorizeUser.ID
	sessionData.CreatedOn = models.GetCurrentDateTimestamp()
	sessionData.LastActivityOn = models.GetCurrentDateTimestamp()
	sessionData.OnetimeAccessToken = utils.RandomStringToken(15)
	if authErr != nil {

	} else {
		response = ResponseController{
			config.SuccessCode,
			config.SuccessFlag,
			config.SuccessMsg,
			authorizeUser,
		}
	}
	GetResponse(c, response)
	return
}

/*
 * Function to create a JWT (JSON Web Token)
 *
 * Returns
 */
// func CreateToken(user models.UserSignup, c *gin.Context) (string, error) {
// 	var ip, userAgent string
// 	keyError := config.InitKeys()
// 	if keyError != nil {
// 		return "", keyError
// 	}

// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.MapClaims{
// 		"email":      user.Email,
// 		"exp":        time.Now().Add(time.Hour * 8760).Unix(),
// 		"role":       user.Role,
// 		"name":       user.FirstName + " " + user.LastName,
// 		"ip":         ip,
// 		"user_agent": userAgent,
// 		"id":         user.Id,
// 	})
// 	config.CurrentUserId = user.Id
// 	models.CurrentUser = user

// 	/* Sign and get the complete encoded token as a string */
// 	tokenString, err := token.SignedString([]byte(config.SignKey))
// 	return tokenString, err
// }

/*
 * Function to valdiate user information while adding/updating
 *
 * Params user type UserSignup
 *
 * Returns errMsg type string
 */
func ValidateUser(user models.UserSignup) string {
	var errMsg string

	// Check first_name if valid.
	if user.FirstName == "" {
		errMsg = "Please enter a valid first name."
		return errMsg
	}

	// Check last_name if valid.
	if user.LastName == "" {
		errMsg = "Please enter a valid last name."
		return errMsg

	}

	// Check email_id if valid.
	if user.Email == "" {
		errMsg = "Please enter a valid email id."
		return errMsg
	}

	//check password if valid
	if user.Password == "" {
		errMsg = "Please enter a valid password"
		return errMsg
	}
	return ""
}

/*
 *	Function  to confirm user account
 *
 *	return ResponseController
 */
func ConfirmationUserAccount(c *gin.Context) {
	response := ResponseController{}
	OTPCode := c.Param("code")
	err := models.UpdateUserJoinedStatus(OTPCode)
	if err != nil {
		response = ResponseController{
			config.FailureCode,
			config.FailureFlag,
			"Something went wrong.",
			nil,
		}
	} else {
		response = ResponseController{
			config.SuccessCode,
			config.SuccessFlag,
			config.SuccessMsg,
			"Successfully joined.",
		}
	}
	GetResponse(c, response)
	return
}

/*
 *	Function to update the user
 *
 *	Retrun responseController
 */
func UpdateUser(c *gin.Context) {

}
