package specification

import (
	"strings"

	"github.com/msmkdenis/yap-gophkeeper/internal/model"
	pb "github.com/msmkdenis/yap-gophkeeper/internal/proto/text_data"
)

type TextDataSpecification struct {
	Text     string
	MetaData string
}

func NewTextDataSpecification(in *pb.GetTextDataRequest) (TextDataSpecification, error) {
	spec := TextDataSpecification{
		Text:     in.Text,
		MetaData: in.Metadata,
	}

	return spec, nil
}

type Predicate func(spec TextDataSpecification, resp model.TextData) bool

func (t *TextDataSpecification) MakeFilterPredicates() []Predicate {
	var predicates []Predicate
	if t.MetaData != "" {
		predicates = append(predicates, makeMetaDataPredicate())
	}
	if t.Text != "" {
		predicates = append(predicates, makeTextPredicate())
	}
	return predicates
}

func makeMetaDataPredicate() Predicate {
	return func(spec TextDataSpecification, resp model.TextData) bool {
		return strings.Contains(strings.ToLower(resp.MetaData), strings.ToLower(spec.MetaData))
	}
}

func makeTextPredicate() Predicate {
	return func(spec TextDataSpecification, resp model.TextData) bool {
		return strings.Contains(strings.ToLower(resp.Text), strings.ToLower(spec.Text))
	}
}
