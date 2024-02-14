# Polytomic Go Library

[![fern shield](https://img.shields.io/badge/%F0%9F%8C%BF-SDK%20generated%20by%20Fern-brightgreen)](https://github.com/fern-api/fern)
[![go shield](https://img.shields.io/badge/go-docs-blue)](https://pkg.go.dev/github.com/polytomic/polytomic-go)

The Polytomic Go library provides convenient access to the Polytomic API from Go.

## Requirements

This module requires Go version >= 1.13.

## Installation

Run the following command to use the Polytomic Go library in your module:

```sh
go get github.com/polytomic/polytomic-go
```

## Usage

```go
import (
  polytomic "github.com/polytomic/polytomic-go"
  polytomicgoclient "github.com/polytomic/polytomic-go/client"
  "github.com/polytomic/polytomic-go/option"
)

client := polytomicgoclient.NewClient(
  option.WithToken("<YOUR_AUTH_TOKEN>"),
)

response, err := client.Organization.Create(
  context.TODO(),
  &polytomic.V2CreateOrganizationRequestSchema{
    Name: "acme",
  },
)
```

## Timeouts

Setting a timeout for each individual request is as simple as using the standard
`context` library. Setting a one second timeout for an individual API call looks
like the following:

```go
ctx, cancel := context.WithTimeout(context.TODO(), time.Second)
defer cancel()

response, err := client.Organization.Create(
  context.TODO(),
  &polytomic.V2CreateOrganizationRequestSchema{
    Name: "acme",
  },
)
```

## Request Options

A variety of request options are included to adapt the behavior of the library, which includes
configuring authorization tokens, or providing your own instrumented `*http.Client`. Both of
these options are shown below:

```go
client := polytomicgoclient.NewClient(
  option.WithToken("<YOUR_AUTH_TOKEN>"),
  option.WithHTTPClient(
    &http.Client{
      Timeout: 5 * time.Second,
    },
  ),
)
```

These request options can either be specified on the client so that they're applied on _every_
request (shown above), or for an individual request like so:

```go
response, err := client.Organization.Create(
  context.TODO(),
  &polytomic.V2CreateOrganizationRequestSchema{
    Name: "acme",
  },
)
```

> Providing your own `*http.Client` is recommended. Otherwise, the `http.DefaultClient` will be used,
> and your client will wait indefinitely for a response (unless the per-request, context-based timeout
> is used).

## Automatic Retries

The Polytomic Go client is instrumented with automatic retries with exponential backoff. A request will be
retried as long as the request is deemed retriable and the number of retry attempts has not grown larger
than the configured retry limit (default: 2).

A request is deemed retriable when any of the following HTTP status codes is returned:

- [408](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/408) (Timeout)
- [409](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/409) (Conflict)
- [429](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/429) (Too Many Requests)
- [5XX](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/500) (Internal Server Errors)

You can use the `option.WithMaxAttempts` option to configure the maximum retry limit to
your liking. For example, if you want to disable retries for the client entirely, you can
set this value to 1 like so:

```go
client := polytomicgoclient.NewClient(
  option.WithMaxAttempts(1),
)
```

This can be done for an individual request, too:

```go
response, err := client.Organization.Create(
  context.TODO(),
  &polytomic.V2CreateOrganizationRequestSchema{
    Name: "acme",
  },
  option.WithMaxAttempts(1),
)
```

## Errors

Structured error types are returned from API calls that return non-success status codes. For example,
you can check if the error was due to an unauthorized error (i.e. status code 401) with the following:

```go
response, err := client.Organization.Create(
  context.TODO(),
  &polytomic.V2CreateOrganizationRequestSchema{
    Name: "invalid organization name",
  },
)
if err != nil {
  if unauthorizedErr, ok := err.(*polytomic.UnauthorizedError);
    // Do something with the unauthorized error ...
  }
  return err
}
```

These errors are also compatible with the `errors.Is` and `errors.As` APIs, so you can access the error
like so:

```go
response, err := client.Organization.Create(
  context.TODO(),
  &polytomic.V2CreateOrganizationRequestSchema{
    Name: "invalid organization name",
  },
)
if err != nil {
  var unauthorizedErr *polytomic.UnauthorizedError
  if errors.As(err, unauthorizedErr) {
    // Do something with the unauthorized error ...
  }
  return err
}
```

If you'd like to wrap the errors with additional information and still retain the ability
to access the type with `errors.Is` and `errors.As`, you can use the `%w` directive:

```go
response, err := client.Organization.Create(
  context.TODO(),
  &polytomic.V2CreateOrganizationRequestSchema{
    Name: "invalid organization name",
  },
)
if err != nil {
  return fmt.Errorf("failed to create organization: %w", err)
}
```

## Contributing

While we value open-source contributions to this SDK, this library is generated programmatically.
Additions made directly to this library would have to be moved over to our generation code,
otherwise they would be overwritten upon the next generated release. Feel free to open a PR as
a proof of concept, but know that we will not be able to merge it as-is. We suggest opening
an issue first to discuss with us!

On the other hand, contributions to the `README.md` are always very welcome!
