package bls

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBLSSigningAndVerification(t *testing.T) {
	// generate private keys
	fooPrivateKey := PrivateKeyGenerate()
	barPrivateKey := PrivateKeyGenerate()

	// get the public keys for the private keys
	fooPublicKey := PrivateKeyPublicKey(fooPrivateKey)
	barPublicKey := PrivateKeyPublicKey(barPrivateKey)

	// make messages to sign with the keys
	fooMessage := Message("hello foo")
	barMessage := Message("hello bar!")

	// calculate the digests of the messages
	fooDigest := Hash(fooMessage)
	barDigest := Hash(barMessage)

	// get the signature when signing the messages with the private keys
	fooSignature := PrivateKeySign(fooPrivateKey, fooMessage)
	barSignature := PrivateKeySign(barPrivateKey, barMessage)

	// assert the foo message was signed with the foo key
	assert.True(t, Verify(fooSignature, []Digest{fooDigest}, []PublicKey{fooPublicKey}))

	// assert the bar message was signed with the bar key
	assert.True(t, Verify(barSignature, []Digest{barDigest}, []PublicKey{barPublicKey}))

	// assert the foo message was not signed by the bar key
	assert.False(t, Verify(fooSignature, []Digest{fooDigest}, []PublicKey{barPublicKey}))

	// assert the bar/foo message was not signed by the foo/bar key
	assert.False(t, Verify(barSignature, []Digest{barDigest}, []PublicKey{fooPublicKey}))
	assert.False(t, Verify(barSignature, []Digest{fooDigest}, []PublicKey{barPublicKey}))
	assert.False(t, Verify(fooSignature, []Digest{barDigest}, []PublicKey{fooPublicKey}))
}

func BenchmarkBLSVerify(b *testing.B) {
	priv := PrivateKeyGenerate()

	msg := Message("this is a message that i will be signing")
	digest := Hash(msg)

	sig := PrivateKeySign(priv, msg)
	pubk := PrivateKeyPublicKey(priv)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if !Verify(sig, []Digest{digest}, []PublicKey{pubk}) {
			b.Fatal("failed to verify")
		}
	}
}
