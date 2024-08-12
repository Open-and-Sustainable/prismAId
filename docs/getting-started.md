---
title: Getting Started
layout: default
toc: true
---

# Getting Started

## Download prismAId
To use prismAId, download the appropriate executable for your operating system from our [GitHub Releases](https://github.com/ricboer0/prismAId/releases) page.

## Setting up a review project
Add text preparation, and ref to prompt design
Follow instructions in the [User Manual](user_manual/manual.md).

## Running prismAId
Simply [download](https://github.com/ricboer0/prismAId/releases) the executable for your OS and platform, place it in a suitable directory, prepare a project configuration (.toml), and run it from your command line, e.g.:

```bash
# For Windows
./prismAId_windows_amd64.exe --project your_project.toml
```
After reading the project configuration, prismAId will print out an estimated cost (without warranty) for running the review using the OpenAI model. The user must enter 'y' to proceed. If the user does not enter 'y', the process will exit without making any API calls, ensuring no costs are incurred.
```bash
Total cost (USD - $): 0.0035965
Do you want to continue? (y/n): 
```

## Specifications
- **Review protocol**: Designed to support any literature review protocol, but our preference is for [Prisma 2020](https://www.prisma-statement.org/prisma-2020) (hence the project name).
- **Compatibility**: Compatible with Windows, MacOS, and Linux operating systems, on AMD64 and ARM64 platforms.
- **Supported LLMs**: OpenAI GPT 3.5 Turbo, GPT 4 Turbo, GPT4o, and GPT4o Mini. You need an OpenAI account and an **API key** to use prismAId.
- **Input format**: Requires TXT files for scientific papers.
- **Project Configuration**: Uses TOML files for easy project setup and parameter configuration.
- **Output format**: Outputs data in CSV and JSON formats.
- **Performance**: Designed to process extensive datasets efficiently with minimal user setup and **no coding**.
- **Programming Language**: Developed in Go.
