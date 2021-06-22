# IAM Service login action

This action logs into HSDP IAM with service credentials

## Inputs

## `region`

**Required** The IAM region to log in. Default `"us-east"`

## `environment`

**Required** The IAM environment to log in. Default `"client-test"`

## `service-id`

**Required** The service ID to use

## `private-key`

**Required** The private key of the service ID

## Outputs

## `token`

The IAM token

## `message`

Informational message on the output

## Example usage

```yml
on: [push]

name: test service tokens

jobs:
  test_token:
    name: Test service tokens
    runs-on: ubuntu-latest

    steps:
      - uses: philips-labs/iam-service-login@v1
        id: service
        with:
          region: us-east
          environment: client-test
          service-id: ${{ secrets.SERVICE_ID }}
          private-key: ${{ secrets.PRIVATE_KEY }}
      - name: Output masked token
        run: echo "The token is ${{ steps.service.outputs.token }}"
```
