package row

import (
	"npmupdate/pkg/entities"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var gap int = 4

type Row struct {
	pkg          entities.Package
	columnWidths []int
}

func New(pkg entities.Package, columnWidths []int) Row {
	return Row{
		pkg:          pkg,
		columnWidths: columnWidths,
	}
}

func (r Row) Init() tea.Cmd {
	return nil
}

func (r Row) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	return r, tea.Batch(cmds...)
}

func (r Row) View() string {
	return lipgloss.JoinHorizontal(
		lipgloss.Left,
		lipgloss.PlaceHorizontal(r.columnWidths[0]+gap, lipgloss.Left, r.pkg.Name),
		lipgloss.PlaceHorizontal(
			r.columnWidths[1]+gap,
			lipgloss.Left,
			r.pkg.Wanted,
		),
		lipgloss.PlaceHorizontal(
			r.columnWidths[2]+gap,
			lipgloss.Left,
			r.pkg.Latest,
		),
		lipgloss.PlaceHorizontal(
			r.columnWidths[3],
			lipgloss.Left,
			r.pkg.Current,
		),
	)
}
