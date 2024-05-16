# <p align="center"><img src="https://github.com/ricboer0/prismAId/blob/main/figures/prismAId_logo.png" alt="logo" width="50"/></p>$$prism{\color{red}A}{\color{blue}I}d$$
# <p align="center">User Manual</p>

## 1. Introduction
### Scope
- PrismAId is a software tool designed to leverage the capabilities of Large Language Models (LLMs) or AI Foundation models in understanding text content for conducting systematic reviews of scientific literature.
- It aims to make the systematic review process easy, requiring no coding.
- PrismAId is designed to be much faster than traditional human-based approaches, offering also a high-speed software implementation.
- It ensures full replicability. Unlike traditional methods, which rely on subjective interpretation and classification of scientific concepts, prismAId addresses the primary issue of replicability in systematic reviews.
- Though running reviews with PrismAId incurs costs associated with using AI models, these costs are limited and lower than alternative approaches such as fine-tuning models or developing ad hoc on-premises models, which also complicate replicability. Indicatively, the cost of extracting information from a paper, as of today, can vary between half a cent to 25 cents (USD or EUR).
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
  - PrismAId simplifies the creation of rigorous and replicable prompts to extract information through the AI model API.
  - The data flow of prismAId is embedded in protocol-based approaches to reviews:
    - Initially, there is a selection of literature to be analyzed through detailed steps. These are defined by protocols and are easily replicable. 
    - Next, the content of these papers is classified, which is where prismAId comes into play.
  - PrismAId allows for parsing the selected literature and extracting the information the researcher is interested in. AI models do not know fatigue and are much faster than humans.

## 3. Technical Requirements
### Hardware and Software Requirements
- Detailed specifications:
  - Users should have an OpenAI account and generate an API key to use the system. Cost management is explained below.
  - Users need to download the executable for their OS and platform from the releases section of this project on GitHub.
  - Users should have the papers to be reviewed in .txt format. PDFs can be converted into .txt using various methods; we suggest using packages such as pdfminer.
  
- Installation steps for necessary software:
  1. **OpenAI Account and API Key:**
     - Create an OpenAI account at [OpenAI](https://www.openai.com/).
     - Generate an API key from the OpenAI dashboard.
  2. **Download Executable:**
     - Visit the releases section of the prismAId GitHub repository.
     - Download the appropriate executable for your operating system and platform.
  3. **Prepare Papers for Review:**
     - Ensure that all papers to be reviewed are in .txt format.
     - Papers in html can be saved as text, to convert PDFs to .txt, you can use packages such as pdfminer. Instructions for using pdfminer can be found [here](https://pdfminersix.readthedocs.io/en/latest/).

## 4. Literature Review Requirements
### Literature Identification and Preparation
- Follow protocols for literature identification, for instance as outlined in [PRISMA 2020](https://doi.org/10.1136/bmj.n71).
- Remove unnecessary elements from the articles. For example, the list of references usually does not provide relevant information. Similarly, the abstract and introductory parts often may (or should) be removed. Reviewing a review should be done with particular care and only if necessary.
- Unnecessary parts of text waste resources and increase analysis costs without any additional value. Actually, uneccessary parts could [negatively affect](https://arxiv.org/abs/2404.08865) model performance.

## 5. Project Configuration
### Way to Configure a Project, Step by Step
- Prepare a project configuratio file in TOML following the sections and conventions presented in the [template.toml](../projects/template.toml) and here.
- The first section of the toml
```toml
[project]
name = "Use of LLM for systematic review"
author = "John Doe"
version = "1.0"
```
- 

```toml
[project.configuration]
input_directory = "/path/to/txt/files" # where the .txt files are
results_file_name = "/path/to/save/results" # where results will be saved, the path must exists, file extension will be added
output_format = "json"  # Could be "csv" [default] or "json"
log_level = "low" # Could be "low" [default], "medium" showing entries on stdout, or "high" saving entries on file, see user manual for details
```

```toml
[project.llm]
provider = "OpenAI" # Only OpenAI is supported so far, this option is actually ignored.
api_key = "" # If left empty, it will search for an API key in env variables. Adding a key here is useful for tracking costs per prokect through project keys
model = "" # Could be 'gpt-4-turbo', 'gpt-3.5-turbo', or '' [default]. Leave empty string (or remove key) if you want cost optimizatoin: it will switch between ChatGPT4 Turbo and ChatGPT3.5 Turbo according to the cost of the model ad the limits on input tokens
temperature = 0.2 # Between 0 and 1. Low model temperature to decrease randomness and ensure replicability
batch_execution = "no" # Not yet implemented, this option is actually ignored.

```

```toml
[prompt]
persona = "Some text telling the model what role should be played." # Personas help in setting the expectation on the model role
task = "You are asked to map the concepts discussed in a scientific paper attached here." # This is the task that needs to be solved
expected_result = "You should output a JSON object with the following keys and possible values: " # This introduces the structure of the output in JSON as specified below in the [review] section
failsafe = "If the concepts neither are clearly discussed in the document nor they can be deduced from the text, respond with an empty '' value." # This is the fail-safe option to ask the model to not force answers in categories provided. PArticularly useful if values to keys below are nto complete.
definitions = "'Interest rate' is the percentage charged by a lender for borrowing money or earned by an investor on a deposit over a specific period, typically expressed annually." # This is a chance to define the concepts we are asking to the model, to avoid misconceptions.
example = "" # This is a chance to provide an example of the concepts we are asking to the model, to avoid misconceptions.
```

```toml
[review] # Review items -- defining the knowledge map that needs to be filled in
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


- Configuration settings.
- Verifying the setup.

## 6. Using the System
### Prompt Engineering and Examples
- Definition and significance.
- Sample prompts with results.
- Crafting your own prompts.

## 7. Cost Management
### Managing Costs
- Cost components.
- Tips for cost reduction.
- Tracking and reporting usage.

## 8. Best Practices
### Summary of Golden Rules of PrismAId
- Do’s and Don’ts.
- Performance optimization tips.


Current foundation models have enough knowledge to understand quite complex scientific concepts  

While model fine tuning and training of as hoc models is possible, it is extremely expensive both because of developing the training set and for actually running the training on dedicated hardware.  

Prompt engineering protocols can be defined and supported to ensure both accuracy and replicability  

The prompt protocol  

One shot

Cleaned text also because of confusion

Clear definition of the output/task + text to be parsed

## How much does prismAId cost?
Cost estimation  
Cost optimization and batch  
Support for project cost
