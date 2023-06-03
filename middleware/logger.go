package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type Middle struct {
}

func init() {
	fmt.Println("this is middleware.")
}

func (m Middle) LoggerToFile() gin.HandlerFunc {
	return nil
}
