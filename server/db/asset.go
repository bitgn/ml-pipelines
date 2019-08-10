package db

import "bytes"

type AssetFactory func() []byte

// GetAsset fetches cached asset from the DB or generates and caches it otherwise
func GetAsset(tx *Tx, key string, digest []byte, getter AssetFactory) []byte {

	cache := &AssetCache{}
	tx.GetProto(CreateKey(Range_ASSET_CACHE, key), cache)

	if bytes.Compare(digest, cache.Digest) == 0 {
		return cache.Body
	}

	cache.Digest = digest
	buf := getter()
	cache.Body = buf

	writer := tx.GetOwner().MustWrite()
	defer writer.MustCommit()

	writer.PutProto(CreateKey(Range_ASSET_CACHE, key), cache)
	return buf
}