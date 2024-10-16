---
title: Advanced Features
layout: default
---

# Advanced Features

## Debugging & Validation
In **Section 1** of the project configuration, there are three parameters supporting the devleopment of projects and testing of prompt configurations.
They are:
  - `log_level`: [`low`], `medium`, or `high`.
  - `duplication`: [`no`], or `yes`.
  - `cot_justification`:  [`no`], or `yes`.
First, if debuggin level is higher than low all API responses can be inspected in details. This means that besides output files, users will be able to access, either on terminal (stdout - `log_level`: `medium`) or in a log file (`log_level`: `high`), the complete reponses and eventual errors from the API and the prismAId execution.

Second, duplication makes possible to test whether a prompt definition is clear enough. In fact, if running twice the same prompt generates different ouptut it is very likely that the prompt is not deifning the model reviewing task clearly enough. Setting `duplication`: `yes` and then checking if answers differ in the two analyses of the same manuscript is a good way to assess whether the prompt is clear enough to be used for the review project. 

Duplicating manuscripts increases the cost of the project run, but the total cost presented at the beginning of the analysis is updated accordingly to let researchers assess the cost to be incurred. Hence, for instance, with Google AI as provider and Gemini 1.5 Flash model, without duplication:
```bash
Unless you are using a free tier with Google AI, the total cost (USD - $) to run this review is at least: 0.0005352
This value is an estimate of the total cost of input tokens only.
Do you want to continue? (y/n): y
Processing file #1/1: lit_test
```
With duplication active:
```bash
Unless you are using a free tier with Google AI, the total cost (USD - $) to run this review is at least: 0.0010704
This value is an estimate of the total cost of input tokens only.
Do you want to continue? (y/n): y
Processing file #1/2: lit_test
Waiting... 30 seconds remaining
Waiting... 25 seconds remaining
Waiting... 20 seconds remaining
Waiting... 15 seconds remaining
Waiting... 10 seconds remaining
Waiting... 5 seconds remaining
Wait completed.
Processing file #2/2: lit_test_duplicate
```

Third, in order to assess if the prompt definition are not only clear but also effective in extracting the information the researcher is looking for, it is is possible to use `cot_justification`: `yes`. This will create an output `.txt` for each manuscript containing the chain of thought (CoT) justification for the answer provided. Technically, the justification is provided by the model in the same chat as the answer, and right after it.

The ouput in the justification output reports the information requested, the answer provided, the modle CoT, and eventually the relevant sentences in the manuscript reviewd, like in:
```md
- **clustering**: "no" - The text does not mention any clustering techniques or grouping of data points based on similarities.
- **copulas**: "yes" - The text explicitly mentions the use of copulas to model the joint distribution of multiple flooding indicators (maximum soil moisture, runoff, and precipitation). "The multidimensional representation of the joint distributions of relevant hydrological climate impacts is based on the concept of statistical copulas [43]."
- **forecasting**: "yes" - The text explicitly mentions the use of models to predict future scenarios of flooding hazards and damage. "Future scenarios use hazard and damage data predicted for the period 2018–2100."

```

## Rate Limits
We enforce usage limits for models through two primary parameters specified in **Section 1** of the project configuration:

- **`tpm_limit`**: Defines the maximum number of tokens that the model can process per minute.
- **`rpm_limit`**: Specifies the maximum number of requests that the model can handle per minute.

For both parameters, a value of `0` is the default and is used if the parameter is not specified in the configuration file. The default value has a special meaning: no delay will be applied. However, if positive numbers are provided, the algorithm will compute delays and wait times between requests to the API accordingly.

Please note that we **do not support automatic enforcement of daily request limits**. If your usage tier includes a maximum number of requests per day, you will need to monitor and manage this limit manually.

