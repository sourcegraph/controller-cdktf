package poller

type PollerHttp struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/poller#body Poller#body}.
	Body *string `field:"optional" json:"body" yaml:"body"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/poller#content_type Poller#content_type}.
	ContentType *string `field:"optional" json:"contentType" yaml:"contentType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/poller#endpoint Poller#endpoint}.
	Endpoint *string `field:"optional" json:"endpoint" yaml:"endpoint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/poller#headers Poller#headers}.
	Headers *map[string]*string `field:"optional" json:"headers" yaml:"headers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/poller#method Poller#method}.
	Method *string `field:"optional" json:"method" yaml:"method"`
	// request block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/poller#request Poller#request}
	Request any `field:"optional" json:"request" yaml:"request"`
	// rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/poller#rule Poller#rule}
	Rule any `field:"optional" json:"rule" yaml:"rule"`
	// template block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/poller#template Poller#template}
	Template *PollerHttpTemplate `field:"optional" json:"template" yaml:"template"`
	// timestamp block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/poller#timestamp Poller#timestamp}
	Timestamp any `field:"optional" json:"timestamp" yaml:"timestamp"`
}
