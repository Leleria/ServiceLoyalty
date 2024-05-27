package SQLite

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	sl "github.com/Leleria/Contract/GeneratedFilesProtoBufGo"
	"github.com/Leleria/ServiceLoyalty/Internal/Domain/Models"
	st "github.com/Leleria/ServiceLoyalty/Internal/Storage"
	"github.com/mattn/go-sqlite3"
	"strconv"
	"time"
)

type Storage struct {
	db *sql.DB
}

var bufferChannel chan string

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

func (s *Storage) GetClient(ctx context.Context, idClient int32) (string, string, int32, string, error) {
	const op = "Storage.Sqlite.GetClient"

	err := s.CheckContainClient(ctx, idClient)
	if err != nil {
		return "", "", 0, "", fmt.Errorf("%s: %w", op, err)
	}
	stmt, err := s.db.Prepare("select Name, Email, Clients.CountBonuses, LoyaltyLevels.NameLevel FROM Clients INNER JOIN LoyaltyLevels ON Clients.LoyaltyLevelFK = LoyaltyLevels.Id WHERE Clients.Id = ?")
	if err != nil {
		return "", "", 0, "", fmt.Errorf("%s: %w", op, err)
	}

	row := stmt.QueryRowContext(ctx, idClient)

	var client Models.Client
	var loyaltyLevelClient Models.LoyaltyLevel
	err = row.Scan(&client.Name, &client.Email, &client.CountBonuses, &loyaltyLevelClient.NameLevel)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", "", 0, "", fmt.Errorf("%s: %w", op, st.ErrClientFound)
		}

		return "", "", 0, "", fmt.Errorf("%s: %w", op, err)
	}

	return client.Name, client.Email, client.CountBonuses, loyaltyLevelClient.NameLevel, nil
}

func (s *Storage) GetAllClients(ctx context.Context) ([]*sl.Client, error) {
	const op = "Storage.Sqlite.GetAllClients"

	var clients []*sl.Client
	stmt, err := s.db.Prepare("select Name, Email, Clients.CountBonuses, LoyaltyLevels.NameLevel FROM Clients INNER JOIN LoyaltyLevels ON Clients.LoyaltyLevelFK = LoyaltyLevels.Id")
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	row, err := stmt.QueryContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	for row.Next() {
		var client Models.Client
		var loyaltyLevelClient Models.LoyaltyLevel
		err := row.Scan(&client.Name, &client.Email, &client.CountBonuses, &loyaltyLevelClient.NameLevel)

		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return nil, fmt.Errorf("%s: %w", op, st.ErrClientFound)
			}
			return nil, fmt.Errorf("%s: %w", op, err)
		}

		clients = append(clients, &sl.Client{
			Name:         client.Name,
			Email:        client.Email,
			CountBonuses: client.CountBonuses,
			LoyaltyLevel: loyaltyLevelClient.NameLevel,
		})
	}

	return clients, nil
}

func (s *Storage) GetOperation(ctx context.Context, idOperation int32) (string, string, int32, string, error) {
	const op = "Storage.Sqlite.GetOperation"

	err := s.CheckContainOperation(ctx, idOperation)
	if err != nil {
		return "", "", 0, "", fmt.Errorf("%s: %w", op, err)
	}
	stmt, err := s.db.Prepare("select TypesOperations.NameTypeOperation, Clients.Name, Operations.CountBonuses, DateAndTimeOperation FROM Operations " +
		"INNER JOIN TypesOperations ON Operations.TypeOperationFK = TypesOperations.Id INNER JOIN Clients ON Operations.ClientFK = Clients.Id  WHERE Operations.Id = ?")
	if err != nil {
		return "", "", 0, "", fmt.Errorf("%s: %w", op, err)
	}

	row := stmt.QueryRowContext(ctx, idOperation)

	var operation Models.Operation
	var client Models.Client
	var typeOperation Models.TypeOperation
	err = row.Scan(&typeOperation.NameTypeOperation, &client.Name, &operation.CountBonuses, &operation.DateAndTimeOperation)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", "", 0, "", fmt.Errorf("%s: %w", op, st.ErrOperationFound)
		}

		return "", "", 0, "", fmt.Errorf("%s: %w", op, err)
	}

	//result := typeOperation.NameTypeOperation + " " + client.Name + " " + strconv.Itoa(int(operation.CountBonuses)) + " " + operation.DateAndTimeOperation
	return typeOperation.NameTypeOperation, client.Name, operation.CountBonuses, operation.DateAndTimeOperation, nil
}

