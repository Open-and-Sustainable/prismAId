---
title: For Developers
layout: default
---

### Contributing
Contributions are welcome! If you'd like to improve prismAId, please create a new branch in the repo and submit a pull request. We encourage you to submit issues for bugs, feature requests, and suggestions. Please make sure to adhere to our coding standards and commit guidelines found in [`CONTRIBUTING.md`](CONTRIBUTING.md). Please also adhere to our [`CODE_OF_CONDUCT.md`](CODE_OF_CONDUCT.md.md).

### Software Dependencies

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
