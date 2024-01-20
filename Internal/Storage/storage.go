package Storage

import "errors"

var (
	ErrPromoCodeExists = errors.New("promo code already exists")
	ErrPromoCodeFound  = errors.New("not found promo code")
)
