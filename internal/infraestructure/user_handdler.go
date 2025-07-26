package infraestructure

import (
	"github.com/gin-gonic/gin"
)

type UserHanddlerInterface interface {
	GetUser(c *gin.Context)
}

type UserHanddler struct {
	UserService UserServiceInterface
}

func (u *Userhanddler) GetUser(c *gin.Context) {

}
