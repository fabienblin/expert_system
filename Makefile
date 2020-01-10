# **************************************************************************** #
#                                                                              #
#                                                         :::      ::::::::    #
#    Makefile                                           :+:      :+:    :+:    #
#                                                     +:+ +:+         +:+      #
#    By: jmonneri <jmonneri@student.42.fr>          +#+  +:+       +#+         #
#                                                 +#+#+#+#+#+   +#+            #
#    Created: 2019/10/30 17:57:13 by jmonneri          #+#    #+#              #
#    Updated: 2020/01/06 18:26:17 by jmonneri         ###   ########.fr        #
#                                                                              #
# **************************************************************************** #

.PHONY: all get install run fclean tests

GONAME = expert_system

TEST_FILE = other/corr1.txt

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

tests: all
	@$(EXECPATH) ./examples/$(filter-out $@,$(MAKECMDGOALS))

run: all
	@$(EXECPATH) ./examples/$(TEST_FILE)
	
%:
	@:

fclean:
	@echo "Cleaning"
	@$(GOENV) go clean
	@rm -rf ./bin/
