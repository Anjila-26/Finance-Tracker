package main

import (
	"fmt"
	"strconv"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	money     float64
	stage     int // 0: menu, 1: input, 2: result
	op        string
	input     textinput.Model
	resultMsg string
}

func initialModel(startMoney float64) model {
	ti := textinput.New()
	ti.Placeholder = "amount"
	ti.Focus()
	ti.CharLimit = 32
	ti.Width = 20

	return model{
		money: startMoney,
		stage: 0,
		input: ti,
	}
}

func (m model) Init() tea.Cmd { return textinput.Blink }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "1":
			m.op = "add"
			m.stage = 1
			m.input.SetValue("")
			m.input.Focus()
			return m, nil
		case "2":
			m.op = "deduct"
			m.stage = 1
			m.input.SetValue("")
			m.input.Focus()
			return m, nil
		case "enter":
			if m.stage == 1 {
				s := m.input.Value()
				v, err := strconv.ParseFloat(s, 64)
				if err != nil {
					m.resultMsg = "invalid amount"
				} else {
					if m.op == "add" {
						m.money = add_money(m.money, v)
						m.resultMsg = fmt.Sprintf("Added: $%.2f", v)
					} else if m.op == "deduct" {
						m.money = deduct_money(m.money, v)
						m.resultMsg = fmt.Sprintf("Deducted: $%.2f", v)
					}
				}
				m.stage = 2
				return m, nil
			}
		}
	}

	if m.stage == 1 {
		m.input, cmd = m.input.Update(msg)
	}

	return m, cmd
}

func (m model) View() string {
	switch m.stage {
	case 0:
		return fmt.Sprintf("Current money: $%.2f\n\n(1) Add\n(2) Deduct\n\nq to quit\n", m.money)
	case 1:
		return fmt.Sprintf("Current money: $%.2f\n\nEnter amount: %s\n\n(enter to submit, q to quit)", m.money, m.input.View())
	case 2:
		return fmt.Sprintf("Current money: $%.2f\n\n%s\n\n(press any key to return to menu, q to quit)", m.money, m.resultMsg)
	}
	return ""
}

// runTUI launches the Bubble Tea TUI.
func runTUI(startMoney float64) error {
	p := tea.NewProgram(initialModel(startMoney))
	if err := p.Start(); err != nil {
		return err
	}
	return nil
}
