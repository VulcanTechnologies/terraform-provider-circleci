package circleci

import (
	"context"
	"fmt"
	"strings"

	"github.com/antihax/optional"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stephenwithph/terraform-provider-circleci/client"
)

func resourceEnvironmentVariable() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceEnvironmentVariableCreate,
		ReadContext:   resourceEnvironmentVariableRead,
		DeleteContext: resourceEnvironmentVariableDelete,
		Schema: map[string]*schema.Schema{
			"project_slug": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				ValidateDiagFunc: assureSlugHasValidVCS,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"value": {
				Type:      schema.TypeString,
				Required:  true,
				ForceNew:  true,
				Sensitive: true,
				// although CircleCI returns masked values for environment variables, no DiffSuppressFunc is needed because resourceEnvironmentVariableRead never sets "value"
			},
		},
	}
}

func resourceEnvironmentVariableCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	provider := m.(*providerContext)
	api := provider.circleCiClient.ProjectApi
	auth := provider.authenticateContext(ctx)

	slug := d.Get("project_slug").(string)
	name := d.Get("name").(string)
	value := d.Get("value").(string)

	envVar := &client.CreateEnvVarOpts{
		EnvironmentVariablePair1: optional.NewInterface(client.EnvironmentVariablePair1{
			Name:  name,
			Value: value,
		}),
	}

	pair, _, err := api.CreateEnvVar(auth, slug, envVar)
	if err != nil {
		return diag.FromErr(err)
	}

	envVarPath := fmt.Sprintf("%s/%s", slug, pair.Name)
	d.SetId(envVarPath)

	return resourceEnvironmentVariableRead(ctx, d, m)
}

func resourceEnvironmentVariableRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	provider := m.(*providerContext)
	api := provider.circleCiClient.ProjectApi
	auth := provider.authenticateContext(ctx)

	id := d.Id()
	slug := id[:strings.LastIndex(id, "/")]
	name := id[strings.LastIndex(id, "/")+1:]

	envVar, _, err := api.GetEnvVar(auth, slug, name)
	if err != nil {
		return diag.FromErr(err)
	}

	d.Set("project_slug", slug)
	d.Set("name", envVar.Name)
	// do NOT set "value" here; CircleCI always masks the values of environment variables when returning them

	return nil
}

func resourceEnvironmentVariableDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	provider := m.(*providerContext)
	api := provider.circleCiClient.ProjectApi
	auth := provider.authenticateContext(ctx)

	id := d.Id()
	slug := id[:strings.LastIndex(id, "/")]
	name := id[strings.LastIndex(id, "/")+1:]

	_, _, err := api.DeleteEnvVar(auth, slug, name)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}
