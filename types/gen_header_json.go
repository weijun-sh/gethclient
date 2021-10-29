package types

import (
	"encoding/json"
	"math/big"

	"github.com/weijun-sh/gethclient/common"
	"github.com/weijun-sh/gethclient/common/hexutil"
)

var _ = (*headerMarshaling)(nil)

// MarshalJSON marshals as JSON.
func (h Header) MarshalJSON() ([]byte, error) {
	type Header struct {
		ParentHash  common.Hash    `json:"parentHash"`
		UncleHash   common.Hash    `json:"sha3Uncles"`
		Coinbase    common.Address `json:"miner"`
		Root        common.Hash    `json:"stateRoot"`
		TxHash      common.Hash    `json:"transactionsRoot"`
		ReceiptHash common.Hash    `json:"receiptsRoot"`
		Bloom       Bloom          `json:"logsBloom"`
		Difficulty  *hexutil.Big   `json:"difficulty"`
		Number      *hexutil.Big   `json:"number"`
		GasLimit    hexutil.Uint64 `json:"gasLimit"`
		GasUsed     hexutil.Uint64 `json:"gasUsed"`
		Time        hexutil.Uint64 `json:"timestamp"`
		Extra       hexutil.Bytes  `json:"extraData"`
		MixDigest   common.Hash    `json:"mixHash"`
		Nonce       BlockNonce     `json:"nonce"`
		BaseFee     *hexutil.Big   `json:"baseFeePerGas" rlp:"optional"`
		Hash        common.Hash    `json:"hash"`
	}
	var enc Header
	enc.ParentHash = h.ParentHash
	enc.UncleHash = h.UncleHash
	enc.Coinbase = h.Coinbase
	enc.Root = h.Root
	enc.TxHash = h.TxHash
	enc.ReceiptHash = h.ReceiptHash
	enc.Bloom = h.Bloom
	enc.Difficulty = (*hexutil.Big)(h.Difficulty)
	enc.Number = (*hexutil.Big)(h.Number)
	enc.GasLimit = hexutil.Uint64(h.GasLimit)
	enc.GasUsed = hexutil.Uint64(h.GasUsed)
	enc.Time = hexutil.Uint64(h.Time)
	enc.Extra = h.Extra
	enc.MixDigest = h.MixDigest
	enc.Nonce = h.Nonce
	enc.BaseFee = (*hexutil.Big)(h.BaseFee)
	enc.Hash = h.Hash()
	return json.Marshal(&enc)
}

// UnmarshalJSON unmarshals from JSON.
func (h *Header) UnmarshalJSON(input []byte) error {
	type Header struct {
		ParentHash  *common.Hash    `json:"parentHash"`
		UncleHash   *common.Hash    `json:"sha3Uncles"`
		Coinbase    *common.Address `json:"miner"`
		Root        *common.Hash    `json:"stateRoot"`
		TxHash      *common.Hash    `json:"transactionsRoot"`
		ReceiptHash *common.Hash    `json:"receiptsRoot"`
		Bloom       *Bloom          `json:"logsBloom"`
		Difficulty  *hexutil.Big    `json:"difficulty"`
		Number      *hexutil.Big    `json:"number"`
		GasLimit    *hexutil.Uint64 `json:"gasLimit"`
		GasUsed     *hexutil.Uint64 `json:"gasUsed"`
		Time        *hexutil.Uint64 `json:"timestamp"`
		Extra       *hexutil.Bytes  `json:"extraData"`
		MixDigest   *common.Hash    `json:"mixHash"`
		Nonce       *BlockNonce     `json:"nonce"`
		BaseFee     *hexutil.Big    `json:"baseFeePerGas" rlp:"optional"`
	}
	var dec Header
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}
	if dec.ParentHash != nil {
		h.ParentHash = *dec.ParentHash
	}
	if dec.UncleHash != nil {
		h.UncleHash = *dec.UncleHash
	}
	if dec.Coinbase != nil {
		h.Coinbase = *dec.Coinbase
	}
	if dec.Root != nil {
		h.Root = *dec.Root
	}
	if dec.TxHash != nil {
		h.TxHash = *dec.TxHash
	}
	if dec.ReceiptHash != nil {
		h.ReceiptHash = *dec.ReceiptHash
	}
	if dec.Bloom != nil {
		h.Bloom = *dec.Bloom
	}
	if dec.Difficulty != nil {
		h.Difficulty = (*big.Int)(dec.Difficulty)
	}
	if dec.Number != nil {
		h.Number = (*big.Int)(dec.Number)
	}
	if dec.GasLimit != nil {
		h.GasLimit = uint64(*dec.GasLimit)
	}
	if dec.GasUsed != nil {
		h.GasUsed = uint64(*dec.GasUsed)
	}
	if dec.Time != nil {
		h.Time = uint64(*dec.Time)
	}
	if dec.Extra != nil {
		h.Extra = *dec.Extra
	}
	if dec.MixDigest != nil {
		h.MixDigest = *dec.MixDigest
	}
	if dec.Nonce != nil {
		h.Nonce = *dec.Nonce
	}
	if dec.BaseFee != nil {
		h.BaseFee = (*big.Int)(dec.BaseFee)
	}
	return nil
}
