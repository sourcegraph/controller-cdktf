package vmwareenginecluster


type VmwareengineClusterDatastoreMountConfig struct {
	// The resource name of the datastore to unmount.
	//
	// The datastore requested to be mounted should be in same region/zone as the
	// cluster.
	// Resource names are schemeless URIs that follow the conventions in
	// https://cloud.google.com/apis/design/resource_names.
	// For example:
	// 'projects/my-project/locations/us-central1/datastores/my-datastore'
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/vmwareengine_cluster#datastore VmwareengineCluster#datastore}
	Datastore *string `field:"required" json:"datastore" yaml:"datastore"`
	// datastore_network block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/vmwareengine_cluster#datastore_network VmwareengineCluster#datastore_network}
	DatastoreNetwork *VmwareengineClusterDatastoreMountConfigDatastoreNetwork `field:"required" json:"datastoreNetwork" yaml:"datastoreNetwork"`
	// Optional. NFS is accessed by hosts in either read or read_write mode Default value used will be READ_WRITE Possible values: READ_ONLY READ_WRITE.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/vmwareengine_cluster#access_mode VmwareengineCluster#access_mode}
	AccessMode *string `field:"optional" json:"accessMode" yaml:"accessMode"`
	// Optional.
	//
	// If set to true, the colocation requirement will be ignored.
	// If set to false, the colocation requirement will be enforced.
	// Colocation requirement is the requirement that the cluster must be in the
	// same region/zone of datastore.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/vmwareengine_cluster#ignore_colocation VmwareengineCluster#ignore_colocation}
	IgnoreColocation interface{} `field:"optional" json:"ignoreColocation" yaml:"ignoreColocation"`
	// Optional. The NFS protocol supported by the NFS volume. Default value used will be NFS_V3 Possible values: NFS_V3.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/vmwareengine_cluster#nfs_version VmwareengineCluster#nfs_version}
	NfsVersion *string `field:"optional" json:"nfsVersion" yaml:"nfsVersion"`
}

