package grpchandlers

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	pb "github.com/msmkdenis/yap-gophkeeper/internal/proto/text_data"
	"github.com/msmkdenis/yap-gophkeeper/internal/server/model"
)

func (c *TextDataHandlerTestSuite) Test_GetLoadTextData() {
	token, err := c.jwtManager.BuildJWTString("050a289a-d10a-417b-ab89-3acfca0f6529")
	require.NoError(c.T(), err)

	texts := []model.TextData{
		{
			ID:        "some id",
			OwnerID:   "some owner id",
			Text:      "some text",
			MetaData:  "some metadata",
			CreatedAt: time.Date(2020, 1, 1, 12, 0, 0, 0, time.UTC),
			UpdatedAt: time.Date(2020, 1, 1, 12, 0, 0, 0, time.UTC),
		},
		{
			ID:        "another id",
			OwnerID:   "another owner id",
			Text:      "another text",
			MetaData:  "another metadata",
			CreatedAt: time.Date(2024, 1, 1, 12, 0, 0, 0, time.UTC),
			UpdatedAt: time.Date(2024, 1, 1, 12, 0, 0, 0, time.UTC),
		},
	}

	textData := make([]*pb.TextData, 0, len(texts))
	for _, v := range texts {
		textData = append(textData, &pb.TextData{
			Id:        v.ID,
			OwnerId:   v.OwnerID,
			Text:      v.Text,
			Metadata:  v.MetaData,
			CreatedAt: v.CreatedAt.Format(time.RFC3339),
			UpdatedAt: v.UpdatedAt.Format(time.RFC3339),
		})
	}

	testCases := []struct {
		name                         string
		token                        string
		body                         *pb.GetTextDataRequest
		expectedCode                 codes.Code
		expectedStatusMessage        string
		expectedViolationField       string
		expectedViolationDescription string
		prepare                      func()
		expectedBody                 *pb.GetTextDataResponse
	}{
		{
			name:  "Unauthorized - token not found",
			token: "",
			body: &pb.GetTextDataRequest{
				Text:     "",
				Metadata: "",
			},
			expectedCode:          codes.Unauthenticated,
			expectedStatusMessage: "authentification by UserID failed",
			prepare: func() {
				c.textDataService.EXPECT().LoadAllTextData(gomock.Any(), gomock.Any()).Times(0)
			},
		},
		{
			name:  "Internal error - unable to load text data",
			token: token,
			body: &pb.GetTextDataRequest{
				Text:     "",
				Metadata: "",
			},
			expectedCode:          codes.Internal,
			expectedStatusMessage: "internal error",
			prepare: func() {
				c.textDataService.EXPECT().LoadAllTextData(gomock.Any(), gomock.Any()).Times(1).Return(nil, errors.New("error"))
			},
		},
		{
			name:  "Success - load credit text data",
			token: token,
			body: &pb.GetTextDataRequest{
				Text:     "",
				Metadata: "",
			},
			expectedCode:          codes.OK,
			expectedStatusMessage: "",
			prepare: func() {
				c.textDataService.EXPECT().LoadAllTextData(gomock.Any(), gomock.Any()).Times(1).Return(texts, nil)
			},
			expectedBody: &pb.GetTextDataResponse{Text: textData},
		},
	}
	for _, test := range testCases {
		c.T().Run(test.name, func(t *testing.T) {
			test.prepare()

			header := metadata.New(map[string]string{"token": test.token})
			ctx := metadata.NewOutgoingContext(context.Background(), header)
			conn, err := grpc.NewClient("passthrough:///bufnet",
				grpc.WithContextDialer(c.dialer),
				grpc.WithTransportCredentials(insecure.NewCredentials()))
			require.NoError(t, err)
			defer conn.Close()

			client := pb.NewTextDataServiceClient(conn)
			resp, err := client.GetLoadTextData(ctx, test.body)
			st := status.Convert(err)
			assert.Equal(t, test.expectedCode, st.Code())
			assert.Equal(t, test.expectedStatusMessage, st.Message())
			for _, detail := range st.Details() {
				switch d := detail.(type) { //nolint:gocritic
				case *errdetails.BadRequest:
					for _, violation := range d.GetFieldViolations() {
						assert.Equal(t, test.expectedViolationField, violation.GetField())
						assert.Equal(t, test.expectedViolationDescription, violation.GetDescription())
					}
				}
			}
			if resp != nil {
				assert.Equal(t, test.expectedBody.GetText(), resp.GetText())
			}
		})
	}
}
