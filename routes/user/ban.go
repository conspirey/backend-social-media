package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	mongof "main/functions/mongo"
	"main/functions/sessions"
	"main/structs"
)

func BanUser(c *gin.Context, db *mongo.Database) {
	session := sessions.Default(c)
	user := session.Get("user")
	idToBan := c.Query("id")
	//bannedUntilUnix := c.Query("time")

	if user == nil {
		c.JSON(401, gin.H{
			"error":         "user_not_logged_in",
			"error_message": "User is not logged in",
		})
		return
	}
	userSTR := structs.User{ID: user.(map[string]any)["id"].(string)}
	if err := userSTR.FetchData(db); err != nil {
		c.JSON(400, Error("failed fetching data for admin", "admin_data_failed_fetch"))
		return
	}
	if !userSTR.Admin {
		c.JSON(400, Error("Admin only endpoint", "not_an_admin"))
		return
	}
	if /*bannedUntilUnix == "" ||*/ idToBan == "" {
		c.JSON(400, Error("id query is empty", "empty_id"))
		return
	}
	userToBan := structs.User{ID: idToBan}
	if err := userToBan.FetchData(db); err != nil {
		c.JSON(400, Error("invalid ban id or failed to fetch", "invalid_or_failed_to_fetch"))
		return
	}
	if userToBan.Admin {
		c.JSON(400, Error("you can not ban an admin", "admin_not_bannable"))

		return
	}
	mongof.UpdateOne(bson.M{
		"$set": bson.M{
			"banned": !userToBan.Banned,
		},
	}, bson.M{
		"id": userToBan.ID,
	}, options.Update(), db, "user")
	c.JSON(200, Success(fmt.Sprintf("Succesfully banned '%s'", userToBan.ID)))
	//bannedTimeUnix, err := strconv.Atoi(bannedUntilUnix)
	//if err != nil {
	//	c.JSON(400, Error("invalid unix time at 'time' query", "invalid_time_query"))
	//	return
	//}
	//curTime := time.Now().Unix()
	//if bannedTimeUnix < int(curTime) {
	//	c.JSON(400, Error())
	//	return
	//}
	//fmt.Println(bannedTimeUnix)
	c.String(200, idToBan)

}
func Error(text string, err string) gin.H {
	return gin.H{
		"error":         err,
		"error_message": text,
	}
}
func Success(text string) gin.H {
	return gin.H{
		"success": text,
	}
}
