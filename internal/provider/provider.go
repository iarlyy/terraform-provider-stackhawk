package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/iarlyy/terraform-provider-stackhawk/stackhawk"
)

// New returns the provider instance.
func New() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"api_key": {
				Description: "Stackhawk API key" +
					"You can set this via the `HAWK_API_KEY` environment variable.",
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("HAWK_API_KEY", nil),
				Sensitive:   true,
			},
			"organization_id": {
				Description: "UUID identifier for this StackHawk Organization",
				Type:        schema.TypeString,
				Required:    true,
				Sensitive:   false,
			},
		},
		ConfigureContextFunc: providerConfigure,

		ResourcesMap: map[string]*schema.Resource{
			"stackhawk_application": resourceApplication(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"stackhawk_application": dataSourceApplication(),
		},
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	api_key := d.Get("api_key").(string)
	organization_id := d.Get("organization_id").(string)

	var diags diag.Diagnostics

	c, err := stackhawk.NewClient(nil, &organization_id, &api_key)
	if err != nil {
		return nil, diag.FromErr(err)
	}

	return c, diags
}
