---
page_title: "additional_ssh_key Resource - terraform-provider-circleci"
subcategory: "Project"
description: "The additional_ssh_key resource allows you to create additional ssh keys in CircleCI."

---

<!---
Copyright (c) 2020 Vulcan, Inc.
All rights reserved.

This Source Code Form is subject to the terms of the Mozilla Public
License, v. 2.0. If a copy of the MPL was not distributed with this
file, You can obtain one at https://mozilla.org/MPL/2.0/.
-->

# `additional_ssh_key` Resource

This resource allows you to upload an ssh key to a CircleCI project. Once the key exists in the project, you may [use it](https://circleci.com/docs/2.0/add-ssh-key/#adding-ssh-keys-to-a-job) in jobs.

~> The implementation of this resource uses a legacy CircleCI endpoint because additional ssh keys are not supported in CircleCI's v2 API. Because of that, using this resource presents additional risk relative to other resources in this provider. Should you wish to see this addressed, please upvote [this item](https://ideas.circleci.com/api-feature-requests/p/create-additional-ssh-keys-using-v2-api) on CircleCI's ideas board.

## Example Usage

```hcl
resource "circleci_additional_ssh_key" "this" {
  project_slug = "gh/circleci/circleci-docs"
  host_name     = "example.com"
  private_key   = file("some_key_id_rsa")
}
```

## Argument Reference

* `project_slug` - (**Required**) a CircleCI [project slug](../guides/project-slugs.md) for a project you have access to.
* `host_name` - (**Required**) the hostname for which this key will be used.
* `private_key` - (**Required**) a [PEM-encoded](https://tools.ietf.org/html/rfc7468) private key. Refer to [CircleCI's documentation](https://circleci.com/docs/2.0/add-ssh-key/#steps) for details. Note that the key **cannot have a passphrase**. Terraform's [`file`](https://www.terraform.io/docs/configuration/functions/file.html) function may be useful for reading an existing key from disk.

## Attribute Reference

* `fingerprint` - the fingerprint of the corresponding public key.
