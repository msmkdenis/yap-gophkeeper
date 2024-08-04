package pbclient

import (
	"context"

	"google.golang.org/grpc/metadata"

	"github.com/msmkdenis/yap-gophkeeper/internal/client/model"
	pb "github.com/msmkdenis/yap-gophkeeper/internal/proto/text_data"
)

func (u *TextDataPBClient) SaveTextData(token string, text model.TextDataPostRequest) (model.TextData, error) {
	req := &pb.PostTextDataRequest{
		Text:     text.Text,
		Metadata: text.MetaData,
	}

	md := metadata.New(map[string]string{"token": token})
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	resp, err := u.textDataService.PostSaveTextData(ctx, req)
	if err != nil {
		return model.TextData{}, err
	}

	txt := model.TextData{
		Text:     resp.Text,
		MetaData: resp.Metadata,
	}

	return txt, nil
}
