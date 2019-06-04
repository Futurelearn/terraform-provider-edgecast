package main

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceCname() *schema.Resource {
	return &schema.Resource{
		Create: resourceCnameCreate,
		Read:   resourceCnameRead,
		Update: resourceCnameUpdate,
		Delete: resourceCnameDelete,

		Schema: map[string]*schema.Schema{
			"hostname": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceCnameCreate(d *schema.ResourceData, m interface{}) error {
	hostname := d.Get("hostname").(string)
	d.SetId(hostname)
	return resourceCnameRead(d, m)
}

func resourceCnameRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceCnameUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceCnameRead(d, m)
}

func resourceCnameDelete(d *schema.ResourceData, m interface{}) error {
	d.SetId("")
	return nil
}
