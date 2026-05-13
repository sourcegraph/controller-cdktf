package googlevmwareenginecluster


type GoogleVmwareengineClusterDatastoreMountConfigDatastoreNetwork struct {
	// The resource name of the subnet Resource names are schemeless URIs that follow the conventions in https://cloud.google.com/apis/design/resource_names. e.g. projects/my-project/locations/us-central1/subnets/my-subnet.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_vmwareengine_cluster#subnet GoogleVmwareengineCluster#subnet}
	Subnet *string `field:"required" json:"subnet" yaml:"subnet"`
	// Optional. The number of connections of the NFS volume. Supported from vsphere 8.0u1. Possible values are 1-4. Default value is 4.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_vmwareengine_cluster#connection_count GoogleVmwareengineCluster#connection_count}
	ConnectionCount *float64 `field:"optional" json:"connectionCount" yaml:"connectionCount"`
	// Optional.
	//
	// The Maximal Transmission Unit (MTU) of the datastore.
	// MTU value can range from 1330-9000. If not set, system sets
	// default MTU size to 1500.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_vmwareengine_cluster#mtu GoogleVmwareengineCluster#mtu}
	Mtu *float64 `field:"optional" json:"mtu" yaml:"mtu"`
}

