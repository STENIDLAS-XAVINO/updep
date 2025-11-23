package main

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type KeyMap struct {
	Quit key.Binding
}

var DefaultKeyMap = KeyMap{
	Quit: key.NewBinding(
		key.WithKeys("q", "ctrl+c", "esc"),
		key.WithHelp("↓/j", "move down"),
	),
}

func (m *AppModel) handleKeyPress(msg tea.KeyMsg) tea.Cmd {
	switch {
	case key.Matches(msg, DefaultKeyMap.Quit):
		return tea.Quit
	}
	return nil
}
