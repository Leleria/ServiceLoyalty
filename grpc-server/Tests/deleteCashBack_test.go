package Tests

import (
	"github.com/Leleria/ServiceLoyalty/Tests/Suite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

func TestDeleteCashBack(t *testing.T) {
	ctx, st := Suite.New(t)

	expected := "complete"
	idCashBack := 6

	result, err := st.DB.DeleteCashBack(ctx, int32(idCashBack))

	if err != nil {
		message := err.Error()
		parts := strings.Split(message, ": ")
		assert.Equal(t, expected, parts[1])
	} else {
		require.NotEmpty(t, result)
		assert.Equal(t, result, "complete")
	}

}
