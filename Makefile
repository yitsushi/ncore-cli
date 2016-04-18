all: cover build

coverFile := $(shell mktemp -u -t cover.out.XXXXXX)
projectName := $(shell basename `pwd`)
hasGocov := #$(shell which gocov)
packages := $(shell go list ./...)

ifndef OUTPUT
	OUTPUT = "text"
endif

ifndef WINDOWS
	WINDOWS = 0
endif

ifndef LINUX
	LINUX = 0
endif

ifndef DARWIN
	DARWIN = 0
endif

ifndef AMD64
	AMD64 = 0
endif

ifndef X86
	X86 = 0
endif

ifndef NOCLEAR
	NOCLEAR = 0
endif

ifeq ($(OS),Windows_NT)
	WINDOWS = 1
else
	UNAME_S := $(shell uname -s)
	ifeq ($(UNAME_S),Linux)
		LINUX = 1
	endif
	ifeq ($(UNAME_S),Darwin)
		DARWIN = 1
	endif
	UNAME_P := $(shell uname -m)
	ifeq ($(UNAME_P),x86_64)
		AMD64 = 1
	endif
	ifneq ($(filter %86,$(UNAME_P)),)
		X86 = 1
	endif
endif
ifeq ($(PROCESSOR_ARCHITECTURE),AMD64)
	AMD64 = 1
endif
ifeq ($(PROCESSOR_ARCHITECTURE),x86)
	X86 = 1
endif

define b
	GOOS=$(1) GOARCH=$(2) go build -o build/bin/$(projectName)$(3)
endef

define r
	./build/bin/$(projectName)$(1)
endef

build: clean build_mac build_linux build_windows

build_windows: clean
ifeq ($(WINDOWS),1)
ifeq ($(X86),1)
	$(call b,windows,386,-win-x86.exe)
endif
ifeq ($(AMD64),1)
	$(call b,windows,amd64,-win-amd64.exe)
endif
endif

run_windows: build_windows
ifeq ($(WINDOWS),1)
ifeq ($(X86),1)
	$(call r,-win-x86.exe)
endif
ifeq ($(AMD64),1)
	$(call r,-win-amd64.exe)
endif
endif

build_linux: clean
ifeq ($(LINUX),1)
ifeq ($(X86),1)
	$(call b,linux,386,-lin-x86.bin)
endif
ifeq ($(AMD64),1)
	$(call b,linux,amd64,-lin-amd64.bin)
endif
endif

run_linux: build_linux
ifeq ($(LINUX),1)
ifeq ($(X86),1)
	$(call r,-lin-x86.bin)
endif
ifeq ($(AMD64),1)
	$(call r,-lin-amd64.bin)
endif
endif

build_mac: clean
ifeq ($(DARWIN),1)
ifeq ($(X86),1)
	$(call b,darwin,386,-mac-x86)
endif
ifeq ($(AMD64),1)
	$(call b,darwin,amd64,-mac-amd64)
endif
endif

run_mac: build_mac
ifeq ($(DARWIN),1)
ifeq ($(X86),1)
	$(call r,-mac-x86)
endif
ifeq ($(AMD64),1)
	$(call r,-mac-amd64)
endif
endif

cover_tools: clean
	@echo $(shell for pkg in $(packages); do \
		go test -coverprofile $(coverFile)-$$(echo "$$pkg" | sed -e 's/\//-/g') $$pkg; \
		done) > /dev/null
	@echo "mode: set" >  $(coverFile)
	@cat $(coverFile)* | grep -v "^mode:" >> $(coverFile)
ifeq ("$(OUTPUT)","html")
	@go tool cover -html=$(coverFile) -o build/test/cover.html
else
	@go tool cover -func=$(coverFile)
endif

cover_gocov: clean
	@gocov test ./indycator > $(coverFile)
ifeq ("$(OUTPUT)","html")
	@gocov-html $(coverFile) > build/test/cover.html
else
	@gocov report $(coverFile)
endif

cover:
ifeq ($(strip $(hasGocov)),)
	@$(MAKE) cover_tools
else
	@$(MAKE) cover_gocov
endif

clean:
ifneq ($(NOCLEAR),1)
	@mkdir -p build/bin
	@mkdir -p build/test
	@rm -f build/bin/*
	@rm -f build/test/*
endif

run: run_linux run_mac run_windows

help:
	@echo ""
	@echo "Epic Makefile of Go"
	@echo ""
	@echo "Try to detect current operation system and architecture"
	@echo "to build at least for the build machine."
	@echo " !!! Generate multi-package cover info."
	@echo ""
	@echo "Available flags: "
	@echo "  // cover and testing"
	@echo "    OUTPUT   can be 'html' otherwise 'text'; used for generate cover report"
	@echo "  // build OS"
	@echo "    LINUX    can be '1' otherwise '0'; if you want to build a binary for Linux"
	@echo "    DARWIN   can be '1' otherwise '0'; if you want to build a binary for MacOS"
	@echo "    WINDOWS  can be '1' otherwise '0'; if you want to build a binary for Windows"
	@echo "  // build OS"
	@echo "    AMD64    can be '1' otherwise '0'; if you want to build a binary for x86-64"
	@echo "    X86      can be '1' otherwise '0'; if you want to build a binary for x86"
	@echo "  // misc"
	@echo "    NOCLEAR  can be '1' otherwise '0'; if you do NOT want to clear build directory"
	@echo ""

.PHONY: test cover build clean build_windows build_linux build_mac help