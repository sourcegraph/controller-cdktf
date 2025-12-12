package catalogentry

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CatalogEntryConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/catalog_entry#attribute_values CatalogEntry#attribute_values}.
	AttributeValues interface{} `field:"required" json:"attributeValues" yaml:"attributeValues"`
	// ID of this catalog type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/catalog_entry#catalog_type_id CatalogEntry#catalog_type_id}
	CatalogTypeId *string `field:"required" json:"catalogTypeId" yaml:"catalogTypeId"`
	// Name is the human readable name of this entry.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/catalog_entry#name CatalogEntry#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Optional aliases that can be used to reference this entry.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/catalog_entry#aliases CatalogEntry#aliases}
	Aliases *[]*string `field:"optional" json:"aliases" yaml:"aliases"`
	// An optional alternative ID for this entry, which is ensured to be unique for the type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/catalog_entry#external_id CatalogEntry#external_id}
	ExternalId *string `field:"optional" json:"externalId" yaml:"externalId"`
	// The set of attributes that are managed by this resource. By default, all attributes are managed by this resource.
	//
	// This can be used to allow other attributes of a catalog entry to be managed elsewhere, for example in another Terraform repository or the incident.io web UI.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/catalog_entry#managed_attributes CatalogEntry#managed_attributes}
	ManagedAttributes *[]*string `field:"optional" json:"managedAttributes" yaml:"managedAttributes"`
	// When catalog type is ranked, this is used to help order things.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/catalog_entry#rank CatalogEntry#rank}
	Rank *float64 `field:"optional" json:"rank" yaml:"rank"`
}

