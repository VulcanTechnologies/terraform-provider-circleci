---
page_title: "CircleCI Project Slugs"

---

<!---
Copyright (c) 2020 Vulcan, Inc.
All rights reserved.

This Source Code Form is subject to the terms of the Mozilla Public
License, v. 2.0. If a copy of the MPL was not distributed with this
file, You can obtain one at https://mozilla.org/MPL/2.0/.
-->

# Constructing a Project Slug
CircleCI's [official documents](https://circleci.com/docs/2.0/api-developers-guide/#getting-started-with-the-api) explain at length how a project slug is constructed. In summary, a project slug has three elements in the form `{x}/{y}/{z}`. The first element, `{x}`, should be CircleCI's _short form_ code for your VCS provider: `gh` for GitHub or `bb` for BitBucket. The second element, `{y}`, is the _owner_ of the repository in question. This could be the name of your personal account, or this could be the name of the organization which owns the repository. The final element, `{z}`, is the name of the repository itself.

## Regarding Case-Sensitivity
CircleCI's documents are ambiguous regarding case-sensitivity for project slugs. If you would like to see this resolved, please upvote [this item](https://ideas.circleci.com/documentation/p/clarify-in-docs-whether-project-slugs-are-case-sensitive) on CircleCI's ideas board. [Anecdotally](https://discuss.circleci.com/t/github-case-insensitive-org-name/23899), it seems project slug is case-sensitive, **but** the second element of the slug (`{y}` above) may have different casing than what is used in GitHub or BitBucket. If in doubt, use the repository owner's name (`{y}` above) as it appears in the URL bar of your browser when viewing the project in CircleCI's web UI.

## When The Project Is Not Found (404)
An invalid or malformed project slug will likely result in a **404 Not Found** response which this provider should surface as an error. A _valid_ project slug will also produce that error _if you have not properly authenticated or if you do not have permissions to view the CircleCI project_.
