package appmeshroute


type AppmeshRouteSpecHttpRouteMatch struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/appmesh_route#prefix AppmeshRoute#prefix}.
	Prefix *string `field:"required" json:"prefix" yaml:"prefix"`
	// header block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/appmesh_route#header AppmeshRoute#header}
	Header interface{} `field:"optional" json:"header" yaml:"header"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/appmesh_route#method AppmeshRoute#method}.
	Method *string `field:"optional" json:"method" yaml:"method"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/appmesh_route#port AppmeshRoute#port}.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/appmesh_route#scheme AppmeshRoute#scheme}.
	Scheme *string `field:"optional" json:"scheme" yaml:"scheme"`
}

