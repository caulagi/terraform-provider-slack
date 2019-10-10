package slack

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

type Config struct {
	webhookURL string
}

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"webhook_url": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("SLACK_WEBHOOK_URL", nil),
				Description: "Webhook url for slack app as in https://api.slack.com/incoming-webhooks",
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"slack_message": resourceServer(),
		},

		ConfigureFunc: configureProvider,
	}
}

func configureProvider(d *schema.ResourceData) (interface{}, error) {
	return Config{
		webhookURL: d.Get("webhook_url").(string),
	}, nil
}
