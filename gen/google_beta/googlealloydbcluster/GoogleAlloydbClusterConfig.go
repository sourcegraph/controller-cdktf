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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.76.0/docs/resources/google_alloydb_cluster#cluster_id GoogleAlloydbCluster#cluster_id}
	ClusterId *string `field:"required" json:"clusterId" yaml:"clusterId"`
	// The location where the alloydb cluster should reside.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.76.0/docs/resources/google_alloydb_cluster#location GoogleAlloydbCluster#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// The relative resource name of the VPC network on which the instance can be accessed.
	//
	// It is specified in the following form:
	//
	// "projects/{projectNumber}/global/networks/{network_id}".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.76.0/docs/resources/google_alloydb_cluster#network GoogleAlloydbCluster#network}
	Network *string `field:"required" json:"network" yaml:"network"`
	// automated_backup_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.76.0/docs/resources/google_alloydb_cluster#automated_backup_policy GoogleAlloydbCluster#automated_backup_policy}
	AutomatedBackupPolicy *GoogleAlloydbClusterAutomatedBackupPolicy `field:"optional" json:"automatedBackupPolicy" yaml:"automatedBackupPolicy"`
	// User-settable and human-readable display name for the Cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.76.0/docs/resources/google_alloydb_cluster#display_name GoogleAlloydbCluster#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// encryption_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.76.0/docs/resources/google_alloydb_cluster#encryption_config GoogleAlloydbCluster#encryption_config}
	EncryptionConfig *GoogleAlloydbClusterEncryptionConfig `field:"optional" json:"encryptionConfig" yaml:"encryptionConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.76.0/docs/resources/google_alloydb_cluster#id GoogleAlloydbCluster#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// initial_user block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.76.0/docs/resources/google_alloydb_cluster#initial_user GoogleAlloydbCluster#initial_user}
	InitialUser *GoogleAlloydbClusterInitialUser `field:"optional" json:"initialUser" yaml:"initialUser"`
	// User-defined labels for the alloydb cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.76.0/docs/resources/google_alloydb_cluster#labels GoogleAlloydbCluster#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.76.0/docs/resources/google_alloydb_cluster#project GoogleAlloydbCluster#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.76.0/docs/resources/google_alloydb_cluster#timeouts GoogleAlloydbCluster#timeouts}
	Timeouts *GoogleAlloydbClusterTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

