package dataincidentcatalogtypeattribute

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataIncidentCatalogTypeAttributeConfig struct {
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
	// ID of this catalog type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.17.1/docs/data-sources/catalog_type_attribute#catalog_type_id DataIncidentCatalogTypeAttribute#catalog_type_id}
	CatalogTypeId *string `field:"required" json:"catalogTypeId" yaml:"catalogTypeId"`
	// The name of this attribute.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.17.1/docs/data-sources/catalog_type_attribute#name DataIncidentCatalogTypeAttribute#name}
	Name *string `field:"required" json:"name" yaml:"name"`
}

