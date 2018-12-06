package ai_util

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/binary"
	"strings"
)

/** ServerId subfield offset in the initialization vector. */
// const INITVSERVERIDOFFSET = 8

/** Integrity signature size. */
// const SIGNATURESIZE = 4

// The following sizes are all in bytes.
const kInitializationVectorSize = 16
const kCiphertextSize = 8
const kSignatureSize = 4
const kEncryptedValueSize = kInitializationVectorSize + kCiphertextSize + kSignatureSize
const kKeySize = 32        // size of SHA-1 HMAC keys.
const kHashOutputSize = 20 // size of SHA-1 hash output.
const kBlockSize = 20      // This is a block cipher with fixed block size.

// hmac ,use sha1
func myhmac(key string, val string) []byte {
	mac := hmac.New(sha1.New, []byte(key))
	mac.Write([]byte(val))
	return mac.Sum(nil)
}

// Price Decoding
func DecodingPrice(kB64EncodedValue string, encryptionKey, integrityKey []byte) (flag bool, res int64) {
	//补全字符串
	if len(kB64EncodedValue)%4 == 3 {
		kB64EncodedValue += "="
	} else if len(kB64EncodedValue)%4 == 2 {
		kB64EncodedValue += "=="
	}

	//替换字符串中的特殊字符
	kB64EncodedValue = strings.Replace(kB64EncodedValue, "-", "/", -1)
	kB64EncodedValue = strings.Replace(kB64EncodedValue, "_", "/", -1)

	// Step 1. find the length of initialization vector and clear text.
	encryptedValue, _ := base64.StdEncoding.DecodeString(kB64EncodedValue)
	if len(encryptedValue) != kEncryptedValueSize {
		return false, 0
	}

	ciphertext := string(encryptedValue)
	cleartextLength := len(ciphertext) - kInitializationVectorSize - kSignatureSize
	if cleartextLength < 0 {
		// The length can't be correct.
		return false, 0
	}
	iv := []byte(ciphertext)[0:kInitializationVectorSize]

	// Step 2. recover clear text
	tmpbyte := []byte(ciphertext)
	ciphertextBegin := len(iv)
	ciphertextEnd := ciphertextBegin + cleartextLength
	cleartext := make([]byte, cleartextLength)
	var addIvCounterByte = true
	for tmpIdx := 0; ciphertextBegin+tmpIdx < ciphertextEnd; {
		encryptionPad := make([]byte, kHashOutputSize)
		encryptionPad = myhmac(string(encryptionKey[0:kKeySize]), string(iv))

		for i := 0; i < kBlockSize && ciphertextBegin+i < ciphertextEnd; i++ {
			cleartext[i] = tmpbyte[ciphertextBegin+tmpIdx] ^ encryptionPad[i]
			tmpIdx++
		}

		if !addIvCounterByte {
			lastByte := kInitializationVectorSize
			lastByte--
			if iv[lastByte] == 0 {
				addIvCounterByte = true
			}
		}

		if addIvCounterByte {
			addIvCounterByte = false
			iv = append(iv, 0)
		}
	}

	// Step 3. Compute integrity hash. The input to the HMAC is cleartext
	// followed by initialization vector, which is stored in the 1st section of
	// ciphertext.
	inputMessage := make([]byte, len(cleartext))
	copy(inputMessage, cleartext)
	for i := 0; i < kInitializationVectorSize; i++ {
		inputMessage = append(inputMessage, ciphertext[i])
	}
	result := int64(binary.BigEndian.Uint64(cleartext))
	integrityHash := myhmac(string(integrityKey[0:kKeySize]), string(inputMessage))
	if ciphertext[ciphertextEnd:ciphertextEnd+kSignatureSize] == string(integrityHash)[0:kSignatureSize] {
		return true, result
	} else {
		return false, 0
	}
}
