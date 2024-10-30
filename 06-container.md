## Kubernetes Cluster mit einem Knoten starten

[Play with Kubernetes](https://labs.play-with-k8s.com/)

Login mit Github-Account

```
kubeadm init --apiserver-advertise-address $(hostname -i) --pod-network-cidr 10.5.0.0/16

kubectl apply -f https://raw.githubusercontent.com/cloudnativelabs/kube-router/master/daemonset/kubeadm-kuberouter.yaml

kubectl apply -f https://raw.githubusercontent.com/kubernetes/website/master/content/en/examples/application/nginx-app.yaml
```

Aktive Deployments und alle ausgeführten Pods auflisten.
Bitte einzelne Befehle ausführen und die Ausgabe beobachten.

```
kubectl get deployments
kubectl get pods
kubectl get deployments.apps my-nginx -o yaml
kubectl get pods --all-namespaces
kubectl get services
```

## Weiteren Knoten hinzufügen

Lese die Ausgabe der ersten Befehle durch und suche in der Ausgabe nach
Text wie

```
Then you can join any number of worker nodes by running the following on each as root:

kubeadm join 192.168.0.13:6443 --token wa0osu.mtkpkmfrtb98ly9y \
        --discovery-token-ca-cert-hash sha256:db37b478ceb02f060cf6e77f011d2fca2c851fa363820a6655b72bc0afb42d7b
```

Das erklärt, wie man den neuen Knoten dem Cluster hinzufügen kann.

Dann "ADD NEW INSTANCE", und auf dem neuen Knoten diesen Befehl
eingeben.

Die neuen Anwendungen könen nun auf beiden Worker-Nodes gestartet
werden. `kubectl` Befehle werden bei unserem Setup aber nur von dem ersten Knoten
akzeptiert.


## Interaktives Kubernetes Management Tool - k9s

Installieren:

    curl -sS https://webinstall.dev/k9s | bash

Folge der Anleitung um `k9s` starten zu können.

[Video-Demo](https://youtu.be/k7zseUhaXeU?si=feTadNGVtWj_r2E6&t=31)


## Rufe die nginx Seite auf

```
kubectl get services
# s. Cluster-Ip für nginx LoadBalancer
# Dann rufe auf ähnlich wie: (verwende die richtige IP Adresse)
curl 10.105.40.125
```

## Service aus dem Internet erreichbar machen

```
kubectl get services
```

Achte auf die Serices, die nach außen veröffentlicht wurden - haben
meistens Portnummer um die 30512.

Mit der Url aus dem Header der Seite kann man auf die Web-Anwendung
jetzt aus dem Internet zugreifen

http://ip172-18-0-134-csh82u0i715g00ffo6r0.direct.labs.play-with-k8s.com:30584



## Andere Beispiel-Anwendungen ausprobieren

https://github.com/kubernetes/examples/blob/master/guestbook/all-in-one/guestbook-all-in-one.yaml

```
kubectl apply -f https://raw.githubusercontent.com/kubernetes/examples/master/guestbook/all-in-one/guestbook-all-in-one.yaml
kubectl get deployments
kubectl get services
```

