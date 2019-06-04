package main

import (
	"encoding/json"
	"io/ioutil"

	"github.com/hashicorp/terraform/helper/schema"
)

func resourceOrigin() *schema.Resource {
	return &schema.Resource{
		Create: resourceOriginCreate,
		Read:   resourceOriginRead,
		Update: resourceOriginUpdate,
		Delete: resourceOriginDelete,

		Schema: map[string]*schema.Schema{
			"directory_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
			"host_header": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"http_hostnames": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeMap},
			},

			"https_hostnames": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeMap},
			},

			"network_configuration": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

type Origin struct {
	DirectoryName      string `json:"DirectoryName"`
	HostHeader         string `json:"HostHeader"`
	HttpHostnames      []HttpHostname
	HttpsHostnames     []HttpsHostname
	HttpsLoadBalancing string `json:"HttpsLoadBalancing"`
}

type HttpHostname struct {
	Name string `json:"HttpHostname"`
}

type HttpsHostname struct {
	Name string `json:"HttpsHostname"`
}

func resourceOriginCreate(d *schema.ResourceData, m interface{}) error {
	e, _ := m.(*edgecast)

	o := Origin{
		DirectoryName: d.Get("directory_name").(string),
		HostHeader:    d.Get("host_header").(string),
	}

	if v, lbOk := d.GetOk("load_balancing"); lbOk {
		o.HttpsLoadBalancing = v.(string)
	}

	if v, httpOk := d.GetOk("http_hostnames"); httpOk {
		hostnames := []HttpHostname{}
		for _, host := range v.([]interface{}) {
			h, ok := host.(map[string]string)
			if !ok {
				continue
			}

			httpHost := HttpHostname{Name: h["name"]}
			hostnames = append(hostnames, httpHost)
		}
		o.HttpHostnames = hostnames
	}

	if v, httpsOk := d.GetOk("https_hostnames"); httpsOk {
		hostnames := []HttpsHostname{}
		for _, host := range v.([]interface{}) {
			h, ok := host.(map[string]string)
			if !ok {
				continue
			}

			httpsHost := HttpsHostname{Name: h["name"]}
			hostnames = append(hostnames, httpsHost)
		}
		o.HttpsHostnames = hostnames
	}

	resp, err := e.Request("POST", "origins/httpsmall", o)

	if err != nil {
		return err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var originID map[string]string
	err = json.Unmarshal(body, &originID)

	d.SetId(originID["CustomerOriginId"])

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
