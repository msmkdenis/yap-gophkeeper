package lib

import (
	"fmt"

	"google.golang.org/genproto/googleapis/rpc/errdetails"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func UnpackGRPCError(err error) {
	st := status.Convert(err)
	if st.Code() == codes.InvalidArgument {
		for _, detail := range st.Details() {
			switch t := detail.(type) { //nolint:gocritic
			case *errdetails.BadRequest:
				for _, violation := range t.GetFieldViolations() {
					fmt.Printf("The %s field %s\n", violation.GetField(), violation.GetDescription())
				}
			}
		}
	} else {
		fmt.Printf("Please try again: %s\n", st.Message())
	}
}