On [OpenAI](https://platform.openai.com/docs/guides/rate-limits/usage-tiers?context=tier-one), for example, as of August 2024 users in tier 1 are subject to the following rate limits:

<table class="table-spacing">
    <thead>
        <tr>
            <th style="text-align: left;">Model</th>
            <th style="text-align: right;">RPM</th>
            <th style="text-align: right;">RPD</th>
            <th style="text-align: right;">TPM</th>
            <th style="text-align: right;">Batch Queue Limit</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td style="text-align: left;">gpt-4o</td>
            <td style="text-align: right;">500</td>
            <td style="text-align: right;">-</td>
            <td style="text-align: right;">30,000</td>
            <td style="text-align: right;">90,000</td>
        </tr>
        <tr>
            <td style="text-align: left;">gpt-4o-mini</td>
            <td style="text-align: right;">500</td>
            <td style="text-align: right;">10,000</td>
            <td style="text-align: right;">200,000</td>
            <td style="text-align: right;">2,000,000</td>
        </tr>
        <tr>
            <td style="text-align: left;">gpt-4-turbo</td>
            <td style="text-align: right;">500</td>
            <td style="text-align: right;">-</td>
            <td style="text-align: right;">30,000</td>
            <td style="text-align: right;">90,000</td>
        </tr>
        <tr>
            <td style="text-align: left;">gpt-3.5-turbo</td>
            <td style="text-align: right;">3,500</td>
            <td style="text-align: right;">10,000</td>
            <td style="text-align: right;">200,000</td>
            <td style="text-align: right;">2,000,000</td>
        </tr>
    </tbody>
</table>


On [GoogleAI](https://ai.google.dev/pricing), as of October 2024 **free of charge** users are subject to the limits:

<table class="table-spacing">
    <thead>
        <tr>
            <th style="text-align: left;">Model</th>
            <th style="text-align: right;">RPM</th>
            <th style="text-align: right;">RPD</th>
            <th style="text-align: right;">TPM</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td style="text-align: left;">Gemini 1.5 Flash</td>
            <td style="text-align: right;">15</td>
            <td style="text-align: right;">1,500</td>
            <td style="text-align: right;">1,000,000</td>
        </tr>
        <tr>
            <td style="text-align: left;">Gemini 1.5 Pro</td>
            <td style="text-align: right;">2</td>
            <td style="text-align: right;">50</td>
            <td style="text-align: right;">32,000</td>
        </tr>
        <tr>
            <td style="text-align: left;">Gemini 1.0 Pro</td>
            <td style="text-align: right;">15</td>
            <td style="text-align: right;">1,500</td>
            <td style="text-align: right;">32,000</td>
        </tr>
    </tbody>
</table>

while **pay-as-you-go** users are subject to:

<table class="table-spacing">
    <thead>
        <tr>
            <th style="text-align: left;">Model</th>
            <th style="text-align: right;">RPM</th>
            <th style="text-align: right;">RPD</th>
            <th style="text-align: right;">TPM</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td style="text-align: left;">Gemini 1.5 Flash</td>
            <td style="text-align: right;">2000</td>
            <td style="text-align: right;">-</td>
            <td style="text-align: right;">4,000,000</td>
        </tr>
        <tr>
            <td style="text-align: left;">Gemini 1.5 Pro</td>
            <td style="text-align: right;">1000</td>
            <td style="text-align: right;">-</td>
            <td style="text-align: right;">4,000,000</td>
        </tr>
        <tr>
            <td style="text-align: left;">Gemini 1.0 Pro</td>
            <td style="text-align: right;">360</td>
            <td style="text-align: right;">30,000</td>
            <td style="text-align: right;">120,000</td>
        </tr>
    </tbody>
</table>

In September 2024 Cohere does not impose rate limits on production keys but trial keys are limited to 20 API calls per minute (refer to the official [documentation](https://docs.cohere.com/docs/rate-limits)).

Anthropic Tier 1 users have the following rate limits:
<table class="table-spacing">
    <thead>
        <tr>
            <th style="text-align: left;">Model</th>
            <th style="text-align: right;">RPM</th>
            <th style="text-align: right;">TPM</th>
            <th style="text-align: right;">TPD</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td style="text-align: left;">Claude 3.5 Sonnet</td>
            <td style="text-align: right;">50</td>
            <td style="text-align: right;">40,000</td>
            <td style="text-align: right;">1,000,000</td>
        </tr>
        <tr>
            <td style="text-align: left;">Claude 3 Opus</td>
            <td style="text-align: right;">50</td>
            <td style="text-align: right;">20,000</td>
            <td style="text-align: right;">1,000,000</td>
        </tr>
        <tr>
            <td style="text-align: left;">Claude 3 Sonnet</td>
            <td style="text-align: right;">50</td>
            <td style="text-align: right;">40,000</td>
            <td style="text-align: right;">1,000,000</td>
        </tr>
        <tr>
            <td style="text-align: left;">Claude 3 Haiku</td>
            <td style="text-align: right;">50</td>
            <td style="text-align: right;">50,000</td>
            <td style="text-align: right;">5,000,000</td>
        </tr>
    </tbody>
</table>


**PLEASE NOTE**: If you choose the cost minimization approach described below you must report in the configuration file the smallest tpm and rpm limits of the models by the provider you selected. This is the only way to ensure respecting limits since there is no authomatic check on them by prismAId and the selected model varies because of number of tokens in requests and model use prices.

## Cost Minimization
In **Section 1** of the project configuration:
 - `model`: Determines the model to use. Options are:
    - Leave empty `''`
This feature allows to always automatically select the cheapest model for the job provided by the provider selected. Please note that this may mean that different manuscripts are analyzed by different models depending on the manuscript length.

### How costs are computed
- The cost of using OpenAI models is calculated based on [tokens](https://help.openai.com/en/articles/4936856-what-are-tokens-and-how-to-count-them).
- prismAId utilizes a [library](https://github.com/pkoukk/tiktoken-go) to compute the input tokens for each single-shot prompt before actually executing the call using another [library](https://github.com/sashabaranov/go-openai). Based on the information provided by OpenAI, the cost of each input token for the different models is used to compute the total cost of the inputs to be used in the review. This estimated cost is presented to the user, allowing them to decide whether to proceed with the analysis and incur the associated cost.
- prismAId calls the Google CountTokens [API](https://cloud.google.com/vertex-ai/generative-ai/docs/model-reference/count-tokens) to compute the input tokens for each single-shot prompt before actually executing the call using a [library](https://github.com/google/generative-ai-go). Based on the information provided by Google AI, the cost of each input token for the different models is used to compute the total cost of the inputs to be used in the review.
- prismAId calls the Cohere API to compute the input tokens for each single-shot prompt before actually executing the call using a [library](https://github.com/cohere-ai/cohere-go/). Please note that different Cohere models are trained with different tokenizers. This means also that the same prompt may be transformed into different number of input tokens depending on the model used. Based on the information provided by Cohere, the cost of each input token for the different models is used to compute the total cost of the inputs to be used in the review.
- Anthropic does not release the tokenizer nor an API free endpoint for counting input tokens. Following suggestions from their own Anthropic models, prismAId estimate the number of input tokens using the OpenAI tokenizer.
- Concise but complete prompts are both cost-effective and efficient in information extraction. Unnecessary text increases costs and may introduce noise, negatively affecting the performance of AI models. While additional explanations and definitions in the prompt engineering part may seem superfluous, they are generally limited in size and do not significantly impact costs.
- By using a project API key, it is possible to track the cost of each project on the OpenAI [dashboard](https://platform.openai.com/usage), the Google AI [dashboard](https://console.cloud.google.com/billing/), or the Cohere [dashboard](https://dashboard.cohere.com/billing).
- **The cost assessment function is indicative.**
  - We strive to maintain up-to-date data for cost estimation, though our estimations currently pertain only to the input aspect of AI model usage. As such, we cannot guarantee precise assessments.
  - Tests should be conducted first, and costs should be estimated more precisely by analyzing the data from the OpenAI [dashboard](https://platform.openai.com/usage) or the Google AI [dashboard](https://console.cloud.google.com/billing/).

<div id="wcb" class="carbonbadge wcb-d"></div>
<script src="https://unpkg.com/website-carbon-badges@1.1.3/b.min.js" defer></script>