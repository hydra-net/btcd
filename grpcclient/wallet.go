package grpcclient

import (
	"context"
	"time"

	"github.com/btcsuite/btcd/btcjson"
	pb "github.com/btcsuite/btcd/grpcclient/protos"
)

///////////////////////////////////////////////////////////////////////////////

func (c *Client) ListUtxos(minConf int32, maxConf int32, addresses []string) ([]btcjson.ListUtxosResult, error) {

	ctx, _ := context.WithTimeout(context.Background(), time.Second * 30)

	listUtxosRequest := &pb.ListUtxosRequest {
		MinConf: minConf,
		MaxConf: maxConf,
		Addresses: addresses,
	}

	utxos, err := c.lwClient.ListUtxos(ctx, listUtxosRequest)
	if err != nil {
		return nil, err
	}

	var result []btcjson.ListUtxosResult
	for _, val := range utxos.Utxos {

		var utxo = btcjson.ListUtxosResult {
			TxID: val.Txid,
			Vout: val.Vout,
			PkScript: val.ScriptPubKey,
			Amount: val.Value,
			Confirmations: val.Confirmations,
		}

		result = append(result, utxo)
	}

	return result, nil
}

///////////////////////////////////////////////////////////////////////////////

func (c *Client) GetLastAddress(change bool) (string, error) {

	ctx, _ := context.WithTimeout(context.Background(), time.Second * 30)

	getLastAddressRequest := &pb.GetLastAddressRequest {
		IsChange: change,
	}

	response, err := c.lwClient.GetLastAddress(ctx, getLastAddressRequest)
	if err != nil {
		return "", err
	}

	return response.Address, nil
}

///////////////////////////////////////////////////////////////////////////////

func (c *Client) LWDumpPrivKey(pkScriptHex string) (*string, error) {

	ctx, _ := context.WithTimeout(context.Background(), time.Second * 30)

	dumpPrivKeyRequest := &pb.DumpPrivKeyRequest {
		ScriptPubKey: pkScriptHex,
	}

	response, err := c.lwClient.DumpPrivKey(ctx, dumpPrivKeyRequest)
	if err != nil {
		return nil, err
	}

	return &response.Hash, nil
}

///////////////////////////////////////////////////////////////////////////////

func (c *Client) DerivePrivKey(family uint32, index uint32, hexPubKey string) (*string, error) {

	ctx, _ := context.WithTimeout(context.Background(), time.Second * 30)

	derivePrivKeyRequest := &pb.KeyDescriptor {
		PubKey: hexPubKey,
		Locator: &pb.KeyLocator {
			KeyFamily: family,
			KeyIndex: index,
		},
	}

	response, err := c.keyChainClient.DerivePrivKey(ctx, derivePrivKeyRequest)
	if err != nil {
		return nil, err
	}

	return &response.Hash, nil
}

///////////////////////////////////////////////////////////////////////////////

func (c *Client) DeriveKey(family uint32, index uint32) (*btcjson.KeyDescriptorResult, error) {

	ctx, _ := context.WithTimeout(context.Background(), time.Second * 30)

	deriveKeyRequest := &pb.KeyLocator {
		KeyFamily: family,
		KeyIndex: index,
	}

	response, err := c.keyChainClient.DeriveKey(ctx, deriveKeyRequest)
	if err != nil {
		return nil, err
	}

	result := &btcjson.KeyDescriptorResult {
		Locator: btcjson.KeyLocator {
			Family: response.Locator.KeyFamily,
			Index: response.Locator.KeyIndex,
		},
		HexPubKey: response.PubKey,
	}

	return result, nil
}

///////////////////////////////////////////////////////////////////////////////

func (c *Client) DeriveNextKey(family uint32) (*btcjson.KeyDescriptorResult, error) {

	ctx, _ := context.WithTimeout(context.Background(), time.Second * 30)

	deriveNextKeyRequest := &pb.DeriveNextKeyReq {
		KeyFamily: family,
	}

	response, err := c.keyChainClient.DeriveNextKey(ctx, deriveNextKeyRequest)
	if err != nil {
		return nil, err
	}

	result := &btcjson.KeyDescriptorResult {
		Locator: btcjson.KeyLocator {
			Family: response.Locator.KeyFamily,
			Index: response.Locator.KeyIndex,
		},
		HexPubKey: response.PubKey,
	}

	return result, nil
}

///////////////////////////////////////////////////////////////////////////////

func (c *Client) IsOurAddress(scriptAddress []byte) bool {

	ctx, _ := context.WithTimeout(context.Background(), time.Second * 30)

	isOurAddressRequest := &pb.IsOurAddressRequest{
		ScriptAddress: scriptAddress,
	}

	response, err := c.keyChainClient.IsOurAddress(ctx, isOurAddressRequest)
	if err != nil {
		return false
	}

	return response.IsOur
}

///////////////////////////////////////////////////////////////////////////////