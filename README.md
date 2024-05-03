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
Simply download the executable for your OS, place it in a suitable directory, and run it from your command line:

```bash
# For Windows
./prismAId_windows_amd64.exe --project your_project.toml

# For Linux
./prismAId_linux_amd64 --project your_project.toml

# For macOS
./prismAId_darwin_amd64 --project your_project.toml
```

### Setting up a review project

Follow instructions in the [user manual](user_manual/manual.md).

* * *

## Specifications
- **Programming Language**: Developed in Go.
- **Compatibility**: Compatible with Windows, MacOS, and Linux operating systems.
- **Input format**: Requires TXT files for scientific papers.
- **Project Configuration**: Uses TOML files for easy project setup and parameter configuration.
- **Output format**: Outputs data in CSV and JSON formats.
- **Performance**: Designed to process extensive datasets efficiently with minimal user setup.

* * *

## Notes
- Ensure that your input files are formatted correctly as per the guidelines to avoid common errors.
- For troubleshooting and support, please refer to the wiki or submit an issue on GitHub.

* * *

### Authors

Riccardo Boero - ribo@nilu.no

## Contributing
Contributions are welcome! If you'd like to improve prismAId, please create a new branch in the repo and submit a pull request. We encourage you to submit issues for bugs, feature requests, and suggestions. Please make sure to adhere to our coding standards and commit guidelines found in [`CONTRIBUTING.md`](CONTRIBUTING.md).


### License
GNU AFFERO GENERAL PUBLIC LICENSE, Version 3

<img src="https://www.gnu.org/graphics/agplv3-155x51.png" alt="license" width="155"/>
