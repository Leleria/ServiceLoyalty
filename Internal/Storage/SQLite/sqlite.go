package SQLite

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/Leleria/ServiceLoyalty/Internal/Domain/Models"
	st "github.com/Leleria/ServiceLoyalty/Internal/Storage"
	"github.com/mattn/go-sqlite3"
	"strconv"
)

type Storage struct {
	db *sql.DB
}

// Конструктор Storage
func New(storagePath string) (*Storage, error) {
	const op = "Storage.SQLite.New"

	// Указываем путь до файла БД
	db, err := sql.Open("sqlite3", storagePath)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &Storage{db: db}, nil
}
func (s *Storage) ChangeNamePromoCode(ctx context.Context, name string, newName string) (string, error) {
	const op = "Storage.SQLite.ChangeNamePromoCode"

	err := s.CheckContainPromoCode(ctx, name)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	stmt, err := s.db.Prepare("UPDATE PromoCodes SET Name = ? WHERE Name = ?")
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	// Выполняем запрос, передав параметры
	_, err = stmt.ExecContext(ctx, newName, name)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	return "complete", nil
}
func (s *Storage) ChangeTypeDiscountPromoCode(ctx context.Context, name string, typeDiscount int32) (string, error) {
	const op = "Storage.SQLite.ChangeTypeDiscountPromoCode"

	err := s.CheckContainPromoCode(ctx, name)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	statement, err := s.db.Prepare(`SELECT ValueDiscount FROM PromoCodes WHERE Name = ?`)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	res := statement.QueryRow(name)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}
	var valueDiscount int
	err = res.Scan(&valueDiscount)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	if valueDiscount > 100 && typeDiscount == 1 {
		return "", fmt.Errorf("%s: %w", op, st.ErrTypeDiscount)
	}

	stmt, err := s.db.Prepare("UPDATE PromoCodes SET TypeDiscountFK = ? WHERE Name = ?")
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	// Выполняем запрос, передав параметры
	_, err = stmt.ExecContext(ctx, typeDiscount, name)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	return "complete", nil
}
func (s *Storage) ChangeValueDiscountPromoCode(ctx context.Context, name string, valueDiscount int32) (string, error) {
	const op = "Storage.SQLite.ChangeValueDiscountPromoCode"

	err := s.CheckContainPromoCode(ctx, name)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	statement, err := s.db.Prepare(`SELECT TypeDiscountFK FROM PromoCodes WHERE Name = ?`)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	res := statement.QueryRow(name)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}
	var typeDiscount int
	err = res.Scan(&typeDiscount)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}
	if typeDiscount == 1 && valueDiscount > 100 {
		return "", fmt.Errorf("%s: %w", op, st.ErrTypeDiscount)
	}

	stmt, err := s.db.Prepare("UPDATE PromoCodes SET ValueDiscount = ? WHERE Name = ?")
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	// Выполняем запрос, передав параметры
	_, err = stmt.ExecContext(ctx, valueDiscount, name)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	return "complete", nil
}
func (s *Storage) ChangeDateStartActivePromoCode(ctx context.Context, name string, dateStartActive string) (string, error) {
	const op = "Storage.SQLite.ChangeDateStartActivePromoCode"

	err := s.CheckContainPromoCode(ctx, name)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	statement, err := s.db.Prepare(`SELECT DateFinishActive FROM PromoCodes WHERE Name = ?`)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	res := statement.QueryRow(name)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}
	var dateFinishActive string
	err = res.Scan(&dateFinishActive)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}
	if dateStartActive > dateFinishActive {
		return "", fmt.Errorf("%s: %w", op, st.ErrDateActive)
	}

	stmt, err := s.db.Prepare("UPDATE PromoCodes SET DateStartActive = ? WHERE Name = ?")
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	// Выполняем запрос, передав параметры
	_, err = stmt.ExecContext(ctx, dateStartActive, name)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	return "complete", nil
}
func (s *Storage) ChangeDateFinishActivePromoCode(ctx context.Context, name string, dateFinishActive string) (string, error) {
	const op = "Storage.SQLite.ChangeDateFinishActivePromoCode"

	err := s.CheckContainPromoCode(ctx, name)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	statement, err := s.db.Prepare(`SELECT DateStartActive FROM PromoCodes WHERE Name = ?`)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	res := statement.QueryRow(name)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}
	var dateStartActive string
	err = res.Scan(&dateStartActive)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}
	if dateFinishActive < dateStartActive {
		return "", fmt.Errorf("%s: %w", op, st.ErrDateActive)
	}

	stmt, err := s.db.Prepare("UPDATE PromoCodes SET DateFinishActive = ? WHERE Name = ?")
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	// Выполняем запрос, передав параметры
	_, err = stmt.ExecContext(ctx, dateFinishActive, name)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	return "complete", nil
}
func (s *Storage) ChangeMaxCountUsesPromoCode(ctx context.Context, name string, maxCountUses int32) (string, error) {
	const op = "Storage.SQLite.ChangeMaxCountUsesPromoCode"

	err := s.CheckContainPromoCode(ctx, name)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	stmt, err := s.db.Prepare("UPDATE PromoCodes SET MaxCountUses = ? WHERE Name = ?")
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	// Выполняем запрос, передав параметры
	_, err = stmt.ExecContext(ctx, maxCountUses, name)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	return "complete", nil
}

