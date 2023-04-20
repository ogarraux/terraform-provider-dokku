package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/melbahja/goph"
)

func resourceMemcachedService() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceMcCreate,
		ReadContext:   resourceMcRead,
		UpdateContext: resourceMcUpdate,
		DeleteContext: resourceMcDelete,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"image": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			// TODO: locked support
			"image_version": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			// We can't support these yet as there's no way to
			// retrieve them from dokku
			// "password": {
			// 	Type:     schema.TypeString,
			// 	Optional: true,
			// },
			// "root_password": {
			// 	Type:     schema.TypeString,
			// 	Optional: true,
			// },
			// "custom_env": {
			// 	Type:     schema.TypeString,
			// 	Optional: true,
			// },
			"stopped": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			// TODO backup related stuff
			// "backup_auth_access_key": {
			// 	Type:     schema.TypeString,
			// 	Optional: true,
			// },
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceMcCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sshClient := m.(*goph.Client)

	var diags diag.Diagnostics

	mc := NewDokkuMemcachedServiceFromResourceData(d)
	err := dokkuMcCreate(mc, sshClient)

	if err != nil {
		return diag.FromErr(err)
	}

	mc.setOnResourceData(d)

	// TODO stop if necessary

	return diags
}

//
func resourceMcRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sshClient := m.(*goph.Client)

	var diags diag.Diagnostics

	var serviceName string
	if d.Id() == "" {
		serviceName = d.Get("name").(string)
	} else {
		serviceName = d.Id()
	}

	mc := NewDokkuMemcachedService(serviceName)
	err := dokkuMcRead(mc, sshClient)

	if err != nil {
		return diag.FromErr(err)
	}

	mc.setOnResourceData(d)

	return diags
}

//
func resourceMcUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sshClient := m.(*goph.Client)

	var diags diag.Diagnostics

	mc := NewDokkuMemcachedServiceFromResourceData(d)
	err := dokkuMcUpdate(mc, d, sshClient)

	if err != nil {
		return diag.FromErr(err)
	}

	mc.setOnResourceData(d)

	return diags
}

//
func resourceMcDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sshClient := m.(*goph.Client)

	var diags diag.Diagnostics

	err := dokkuMcDestroy(NewDokkuMemcachedService(d.Id()), sshClient)

	if err != nil {
		return diag.FromErr(err)
	}

	return diags
}
