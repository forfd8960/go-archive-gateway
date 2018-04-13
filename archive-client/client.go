package archive_client

import (
	"context"
	"time"

	pb "github.com/forfd8960/go-archive/pb"
)

type archiveClient struct {
	ctx     context.Context
	timeout time.Duration
}

func NewArchiveClient(ctx context.Context, timeout time.Duration) *archiveClient {
	return &archiveClient{ctx: ctx, timeout: timeout}
}

func (c *archiveClient) GetArchiveList(int32, int32) (*pb.GetArchiveListRes, error) {
	return new(pb.GetArchiveListRes), nil
}

func (c *archiveClient) GetArchiveItemCount() (int64, error) {
	return 0, nil
}
