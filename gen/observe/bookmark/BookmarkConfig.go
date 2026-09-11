package bookmark

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type BookmarkConfig struct {
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
	// OID of the bookmark group used to organize the bookmark.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/bookmark#group Bookmark#group}
	Group *string `field:"required" json:"group" yaml:"group"`
	// Name of the bookmark.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/bookmark#name Bookmark#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The OID of the resource to bookmark. Datasets and dashboards are currently supported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/bookmark#target Bookmark#target}
	Target *string `field:"required" json:"target" yaml:"target"`
	// The target page for your bookmark.
	//
	// An appropriate default will be used if not set.
	//  Accepted values: `dataset`, `dashboard`, `log_explorer`, `metric_explorer`
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/bookmark#bookmark_kind Bookmark#bookmark_kind}
	BookmarkKind *string `field:"optional" json:"bookmarkKind" yaml:"bookmarkKind"`
	// Icon to be displayed for this object. Icons are sourced from the [fluency-filled](https://icons8.com/icons/fluency-systems-filled) icon set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/bookmark#icon_url Bookmark#icon_url}
	IconUrl *string `field:"optional" json:"iconUrl" yaml:"iconUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/bookmark#id Bookmark#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

