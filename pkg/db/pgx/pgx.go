package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net"
	"os"
	"pronaces_back/pkg/domain"

	"cloud.google.com/go/cloudsqlconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/stdlib"
	migrate "github.com/rubenv/sql-migrate"
)

type psxStore struct {
	db *sql.DB
}

func NewPsxStore() (domain.FormDB, error) {

	dbUser := os.Getenv("DB_USER")                                  // e.g. 'my-db-user'
	dbPwd := os.Getenv("DB_PASS")                                   // e.g. 'my-db-password'
	dbName := os.Getenv("DB_NAME")                                  // e.g. 'my-database'
	instanceConnectionName := os.Getenv("INSTANCE_CONNECTION_NAME") // e.g. 'project:region:instance'
	usePrivate := os.Getenv("PRIVATE_IP")

	dsn := fmt.Sprintf("user=%s password=%s database=%s sslmode=disable",
		dbUser,
		dbPwd,
		dbName,
	)

	config, err := pgx.ParseConfig(dsn)
	if err != nil {
		return nil, err
	}

	var opts []cloudsqlconn.Option
	if usePrivate != "" {
		opts = append(opts, cloudsqlconn.WithDefaultDialOptions(cloudsqlconn.WithPrivateIP()))
	}

	d, err := cloudsqlconn.NewDialer(context.Background(), opts...)
	if err != nil {
		return nil, err
	}

	config.DialFunc = func(ctx context.Context, network, instance string) (net.Conn, error) {
		return d.Dial(ctx, instanceConnectionName)
	}

	dbURI := stdlib.RegisterConnConfig(config)

	db, err := sql.Open("pgx", dbURI)
	if err != nil {
		log.Fatalf("Error on sql.Open: %v", err)
	}

	fmt.Println("Successfully connected to database!")
	return &psxStore{db: db}, nil
}

func migrateDB(db *sql.DB) error {
	migrations := &migrate.FileMigrationSource{
		Dir: "migrations", // Carpeta donde están las migraciones SQL
	}

	_, err := migrate.Exec(db, "postgres", migrations, migrate.Up)
	if err != nil {
		return err
	}

	return nil
}
