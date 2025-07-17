.PHONY: all rational_calculator mcp_server clean

all: rational_calculator mcp_server

rational_calculator:
	@echo "Building rational_calculator from RationalSystem/src..."
	g++ -o RationalSystem/rational_calculator RationalSystem/src/*.cpp -I RationalSystem/src
	cp RationalSystem/rational_calculator server/rational_calculator

mcp_server:
	@echo "Building mcp_server..."
	go build -o server/mcp_server server/main.go

clean:
	@echo "Cleaning up..."
	rm -f RationalSystem/rational_calculator
	rm -f server/rational_calculator server/mcp_server
