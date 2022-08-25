package handlers

import (
	"advert"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)



func (h *Handler) create(c *gin.Context) {
	var input advert.Advert

	if err := c.BindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}
	maxLenOfTitle := 200
	maxLenOfDescription := 1000

	if len(input.Title) > maxLenOfTitle || len(input.Description) > maxLenOfDescription{
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{"message": "Check title and desc's len"})
		return
	}
	input.CreateDate = time.Now()
	id, err := h.services.AdvertList.Create(input)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{"id": id})
}

func (h *Handler) getById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	fields := c.Request.URL.Query().Get("fields")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}
	advert, err := h.services.GetById(id, fields)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, advert)
}

func (h *Handler) getAll(c *gin.Context) {
	typeSort := c.Request.URL.Query().Get("typesort")
	subjectSort := c.Request.URL.Query().Get("subjectsort")
	page, err := strconv.Atoi(c.Request.URL.Query().Get("page"))
	if err != nil {
		page = 1
	}

	if typeSort != "asc" && typeSort != "desc" {
		typeSort = "asc"
	}

	if subjectSort != "create_date" && subjectSort != "price" {
		typeSort = "create_date"
	}
	adverts, err := h.services.AdvertList.GetAll(page, typeSort, subjectSort)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, adverts)
}
