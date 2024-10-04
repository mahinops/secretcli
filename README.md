# Secret CLI

**Secret CLI** is a command-line tool designed for securely managing secrets and sensitive information. It provides a simple and efficient interface for creating, storing, and retrieving secrets in a secure manner.

## Features

- **Secure Storage**: Safely store your secrets using encryption.
- **User Authentication**: Register and authenticate users to access their secrets.
- **Command-Line Interface**: Simple and intuitive command-line interface for managing secrets.
- **Cross-Platform Compatibility**: Works on any system with a compatible terminal.
- **Easy to Use**: Minimal setup required for quick access to your secrets.

## Installation

To install Secret CLI, follow these steps:

1. Download the Binary
    ```bash
    curl -L -o secretcli https://github.com/mahinops/secretcli/releases/download/v1.0.0/secretcli
    ```  

2. Make the Binary Executable
    ```bash
    chmod +x secretcli
    ```

3. Move the Binary to a Directory in PATH
    ```bash
    sudo mv secretcli /usr/local/bin/
    ```

4. Verify the Installation and Set Master Password
    ```bash
    secretcli
    ```

## Usage
1. List All Secrets
    ```bash
    secretcli -list
    ```

2. Add a Secret
    ```bash
    secretcli -add //Follow the prompt and insert values
    ```

3. Edit a Secret
    ```bash
    secretcli -edit <id_number>
    ```
4. Delete a Secret
    ```bash
    secretcli -del <id_number>
    ```

## Contribution
- **Contributing:** Contributions are welcome! Please open an issue or submit a pull request to enhance the functionality of Secret CLI.

## License
- **License:** This project is licensed under the MIT License. See the [LICENSE](https://github.com/mahinops/secretcli/blob/main/LICENSE) file for more details.
