package db

import (
	"bytes"
	"github.com/bmatsuo/lmdb-go/lmdb"
	"github.com/bmatsuo/lmdb-go/lmdbscan"
	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"

	"github.com/abdullin/lex-go/tuple"
)

func CreateKey(r Range, args ...tuple.Element) []byte {

	var slice []tuple.Element
	slice = append(slice, int(r))
	slice = append(slice, args...)


	tpl := tuple.Tuple(slice)
	return tpl.Pack()
}

func (tx *Tx) PutProto(key []byte, val proto.Message) {
	var err error
	var data []byte

	if data, err = proto.Marshal(val); err != nil {
		panic(errors.Wrap(err, "Marshal"))
	}
	tx.Put(key, data)
}


func (tx *Tx) GetProto(key []byte, pb proto.Message) bool {


	data := tx.Get(key)


	if data == nil {
		return false
	}

	if err := proto.Unmarshal(data, pb); err != nil {
		panic(errors.Wrap(err, "Unmarshal"))
	}
	return true
}

func (tx *Tx) GetNext(key []byte) (k, v []byte, err error) {
	scanner := lmdbscan.New(tx.Tx, tx.DB)
	defer scanner.Close()
	if !scanner.Set(key, nil, lmdb.SetRange) {
		err = lmdb.NotFound
		return
	}

	if !scanner.Scan() {
		err = lmdb.NotFound
		return
	}

	k = scanner.Key()
	v = scanner.Val()
	err = scanner.Err()
	return
}

func (tx *Tx) GetPrev(key []byte) (k, v []byte, err error) {

	scanner := lmdbscan.New(tx.Tx, tx.DB)
	defer scanner.Close()
	if !scanner.Set(key, nil, lmdb.SetRange) {
		err = lmdb.NotFound
		return
	}
	if !scanner.Set(nil, nil, lmdb.Prev) {
		err = lmdb.NotFound
		return
	}

	if !scanner.Scan() {
		err = lmdb.NotFound
		return
	}

	k = scanner.Key()
	v = scanner.Val()
	err = scanner.Err()
	return
}

func (tx *Tx) ScanRange(key []byte, row func(k, v []byte))  {
	scanner := lmdbscan.New(tx.Tx, tx.DB)
	defer scanner.Close()
	if !scanner.Set(key, nil, lmdb.SetRange) {
		return
	}

	for scanner.Scan() {
		if !bytes.HasPrefix(scanner.Key(), key) {
			break
		}
		row(scanner.Key(), scanner.Val())
	}

	if err := scanner.Err(); err != nil {
		panic(errors.Wrap(err, "Scanner					"))
	}
}
func (t *Tx) DelRange(key []byte) error {

	scanner := lmdbscan.New(t.Tx, t.DB)
	defer scanner.Close()

	if !scanner.Set(key, nil, lmdb.SetRange) {
		return nil
	}

	for scanner.Scan() {
		if !bytes.HasPrefix(scanner.Key(), key) {
			break
		}

		err := t.Tx.Del(t.DB, scanner.Key(), nil)
		if err != nil {
			return err
		}
	}

	return scanner.Err()
}
