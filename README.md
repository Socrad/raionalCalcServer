# Rational Calculator MCP Server

This project integrates a C++ rational number calculator with a Go-based Model Context Protocol (MCP) server. The Go server exposes the C++ calculator as an MCP tool, allowing other MCP-compatible clients, such as `gemini_cli`, to perform rational number arithmetic.

## Project Structure

- `RationalSystem/`: A Git submodule containing the C++ rational number calculator library.
- `server/`: Contains the Go MCP server implementation.
- `Makefile`: Provides commands for building both the C++ calculator and the Go server.

## Features

- **Rational Number Arithmetic:** Utilizes a C++ library for precise calculations with rational numbers.
- **Model Context Protocol (MCP) Integration:** Exposes the calculator functionality as an MCP tool, enabling seamless integration with MCP-compatible systems like `gemini_cli`.
- **Standard I/O Communication:** The MCP server communicates via standard input/output (stdio)
- **Simple Build Process:** A `Makefile` simplifies the compilation and setup.

## Getting Started

### Prerequisites

- Go (version 1.24.5 or higher)
- C++ Compiler (g++ recommended)
- Git

### Setup

1.  **Clone the repository:**
    ```bash
    git clone --recurse-submodules https://github.com/Socrad/rationalCalcServer.git
    cd rationalCalcServer
    ```
    If you cloned without `--recurse-submodules`, initialize and update the submodule:
    ```bash
    git submodule update --init --recursive
    ```

2.  **Build the project:**
    The `Makefile` will compile both the C++ rational calculator and the Go MCP server.
    ```bash
    make all
    ```
    This will create `server/rational_calculator` (the C++ executable) and `server/mcp_server` (the Go executable).

## Integration with `gemini_cli`

This MCP server is designed to be integrated with `gemini_cli` as a custom tool. After building the `mcp_server` executable, you can configure `gemini_cli`'s `settings.json` (or similar configuration file) to register this server as a tool provider.

*Note: The exact configuration might vary based on your `gemini_cli` version and setup. Refer to `gemini_cli` documentation for precise tool registration details.*

Once configured, the `gemini_cli` (and the underlying LLM) will be able to automatically invoke the `rational_calculator` tool for rational number operations.

## Usage (MCP Tool)

Once the server is running and registered with an MCP client (like `gemini_cli`), you can interact with the `rational_calculator` tool. The tool expects three arguments: `operand1` (string), `operator` (string), and `operand2` (string).

**Tool Name:** `rational_calculator`
**Description:** `Performs calculations with integers and rational numbers. Use this tool for everyday arithmetic and comparisons.`

The tool will return the result of the operation as a string.

## Cleaning Up

To remove the compiled binaries:

```bash
make clean
```

## Dependencies

The `rationalCalcServer` project relies on the following external repositories and tools:

- **[RationalSystem](https://github.com/Socrad/RationalSystem)**: Provides the `Rational` class and `calculator_main.cpp` for rational number arithmetic and comparison operations. Licensed under the MIT License.
- **[MCP Go SDK](https://pkg.go.dev/github.com/modelcontextprotocol/go-sdk/mcp)**: Implements the Model Context Protocol (MCP) server to expose `calculator_main.cpp` operations as tools for LLM clients. Licensed under the MIT License.
- **Go**: Requires Go 1.18 or higher for building the MCP server.
- **C++17 Compiler**: Requires a C++17-compatible compiler (e.g., GCC, Clang) to build the `rational_calculator` executable.
- **CMake** (optional): Used to build the `RationalSystem` components if included in the build process.

## Acknowledgements

This project builds upon the following open-source contributions:
- **RationalSystem**: For the robust `Rational` class implementation, developed with assistance from Google's Gemini.
- **Model Context Protocol (MCP)**: For the standardized protocol enabling LLM integration, maintained by Anthropic and the Go team.

## License

Distributed under the MIT License. See [LICENSE](./LICENSE) for more information.
