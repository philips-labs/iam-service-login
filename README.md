# IAM Service login action

This action logs into HSDP IAM with service credentials

## Inputs

## `region`

**Required** The IAM region to log in. Default `"us-east"`

## `environment`

**Required** The IAM environment to log in. Default `"client-test"`

## `service-id`

**Required** The service ID to use

## `environment`

**Required** The private key of the service ID

## Outputs

## `token`

The IAM token

## Example usage

```yml
uses: loafoe/iam-service-login@v1
  with:
    who-to-greet: 'Mona the Octocat'
```
