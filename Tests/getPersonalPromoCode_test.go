package Tests

import (
	"github.com/Leleria/ServiceLoyalty/Tests/Suite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

func TestGetPersonalPromoCode(t *testing.T) {
	ctx, st := Suite.New(t)

	expected := "User1 постоянные клиенты процентная 89 2023-10-24 2024-05-15"
	namePromoCode := "AqVPu"

	result, err := st.DB.GetPersonalPromoCode(ctx, namePromoCode)

	if err != nil {
		message := err.Error()
		parts := strings.Split(message, ": ")
		assert.Equal(t, expected, parts[1])
	} else {
		require.NotEmpty(t, result)
		assert.Equal(t, result, expected)
	}

}
