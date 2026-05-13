package googlebiglakeicebergnamespace

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleBiglakeIcebergNamespaceConfig struct {
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
	// The name of the IcebergCatalog.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_biglake_iceberg_namespace#catalog GoogleBiglakeIcebergNamespace#catalog}
	Catalog *string `field:"required" json:"catalog" yaml:"catalog"`
	// The unique identifier of the namespace.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_biglake_iceberg_namespace#namespace_id GoogleBiglakeIcebergNamespace#namespace_id}
	NamespaceId *string `field:"required" json:"namespaceId" yaml:"namespaceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_biglake_iceberg_namespace#id GoogleBiglakeIcebergNamespace#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_biglake_iceberg_namespace#project GoogleBiglakeIcebergNamespace#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// User-defined properties for the namespace.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_biglake_iceberg_namespace#properties GoogleBiglakeIcebergNamespace#properties}
	Properties *map[string]*string `field:"optional" json:"properties" yaml:"properties"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_biglake_iceberg_namespace#timeouts GoogleBiglakeIcebergNamespace#timeouts}
	Timeouts *GoogleBiglakeIcebergNamespaceTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

