package Tests

import (
	"github.com/Leleria/ServiceLoyalty/Tests/Suite"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

func TestAccrualBonusesCashBack(t *testing.T) {
	ctx, st := Suite.New(t)

	var expected string
	idClient := gofakeit.IntRange(1, 5)
	idCashBack := gofakeit.IntRange(1, 5)

	result, err := st.DB.AccrualBonusesCashback(ctx, int32(idClient), int32(idCashBack))

	if err != nil {
		message := err.Error()
		parts := strings.Split(message, ": ")
		assert.Equal(t, expected, parts[1])
	} else {
		require.NotEmpty(t, result)
		assert.Equal(t, result, "complete")
	}

}
