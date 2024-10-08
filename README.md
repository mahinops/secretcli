# Secret CLI

![GitHub stars](https://img.shields.io/github/stars/mahinops/secretcli?style=social)
![GitHub release (latest by date)](https://img.shields.io/github/v/release/mahinops/secretcli?color=blue)
[![Go Report Card](https://goreportcard.com/badge/github.com/mahinops/secretcli)](https://goreportcard.com/report/github.com/mahinops/secretcli)
![GitHub](https://img.shields.io/github/license/mahinops/secretcli?color=brightgreen)
![GitHub Workflow Status](https://img.shields.io/github/actions/workflow/status/mahinops/secretcli/release.yml?branch=main&color=success)
![Go Version](https://img.shields.io/github/go-mod/go-version/mahinops/secretcli?color=blue)
![GitHub contributors](https://img.shields.io/github/contributors/mahinops/secretcli)
![GitHub issues](https://img.shields.io/github/issues/mahinops/secretcli)
![GitHub forks](https://img.shields.io/github/forks/mahinops/secretcli?style=social)
![GitHub pull requests](https://img.shields.io/github/issues-pr/mahinops/secretcli)
![GitHub last commit](https://img.shields.io/github/last-commit/mahinops/secretcli)
![GitHub repo size](https://img.shields.io/github/repo-size/mahinops/secretcli)


**Secret CLI** is a command-line tool designed for securely managing secrets and sensitive information. It provides a simple and efficient interface for creating, storing, and retrieving secrets in a secure manner.

[![Buy Me a Coffee](https://img.buymeacoffee.com/button-api/?text=Buy%20me%20a%20coffee&emoji=&slug=your-username&button_colour=FFDD00&font_colour=000000&font_family=Cookie&outline_colour=000000&coffee_colour=ffffff)](https://www.buymeacoffee.com/mahinops)


## Features

- **Secure Storage**: Safely store your secrets using encryption.
- **User Authentication**: Register and authenticate users to access their secrets.
- **Command-Line Interface**: Simple and intuitive command-line interface for managing secrets.
- **Cross-Platform Compatibility**: Works on any system with a compatible terminal.
- **Easy to Use**: Minimal setup required for quick access to your secrets.

## Installation

To install Secret CLI, follow these steps:

1. Download the Binary from [here](https://github.com/mahinops/secretcli/releases) based on your machine architecture.

    If you are not sure. Try these commands_
    **Ubuntu:**
    ```bash
    uname -m
    ```
    - `x86_64` indicates an AMD/Intel architecture.
    - `armv7l` or aarch64 indicates an ARM architecture.


    **Mac:**
    ```bash
    uname -a
    ```

    **Windows:**
    ```bash
    echo %PROCESSOR_ARCHITECTURE%
    ```


2. Move the Binary to a Directory in PATH
    ```bash
    sudo mv <file_name> /usr/local/bin/secretcli
    ```
3. Make the Binary Executable
    ```bash
    chmod +x /usr/local/bin/secretcli
    ```

4. Verify the Installation and Set Master Password
    ```bash
    secretcli
    ```

`Note:` If you face any issue like permission error or not trusted, then follow this alternative way_

1. Clone This [Repo](https://github.com/mahinops/secretcli) in Your PC.
2. Install Golang:
    ```bash
    version >= 1.23
    ```
3. `cd` to the repo => 
    ```bash
    cd secretcli
    ```
4. Build the Project => 
    ```bash
    go build -o secretcli cmd/secretcli.go
    ```
5. Make the Generated Binary Executable => 
    ```
    chmod +x secretcli
    ```
6. Move the Binary => 
    ```
    sudo mv secretcli /usr/local/bin/
    ```
7. Verify the Installation and Set Master Password
    ```bash
    secretcli
    ```

## Usage
1. List All Secrets
    ```bash
    secretcli secret -list
    ```

2. Add a Secret
    ```bash
    secretcli secret -add //Follow the prompt and insert values
    ```

3. Edit a Secret
    ```bash
    secretcli secret -edit <id_number>
    ```
4. Delete a Secret
    ```bash
    secretcli secret -del <id_number>
    ```

## User Session Expiry
By default the user session expiry limit is set to `5 minutes`. But if any user want to update the expiry they can use the following commands_
- Update Session Expiry Command

    ```bash
    secretcli auth -set-expiry <expiry_time>
    ```

    Here, the available options for expiry are:
    - `s` -> seconds
    - `m` -> minutes
    - `h` -> hours
    - `o` -> no-expiry (never need to authenticate again)

- To set a `10 minutes expiry` user can use the following command:

    ```bash
    secretcli auth -set-expiry 10m
    ```
- To set `no expirty` user can use the following command:

    ```bash
    secretcli auth -set-expiry o
    ```

## Export Secrets List as JSON file
If you want to take a backup of the saved secret list, you can export them in a json file. To do so, use the following command_

```bash
secretcli secret -export
```
This is applicable for exporting the secrets only. User data cannot be exported. 


## Contribution
Contributions are welcome! Please open an issue or submit a pull request to enhance the functionality of Secret CLI.
Here are the reference links_

- [Contributing Guidelines](CONTRIBUTING.md) - Learn how to contribute to the project.
- [Code of Conduct](CODE_OF_CONDUCT.md) - Be sure to follow our code of conduct when participating in the project.


## License
- **License:** This project is licensed under the MIT License. See the [LICENSE](https://github.com/mahinops/secretcli/blob/main/LICENSE) file for more details.
