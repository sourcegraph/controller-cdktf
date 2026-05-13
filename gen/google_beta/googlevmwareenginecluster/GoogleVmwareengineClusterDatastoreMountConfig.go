package googlevmwareenginecluster


type GoogleVmwareengineClusterDatastoreMountConfig struct {
	// The resource name of the datastore to unmount.
	//
	// The datastore requested to be mounted should be in same region/zone as the
	// cluster.
	// Resource names are schemeless URIs that follow the conventions in
	// https://cloud.google.com/apis/design/resource_names.
	// For example:
	// 'projects/my-project/locations/us-central1/datastores/my-datastore'
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_vmwareengine_cluster#datastore GoogleVmwareengineCluster#datastore}
	Datastore *string `field:"required" json:"datastore" yaml:"datastore"`
	// datastore_network block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_vmwareengine_cluster#datastore_network GoogleVmwareengineCluster#datastore_network}
	DatastoreNetwork *GoogleVmwareengineClusterDatastoreMountConfigDatastoreNetwork `field:"required" json:"datastoreNetwork" yaml:"datastoreNetwork"`
	// Optional. NFS is accessed by hosts in either read or read_write mode Default value used will be READ_WRITE Possible values: READ_ONLY READ_WRITE.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_vmwareengine_cluster#access_mode GoogleVmwareengineCluster#access_mode}
	AccessMode *string `field:"optional" json:"accessMode" yaml:"accessMode"`
	// Optional.
	//
	// If set to true, the colocation requirement will be ignored.
	// If set to false, the colocation requirement will be enforced.
	// Colocation requirement is the requirement that the cluster must be in the
	// same region/zone of datastore.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_vmwareengine_cluster#ignore_colocation GoogleVmwareengineCluster#ignore_colocation}
	IgnoreColocation interface{} `field:"optional" json:"ignoreColocation" yaml:"ignoreColocation"`
	// Optional. The NFS protocol supported by the NFS volume. Default value used will be NFS_V3 Possible values: NFS_V3.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_vmwareengine_cluster#nfs_version GoogleVmwareengineCluster#nfs_version}
	NfsVersion *string `field:"optional" json:"nfsVersion" yaml:"nfsVersion"`
}

