package Tests

import (
	sl "github.com/Leleria/Contract/GeneratedFilesProtoBufGo"
	"github.com/Leleria/ServiceLoyalty/Tests/Suite"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestAddPromoCode_AddToDatabase(t *testing.T) {
	ctx, st := Suite.New(t)

	startDate := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)

	testTable := []struct {
		namePromoCode    string
		typeDiscountId   int
		valueDiscount    int
		dateStartActive  string
		dateFinishActive string
		maxCountUses     int
		expected         string
	}{
		{
			namePromoCode:    gofakeit.Lexify("?????"),
			typeDiscountId:   1,
			valueDiscount:    gofakeit.Number(1, 99),
			dateStartActive:  gofakeit.DateRange(startDate, endDate).Format("2006-01-02"),
			dateFinishActive: gofakeit.PastDate().Format("2006-01-02"),
			maxCountUses:     gofakeit.Number(1, 500),
			expected:         "complete",
		},
		{
			namePromoCode:    gofakeit.Lexify("?????"),
			typeDiscountId:   2,
			valueDiscount:    gofakeit.Number(1, 500),
			dateStartActive:  gofakeit.DateRange(startDate, endDate).Format("2006-01-02"),
			dateFinishActive: gofakeit.PastDate().Format("2006-02-03"),
			maxCountUses:     gofakeit.Number(1, 500),
			expected:         "complete",
		},
		{
			namePromoCode:    "TkyzD",
			typeDiscountId:   1,
			valueDiscount:    gofakeit.Number(1, 99),
			dateStartActive:  gofakeit.DateRange(startDate, endDate).Format("2006-01-02"),
			dateFinishActive: gofakeit.PastDate().Format("2006-02-03"),
			maxCountUses:     gofakeit.Number(1, 500),
			expected:         "promo code already exists",
		},
	}
	for _, equal := range testTable {
		result, err := st.DB.SavePromoCode(ctx, equal.namePromoCode, int32(equal.typeDiscountId),
			int32(equal.valueDiscount), equal.dateStartActive, equal.dateFinishActive, int32(equal.maxCountUses))

		if err != nil {
			message := err.Error()
			parts := strings.Split(message, ": ")
			assert.Equal(t, equal.expected, parts[1])
		} else {
			require.NotEmpty(t, result)
			assert.Equal(t, result, "complete")
		}
	}
}
func TestAddPromoCode_CheckNamePromoCode(t *testing.T) {
	ctx, st := Suite.New(t)

	currentDate := time.Now()
	startDate := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(currentDate.Year(), currentDate.Month(), currentDate.Day(), 0, 0, 0, 0, time.UTC)
	char := gofakeit.Password(false, false, false, true, false, 1)

	testTable := []struct {
		namePromoCode    string
		typeDiscountId   int
		valueDiscount    int
		dateStartActive  string
		dateFinishActive string
		maxCountUses     int
		expected         string
	}{
		{
			namePromoCode:    gofakeit.Lexify("??????"),
			typeDiscountId:   1,
			valueDiscount:    gofakeit.Number(1, 99),
			dateStartActive:  gofakeit.DateRange(startDate, endDate).Format("2006-01-02"),
			dateFinishActive: gofakeit.PastDate().Format("2006-01-02"),
			maxCountUses:     gofakeit.Number(1, 500),
			expected:         "incorrect name",
		},
		{
			namePromoCode:    gofakeit.Lexify("????"),
			typeDiscountId:   1,
			valueDiscount:    gofakeit.Number(1, 99),
			dateStartActive:  gofakeit.DateRange(startDate, endDate).Format("2006-01-02"),
			dateFinishActive: gofakeit.PastDate().Format("2006-02-03"),
			maxCountUses:     gofakeit.Number(1, 500),
			expected:         "incorrect name",
		},
		{
			namePromoCode:    gofakeit.Lexify("????") + string(char[0]),
			typeDiscountId:   1,
			valueDiscount:    gofakeit.Number(1, 99),
			dateStartActive:  gofakeit.DateRange(startDate, endDate).Format("2006-01-02"),
			dateFinishActive: gofakeit.PastDate().Format("2006-02-03"),
			maxCountUses:     gofakeit.Number(1, 500),
			expected:         "incorrect name",
		},
		{
			namePromoCode:    gofakeit.Lexify("????") + strconv.Itoa(gofakeit.Number(0, 9)),
			typeDiscountId:   1,
			valueDiscount:    gofakeit.Number(1, 99),
			dateStartActive:  gofakeit.DateRange(startDate, endDate).Format("2006-01-02"),
			dateFinishActive: gofakeit.PastDate().Format("2006-02-03"),
			maxCountUses:     gofakeit.Number(1, 500),
			expected:         "incorrect name",
		},
		{
			namePromoCode:    gofakeit.Lexify("????") + " ",
			typeDiscountId:   1,
			valueDiscount:    gofakeit.Number(1, 99),
			dateStartActive:  gofakeit.DateRange(startDate, endDate).Format("2006-01-02"),
			dateFinishActive: gofakeit.PastDate().Format("2006-02-03"),
			maxCountUses:     gofakeit.Number(1, 500),
			expected:         "incorrect name",
		},
	}

	for _, equal := range testTable {
		_, err := st.LoyaltyServiceClient.AddNewPromoCode(ctx, &sl.AddNewPromoCodeRequest{
			Name:             equal.namePromoCode,
			TypeDiscount:     int32(equal.typeDiscountId),
			ValueDiscount:    int32(equal.valueDiscount),
			DateStartActive:  equal.dateStartActive,
			DateFinishActive: equal.dateFinishActive,
			MaxCountUses:     int32(equal.maxCountUses),
		})
		message := err.Error()
		parts := strings.Split(message, " = ")
		assert.Equal(t, equal.expected, parts[2])

	}
}
func TestAddPromoCode_CheckTypeAndValueDiscountPromoCode(t *testing.T) {
	ctx, st := Suite.New(t)

	currentDate := time.Now()
	startDate := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(currentDate.Year(), currentDate.Month(), currentDate.Day(), 0, 0, 0, 0, time.UTC)

	testTable := []struct {
		namePromoCode    string
		typeDiscountId   int
		valueDiscount    int
		dateStartActive  string
		dateFinishActive string
		maxCountUses     int
		expected         string
	}{
		{
			namePromoCode:    gofakeit.Lexify("?????"),
			typeDiscountId:   gofakeit.Number(-10, 0),
			valueDiscount:    gofakeit.Number(1, 99),
			dateStartActive:  gofakeit.DateRange(startDate, endDate).Format("2006-01-02"),
			dateFinishActive: gofakeit.PastDate().Format("2006-01-02"),
			maxCountUses:     gofakeit.Number(1, 500),
			expected:         "incorrect type discount",
		},
		{
			namePromoCode:    gofakeit.Lexify("?????"),
			typeDiscountId:   gofakeit.Number(3, 10),
			valueDiscount:    gofakeit.Number(1, 99),
			dateStartActive:  gofakeit.DateRange(startDate, endDate).Format("2006-01-02"),
			dateFinishActive: gofakeit.PastDate().Format("2006-02-03"),
			maxCountUses:     gofakeit.Number(1, 500),
			expected:         "incorrect type discount",
		},
		{
			namePromoCode:    gofakeit.Lexify("?????"),
			typeDiscountId:   gofakeit.Number(1, 2),
			valueDiscount:    gofakeit.Number(-10, 0),
			dateStartActive:  gofakeit.DateRange(startDate, endDate).Format("2006-01-02"),
			dateFinishActive: gofakeit.PastDate().Format("2006-02-03"),
			maxCountUses:     gofakeit.Number(1, 500),
			expected:         "value discount cannot be less then 0",
		},
		{
			namePromoCode:    gofakeit.Lexify("?????"),
			typeDiscountId:   1,
			valueDiscount:    gofakeit.Number(100, 500),
			dateStartActive:  gofakeit.DateRange(startDate, endDate).Format("2006-01-02"),
			dateFinishActive: gofakeit.PastDate().Format("2006-02-03"),
			maxCountUses:     gofakeit.Number(1, 500),
			expected:         "percentage discount cannot be more than 100",
		},
	}

	for _, equal := range testTable {
		_, err := st.LoyaltyServiceClient.AddNewPromoCode(ctx, &sl.AddNewPromoCodeRequest{
			Name:             equal.namePromoCode,
			TypeDiscount:     int32(equal.typeDiscountId),
			ValueDiscount:    int32(equal.valueDiscount),
			DateStartActive:  equal.dateStartActive,
			DateFinishActive: equal.dateFinishActive,
			MaxCountUses:     int32(equal.maxCountUses),
		})

		message := err.Error()
		parts := strings.Split(message, " = ")
		assert.Equal(t, equal.expected, parts[2])
	}
}
func TestAddPromoCode_CheckDateStart(t *testing.T) {
	ctx, st := Suite.New(t)

	startDate := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)

	testTable := []struct {
		namePromoCode    string
		typeDiscountId   int
		valueDiscount    int
		dateStartActive  string
		dateFinishActive string
		maxCountUses     int
		expected         string
	}{
		{
			namePromoCode:    gofakeit.Lexify("?????"),
			typeDiscountId:   1,
			valueDiscount:    gofakeit.Number(1, 99),
			dateStartActive:  gofakeit.DateRange(startDate, endDate).Format("02-Jan-2006"),
			dateFinishActive: gofakeit.FutureDate().Format("2006-01-02"),
			maxCountUses:     gofakeit.Number(1, 500),
			expected:         "incorrect format date start",
		},
		{
			namePromoCode:    gofakeit.Lexify("?????"),
			typeDiscountId:   1,
			valueDiscount:    gofakeit.Number(1, 99),
			dateStartActive:  gofakeit.DateRange(startDate, endDate).Format("20060102"),
			dateFinishActive: gofakeit.FutureDate().Format("2006-01-02"),
			maxCountUses:     gofakeit.Number(1, 500),
			expected:         "incorrect format date start",
		},
		{
			namePromoCode:    gofakeit.Lexify("?????"),
			typeDiscountId:   1,
			valueDiscount:    gofakeit.Number(1, 99),
			dateStartActive:  gofakeit.DateRange(startDate, endDate).Format("January 02, 2006"),
			dateFinishActive: gofakeit.FutureDate().Format("2006-01-02"),
			maxCountUses:     gofakeit.Number(1, 500),
			expected:         "incorrect format date start",
		},
		{
			namePromoCode:    gofakeit.Lexify("?????"),
			typeDiscountId:   1,
			valueDiscount:    gofakeit.Number(1, 99),
			dateStartActive:  gofakeit.DateRange(startDate, endDate).Format("01/02/06"),
			dateFinishActive: gofakeit.FutureDate().Format("2006-01-02"),
			maxCountUses:     gofakeit.Number(1, 500),
			expected:         "incorrect format date start",
		}, {
			namePromoCode:    gofakeit.Lexify("?????"),
			typeDiscountId:   1,
			valueDiscount:    gofakeit.Number(1, 99),
			dateStartActive:  gofakeit.DateRange(startDate, endDate).Format("01/02/2006"),
			dateFinishActive: gofakeit.FutureDate().Format("2006-01-02"),
			maxCountUses:     gofakeit.Number(1, 500),
			expected:         "incorrect format date start",
		},
		{
			namePromoCode:    gofakeit.Lexify("?????"),
			typeDiscountId:   1,
			valueDiscount:    gofakeit.Number(1, 99),
			dateStartActive:  gofakeit.DateRange(startDate, endDate).Format("010206"),
			dateFinishActive: gofakeit.FutureDate().Format("2006-01-02"),
			maxCountUses:     gofakeit.Number(1, 500),
			expected:         "incorrect format date start",
		},
		{
			namePromoCode:    gofakeit.Lexify("?????"),
			typeDiscountId:   1,
			valueDiscount:    gofakeit.Number(1, 99),
			dateStartActive:  gofakeit.DateRange(startDate, endDate).Format("Jan-02-06"),
			dateFinishActive: gofakeit.FutureDate().Format("2006-01-02"),
			maxCountUses:     gofakeit.Number(1, 500),
			expected:         "incorrect format date start",
		},
		{
			namePromoCode:    gofakeit.Lexify("?????"),
			typeDiscountId:   1,
			valueDiscount:    gofakeit.Number(1, 99),
			dateStartActive:  gofakeit.DateRange(startDate, endDate).Format("Jan-02-2006"),
			dateFinishActive: gofakeit.FutureDate().Format("2006-01-02"),
			maxCountUses:     gofakeit.Number(1, 500),
			expected:         "incorrect format date start",
		},
		{
			namePromoCode:    gofakeit.Lexify("?????"),
			typeDiscountId:   1,
			valueDiscount:    gofakeit.Number(1, 99),
			dateStartActive:  gofakeit.FutureDate().Format("2006-01-02"),
			dateFinishActive: gofakeit.DateRange(startDate, endDate).Format("2006-01-02"),
			maxCountUses:     gofakeit.Number(1, 500),
			expected:         "date start cannot be more then date finish",
		},
	}
	for _, equal := range testTable {
		_, err := st.LoyaltyServiceClient.AddNewPromoCode(ctx, &sl.AddNewPromoCodeRequest{
			Name:             equal.namePromoCode,
			TypeDiscount:     int32(equal.typeDiscountId),
			ValueDiscount:    int32(equal.valueDiscount),
			DateStartActive:  equal.dateStartActive,
			DateFinishActive: equal.dateFinishActive,
			MaxCountUses:     int32(equal.maxCountUses),
		})

		message := err.Error()
		parts := strings.Split(message, " = ")
		assert.Equal(t, equal.expected, parts[2])
	}
}
func TestAddPromoCode_CheckDateFinish(t *testing.T) {
	ctx, st := Suite.New(t)

	startDate := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)

	testTable := []struct {
		namePromoCode    string
		typeDiscountId   int
		valueDiscount    int
		dateStartActive  string
		dateFinishActive string
		maxCountUses     int
		expected         string
	}{
		{
			namePromoCode:    gofakeit.Lexify("?????"),
			typeDiscountId:   1,
			valueDiscount:    gofakeit.Number(1, 99),
			dateStartActive:  gofakeit.DateRange(startDate, endDate).Format("2006-01-02"),
			dateFinishActive: gofakeit.FutureDate().Format("02-Jan-2006"),
			maxCountUses:     gofakeit.Number(1, 500),
			expected:         "incorrect format date finish",
		},
		{
			namePromoCode:    gofakeit.Lexify("?????"),
			typeDiscountId:   1,
			valueDiscount:    gofakeit.Number(1, 99),
			dateStartActive:  gofakeit.DateRange(startDate, endDate).Format("2006-01-02"),
			dateFinishActive: gofakeit.FutureDate().Format("20060102"),
			maxCountUses:     gofakeit.Number(1, 500),
			expected:         "incorrect format date finish",
		},
		{
			namePromoCode:    gofakeit.Lexify("?????"),
			typeDiscountId:   1,
			valueDiscount:    gofakeit.Number(1, 99),
			dateStartActive:  gofakeit.DateRange(startDate, endDate).Format("2006-01-02"),
			dateFinishActive: gofakeit.FutureDate().Format("January 02, 2006"),
			maxCountUses:     gofakeit.Number(1, 500),
			expected:         "incorrect format date finish",
		},
		{
			namePromoCode:    gofakeit.Lexify("?????"),
			typeDiscountId:   1,
			valueDiscount:    gofakeit.Number(1, 99),
			dateStartActive:  gofakeit.DateRange(startDate, endDate).Format("2006-01-02"),
			dateFinishActive: gofakeit.FutureDate().Format("01/02/06"),
			maxCountUses:     gofakeit.Number(1, 500),
			expected:         "incorrect format date finish",
		}, {
			namePromoCode:    gofakeit.Lexify("?????"),
			typeDiscountId:   1,
			valueDiscount:    gofakeit.Number(1, 99),
			dateStartActive:  gofakeit.DateRange(startDate, endDate).Format("2006-01-02"),
			dateFinishActive: gofakeit.FutureDate().Format("01/02/2006"),
			maxCountUses:     gofakeit.Number(1, 500),
			expected:         "incorrect format date finish",
		},
		{
			namePromoCode:    gofakeit.Lexify("?????"),
			typeDiscountId:   1,
			valueDiscount:    gofakeit.Number(1, 99),
			dateStartActive:  gofakeit.DateRange(startDate, endDate).Format("2006-01-02"),
			dateFinishActive: gofakeit.FutureDate().Format("010206"),
			maxCountUses:     gofakeit.Number(1, 500),
			expected:         "incorrect format date finish",
		},
		{
			namePromoCode:    gofakeit.Lexify("?????"),
			typeDiscountId:   1,
			valueDiscount:    gofakeit.Number(1, 99),
			dateStartActive:  gofakeit.DateRange(startDate, endDate).Format("2006-01-02"),
			dateFinishActive: gofakeit.FutureDate().Format("Jan-02-06"),
			maxCountUses:     gofakeit.Number(1, 500),
			expected:         "incorrect format date finish",
		},
		{
			namePromoCode:    gofakeit.Lexify("?????"),
			typeDiscountId:   1,
			valueDiscount:    gofakeit.Number(1, 99),
			dateStartActive:  gofakeit.DateRange(startDate, endDate).Format("2006-01-02"),
			dateFinishActive: gofakeit.FutureDate().Format("Jan-02-2006"),
			maxCountUses:     gofakeit.Number(1, 500),
			expected:         "incorrect format date finish",
		},
	}
	for _, equal := range testTable {
		_, err := st.LoyaltyServiceClient.AddNewPromoCode(ctx, &sl.AddNewPromoCodeRequest{
			Name:             equal.namePromoCode,
			TypeDiscount:     int32(equal.typeDiscountId),
			ValueDiscount:    int32(equal.valueDiscount),
			DateStartActive:  equal.dateStartActive,
			DateFinishActive: equal.dateFinishActive,
			MaxCountUses:     int32(equal.maxCountUses),
		})

		message := err.Error()
		parts := strings.Split(message, " = ")
		assert.Equal(t, equal.expected, parts[2])
	}
}
func TestAddPromoCode_CheckMaxCountUses(t *testing.T) {
	ctx, st := Suite.New(t)

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
	maxCountUses := gofakeit.Number(-10, 0)
	expected := "max count uses cannot be less then 0"

	_, err := st.LoyaltyServiceClient.AddNewPromoCode(ctx, &sl.AddNewPromoCodeRequest{
		Name:             namePromoCode,
		TypeDiscount:     int32(typeDiscountId),
		ValueDiscount:    int32(valueDiscount),
		DateStartActive:  dateStartActive,
		DateFinishActive: dateFinishActive,
		MaxCountUses:     int32(maxCountUses),
	})
	message := err.Error()
	parts := strings.Split(message, " = ")
	assert.Equal(t, expected, parts[2])
}
