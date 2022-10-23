package ui

import (
	"fmt"
	"strconv"
	"time"

	"github.com/LeonardsonCC/blockchain-using-go/blockchain"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func NewAddTx() View {
	return View{
		Init:   addTxInit,
		View:   addTxView,
		Update: addTxUpdate,
	}
}

const (
	hotPink  = lipgloss.Color("#FF06B7")
	darkGray = lipgloss.Color("#767676")
)

var (
	inputStyle    = lipgloss.NewStyle().Foreground(hotPink)
	continueStyle = lipgloss.NewStyle().Foreground(darkGray)
)

const (
	sender = iota
	receiver
	value
)

func addTxInit(m model) model {
	m.inputs = []textinput.Model{
		sender:   textinput.New(),
		receiver: textinput.New(),
		value:    textinput.New(),
	}

	m.inputs[sender].Placeholder = "Sender"

	m.inputs[receiver].Placeholder = "Receiver"

	m.inputs[value].Placeholder = "Value"

	m.selectedInput = sender

	return m
}

func addTxView(m model) string {
	return fmt.Sprintf(
		"New transaction:\n%s\n%s\n%s\n%s\n%s\n%s\n\nPress Enter to finish\n",
		inputStyle.Width(30).Render("Sender"),
		m.inputs[sender].View(),
		inputStyle.Width(30).Render("Receiver"),
		m.inputs[receiver].View(),
		inputStyle.Width(30).Render("Value"),
		m.inputs[value].View())
}

func addTxUpdate(m model, msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd = make([]tea.Cmd, len(m.inputs))

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC:
			m = changeView(m, menu)
		case tea.KeyShiftTab:
			m = prevInput(m)
		case tea.KeyTab:
			m = nextInput(m)
		case tea.KeyEnter:
			if m.selectedInput == len(m.inputs)-1 {
				value, err := strconv.ParseInt(m.inputs[value].Value(), 10, 64)
				if err == nil {
					data := blockchain.NewData(m.inputs[sender].Value(), m.inputs[receiver].Value(), int(value))
					m.blockchain.AddTx(time.Now(), data)
					m = changeView(m, menu)
				}
			} else {
				m = nextInput(m)
			}
		}

		// clear focus and then set the focus
		for i := range m.inputs {
			m.inputs[i].Blur()
		}
		m.inputs[m.selectedInput].Focus()
	}

	for i := range m.inputs {
		m.inputs[i], cmds[i] = m.inputs[i].Update(msg)
	}

	return m, tea.Batch(cmds...)
}

func nextInput(m model) model {
	if m.selectedInput < len(m.inputs)-1 {
		m.selectedInput = m.selectedInput + 1
	}
	return m
}

func prevInput(m model) model {
	if m.selectedInput > 0 {
		m.selectedInput = m.selectedInput - 1
	}
	return m
}
