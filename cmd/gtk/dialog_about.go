//go:build gtk

package main

import (
	_ "embed"

	"github.com/aerium-network/aerium/cmd"
	"github.com/aerium-network/aerium/version"
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
)

var (
	//go:embed assets/ui/dialog_about.ui
	uiAboutDialog []byte

	//go:embed assets/images/logo.png
	aeriumLogo []byte
)

func aboutDialog() *gtk.AboutDialog {
	builder, err := gtk.BuilderNewFromString(string(uiAboutDialog))
	fatalErrorCheck(err)

	dlg := getAboutDialogObj(builder, "id_dialog_about")

	pxLogo, err := gdk.PixbufNewFromBytesOnly(aeriumLogo)
	if err != nil {
		cmd.PrintErrorMsgf("Failed to load Logo Pixbuf: %v", err)
	} else {
		dlg.SetLogo(pxLogo)
	}

	dlg.SetVersion(version.NodeVersion().StringWithAlias())

	return dlg
}
