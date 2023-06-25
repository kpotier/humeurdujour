# Banking
A tool to track your bank accounts for you and your family.

**WARNING: under development**

## Capabilities

### Supported banks
1. Boursorama Banque
2. ... You can add your bank by doing a PR!

## Installation

## Development
### How to compile API files
    make proto
    make swagger
    make swagger-build

### How to launch the server
    mkdir -p test/dist
    go run cmd/main.go -dev -path test

## Security concerns

By default SignUp is not in invite-only mode to allow you to create an admin account.
It is recommended to set it to invite-only mode or to add a captcha to prevent mass inserts.

## Screenshots
![screen1.png](screen1.png)
![screen2.png](screen2.png)