package service

import (
	"blog-go/global"
	"blog-go/model/appTypes"
	"blog-go/model/database"
	"blog-go/model/elasticsearch"
	"blog-go/model/other"
	"blog-go/model/request"
	"blog-go/utils"
	"context"
	"errors"
	"strconv"
	"time"

	"github.com/elastic/go-elasticsearch/v8/typedapi/core/search"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/scriptlanguage"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/sortorder"
	"gorm.io/gorm"
)

type ArticleService struct {
}

func (articleService *ArticleService) ArticleInfoByID(id string) (elasticsearch.Article, error) {
	go func() {
		articleView := articleService.NewArticleView()
		_ = articleView.Set(id)
	}()
	return articleService.Get(id)
}
func (articleService *ArticleService) ArticleSearch(req request.ArticleSearch) (interface{}, int64, error) {
	searchReq := &search.Request{
		Query: &types.Query{},
	}

	boolReq := &types.BoolQuery{}

	if req.Query != "" {
		boolReq.Should = []types.Query{
			{Match: map[string]types.MatchQuery{"title": {Query: req.Query}}},
			{Match: map[string]types.MatchQuery{"keyword": {Query: req.Query}}},
			{Match: map[string]types.MatchQuery{"abstract": {Query: req.Query}}},
			{Match: map[string]types.MatchQuery{"content": {Query: req.Query}}},
		}
	}

	if req.Tag != "" {
		boolReq.Must = []types.Query{
			{Match: map[string]types.MatchQuery{"tags": {Query: req.Tag}}},
		}
	}

	if req.Category != "" {
		boolReq.Filter = []types.Query{
			{Term: map[string]types.TermQuery{"category": {Value: req.Category}}},
		}
	}

	if boolReq.Filter != nil || boolReq.Must != nil || boolReq.Should != nil {
		searchReq.Query.Bool = boolReq
	} else {
		searchReq.Query.MatchAll = &types.MatchAllQuery{}
	}

	if req.Sort != "" {
		var sortField string
		switch req.Sort {
		case "time":
			sortField = "created_at"
		case "view":
			sortField = "views"
		case "comment":
			sortField = "comments"
		case "like":
			sortField = "likes"
		default:
			sortField = "created_at"
		}

		var order sortorder.SortOrder

		if req.Order != "asc" {
			order = sortorder.Desc
		} else {
			order = sortorder.Asc
		}

		searchReq.Sort = []types.SortCombinations{
			types.SortOptions{
				SortOptions: map[string]types.FieldSort{
					sortField: {Order: &order},
				},
			},
		}
	}

	option := other.EsOption{
		PageInfo:       req.PageInfo,
		Index:          elasticsearch.ArticleIndex(),
		Request:        searchReq,
		SourceIncludes: []string{"created_at", "cover", "title", "abstract", "category", "tags", "views", "comments", "likes"},
	}

	return utils.EsPagination(context.TODO(), option)

}

func (articleService *ArticleService) ArticleCategory() ([]database.ArticleCategory, error) {
	var categorys []database.ArticleCategory
	if err := global.DB.Find(&categorys).Error; err != nil {
		return nil, err
	}
	return categorys, nil
}

func (articleService *ArticleService) ArticleTags() ([]database.ArticleTag, error) {
	var tags []database.ArticleTag
	if err := global.DB.Find(&tags).Error; err != nil {
		return nil, err
	}
	return tags, nil
}

func (articleService *ArticleService) ArticleLike(req request.ArticleLike) error {
	return global.DB.Transaction(func(tx *gorm.DB) error {
		var al database.ArticleLike
		var num int
		if errors.Is(tx.Where("user_id = ? AND article_id = ?", req.UserID, req.ArticleID).First(&al).Error, gorm.ErrRecordNotFound) {
			if err := tx.Create(&database.ArticleLike{UserID: req.UserID, ArticleID: req.ArticleID}).Error; err != nil {
				return err
			}
			num = 1
		} else {
			if err := tx.Delete(&al).Error; err != nil {
				return err
			}
			num = -1
		}
		source := "ctx._source.likes += " + strconv.Itoa(num)
		script := types.Script{Source: &source, Lang: &scriptlanguage.Painless}
		_, err := global.ESClient.Update(elasticsearch.ArticleIndex(), req.ArticleID).Script(&script).Do(context.TODO())
		return err
	})
}

