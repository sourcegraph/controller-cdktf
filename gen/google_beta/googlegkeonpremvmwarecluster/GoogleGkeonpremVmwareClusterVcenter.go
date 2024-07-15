package googlegkeonpremvmwarecluster


type GoogleGkeonpremVmwareClusterVcenter struct {
	// Contains the vCenter CA certificate public key for SSL verification.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gkeonprem_vmware_cluster#ca_cert_data GoogleGkeonpremVmwareCluster#ca_cert_data}
	CaCertData *string `field:"optional" json:"caCertData" yaml:"caCertData"`
	// The name of the vCenter cluster for the user cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gkeonprem_vmware_cluster#cluster GoogleGkeonpremVmwareCluster#cluster}
	Cluster *string `field:"optional" json:"cluster" yaml:"cluster"`
	// The name of the vCenter datacenter for the user cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gkeonprem_vmware_cluster#datacenter GoogleGkeonpremVmwareCluster#datacenter}
	Datacenter *string `field:"optional" json:"datacenter" yaml:"datacenter"`
	// The name of the vCenter datastore for the user cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gkeonprem_vmware_cluster#datastore GoogleGkeonpremVmwareCluster#datastore}
	Datastore *string `field:"optional" json:"datastore" yaml:"datastore"`
	// The name of the vCenter folder for the user cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gkeonprem_vmware_cluster#folder GoogleGkeonpremVmwareCluster#folder}
	Folder *string `field:"optional" json:"folder" yaml:"folder"`
	// The name of the vCenter resource pool for the user cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gkeonprem_vmware_cluster#resource_pool GoogleGkeonpremVmwareCluster#resource_pool}
	ResourcePool *string `field:"optional" json:"resourcePool" yaml:"resourcePool"`
	// The name of the vCenter storage policy for the user cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gkeonprem_vmware_cluster#storage_policy_name GoogleGkeonpremVmwareCluster#storage_policy_name}
	StoragePolicyName *string `field:"optional" json:"storagePolicyName" yaml:"storagePolicyName"`
}

