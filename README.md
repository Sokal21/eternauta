# Eternauta

This project has the objective of investigating and explore the challenges of building and deploying ephemeral testing environments for testing, development etc.

## Local development

If you want to run and test the project locally you will need to install minikube https://minikube.sigs.k8s.io/docs/start/ and helm https://helm.sh/docs/intro/install/.

After this you should apply the kong-gateway manifests

```bash
kubectl kustomize kube/kong-gateway --enable-helm | kubectl apply -f -
```

It will take some time to start all the necesary things

### Local DNS resolution

We realy a lot in DNS resolution. So if you want to run it locally it is necessary to configure a few things.

Activate minikube tunnel.

```bash
minikube tunnel
```

After this get the services of the kong namespace and retrieve the external ip of the kong proxy.

```bash
kubectl get svc -n kong
NAME                           TYPE           CLUSTER-IP       EXTERNAL-IP   PORT(S)                    
kong-kong-admin                NodePort       10.105.78.135    <none>        80:32000/TCP               
kong-kong-manager              NodePort       10.110.62.154    <none>        80:32001/TCP               
kong-kong-proxy                LoadBalancer   10.110.3.11      {proxy-ip}    80:30000/TCP,443:30001/TCP 
```

Install dnsmasq. Maybe you will need to stop and disabled system-resolved:

```bash
sudo systemctl stop systemd-resolved
sudo systemctl disable systemd-resolved
sudo systemctl mask systemd-resolved
```

if you want to enable it

```bash
sudo systemctl unmask systemd-resolved
sudo systemctl enable systemd-resolved
sudo systemctl start systemd-resolved
```

this will probably disable you internet acces so follow the nexts steps.

open `sudo nano /etc/dnsmasq.conf` and add 

```
address=/.minikube/{proxy-ip}
server=8.8.8.8      # Google DNS
server=8.8.4.4      # Google DNS
```

after this `sudo systemctl restart dnsmasq`

### Instaling argo (DO NOT)

This is compleately unnecesary

```bash
kubectl create ns argocd
kubectl apply -n argocd -f https://raw.githubusercontent.com/argoproj/argo-cd/v2.5.8/manifests/install.yaml
kubectl apply -f ./kube/ingress/manifest.yaml
```

retrieve the admin password

```bash
kubectl -n argocd get secret argocd-initial-admin-secret -o jsonpath="{.data.password}" | base64 -d; echo
```