func (s *Storage) GetAllOperations(ctx context.Context) ([]*sl.Operation, error) {
	const op = "Storage.Sqlite.GetAllOperations"

	var operations []*sl.Operation
	stmt, err := s.db.Prepare("select TypesOperations.NameTypeOperation, Clients.Name, Operations.CountBonuses, DateAndTimeOperation FROM Operations " +
		"INNER JOIN TypesOperations ON Operations.TypeOperationFK = TypesOperations.Id INNER JOIN Clients ON Operations.ClientFK = Clients.Id")
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	row, err := stmt.QueryContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	var result string
	for row.Next() {
		var operation Models.Operation
		var client Models.Client
		var typeOperation Models.TypeOperation
		err := row.Scan(&typeOperation.NameTypeOperation, &client.Name, &operation.CountBonuses, &operation.DateAndTimeOperation)

		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return nil, fmt.Errorf("%s: %w", op, st.ErrOperationFound)
			}
			return nil, fmt.Errorf("%s: %w", op, err)
		}

		result = result + typeOperation.NameTypeOperation + " " + client.Name + " " + strconv.Itoa(int(operation.CountBonuses)) + " " + operation.DateAndTimeOperation + ", "
		operations = append(operations, &sl.Operation{
			TypeDiscount: typeOperation.NameTypeOperation,
			Client:       client.Name,
			CountBonuses: operation.CountBonuses,
			DateAndTime:  operation.DateAndTimeOperation,
		})
	}

	return operations, nil
}

func (s *Storage) ChangeClientPersonalPromoCode(ctx context.Context, name string, idClient int32) (string, error) {
	const op = "Storage.SQLite.ChangeClientPersonalPromoCode"

	err := s.CheckContainPersonalPromoCode(ctx, name)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	statement, err := s.db.Prepare(`SELECT NamePromoCode FROM PersonalPromoCodes WHERE ClientFK = ?`)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	res := statement.QueryRow(idClient)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}
	var dataFromDB string
	err = res.Scan(&dataFromDB)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	stmt, err := s.db.Prepare("UPDATE PersonalPromoCodes SET ClientFK = ? WHERE NamePromoCode = ?")
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	// Выполняем запрос, передав параметры
	_, err = stmt.ExecContext(ctx, idClient, name)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	return "complete", nil
}

