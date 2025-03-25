package courierrepo

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"lisichkinuriy/delivery/internal/domain/courier"
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
	err = db.AutoMigrate(&CourierDTO{})
	require.NoError(t, err)

	err = db.AutoMigrate(&TransportDTO{})
	require.NoError(t, err)

	tx := db.Begin()
	return ctx, tx, nil
}

func Test_CourierRepositoryShouldCanAddCourier(t *testing.T) {
	// Инициализируем окружение
	ctx, db, err := setupTest(t)
	require.NoError(t, err)

	// Создаем репозиторий
	courierRepo, err := NewRepository(db)
	require.NoError(t, err)

	// Вызываем Add

	name, _ := vo.FakeCourierName()
	transportName, _ := vo.FakeTransportName()
	speed, _ := vo.FakeSpeed()
	location, _ := vo.FakeLocation()
	courierAggregate, _ := courier.NewCourier(name, transportName, speed, location)
	err = courierRepo.Add(ctx, courierAggregate)
	require.NoError(t, err)

	dbAggregate, err := courierRepo.Get(ctx, courierAggregate.ID())
	assert.NoError(t, err)

	// Проверяем эквивалентность
	require.Equal(t, name, dbAggregate.Name())
	require.Equal(t, speed, dbAggregate.Transport().Speed())
	require.Equal(t, transportName, dbAggregate.Transport().Name())
}
