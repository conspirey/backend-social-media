package updatedata

import (
	"os"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func UpdateUserData(c *gin.Context, db *mongo.Database) {
	c.String(200, os.Getenv("ENCRYPTION_KEY"))
}