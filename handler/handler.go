package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ArchiveHandler struct {
	item archiveItem
}

func NewArchiveHandler(item archiveItem) *ArchiveHandler {
	return &ArchiveHandler{item: item}
}

func (hdl *ArchiveHandler) GetArchiveList(c *gin.Context) {
	offset, ok := c.Get("offset")
	if !ok {
		hdl.errorPage(c)
	}

	limit, ok := c.Get("limit")
	if !ok {
		hdl.errorPage(c)
	}

	oft, ok := offset.(int32)
	if !ok {
		hdl.errorPage(c)
	}

	lmt, ok := limit.(int32)
	if !ok {
		hdl.errorPage(c)
	}

	listRes, err := hdl.item.GetArchiveList(oft, lmt)
	if err != nil {
		hdl.errorPage(c)
	}

	count, err := hdl.item.GetArchiveItemCount()
	if err != nil {
		hdl.errorPage(c)
	}

	c.HTML(http.StatusOK, "index.html", indexPageData{Items: listRes.ArchiveItems, Count: count})
}

func (hdl *ArchiveHandler) SearchItem(c *gin.Context) {
	c.HTML(http.StatusOK, "search.html", nil)
}

func (hdl *ArchiveHandler) errorPage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}
