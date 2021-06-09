package main

import (
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	searchBox string
	choices   []string
	cursor    int
	selected  map[int]struct{}
}

var initialModed = model{
	searchBox: "",
	choices:   []string{},
	selected:  make(map[int]struct{}),
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:

	}
}

func main() {
	p := tea.NewProgram()
}
