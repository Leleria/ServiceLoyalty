package Loyalty

import (
	"context"
	"fmt"
	"github.com/Leleria/ServiceLoyalty/Internal/Lib/Logger/Sl"
	"log/slog"
)

type PromoCodeChanger interface {
	SavePromoCode(ctx context.Context,
		name string, typeDiscount int32,
		valueDiscount int32, dateStartActive string,
		dateFinishActive string, maxCountUses int32) (
		result string, err error)
	//PromoCode(ctx context.Context, name string) (Models.PromoCode, error)
	DeletePromoCode(ctx context.Context, name string) (result string, err error)
}

type Loyalty struct {
	log              *slog.Logger
	promoCodeChanger PromoCodeChanger
}

func New(log *slog.Logger,
	promoCodeChanger PromoCodeChanger) *Loyalty {
	return &Loyalty{log: log,
		promoCodeChanger: promoCodeChanger,
	}
}
func (l *Loyalty) DeletePromoCode(ctx context.Context, name string) (result string, err error) {
	const op = "Loyalty.DeletePromoCode"
	log := l.log.With(
		slog.String("op", op),
		slog.String("name", name),
	)
	log.Info("deleted " + "\"" + name + "\"" + " promo code")
	result, err = l.promoCodeChanger.DeletePromoCode(ctx, name)
	if err != nil {
		log.Error("failed to delete promo code", Sl.Err(err))
		return "", fmt.Errorf("%s: %w", op, err)
	}
	return result, nil
}
func (l *Loyalty) AddNewPromoCode(ctx context.Context, name string, typeDiscount int32,
	valueDiscount int32, dateStartActive string,
	dateFinishActive string, maxCountUses int32) (string, error) {
	const op = "Loyalty.AddNewPromoCode"
	log := l.log.With(
		slog.String("op", op),
		slog.String("name", name),
	)
	log.Info("added new promo code")

	result, err := l.promoCodeChanger.SavePromoCode(ctx, name, typeDiscount, valueDiscount,
		dateStartActive, dateFinishActive, maxCountUses)
	if err != nil {
		log.Error("failed to save promo code", Sl.Err(err))
		return "", fmt.Errorf("%s: %w", op, err)
	}
	return result, nil
}
