package pbclient

import (
	pb "github.com/msmkdenis/yap-gophkeeper/internal/proto/text_data"
)

type TextDataPBClient struct {
	textDataService pb.TextDataServiceClient
}

func NewCreditCardPBClient(u pb.TextDataServiceClient) *TextDataPBClient {
	return &TextDataPBClient{
		textDataService: u,
	}
}
