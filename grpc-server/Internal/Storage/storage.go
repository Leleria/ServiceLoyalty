package Storage

import "errors"

var (
	ErrPromoCodeExists         = errors.New("promo code already exists")
	ErrPersonalPromoCodeExists = errors.New("personal promo code already exists")
	ErrClientExists            = errors.New("client already exists")
	ErrSettingUpBudgetExists   = errors.New("budget already exists")
	ErrCashBackExists          = errors.New("cashback already exists")
	ErrPromoCodeFound          = errors.New("not found promo code")
	ErrOperationFound          = errors.New("not found operation")
	ErrClientFound             = errors.New("not found client")
	ErrGroupFound              = errors.New("not found group")
	ErrCashBackFound           = errors.New("not found cashback")
	ErrTypeDiscount            = errors.New("percent discount cant be more 100")
	ErrDateActive              = errors.New("date start more than date finish")
)
