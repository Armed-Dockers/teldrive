package routes

import (
	"net/http"

	"github.com/divyam234/teldrive/database"
	"github.com/divyam234/teldrive/services"
	"github.com/divyam234/teldrive/utils"

	"github.com/gin-gonic/gin"
)

func addFileRoutes(rg *gin.RouterGroup) {

	r := rg.Group("/files")
	r.Use(Authmiddleware)
	fileService := services.FileService{Db: database.DB, ChannelID: utils.GetConfig().ChannelID}

	r.GET("", func(c *gin.Context) {
		res, err := fileService.ListFiles(c)

		if err != nil {
			c.AbortWithError(err.Code, err.Error)
			return
		}

		c.JSON(http.StatusOK, res)
	})

	r.POST("", func(c *gin.Context) {

		res, err := fileService.CreateFile(c)

		if err != nil {
			c.AbortWithError(err.Code, err.Error)
			return
		}

		c.JSON(http.StatusOK, res)
	})

	r.GET("/:fileID", func(c *gin.Context) {

		res, err := fileService.GetFileByID(c)

		if err != nil {
			c.AbortWithError(http.StatusNotFound, err)
			return
		}

		c.JSON(http.StatusOK, res)
	})

	r.PATCH("/:fileID", func(c *gin.Context) {

		res, err := fileService.UpdateFile(c)

		if err != nil {
			c.AbortWithError(err.Code, err.Error)
			return
		}

		c.JSON(http.StatusOK, res)
	})

	r.GET("/:fileID/:fileName", func(c *gin.Context) {

		fileService.GetFileStream(c)
	})
	r.GET("/songMetadata/:fileID", func(c *gin.Context) {

		res, err := fileService.GetSongMainMetadata(c)
		if err != nil {
			c.AbortWithError(err.Code, err.Error)
			return
		}

		c.JSON(http.StatusOK, res)
	})
	r.GET("/songCover/:fileID/:coverName", func(c *gin.Context) {

		fileService.GetSongCoverStream(c)
	})

	r.POST("/movefiles", func(c *gin.Context) {

		res, err := fileService.MoveFiles(c)

		if err != nil {
			c.AbortWithError(err.Code, err.Error)
			return
		}

		c.JSON(http.StatusOK, res)
	})

	r.POST("/makedir", func(c *gin.Context) {

		res, err := fileService.MakeDirectory(c)

		if err != nil {
			c.AbortWithError(err.Code, err.Error)
			return
		}

		c.JSON(http.StatusOK, res)
	})

	r.POST("/deletefiles", func(c *gin.Context) {

		res, err := fileService.DeleteFiles(c)

		if err != nil {
			c.AbortWithError(err.Code, err.Error)
			return
		}

		c.JSON(http.StatusOK, res)
	})

}
