package provider

import (
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/iarlyy/terraform-provider-stackhawk/stackhawk"
)

func dataSourceApplication() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceApplicationRead,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"application_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"env": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"env_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"data_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"risk_level": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceApplicationRead(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[INFO] Data source for Application ID %s", d.Get("application_id").(string))

	client := meta.(*stackhawk.Client)

	application, err := client.GetApplication(d.Get("application_id").(string))
	if err != nil {
		return nil
	}

	d.SetId(application.ApplicationId)
	d.Set("name", application.Name)
	d.Set("application_id", application.ApplicationId)
	d.Set("env", application.Env)
	d.Set("env_id", application.EnvId)
	d.Set("data_type", application.DataType)
	d.Set("risk_level", application.RiskLevel)

	return nil
}
