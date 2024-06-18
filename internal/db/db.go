package db

import (
    "context"
    "github.com/jackc/pgx/v4"
    "log"
)

func Connect() (*pgx.Conn, error) {
    conn, err := pgx.Connect(context.Background(), "postgres://postgres:adminadmin@db:5432/db")
    if err != nil {
        return nil, err
    }
    return conn, nil
}