func (articleService *ArticleService) ArticleIsLike(req request.ArticleLike) (bool, error) {
	err := global.DB.Where("user_id = ? AND article_id = ?", req.UserID, req.ArticleID).
		First(&database.ArticleLike{}).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	}
	if err != nil {
		return false, err
	}

	return true, nil
}

func (articleService *ArticleService) ArticleLikesList(req request.ArticleLikesList) (interface{}, int64, error) {
	db := global.DB.Where("user_id = ?", req.UserID)
	option := other.MySQLOption{
		PageInfo: req.PageInfo,
		Where:    db,
	}
	// 拿到指定用户id的所有收藏的文章
	l, total, err := utils.MySQLPagination(&database.ArticleLike{}, option)
	if err != nil {
		return nil, 0, err
	}
	var list []struct {
		Id_     string                `json:"_id"`
		Source_ elasticsearch.Article `json:"_source"`
	}
	for _, articleLike := range l {
		article, err := articleService.Get(articleLike.ArticleID)
		if err != nil {
			return nil, 0, err
		}
		article.UpdatedAt = ""
		article.Keyword = ""
		article.Content = ""
		list = append(list, struct {
			Id_     string                `json:"_id"`
			Source_ elasticsearch.Article `json:"_source"`
		}{
			Id_:     articleLike.ArticleID,
			Source_: article,
		})
	}
	return list, total, nil
}

func (articleService *ArticleService) ArticleCreate(req request.ArticleCreate) error {
	b, err := articleService.Exits(req.Title)
	if err != nil {
		return err
	}
	if b {
		return errors.New("the article already exists")
	}
	now := time.Now().Format("2006-01-02 15:04:05")
	articleToCreate := elasticsearch.Article{
		CreatedAt: now,
		UpdatedAt: now,
		Cover:     req.Cover,
		Title:     req.Title,
		Keyword:   req.Title,
		Category:  req.Category,
		Tags:      req.Tags,
		Abstract:  req.Abstract,
		Content:   req.Content,
	}

	return global.DB.Transaction(func(tx *gorm.DB) error {
		if err := articleService.UpdateCategoryCount(tx, "", articleToCreate.Category); err != nil {
			return err
		}
		if err := articleService.UpdateTagsCount(tx, []string{}, articleToCreate.Tags); err != nil {
			return err
		}
		if err := utils.ChangeImagesCategory(tx, []string{articleToCreate.Cover}, appTypes.Cover); err != nil {
			return err
		}
		illustrations, err := utils.FindIllustrations(articleToCreate.Content)
		if err != nil {
			return err
		}
		if err := utils.ChangeImagesCategory(tx, illustrations, appTypes.Illustration); err != nil {
			return err
		}
		return articleService.Create(&articleToCreate)
	})
}

func (articleService *ArticleService) ArticleDelete(req request.ArticleDelete) error {
	if len(req.IDs) == 0 {
		return nil
	}
	return global.DB.Transaction(func(tx *gorm.DB) error {
		for _, id := range req.IDs {
			articleToDelete, err := articleService.Get(id)
			if err != nil {
				return err
			}
			if err := articleService.UpdateCategoryCount(tx, articleToDelete.Category, ""); err != nil {
				return err
			}
			if err := articleService.UpdateTagsCount(tx, articleToDelete.Tags, []string{}); err != nil {
				return err
			}
			if err := utils.InitImagesCategory(tx, []string{articleToDelete.Cover}); err != nil {
				return err
			}
			illustrations, err := utils.FindIllustrations(articleToDelete.Content)
			if err != nil {
				return err
			}
			if err := utils.InitImagesCategory(tx, illustrations); err != nil {
				return err
			}
			comments, err := ServiceGroupApp.CommentService.CommentInfoByArticleID(request.CommentInfoByArticleID{ArticleID: id})
			if err != nil {
				return err
			}
			for _, comment := range comments {
				if err := ServiceGroupApp.CommentService.DeleteCommentChildren(tx, comment.ID); err != nil {
					return err
				}
			}
		}
		return articleService.Delete(req.IDs)
	})
}

