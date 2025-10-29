package service

import (
	"blog-go/global"
	"blog-go/model/database"

	"gorm.io/gorm"
)

func (commentService *CommentService) LoadChildren(comment *database.Comment) error {
	var children []database.Comment
	if err := global.DB.Where("p_id = ?", comment.ID).Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("uuid, username, avatar, address, signature")
	}).Find(&children).Error; err != nil {
		return err
	}

	for i := range children {
		if err := commentService.LoadChildren(&children[i]); err != nil {
			return err
		}
	}
	comment.Children = children
	return nil
}

func (commentService *CommentService) DeleteCommentChildren(tx *gorm.DB, commentId uint) error {
	var children []database.Comment
	if err := tx.Where("p_id = ?", commentId).Find(&children).Error; err != nil {
		return nil
	}

	for _, child := range children {
		if err := commentService.DeleteCommentChildren(tx, child.ID); err != nil {
			return err
		}
	}

	if err := tx.Delete(&database.Comment{}, commentId).Error; err != nil {
		return err
	}

	return nil
}

func (commentService *CommentService) FindChildCommentsIDByRootCommentUserUUID(comments []database.Comment) map[uint]struct{} {
	result := make(map[uint]struct{})

	for _, rootComment := range comments {
		var findChildren func([]database.Comment)

		findChildren = func(children []database.Comment) {
			for _, child := range children {
				if child.UserUUID == rootComment.UserUUID {
					result[child.ID] = struct{}{}
				}

				if len(child.Children) > 0 {
					findChildren(child.Children)
				}
			}
		}

		findChildren(rootComment.Children)
	}

	return result
}
