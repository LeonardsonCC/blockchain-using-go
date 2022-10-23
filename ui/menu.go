package ui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

func NewMenu() View {
	return View{
		View:   menuView,
		Update: menuUpdate,
	}
}

func menuView(m model) string {
	s := "What do you want to do?\n"

	for i, option := range m.options {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		s += fmt.Sprintf("%s %s\n", cursor, option.Text)
	}

	s += "\nPress q to quit.\n"

	return s
}

func menuUpdate(m model, msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.options)-1 {
				m.cursor++
			}
		case "enter", " ":
			m := m.options[m.cursor].Fn(m)
			return m, nil
		}
	}

	return m, nil
}
