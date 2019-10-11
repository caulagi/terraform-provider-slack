package slack

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/nlopes/slack"
	"log"
)

func slackMessage() *schema.Resource {
	return &schema.Resource{
		Create: slackMessageCreate,
		Read:   slackMessageRead,
		Update: slackMessageUpdate,
		Delete: slackMessageDelete,

		Importer: nil,

		Schema: map[string]*schema.Schema{
			"message": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The human readable name for the organization",
			},
		},
	}
}

func slackMessageCreate(d *schema.ResourceData, m interface{}) error {
	message := d.Get("message").(string)
	err := sendMessage(message, m.(Config).webhookURL)
	if err != nil {

		return err
	}
	d.SetId(message)
	return slackMessageRead(d, m)
}

func slackMessageRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func slackMessageUpdate(d *schema.ResourceData, m interface{}) error {
	return slackMessageRead(d, m)
}

func slackMessageDelete(d *schema.ResourceData, m interface{}) error {
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
