provider "slack" {
  # webhook_url = var.webhook_url
}

provider "google" {
  project = var.project
}

resource "slack_message" "message" {
  message = "Creating a preemptible instance"
}

resource "google_compute_instance" "slack-test" {
  depends_on = ["slack_message.message"]

  name     = "tf-slack"
  machine_type = "f1-micro"
  zone         = "europe-west1-b"

  boot_disk {
    initialize_params {
      image = "ubuntu-os-cloud/ubuntu-1804-lts"
    }
  }

  scheduling {
    automatic_restart = false
    preemptible       = true
  }

  network_interface {
    network = "default"
  }
}
