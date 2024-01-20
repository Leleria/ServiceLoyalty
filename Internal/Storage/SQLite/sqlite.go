package SQLite

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/Leleria/ServiceLoyalty/Internal/Domain/Models"
	st "github.com/Leleria/ServiceLoyalty/Internal/Storage"
	"github.com/mattn/go-sqlite3"
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

func (s *Storage) SavePromoCode(ctx context.Context, name string, typeDiscount int32,
	valueDiscount int32, dateStartActive string,
	dateFinishActive string, maxCountUses int32) (string, error) {
	const op = "Storage.SQLite.SavePromoCode"

	// Простой запрос на добавление пользователя
	stmt, err := s.db.Prepare("INSERT INTO PromoCodes(Name, TypeDiscountFK, " +
		"ValueDiscount, DateStartActive, DateFinishActive, MaxCountUses) VALUES(?, ?, ?, ?, ?, ?)")
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	// Выполняем запрос, передав параметры
	_, err = stmt.ExecContext(ctx, name, typeDiscount, valueDiscount, dateStartActive, dateFinishActive, maxCountUses)
	if err != nil {
		var sqliteErr sqlite3.Error

		// Небольшое кунг-фу для выявления ошибки ErrConstraintUnique
		// (см. подробности ниже)
		if errors.As(err, &sqliteErr) && sqliteErr.ExtendedCode == sqlite3.ErrConstraintUnique {
			return "", fmt.Errorf("%s: %w", op, st.ErrPromoCodeExists)
		}

		return "", fmt.Errorf("%s: %w", op, err)
	}

	return "complete", nil
}

func (s *Storage) DeletePromoCode(ctx context.Context, name string) (string, error) {
	const op = "Storage.SQLite.DeletePromoCode"

	err := s.CheckContain(ctx, name)
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

func (s *Storage) PromoCode(ctx context.Context, name string) (Models.PromoCode, error) {
	const op = "Storage.Sqlite.PromoCode"

	stmt, err := s.db.Prepare("SELECT TypeDiscountFK, ValueDiscount, DateStartActive, DateFinishActive" +
		", MaxCountUses FROM PromoCodes WHERE name = ?")
	if err != nil {
		return Models.PromoCode{}, fmt.Errorf("%s: %w", op, err)
	}

	row := stmt.QueryRowContext(ctx, name)

	var promoCode Models.PromoCode
	err = row.Scan(&promoCode.TypeDiscount, &promoCode.ValueDiscount, &promoCode.DateStartActive,
		&promoCode.DateFinishActive, &promoCode.MaxCountUses)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return Models.PromoCode{}, fmt.Errorf("%s: %w", op, st.ErrPromoCodeFound)
		}

		return Models.PromoCode{}, fmt.Errorf("%s: %w", op, err)
	}

	return promoCode, nil
}

func (s *Storage) CheckContain(ctx context.Context, elementForSearch string) error {
	const op = "Storage.SQLite.CheckContain"
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
