package pgstorage

import (
	pgxpool "github.com/jackc/pgx/v5/pgxpool"
	logging "github.com/kthucydi/bs_go_logrus"
)

var Log = &logging.Log

type DB struct {
	dataConnect map[string]string
	Database    *pgxpool.Pool
}

type PostGreStorage struct {
	DataBase DB
}

var PGDB = DB{}
