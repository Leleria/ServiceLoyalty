package Tests

import (
	"github.com/Leleria/ServiceLoyalty/Tests/Suite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

func TestDeletePromoCode(t *testing.T) {
	ctx, st := Suite.New(t)

	var expected string
	namePromoCode := "qwert"

	result, err := st.DB.DeletePromoCode(ctx, namePromoCode)

	if err != nil {
		message := err.Error()
		parts := strings.Split(message, ": ")
		assert.Equal(t, expected, parts[1])
	} else {
		require.NotEmpty(t, result)
		assert.Equal(t, result, "complete")
	}

}
