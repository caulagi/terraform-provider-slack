package slack

import (
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/nlopes/slack"
)

func resourceServer() *schema.Resource {
	log.Printf("[INFO] Creating Sentry organization %s", "params.Name")
	return &schema.Resource{
		Create: resourceServerCreate,
		Read:   resourceServerRead,
		Update: resourceServerUpdate,
		Delete: resourceServerDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"message": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The human readable name for the organization",
			},
		},
	}
}

func resourceServerCreate(d *schema.ResourceData, m interface{}) error {
	message := d.Get("message").(string)
	d.SetId(message)
	sendMessage(message)
	return resourceServerRead(d, m)
}

func resourceServerRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceServerUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceServerRead(d, m)
}

func resourceServerDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}

func sendMessage(message string) error {
	attachment := slack.Attachment{
		Color: "good",
		Text:  message,
	}
	msg := slack.WebhookMessage{
		Attachments: []slack.Attachment{attachment},
	}

	err := slack.PostWebhook("FIXME", &msg)
	if err != nil {
		return err
	}
	return nil
}
