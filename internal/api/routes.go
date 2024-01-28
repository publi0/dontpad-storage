package api

import (
	"dontpad-storage/internal/files"
	"github.com/gin-gonic/gin"
)

type RoutesAPI interface {
	SetupRoutes()
}

type Routes struct {
	engine      *gin.Engine
	fileHandler files.HandlerAPI
}

func NewRoutes(engine *gin.Engine, fileHandler files.HandlerAPI) *Routes {
	return &Routes{
		engine:      engine,
		fileHandler: fileHandler,
	}
}

func (r *Routes) SetupRoutes() {

	r.engine.PUT("/files", r.fileHandler.UploadFile)

	r.engine.GET("/files", r.fileHandler.ListFiles)

	r.engine.GET("/files/:id", r.fileHandler.DownloadFile)

	r.engine.DELETE("/files/:id", r.fileHandler.DeleteFile)
}
