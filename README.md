# Kustomize Markdown Graph

## Disclaimer

*At the moment, only kustomize dependencies under the `resources` section are inlcuded in the resulting markdown graph.*

## Binary build

The local build (see below) will create a executable with the name `kustomize-markdown` under `/bin`.

### MacOS
Run `./build-darwin.sh`.

### Linux
Run `./build-linux.sh`.

For easy use, add the binary folder to your `PATH`:
```sh
PATH=$PATH:$(pwd)/bin
```

## Binary execution

Run the following steps:
* Navigate to the directory, where your `kustomization.yaml` file is located
* Run the executable file you created previously inside that directory

By now, the resulting markdown will be printed on your console.

## Graph generation

```sh
# Left-Right oriented graph
kustomize-markdown

# Top-Bottom oriented graph
kustomize-markdown -tb
```

## Examples

### Simple example
The examples below are related to `./sample/overlays/production/`.

```sh
cd ./sample/overlays/production/
```

**Left-Right**

The command...

```sh
kustomize-markdown
```

...will generate the following output...

<pre>
```mermaid
flowchart LR
K2388409362 --> K1632387892
K1632387892 --> K244167769
K244167769[[./moduleA<br/><br/>deploymentModuleA.yaml]]
K1632387892[[../../base<br/><br/>deployment.yaml<br/>namespace.yaml]]
K2388409362[[./production]]
```
</pre>

...which will create the following graph:

```mermaid
flowchart LR
K2388409362 --> K1632387892
K1632387892 --> K244167769
K244167769[[./moduleA<br/><br/>deploymentModuleA.yaml]]
K1632387892[[../../base<br/><br/>deployment.yaml<br/>namespace.yaml]]
K2388409362[[./production]]
```

**Top-Down**

The command...

```sh
kustomize-markdown -tb
```

...will generate the following output...

<pre>
```mermaid
flowchart TB
K2388409362 --> K1632387892
K1632387892 --> K244167769
K244167769[[./moduleA<br/><br/>deploymentModuleA.yaml]]
K1632387892[[../../base<br/><br/>deployment.yaml<br/>namespace.yaml]]
K2388409362[[./production]]
```
</pre>

...which will create the following graph:

```mermaid
flowchart TB
K2388409362 --> K1632387892
K1632387892 --> K244167769
K244167769[[./moduleA<br/><br/>deploymentModuleA.yaml]]
K1632387892[[../../base<br/><br/>deployment.yaml<br/>namespace.yaml]]
K2388409362[[./production]]
```

### More complex example

```mermaid
flowchart LR
K3692039636 --> K4134617231
K4134617231 --> K1073288559
K1073288559[[./ingress-nginx<br/><br/>./helm-release.yaml<br/>./network-policies.yaml]]
K4134617231 --> K253602071
K253602071[[./cert-manager<br/><br/>./helm-release.yaml<br/>./network-policies.yaml]]
K4134617231 --> K3117942593
K3117942593[[./reloader<br/><br/>./helm-release.yaml<br/>./network-policy.yaml]]
K4134617231 --> K1199196442
K1199196442[[./elastic-logstash<br/><br/>./config.yaml<br/>./secretprovider.yaml<br/>./statefulset.yaml]]
K4134617231 --> K1738802803
K1738802803[[./fluent-bit<br/><br/>./helm-release.yaml<br/>./secret-provider.yaml]]
K4134617231 --> K3105005572
K3105005572[[./kube-prometheus-stack<br/><br/>./helm-release.yaml<br/>./secret-provider.yaml<br/>./network-policy.yaml]]
K4134617231 --> K888054092
K888054092[[./prometheus-msteams<br/><br/>./secretprovider.yaml<br/>./helm-release.yaml]]
K4134617231 --> K548728024
K548728024[[./prometheus-pushgateway<br/><br/>./helm-release.yaml]]
K4134617231 --> K2407379872
K2407379872[[./flux-notification<br/><br/>./msteams.yaml<br/>./alert.yaml]]
K4134617231 --> K2950914330
K2950914330[[./kured<br/><br/>./helm-release.yaml]]
K4134617231 --> K453883029
K453883029[[./kubeaudit<br/><br/>kubeaudit-cron-job.yaml<br/>image-automation.yaml<br/>rbac.yaml]]
K4134617231 --> K3736359115
K3736359115 --> K3522928892
K3522928892 --> K2925748861
K2925748861[[../crd<br/><br/>bases/azure-networking.platform.markant.com_publicapis.yaml]]
K3522928892 --> K1445474874
K1445474874[[../rbac<br/><br/>service_account.yaml<br/>role.yaml<br/>role_binding.yaml<br/>leader_election_role.yaml<br/>leader_election_role_binding.yaml<br/>auth_proxy_service.yaml<br/>auth_proxy_role.yaml<br/>auth_proxy_role_binding.yaml<br/>auth_proxy_client_clusterrole.yaml]]
K3522928892 --> K3160989825
K3160989825[[../manager<br/><br/>./secret-provider.yaml<br/>./manager.yaml]]
K3522928892 --> K527266939
K527266939[[../webhook<br/><br/>manifests.yaml<br/>service.yaml]]
K3522928892 --> K1457530395
K1457530395[[../certmanager<br/><br/>certificate.yaml]]
K3522928892 --> K1312330766
K1312330766[[../prometheus<br/><br/>monitor.yaml]]
K3522928892[[./default]]
K3736359115[[./operandi<br/><br/>./image-automation.yaml<br/>./network-policy.yaml]]
K4134617231 --> K3083283366
K3083283366[[./dapr<br/><br/>./helm-release.yaml<br/>./secretprovider.yaml<br/>./pod-monitor.yaml<br/>./network-policy.yaml]]
K4134617231 --> K2644734951
K2644734951[[./azure-service-bus<br/><br/>./metrics-cronjob.yaml]]
K4134617231 --> K1834586704
K1834586704[[./aks<br/><br/>./metrics-cronjob.yaml]]
K4134617231[[../base<br/><br/>./image-reflector-controller-patch.yaml<br/>./image-update-automation.yaml<br/>./namespaces.yaml<br/>./network-policies.yaml<br/>./roles.yaml<br/>./coredns.yaml]]
K3692039636[[./dev-spoke]]
```
