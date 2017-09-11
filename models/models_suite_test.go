package models_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/treacher/step-warrior-api/database"
	"gopkg.in/pg.v5"
)

func TestModels(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Models Suite")
}

var TestDatabase *pg.DB

var _ = BeforeSuite(func() {
	TestDatabase = database.NewDatabaseConnection("test")
})

var _ = AfterSuite(func() {
	TestDatabase.Close()
})
