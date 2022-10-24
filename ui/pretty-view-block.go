package ui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

func NewPrettyViewBlock() View {
	return View{
		View:   prettyViewBlockView,
		Update: prettyViewBlockUpdate,
	}
}

func prettyViewBlockView(m model) string {
	s := "Txs:\n\n"

	for _, tx := range m.blockchain.Blocks[m.selectedBox].Txs {
		s += fmt.Sprintf("%s\n%s %s\n%s %s\n%s %d\n",
			tx.Hash,
			title.Render("Sender"),
			tx.Data.Sender,
			title.Render("Receiver"),
			tx.Data.Receiver,
			title.Render("Value"),
			tx.Data.Value)

		s += "\n"
	}

	s += "\n\nPress q to quit.\n"

	return s
}

func prettyViewBlockUpdate(m model, msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			m = changeView(m, prettyListBlocks)
		}
	}

	return m, nil
}
