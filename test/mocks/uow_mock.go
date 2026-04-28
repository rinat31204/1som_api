package mocks

import (
	"1som/internal/domain/account"
	"1som/internal/domain/transaction"
	"1som/internal/domain/user"
	"1som/internal/domain/wallet"
	"context"

	"github.com/stretchr/testify/mock"
)

// MockUnitOfWork мок для unitOfWork.UnitOfWork
type MockUnitOfWork struct {
	mock.Mock
	userRepo        *MockUserRepo
	walletRepo      *MockWalletRepo
	accountRepo     *MockAccountRepo
	transactionRepo *MockTransactionRepo
}

// NewMockUnitOfWork конструктор с опциональными репозиториями (если nil, создаются новые пустые моки)
func NewMockUnitOfWork(
	userRepo *MockUserRepo,
	walletRepo *MockWalletRepo,
	accountRepo *MockAccountRepo,
	transactionRepo *MockTransactionRepo,
) *MockUnitOfWork {
	if userRepo == nil {
		userRepo = new(MockUserRepo)
	}
	if walletRepo == nil {
		walletRepo = new(MockWalletRepo)
	}
	if accountRepo == nil {
		accountRepo = new(MockAccountRepo)
	}
	if transactionRepo == nil {
		transactionRepo = new(MockTransactionRepo)
	}
	return &MockUnitOfWork{
		userRepo:        userRepo,
		walletRepo:      walletRepo,
		accountRepo:     accountRepo,
		transactionRepo: transactionRepo,
	}
}

func (m *MockUnitOfWork) Begin(ctx context.Context) error {
	args := m.Called(ctx)
	return args.Error(0)
}

func (m *MockUnitOfWork) Rollback(ctx context.Context) error {
	args := m.Called(ctx)
	return args.Error(0)
}

func (m *MockUnitOfWork) Commit(ctx context.Context) error {
	args := m.Called(ctx)
	return args.Error(0)
}

func (m *MockUnitOfWork) Complete() error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockUnitOfWork) UserRepo() user.Repository {
	return m.userRepo
}

func (m *MockUnitOfWork) WalletRepo() wallet.Repository {
	return m.walletRepo
}

func (m *MockUnitOfWork) AccountRepo() account.Repository {
	return m.accountRepo
}

func (m *MockUnitOfWork) TransactionRepo() transaction.Repository {
	return m.transactionRepo
}

// Методы для доступа к конкретным мокам репозиториев (упрощают настройку ожиданий)
func (m *MockUnitOfWork) MockUserRepo() *MockUserRepo {
	return m.userRepo
}

func (m *MockUnitOfWork) MockWalletRepo() *MockWalletRepo {
	return m.walletRepo
}

func (m *MockUnitOfWork) MockAccountRepo() *MockAccountRepo {
	return m.accountRepo
}

func (m *MockUnitOfWork) MockTransactionRepo() *MockTransactionRepo {
	return m.transactionRepo
}
