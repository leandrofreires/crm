package controller

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/leandrofreires/crm/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateArticle(c *gin.Context) {
	var article model.Article
	if err := c.ShouldBind(&article); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := article.Create(); err != nil {
		merr := err.(mongo.WriteException)
		c.JSON(http.StatusBadRequest, gin.H{"error": merr.WriteErrors[0].Message})
		return
	}
	c.JSON(http.StatusCreated, article)
}

//GetArticles is a handler func for router of Article list
func GetArticles(c *gin.Context) {
	var article model.Article
	page, err := strconv.ParseInt(c.Query("page"), 10, 64)
	if err != nil {
		page = 0
	}
	limit, err := strconv.ParseInt(c.Query("limit"), 10, 64)
	if err != nil {
		limit = 5
	}

	articles, err := article.GetArticles(page, limit)
	if err != nil {
		log.Printf("ocorreu um erro ao trazer a lista de articles: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "ocorreu um erro ao trazer a lista de article"})
		return
	}
	c.JSON(http.StatusOK, articles)
}

func GetArticle(c *gin.Context) {
	var article model.Article
	objectID, err := primitive.ObjectIDFromHex(html.EscapeString(c.Param("id")))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "o parametro passado como id não corresponde a um id válido"})
		return
	}
	article.ID = objectID
	if err := article.GetArticle(); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "unexpected error on find this article on database"})
		return
	}
	c.JSON(http.StatusOK, article)
}

func UpdateArticle(c *gin.Context) {
	var article model.Article
	if err := c.ShouldBind(&article); err != nil {
		e := make(map[string]string)

		for _, fieldErr := range err.(validator.ValidationErrors) {
			e[strings.ToLower(fieldErr.Field())] = fmt.Sprintf("validation failed on field condition: %v", fieldErr.ActualTag())
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": e})
		return

	}
	objectID, err := primitive.ObjectIDFromHex(html.EscapeString(c.Param("id")))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "o parametro passado como id não corresponde a um id válido"})
		return
	}
	if objectID != article.ID {
		article.ID = objectID
	}

	if err := article.UpdateArticle(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ocorreu um erro ao salvar a atualização"})
		return
	}
	c.JSON(http.StatusOK, article)
	return
}
