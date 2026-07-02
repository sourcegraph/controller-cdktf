package apphubapplication

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/apphubapplication/internal"
)

type ApphubApplicationAttributesOutputReference interface {
	cdktf.ComplexObject
	BusinessOwners() ApphubApplicationAttributesBusinessOwnersList
	BusinessOwnersInput() any
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
	Criticality() ApphubApplicationAttributesCriticalityOutputReference
	CriticalityInput() *ApphubApplicationAttributesCriticality
	DeveloperOwners() ApphubApplicationAttributesDeveloperOwnersList
	DeveloperOwnersInput() any
	Environment() ApphubApplicationAttributesEnvironmentOutputReference
	EnvironmentInput() *ApphubApplicationAttributesEnvironment
	// Experimental.
	Fqn() *string
	InternalValue() *ApphubApplicationAttributes
	SetInternalValue(val *ApphubApplicationAttributes)
	OperatorOwners() ApphubApplicationAttributesOperatorOwnersList
	OperatorOwnersInput() any
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
	PutBusinessOwners(value any)
	PutCriticality(value *ApphubApplicationAttributesCriticality)
	PutDeveloperOwners(value any)
	PutEnvironment(value *ApphubApplicationAttributesEnvironment)
	PutOperatorOwners(value any)
	ResetBusinessOwners()
	ResetCriticality()
	ResetDeveloperOwners()
	ResetEnvironment()
	ResetOperatorOwners()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ApphubApplicationAttributesOutputReference
type jsiiProxy_ApphubApplicationAttributesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ApphubApplicationAttributesOutputReference) BusinessOwners() ApphubApplicationAttributesBusinessOwnersList {
	var returns ApphubApplicationAttributesBusinessOwnersList
	_jsii_.Get(
		j,
		"businessOwners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApphubApplicationAttributesOutputReference) BusinessOwnersInput() any {
	var returns any
	_jsii_.Get(
		j,
		"businessOwnersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApphubApplicationAttributesOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApphubApplicationAttributesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApphubApplicationAttributesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApphubApplicationAttributesOutputReference) Criticality() ApphubApplicationAttributesCriticalityOutputReference {
	var returns ApphubApplicationAttributesCriticalityOutputReference
	_jsii_.Get(
		j,
		"criticality",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApphubApplicationAttributesOutputReference) CriticalityInput() *ApphubApplicationAttributesCriticality {
	var returns *ApphubApplicationAttributesCriticality
	_jsii_.Get(
		j,
		"criticalityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApphubApplicationAttributesOutputReference) DeveloperOwners() ApphubApplicationAttributesDeveloperOwnersList {
	var returns ApphubApplicationAttributesDeveloperOwnersList
	_jsii_.Get(
		j,
		"developerOwners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApphubApplicationAttributesOutputReference) DeveloperOwnersInput() any {
	var returns any
	_jsii_.Get(
		j,
		"developerOwnersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApphubApplicationAttributesOutputReference) Environment() ApphubApplicationAttributesEnvironmentOutputReference {
	var returns ApphubApplicationAttributesEnvironmentOutputReference
	_jsii_.Get(
		j,
		"environment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApphubApplicationAttributesOutputReference) EnvironmentInput() *ApphubApplicationAttributesEnvironment {
	var returns *ApphubApplicationAttributesEnvironment
	_jsii_.Get(
		j,
		"environmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApphubApplicationAttributesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApphubApplicationAttributesOutputReference) InternalValue() *ApphubApplicationAttributes {
	var returns *ApphubApplicationAttributes
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApphubApplicationAttributesOutputReference) OperatorOwners() ApphubApplicationAttributesOperatorOwnersList {
	var returns ApphubApplicationAttributesOperatorOwnersList
	_jsii_.Get(
		j,
		"operatorOwners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApphubApplicationAttributesOutputReference) OperatorOwnersInput() any {
	var returns any
	_jsii_.Get(
		j,
		"operatorOwnersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApphubApplicationAttributesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApphubApplicationAttributesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewApphubApplicationAttributesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ApphubApplicationAttributesOutputReference {
	_init_.Initialize()

	if err := validateNewApphubApplicationAttributesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ApphubApplicationAttributesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.apphubApplication.ApphubApplicationAttributesOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewApphubApplicationAttributesOutputReference_Override(a ApphubApplicationAttributesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.apphubApplication.ApphubApplicationAttributesOutputReference",
		[]any{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_ApphubApplicationAttributesOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ApphubApplicationAttributesOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ApphubApplicationAttributesOutputReference) SetInternalValue(val *ApphubApplicationAttributes) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ApphubApplicationAttributesOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ApphubApplicationAttributesOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_ApphubApplicationAttributesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApphubApplicationAttributesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]any

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApphubApplicationAttributesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApphubApplicationAttributesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApphubApplicationAttributesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApphubApplicationAttributesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApphubApplicationAttributesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApphubApplicationAttributesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApphubApplicationAttributesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApphubApplicationAttributesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApphubApplicationAttributesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApphubApplicationAttributesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]any{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApphubApplicationAttributesOutputReference) PutBusinessOwners(value any) {
	if err := a.validatePutBusinessOwnersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putBusinessOwners",
		[]any{value},
	)
}

func (a *jsiiProxy_ApphubApplicationAttributesOutputReference) PutCriticality(value *ApphubApplicationAttributesCriticality) {
	if err := a.validatePutCriticalityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCriticality",
		[]any{value},
	)
}

func (a *jsiiProxy_ApphubApplicationAttributesOutputReference) PutDeveloperOwners(value any) {
	if err := a.validatePutDeveloperOwnersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDeveloperOwners",
		[]any{value},
	)
}

func (a *jsiiProxy_ApphubApplicationAttributesOutputReference) PutEnvironment(value *ApphubApplicationAttributesEnvironment) {
	if err := a.validatePutEnvironmentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEnvironment",
		[]any{value},
	)
}

func (a *jsiiProxy_ApphubApplicationAttributesOutputReference) PutOperatorOwners(value any) {
	if err := a.validatePutOperatorOwnersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putOperatorOwners",
		[]any{value},
	)
}

func (a *jsiiProxy_ApphubApplicationAttributesOutputReference) ResetBusinessOwners() {
	_jsii_.InvokeVoid(
		a,
		"resetBusinessOwners",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApphubApplicationAttributesOutputReference) ResetCriticality() {
	_jsii_.InvokeVoid(
		a,
		"resetCriticality",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApphubApplicationAttributesOutputReference) ResetDeveloperOwners() {
	_jsii_.InvokeVoid(
		a,
		"resetDeveloperOwners",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApphubApplicationAttributesOutputReference) ResetEnvironment() {
	_jsii_.InvokeVoid(
		a,
		"resetEnvironment",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApphubApplicationAttributesOutputReference) ResetOperatorOwners() {
	_jsii_.InvokeVoid(
		a,
		"resetOperatorOwners",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApphubApplicationAttributesOutputReference) Resolve(_context cdktf.IResolveContext) any {
	if err := a.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns any

	_jsii_.Invoke(
		a,
		"resolve",
		[]any{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApphubApplicationAttributesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
