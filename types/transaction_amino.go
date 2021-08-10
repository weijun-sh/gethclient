package types

import (
	"crypto/sha256"
	"errors"
	"sync"

	"github.com/jowenshaw/gethclient/common"
	amino "github.com/tendermint/go-amino"
)

var (
	aminoCdc       *amino.Codec
	aminoCdcInitor sync.Once

	ErrTxIsNotAminoCodec = errors.New("transaction is not amino codec")
)

// IsAminoCodec is amino codec
func (tx *Transaction) IsAminoCodec() bool {
	chainID := tx.ChainId()
	return chainID.Uint64() == 66 // okex chain
}

func getAminoCdc() *amino.Codec {
	if aminoCdc == nil {
		aminoCdcInitor.Do(func() {
			aminoCdc = amino.NewCodec()
			aminoCdc.RegisterConcrete(MsgEthereumTx{}, "ethermint/MsgEthereumTx", nil)
			aminoCdc.Seal()
		})
	}
	return aminoCdc
}

// MsgEthereumTx encapsulates an Ethereum transaction as an SDK message.
type MsgEthereumTx struct {
	Data *LegacyTx
}

// CalcAminoHash calc amino tx hash
func CalcAminoHash(tx *Transaction) (hash common.Hash, err error) {
	legacyTx, ok := tx.inner.(*LegacyTx)
	if !ok {
		return hash, ErrTxIsNotAminoCodec
	}
	txBytes, err := getAminoCdc().MarshalBinaryLengthPrefixed(MsgEthereumTx{legacyTx})
	if err != nil {
		return hash, err
	}

	hash = common.Hash(sha256.Sum256(txBytes))
	return hash, nil
}

// MarshalAmino defines custom encoding scheme
func (td LegacyTx) MarshalAmino() ([]byte, error) {
	gasPrice, err := common.MarshalBigInt(td.GasPrice)
	if err != nil {
		return nil, err
	}

	amount, err := common.MarshalBigInt(td.Value)
	if err != nil {
		return nil, err
	}

	v, err := common.MarshalBigInt(td.V)
	if err != nil {
		return nil, err
	}

	r, err := common.MarshalBigInt(td.R)
	if err != nil {
		return nil, err
	}

	s, err := common.MarshalBigInt(td.S)
	if err != nil {
		return nil, err
	}

	e := encodableTxData{
		AccountNonce: td.Nonce,
		Price:        gasPrice,
		GasLimit:     td.Gas,
		Recipient:    td.To,
		Amount:       amount,
		Payload:      td.Data,
		V:            v,
		R:            r,
		S:            s,
	}

	return getAminoCdc().MarshalBinaryBare(e)
}

// encodableTxData implements the Ethereum transaction data structure. It is used
// solely as intended in Ethereum abiding by the protocol.
type encodableTxData struct {
	AccountNonce uint64
	Price        string
	GasLimit     uint64
	Recipient    *common.Address
	Amount       string
	Payload      []byte

	// signature values
	V string
	R string
	S string
}
