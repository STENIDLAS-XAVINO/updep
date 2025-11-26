package row

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type KeyMap struct {
	Wanted, Latest, ToggleSelect key.Binding
}

var keyMap = KeyMap{
	Wanted: key.NewBinding(
		key.WithKeys("w"),
		key.WithHelp("w", "update to wanted version"),
	),
	Latest: key.NewBinding(
		key.WithKeys("l"),
		key.WithHelp("l", "update to latest version"),
	),
	ToggleSelect: key.NewBinding(
		key.WithKeys(" "),
		key.WithHelp("space", "toggle"),
	),
}

func (r *Row) handleKeyPress(msg tea.KeyMsg) tea.Cmd {
	switch {
	case key.Matches(msg, keyMap.Wanted):
		r.Target = &r.Pkg.Wanted
	case key.Matches(msg, keyMap.Latest):
		r.Target = &r.Pkg.Latest
	case key.Matches(msg, keyMap.ToggleSelect):
		if r.Target != nil {
			r.Target = nil
			break
		}
		if r.Pkg.Current.Compare(r.Pkg.Wanted) >= 0 {
			r.Target = &r.Pkg.Latest
		} else {
			r.Target = &r.Pkg.Wanted
		}
	}
	return nil
}
