package extract

import (
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
)

var errPEMDecode = errors.New("cannot decode PEM encoded block")

func PEM(b []byte) (encodedData string, err error) {
	pemBlock, _ := pem.Decode(b)
	if pemBlock == nil {
		return "", fmt.Errorf("%w", errPEMDecode)
	}

	der := pemBlock.Bytes
	encodedData = base64.StdEncoding.EncodeToString(der)
	return encodedData, nil
}
