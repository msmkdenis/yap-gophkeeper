package pbclient

import (
	"context"
	"fmt"

	"google.golang.org/grpc/metadata"

	"github.com/msmkdenis/yap-gophkeeper/internal/client/model"
	pb "github.com/msmkdenis/yap-gophkeeper/internal/proto/binary_data"
)

func (u *BinaryDataPBClient) SaveBinaryData(token string, bData model.BinaryDataPostRequest) (model.BinaryData, error) {
	req := &pb.PostBinaryDataRequest{
		Data:      bData.Data,
		Name:      bData.Name,
		Extension: bData.Extension,
		Metadata:  bData.MetaData,
	}

	md := metadata.New(map[string]string{"token": token})
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	resp, err := u.binaryDataService.PostSaveBinaryData(ctx, req)
	if err != nil {
		return model.BinaryData{}, fmt.Errorf("save binary data: %w", err)
	}

	binaryData := model.BinaryData{
		Name:      resp.Name,
		Extension: resp.Extension,
		MetaData:  resp.Metadata,
	}

	return binaryData, nil
}
