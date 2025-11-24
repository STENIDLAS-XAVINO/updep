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
	nameCellStyle := lipgloss.NewStyle().Render
	wantedCellStyle := lipgloss.NewStyle().Render
	latestCellStyle := lipgloss.NewStyle().Render

	nameCell := lipgloss.PlaceHorizontal(
		r.columnWidths[0]+gap,
		lipgloss.Left,
		nameCellStyle(r.pkg.Name),
	)
	wantedCell := lipgloss.PlaceHorizontal(
		r.columnWidths[1]+gap,
		lipgloss.Left,
		wantedCellStyle(r.pkg.Wanted.String()),
	)
	latestCell := lipgloss.PlaceHorizontal(
		r.columnWidths[2]+gap,
		lipgloss.Left,
		latestCellStyle(r.pkg.Latest.String()),
	)
	currentCell := lipgloss.PlaceHorizontal(
		r.columnWidths[3],
		lipgloss.Left,
		r.pkg.Current.String(),
	)

	return lipgloss.JoinHorizontal(
		lipgloss.Center,
		nameCell,
		wantedCell,
		latestCell,
		currentCell,
	)
}
