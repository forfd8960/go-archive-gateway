package archive_client

import (
	"context"
	"time"

	"google.golang.org/grpc"

	pb "github.com/forfd8960/go-archive/pb"
)

type archiveClient struct {
	ctx     context.Context
	addrs   string
	timeout time.Duration
}

func NewArchiveClient(ctx context.Context, addrs string, timeout time.Duration) *archiveClient {
	return &archiveClient{ctx: ctx, addrs: addrs, timeout: timeout}
}

func (c *archiveClient) GetArchiveList(offset, limit int32) (*pb.GetArchiveListRes, error) {
	var cc *grpc.ClientConn
	cc, err := grpc.DialContext(c.ctx, c.addrs, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	defer cc.Close()

	timeOutCtx, fn := context.WithTimeout(c.ctx, c.timeout)
	defer fn()

	res, err := pb.NewGoArchiveClient(cc).GetArchiveList(timeOutCtx, &pb.GetArchiveListReq{Offset: offset, Limit: limit})
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c *archiveClient) GetArchiveItemCount() (int64, error) {
	var cc *grpc.ClientConn
	cc, err := grpc.DialContext(c.ctx, c.addrs, grpc.WithInsecure())
	if err != nil {
		return -1, err
	}
	defer cc.Close()

	timeOutCtx, fn := context.WithTimeout(c.ctx, c.timeout)
	defer fn()

	res, err := pb.NewGoArchiveClient(cc).GetArchiveItemCount(timeOutCtx, &pb.GetArchiveItemCountReq{})
	if err != nil {
		return -1, err
	}

	return res.GetCount(), nil
}
