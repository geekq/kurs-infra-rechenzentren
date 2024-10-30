terraform {
  required_providers {
    hcloud = {
      source  = "hetznercloud/hcloud"
      version = "~> 1.48.1" // newest version as of 2024-10-19
    }
  }
}

variable "hcloud_token" {
  description = "Hetzner Cloud API token"
  type        = string
  sensitive   = true
}

provider "hcloud" {
  token = var.hcloud_token
}

// We'll put all the application related servers into a private
// network for security reasons
resource "hcloud_network" "private_network" {
  name     = "private-network"
  ip_range = "10.0.12.0/24"
}

resource "hcloud_network_subnet" "private_subnet" {
  network_id   = hcloud_network.private_network.id
  type         = "cloud"
  network_zone = "eu-central"
  ip_range     = "10.0.12.0/24"
}

// Get the list of SSH keys already created
// for this hetzner project
data "hcloud_ssh_keys" "all_keys" {
}

// possible extension for selecting keys by some tag
data "hcloud_ssh_keys" "keys_by_selector" {
  with_selector = "foo=bar"
}

resource "hcloud_server" "app1" {
  name        = "app1"
  server_type = "cx11" # Cheapest VM type
  image       = "debian-12"
  labels = {
    project = "schulung1"
  }
  public_net {
    ipv4_enabled = false
    ipv6_enabled = false
  }

  network {
    network_id = hcloud_network.private_network.id
    ip         = "10.0.12.11"
  }
  depends_on = [
    hcloud_network_subnet.private_subnet
  ]
  ssh_keys = data.hcloud_ssh_keys.all_keys.ssh_keys.*.name
}

resource "hcloud_server_network" "app1network" {
  server_id = hcloud_server.app1.id
  // network_id = hcloud_network.private_network.id
  subnet_id = hcloud_network_subnet.private_subnet.id
  ip        = "10.0.12.11"
}

resource "hcloud_server" "db1" {
  name        = "db1"
  server_type = "cx11" # Cheapest VM type
  image       = "ubuntu-24.04"
  labels = {
    project = "schulung1"
  }

  network {
    network_id = hcloud_network.private_network.id
    ip         = "10.0.12.21"
  }

  public_net {
    ipv4_enabled = false
    ipv6_enabled = false
  }
  depends_on = [
    hcloud_network_subnet.private_subnet
  ]
  ssh_keys = data.hcloud_ssh_keys.all_keys.ssh_keys.*.name
}

resource "hcloud_server" "jump1" {
  name        = "jump1"
  server_type = "cx11" # Cheapest VM type
  image       = "ubuntu-24.04"
  labels = {
    project = "schulung1"
  }

  network {
    network_id = hcloud_network.private_network.id
    ip         = "10.0.12.51"
  }
  depends_on = [
    hcloud_network_subnet.private_subnet
  ]
  public_net {
    ipv4_enabled = true
    // can use a floating IP in the future ipv4 = hcloud_primary_ip.primary_ip_1.id
    ipv6_enabled = false
  }
  ssh_keys = data.hcloud_ssh_keys.all_keys.ssh_keys.*.name
}

// Load balancer to point to the app server,
// connect LB port 80 to app1 port 80
resource "hcloud_load_balancer" "lb1" {
  name               = "applb1"
  load_balancer_type = "lb11"
  location           = "fsn1"
}


resource "hcloud_load_balancer_network" "srvnetwork" {
  load_balancer_id = hcloud_load_balancer.lb1.id
  network_id       = hcloud_network.private_network.id
  ip               = "10.0.12.5"

  # **Note**: the depends_on is important when directly attaching the
  # server to a network. Otherwise Terraform will attempt to create
  # server and sub-network in parallel. This may result in the server
  # creation failing randomly.
  depends_on = [
    hcloud_network_subnet.private_subnet
  ]
}

resource "hcloud_load_balancer_target" "load_balancer_target" {
  type             = "server"
  load_balancer_id = hcloud_load_balancer.lb1.id
  server_id        = hcloud_server.app1.id
  use_private_ip   = true
  depends_on = [
    hcloud_load_balancer_network.srvnetwork
  ]
}
/*
data "hcloud_load_balancer" "datalb1" {
  name = "applb1"
}

output "datalb1" {
  value = data.hcloud_load_balancer.datalb1
}
*/
