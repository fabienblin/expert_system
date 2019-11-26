# **************************************************************************** #
#                                                                              #
#                                                         :::      ::::::::    #
#    Makefile                                           :+:      :+:    :+:    #
#                                                     +:+ +:+         +:+      #
#    By: jmonneri <jmonneri@student.42.fr>          +#+  +:+       +#+         #
#                                                 +#+#+#+#+#+   +#+            #
#    Created: 2019/10/30 17:57:13 by jmonneri          #+#    #+#              #
#    Updated: 2019/11/26 20:01:31 by jmonneri         ###   ########.fr        #
#                                                                              #
# **************************************************************************** #

.PHONY: all get install run fclean

GONAME = expert_system

TEST_FILE = operators/and2.txt

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
