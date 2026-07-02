package lustreinstance

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/lustreinstance/internal"
)

type LustreInstanceAccessRulesOptionsOutputReference interface {
	cdktf.ComplexObject
	AccessRules() LustreInstanceAccessRulesOptionsAccessRulesList
	AccessRulesInput() any
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
	DefaultSquashGid() *float64
	SetDefaultSquashGid(val *float64)
	DefaultSquashGidInput() *float64
	DefaultSquashMode() *string
	SetDefaultSquashMode(val *string)
	DefaultSquashModeInput() *string
	DefaultSquashUid() *float64
	SetDefaultSquashUid(val *float64)
	DefaultSquashUidInput() *float64
	// Experimental.
	Fqn() *string
	InternalValue() *LustreInstanceAccessRulesOptions
	SetInternalValue(val *LustreInstanceAccessRulesOptions)
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
	PutAccessRules(value any)
	ResetAccessRules()
	ResetDefaultSquashGid()
	ResetDefaultSquashUid()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for LustreInstanceAccessRulesOptionsOutputReference
type jsiiProxy_LustreInstanceAccessRulesOptionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_LustreInstanceAccessRulesOptionsOutputReference) AccessRules() LustreInstanceAccessRulesOptionsAccessRulesList {
	var returns LustreInstanceAccessRulesOptionsAccessRulesList
	_jsii_.Get(
		j,
		"accessRules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LustreInstanceAccessRulesOptionsOutputReference) AccessRulesInput() any {
	var returns any
	_jsii_.Get(
		j,
		"accessRulesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LustreInstanceAccessRulesOptionsOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LustreInstanceAccessRulesOptionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LustreInstanceAccessRulesOptionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LustreInstanceAccessRulesOptionsOutputReference) DefaultSquashGid() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"defaultSquashGid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LustreInstanceAccessRulesOptionsOutputReference) DefaultSquashGidInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"defaultSquashGidInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LustreInstanceAccessRulesOptionsOutputReference) DefaultSquashMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultSquashMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LustreInstanceAccessRulesOptionsOutputReference) DefaultSquashModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultSquashModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LustreInstanceAccessRulesOptionsOutputReference) DefaultSquashUid() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"defaultSquashUid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LustreInstanceAccessRulesOptionsOutputReference) DefaultSquashUidInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"defaultSquashUidInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LustreInstanceAccessRulesOptionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LustreInstanceAccessRulesOptionsOutputReference) InternalValue() *LustreInstanceAccessRulesOptions {
	var returns *LustreInstanceAccessRulesOptions
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LustreInstanceAccessRulesOptionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LustreInstanceAccessRulesOptionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewLustreInstanceAccessRulesOptionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) LustreInstanceAccessRulesOptionsOutputReference {
	_init_.Initialize()

	if err := validateNewLustreInstanceAccessRulesOptionsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_LustreInstanceAccessRulesOptionsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.lustreInstance.LustreInstanceAccessRulesOptionsOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewLustreInstanceAccessRulesOptionsOutputReference_Override(l LustreInstanceAccessRulesOptionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.lustreInstance.LustreInstanceAccessRulesOptionsOutputReference",
		[]any{terraformResource, terraformAttribute},
		l,
	)
}

func (j *jsiiProxy_LustreInstanceAccessRulesOptionsOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_LustreInstanceAccessRulesOptionsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_LustreInstanceAccessRulesOptionsOutputReference) SetDefaultSquashGid(val *float64) {
	if err := j.validateSetDefaultSquashGidParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultSquashGid",
		val,
	)
}

func (j *jsiiProxy_LustreInstanceAccessRulesOptionsOutputReference) SetDefaultSquashMode(val *string) {
	if err := j.validateSetDefaultSquashModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultSquashMode",
		val,
	)
}

func (j *jsiiProxy_LustreInstanceAccessRulesOptionsOutputReference) SetDefaultSquashUid(val *float64) {
	if err := j.validateSetDefaultSquashUidParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultSquashUid",
		val,
	)
}

func (j *jsiiProxy_LustreInstanceAccessRulesOptionsOutputReference) SetInternalValue(val *LustreInstanceAccessRulesOptions) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_LustreInstanceAccessRulesOptionsOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LustreInstanceAccessRulesOptionsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (l *jsiiProxy_LustreInstanceAccessRulesOptionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LustreInstanceAccessRulesOptionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
	if err := l.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]any

	_jsii_.Invoke(
		l,
		"getAnyMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LustreInstanceAccessRulesOptionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := l.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"getBooleanAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LustreInstanceAccessRulesOptionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := l.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		l,
		"getBooleanMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LustreInstanceAccessRulesOptionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := l.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"getListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LustreInstanceAccessRulesOptionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := l.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		l,
		"getNumberAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LustreInstanceAccessRulesOptionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := l.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		l,
		"getNumberListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LustreInstanceAccessRulesOptionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := l.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		l,
		"getNumberMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LustreInstanceAccessRulesOptionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := l.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		l,
		"getStringAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LustreInstanceAccessRulesOptionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := l.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		l,
		"getStringMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LustreInstanceAccessRulesOptionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LustreInstanceAccessRulesOptionsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := l.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]any{property},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LustreInstanceAccessRulesOptionsOutputReference) PutAccessRules(value any) {
	if err := l.validatePutAccessRulesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putAccessRules",
		[]any{value},
	)
}

func (l *jsiiProxy_LustreInstanceAccessRulesOptionsOutputReference) ResetAccessRules() {
	_jsii_.InvokeVoid(
		l,
		"resetAccessRules",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LustreInstanceAccessRulesOptionsOutputReference) ResetDefaultSquashGid() {
	_jsii_.InvokeVoid(
		l,
		"resetDefaultSquashGid",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LustreInstanceAccessRulesOptionsOutputReference) ResetDefaultSquashUid() {
	_jsii_.InvokeVoid(
		l,
		"resetDefaultSquashUid",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LustreInstanceAccessRulesOptionsOutputReference) Resolve(_context cdktf.IResolveContext) any {
	if err := l.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns any

	_jsii_.Invoke(
		l,
		"resolve",
		[]any{_context},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LustreInstanceAccessRulesOptionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
