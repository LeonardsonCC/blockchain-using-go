package ui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const primary = lipgloss.Color("69")

var (
	square = lipgloss.NewStyle().
		Width(25).
		Height(10).
		Align(lipgloss.Center, lipgloss.Center)
	focusedSquare = lipgloss.NewStyle().
			Width(25).
			Height(10).
			Align(lipgloss.Center, lipgloss.Center).
			Background(primary)
)

func NewPrettyListBlocks() View {
	return View{
		View:   prettyListBlocksView,
		Update: prettyListBlocksUpdate,
	}
}

func prettyListBlocksView(m model) string {
	s := "Blocks:\n\n"

	blocks := []string{}
	for i, block := range m.blockchain.Blocks {
		if i == m.selectedBox {
			blocks = append(blocks, fmt.Sprintf("%s", focusedSquare.Render(block.Hex())))
		} else {
			blocks = append(blocks, fmt.Sprintf("%s", square.Render(block.Hex())))
		}
	}

	s += lipgloss.JoinHorizontal(lipgloss.Top, blocks...)
	s += "\n\nPress q to quit.\n"

	return s
}

func prettyListBlocksUpdate(m model, msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			m = changeView(m, menu)
		case "down", "right", "j", "l":
			m = nextBlock(m)
		case "up", "left", "k", "h":
			m = prevBlock(m)
		case "enter":
			m = changeView(m, prettyViewBlock)
		}
	}

	return m, nil
}

func nextBlock(m model) model {
	if m.selectedBox < len(m.blockchain.Blocks)-1 {
		m.selectedBox = m.selectedBox + 1
	}
	return m
}

func prevBlock(m model) model {
	if m.selectedBox > 0 {
		m.selectedBox = m.selectedBox - 1
	}
	return m
}
