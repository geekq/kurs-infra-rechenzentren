## Aufgabe 2

![Aufgabe2](aufgabeN2.mermaid.png)

SSH auf den Jump-Host:

    ssh root@<public-ip>

Tunnel auf die Server im privaten Netzwerk

    ssh -A -J root@<public-ip> root@<private-ip>

Auf beiden Web/App-Servern nginx installieren und eine Homepage mit verschiedenen Inhalten vorbereiten:

```
apt-get install nginx
nano /var/www/html/...
```

Beide Web/App-Server in das Load-Balancing einbinden.
Die IP von dem Load-Balancer aufrufen und die Seite im Browser auffrischen. Kommen da mal die, mal die anderen Inhalte?

---

## Infrastruktur per Terraform erstellen

Terraform Datei herunterladen

https://raw.githubusercontent.com/geekq/kurs-infra-rechenzentren/refs/heads/main/tf-example-web-app-hetzner-vms/main.tf

Und Projektnamen anpassen von schulung1 auf schulung-teamN

```
terraform plan
terraform apply
```

---

## Http Proxy

### Auf der Machine mit public IP

    vim /etc/squid/squid.conf

Zeile hinzufügen / einkommentieren und die derzeitige Liste für `localnet`
überprüfen, `10...` Netz sollte dabei sein:

    http_access allow localnet

Den Proxy-Server neu starten

    systemctl restart squid


### Auf den Servern im privaten Netz

```
export http_proxy="http://10.0.12.51:3128"
export https_proxy="http://10.0.12.51:3128"
apt-get update
```

Man kann während des Updates auf dem proxy Server die Aufrufe verfolgen

    tail -f /var/log/squid/access.log
