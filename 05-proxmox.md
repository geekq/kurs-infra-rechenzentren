## Host für proxmox auf Hetzner erstellen

VM vom Typ CX22 erstellen: 4 GB RAM, 2 VCPUs, 40 GB Disk lokal
Image: Debian 12


## Static IP configuration für Hetzner

Notwendig, damit das software-defined-networking und Bridges der Netzwerk-Virtualisierung funktionieren.
Sieh auch https://docs.hetzner.com/cloud/servers/static-configuration/
und https://forum.proxmox.com/threads/installation-stuck-at-60.118540/#post-580603


## Proxmox auf Debian 12 installieren

sieh auch https://pve.proxmox.com/wiki/Install_Proxmox_VE_on_Debian_12_Bookworm

    ssh root@.........

    echo "deb http://download.proxmox.com/debian/pve bookworm pve-no-subscription" > /etc/apt/sources.list.d/pve-install-repo.list
    wget https://enterprise.proxmox.com/debian/proxmox-release-bookworm.gpg -O /etc/apt/trusted.gpg.d/proxmox-release-bookworm.gpg
    apt update && apt dist-upgrade -y

<!--
TODO check or are installing via:

wget http://download.proxmox.com/debian/proxmox-ve-release-7.x.gpg -O /etc/apt/trusted.gpg.d/proxmox-ve-release-7.x.gpg
chmod +r /etc/apt/trusted.gpg.d/proxmox-ve-release-7.x.gpg
-->

```
apt install proxmox-default-kernel
uname -a
systemctl reboot

uname -a # sollte kernel mit proxmox-Erweiterungen zeigen
apt-get install proxmox-ve
```

```
cat /etc/hosts
127.0.0.1 localhost
<server-ip> proxmox2 proxmox2

$ hostname --ip-address
sollte die public IP des Hosts zeigen - vergleiche mit dem Wert in Hetzner konsoleh
```


## Proxmox management UI aufrufen

https://<your-server-ip-address>:8006

## Firewall - Portmapper Port schützen

Port 111 aus dem Internet sperren - s. auch die BSI-Warnung (Email).

## Virtuelles Netzwerk mit VLAN erstellen

[Erklärung QinQ vs. VLAN vs. VXLAN ](https://community.fs.com/de/article/qinq-vs-vlan-vs-vxlan.html)


## Guest VMs hinzufügen

Vorbereitung - ISO Image hinzufügen - download über URL:
`http://ftp.hosteurope.de/mirror/releases.ubuntu.com/24.04.1/ubuntu-24.04.1-live-server-amd64.iso`

