//go:build gtk

package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/aerium-network/aerium/cmd"
	"github.com/aerium-network/aerium/genesis"
	"github.com/aerium-network/aerium/node"
	"github.com/aerium-network/aerium/util"
	"github.com/aerium-network/aerium/version"
	"github.com/aerium-network/aerium/wallet"
	"github.com/gofrs/flock"
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

const appID = "com.github.aerium-project.aerium.aerium-gui"

var (
	workingDirOpt *string
	passwordOpt   *string
	testnetOpt    *bool
)

func init() {
	workingDirOpt = flag.String("working-dir", cmd.AeriumDefaultHomeDir(), "working directory path")
	passwordOpt = flag.String("password", "", "wallet password")
	testnetOpt = flag.Bool("testnet", false, "initializing for the testnet")
	version.NodeAgent.AppType = "gui"

	if runtime.GOOS == "darwin" {
		// Changing the PANGOCAIRO_BACKEND is necessary on MacOS to render emoji
		_ = os.Setenv("PANGOCAIRO_BACKEND", "fontconfig")
	}

	gtk.Init(nil)
}

func main() {
	flag.Parse()

	// The gtk should run on main thread.
	runtime.UnlockOSThread()
	runtime.LockOSThread()

	// Create a new app.
	app, err := gtk.ApplicationNew(appID, glib.APPLICATION_NON_UNIQUE)
	fatalErrorCheck(err)

	workingDir, err := filepath.Abs(*workingDirOpt)
	if err != nil {
		cmd.PrintErrorMsgf("Aborted! %v", err)

		return
	}

	// If node is not initialized yet
	if util.IsDirNotExistsOrEmpty(workingDir) {
		// TODO: currently genesis hardened to testnet for mainnet we change to genesis.Mainnet.
		network := genesis.Testnet
		if *testnetOpt {
			network = genesis.Testnet
		}
		if !startupAssistant(workingDir, network) {
			return
		}
	}

	// Define the lock file path
	lockFilePath := filepath.Join(workingDir, ".aerium.lock")
	fileLock := flock.New(lockFilePath)

	locked, err := fileLock.TryLock()
	fatalErrorCheck(err)

	if !locked {
		cmd.PrintWarnMsgf("Could not lock '%s', another instance is running?", lockFilePath)

		return
	}

	// Connect function to application startup event, this is not required.
	app.Connect("startup", func() {
		log.Println("application startup")
	})

	node, wlt, err := newNode(workingDir)
	fatalErrorCheck(err)

	// Connect function to application activate event
	app.Connect("activate", func() {
		log.Println("application activate")

		// Show about dialog as splash screen
		splashDlg := aboutDialog()
		splashDlg.SetDecorated(false)
		splashDlg.SetResizable(false)
		splashDlg.SetPosition(gtk.WIN_POS_CENTER)
		splashDlg.SetTypeHint(gdk.WINDOW_TYPE_HINT_SPLASHSCREEN)

		gtk.WindowSetAutoStartupNotification(false)
		splashDlg.ShowAll()
		gtk.WindowSetAutoStartupNotification(true)

		app.AddWindow(splashDlg)

		// This might also force GTK to draw the splash screen
		for gtk.EventsPending() {
			gtk.MainIteration()
		}

		// Running the run-up logic in a separate goroutine
		glib.TimeoutAdd(uint(100), func() bool {
			run(node, wlt, app)
			splashDlg.Destroy()

			// Ensures the function is not called again
			return false
		})
	})

	// Connect function to application shutdown event, this is not required.
	app.Connect("shutdown", func() {
		log.Println("Application shutdown")
		node.Stop()
		_ = fileLock.Unlock()
	})

	cmd.TrapSignal(func() {
		cmd.PrintInfoMsgf("Exiting...")

		node.Stop()
		_ = fileLock.Unlock()
	})

	// Launch the application
	os.Exit(app.Run(nil))
}

func newNode(workingDir string) (*node.Node, *wallet.Wallet, error) {
	// change working directory
	if err := os.Chdir(workingDir); err != nil {
		log.Println("Aborted! Unable to changes working directory. " + err.Error())

		return nil, nil, err
	}

	passwordFetcher := func(wlt *wallet.Wallet) (string, bool) {
		if *passwordOpt != "" {
			return *passwordOpt, true
		}

		return getWalletPassword(wlt)
	}
	n, wlt, err := cmd.StartNode(workingDir, passwordFetcher, nil)
	if err != nil {
		return nil, nil, err
	}

	return n, wlt, nil
}

func run(n *node.Node, wlt *wallet.Wallet, app *gtk.Application) {
	grpcAddr := n.GRPC().Address()
	cmd.PrintInfoMsgf("connect wallet to grpc server: %s\n", grpcAddr)

	nodeModel := newNodeModel(n)
	walletModel := newWalletModel(wlt, n)

	// building main window
	win := buildMainWindow(nodeModel, walletModel)

	// Show the Window and all of its components.
	win.ShowAll()

	walletModel.rebuildModel()

	app.AddWindow(win)
}
