package service

import (
	"1som/internal/domain/account"
	"1som/internal/domain/account/vo"
	"1som/internal/domain/currency"
	"1som/internal/domain/transaction"
	"1som/internal/domain/unitOfWork"
	"1som/internal/domain/user"
	"1som/internal/domain/wallet"
	"context"
	"errors"
	"log/slog"
	"time"

	"github.com/google/uuid"
)

type UserService struct {
	uowFactory func() unitOfWork.UnitOfWork
	eventBus   EventBus
}

func NewUserService(uowFactory func() unitOfWork.UnitOfWork, eventBus EventBus) *UserService {
	s := &UserService{
		uowFactory: uowFactory,
		eventBus:   eventBus,
	}
	s.subscribe()
	return s
}

func (u *UserService) subscribe() {
	u.eventBus.Subscribe(EventUserVerified, func(event any) {
		ev, ok := event.(UserVerifiedEvent)
		if !ok {
			return
		}

		domainUser := user.User{
			Id:        ev.UserID,
			Email:     ev.Email,
			CreatedAt: time.Now(),
		}

		err := u.createUser(context.Background(), &domainUser)
		if err != nil {
			return
		}
		slog.Info("User created: ", domainUser)
	})
}

func (u *UserService) createUser(ctx context.Context, user *user.User) error {
	uow := u.uowFactory()

	if err := uow.Begin(ctx); err != nil {
		return err
	}
	defer func() {
		_ = uow.Rollback(ctx)
	}()

	exists, err := uow.UserRepo().ExistsUser(user.Id)
	if err != nil {
		return err
	}
	if exists {
		return errors.New("user already exists")
	}

	if err := uow.UserRepo().Create(user); err != nil {
		return err
	}

	wal := wallet.Wallet{
		Id:        uuid.New(),
		UserId:    user.Id,
		CreatedAt: time.Now(),
	}
	if err := uow.WalletRepo().Create(&wal); err != nil {
		return err
	}

	an, _ := vo.GenerateAccountNumber()

	acc := account.Account{
		Id:            uuid.New(),
		WalletId:      wal.Id,
		AccountNumber: an,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
		Balance:       1000,
		Currency:      currency.DefaultCurrency,
	}
	if err := uow.AccountRepo().Create(&acc); err != nil {
		return err
	}

	tx := transaction.Transaction{
		ID:            uuid.New(),
		AccountNumber: an,
		Amount:        1000,
		Currency:      currency.DefaultCurrency,
		Description:   "initial transaction",
		Type:          transaction.Debit,
		CreatedAt:     time.Now(),
	}
	if err := uow.TransactionRepo().Create(&tx); err != nil {
		return err
	}

	if err := uow.Commit(ctx); err != nil {
		return err
	}
	return nil
}
