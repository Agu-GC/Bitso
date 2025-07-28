package infraestructure

import (
	"github.com/Agu-GC/Bitso/internal/application"
	"github.com/gin-gonic/gin"
)

type UserHanddlerInterface interface {
	GetUser(c *gin.Context)
}

type UserHanddler struct {
	UserService application.UserServiceInterface
}

func (u *UserHanddler) GetUser(c *gin.Context) {
	id := c.Param("id")
	surname := c.Query("surname")
	user := u.UserService.GetUserInfo(id, surname)

	if user.Id == "" {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}

	c.JSON(200, user)
}
