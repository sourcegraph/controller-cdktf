package vpcpeeringconnectionaccepter

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/aws/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/aws/vpcpeeringconnectionaccepter/internal"
)

type VpcPeeringConnectionAccepterRequesterOutputReference interface {
	cdktf.ComplexObject
	AllowClassicLinkToRemoteVpc() any
	SetAllowClassicLinkToRemoteVpc(val any)
	AllowClassicLinkToRemoteVpcInput() any
	AllowRemoteVpcDnsResolution() any
	SetAllowRemoteVpcDnsResolution(val any)
	AllowRemoteVpcDnsResolutionInput() any
	AllowVpcToRemoteClassicLink() any
	SetAllowVpcToRemoteClassicLink(val any)
	AllowVpcToRemoteClassicLinkInput() any
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
	InternalValue() *VpcPeeringConnectionAccepterRequester
	SetInternalValue(val *VpcPeeringConnectionAccepterRequester)
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
	ResetAllowClassicLinkToRemoteVpc()
	ResetAllowRemoteVpcDnsResolution()
	ResetAllowVpcToRemoteClassicLink()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for VpcPeeringConnectionAccepterRequesterOutputReference
type jsiiProxy_VpcPeeringConnectionAccepterRequesterOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_VpcPeeringConnectionAccepterRequesterOutputReference) AllowClassicLinkToRemoteVpc() any {
	var returns any
	_jsii_.Get(
		j,
		"allowClassicLinkToRemoteVpc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcPeeringConnectionAccepterRequesterOutputReference) AllowClassicLinkToRemoteVpcInput() any {
	var returns any
	_jsii_.Get(
		j,
		"allowClassicLinkToRemoteVpcInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcPeeringConnectionAccepterRequesterOutputReference) AllowRemoteVpcDnsResolution() any {
	var returns any
	_jsii_.Get(
		j,
		"allowRemoteVpcDnsResolution",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcPeeringConnectionAccepterRequesterOutputReference) AllowRemoteVpcDnsResolutionInput() any {
	var returns any
	_jsii_.Get(
		j,
		"allowRemoteVpcDnsResolutionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcPeeringConnectionAccepterRequesterOutputReference) AllowVpcToRemoteClassicLink() any {
	var returns any
	_jsii_.Get(
		j,
		"allowVpcToRemoteClassicLink",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcPeeringConnectionAccepterRequesterOutputReference) AllowVpcToRemoteClassicLinkInput() any {
	var returns any
	_jsii_.Get(
		j,
		"allowVpcToRemoteClassicLinkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcPeeringConnectionAccepterRequesterOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcPeeringConnectionAccepterRequesterOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcPeeringConnectionAccepterRequesterOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcPeeringConnectionAccepterRequesterOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcPeeringConnectionAccepterRequesterOutputReference) InternalValue() *VpcPeeringConnectionAccepterRequester {
	var returns *VpcPeeringConnectionAccepterRequester
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcPeeringConnectionAccepterRequesterOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcPeeringConnectionAccepterRequesterOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewVpcPeeringConnectionAccepterRequesterOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) VpcPeeringConnectionAccepterRequesterOutputReference {
	_init_.Initialize()

	if err := validateNewVpcPeeringConnectionAccepterRequesterOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_VpcPeeringConnectionAccepterRequesterOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.vpcPeeringConnectionAccepter.VpcPeeringConnectionAccepterRequesterOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewVpcPeeringConnectionAccepterRequesterOutputReference_Override(v VpcPeeringConnectionAccepterRequesterOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.vpcPeeringConnectionAccepter.VpcPeeringConnectionAccepterRequesterOutputReference",
		[]any{terraformResource, terraformAttribute},
		v,
	)
}

func (j *jsiiProxy_VpcPeeringConnectionAccepterRequesterOutputReference) SetAllowClassicLinkToRemoteVpc(val any) {
	if err := j.validateSetAllowClassicLinkToRemoteVpcParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowClassicLinkToRemoteVpc",
		val,
	)
}

func (j *jsiiProxy_VpcPeeringConnectionAccepterRequesterOutputReference) SetAllowRemoteVpcDnsResolution(val any) {
	if err := j.validateSetAllowRemoteVpcDnsResolutionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowRemoteVpcDnsResolution",
		val,
	)
}

func (j *jsiiProxy_VpcPeeringConnectionAccepterRequesterOutputReference) SetAllowVpcToRemoteClassicLink(val any) {
	if err := j.validateSetAllowVpcToRemoteClassicLinkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowVpcToRemoteClassicLink",
		val,
	)
}

func (j *jsiiProxy_VpcPeeringConnectionAccepterRequesterOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_VpcPeeringConnectionAccepterRequesterOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_VpcPeeringConnectionAccepterRequesterOutputReference) SetInternalValue(val *VpcPeeringConnectionAccepterRequester) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_VpcPeeringConnectionAccepterRequesterOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_VpcPeeringConnectionAccepterRequesterOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (v *jsiiProxy_VpcPeeringConnectionAccepterRequesterOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpcPeeringConnectionAccepterRequesterOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
	if err := v.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]any

	_jsii_.Invoke(
		v,
		"getAnyMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpcPeeringConnectionAccepterRequesterOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := v.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"getBooleanAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpcPeeringConnectionAccepterRequesterOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := v.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		v,
		"getBooleanMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpcPeeringConnectionAccepterRequesterOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := v.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		v,
		"getListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpcPeeringConnectionAccepterRequesterOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := v.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		v,
		"getNumberAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpcPeeringConnectionAccepterRequesterOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := v.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		v,
		"getNumberListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpcPeeringConnectionAccepterRequesterOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := v.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		v,
		"getNumberMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpcPeeringConnectionAccepterRequesterOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := v.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		v,
		"getStringAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpcPeeringConnectionAccepterRequesterOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := v.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		v,
		"getStringMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpcPeeringConnectionAccepterRequesterOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpcPeeringConnectionAccepterRequesterOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := v.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationForAttribute",
		[]any{property},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpcPeeringConnectionAccepterRequesterOutputReference) ResetAllowClassicLinkToRemoteVpc() {
	_jsii_.InvokeVoid(
		v,
		"resetAllowClassicLinkToRemoteVpc",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpcPeeringConnectionAccepterRequesterOutputReference) ResetAllowRemoteVpcDnsResolution() {
	_jsii_.InvokeVoid(
		v,
		"resetAllowRemoteVpcDnsResolution",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpcPeeringConnectionAccepterRequesterOutputReference) ResetAllowVpcToRemoteClassicLink() {
	_jsii_.InvokeVoid(
		v,
		"resetAllowVpcToRemoteClassicLink",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpcPeeringConnectionAccepterRequesterOutputReference) Resolve(_context cdktf.IResolveContext) any {
	if err := v.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns any

	_jsii_.Invoke(
		v,
		"resolve",
		[]any{_context},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpcPeeringConnectionAccepterRequesterOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
