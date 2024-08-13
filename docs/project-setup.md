---
title: Project Setup
layout: default
---

# Project Setup

## Walktrhrough
1. **OpenAI Account and API Key:**
    - Create an OpenAI account at [OpenAI](https://www.openai.com/).
    - Generate an API key from the OpenAI dashboard.
2. **Download Executable:**
    - Visit the [releases](https://github.com/Open-and-Sustainable/prismAId/releases) section of the prismAId GitHub repository.
    - Download the appropriate executable for your operating system and platform.
3. **Prepare Papers for Review:**
    - Ensure that all papers to be reviewed are in .txt format.
    - Papers in html can be saved as text. To convert PDFs to .txt, there are many good options. A good one is the Python solution provided by pdfminer: instructions can be found [here](https://pdfminersix.readthedocs.io/en/latest/).

## Literature Review Requirements
- Follow protocols for literature search and identification, for instance as outlined in [PRISMA 2020](https://doi.org/10.1136/bmj.n71).
- Remove unnecessary elements from the articles. For example, the list of references usually does not provide relevant information. Similarly, the abstract and introductory parts often may (or should) be removed. Reviewing a review should be done with particular care and only if necessary.
- Unnecessary parts of text waste resources and increase analysis costs without any additional value. Actually, uneccessary parts could [negatively affect](https://arxiv.org/abs/2404.08865) model performance.

## How to Configure a Review Project
Prepare a project configuration file in [TOML](https://toml.io/en/) following the 3-sections structure, explanations, and suggestions presented in the [template.toml](https://github.com/Open-and-Sustainable/prismAId/blob/main/projects/template.toml) and here.

**Section 1** is introduced below focusing on the basic settings to configure a project. **Section 2 and 3** of the configuraiton file are discussed in [Prompt Design](prompt-design).

Additional parameters in **Section 1** can be used to activate the most advanced features of prismAId. these are discussed in [Advanced Features](advanced-features).

## Section 1: 'Project' Details

### Project Information:
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

### Configuration Details:
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

### LLM Configuration:
```toml
[project.llm]
provider = "OpenAI"
api_key = ""
model = ""
temperature = 0.2
tpm_limit = 0
```
- The `[project.llm]` section includes parameters for managing the use of the LLM:
  - `provider`: Currently irrelevant as only OpenAI is supported.
  - `api_key`: The API key can be specified here for tracking project-specific keys. If not provided, the software will look for the key in environment variables.
  - `model`: Determines the model to use. Options are:
    - Leave empty `''` for cost optimization (automatically selects the cheapest model based on token limits).
    - `gpt-4o-mini`, `gpt-4o` or `gpt-4-turbo` or `gpt-3.5-turbo` for specific model selection.
  - `temperature`: A value between 0 and 1 to control randomness. A lower value ensures replicability and accurate responses.
  - `tpm_limit`: Specifies the maximum number of tokens per minute that can be processed. The default value is `0`, which indicates that there is no delay in processing prompts by prismAId. If set to a non-zero value, this parameter should reflect the minimum tokens per minute allowed by the OpenAI API for your specific model(s) and user tier. To determine the appropriate TPM limit for your use case, consult the TPM limits section in the [OpenAI API documentation](https://platform.openai.com/settings/organization/limits).


