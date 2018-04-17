package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	pb "github.com/forfd8960/go-archive/pb"
)

const defaultLimit = 10

type ArchiveHandler struct {
	item archiveItem
}

func NewArchiveHandler(item archiveItem) *ArchiveHandler {
	return &ArchiveHandler{item: item}
}

func (hdl *ArchiveHandler) GetArchiveList(c *gin.Context) {
	page := c.Query("page")

	var err error
	var oft = 0
	var lmt = defaultLimit
	if page == "" {
		listRes, count, err := hdl.getIndexData(int32(oft), int32(lmt))
		if err != nil {
			hdl.IndexPage(c)
			return
		}

		c.HTML(http.StatusOK,
			"index.html",
			indexPageData{
				Items:        listRes.ArchiveItems,
				Count:        count,
				PreviousPage: 0,
				CurrentPage:  1,
				NextPage:     2,
			})
		return
	}

	var pageNum = 1
	pageNum, err = strconv.Atoi(page)
	if err != nil {
		hdl.IndexPage(c)
		return
	}

	listRes, count, err := hdl.getIndexData(int32(pageNum*defaultLimit), int32(defaultLimit))
	if err != nil {
		hdl.IndexPage(c)
		return
	}

	c.HTML(http.StatusOK,
		"index.html",
		indexPageData{
			Items:        listRes.ArchiveItems,
			Count:        count,
			PreviousPage: int64(pageNum - 1),
			CurrentPage:  int64(pageNum),
			NextPage:     int64(pageNum + 1),
		},
	)
}

func (hdl *ArchiveHandler) SearchItem(c *gin.Context) {
	c.HTML(http.StatusOK, "search.html", nil)
}

func (hdl *ArchiveHandler) IndexPage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func (hdl *ArchiveHandler) getIndexData(oft, lmt int32) (*pb.GetArchiveListRes, int64, error) {
	listRes, err := hdl.item.GetArchiveList(int32(oft), int32(lmt))
	if err != nil {
		return nil, -1, err
	}

	count, err := hdl.item.GetArchiveItemCount()
	if err != nil {
		return nil, -1, err
	}

	return listRes, count, nil
}
