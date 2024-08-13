---
title: Prompt Design
layout: default
---

# Prompt Design
**Section 2 and 3** of the project configuration file define the prompts used to run the generative AI models to extract the information researchers are looking for. This is the key of a review project andthe prismAId robust approach to this part enables the many [Open Science](open-science) advantages provided by the tool.

## Section 2: 'Prompt' Details
The `[prompt]` section is aimed at defining the building blocks of the prompt, ensuring high accuracy in information extraction and minimizing hallucinations and misinterpretations.

### Logic of the Prompt Section
- The prompt section allows the user providing clear instructions and context to the AI model.
- The prompt structure is made of these blocks:
![prompt structure](https://raw.githubusercontent.com/ricboer0/prismAId/main/figures/prompt_struct.png)
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

### Examples and Explanation of Entries
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

## Section 3: 'Review' Details
The `[review]` section is focused on defining the information to be extracted from the text. It outlines the structure of the JSON file to be returned by the LLM, specifying the keys and possible values for the extracted information.

### Logic of the Review Section
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
### Examples and Explanation of Entries
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
