# Kustomize Markdown Graph

## Disclaimer

*At the moment, only kustomize dependencies under the `resources:` section are inlcuded in the resulting markdown graph.*

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

# Top-Down oriented graph
kustomize-markdown -tb
```

## Examples

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
