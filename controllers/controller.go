package controllers

import (
	"fmt"
	"net/http"
	"tourism/db"
	"tourism/model"

	"github.com/gin-gonic/gin"
)

func DetailsCreate(ctx *gin.Context) {
	var body struct {
		Name     string `99`
		Email    string `json:"email" binding:"required,email"`
		Language string `json:"language" binding:"required"`
	}
	err := ctx.ShouldBind(&body)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Message": "Enter Right Data"})
		return
	}

	post := model.Tourism{Name: body.Name, Email: body.Email, Language: body.Language}
	result := db.DB.Create(&post)
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Error creating post"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Details Added Successfully", "Post:": post})

}

func GetAllDetails(ctx *gin.Context) {
	var detail []model.Tourism
	db.DB.Find(&detail)
	ctx.JSON(http.StatusOK, gin.H{
		"message": "All Details",
		"Details": detail,
	})

}

func FindOneDetail(ctx *gin.Context) {
	var detail model.Tourism
	id := ctx.Param("id")

	db.DB.First(&detail, id)
	err := db.DB.First(&detail, id).Error

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"Message": "Details Not Found",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"Message": "Details Found",
		"Post":    detail,
	})
}
func UpdateDetails(ctx *gin.Context) {
	id := ctx.Param("id")
	var input struct {
		Name     string `json:"name" `
		Email    string `json:"email" `
		Language string `json:"language" `
	}

	ctx.Bind(&input)

	var det model.Tourism

	if err := db.DB.First(&det, id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"Error": "Details not found"})
		return
	}

	if err := db.DB.Model(&det).Updates(&model.Tourism{Name: input.Name, Language: input.Language, Email: input.Email}).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update details"})
		return
	}

	s := fmt.Sprintf("Details updated of id = %v", id)
	ctx.JSON(http.StatusOK, gin.H{"message": s, "post": det})
}

func DetailsDelete(ctx *gin.Context) {
	id := ctx.Param("id")

	if (db.DB.First(&model.Tourism{}, id).RowsAffected == 0) {
		ctx.JSON(404, gin.H{"message": "Details not found"})
		return
	}

	db.DB.Delete(&model.Tourism{}, id)
	s := fmt.Sprintf("Post deleted successfully of id %v", id)
	ctx.JSON(200, gin.H{"Message": s})
}
