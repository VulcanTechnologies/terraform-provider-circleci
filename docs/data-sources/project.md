---
page_title: "project Data Source - terraform-provider-circleci"
subcategory: "Project"
description: "The project data source allows you to retrieve information about a CircleCI project by project slug."

---

<!---
Copyright (c) 2020 Vulcan, Inc.
All rights reserved.

This Source Code Form is subject to the terms of the Mozilla Public
License, v. 2.0. If a copy of the MPL was not distributed with this
file, You can obtain one at https://mozilla.org/MPL/2.0/.
-->

# `project` Data Source

This data source allows you to retrieve basic information about a CircleCI project by the project's slug.

## Example Usage

```hcl
data "circleci_project" "this" {
  project_slug = "gh/circleci/circleci-docs"
}
```

## Argument Reference

* `project_slug` - (**Required**) a CircleCI [project slug](../guides/project-slugs.md) for a project you have access to.

## Attribute Reference

* `name` - The name of the CircleCI project. This corresponds with the repository name.
* `organization_name` - The name of the repository owner. This could be an organization or individual account with the VCS provider specified in `project_slug`.
