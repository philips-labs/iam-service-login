# IAM Service login action

<!-- action-docs-description -->

<!-- action-docs-inputs -->

<!-- action-docs-outputs -->

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
