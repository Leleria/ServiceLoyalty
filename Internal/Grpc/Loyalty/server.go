package Loyalty

import (
	"context"
	"errors"
	"fmt"
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

//go:generate go run github.com/vektra/mockery/v2@v2.20.2 --name=Loyalty
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
	AddPersonalPromoCode(ctx context.Context, idClient int32, idGroup int32, namePromoCode string,
		typeDiscount int32, valueDiscount int32, dateStartActive string, dateFinishActive string) (
		result string, err error)

	SettingUpBudget(ctx context.Context, typeCashBack int32, condition string, valueBudget int32) (result string, err error)
	ChangeBudgetCashBack(ctx context.Context, idCashBack int32, budget int32) (result string, err error)
	ChangeTypeCashBack(ctx context.Context, idCashBack int32, typeCashBack int32) (result string, err error)
	ChangeConditionCashBack(ctx context.Context, idCashBack int32, condition string) (result string, err error)
	GetCashBack(ctx context.Context, idCashBack int32) (result string, err error)
	GetAllCashBack(ctx context.Context) (result string, err error)
	DeleteCashBack(ctx context.Context, idCashBack int32) (result string, err error)

	CalculatePriceWithPromoCode(ctx context.Context, idClient int32, namePromoCode string, amountProduct float32) (
		result string, finalAmountProduct float32, amountDiscount float32, err error)
	CalculatePriceWithBonuses(ctx context.Context, idClient int32, amountProduct float32) (
		result string, finalAmountProduct float32, numberBonusesDebited float32, err error)
	DebitingPromoBonuses(ctx context.Context, idClient int32, paymentStatus bool) (result string, err error)
	AccrualBonusesCashback(ctx context.Context, idClient int32, idCashBack int32) (result string, err error)
}

func (s *ServerAPI) AccrualBonusesCashback(ctx context.Context,
	in *sl.AccrualBonusesCashbackRequest) (*sl.AccrualBonusesCashbackResponse, error) {
	if !CheckIdForTable(in.GetIdClient()) {
		return nil, status.Error(codes.InvalidArgument, "incorrect id client")
	}
	if !CheckIdForTable(in.GetIdCashBack()) {
		return nil, status.Error(codes.InvalidArgument, "incorrect id cashback")
	}

	result, err := s.loyalty.AccrualBonusesCashback(ctx, in.GetIdClient(), in.GetIdCashBack())
	if err != nil {
		if errors.Is(err, Storage.ErrCashBackExists) {
			return nil, status.Error(codes.NotFound, "cashback not found")
		}
		return nil, status.Error(codes.Internal, "failed")
	}
	return &sl.AccrualBonusesCashbackResponse{Result: result}, nil
}

func (s *ServerAPI) DebitingPromoBonuses(ctx context.Context,
	in *sl.DebitingPromoBonusesRequest) (*sl.DebitingPromoBonusesResponse, error) {
	if !CheckIdForTable(in.GetIdClient()) {
		return nil, status.Error(codes.InvalidArgument, "incorrect id")
	}

	result, err := s.loyalty.DebitingPromoBonuses(ctx, in.GetIdClient(), in.GetPaymentStatus())
	if err != nil {
		if errors.Is(err, Storage.ErrPromoCodeExists) {
			return nil, status.Error(codes.NotFound, "client not found")
		}
		return nil, status.Error(codes.Internal, "failed")
	}
	return &sl.DebitingPromoBonusesResponse{Result: result}, nil
}
func (s *ServerAPI) CalculatePriceWithBonuses(ctx context.Context,
	in *sl.CalculatePriceWithBonusesRequest) (*sl.CalculatePriceWithBonusesResponse, error) {
	if !CheckIdForTable(in.GetIdClient()) {
		return nil, status.Error(codes.InvalidArgument, "incorrect id")
	}
	if !CheckAmountProduct(in.GetAmountProduct()) {
		return nil, status.Error(codes.InvalidArgument, "amount must be more 100")
	}
	resultMessage, resultFinalAmount, resultNumberBonusesDebited, err := s.loyalty.CalculatePriceWithBonuses(ctx, in.GetIdClient(), in.GetAmountProduct())
	if err != nil {
		if errors.Is(err, Storage.ErrClientExists) {
			return nil, status.Error(codes.NotFound, "client not found")
		}
		return nil, status.Error(codes.Internal, "failed")
	}
	return &sl.CalculatePriceWithBonusesResponse{Result: resultMessage + ", " + fmt.Sprintf("%g", resultFinalAmount) + ", " + fmt.Sprintf("%g", resultNumberBonusesDebited)}, nil

}

