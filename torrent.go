package main

import (
	"fmt"
	"time"
)

type Torrent struct {
	Id         int64
	Name       string
	Type       string
	UploadedAt string
	Size       string
	Seed       int64
	Leech      int64
	Download   string
}

func (t *Torrent) ToString() string {
	return fmt.Sprintf(
		"[%8d] [%14s] (%4d↑  %4d↓) %s",
		t.Id,
		t.Type,
		t.Seed,
		t.Leech,
		t.Name,
	)
}

func (t *Torrent) ToStringMultiLine() string {
	return fmt.Sprintf(
		"%s%14s  %s%-9d %s%s\n%26s%s%-12s  %s%4d↑  %s%4d↓   %s%-3s%s",
		GetTerminalColor(CLI_GREEN),
		fmt.Sprintf("[%s]", t.Type),
		GetTerminalColor(CLI_MAGENTA),
		t.Id,
		GetTerminalColor(CLI_YELLOW),
		t.Name,
		"",
		GetTerminalColor(CLI_BLUE),
		t.Size,
		GetTerminalColor(CLI_GREEN),
		t.Seed,
		GetTerminalColor(CLI_RED),
		t.Leech,
		GetTerminalColor(CLI_NOCOLOR),
		t.Download,
		GetTerminalColor(CLI_NOCOLOR),
	)
}

func (t *Torrent) UploadedTime() time.Time {
	tt, _ := time.Parse("2006-01-0215:04:05", t.UploadedAt)

	return tt
}
