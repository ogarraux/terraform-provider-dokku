package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/melbahja/goph"
)

func resourceRabbitmqService() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceRmCreate,
		ReadContext:   resourceRmRead,
		UpdateContext: resourceRmUpdate,
		DeleteContext: resourceRmDelete,
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

func resourceRmCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sshClient := m.(*goph.Client)

	var diags diag.Diagnostics

	rm := NewDokkuRabbitmqServiceFromResourceData(d)
	err := dokkuRmCreate(rm, sshClient)

	if err != nil {
		return diag.FromErr(err)
	}

	rm.setOnResourceData(d)

	// TODO stop if necessary

	return diags
}

//
func resourceRmRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sshClient := m.(*goph.Client)

	var diags diag.Diagnostics

	var serviceName string
	if d.Id() == "" {
		serviceName = d.Get("name").(string)
	} else {
		serviceName = d.Id()
	}

	rm := NewDokkuRabbitmqService(serviceName)
	err := dokkuRmRead(rm, sshClient)

	if err != nil {
		return diag.FromErr(err)
	}

	rm.setOnResourceData(d)

	return diags
}

//
func resourceRmUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sshClient := m.(*goph.Client)

	var diags diag.Diagnostics

	rm := NewDokkuRabbitmqServiceFromResourceData(d)
	err := dokkuRmUpdate(rm, d, sshClient)

	if err != nil {
		return diag.FromErr(err)
	}

	rm.setOnResourceData(d)

	return diags
}

//
func resourceRmDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sshClient := m.(*goph.Client)

	var diags diag.Diagnostics

	err := dokkuRmDestroy(NewDokkuRabbitmqService(d.Id()), sshClient)

	if err != nil {
		return diag.FromErr(err)
	}

	return diags
}
