package main

import (
	"fmt"

	tea "charm.land/bubbletea/v2"
	"github.com/SEKOIA-IO/tail"
)

type guideModel struct {
	guide    *Guide
	tail     *tail.Tail
	quitting bool
	err      error
}

func NewGuideModel(g *Guide, t *tail.Tail) guideModel {
	return guideModel{
		guide: g,
		tail:  t,
	}
}

func (m guideModel) Init() tea.Cmd {
	pos, _ := readState()
	m.guide.SetPos(pos)

	return waitForLine(m.tail)
}

func (m guideModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case zoneMsg:
		if m.guide.IsNextZone(string(msg)) {
			m.guide.Next()
		}
		return m, waitForLine(m.tail)

	case errMsg:
		m.err = msg
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			m.quitting = true
			return m, tea.Quit
		case "up", "r":
			m.guide.Start()
			return m, nil
		case "left":
			m.guide.Prev()
			return m, nil
		case "right":
			m.guide.Next()
			return m, nil
		case "down":
			m.guide.End()
			return m, nil
		}
	}
	return m, nil
}

func (m guideModel) View() tea.View {
	if m.quitting {
		m.tail.Cleanup()

		if err := writeState(m.guide.Pos); err != nil {
			return tea.NewView(fmt.Sprintf("failed to save state: %v", err))
		}

		return tea.NewView("")
	}

	if m.err != nil {
		return tea.NewView(m.err.Error())
	}

	v := tea.NewView(m.guide.Display())
	v.AltScreen = true
	return v
}
