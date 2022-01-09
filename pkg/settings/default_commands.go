package settings

import (
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/theme"
	"github.com/ventsislav-georgiev/prosper/pkg/global"
	"github.com/ventsislav-georgiev/prosper/pkg/helpers"
	"github.com/ventsislav-georgiev/prosper/pkg/tools/b64"
	"github.com/ventsislav-georgiev/prosper/pkg/tools/cb"
)

const maxNameLen = 16

var (
	CommandRunnerName = "Command Runner"
	defaultCommands   map[string]*shortcut
)

func init() {
	optionKey := "Alt"
	if helpers.IsDarwin {
		optionKey = "Option"
	}

	commands := []*shortcut{
		{
			Command: &Command{
				ID:   CommandRunnerName,
				Name: CommandRunnerName,
				icon: func() []byte { return theme.RadioButtonCheckedIcon().Content() },
				run:  func() { global.AppWindow.Show() },
			},
			KeyNames:        []fyne.KeyName{desktop.KeyAltLeft, fyne.KeySpace},
			DisplayKeyNames: optionKey + "+Space",
		},
		{
			Command: &Command{
				ID:   "Open " + WindowName,
				Name: "Open " + WindowName,
				icon: func() []byte { return theme.SettingsIcon().Content() },
				run:  func() { Show() },
			},
			KeyNames:        []fyne.KeyName{desktop.KeyAltLeft, fyne.KeyBackslash},
			DisplayKeyNames: optionKey + "+\\",
		},
		{
			Command: &Command{
				ID:   b64.WindowName,
				Name: b64.WindowName,
				icon: func() []byte { return theme.StorageIcon().Content() },
				run:  func() { b64.Show() },
			},
			KeyNames:        []fyne.KeyName{desktop.KeyAltLeft, fyne.KeySlash},
			DisplayKeyNames: optionKey + "+/",
		},
		{
			Command: &Command{
				ID:   cb.WindowName,
				Name: cb.WindowName,
				icon: func() []byte { return theme.StorageIcon().Content() },
				run:  func() { cb.Show() },
			},
			KeyNames:        []fyne.KeyName{desktop.KeyAltLeft, desktop.KeyShiftLeft, fyne.KeyA},
			DisplayKeyNames: optionKey + "Shift+A",
		},
	}

	defaultCommands = make(map[string]*shortcut, len(commands))
	for i, v := range commands {
		v.Command.ID = strconv.Itoa(i) + ". " + v.Command.ID

		if len(v.Command.Name) > maxNameLen {
			v.Command.Name = v.Command.Name[:int(maxNameLen)] + "..."
		}

		defaultCommands[v.Command.ID] = v
	}
}