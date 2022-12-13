package merkele

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
)

type Merkele struct {
	Hash      [32]byte `json:"-"`
	Signature string `json:"hash"`
	Left      *Merkele
	Right     *Merkele
}

func CreateFullTree(hashes [][32]byte) (*Merkele, error) {
	left, right, hash, sig, err := CreateTree(hashes)
	if err != nil {
		return nil, err
	}

	return &Merkele{
		Left:      left,
		Right:     right,
		Hash:      hash,
		Signature: sig,
	},nil
}

func CreateTree(hashes [][32]byte) (*Merkele, *Merkele, [32]byte, string, error) {
	totalLength := len(hashes)
	// enforcing binary tree structure
	if totalLength %2 != 0 {
		return nil, nil, [32]byte{}, "", errors.New("The size of the hashes its not even.")
	}
	if len(hashes) == 2 {
		// Concatinating the two hash values.
		concat := make([]byte, 0)
		concat = append(concat, hashes[0][:]...)
		concat = append(concat, hashes[1][:]...)
		hash := sha256.Sum256(concat)

		return &Merkele{
				Hash:      hashes[0],
				Signature: hex.EncodeToString(hashes[0][:]),
			}, &Merkele{
				Hash:      hashes[1],
				Signature: hex.EncodeToString(hashes[1][:]),
			},
			hash,
			hex.EncodeToString(hash[:]),
			nil
	} else if len(hashes) > 0 {
		//length calculations
		length := len(hashes)/2
		// first half of the array
		left, right, hash, sig, err := CreateTree(hashes[:length])
		// unexpected error handling
		if err != nil {
			return nil, nil, [32]byte{}, "", err
		}
		// initialization
		Left := &Merkele{
			Hash:      hash,
			Signature: sig,
			Left:      left,
			Right:     right,
		}
		// second half
		left, right, hash, sig, err = CreateTree(hashes[length:])
		// unexpected error handling
		if err != nil {
			return nil, nil, [32]byte{}, "", err
		}
		// initialization
		Right := &Merkele{
			Hash:      hash,
			Signature: sig,
			Left:      left,
			Right:     right,
		}
		// its hash
		concat := make([]byte, 0)
		concat = append(concat, Left.Hash[:]...)
		concat = append(concat, Right.Hash[:]...)
		hash = sha256.Sum256(concat)
		return Left, Right, hash, hex.EncodeToString(hash[:]), nil
	}

	// Unknow path of code flow error.
	return nil, nil, [32]byte{}, "", errors.New("Unable to create tree")
}

func CalculateRoot(hashes [][32]byte) ([]byte,error) {
	if len(hashes) % 2 == 0 {
		return nil, errors.New("length of hashes is not even cannot calculate root.")
	} else {
		hash := calculateRoot(hashes)
		return hash[:], nil
	}
}

func calculateRoot(hashes [][32]byte) [32]byte {
	if len(hashes) == 2 {
		concat := make([]byte, 0)
		concat = append(concat, hashes[0][:]...)
		concat = append(concat, hashes[1][:]...)
		return sha256.Sum256(concat)
	} else {
		concat := make([]byte, 0)
		concat = append(concat, hashes[0][:]...)
		concat = append(concat, hashes[1][:]...)
		hashValue := sha256.Sum256(concat)
		hashes = append(hashes[2:], hashValue)
		return calculateRoot(hashes)
	}
}