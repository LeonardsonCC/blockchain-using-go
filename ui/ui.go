package ui

import (
	"fmt"
	"os"

	"github.com/LeonardsonCC/blockchain-using-go/blockchain"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type Option struct {
	Fn   func(m model) model
	Text string
}

type Options []Option

type View struct {
	Init   func(m model) model
	View   func(m model) string
	Update func(m model, msg tea.Msg) (tea.Model, tea.Cmd)
}

type Views map[int]View

const (
	menu = iota
	addTx
	listBlocks
	prettyListBlocks
	prettyViewBlock
)

type model struct {
	views         Views
	options       Options
	blockchain    *blockchain.Blockchain
	cursor        int
	currentView   int
	inputs        []textinput.Model
	selectedInput int
	selectedBox   int
}

func Start(blockchain *blockchain.Blockchain) *tea.Program {
	p := tea.NewProgram(initialModel(blockchain))

	if err := p.Start(); err != nil {
		fmt.Printf("there's been an error: %v", err)
		os.Exit(1)
	}
	return p
}

func changeView(m model, view int) model {
	if m.views[view].Init != nil {
		m = m.views[view].Init(m)
	}
	m.currentView = view
	return m
}

func initialModel(blockchain *blockchain.Blockchain) model {
	options := []Option{
		{
			Text: "Create Tx",
			Fn: func(m model) model {
				m = changeView(m, addTx)
				return m
			},
		},
		{
			Text: "List Blocks",
			Fn: func(m model) model {
				m = changeView(m, listBlocks)
				return m
			},
		},
		{
			Text: "Pretty - List Blocks",
			Fn: func(m model) model {
				m = changeView(m, prettyListBlocks)
				return m
			},
		},
	}

	views := Views{
		menu:             NewMenu(),
		addTx:            NewAddTx(),
		listBlocks:       NewListBlocks(),
		prettyListBlocks: NewPrettyListBlocks(),
		prettyViewBlock:  NewPrettyViewBlock(),
	}

	return model{
		views:       views,
		options:     options,
		blockchain:  blockchain,
		currentView: menu,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m.views[m.currentView].Update(m, msg)
}

func (m model) View() string {
	return m.views[m.currentView].View(m)
}
