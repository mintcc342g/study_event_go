PACKAGE = study-event-go
BUILDPATH ?= $(CURDIR)
BASE	= $(BUILDPATH)
BIN		= $(BASE)/bin
GORACE  = -race

ifeq ($(OS),Windows_NT)
	PACKAGE = study-event-go.exe
	ifeq ($(PROCESSOR_ARCHITECTURE), x86)
		GORACE =
	endif
else
	UNAME := $(shell uname)
	ifeq ($(UNAME), Linux)
		GOENV   ?= CGO_ENABLED=0 GOOS=linux
	endif
endif

GOBUILD = ${GOENV} go
GO      = go

BUILDTAG=-tags 'studyEventGo'
export GO111MODULE=on

V = 0
Q = $(if $(filter 1,$V),,@)
M = $(shell printf "\033[34;1m▶\033[0m")


## Build all steps
.PHONY: all
all: all_info get tidy gen build

ifeq ($(OS),Windows_NT)
all_info: ; @echo building all steps...
else
all_info: ; $(info $(M) build all steps… ) @
endif


## test
.PHONY: test
test:
ifeq ($(OS),Windows_NT)
	@echo run test...
else
	$(info $(M) run test…) @
endif
	$Q cd $(BASE) && $(GO) test -p 1 -failfast $(GORACE) -cover ./...


## build
.PHONY: build
build:
ifeq ($(OS),Windows_NT)
	@echo building executable...
else
	$(info $(M) building executable… ) @
endif
	$Q cd $(BASE)/cmd && $(GOBUILD) build \
	$(BUILDTAG) \
	-o $(BIN)/$(PACKAGE)


## get
.PHONY: get
get:
ifeq ($(OS),Windows_NT)
	@echo get and install packages...
else
	$(info $(M) get and install packages… ) @
endif
	$Q $(GO) get ./...


## tidy
.PHONY: tidy
tidy:
ifeq ($(OS),Windows_NT)
	@echo tidy packages...
else
    $(info $(M) tidy packages… ) @
endif
	$Q $(GO) mod tidy


## ent # TODO: edit filter-out
.PHONY: ent
ent:
ifeq ($(OS),Windows_NT)
	@echo make entities...
else
	$(info $(M) make entities… ) @
endif
	$Q $(GO) run entgo.io/ent/cmd/ent init $(filter-out $@,$(MAKECMDGOALS))


.PHONY: gen
gen:
ifeq ($(OS),Windows_NT)
	@echo tidy packages...
else
	$(info $(M) generate ent… ) @
endif
	$Q $(GO) generate ./ent
