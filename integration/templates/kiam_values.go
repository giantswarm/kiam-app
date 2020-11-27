// +build k8srequired

package templates

// KiamValues values used for kiam-app in integration test
const KiamValues = `
server:
  nodeSelector:
    node-role.kubernetes.io/master: ""
`
