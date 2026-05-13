package oracledatabaseexadbvmcluster


type OracleDatabaseExadbVmClusterPropertiesVmFileSystemStorage struct {
	// The storage allocation for the exadbvmcluster per node, in gigabytes (GB).
	//
	// This field is used to calculate the total storage allocation for the
	// exadbvmcluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/oracle_database_exadb_vm_cluster#size_in_gbs_per_node OracleDatabaseExadbVmCluster#size_in_gbs_per_node}
	SizeInGbsPerNode *float64 `field:"required" json:"sizeInGbsPerNode" yaml:"sizeInGbsPerNode"`
}

