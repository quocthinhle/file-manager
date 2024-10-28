package pgoutadapter_test

import (
	"context"
	"embed"
	"errors"
	"fmt"
	postgresdb "github.com/quocthinhle/file-manager-api/internal/database/postgres"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
	"testing"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/source/iofs"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/lib/pq"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestPg(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Pg Suite")
}

const (
	dbName = "file_manager"
	dbUser = "user"
	dbPass = "password"
)

var postgresContainer *postgres.PostgresContainer
var dbURL string

func migrateUp(fs embed.FS) error {
	GinkgoHelper()
	source, err := iofs.New(fs, "migrations")
	port, err := postgresContainer.MappedPort(context.Background(), "5432")
	Expect(err).NotTo(HaveOccurred())

	dbURL = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		dbUser,
		dbPass,
		"127.0.0.1",
		port.Port(),
		dbName,
	)
	Expect(err).NotTo(HaveOccurred())
	m, err := migrate.NewWithSourceInstance("iofs", source, dbURL)
	if err != nil {
		return err
	}

	defer func(m *migrate.Migrate) {
		err, _ := m.Close()
		if err != nil {
			fmt.Println("Error closing migration", err)
		}
	}(m)

	if err = m.Up(); err != nil {
		if !errors.Is(err, migrate.ErrNoChange) {
			return err
		}
	}

	return nil
}

var _ = BeforeSuite(func() {
	var err error
	postgresContainer, err = postgres.Run(context.Background(), "postgres:16-alpine",
		postgres.WithDatabase(dbName),
		postgres.WithUsername(dbUser),
		postgres.WithPassword(dbPass),
		testcontainers.WithWaitStrategy(
			wait.ForExposedPort(),
		),
	)
	Expect(err).NotTo(HaveOccurred())
	Expect(migrateUp(postgresdb.FS)).NotTo(HaveOccurred())
})
