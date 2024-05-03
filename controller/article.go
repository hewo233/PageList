package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/hewo233/PageList/model"
	"strconv"
)

func ShowIndexPage(c *gin.Context) {
	articles := model.GetAllArticles()
	c.HTML(200, "index.html", gin.H{
		"title":   "Home Page",
		"payload": articles,
	})
}

func GetArticleByID(id int) (*model.Article, error) {
	for _, a := range model.ArticleList {
		if a.ID == id {
			return &a, nil
		}
	}
	return nil, errors.New("article not found")
}

func ShowArticlePage(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)
	article, err := GetArticleByID(id)
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.HTML(200, "article.html", gin.H{
		"title":   article.Title,
		"payload": article,
	})
}
