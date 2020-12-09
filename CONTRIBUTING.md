<!---
Copyright (c) 2020 Vulcan, Inc.
All rights reserved.

This Source Code Form is subject to the terms of the Mozilla Public
License, v. 2.0. If a copy of the MPL was not distributed with this
file, You can obtain one at https://mozilla.org/MPL/2.0/.
-->


<!-- The two trailing spaces at the end of the numbered list items are there to force formatting. Please leave them alone. -->
# How to Contribute

## Before Opening a PR
1. Run the tests per [the README](./README.md#test-the-provider).

1. [Changelog](./CHANGELOG.md)  
  Regardless of what you contribute (new functionality, bug fixes, documentation), please [update the changelog](https://keepachangelog.com/en/1.0.0/#how).

1. Documentation  
  Be sure to show (or update) inputs, outputs, and example usage in the module's [docs](./docs).

## Releasing
1. Cut a release branch.

1. Run the tests per [the README](./README.md#test-the-provider).

1. [Update the Changelog](https://keepachangelog.com/en/1.0.0/#effort)  
  Move unreleased items to a new, properly versioned section and update the `Unreleased` GitHub link's ref to compare `HEAD` against the latest _released_ tag.

1. Merge release code to master  
  Once your feature branch has the above release updates and PR approval, merge the updates to master.

1. Create and push a tag  
  [Create a lightweight git tag](https://git-scm.com/book/en/v2/Git-Basics-Tagging) using the appropriate version number. Push it. There is a GitHub action which will take care of building and signing the binaries.