func (articleService *ArticleService) ArticleUpdate(req request.ArticleUpdate) error {
	now := time.Now().Format("2006-01-02 15:04:05")
	articleToUpdate := struct {
		UpdatedAt string   `json:"updated_at"` // Last update time
		Cover     string   `json:"cover"`      // Cover image
		Title     string   `json:"title"`      // Title
		Keyword   string   `json:"keyword"`    // Keywords
		Category  string   `json:"category"`   // Category
		Tags      []string `json:"tags"`       // Tags
		Abstract  string   `json:"abstract"`   // Abstract
		Content   string   `json:"content"`    // Content
	}{
		UpdatedAt: now,
		Cover:     req.Cover,
		Title:     req.Title,
		Keyword:   req.Title,
		Category:  req.Category,
		Tags:      req.Tags,
		Abstract:  req.Abstract,
		Content:   req.Content,
	}
	return global.DB.Transaction(func(tx *gorm.DB) error {
		oldArticle, err := articleService.Get(req.ID)
		if err != nil {
			return err
		}
		if err := articleService.UpdateCategoryCount(tx, oldArticle.Category, articleToUpdate.Category); err != nil {
			return err
		}
		if err := articleService.UpdateTagsCount(tx, oldArticle.Tags, articleToUpdate.Tags); err != nil {
			return err
		}
		if articleToUpdate.Cover != oldArticle.Cover {
			if err := utils.InitImagesCategory(tx, []string{oldArticle.Cover}); err != nil {
				return err
			}
			if err := utils.ChangeImagesCategory(tx, []string{articleToUpdate.Cover}, appTypes.Cover); err != nil {
				return err
			}
		}
		oldIllustrations, err := utils.FindIllustrations(oldArticle.Content)
		if err != nil {
			return err
		}
		newIllustrations, err := utils.FindIllustrations(articleToUpdate.Content)
		if err != nil {
			return err
		}
		addedIllustrations, removedIllustrations := utils.DiffArrays(oldIllustrations, newIllustrations)
		if err := utils.InitImagesCategory(tx, removedIllustrations); err != nil {
			return err
		}
		if err := utils.ChangeImagesCategory(tx, addedIllustrations, appTypes.Illustration); err != nil {
			return err
		}
		return articleService.Update(req.ID, articleToUpdate)
	})
}

func (articleService *ArticleService) ArticleList(req request.ArticleList) (interface{}, int64, error) {
	searchReq := &search.Request{
		Query: &types.Query{},
	}

	boolQuery := &types.BoolQuery{}

	if req.Title != nil {
		boolQuery.Must = append(boolQuery.Must, types.Query{Match: map[string]types.MatchQuery{"title": {Query: *req.Title}}})
	}

	if req.Abstract != nil {
		boolQuery.Must = append(boolQuery.Must, types.Query{Match: map[string]types.MatchQuery{"abstract": {Query: *req.Abstract}}})
	}

	if req.Category != nil {
		boolQuery.Filter = []types.Query{
			{
				Term: map[string]types.TermQuery{
					"category": {Value: req.Category},
				},
			},
		}
	}

	if boolQuery.Must != nil || boolQuery.Filter != nil {
		searchReq.Query.Bool = boolQuery
	} else {
		searchReq.Query.MatchAll = &types.MatchAllQuery{}
		searchReq.Sort = []types.SortCombinations{
			types.SortOptions{
				SortOptions: map[string]types.FieldSort{
					"created_at": {Order: &sortorder.Desc},
				},
			},
		}
	}

	option := other.EsOption{
		PageInfo: req.PageInfo,
		Index:    elasticsearch.ArticleIndex(),
		Request:  searchReq,
	}
	return utils.EsPagination(context.TODO(), option)
}
