[![CircleCI](https://circleci.com/gh/giantswarm/kiam-app.svg?style=svg)](https://circleci.com/gh/giantswarm/kiam-app)

# kiam-app
Helm chart for kiam service running in guest clusters

## Installing the Chart

To install the chart locally:

```bash
$ git clone https://github.com/giantswarm/kiam-app.git
$ cd kiam-app
$ helm install helm/kiam-app
```

Provide a custom `values.yaml`:

```bash
$ helm install kiam-app -f values.yaml
```

Deployment to Tenant Clusters is handled by [app-operator](https://github.com/giantswarm/app-operator).

## Configuration

Configuration options are documented in [Configuration.md](helm/kiam-app/Configuration.md) document.
