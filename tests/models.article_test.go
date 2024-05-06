package tests

import (
	"github.com/hewo233/PageList/model"
	"testing"
)

func TestGetAllArticles(t *testing.T) {

	alist := model.GetAllArticles()
	if len(alist) != len(model.ArticleList) {
		t.Fail()
	}

	for i, v := range alist {
		if v.Content != model.ArticleList[i].Content ||
			v.ID != model.ArticleList[i].ID ||
			v.Title != model.ArticleList[i].Title {
			t.Fail()
			break
		}
	}

}
