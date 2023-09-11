package controller

import (
	"fmt"
	"net/http"

	"github.com/UNWorkout/Videos/orm"
	"github.com/gin-gonic/gin"
)

// Get all videos
func GetVideos(ctx *gin.Context) {
	db := orm.GetDB()
	var videos []orm.Video
	result := db.Model(&orm.Video{}).Preload("Musculos").Preload("Dificultad").Preload("Grupo").Preload("Objetivo").Preload("Equipamento").Preload("Disciplina").Find(&videos)
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
	var options orm.CreateParamsVideo
	var newVideo orm.Video
	var exist orm.Video

	if err := ctx.BindJSON(&options); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Error in db",
		})
		return
	}
	db := orm.GetDB()

	newVideo.Titulo = options.Titulo
	newVideo.Link = options.Link
	newVideo.Imagen = options.Imagen
	newVideo.Duracion = options.Duracion
	newVideo.Series = options.Series

	result := db.Where("titulo = ?", options.Titulo).Find(&exist)
	if result.RowsAffected != 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Titulo de video ya existente",
		})
		return
	}
	if len(options.Musculos) != 0 {
		var musculos []*orm.Musculos
		for _, element := range options.Musculos {
			var musculo orm.Musculos
			db.First(&musculo, element)
			musculos = append(musculos, &musculo)
		}
		newVideo.Musculos = musculos
	}
	if len(options.Grupo) != 0 {
		var grupos []*orm.Grupo
		fmt.Println(options.Grupo)
		for _, element := range options.Grupo {
			var grupo orm.Grupo
			db.First(&grupo, element)
			grupos = append(grupos, &grupo)
		}
		newVideo.Grupo = grupos
	}
	if len(options.Objetivo) != 0 {
		var objetivos []*orm.Objetivo
		for _, element := range options.Objetivo {
			var objetivo orm.Objetivo
			db.First(&objetivo, element)
			objetivos = append(objetivos, &objetivo)
		}
		newVideo.Objetivo = objetivos
	}
	if len(options.Dificultad) != 0 {
		var dificultades []*orm.Dificultad
		for _, element := range options.Dificultad {
			var dificultad orm.Dificultad
			db.First(&dificultad, element)
			dificultades = append(dificultades, &dificultad)
		}
		newVideo.Dificultad = dificultades
	}
	if len(options.Equipamento) != 0 {
		var equipamentos []*orm.Equipamento
		fmt.Println(options.Equipamento)
		for _, element := range options.Equipamento {
			var equipamento orm.Equipamento
			db.First(&equipamento, element)
			equipamentos = append(equipamentos, &equipamento)
		}
		newVideo.Equipamento = equipamentos
	}
	if len(options.Disciplina) != 0 {
		var disciplinas []*orm.Disciplina
		fmt.Println(options.Disciplina)
		for _, element := range options.Disciplina {
			var disciplina orm.Disciplina
			db.First(&disciplina, element)
			disciplinas = append(disciplinas, &disciplina)
		}
		newVideo.Disciplina = disciplinas
	}
	db.Create(&newVideo)
	ctx.IndentedJSON(http.StatusCreated, newVideo)
}

// Get all videos
func GetVideoID(ctx *gin.Context) {
	id := ctx.Param("id")
	db := orm.GetDB()
	var video orm.Video

	result := db.Model(&orm.Video{}).Preload("Musculos").Preload("Dificultad").Preload("Grupo").Preload("Objetivo").Preload("Equipamento").Preload("Disciplina").First(&video, id)
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
		var inf orm.CreateParamsVideo
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
		if len(inf.Musculos) != 0 {
			var musculos []*orm.Musculos
			for _, element := range inf.Musculos {
				var musculo orm.Musculos
				db.First(&musculo, element)
				musculos = append(musculos, &musculo)
			}
			video.Musculos = musculos
		}
		if len(inf.Grupo) != 0 {
			var grupos []*orm.Grupo
			fmt.Println(inf.Grupo)
			for _, element := range inf.Grupo {
				var grupo orm.Grupo
				db.First(&grupo, element)
				grupos = append(grupos, &grupo)
			}
			video.Grupo = grupos
		}
		if len(inf.Objetivo) != 0 {
			var objetivos []*orm.Objetivo
			for _, element := range inf.Objetivo {
				var objetivo orm.Objetivo
				db.First(&objetivo, element)
				objetivos = append(objetivos, &objetivo)
			}
			video.Objetivo = objetivos
		}
		if len(inf.Dificultad) != 0 {
			var dificultades []*orm.Dificultad
			for _, element := range inf.Dificultad {
				var dificultad orm.Dificultad
				db.First(&dificultad, element)
				dificultades = append(dificultades, &dificultad)
			}
			video.Dificultad = dificultades
		}
		if len(inf.Equipamento) != 0 {
			var equipamentos []*orm.Equipamento
			fmt.Println(inf.Equipamento)
			for _, element := range inf.Equipamento {
				var equipamento orm.Equipamento
				db.First(&equipamento, element)
				equipamentos = append(equipamentos, &equipamento)
			}
			video.Equipamento = equipamentos
		}
		if len(inf.Disciplina) != 0 {
			var disciplinas []*orm.Disciplina
			fmt.Println(inf.Disciplina)
			for _, element := range inf.Disciplina {
				var disciplina orm.Disciplina
				db.First(&disciplina, element)
				disciplinas = append(disciplinas, &disciplina)
			}
			video.Disciplina = disciplinas
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
