package row

import (
	"npmupdate/pkg/config"

	"github.com/charmbracelet/lipgloss"
)

var (
	ActiveRowStyle = lipgloss.NewStyle().
			Bold(true).
			Background(config.Theme.Surface1)
	latestStyle       = lipgloss.NewStyle().Foreground(config.Theme.Red)
	errorVersionStyle = lipgloss.NewStyle().Foreground(config.Theme.Red)
	needUpdateStyle   = lipgloss.NewStyle().
				Foreground(config.Theme.Green)
	wantedStyle = lipgloss.NewStyle().
			Foreground(config.Theme.Green)
	optionalUpdateStyle = lipgloss.NewStyle().
				Foreground(config.Theme.Peach)
)
