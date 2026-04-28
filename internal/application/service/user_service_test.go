package service_test

import (
	"1som/internal/application/service"
	"1som/internal/domain/account"
	"1som/internal/domain/currency"
	"1som/internal/domain/transaction"
	"1som/internal/domain/unitOfWork"
	"1som/internal/domain/user"
	"1som/internal/domain/wallet"
	"1som/test/mocks"
	"context"
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestUserService_HandleUserVerifiedEvent_Success(t *testing.T) {
	userRepo := &mocks.MockUserRepo{}
	walletRepo := &mocks.MockWalletRepo{}
	accountRepo := &mocks.MockAccountRepo{}
	txRepo := &mocks.MockTransactionRepo{}

	uow := mocks.NewMockUnitOfWork(userRepo, walletRepo, accountRepo, txRepo)

	uowFactory := func() unitOfWork.UnitOfWork { return uow }

	fakeBus := mocks.NewFakeEventBus()

	_ = service.NewUserService(uowFactory, fakeBus)

	userID := ""
	email := "test@example.com"
	event := service.UserVerifiedEvent{UserID: userID, Email: email}

	uow.On("Begin", mock.Anything).Return(nil).Once()
	userRepo.On("ExistsUserByIdOrEmail", userID, email).Return(false, nil).Once()
	userRepo.On("Create", mock.MatchedBy(func(u *user.User) bool {
		return u.Id == userID && u.Email == email
	})).Return(nil).Once()
	walletRepo.On("Create", mock.Anything).Return(nil).Once()
	accountRepo.On("Create", mock.Anything).Return(nil).Once()
	txRepo.On("Create", mock.Anything).Return(nil).Once()
	uow.On("Commit", mock.Anything).Return(nil).Once()
	uow.On("Rollback", mock.Anything).Return(nil).Once()

	fakeBus.Publish(service.EventUserVerified, event)

	uow.AssertExpectations(t)
	userRepo.AssertExpectations(t)
	walletRepo.AssertExpectations(t)
	accountRepo.AssertExpectations(t)
	txRepo.AssertExpectations(t)
}

func TestUserService_createUserEntity(t *testing.T) {
	tests := []struct {
		name          string
		user          *user.User
		existsReturn  bool
		existsErr     error
		createErr     error
		expectedError error
	}{
		{
			name:          "success",
			user:          &user.User{Id: "123", Email: "test@example.com"},
			existsReturn:  false,
			existsErr:     nil,
			createErr:     nil,
			expectedError: nil,
		},
		{
			name:          "user already exists",
			user:          &user.User{Id: "123", Email: "test@example.com"},
			existsReturn:  true,
			existsErr:     nil,
			createErr:     nil,
			expectedError: errors.New("user already exists"),
		},
		{
			name:          "exists check database error",
			user:          &user.User{Id: "123", Email: "test@example.com"},
			existsReturn:  false,
			existsErr:     errors.New("db error"),
			createErr:     nil,
			expectedError: errors.New("db error"),
		},
		{
			name:          "create user database error",
			user:          &user.User{Id: "123", Email: "test@example.com"},
			existsReturn:  false,
			existsErr:     nil,
			createErr:     errors.New("insert error"),
			expectedError: errors.New("insert error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userRepo := &mocks.MockUserRepo{}
			uow := mocks.NewMockUnitOfWork(userRepo, nil, nil, nil)
			uowFactory := func() unitOfWork.UnitOfWork { return uow }
			fakeBus := mocks.NewFakeEventBus()

			userRepo.On("ExistsUserByIdOrEmail", tt.user.Id, tt.user.Email).
				Return(tt.existsReturn, tt.existsErr).Once()

			if tt.existsErr == nil && !tt.existsReturn {
				userRepo.On("Create", tt.user).Return(tt.createErr).Once()
			}

			svc := service.NewUserService(uowFactory, fakeBus)

			err := svc.CreateUserEntity(context.Background(), uow, tt.user)

			if tt.expectedError != nil {
				require.EqualError(t, err, tt.expectedError.Error())
			} else {
				require.NoError(t, err)
			}
			userRepo.AssertExpectations(t)
		})
	}
}

func TestUserService_CreateUser_FailOnBegin(t *testing.T) {
	uow := &mocks.MockUnitOfWork{}
	uowFactory := func() unitOfWork.UnitOfWork { return uow }
	fakeBus := mocks.NewFakeEventBus()
	svc := service.NewUserService(uowFactory, fakeBus)

	beginErr := errors.New("begin failed")
	uow.On("Begin", mock.Anything).Return(beginErr).Once()

	err := svc.CreateUser(context.Background(), &user.User{})

	require.Error(t, err)
	require.Equal(t, beginErr, err)
	uow.AssertExpectations(t)
}

func TestUserService_createWallet_CreateError(t *testing.T) {
	walletRepo := &mocks.MockWalletRepo{}
	uow := mocks.NewMockUnitOfWork(nil, walletRepo, nil, nil)
	fakeBus := mocks.NewFakeEventBus()
	svc := service.NewUserService(func() unitOfWork.UnitOfWork { return uow }, fakeBus)

	testUser := &user.User{Id: "user-123"}
	expectedErr := errors.New("wallet insert failed")

	walletRepo.On("Create", mock.Anything).Return(expectedErr).Once()

	wal, err := svc.CreateWallet(context.Background(), uow, testUser)

	require.Error(t, err)
	require.Equal(t, expectedErr, err)
	require.Nil(t, wal)
	walletRepo.AssertExpectations(t)
}

func TestUserService_CreateAccount_Success(t *testing.T) {
	accountRepo := &mocks.MockAccountRepo{}
	uow := mocks.NewMockUnitOfWork(nil, nil, accountRepo, nil)
	fakeBus := mocks.NewFakeEventBus()
	svc := service.NewUserService(func() unitOfWork.UnitOfWork { return uow }, fakeBus)

	testWallet := &wallet.Wallet{Id: uuid.New()}

	accountRepo.On("Create", mock.MatchedBy(func(a *account.Account) bool {
		return a.WalletId == testWallet.Id &&
			a.Balance == 1000 &&
			a.Currency == currency.DefaultCurrency &&
			a.AccountNumber != "" &&
			a.Id != uuid.Nil &&
			!a.CreatedAt.IsZero() &&
			!a.UpdatedAt.IsZero()
	})).Return(nil).Once()

	acc, err := svc.CreateAccount(context.Background(), uow, testWallet)

	require.NoError(t, err)
	require.NotNil(t, acc)
	require.Equal(t, testWallet.Id, acc.WalletId)
	require.Equal(t, int64(1000), acc.Balance)
	require.Equal(t, currency.DefaultCurrency, acc.Currency)
	require.NotEmpty(t, acc.AccountNumber)
	accountRepo.AssertExpectations(t)
}

func TestUserService_createAccount_CreateError(t *testing.T) {
	accountRepo := &mocks.MockAccountRepo{}
	uow := mocks.NewMockUnitOfWork(nil, nil, accountRepo, nil)
	fakeBus := mocks.NewFakeEventBus()
	svc := service.NewUserService(func() unitOfWork.UnitOfWork { return uow }, fakeBus)

	testWallet := &wallet.Wallet{Id: uuid.New()}
	expectedErr := errors.New("account insert failed")

	accountRepo.On("Create", mock.Anything).Return(expectedErr).Once()

	acc, err := svc.CreateAccount(context.Background(), uow, testWallet)

	require.Error(t, err)
	require.Equal(t, expectedErr, err)
	require.Nil(t, acc)
	accountRepo.AssertExpectations(t)
}

func TestUserService_createInitialTransaction_Success(t *testing.T) {
	txRepo := &mocks.MockTransactionRepo{}
	uow := mocks.NewMockUnitOfWork(nil, nil, nil, txRepo)
	fakeBus := mocks.NewFakeEventBus()
	svc := service.NewUserService(func() unitOfWork.UnitOfWork { return uow }, fakeBus)

	testAccount := &account.Account{
		Id:            uuid.New(),
		AccountNumber: "1234567890",
	}

	var createdTx *transaction.Transaction
	txRepo.On("Create", mock.MatchedBy(func(tx *transaction.Transaction) bool {
		createdTx = tx
		return tx.AccountNumber == testAccount.AccountNumber &&
			tx.Amount == 1000 &&
			tx.Currency == currency.DefaultCurrency &&
			tx.Description == "initial transaction" &&
			tx.Type == transaction.Debit &&
			tx.ID != uuid.Nil &&
			!tx.CreatedAt.IsZero()
	})).Return(nil).Once()

	err := svc.CreateInitialTransaction(context.Background(), uow, testAccount)

	require.NoError(t, err)
	require.NotNil(t, createdTx)
	txRepo.AssertExpectations(t)
}

func TestUserService_createInitialTransaction_CreateError(t *testing.T) {
	txRepo := &mocks.MockTransactionRepo{}
	uow := mocks.NewMockUnitOfWork(nil, nil, nil, txRepo)
	fakeBus := mocks.NewFakeEventBus()
	svc := service.NewUserService(func() unitOfWork.UnitOfWork { return uow }, fakeBus)

	testAccount := &account.Account{AccountNumber: "1234567890"}
	expectedErr := errors.New("transaction insert failed")

	txRepo.On("Create", mock.Anything).Return(expectedErr).Once()

	err := svc.CreateInitialTransaction(context.Background(), uow, testAccount)

	require.Error(t, err)
	require.Equal(t, expectedErr, err)
	txRepo.AssertExpectations(t)
}

func TestUserService_CreateWallet_Error_RollbackOnce(t *testing.T) {
	userRepo := &mocks.MockUserRepo{}
	walletRepo := &mocks.MockWalletRepo{}
	accountRepo := &mocks.MockAccountRepo{}
	txRepo := &mocks.MockTransactionRepo{}

	uow := mocks.NewMockUnitOfWork(userRepo, walletRepo, accountRepo, txRepo)

	fakeBus := mocks.NewFakeEventBus()
	_ = service.NewUserService(func() unitOfWork.UnitOfWork { return uow }, fakeBus)

	userID := ""
	email := "test@example.com"
	event := service.UserVerifiedEvent{UserID: userID, Email: email}
	expectedErr := errors.New("wallet insert failed")

	uow.On("Begin", mock.Anything).Return(nil).Once()
	userRepo.On("ExistsUserByIdOrEmail", userID, email).Return(false, nil).Once()
	userRepo.On("Create", mock.MatchedBy(func(u *user.User) bool {
		return u.Id == userID && u.Email == email
	})).Return(nil).Once()
	walletRepo.On("Create", mock.Anything).Return(expectedErr).Once()
	uow.On("Rollback", mock.Anything).Return(nil).Once()

	fakeBus.Publish(service.EventUserVerified, event)

	uow.AssertExpectations(t)
	walletRepo.AssertExpectations(t)
}
