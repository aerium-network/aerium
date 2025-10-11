package grpc

import (
	"context"
	"net"

	"github.com/aerium-network/aerium/consensus/manager"
	"github.com/aerium-network/aerium/network"
	"github.com/aerium-network/aerium/state"
	"github.com/aerium-network/aerium/sync"
	"github.com/aerium-network/aerium/util"
	"github.com/aerium-network/aerium/util/logger"
	"github.com/aerium-network/aerium/wallet"
	aerium "github.com/aerium-network/aerium/www/grpc/gen/go"
	"github.com/aerium-network/aerium/www/zmq"
	"google.golang.org/grpc"
)

type Server struct {
	ctx           context.Context
	config        *Config
	listener      net.Listener
	server        *grpc.Server
	address       string
	state         state.Facade
	net           network.Network
	sync          sync.Synchronizer
	consMgr       manager.ManagerReader
	walletMgr     *wallet.Manager
	zmqPublishers []zmq.Publisher
	logger        *logger.SubLogger
}

func NewServer(ctx context.Context, conf *Config, state state.Facade, sync sync.Synchronizer,
	network network.Network, consMgr manager.ManagerReader,
	walletMgr *wallet.Manager,
	zmqPublishers []zmq.Publisher,
) *Server {
	return &Server{
		ctx:           ctx,
		config:        conf,
		state:         state,
		sync:          sync,
		net:           network,
		consMgr:       consMgr,
		walletMgr:     walletMgr,
		zmqPublishers: zmqPublishers,
		logger:        logger.NewSubLogger("_grpc", nil),
	}
}

func (s *Server) Address() string {
	return s.address
}

func (s *Server) StartServer() error {
	if !s.config.Enable {
		return nil
	}

	listener, err := util.NetworkListen(s.ctx, "tcp", s.config.Listen)
	if err != nil {
		return err
	}

	return s.startListening(listener)
}

func (s *Server) startListening(listener net.Listener) error {
	opts := make([]grpc.UnaryServerInterceptor, 0)

	if s.config.BasicAuth != "" {
		opts = append(opts, BasicAuth(s.config.BasicAuth))
	}

	opts = append(opts, s.Recovery())

	grpcServer := grpc.NewServer(grpc.ChainUnaryInterceptor(opts...))

	blockchainServer := newBlockchainServer(s)
	transactionServer := newTransactionServer(s)
	networkServer := newNetworkServer(s)
	utilServer := newUtilsServer(s)

	aerium.RegisterBlockchainServer(grpcServer, blockchainServer)
	aerium.RegisterTransactionServer(grpcServer, transactionServer)
	aerium.RegisterNetworkServer(grpcServer, networkServer)
	aerium.RegisterUtilsServer(grpcServer, utilServer)

	if s.config.EnableWallet {
		walletServer := newWalletServer(s, s.walletMgr)

		aerium.RegisterWalletServer(grpcServer, walletServer)
	}

	s.listener = listener
	s.address = listener.Addr().String()
	s.server = grpcServer

	go func() {
		s.logger.Info("gRPC server start listening", "address", listener.Addr())
		if err := s.server.Serve(listener); err != nil {
			s.logger.Debug("error on gRPC server", "error", err)
		}
	}()

	return nil
}

func (s *Server) StopServer() {
	if s.server != nil {
		s.server.Stop()
		_ = s.listener.Close()
	}
}
