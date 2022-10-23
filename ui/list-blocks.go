package ui

import (
	"encoding/json"
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

func NewListBlocks() View {
	return View{
		View:   listBlocksView,
		Update: listBlocksUpdate,
	}
}

func listBlocksView(m model) string {
	s := "Blocks:\n"

	for _, block := range m.blockchain.Blocks {
		s += fmt.Sprintf("Block %s:\n", block.Hash)
		for _, tx := range block.Txs {
			b, _ := json.Marshal(tx)
			s += fmt.Sprintf("%s\n", b)
		}
		s += "\n\n"
	}

	s += "\nPress q to quit.\n"

	return s
}

func listBlocksUpdate(m model, msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			m = changeView(m, menu)
		}
	}

	return m, nil
}
