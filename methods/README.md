# methods

This provides a simple wrapper around the most basic (and common) operations.  This example is geared towards a k8s environment, and takes yaml input as the config which specifies the MVIP, SVIP, Login credentials and the account name to use for volume operations.  It provides the ability to create, delete, list, update (size or QoS) as well as perform iscsi connect operations (uses the kubernetes-csi iscsi-lib).

## Example yaml config
`
endpoint: https://administrator:SolidF1r3!@10.117.28.84/json-rpc/10.0
svip: 10.10.64.3:3260
tenantname: px-admin
defaultvolumesize: 1
initiatoriface: default
`

The remaining needed items (like CHAP credentials) should be able to be collected by the init routine itself so long as the endpoint and tenant info is correct.
