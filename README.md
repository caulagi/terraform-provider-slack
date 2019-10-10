# terraform-provider-slack

A terraform provider which allows us to send messages to slack
from terraform files (for instance, before creating/updating
a Kubernetes cluster)

**Note:** You will need to create an [incoming webhook][1]

[1]: (https://api.slack.com/incoming-webhooks) 

## Example usage

```shell
# build the plugin
$ go build -o terraform-provider-slack

# Run the example which will send a slack message
$ cd example
$ terraform init -plugin-dir ..
$ terraform plan -var webhook_url=<fixme>
$ terraform apply -var webhook_url=<fixme>
```

## License

This project is licensed under [MIT](LICENSE).