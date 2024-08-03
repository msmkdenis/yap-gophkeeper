package grpchandlers

import (
	"context"
	"log"
	"net"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"

	pb "github.com/msmkdenis/yap-gophkeeper/internal/proto/text_data"
	"github.com/msmkdenis/yap-gophkeeper/internal/server/config"
	"github.com/msmkdenis/yap-gophkeeper/internal/server/interceptors/auth"
	mocks "github.com/msmkdenis/yap-gophkeeper/internal/server/mocks/text_data"
	textDataValidation "github.com/msmkdenis/yap-gophkeeper/internal/server/text_data/api/v1/validation"
	"github.com/msmkdenis/yap-gophkeeper/pkg/jwtmanager"
)

var cfgMock = &config.Config{
	TokenName:     "token",
	TokenSecret:   "secret",
	TokenExpHours: 24,
}

type TextDataHandlerTestSuite struct {
	suite.Suite
	textDataService *mocks.MockTextDataService
	dialer          func(ctx context.Context, address string) (net.Conn, error)
	jwtManager      *jwtmanager.JWTManager
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(TextDataHandlerTestSuite))
}

func (c *TextDataHandlerTestSuite) SetupSuite() {
	ctrl := gomock.NewController(c.T())
	c.textDataService = mocks.NewMockTextDataService(ctrl)
	c.jwtManager = jwtmanager.New(cfgMock.TokenName, cfgMock.TokenSecret, cfgMock.TokenExpHours)
	authentication := auth.New(c.jwtManager).GRPCJWTAuth
	validate := validator.New()
	textDataValidator, err := textDataValidation.New(validate)
	require.NoError(c.T(), err)

	buffer := 1024 * 1024
	lis := bufconn.Listen(buffer)
	server := grpc.NewServer(grpc.ChainUnaryInterceptor(authentication))
	pb.RegisterTextDataServiceServer(server, New(c.textDataService, textDataValidator))

	c.dialer = func(ctx context.Context, address string) (net.Conn, error) {
		return lis.Dial()
	}

	go func() {
		if err := server.Serve(lis); err != nil {
			log.Fatal(err)
		}
	}()
}
