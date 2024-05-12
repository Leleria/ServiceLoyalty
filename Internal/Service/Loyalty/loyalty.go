package Loyalty

import (
	"context"
	"fmt"
	"github.com/Leleria/ServiceLoyalty/Internal/Lib/Logger/Sl"
	"log/slog"
	"strconv"
)

type PromoCodeChanger interface {
	SavePersonalPromoCode(ctx context.Context, idClient int32, idGroup int32,
		IdPromoCode string, typeDiscount int32, valueDiscount int32, dateStartActive string, dateFinishActive string) (result string, err error)
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

	DeletePersonalPromoCode(ctx context.Context, name string) (result string, err error)
	ChangeClientPersonalPromoCode(ctx context.Context, name string, idClient int32) (result string, err error)
	ChangeGroupPersonalPromoCode(ctx context.Context, name string, idGroup int32) (result string, err error)
	ChangeNamePersonalPromoCode(ctx context.Context, name string, newName string) (result string, err error)
	ChangeTypeDiscountPersonalPromoCode(ctx context.Context, name string, typeDiscount int32) (result string, err error)
	ChangeValueDiscountPersonalPromoCode(ctx context.Context, name string, valueDiscount int32) (result string, err error)
	ChangeDateStartActivePersonalPromoCode(ctx context.Context, name string, dateStartActive string) (result string, err error)
	ChangeDateFinishActivePersonalPromoCode(ctx context.Context, name string, dateFinish string) (result string, err error)
	GetPersonalPromoCode(ctx context.Context, name string) (result string, err error)
	GetAllPersonalPromoCodes(ctx context.Context) (result string, err error)

	GetClient(ctx context.Context, idClient int32) (result string, err error)
	GetAllClients(ctx context.Context) (result string, err error)

	GetOperation(ctx context.Context, idOperation int32) (result string, err error)
	GetAllOperations(ctx context.Context) (result string, err error)

	SaveSettingUpBudget(ctx context.Context, typeCashBack int32, condition string, valueBudget int32) (result string, err error)
	ChangeBudgetCashBack(ctx context.Context, idCashBack int32, budget int32) (result string, err error)
	ChangeTypeCashBack(ctx context.Context, idCashBack int32, typeCashBack int32) (result string, err error)
	ChangeConditionCashBack(ctx context.Context, idCashBack int32, condition string) (result string, err error)
	GetCashBack(ctx context.Context, idCashBack int32) (result string, err error)
	GetAllCashBack(ctx context.Context) (result string, err error)
	DeleteCashBack(ctx context.Context, idCashBack int32) (result string, err error)

	CalculatePriceWithPromoCode(ctx context.Context, idClient int32, namePromoCode string, amountProduct float32) (
		result string, finalAmountProduct float32, amountDiscount float32, err error)
	CalculatePriceWithBonuses(ctx context.Context, idClient int32, amountProduct float32) (
		result string, finalAmountProduct float32, numberBonusesDebited float32, err error)
	DebitingPromoBonuses(ctx context.Context, idClient int32, paymentStatus bool) (result string, err error)
	AccrualBonusesCashback(ctx context.Context, idClient int32, idCaskBack int32) (result string, err error)
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

func (l *Loyalty) GetClient(ctx context.Context, idClient int32) (result string, err error) {
	const op = "Loyalty.GetClient"
	log := l.log.With(
		slog.String("op", op),
		slog.String("id", strconv.Itoa(int(idClient))),
	)
	result, err = l.promoCodeChanger.GetClient(ctx, idClient)
	if err != nil {
		log.Error("failed to get client", Sl.Err(err))
		return "", fmt.Errorf("%s: %w", op, err)
	}
	log.Info("received client " + "\"" + strconv.Itoa(int(idClient)) + "\"")
	return result, nil
}

func (l *Loyalty) GetAllClients(ctx context.Context) (result string, err error) {
	const op = "Loyalty.GetAllClients"
	log := l.log.With(
		slog.String("op", op),
	)
	result, err = l.promoCodeChanger.GetAllClients(ctx)
	if err != nil {
		log.Error("failed to get all clients", Sl.Err(err))
		return "", fmt.Errorf("%s: %w", op, err)
	}
	log.Info("received all clients " + "\"")
	return result, nil
}

func (l *Loyalty) GetOperation(ctx context.Context, idOperation int32) (result string, err error) {
	const op = "Loyalty.GetOperation"
	log := l.log.With(
		slog.String("op", op),
		slog.String("id", strconv.Itoa(int(idOperation))),
	)
	result, err = l.promoCodeChanger.GetOperation(ctx, idOperation)
	if err != nil {
		log.Error("failed to get operation", Sl.Err(err))
		return "", fmt.Errorf("%s: %w", op, err)
	}
	log.Info("received operation " + "\"" + strconv.Itoa(int(idOperation)) + "\"")
	return result, nil
}

func (l *Loyalty) GetAllOperations(ctx context.Context) (result string, err error) {
	const op = "Loyalty.GetAllOperations"
	log := l.log.With(
		slog.String("op", op),
	)
	result, err = l.promoCodeChanger.GetAllOperations(ctx)
	if err != nil {
		log.Error("failed to get all operations", Sl.Err(err))
		return "", fmt.Errorf("%s: %w", op, err)
	}
	log.Info("received all operations " + "\"")
	return result, nil
}

func (l *Loyalty) ChangeClientPersonalPromoCode(ctx context.Context, name string, idClient int32) (result string, err error) {
	const op = "Loyalty.ChangeClientPersonalPromoCode"
	log := l.log.With(
		slog.String("op", op),
		slog.String("name", name),
	)

	result, err = l.promoCodeChanger.ChangeClientPersonalPromoCode(ctx, name, idClient)
	if err != nil {
		log.Error("failed to change id client personal promo code", Sl.Err(err))
		return "", fmt.Errorf("%s: %w", op, err)
	}
	log.Info("changed id client personal promo code ")
	return result, nil
}

func (l *Loyalty) ChangeGroupPersonalPromoCode(ctx context.Context, name string, idGroup int32) (result string, err error) {
	const op = "Loyalty.ChangeGroupPersonalPromoCode"
	log := l.log.With(
		slog.String("op", op),
		slog.String("name", name),
	)

	result, err = l.promoCodeChanger.ChangeGroupPersonalPromoCode(ctx, name, idGroup)
	if err != nil {
		log.Error("failed to change id group personal promo code", Sl.Err(err))
		return "", fmt.Errorf("%s: %w", op, err)
	}
	log.Info("changed id group personal promo code ")
	return result, nil
}

func (l *Loyalty) ChangeNamePersonalPromoCode(ctx context.Context, name string, newName string) (result string, err error) {
	const op = "Loyalty.ChangeNamePersonalPromoCode"
	log := l.log.With(
		slog.String("op", op),
		slog.String("name", name),
	)

	result, err = l.promoCodeChanger.ChangeNamePersonalPromoCode(ctx, name, newName)
	if err != nil {
		log.Error("failed to change name personal promo code", Sl.Err(err))
		return "", fmt.Errorf("%s: %w", op, err)
	}
	log.Info("changed name personal promo code " + "\"" + name + "\"" + " --> " + "\"" + newName + "\"")
	return result, nil
}
func (l *Loyalty) ChangeTypeDiscountPersonalPromoCode(ctx context.Context, name string, typeDiscount int32) (result string, err error) {
	const op = "Loyalty.ChangeTypeDiscountPersonalPromoCode"
	log := l.log.With(
		slog.String("op", op),
		slog.String("name", name),
	)

	result, err = l.promoCodeChanger.ChangeTypeDiscountPersonalPromoCode(ctx, name, typeDiscount)
	if err != nil {
		log.Error("failed to change type discount personal promo code", Sl.Err(err))
		return "", fmt.Errorf("%s: %w", op, err)
	}
	log.Info("changed type discount personal promo code ")
	return result, nil
}
func (l *Loyalty) ChangeValueDiscountPersonalPromoCode(ctx context.Context, name string, valueDiscount int32) (result string, err error) {
	const op = "Loyalty.ChangeValueDiscountPersonalPromoCode"
	log := l.log.With(
		slog.String("op", op),
		slog.String("name", name),
	)

	result, err = l.promoCodeChanger.ChangeValueDiscountPersonalPromoCode(ctx, name, valueDiscount)
	if err != nil {
		log.Error("failed to change value discount personal promo code", Sl.Err(err))
		return "", fmt.Errorf("%s: %w", op, err)
	}
	log.Info("changed value discount personal promo code ")
	return result, nil
}
func (l *Loyalty) ChangeDateStartActivePersonalPromoCode(ctx context.Context, name string, dateStartActive string) (result string, err error) {
	const op = "Loyalty.ChangeDateStartActivePersonalPromoCode"
	log := l.log.With(
		slog.String("op", op),
		slog.String("name", name),
	)

	result, err = l.promoCodeChanger.ChangeDateStartActivePersonalPromoCode(ctx, name, dateStartActive)
	if err != nil {
		log.Error("failed to change date start personal promo code", Sl.Err(err))
		return "", fmt.Errorf("%s: %w", op, err)
	}
	log.Info("changed date start activation personal promo code ")
	return result, nil
}
func (l *Loyalty) ChangeDateFinishActivePersonalPromoCode(ctx context.Context, name string, dateFinishActive string) (result string, err error) {
	const op = "Loyalty.ChangeDateFinishPersonalPromoCode"
	log := l.log.With(
		slog.String("op", op),
		slog.String("name", name),
	)

	result, err = l.promoCodeChanger.ChangeDateFinishActivePersonalPromoCode(ctx, name, dateFinishActive)
	if err != nil {
		log.Error("failed to change date finish personal promo code", Sl.Err(err))
		return "", fmt.Errorf("%s: %w", op, err)
	}
	log.Info("changed date finish activation personal promo code ")
	return result, nil
}

func (l *Loyalty) DeletePersonalPromoCode(ctx context.Context, name string) (result string, err error) {
	const op = "Loyalty.DeletePersonalPromoCode"
	log := l.log.With(
		slog.String("op", op),
		slog.String("name", name),
	)
	log.Info("deleted " + "\"" + name + "\"" + " personal promo code")
	result, err = l.promoCodeChanger.DeletePersonalPromoCode(ctx, name)
	if err != nil {
		log.Error("failed to delete personal promo code", Sl.Err(err))
		return "", fmt.Errorf("%s: %w", op, err)
	}
	return result, nil
}

func (l *Loyalty) AccrualBonusesCashback(ctx context.Context, idClient int32, idCashBack int32) (result string, err error) {
	const op = "Loyalty.AccrualBonusesCashback"
	log := l.log.With(
		slog.String("op", op),
		slog.String("idClient", strconv.Itoa(int(idClient))),
		slog.String("idCashBack", strconv.Itoa(int(idCashBack))),
	)
	result, err = l.promoCodeChanger.AccrualBonusesCashback(ctx, idClient, idCashBack)
	if err != nil {
		log.Error("failed to accrual bonuses", Sl.Err(err))
		return "", fmt.Errorf("%s: %w", op, err)
	}
	return result, nil
}

func (l *Loyalty) DebitingPromoBonuses(ctx context.Context, idClient int32, paymentStatus bool) (result string, err error) {
	const op = "Loyalty.DebitingPromoBonuses"
	log := l.log.With(
		slog.String("op", op),
		slog.String("idClient", strconv.Itoa(int(idClient))),
	)
	log.Info("paid by client " + "\"" + strconv.Itoa(int(idClient)))
	result, err = l.promoCodeChanger.DebitingPromoBonuses(ctx, idClient, paymentStatus)
	if err != nil {
		log.Error("failed to debiting promo code oe bonuses", Sl.Err(err))
		return "", fmt.Errorf("%s: %w", op, err)
	}
	return result, nil
}

func (l *Loyalty) CalculatePriceWithPromoCode(ctx context.Context, idClient int32, namePromoCode string,
	amountProduct float32) (result string, finalAmountProduct float32, amountDiscount float32, err error) {
	const op = "Loyalty.CalculatePriceWithPromoCode"
	log := l.log.With(
		slog.String("op", op),
		slog.String("namePromoCode", namePromoCode),
		slog.String("idClient", strconv.Itoa(int(idClient))),
	)
	log.Info("calculated " + "\"" + strconv.Itoa(int(amountProduct)) + "\"" + " product amount")
	result, finalAmountProduct, amountDiscount, err = l.promoCodeChanger.CalculatePriceWithPromoCode(ctx, idClient, namePromoCode, amountProduct)
	if err != nil {
		log.Error("failed to calculated price", Sl.Err(err))
		return "", 0, 0, fmt.Errorf("%s: %w", op, err)
	}
	return result, finalAmountProduct, amountDiscount, nil
}

func (l *Loyalty) CalculatePriceWithBonuses(ctx context.Context, idClient int32,
	amountProduct float32) (result string, finalAmountProduct float32, numberBonusesDebited float32, err error) {
	const op = "Loyalty.CalculatePriceWithBonuses"
	log := l.log.With(
		slog.String("op", op),
		slog.String("idClient", strconv.Itoa(int(idClient))),
	)
	log.Info("calculated " + "\"" + strconv.Itoa(int(amountProduct)) + "\"" + " product amount")
	result, finalAmountProduct, numberBonusesDebited, err = l.promoCodeChanger.CalculatePriceWithBonuses(ctx, idClient, amountProduct)
	if err != nil {
		log.Error("failed to calculated price", Sl.Err(err))
		return "", 0, 0, fmt.Errorf("%s: %w", op, err)
	}
	return result, finalAmountProduct, numberBonusesDebited, nil
}

func (l *Loyalty) DeleteCashBack(ctx context.Context, idCashBack int32) (result string, err error) {
	const op = "Loyalty.DeleteCashBack"
	log := l.log.With(
		slog.String("op", op),
		slog.String("id", strconv.Itoa(int(idCashBack))),
	)
	log.Info("deleted " + "\"" + strconv.Itoa(int(idCashBack)) + "\"" + " cashback")
	result, err = l.promoCodeChanger.DeleteCashBack(ctx, idCashBack)
	if err != nil {
		log.Error("failed to delete cashback", Sl.Err(err))
		return "", fmt.Errorf("%s: %w", op, err)
	}
	return result, nil
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

func (l *Loyalty) GetPersonalPromoCode(ctx context.Context, name string) (result string, err error) {
	const op = "Loyalty.GetPersonalPromoCode"
	log := l.log.With(
		slog.String("op", op),
		slog.String("name", name),
	)
	result, err = l.promoCodeChanger.GetPersonalPromoCode(ctx, name)
	if err != nil {
		log.Error("failed to get personal promo code", Sl.Err(err))
		return "", fmt.Errorf("%s: %w", op, err)
	}
	log.Info("received personal promo code " + "\"" + name + "\"")
	return result, nil
}

func (l *Loyalty) GetAllPersonalPromoCodes(ctx context.Context) (result string, err error) {
	const op = "Loyalty.GetAllPersonalPromoCodes"
	log := l.log.With(
		slog.String("op", op),
	)
	result, err = l.promoCodeChanger.GetAllPersonalPromoCodes(ctx)
	if err != nil {
		log.Error("failed to get all personal promo codes", Sl.Err(err))
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
		log.Error("failed to change date start promo code", Sl.Err(err))
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
		log.Error("failed to change date finish promo code", Sl.Err(err))
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
		log.Error("failed to change max count uses promo code", Sl.Err(err))
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
		log.Error("failed to change value discount promo code", Sl.Err(err))
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
	IdPromoCode string, typeDiscount int32, valueDiscount int32, dateStartActive string, dateFinishActive string) (string, error) {
	const op = "Loyalty.AddPersonalPromoCode"
	log := l.log.With(
		slog.String("op", op),
	)
	log.Info("added personal promo code")

	result, err := l.promoCodeChanger.SavePersonalPromoCode(ctx, idClient, idGroup, IdPromoCode, typeDiscount,
		valueDiscount, dateStartActive, dateFinishActive)
	if err != nil {
		log.Error("failed to save personal promo code", Sl.Err(err))
		return "", fmt.Errorf("%s: %w", op, err)
	}
	return result, nil
}

func (l *Loyalty) SettingUpBudget(ctx context.Context, typeCashBack int32, condition string, valueBudget int32) (string, error) {
	const op = "Loyalty.SettingUpBudget"
	log := l.log.With(
		slog.String("op", op),
	)
	log.Info("setting up a budget")

	result, err := l.promoCodeChanger.SaveSettingUpBudget(ctx, typeCashBack, condition, valueBudget)
	if err != nil {
		log.Error("failed to setting up a budget", Sl.Err(err))
		return "", fmt.Errorf("%s: %w", op, err)
	}
	return result, nil
}

func (l *Loyalty) GetCashBack(ctx context.Context, idCashBack int32) (result string, err error) {
	const op = "Loyalty.GetCashBack"
	log := l.log.With(
		slog.String("op", op),
		slog.String("id", strconv.Itoa(int(idCashBack))),
	)
	result, err = l.promoCodeChanger.GetCashBack(ctx, idCashBack)
	if err != nil {
		log.Error("failed to get cashback", Sl.Err(err))
		return "", fmt.Errorf("%s: %w", op, err)
	}
	log.Info("received cashback " + "\"" + strconv.Itoa(int(idCashBack)) + "\"")
	return result, nil
}

func (l *Loyalty) GetAllCashBack(ctx context.Context) (result string, err error) {
	const op = "Loyalty.GetAllCashBack"
	log := l.log.With(
		slog.String("op", op),
	)
	result, err = l.promoCodeChanger.GetAllCashBack(ctx)
	if err != nil {
		log.Error("failed to get all cashbacks", Sl.Err(err))
		return "", fmt.Errorf("%s: %w", op, err)
	}
	log.Info("received all cashbacks " + "\"")
	return result, nil
}

func (l *Loyalty) ChangeBudgetCashBack(ctx context.Context, idCashBack int32, budget int32) (result string, err error) {
	const op = "Loyalty.ChangeBudgetCashBack"
	log := l.log.With(
		slog.String("op", op),
		slog.String("id", strconv.Itoa(int(idCashBack))),
	)

	result, err = l.promoCodeChanger.ChangeBudgetCashBack(ctx, idCashBack, budget)
	if err != nil {
		log.Error("failed to change budget cashback", Sl.Err(err))
		return "", fmt.Errorf("%s: %w", op, err)
	}
	log.Info("changed budget cashback ")
	return result, nil
}

func (l *Loyalty) ChangeTypeCashBack(ctx context.Context, idCashBack int32, typeCashBack int32) (result string, err error) {
	const op = "Loyalty.ChangeTypeCashBack"
	log := l.log.With(
		slog.String("op", op),
		slog.String("id", strconv.Itoa(int(idCashBack))),
	)

	result, err = l.promoCodeChanger.ChangeTypeCashBack(ctx, idCashBack, typeCashBack)
	if err != nil {
		log.Error("failed to change type cashback", Sl.Err(err))
		return "", fmt.Errorf("%s: %w", op, err)
	}
	log.Info("changed type cashback ")
	return result, nil
}

func (l *Loyalty) ChangeConditionCashBack(ctx context.Context, idCashBack int32, condition string) (result string, err error) {
	const op = "Loyalty.ChangeConditionCashBack"
	log := l.log.With(
		slog.String("op", op),
		slog.String("id", strconv.Itoa(int(idCashBack))),
	)

	result, err = l.promoCodeChanger.ChangeConditionCashBack(ctx, idCashBack, condition)
	if err != nil {
		log.Error("failed to change condition cashback", Sl.Err(err))
		return "", fmt.Errorf("%s: %w", op, err)
	}
	log.Info("changed condition cashback ")
	return result, nil
}
