# INSTRUCTION FOR RUNNING SOURCE CODE ON LOCAL

# Preparation
Note: All files had already initialized in this source code. 

## Initialize campaigns:
- In this demo phase, I initialize data from `./user-service/config/data/campaign.json`.
- For real, we must initialize by Back Office service.

## Initialize subscription plans:
- In this demo phase, I initialize data from `./purchase-service/config/data/subscription_plan.json`.
- For real, we must initialize by Back Office service.

## Initialize voucher configurations:
- In this demo phase, I initialize data from `./purchase-service/config/data/voucher_configuration.json`.
- For real, we must initialize by Back Office service.

## Create a key pair for JWT using Ed25519 algorithm:

```shell
openssl genpkey -algorithm Ed25519 -out sk.pem
openssl pkey -in private_key.pem -pubout -out pk.pem
```
- Move both keys to `./user-service/config/local_jwt` folder.
- Move `pk.pem` to `./purchase-service/config/local_jwt` folder.

---
# Step-by-step to run locally

- Install `go 1.23.0`: https://go.dev/doc/install
- Install Docker: version `4.35.1`
- Install Buf:
  - Use .sh file
  ```shell
  # Create install_buf.sh file
  # Copy and paste this content to it
  # Then save it, run "chmod +x install_buf.sh" and "source install_buf.sh"
  
  #!/bin/bash
  version="1.30.1"
  
  BIN="/usr/local/bin" && \
  VERSION=$version && \
  curl -sSL \
  "https://github.com/bufbuild/buf/releases/download/v${VERSION}/buf-$(uname -s)-$(uname -m)" \
  -o "${BIN}/buf" && \
  chmod +x "${BIN}/buf"
  ```
    - Or see instruction here: https://buf.build/docs/installation/
- Install Custom protoc plugins:
  ```shell
  # Fetch code to install
  go get -u github.com/ndtdat/protoc-gen-gorm-enum@latest
  
  # Install custom protoc plugin to generate gorm enum from proto
  go install github.com/ndtdat/protoc-gen-gorm-enum@latest
  
  # Check if protoc-gen-gorm-enum was installed
  which protoc-gen-gorm-enum
  ```
- Run base service: MySQL, Redis:
  ```shell
  cd network-social-monorepo/tools
  make run-base
  ```
- Run `user-service`:
  ```shell
  cd network-social-monorepo/user-service
  go run cmd/main.go
  ```
- Run `purchase-service`:
  ```shell
  cd network-social-monorepo/purchase-service
  go run cmd/main.go
  ```
- And now, you're able to test APIs by cURL, Postman,... Here is cURL:
  - Register:
  ```shell
  curl --location 'localhost:9091/api/user/register' \
  --header 'Content-Type: application/json' \
  --data-raw '{
  "email": "test@abc.com",
  "password": "12345678",
  "campaign_code": "SILVER-PROMOTION"
  }'
  ```
  - Login:
  ```shell
  curl --location 'localhost:9091/api/user/login' \
  --header 'Content-Type: application/json' \
  --data-raw '{
  "email": "test@abc.com",
  "password": "12345678"
  }'
  ```
  - Buy Subscription Plan, replace ****** to your access token:
  ```shell
  curl --location 'localhost:9092/api/purchase/buy' \
  --header 'Content-Type: application/json' \
  --header 'Authorization: Bearer ******' \
  --data '{
  "subscription_plan_tier": 2
  }'
  ```