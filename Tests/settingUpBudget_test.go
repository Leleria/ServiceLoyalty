package Tests

import (
	sl "github.com/Leleria/Contract/GeneratedFilesProtoBufGo"
	"github.com/Leleria/ServiceLoyalty/Tests/Suite"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

func TestSettingUpBudget_HappyPath(t *testing.T) {
	ctx, st := Suite.New(t)

	budget := gofakeit.Number(1, 500)
	typeCashBackId := gofakeit.Number(1, 2)
	valueCondition := "в конце месяца"

	respCashback, err := st.LoyaltyServiceClient.SettingUpBudget(ctx, &sl.SettingUpBudgetRequest{
		ValueBudget:  int32(budget),
		TypeCashBack: int32(typeCashBackId),
		Condition:    valueCondition,
	})
	require.NoError(t, err)
	result := respCashback.GetResult()
	require.NotEmpty(t, result)
	assert.Equal(t, result, "complete")
}
func TestSettingUpBudget_CheckBudgetCashBack(t *testing.T) {
	ctx, st := Suite.New(t)

	expected := "incorrect value budget"
	budget := gofakeit.Number(-10, 0)
	typeCashBackId := gofakeit.Number(1, 2)
	valueCondition := "в конце месяца"

	_, err := st.LoyaltyServiceClient.SettingUpBudget(ctx, &sl.SettingUpBudgetRequest{
		ValueBudget:  int32(budget),
		TypeCashBack: int32(typeCashBackId),
		Condition:    valueCondition,
	})
	message := err.Error()
	parts := strings.Split(message, " = ")
	assert.Equal(t, expected, parts[2])
}
func TestSettingUpBudget_CheckTypeCashBack(t *testing.T) {
	ctx, st := Suite.New(t)

	expected := "incorrect type cashback"
	budget := gofakeit.Number(1, 500)
	typeCashBackId := gofakeit.Number(-10, 0)
	valueCondition := "в конце месяца"

	_, err := st.LoyaltyServiceClient.SettingUpBudget(ctx, &sl.SettingUpBudgetRequest{
		ValueBudget:  int32(budget),
		TypeCashBack: int32(typeCashBackId),
		Condition:    valueCondition,
	})
	message := err.Error()
	parts := strings.Split(message, " = ")
	assert.Equal(t, expected, parts[2])
}
