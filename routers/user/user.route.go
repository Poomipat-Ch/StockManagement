package user

import (
	"strconv"

	"github.com/Poomipat-Ch/StockManagement/dto"
	"github.com/Poomipat-Ch/StockManagement/dto/common"
	"github.com/Poomipat-Ch/StockManagement/models"
	"github.com/Poomipat-Ch/StockManagement/routers"
	"github.com/Poomipat-Ch/StockManagement/services"
	"github.com/gin-gonic/gin"
)

type userRouter struct {
	r *gin.RouterGroup
	s services.UserService
}

// MapRoutes implements routers.Routers.
func (u *userRouter) MapRoutes() {
	u.r.GET("", u.getUsers)
	u.r.POST("", u.createUser)
	u.r.GET("/calculate", u.calculate)
}

func NewUserRouter(r *gin.RouterGroup, s services.UserService) routers.Routers {
	return &userRouter{
		r: r.Group("/user"),
		s: s,
	}
}

func (u *userRouter) getUsers(c *gin.Context) {
	users, err := u.s.GetUsers()

	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, &common.Response[[]*models.User]{Message: "Success", Data: users})
}

func (u *userRouter) createUser(c *gin.Context) {
	var userRequest dto.CreateUserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	user, validateErrs, err := u.s.CreateUser(&userRequest)

	if validateErrs != nil {
		c.JSON(400, gin.H{"message": validateErrs})
		return
	}

	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, user)
}

func (u *userRouter) calculate(c *gin.Context) {
	// get number from query string
	number := c.Query("number")

	numberInt, err := strconv.ParseInt(number, 10, 64)

	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
	}

	result := sum(numberInt)

	c.JSON(200, gin.H{"result": result})
}

func sum(number int64) int64 {
	var result int64 = 0

	for i := int64(0); i < number; i++ {
		result += i
	}

	return result
}
