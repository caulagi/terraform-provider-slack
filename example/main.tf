provider "slack" {
  webhook_url = var.webhook_url
}

resource "slack_message" "message" {
  message = "Terraform slack provider"
}
