package db

import (
	"github.com/pkg/errors"
	"github.com/bmatsuo/lmdb-go/lmdb"
	"log"
	"os"
)

type Config struct {
	EnvFlags uint
	SizeMbs  int64
}

func NewConfig() *Config {

	return &Config{
		EnvFlags: lmdb.NoSync,
		SizeMbs:  1024,

	}

}

type DB struct {
	Env *lmdb.Env
	DBI lmdb.DBI
}


const (
	max_dbs = 2
	db_mode = 0644
)



// New creates a new DB wrapper around LMDB
func New(folder string, cfg *Config) (*DB, error) {
	env, err := lmdb.NewEnv()

	if err != nil {
		return nil, errors.Wrap(err, "env create failed")
	}

	err = env.SetMaxDBs(max_dbs)
	if err != nil {
		return nil, errors.Wrap(err, "env config failed")
	}
	err = env.SetMapSize(cfg.SizeMbs * 1024 * 1024)
	if err != nil {
		return nil, errors.Wrap(err, "map size failed")
	}

	if err = env.SetFlags(cfg.EnvFlags); err != nil {
		return nil, errors.Wrap(err, "set flag")
	}

	os.MkdirAll(folder, os.ModePerm)
	err = env.Open(folder, 0, db_mode)
	if err != nil {
		return nil, errors.Wrap(err, "open env")
	}

	var staleReaders int
	if staleReaders, err = env.ReaderCheck(); err != nil {
		return nil, errors.Wrap(err, "reader check")
	}
	if staleReaders > 0 {
		log.Printf("cleared %d reader slots from dead processes", staleReaders)
	}

	var dbi lmdb.DBI
	err = env.Update(func(txn *lmdb.Txn) (err error) {
		dbi, err = txn.CreateDBI("agg")
		return err
	})
	if err != nil {
		return nil, errors.Wrap(err, "create DB")
	}

	return &DB{env, dbi}, nil
}

func (db *DB) Read() (tx *Tx, err error) {
	return db.CreateTransaction(lmdb.Readonly)
}

func (db *DB) MustRead() *Tx{
	tx, err := db.CreateTransaction(lmdb.Readonly)
	if err != nil {
		panic(err)
	}
	return tx
}




func (db *DB) MustWrite() *Tx{
	tx, err := db.CreateTransaction(0)

	if err != nil{
		panic(err)
	}
	return tx
}

func (db *DB) BeginWrite() (tx *Tx, err error) {
	return db.CreateTransaction(0)
}

func (db *DB) CreateTransaction(flags uint) (tx *Tx, err error) {

	var txn *lmdb.Txn
	if txn, err = db.Env.BeginTxn(nil, flags); err != nil {
		return nil, errors.Wrap(err, "BeginTxn(flags)")
	}

	return &Tx{db.DBI, db.Env, txn}, nil
}

// Close the environment
func (db *DB) Close() error {

	if db.Env == nil {
		return nil
	}

	//db.Env.CloseDBI(db.DBI)
	err := db.Env.Close()
	db.Env = nil
	return errors.Wrap(err, "Env.Close")
}