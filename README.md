# Kustomize Markdown Graph

## Disclaimer

*At the moment, only kustomize dependencies under the `resources` section are inlcuded in the resulting markdown graph.*

## Binary build

For local build, run `./build.sh` on your machine. This will create a linux executable with the name `kustomize-markdown`.

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
subgraph ./production
direction LR
K4108157276{{kustomization.yaml}}
subgraph ../../base
direction LR
K2125297382{{kustomization.yaml}}
K2125297382 --> K2125297382R0(deployment.yaml)
K2125297382 --> K2125297382R1(namespace.yaml)
end
K4108157276 --> |resources| ../../base
end
```
</pre>

...which will create the following graph:

```mermaid
flowchart LR
subgraph ./production
direction LR
K4108157276{{kustomization.yaml}}
subgraph ../../base
direction LR
K2125297382{{kustomization.yaml}}
K2125297382 --> K2125297382R0(deployment.yaml)
K2125297382 --> K2125297382R1(namespace.yaml)
end
K4108157276 --> |resources| ../../base
end
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
subgraph ./production
direction TB
K4108157276{{kustomization.yaml}}
subgraph ../../base
direction TB
K2125297382{{kustomization.yaml}}
K2125297382 --> K2125297382R0(deployment.yaml)
K2125297382 --> K2125297382R1(namespace.yaml)
end
K4108157276 --> |resources| ../../base
end
```
</pre>

...which will create the following graph:

```mermaid
flowchart TB
subgraph ./production
direction TB
K4108157276{{kustomization.yaml}}
subgraph ../../base
direction TB
K2125297382{{kustomization.yaml}}
K2125297382 --> K2125297382R0(deployment.yaml)
K2125297382 --> K2125297382R1(namespace.yaml)
end
K4108157276 --> |resources| ../../base
end
```

### More complex example

```mermaid
flowchart LR
subgraph ./dev
direction LR
K3967967097{{kustomization.yaml}}
subgraph ../base
direction LR
K159962798{{kustomization.yaml}}
K159962798 --> K159962798R0(./image-reflector-controller-patch.yaml)
subgraph ./ingress-nginx
direction LR
K4062763026{{kustomization.yaml}}
K4062763026 --> K4062763026R0(./helm-release.yaml)
end
K159962798 --> |resources| ./ingress-nginx
subgraph ./cert-manager
direction LR
K4253232428{{kustomization.yaml}}
K4253232428 --> K4253232428R0(./helm-release.yaml)
end
K159962798 --> |resources| ./cert-manager
subgraph ./reloader
direction LR
K3627363822{{kustomization.yaml}}
K3627363822 --> K3627363822R0(./helm-release.yaml)
end
K159962798 --> |resources| ./reloader
subgraph ./elastic-logstash
direction LR
K2666373321{{kustomization.yaml}}
K2666373321 --> K2666373321R0(./namespace.yaml)
K2666373321 --> K2666373321R1(./config.yaml)
K2666373321 --> K2666373321R2(./secretprovider.yaml)
K2666373321 --> K2666373321R3(./statefulset.yaml)
end
K159962798 --> |resources| ./elastic-logstash
K159962798 --> K159962798R5(./monitoring-namespace.yaml)
subgraph ./fluent-bit
direction LR
K685186192{{kustomization.yaml}}
K685186192 --> K685186192R0(./helm-release.yaml)
K685186192 --> K685186192R1(./secret-provider.yaml)
end
K159962798 --> |resources| ./fluent-bit
subgraph ./kube-prometheus-stack
direction LR
K2863649673{{kustomization.yaml}}
K2863649673 --> K2863649673R0(./helm-release.yaml)
K2863649673 --> K2863649673R1(./secret-provider.yaml)
end
K159962798 --> |resources| ./kube-prometheus-stack
subgraph ./flux-notification
direction LR
K3220439621{{kustomization.yaml}}
K3220439621 --> K3220439621R0(./msteams.yaml)
K3220439621 --> K3220439621R1(./alert.yaml)
end
K159962798 --> |resources| ./flux-notification
subgraph ./kured
direction LR
K1187378851{{kustomization.yaml}}
K1187378851 --> K1187378851R0(./helm-release.yaml)
end
K159962798 --> |resources| ./kured
end
K3967967097 --> |resources| ../base
end
```

