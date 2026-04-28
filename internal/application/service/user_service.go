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

		err := u.CreateUser(context.Background(), &domainUser)
		if err != nil {
			return
		}
		slog.Info("User created: ", slog.Any("user", domainUser))
	})
}

func (u *UserService) CreateUser(ctx context.Context, user *user.User) error {
	uow := u.uowFactory()

	if err := uow.Begin(ctx); err != nil {
		return err
	}
	defer func() {
		_ = uow.Rollback(ctx)
	}()

	if err := u.CreateUserEntity(ctx, uow, user); err != nil {
		return err
	}

	wal, err := u.CreateWallet(ctx, uow, user)
	if err != nil {
		return err
	}

	acc, err := u.CreateAccount(ctx, uow, wal)
	if err != nil {
		return err
	}

	if err := u.CreateInitialTransaction(ctx, uow, acc); err != nil {
		return err
	}

	if err := uow.Commit(ctx); err != nil {
		return err
	}
	return nil
}

func (u *UserService) CreateUserEntity(
	ctx context.Context,
	uow unitOfWork.UnitOfWork,
	user *user.User) error {
	exists, err := uow.UserRepo().ExistsUserByIdOrEmail(user.Id, user.Email)
	if err != nil {
		return err
	}
	if exists {
		return errors.New("user already exists")
	}

	return uow.UserRepo().Create(user)
}

func (u *UserService) CreateWallet(
	ctx context.Context,
	uow unitOfWork.UnitOfWork,
	user *user.User) (*wallet.Wallet, error) {
	wal := &wallet.Wallet{
		Id:        uuid.New(),
		UserId:    user.Id,
		CreatedAt: time.Now(),
	}

	if err := uow.WalletRepo().Create(wal); err != nil {
		return nil, err
	}

	return wal, nil
}

func (u *UserService) CreateAccount(
	ctx context.Context,
	uow unitOfWork.UnitOfWork,
	wal *wallet.Wallet) (*account.Account, error) {
	an, err := vo.GenerateAccountNumber()
	if err != nil {
		return nil, err
	}

	acc := &account.Account{
		Id:            uuid.New(),
		WalletId:      wal.Id,
		AccountNumber: an,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
		Balance:       1000,
		Currency:      currency.DefaultCurrency,
	}

	if err := uow.AccountRepo().Create(acc); err != nil {
		return nil, err
	}

	return acc, nil
}

func (u *UserService) CreateInitialTransaction(
	ctx context.Context,
	uow unitOfWork.UnitOfWork,
	acc *account.Account) error {
	tx := &transaction.Transaction{
		ID:            uuid.New(),
		AccountNumber: acc.AccountNumber,
		Amount:        1000,
		Currency:      currency.DefaultCurrency,
		Description:   "initial transaction",
		Type:          transaction.Debit,
		CreatedAt:     time.Now(),
	}

	return uow.TransactionRepo().Create(tx)
}
