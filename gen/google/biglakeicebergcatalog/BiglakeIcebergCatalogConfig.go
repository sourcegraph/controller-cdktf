package biglakeicebergcatalog

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type BiglakeIcebergCatalogConfig struct {
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
	// The catalog type of the IcebergCatalog. Currently only supports the type for Google Cloud Storage Buckets. Possible values: ["CATALOG_TYPE_GCS_BUCKET"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/biglake_iceberg_catalog#catalog_type BiglakeIcebergCatalog#catalog_type}
	CatalogType *string `field:"required" json:"catalogType" yaml:"catalogType"`
	// The name of the IcebergCatalog.
	//
	// For CATALOG_TYPE_GCS_BUCKET typed catalogs, the name needs to be the
	// exact same value of the GCS bucket's name. For example, for a bucket:
	// gs://bucket-name, the catalog name will be exactly "bucket-name".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/biglake_iceberg_catalog#name BiglakeIcebergCatalog#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The credential mode used for the catalog.
	//
	// CREDENTIAL_MODE_END_USER - End user credentials, default. The authenticating user must have access to the catalog resources and the corresponding Google Cloud Storage files. CREDENTIAL_MODE_VENDED_CREDENTIALS - Use credential vending. The authenticating user must have access to the catalog resources and the system will provide the caller with downscoped credentials to access the Google Cloud Storage files. All table operations in this mode would require 'X-Iceberg-Access-Delegation' header with 'vended-credentials' value included. System will generate a service account and the catalog administrator must grant the service account appropriate permissions. Possible values: ["CREDENTIAL_MODE_END_USER", "CREDENTIAL_MODE_VENDED_CREDENTIALS"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/biglake_iceberg_catalog#credential_mode BiglakeIcebergCatalog#credential_mode}
	CredentialMode *string `field:"optional" json:"credentialMode" yaml:"credentialMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/biglake_iceberg_catalog#id BiglakeIcebergCatalog#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The primary location for mirroring the remote catalog metadata.
	//
	// It must be
	// a BigLake-supported location, and it should be proximate to the remote
	// catalog's location.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/biglake_iceberg_catalog#primary_location BiglakeIcebergCatalog#primary_location}
	PrimaryLocation *string `field:"optional" json:"primaryLocation" yaml:"primaryLocation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/biglake_iceberg_catalog#project BiglakeIcebergCatalog#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/biglake_iceberg_catalog#timeouts BiglakeIcebergCatalog#timeouts}
	Timeouts *BiglakeIcebergCatalogTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

