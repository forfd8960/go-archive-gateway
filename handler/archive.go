package handler

import (
	pb "github.com/forfd8960/go-archive/pb"
)

type archiveItem interface {
	GetArchiveList(int32, int32) (*pb.GetArchiveListRes, error)
	GetArchiveItemCount() (int64, error)
}

type indexPageData struct {
	Items        []*pb.ArchiveItem
	Count        int64
	PreviousPage int64
	CurrentPage  int64
	NextPage     int64
}
