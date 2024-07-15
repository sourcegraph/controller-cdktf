package googlebiglakedatabase

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleBiglakeDatabaseConfig struct {
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
	// The parent catalog.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_biglake_database#catalog GoogleBiglakeDatabase#catalog}
	Catalog *string `field:"required" json:"catalog" yaml:"catalog"`
	// hive_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_biglake_database#hive_options GoogleBiglakeDatabase#hive_options}
	HiveOptions *GoogleBiglakeDatabaseHiveOptions `field:"required" json:"hiveOptions" yaml:"hiveOptions"`
	// The name of the database.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_biglake_database#name GoogleBiglakeDatabase#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The database type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_biglake_database#type GoogleBiglakeDatabase#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_biglake_database#id GoogleBiglakeDatabase#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_biglake_database#timeouts GoogleBiglakeDatabase#timeouts}
	Timeouts *GoogleBiglakeDatabaseTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

