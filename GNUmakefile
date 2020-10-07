#
# Copyright (c) 2020 Vulcan, Inc.
# All rights reserved.
#
# This Source Code Form is subject to the terms of the Mozilla Public
# License, v. 2.0. If a copy of the MPL was not distributed with this
# file, You can obtain one at https://mozilla.org/MPL/2.0/.
#

this_file := $(lastword $(MAKEFILE_LIST))

MAKEFLAGS += --warn-undefined-variables

minimum_make_version := 4.1
current_make_version := $(MAKE_VERSION)

ifneq ($(minimum_make_version), $(firstword $(sort $(current_make_version) $(minimum_make_version))))
$(error You need GNU make version $(minimum_make_version) or greater. You have $(current_make_version))
endif

.POSIX:
SHELL := /bin/sh

.DEFAULT_GOAL := help

column1_helptext_width := 20
column2_helptext_width := 55
column3_helptext_width := 25

circleci_spec_url := https://circleci.com/api/v2/openapi.json
circleci_spec_path := $(CURDIR)/spec/openapi.json
circleci_non_preview_spec_path := $(CURDIR)/spec/openapi-non-preview.json

spectral_docker_image := docker.io/stoplight/spectral:5.6.0
container_runtime := docker


.PHONY: help
help: ## show this help
	@ printf "\033[36m%-$(column1_helptext_width)s\033[0m%-$(column2_helptext_width)s\033[93m%-$(column3_helptext_width)s\033[92m%s\033[0m\n" "target" "description" "arguments" "defaults" >&2
	@ printf "%s\n" "------------------------------------------------------------------------------------------------------------------------------------" >&2
	@ grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk -F ":.*?## |[\\\|] " '{printf "\033[36m%-$(column1_helptext_width)s\033[0m%-$(column2_helptext_width)s\033[93m%-$(column3_helptext_width)s\033[92m%s\033[0m\n", $$1, $$2, $$3, $$4}' >&2

.PHONY: generate_spec
generate_spec: ## download api v2 spec and remove preview paths from it
	@ $(MAKE) --file '$(this_file)' --no-print-directory check_command 'command=jq'
	@ $(MAKE) --file '$(this_file)' --no-print-directory check_command 'command=$(container_runtime)'
	curl --fail --output '$(circleci_spec_path)' --silent '$(circleci_spec_url)'
	jq '. as $$in | $$in.paths |= (map_values(with_entries(select(.value.tags | inside($$in | [.tags[].name] - [.tags[] | select(has("x-displayName")) | select(."x-displayName" | contains("Preview")) | .name])))) | with_entries(select(.value != {})))' < '$(circleci_spec_path)' > '$(circleci_non_preview_spec_path)'
	$(container_runtime) run \
		--mount='type=bind,src=$(CURDIR),target=$(CURDIR),ro' \
		--rm \
		--tty \
		--workdir '$(CURDIR)' \
		'$(spectral_docker_image)' lint --fail-severity error --verbose '$(circleci_non_preview_spec_path)'

.PHONY: check_command
check_command: command ?=
check_command:
	@ command -v '$(command)' > /dev/null || { printf "\n\nrequires %s but that command was not found\n\n" '$(command)' >&2 ; exit 1; }
