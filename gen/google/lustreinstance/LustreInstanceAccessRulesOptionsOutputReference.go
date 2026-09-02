package lustreinstance

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google/lustreinstance/internal"
)

type LustreInstanceAccessRulesOptionsOutputReference interface {
	cdktn.ComplexObject
	AccessRules() LustreInstanceAccessRulesOptionsAccessRulesList
	AccessRulesInput() interface{}
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
	PutAccessRules(value interface{})
	ResetAccessRules()
	ResetDefaultSquashGid()
	ResetDefaultSquashUid()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for LustreInstanceAccessRulesOptionsOutputReference
type jsiiProxy_LustreInstanceAccessRulesOptionsOutputReference struct {
	internal.Type__cdktnComplexObject
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

func (j *jsiiProxy_LustreInstanceAccessRulesOptionsOutputReference) AccessRulesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"accessRulesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LustreInstanceAccessRulesOptionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
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

func (j *jsiiProxy_LustreInstanceAccessRulesOptionsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewLustreInstanceAccessRulesOptionsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) LustreInstanceAccessRulesOptionsOutputReference {
	_init_.Initialize()

	if err := validateNewLustreInstanceAccessRulesOptionsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_LustreInstanceAccessRulesOptionsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.lustreInstance.LustreInstanceAccessRulesOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewLustreInstanceAccessRulesOptionsOutputReference_Override(l LustreInstanceAccessRulesOptionsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.lustreInstance.LustreInstanceAccessRulesOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		l,
	)
}

func (j *jsiiProxy_LustreInstanceAccessRulesOptionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_LustreInstanceAccessRulesOptionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_LustreInstanceAccessRulesOptionsOutputReference)SetDefaultSquashGid(val *float64) {
	if err := j.validateSetDefaultSquashGidParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultSquashGid",
		val,
	)
}

func (j *jsiiProxy_LustreInstanceAccessRulesOptionsOutputReference)SetDefaultSquashMode(val *string) {
	if err := j.validateSetDefaultSquashModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultSquashMode",
		val,
	)
}

func (j *jsiiProxy_LustreInstanceAccessRulesOptionsOutputReference)SetDefaultSquashUid(val *float64) {
	if err := j.validateSetDefaultSquashUidParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultSquashUid",
		val,
	)
}

func (j *jsiiProxy_LustreInstanceAccessRulesOptionsOutputReference)SetInternalValue(val *LustreInstanceAccessRulesOptions) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_LustreInstanceAccessRulesOptionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LustreInstanceAccessRulesOptionsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
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

func (l *jsiiProxy_LustreInstanceAccessRulesOptionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := l.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LustreInstanceAccessRulesOptionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := l.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		l,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
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
		[]interface{}{terraformAttribute},
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
		[]interface{}{terraformAttribute},
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
		[]interface{}{terraformAttribute},
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
		[]interface{}{terraformAttribute},
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
		[]interface{}{terraformAttribute},
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
		[]interface{}{terraformAttribute},
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
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LustreInstanceAccessRulesOptionsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LustreInstanceAccessRulesOptionsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := l.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LustreInstanceAccessRulesOptionsOutputReference) PutAccessRules(value interface{}) {
	if err := l.validatePutAccessRulesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putAccessRules",
		[]interface{}{value},
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

func (l *jsiiProxy_LustreInstanceAccessRulesOptionsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := l.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		l,
		"resolve",
		[]interface{}{context},
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

