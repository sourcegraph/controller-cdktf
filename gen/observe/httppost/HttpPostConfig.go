package httppost

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type HttpPostConfig struct {
	// Experimental.
	Connection any `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count any `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]any `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Data to submit to Observe collector.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/http_post#data HttpPost#data}
	Data *string `field:"required" json:"data" yaml:"data"`
	// Content Type for HTTP POST request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/http_post#content_type HttpPost#content_type}
	ContentType *string `field:"optional" json:"contentType" yaml:"contentType"`
	// Additional HTTP headers.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/http_post#headers HttpPost#headers}
	Headers *map[string]*string `field:"optional" json:"headers" yaml:"headers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/http_post#id HttpPost#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Key used to tag submitted observations with unique ID. Set to empty string to omit tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/http_post#id_tag HttpPost#id_tag}
	IdTag *string `field:"optional" json:"idTag" yaml:"idTag"`
	// Path under which to submit observations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/http_post#path HttpPost#path}
	Path *string `field:"optional" json:"path" yaml:"path"`
	// Tags to set on submitted observations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/http_post#tags HttpPost#tags}
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}
