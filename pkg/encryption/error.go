package encryption

import "errors"

var (
	CIPHER_TEXT_IMPROPER_SIZE = errors.New("cipher text is not of the right size.")
)