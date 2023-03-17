package auth

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

// Generates a random hex string of n bytes
func Nonce(n int) (string, error) {
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func VerifySignature(addr string, sig string, nonce string) error {
	// Decode to hex and truncate recovery offset bytes
	decodedSig := hexutil.MustDecode(sig)
	decodedSig[crypto.RecoveryIDOffset] -= 27

	// Hash nonce
	hash := accounts.TextHash([]byte(nonce))

	// Verify signature
	sigPublicKeyECDSA, err := crypto.SigToPub(hash, decodedSig)
	if err != nil {
		return fmt.Errorf("Bad signature")
	}

	derivedAddr := crypto.PubkeyToAddress(*sigPublicKeyECDSA).Hex()
	if derivedAddr != addr {
		return fmt.Errorf("Invalid address & signature combination")
	}

	return nil
}
