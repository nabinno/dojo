# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Repository Overview

This is a comprehensive coding practice and learning repository called "Nab's dojo" containing exercises and implementations across multiple programming languages, platforms, and technologies.

## Build and Test Commands

### Go Projects
Go projects use Ginkgo/Gomega BDD testing framework. All Go projects have similar structure:

```bash
# Navigate to any Go project directory (a_tour_of_go, atcoder, codewars, leetcode, project_euler)
cd <go_project_directory>

# Run tests
go test

# Run with verbose output
ginkgo -v

# Build
go build
```

### AWS Lambda Project
Complex project with both Go Lambda functions and TypeScript CDK infrastructure:

```bash
cd aws_lambda

# Build all (Go lambdas + TypeScript CDK)
npm run build

# Build only Go lambdas
npm run build:go

# Build only TypeScript
npm run build:ts

# Run tests
npm test

# Deploy infrastructure
npm run deploy

# Local Lambda invocation (requires AWS SAM)
npm run invoke -- <function_name> <event_type>
```

### AWS CDK Projects
```bash
cd aws_cdk/cdk-workshop

# Build TypeScript
npm run build

# Watch mode
npm run watch

# Run tests
npm test

# CDK commands
npm run cdk diff
npm run cdk deploy
```

### DataCamp Python Projects
Uses Poetry for dependency management:

```bash
cd datacamp

# Install dependencies
poetry install

# Activate virtual environment
poetry shell

# Run Python scripts or notebooks
python <script_name>.py
```

### Haskell Project (lyhgg)
```bash
cd lyhgg

# Build with Cabal
cabal build

# Run tests
cabal test

# Run executable
cabal run
```

## Code Architecture

### Go Projects Structure
- Each algorithmic challenge directory follows same pattern:
  - `<name>.go` - main implementation file
  - `<name>_suite_test.go` - BDD test suite using Ginkgo/Gomega
  - `go.mod` - Go module definition with dependencies
- All Go projects target Go 1.12
- Tests use descriptive BDD-style descriptions with nested contexts

### AWS Lambda Architecture
Multi-language serverless application with:
- **Go Lambda Functions**: Located in `lambda/` directory
  - `lambda/authorizer/` - Custom API Gateway authorizer
  - `lambda/pets/` - CRUD operations with Cognito integration
  - `lambda/pretokengen/` - Cognito pre-token generation
- **TypeScript CDK Infrastructure**: Located in `lib/` directory
  - Modular constructs for different AWS services
  - Separate stacks for S3 deployment and API Gateway
- **Build System**: Custom scripts for Go compilation and TypeScript compilation

### React Applications (O'Reilly Projects)
- Modern React with hooks and functional components
- Standard Create React App structure
- Uses yarn for dependency management
- Implements real-world patterns (forms, routing, state management)

### Testing Patterns
1. **Go**: Ginkgo BDD framework with descriptive test suites
2. **TypeScript/JavaScript**: Jest for unit testing
3. **AWS CDK**: CDK testing constructs for infrastructure validation

## Key Dependencies and Versions

- **Go**: Version 1.12 across all projects
- **Node.js**: Various projects using npm/yarn
- **Python**: 3.6+ with data science stack (numpy, pandas, scikit-learn)
- **AWS CDK**: Version 1.x (legacy version)
- **React**: Modern functional components with hooks

## Development Patterns

### Go Development
- Use BDD testing approach with Ginkgo/Gomega
- Follow standard Go project layout
- Implement solutions for algorithmic challenges in separate files
- Test functions are prefixed with `test` and use descriptive names

### AWS Development
- Infrastructure as Code using AWS CDK with TypeScript
- Serverless functions in Go for performance
- Multi-stack deployment strategy
- Local development support with AWS SAM

### Learning Materials
- Extensive markdown documentation in platform-specific directories
- Practical implementations accompanying theoretical learning
- Real-world project examples demonstrating learned concepts

## Repository Navigation

- **Algorithmic Challenges**: `atcoder/`, `codewars/`, `leetcode/`, `project_euler/`, `hackerrank/`
- **Language Learning**: `a_tour_of_go/`, `lyhgg/` (Haskell)
- **Cloud Platforms**: `aws_*/`, `qwiklabs/`, `microsoft_learn/`
- **Educational Content**: `datacamp/`, `coursera/`, `oreilly/`, `pluralsight/`
- **Project Examples**: `aws_lambda/`, `aws_cdk/`, React apps in `oreilly/`