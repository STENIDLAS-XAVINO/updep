package row

import (
	"npmupdate/pkg/config"

	"github.com/charmbracelet/lipgloss"
)

var (
	ActiveRowStyle = lipgloss.NewStyle().
			Bold(true).
			Background(config.Theme.Surface1).Render
)
