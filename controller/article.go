package controller

import (
	"html"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
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
		s := strings.SplitN(err.Error(), "\n", -1)
		c.JSON(http.StatusBadRequest, s)
		return
	}
	c.JSON(http.StatusOK, article)
	return
	// objectID, err := primitive.ObjectIDFromHex(html.EscapeString(c.Param("id")))
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "o parametro passado como id não corresponde a um id válido"})
	// 	return
	// }

}
