package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"html/template"
	"net/http"
)

func (h *Handler) Get(ctx *gin.Context) {
	searchOrder := ctx.Request.FormValue("order_uid")
	model, err := h.services.GetModelCache(searchOrder)
	if err != nil {
		logrus.Fatalln(err)
		newErrorRespone(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	page, err := template.ParseFiles("message.html")
	if err != nil {
		newErrorRespone(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	err = page.Execute(ctx.Writer, model)
	if err != nil {
		newErrorRespone(ctx, http.StatusInternalServerError, err.Error())
		return
	}
}

func (h *Handler) CreateHTML(ctx *gin.Context) {
	page, err := template.ParseFiles("message.html")
	if err != nil {
		newErrorRespone(ctx, http.StatusInternalServerError, err.Error())
	}
	err = page.Execute(ctx.Writer, nil)
	if err != nil {
		newErrorRespone(ctx, http.StatusInternalServerError, err.Error())
	}
}
