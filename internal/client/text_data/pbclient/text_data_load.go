package pbclient

import (
	"context"
	"fmt"

	"google.golang.org/grpc/metadata"

	"github.com/msmkdenis/yap-gophkeeper/internal/client/model"
	pb "github.com/msmkdenis/yap-gophkeeper/internal/proto/text_data"
)

func (u *TextDataPBClient) LoadTextData(token string, textData model.TextDataLoadRequest) ([]model.TextData, error) {
	req := &pb.GetTextDataRequest{
		Text:     textData.Text,
		Metadata: textData.MetaData,
	}

	md := metadata.New(map[string]string{"token": token})
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	resp, err := u.textDataService.GetLoadTextData(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("load text data: %w", err)
	}

	texts := make([]model.TextData, 0, len(resp.Text))
	for _, txt := range resp.Text {
		texts = append(texts, model.TextData{
			Text:     txt.Text,
			MetaData: txt.Metadata,
		})
	}

	return texts, nil
}
