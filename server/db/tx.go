package db

import (
	"github.com/bmatsuo/lmdb-go/lmdb"
	"github.com/pkg/errors"
)

type Tx struct {
	DB  lmdb.DBI
	Env *lmdb.Env
	Tx  *lmdb.Txn
}

func (tx *Tx) Get(key []byte) (data []byte) {
	var err error
	if data, err = tx.Tx.Get(tx.DB, key); err != nil {
		if lmdb.IsNotFound(err) {
			return nil
		}

		panic(errors.Wrap(err, "Tx.Get"))
	}
	return data
}


func (tx *Tx) GetOwner() *DB{
	return &DB{
		Env:tx.Env,
		DBI:tx.DB,
	}
}






func (tx *Tx) Commit() error {
	return tx.Tx.Commit()
}

func (tx *Tx) MustCommit() {
	err := tx.Tx.Commit()
	if err != nil{
		panic(err)
	}
}

func (tx *Tx) Put(key []byte, val []byte)  {
	if err := tx.Tx.Put(tx.DB, key, val, 0); err != nil {
		panic(errors.Wrap(err, "tx.Put"))
	}

}

func (tx *Tx) Del(key []byte) error {
	if err := tx.Tx.Del(tx.DB, key, nil); err != nil {
		return err

	}
	return nil
}

func (tx *Tx) PutReserve(key []byte, size int) ([]byte, error) {
	return tx.Tx.PutReserve(tx.DB, key, size, 0)
}

func (tx *Tx) Close() (err error) {
	return nil
}


func (tx *Tx) MustAbort(){
	tx.Tx.Abort()
}

func (tx *Tx) MustClose(){

	err := tx.Close()
	if err != nil {
		panic(err)
	}

}


