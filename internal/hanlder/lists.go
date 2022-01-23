package hanlder

import (
	"github.com/gin-gonic/gin"
	"http_server/model"
	"net/http"
	"strconv"
)

func (h *Handler) createList(c *gin.Context) {
	var lists model.Lists
	if err := c.BindJSON(&lists); err != nil {
		newErrResp(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.service.Lists.CreateList(lists.Title, lists.Description)
	if err != nil {
		newErrResp(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type getAllLists struct {
	Data []model.Lists `json:"data"`
}

func (h *Handler) getLists(c *gin.Context) {
	lists, err := h.service.Lists.GetLists()
	if err != nil {
		newErrResp(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, &getAllLists{Data: lists})
}

func (h *Handler) getListById(c *gin.Context) {
	getId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrResp(c, http.StatusBadRequest, err.Error())
		return
	}
	lists, err := h.service.Lists.GetListById(getId)
	if err != nil {
		newErrResp(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, model.Lists{
		Id:          lists.Id,
		Title:       lists.Title,
		Description: lists.Description,
	})
}

func (h *Handler) updateListById(c *gin.Context) {
	getId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrResp(c, http.StatusBadRequest, err.Error())
		return
	}
	var lists model.Lists
	if err := c.BindJSON(&lists); err != nil {
		newErrResp(c, http.StatusBadRequest, err.Error())
		return
	}
	list, err := h.service.Lists.UpdateListById(getId, lists.Title, lists.Description)
	if err != nil {
		newErrResp(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id":      list,
		"message": "success",
	})
}

func (h *Handler) deleteListById(c *gin.Context) {
	getId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrResp(c, http.StatusBadRequest, err.Error())
		return
	}
	if err := h.service.Lists.DeleteListById(getId); err != nil {
		newErrResp(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
	})
}
