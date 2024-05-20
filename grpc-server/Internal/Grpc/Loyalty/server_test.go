package Loyalty

import (
	"context"
	sl "github.com/Leleria/Contract/GeneratedFilesProtoBufGo"
	"github.com/Leleria/ServiceLoyalty/Internal/Grpc/Loyalty/mocks"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestServerAPI_AddNewPromoCode(t *testing.T) {
	startDate := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)

	namePromoCode := gofakeit.Lexify("?????")
	typeDiscountId := gofakeit.Number(1, 2)
	var valueDiscount int
	if typeDiscountId == 1 {
		valueDiscount = gofakeit.Number(1, 99)
	}
	if typeDiscountId == 2 {
		valueDiscount = gofakeit.Number(1, 500)
	}
	dateStartActive := gofakeit.DateRange(startDate, endDate).Format("2006-01-02")
	dateFinishActive := gofakeit.FutureDate().Format("2006-01-02")
	maxCountUses := gofakeit.Number(1, 500)

	ctx := context.Background()
	loyalty := mocks.NewLoyalty(t)

	req := &sl.AddNewPromoCodeRequest{
		Name:             namePromoCode,
		TypeDiscount:     int32(typeDiscountId),
		ValueDiscount:    int32(valueDiscount),
		DateStartActive:  dateStartActive,
		DateFinishActive: dateFinishActive,
		MaxCountUses:     int32(maxCountUses),
	}
	loyalty.On("AddNewPromoCode", ctx, namePromoCode, int32(typeDiscountId), int32(valueDiscount),
		dateStartActive, dateFinishActive, int32(maxCountUses)).Return(nil, nil)
	server := &ServerAPI{
		loyalty: loyalty,
	}
	resp, _ := server.AddNewPromoCode(ctx, req)
	assert.Equal(t, resp, "complete")
}
