<!---
This file adheres to https://www.terraform.io/docs/extend/best-practices/versioning.html#changelog-specification
-->

## 0.1.0 (Unreleased)

FEATURES:

* **New Resource:** `additional_ssh_key`
* **New Resource:** `environment_variable`
* **New Data Source:** `project`

NOTES:

* The new resource `additional_ssh_key` does not use CircleCI's v2 API, because that version of the API does not (yet?) support this functionality.
