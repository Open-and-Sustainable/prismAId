---
title: Getting Started
layout: default
toc: true
---

# Getting Started

## Purpose and Benefits
prismAId leverages generative AI to optimize the extraction and management of data from scientific literature. It extracts a structured database of the specific information researchers seek from the literature.

This tool is tailored to assist researchers by offering a simple-to-use, efficient, and replicable method for analyzing literature when conducting systematic reviews. No coding skills are required to use prismAId.

By significantly reducing the time and effort needed for data collection and analysis, prismAId enhances research efficiency. Through the use of customized prompts, prismAId automates information extraction, ensuring high accuracy and reliability. By formalizing concepts and information extraction, prismAId allows for the first time ever 100% replicable systematic literature reviews.

Most AI tools for systematic literature reviews focus on the literature search phase. While a few tools address the analysis phase, they do not fully leverage the capabilities of generative AI models. prismAId brings generative AI and Open Science where they matter most — in the analysis and data extraction phases.
![workflow](https://raw.githubusercontent.com/ricboer0/prismAId/main/figures/prismAId_workflow.png)

## Audience

## AI Requirements

## Installation Instructions

## Initial Configuration

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


