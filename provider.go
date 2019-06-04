package main

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"account_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"api_key": {
				Type:     schema.TypeString,
				Required: true,
			},
			"base_url": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "https://api.edgecast.com/v2/mcc/customers",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"edgecast_origin": resourceOrigin(),
			"edgecast_cname":  resourceCname(),
		},

		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	url := d.Get("base_url").(string) + "/" + d.Get("account_id").(string)

	config := edgecast{
		apiKey:  d.Get("api_key").(string),
		baseURL: url,
	}

	// TODO: Probably should do some validation of the account id and API key
	// and return errors here
	return &config, nil
}
