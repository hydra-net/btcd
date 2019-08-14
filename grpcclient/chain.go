package grpcclient

import (
	"context"
	"encoding/hex"
	"time"
	"bytes"

	"github.com/btcsuite/btcd/btcjson"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/wire"
	"github.com/btcsuite/btcutil"
	"github.com/btcsuite/btcutil/gcs"
	pb "github.com/btcsuite/btcd/grpcclient/protos"
	"google.golang.org/grpc/status"
)

///////////////////////////////////////////////////////////////////////////////

func (c *Client) GetBestBlock() (*chainhash.Hash, int32, error) {

	ctx, _ := context.WithTimeout(context.Background(), time.Second * 30)

	chainInfo, err := c.lwClient.GetChainInfo(ctx, &pb.Empty{})
	if err != nil {
		return nil, 0, err
	}

	hash, err := chainhash.NewHashFromStr(chainInfo.BestBlockHash)
	if err != nil {
		return nil, 0, err
	}

	return hash, chainInfo.Height, nil
}

///////////////////////////////////////////////////////////////////////////////

func (c *Client) GetBlock(hash *chainhash.Hash) (*wire.MsgBlock, error) {

	ctx, _ := context.WithTimeout(context.Background(), time.Second * 30)

	blockHash := &pb.BlockHash {
		Hash: hash.String(),
	}

	response, err := c.lwClient.GetBlock(ctx, blockHash)
	if err != nil {
		return nil, err
	}

	// get blockheader
	serializedBH, err := hex.DecodeString(response.BlockHeader)
	if err != nil {
		return nil, err
	}

	// Deserialize the blockheader and return it.
	var bh wire.BlockHeader
	err = bh.Deserialize(bytes.NewReader(serializedBH))
	if err != nil {
		return nil, err
	}

	// get transactions
	var txns []*wire.MsgTx

	for _, val := range response.Transactions {

		rawBytes, _ := hex.DecodeString(val)

		var tx wire.MsgTx
		err := tx.Deserialize(bytes.NewReader(rawBytes))

		if err != nil {
			return nil, err
		}

		txns = append(txns, &tx)
	}

	block := wire.MsgBlock {
		Header: bh,
		Transactions: txns,
	}

	return &block, nil
}


///////////////////////////////////////////////////////////////////////////////

func (c *Client) GetFilterBlock(hash *chainhash.Hash) ([]*wire.MsgTx, error) {

	ctx, _ := context.WithTimeout(context.Background(), time.Second * 30)

	blockHash := &pb.BlockHash {
		Hash: hash.String(),
	}

	response, err := c.lwClient.GetFilterBlock(ctx, blockHash)
	if err != nil {
		return nil, err
	}

	var result []*wire.MsgTx

	for _, val := range response.Transactions {

		rawBytes, _ := hex.DecodeString(val)

		var tx wire.MsgTx
		err := tx.Deserialize(bytes.NewReader(rawBytes))

		if err != nil {
			return nil, err
		}

		result = append(result, &tx)
	}

	return result, nil
}

///////////////////////////////////////////////////////////////////////////////

func (c *Client) GetUnspentOutput(hash *chainhash.Hash, index uint32) (*btcjson.GetUnspentOutputResult, error) {

	ctx, _ := context.WithTimeout(context.Background(), time.Second * 30)

	outpoint := &pb.Outpoint {
		Hash: hash.String(),
		Index: index,
	}
	response, err := c.lwClient.GetTxOut(ctx, outpoint)
	if err != nil {

		stat := status.Convert(err)
		if stat.Code() == 14 {
			return nil, nil
		}

		return nil, err
	}

	result := &btcjson.GetUnspentOutputResult{
		response.ScriptPubkey,
		response.Value,
	}

	return result, err
}

///////////////////////////////////////////////////////////////////////////////

func (c *Client) GetBlockHeight(hash *chainhash.Hash) (int32, error) {

	ctx, _ := context.WithTimeout(context.Background(), time.Second * 30)

	blockHash := &pb.BlockHash {
		Hash: hash.String(),
	}

	header, err := c.lwClient.GetBlockHeaderVerbose(ctx, blockHash)
	if err != nil {
		return 0, err
	}

	return header.Height, nil
}

///////////////////////////////////////////////////////////////////////////////

func (c *Client) GetBlockHash(height int64) (*chainhash.Hash, error) {
	ctx, _ := context.WithTimeout(context.Background(), time.Second * 30)

	hashHeight := &pb.Height {
		Height: int32(height),
	}

	response, err := c.lwClient.GetBlockHash(ctx, hashHeight)
	if err != nil {
		return nil, err
	}

	result, err := chainhash.NewHashFromStr(response.Hash)
	if err != nil {
		return nil, err
	}

	return result, err
}

