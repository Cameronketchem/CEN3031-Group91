package auth

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestVerifyJWT(t *testing.T) {
	secret := "0x123456789ABCDEF"
	for i := 0; i < 100; i++ {
		address, err := Nonce(20)
		require.NoError(t, err)

		token, err := NewJWT(secret, address, 999)
		require.NoError(t, err)

		derivedAddr, derivedFltId, err := VerifyJWT(secret, token)
		require.NoError(t, err)

		derivedId := int(derivedFltId)
		require.Equal(t, derivedId, 999)
		require.Equal(t, derivedAddr, address)
	}
}

func TestVerifySignature(t *testing.T) {
	addr := "0x15A855A37A27eaA9514D3952469806EC03a117eD"

	// Signatures generated via metamask
	nonce1 := "555c415f4b2c13e8be8c7c99b8c2546ed29a0db9"
	sig1 := "0xb2edfd07051326726bc6f313b4911e1608e9fddaef2f0affd78e0476b14ab2a0536e8ed7d3045eadcd5dac1c49d18187126ac2bcf42e505bd47d6039f4d242421b"

	nonce2 := "d1bd9993f9b4d49a468674c5ae9a1d9186478249"
	sig2 := "0x7f0af3a63d287a6f048858e33847d55b5a2c68529185eeeb221685864903814d1fe1aaa4929c0e243f204dadff3cbb5d00aadc815a641bcf6b5bcbddbca9452b1c"

	// Test Verify Signature
	err := VerifySignature(addr, sig1, nonce1)
	require.NoError(t, err)

	err = VerifySignature(addr, sig2, nonce2)
	require.NoError(t, err)
}
