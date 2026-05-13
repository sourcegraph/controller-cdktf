package vmwareenginedatastore


type VmwareengineDatastoreNfsDatastoreThirdPartyFileService struct {
	// Required Mount Folder name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/vmwareengine_datastore#file_share VmwareengineDatastore#file_share}
	FileShare *string `field:"required" json:"fileShare" yaml:"fileShare"`
	// Required to identify vpc peering used for NFS access network name of NFS's vpc e.g. projects/project-id/global/networks/my-network_id.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/vmwareengine_datastore#network VmwareengineDatastore#network}
	Network *string `field:"required" json:"network" yaml:"network"`
	// Server IP addresses of the NFS file service.
	//
	// NFS v3, provide a single IP address or DNS name.
	// Multiple servers can be supported in future when NFS 4.1 protocol support
	// is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/vmwareengine_datastore#servers VmwareengineDatastore#servers}
	Servers *[]*string `field:"required" json:"servers" yaml:"servers"`
}