///////////////////////////////////////////////////////////////////////////////

func (c *Client) GetBlockHeader(hash *chainhash.Hash) (*wire.BlockHeader, error) {
	ctx, _ := context.WithTimeout(context.Background(), time.Second * 30)

	blockHash := &pb.BlockHash {
		Hash: hash.String(),
	}

	response, err := c.lwClient.GetBlockHeader(ctx, blockHash)
	if err != nil {
		return nil, err
	}

	serializedBH, err := hex.DecodeString(response.Hash)
	if err != nil {
		return nil, err
	}

	// Deserialize the blockheader and return it.
	var bh wire.BlockHeader
	err = bh.Deserialize(bytes.NewReader(serializedBH))
	if err != nil {
		return nil, err
	}

	return &bh, err
}

///////////////////////////////////////////////////////////////////////////////

func (c *Client) GetRawTransaction(hash *chainhash.Hash) (*btcutil.Tx, error) {
	ctx, _ := context.WithTimeout(context.Background(), time.Second * 30)

	txid := &pb.TxID {
		Txid: hash.String(),
	}

	response, err := c.lwClient.GetRawTransaction(ctx, txid)
	if err != nil {
		return nil, err
	}

	// Decode the serialized transaction hex to raw bytes.
	serializedTx, err := hex.DecodeString(response.Hash)
	if err != nil {
		return nil, err
	}

	// Deserialize the transaction and return it.
	var msgTx wire.MsgTx
	if err := msgTx.Deserialize(bytes.NewReader(serializedTx)); err != nil {
		return nil, err
	}

	return btcutil.NewTx(&msgTx), err
}

///////////////////////////////////////////////////////////////////////////////

func (c *Client) SendRawTransaction(tx *wire.MsgTx, allowHighFees bool) (*chainhash.Hash, error) {

	ctx, _ := context.WithTimeout(context.Background(), time.Second * 30)

	txHex := ""
	if tx != nil {
		// Serialize the transaction and convert to hex string.
		buf := bytes.NewBuffer(make([]byte, 0, tx.SerializeSize()))
		if err := tx.Serialize(buf); err != nil {
			return nil, err
		}
		txHex = hex.EncodeToString(buf.Bytes())
	}

	sendRequest := &pb.HexEncoded {
		Hash: txHex,
	}

	response, err := c.lwClient.SendRawTransaction(ctx, sendRequest)
	if err != nil {
		return nil, err
	}

	return chainhash.NewHashFromStr(response.Txid)
}

///////////////////////////////////////////////////////////////////////////////

func (c *Client) GetBlockFilter(hash *chainhash.Hash) (*gcs.Filter, error) {

	ctx, _ := context.WithTimeout(context.Background(), time.Second * 30)

	blockHash := &pb.BlockHash {
		Hash: hash.String(),
	}

	response, err := c.lwClient.GetBlockFilter(ctx, blockHash)
	if err != nil {
		return nil, err
	}

	// Decode the serialized block hex to raw bytes.
	bytes, err := hex.DecodeString(response.Bytes)
	if err != nil {
		return nil, err
	}

	filter, err := gcs.FromBytes(response.N, uint8(response.P), response.M, bytes)
	if err != nil {
		return nil, err
	}

	return filter, err
}

///////////////////////////////////////////////////////////////////////////////

func (c *Client) Generate(numBlocks uint32) ([]*chainhash.Hash, error) {

	ctx, _ := context.WithTimeout(context.Background(), time.Second * 30)

	genRequest := &pb.GenerateRequest {
		NumBlocks: numBlocks,
	}

	response, err := c.lwClient.Generate(ctx, genRequest)
	if err != nil {
		return nil, err
	}

	// Convert each block hash to a chainhash.Hash and store a pointer to
	// each.
	convertedResult := make([]*chainhash.Hash, len(response.BlockHash))
	for i, hashString := range response.BlockHash {
		convertedResult[i], err = chainhash.NewHashFromStr(hashString)
		if err != nil {
			return nil, err
		}
	}

	return convertedResult, nil
}

///////////////////////////////////////////////////////////////////////////////

func (c *Client) GetBlockHeaderVerbose(hash *chainhash.Hash) (*btcjson.GetBlockHeaderVerboseResult, error) {

	ctx, _ := context.WithTimeout(context.Background(), time.Second * 30)

	blockHash := &pb.BlockHash {
		Hash: hash.String(),
	}
	response, err := c.lwClient.GetBlockHeaderVerbose(ctx, blockHash)
	if err != nil {
		return nil, err
	}

	result := &btcjson.GetBlockHeaderVerboseResult {
		Hash: response.Hash,
		Height: response.Height,
		PreviousHash: response.PrevBlockHash,
		Time: response.Time,
	}

	return result, nil
}