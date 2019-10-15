variable "webhook_url" {
  type        = string
  description = "Webhook url for the app as described in https://api.slack.com/incoming-webhooks"
}

variable "project" {
  type        = string
  description = "Google project id"
}
