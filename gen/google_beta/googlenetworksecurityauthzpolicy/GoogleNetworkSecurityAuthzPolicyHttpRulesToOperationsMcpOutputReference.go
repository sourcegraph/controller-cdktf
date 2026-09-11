package googlenetworksecurityauthzpolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlenetworksecurityauthzpolicy/internal"
)

type GoogleNetworkSecurityAuthzPolicyHttpRulesToOperationsMcpOutputReference interface {
	cdktn.ComplexObject
	BaseProtocolMethodsOption() *string
	SetBaseProtocolMethodsOption(val *string)
	BaseProtocolMethodsOptionInput() *string
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
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleNetworkSecurityAuthzPolicyHttpRulesToOperationsMcp
	SetInternalValue(val *GoogleNetworkSecurityAuthzPolicyHttpRulesToOperationsMcp)
	Methods() GoogleNetworkSecurityAuthzPolicyHttpRulesToOperationsMcpMethodsList
	MethodsInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable
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
	InterpolationAsList() cdktn.IResolvable
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable
	PutMethods(value interface{})
	ResetBaseProtocolMethodsOption()
	ResetMethods()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleNetworkSecurityAuthzPolicyHttpRulesToOperationsMcpOutputReference
type jsiiProxy_GoogleNetworkSecurityAuthzPolicyHttpRulesToOperationsMcpOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleNetworkSecurityAuthzPolicyHttpRulesToOperationsMcpOutputReference) BaseProtocolMethodsOption() *string {
	var returns *string
	_jsii_.Get(
		j,
		"baseProtocolMethodsOption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkSecurityAuthzPolicyHttpRulesToOperationsMcpOutputReference) BaseProtocolMethodsOptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"baseProtocolMethodsOptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkSecurityAuthzPolicyHttpRulesToOperationsMcpOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkSecurityAuthzPolicyHttpRulesToOperationsMcpOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkSecurityAuthzPolicyHttpRulesToOperationsMcpOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkSecurityAuthzPolicyHttpRulesToOperationsMcpOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkSecurityAuthzPolicyHttpRulesToOperationsMcpOutputReference) InternalValue() *GoogleNetworkSecurityAuthzPolicyHttpRulesToOperationsMcp {
	var returns *GoogleNetworkSecurityAuthzPolicyHttpRulesToOperationsMcp
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkSecurityAuthzPolicyHttpRulesToOperationsMcpOutputReference) Methods() GoogleNetworkSecurityAuthzPolicyHttpRulesToOperationsMcpMethodsList {
	var returns GoogleNetworkSecurityAuthzPolicyHttpRulesToOperationsMcpMethodsList
	_jsii_.Get(
		j,
		"methods",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkSecurityAuthzPolicyHttpRulesToOperationsMcpOutputReference) MethodsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"methodsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkSecurityAuthzPolicyHttpRulesToOperationsMcpOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkSecurityAuthzPolicyHttpRulesToOperationsMcpOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleNetworkSecurityAuthzPolicyHttpRulesToOperationsMcpOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleNetworkSecurityAuthzPolicyHttpRulesToOperationsMcpOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleNetworkSecurityAuthzPolicyHttpRulesToOperationsMcpOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleNetworkSecurityAuthzPolicyHttpRulesToOperationsMcpOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleNetworkSecurityAuthzPolicy.GoogleNetworkSecurityAuthzPolicyHttpRulesToOperationsMcpOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleNetworkSecurityAuthzPolicyHttpRulesToOperationsMcpOutputReference_Override(g GoogleNetworkSecurityAuthzPolicyHttpRulesToOperationsMcpOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleNetworkSecurityAuthzPolicy.GoogleNetworkSecurityAuthzPolicyHttpRulesToOperationsMcpOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleNetworkSecurityAuthzPolicyHttpRulesToOperationsMcpOutputReference)SetBaseProtocolMethodsOption(val *string) {
	if err := j.validateSetBaseProtocolMethodsOptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"baseProtocolMethodsOption",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkSecurityAuthzPolicyHttpRulesToOperationsMcpOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkSecurityAuthzPolicyHttpRulesToOperationsMcpOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkSecurityAuthzPolicyHttpRulesToOperationsMcpOutputReference)SetInternalValue(val *GoogleNetworkSecurityAuthzPolicyHttpRulesToOperationsMcp) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkSecurityAuthzPolicyHttpRulesToOperationsMcpOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkSecurityAuthzPolicyHttpRulesToOperationsMcpOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleNetworkSecurityAuthzPolicyHttpRulesToOperationsMcpOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleNetworkSecurityAuthzPolicyHttpRulesToOperationsMcpOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleNetworkSecurityAuthzPolicyHttpRulesToOperationsMcpOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleNetworkSecurityAuthzPolicyHttpRulesToOperationsMcpOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleNetworkSecurityAuthzPolicyHttpRulesToOperationsMcpOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleNetworkSecurityAuthzPolicyHttpRulesToOperationsMcpOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleNetworkSecurityAuthzPolicyHttpRulesToOperationsMcpOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleNetworkSecurityAuthzPolicyHttpRulesToOperationsMcpOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleNetworkSecurityAuthzPolicyHttpRulesToOperationsMcpOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleNetworkSecurityAuthzPolicyHttpRulesToOperationsMcpOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleNetworkSecurityAuthzPolicyHttpRulesToOperationsMcpOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleNetworkSecurityAuthzPolicyHttpRulesToOperationsMcpOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleNetworkSecurityAuthzPolicyHttpRulesToOperationsMcpOutputReference) PutMethods(value interface{}) {
	if err := g.validatePutMethodsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putMethods",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleNetworkSecurityAuthzPolicyHttpRulesToOperationsMcpOutputReference) ResetBaseProtocolMethodsOption() {
	_jsii_.InvokeVoid(
		g,
		"resetBaseProtocolMethodsOption",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetworkSecurityAuthzPolicyHttpRulesToOperationsMcpOutputReference) ResetMethods() {
	_jsii_.InvokeVoid(
		g,
		"resetMethods",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetworkSecurityAuthzPolicyHttpRulesToOperationsMcpOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := g.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		g,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleNetworkSecurityAuthzPolicyHttpRulesToOperationsMcpOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

