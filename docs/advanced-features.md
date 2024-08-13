---
title: Advanced Features
layout: default
---


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


