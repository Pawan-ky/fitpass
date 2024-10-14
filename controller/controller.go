package controller

import (
	db "fitpass/database"
	"fitpass/models"
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserPlan struct {
	UserID       int `json:"user_id"`       
	PlanDuration int `json:"plan_duration"` 
}

func GetUsers(ctx *gin.Context) {
	var users []models.User
	db.Instance.Find(&users)
	ctx.JSON(200, users)
}

func AddUser(ctx *gin.Context) {
	var req models.User
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(400, "invalid request")
	}
	db.Instance.Create(&req)
	ctx.JSON(200, "created successfully")
}

func DeleteUser(ctx *gin.Context) {
	id := ctx.Query("id")
	num, _ := strconv.ParseUint(id, 10, 64)
	fmt.Println(num)
	db.Instance.Where("id =?", num).Delete(&models.User{})
	ctx.JSON(200, "deleted succesfully")
}

func AddUserPlan(ctx *gin.Context) {
	var req UserPlan
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(400, "invalid request")
	}

	validDurations := map[int]string{
		1:  models.Monthly,
		3:  models.TriMonthly,
		6:  models.HalfYearly,
		12: models.Yearly,
	}

	planName, valid := validDurations[req.PlanDuration]
	if !valid {
		ctx.JSON(400, gin.H{"error": "invalid plan duration, must be 1, 3, 6, or 12 months"})
		return
	}

	var user models.User
	if err := db.Instance.First(&user, req.UserID).Error; err != nil {
		ctx.JSON(404, gin.H{"error": "user not found"})
		return
	}
	expiryDate := time.Now().AddDate(0, 0, req.PlanDuration*30)

	// Create the new subscription
	subscription := models.Subscription{
		PlanName:  planName,
		ExpiresAt: expiryDate,
		UserID:    uint(req.UserID),
	}

	// Save the subscription to the database
	if err := db.Instance.Create(&subscription).Error; err != nil {
		ctx.JSON(500, gin.H{"error": "failed to create subscription"})
		return
	}

	// Return a success response
	ctx.JSON(200, gin.H{"message": "subscription added successfully", "subscription": subscription})

}

func GetUserSubscription(ctx *gin.Context) {
    userIDParam := ctx.Query("user_id")
    id,_ := strconv.ParseUint(userIDParam, 10, 64)
	fmt.Println("id. ------", id)
	userID := uint(id)
    var subscription models.Subscription
    if err := db.Instance.Where("user_id = ?", userID).Preload("User").First(&subscription).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            ctx.JSON(404, gin.H{"error": "no subscription found for this user"})
        } else {
            ctx.JSON(500, gin.H{"error": "failed to fetch subscription"})
        }
        return
    }
    ctx.JSON(200, subscription)
}
