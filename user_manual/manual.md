# ![logo](https://raw.githubusercontent.com/ricboer0/prismAId/main/figures/prismAId_logo.png) prismAId - User Manual
***
## Table of Contents
- [1. Introduction](#1-introduction)
- [2. Project Overview](#2-project-overview)
- [3. Technical Requirements](#3-technical-requirements)
- [4. Literature Review Requirements](#4-literature-review-requirements)
- [5. Project Configuration](#5-project-configuration)
  - [Section 1 'Project' Details](#section-1-project-details)
  - [Section 2 'Prompt' Details](#section-2-prompt-details)
  - [Section 3 'Review' Details](#section-3-review-details)
- [6. Cost Management](#6-cost-management)
- [7. Further Resources](#7-further-resources)
- [8. Best Practices](#8-best-practices)


***
## 1. Introduction
### Scope
- prismAId is a software tool designed to leverage the capabilities of Large Language Models (LLMs) or AI Foundation models in understanding text content for conducting systematic reviews of scientific literature.
- It aims to make the systematic review process easy, requiring no coding.
- prismAId is designed to be much faster than traditional human-based approaches, offering also a high-speed software implementation.
- It ensures full replicability. Unlike traditional methods, which rely on subjective interpretation and classification of scientific concepts, prismAId addresses the primary issue of replicability in systematic reviews.
- Though running reviews with prismAId incurs costs associated with using AI models, these costs are limited and lower than alternative approaches such as fine-tuning models or developing ad hoc on-premises models, which also complicate replicability. Indicatively, the cost of extracting information from a paper, as of today, can vary between a quarter of a cent to 10 cents (USD or EUR).
- Beneficiaries: Any scientist conducting a literature review or meta-analysis for developing projects, proposals, or manuscripts.
***
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
***
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
    - `gpt-4o` or `gpt-4-turbo` or `gpt-3.5-turbo` for specific model selection.
  - `temperature`: A value between 0 and 1 to control randomness. A lower value ensures replicability and accurate responses.
  - `batch_execution`: Not yet supported. Once implemented, it will allow running API calls with a delay for cost savings. Results will need to be retrieved from the OpenAI platform differently.
  - `tpm_limit`: Specifies the maximum number of tokens per minute that can be processed. The default value is `0`, which indicates that there is no delay in processing prompts by prismAId. If set to a non-zero value, this parameter should reflect the minimum tokens per minute allowed by the OpenAI API for your specific model(s) and user tier. To determine the appropriate TPM limit for your use case, consult the TPM limits section in the [OpenAI API documentation](https://platform.openai.com/settings/organization/limits).

### Section 2 'Prompt' Details
The `[prompt]` section is aimed at defining the building blocks of the prompt, ensuring high accuracy in information extraction and minimizing hallucinations and misinterpretations.

#### Logic of the Prompt Section
- The prompt section allows the user providing clear instructions and context to the AI model.
- The prompt structure is made of these blocks:
![prompt structure](https://raw.githubusercontent.com/ricboer0/prismAId/main/figures/prompt_struct.PNG)
- It ensures that the model understands the role it needs to play, the task it needs to perform, and the format of the expected output.
- By providing definitions and examples, it minimizes the risk of misinterpretation and improves the accuracy of the information extracted.
- A failsafe mechanism is included to prevent the model from forcing answers when information is not available.

```toml
[prompt]
persona = "You are an experienced scientist working on a systematic review of the literature."
task = "You are asked to map the concepts discussed in a scientific paper attached here."
expected_result = "You should output a JSON object with the following keys and possible values: "
failsafe = "If the concepts neither are clearly discussed in the document nor they can be deduced from the text, respond with an empty '' value."
definitions = "'Interest rate' is the percentage charged by a lender for borrowing money or earned by an investor on a deposit over a specific period, typically expressed annually."
example = ""
```

#### Examples and Explanation of Entries
- `persona`:
  - "You are an experienced scientist working on a systematic review of the literature."
  - Personas help in setting the expectation on the model's role, providing context for the responses.
- `task`:
  - "You are asked to map the concepts discussed in a scientific paper attached here."
  - This entry defines the specific task the model needs to accomplish.
- `expected_result`:
  - "You should output a JSON object with the following keys and possible values: "
  - This introduces the expected output format, specifying that the result should be a JSON object with particular keys and values.
- `failsafe`:
  - "If the concepts neither are clearly discussed in the document nor they can be deduced from the text, respond with an empty '' value."
  - This entry provides a fail-safe mechanism to avoid forcing answers when the required information is not present, ensuring accuracy and avoiding misinterpretation.
- `definitions`:
  - "'Interest rate' is the percentage charged by a lender for borrowing money or earned by an investor on a deposit over a specific period, typically expressed annually."
  - This allows for defining specific concepts to avoid misconceptions, helping the model understand precisely what is being asked.
- `example`:
  - ""
  - This is an opportunity to provide an example of the desired output, further reducing the risk of misinterpretation and guiding the model towards the correct response.

### Section 3 'Review' Details
The `[review]` section is focused on defining the information to be extracted from the text. It outlines the structure of the JSON file to be returned by the LLM, specifying the keys and possible values for the extracted information.

#### Logic of the Review Section
- The review section defines the knowledge map that the model needs to fill in, guiding the extraction process.
- Each review item specifies a key, representing a concept or topic of interest, and possible values that the model can assign to that key.
- This structured approach ensures that the extracted information is consistent and adheres to the predefined schema.
- There can be as many review items as needed.

```toml
[review]
[review.1]
key = "interest rate"
values = [""]
[review.2]
key = "regression models"
values = ["yes", "no"]
[review.3]
key = "geographical scale"
values = ["world", "continent", "river basin"]
```
#### Examples and Explanation of Entries
- `[review]`:
  - This section header indicates the beginning of the review items configuration, which defines the structure of the knowledge map.
- `[review.1]`:
  - Defines the first item to be reviewed.
  - `key`: "interest rate"
    - The concept or topic to be extracted.
  - `values`: [""]
    - Possible values for this key. An empty string indicates that any value can be assigned.
- `[review.2]`:
  - Defines the second item to be reviewed.
  - `key`: "regression models"
    - The concept or topic to be extracted.
  - `values`: ["yes", "no"]
    - The key "regression models" can take either "yes" or "no" as its value, providing a clear binary choice.
- `[review.3]`:
  - Defines the third item to be reviewed.
  - `key`: "geographical scale"
    - The concept or topic to be extracted.
  - `values`: ["world", "continent", "river basin"]
    - The key "geographical scale" can take one of these specific values, indicating the scale of the geographical analysis.
***
## 6. Cost Management
### Managing Costs
- The cost of using OpenAI models is calculated based on [tokens](https://help.openai.com/en/articles/4936856-what-are-tokens-and-how-to-count-them).
- prismAId utilizes a [library](https://github.com/pkoukk/tiktoken-go) to compute the input tokens for each single-shot prompt before actually executing the call using another [library](https://github.com/sashabaranov/go-openai). Based on the information provided by OpenAI, the cost of each input token for the different models is used to compute the total cost of the review. This estimated cost is presented to the user, allowing them to decide whether to proceed with the analysis and incur the associated cost.
- Concise but complete prompts are both cost-effective and efficient in information extraction. Unnecessary text increases costs and may introduce noise, negatively affecting the performance of AI models. While additional explanations and definitions in the prompt engineering part may seem superfluous, they are generally limited in size and do not significantly impact costs.
- By using a project API key, it is possible to track the cost of each project on the OpenAI [dashboard](https://platform.openai.com/usage).
- **The cost assessment function is indicative.**
  - We strive to maintain up-to-date data for cost estimation, though our estimations currently pertain only to the input aspect of AI model usage. As such, we cannot guarantee precise assessments.
  - Tests should be conducted first, and costs should be estimated more precisely by analyzing the data from the OpenAI [dashboard](https://platform.openai.com/usage).
***
## 7. Further Resources
### Mastering the Art of prismAId
- Carefully read the [technical FAQs](technical_faqs.md) to avoid misusing the tool and to access emerging scientific references on issues related to the use of generative AI similar to those you may encounter in prismAId.
- We provide an additional example in the [projects](https://github.com/Open-and-Sustainable/prismAId/blob/main/projects/test.toml) directory. This includes not only the project configuration but also [input files](https://github.com/Open-and-Sustainable/prismAId/tree/main/projects/input/test) and [output files](https://github.com/Open-and-Sustainable/prismAId/tree/main/projects/output/test). The input text is extracted from a study we conducted [doi.org/10.3390/cli10020027](https://doi.org/10.3390/cli10020027).
- Multiple protocols for reporting systematic literature reviews are supported by prismAId [https://doi.org/10.1186/s13643-023-02255-9](https://doi.org/10.1186/s13643-023-02255-9). Users are encouraged to experiment and define their own prismAId methodologies.
<details>
<summary>Software Dependencies</summary>

```text
command-line-arguments
  ├ flag
  ├ fmt
  ├ io
  ├ log
  ├ os
  ├ path/filepath
  ├ prismAId/config
    ├ os
    └ github.com/BurntSushi/toml
  ├ prismAId/cost
    ├ bufio
    ├ fmt
    ├ log
    ├ os
    ├ strings
    ├ github.com/pkoukk/tiktoken-go
    ├ github.com/sashabaranov/go-openai
    ├ github.com/shopspring/decimal
    └ prismAId/config
  ├ prismAId/llm
    ├ context
    ├ encoding/json
    ├ fmt
    ├ log
    ├ github.com/sashabaranov/go-openai
    ├ prismAId/config
    └ prismAId/cost
  ├ prismAId/prompt
    ├ encoding/json
    ├ fmt
    ├ log
    ├ os
    ├ path/filepath
    ├ sort
    ├ strings
    └ prismAId/config
  └ prismAId/results
    ├ bytes
    ├ encoding/csv
    ├ encoding/json
    ├ log
    ├ os
    └ strings
```
</details>

***
## 8. Best Practices
### The Golden Rules of prismAId
1. Remove unnecessary sections from the literature to be reviewed.
2. It's better to risk repeating an explanation of the information you are seeking through examples than not defining it clearly enough.
3. If the budget allows, conduct a separate review process for each piece of information you want to extract. This allows for more detailed definitions for each information piece.
4. Try to avoid using open-ended answers and define and explain all possible answers the AI model can provide.
5. First, run a test on a single paper. Once the results are satisfactory, repeat the process with a different batch of papers. If the results are still satisfactory, proceed with the rest of the literature.
6. Focus on primary sources and avoid reviewing reviews unless it is intentional and carefully planned. Do not mix primary and secondary sources in the same review process.
7. Include the project configuration (the .toml file) in the appendix of your paper.
8. Properly cite prismAId [doi.org/10.5281/zenodo.11210796](https://doi.org/10.5281/zenodo.11210796).


