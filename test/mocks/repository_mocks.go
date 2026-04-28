package mocks

import (
	"1som/internal/domain/account"
	"1som/internal/domain/transaction"
	"1som/internal/domain/user"
	"1som/internal/domain/wallet"

	"github.com/stretchr/testify/mock"
)

// MockUserRepo мок для user.Repository
type MockUserRepo struct {
	mock.Mock
}

func (m *MockUserRepo) ExistsUserByIdOrEmail(id, email string) (bool, error) {
	args := m.Called(id, email)
	return args.Bool(0), args.Error(1)
}

func (m *MockUserRepo) Create(u *user.User) error {
	args := m.Called(u)
	return args.Error(0)
}

func (m *MockUserRepo) FindById(id string) (*user.User, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*user.User), args.Error(1)
}

// MockWalletRepo мок для wallet.Repository
type MockWalletRepo struct {
	mock.Mock
}

func (m *MockWalletRepo) Create(w *wallet.Wallet) error {
	args := m.Called(w)
	return args.Error(0)
}

// MockAccountRepo мок для account.Repository
type MockAccountRepo struct {
	mock.Mock
}

func (m *MockAccountRepo) Create(a *account.Account) error {
	args := m.Called(a)
	return args.Error(0)
}

// MockTransactionRepo мок для transaction.Repository
type MockTransactionRepo struct {
	mock.Mock
}

func (m *MockTransactionRepo) Create(tx *transaction.Transaction) error {
	args := m.Called(tx)
	return args.Error(0)
}
