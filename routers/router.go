package routers

import (
	"github.com/go-gin-example/middleware/cors"
	"net/http"

	"github.com/gin-gonic/gin"

	_ "github.com/go-gin-example/docs"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	"github.com/go-gin-example/middleware/jwt"
	"github.com/go-gin-example/pkg/export"
	"github.com/go-gin-example/pkg/qrcode"
	"github.com/go-gin-example/pkg/upload"
	"github.com/go-gin-example/routers/api"
	"github.com/go-gin-example/routers/api/v1"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.StaticFS("/export", http.Dir(export.GetExcelFullPath()))
	r.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))
	r.StaticFS("/qrcode", http.Dir(qrcode.GetQrCodeFullPath()))

	r.POST("/login", api.GetAuth)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.POST("/upload", api.UploadImage)

	r.GET("/ping", func(context *gin.Context) {
		context.JSON(200,"msg:ok")
	})



	apiv1 := r.Group("/api/v1")
	apiv1.Use(jwt.JWT())
	apiv1.Use(cors.Cors())
	{
		//获取标签列表
		r.GET("/tags", v1.GetTags)
		//新建标签
		apiv1.POST("/tags", v1.AddTag)
		//更新指定标签
		apiv1.PUT("/tags/:id", v1.EditTag)
		//删除指定标签
		apiv1.DELETE("/tags/:id", v1.DeleteTag)
		//导出标签
		r.POST("/tags/export", v1.ExportTag)
		//导入标签
		r.POST("/tags/import", v1.ImportTag)

		//获取文章列表
		r.GET("/articles", v1.GetArticles)
		//获取指定文章
		r.GET("/articles/:id", v1.GetArticle)
		//新建文章
		apiv1.POST("/articles", v1.AddArticle)
		//更新指定文章
		apiv1.PUT("/articles/:id", v1.EditArticle)
		//删除指定文章
		apiv1.DELETE("/articles/:id", v1.DeleteArticle)

		// 评论测试
		r.GET("/comment",v1.GetComments)
		r.POST("/comment",v1.AddComments)
		r.DELETE("/comment/:id",v1.DeleteComment)

		//生成文章海报
		apiv1.POST("/articles/poster/generate", v1.GenerateArticlePoster)
	}

	return r
}
