#!/bin/bash

export APP_VERSION=$(shell cat ./package.json | grep version | head -2 | tail -1 | grep -o '".*"' | sed 's/"//g')

APPNAME="todos"

all: version build run

version:
	@echo "App Version: ${APP_VERSION}"

build:
	@echo "Building application..."
	@go build -v -o todos cmd/serve/main.go
	@echo "😁  Success 😁"

run:
	@go run cmd/serve/main.go

install:
	@dep ensure -v

test:
	@go test ./... -cover -race

docker-build:
	@CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ${APPNAME} cmd/serve/main.go
	@docker build -t ${APPNAME} -f Dockerfile .

git-hook:
	@echo " >> adding git hooks"
	@rm -f -- .git/hooks/pre-commit || true
	@rm -f -- .git/hooks/pre-push || true
	@touch .git/hooks/pre-commit
	@> ./.git/hooks/pre-commit
	@chmod +x ./.git/hooks/pre-commit

	@echo "#!/bin/bash" >> ./.git/hooks/pre-commit
	@echo "set -e" >> ./.git/hooks/pre-commit
	@echo "RED=\`tput setaf 1\`" >> ./.git/hooks/pre-commit
	@echo "GREEN=\`tput setaf 2\`" >> ./.git/hooks/pre-commit
	@echo "YELLOW=\`tput setaf 3\`" >> ./.git/hooks/pre-commit
	@echo "NC=\`tput sgr0\`" >> ./.git/hooks/pre-commit
	@echo 'if [ $$( git diff HEAD --staged --name-only | grep -m 1 ".go$$") ]' >> ./.git/hooks/pre-commit
	@echo "then" >> ./.git/hooks/pre-commit
	@echo "# run go vet w/ go tool to check that the code can be built, etc." >> ./.git/hooks/pre-commit
	@echo 'printf "\\n\\n$${YELLOW}RUNNING GO BUILD...\\n\\n$${NC}"' >> ./.git/hooks/pre-commit
	@echo "make build" >> ./.git/hooks/pre-commit
	@echo 'printf "\\n\\n$${GREEN}OK..\\n\\n$${NC}"' >> ./.git/hooks/pre-commit
	@echo "# check this build w/ the go build race detector" >> ./.git/hooks/pre-commit
	@echo "#go build -race $SOMEDIR" >> ./.git/hooks/pre-commit
	@echo "# run go tests" >> ./.git/hooks/pre-commit
	@echo 'printf "\\n$${YELLOW}RUNNING GO TEST...\\n\\n$${NC}"' >> ./.git/hooks/pre-commit
	@echo 'go test ./... -v | grep "^FAIL" --color=always && exit 1' >> ./.git/hooks/pre-commit
	@echo 'printf "\\n\\n$${GREEN}PASSED...\\n\\n\\n" && exit 0' >> ./.git/hooks/pre-commit
	@echo "fi" >> ./.git/hooks/pre-commit
	@echo " >> git hook added" 

rm-git-hook: 
	@rm ./.git/hooks/pre-commit
	@echo " >> git-hook removed"