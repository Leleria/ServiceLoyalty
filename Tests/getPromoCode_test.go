package Tests

import (
	"github.com/Leleria/ServiceLoyalty/Tests/Suite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

func TestGetPromoCode(t *testing.T) {
	ctx, st := Suite.New(t)

	expected := "процентная 12 2013-02-12 2024-12-12 15"
	namePromoCode := "HbnGh"

	result, err := st.DB.GetPromoCode(ctx, namePromoCode)

	if err != nil {
		message := err.Error()
		parts := strings.Split(message, ": ")
		assert.Equal(t, expected, parts[1])
	} else {
		require.NotEmpty(t, result)
		assert.Equal(t, result, expected)
	}

}
