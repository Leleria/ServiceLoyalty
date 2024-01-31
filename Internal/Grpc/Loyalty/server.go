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
	GetPromoCode(ctx context.Context, name string) (result string, err error)
	GetAllPromoCodes(ctx context.Context) (result string, err error)
	DeletePromoCode(ctx context.Context, name string) (result string, err error)
	ChangeNamePromoCode(ctx context.Context, name string, newName string) (result string, err error)
	ChangeTypeDiscountPromoCode(ctx context.Context, name string, typeDiscount int32) (result string, err error)
	ChangeValueDiscountPromoCode(ctx context.Context, name string, valueDiscount int32) (result string, err error)
	ChangeDateStartActivePromoCode(ctx context.Context, name string, dateStartActive string) (result string, err error)
	ChangeDateFinishActivePromoCode(ctx context.Context, name string, dateFinish string) (result string, err error)
	ChangeMaxCountUsesPromoCode(ctx context.Context, name string, maxCountUses int32) (result string, err error)
	AddPersonalPromoCode(ctx context.Context, idClient int32, idGroup int32, namePromoCode string) (result string, err error)

	SettingUpBudget(ctx context.Context, typeCashBack int32, condition string, valueBudget int32) (result string, err error)
}

func (s *ServerAPI) SettingUpBudget(ctx context.Context,
	in *sl.SettingUpBudgetRequest) (*sl.SettingUpBudgetResponse, error) {

	if CheckArgsForSettingUpBudget(in) {
		return nil, status.Error(codes.InvalidArgument, "incorrect data")
	}
	result, err := s.loyalty.SettingUpBudget(ctx, in.GetTypeCashBack(), in.GetCondition(), in.GetValueBudget())
	if err != nil {
		if errors.Is(err, Storage.ErrPromoCodeExists) {
			return nil, status.Error(codes.AlreadyExists, "cashback exists")
		}
		return nil, status.Error(codes.Internal, "incorrect data")
	}
	return &sl.SettingUpBudgetResponse{Result: result}, nil
}

func (s *ServerAPI) AddPersonalPromoCode(ctx context.Context,
	in *sl.AddPersonalPromoCodeRequest) (*sl.AddPersonalPromoCodeResponse, error) {

	if CheckArgsForAddedPersonalPromoCodeToDB(in) {
		return nil, status.Error(codes.InvalidArgument, "incorrect data")
	}
	result, err := s.loyalty.AddPersonalPromoCode(ctx, in.GetIdClient(), in.GetIdGroup(), in.GetNamePromoCode())
	if err != nil {
		if errors.Is(err, Storage.ErrPromoCodeExists) {
			return nil, status.Error(codes.AlreadyExists, "personal promo code exists")
		}
		return nil, status.Error(codes.Internal, "incorrect data")
	}
	return &sl.AddPersonalPromoCodeResponse{Result: result}, nil
}

func (s *ServerAPI) GetPromoCode(ctx context.Context,
	in *sl.GetPromoCodeRequest) (*sl.GetPromoCodeResponse, error) {

	if !CheckName(in.GetName()) {
		return nil, status.Error(codes.InvalidArgument, "incorrect name")
	}

	result, err := s.loyalty.GetPromoCode(ctx, in.GetName())
	if err != nil {
		if errors.Is(err, Storage.ErrPromoCodeExists) {
			return nil, status.Error(codes.NotFound, "promo code not found")
		}
		return nil, status.Error(codes.Internal, "incorrect name")
	}
	return &sl.GetPromoCodeResponse{Result: result}, nil
}
func (s *ServerAPI) GetAllPromoCodes(ctx context.Context,
	in *sl.GetAllPromoCodesRequest) (*sl.GetAllPromoCodesResponse, error) {

	result, err := s.loyalty.GetAllPromoCodes(ctx)
	if err != nil {
		if errors.Is(err, Storage.ErrPromoCodeExists) {
			return nil, status.Error(codes.NotFound, "promo code not found")
		}
		return nil, status.Error(codes.Internal, "incorrect name")
	}
	return &sl.GetAllPromoCodesResponse{Result: result}, nil
}

