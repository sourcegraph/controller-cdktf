package hypercomputeclustercluster


type HypercomputeclusterClusterStorageResourcesConfigNewFilestore struct {
	// file_shares block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/hypercomputecluster_cluster#file_shares HypercomputeclusterCluster#file_shares}
	FileShares interface{} `field:"required" json:"fileShares" yaml:"fileShares"`
	// Name of the Filestore instance to create, in the format 'projects/{project}/locations/{location}/instances/{instance}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/hypercomputecluster_cluster#filestore HypercomputeclusterCluster#filestore}
	Filestore *string `field:"required" json:"filestore" yaml:"filestore"`
	// Service tier to use for the instance. Possible values: ZONAL REGIONAL Possible values: ["TIER_UNSPECIFIED", "ZONAL", "REGIONAL"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/hypercomputecluster_cluster#tier HypercomputeclusterCluster#tier}
	Tier *string `field:"required" json:"tier" yaml:"tier"`
	// Description of the instance. Maximum of 2048 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/hypercomputecluster_cluster#description HypercomputeclusterCluster#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Access protocol to use for all file shares in the instance.
	//
	// Defaults to NFS
	// V3 if not set.
	// Possible values:
	// NFSV3
	// NFSV41 Possible values: ["PROTOCOL_UNSPECIFIED", "NFSV3", "NFSV41"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/hypercomputecluster_cluster#protocol HypercomputeclusterCluster#protocol}
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
}