func (s *Storage) ChangeGroupPersonalPromoCode(ctx context.Context, name string, idGroup int32) (string, error) {
	const op = "Storage.SQLite.ChangeGroupPersonalPromoCode"

	err := s.CheckContainPersonalPromoCode(ctx, name)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	statement, err := s.db.Prepare(`SELECT NamePromoCode FROM PersonalPromoCodes WHERE GroupFK = ?`)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	res := statement.QueryRow(idGroup)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}
	var dataFromDB string
	err = res.Scan(&dataFromDB)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	stmt, err := s.db.Prepare("UPDATE PersonalPromoCodes SET GroupFK = ? WHERE NamePromoCode = ?")
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	// Выполняем запрос, передав параметры
	_, err = stmt.ExecContext(ctx, idGroup, name)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	return "complete", nil
}
func (s *Storage) ChangeNamePersonalPromoCode(ctx context.Context, name string, newName string) (string, error) {
	const op = "Storage.SQLite.ChangeNamePersonalPromoCode"

	err := s.CheckContainPersonalPromoCode(ctx, name)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	stmt, err := s.db.Prepare("UPDATE PersonalPromoCodes SET NamePromoCode = ? WHERE NamePromoCode = ?")
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
func (s *Storage) ChangeTypeDiscountPersonalPromoCode(ctx context.Context, name string, typeDiscount int32) (string, error) {
	const op = "Storage.SQLite.ChangeTypeDiscountPromoCode"

	err := s.CheckContainPersonalPromoCode(ctx, name)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	statement, err := s.db.Prepare(`SELECT ValueDiscount FROM PersonalPromoCodes WHERE NamePromoCode = ?`)
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

	stmt, err := s.db.Prepare("UPDATE PersonalPromoCodes SET TypeDiscountFK = ? WHERE NamePromoCode = ?")
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
func (s *Storage) ChangeValueDiscountPersonalPromoCode(ctx context.Context, name string, valueDiscount int32) (string, error) {
	const op = "Storage.SQLite.ChangeValueDiscountPersonalPromoCode"

	err := s.CheckContainPersonalPromoCode(ctx, name)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	statement, err := s.db.Prepare(`SELECT TypeDiscountFK FROM PersonalPromoCodes WHERE NamePromoCode = ?`)
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

	stmt, err := s.db.Prepare("UPDATE PersonalPromoCodes SET ValueDiscount = ? WHERE NamePromoCode = ?")
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
func (s *Storage) ChangeDateStartActivePersonalPromoCode(ctx context.Context, name string, dateStartActive string) (string, error) {
	const op = "Storage.SQLite.ChangeDateStartActivePersonalPromoCode"

	err := s.CheckContainPersonalPromoCode(ctx, name)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	statement, err := s.db.Prepare(`SELECT DateFinishActive FROM PersonalPromoCodes WHERE NamePromoCode = ?`)
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

	stmt, err := s.db.Prepare("UPDATE PersonalPromoCodes SET DateStartActive = ? WHERE NamePromoCode = ?")
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
func (s *Storage) ChangeDateFinishActivePersonalPromoCode(ctx context.Context, name string, dateFinishActive string) (string, error) {
	const op = "Storage.SQLite.ChangeDateFinishActivePersonalPromoCode"

	err := s.CheckContainPersonalPromoCode(ctx, name)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	statement, err := s.db.Prepare(`SELECT DateStartActive FROM PersonalPromoCodes WHERE NamePromoCode = ?`)
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

	stmt, err := s.db.Prepare("UPDATE PersonalPromoCodes SET DateFinishActive = ? WHERE NamePromoCode = ?")
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

func (s *Storage) AccrualBonusesCashback(ctx context.Context, idClient int32, idCashBack int32) (string, error) {
	const op = "Storage.SQLite.AccrualBonusesCashback"

	err := s.CheckContainClient(ctx, idClient)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	err = s.CheckContainCashBack(ctx, idCashBack)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	stmt, err := s.db.Prepare("select Budget FROM CashBack WHERE Id = ?")
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	row := stmt.QueryRowContext(ctx, idCashBack)
	var budget int32

	err = row.Scan(&budget)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", fmt.Errorf("%s: %w", op, st.ErrCashBackFound)
		}
		return "", fmt.Errorf("%s: %w", op, err)
	}

	stmt, err = s.db.Prepare("select CountBonuses FROM Clients WHERE Id = ?")
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	row = stmt.QueryRowContext(ctx, idClient)
	var countBonuses int32

	err = row.Scan(&countBonuses)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", fmt.Errorf("%s: %w", op, st.ErrClientFound)
		}
		return "", fmt.Errorf("%s: %w", op, err)
	}

	stmt, err = s.db.Prepare("UPDATE Clients SET CountBonuses = ? WHERE Id = ?")
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	// Выполняем запрос, передав параметры
	_, err = stmt.ExecContext(ctx, countBonuses+budget, idClient)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	stmt, err = s.db.Prepare("INSERT INTO Operations(TypeOperationFK, ClientFK, CountBonuses, DateAndTimeOperation) VALUES(?, ?, ?, ?)")
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	timeNow := time.Now()
	currentTime := timeNow.Format("2006-01-02 15:04:05")
	// Выполняем запрос, передав параметры
	_, err = stmt.ExecContext(ctx, 1, idClient, budget, currentTime)
	if err != nil {
		var sqliteErr sqlite3.Error

		if errors.As(err, &sqliteErr) && sqliteErr.ExtendedCode == sqlite3.ErrConstraintUnique {
			return "", fmt.Errorf("%s: %w", op, st.ErrPromoCodeExists)
		}

		return "", fmt.Errorf("%s: %w", op, err)
	}
	return "complete", nil
}
func (s *Storage) CalculatePriceWithPromoCode(ctx context.Context, idClient int32, namePromoCode string,
	amountProduct float32) (string, float32, float32, error) {
	const op = "Storage.SQLite.CalculatePriceWithPromoCode"

	err := s.CheckContainPromoCode(ctx, namePromoCode)
	if err != nil {
		return "", 0, 0, fmt.Errorf("%s: %w", op, err)
	}
	err = s.CheckContainClient(ctx, idClient)
	if err != nil {
		return "", 0, 0, fmt.Errorf("%s: %w", op, err)
	}

	stmt, err := s.db.Prepare("select PromoCodes.Id, TypesOfDiscounts.NameType, ValueDiscount, DateFinishActive," +
		" MaxCountUses FROM PromoCodes INNER JOIN TypesOfDiscounts ON PromoCodes.TypeDiscountFK = TypesOfDiscounts.Id  WHERE PromoCodes.Name = ?")
	if err != nil {
		return "", 0, 0, fmt.Errorf("%s: %w", op, err)
	}

	row := stmt.QueryRowContext(ctx, namePromoCode)

	var promoCode Models.PromoCode
	var typeDiscountPromoCode Models.TypeDiscount
	err = row.Scan(&promoCode.Id, &typeDiscountPromoCode.NameType, &promoCode.ValueDiscount,
		&promoCode.DateFinishActive, &promoCode.MaxCountUses)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", 0, 0, fmt.Errorf("%s: %w", op, st.ErrPromoCodeFound)
		}

		return "", 0, 0, fmt.Errorf("%s: %w", op, err)
	}

	//resultIdPromoCode := promoCode.Id
	resultTypeDiscountPromoCode := typeDiscountPromoCode.NameType
	resultValueDiscountPromoCode := promoCode.ValueDiscount
	resultDateFinishActivePromoCode := promoCode.DateFinishActive
	resultMaxCountUsesPromoCode := promoCode.MaxCountUses

	stmt, err = s.db.Prepare("select Id FROM PersonalPromoCodes WHERE ClientFK = ? AND NamePromoCode = ?")
	if err != nil {
		return "", 0, 0, fmt.Errorf("%s: %w", op, err)
	}

	row = stmt.QueryRowContext(ctx, idClient, namePromoCode)
	var idPersonal int32

	var finalAmount float32
	var amountDiscount float32
	err = row.Scan(&idPersonal)
	if err != nil {
		timeNow := time.Now()
		currentTime := timeNow.Format("2006-01-02")
		if resultDateFinishActivePromoCode < currentTime {
			return "promo code is expired", 0, 0, fmt.Errorf("%s: %w", op, "promo code is expired")
		}

		if resultMaxCountUsesPromoCode == 0 {
			return "max count uses is zero", 0, 0, fmt.Errorf("%s: %w", op, "max count uses is zero")
		}

		if resultTypeDiscountPromoCode == "фиксированная сумма" {

			amountDiscount = float32(resultValueDiscountPromoCode)

			if amountProduct > float32(resultValueDiscountPromoCode) {

				bufferChannel = make(chan string, 3)
				finalAmount = amountProduct - float32(resultValueDiscountPromoCode)
				bufferChannel <- strconv.Itoa(int(idClient))
				bufferChannel <- namePromoCode
				bufferChannel <- strconv.Itoa(int(resultMaxCountUsesPromoCode))

			} else {
				return "base price is less than the fixed amount of the promo code", finalAmount, amountDiscount, nil
			}
		}
		if resultTypeDiscountPromoCode == "процентная" {

			amountDiscount = (amountProduct * float32(resultValueDiscountPromoCode)) / 100
			finalAmount = amountProduct - amountDiscount
			bufferChannel = make(chan string, 3)
			bufferChannel <- strconv.Itoa(int(idClient))
			bufferChannel <- namePromoCode
			bufferChannel <- strconv.Itoa(int(resultMaxCountUsesPromoCode))

		}
	} else {
		timeNow := time.Now()
		currentTime := timeNow.Format("2006-01-02")
		if resultDateFinishActivePromoCode < currentTime {
			return "promo code is expired", 0, 0, fmt.Errorf("%s: %w", op, "promo code is expired")
		}

		if resultMaxCountUsesPromoCode == 0 {
			return "max count uses is zero", 0, 0, fmt.Errorf("%s: %w", op, "max count uses is zero")
		}

		if resultTypeDiscountPromoCode == "фиксированная сумма" {

			amountDiscount = float32(resultValueDiscountPromoCode)

			if amountProduct > float32(resultValueDiscountPromoCode) {

				bufferChannel = make(chan string, 4)
				finalAmount = amountProduct - float32(resultValueDiscountPromoCode)
				bufferChannel <- strconv.Itoa(int(idClient))
				bufferChannel <- namePromoCode
				bufferChannel <- strconv.Itoa(int(resultMaxCountUsesPromoCode))
				bufferChannel <- strconv.Itoa(int(idPersonal))

			} else {
				return "base price is less than the fixed amount of the promo code", finalAmount, amountDiscount, nil
			}
		}
		if resultTypeDiscountPromoCode == "процентная" {

			amountDiscount = (amountProduct * float32(resultValueDiscountPromoCode)) / 100
			finalAmount = amountProduct - amountDiscount
			bufferChannel = make(chan string, 4)
			bufferChannel <- strconv.Itoa(int(idClient))
			bufferChannel <- namePromoCode
			bufferChannel <- strconv.Itoa(int(resultMaxCountUsesPromoCode))
			bufferChannel <- strconv.Itoa(int(idPersonal))

		}

	}

	return "complete", finalAmount, amountDiscount, nil

}
func (s *Storage) CalculatePriceWithBonuses(ctx context.Context, idClient int32,
	amountProduct float32) (string, float32, float32, error) {
	const op = "Storage.SQLite.CalculatePriceWithBonuses"

	err := s.CheckContainClient(ctx, idClient)
	if err != nil {
		return "", 0, 0, fmt.Errorf("%s: %w", op, err)
	}

	stmt, err := s.db.Prepare("select CountBonuses FROM Clients WHERE Id = ?")
	if err != nil {
		return "", 0, 0, fmt.Errorf("%s: %w", op, err)
	}

	row := stmt.QueryRowContext(ctx, idClient)

	var countBonuses int32
	err = row.Scan(&countBonuses)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", 0, 0, fmt.Errorf("%s: %w", op, st.ErrPromoCodeFound)
		}

		return "", 0, 0, fmt.Errorf("%s: %w", op, err)
	}

	resultCountBonuses := countBonuses
	var finalAmount float32
	var numberBonusesDebited float32
	var amountDiscount float32

	if resultCountBonuses > 0 {
		bufferChannel = make(chan string, 2)
		amountDiscount = (amountProduct * 30) / 100

		if countBonuses <= int32(amountDiscount) {
			//списываем бонусы все
			finalAmount = amountProduct - float32(countBonuses)
			numberBonusesDebited = float32(countBonuses)
			bufferChannel <- strconv.Itoa(int(idClient))
			bufferChannel <- strconv.Itoa(int(countBonuses))

		} else {
			//списываем discount
			finalAmount = amountProduct - amountDiscount
			numberBonusesDebited = amountDiscount
			bufferChannel <- strconv.Itoa(int(idClient))
			bufferChannel <- strconv.Itoa(int(numberBonusesDebited))
		}
	} else {
		return "there are not enough bonuses in the account", finalAmount, numberBonusesDebited, nil
	}
	return "complete", finalAmount, numberBonusesDebited, nil

}
func (s *Storage) DebitingPromoBonuses(ctx context.Context, idClient int32, paymentStatus bool) (string, error) {
	const op = "Storage.SQLite.DebitingPromoBonuses"

	err := s.CheckContainClient(ctx, idClient)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	if paymentStatus {
		if bufferChannel == nil {
			return "no information about the shares has been entered", nil
		} else if len(bufferChannel) == 4 {

			bufferIdClient := <-bufferChannel
			bIdClient, _ := strconv.Atoi(bufferIdClient)
			if idClient == int32(bIdClient) {
				namePromoCode := <-bufferChannel
				maxCountUses := <-bufferChannel

				countUses, _ := strconv.Atoi(maxCountUses)
				idPersonalPromoCode := <-bufferChannel

				_, err := s.ChangeMaxCountUsesPromoCode(ctx, namePromoCode, int32(countUses)-1)
				if err != nil {
					return "", fmt.Errorf("%s: %w", op, err)
				}

				stmt, err := s.db.Prepare("DELETE FROM PersonalPromoCodes WHERE Id = ?")
				if err != nil {
					return "", fmt.Errorf("%s: %w", op, err)
				}

				// Выполняем запрос, передав параметры
				_, err = stmt.ExecContext(ctx, idPersonalPromoCode)
				if err != nil {
					return "", fmt.Errorf("%s: %w", op, err)
				}
				bufferChannel = nil
			} else {
				return "client was not found", nil
			}
		} else if len(bufferChannel) == 3 {

			bufferIdClient := <-bufferChannel
			bIdClient, _ := strconv.Atoi(bufferIdClient)
			if idClient == int32(bIdClient) {
				namePromoCode := <-bufferChannel
				maxCountUses := <-bufferChannel

				countUses, _ := strconv.Atoi(maxCountUses)
				_, err := s.ChangeMaxCountUsesPromoCode(ctx, namePromoCode, int32(countUses)-1)
				if err != nil {
					return "", fmt.Errorf("%s: %w", op, err)
				}
				bufferChannel = nil
			} else {
				return "client was not found", nil
			}
		} else if len(bufferChannel) == 2 {
			bufferClient := <-bufferChannel
			bIdClient, _ := strconv.Atoi(bufferClient)
			if idClient == int32(bIdClient) {
				countBonuses := <-bufferChannel
				countBonusesInt, _ := strconv.Atoi(countBonuses)

				stmt, err := s.db.Prepare("SELECT CountBonuses FROM Clients WHERE Id = ?")
				if err != nil {
					return "", fmt.Errorf("%s: %w", op, err)
				}

				row := stmt.QueryRowContext(ctx, bIdClient)

				var countBonusesClient int
				err = row.Scan(&countBonusesClient)
				if err != nil {
					if errors.Is(err, sql.ErrNoRows) {
						return "", fmt.Errorf("%s: %w", op, st.ErrClientFound)
					}

					return "", fmt.Errorf("%s: %w", op, err)
				}

				stmt, err = s.db.Prepare("UPDATE Clients SET CountBonuses = ? WHERE Id = ?")
				if err != nil {
					return "", fmt.Errorf("%s: %w", op, err)
				}

				// Выполняем запрос, передав параметры
				_, err = stmt.ExecContext(ctx, countBonusesClient-countBonusesInt, idClient)
				if err != nil {
					return "", fmt.Errorf("%s: %w", op, err)
				}

				stmt, err = s.db.Prepare("INSERT INTO Operations(TypeOperationFK, ClientFK, CountBonuses, DateAndTimeOperation) VALUES(?, ?, ?, ?)")
				if err != nil {
					return "", fmt.Errorf("%s: %w", op, err)
				}

				timeNow := time.Now()
				currentTime := timeNow.Format("2006-01-02 15:04:05")
				// Выполняем запрос, передав параметры
				_, err = stmt.ExecContext(ctx, 2, idClient, countBonusesInt, currentTime)
				if err != nil {
					var sqliteErr sqlite3.Error

					if errors.As(err, &sqliteErr) && sqliteErr.ExtendedCode == sqlite3.ErrConstraintUnique {
						return "", fmt.Errorf("%s: %w", op, st.ErrPromoCodeExists)
					}

					return "", fmt.Errorf("%s: %w", op, err)
				}

			} else {
				return "client was not found", nil
			}
			bufferChannel = nil
		}

	} else {
		return "payment has not been confirmed", nil
	}

	return "complete", nil
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
	namePromoCode string, typeDiscount int32, valueDiscount int32, dateStartActive string, dateFinishActive string) (string, error) {
	const op = "Storage.SQLite.SavePersonalPromoCode"

	err := s.CheckContainClient(ctx, idClient)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, st.ErrClientFound)
	}

	err = s.CheckContainGroup(ctx, idGroup)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, st.ErrGroupFound)
	}

	stmt, err := s.db.Prepare("INSERT INTO PersonalPromoCodes(ClientFK, GroupFK, NamePromoCode, TypeDiscountFK, ValueDiscount, DateStartActive, DateFinishActive) VALUES(?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	// Выполняем запрос, передав параметры
	_, err = stmt.ExecContext(ctx, idClient, idGroup, namePromoCode, typeDiscount, valueDiscount, dateStartActive, dateFinishActive)
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

func (s *Storage) DeletePersonalPromoCode(ctx context.Context, name string) (string, error) {
	const op = "Storage.SQLite.DeletePersonalPromoCode"

	err := s.CheckContainPersonalPromoCode(ctx, name)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	stmt, err := s.db.Prepare("DELETE FROM PersonalPromoCodes WHERE NamePromoCode = ?")
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

func (s *Storage) GetPromoCode(ctx context.Context, name string) (string, string, int32, string, string, int32, error) {
	const op = "Storage.Sqlite.GetPromoCode"

	err := s.CheckContainPromoCode(ctx, name)
	if err != nil {
		return "", "", 0, "", "", 0, fmt.Errorf("%s: %w", op, err)
	}
	stmt, err := s.db.Prepare("select TypesOfDiscounts.NameType, ValueDiscount, DateStartActive, DateFinishActive," +
		" MaxCountUses FROM PromoCodes INNER JOIN TypesOfDiscounts ON PromoCodes.TypeDiscountFK = TypesOfDiscounts.Id  WHERE PromoCodes.Name = ?")
	if err != nil {
		return "", "", 0, "", "", 0, fmt.Errorf("%s: %w", op, err)
	}

	row := stmt.QueryRowContext(ctx, name)

	var promoCode Models.PromoCode
	var typeDiscountPromoCode Models.TypeDiscount
	err = row.Scan(&typeDiscountPromoCode.NameType, &promoCode.ValueDiscount, &promoCode.DateStartActive,
		&promoCode.DateFinishActive, &promoCode.MaxCountUses)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", "", 0, "", "", 0, fmt.Errorf("%s: %w", op, st.ErrPromoCodeFound)
		}

		return "", "", 0, "", "", 0, fmt.Errorf("%s: %w", op, err)
	}

	return name, typeDiscountPromoCode.NameType, promoCode.ValueDiscount, promoCode.DateStartActive, promoCode.DateFinishActive, promoCode.MaxCountUses, nil
}

