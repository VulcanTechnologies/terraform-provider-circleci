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
