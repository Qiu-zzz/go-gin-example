package comment_service

import (
	"github.com/go-gin-example/models"
	"github.com/jinzhu/gorm"
)

type Comment struct {
	ID            int

	Content       string
	ParentId      int
	ArticleId     int
	CreatedBy     string
	ModifiedBy    string

	PageNum  int
	PageSize int
}

func (c *Comment) GetAll()([]*models.Comment,error) {
	var comments []*models.Comment

	comments, err := models.GetComments(c.PageNum, c.PageSize, c.getMaps(),c.ArticleId,c.ParentId)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return comments, nil
}

func (c *Comment) getMaps() map[string]interface{} {
	maps := make(map[string]interface{})
	maps["deleted_on"] = 0
	if c.ArticleId != -1 {
		maps["article_id"] = c.ArticleId
	}
	if c.ParentId != -1 {
		maps["parent_id"] = c.ParentId
	}

	return maps
}

func (c *Comment) Add() error {
	comment := map[string]interface{}{
		"content": c.Content,
		"created_by": c.CreatedBy,
		"article_id": c.ArticleId,
		"parent_id": c.ParentId,
	}
	if err :=models.AddComment(comment); err !=nil {
		return err
	}
	return nil
}
