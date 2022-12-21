WORKDIR = $(shell pwd)
DOCKER_DIR = $(WORKDIR)/docker
ANTLR_IMAGE = antlr4
MOUNT_DIR = /datastx
SNOWFLAKE_MOUNT_PARSER_DIR = $(MOUNT_DIR)/internal/databases/snowflake/snowflake-parser
SNOWFLAKE_PARSER_DIR = $(WORKDIR)/internal/databases/snowflake/snowflake-parser
POSTGRES_MOUNT_PARSER_DIR = $(MOUNT_DIR)/internal/databases/postgres/postgres-parser
POSTGRES_PARSER_DIR = $(WORKDIR)/internal/databases/postgres/postgres-parser
LOOKER_MOUNT_PARSER_DIR = $(MOUNT_DIR)/internal/looker/looker-parser
LOOKER_PARSER_DIR = $(WORKDIR)/internal/looker/looker-parser
PARSER_PATTERN = -visitor

test:
	go test ./internal/...

repl:
	go run cmd/main.go

purge-branches:
	git branch | grep -v "main" | xargs git branch -D 

antlr-build:
	echo "Building $(ANTLR_IMAGE) docker image"
	echo "WORKDIR: $(WORKDIR)"
	docker build -t antlr4 $(WORKDIR) -f $(DOCKER_DIR)/Dockerfile

antlr-snowflake: antlr-build
	docker run -it -v $(PWD):/datastx antlr4 -Dlanguage=Go $(SNOWFLAKE_MOUNT_PARSER_DIR)/SnowflakeLexer.g4 $(PARSER_PATTERN) -o $(SNOWFLAKE_MOUNT_PARSER_DIR)
	docker run -it -v $(PWD):/datastx antlr4 -Dlanguage=Go $(SNOWFLAKE_MOUNT_PARSER_DIR)/SnowflakeParser.g4 $(PARSER_PATTERN) -o $(SNOWFLAKE_MOUNT_PARSER_DIR)

antlr-postres: antlr-build
	docker run -it -v $(PWD):/datastx antlr4 -Dlanguage=Go $(POSTGRES_MOUNT_PARSER_DIR)/PostgreSQLLexer.g4 $(PARSER_PATTERN) -o $(POSTGRES_MOUNT_PARSER_DIR)
	docker run -it -v $(PWD):/datastx antlr4 -Dlanguage=Go $(POSTGRES_MOUNT_PARSER_DIR)/PostgreSQLParser.g4 $(PARSER_PATTERN) -o $(POSTGRES_MOUNT_PARSER_DIR)

antlr-looker: antlr-build
	docker run -it -v $(PWD):/datastx antlr4 -Dlanguage=Go $(LOOKER_MOUNT_PARSER_DIR)/LookML.g4 $(PARSER_PATTERN) -o $(LOOKER_MOUNT_PARSER_DIR)
	

clean-docker:
	docker system prune -a