func (s *Storage) GetAllPromoCodes(ctx context.Context) ([]*sl.PromoCode, error) {
	const op = "Storage.Sqlite.GetAllPromoCodes"

	var promoCodes []*sl.PromoCode
	stmt, err := s.db.Prepare("select Name, TypesOfDiscounts.NameType, ValueDiscount, DateStartActive, DateFinishActive," +
		" MaxCountUses FROM PromoCodes INNER JOIN TypesOfDiscounts ON PromoCodes.TypeDiscountFK = TypesOfDiscounts.Id ")
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	row, err := stmt.QueryContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	for row.Next() {
		var promoCode Models.PromoCode
		var typeDiscountPromoCode Models.TypeDiscount
		err := row.Scan(&promoCode.Name, &typeDiscountPromoCode.NameType, &promoCode.ValueDiscount, &promoCode.DateStartActive,
			&promoCode.DateFinishActive, &promoCode.MaxCountUses)

		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return nil, fmt.Errorf("%s: %w", op, st.ErrPromoCodeFound)
			}
			return nil, fmt.Errorf("%s: %w", op, err)
		}

		promoCodes = append(promoCodes, &sl.PromoCode{
			NamePromoCode: promoCode.Name,
			TypeDiscount:  typeDiscountPromoCode.NameType,
			ValueDiscount: promoCode.ValueDiscount,
			DateStart:     promoCode.DateStartActive,
			DateFinish:    promoCode.DateFinishActive,
			MaxCountUses:  promoCode.MaxCountUses,
		})
	}

	return promoCodes, nil
}

