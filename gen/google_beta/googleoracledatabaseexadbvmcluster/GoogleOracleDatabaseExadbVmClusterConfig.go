package googleoracledatabaseexadbvmcluster

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleOracleDatabaseExadbVmClusterConfig struct {
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
	// The name of the backup OdbSubnet associated with the ExadbVmCluster. Format: projects/{project}/locations/{location}/odbNetworks/{odb_network}/odbSubnets/{odb_subnet}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_oracle_database_exadb_vm_cluster#backup_odb_subnet GoogleOracleDatabaseExadbVmCluster#backup_odb_subnet}
	BackupOdbSubnet *string `field:"required" json:"backupOdbSubnet" yaml:"backupOdbSubnet"`
	// The display name for the ExadbVmCluster.
	//
	// The name does not have to
	// be unique within your project. The name must be 1-255 characters long and
	// can only contain alphanumeric characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_oracle_database_exadb_vm_cluster#display_name GoogleOracleDatabaseExadbVmCluster#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// The ID of the ExadbVmCluster to create.
	//
	// This value is
	// restricted to (^[a-z]([a-z0-9-]{0,61}[a-z0-9])?$) and must be a maximum of
	// 63 characters in length. The value must start with a letter and end with a
	// letter or a number.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_oracle_database_exadb_vm_cluster#exadb_vm_cluster_id GoogleOracleDatabaseExadbVmCluster#exadb_vm_cluster_id}
	ExadbVmClusterId *string `field:"required" json:"exadbVmClusterId" yaml:"exadbVmClusterId"`
	// Resource ID segment making up resource 'name'. It identifies the resource within its parent collection as described in https://google.aip.dev/122.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_oracle_database_exadb_vm_cluster#location GoogleOracleDatabaseExadbVmCluster#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// The name of the OdbSubnet associated with the ExadbVmCluster for IP allocation. Format: projects/{project}/locations/{location}/odbNetworks/{odb_network}/odbSubnets/{odb_subnet}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_oracle_database_exadb_vm_cluster#odb_subnet GoogleOracleDatabaseExadbVmCluster#odb_subnet}
	OdbSubnet *string `field:"required" json:"odbSubnet" yaml:"odbSubnet"`
	// properties block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_oracle_database_exadb_vm_cluster#properties GoogleOracleDatabaseExadbVmCluster#properties}
	Properties *GoogleOracleDatabaseExadbVmClusterProperties `field:"required" json:"properties" yaml:"properties"`
	// Whether or not to allow Terraform to destroy the instance.
	//
	// Unless this field is set to false in Terraform state, a terraform destroy or terraform apply that would delete the instance will fail.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_oracle_database_exadb_vm_cluster#deletion_protection GoogleOracleDatabaseExadbVmCluster#deletion_protection}
	DeletionProtection interface{} `field:"optional" json:"deletionProtection" yaml:"deletionProtection"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_oracle_database_exadb_vm_cluster#id GoogleOracleDatabaseExadbVmCluster#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The labels or tags associated with the ExadbVmCluster.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_oracle_database_exadb_vm_cluster#labels GoogleOracleDatabaseExadbVmCluster#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// The name of the OdbNetwork associated with the ExadbVmCluster.
	//
	// Format: projects/{project}/locations/{location}/odbNetworks/{odb_network}
	// It is optional but if specified, this should match the parent ODBNetwork of
	// the OdbSubnet.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_oracle_database_exadb_vm_cluster#odb_network GoogleOracleDatabaseExadbVmCluster#odb_network}
	OdbNetwork *string `field:"optional" json:"odbNetwork" yaml:"odbNetwork"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_oracle_database_exadb_vm_cluster#project GoogleOracleDatabaseExadbVmCluster#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_oracle_database_exadb_vm_cluster#timeouts GoogleOracleDatabaseExadbVmCluster#timeouts}
	Timeouts *GoogleOracleDatabaseExadbVmClusterTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

