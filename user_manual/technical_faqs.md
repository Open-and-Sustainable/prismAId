# ![logo](https://raw.githubusercontent.com/ricboer0/prismAId/main/figures/prismAId_logo.png) prismAId - Technical FAQs
***
## List of FAQs

- [Will I always get the same answer if I set the model temperature to zero?](#q-will-i-always-get-the-same-answer-if-i-set-the-model-temperature-to-zero)
- [Does noise (or how much information is hidden in the literature to be reviewed) have an impact?](#q-does-noise-or-how-much-information-is-hidden-in-the-literature-to-be-reviewed-have-an-impact)
- [What happens if the literature to be reviewed says something different from the data used to train the model?](#q-what-happens-if-the-literature-to-be-reviewed-says-something-different-from-the-data-used-to-train-the-model)
- [Are there reasoning biases I should expect when analyzing literature with generative AI models?](#q-are-there-reasoning-biases-i-should-expect-when-analyzing-literature-with-generative-ai-models)
- [Is it always better to analyze literature by extracting one piece of information at a time (one piece of information per prismAId project)?](#q-is-it-always-better-to-analyze-literature-by-extracting-one-piece-of-information-at-a-time-one-piece-of-information-per-prismaid-project)

***
### Q: Will I always get the same answer if I set the model temperature to zero?

**A:** No, setting the model temperature to zero does not guarantee identical answers every time. Generative AI models, including GPT-4, can exhibit some variability even with a temperature setting of zero. While the temperature parameter influences the randomness of the output, setting it to zero aims to make the model more deterministic. However, GPT-4 and similar models are sparse mixture-of-experts models, meaning they may still show some probabilistic behavior at higher levels.

This probabilistic behavior becomes more pronounced when the prompts are near the maximum token limit. In such cases, the content within the model's attention window may change due to space constraints, leading to different outputs. Additionally, there are other mechanisms within the model that can affect its determinism.

Nevertheless, using a lower temperature is a good strategy to minimize probabilistic behavior. During the development phase of your project, repeating prompts can help achieve more consistent and replicable answers, contributing to a robust reviewing process. Robustness of answers could further be tested by modifying the sequence order of prompt building blocks, i.e., the order by which information is presented to the model.

**Further reading:** [https://doi.org/10.48550/arXiv.2308.00951](https://doi.org/10.48550/arXiv.2308.00951)
***
### Q: Does noise (or how much information is hidden in the literature to be reviewed) have an impact?

**A:** Yes, the presence of noise and the degree to which information is hidden significantly impact the quality of information extraction. The more obscured the information and the higher the noise level in the prompt, the more challenging it becomes to extract accurate and high-quality information. 

During the project configuration development phase, thorough testing can help identify the most effective prompt structures and content. This process helps in refining the prompts to minimize noise and enhance clarity, thereby improving the ability to find critical information even when it is well-hidden, akin to finding a "needle in the haystack."

**Further reading:** [https://doi.org/10.48550/arXiv.2404.08865](https://doi.org/10.48550/arXiv.2404.08865)
***
### Q: What happens if the literature to be reviewed says something different from the data used to train the model?

**A:** This is a challenge that cannot be completely avoided. We do not have full transparency on the exact data used for training the model. If the literature and the training data conflict, information extraction from the literature could be biased, transformed, or augmented by the training data. 

While focusing strictly on providing results directly from the prompt can help minimize these risks, they cannot be entirely eliminated. These biases must be carefully considered, especially when reviewing information on topics that lack long-term or widely agreed-upon consensus. These biases are similar to those of human reviewers and are very difficult to control. 

However, the prismAId ability to replicate reviews and experiment with different prompts provides an additional tool for checking these biases, offering an advantage over traditional human reviewers.

**Further reading:** [https://doi.org/10.48550/arXiv.2404.08865](https://doi.org/10.48550/arXiv.2404.08865)
***
### Q: Are there reasoning biases I should expect when analyzing literature with generative AI models?

**A:** Yes, similar to human reasoning biases, AI models trained on human texts can replicate these biases and may lead to false information extraction if the prompts steer them in that direction. This is because the models learn from and mimic the patterns found in the training data, which includes human reasoning biases. 

A good strategy to address this in prismAId is to ensure prompts are carefully crafted to be as precise, neutral, and unbiased as possible. prismAId's prompt structure supports the minimization of these problems, but there is still no guarantee against misuse by researchers.

**Further reading:** [https://doi.org/10.1038/s43588-023-00527-x](https://doi.org/10.1038/s43588-023-00527-x)
***
### Q: Is it always better to analyze literature by extracting one piece of information at a time (one piece of information per prismAId project)?

**A:** Yes, creating a separate prismAId project for each piece of information to be analyzed is a viable and highly effective approach. The main advantage is that it allows you to tailor the prompt structure and content, effectively guiding the AI model to provide accurate answers. Adding multiple information retrieval tasks within a single prismAId project requires writing much longer prompts, which can lead to more complex, potentially noisier, and confused requests for the AI model. 

The only drawback is the higher cost incurred from using the model API. Separating two questions into distinct projects approximately doubles the cost of the analysis, as most of the tokens in each project are comprised of the literature text. Therefore, the only constraint to quality is the budget.

**Further reading:** [OpenAI API Prices](https://openai.com/api/pricing/) - [https://doi.org/10.48550/arXiv.2404.08865](https://doi.org/10.48550/arXiv.2404.08865)
***
