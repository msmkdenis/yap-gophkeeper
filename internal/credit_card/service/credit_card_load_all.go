package service

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/msmkdenis/yap-gophkeeper/internal/credit_card/specification"
	"github.com/msmkdenis/yap-gophkeeper/internal/interceptors/auth"
	"github.com/msmkdenis/yap-gophkeeper/internal/interceptors/keyextraction"
	"github.com/msmkdenis/yap-gophkeeper/internal/model"
)

func (s *CreditCardService) LoadAllCreditCard(ctx context.Context, spec specification.CreditCardSpecification) ([]model.CreditCardPostResponse, error) {
	userID, ok := ctx.Value(auth.UserIDContextKey("userID")).(string)
	if !ok {
		return nil, fmt.Errorf("failed to get userID from context")
	}

	userKey, ok := ctx.Value(keyextraction.UserKeyContextKey("userKey")).([]byte)
	if !ok {
		return nil, fmt.Errorf("failed to get userKey from context")
	}

	encryptedCards, err := s.repository.SelectAll(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("select all cards: %w", err)
	}

	cards := make([]model.CreditCardPostResponse, 0, len(encryptedCards))
	for _, encryptedCard := range encryptedCards {
		decryptedData, err := s.crypt.Decrypt(userKey, encryptedCard.CryptData)
		if err != nil {
			return nil, fmt.Errorf("decrypt data: %w", err)
		}

		var card model.CreditCardCryptData
		err = json.Unmarshal(decryptedData, &card)
		if err != nil {
			return nil, fmt.Errorf("unmarshal data: %w", err)
		}

		cards = append(cards, model.CreditCardPostResponse{
			ID:        encryptedCard.ID,
			OwnerID:   encryptedCard.OwnerID,
			Number:    card.Number,
			OwnerName: card.OwnerName,
			ExpiresAt: card.ExpiresAt,
			CVV:       card.CVV,
			PinCode:   card.PinCode,
			MetaData:  encryptedCard.MetaData,
			CreatedAt: encryptedCard.CreatedAt,
			UpdatedAt: encryptedCard.UpdatedAt,
		})
	}

	predicates := spec.MakeFilterPredicates()
	var filteredCards []model.CreditCardPostResponse
	for _, card := range cards {
		take := true
		for _, filterCardWithSpec := range predicates {
			if !filterCardWithSpec(spec, card) {
				take = false
				break
			}
		}
		if take {
			filteredCards = append(filteredCards, card)
		}
	}

	return filteredCards, nil
}