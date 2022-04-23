// SPDX-License-Identifier: MIT
package web3lib

import (
	"encoding/hex"
	"fmt"
	"github.com/sonsofcrypto/bip39"
	keyStore "github.com/sonsofcrypto/keyStore"
	_ "golang.org/x/mobile/bind"
	"log"
)

// SomeType is for testing
type SomeType struct {
	foo string
}

// NewSomeType create new instance
func NewSomeType() *SomeType {
	return &SomeType{
		foo: "SomeType",
	}
}

// NewKeyStore create new default keyStore
func NewKeyStore() *keyStore.KeyStore {
	return keyStore.NewKeyStore(
		[]keyStore.Backend{keyStore.NewDiskBackEnd("~/.soc/keystore")},
	)
}

//HelloWorld export Hello world
func HelloWorld() {
	log.Println("Hello World")
}

// PrintMeAString export print me as string
func PrintMeAString(str string) {
	log.Println("Printing a string:", str)
}

func NewWrapperMnemonic(wordListPath string) string {
	// Load word list
	wordList, _ := bip39.LoadWordList(bip39.ListLangEn, wordListPath)
	// Generate bip39 compliant entropy
	entropy, _ := bip39.NewEntropy(bip39.EntropySize128)
	// Generate mnemonic from entropy
	mnemonic, _ := bip39.NewMnemonic(entropy, wordList)
	// Generate seed from mnemonic
	seed := bip39.NewSeed(mnemonic, "Secret Passphrase")

	fmt.Println("Hex:", hex.EncodeToString(seed))

	return mnemonic
}

//func main() {
//	// Load word list
//	wordList, _ := bip39.LoadWordList(bip39.ListLangEn, "wordLists")
//	// Generate bip39 compliant entropy
//	entropy, _ := bip39.NewEntropy(bip39.EntropySize128)
//	// Generate mnemonic from entropy
//	mnemonic, _ := bip39.NewMnemonic(entropy, wordList)
//	// Generate seed from mnemonic
//	seed := bip39.NewSeed(mnemonic, "Secret Passphrase")
//
//	// Display mnemonic and keys
//	fmt.Println("Mnemonic:", mnemonic)
//	fmt.Println("Seed:", seed)
//	fmt.Println("Hex:", hexutil.Encode(seed))
//	fmt.Println("Int:", big.NewInt(0).SetBytes(seed[:]))
//}
