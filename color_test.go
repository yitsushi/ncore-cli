package main

import (
    "testing"
)

func Test_setColorMode_enable_nocolor(t *testing.T) {
    noColor = false

    setColorMode(true);

    if noColor != true {
        t.Error("noColor:false; expected: true")
    }
}

func Test_setColorMode_disable_nocolor(t *testing.T) {
    noColor = true

    setColorMode(false);

    if noColor != false {
        t.Error("noColor:true; expected: false")
    }
}

func Test_setColorMode_named_enable_nocolor(t *testing.T) {
    noColor = true

    setColorMode(MODE_NOCOLOR);

    if noColor != true {
        t.Error("noColor:false; expected: true")
    }
}

func Test_setColorMode_named_disable_nocolor(t *testing.T) {
    noColor = true

    setColorMode(MODE_COLOR);

    if noColor != false {
        t.Error("noColor:true; expected: false")
    }
}

func Test_GetTerminalColor(t *testing.T) {
    setColorMode(MODE_NOCOLOR)

    if GetTerminalColor(CLI_RED) != "" {
        t.Error("color mode is disabled, but the return value was a color code")
    }

    setColorMode(MODE_COLOR)

    if GetTerminalColor(CLI_RED) == "" {
        t.Error("color mode is enabled, but the return value was not a color code")
    }
}
