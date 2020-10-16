package circleci

import (
	"context"
	"net/http"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/stephenwithph/terraform-provider-circleci/client"
)

type providerContext struct {
	apiKey              string // this is an allowance for the legacy ssh key endpoint and can go away once that goes away
	authenticateContext func(context.Context) context.Context
	circleCiClient      *client.APIClient
}

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"api_key": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("CIRCLECI_API_KEY", nil),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"circleci_environment_variable": resourceEnvironmentVariable(),
			"circleci_additional_ssh_key":   resourceAdditionalSSHKey(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"circleci_project": dataSourceProject(),
		},
		ConfigureContextFunc: createProviderContext,
	}
}

func createProviderContext(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	config := client.NewConfiguration()

	config.HTTPClient = &http.Client{
		Timeout: 10 * time.Second,
	}

	provider := &providerContext{
		apiKey: d.Get("api_key").(string),
		authenticateContext: func(ctx context.Context) context.Context {
			return context.WithValue(ctx, client.ContextAPIKey, client.APIKey{Key: d.Get("api_key").(string)})
		},
		circleCiClient: client.NewAPIClient(client.NewConfiguration()),
	}

	return provider, nil
}
