# KASPIN-API
A short and simple API functions for Nicepay

## Installation
1. Clone the repo:
    ```bash
    git clone https://github.com/dickygiancini/kaspin-api
    ```
2. Install all the dependencies
    ```bash
    cd kaspin-api
    go mod download
    ```

3. Config:
    Rename `.env-example` to `.env` and setup all the necesary environment

4. Build:
    After that, go run:
    ```bash
    go build
    ```

5. Running
    After finished building, run the executable file in terminal and go to `localhost:3000/{any-link}`

## Available Routes
There are 3 routes to be tested:
1. `localhost:3000/register`
    This will register a new transaction to NICEPAY
2. `localhost:3000/payment`
    This will send a payment request to NICEPAY
3. `localhost:3000/status-inquiry`
    This will check the status in NICEPAY