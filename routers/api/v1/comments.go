package v1

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/go-gin-example/models"
	"github.com/go-gin-example/pkg/app"
	"github.com/go-gin-example/pkg/e"
	"github.com/go-gin-example/pkg/setting"
	"github.com/go-gin-example/pkg/util"
	"github.com/go-gin-example/service/article_service"
	"github.com/go-gin-example/service/comment_service"
	"github.com/unknwon/com"
	"net/http"
)

// @Summary Get multiple comments
// @Produce  json
// @Param article_id body int false "ArticleId"
// @Param parent_id body int false "ParentId"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /comments [get]
func GetComments(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}

	articleId := com.StrTo(c.Query("articleId")).MustInt()
	parentId := com.StrTo(c.Query("parentId")).MustInt()
	commentService := comment_service.Comment{
		ArticleId: articleId,
		ParentId:  parentId,
		PageNum:   util.GetPage(c),
		PageSize:  setting.AppSetting.PageSize,
	}
	ok, _ := valid.Valid(&commentService)
	if !ok {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}
	var (
		comments []*models.Comment
		err      error
	)
	comments, err = commentService.GetAll()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_COMMENTS_FAIL, nil)
		return
	}
	data := make(map[string]interface{})
	data["lists"] = comments
	appG.Response(http.StatusOK, e.SUCCESS, data)
}

type AddCommentForm struct {
	Content   string `form:"content" valid:"Required;MaxSize(255)"`
	CreatedBy string `form:"created_by" valid:"Required;MaxSize(100)"`
	ParentId int `form:"parent_id" `
	ArticleId int `form:"article_id" valid:"Required;Min(1)"`
}
// @Summary Get multiple comments
// @Produce  json
// @Param article_id body int false "ArticleId"
// @Param parent_id body int false "ParentId"
// @Param content body string false "Content"
// @Param created_by body string false "CreatedBy"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /comments [post]
func AddComments(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form AddCommentForm
	)
	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}
	articleService := article_service.Article{ID: form.ArticleId}
	exists, err := articleService.ExistByID()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_CHECK_EXIST_ARTICLE_FAIL, nil)
		return
	}
	if !exists {
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_ARTICLE, nil)
		return
	}
	commentService := comment_service.Comment{
		Content:   form.Content,
		ParentId:  form.ParentId,
		ArticleId: form.ArticleId,
		CreatedBy: form.CreatedBy,
	}
	if err := commentService.Add(); err !=nil{
		appG.Response(http.StatusInternalServerError, e.ERROR_ADD_COMMENT_FAIL, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, nil)
}



// @Summary delete comment
// @Produce  json
// @Param id path int true "ID"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /comments [delete]
func DeleteComment(c *gin.Context){
	appG := app.Gin{C: c}
	valid := validation.Validation{}
	id := com.StrTo(c.Param("id")).MustInt()
	valid.Min(id, 1, "id").Message("ID必须大于0")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusOK, e.INVALID_PARAMS, nil)
		return
	}
	commentService := comment_service.Comment{
		ID:   id,
	}
	if err := commentService.Delete(); err !=nil{
		appG.Response(http.StatusInternalServerError, e.ERROR_ADD_COMMENT_FAIL, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, nil)
}