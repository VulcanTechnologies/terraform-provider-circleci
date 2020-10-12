/*
 * Copyright (c) 2020 Vulcan, Inc.
 * All rights reserved.
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 *
 */

package circleci

import (
	"strings"

	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
)

func assureSlugHasValidVCS(slug interface{}, _ cty.Path) diag.Diagnostics {
	stringifiedSlug := slug.(string)
	if strings.HasPrefix(stringifiedSlug, "gh/") || strings.HasPrefix(stringifiedSlug, "bb/") {
		return nil
	}

	return diag.Errorf("A project_slug must begin with 'gh/' or 'bb/' depending on your vcs provider.")
}