func (s *Storage) GetPersonalPromoCode(ctx context.Context, name string) (string, string, string, string, int32, string, string, error) {
	const op = "Storage.Sqlite.GetPersonalPromoCode"

	err := s.CheckContainPersonalPromoCode(ctx, name)
	if err != nil {
		return "", "", "", "", 0, "", "", fmt.Errorf("%s: %w", op, err)
	}
	stmt, err := s.db.Prepare("select Clients.Name, TypesOfGroups.NameType, TypesOfDiscounts.NameType, ValueDiscount, DateStartActive, DateFinishActive FROM PersonalPromoCodes " +
		" INNER JOIN Clients ON PersonalPromoCodes.ClientFK = Clients.Id INNER JOIN TypesOfGroups ON PersonalPromoCodes.GroupFK = TypesOfGroups.Id " +
		"INNER JOIN TypesOfDiscounts ON PersonalPromoCodes.TypeDiscountFK = TypesOfDiscounts.Id WHERE PersonalPromoCodes.NamePromoCode = ?")
	if err != nil {
		return "", "", "", "", 0, "", "", fmt.Errorf("%s: %w", op, err)
	}

	row := stmt.QueryRowContext(ctx, name)

	var personalPromoCode Models.PersonalPromoCode
	var typeDiscountPersonalPromoCode Models.TypeDiscount
	var clientPersonalPromoCode Models.Client
	var groupPersonalPromoCode Models.Group
	err = row.Scan(&clientPersonalPromoCode.Name, &groupPersonalPromoCode.NameType, &typeDiscountPersonalPromoCode.NameType, &personalPromoCode.ValueDiscount, &personalPromoCode.DateStartActive,
		&personalPromoCode.DateFinishActive)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", "", "", "", 0, "", "", fmt.Errorf("%s: %w", op, st.ErrPromoCodeFound)
		}

		return "", "", "", "", 0, "", "", fmt.Errorf("%s: %w", op, err)
	}

	return clientPersonalPromoCode.Name, groupPersonalPromoCode.NameType, name, typeDiscountPersonalPromoCode.NameType, personalPromoCode.ValueDiscount, personalPromoCode.DateStartActive, personalPromoCode.DateFinishActive, nil
}

