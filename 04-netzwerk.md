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
    