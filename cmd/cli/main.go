package main

import (
	tea "github.com/charmbracelet/bubbletea"
)

type AppModel struct {
}

func NewAppModel() AppModel {
	return AppModel{
	}
}

func (m AppModel) Init() tea.Cmd {
	return nil
}

func (m AppModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		cmds = append(cmds, m.handleKeyPress(msg))
	}

	return m, tea.Batch(cmds...)
}

func (m AppModel) View() string {
	return ""
}

func main() {
	p := tea.NewProgram(NewAppModel())
	if _, err := p.Run(); err != nil {
		panic(err)
	}
}
