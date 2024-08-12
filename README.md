# ![logo](https://raw.githubusercontent.com/ricboer0/prismAId/main/figures/prismAId_logo.png) prismAId
# Open Science AI Tools for Systematic, Protocol-Based Literature Reviews

## Introduction
### Purpose
prismAId leverages generative AI to optimize the extraction and management of data from scientific literature. It extracts a structured database of the specific information researchers seek from the literature.

This tool is tailored to assist researchers by offering a simple-to-use, efficient, and replicable method for analyzing literature when conducting systematic reviews. No coding skills are required to use prismAId.

By significantly reducing the time and effort needed for data collection and analysis, prismAId enhances research efficiency. Through the use of customized prompts, prismAId automates information extraction, ensuring high accuracy and reliability. By formalizing concepts and information extraction, prismAId allows for the first time ever 100% replicable systematic literature reviews.

Most AI tools for systematic literature reviews focus on the literature search phase. While a few tools address the analysis phase, they do not fully leverage the capabilities of generative AI models. prismAId brings generative AI and Open Science where they matter most — in the analysis and data extraction phases.
![workflow](https://raw.githubusercontent.com/ricboer0/prismAId/main/figures/prismAId_workflow.png)
  
* * *
### Notes
- Ensure that you have fully read the [User Manual](user_manual/manual.md) and the [technical FAQs](user_manual/technical_faqs.md).
- For troubleshooting with software bugs, to submit requests for new functionalities, or to engage in discussions, submit an [issue](/../../issues) or participate in [GitHub Discussions](/../../discussions).

### Specifications
- **Review protocol**: Designed to support any literature review protocol, but our preference is for [Prisma 2020](https://www.prisma-statement.org/prisma-2020) (hence the project name).
- **Compatibility**: Compatible with Windows, MacOS, and Linux operating systems, on AMD64 and ARM64 platforms.
- **Supported LLMs**: OpenAI GPT 3.5 Turbo, GPT 4 Turbo, GPT4o, and GPT4o Mini. You need an OpenAI account and an **API key** to use prismAId.
- **Input format**: Requires TXT files for scientific papers.
- **Project Configuration**: Uses TOML files for easy project setup and parameter configuration.
- **Output format**: Outputs data in CSV and JSON formats.
- **Performance**: Designed to process extensive datasets efficiently with minimal user setup and **no coding**.
- **Programming Language**: Developed in Go.

* * *
## Credits
### Authors
Riccardo Boero - ribo@nilu.no

### Acknowledgments
This project was initiated with the generous support of a SIS internal project from [NILU](https://nilu.com). Their support was crucial in starting this research and development effort.

### Contributing
Contributions are welcome! If you'd like to improve prismAId, please create a new branch in the repo and submit a pull request. We encourage you to submit issues for bugs, feature requests, and suggestions. Please make sure to adhere to our coding standards and commit guidelines found in [`CONTRIBUTING.md`](CONTRIBUTING.md). Please also adhere to our [`CODE_OF_CONDUCT.md`](CODE_OF_CONDUCT.md.md).

* * *
## License
GNU AFFERO GENERAL PUBLIC LICENSE, Version 3

![license](https://www.gnu.org/graphics/agplv3-155x51.png)

* * *
## Citation
Boero, R. (2024). prismAId - Open Science AI Tools for Systematic, Protocol-Based Literature Reviews. Zenodo. https://doi.org/10.5281/zenodo.11210796

[![DOI](https://zenodo.org/badge/DOI/10.5281/zenodo.11210796.svg)](https://doi.org/10.5281/zenodo.11210796)
