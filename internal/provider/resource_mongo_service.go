package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/melbahja/goph"
)

func resourceMongoService() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceMongoCreate,
		ReadContext:   resourceMongoRead,
		UpdateContext: resourceMongoUpdate,
		DeleteContext: resourceMongoDelete,
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

func resourceMongoCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sshClient := m.(*goph.Client)

	var diags diag.Diagnostics

	mongo := NewDokkuMongoServiceFromResourceData(d)
	err := dokkuMongoCreate(mongo, sshClient)

	if err != nil {
		return diag.FromErr(err)
	}

	mongo.setOnResourceData(d)

	// TODO stop if necessary

	return diags
}

func resourceMongoRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sshClient := m.(*goph.Client)

	var diags diag.Diagnostics

	var serviceName string
	if d.Id() == "" {
		serviceName = d.Get("name").(string)
	} else {
		serviceName = d.Id()
	}

	mongo := NewDokkuMongoService(serviceName)
	err := dokkuMongoRead(mongo, sshClient)

	if err != nil {
		return diag.FromErr(err)
	}

	mongo.setOnResourceData(d)

	return diags
}

func resourceMongoUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sshClient := m.(*goph.Client)

	var diags diag.Diagnostics

	mongo := NewDokkuMongoServiceFromResourceData(d)
	err := dokkuMongoUpdate(mongo, d, sshClient)

	if err != nil {
		return diag.FromErr(err)
	}

	mongo.setOnResourceData(d)

	return diags
}

func resourceMongoDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sshClient := m.(*goph.Client)

	var diags diag.Diagnostics

	err := dokkuMongoDestroy(NewDokkuMongoService(d.Id()), sshClient)

	if err != nil {
		return diag.FromErr(err)
	}

	return diags
}
