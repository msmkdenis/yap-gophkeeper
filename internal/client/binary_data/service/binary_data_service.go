package service

import (
	"github.com/msmkdenis/yap-gophkeeper/internal/client/model"
	"github.com/msmkdenis/yap-gophkeeper/internal/client/state"
)

type BinaryDataService interface {
	SaveBinaryData(token string, bData model.BinaryDataPostRequest) (model.BinaryData, error)
	LoadBinaryData(token string, bData model.BinaryDataLoadRequest) ([]model.BinaryData, error)
}

type BinaryDataProvider struct {
	binaryDataService BinaryDataService
	state             *state.ClientState
}

func NewBinaryDataService(u BinaryDataService, state *state.ClientState) *BinaryDataProvider {
	return &BinaryDataProvider{
		binaryDataService: u,
		state:             state,
	}
}
