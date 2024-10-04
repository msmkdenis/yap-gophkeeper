package pbclient

import (
	pb "github.com/msmkdenis/yap-gophkeeper/internal/proto/binary_data"
)

type BinaryDataPBClient struct {
	binaryDataService pb.BinaryDataServiceClient
}

func NewBinaryDataPBClient(u pb.BinaryDataServiceClient) *BinaryDataPBClient {
	return &BinaryDataPBClient{
		binaryDataService: u,
	}
}