func (s *ServerAPI) CalculatePriceWithPromoCode(ctx context.Context,
	in *sl.CalculatePriceWithPromoCodeRequest) (*sl.CalculatePriceWithPromoCodeResponse, error) {

	if !CheckIdForTable(in.GetIdClient()) {
		return nil, status.Error(codes.InvalidArgument, "incorrect id")
	}
	if !CheckName(in.GetPromoCode()) {
		return nil, status.Error(codes.InvalidArgument, "incorrect name")
	}
	if !CheckAmountProduct(in.GetAmountProduct()) {
		return nil, status.Error(codes.InvalidArgument, "amount must be more 100")
	}

	resultMessage, resultFinalAmount, resultAmountDiscount, err := s.loyalty.CalculatePriceWithPromoCode(ctx, in.GetIdClient(), in.GetPromoCode(), in.GetAmountProduct())
	if err != nil {
		if errors.Is(err, Storage.ErrPromoCodeExists) {
			return nil, status.Error(codes.NotFound, "promo code not found")
		}
		return nil, status.Error(codes.Internal, "failed")
	}
	return &sl.CalculatePriceWithPromoCodeResponse{Result: resultMessage + ", " + fmt.Sprintf("%g", resultFinalAmount) + ", " + fmt.Sprintf("%g", resultAmountDiscount)}, nil

}

func (s *ServerAPI) DeleteCashBack(ctx context.Context,
	in *sl.DeleteCashBackRequest) (*sl.DeleteCashBackResponse, error) {
	if !CheckIdForTable(in.GetIdCashBack()) {
		return nil, status.Error(codes.InvalidArgument, "incorrect id")
	}
	result, err := s.loyalty.DeleteCashBack(ctx, in.GetIdCashBack())
	if err != nil {
		if errors.Is(err, Storage.ErrPromoCodeExists) {
			return nil, status.Error(codes.NotFound, "cashback not found")
		}
		return nil, status.Error(codes.Internal, "incorrect id")
	}
	return &sl.DeleteCashBackResponse{Result: result}, nil
}

func (s *ServerAPI) ChangeBudgetCashBack(ctx context.Context,
	in *sl.ChangeBudgetCashBackRequest) (*sl.ChangeBudgetCashBackResponse, error) {
	if !CheckIdForTable(in.GetIdCashBack()) {
		return nil, status.Error(codes.InvalidArgument, "incorrect id")
	}
	if !CheckBudgetCashBack(in.GetBudget()) {
		return nil, status.Error(codes.InvalidArgument, "incorrect budget")
	}
	result, err := s.loyalty.ChangeBudgetCashBack(ctx, in.GetIdCashBack(), in.GetBudget())
	if err != nil {
		if errors.Is(err, Storage.ErrCashBackExists) {
			return nil, status.Error(codes.NotFound, "cashback not found")
		}
		return nil, status.Error(codes.Internal, "incorrect data")
	}
	return &sl.ChangeBudgetCashBackResponse{Result: result}, nil
}

func (s *ServerAPI) ChangeTypeCashBack(ctx context.Context,
	in *sl.ChangeTypeCashBackRequest) (*sl.ChangeTypeCashBackResponse, error) {
	if !CheckIdForTable(in.GetIdCashBack()) {
		return nil, status.Error(codes.InvalidArgument, "incorrect id")
	}
	result, err := s.loyalty.ChangeTypeCashBack(ctx, in.GetIdCashBack(), in.GetTypeCashBack())
	if err != nil {
		if errors.Is(err, Storage.ErrCashBackExists) {
			return nil, status.Error(codes.NotFound, "cashback not found")
		}
		return nil, status.Error(codes.Internal, "incorrect data")
	}
	return &sl.ChangeTypeCashBackResponse{Result: result}, nil
}

func (s *ServerAPI) ChangeConditionCashBack(ctx context.Context,
	in *sl.ChangeConditionCashBackRequest) (*sl.ChangeConditionCashBackResponse, error) {
	if !CheckIdForTable(in.GetIdCashBack()) {
		return nil, status.Error(codes.InvalidArgument, "incorrect id")
	}
	result, err := s.loyalty.ChangeConditionCashBack(ctx, in.GetIdCashBack(), in.GetCondition())
	if err != nil {
		if errors.Is(err, Storage.ErrCashBackExists) {
			return nil, status.Error(codes.NotFound, "cashback not found")
		}
		return nil, status.Error(codes.Internal, "incorrect data")
	}
	return &sl.ChangeConditionCashBackResponse{Result: result}, nil
}