func (s *Storage) SavePromoCode(ctx context.Context, name string, typeDiscount int32,
	valueDiscount int32, dateStartActive string,
	dateFinishActive string, maxCountUses int32) (string, error) {
	const op = "Storage.SQLite.SavePromoCode"

	stmt, err := s.db.Prepare("INSERT INTO PromoCodes(Name, TypeDiscountFK, " +
		"ValueDiscount, DateStartActive, DateFinishActive, MaxCountUses) VALUES(?, ?, ?, ?, ?, ?)")
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	_, err = stmt.ExecContext(ctx, name, typeDiscount, valueDiscount, dateStartActive, dateFinishActive, maxCountUses)
	if err != nil {
		var sqliteErr sqlite3.Error

		if errors.As(err, &sqliteErr) && sqliteErr.ExtendedCode == sqlite3.ErrConstraintUnique {
			return "", fmt.Errorf("%s: %w", op, st.ErrPromoCodeExists)
		}

		return "", fmt.Errorf("%s: %w", op, err)
	}

	return "complete", nil
}

func (s *Storage) SavePersonalPromoCode(ctx context.Context, idClient int32, idGroup int32,
	namePromoCode string) (string, error) {
	const op = "Storage.SQLite.SavePersonalPromoCode"

	err := s.CheckContainClient(ctx, idClient)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	err = s.CheckContainGroup(ctx, idGroup)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	err = s.CheckContainPromoCode(ctx, namePromoCode)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	stmt, err := s.db.Prepare("SELECT Id FROM PromoCodes WHERE Name = ?")
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	row := stmt.QueryRowContext(ctx, namePromoCode)

	var promoCode Models.PromoCode
	err = row.Scan(&promoCode.Id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", fmt.Errorf("%s: %w", op, st.ErrPromoCodeFound)
		}

		return "", fmt.Errorf("%s: %w", op, err)
	}

	result := strconv.Itoa(int(promoCode.Id))

	stmt, err = s.db.Prepare("INSERT INTO PersonalPromoCodes(ClientFK, GroupFK, PromoCodeFK) VALUES(?, ?, ?)")
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	// Выполняем запрос, передав параметры
	_, err = stmt.ExecContext(ctx, idClient, idGroup, result)
	if err != nil {
		var sqliteErr sqlite3.Error

		if errors.As(err, &sqliteErr) && sqliteErr.ExtendedCode == sqlite3.ErrConstraintUnique {
			return "", fmt.Errorf("%s: %w", op, st.ErrPromoCodeExists)
		}

		return "", fmt.Errorf("%s: %w", op, err)
	}

	return "complete", nil
}

func (s *Storage) SaveSettingUpBudget(ctx context.Context, typeCashBack int32, condition string, valueBudget int32) (string, error) {
	const op = "Storage.SQLite.SaveSettingUpBudget"

	err := s.CheckContainCashBackType(ctx, typeCashBack)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	statement, err := s.db.Prepare("SELECT Id FROM CashBack WHERE Budget = ? AND TypeCashBackFK = ? AND ValueCondition = ?")
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	res := statement.QueryRowContext(ctx, valueBudget, typeCashBack, condition)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}
	var dataFromDB int32
	err = res.Scan(&dataFromDB)
	if err != nil {
		stmt, err := s.db.Prepare("INSERT INTO CashBack(TypeCashBackFK, ValueCondition, Budget) VALUES(?, ?, ?)")
		if err != nil {
			return "", fmt.Errorf("%s: %w", op, err)
		}

		// Выполняем запрос, передав параметры
		_, err = stmt.ExecContext(ctx, typeCashBack, condition, valueBudget)
		if err != nil {
			var sqliteErr sqlite3.Error

			if errors.As(err, &sqliteErr) && sqliteErr.ExtendedCode == sqlite3.ErrConstraintUnique {
				return "", fmt.Errorf("%s: %w", op, st.ErrPromoCodeExists)
			}

			return "", fmt.Errorf("%s: %w", op, err)
		}
		return "complete", nil
	}

	return "", fmt.Errorf("%s: %w", op, st.ErrCashBackExists)
}

