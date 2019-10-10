package slack

import (
	"log"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/nlopes/slack"
)

func resourceServer() *schema.Resource {
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
	config := m.(Config)
	message := d.Get("message").(string)
	d.SetId(message)
	err := sendMessage(message, config.webhookURL)
	if err != nil {
		return resourceServerRead(d, m)
	}
	return err
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

func sendMessage(message string, webhookURL string) error {
	log.Printf("[DEBUG] Sending message: %s to webhook: %s", message, webhookURL)
	attachment := slack.Attachment{
		Color: "good",
		Text:  message,
	}
	msg := slack.WebhookMessage{
		Attachments: []slack.Attachment{attachment},
	}

	err := slack.PostWebhook(webhookURL, &msg)
	if err != nil {
		log.Printf("[ERROR] Failed to send to slack: %v", err)
		return err
	}
	return nil
}
