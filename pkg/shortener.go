package pkg

import (
	"crypto/sha256"
	"encoding/hex"
)

type UrlShortener struct {
	store map[string]string
}

func New() *UrlShortener {
	return &UrlShortener{store: make(map[string]string)}
}

func (us *UrlShortener) Encode(longUrl string) string {
	hash := sha256.Sum256([]byte(longUrl))
	shortUrl := hex.EncodeToString(hash[:3])  // берем первые 3 байта для короткого URL
	us.store[shortUrl] = longUrl
	return shortUrl
}

func (us *UrlShortener) Decode(shortUrl string) (string, bool) {
	longUrl, exists := us.store[shortUrl]
	return longUrl, exists
}
