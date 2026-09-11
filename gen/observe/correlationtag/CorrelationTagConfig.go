package correlationtag

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type CorrelationTagConfig struct {
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
	// The column to which the correlation tag should be attached.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/correlation_tag#column CorrelationTag#column}
	Column *string `field:"required" json:"column" yaml:"column"`
	// OID of the dataset to which the correlation tag should be attached.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/correlation_tag#dataset CorrelationTag#dataset}
	Dataset *string `field:"required" json:"dataset" yaml:"dataset"`
	// The name to attach.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/correlation_tag#name CorrelationTag#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/correlation_tag#id CorrelationTag#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// If the column is of type "object", a correlation tag can be attached to a key nested within the object.
	//
	// Standard Javascript notation can be used to specify the path to the key.
	// For example, say the object has the following structure -
	// {
	//   "a": {
	//     "b": {
	//       "c": "value"
	//     }
	//   }
	// }
	// Then the path to the key "c" would be "a.b.c" or "a['b']['c']"
	//
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/correlation_tag#path CorrelationTag#path}
	Path *string `field:"optional" json:"path" yaml:"path"`
}

