package taro

import (
	"net"
	"time"

	"github.com/btcsuite/btcd/chaincfg"
	"github.com/lightninglabs/taro/address"
	"github.com/lightninglabs/taro/proof"
	"github.com/lightninglabs/taro/tarodb"
	"github.com/lightninglabs/taro/tarofreighter"
	"github.com/lightninglabs/taro/tarogarden"
	"github.com/lightninglabs/taro/universe"
	"github.com/lightningnetwork/lnd"
	"github.com/lightningnetwork/lnd/build"
	"github.com/lightningnetwork/lnd/signal"
	"google.golang.org/grpc"
)

// RPCConfig is a sub-config of the main server that packages up everything
// needed to start the RPC server.
type RPCConfig struct {
	LisCfg *lnd.ListenerCfg

	RPCListeners []net.Addr

	RESTListeners []net.Addr

	GrpcServerOpts []grpc.ServerOption

	RestDialOpts []grpc.DialOption

	RestListenFunc func(net.Addr) (net.Listener, error)

	WSPingInterval time.Duration

	WSPongWait time.Duration

	RestCORS []string

	NoMacaroons bool

	MacaroonPath string

	LetsEncryptDir string

	LetsEncryptListen string

	LetsEncryptDomain string

	LetsEncryptEmail string
}

// DatabaseConfig is the config that holds all the persistence related structs
// and interfaces needed for tarod to function.
type DatabaseConfig struct {
	RootKeyStore *tarodb.RootKeyStore

	MintingStore tarogarden.MintingStore

	AssetStore *tarodb.AssetStore

	TaroAddrBook *tarodb.TaroAddressBook

	UniverseForest *tarodb.BaseUniverseForest

	FederationDB *tarodb.UniverseFederationDB
}

// Config is the main config of the Taro server.
type Config struct {
	DebugLevel string

	// TODO(roasbeef): use the taro chain param wrapper here?
	ChainParams chaincfg.Params

	SignalInterceptor signal.Interceptor

	AssetMinter tarogarden.Planter

	AssetCustodian *tarogarden.Custodian

	ChainBridge tarogarden.ChainBridge

	AddrBook *address.Book

	ProofArchive proof.Archiver

	AssetWallet tarofreighter.Wallet

	ChainPorter tarofreighter.Porter

	BaseUniverse *universe.MintingArchive

	UniverseSyncer universe.Syncer

	UniverseFederation *universe.FederationEnvoy

	// LogWriter is the root logger that all of the daemon's subloggers are
	// hooked up to.
	LogWriter *build.RotatingLogWriter

	*RPCConfig

	*DatabaseConfig
}