func (s *ServerAPI) GetCashBack(ctx context.Context,
	in *sl.GetCashBackRequest) (*sl.GetCashBackResponse, error) {

	if !CheckIdForTable(in.GetIdCashBack()) {
		return nil, status.Error(codes.InvalidArgument, "incorrect id")
	}

	result, err := s.loyalty.GetCashBack(ctx, in.GetIdCashBack())
	if err != nil {
		if errors.Is(err, Storage.ErrCashBackExists) {
			return nil, status.Error(codes.NotFound, "cashback not found")
		}
		return nil, status.Error(codes.Internal, "incorrect data")
	}
	return &sl.GetCashBackResponse{Result: result}, nil
}
func (s *ServerAPI) GetAllCashBack(ctx context.Context,
	in *sl.GetAllCashBackRequest) (*sl.GetAllCashBackResponse, error) {

	result, err := s.loyalty.GetAllCashBack(ctx)
	if err != nil {
		if errors.Is(err, Storage.ErrCashBackExists) {
			return nil, status.Error(codes.NotFound, "cashback not found")
		}
		return nil, status.Error(codes.Internal, "incorrect data")
	}
	return &sl.GetAllCashBackResponse{Result: result}, nil
}

func (s *ServerAPI) SettingUpBudget(ctx context.Context,
	in *sl.SettingUpBudgetRequest) (*sl.SettingUpBudgetResponse, error) {

	flag, message := CheckArgsForSettingUpBudget(in)
	if flag && message == "budget" {
		return nil, status.Error(codes.InvalidArgument, "incorrect value budget")
	}
	if flag && message == "type" {
		return nil, status.Error(codes.InvalidArgument, "incorrect type cashback")
	}
	result, err := s.loyalty.SettingUpBudget(ctx, in.GetTypeCashBack(), in.GetCondition(), in.GetValueBudget())
	if err != nil {
		if errors.Is(err, Storage.ErrSettingUpBudgetExists) {
			return nil, status.Error(codes.AlreadyExists, "cashback exists")
		}
		return nil, status.Error(codes.Internal, "incorrect data")
	}
	return &sl.SettingUpBudgetResponse{Result: result}, nil
}

