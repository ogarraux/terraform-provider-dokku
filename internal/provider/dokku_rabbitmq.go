package provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/melbahja/goph"
)

//
type DokkuRabbitmqService struct {
	DokkuGenericService
}

//
func NewDokkuRabbitmqService(name string) *DokkuRabbitmqService {
	return &DokkuRabbitmqService{
		DokkuGenericService: DokkuGenericService{
			Name:    name,
			CmdName: "rabbitmq",
		},
	}
}

//
func NewDokkuRabbitmqServiceFromResourceData(d *schema.ResourceData) *DokkuRabbitmqService {
	return &DokkuRabbitmqService{
		DokkuGenericService: DokkuGenericService{
			Name:         d.Get("name").(string),
			Image:        d.Get("image").(string),
			ImageVersion: d.Get("image_version").(string),
			// Password:     d.Get("password").(string),
			// RootPassword: d.Get("root_password").(string),
			// CustomEnv:    d.Get("custom_env").(string),
			Stopped: d.Get("stopped").(bool),

			CmdName: "rabbitmq",
		},
	}
}

func dokkuRmRead(rm *DokkuRabbitmqService, client *goph.Client) error {
	return dokkuServiceRead(&rm.DokkuGenericService, client)
}

func dokkuRmCreate(rm *DokkuRabbitmqService, client *goph.Client) error {
	return dokkuServiceCreate(&rm.DokkuGenericService, client)
}

func dokkuRmUpdate(rm *DokkuRabbitmqService, d *schema.ResourceData, client *goph.Client) error {
	return dokkuServiceUpdate(&rm.DokkuGenericService, d, client)
}

func dokkuRmDestroy(rm *DokkuRabbitmqService, client *goph.Client) error {
	return dokkuServiceDestroy(rm.CmdName, rm.Name, client)
}
