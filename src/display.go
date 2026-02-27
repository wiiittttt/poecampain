package main

import (
	"fmt"
	"strings"

	"charm.land/lipgloss/v2"
)

var (
	red      = lipgloss.Color("#f38ba8")
	peach    = lipgloss.Color("#fab387")
	yellow   = lipgloss.Color("#f9e2af")
	green    = lipgloss.Color("#a6e3a1")
	blue     = lipgloss.Color("#89b4fa")
	sapphire = lipgloss.Color("#74c7ec")
	teal     = lipgloss.Color("#94e2d5")
	mauve    = lipgloss.Color("#cba6f7")
	text     = lipgloss.Color("#cdd6f4")
	base     = lipgloss.Color("#1e1e2e")
	mantle   = lipgloss.Color("#181825")
	overlay0 = lipgloss.Color("#6c7086")

	textStyle     = lipgloss.NewStyle().Foreground(text)
	killStyle     = lipgloss.NewStyle().Foreground(red).Background(mantle)
	npcStyle      = lipgloss.NewStyle().Foreground(mauve).Background(mantle)
	zoneStyle     = lipgloss.NewStyle().Foreground(yellow).Background(mantle)
	waypointStyle = lipgloss.NewStyle().Foreground(sapphire).Background(mantle)
	portalStyle   = lipgloss.NewStyle().Foreground(blue).Background(mantle)
	itemStyle     = lipgloss.NewStyle().Foreground(green).Background(mantle)
	craftStyle    = lipgloss.NewStyle().Foreground(peach).Background(mantle)
	trialStyle    = lipgloss.NewStyle().Foreground(teal).Background(mantle)
	errorStyle    = lipgloss.NewStyle().Foreground(overlay0).Background(red)
)

func replaceUS(s string) string {
	return strings.ReplaceAll(s, "_", " ")
}

type displayFunc func(...string) string

var keywords = map[string]displayFunc{
	"kill": func(s ...string) string {
		return killStyle.Render(fmt.Sprintf("Kill %s", replaceUS(s[0])))
	},
	"red": func(s ...string) string {
		return killStyle.Render(s[0])
	},
	"npc": func(s ...string) string {
		return npcStyle.Render(replaceUS(s[0]))
	},
	"quest": func(s ...string) string {
		return npcStyle.Render(fmt.Sprintf("%s:", replaceUS(s[0])), replaceUS(s[1]))
	},
	"waypoint": func(s ...string) string {
		return waypointStyle.Render(fmt.Sprintf("Waypoint to %s", Areas[s[0]]))
	},
	"relog": func(s ...string) string {
		return zoneStyle.Render(fmt.Sprintf("Relog to %s", Areas[s[0]]))
	},
	"zone": func(s ...string) string {
		return zoneStyle.Render(fmt.Sprintf("Enter %s", Areas[s[0]]))
	},
	"portal": func(s ...string) string {
		return portalStyle.Render(fmt.Sprintf("Portal to %s", Areas[s[0]]))
	},
	"item": func(s ...string) string {
		return itemStyle.Render(replaceUS(s[0]))
	},
	"craft": func(s ...string) string {
		if len(s) == 2 {
			return craftStyle.Render(fmt.Sprintf("󰖼 %s - %s", replaceUS(s[0]), replaceUS(s[1])))
		}
		return craftStyle.Render(fmt.Sprintf("󰖼 %s", replaceUS(s[0])))
	},
	"setportal": func(s ...string) string {
		return portalStyle.Render(fmt.Sprintf("Set portal at the %s", replaceUS(s[0])))
	},
	"trial": func(s ...string) string {
		return trialStyle.Render("Complete the trial")
	},
	"lab": func(s ...string) string {
		return trialStyle.Render(fmt.Sprintf("Complete %s Lab", s[0]))
	},
	"text": func(s ...string) string {
		return textStyle.Render(replaceUS(s[0]))
	},
	"error": func(s ...string) string {
		return errorStyle.Render(fmt.Sprintf("ERROR %s", replaceUS(s[0])))
	},
}

func keywordOrError(s string) displayFunc {
	f, ok := keywords[s]
	if !ok {
		f = keywords["error"]
	}
	return f
}
