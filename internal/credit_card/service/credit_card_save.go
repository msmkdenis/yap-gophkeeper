package service

import (
	"context"
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

	cryptNumber, err := s.crypt.Encrypt(userKey, []byte(req.Number))
	if err != nil {
		return model.CreditCardPostResponse{}, fmt.Errorf("encrypt number: %w", err)
	}

	cryptOwnerName, err := s.crypt.Encrypt(userKey, []byte(req.OwnerName))
	if err != nil {
		return model.CreditCardPostResponse{}, fmt.Errorf("encrypt owner_name: %w", err)
	}

	cryptExpiresAt, err := s.crypt.Encrypt(userKey, []byte(req.ExpiresAt))
	if err != nil {
		return model.CreditCardPostResponse{}, fmt.Errorf("encrypt expires_at: %w", err)
	}

	cryptCVV, err := s.crypt.Encrypt(userKey, []byte(req.CVV))
	if err != nil {
		return model.CreditCardPostResponse{}, fmt.Errorf("encrypt cvv: %w", err)
	}

	cryptPIN, err := s.crypt.Encrypt(userKey, []byte(req.PinCode))
	if err != nil {
		return model.CreditCardPostResponse{}, fmt.Errorf("encrypt pin: %w", err)
	}

	creditCard := model.CreditCard{
		ID:        id.String(),
		OwnerID:   userID,
		Number:    cryptNumber,
		OwnerName: cryptOwnerName,
		ExpiresAt: cryptExpiresAt,
		CVV:       cryptCVV,
		PinCode:   cryptPIN,
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
