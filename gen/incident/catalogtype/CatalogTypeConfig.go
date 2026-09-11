package catalogtype

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type CatalogTypeConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Human readble description of this type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/catalog_type#description CatalogType#description}
	Description *string `field:"required" json:"description" yaml:"description"`
	// Name is the human readable name of this type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/catalog_type#name CatalogType#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The url of the external repository where this type is managed.
	//
	// Users will not be able to edit the catalog type (or its entries) via the UI, and will instead be provided a link to this URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/catalog_type#source_repo_url CatalogType#source_repo_url}
	SourceRepoUrl *string `field:"required" json:"sourceRepoUrl" yaml:"sourceRepoUrl"`
	// The categories that this type belongs to, to be shown in the web dashboard.
	//
	// Possible values are: `customer`, `issue-tracker`, `product-feature`, `service`, `on-call`, `team`, `user`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/catalog_type#categories CatalogType#categories}
	Categories *[]*string `field:"optional" json:"categories" yaml:"categories"`
	// The type name of this catalog type, to be used when defining attributes.
	//
	// This is immutable once a CatalogType has been created. For non-externally sync types, it must follow the pattern Custom["SomeName"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/catalog_type#type_name CatalogType#type_name}
	TypeName *string `field:"optional" json:"typeName" yaml:"typeName"`
	// If enabled, you can refer to entries of this type by their name, as well as their external ID and any aliases.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/catalog_type#use_name_as_identifier CatalogType#use_name_as_identifier}
	UseNameAsIdentifier interface{} `field:"optional" json:"useNameAsIdentifier" yaml:"useNameAsIdentifier"`
}

