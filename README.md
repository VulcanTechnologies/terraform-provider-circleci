<!---
Copyright (c) 2020 Vulcan, Inc.
All rights reserved.

This Source Code Form is subject to the terms of the Mozilla Public
License, v. 2.0. If a copy of the MPL was not distributed with this
file, You can obtain one at https://mozilla.org/MPL/2.0/.
-->

# Terraform Provider CircleCI

## Context
Terraform best practices suggest using an existing API client library rather than commingling client logic with Terraform provider logic. At this time, there does not appear to be a CircleCI v2 API client library. Fortunately, CircleCI publishes their [OpenAPI v3 spec](https://circleci.com/api/v2/openapi.json). This allows for the generation of client libraries. Although the resulting client code is included in this repository, such an approach seems spiritually consistent with the principle of separating the client logic from the provider. Of course, an official API client library would be better. If you want to see this happen, please upvote [this item](https://ideas.circleci.com/api-feature-requests/p/provide-an-official-go-client-library-for-v2-api) on CircleCI's ideas board.

As of this writing, CircleCI's OpenAPI spec needs cleansing before generating the client:
* removing preview paths
* making accommodation for the missing, non-preview `Project` tag
* removing `anyOf` from the spec because it leads to invalid go code [GitHub issue](https://github.com/OpenAPITools/openapi-generator/issues/2164)

## Prerequisites
* [GNUmake](http://www.gnu.org/software/make/) 4.1 or greater- in order to use any make targets provided for convenience.
* [jq](https://stedolan.github.io/jq/)- used while cleansing CircleCI's OpenAPI spec.
* [docker](https://www.docker.com/) or [podman](https://podman.io/)- used in the make recipes which validate the cleansed OpenAPI spec and generate go code from it.
* [go](https://golang.org/) 1.15

## Refresh the Spec and Client

```shell
$  make generate_spec
$  make generate_client
```

Or if you use podman rather than docker:

```shell
$  make generate_spec container_runtime=podman
$  make generate_client container_runtime=podman
```

## Build the provider

```shell
$ make build
```

## Test the provider

### Unit Tests

```shell
$ make test
```

### Acceptance Tests

```shell
$ make acceptance_test
```

Note that the tests will fail unless you have set the environment variable `CIRCLECI_API_KEY` with a valid API key.
