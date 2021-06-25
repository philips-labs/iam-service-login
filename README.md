# IAM Service login action

<!-- action-docs-description -->
## Description

Login to IAM using service credentials


<!-- action-docs-description -->

<!-- action-docs-inputs -->
## Inputs

| parameter | description | required | default |
| - | - | - | - |
| service-id | The IAM service id | `true` | notset |
| private-key | The IAM service private key | `true` | notset |
| region | The IAM region to log into | `true` | us-east |
| environment | The IAM environment to log into | `true` | client-test |



<!-- action-docs-inputs -->

<!-- action-docs-outputs -->
## Outputs

| parameter | description |
| - | - |
| token | The IAM token |
| message | An informational message about the request |



<!-- action-docs-outputs -->

<!-- action-docs-runs -->
## Runs

This action is an `docker` action.


<!-- action-docs-runs -->

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
