# Quartz Project Specification

## 1. Project Overview

Quartz is a modular Golang Service Framework designed for building scalable, event-driven applications. Its architecture is centered around a flexible, configuration-driven engine that composes and runs various backend services and data-processing pipelines. The framework heavily utilizes a command-based structure, dependency injection, and a component model to promote modularity and extensibility.

## 2. Core Architecture

The core of the framework resides in the `qz` module. It provides the foundation for defining, configuring, and running application services.

### 2.1. Entry Point

The main entry point for any application built with the framework is **`qz/runcmd.go`**.

- The `main()` function in this file bootstraps the system.
- The `init()` function is responsible for registering all available commands (services) that the application can run.
- It invokes a generic `run.Runner`, which handles configuration parsing, component initialization, and the execution of the requested command.

### 2.2. Command-Based Execution

Quartz operates by executing named commands. Each command represents a distinct service or task. This design allows developers to build complex applications by combining multiple, self-contained commands.

Key examples include:

- **`httpssrv`** (`qz/cmd/httpsrv.go`): A dynamically configurable HTTP server.
- **`ChanExec`** (`qz/cmd/ChanExec.go`): A service for creating and running channel-based data processing pipelines, consisting of a source, transformers, and a sink.

### 2.3. Component Model & Dependency Injection

The framework is designed to be highly extensible. Functionality is encapsulated in components, which are initialized and wired together at runtime. This dependency injection model allows developers to easily swap implementations and add new features without modifying the core framework.

## 3. Key Modules

The project is organized into several Go modules, each serving a distinct purpose.

- **`qz/`**: The core framework module.
  - `cmd/`: Contains the implementations for the various commands (services) the application can run.
  - `run/`: Holds the core command execution logic (`runner.go`).
  - `comp/` & `compfact/`: Likely define the components and their factories for the dependency injection system.
  - `dsl/`: Potentially a Domain-Specific Language for configuration or scripting.

- **`arango/`**: Provides components and tools for interacting with an ArangoDB database.

- **`bc/`**: Contains blockchain-related functionality, including data structures for Merkle graphs and transaction data.

- **`chansim/`**: A simulation module with dummy components (Source, Sink, Transformer) for testing channel-based pipelines.

- **`clarity/`**: Implements a parser for the Clarity smart contract language.

- **`gcp/`**: Contains modules for integrating with Google Cloud Platform services like Firestore and Pub/Sub.

- **`gnovm/`**: Integration with the Gno Virtual Machine.

- **`jshttp/`**: A component for serving HTTP requests using an embedded JavaScript engine.

- **`merkle/`**: A standalone module providing data structures and algorithms for Merkle trees.

- **`multivm/`**: Appears to be a multi-virtual machine execution environment, with support for JavaScript and Starlark.

## 4. How to Extend the Codebase

To extend the Quartz framework, developers should follow the existing patterns:

1.  **Create a New Command**: To add a new top-level service, create a new file in the `qz/cmd/` directory. This file should define the command, its configuration flags, and the main execution logic.
2.  **Register the Command**: Add the new command to the `init()` function in `qz/runcmd.go` so the framework recognizes it.
3.  **Develop New Components**: To add new, reusable functionality (e.g., a new database client, a new message queue source), create new components and factories, likely within a dedicated module or the `qz/comp/` directory.
4.  **Utilize Configuration**: Leverage the framework's configuration loading to make your components flexible and configurable from the application's main config file.
