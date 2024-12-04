Create a key pair for JWT using Ed25519 algorithm:

```shell
openssl genpkey -algorithm Ed25519 -out private_key.pem
openssl pkey -in private_key.pem -pubout -out public_key.pem
```