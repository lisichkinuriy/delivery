package orderrepo

import (
	"context"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"lisichkinuriy/delivery/internal/domain/order"
	"lisichkinuriy/delivery/internal/domain/vo"
	"testing"
)

func setupTest(t *testing.T) (context.Context, *gorm.DB, error) {
	ctx := context.Background()

	dsn := "host=localhost user=test password=test dbname=test port=5491 sslmode=disable TimeZone=Europe/Moscow"
	// Подключаемся к БД через Gorm
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	require.NoError(t, err)

	// Авто миграция (создаём таблицу)
	err = db.AutoMigrate(&OrderDTO{})
	require.NoError(t, err)

	tx := db.Begin()
	return ctx, tx, nil
}

func Test_OrderRepositoryShouldCanAddOrder(t *testing.T) {
	// Инициализируем окружение
	ctx, db, err := setupTest(t)
	require.NoError(t, err)

	// Создаем репозиторий
	orderRepository, err := NewRepository(db)
	require.NoError(t, err)

	// Вызываем Add
	location := vo.MaxLocation()
	orderAggregate, err := order.NewOrder(uuid.New(), location)
	err = orderRepository.Add(ctx, orderAggregate)
	require.NoError(t, err)

	dbAggregate, err := orderRepository.Get(ctx, orderAggregate.ID())
	assert.NoError(t, err)

	// Проверяем эквивалентность
	require.Equal(t, orderAggregate.ID(), dbAggregate.ID())
	require.Equal(t, orderAggregate.Location(), dbAggregate.Location())
}
