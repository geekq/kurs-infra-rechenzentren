# Setup

Auf frischem Ubuntu 24.04

```
sudo apt-get install -y python3 git ansible gnupg software-properties-common
```

Terraform installieren:
(aus [Terraform quick start tutorial](https://developer.hashicorp.com/terraform/tutorials/aws-get-started/install-cli))

```
wget -O- https://apt.releases.hashicorp.com/gpg | \
gpg --dearmor | \
sudo tee /usr/share/keyrings/hashicorp-archive-keyring.gpg > /dev/null

gpg --no-default-keyring \
--keyring /usr/share/keyrings/hashicorp-archive-keyring.gpg \
--fingerprint

echo "deb [signed-by=/usr/share/keyrings/hashicorp-archive-keyring.gpg] \
https://apt.releases.hashicorp.com $(lsb_release -cs) main" | \
sudo tee /etc/apt/sources.list.d/hashicorp.list

sudo apt update
sudo apt-get install -y terraform
```

SSH Schlüsselpaar erstellen:

```
```
