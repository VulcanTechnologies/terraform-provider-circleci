---
page_title: "environment_variable Resource - terraform-provider-circleci"
subcategory: "Project"
description: "The environment_variable resource allows you to create environment variables in a CircleCI project."

---

<!---
Copyright (c) 2020 Vulcan, Inc.
All rights reserved.

This Source Code Form is subject to the terms of the Mozilla Public
License, v. 2.0. If a copy of the MPL was not distributed with this
file, You can obtain one at https://mozilla.org/MPL/2.0/.
-->

# `environment_variable` Resource

This resource allows you to create [environment variables](https://circleci.com/docs/2.0/env-vars/) for use in jobs of a CircleCI project.

Note that CircleCI's API always masks environment variables' values. This means that Terraform **cannot check [check for drift](https://ideas.circleci.com/api-feature-requests/p/create-additional-ssh-keys-using-v2-api) in those values**.

In order to behave properly, Terraform does store **the original, unmasked value** of the environment variable in its state. Refer to Terraform's [documentation](https://www.terraform.io/docs/state/sensitive-data.html) regarding sensitive values in state.

## Example Usage

```hcl
resource "circleci_environment_variable" "this" {
  project_slug = "gh/circleci/circleci-docs"
  name         = "FOO"
  value        = "BAR"
}
```

## Argument Reference

* `project_slug` - (**Required**) a CircleCI [project slug](../guides/project-slugs.md) for a project you have access to.
* `name` - (**Required**) the name you with to give this environment variable.
* `value` - (**Required**) the value of the environment variable.
