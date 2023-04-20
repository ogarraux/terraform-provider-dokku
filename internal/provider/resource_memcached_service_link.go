package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/melbahja/goph"
)

func resourceMemcachedServiceLink() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceMemcachedServiceLinkCreate,
		ReadContext:   resourceMemcachedServiceLinkRead,
		DeleteContext: resourceMemcachedServiceLinkDelete,
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

const mcServiceCmd = "memcached"

//
func resourceMemcachedServiceLinkCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	err := serviceLinkCreate(d, mcServiceCmd, m.(*goph.Client))

	var diags diag.Diagnostics

	if err != nil {
		return diag.FromErr(err)
	}

	return diags
}

//
func resourceMemcachedServiceLinkRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	err := serviceLinkRead(d, mcServiceCmd, m.(*goph.Client))

	var diags diag.Diagnostics

	if err != nil {
		return diag.FromErr(err)
	}

	return diags
}

//
func resourceMemcachedServiceLinkDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	err := serviceLinkDelete(d, mcServiceCmd, m.(*goph.Client))

	var diags diag.Diagnostics

	if err != nil {
		return diag.FromErr(err)
	}

	return diags
}
