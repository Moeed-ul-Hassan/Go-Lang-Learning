package tui

import (
	"fmt"
	"gen-code/internal/matrix"
	"gen-code/internal/scaffold"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type state int

const (
	stateAppName state = iota
	stateLanguageSelection
	stateFrameworkSelection
	stateEnvSelection
	stateComplexitySelection
	statePath
	stateScaffolding
	stateDone
)

const asciiHeader = `
  ____ _____ _   _    ____ ___  ____  _____ 
 / ___| ____| \ | |  / ___/ _ \|  _ \| ____|
| |  _|  _| |  \| | | |  | | | | | | |  _|  
| |_| | |___| |\  | | |__| |_| | |_| | |___ 
 \____|_____|_| \_|  \____\___/|____/|_____|
`

type Model struct {
	state        state
	choice       int
	quitting     bool
	windowWidth  int
	windowHeight int

	// Input fields
	textInput    textinput.Model
	appName      string
	selectedLang matrix.Language
	selectedFW   matrix.Framework
	outputPath   string

	env  matrix.ProjectType
	comp matrix.Complexity
}

type scaffoldingMsg struct {
	err error
}

func InitialModel() Model {
	ti := textinput.New()
	ti.Placeholder = "Enter app name..."
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 30

	return Model{
		state:     stateAppName,
		textInput: ti,
		choice:    0,
	}
}

func (m Model) Init() tea.Cmd {
	return textinput.Blink
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			m.quitting = true
			return m, tea.Quit
		}

		if m.state == stateAppName || m.state == statePath {
			if msg.String() == "enter" {
				if m.state == stateAppName {
					m.appName = m.textInput.Value()
					m.state = stateLanguageSelection
					m.choice = 0
				} else {
					m.outputPath = m.textInput.Value()
					m.state = stateScaffolding
					return m, func() tea.Msg {
						err := scaffold.Scaffold(m.appName, m.selectedLang, m.selectedFW, m.env, m.comp, m.outputPath)
						return scaffoldingMsg{err: err}
					}
				}
				return m, nil
			}
			m.textInput, cmd = m.textInput.Update(msg)
			return m, cmd
		}

		switch msg.String() {
		case "q":
			m.quitting = true
			return m, tea.Quit
		case "up", "k":
			if m.choice > 0 {
				m.choice--
			}
		case "down", "j":
			max := 2
			if m.state == stateFrameworkSelection {
				switch m.selectedLang {
				case matrix.Go:
					max = 2
				case matrix.JS:
					max = 1
				case matrix.Python:
					max = 2
				}
			}
			if m.choice < max {
				m.choice++
			}
		case "enter":
			switch m.state {
			case stateLanguageSelection:
				m.selectedLang = []matrix.Language{matrix.Go, matrix.JS, matrix.Python}[m.choice]
				m.state = stateFrameworkSelection
				m.choice = 0
			case stateFrameworkSelection:
				switch m.selectedLang {
				case matrix.Go:
					m.selectedFW = []matrix.Framework{matrix.Gin, matrix.Echo, matrix.Fiber}[m.choice]
				case matrix.JS:
					m.selectedFW = []matrix.Framework{matrix.Express, matrix.Fastify}[m.choice]
				case matrix.Python:
					m.selectedFW = []matrix.Framework{matrix.Flask, matrix.FastAPI, matrix.Django}[m.choice]
				}
				m.state = stateEnvSelection
				m.choice = 0
			case stateEnvSelection:
				m.env = []matrix.ProjectType{matrix.WebApp, matrix.CLI, matrix.Backend}[m.choice]
				m.state = stateComplexitySelection
				m.choice = 0
			case stateComplexitySelection:
				m.comp = []matrix.Complexity{matrix.Minimal, matrix.Standard, matrix.Enterprise}[m.choice]
				m.state = statePath
				m.textInput.Placeholder = "Enter output path..."
				m.textInput.SetValue(".")
				m.textInput.Focus()
			}
		default:
			if m.state == stateDone {
				m.quitting = true
				return m, tea.Quit
			}
		}

	case scaffoldingMsg:
		if msg.err != nil {
			fmt.Printf("Error during scaffolding: %v\n", msg.err)
			return m, tea.Quit
		}
		m.state = stateDone
		return m, nil

	case tea.WindowSizeMsg:
		m.windowWidth = msg.Width
		m.windowHeight = msg.Height
	}

	return m, nil
}

