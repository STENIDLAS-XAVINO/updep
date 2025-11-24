package main

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type KeyMap struct {
	Up, Down, Quit, ExpandHelp key.Binding
}

var keyMap = KeyMap{
	Up: key.NewBinding(
		key.WithKeys("k", "up"),
		key.WithHelp("↑/k", "up"),
	),
	Down: key.NewBinding(
		key.WithKeys("j", "down"),
		key.WithHelp("↓/j", "down"),
	),
	Quit: key.NewBinding(
		key.WithKeys("q", "ctrl+c", "esc"),
		key.WithHelp("q/esc", "quit"),
	),
	ExpandHelp: key.NewBinding(
		key.WithKeys("?"),
		key.WithHelp("?", "help"),
	),
}

func (k KeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Up, k.Down, k.ExpandHelp}
}

func (k KeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Up, k.Quit},
		{k.Down},
		{k.ExpandHelp},
	}
}

func (m *AppModel) handleKeyPress(msg tea.KeyMsg) tea.Cmd {
	switch {
	case key.Matches(msg, keyMap.Up):
		if m.cursor > 0 {
			m.cursor -= 1
		}
	case key.Matches(msg, keyMap.Down):
		if m.cursor < len(m.rows)-1 {
			m.cursor += 1
		}
	case key.Matches(msg, keyMap.Quit):
		return tea.Quit
	case key.Matches(msg, keyMap.ExpandHelp):
		m.help.ShowAll = !m.help.ShowAll
	}
	return nil
}
