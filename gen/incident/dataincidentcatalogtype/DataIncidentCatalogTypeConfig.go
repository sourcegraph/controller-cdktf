package dataincidentcatalogtype

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataIncidentCatalogTypeConfig struct {
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
	// The categories that this type belongs to, to be shown in the web dashboard.
	//
	// Possible values are: `customer`, `issue-tracker`, `product-feature`, `service`, `on-call`, `team`, `user`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/data-sources/catalog_type#categories DataIncidentCatalogType#categories}
	Categories *[]*string `field:"optional" json:"categories" yaml:"categories"`
	// Name is the human readable name of this type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/data-sources/catalog_type#name DataIncidentCatalogType#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The type name of this catalog type, to be used when defining attributes.
	//
	// This is immutable once a CatalogType has been created. For non-externally sync types, it must follow the pattern Custom["SomeName"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/data-sources/catalog_type#type_name DataIncidentCatalogType#type_name}
	TypeName *string `field:"optional" json:"typeName" yaml:"typeName"`
}

