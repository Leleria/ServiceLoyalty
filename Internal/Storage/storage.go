package Storage

import "errors"

var (
	ErrPromoCodeExists       = errors.New("promo code already exists")
	ErrSettingUpBudgetExists = errors.New("budget already exists")
	ErrCashBackExists        = errors.New("cashback already exists")
	ErrPromoCodeFound        = errors.New("not found promo code")
	ErrCashBackFound         = errors.New("not found cashback")
	ErrTypeDiscount          = errors.New("percent discount cant be more 100")
	ErrDateActive            = errors.New("date start more than date finish")
)
