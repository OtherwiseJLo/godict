package main

import (
	"fmt"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"log"
)

type model struct {
	searchBox textinput.Model
	choices   []string
	cursor    int
	selected  map[int]struct{}
	err       error
}

func initialModel() model {
	search := textinput.NewModel()
	search.Placeholder = ""
	search.Focus()
	search.CharLimit = 100
	search.Width = 30

	return model{
		searchBox: search,
		choices:   []string{},
		selected:  make(map[int]struct{}),
	}
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter, tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		}

	// We handle errors just like any other message
	case tea.errMsg:
		m.err = msg
		return m, nil
	}

	m.searchBox, cmd = m.searchBox.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return fmt.Sprintf(
		"What’s your favorite Pokémon?\n\n%s\n\n%s",
		m.searchBox.View(),
		"(esc to quit)",
	) + "\n"
}
func main() {
	p := tea.NewProgram(initialModel())
	if err := p.Start(); err != nil {
		log.Fatal(err)
	}
}