func (s *Storage) GetAllPersonalPromoCodes(ctx context.Context) ([]*sl.PersonalPromoCode, error) {
	const op = "Storage.Sqlite.GetAllPromoCodes"

	var personalPromoCodes []*sl.PersonalPromoCode
	stmt, err := s.db.Prepare("select NamePromoCode, Clients.Name, TypesOfGroups.NameType, TypesOfDiscounts.NameType, ValueDiscount, DateStartActive, DateFinishActive FROM PersonalPromoCodes " +
		" INNER JOIN Clients ON PersonalPromoCodes.ClientFK = Clients.Id INNER JOIN TypesOfGroups ON PersonalPromoCodes.GroupFK = TypesOfGroups.Id " +
		"INNER JOIN TypesOfDiscounts ON PersonalPromoCodes.TypeDiscountFK = TypesOfDiscounts.Id")
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	row, err := stmt.QueryContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	var result string
	for row.Next() {
		var personalPromoCode Models.PersonalPromoCode
		var typeDiscountPersonalPromoCode Models.TypeDiscount
		var clientPersonalPromoCode Models.Client
		var groupPersonalPromoCode Models.Group
		err := row.Scan(&personalPromoCode.NamePromoCode, &clientPersonalPromoCode.Name, &groupPersonalPromoCode.NameType, &typeDiscountPersonalPromoCode.NameType, &personalPromoCode.ValueDiscount, &personalPromoCode.DateStartActive,
			&personalPromoCode.DateFinishActive)

		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return nil, fmt.Errorf("%s: %w", op, st.ErrPromoCodeFound)
			}
			return nil, fmt.Errorf("%s: %w", op, err)
		}

		result = result + personalPromoCode.NamePromoCode + " " + clientPersonalPromoCode.Name + " " + groupPersonalPromoCode.NameType + " " + typeDiscountPersonalPromoCode.NameType +
			" " + strconv.Itoa(int(personalPromoCode.ValueDiscount)) + " " + personalPromoCode.DateStartActive +
			" " + personalPromoCode.DateFinishActive + " " + ", "

		personalPromoCodes = append(personalPromoCodes, &sl.PersonalPromoCode{
			Client:        clientPersonalPromoCode.Name,
			Group:         groupPersonalPromoCode.NameType,
			NamePromoCode: personalPromoCode.NamePromoCode,
			TypeDiscount:  typeDiscountPersonalPromoCode.NameType,
			ValueDiscount: personalPromoCode.ValueDiscount,
			DateStart:     personalPromoCode.DateStartActive,
			DateFinish:    personalPromoCode.DateFinishActive,
		})
	}

	return personalPromoCodes, nil
}

