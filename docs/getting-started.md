---
title: Getting Started
layout: default
---

# Getting Started

## Purpose and Benefits
prismAId leverages generative AI to optimize the extraction and management of data from scientific literature. It extracts a structured database of the specific information researchers seek from the literature.

This tool is tailored to assist researchers by offering a simple-to-use, efficient, and replicable method for analyzing literature when conducting systematic reviews. No coding skills are required to use prismAId.

By significantly reducing the time and effort needed for data collection and analysis, prismAId enhances research efficiency. Through the use of customized prompts, prismAId automates information extraction, ensuring high accuracy and reliability. By formalizing concepts and information extraction, prismAId allows for the first time ever 100% replicable systematic literature reviews.

Most AI tools for systematic literature reviews focus on the literature search phase. While a few tools address the analysis phase, they do not fully leverage the capabilities of generative AI models. prismAId brings generative AI and Open Science where they matter most — in the analysis and data extraction phases.
![workflow](https://raw.githubusercontent.com/ricboer0/prismAId/main/figures/prismAId_workflow.png)

### Scope
- prismAId is a software tool designed to leverage the capabilities of Large Language Models (LLMs) or AI Foundation models in understanding text content for conducting systematic reviews of scientific literature.
- It aims to make the systematic review process easy, requiring no coding.
- prismAId is designed to be much faster than traditional human-based approaches, offering also a high-speed software implementation.
- It ensures full replicability. Unlike traditional methods, which rely on subjective interpretation and classification of scientific concepts, prismAId addresses the primary issue of replicability in systematic reviews.
- Though running reviews with prismAId incurs costs associated with using AI models, these costs are limited and lower than alternative approaches such as fine-tuning models or developing ad hoc on-premises models, which also complicate replicability. Indicatively, the cost of extracting information from a paper, as of today, can vary between a quarter of a cent to 10 cents (USD or EUR).
- Beneficiaries: Any scientist conducting a literature review or meta-analysis for developing projects, proposals, or manuscripts.

### Description of Underlying Mechanism
- How LLMs work:
  - LLMs (Large Language Models) are AI models trained on vast amounts of text data to understand and generate human-like text.
  - These models can perform a variety of language tasks such as text completion, summarization, translation, and more.  
- Data flow and processing steps:
  - Contemporary state-of-the-art LLMs offer subscription-based API access.
  - While foundation models can be used in various ways, prismAId focuses solely on prompt engineering or prompting.
  - Prompt engineering involves crafting precise prompts to extract specific information from the AI model via the API.
  - prismAId simplifies the creation of rigorous and replicable prompts to extract information through the AI model API.
  - The data flow of prismAId is embedded in protocol-based approaches to reviews:
    - Initially, there is a selection of literature to be analyzed through detailed steps. These are defined by protocols and are easily replicable. 
    - Next, the content of these papers is classified, which is where prismAId comes into play.
  - prismAId allows for parsing the selected literature and extracting the information the researcher is interested in. AI models do not know fatigue and are much faster than humans.
  - prismAId users define the information extraction tasks using the prompt engineering template provided by prismAId.
  - prismAId utilizes multiple single-shot prompt API calls to individually parse each scientific paper.
  - prismAId processes the JSON files returned by the AI model, converting the extracted information into the user-specified format.
  - To facilitate cost management, prismAId tokenizes each single-shot prompt and estimates the execution cost, allowing users to understand the total review cost before proceeding.

## Installation
To use prismAId, you have two options.

### Option 1. Binaries 
Download the appropriate executable for your operating system and platform from our [GitHub Releases](https://github.com/open-and-sustainable/prismAId/releases) page. Using executables does not require any coding skill.
### Option 2. Go Package
You can download the prismAId Go package for developing your own software or review project. To add the package to yoru project:
```bash
go get "github.com/Open-and-Sustainable/prismaid"
```
Once added, it can be imported when needed with:
```go
import "github.com/Open-and-Sustainable/prismaid"
```
The package documentation is available on [pkg.go.dev](https://pkg.go.dev/github.com/Open-and-Sustainable/prismaid).
## Running prismAId binaries
The tool uses humaly readable project configuration files (.toml) to configure and run the reviews.

You can find a template and an example on the [GitHub repository](https://github.com/Open-and-Sustainable/prismAId/tree/main/projects).

prismAId provides an interactive terminal application guiding users in the creation of draft configuration files. This function is activated by calling binaries with the '-init' flag, for instance: 

```bash
# For Linux on Intel
./prismAId_linux_amd64 -init
```

![Terminal app for drafting project configuration file](https://raw.githubusercontent.com/ricboer0/prismAId/main/figures/terminal.gif)

Once the project configuration (.toml) is ready, the project can be executed from your command line, e.g.:

```bash
# For Windows
./prismAId_windows_amd64.exe -project your_project.toml
```

After reading the project configuration, prismAId will print out an estimated cost (without warranty) for running the review using the OpenAI model. The user must enter 'y' to proceed. If the user does not enter 'y', the process will exit without making any API calls, ensuring no costs are incurred.
```bash
Total cost (USD - $): 0.0035965
Do you want to continue? (y/n): 
```

<div id="wcb" class="carbonbadge"></div>
<script src="https://unpkg.com/website-carbon-badges@1.1.3/b.min.js" defer></script>