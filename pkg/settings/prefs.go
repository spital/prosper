package settings

import (
	"encoding/json"
	"sync"

	"fyne.io/fyne/v2"
	"github.com/ventsislav-georgiev/prosper/pkg/global"
	"github.com/ventsislav-georgiev/prosper/pkg/open"
	"github.com/ventsislav-georgiev/prosper/pkg/open/exec"
)

const shortcutsStore = "shortcuts.json"

var (
	prefs *sync.Map
)

type shortcut struct {
	ExecInfo        *exec.Info
	Command         *Command
	KeyNames        []fyne.KeyName
	DisplayKeyNames string
	unregister      func()
}

func Load() error {
	if prefs != nil {
		return nil
	}

	prefs = &sync.Map{}

	r, err := global.AppInstance.Storage().Open(shortcutsStore)
	if err != nil {
		w, err := global.AppInstance.Storage().Create(shortcutsStore)
		if err != nil {
			return err
		}
		w.Close()
		return nil
	}

	defer r.Close()

	var s map[string]*shortcut
	e := json.NewDecoder(r)
	err = e.Decode(&s)
	if err != nil {
		return err
	}

	for k, v := range s {
		prefs.Store(k, v)
	}

	return nil
}

func Save() error {
	w, err := global.AppInstance.Storage().Save(shortcutsStore)
	if err != nil {
		return err
	}

	defer w.Close()

	s := make(map[string]*shortcut)
	prefs.Range(func(k, v interface{}) bool {
		sh := v.(*shortcut)
		if sh.Command != nil && sh.Command.run == nil {
			return true
		}
		s[k.(string)] = sh
		return true
	})

	e := json.NewEncoder(w)
	e.SetIndent("", "\t")
	err = e.Encode(s)
	if err != nil {
		return err
	}

	return nil
}

type Command struct {
	ID   string
	Name string

	icon func() []byte
	run  func()
}

func (e shortcut) ID() string {
	if e.ExecInfo != nil {
		return e.ExecInfo.Filepath()
	}
	return e.Command.ID
}

func (e shortcut) Name() string {
	if e.ExecInfo != nil {
		return e.ExecInfo.DisplayName
	}
	return e.Command.Name
}

func (e shortcut) Icon() (icon []byte) {
	if e.ExecInfo != nil {
		_, icon, _, _ = open.EvalApp(*e.ExecInfo)
	} else {
		icon = e.Command.icon()
	}
	return
}

func (e shortcut) Run() {
	if e.ExecInfo != nil {
		e.ExecInfo.Exec()
		return
	}

	e.Command.run()
}

func (e shortcut) Data() shortcut {
	return e
}