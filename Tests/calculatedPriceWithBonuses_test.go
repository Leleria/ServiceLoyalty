package Tests

import (
	sl "github.com/Leleria/Contract/GeneratedFilesProtoBufGo"
	"github.com/Leleria/ServiceLoyalty/Tests/Suite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCalculatedPriceWthBonuses(t *testing.T) {

	ctx, st := Suite.New(t)

	idClient := 1
	amountProduct := 500

	respCalcPrice, err := st.LoyaltyServiceClient.CalculatePriceWithBonuses(ctx, &sl.CalculatePriceWithBonusesRequest{
		IdClient:      int32(idClient),
		AmountProduct: float32(amountProduct),
	})
	require.NoError(t, err)
	result := respCalcPrice.GetResult()
	require.NotEmpty(t, result)
	assert.Equal(t, result, "complete, 350, 150")
}
