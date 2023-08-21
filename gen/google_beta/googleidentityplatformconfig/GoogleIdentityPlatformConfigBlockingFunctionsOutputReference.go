package googleidentityplatformconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googleidentityplatformconfig/internal"
)

type GoogleIdentityPlatformConfigBlockingFunctionsOutputReference interface {
	cdktf.ComplexObject
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
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	ForwardInboundCredentials() GoogleIdentityPlatformConfigBlockingFunctionsForwardInboundCredentialsOutputReference
	ForwardInboundCredentialsInput() *GoogleIdentityPlatformConfigBlockingFunctionsForwardInboundCredentials
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleIdentityPlatformConfigBlockingFunctions
	SetInternalValue(val *GoogleIdentityPlatformConfigBlockingFunctions)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Triggers() GoogleIdentityPlatformConfigBlockingFunctionsTriggersList
	TriggersInput() interface{}
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
	PutForwardInboundCredentials(value *GoogleIdentityPlatformConfigBlockingFunctionsForwardInboundCredentials)
	PutTriggers(value interface{})
	ResetForwardInboundCredentials()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleIdentityPlatformConfigBlockingFunctionsOutputReference
type jsiiProxy_GoogleIdentityPlatformConfigBlockingFunctionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleIdentityPlatformConfigBlockingFunctionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIdentityPlatformConfigBlockingFunctionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIdentityPlatformConfigBlockingFunctionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIdentityPlatformConfigBlockingFunctionsOutputReference) ForwardInboundCredentials() GoogleIdentityPlatformConfigBlockingFunctionsForwardInboundCredentialsOutputReference {
	var returns GoogleIdentityPlatformConfigBlockingFunctionsForwardInboundCredentialsOutputReference
	_jsii_.Get(
		j,
		"forwardInboundCredentials",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIdentityPlatformConfigBlockingFunctionsOutputReference) ForwardInboundCredentialsInput() *GoogleIdentityPlatformConfigBlockingFunctionsForwardInboundCredentials {
	var returns *GoogleIdentityPlatformConfigBlockingFunctionsForwardInboundCredentials
	_jsii_.Get(
		j,
		"forwardInboundCredentialsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIdentityPlatformConfigBlockingFunctionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIdentityPlatformConfigBlockingFunctionsOutputReference) InternalValue() *GoogleIdentityPlatformConfigBlockingFunctions {
	var returns *GoogleIdentityPlatformConfigBlockingFunctions
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIdentityPlatformConfigBlockingFunctionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIdentityPlatformConfigBlockingFunctionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIdentityPlatformConfigBlockingFunctionsOutputReference) Triggers() GoogleIdentityPlatformConfigBlockingFunctionsTriggersList {
	var returns GoogleIdentityPlatformConfigBlockingFunctionsTriggersList
	_jsii_.Get(
		j,
		"triggers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIdentityPlatformConfigBlockingFunctionsOutputReference) TriggersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"triggersInput",
		&returns,
	)
	return returns
}


func NewGoogleIdentityPlatformConfigBlockingFunctionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleIdentityPlatformConfigBlockingFunctionsOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleIdentityPlatformConfigBlockingFunctionsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleIdentityPlatformConfigBlockingFunctionsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleIdentityPlatformConfig.GoogleIdentityPlatformConfigBlockingFunctionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleIdentityPlatformConfigBlockingFunctionsOutputReference_Override(g GoogleIdentityPlatformConfigBlockingFunctionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleIdentityPlatformConfig.GoogleIdentityPlatformConfigBlockingFunctionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleIdentityPlatformConfigBlockingFunctionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleIdentityPlatformConfigBlockingFunctionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleIdentityPlatformConfigBlockingFunctionsOutputReference)SetInternalValue(val *GoogleIdentityPlatformConfigBlockingFunctions) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleIdentityPlatformConfigBlockingFunctionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleIdentityPlatformConfigBlockingFunctionsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleIdentityPlatformConfigBlockingFunctionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleIdentityPlatformConfigBlockingFunctionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleIdentityPlatformConfigBlockingFunctionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleIdentityPlatformConfigBlockingFunctionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleIdentityPlatformConfigBlockingFunctionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleIdentityPlatformConfigBlockingFunctionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleIdentityPlatformConfigBlockingFunctionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleIdentityPlatformConfigBlockingFunctionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleIdentityPlatformConfigBlockingFunctionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleIdentityPlatformConfigBlockingFunctionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleIdentityPlatformConfigBlockingFunctionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleIdentityPlatformConfigBlockingFunctionsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleIdentityPlatformConfigBlockingFunctionsOutputReference) PutForwardInboundCredentials(value *GoogleIdentityPlatformConfigBlockingFunctionsForwardInboundCredentials) {
	if err := g.validatePutForwardInboundCredentialsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putForwardInboundCredentials",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleIdentityPlatformConfigBlockingFunctionsOutputReference) PutTriggers(value interface{}) {
	if err := g.validatePutTriggersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTriggers",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleIdentityPlatformConfigBlockingFunctionsOutputReference) ResetForwardInboundCredentials() {
	_jsii_.InvokeVoid(
		g,
		"resetForwardInboundCredentials",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleIdentityPlatformConfigBlockingFunctionsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleIdentityPlatformConfigBlockingFunctionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

