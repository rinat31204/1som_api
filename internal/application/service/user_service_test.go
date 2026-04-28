package service_test

import (
	"1som/internal/application/service"
	"1som/internal/domain/unitOfWork"
	"1som/internal/domain/user"
	"1som/test/mocks"
	"testing"

	"github.com/stretchr/testify/mock"
)

func TestUserService_HandleUserVerifiedEvent_Success(t *testing.T) {
	// Создаём моки репозиториев
	userRepo := &mocks.MockUserRepo{}
	walletRepo := &mocks.MockWalletRepo{}
	accountRepo := &mocks.MockAccountRepo{}
	txRepo := &mocks.MockTransactionRepo{}

	// Создаём мок UnitOfWork
	uow := mocks.NewMockUnitOfWork(userRepo, walletRepo, accountRepo, txRepo)

	// Фабрика возвращает этот uow
	uowFactory := func() unitOfWork.UnitOfWork { return uow }

	// Фейк-шина событий
	fakeBus := mocks.NewFakeEventBus()

	// Создаём сервис – он подписывается на EventUserVerified
	_ = service.NewUserService(uowFactory, fakeBus) // svc не используется напрямую, но подписка произошла

	// Подготавливаем событие
	userID := ""
	email := "test@example.com"
	event := service.UserVerifiedEvent{UserID: userID, Email: email}

	// Настраиваем ожидания
	uow.On("Begin", mock.Anything).Return(nil).Once()
	userRepo.On("ExistsUserByIdOrEmail", userID, email).Return(false, nil).Once()
	userRepo.On("Create", mock.MatchedBy(func(u *user.User) bool {
		return u.Id == userID && u.Email == email
	})).Return(nil).Once()
	walletRepo.On("Create", mock.Anything).Return(nil).Once()
	accountRepo.On("Create", mock.Anything).Return(nil).Once()
	txRepo.On("Create", mock.Anything).Return(nil).Once()
	uow.On("Commit", mock.Anything).Return(nil).Once()
	uow.On("Rollback", mock.Anything).Return(nil).Once() // вызывается из defer после коммита

	// Публикуем событие – обработчик должен выполниться
	fakeBus.Publish(service.EventUserVerified, event)

	// Проверяем, что все ожидания выполнены
	uow.AssertExpectations(t)
	userRepo.AssertExpectations(t)
	walletRepo.AssertExpectations(t)
	accountRepo.AssertExpectations(t)
	txRepo.AssertExpectations(t)
}

// Остальные тесты (ошибки, валидация) пишутся аналогично
