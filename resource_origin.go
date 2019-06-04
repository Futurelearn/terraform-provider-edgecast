package main

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceOrigin() *schema.Resource {
	return &schema.Resource{
		Create: resourceOriginCreate,
		Read:   resourceOriginRead,
		Update: resourceOriginUpdate,
		Delete: resourceOriginDelete,

		Schema: map[string]*schema.Schema{
			"hostname": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceOriginCreate(d *schema.ResourceData, m interface{}) error {
	hostname := d.Get("hostname").(string)
	d.SetId(hostname)
	return resourceOriginRead(d, m)
}

func resourceOriginRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceOriginUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceOriginRead(d, m)
}

func resourceOriginDelete(d *schema.ResourceData, m interface{}) error {
	d.SetId("")
	return nil
}
