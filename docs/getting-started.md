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
