package utils

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

type Postgres struct {
	Host        string
	Port        string
	User        string
	Password    string
	DBName      string
	MaxOpenConn int
	MaxIdleConn int
	MaxIdleTime time.Duration
	SSLMode     string
}

type Option func(*Postgres)

func WithHost(host string) Option {
	return func(p *Postgres) {
		p.Host = host
	}
}

func WithPort(port string) Option {
	return func(p *Postgres) {
		p.Port = port
	}
}

func WithUser(user string) Option {
	return func(p *Postgres) {
		p.User = user
	}
}

func WithPassword(password string) Option {
	return func(p *Postgres) {
		p.Password = password
	}
}

func WithDBName(dbName string) Option {
	return func(p *Postgres) {
		p.DBName = dbName
	}
}

func WithMaxOpenConn(maxOpenConn int) Option {
	return func(p *Postgres) {
		p.MaxOpenConn = maxOpenConn
	}
}

func WithMaxIdleConn(maxIdleConn int) Option {
	return func(p *Postgres) {
		p.MaxIdleConn = maxIdleConn
	}
}

func WithMaxIdleTime(maxIdleTime time.Duration) Option {
	return func(p *Postgres) {
		p.MaxIdleTime = maxIdleTime
	}
}

func WithSSLMode(mode string) Option {
	return func(p *Postgres) {
		p.SSLMode = mode
	}
}

func (p *Postgres) URI() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", p.Host, p.Port, p.User, p.Password, p.DBName, p.SSLMode)
}

func (p *Postgres) Connect() (*sql.DB, error) {
	db, err := sql.Open("postgres", p.URI())
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(p.MaxOpenConn)
	db.SetMaxIdleConns(p.MaxIdleConn)
	db.SetConnMaxLifetime(p.MaxIdleTime)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		db.Close()
		return nil, err
	}
	return db, nil
}

func NewPostgres(opts ...Option) *Postgres {
	p := &Postgres{}
	for _, opt := range opts {
		opt(p)
	}
	return p
}
