---
title: Advanced Features
layout: default
---

# Advanced Features

## Rate Limits
We enforce usage limits for models through two primary parameters specified in **Section 1** of the project configuration:

- **`tpm_limit`**: Defines the maximum number of tokens that the model can process per minute.
- **`rpm_limit`**: Specifies the maximum number of requests that the model can handle per minute.

For both parameters, a value of `0` is the default and is used if the parameter is not specified in the configuration file. The default value has a special meaning: no delay will be applied. However, if positive numbers are provided, the algorithm will compute delays and wait times between requests to the API accordingly.

Please note that we **do not support automatic enforcement of daily request limits**. If your usage tier includes a maximum number of requests per day, you will need to monitor and manage this limit manually.

On [OpenAI](https://platform.openai.com/docs/guides/rate-limits/usage-tiers?context=tier-one), for example, as of August 2024 users in tier 1 are subject to the following rate limits:
| Model          | RPM  | RPD    | TPM     | Batch Queue Limit |
|----------------|------|--------|---------|-------------------|
| gpt-4o         | 500  | -      | 30,000  | 90,000            |
| gpt-4o-mini    | 500  | 10,000 | 200,000 | 2,000,000          |
| gpt-4-turbo    | 500  | -      | 30,000  | 90,000            |
| gpt-3.5-turbo  | 3,500| 10,000 | 200,000 | 2,000,000          |

On [GoogleAI](https://ai.google.dev/pricing), as of August 2024 free of charge users are subject to the limits:
| Model           | RPM  | RPD   | TPM       |
|-----------------|------|-------|-----------|
| Gemini 1.5 Flash | 15   | 1,500 | 1,000,000 |
| Gemini 1.5 Pro   | 2    | 50    | 32,000    |
| Gemini 1.0 Pro   | 15   | 1,500 | 32,000    |

## Cost Minimization
In **Section 1** of the project configuration:
 - `model`: Determines the model to use. Options are:
    - Leave empty `''`

- The cost of using OpenAI models is calculated based on [tokens](https://help.openai.com/en/articles/4936856-what-are-tokens-and-how-to-count-them).
- prismAId utilizes a [library](https://github.com/pkoukk/tiktoken-go) to compute the input tokens for each single-shot prompt before actually executing the call using another [library](https://github.com/sashabaranov/go-openai). Based on the information provided by OpenAI, the cost of each input token for the different models is used to compute the total cost of the review. This estimated cost is presented to the user, allowing them to decide whether to proceed with the analysis and incur the associated cost.
- Concise but complete prompts are both cost-effective and efficient in information extraction. Unnecessary text increases costs and may introduce noise, negatively affecting the performance of AI models. While additional explanations and definitions in the prompt engineering part may seem superfluous, they are generally limited in size and do not significantly impact costs.
- By using a project API key, it is possible to track the cost of each project on the OpenAI [dashboard](https://platform.openai.com/usage).
- **The cost assessment function is indicative.**
  - We strive to maintain up-to-date data for cost estimation, though our estimations currently pertain only to the input aspect of AI model usage. As such, we cannot guarantee precise assessments.
  - Tests should be conducted first, and costs should be estimated more precisely by analyzing the data from the OpenAI [dashboard](https://platform.openai.com/usage).

## Batch Execution
In **Section 1** of the project configuration:
  - `batch_execution`: Not yet supported. Once implemented, it will allow running API calls with a delay for cost savings. Results will need to be retrieved from the OpenAI platform differently.


