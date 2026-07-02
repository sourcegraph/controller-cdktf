package networksecurityauthzpolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/networksecurityauthzpolicy/internal"
)

type NetworkSecurityAuthzPolicyHttpRulesToOperationsMcpOutputReference interface {
	cdktf.ComplexObject
	BaseProtocolMethodsOption() *string
	SetBaseProtocolMethodsOption(val *string)
	BaseProtocolMethodsOptionInput() *string
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() any
	// Experimental.
	SetComplexObjectIndex(val any)
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
	InternalValue() *NetworkSecurityAuthzPolicyHttpRulesToOperationsMcp
	SetInternalValue(val *NetworkSecurityAuthzPolicyHttpRulesToOperationsMcp)
	Methods() NetworkSecurityAuthzPolicyHttpRulesToOperationsMcpMethodsList
	MethodsInput() any
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]any
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
	PutMethods(value any)
	ResetBaseProtocolMethodsOption()
	ResetMethods()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for NetworkSecurityAuthzPolicyHttpRulesToOperationsMcpOutputReference
type jsiiProxy_NetworkSecurityAuthzPolicyHttpRulesToOperationsMcpOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_NetworkSecurityAuthzPolicyHttpRulesToOperationsMcpOutputReference) BaseProtocolMethodsOption() *string {
	var returns *string
	_jsii_.Get(
		j,
		"baseProtocolMethodsOption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityAuthzPolicyHttpRulesToOperationsMcpOutputReference) BaseProtocolMethodsOptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"baseProtocolMethodsOptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityAuthzPolicyHttpRulesToOperationsMcpOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityAuthzPolicyHttpRulesToOperationsMcpOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityAuthzPolicyHttpRulesToOperationsMcpOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityAuthzPolicyHttpRulesToOperationsMcpOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityAuthzPolicyHttpRulesToOperationsMcpOutputReference) InternalValue() *NetworkSecurityAuthzPolicyHttpRulesToOperationsMcp {
	var returns *NetworkSecurityAuthzPolicyHttpRulesToOperationsMcp
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityAuthzPolicyHttpRulesToOperationsMcpOutputReference) Methods() NetworkSecurityAuthzPolicyHttpRulesToOperationsMcpMethodsList {
	var returns NetworkSecurityAuthzPolicyHttpRulesToOperationsMcpMethodsList
	_jsii_.Get(
		j,
		"methods",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityAuthzPolicyHttpRulesToOperationsMcpOutputReference) MethodsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"methodsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityAuthzPolicyHttpRulesToOperationsMcpOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityAuthzPolicyHttpRulesToOperationsMcpOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewNetworkSecurityAuthzPolicyHttpRulesToOperationsMcpOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) NetworkSecurityAuthzPolicyHttpRulesToOperationsMcpOutputReference {
	_init_.Initialize()

	if err := validateNewNetworkSecurityAuthzPolicyHttpRulesToOperationsMcpOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_NetworkSecurityAuthzPolicyHttpRulesToOperationsMcpOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.networkSecurityAuthzPolicy.NetworkSecurityAuthzPolicyHttpRulesToOperationsMcpOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewNetworkSecurityAuthzPolicyHttpRulesToOperationsMcpOutputReference_Override(n NetworkSecurityAuthzPolicyHttpRulesToOperationsMcpOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.networkSecurityAuthzPolicy.NetworkSecurityAuthzPolicyHttpRulesToOperationsMcpOutputReference",
		[]any{terraformResource, terraformAttribute},
		n,
	)
}

func (j *jsiiProxy_NetworkSecurityAuthzPolicyHttpRulesToOperationsMcpOutputReference) SetBaseProtocolMethodsOption(val *string) {
	if err := j.validateSetBaseProtocolMethodsOptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"baseProtocolMethodsOption",
		val,
	)
}

func (j *jsiiProxy_NetworkSecurityAuthzPolicyHttpRulesToOperationsMcpOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_NetworkSecurityAuthzPolicyHttpRulesToOperationsMcpOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_NetworkSecurityAuthzPolicyHttpRulesToOperationsMcpOutputReference) SetInternalValue(val *NetworkSecurityAuthzPolicyHttpRulesToOperationsMcp) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_NetworkSecurityAuthzPolicyHttpRulesToOperationsMcpOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_NetworkSecurityAuthzPolicyHttpRulesToOperationsMcpOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (n *jsiiProxy_NetworkSecurityAuthzPolicyHttpRulesToOperationsMcpOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkSecurityAuthzPolicyHttpRulesToOperationsMcpOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
	if err := n.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]any

	_jsii_.Invoke(
		n,
		"getAnyMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkSecurityAuthzPolicyHttpRulesToOperationsMcpOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := n.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"getBooleanAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkSecurityAuthzPolicyHttpRulesToOperationsMcpOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := n.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		n,
		"getBooleanMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkSecurityAuthzPolicyHttpRulesToOperationsMcpOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := n.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		n,
		"getListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkSecurityAuthzPolicyHttpRulesToOperationsMcpOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := n.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		n,
		"getNumberAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkSecurityAuthzPolicyHttpRulesToOperationsMcpOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := n.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		n,
		"getNumberListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkSecurityAuthzPolicyHttpRulesToOperationsMcpOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := n.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		n,
		"getNumberMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkSecurityAuthzPolicyHttpRulesToOperationsMcpOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := n.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		n,
		"getStringAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkSecurityAuthzPolicyHttpRulesToOperationsMcpOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := n.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		n,
		"getStringMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkSecurityAuthzPolicyHttpRulesToOperationsMcpOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkSecurityAuthzPolicyHttpRulesToOperationsMcpOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := n.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationForAttribute",
		[]any{property},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkSecurityAuthzPolicyHttpRulesToOperationsMcpOutputReference) PutMethods(value any) {
	if err := n.validatePutMethodsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putMethods",
		[]any{value},
	)
}

func (n *jsiiProxy_NetworkSecurityAuthzPolicyHttpRulesToOperationsMcpOutputReference) ResetBaseProtocolMethodsOption() {
	_jsii_.InvokeVoid(
		n,
		"resetBaseProtocolMethodsOption",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkSecurityAuthzPolicyHttpRulesToOperationsMcpOutputReference) ResetMethods() {
	_jsii_.InvokeVoid(
		n,
		"resetMethods",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkSecurityAuthzPolicyHttpRulesToOperationsMcpOutputReference) Resolve(_context cdktf.IResolveContext) any {
	if err := n.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns any

	_jsii_.Invoke(
		n,
		"resolve",
		[]any{_context},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkSecurityAuthzPolicyHttpRulesToOperationsMcpOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
