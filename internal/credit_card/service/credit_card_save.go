package service

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/google/uuid"

	"github.com/msmkdenis/yap-gophkeeper/internal/interceptors/auth"
	"github.com/msmkdenis/yap-gophkeeper/internal/interceptors/keyextraction"
	"github.com/msmkdenis/yap-gophkeeper/internal/model"
)

func (s *CreditCardService) SaveCreditCard(ctx context.Context, req model.CreditCardPostRequest) (model.CreditCardPostResponse, error) {
	userID, ok := ctx.Value(auth.UserIDContextKey("userID")).(string)
	if !ok {
		return model.CreditCardPostResponse{}, fmt.Errorf("failed to get userID from context")
	}

	userKey, ok := ctx.Value(keyextraction.UserKeyContextKey("userKey")).([]byte)
	if !ok {
		return model.CreditCardPostResponse{}, fmt.Errorf("failed to get userKey from context")
	}

	id, err := uuid.NewUUID()
	if err != nil {
		return model.CreditCardPostResponse{}, fmt.Errorf("new uuid: %w", err)
	}

	card := model.CreditCardCryptData{
		Number:    req.Number,
		OwnerName: req.OwnerName,
		ExpiresAt: req.ExpiresAt,
		CVV:       req.CVV,
		PinCode:   req.PinCode,
	}

	data, err := json.Marshal(card)
	if err != nil {
		return model.CreditCardPostResponse{}, fmt.Errorf("marshal: %w", err)
	}

	cryptData, err := s.crypt.Encrypt(userKey, data)
	if err != nil {
		return model.CreditCardPostResponse{}, fmt.Errorf("encrypt data: %w", err)
	}

	creditCard := model.CreditCard{
		ID:        id.String(),
		OwnerID:   userID,
		CryptData: cryptData,
		MetaData:  req.MetaData,
	}

	savedCreditCard, err := s.repository.Insert(ctx, creditCard)
	if err != nil {
		return model.CreditCardPostResponse{}, fmt.Errorf("insert credit card: %w", err)
	}

	return model.CreditCardPostResponse{
		ID:        savedCreditCard.ID,
		OwnerID:   savedCreditCard.OwnerID,
		Number:    req.Number,
		OwnerName: req.OwnerName,
		ExpiresAt: req.ExpiresAt,
		CVV:       req.CVV,
		PinCode:   req.PinCode,
		MetaData:  savedCreditCard.MetaData,
		CreatedAt: savedCreditCard.CreatedAt,
		UpdatedAt: savedCreditCard.UpdatedAt,
	}, nil
}
