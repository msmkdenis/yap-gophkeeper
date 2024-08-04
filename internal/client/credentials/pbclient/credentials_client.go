package pbclient

import (
	pb "github.com/msmkdenis/yap-gophkeeper/internal/proto/credentials"
)

type CredentialsPBClient struct {
	credentialsService pb.CredentialsServiceClient
}

func NewCredentialsPBClient(u pb.CredentialsServiceClient) *CredentialsPBClient {
	return &CredentialsPBClient{
		credentialsService: u,
	}
}
