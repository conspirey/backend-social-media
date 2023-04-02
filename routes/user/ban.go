package user
import (
	"fmt"
	mongof "main/functions/mongo"
	"main/functions/sessions"
	"os"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
func BanUser(c *gin.Context, db *mongo.Database) {

}