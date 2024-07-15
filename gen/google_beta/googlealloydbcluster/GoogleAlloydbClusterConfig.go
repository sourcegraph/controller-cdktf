package googlealloydbcluster

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleAlloydbClusterConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The ID of the alloydb cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_alloydb_cluster#cluster_id GoogleAlloydbCluster#cluster_id}
	ClusterId *string `field:"required" json:"clusterId" yaml:"clusterId"`
	// The location where the alloydb cluster should reside.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_alloydb_cluster#location GoogleAlloydbCluster#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// Annotations to allow client tools to store small amount of arbitrary data.
	//
	// This is distinct from labels. https://google.aip.dev/128
	// An object containing a list of "key": value pairs. Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
	//
	//
	// **Note**: This field is non-authoritative, and will only manage the annotations present in your configuration.
	// Please refer to the field 'effective_annotations' for all of the annotations present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_alloydb_cluster#annotations GoogleAlloydbCluster#annotations}
	Annotations *map[string]*string `field:"optional" json:"annotations" yaml:"annotations"`
	// automated_backup_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_alloydb_cluster#automated_backup_policy GoogleAlloydbCluster#automated_backup_policy}
	AutomatedBackupPolicy *GoogleAlloydbClusterAutomatedBackupPolicy `field:"optional" json:"automatedBackupPolicy" yaml:"automatedBackupPolicy"`
	// The type of cluster. If not set, defaults to PRIMARY. Default value: "PRIMARY" Possible values: ["PRIMARY", "SECONDARY"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_alloydb_cluster#cluster_type GoogleAlloydbCluster#cluster_type}
	ClusterType *string `field:"optional" json:"clusterType" yaml:"clusterType"`
	// continuous_backup_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_alloydb_cluster#continuous_backup_config GoogleAlloydbCluster#continuous_backup_config}
	ContinuousBackupConfig *GoogleAlloydbClusterContinuousBackupConfig `field:"optional" json:"continuousBackupConfig" yaml:"continuousBackupConfig"`
	// The database engine major version.
	//
	// This is an optional field and it's populated at the Cluster creation time. This field cannot be changed after cluster creation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_alloydb_cluster#database_version GoogleAlloydbCluster#database_version}
	DatabaseVersion *string `field:"optional" json:"databaseVersion" yaml:"databaseVersion"`
	// Policy to determine if the cluster should be deleted forcefully.
	//
	// Deleting a cluster forcefully, deletes the cluster and all its associated instances within the cluster.
	// Deleting a Secondary cluster with a secondary instance REQUIRES setting deletion_policy = "FORCE" otherwise an error is returned. This is needed as there is no support to delete just the secondary instance, and the only way to delete secondary instance is to delete the associated secondary cluster forcefully which also deletes the secondary instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_alloydb_cluster#deletion_policy GoogleAlloydbCluster#deletion_policy}
	DeletionPolicy *string `field:"optional" json:"deletionPolicy" yaml:"deletionPolicy"`
	// User-settable and human-readable display name for the Cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_alloydb_cluster#display_name GoogleAlloydbCluster#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// encryption_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_alloydb_cluster#encryption_config GoogleAlloydbCluster#encryption_config}
	EncryptionConfig *GoogleAlloydbClusterEncryptionConfig `field:"optional" json:"encryptionConfig" yaml:"encryptionConfig"`
	// For Resource freshness validation (https://google.aip.dev/154).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_alloydb_cluster#etag GoogleAlloydbCluster#etag}
	Etag *string `field:"optional" json:"etag" yaml:"etag"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_alloydb_cluster#id GoogleAlloydbCluster#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// initial_user block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_alloydb_cluster#initial_user GoogleAlloydbCluster#initial_user}
	InitialUser *GoogleAlloydbClusterInitialUser `field:"optional" json:"initialUser" yaml:"initialUser"`
	// User-defined labels for the alloydb cluster.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_alloydb_cluster#labels GoogleAlloydbCluster#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// maintenance_update_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_alloydb_cluster#maintenance_update_policy GoogleAlloydbCluster#maintenance_update_policy}
	MaintenanceUpdatePolicy *GoogleAlloydbClusterMaintenanceUpdatePolicy `field:"optional" json:"maintenanceUpdatePolicy" yaml:"maintenanceUpdatePolicy"`
	// The relative resource name of the VPC network on which the instance can be accessed.
	//
	// It is specified in the following form:
	//
	// "projects/{projectNumber}/global/networks/{network_id}".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_alloydb_cluster#network GoogleAlloydbCluster#network}
	Network *string `field:"optional" json:"network" yaml:"network"`
	// network_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_alloydb_cluster#network_config GoogleAlloydbCluster#network_config}
	NetworkConfig *GoogleAlloydbClusterNetworkConfig `field:"optional" json:"networkConfig" yaml:"networkConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_alloydb_cluster#project GoogleAlloydbCluster#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// psc_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_alloydb_cluster#psc_config GoogleAlloydbCluster#psc_config}
	PscConfig *GoogleAlloydbClusterPscConfig `field:"optional" json:"pscConfig" yaml:"pscConfig"`
	// restore_backup_source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_alloydb_cluster#restore_backup_source GoogleAlloydbCluster#restore_backup_source}
	RestoreBackupSource *GoogleAlloydbClusterRestoreBackupSource `field:"optional" json:"restoreBackupSource" yaml:"restoreBackupSource"`
	// restore_continuous_backup_source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_alloydb_cluster#restore_continuous_backup_source GoogleAlloydbCluster#restore_continuous_backup_source}
	RestoreContinuousBackupSource *GoogleAlloydbClusterRestoreContinuousBackupSource `field:"optional" json:"restoreContinuousBackupSource" yaml:"restoreContinuousBackupSource"`
	// secondary_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_alloydb_cluster#secondary_config GoogleAlloydbCluster#secondary_config}
	SecondaryConfig *GoogleAlloydbClusterSecondaryConfig `field:"optional" json:"secondaryConfig" yaml:"secondaryConfig"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_alloydb_cluster#timeouts GoogleAlloydbCluster#timeouts}
	Timeouts *GoogleAlloydbClusterTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

