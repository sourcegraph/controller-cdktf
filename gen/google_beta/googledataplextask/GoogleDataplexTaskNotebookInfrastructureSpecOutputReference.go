package googledataplextask

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googledataplextask/internal"
)

type GoogleDataplexTaskNotebookInfrastructureSpecOutputReference interface {
	cdktf.ComplexObject
	Batch() GoogleDataplexTaskNotebookInfrastructureSpecBatchOutputReference
	BatchInput() *GoogleDataplexTaskNotebookInfrastructureSpecBatch
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
	ContainerImage() GoogleDataplexTaskNotebookInfrastructureSpecContainerImageOutputReference
	ContainerImageInput() *GoogleDataplexTaskNotebookInfrastructureSpecContainerImage
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleDataplexTaskNotebookInfrastructureSpec
	SetInternalValue(val *GoogleDataplexTaskNotebookInfrastructureSpec)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	VpcNetwork() GoogleDataplexTaskNotebookInfrastructureSpecVpcNetworkOutputReference
	VpcNetworkInput() *GoogleDataplexTaskNotebookInfrastructureSpecVpcNetwork
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
	PutBatch(value *GoogleDataplexTaskNotebookInfrastructureSpecBatch)
	PutContainerImage(value *GoogleDataplexTaskNotebookInfrastructureSpecContainerImage)
	PutVpcNetwork(value *GoogleDataplexTaskNotebookInfrastructureSpecVpcNetwork)
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

// The jsii proxy struct for GoogleDataplexTaskNotebookInfrastructureSpecOutputReference
type jsiiProxy_GoogleDataplexTaskNotebookInfrastructureSpecOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleDataplexTaskNotebookInfrastructureSpecOutputReference) Batch() GoogleDataplexTaskNotebookInfrastructureSpecBatchOutputReference {
	var returns GoogleDataplexTaskNotebookInfrastructureSpecBatchOutputReference
	_jsii_.Get(
		j,
		"batch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexTaskNotebookInfrastructureSpecOutputReference) BatchInput() *GoogleDataplexTaskNotebookInfrastructureSpecBatch {
	var returns *GoogleDataplexTaskNotebookInfrastructureSpecBatch
	_jsii_.Get(
		j,
		"batchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexTaskNotebookInfrastructureSpecOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexTaskNotebookInfrastructureSpecOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexTaskNotebookInfrastructureSpecOutputReference) ContainerImage() GoogleDataplexTaskNotebookInfrastructureSpecContainerImageOutputReference {
	var returns GoogleDataplexTaskNotebookInfrastructureSpecContainerImageOutputReference
	_jsii_.Get(
		j,
		"containerImage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexTaskNotebookInfrastructureSpecOutputReference) ContainerImageInput() *GoogleDataplexTaskNotebookInfrastructureSpecContainerImage {
	var returns *GoogleDataplexTaskNotebookInfrastructureSpecContainerImage
	_jsii_.Get(
		j,
		"containerImageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexTaskNotebookInfrastructureSpecOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexTaskNotebookInfrastructureSpecOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexTaskNotebookInfrastructureSpecOutputReference) InternalValue() *GoogleDataplexTaskNotebookInfrastructureSpec {
	var returns *GoogleDataplexTaskNotebookInfrastructureSpec
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexTaskNotebookInfrastructureSpecOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexTaskNotebookInfrastructureSpecOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexTaskNotebookInfrastructureSpecOutputReference) VpcNetwork() GoogleDataplexTaskNotebookInfrastructureSpecVpcNetworkOutputReference {
	var returns GoogleDataplexTaskNotebookInfrastructureSpecVpcNetworkOutputReference
	_jsii_.Get(
		j,
		"vpcNetwork",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexTaskNotebookInfrastructureSpecOutputReference) VpcNetworkInput() *GoogleDataplexTaskNotebookInfrastructureSpecVpcNetwork {
	var returns *GoogleDataplexTaskNotebookInfrastructureSpecVpcNetwork
	_jsii_.Get(
		j,
		"vpcNetworkInput",
		&returns,
	)
	return returns
}


func NewGoogleDataplexTaskNotebookInfrastructureSpecOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleDataplexTaskNotebookInfrastructureSpecOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleDataplexTaskNotebookInfrastructureSpecOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleDataplexTaskNotebookInfrastructureSpecOutputReference{}

	_jsii_.Create(
		"google-beta.googleDataplexTask.GoogleDataplexTaskNotebookInfrastructureSpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleDataplexTaskNotebookInfrastructureSpecOutputReference_Override(g GoogleDataplexTaskNotebookInfrastructureSpecOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"google-beta.googleDataplexTask.GoogleDataplexTaskNotebookInfrastructureSpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleDataplexTaskNotebookInfrastructureSpecOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleDataplexTaskNotebookInfrastructureSpecOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleDataplexTaskNotebookInfrastructureSpecOutputReference)SetInternalValue(val *GoogleDataplexTaskNotebookInfrastructureSpec) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleDataplexTaskNotebookInfrastructureSpecOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleDataplexTaskNotebookInfrastructureSpecOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleDataplexTaskNotebookInfrastructureSpecOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataplexTaskNotebookInfrastructureSpecOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleDataplexTaskNotebookInfrastructureSpecOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDataplexTaskNotebookInfrastructureSpecOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleDataplexTaskNotebookInfrastructureSpecOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleDataplexTaskNotebookInfrastructureSpecOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleDataplexTaskNotebookInfrastructureSpecOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleDataplexTaskNotebookInfrastructureSpecOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleDataplexTaskNotebookInfrastructureSpecOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleDataplexTaskNotebookInfrastructureSpecOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleDataplexTaskNotebookInfrastructureSpecOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataplexTaskNotebookInfrastructureSpecOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDataplexTaskNotebookInfrastructureSpecOutputReference) PutBatch(value *GoogleDataplexTaskNotebookInfrastructureSpecBatch) {
	if err := g.validatePutBatchParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putBatch",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataplexTaskNotebookInfrastructureSpecOutputReference) PutContainerImage(value *GoogleDataplexTaskNotebookInfrastructureSpecContainerImage) {
	if err := g.validatePutContainerImageParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putContainerImage",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataplexTaskNotebookInfrastructureSpecOutputReference) PutVpcNetwork(value *GoogleDataplexTaskNotebookInfrastructureSpecVpcNetwork) {
	if err := g.validatePutVpcNetworkParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putVpcNetwork",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataplexTaskNotebookInfrastructureSpecOutputReference) ResetBatch() {
	_jsii_.InvokeVoid(
		g,
		"resetBatch",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataplexTaskNotebookInfrastructureSpecOutputReference) ResetContainerImage() {
	_jsii_.InvokeVoid(
		g,
		"resetContainerImage",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataplexTaskNotebookInfrastructureSpecOutputReference) ResetVpcNetwork() {
	_jsii_.InvokeVoid(
		g,
		"resetVpcNetwork",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataplexTaskNotebookInfrastructureSpecOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleDataplexTaskNotebookInfrastructureSpecOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

