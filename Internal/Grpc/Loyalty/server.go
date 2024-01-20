package Loyalty

import (
	"context"
	"errors"
	sl "github.com/Leleria/Contract/GeneratedFilesProtoBufGo"
	"github.com/Leleria/ServiceLoyalty/Internal/Storage"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
	"unicode"
)

type ServerAPI struct {
	sl.UnimplementedLoyaltyServiceServer
	loyalty Loyalty
}

type Loyalty interface {
	AddNewPromoCode(
		ctx context.Context,
		name string,
		typeDiscount int32,
		valueDiscount int32,
		dateStartActive string,
		dateFinishActive string,
		maxCountUses int32,
	) (result string, err error)
	DeletePromoCode(ctx context.Context, name string) (result string, err error)
}

func (s *ServerAPI) DeletePromoCode(ctx context.Context,
	in *sl.DeletePromoCodeRequest) (*sl.DeletePromoCodeResponse, error) {
	if !CheckName(in.Name) {
		return nil, status.Error(codes.InvalidArgument, "incorrect name")
	}
	result, err := s.loyalty.DeletePromoCode(ctx, in.GetName())
	if err != nil {
		if errors.Is(err, Storage.ErrPromoCodeExists) {
			return nil, status.Error(codes.NotFound, "promo code not found")
		}
		return nil, status.Error(codes.Internal, "incorrect name")
	}
	return &sl.DeletePromoCodeResponse{Result: result}, nil
}
func (s *ServerAPI) AddNewPromoCode(
	ctx context.Context,
	in *sl.AddNewPromoCodeRequest) (*sl.AddNewPromoCodeResponse, error) {

	if CheckArgsForAddedToDB(in) {
		return nil, status.Error(codes.InvalidArgument, "incorrect data")
	}
	result, err := s.loyalty.AddNewPromoCode(ctx, in.GetName(), in.GetTypeDiscount(), in.GetValueDiscount(),
		in.GetDateStartActive(), in.GetDateFinishActive(), in.GetMaxCountUses())
	if err != nil {
		if errors.Is(err, Storage.ErrPromoCodeExists) {
			return nil, status.Error(codes.AlreadyExists, "promo code exists")
		}
		return nil, status.Error(codes.Internal, "incorrect data")
	}
	return &sl.AddNewPromoCodeResponse{Result: result}, nil
}

func Register(gRPCServer *grpc.Server, loyalty Loyalty) {
	sl.RegisterLoyaltyServiceServer(gRPCServer, &ServerAPI{loyalty: loyalty})
}
func CheckArgsForAddedToDB(request *sl.AddNewPromoCodeRequest) bool {

	if !CheckName(request.Name) {
		return true
	}
	if request.TypeDiscount != 1 && request.TypeDiscount != 2 {
		return true
	}
	if request.ValueDiscount <= 0 {
		return true
	}
	if request.TypeDiscount == 1 && request.ValueDiscount > 100 {
		return true
	}
	if request.MaxCountUses <= 0 {
		return true
	}
	if request.DateStartActive > request.DateFinishActive {
		return true
	}
	layout := "2006-01-02"
	timeStart, err := time.Parse(layout, request.DateStartActive)
	if err != nil {
		return true
	}
	timeFinish, err := time.Parse(layout, request.DateFinishActive)
	if err != nil {
		return true
	}
	request.DateStartActive = timeStart.Format(layout)
	request.DateFinishActive = timeFinish.Format(layout)
	return false
} //true its error
func CheckName(name string) bool {
	size := len(name)
	if size == 5 && IsLetter(name) {
		return true
	}
	return false
}
func IsLetter(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}
