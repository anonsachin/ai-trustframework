package lamportsig

import (
	"bytes"
	"crypto/rand"
	"crypto/sha256"
	"errors"
	"fmt"
	"math"
	"strconv"
)

type LamportSignature struct{
	Signature []LamportNumber
	PublicKey LamportList
}

// Verify validates the lengths of all the elements
func (l *LamportSignature) VerifyElements() bool {

	// Verifying the signature length
	if len(l.Signature) != LamportHashLength {
		return false
	}

	// Verifing public key lengths

	if len(l.PublicKey[0]) != LamportHashLength {
		return false
	}

	if len(l.PublicKey[1]) != LamportHashLength {
		return false
	}

	// all the elements are the right sizes
	return true
}

const LamportHashLength = 256
type LamportNumber [32]byte
type LamportList [2][]LamportNumber

// LamportSig takes the message hash and the private key and generates
// a signature correlating the message and private key
func LamportSig(messageHassh [32]byte, PrivateKeys []LamportNumber) ([]LamportNumber, error) {
	// length check
	if len(PrivateKeys) != LamportHashLength * 2 {
		return nil, errors.New("There are not an exact 512 numbers.")
	}

	// iterate over the message
	return selectKeys(messageHassh,PrivateKeys,0), nil
}

// select the private keys for sig
func selectKeys(messageHash [32]byte, PrivateKeys []LamportNumber, position int) []LamportNumber {
	// if it reaches the last position in this case
	// 255 or 256th bit of the hash return the last
	// selected key
	if position == LamportHashLength - 1 {
		if BitPresence(messageHash[int(math.Floor(float64(position)/8.0))], position % 8) {
			return []LamportNumber{PrivateKeys[256 + position]}
		} else {
			return []LamportNumber{PrivateKeys[position]}
		}
	} else {
		// based on the bit in the current position select the key
		// the append the next key for the next position into the 
		// array after the current key- 
		// 
		// should look like [priv_key_1, priv_key_2, .... , priv_key_256]
		if BitPresence(messageHash[int(math.Floor(float64(position)/8.0))], position % 8) {
			return append([]LamportNumber{PrivateKeys[256 + position]}, selectKeys(messageHash,PrivateKeys,position +1)...)
		} else {
			return append([]LamportNumber{PrivateKeys[position]}, selectKeys(messageHash,PrivateKeys,position +1)...)
		}
	}
}

func BinaryRep(data [32]byte, position int) string {
	if position == 31 {
		return strconv.FormatUint(uint64(data[position]),2)
	} else {
		return fmt.Sprintf("%s:%s",strconv.FormatUint(uint64(data[position]),2),BinaryRep(data,position+1))
	}
}


// BitPresence checks if a bit is one in a paricular location.
func BitPresence(data byte, position int) bool {
	one := 1
	if position > 8 {
		return false
	}

	// left shifting 00000001 to a particular position
	// ex :- position = 2; 00000001 << 2 = 00000100
	mask := one << position

	// Checking if bit in a position is present of not
	// by anding with a mask : data & mask (00000100) == mask
	if int(data&byte(mask)) == mask {
		return true
	}

	return false
}

// GenerateLamportKey creates 512 random numbers
func GenerateLamportKey() ([]LamportNumber,error) {
	Keys := make([]LamportNumber, 0)
	for i := 0; i< LamportHashLength * 2 ; i++ {
		key, err := GenerateRandom256bitNumber()
		if err != nil {
			return nil, err
		}
		Keys = append(Keys, key)
	}

	return Keys, nil
}

// LamportPublicKeyFromPrivateKey takes the private keys and 
// returns the hashes of the private key which is the public key
func LamportPublicKeyFromPrivateKey(PrivateKeys []LamportNumber) (*LamportList,error) {
	// length check
	if len(PrivateKeys) != LamportHashLength * 2 {
		return nil, errors.New("There are not an exact 512 numbers.")
	}

	Keys1 := make([]LamportNumber, 0)
	Keys2 := make([]LamportNumber, 0)
	for i := 0; i< LamportHashLength; i++ {
		key1 := sha256.Sum256(PrivateKeys[i][:])
		Keys1 = append(Keys1, key1)
		key2 := sha256.Sum256(PrivateKeys[i+256][:])
		Keys2 = append(Keys2, key2)
	}

	return &LamportList{Keys1,Keys2}, nil
}

// GenerateRandom256bitNumber generates a single random number of length of 256 bits
func GenerateRandom256bitNumber() (LamportNumber, error){
	r := [32]byte{}
	_, err := rand.Read(r[:])

	if err != nil {
		return [32]byte{},err
	}

	return r, nil
}

// VerifySignature takes the hash of the message and then
// validates if the signature matches the public keys for this 
// message
func (l *LamportSignature) VerifySignature(messageHash [32]byte) bool {
	return l.VerifySignatureRecurcive(messageHash,0)
}

// VerifySignatureRecurcive is the recurcive implementation of the verification algorithm.
func (l *LamportSignature) VerifySignatureRecurcive(messageHash [32]byte, position int) bool {
	// if it reaches the last position in this case
	// 255 or 256th bit of the hash return the last
	// selected key
	if position == LamportHashLength - 1 {
		// Hashing the signature value at position
		hashValue := sha256.Sum256(l.Signature[position][:])
		// Checking if it is 0 or 1 at the position
		if BitPresence(messageHash[int(math.Floor(float64(position)/8.0))], position % 8) {
			// checking equality
			return  bytes.Compare(hashValue[:],l.PublicKey[1][position][:]) == 0
		} else {
			// checking equality
			return bytes.Compare(hashValue[:],l.PublicKey[0][position][:]) == 0
		}
	} else {
		// Hashing the signature value at position
		hashValue := sha256.Sum256(l.Signature[position][:])
		// Checking if it is 0 or 1 at the position
		if BitPresence(messageHash[int(math.Floor(float64(position)/8.0))], position % 8) {
			// checking equality
			if  bytes.Compare(hashValue[:],l.PublicKey[1][position][:]) == 0 {
				return l.VerifySignatureRecurcive(messageHash,position+1)
			} else {
				return false
			}
		} else {
			// checking equality
			if  bytes.Compare(hashValue[:],l.PublicKey[0][position][:]) == 0 {
				return l.VerifySignatureRecurcive(messageHash,position+1)
			} else {
				return false
			}
		}

	}
}