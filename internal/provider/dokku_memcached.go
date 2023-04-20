package provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/melbahja/goph"
)

//
type DokkuMemcachedService struct {
	DokkuGenericService
}

//
func NewDokkuMemcachedService(name string) *DokkuMemcachedService {
	return &DokkuMemcachedService{
		DokkuGenericService: DokkuGenericService{
			Name:    name,
			CmdName: "memcached",
		},
	}
}

//
func NewDokkuMemcachedServiceFromResourceData(d *schema.ResourceData) *DokkuMemcachedService {
	return &DokkuMemcachedService{
		DokkuGenericService: DokkuGenericService{
			Name:         d.Get("name").(string),
			Image:        d.Get("image").(string),
			ImageVersion: d.Get("image_version").(string),
			// Password:     d.Get("password").(string),
			// RootPassword: d.Get("root_password").(string),
			// CustomEnv:    d.Get("custom_env").(string),
			Stopped: d.Get("stopped").(bool),

			CmdName: "memcached",
		},
	}
}

func dokkuMcRead(mc *DokkuMemcachedService, client *goph.Client) error {
	return dokkuServiceRead(&mc.DokkuGenericService, client)
}

func dokkuMcCreate(mc *DokkuMemcachedService, client *goph.Client) error {
	return dokkuServiceCreate(&mc.DokkuGenericService, client)
}

func dokkuMcUpdate(mc *DokkuMemcachedService, d *schema.ResourceData, client *goph.Client) error {
	return dokkuServiceUpdate(&mc.DokkuGenericService, d, client)
}

func dokkuMcDestroy(mc *DokkuMemcachedService, client *goph.Client) error {
	return dokkuServiceDestroy(mc.CmdName, mc.Name, client)
}
