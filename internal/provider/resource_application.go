package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/iarlyy/terraform-provider-stackhawk/stackhawk"
)

func resourceApplication() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceApplicationCreate,
		ReadContext:   resourceApplicationRead,
		UpdateContext: resourceApplicationUpdate,
		DeleteContext: resourceApplicationDelete,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"application_id": {
				Type:     schema.TypeString,
				Computed: true,
				Required: false,
			},
			"env": {
				Type:     schema.TypeString,
				Required: true,
			},
			"env_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"data_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"risk_level": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceApplicationCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Creating application: %s", d.Get("name"))

	client := meta.(*stackhawk.Client)

	var diags diag.Diagnostics

	application, err := client.CreateApplication(stackhawk.Application{
		Name:           d.Get("name").(string),
		DataType:       d.Get("data_type").(string),
		Env:            d.Get("env").(string),
		EnvId:          d.Get("env_id").(string),
		RiskLevel:      d.Get("risk_level").(string),
		OrganizationId: client.OrgId,
	})
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(application.ApplicationId)
	d.Set("application_id", application.ApplicationId)

	return diags
}

func resourceApplicationRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[INFO] Refreshing application: %s", d.Id())

	client := meta.(*stackhawk.Client)

	var diags diag.Diagnostics

	application, err := client.GetApplication(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	d.Set("name", application.Name)
	d.Set("application_id", application.ApplicationId)
	d.Set("env", application.Env)
	d.Set("env_id", application.EnvId)
	d.Set("data_type", application.DataType)
	d.Set("risk_level", application.RiskLevel)

	return diags
}

func resourceApplicationUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Updating application: %s", d.Get("name"))

	client := meta.(*stackhawk.Client)

	var diags diag.Diagnostics

	application, err := client.UpdateApplication(stackhawk.Application{
		ApplicationId: d.Id(),
		Name:          d.Get("name").(string),
		DataType:      d.Get("data_type").(string),
		Env:           d.Get("env").(string),
		EnvId:         d.Get("env_id").(string),
		RiskLevel:     d.Get("risk_level").(string),
	})
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(application.ApplicationId)
	d.Set("application_id", application.ApplicationId)

	return diags
}

func resourceApplicationDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Deleting application: %s", d.Get("name"))

	client := meta.(*stackhawk.Client)

	var diags diag.Diagnostics

	err := client.DeleteApplication(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	return diags
}