func (s *Storage) DeletePromoCode(ctx context.Context, name string) (string, error) {
	const op = "Storage.SQLite.DeletePromoCode"

	err := s.CheckContainPromoCode(ctx, name)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	stmt, err := s.db.Prepare("DELETE FROM PromoCodes WHERE Name = ?")
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	// Выполняем запрос, передав параметры
	_, err = stmt.ExecContext(ctx, name)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	return "complete", nil
}

func (s *Storage) GetPromoCode(ctx context.Context, name string) (string, error) {
	const op = "Storage.Sqlite.GetPromoCode"

	err := s.CheckContainPromoCode(ctx, name)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}
	stmt, err := s.db.Prepare("select TypesOfDiscounts.NameType, ValueDiscount, DateStartActive, DateFinishActive," +
		" MaxCountUses FROM PromoCodes INNER JOIN TypesOfDiscounts ON PromoCodes.TypeDiscountFK = TypesOfDiscounts.Id  WHERE PromoCodes.Name = ?")
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	row := stmt.QueryRowContext(ctx, name)

	var promoCode Models.PromoCode
	var typeDiscountPromoCode Models.TypeDiscount
	err = row.Scan(&typeDiscountPromoCode.NameType, &promoCode.ValueDiscount, &promoCode.DateStartActive,
		&promoCode.DateFinishActive, &promoCode.MaxCountUses)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", fmt.Errorf("%s: %w", op, st.ErrPromoCodeFound)
		}

		return "", fmt.Errorf("%s: %w", op, err)
	}

	result := typeDiscountPromoCode.NameType + " " + strconv.Itoa(int(promoCode.ValueDiscount)) + " " + promoCode.DateStartActive +
		" " + promoCode.DateFinishActive + " " + strconv.Itoa(int(promoCode.MaxCountUses))
	return result, nil
}

func (s *Storage) GetAllPromoCodes(ctx context.Context) (string, error) {
	const op = "Storage.Sqlite.GetAllPromoCodes"

	stmt, err := s.db.Prepare("select Name, TypesOfDiscounts.NameType, ValueDiscount, DateStartActive, DateFinishActive," +
		" MaxCountUses FROM PromoCodes INNER JOIN TypesOfDiscounts ON PromoCodes.TypeDiscountFK = TypesOfDiscounts.Id ")
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	row, err := stmt.QueryContext(ctx)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}
	var result string
	for row.Next() {
		var promoCode Models.PromoCode
		var typeDiscountPromoCode Models.TypeDiscount
		err := row.Scan(&promoCode.Name, &typeDiscountPromoCode.NameType, &promoCode.ValueDiscount, &promoCode.DateStartActive,
			&promoCode.DateFinishActive, &promoCode.MaxCountUses)

		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return "", fmt.Errorf("%s: %w", op, st.ErrPromoCodeFound)
			}
			return "", fmt.Errorf("%s: %w", op, err)
		}

		result = result + promoCode.Name + " " + typeDiscountPromoCode.NameType + " " + strconv.Itoa(int(promoCode.ValueDiscount)) + " " + promoCode.DateStartActive +
			" " + promoCode.DateFinishActive + " " + strconv.Itoa(int(promoCode.MaxCountUses)) + ", "
	}

	return result, nil
}

func (s *Storage) GetCashBack(ctx context.Context, idCashBack int32) (string, error) {
	const op = "Storage.Sqlite.GetCashBack"

	err := s.CheckContainCashBack(ctx, idCashBack)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}
	stmt, err := s.db.Prepare("SELECT Budget, NameType, ValueCondition FROM CashBack INNER JOIN CashBackTypes ON  CashBack.TypeCashBackFK = CashBackTypes.Id WHERE CashBack.Id = ?")
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	row := stmt.QueryRowContext(ctx, idCashBack)

	var cashback Models.CashBack
	var cashbackType Models.TypeCashBack
	err = row.Scan(&cashback.Budget, &cashbackType.NameType, &cashback.ValueCondition)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", fmt.Errorf("%s: %w", op, st.ErrCashBackFound)
		}

		return "", fmt.Errorf("%s: %w", op, err)
	}

	result := strconv.Itoa(int(cashback.Budget)) + " " + cashbackType.NameType + " " + cashback.ValueCondition
	return result, nil
}

func (s *Storage) GetAllCashBack(ctx context.Context) (string, error) {
	const op = "Storage.Sqlite.GetAllCashBack"

	stmt, err := s.db.Prepare("SELECT Budget, CashBackTypes.NameType, ValueCondition FROM CashBack INNER JOIN CashBackTypes ON CashBack.Id = CashBackTypes.Id ")
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	row, err := stmt.QueryContext(ctx)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}
	var result string
	for row.Next() {
		var cashback Models.CashBack
		var cashbackType Models.TypeCashBack
		err := row.Scan(&cashback.Budget, &cashbackType.NameType, &cashback.ValueCondition)

		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return "", fmt.Errorf("%s: %w", op, st.ErrCashBackFound)
			}
			return "", fmt.Errorf("%s: %w", op, err)
		}

		result = result + strconv.Itoa(int(cashback.Budget)) + " " + cashbackType.NameType + " " + cashback.ValueCondition + ", "
	}

	return result, nil
}

