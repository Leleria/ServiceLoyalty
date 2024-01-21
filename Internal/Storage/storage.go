package Storage

import "errors"

var (
	ErrPromoCodeExists = errors.New("promo code already exists")
	ErrPromoCodeFound  = errors.New("not found promo code")
	ErrTypeDiscount    = errors.New("percent discount cant be more 100")
	ErrDateActive      = errors.New("date start more than date finish")
)
