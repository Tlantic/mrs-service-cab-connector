package auth

import (
	"github.com/Tlantic/go-util/errors"
	pb "github.com/Tlantic/mrs-service-cab-connector/proto/store_pb"
	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func errorFromGRPC(err error) error {
	st := status.Convert(err)
	if st.Code() == codes.OK {
		return nil
	}
	for _, detail := range st.Proto().GetDetails() {
		detailErr := pb.Error{}
		if !ptypes.Is(detail, &detailErr) {
			continue
		}
		if unmarshalErr := ptypes.UnmarshalAny(detail, &detailErr); unmarshalErr != nil {
			return errors.Wrap(unmarshalErr, err.Error())
		}

		break
	}

	return errors.WithStack(err)
}