func (s *Storage) ChangeBudgetCashBack(ctx context.Context, idCashBack int32, budget int32) (string, error) {
	const op = "Storage.SQLite.ChangeBudgetCashBack"

	err := s.CheckContainCashBack(ctx, idCashBack)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	stmt, err := s.db.Prepare("UPDATE CashBack SET Budget = ? WHERE Id = ?")
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	// Выполняем запрос, передав параметры
	_, err = stmt.ExecContext(ctx, budget, idCashBack)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	return "complete", nil
}

func (s *Storage) ChangeTypeCashBack(ctx context.Context, idCashBack int32, typeCashBack int32) (string, error) {
	const op = "Storage.SQLite.ChangeTypeCashBack"

	err := s.CheckContainCashBack(ctx, idCashBack)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	statement, err := s.db.Prepare(`SELECT Id FROM CashBack WHERE TypeCashBackFK = ?`)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	res := statement.QueryRow(typeCashBack)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}
	var dataFromDB int32
	err = res.Scan(&dataFromDB)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	stmt, err := s.db.Prepare("UPDATE CashBack SET TypeCashBackFK = ? WHERE Id = ?")
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	// Выполняем запрос, передав параметры
	_, err = stmt.ExecContext(ctx, typeCashBack, idCashBack)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	return "complete", nil
}

func (s *Storage) ChangeConditionCashBack(ctx context.Context, idCashBack int32, condition string) (string, error) {
	const op = "Storage.SQLite.ChangeConditionCashBack"

	err := s.CheckContainCashBack(ctx, idCashBack)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	stmt, err := s.db.Prepare("UPDATE CashBack SET ValueCondition = ? WHERE Id = ?")
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	// Выполняем запрос, передав параметры
	_, err = stmt.ExecContext(ctx, condition, idCashBack)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	return "complete", nil
}

func (s *Storage) DeleteCashBack(ctx context.Context, idCashBack int32) (string, error) {
	const op = "Storage.SQLite.DeleteCashBack"

	err := s.CheckContainCashBack(ctx, idCashBack)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	stmt, err := s.db.Prepare("DELETE FROM CashBack WHERE Id = ?")
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	// Выполняем запрос, передав параметры
	_, err = stmt.ExecContext(ctx, idCashBack)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	return "complete", nil
}

func (s *Storage) CheckContainPromoCode(ctx context.Context, elementForSearch string) error {
	const op = "Storage.SQLite.CheckContainPromoCode"
	statement, err := s.db.Prepare("SELECT Name FROM PromoCodes WHERE Name = ?")
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	res := statement.QueryRowContext(ctx, elementForSearch)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	var dataFromDB string
	err = res.Scan(&dataFromDB)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}
func (s *Storage) CheckContainCashBack(ctx context.Context, elementForSearch int32) error {
	const op = "Storage.SQLite.CheckContainCashBack"
	statement, err := s.db.Prepare("SELECT Id FROM CashBack WHERE Id = ?")
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	res := statement.QueryRowContext(ctx, elementForSearch)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	var dataFromDB int32
	err = res.Scan(&dataFromDB)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}

func (s *Storage) CheckContainClient(ctx context.Context, idClient int32) error {
	const op = "Storage.SQLite.CheckContainClient"
	statement, err := s.db.Prepare("SELECT Id FROM Clients WHERE Id = ?")
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	res := statement.QueryRowContext(ctx, idClient)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	var dataFromDB int32
	err = res.Scan(&dataFromDB)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}

func (s *Storage) CheckContainGroup(ctx context.Context, elementForSearch int32) error {
	const op = "Storage.SQLite.CheckContainGroup"
	statement, err := s.db.Prepare("SELECT Id FROM TypesOfGroups WHERE Id = ?")
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	res := statement.QueryRowContext(ctx, elementForSearch)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	var dataFromDB int32
	err = res.Scan(&dataFromDB)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}

func (s *Storage) CheckContainCashBackType(ctx context.Context, elementForSearch int32) error {
	const op = "Storage.SQLite.CheckContainCashBackType"
	statement, err := s.db.Prepare("SELECT Id FROM CashBackTypes WHERE Id = ?")
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	res := statement.QueryRowContext(ctx, elementForSearch)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	var dataFromDB int32
	err = res.Scan(&dataFromDB)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}
