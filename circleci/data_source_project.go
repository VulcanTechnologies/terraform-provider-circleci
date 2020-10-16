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
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceProject() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceProjectRead,
		Schema: map[string]*schema.Schema{
			"project_slug": {
				Type:             schema.TypeString,
				Required:         true,
				ValidateDiagFunc: assureSlugHasValidVCS,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"organization_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceProjectRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	provider := m.(*providerContext)
	api := provider.circleCiClient.ProjectApi
	auth := provider.authenticateContext(ctx)

	slug := d.Get("project_slug").(string)

	project, _, err := api.GetProjectBySlug(auth, slug)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(project.Slug)
	d.Set("project_slug", project.Slug)
	d.Set("name", project.Name)
	d.Set("organization_name", project.OrganizationName)

	return nil
}
