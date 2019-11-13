#******************************************************************************#
#*                                                                            *#
#*          ▄▄▄██▀▀▀      ███▄ ▄███▓         ██▓    ▄▄▄       ▄▄▄▄            *#
#*            ▒██        ▓██▒▀█▀ ██▒        ▓██▒   ▒████▄    ▓█████▄          *#
#*            ░██        ▓██    ▓██░        ▒██░   ▒██  ▀█▄  ▒██▒ ▄██         *#
#*         ▓██▄██▓       ▒██    ▒██         ▒██░   ░██▄▄▄▄██ ▒██░█▀           *#
#*          ▓███▒    ██▓ ▒██▒   ░██▒ ██▓    ░██████▒▓█   ▓██▒░▓█  ▀█▓         *#
#*          ▒▓▒▒░    ▒▓▒ ░ ▒░   ░  ░ ▒▓▒    ░ ▒░▓  ░▒▒   ▓▒█░░▒▓███▀▒         *#
#*          ▒ ░▒░    ░▒  ░  ░      ░ ░▒     ░ ░ ▒  ░ ▒   ▒▒ ░▒░▒   ░          *#
#*          ░ ░ ░    ░   ░      ░    ░        ░ ░    ░   ▒    ░    ░          *#
#*          ░   ░     ░         ░     ░         ░  ░     ░  ░ ░               *#
#*                    ░               ░                            ░          *#
#*                                                                            *#
#******************************************************************************#
                                   #* Makefile *#


SRC_FILES = engine.go \
			header.go \
			lexer.go \
			main.go \
			parser.go \
			infTree.go \
			utils.go

.PHONY: all get install run fclean


GOPATH = $(shell pwd)
GOBIN = $(GOPATH)/bin
GOENV = GOPATH=$(GOPATH) GOBIN=$(GOBIN)  
GOFILES = $(wildcard cmd/*.go)
GONAME = expert_system
EXECPATH = ./bin/$(GONAME)
TEST_FILE = example_input.txt

all: $(EXECPATH)

$(EXECPATH): $(GOFILES)
	@printf "0️⃣  Building $(GOFILES) to ./bin \n"
	@$(GOENV) go build -o $(EXECPATH) $(GOFILES)
	@printf "✅  Built! ✅\n"

get:
	@$(GOENV) go get .

install:
	@$(GOENV) go install $(GOFILES) 

run: all
	@$(EXECPATH) ./examples/$(TEST_FILE)

fclean:s
	@echo "Cleaning"
	@$(GOENV) go clean
	@rm -rf ./bin/
