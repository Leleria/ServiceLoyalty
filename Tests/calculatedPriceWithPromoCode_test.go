package Tests

import (
	sl "github.com/Leleria/Contract/GeneratedFilesProtoBufGo"
	"github.com/Leleria/ServiceLoyalty/Tests/Suite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCalculatedPriceWthPromoCode(t *testing.T) {

	ctx, st := Suite.New(t)

	idClient := 1
	namePromoCode := "PPPPP"
	amountProduct := 500

	respCalcPrice, err := st.LoyaltyServiceClient.CalculatePriceWithPromoCode(ctx, &sl.CalculatePriceWithPromoCodeRequest{
		IdClient:      int32(idClient),
		PromoCode:     namePromoCode,
		AmountProduct: float32(amountProduct),
	})
	require.NoError(t, err)
	result := respCalcPrice.GetResult()
	require.NotEmpty(t, result)
	assert.Equal(t, result, "complete, 440, 60")
}
