PACKAGE = study-event-go
BUILDPATH ?= $(CURDIR)
BASE	= $(BUILDPATH)
BIN		= $(BASE)/bin

ifeq ($(OS),Windows_NT)
	PACKAGE = study-event-go.exe
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

.PHONY: all
all: gen build ; $(info $(M) building all steps… ) @ ## Build all steps

.PHONY: build
build: ; $(info $(M) building executable… ) @ ## Build program binary
	$Q cd $(BASE)/cmd && $(GOBUILD) build -i \
		$(BUILDTAG) \
		-o $(BIN)/$(PACKAGE)

.PHONY: buildw
buildw:
	$Q cd $(BASE)/cmd && $(GOBUILD) build -i \
		$(BUILDTAG) \
		-o $(BIN)/$(PACKAGE)

## TODO: edit filter-out
.PHONY: ent
ent:
	$(GO) run entgo.io/ent/cmd/ent init $(filter-out $@,$(MAKECMDGOALS))

.PHONY: gen
gen: ; $(info $(M) generate ent… )
	$(GO) generate ./ent