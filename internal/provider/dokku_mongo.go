package provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/melbahja/goph"
)

type DokkuMongoService struct {
	DokkuGenericService
}

func NewDokkuMongoService(name string) *DokkuMongoService {
	return &DokkuMongoService{
		DokkuGenericService: DokkuGenericService{
			Name:    name,
			CmdName: "mongo",
		},
	}
}

func NewDokkuMongoServiceFromResourceData(d *schema.ResourceData) *DokkuMongoService {
	return &DokkuMongoService{
		DokkuGenericService: DokkuGenericService{
			Name:         d.Get("name").(string),
			Image:        d.Get("image").(string),
			ImageVersion: d.Get("image_version").(string),
			// Password:     d.Get("password").(string),
			// RootPassword: d.Get("root_password").(string),
			// CustomEnv:    d.Get("custom_env").(string),
			Stopped: d.Get("stopped").(bool),

			CmdName: "mongo",
		},
	}
}

func dokkuMongoRead(rm *DokkuMongoService, client *goph.Client) error {
	return dokkuServiceRead(&rm.DokkuGenericService, client)
}

func dokkuMongoCreate(rm *DokkuMongoService, client *goph.Client) error {
	return dokkuServiceCreate(&rm.DokkuGenericService, client)
}

func dokkuMongoUpdate(rm *DokkuMongoService, d *schema.ResourceData, client *goph.Client) error {
	return dokkuServiceUpdate(&rm.DokkuGenericService, d, client)
}

func dokkuMongoDestroy(rm *DokkuMongoService, client *goph.Client) error {
	return dokkuServiceDestroy(rm.CmdName, rm.Name, client)
}
