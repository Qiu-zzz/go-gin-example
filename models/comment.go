package models

import "github.com/jinzhu/gorm"

type Comment struct {
	Model

	Content string `json:"content"`
	ArticleId int     `json:"article_id" gorm:"index"`
	CreatedBy    string `json:"created_by"`
	ParentId     int    `json:"parent_id"`
	LikeCount    int    `json:"like_count"`
	DislikeCount int    `json:"dislike_count"`
}

// ExistCommentById checks if a Comment exists based on ID
func ExistCommentById(id int) (bool, error) {
	var comment Comment
	err := db.Select("id").Where("id = ? AND deleted_on = ?", id, 0).First(&comment).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if comment.ID > 0 {
		return true, nil
	}
	return false, nil
}


func GetComments(pageNum int, pageSize int,  articleId int) ([]*Comment, error) {
	var comments []*Comment
	err := db.Where("article_id = ? AND deleted_on = ?",articleId,0).Find(&comments).Error
	if err != nil {
		return nil, err
	}
	return comments, nil
}

func AddComment(data map[string]interface{}) error{
	comment := Comment{
		Content:      data["content"].(string),
		ArticleId:    data["article_id"].(int),
		CreatedBy:    data["created_by"].(string),
		ParentId:     data["parent_id"].(int),
	}
	if err := db.Create(&comment).Error; err != nil {
		return err
	}
	return nil
}

func DeleteComment(id int)error  {
	if err := db.Where("id = ?", id).Delete(Comment{}).Error; err != nil {
		return err
	}
	return nil
}