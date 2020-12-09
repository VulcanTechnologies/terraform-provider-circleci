<!---
This file adheres to https://www.terraform.io/docs/extend/best-practices/versioning.html#changelog-specification
-->

## [Unreleased](https://github.com/VulcanTechnologies/terraform-provider-circleci/compare/v1.0.0...HEAD)

## [v1.0.0](https://github.com/VulcanTechnologies/terraform-provider-circleci/releases/tag/v1.0.0)

FEATURES:

* **New Resource:** `additional_ssh_key`
* **New Resource:** `environment_variable`
* **New Data Source:** `project`

NOTES:

* The new resource `additional_ssh_key` does not use CircleCI's v2 API, because that version of the API does not (yet?) support this functionality.
