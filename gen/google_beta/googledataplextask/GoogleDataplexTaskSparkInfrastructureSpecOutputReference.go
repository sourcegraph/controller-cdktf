package googledataplextask

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googledataplextask/internal"
)

type GoogleDataplexTaskSparkInfrastructureSpecOutputReference interface {
	cdktf.ComplexObject
	Batch() GoogleDataplexTaskSparkInfrastructureSpecBatchOutputReference
	BatchInput() *GoogleDataplexTaskSparkInfrastructureSpecBatch
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	ContainerImage() GoogleDataplexTaskSparkInfrastructureSpecContainerImageOutputReference
	ContainerImageInput() *GoogleDataplexTaskSparkInfrastructureSpecContainerImage
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleDataplexTaskSparkInfrastructureSpec
	SetInternalValue(val *GoogleDataplexTaskSparkInfrastructureSpec)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	VpcNetwork() GoogleDataplexTaskSparkInfrastructureSpecVpcNetworkOutputReference
	VpcNetworkInput() *GoogleDataplexTaskSparkInfrastructureSpecVpcNetwork
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	PutBatch(value *GoogleDataplexTaskSparkInfrastructureSpecBatch)
	PutContainerImage(value *GoogleDataplexTaskSparkInfrastructureSpecContainerImage)
	PutVpcNetwork(value *GoogleDataplexTaskSparkInfrastructureSpecVpcNetwork)
	ResetBatch()
	ResetContainerImage()
	ResetVpcNetwork()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleDataplexTaskSparkInfrastructureSpecOutputReference
type jsiiProxy_GoogleDataplexTaskSparkInfrastructureSpecOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleDataplexTaskSparkInfrastructureSpecOutputReference) Batch() GoogleDataplexTaskSparkInfrastructureSpecBatchOutputReference {
	var returns GoogleDataplexTaskSparkInfrastructureSpecBatchOutputReference
	_jsii_.Get(
		j,
		"batch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexTaskSparkInfrastructureSpecOutputReference) BatchInput() *GoogleDataplexTaskSparkInfrastructureSpecBatch {
	var returns *GoogleDataplexTaskSparkInfrastructureSpecBatch
	_jsii_.Get(
		j,
		"batchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexTaskSparkInfrastructureSpecOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexTaskSparkInfrastructureSpecOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexTaskSparkInfrastructureSpecOutputReference) ContainerImage() GoogleDataplexTaskSparkInfrastructureSpecContainerImageOutputReference {
	var returns GoogleDataplexTaskSparkInfrastructureSpecContainerImageOutputReference
	_jsii_.Get(
		j,
		"containerImage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexTaskSparkInfrastructureSpecOutputReference) ContainerImageInput() *GoogleDataplexTaskSparkInfrastructureSpecContainerImage {
	var returns *GoogleDataplexTaskSparkInfrastructureSpecContainerImage
	_jsii_.Get(
		j,
		"containerImageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexTaskSparkInfrastructureSpecOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexTaskSparkInfrastructureSpecOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexTaskSparkInfrastructureSpecOutputReference) InternalValue() *GoogleDataplexTaskSparkInfrastructureSpec {
	var returns *GoogleDataplexTaskSparkInfrastructureSpec
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexTaskSparkInfrastructureSpecOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexTaskSparkInfrastructureSpecOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexTaskSparkInfrastructureSpecOutputReference) VpcNetwork() GoogleDataplexTaskSparkInfrastructureSpecVpcNetworkOutputReference {
	var returns GoogleDataplexTaskSparkInfrastructureSpecVpcNetworkOutputReference
	_jsii_.Get(
		j,
		"vpcNetwork",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexTaskSparkInfrastructureSpecOutputReference) VpcNetworkInput() *GoogleDataplexTaskSparkInfrastructureSpecVpcNetwork {
	var returns *GoogleDataplexTaskSparkInfrastructureSpecVpcNetwork
	_jsii_.Get(
		j,
		"vpcNetworkInput",
		&returns,
	)
	return returns
}


func NewGoogleDataplexTaskSparkInfrastructureSpecOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleDataplexTaskSparkInfrastructureSpecOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleDataplexTaskSparkInfrastructureSpecOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleDataplexTaskSparkInfrastructureSpecOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleDataplexTask.GoogleDataplexTaskSparkInfrastructureSpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleDataplexTaskSparkInfrastructureSpecOutputReference_Override(g GoogleDataplexTaskSparkInfrastructureSpecOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleDataplexTask.GoogleDataplexTaskSparkInfrastructureSpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleDataplexTaskSparkInfrastructureSpecOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleDataplexTaskSparkInfrastructureSpecOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleDataplexTaskSparkInfrastructureSpecOutputReference)SetInternalValue(val *GoogleDataplexTaskSparkInfrastructureSpec) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleDataplexTaskSparkInfrastructureSpecOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleDataplexTaskSparkInfrastructureSpecOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleDataplexTaskSparkInfrastructureSpecOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataplexTaskSparkInfrastructureSpecOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := g.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataplexTaskSparkInfrastructureSpecOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataplexTaskSparkInfrastructureSpecOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := g.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		g,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataplexTaskSparkInfrastructureSpecOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := g.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		g,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataplexTaskSparkInfrastructureSpecOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := g.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		g,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataplexTaskSparkInfrastructureSpecOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := g.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		g,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataplexTaskSparkInfrastructureSpecOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := g.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		g,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataplexTaskSparkInfrastructureSpecOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := g.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		g,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataplexTaskSparkInfrastructureSpecOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := g.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		g,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataplexTaskSparkInfrastructureSpecOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataplexTaskSparkInfrastructureSpecOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataplexTaskSparkInfrastructureSpecOutputReference) PutBatch(value *GoogleDataplexTaskSparkInfrastructureSpecBatch) {
	if err := g.validatePutBatchParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putBatch",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataplexTaskSparkInfrastructureSpecOutputReference) PutContainerImage(value *GoogleDataplexTaskSparkInfrastructureSpecContainerImage) {
	if err := g.validatePutContainerImageParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putContainerImage",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataplexTaskSparkInfrastructureSpecOutputReference) PutVpcNetwork(value *GoogleDataplexTaskSparkInfrastructureSpecVpcNetwork) {
	if err := g.validatePutVpcNetworkParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putVpcNetwork",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataplexTaskSparkInfrastructureSpecOutputReference) ResetBatch() {
	_jsii_.InvokeVoid(
		g,
		"resetBatch",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataplexTaskSparkInfrastructureSpecOutputReference) ResetContainerImage() {
	_jsii_.InvokeVoid(
		g,
		"resetContainerImage",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataplexTaskSparkInfrastructureSpecOutputReference) ResetVpcNetwork() {
	_jsii_.InvokeVoid(
		g,
		"resetVpcNetwork",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataplexTaskSparkInfrastructureSpecOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := g.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		g,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataplexTaskSparkInfrastructureSpecOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

