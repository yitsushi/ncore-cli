package main

import "fmt"

type HitAndRunTorrent struct {
	Id         int64
	Name       string
	Start      string
	Updated    string
	Status     string
	Uploaded   string
	Downloaded string
	SeedUntil  string
	Ratio      float64
}

// [Leech] (46ó 47p) Anchorman.2.The.Legend.Continues.2013.UNRATED.720p... (0 B / 2.61 GB)
func (t *HitAndRunTorrent) ToString() string {
	return fmt.Sprintf(
		"[%5s] (%7s) [%8d] %s (%s / %s)",
		t.Status,
		t.SeedUntil,
		t.Id,
		t.Name,
		t.Uploaded,
		t.Downloaded,
	)
}

func (t *HitAndRunTorrent) ToStringMultiLine() string {
	return fmt.Sprintf(
		"%s%9s  %s%9s %s%s\n%21s%s↑%-10s  %s↓%-10s  %s%%%.3f%s",
		GetTerminalColor(CLI_YELLOW),
		fmt.Sprintf("[%s]", t.Status),
		GetTerminalColor(CLI_BLUE),
		fmt.Sprintf("[%s]", t.SeedUntil),
		GetTerminalColor(CLI_NOCOLOR),
		t.Name,
		"",
		GetTerminalColor(CLI_GREEN),
		t.Uploaded,
		GetTerminalColor(CLI_RED),
		t.Downloaded,
		GetTerminalColor(CLI_YELLOW),
		t.Ratio,
		GetTerminalColor(CLI_NOCOLOR),
	)
}
