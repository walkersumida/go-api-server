package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pkg/errors"
)

type Database struct {
	Cli *sql.DB
}

type Configs struct {
	Host     string
	Name     string
	Password string
	Port     string
	SSLMode  string
	User     string
}

func NewClient(configs Configs) (*Database, error) {
	url := BuildDBURL(configs)

	db, err := sql.Open("pgx", url)
	if err != nil {
		return nil, errors.Wrap(err, "failed to connect a database")
	}

	return &Database{
		Cli: db,
	}, nil
}

func BuildDBURL(configs Configs) string {
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbPort := os.Getenv("DB_PORT")
	dbSSLMode := os.Getenv("DB_SSL_MODE")
	dbUser := os.Getenv("DB_USER")

	if configs.Host != "" {
		dbHost = configs.Host
	}
	if configs.Name != "" {
		dbName = configs.Name
	}
	if configs.Password != "" {
		dbPassword = configs.Password
	}
	if configs.Port != "" {
		dbPort = configs.Port
	}
	if configs.SSLMode != "" {
		dbSSLMode = configs.SSLMode
	}
	if configs.User != "" {
		dbUser = configs.User
	}

	url := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=%s", dbUser, dbPassword, dbHost, dbPort, dbName, dbSSLMode)

	return url
}

func (d *Database) Create(name string) error {
	q := fmt.Sprintf("CREATE DATABASE %s ENCODING = 'UTF8' template=template0", name)
	if _, err := d.Cli.Exec(q); err != nil {
		return errors.Wrapf(err, "failed to create database %s", name)
	}

	q = fmt.Sprintf("ALTER DATABASE %s SET timezone TO 'Etc/UTC'", name)
	if _, err := d.Cli.Exec(q); err != nil {
		return errors.Wrapf(err, "failed to alter a database: %s", name)
	}

	return nil
}

func (d *Database) Drop(name string) error {
	q := fmt.Sprintf("DROP DATABASE IF EXISTS %s", name)
	if _, err := d.Cli.Exec(q); err != nil {
		return errors.Wrapf(err, "failed to drop database %s", name)
	}

	return nil
}
