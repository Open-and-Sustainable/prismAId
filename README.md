# <p align="center"><img src="https://github.com/ricboer0/prismAId/blob/main/figures/prismAId_logo.png" alt="logo" width="50"/></p>$$prism{\color{red}A}{\color{blue}I}d$$
# <p align="center">Open Science AI Tools for Systematic, Protocol-Based Literature Reviews</p>
Innovate and Accelerate Science with AI: Open and Replicable Tools for Systematic, Protocol-Based Literature Reviews.

* * *

## Purpose
prismAId harnesses the power of AI to streamline the process of extracting and managing data from scientific literature. This tool is designed to support researchers by providing a systematic, replicable approach to literature reviews, significantly reducing the time and effort required to gather and analyze data. By automating the extraction of information using customized prompts, prismAId ensures high accuracy and efficiency, enabling users to focus more on analysis and less on data processing.

* * *

## Getting Started
To use prismAId, download the appropriate executable for your operating system from our [GitHub Releases](https://github.com/ricboer0/prismAId/releases) page.

### Running prismAId
Simply download the executable for your OS and platform, place it in a suitable directory, prepare a project configuration (.toml), and run it from your command line, e.g.:

```bash
# For Windows
./prismAId_windows_amd64.exe --project your_project.toml

# For Linux
./prismAId_linux_amd64 --project your_project.toml

# For macOS
./prismAId_darwin_amd64 --project your_project.toml
```

### Setting up a review project

Follow instructions in the [User Manual](user_manual/manual.md).

* * *

## Specifications
- **Review protocol**: Designed to support any literature review protocol, but our preference is for [Prisma 2020](https://www.prisma-statement.org/prisma-2020) (hence the project name).
- **Compatibility**: Compatible with Windows, MacOS, and Linux operating systems, on AMD64 and ARM64 platforms.
- **Supported LLMs**: OpenAI ChatGPT Turbo 3.5 and 4. You need an OpenAI account and an **API key** to use prismAId.
- **Input format**: Requires TXT files for scientific papers.
- **Project Configuration**: Uses TOML files for easy project setup and parameter configuration.
- **Output format**: Outputs data in CSV and JSON formats.
- **Performance**: Designed to process extensive datasets efficiently with minimal user setup and **no coding**.
- **Programming Language**: Developed in Go.

* * *

## Notes
- Ensure that you have fully read the [User Manual](user_manual/manual.md).
- For troubleshooting and support not covered in the [User Manual](user_manual/manual.md), submit an [issue](/../../issues) on GitHub.

* * *

## Contributing
Contributions are welcome! If you'd like to improve prismAId, please create a new branch in the repo and submit a pull request. We encourage you to submit issues for bugs, feature requests, and suggestions. Please make sure to adhere to our coding standards and commit guidelines found in [`CONTRIBUTING.md`](CONTRIBUTING.md).

### Authors

Riccardo Boero - ribo@nilu.no

### License
GNU AFFERO GENERAL PUBLIC LICENSE, Version 3

<img src="https://www.gnu.org/graphics/agplv3-155x51.png" alt="license" width="155"/>
