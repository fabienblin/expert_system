# **************************************************************************** #
#                                                                              #
#                                                         :::      ::::::::    #
#    Makefile                                           :+:      :+:    :+:    #
#                                                     +:+ +:+         +:+      #
#    By: jmonneri <jmonneri@student.42.fr>          +#+  +:+       +#+         #
#                                                 +#+#+#+#+#+   +#+            #
#    Created: 2019/10/30 17:57:13 by jmonneri          #+#    #+#              #
#    Updated: 2020/02/17 15:44:01 by jmonneri         ###   ########.fr        #
#                                                                              #
# **************************************************************************** #

.PHONY: all get install run fclean tests

GONAME = expert_system

TEST_FILE = other/corr1.txt

GOPATH = $(shell pwd)
GOBIN = $(GOPATH)/bin
GOENV = GOPATH=$(GOPATH) GOBIN=$(GOBIN)  
FILES = engine.go\
		forwardChaining.go\
		header.go\
		infTree.go\
		lexer.go\
		main.go\
		opeFuncForward.go\
		opeFuncs.go\
		parser.go\
		utils.go
GOFILES = $(addprefix cmd/, $(FILES))
EXECPATH = ./bin/$(GONAME)

all: $(EXECPATH)

$(EXECPATH): $(GOFILES)
	@printf "0️⃣  Building $(GOFILES) to ./bin \n"
	@$(GOENV) go build -o $(EXECPATH) $(GOFILES)
	@printf "✅  Built! ✅\n"

get:
	@$(GOENV) go get .

install:
	@$(GOENV) go install $(GOFILES) 

tests: all
	@$(EXECPATH) ./examples/$(filter-out $@,$(MAKECMDGOALS))

testv: all
	@$(EXECPATH) -v ./examples/$(filter-out $@,$(MAKECMDGOALS))

testf: all
	@$(EXECPATH) -f ./examples/$(filter-out $@,$(MAKECMDGOALS))

script: all
	@sh tests.sh

run: all
	@$(EXECPATH) ./examples/$(TEST_FILE)
	
%:
	@:

fclean:
	@echo "Cleaning"
	@$(GOENV) go clean
	@rm -rf ./bin/

re: fclean all