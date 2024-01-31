package Loyalty

import (
	"context"
	"fmt"
	"github.com/Leleria/ServiceLoyalty/Internal/Lib/Logger/Sl"
	"log/slog"
)

type PromoCodeChanger interface {
	SavePersonalPromoCode(ctx context.Context, idClient int32, idGroup int32,
		IdPromoCode string) (result string, err error)
	SavePromoCode(ctx context.Context,
		name string, typeDiscount int32,
		valueDiscount int32, dateStartActive string,
		dateFinishActive string, maxCountUses int32) (
		result string, err error)
	GetPromoCode(ctx context.Context, name string) (string, error)
	GetAllPromoCodes(ctx context.Context) (string, error)
	DeletePromoCode(ctx context.Context, name string) (result string, err error)
	ChangeNamePromoCode(ctx context.Context, name string, newName string) (result string, err error)
	ChangeTypeDiscountPromoCode(ctx context.Context, name string, typeDiscount int32) (result string, err error)
	ChangeValueDiscountPromoCode(ctx context.Context, name string, valueDiscount int32) (result string, err error)
	ChangeDateStartActivePromoCode(ctx context.Context, name string, dateStartActive string) (result string, err error)
	ChangeDateFinishActivePromoCode(ctx context.Context, name string, dateFinish string) (result string, err error)
	ChangeMaxCountUsesPromoCode(ctx context.Context, name string, maxCountUses int32) (result string, err error)
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

func (l *Loyalty) GetPromoCode(ctx context.Context, name string) (result string, err error) {
	const op = "Loyalty.GetPromoCode"
	log := l.log.With(
		slog.String("op", op),
		slog.String("name", name),
	)
	result, err = l.promoCodeChanger.GetPromoCode(ctx, name)
	if err != nil {
		log.Error("failed to get promo code", Sl.Err(err))
		return "", fmt.Errorf("%s: %w", op, err)
	}
	log.Info("received promo code " + "\"" + name + "\"")
	return result, nil
}

func (l *Loyalty) GetAllPromoCodes(ctx context.Context) (result string, err error) {
	const op = "Loyalty.GetAllPromoCodes"
	log := l.log.With(
		slog.String("op", op),
	)
	result, err = l.promoCodeChanger.GetAllPromoCodes(ctx)
	if err != nil {
		log.Error("failed to get all promo codes", Sl.Err(err))
		return "", fmt.Errorf("%s: %w", op, err)
	}
	log.Info("received all promo codes " + "\"")
	return result, nil
}

func (l *Loyalty) ChangeNamePromoCode(ctx context.Context, name string, newName string) (result string, err error) {
	const op = "Loyalty.ChangeNamePromoCode"
	log := l.log.With(
		slog.String("op", op),
		slog.String("name", name),
	)

	result, err = l.promoCodeChanger.ChangeNamePromoCode(ctx, name, newName)
	if err != nil {
		log.Error("failed to change name promo code", Sl.Err(err))
		return "", fmt.Errorf("%s: %w", op, err)
	}
	log.Info("changed name promo code " + "\"" + name + "\"" + " --> " + "\"" + newName + "\"")
	return result, nil
}
func (l *Loyalty) ChangeTypeDiscountPromoCode(ctx context.Context, name string, typeDiscount int32) (result string, err error) {
	const op = "Loyalty.ChangeTypeDiscountPromoCode"
	log := l.log.With(
		slog.String("op", op),
		slog.String("name", name),
	)

	result, err = l.promoCodeChanger.ChangeTypeDiscountPromoCode(ctx, name, typeDiscount)
	if err != nil {
		log.Error("failed to change type discount promo code", Sl.Err(err))
		return "", fmt.Errorf("%s: %w", op, err)
	}
	log.Info("changed type discount promo code ")
	return result, nil
}
func (l *Loyalty) ChangeDateStartActivePromoCode(ctx context.Context, name string, dateStartActive string) (result string, err error) {
	const op = "Loyalty.ChangeDateStartActivePromoCode"
	log := l.log.With(
		slog.String("op", op),
		slog.String("name", name),
	)

	result, err = l.promoCodeChanger.ChangeDateStartActivePromoCode(ctx, name, dateStartActive)
	if err != nil {
		log.Error("failed to delete promo code", Sl.Err(err))
		return "", fmt.Errorf("%s: %w", op, err)
	}
	log.Info("changed date start activation promo code ")
	return result, nil
}
func (l *Loyalty) ChangeDateFinishActivePromoCode(ctx context.Context, name string, dateFinishActive string) (result string, err error) {
	const op = "Loyalty.ChangeDateFinishPromoCode"
	log := l.log.With(
		slog.String("op", op),
		slog.String("name", name),
	)

	result, err = l.promoCodeChanger.ChangeDateFinishActivePromoCode(ctx, name, dateFinishActive)
	if err != nil {
		log.Error("failed to delete promo code", Sl.Err(err))
		return "", fmt.Errorf("%s: %w", op, err)
	}
	log.Info("changed date finish activation promo code ")
	return result, nil
}
func (l *Loyalty) ChangeMaxCountUsesPromoCode(ctx context.Context, name string, maxCountUses int32) (result string, err error) {
	const op = "Loyalty.ChangeMaxCountUsesPromoCode"
	log := l.log.With(
		slog.String("op", op),
		slog.String("name", name),
	)

	result, err = l.promoCodeChanger.ChangeMaxCountUsesPromoCode(ctx, name, maxCountUses)
	if err != nil {
		log.Error("failed to delete promo code", Sl.Err(err))
		return "", fmt.Errorf("%s: %w", op, err)
	}
	log.Info("changed max count uses promo code ")
	return result, nil
}
func (l *Loyalty) ChangeValueDiscountPromoCode(ctx context.Context, name string, valueDiscount int32) (result string, err error) {
	const op = "Loyalty.ChangeValueDiscountPromoCode"
	log := l.log.With(
		slog.String("op", op),
		slog.String("name", name),
	)

	result, err = l.promoCodeChanger.ChangeValueDiscountPromoCode(ctx, name, valueDiscount)
	if err != nil {
		log.Error("failed to delete promo code", Sl.Err(err))
		return "", fmt.Errorf("%s: %w", op, err)
	}
	log.Info("changed value discount promo code ")
	return result, nil
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

func (l *Loyalty) AddPersonalPromoCode(ctx context.Context, idClient int32, idGroup int32,
	IdPromoCode string) (string, error) {
	const op = "Loyalty.AddPersonalPromoCode"
	log := l.log.With(
		slog.String("op", op),
	)
	log.Info("added personal promo code")

	result, err := l.promoCodeChanger.SavePersonalPromoCode(ctx, idClient, idGroup, IdPromoCode)
	if err != nil {
		log.Error("failed to save personal promo code", Sl.Err(err))
		return "", fmt.Errorf("%s: %w", op, err)
	}
	return result, nil
}
