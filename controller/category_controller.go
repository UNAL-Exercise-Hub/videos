package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jbautistas/videos_ms/orm"
)

func GetMusculos(ctx *gin.Context) {
	db := orm.GetDB()
	var musculos []orm.Musculos
	result := db.Find(&musculos)
	if result.Error == nil {
		if result.RowsAffected == 0 {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "Musculos no encontrados",
			})
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "Musculos encontrados",
				"videos":  musculos,
			})
		}
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Error in db",
		})
	}
}

func GetGrupoMuscular(ctx *gin.Context) {
	db := orm.GetDB()
	var grupos []orm.Grupo
	result := db.Find(&grupos)
	if result.Error == nil {
		if result.RowsAffected == 0 {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "Grupos no encontrados",
			})
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "Grupos musculares encontrados",
				"videos":  grupos,
			})
		}
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Error in db",
		})
	}
}

func GetObjetivos(ctx *gin.Context) {
	db := orm.GetDB()
	var objetivos []orm.Objetivo
	result := db.Find(&objetivos)
	if result.Error == nil {
		if result.RowsAffected == 0 {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "Objetivos no encontrados",
			})
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "Objetivos encontrados",
				"videos":  objetivos,
			})
		}
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Error in bd",
		})
	}
}

func GetDificultad(ctx *gin.Context) {
	db := orm.GetDB()
	var dificultades []orm.Dificultad
	result := db.Find(&dificultades)
	if result.Error == nil {
		if result.RowsAffected == 0 {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "Dificultades no encontradas",
			})
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "Dificultades encontradas",
				"videos":  dificultades,
			})
		}
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Error in db",
		})
	}
}

func GetEquipamento(ctx *gin.Context) {
	db := orm.GetDB()
	var equipamentos []orm.Equipamento
	result := db.Find(&equipamentos)
	if result.Error == nil {
		if result.RowsAffected == 0 {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "Equipamentos no encontrados",
			})
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "Equipamentos encontrados",
				"videos":  equipamentos,
			})
		}
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Error in db",
		})
	}
}

func GetDisciplina(ctx *gin.Context) {
	db := orm.GetDB()
	var disciplina []orm.Disciplina
	result := db.Find(&disciplina)
	if result.Error == nil {
		if result.RowsAffected == 0 {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "Disciplinas no encontradas",
			})
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "Disciplinas encontradas",
				"videos":  disciplina,
			})
		}
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Error in db",
		})
	}
}
