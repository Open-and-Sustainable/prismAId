# ![logo](https://raw.githubusercontent.com/ricboer0/prismAId/main/figures/prismAId_logo.png) prismAId - Technical FAQs

### Q: Will I always get the same answer if I set the model temperature to zero?

**A:** No, setting the model temperature to zero does not guarantee identical answers every time. Generative AI models, including GPT-4, can exhibit some variability even with a temperature setting of zero. While the temperature parameter influences the randomness of the output, setting it to zero aims to make the model more deterministic. However, GPT-4 and similar models are sparse mixture-of-experts models, meaning they may still show some probabilistic behavior at higher levels.

This probabilistic behavior becomes more pronounced when the prompts are near the maximum token limit. In such cases, the content within the model's attention window may change due to space constraints, leading to different outputs. Additionally, there are other mechanisms within the model that can affect its determinism.

Nevertheless, using a lower temperature is a good strategy to minimize probabilistic behavior. During the development phase of your project, repeating prompts can help achieve more consistent and replicable answers, contributing to a robust reviewing process. Robustness of answers could further be tested by modifying the sequence order of prompt building blocks, i.e., the order by which information is presented to the model.

**Further reading:** [http://arxiv.org/abs/2308.00951](http://arxiv.org/abs/2308.00951)

