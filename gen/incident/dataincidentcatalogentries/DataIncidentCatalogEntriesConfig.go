package dataincidentcatalogentries

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataIncidentCatalogEntriesConfig struct {
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
	// The catalog type ID to list entries for.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/data-sources/catalog_entries#catalog_type_id DataIncidentCatalogEntries#catalog_type_id}
	CatalogTypeId *string `field:"required" json:"catalogTypeId" yaml:"catalogTypeId"`
}

