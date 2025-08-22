package main

import (
	_ "fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	_ "github.com/stretchr/testify/assert"
	"regexp"
	"testing"
	_ "testing"
)

func TestRetrieveSchrodingerDataRandomKey(t *testing.T) {
	mockRandomGenerator := &MockRandomNumbersGenerator{}
	mockRandomGenerator.value = 0.2
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Errorf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	rows := sqlmock.NewRows([]string{"value"}).AddRow("banana")
	mock.ExpectQuery(regexp.QuoteMeta(SelectByRandomKeySQL)).WillReturnRows(rows)
	err, _ = retrieveSchrodingerData(db, mockRandomGenerator)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestRetrieveSchrodingerDataExpectedKey(t *testing.T) {
	mockRandomGenerator := &MockRandomNumbersGenerator{}
	mockRandomGenerator.value = 0.4
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Errorf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	rows := sqlmock.NewRows([]string{"key", "value"}).AddRow("apple", "banana")
	mock.ExpectQuery(regexp.QuoteMeta(SelectByKeySQL)).WithArgs("apple").WillReturnRows(rows)
	err, _ = retrieveSchrodingerData(db, mockRandomGenerator, "apple")
	assert.NoError(t, mock.ExpectationsWereMet())
}
