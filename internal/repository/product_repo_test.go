package repository_test

import (
	"context"
	"testing"

	"github.com/DucTran999/go-clean-archx/internal/entity"
	"github.com/DucTran999/go-clean-archx/internal/repository"
	"github.com/DucTran999/go-clean-archx/test/datatest"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func newMockDB(t *testing.T) (*gorm.DB, sqlmock.Sqlmock) {
	t.Helper()
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)

	dialector := postgres.New(postgres.Config{
		Conn: db,
	})
	gormDB, err := gorm.Open(dialector, &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	assert.NoError(t, err)

	return gormDB, mock
}

func TestProductRepo_CreateFailed(t *testing.T) {
	t.Parallel()

	// Arrange
	db, mock := newMockDB(t)
	repo := repository.NewProductRepository(db)

	product := &entity.Product{
		Name:  "Mock Product",
		Qty:   5,
		Price: 99.99,
	}

	mock.ExpectBegin()
	mock.ExpectQuery(`INSERT INTO "products"`).
		WithArgs(
			product.Name,
			product.Qty,
			product.Price,
			sqlmock.AnyArg(),
		).
		WillReturnError(datatest.ErrUnexpectedDB)
	mock.ExpectRollback()

	// Act
	err := repo.Create(t.Context(), product)

	// Assert
	assert.ErrorIs(t, err, datatest.ErrUnexpectedDB)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestProductRepo_CreateSuccess(t *testing.T) {
	t.Parallel()

	// Arrange
	db, mock := newMockDB(t)
	repo := repository.NewProductRepository(db)

	product := &entity.Product{
		Name:  "Test Product",
		Qty:   10,
		Price: 99.99,
	}

	mock.ExpectBegin()
	mock.ExpectQuery(`INSERT INTO "products"`).
		WithArgs(
			product.Name,
			product.Qty,
			product.Price,
			sqlmock.AnyArg(),
		).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(datatest.FakeProductID))
	mock.ExpectCommit()

	// Act
	err := repo.Create(context.Background(), product)

	// Assert
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}
