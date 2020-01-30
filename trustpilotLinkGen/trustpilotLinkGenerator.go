package trustpilotLinkGen

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/url"
)

type TrustpilotUserData struct {
	Email string   `json:"email"`
	Name  string   `json:"name"`
	Ref   string   `json:"ref"`
	Skus  []string `json:"skus,omitempty"`
	Tags  []string `json:"tags,omitempty"`
}

const (
	trustpilotFullUrl = "https://www.trustpilot.com/evaluate-bgl/"
)

var (
	errorDecoding       = errors.New("DECODING_ERROR")
	errorMarshalling    = errors.New("MARSHALLING_ERROR")
	errorLinkGenerating = errors.New("ERROR_GENERATING_TRUSTPILOT_URL")
)

type TrustpilotLinkGenerator interface {
	GenerateBusinessLink(userData interface{}) (string, error)
}

type trustpilotLinkGenerator struct {
	encrKeyByte []byte
	authKeyByte []byte
	domain      string
}

func NewTrustpilotLinkGenerator(encrKey, authKey, domain string) (TrustpilotLinkGenerator, error) {
	//Step 1
	encrKeyByte, err := base64.StdEncoding.DecodeString(encrKey)
	if err != nil {
		return nil, errorDecoding
	}
	authKeyByte, err := base64.StdEncoding.DecodeString(authKey)
	if err != nil {
		return nil, errorDecoding
	}
	return &trustpilotLinkGenerator{
		encrKeyByte: encrKeyByte,
		authKeyByte: authKeyByte,
		domain:      domain,
	}, nil
}

func (t *trustpilotLinkGenerator) GenerateBusinessLink(userData interface{}) (string, error) {
	//Step 2
	//Create JSON
	payload, err := json.Marshal(userData)
	if err != nil {
		return "", errorMarshalling
	}
	fmt.Println(string(payload))

	//PKCS7 padding mode
	block, err := aes.NewCipher(t.encrKeyByte)
	if err != nil {
		return "", errorLinkGenerating
	}
	payload = pkcs7(payload, block.BlockSize())

	//Step 3
	//Generating of ivCipher & IV at the beginning of it
	//ivCipher = IV + encoded payload
	ivCipher := make([]byte, aes.BlockSize+len(payload))
	iv := ivCipher[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ivCipher[aes.BlockSize:], payload)

	//Creating of HMAC-SHA256 with authKey
	h := hmac.New(sha256.New, t.authKeyByte)

	//Hashing IV + Ciphertext
	h.Write(ivCipher)

	//Step 4
	//base64 encoding
	//Creating of []byte IV + Ciphertext + HMAC
	ivCipherHmac := make([]byte, len(h.Sum(nil))+len(ivCipher))
	for i, v := range ivCipher {
		ivCipherHmac[i] = v
	}
	for i, v := range h.Sum(nil) {
		ivCipherHmac[i+len(ivCipher)] = v
	}

	//base64 encoding of IV + Ciphertext + HMAC
	base64tot := base64.StdEncoding.EncodeToString(ivCipherHmac)

	//Step 5
	//url construction
	finalUrl, err := url.Parse(trustpilotFullUrl + t.domain)
	if err != nil {
		return "", errorLinkGenerating
	}
	params := url.Values{}
	params.Add("p", base64tot)
	finalUrl.RawQuery = params.Encode()

	return finalUrl.String(), nil
}

func pkcs7(data []byte, blockSize int) []byte {
	neededBytes := blockSize - (len(data))%blockSize
	return append(data, bytes.Repeat([]byte{byte(neededBytes)}, neededBytes)...)
}
