package files

import (
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
)

type HandlerAPI interface {
	UploadFile(c *gin.Context)
	DownloadFile(c *gin.Context)
	ListFiles(c *gin.Context)
	DeleteFile(c *gin.Context)
}

type Handler struct {
	processor ProcessorAPI
}

func NewHandler(processor ProcessorAPI) *Handler {
	return &Handler{processor: processor}
}

func (h *Handler) UploadFile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "No file provided",
		})
		return
	}

	fileReader, err := file.Open()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to open file",
		})
		return
	}
	defer fileReader.Close()

	fileBytes, err := io.ReadAll(fileReader)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to read file",
		})
		return
	}

	err = h.processor.UploadFile(fileBytes, file.Filename)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to process file",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "File uploaded successfully",
	})
}

func (h *Handler) DownloadFile(c *gin.Context) {
	id := c.Param("id")

	file, name, err := h.processor.GetFileContentAndName(id)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to download file",
		})
		return
	}

	c.Header("Content-Disposition", "attachment; filename="+name)
	c.Data(http.StatusOK, "application/octet-stream", file)
}

func (h *Handler) ListFiles(c *gin.Context) {
	files, err := h.processor.ListFiles()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to list files",
		})
		return
	}

	c.JSON(http.StatusOK, files)
}

func (h *Handler) DeleteFile(c *gin.Context) {
	id := c.Param("id")

	err := h.processor.DeleteFile(id)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to delete file",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "File deleted successfully",
	})
}