func (m Model) View() string {
	if m.quitting {
		return "Bye!\n"
	}

	headerStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#7D56F4")).
		MarginBottom(1)

	header := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#00D7FF")).
		Bold(true).
		SetString(asciiHeader).
		PaddingLeft(2).
		Render()

	var s string
	switch m.state {
	case stateAppName:
		s = header + "\n" + headerStyle.Render("Step 1: Application Name") + "\n\n"
		s += m.textInput.View() + "\n\n"
		s += "(press enter to continue)"
	case stateLanguageSelection:
		s = header + "\n" + headerStyle.Render("Step 2: Select Language") + "\n\n"
		choices := []string{"Go", "JavaScript", "Python"}
		for i, choice := range choices {
			cursor := " "
			if m.choice == i {
				cursor = ">"
				choice = lipgloss.NewStyle().Foreground(lipgloss.Color("#FF00FF")).Render(choice)
			}
			s += fmt.Sprintf("%s %s\n", cursor, choice)
		}
	case stateFrameworkSelection:
		s = header + "\n" + headerStyle.Render("Step 3: Select Framework") + "\n\n"
		var choices []string
		switch m.selectedLang {
		case matrix.Go:
			choices = []string{"Gin", "Echo", "Fiber"}
		case matrix.JS:
			choices = []string{"Express", "Fastify"}
		case matrix.Python:
			choices = []string{"Flask", "FastAPI", "Django"}
		}
		for i, choice := range choices {
			cursor := " "
			if m.choice == i {
				cursor = ">"
				choice = lipgloss.NewStyle().Foreground(lipgloss.Color("#00D7FF")).Render(choice)
			}
			s += fmt.Sprintf("%s %s\n", cursor, choice)
		}
	case stateEnvSelection:
		s = header + "\n" + headerStyle.Render("Step 4: Select Environment") + "\n\n"
		choices := []string{"Web Application", "CLI Tool", "Backend Service"}
		for i, choice := range choices {
			cursor := " "
			if m.choice == i {
				cursor = ">"
				choice = lipgloss.NewStyle().Foreground(lipgloss.Color("#FF00FF")).Render(choice)
			}
			s += fmt.Sprintf("%s %s\n", cursor, choice)
		}
	case stateComplexitySelection:
		s = header + "\n" + headerStyle.Render("Step 5: Select Complexity") + "\n\n"
		choices := []string{"Minimal (MVP)", "Standard (Clean Architecture)", "Enterprise (Microservices Ready)"}
		for i, choice := range choices {
			cursor := " "
			if m.choice == i {
				cursor = ">"
				choice = lipgloss.NewStyle().Foreground(lipgloss.Color("#00FFFF")).Render(choice)
			}
			s += fmt.Sprintf("%s %s\n", cursor, choice)
		}
	case statePath:
		s = header + "\n" + headerStyle.Render("Step 6: Output Path") + "\n\n"
		s += m.textInput.View() + "\n\n"
		s += "(press enter to generate)"
	case stateScaffolding:
		s = "Generating your masterpiece..."
	case stateDone:
		s = header + "\n" + lipgloss.NewStyle().Foreground(lipgloss.Color("#00FF00")).Render("Success! Your project has been scaffolded.") + "\n\n"
		s += "Project: " + m.appName + "\n"
		s += "Language: " + string(m.selectedLang) + "\n"
		s += "Framework: " + string(m.selectedFW) + "\n"
		s += "Created by: Moeed ul Hassan\n\n"
		s += "Check " + m.outputPath + " for the new structure.\n"
		s += "Press any key to exit."
	}

	return lipgloss.NewStyle().Margin(1, 2).Render(s)
}
