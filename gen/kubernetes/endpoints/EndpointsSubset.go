package endpoints

type EndpointsSubset struct {
	// address block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.15.0/docs/resources/endpoints#address Endpoints#address}
	Address any `field:"optional" json:"address" yaml:"address"`
	// not_ready_address block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.15.0/docs/resources/endpoints#not_ready_address Endpoints#not_ready_address}
	NotReadyAddress any `field:"optional" json:"notReadyAddress" yaml:"notReadyAddress"`
	// port block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.15.0/docs/resources/endpoints#port Endpoints#port}
	Port any `field:"optional" json:"port" yaml:"port"`
}
