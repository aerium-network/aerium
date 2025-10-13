package wallet

import (
	"context"
	"encoding/hex"
	"errors"
	"net"
	"time"

	"github.com/aerium-network/aerium/crypto/hash"
	"github.com/aerium-network/aerium/types/tx"
	"github.com/aerium-network/aerium/util"
	aerium "github.com/aerium-network/aerium/www/grpc/gen/go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// grpcClient is a gRPC client that randomly establishes a connection to a gRPC server.
// It is used to get information such as account balance or transaction data from the server.
type grpcClient struct {
	ctx               context.Context
	servers           []string
	conn              *grpc.ClientConn
	timeout           time.Duration
	blockchainClient  aerium.BlockchainClient
	transactionClient aerium.TransactionClient
}

func newGrpcClient(timeout time.Duration, servers []string) *grpcClient {
	ctx := context.Background()

	cli := &grpcClient{
		ctx:               ctx,
		timeout:           timeout,
		conn:              nil,
		blockchainClient:  nil,
		transactionClient: nil,
	}

	if len(servers) > 0 {
		cli.servers = servers
	}

	return cli
}

func (c *grpcClient) connect() error {
	if c.conn != nil {
		return nil
	}

	for _, server := range c.servers {
		conn, err := grpc.NewClient(server,
			grpc.WithTransportCredentials(insecure.NewCredentials()),
			grpc.WithContextDialer(func(ctx context.Context, address string) (net.Conn, error) {
				return util.NetworkDialTimeout(ctx, "tcp", address, c.timeout)
			}))
		if err != nil {
			continue
		}

		blockchainClient := aerium.NewBlockchainClient(conn)
		transactionClient := aerium.NewTransactionClient(conn)

		// Check if client is responding
		_, err = blockchainClient.GetBlockchainInfo(c.ctx,
			&aerium.GetBlockchainInfoRequest{})
		if err != nil {
			_ = conn.Close()

			continue
		}

		c.conn = conn
		c.blockchainClient = blockchainClient
		c.transactionClient = transactionClient

		return nil
	}

	return errors.New("unable to connect to the servers")
}

func (c *grpcClient) getBlockchainInfo() (*aerium.GetBlockchainInfoResponse, error) {
	if err := c.connect(); err != nil {
		return nil, err
	}

	info, err := c.blockchainClient.GetBlockchainInfo(c.ctx,
		&aerium.GetBlockchainInfoRequest{})
	if err != nil {
		return nil, err
	}

	return info, nil
}

func (c *grpcClient) getAccount(addrStr string) (*aerium.AccountInfo, error) {
	if err := c.connect(); err != nil {
		return nil, err
	}

	res, err := c.blockchainClient.GetAccount(c.ctx,
		&aerium.GetAccountRequest{Address: addrStr})
	if err != nil {
		return nil, err
	}

	return res.Account, nil
}

func (c *grpcClient) getValidator(addrStr string) (*aerium.ValidatorInfo, error) {
	if err := c.connect(); err != nil {
		return nil, err
	}

	res, err := c.blockchainClient.GetValidator(c.ctx,
		&aerium.GetValidatorRequest{Address: addrStr})
	if err != nil {
		return nil, err
	}

	return res.Validator, nil
}

func (c *grpcClient) sendTx(trx *tx.Tx) (tx.ID, error) {
	if err := c.connect(); err != nil {
		return hash.UndefHash, err
	}

	data, err := trx.Bytes()
	if err != nil {
		return hash.UndefHash, err
	}
	res, err := c.transactionClient.BroadcastTransaction(c.ctx,
		&aerium.BroadcastTransactionRequest{SignedRawTransaction: hex.EncodeToString(data)})
	if err != nil {
		return hash.UndefHash, err
	}

	return hash.FromString(res.Id)
}

// TODO: check the return value type.
func (c *grpcClient) getTransaction(txID tx.ID) (*aerium.GetTransactionResponse, error) {
	if err := c.connect(); err != nil {
		return nil, err
	}

	res, err := c.transactionClient.GetTransaction(c.ctx,
		&aerium.GetTransactionRequest{
			Id:        txID.String(),
			Verbosity: aerium.TransactionVerbosity_TRANSACTION_VERBOSITY_INFO,
		})
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c *grpcClient) getPublicKeyByAddress(ctx context.Context, addr string) (string, error) {
	if err := c.connect(); err != nil {
		return "", err
	}

	res, err := c.blockchainClient.GetPublicKey(ctx, &aerium.GetPublicKeyRequest{Address: addr})
	if err != nil {
		return "", err
	}

	return res.PublicKey, nil
}
