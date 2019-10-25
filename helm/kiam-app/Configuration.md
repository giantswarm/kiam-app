# kiam

This chart installs kiam as managed applications. kiam runs as an agent on each node in your Kubernetes cluster and allows cluster users to associate IAM roles to Pods.


## Configuration

The following table lists the configurable parameters of the kiam chart, its dependencies and default values.

Parameter | Description | Default
--- | --- | ---
`clusterID` | Cluster identifier. Applies only to Giant Swarm managed clusters | 'testid'
`provider` | Provider identifier (`aws`/`azure`/`kvm`). `kiam` applies only to `aws` provider | 'aws'
`iam.managed` | If `true` - app will use precreated IAM role in Giant Swarm cluster | `true`
`iam.assumeRoleARN` | ARN of the role to assume if `iam.managed` is false | ""
