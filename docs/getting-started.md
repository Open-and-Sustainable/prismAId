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

## 1. Introduction
### Scope
- prismAId is a software tool designed to leverage the capabilities of Large Language Models (LLMs) or AI Foundation models in understanding text content for conducting systematic reviews of scientific literature.
- It aims to make the systematic review process easy, requiring no coding.
- prismAId is designed to be much faster than traditional human-based approaches, offering also a high-speed software implementation.
- It ensures full replicability. Unlike traditional methods, which rely on subjective interpretation and classification of scientific concepts, prismAId addresses the primary issue of replicability in systematic reviews.
- Though running reviews with prismAId incurs costs associated with using AI models, these costs are limited and lower than alternative approaches such as fine-tuning models or developing ad hoc on-premises models, which also complicate replicability. Indicatively, the cost of extracting information from a paper, as of today, can vary between a quarter of a cent to 10 cents (USD or EUR).
- Beneficiaries: Any scientist conducting a literature review or meta-analysis for developing projects, proposals, or manuscripts.

## 2. Project Overview
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
***
## 3. Technical Requirements
### Hardware and Software Requirements
- Detailed specifications:
  - Users should have an OpenAI account and generate an API key to use the system. Cost management is explained below.
  - Users need to download the executable for their OS and platform from the releases section of this project on GitHub.
  - Users should have the papers to be reviewed in .txt format. PDFs can be converted into .txt using various methods; we suggest using packages such as [pdfminer](https://github.com/pdfminer/pdfminer.six).
- Installation steps for necessary software:
  1. **OpenAI Account and API Key:**
     - Create an OpenAI account at [OpenAI](https://www.openai.com/).
     - Generate an API key from the OpenAI dashboard.
  2. **Download Executable:**
     - Visit the [releases](https://github.com/Open-and-Sustainable/prismAId/releases) section of the prismAId GitHub repository.
     - Download the appropriate executable for your operating system and platform.
  3. **Prepare Papers for Review:**
     - Ensure that all papers to be reviewed are in .txt format.
     - Papers in html can be saved as text. To convert PDFs to .txt, there are many good options. A good one is the Python solution provided by pdfminer: instructions can be found [here](https://pdfminersix.readthedocs.io/en/latest/).

## 4. Literature Review Requirements
### Literature Search, Identification, and Preparation
- Follow protocols for literature search and identification, for instance as outlined in [PRISMA 2020](https://doi.org/10.1136/bmj.n71).
- Remove unnecessary elements from the articles. For example, the list of references usually does not provide relevant information. Similarly, the abstract and introductory parts often may (or should) be removed. Reviewing a review should be done with particular care and only if necessary.
- Unnecessary parts of text waste resources and increase analysis costs without any additional value. Actually, uneccessary parts could [negatively affect](https://arxiv.org/abs/2404.08865) model performance.
***
## 5. Project Configuration
- Prepare a project configuration file in [TOML](https://toml.io/en/) following the 3-sections structure, explanations, and suggestions presented in the [template.toml](../projects/template.toml) and here.
### Section 1 'Project' Details
#### Project Information:
```toml
[project]
name = "Use of LLM for Systematic Review"
author = "John Doe"
version = "1.0"
```
- The first section `[project]` contains basic information about the project. This includes:
  - `name`: The name of the project.
  - `author`: The author of the project.
  - `version`: The version of the project configuration.
#### Configuration Details:
```toml
[project.configuration]
input_directory = "/path/to/txt/files"
results_file_name = "/path/to/save/results"
output_format = "json"
log_level = "low"
```
- The subsection `[project.configuration]` contains settings related to the project's execution environment:
  - `input_directory`: The directory where the .txt files to be reviewed are located.
  - `results_file_name`: The path where the results will be saved. Ensure the path exists in the filesystem.
  - `output_format`: The format of the output file, either `csv` or `json`.
  - `log_level`: The level of logging:
    - `low`: Minimal logging, making debugging difficult.
    - `medium`: Logs entries to stdout.
    - `high`: Saves logs to a file in the same directory as 'project_name.toml', named 'project_name.log'.
#### LLM Configuration:
```toml
[project.llm]
provider = "OpenAI"
api_key = ""
model = ""
temperature = 0.2
batch_execution = "no"
tpm_limit = 0
```
- The `[project.llm]` section includes parameters for managing the use of the LLM:
  - `provider`: Currently irrelevant as only OpenAI is supported.
  - `api_key`: The API key can be specified here for tracking project-specific keys. If not provided, the software will look for the key in environment variables.
  - `model`: Determines the model to use. Options are:
    - Leave empty `''` for cost optimization (automatically selects the cheapest model based on token limits).
    - `gpt-4o-mini`, `gpt-4o` or `gpt-4-turbo` or `gpt-3.5-turbo` for specific model selection.
  - `temperature`: A value between 0 and 1 to control randomness. A lower value ensures replicability and accurate responses.
  - `batch_execution`: Not yet supported. Once implemented, it will allow running API calls with a delay for cost savings. Results will need to be retrieved from the OpenAI platform differently.
  - `tpm_limit`: Specifies the maximum number of tokens per minute that can be processed. The default value is `0`, which indicates that there is no delay in processing prompts by prismAId. If set to a non-zero value, this parameter should reflect the minimum tokens per minute allowed by the OpenAI API for your specific model(s) and user tier. To determine the appropriate TPM limit for your use case, consult the TPM limits section in the [OpenAI API documentation](https://platform.openai.com/settings/organization/limits).


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


