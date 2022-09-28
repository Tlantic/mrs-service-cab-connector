package stores

import (
	"github.com/Tlantic/go-util/errors"
	pb "github.com/Tlantic/mrs-service-cab-connector/proto/store_pb"
	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

/*
 SERVICE CODE    ENTITY CODE    ERROR CODE
 	 __              __            ___
*/

const (
	UnknownErrorCode = 2000000

	StoreLockedErrorCode                = 2010100
	StoreAlreadyExistsErrorCode         = 2010200
	StoreNotFoundErrorCode              = 2010300
	StoreInvalidArgumentErrorCode       = 2010400
	StoreRequiredFieldMissingErrorCode  = 2010401
	StoreInvalidLatitudeErrorCode       = 2010402
	StoreInvalidLongitudeErrorCode      = 2010403
	StoreInvalidTimezoneErrorCode       = 2010404
	UserIDRequiredFieldMissingErrorCode = 2010405
)

var (
	ErrUnknown = errors.WithCode(errors.New("unknown error"), UnknownErrorCode)

	ErrStoreLocked                = errors.WithCode(errors.New("store currently locked"), StoreLockedErrorCode)
	ErrStoreAlreadyExists         = errors.WithCode(errors.New("store already exists"), StoreAlreadyExistsErrorCode)
	ErrStoreNotFound              = errors.WithCode(errors.New("store not found"), StoreNotFoundErrorCode)
	ErrStoreInvalidArgument       = errors.WithCode(errors.New("invalid argument"), StoreInvalidArgumentErrorCode)
	ErrStoreRequiredFieldMissing  = errors.WithCode(errors.New("required field is empty or missing"), StoreRequiredFieldMissingErrorCode)
	ErrStoreInvalidLatitude       = errors.WithCode(errors.New("invalid latitude value"), StoreInvalidLatitudeErrorCode)
	ErrStoreInvalidLongitude      = errors.WithCode(errors.New("invalid longitude value"), StoreInvalidLongitudeErrorCode)
	ErrStoreInvalidTimezone       = errors.WithCode(errors.New("invalid timezone value"), StoreInvalidTimezoneErrorCode)
	ErrUserIDRequiredFieldMissing = errors.WithCode(errors.New("required field is empty or missing"), UserIDRequiredFieldMissingErrorCode)
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
		switch detailErr.GetCode() {
		case StoreLockedErrorCode:
			return errors.As(st.Err(), errors.WithStack(ErrStoreLocked))
		case StoreAlreadyExistsErrorCode:
			return errors.As(st.Err(), errors.WithStack(ErrStoreAlreadyExists))
		case StoreNotFoundErrorCode:
			return errors.As(st.Err(), errors.WithStack(ErrStoreNotFound))
		case StoreInvalidArgumentErrorCode:
			return errors.As(st.Err(), errors.WithStack(ErrStoreInvalidArgument))
		case StoreRequiredFieldMissingErrorCode:
			return errors.As(st.Err(), errors.WithStack(ErrStoreRequiredFieldMissing))
		case StoreInvalidLatitudeErrorCode:
			return errors.As(st.Err(), errors.WithStack(ErrStoreInvalidLatitude))
		case StoreInvalidLongitudeErrorCode:
			return errors.As(st.Err(), errors.WithStack(ErrStoreInvalidLongitude))
		case StoreInvalidTimezoneErrorCode:
			return errors.As(st.Err(), errors.WithStack(ErrStoreInvalidTimezone))
		case UserIDRequiredFieldMissingErrorCode:
			return errors.As(st.Err(), errors.WithStack(ErrUserIDRequiredFieldMissing))
		case UnknownErrorCode:
			return errors.As(st.Err(), errors.WithStack(ErrUnknown))
		default:
			break
		}
	}
	return errors.AsTemporary(errors.WithStack(err))
}
