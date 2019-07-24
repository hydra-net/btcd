package grpcclient

import (

	"github.com/btcsuite/btcd/rpcclient"
	"github.com/btcsuite/btcd/rpcclient/grpcclient/protos"
	"google.golang.org/grpc"
)

// New creates a new RPC client based on the provided connection configuration
// details.  The notification handlers parameter may be nil if you are not
// interested in receiving notifications and will be ignored if the
// configuration is set to run in HTTP POST mode.
func New(config *rpcclient.ConnConfig) (lightwalletrpc.LightWalletServiceClient, error) {

	clientConn, err := grpc.Dial(config.Host, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	//defer conn.Close()
	client := lightwalletrpc.NewLightWalletServiceClient(clientConn)

	return client, nil
}