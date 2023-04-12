package googleprojectorganizationpolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/googlebeta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/googlebeta/googleprojectorganizationpolicy/internal"
)

type GoogleProjectOrganizationPolicyListPolicyOutputReference interface {
	cdktf.ComplexObject
	Allow() GoogleProjectOrganizationPolicyListPolicyAllowOutputReference
	AllowInput() *GoogleProjectOrganizationPolicyListPolicyAllow
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
	Deny() GoogleProjectOrganizationPolicyListPolicyDenyOutputReference
	DenyInput() *GoogleProjectOrganizationPolicyListPolicyDeny
	// Experimental.
	Fqn() *string
	InheritFromParent() interface{}
	SetInheritFromParent(val interface{})
	InheritFromParentInput() interface{}
	InternalValue() *GoogleProjectOrganizationPolicyListPolicy
	SetInternalValue(val *GoogleProjectOrganizationPolicyListPolicy)
	SuggestedValue() *string
	SetSuggestedValue(val *string)
	SuggestedValueInput() *string
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
	PutAllow(value *GoogleProjectOrganizationPolicyListPolicyAllow)
	PutDeny(value *GoogleProjectOrganizationPolicyListPolicyDeny)
	ResetAllow()
	ResetDeny()
	ResetInheritFromParent()
	ResetSuggestedValue()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleProjectOrganizationPolicyListPolicyOutputReference
type jsiiProxy_GoogleProjectOrganizationPolicyListPolicyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleProjectOrganizationPolicyListPolicyOutputReference) Allow() GoogleProjectOrganizationPolicyListPolicyAllowOutputReference {
	var returns GoogleProjectOrganizationPolicyListPolicyAllowOutputReference
	_jsii_.Get(
		j,
		"allow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProjectOrganizationPolicyListPolicyOutputReference) AllowInput() *GoogleProjectOrganizationPolicyListPolicyAllow {
	var returns *GoogleProjectOrganizationPolicyListPolicyAllow
	_jsii_.Get(
		j,
		"allowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProjectOrganizationPolicyListPolicyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProjectOrganizationPolicyListPolicyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProjectOrganizationPolicyListPolicyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProjectOrganizationPolicyListPolicyOutputReference) Deny() GoogleProjectOrganizationPolicyListPolicyDenyOutputReference {
	var returns GoogleProjectOrganizationPolicyListPolicyDenyOutputReference
	_jsii_.Get(
		j,
		"deny",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProjectOrganizationPolicyListPolicyOutputReference) DenyInput() *GoogleProjectOrganizationPolicyListPolicyDeny {
	var returns *GoogleProjectOrganizationPolicyListPolicyDeny
	_jsii_.Get(
		j,
		"denyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProjectOrganizationPolicyListPolicyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProjectOrganizationPolicyListPolicyOutputReference) InheritFromParent() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"inheritFromParent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProjectOrganizationPolicyListPolicyOutputReference) InheritFromParentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"inheritFromParentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProjectOrganizationPolicyListPolicyOutputReference) InternalValue() *GoogleProjectOrganizationPolicyListPolicy {
	var returns *GoogleProjectOrganizationPolicyListPolicy
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProjectOrganizationPolicyListPolicyOutputReference) SuggestedValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"suggestedValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProjectOrganizationPolicyListPolicyOutputReference) SuggestedValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"suggestedValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProjectOrganizationPolicyListPolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProjectOrganizationPolicyListPolicyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleProjectOrganizationPolicyListPolicyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleProjectOrganizationPolicyListPolicyOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleProjectOrganizationPolicyListPolicyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleProjectOrganizationPolicyListPolicyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-googlebeta.googleProjectOrganizationPolicy.GoogleProjectOrganizationPolicyListPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleProjectOrganizationPolicyListPolicyOutputReference_Override(g GoogleProjectOrganizationPolicyListPolicyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-googlebeta.googleProjectOrganizationPolicy.GoogleProjectOrganizationPolicyListPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleProjectOrganizationPolicyListPolicyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleProjectOrganizationPolicyListPolicyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleProjectOrganizationPolicyListPolicyOutputReference)SetInheritFromParent(val interface{}) {
	if err := j.validateSetInheritFromParentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inheritFromParent",
		val,
	)
}

func (j *jsiiProxy_GoogleProjectOrganizationPolicyListPolicyOutputReference)SetInternalValue(val *GoogleProjectOrganizationPolicyListPolicy) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleProjectOrganizationPolicyListPolicyOutputReference)SetSuggestedValue(val *string) {
	if err := j.validateSetSuggestedValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"suggestedValue",
		val,
	)
}

func (j *jsiiProxy_GoogleProjectOrganizationPolicyListPolicyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleProjectOrganizationPolicyListPolicyOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleProjectOrganizationPolicyListPolicyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleProjectOrganizationPolicyListPolicyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleProjectOrganizationPolicyListPolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleProjectOrganizationPolicyListPolicyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleProjectOrganizationPolicyListPolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleProjectOrganizationPolicyListPolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleProjectOrganizationPolicyListPolicyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleProjectOrganizationPolicyListPolicyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleProjectOrganizationPolicyListPolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleProjectOrganizationPolicyListPolicyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleProjectOrganizationPolicyListPolicyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleProjectOrganizationPolicyListPolicyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleProjectOrganizationPolicyListPolicyOutputReference) PutAllow(value *GoogleProjectOrganizationPolicyListPolicyAllow) {
	if err := g.validatePutAllowParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAllow",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleProjectOrganizationPolicyListPolicyOutputReference) PutDeny(value *GoogleProjectOrganizationPolicyListPolicyDeny) {
	if err := g.validatePutDenyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDeny",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleProjectOrganizationPolicyListPolicyOutputReference) ResetAllow() {
	_jsii_.InvokeVoid(
		g,
		"resetAllow",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleProjectOrganizationPolicyListPolicyOutputReference) ResetDeny() {
	_jsii_.InvokeVoid(
		g,
		"resetDeny",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleProjectOrganizationPolicyListPolicyOutputReference) ResetInheritFromParent() {
	_jsii_.InvokeVoid(
		g,
		"resetInheritFromParent",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleProjectOrganizationPolicyListPolicyOutputReference) ResetSuggestedValue() {
	_jsii_.InvokeVoid(
		g,
		"resetSuggestedValue",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleProjectOrganizationPolicyListPolicyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleProjectOrganizationPolicyListPolicyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

