# terraform-provider-slack

A terraform provider which allows us to send messages to slack
from terraform files (for instance, before creating/updating
a Kubernetes cluster)

**Note:** You will need to create an [incoming webhook][webhook] for Slack before proceeding.

## Example usage

```shell
# build the plugin
$ go build -o example/.terraform/plugins/darwin_amd64/terraform-provider-slack

# Run the example which will send a slack message
$ cd example
$ terraform init -plugin-dir ..
$ terraform plan -var webhook_url=<fixme>
$ terraform apply -var webhook_url=<fixme>
```

## License

This project is licensed under [MIT](LICENSE).

[webhook]: https://api.slack.com/incoming-webhooks
