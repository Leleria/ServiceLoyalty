package Tests

import (
	sl "github.com/Leleria/Contract/GeneratedFilesProtoBufGo"
	"github.com/Leleria/ServiceLoyalty/Tests/Suite"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"strconv"
	"strings"
	"testing"
	"time"
)

//	func TestAddPersonalPromoCode_AddToDatabase(t *testing.T) {
//		ctx, st := Suite.New(t)
//
//		startDate := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
//		endDate := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
//
//		testTable := []struct {
//			namePromoCode    string
//			typeDiscountId   int
//			valueDiscount    int
//			dateStartActive  string
//			dateFinishActive string
//			maxCountUses     int
//			expected         string
//		}{
//			{
//				namePromoCode:    gofakeit.Lexify("?????"),
//				typeDiscountId:   1,
//				valueDiscount:    gofakeit.Number(1, 99),
//				dateStartActive:  gofakeit.DateRange(startDate, endDate).Format("2006-01-02"),
//				dateFinishActive: gofakeit.PastDate().Format("2006-01-02"),
//				maxCountUses:     gofakeit.Number(1, 500),
//				expected:         "complete",
//			},
//			{
//				namePromoCode:    gofakeit.Lexify("?????"),
//				typeDiscountId:   2,
//				valueDiscount:    gofakeit.Number(1, 500),
//				dateStartActive:  gofakeit.DateRange(startDate, endDate).Format("2006-01-02"),
//				dateFinishActive: gofakeit.PastDate().Format("2006-02-03"),
//				maxCountUses:     gofakeit.Number(1, 500),
//				expected:         "complete",
//			},
//			{
//				namePromoCode:    "TkyzD",
//				typeDiscountId:   1,
//				valueDiscount:    gofakeit.Number(1, 99),
//				dateStartActive:  gofakeit.DateRange(startDate, endDate).Format("2006-01-02"),
//				dateFinishActive: gofakeit.PastDate().Format("2006-02-03"),
//				maxCountUses:     gofakeit.Number(1, 500),
//				expected:         "promo code already exists",
//			},
//		}
//		for _, equal := range testTable {
//			result, err := st.DB.SavePersoPromoCode(ctx, equal.namePromoCode, int32(equal.typeDiscountId),
//				int32(equal.valueDiscount), equal.dateStartActive, equal.dateFinishActive, int32(equal.maxCountUses))
//
//			if err != nil {
//				message := err.Error()
//				parts := strings.Split(message, ": ")
//				assert.Equal(t, equal.expected, parts[1])
//			} else {
//				require.NotEmpty(t, result)
//				assert.Equal(t, result, "complete")
//			}
//		}
//	}
func TestAddPersonalPromoCode_CheckIdClientPersonalPromoCode(t *testing.T) {
	ctx, st := Suite.New(t)

	currentDate := time.Now()
	startDate := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(currentDate.Year(), currentDate.Month(), currentDate.Day(), 0, 0, 0, 0, time.UTC)
	expected := "incorrect id client"

	idClient := gofakeit.Number(-10, 0)
	idGroup := 1
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

	_, err := st.LoyaltyServiceClient.AddPersonalPromoCode(ctx, &sl.AddPersonalPromoCodeRequest{
		IdClient:         int32(idClient),
		IdGroup:          int32(idGroup),
		NamePromoCode:    namePromoCode,
		TypeDiscount:     int32(typeDiscountId),
		ValueDiscount:    int32(valueDiscount),
		DateStartActive:  dateStartActive,
		DateFinishActive: dateFinishActive,
	})

	message := err.Error()
	parts := strings.Split(message, " = ")
	assert.Equal(t, expected, parts[2])
}
func TestAddPersonalPromoCode_CheckIdGroupPersonalPromoCode(t *testing.T) {
	ctx, st := Suite.New(t)

	currentDate := time.Now()
	startDate := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(currentDate.Year(), currentDate.Month(), currentDate.Day(), 0, 0, 0, 0, time.UTC)
	expected := "incorrect id group"

	idClient := 1
	idGroup := gofakeit.Number(-10, 0)
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

	_, err := st.LoyaltyServiceClient.AddPersonalPromoCode(ctx, &sl.AddPersonalPromoCodeRequest{
		IdClient:         int32(idClient),
		IdGroup:          int32(idGroup),
		NamePromoCode:    namePromoCode,
		TypeDiscount:     int32(typeDiscountId),
		ValueDiscount:    int32(valueDiscount),
		DateStartActive:  dateStartActive,
		DateFinishActive: dateFinishActive,
	})

	message := err.Error()
	parts := strings.Split(message, " = ")
	assert.Equal(t, expected, parts[2])
}
func TestAddPersonalPromoCode_CheckNamePersonalPromoCode(t *testing.T) {
	ctx, st := Suite.New(t)

	currentDate := time.Now()
	startDate := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(currentDate.Year(), currentDate.Month(), currentDate.Day(), 0, 0, 0, 0, time.UTC)
	char := gofakeit.Password(false, false, false, true, false, 1)

	testTable := []struct {
		idClient         int
		idGroup          int
		namePromoCode    string
		typeDiscountId   int
		valueDiscount    int
		dateStartActive  string
		dateFinishActive string
		expected         string
	}{
		{
			idClient:         1,
			idGroup:          1,
			namePromoCode:    gofakeit.Lexify("??????"),
			typeDiscountId:   1,
			valueDiscount:    gofakeit.Number(1, 99),
			dateStartActive:  gofakeit.DateRange(startDate, endDate).Format("2006-01-02"),
			dateFinishActive: gofakeit.PastDate().Format("2006-01-02"),
			expected:         "incorrect name",
		},
		{
			idClient:         1,
			idGroup:          1,
			namePromoCode:    gofakeit.Lexify("????"),
			typeDiscountId:   1,
			valueDiscount:    gofakeit.Number(1, 99),
			dateStartActive:  gofakeit.DateRange(startDate, endDate).Format("2006-01-02"),
			dateFinishActive: gofakeit.PastDate().Format("2006-02-03"),
			expected:         "incorrect name",
		},
		{
			idClient:         1,
			idGroup:          1,
			namePromoCode:    gofakeit.Lexify("????") + string(char[0]),
			typeDiscountId:   1,
			valueDiscount:    gofakeit.Number(1, 99),
			dateStartActive:  gofakeit.DateRange(startDate, endDate).Format("2006-01-02"),
			dateFinishActive: gofakeit.PastDate().Format("2006-02-03"),
			expected:         "incorrect name",
		},
		{
			idClient:         1,
			idGroup:          1,
			namePromoCode:    gofakeit.Lexify("????") + strconv.Itoa(gofakeit.Number(0, 9)),
			typeDiscountId:   1,
			valueDiscount:    gofakeit.Number(1, 99),
			dateStartActive:  gofakeit.DateRange(startDate, endDate).Format("2006-01-02"),
			dateFinishActive: gofakeit.PastDate().Format("2006-02-03"),
			expected:         "incorrect name",
		},
		{
			idClient:         1,
			idGroup:          1,
			namePromoCode:    gofakeit.Lexify("????") + " ",
			typeDiscountId:   1,
			valueDiscount:    gofakeit.Number(1, 99),
			dateStartActive:  gofakeit.DateRange(startDate, endDate).Format("2006-01-02"),
			dateFinishActive: gofakeit.PastDate().Format("2006-02-03"),
			expected:         "incorrect name",
		},
	}

	for _, equal := range testTable {
		_, err := st.LoyaltyServiceClient.AddPersonalPromoCode(ctx, &sl.AddPersonalPromoCodeRequest{
			IdClient:         int32(equal.idClient),
			IdGroup:          int32(equal.idGroup),
			NamePromoCode:    equal.namePromoCode,
			TypeDiscount:     int32(equal.typeDiscountId),
			ValueDiscount:    int32(equal.valueDiscount),
			DateStartActive:  equal.dateStartActive,
			DateFinishActive: equal.dateFinishActive,
		})
		message := err.Error()
		parts := strings.Split(message, " = ")
		assert.Equal(t, equal.expected, parts[2])

	}
}
func TestAddPersonalPromoCode_CheckTypeAndValueDiscountPersonalPromoCode(t *testing.T) {
	ctx, st := Suite.New(t)

	currentDate := time.Now()
	startDate := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(currentDate.Year(), currentDate.Month(), currentDate.Day(), 0, 0, 0, 0, time.UTC)

	testTable := []struct {
		idClient         int
		idGroup          int
		namePromoCode    string
		typeDiscountId   int
		valueDiscount    int
		dateStartActive  string
		dateFinishActive string
		expected         string
	}{
		{
			idClient:         1,
			idGroup:          1,
			namePromoCode:    gofakeit.Lexify("?????"),
			typeDiscountId:   gofakeit.Number(-10, 0),
			valueDiscount:    gofakeit.Number(1, 99),
			dateStartActive:  gofakeit.DateRange(startDate, endDate).Format("2006-01-02"),
			dateFinishActive: gofakeit.PastDate().Format("2006-01-02"),
			expected:         "incorrect type discount",
		},
		{
			idClient:         1,
			idGroup:          1,
			namePromoCode:    gofakeit.Lexify("?????"),
			typeDiscountId:   gofakeit.Number(3, 10),
			valueDiscount:    gofakeit.Number(1, 99),
			dateStartActive:  gofakeit.DateRange(startDate, endDate).Format("2006-01-02"),
			dateFinishActive: gofakeit.PastDate().Format("2006-02-03"),
			expected:         "incorrect type discount",
		},
		{
			idClient:         1,
			idGroup:          1,
			namePromoCode:    gofakeit.Lexify("?????"),
			typeDiscountId:   gofakeit.Number(1, 2),
			valueDiscount:    gofakeit.Number(-10, 0),
			dateStartActive:  gofakeit.DateRange(startDate, endDate).Format("2006-01-02"),
			dateFinishActive: gofakeit.PastDate().Format("2006-02-03"),
			expected:         "value discount cannot be less then 0",
		},
		{
			idClient:         1,
			idGroup:          1,
			namePromoCode:    gofakeit.Lexify("?????"),
			typeDiscountId:   1,
			valueDiscount:    gofakeit.Number(100, 500),
			dateStartActive:  gofakeit.DateRange(startDate, endDate).Format("2006-01-02"),
			dateFinishActive: gofakeit.PastDate().Format("2006-02-03"),
			expected:         "percentage discount cannot be more than 100",
		},
	}

	for _, equal := range testTable {
		_, err := st.LoyaltyServiceClient.AddPersonalPromoCode(ctx, &sl.AddPersonalPromoCodeRequest{
			IdClient:         int32(equal.idClient),
			IdGroup:          int32(equal.idGroup),
			NamePromoCode:    equal.namePromoCode,
			TypeDiscount:     int32(equal.typeDiscountId),
			ValueDiscount:    int32(equal.valueDiscount),
			DateStartActive:  equal.dateStartActive,
			DateFinishActive: equal.dateFinishActive,
		})

		message := err.Error()
		parts := strings.Split(message, " = ")
		assert.Equal(t, equal.expected, parts[2])
	}
}
func TestAddPersonalPromoCode_CheckDateStartPersonalPromoCode(t *testing.T) {
	ctx, st := Suite.New(t)

	startDate := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)

	testTable := []struct {
		idClient         int
		idGroup          int
		namePromoCode    string
		typeDiscountId   int
		valueDiscount    int
		dateStartActive  string
		dateFinishActive string
		expected         string
	}{
		{
			idClient:         1,
			idGroup:          1,
			namePromoCode:    gofakeit.Lexify("?????"),
			typeDiscountId:   1,
			valueDiscount:    gofakeit.Number(1, 99),
			dateStartActive:  gofakeit.DateRange(startDate, endDate).Format("02-Jan-2006"),
			dateFinishActive: gofakeit.FutureDate().Format("2006-01-02"),
			expected:         "incorrect format date start",
		},
		{
			idClient:         1,
			idGroup:          1,
			namePromoCode:    gofakeit.Lexify("?????"),
			typeDiscountId:   1,
			valueDiscount:    gofakeit.Number(1, 99),
			dateStartActive:  gofakeit.DateRange(startDate, endDate).Format("20060102"),
			dateFinishActive: gofakeit.FutureDate().Format("2006-01-02"),
			expected:         "incorrect format date start",
		},
		{
			idClient:         1,
			idGroup:          1,
			namePromoCode:    gofakeit.Lexify("?????"),
			typeDiscountId:   1,
			valueDiscount:    gofakeit.Number(1, 99),
			dateStartActive:  gofakeit.DateRange(startDate, endDate).Format("January 02, 2006"),
			dateFinishActive: gofakeit.FutureDate().Format("2006-01-02"),
			expected:         "incorrect format date start",
		},
		{
			idClient:         1,
			idGroup:          1,
			namePromoCode:    gofakeit.Lexify("?????"),
			typeDiscountId:   1,
			valueDiscount:    gofakeit.Number(1, 99),
			dateStartActive:  gofakeit.DateRange(startDate, endDate).Format("01/02/06"),
			dateFinishActive: gofakeit.FutureDate().Format("2006-01-02"),
			expected:         "incorrect format date start",
		}, {
			idClient:         1,
			idGroup:          1,
			namePromoCode:    gofakeit.Lexify("?????"),
			typeDiscountId:   1,
			valueDiscount:    gofakeit.Number(1, 99),
			dateStartActive:  gofakeit.DateRange(startDate, endDate).Format("01/02/2006"),
			dateFinishActive: gofakeit.FutureDate().Format("2006-01-02"),
			expected:         "incorrect format date start",
		},
		{
			idClient:         1,
			idGroup:          1,
			namePromoCode:    gofakeit.Lexify("?????"),
			typeDiscountId:   1,
			valueDiscount:    gofakeit.Number(1, 99),
			dateStartActive:  gofakeit.DateRange(startDate, endDate).Format("010206"),
			dateFinishActive: gofakeit.FutureDate().Format("2006-01-02"),
			expected:         "incorrect format date start",
		},
		{
			idClient:         1,
			idGroup:          1,
			namePromoCode:    gofakeit.Lexify("?????"),
			typeDiscountId:   1,
			valueDiscount:    gofakeit.Number(1, 99),
			dateStartActive:  gofakeit.DateRange(startDate, endDate).Format("Jan-02-06"),
			dateFinishActive: gofakeit.FutureDate().Format("2006-01-02"),
			expected:         "incorrect format date start",
		},
		{
			idClient:         1,
			idGroup:          1,
			namePromoCode:    gofakeit.Lexify("?????"),
			typeDiscountId:   1,
			valueDiscount:    gofakeit.Number(1, 99),
			dateStartActive:  gofakeit.DateRange(startDate, endDate).Format("Jan-02-2006"),
			dateFinishActive: gofakeit.FutureDate().Format("2006-01-02"),
			expected:         "incorrect format date start",
		},
		{
			idClient:         1,
			idGroup:          1,
			namePromoCode:    gofakeit.Lexify("?????"),
			typeDiscountId:   1,
			valueDiscount:    gofakeit.Number(1, 99),
			dateStartActive:  gofakeit.FutureDate().Format("2006-01-02"),
			dateFinishActive: gofakeit.DateRange(startDate, endDate).Format("2006-01-02"),
			expected:         "date start cannot be more then date finish",
		},
	}
	for _, equal := range testTable {
		_, err := st.LoyaltyServiceClient.AddPersonalPromoCode(ctx, &sl.AddPersonalPromoCodeRequest{
			IdClient:         int32(equal.idClient),
			IdGroup:          int32(equal.idGroup),
			NamePromoCode:    equal.namePromoCode,
			TypeDiscount:     int32(equal.typeDiscountId),
			ValueDiscount:    int32(equal.valueDiscount),
			DateStartActive:  equal.dateStartActive,
			DateFinishActive: equal.dateFinishActive,
		})

		message := err.Error()
		parts := strings.Split(message, " = ")
		assert.Equal(t, equal.expected, parts[2])
	}
}
func TestAddPersonalPromoCode_CheckDateFinishPersonalPromoCode(t *testing.T) {
	ctx, st := Suite.New(t)

	startDate := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)

	testTable := []struct {
		idClient         int
		idGroup          int
		namePromoCode    string
		typeDiscountId   int
		valueDiscount    int
		dateStartActive  string
		dateFinishActive string
		expected         string
	}{
		{
			idClient:         1,
			idGroup:          1,
			namePromoCode:    gofakeit.Lexify("?????"),
			typeDiscountId:   1,
			valueDiscount:    gofakeit.Number(1, 99),
			dateStartActive:  gofakeit.DateRange(startDate, endDate).Format("2006-01-02"),
			dateFinishActive: gofakeit.FutureDate().Format("02-Jan-2006"),
			expected:         "incorrect format date finish",
		},
		{
			idClient:         1,
			idGroup:          1,
			namePromoCode:    gofakeit.Lexify("?????"),
			typeDiscountId:   1,
			valueDiscount:    gofakeit.Number(1, 99),
			dateStartActive:  gofakeit.DateRange(startDate, endDate).Format("2006-01-02"),
			dateFinishActive: gofakeit.FutureDate().Format("20060102"),
			expected:         "incorrect format date finish",
		},
		{
			idClient:         1,
			idGroup:          1,
			namePromoCode:    gofakeit.Lexify("?????"),
			typeDiscountId:   1,
			valueDiscount:    gofakeit.Number(1, 99),
			dateStartActive:  gofakeit.DateRange(startDate, endDate).Format("2006-01-02"),
			dateFinishActive: gofakeit.FutureDate().Format("January 02, 2006"),
			expected:         "incorrect format date finish",
		},
		{
			idClient:         1,
			idGroup:          1,
			namePromoCode:    gofakeit.Lexify("?????"),
			typeDiscountId:   1,
			valueDiscount:    gofakeit.Number(1, 99),
			dateStartActive:  gofakeit.DateRange(startDate, endDate).Format("2006-01-02"),
			dateFinishActive: gofakeit.FutureDate().Format("01/02/06"),
			expected:         "incorrect format date finish",
		}, {
			idClient:         1,
			idGroup:          1,
			namePromoCode:    gofakeit.Lexify("?????"),
			typeDiscountId:   1,
			valueDiscount:    gofakeit.Number(1, 99),
			dateStartActive:  gofakeit.DateRange(startDate, endDate).Format("2006-01-02"),
			dateFinishActive: gofakeit.FutureDate().Format("01/02/2006"),
			expected:         "incorrect format date finish",
		},
		{
			idClient:         1,
			idGroup:          1,
			namePromoCode:    gofakeit.Lexify("?????"),
			typeDiscountId:   1,
			valueDiscount:    gofakeit.Number(1, 99),
			dateStartActive:  gofakeit.DateRange(startDate, endDate).Format("2006-01-02"),
			dateFinishActive: gofakeit.FutureDate().Format("010206"),
			expected:         "incorrect format date finish",
		},
		{
			idClient:         1,
			idGroup:          1,
			namePromoCode:    gofakeit.Lexify("?????"),
			typeDiscountId:   1,
			valueDiscount:    gofakeit.Number(1, 99),
			dateStartActive:  gofakeit.DateRange(startDate, endDate).Format("2006-01-02"),
			dateFinishActive: gofakeit.FutureDate().Format("Jan-02-06"),
			expected:         "incorrect format date finish",
		},
		{
			idClient:         1,
			idGroup:          1,
			namePromoCode:    gofakeit.Lexify("?????"),
			typeDiscountId:   1,
			valueDiscount:    gofakeit.Number(1, 99),
			dateStartActive:  gofakeit.DateRange(startDate, endDate).Format("2006-01-02"),
			dateFinishActive: gofakeit.FutureDate().Format("Jan-02-2006"),
			expected:         "incorrect format date finish",
		},
	}
	for _, equal := range testTable {
		_, err := st.LoyaltyServiceClient.AddPersonalPromoCode(ctx, &sl.AddPersonalPromoCodeRequest{
			IdClient:         int32(equal.idClient),
			IdGroup:          int32(equal.idGroup),
			NamePromoCode:    equal.namePromoCode,
			TypeDiscount:     int32(equal.typeDiscountId),
			ValueDiscount:    int32(equal.valueDiscount),
			DateStartActive:  equal.dateStartActive,
			DateFinishActive: equal.dateFinishActive,
		})

		message := err.Error()
		parts := strings.Split(message, " = ")
		assert.Equal(t, equal.expected, parts[2])
	}
}
