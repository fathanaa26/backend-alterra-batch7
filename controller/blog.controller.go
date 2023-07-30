package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"

	"blog-app/model"
)

func BlogIdHandler(r *gin.Context) {
	id := r.Param("id")
	isAvl := r.Query("isAvl")
	r.JSON(
		http.StatusOK,
		gin.H{
			"id":     id,
			"isAvl":  isAvl,
			"title":  "Manusia Bumi",
			"author": "Pidie Baiq",
			"text":   "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.",
		})
}

func BlogsHandler(r *gin.Context) {
	r.JSON(
		http.StatusOK,
		gin.H{
			"id":     1001,
			"title":  "Berulah",
			"author": "Mr. Boston",
			"text":   "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.",
		})
}

func AddBlogHandler(r *gin.Context) {
	var payloadProp model.Blog

	err := r.ShouldBindJSON(&payloadProp)
	if err != nil {

		errMsgStack := []string{}

		for _, e := range err.(validator.ValidationErrors) {
			errMsg := fmt.Sprintf("Error on %s field: %s", e.Field(), e.ActualTag())
			errMsgStack = append(errMsgStack, errMsg)
		}

		r.JSON(http.StatusBadRequest, gin.H{
			"errors": errMsgStack,
		})

		return
	}

	r.JSON(
		http.StatusCreated,
		gin.H{
			"title":     payloadProp.Title,
			"author":    payloadProp.Author,
			"text":      payloadProp.Text,
			"stock":     payloadProp.Stock,
			"createdAt": payloadProp.CreatedAt,
		})

}
