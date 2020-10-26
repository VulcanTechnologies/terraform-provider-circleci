---
page_title: "Provider: CircleCI"
subcategory: ""
description: "Terraform provider for interacting with CircleCI's v2 API."

---

<!---
Copyright (c) 2020 Vulcan, Inc.
All rights reserved.

This Source Code Form is subject to the terms of the Mozilla Public
License, v. 2.0. If a copy of the MPL was not distributed with this
file, You can obtain one at https://mozilla.org/MPL/2.0/.
-->

# `circleci` Provider

This provider allows you to interact with the _non-preview_ paths of [CircleCI's v2 API](https://circleci.com/docs/api/v2/).

~> The [`additional_ssh_key` resource](./resources/additional_ssh_key.md) uses CircleCI's legacy endpoint because additional ssh keys are not yet supported in the v2 API.  Because of that, using this resource presents additional risk relative to other resources in this provider. Should you wish to see this addressed, please upvote [this item](https://ideas.circleci.com/api-feature-requests/p/create-additional-ssh-keys-using-v2-api) on CircleCI's ideas board.

## Example Usage

```hcl
provider "circleci" {
  api_key = "<api key>" # This may be provided by setting the `CIRCLECI_API_KEY` environment variable.
}
```

## Argument Reference

* `api_key` - your [personal API token](https://circleci.com/docs/2.0/managing-api-tokens/#creating-a-personal-api-token) for CircleCI. This may be provided by setting the `CIRCLECI_API_KEY` environment variable.
