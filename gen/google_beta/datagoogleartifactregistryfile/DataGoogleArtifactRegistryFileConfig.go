package datagoogleartifactregistryfile

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataGoogleArtifactRegistryFileConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/data-sources/google_artifact_registry_file#file_id DataGoogleArtifactRegistryFile#file_id}.
	FileId *string `field:"required" json:"fileId" yaml:"fileId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/data-sources/google_artifact_registry_file#location DataGoogleArtifactRegistryFile#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/data-sources/google_artifact_registry_file#output_path DataGoogleArtifactRegistryFile#output_path}.
	OutputPath *string `field:"required" json:"outputPath" yaml:"outputPath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/data-sources/google_artifact_registry_file#repository_id DataGoogleArtifactRegistryFile#repository_id}.
	RepositoryId *string `field:"required" json:"repositoryId" yaml:"repositoryId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/data-sources/google_artifact_registry_file#id DataGoogleArtifactRegistryFile#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/data-sources/google_artifact_registry_file#overwrite DataGoogleArtifactRegistryFile#overwrite}.
	Overwrite interface{} `field:"optional" json:"overwrite" yaml:"overwrite"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/data-sources/google_artifact_registry_file#project DataGoogleArtifactRegistryFile#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/data-sources/google_artifact_registry_file#timeouts DataGoogleArtifactRegistryFile#timeouts}
	Timeouts *DataGoogleArtifactRegistryFileTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

