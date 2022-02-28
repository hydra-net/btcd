package grpcclient

import (
	"context"
	"github.com/btcsuite/btcd/grpcclient/protos"
	"github.com/btcsuite/btcd/rpcclient"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type Client struct {

	// config holds the connection configuration assoiated with this client.
	config *rpcclient.ConnConfig

	// Conn is the underlying connection.
	conn *grpc.ClientConn

	lwClient 	lightwalletrpc.LightWalletServiceClient
	keyChainClient 	lightwalletrpc.KeychainServiceClient
}

// New creates a new RPC client based on the provided connection configuration
// details.  The notification handlers parameter may be nil if you are not
// interested in receiving notifications and will be ignored if the
// configuration is set to run in HTTP POST mode.
func New(config *rpcclient.ConnConfig) (*Client, error) {

	clientConn, err := grpc.Dial(config.Host, grpc.WithInsecure(), grpc.WithUnaryInterceptor(interceptor))
	if err != nil {
		return nil, err
	}

	lwClient := lightwalletrpc.NewLightWalletServiceClient(clientConn)
	keyChainClient := lightwalletrpc.NewKeychainServiceClient(clientConn)

	client := &Client {
		config: config,
		conn: clientConn,
		lwClient: lwClient,
		keyChainClient: keyChainClient,
	}

	return client, nil
}

func (c *Client) Disconnect() {
	c.conn.Close()
}

// unaryInterceptor is an example unary interceptor.
func interceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {

	ctx = metadata.AppendToOutgoingContext(ctx, "assetid", "0")
	return invoker(ctx, method, req, reply, cc, opts...)
}