func (s *Storage) GetCashBack(ctx context.Context, idCashBack int32) (int32, string, string, error) {
	const op = "Storage.Sqlite.GetCashBack"

	err := s.CheckContainCashBack(ctx, idCashBack)
	if err != nil {
		return 0, "", "", fmt.Errorf("%s: %w", op, err)
	}
	stmt, err := s.db.Prepare("SELECT Budget, NameType, ValueCondition FROM CashBack INNER JOIN CashBackTypes ON  CashBack.TypeCashBackFK = CashBackTypes.Id WHERE CashBack.Id = ?")
	if err != nil {
		return 0, "", "", fmt.Errorf("%s: %w", op, err)
	}

	row := stmt.QueryRowContext(ctx, idCashBack)

	var cashback Models.CashBack
	var cashbackType Models.TypeCashBack
	err = row.Scan(&cashback.Budget, &cashbackType.NameType, &cashback.ValueCondition)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, "", "", fmt.Errorf("%s: %w", op, st.ErrCashBackFound)
		}

		return 0, "", "", fmt.Errorf("%s: %w", op, err)
	}

	return cashback.Budget, cashbackType.NameType, cashback.ValueCondition, nil
}

func (s *Storage) GetAllCashBack(ctx context.Context) ([]*sl.CashBack, error) {
	const op = "Storage.Sqlite.GetAllCashBack"

	var cashBacks []*sl.CashBack
	stmt, err := s.db.Prepare("SELECT Budget, CashBackTypes.NameType, ValueCondition FROM CashBack INNER JOIN CashBackTypes ON CashBack.TypeCashBackFK = CashBackTypes.Id ")
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	row, err := stmt.QueryContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	for row.Next() {
		var cashback Models.CashBack
		var cashbackType Models.TypeCashBack
		err := row.Scan(&cashback.Budget, &cashbackType.NameType, &cashback.ValueCondition)

		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return nil, fmt.Errorf("%s: %w", op, st.ErrCashBackFound)
			}
			return nil, fmt.Errorf("%s: %w", op, err)
		}

		cashBacks = append(cashBacks, &sl.CashBack{
			Budget:         cashback.Budget,
			TypeCashBack:   cashbackType.NameType,
			ValueCondition: cashback.ValueCondition,
		})
	}

	return cashBacks, nil
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
func (s *Storage) CheckContainPersonalPromoCode(ctx context.Context, elementForSearch string) error {
	const op = "Storage.SQLite.CheckContainPromoCode"
	statement, err := s.db.Prepare("SELECT NamePromoCode FROM PersonalPromoCodes WHERE NamePromoCode = ?")
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

func (s *Storage) CheckContainOperation(ctx context.Context, idOperation int32) error {
	const op = "Storage.SQLite.CheckContainOperation"
	statement, err := s.db.Prepare("SELECT Id FROM Operations WHERE Id = ?")
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	res := statement.QueryRowContext(ctx, idOperation)
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
func (s *Storage) CheckContainIdPromoCode(ctx context.Context, idPromoCode int32) error {
	const op = "Storage.SQLite.CheckContainIdPromoCode"
	statement, err := s.db.Prepare("SELECT Id FROM PromoCodes WHERE Id = ?")
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	res := statement.QueryRowContext(ctx, idPromoCode)
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