func (s *ServerAPI) AddPersonalPromoCode(ctx context.Context,
	in *sl.AddPersonalPromoCodeRequest) (*sl.AddPersonalPromoCodeResponse, error) {

	flag, message := CheckArgsForAddedPersonalPromoCodeToDB(in)
	if flag && message == "client" {
		return nil, status.Error(codes.InvalidArgument, "incorrect id client")
	}
	if flag && message == "group" {
		return nil, status.Error(codes.InvalidArgument, "incorrect id group")
	}
	if flag && message == "name" {
		return nil, status.Error(codes.InvalidArgument, "incorrect name")
	}
	if flag && message == "type" {
		return nil, status.Error(codes.InvalidArgument, "incorrect type discount")
	}
	if flag && message == "value" {
		return nil, status.Error(codes.InvalidArgument, "value discount cannot be less then 0")
	}
	if flag && message == "value percent" {
		return nil, status.Error(codes.InvalidArgument, "percentage discount cannot be more than 100")
	}
	if flag && message == "format date start" {
		return nil, status.Error(codes.InvalidArgument, "incorrect format date start")
	}
	if flag && message == "format date finish" {
		return nil, status.Error(codes.InvalidArgument, "incorrect format date finish")
	}
	if flag && message == "date start > date finish" {
		return nil, status.Error(codes.InvalidArgument, "date start cannot be more then date finish")
	}

	result, err := s.loyalty.AddPersonalPromoCode(ctx, in.GetIdClient(), in.GetIdGroup(), in.GetNamePromoCode(), in.TypeDiscount, in.ValueDiscount,
		in.DateStartActive, in.DateFinishActive)
	if err != nil {
		if errors.Is(err, Storage.ErrPersonalPromoCodeExists) {
			return nil, status.Error(codes.AlreadyExists, "personal promo code exists")
		}
		if errors.Is(err, Storage.ErrClientFound) {
			return nil, status.Error(codes.AlreadyExists, "client not found")
		}
		if errors.Is(err, Storage.ErrGroupFound) {
			return nil, status.Error(codes.AlreadyExists, "group not found")
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
		if errors.Is(err, Storage.ErrPromoCodeFound) {
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
		if errors.Is(err, Storage.ErrCashBackFound) {
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
	ctx context.Context, in *sl.AddNewPromoCodeRequest) (*sl.AddNewPromoCodeResponse, error) {

	flag, message := CheckArgsForAddedToDB(in)
	if flag && message == "name" {
		return nil, status.Error(codes.InvalidArgument, "incorrect name")
	}
	if flag && message == "type" {
		return nil, status.Error(codes.InvalidArgument, "incorrect type discount")
	}
	if flag && message == "value" {
		return nil, status.Error(codes.InvalidArgument, "value discount cannot be less then 0")
	}
	if flag && message == "value percent" {
		return nil, status.Error(codes.InvalidArgument, "percentage discount cannot be more than 100")
	}
	if flag && message == "max count" {
		return nil, status.Error(codes.InvalidArgument, "max count uses cannot be less then 0")
	}
	if flag && message == "format date start" {
		return nil, status.Error(codes.InvalidArgument, "incorrect format date start")
	}
	if flag && message == "format date finish" {
		return nil, status.Error(codes.InvalidArgument, "incorrect format date finish")
	}
	if flag && message == "date start > date finish" {
		return nil, status.Error(codes.InvalidArgument, "date start cannot be more then date finish")
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
func CheckArgsForAddedToDB(request *sl.AddNewPromoCodeRequest) (bool, string) {

	if !CheckName(request.Name) {
		return true, "name"
	}
	if request.TypeDiscount != 1 && request.TypeDiscount != 2 {
		return true, "type"
	}
	if request.ValueDiscount <= 0 {
		return true, "value"
	}
	if request.TypeDiscount == 1 && request.ValueDiscount > 100 {
		return true, "value percent"
	}
	if request.MaxCountUses <= 0 {
		return true, "max count"
	}
	layout := "2006-01-02"
	timeStart, err := time.Parse(layout, request.DateStartActive)
	if err != nil {
		return true, "format date start"
	}
	timeFinish, err := time.Parse(layout, request.DateFinishActive)
	if err != nil {
		return true, "format date finish"
	}
	if request.DateStartActive > request.DateFinishActive {
		return true, "date start > date finish"
	}

	request.DateStartActive = timeStart.Format(layout)
	request.DateFinishActive = timeFinish.Format(layout)
	return false, ""
} //true its error

func CheckArgsForAddedPersonalPromoCodeToDB(request *sl.AddPersonalPromoCodeRequest) (bool, string) {

	if !CheckName(request.NamePromoCode) {
		return true, "name"
	}
	if request.IdClient <= 0 {
		return true, "client"
	}
	if request.IdGroup <= 0 {
		return true, "group"
	}
	if request.TypeDiscount != 1 && request.TypeDiscount != 2 {
		return true, "type"
	}
	if request.ValueDiscount <= 0 {
		return true, "value"
	}
	if request.TypeDiscount == 1 && request.ValueDiscount > 100 {
		return true, "value percent"
	}
	layout := "2006-01-02"
	timeStart, err := time.Parse(layout, request.DateStartActive)
	if err != nil {
		return true, "format date start"
	}
	timeFinish, err := time.Parse(layout, request.DateFinishActive)
	if err != nil {
		return true, "format date finish"
	}
	if request.DateStartActive > request.DateFinishActive {
		return true, "date start > date finish"
	}

	request.DateStartActive = timeStart.Format(layout)
	request.DateFinishActive = timeFinish.Format(layout)
	return false, ""
} //true its error

func CheckArgsForSettingUpBudget(request *sl.SettingUpBudgetRequest) (bool, string) {
	if request.ValueBudget <= 0 {
		return true, "budget"
	}
	if request.TypeCashBack <= 0 {
		return true, "type"
	}
	return false, ""
} //true its error

func CheckName(name string) bool {
	size := len(name)
	if size == 5 && IsLetter(name) {
		return true
	}
	return false
}
func CheckIdForTable(id int32) bool {
	if id > 0 {
		return true
	}
	return false
}

func CheckBudgetCashBack(budget int32) bool {
	if budget > 0 {
		return true
	}
	return false
}
func CheckAmountProduct(amount float32) bool {
	if amount >= 100 {
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
