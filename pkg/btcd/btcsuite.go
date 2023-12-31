package btcd

import (
	"fmt"
	"github.com/tyler-smith/go-bip39"
	"log"

	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/btcutil/hdkeychain"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/rpcclient"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcd/wire"
)

func CreateClient(rpcConfig *rpcclient.ConnConfig) *rpcclient.Client {
	client, err := rpcclient.New(rpcConfig, nil)
	if err != nil {
		log.Panicf("error creating new btcclient: %v", err)
	}
	return client
}

func CreateHDWallet(seed []byte, params *chaincfg.Params) (*hdkeychain.ExtendedKey, error) {
	return hdkeychain.NewMaster(seed, params)
}

func GetChildFromRoot(root *hdkeychain.ExtendedKey, index uint32) *hdkeychain.ExtendedKey {
	derive, err := root.Derive(index)
	if err != nil {
		log.Panicf("failed to derive %v", err)
	}
	return derive
}

func GetParams(mode string) *chaincfg.Params {
	var params *chaincfg.Params
	switch mode {
	case "LOCAL":
		params = GetTestNetworkParams()
	case "PROD":
		params = GetMainNetworkParams()
	default:
		return GetTestNetworkParams()
	}
	return params
}

func GetMasterKey(key string, params *chaincfg.Params) (*hdkeychain.ExtendedKey, error) {
	seed := bip39.NewSeed(key, "")
	masterKey, err := hdkeychain.NewMaster(seed, params)
	if err != nil {
		return nil, fmt.Errorf("error creating master key: %v", err)
	}
	return masterKey, nil
}

func GetMainNetworkParams() *chaincfg.Params {
	return &chaincfg.MainNetParams
}

func GetTestNetworkParams() *chaincfg.Params {
	return &chaincfg.TestNet3Params
}

func GetPrivateKey(extended *hdkeychain.ExtendedKey) (*btcec.PrivateKey, error) {
	return extended.ECPrivKey()
}

func GetPubKeyHash(privateKey *btcec.PrivateKey) []byte {
	return btcutil.Hash160(privateKey.PubKey().SerializeCompressed())
}

func GetRedeemScript(pubKeyHash []byte) ([]byte, error) {
	return txscript.NewScriptBuilder().AddOp(txscript.OP_0).AddData(pubKeyHash).Script()
}

func GetP2SHAddress(redeemScript []byte, params *chaincfg.Params) (*btcutil.AddressScriptHash, error) {
	return btcutil.NewAddressScriptHash(redeemScript, params)
}

func CreateP2SHAddress(extended *hdkeychain.ExtendedKey, params *chaincfg.Params) (*btcutil.AddressScriptHash, error) {
	privateKey, err := GetPrivateKey(extended)
	if err != nil {
		return nil, err
	}
	pubKeyHash := GetPubKeyHash(privateKey)
	redeemScript, err := GetRedeemScript(pubKeyHash)
	if err != nil {
		return nil, err
	}
	p2shAddr, err := GetP2SHAddress(redeemScript, params)
	if err != nil {
		return nil, err
	}
	return p2shAddr, nil
}

func GetWitness(privateKey *btcec.PrivateKey) []byte {
	return btcutil.Hash160(privateKey.PubKey().SerializeCompressed())
}

func CreateP2WPKHAddress(extended *hdkeychain.ExtendedKey, params *chaincfg.Params) (*btcutil.AddressWitnessPubKeyHash, error) {
	privateKey, err := GetPrivateKey(extended)
	if err != nil {
		return nil, err
	}
	witness := GetWitness(privateKey)
	return btcutil.NewAddressWitnessPubKeyHash(witness, params)
}

func CreateNewMessageTransaction() *wire.MsgTx {
	return wire.NewMsgTx(wire.TxVersion)
}

func NewHashFromString(hash string) (*chainhash.Hash, error) {
	return chainhash.NewHashFromStr(hash)
}

func AddInput(tx *wire.MsgTx, hash string) error {
	txHash, err := NewHashFromString(hash)
	if err != nil {
		return err
	}
	txIn := wire.NewTxIn(wire.NewOutPoint(txHash, 0), nil, nil)
	tx.AddTxIn(txIn)
	return nil
}

func GetPayToAddScript(addr btcutil.Address) ([]byte, error) {
	return txscript.PayToAddrScript(addr)
}

func CreateNewTransactionOutScript(value int64, script []byte) *wire.TxOut {
	return wire.NewTxOut(value, script)
}

func AddOutput(value int64, p2shAddr, p2wpkhAddr btcutil.Address, tx *wire.MsgTx) error {
	p2shScript, err := GetPayToAddScript(p2shAddr)
	if err != nil {
		return err
	}
	p2wpkhScript, err := GetPayToAddScript(p2wpkhAddr)
	if err != nil {
		return err
	}
	txOut1 := CreateNewTransactionOutScript(value, p2shScript)
	tx.AddTxOut(txOut1)
	txOut2 := CreateNewTransactionOutScript(value, p2wpkhScript)
	tx.AddTxOut(txOut2)
	return nil
}

func Sign(tx *wire.MsgTx, txIn *wire.TxIn, pubKeyHash []byte, privateKey *btcec.PrivateKey) {
	subScript, err := txscript.NewScriptBuilder().
		AddOp(txscript.OP_0).
		AddData(pubKeyHash).
		Script()

	if err != nil {

	}
	// Calculate the signature hash for the input being signed.
	hashType := txscript.SigHashAll
	inputIndex := 0
	sig, _ := txscript.RawTxInSignature(tx, inputIndex, subScript, hashType, privateKey)
	sigScript, err := txscript.NewScriptBuilder().
		AddData(sig).
		AddData(privateKey.PubKey().SerializeCompressed()).
		Script()

	if err != nil {

	}
	// Set the signature script for the input being signed.
	txIn.SignatureScript = sigScript
}
