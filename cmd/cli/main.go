package main

import (
	"npmupdate/pkg/components/row"
	"npmupdate/pkg/entities"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type AppModel struct {
	packages []entities.Package
	header   row.Row
	rows     []row.Row
	cursor   int
}

func NewAppModel() AppModel {
	packages := entities.ParseJSON()

	columnWidth := []int{0, 0, 0, 0}
	for _, p := range packages {
		columnWidth[0] = max(columnWidth[0], lipgloss.Width(p.Name))
		columnWidth[1] = max(columnWidth[1], lipgloss.Width(p.Wanted))
		columnWidth[2] = max(columnWidth[2], lipgloss.Width(p.Latest))
		columnWidth[3] = max(columnWidth[3], lipgloss.Width(p.Current))
	}

	header := row.New(entities.Package{
		Name:    "Package Name",
		Current: "Wanted",
		Wanted:  "Latest",
		Latest:  "Current",
	}, columnWidth)

	rows := []row.Row{}
	for _, p := range packages {
		rows = append(rows, row.New(p, columnWidth))
	}

	return AppModel{
		packages: packages,
		header:   header,
		rows:     rows,
		cursor:   0,
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
	renderRows := []string{}
	for i, p := range m.rows {
		if i != m.cursor {
			renderRows = append(renderRows, p.View())
			continue
		}

		renderRows = append(
			renderRows,
			row.ActiveRowStyle.Render(p.View()),
		)
	}

	return lipgloss.JoinVertical(lipgloss.Left, renderRows...)
}

func main() {
	p := tea.NewProgram(NewAppModel())
	if _, err := p.Run(); err != nil {
		panic(err)
	}
}
