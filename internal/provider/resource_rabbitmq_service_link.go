package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/melbahja/goph"
)

func resourceRabbitmqServiceLink() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceRabbitmqServiceLinkCreate,
		ReadContext:   resourceRabbitmqServiceLinkRead,
		DeleteContext: resourceRabbitmqServiceLinkDelete,
		Schema: map[string]*schema.Schema{
			"service": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"app": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"alias": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"query_string": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
		},
	}
}

const rmServiceCmd = "rabbitmq"

//
func resourceRabbitmqServiceLinkCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	err := serviceLinkCreate(d, rmServiceCmd, m.(*goph.Client))

	var diags diag.Diagnostics

	if err != nil {
		return diag.FromErr(err)
	}

	return diags
}

//
func resourceRabbitmqServiceLinkRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	err := serviceLinkRead(d, rmServiceCmd, m.(*goph.Client))

	var diags diag.Diagnostics

	if err != nil {
		return diag.FromErr(err)
	}

	return diags
}

//
func resourceRabbitmqServiceLinkDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	err := serviceLinkDelete(d, rmServiceCmd, m.(*goph.Client))

	var diags diag.Diagnostics

	if err != nil {
		return diag.FromErr(err)
	}

	return diags
}
