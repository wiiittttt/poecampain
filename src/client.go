package main

import (
	"regexp"

	tea "charm.land/bubbletea/v2"
	"github.com/SEKOIA-IO/tail"
)

var zoneRe = regexp.MustCompile(`You\shave\sentered\s(.*?)\.`)

type zoneMsg string
type errMsg error

func startTail(path string) (*tail.Tail, error) {
	return tail.TailFile(path, tail.Config{
		Follow:    true,
		ReOpen:    true,
		MustExist: true,
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2},
	})
}

func waitForLine(t *tail.Tail) tea.Cmd {
	return func() tea.Msg {
		for line := range t.Lines {
			if line.Err != nil {
				return line.Err
			}

			match := zoneRe.FindStringSubmatch(line.Text)
			if match != nil {
				return zoneMsg(match[1])
			}
		}
		return nil
	}
}
