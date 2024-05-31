# ![logo](https://raw.githubusercontent.com/ricboer0/prismAId/main/figures/prismAId_logo.png) prismAId
# Open Science AI Tools for Systematic, Protocol-Based Literature Reviews
<!-- Innovate and Accelerate Science with AI: Open and Replicable Tools for Systematic, Protocol-Based Literature Reviews. -->
* * *
## Purpose
prismAId leverages generative AI to optimize the extraction and management of data from scientific literature. It extracts a structured database of the specific information researchers seek from the literature.

This tool is tailored to assist researchers by offering a simple-to-use, efficient, and replicable method for analyzing literature when conducting systematic reviews. No coding skills are required to use prismAId.

By significantly reducing the time and effort needed for data collection and analysis, prismAId enhances research efficiency. Through the use of customized prompts, prismAId automates information extraction, ensuring high accuracy and reliability. By formalizing concepts and information extraction, prismAId allows for the first time ever 100% replicable systematic literature reviews.

Most AI tools for systematic literature reviews focus on the literature search phase. While a few tools address the analysis phase, they do not fully leverage the capabilities of generative AI models. prismAId brings generative AI and Open Science where they matter most — in the analysis and data extraction phases.
![workflow](https://raw.githubusercontent.com/ricboer0/prismAId/main/figures/prismAId_workflow.PNG)
***
## Open Science
prismAId supports Open Science in many aspects:

1. **Transparency and Reproducibility**
   - prismAId ensures transparency in the analysis process, making it easier for other researchers to understand, replicate, and validate the findings.
   - prismAId removes the subjectivity of individual interpretations, making systematic literature reviews 100% reproducible.
   - As a software tool, prismAId helps maintain detailed logs and records of the analysis process, enhancing reproducibility.

2. **Accessibility and Collaboration**
   - prismAId facilitates collaboration among researchers by providing an open tool that makes it possible to share data, analysis methods, and results.
   - prismAId is open source and openly licensed. Making analysis tools openly available promotes wider participation and contribution from the scientific community.
   - prismAId releases and their source code are archived on [Zenodo](https://zenodo.org/doi/10.5281/zenodo.11210796), ensuring long-term accessibility and referencability. This helps address legacy issues for analyses conducted using prismAId, making the tool and its results open, replicable, and understandable over the long run.

3. **Efficiency and Scalability**
   - prismAId can handle large volumes of data efficiently, making the analysis phase quicker and more scalable compared to traditional methods.
   - This efficiency supports open science by allowing more comprehensive and timely reviews, reducing the time society needs to properly 'digest' scientific innovations.

4. **Quality and Accuracy**
   - By explicitly defining each piece of reviewed information through prompt configurations, prismAId enhances the quality and accuracy of data extraction and analysis, leading to more reliable systematic reviews.
   - Publishing prismAId project configuration files ensures that approaches, biases, and methods are visible and verifiable by the broader research community. By doing so, they are also reusable and extendable.

5. **Ethical Considerations and Bias Reduction**
   - Using prismAId means explicitly addressing biases and incorporating ethical considerations in its design and implementation to minimize biases in data analysis.
   - prismAId enables open science approaches, ensuring full transparency on ethical standards, with community oversight and input helping to identify and mitigate potential biases.

6. **Scientific Innovation**
   - prismAId promotes scientific innovation by fully formalizing the analysis process, creating standardized procedures that ensure consistency and accuracy in systematic reviews.
   - This formalization makes methods and procedures reusable and extendable, allowing researchers to build upon previous analyses and adapt methods to new contexts.
   - By facilitating incremental discoveries, prismAId supports the cumulative advancement of science, where each study contributes to a larger body of knowledge.
   - prismAId's commitment to open science principles ensures that all tools, methods, and data are openly accessible, fostering collaboration and rapid dissemination of innovations.
  
* * *

## Getting Started
To use prismAId, download the appropriate executable for your operating system from our [GitHub Releases](https://github.com/ricboer0/prismAId/releases) page.

### Running prismAId
Simply [download](https://github.com/ricboer0/prismAId/releases) the executable for your OS and platform, place it in a suitable directory, prepare a project configuration (.toml), and run it from your command line, e.g.:

```bash
# For Windows
./prismAId_windows_amd64.exe --project your_project.toml
```
After reading the project configuration, prismAId will print out an estimated cost (without warranty) for running the review using the OpenAI model. The user must enter 'y' to proceed. If the user does not enter 'y', the process will exit without making any API calls, ensuring no costs are incurred.
```bash
Total cost (USD - $): 0.0035965
Do you want to continue? (y/n): 
```

### Setting up a review project

Follow instructions in the [User Manual](user_manual/manual.md).

* * *

## Specifications
- **Review protocol**: Designed to support any literature review protocol, but our preference is for [Prisma 2020](https://www.prisma-statement.org/prisma-2020) (hence the project name).
- **Compatibility**: Compatible with Windows, MacOS, and Linux operating systems, on AMD64 and ARM64 platforms.
- **Supported LLMs**: OpenAI GPT 3.5 Turbo, GPT 4 Turbo, and GPT4o. You need an OpenAI account and an **API key** to use prismAId.
- **Input format**: Requires TXT files for scientific papers.
- **Project Configuration**: Uses TOML files for easy project setup and parameter configuration.
- **Output format**: Outputs data in CSV and JSON formats.
- **Performance**: Designed to process extensive datasets efficiently with minimal user setup and **no coding**.
- **Programming Language**: Developed in Go.

* * *

## Notes
- Ensure that you have fully read the [User Manual](user_manual/manual.md) and the [technical FAQs](user_manual/technical_faqs.md).
- For troubleshooting with software bugs, to submit requests for new functionalities, or to engage in discussions, submit an [issue](/../../issues) or participate in [GitHub Discussions](/../../discussions).
- Forthcoming features include support for batch execution processes, which can halve the cost of reviews at the expense of a delay of up to 24 hours.

### Contributing
Contributions are welcome! If you'd like to improve prismAId, please create a new branch in the repo and submit a pull request. We encourage you to submit issues for bugs, feature requests, and suggestions. Please make sure to adhere to our coding standards and commit guidelines found in [`CONTRIBUTING.md`](CONTRIBUTING.md). Please also adhere to our [`CODE_OF_CONDUCT.md`](CODE_OF_CONDUCT.md.md).

* * *

## Authors
Riccardo Boero - ribo@nilu.no

### Acknowledgments
This project was initiated with the generous support of a SIS internal project from [NILU](https://nilu.com). Their support was crucial in starting this research and development effort.

* * *

## License
GNU AFFERO GENERAL PUBLIC LICENSE, Version 3

![license](https://www.gnu.org/graphics/agplv3-155x51.png)

* * *

## Citation
Boero, R. (2024). prismAId - Open Science AI Tools for Systematic, Protocol-Based Literature Reviews. Zenodo. https://doi.org/10.5281/zenodo.11210796

[![DOI](https://zenodo.org/badge/DOI/10.5281/zenodo.11210796.svg)](https://doi.org/10.5281/zenodo.11210796)
