package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/sonsofcrypto/bip39"
	"math/big"
)

func main() {
	// Load word list
	wordList, _ := bip39.LoadWordList(bip39.ListLangEn, "wordLists")
	// Generate bip39 compliant entropy
	entropy, _ := bip39.NewEntropy(bip39.EntropySize128)
	// Generate mnemonic from entropy
	mnemonic, _ := bip39.NewMnemonic(entropy, wordList)
	// Generate seed from mnemonic
	seed := bip39.NewSeed(mnemonic, "Secret Passphrase")

	// Display mnemonic and keys
	fmt.Println("Mnemonic:", mnemonic)
	fmt.Println("Seed:", seed)
	fmt.Println("Hex:", hexutil.Encode(seed))
	fmt.Println("Int:", big.NewInt(0).SetBytes(seed[:]))
}
