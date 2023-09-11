package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jbautistas/videos_ms/orm"
)

// Get all videos
func GetVideos(ctx *gin.Context) {
	db := orm.GetDB()
	var videos []orm.Video
	result := db.Find(&videos)
	if result.Error == nil {
		if result.RowsAffected == 0 {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "No videos found",
			})
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "Videos found!!",
				"videos":  videos,
			})
		}
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Error in db",
		})
	}
}

// Post video
func PostVideo(ctx *gin.Context) {
	var newVideo orm.Video
	if err := ctx.BindJSON(&newVideo); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Error in db",
		})
	}
	db := orm.GetDB()
	db.Create(&newVideo)
	ctx.IndentedJSON(http.StatusCreated, newVideo)
}

// Get all videos
func GetVideoID(ctx *gin.Context) {
	id := ctx.Param("id")
	db := orm.GetDB()
	var video orm.Video

	result := db.First(&video, id)
	if result.Error == nil {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Video encontrado",
			"videos":  video,
		})
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Video no encontrado con id:" + id,
		})
	}
}

func UpdateVideo(ctx *gin.Context) {
	id := ctx.Param("id")
	db := orm.GetDB()
	var video orm.Video

	result := db.First(&video, id)
	if result.Error == nil {
		var inf orm.Video
		if err := ctx.BindJSON(&inf); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "Error in db",
			})
		}
		if inf.Titulo != "" {
			video.Titulo = inf.Titulo
		}
		if inf.Link != "" {
			video.Link = inf.Link
		}
		if inf.Imagen != "" {
			video.Imagen = inf.Imagen
		}
		if inf.Duracion != 0 {
			video.Duracion = inf.Duracion
		}
		if inf.Series != 0 {
			video.Series = inf.Series
		}
		db.Save(&video)

		ctx.JSON(http.StatusOK, gin.H{
			"message": "Video actualizado",
			"videos":  video,
		})
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Video no encontrado con id:" + id,
		})
	}
}

func DeleteVideo(ctx *gin.Context) {
	id := ctx.Param("id")
	db := orm.GetDB()
	var video orm.Video

	result := db.First(&video, id)
	if result.Error == nil {
		db.Delete(&video)
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Video eliminado correctamente",
			"videos":  video,
		})
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Video no encontrado con id:" + id,
		})
	}
}
