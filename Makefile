# **************************************************************************** #
#                                                                              #
#                                                         :::      ::::::::    #
#    Makefile                                           :+:      :+:    :+:    #
#                                                     +:+ +:+         +:+      #
#    By: jmonneri <jmonneri@student.42.fr>          +#+  +:+       +#+         #
#                                                 +#+#+#+#+#+   +#+            #
#    Created: 2019/10/30 17:57:13 by jmonneri          #+#    #+#              #
#    Updated: 2019/10/31 00:44:03 by jmonneri         ###   ########.fr        #
#                                                                              #
# **************************************************************************** #

.PHONY: all get install run fclean

GONAME = expert_system

TEST_FILE = example_input.txt
#TEST_FILE = errors_input.txt

GOPATH = $(shell pwd)
GOBIN = $(GOPATH)/bin
GOENV = GOPATH=$(GOPATH) GOBIN=$(GOBIN)  
GOFILES = $(wildcard cmd/*.go)
EXECPATH = ./bin/$(GONAME)

all: $(EXECPATH)

$(EXECPATH): $(GOFILES)
	@printf "0️⃣  Building $(GOFILES) to ./bin \n"
	@$(GOENV) go build -o $(EXECPATH) $(GOFILES)
	@printf "✅  Builded! ✅\n"

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
