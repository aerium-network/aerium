package grpc

import (
	"context"
	"net"
	"path/filepath"
	"testing"

	"github.com/aerium-network/aerium/consensus"
	"github.com/aerium-network/aerium/consensus/manager"
	"github.com/aerium-network/aerium/crypto/bls"
	"github.com/aerium-network/aerium/genesis"
	"github.com/aerium-network/aerium/network"
	"github.com/aerium-network/aerium/state"
	"github.com/aerium-network/aerium/sync"
	"github.com/aerium-network/aerium/util"
	"github.com/aerium-network/aerium/util/testsuite"
	"github.com/aerium-network/aerium/wallet"
	aerium "github.com/aerium-network/aerium/www/grpc/gen/go"
	"github.com/aerium-network/aerium/www/zmq"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
)

type testData struct {
	*testsuite.TestSuite

	mockState     *state.MockState
	mockSync      *sync.MockSync
	consMocks     []*consensus.MockConsensus
	mockConsMgr   manager.Manager
	defaultWallet *wallet.Wallet
	listener      *bufconn.Listener
	server        *Server
}

func testConfig() *Config {
	conf := DefaultConfig()
	conf.WalletsDir = util.TempDirPath()

	return conf
}

func setup(t *testing.T, conf *Config) *testData {
	t.Helper()

	if conf == nil {
		conf = testConfig()
	}

	ts := testsuite.NewTestSuite(t)

	// for saving test wallets in temp directory
	t.Chdir(util.TempDirPath())

	const bufSize = 1024 * 1024

	listener := bufconn.Listen(bufSize)
	valKeys := []*bls.ValidatorKey{ts.RandValKey(), ts.RandValKey()}
	mockState := state.MockingState(ts)
	mockNet := network.MockingNetwork(ts, ts.RandPeerID())
	mockSync := sync.MockingSync(ts)
	mockConsMgr, consMocks := manager.MockingManager(ts, mockState, valKeys)

	mockState.CommitTestBlocks(10)

	wltPath := filepath.Join(conf.WalletsDir, "default_wallet")
	mnemonic, _ := wallet.GenerateMnemonic(128)
	defaultWallet, err := wallet.Create(wltPath, mnemonic, "", genesis.Mainnet)
	require.NoError(t, err)
	require.NoError(t, defaultWallet.Save())

	mockWalletMgrConf := wallet.DefaultConfig()
	mockWalletMgrConf.WalletsDir = conf.WalletsDir
	mockWalletMgrConf.ChainType = mockState.Genesis().ChainType()

	zmqPublishers := []zmq.Publisher{
		zmq.MockingPublisher("zmq_address", "zmq_topic", 100),
	}

	server := NewServer(context.Background(), conf,
		mockState, mockSync, mockNet, mockConsMgr,
		wallet.NewWalletManager(mockWalletMgrConf), zmqPublishers,
	)
	err = server.startListening(listener)
	assert.NoError(t, err)

	return &testData{
		TestSuite:     ts,
		mockState:     mockState,
		mockSync:      mockSync,
		consMocks:     consMocks,
		mockConsMgr:   mockConsMgr,
		defaultWallet: defaultWallet,
		server:        server,
		listener:      listener,
	}
}

func (td *testData) StopServer() {
	td.server.StopServer()
	_ = td.listener.Close()
}

func (td *testData) bufDialer(context.Context, string) (net.Conn, error) {
	return td.listener.Dial()
}

func (td *testData) blockchainClient(t *testing.T) (*grpc.ClientConn, aerium.BlockchainClient) {
	t.Helper()

	conn, err := grpc.NewClient("passthrough://bufnet",
		grpc.WithContextDialer(td.bufDialer),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	assert.NoError(t, err)

	return conn, aerium.NewBlockchainClient(conn)
}

func (td *testData) networkClient(t *testing.T) (*grpc.ClientConn, aerium.NetworkClient) {
	t.Helper()

	conn, err := grpc.NewClient("passthrough://bufnet",
		grpc.WithContextDialer(td.bufDialer),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	assert.NoError(t, err)

	return conn, aerium.NewNetworkClient(conn)
}

func (td *testData) transactionClient(t *testing.T) (*grpc.ClientConn, aerium.TransactionClient) {
	t.Helper()

	conn, err := grpc.NewClient("passthrough://bufnet",
		grpc.WithContextDialer(td.bufDialer),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	assert.NoError(t, err)

	return conn, aerium.NewTransactionClient(conn)
}

func (td *testData) walletClient(t *testing.T) (*grpc.ClientConn, aerium.WalletClient) {
	t.Helper()

	conn, err := grpc.NewClient("passthrough://bufnet",
		grpc.WithContextDialer(td.bufDialer),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	assert.NoError(t, err)

	return conn, aerium.NewWalletClient(conn)
}

func (td *testData) utilClient(t *testing.T) (*grpc.ClientConn, aerium.UtilsClient) {
	t.Helper()

	conn, err := grpc.NewClient("passthrough://bufnet",
		grpc.WithContextDialer(td.bufDialer),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	assert.NoError(t, err)

	return conn, aerium.NewUtilsClient(conn)
}