func (s *ServerAPI) ChangeNamePromoCode(ctx context.Context,
	in *sl.ChangeNamePromoCodeRequest) (*sl.ChangeNamePromoCodeResponse, error) {
	if !CheckName(in.GetName()) {
		return nil, status.Error(codes.InvalidArgument, "incorrect name")
	}
	if !CheckName(in.GetNewName()) {
		return nil, status.Error(codes.InvalidArgument, "incorrect newName")
	}
	result, err := s.loyalty.ChangeNamePromoCode(ctx, in.GetName(), in.GetNewName())
	if err != nil {
		if errors.Is(err, Storage.ErrPromoCodeExists) {
			return nil, status.Error(codes.NotFound, "promo code not found")
		}
		return nil, status.Error(codes.Internal, "incorrect name")
	}
	return &sl.ChangeNamePromoCodeResponse{Result: result}, nil
}
func (s *ServerAPI) ChangeTypeDiscountPromoCode(ctx context.Context,
	in *sl.ChangeTypeDiscountPromoCodeRequest) (*sl.ChangeTypeDiscountPromoCodeResponse, error) {
	if !CheckName(in.Name) {
		return nil, status.Error(codes.InvalidArgument, "incorrect name")
	}
	if in.TypeDiscount != 1 && in.TypeDiscount != 2 {
		return nil, status.Error(codes.InvalidArgument, "incorrect name")
	}
	result, err := s.loyalty.ChangeTypeDiscountPromoCode(ctx, in.GetName(), in.TypeDiscount)
	if err != nil {
		if errors.Is(err, Storage.ErrPromoCodeExists) {
			return nil, status.Error(codes.NotFound, "promo code not found")
		}
		return nil, status.Error(codes.Internal, "incorrect type promo code")
	}
	return &sl.ChangeTypeDiscountPromoCodeResponse{Result: result}, nil
}
func (s *ServerAPI) ChangeValueDiscountPromoCode(ctx context.Context,
	in *sl.ChangeValueDiscountPromoCodeRequest) (*sl.ChangeValueDiscountPromoCodeResponse, error) {
	if !CheckName(in.GetName()) {
		return nil, status.Error(codes.InvalidArgument, "incorrect name")
	}
	if in.GetValueDiscount() <= 0 {
		return nil, status.Error(codes.InvalidArgument, "incorrect value")
	}
	result, err := s.loyalty.ChangeValueDiscountPromoCode(ctx, in.GetName(), in.GetValueDiscount())
	if err != nil {
		if errors.Is(err, Storage.ErrPromoCodeExists) {
			return nil, status.Error(codes.NotFound, "promo code not found")
		}
		return nil, status.Error(codes.Internal, "incorrect value")
	}
	return &sl.ChangeValueDiscountPromoCodeResponse{Result: result}, nil
}
func (s *ServerAPI) ChangeDateStartActivePromoCode(ctx context.Context,
	in *sl.ChangeDateStartActivePromoCodeRequest) (*sl.ChangeDateStartActivePromoCodeResponse, error) {
	if !CheckName(in.Name) {
		return nil, status.Error(codes.InvalidArgument, "incorrect name")
	}
	layout := "2006-01-02"
	timeStart, err := time.Parse(layout, in.GetDateStartActive())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "incorrect date")
	}

	result, err := s.loyalty.ChangeDateStartActivePromoCode(ctx, in.GetName(), timeStart.String())
	if err != nil {
		if errors.Is(err, Storage.ErrPromoCodeExists) {
			return nil, status.Error(codes.NotFound, "promo code not found")
		}
		return nil, status.Error(codes.Internal, "incorrect date")
	}
	return &sl.ChangeDateStartActivePromoCodeResponse{Result: result}, nil
}
func (s *ServerAPI) ChangeDateFinishActivePromoCode(ctx context.Context,
	in *sl.ChangeDateFinishActivePromoCodeRequest) (*sl.ChangeDateFinishActivePromoCodeResponse, error) {
	if !CheckName(in.Name) {
		return nil, status.Error(codes.InvalidArgument, "incorrect name")
	}
	layout := "2006-01-02"
	timeFinish, err := time.Parse(layout, in.GetDateFinishActive())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "incorrect date")
	}
	result, err := s.loyalty.ChangeDateFinishActivePromoCode(ctx, in.GetName(), timeFinish.String())
	if err != nil {
		if errors.Is(err, Storage.ErrPromoCodeExists) {
			return nil, status.Error(codes.NotFound, "promo code not found")
		}
		return nil, status.Error(codes.Internal, "incorrect date")
	}
	return &sl.ChangeDateFinishActivePromoCodeResponse{Result: result}, nil
}
func (s *ServerAPI) ChangeMaxCountUsesPromoCode(ctx context.Context,
	in *sl.ChangeMaxCountUsesPromoCodeRequest) (*sl.ChangeMaxCountUsesPromoCodeResponse, error) {
	if !CheckName(in.Name) {
		return nil, status.Error(codes.InvalidArgument, "incorrect name")
	}
	if in.GetMaxCountUses() <= 0 {
		return nil, status.Error(codes.InvalidArgument, "incorrect value")
	}
	result, err := s.loyalty.ChangeMaxCountUsesPromoCode(ctx, in.GetName(), in.GetMaxCountUses())
	if err != nil {
		if errors.Is(err, Storage.ErrPromoCodeExists) {
			return nil, status.Error(codes.NotFound, "promo code not found")
		}
		return nil, status.Error(codes.Internal, "incorrect value")
	}
	return &sl.ChangeMaxCountUsesPromoCodeResponse{Result: result}, nil
}
func (s *ServerAPI) DeletePromoCode(ctx context.Context,
	in *sl.DeletePromoCodeRequest) (*sl.DeletePromoCodeResponse, error) {
	if !CheckName(in.GetName()) {
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

func CheckArgsForAddedPersonalPromoCodeToDB(request *sl.AddPersonalPromoCodeRequest) bool {

	if !CheckName(request.NamePromoCode) {
		return true
	}
	if request.IdClient <= 0 {
		return true
	}
	if request.IdGroup <= 0 {
		return true
	}
	return false
} //true its error

func CheckArgsForSettingUpBudget(request *sl.SettingUpBudgetRequest) bool {
	if request.ValueBudget <= 0 {
		return true
	}
	if request.TypeCashBack <= 0 {
		return true
	}
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
