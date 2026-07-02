package teamsaccount

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/cloudflare/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/cloudflare/teamsaccount/internal"
)

type TeamsAccountLoggingSettingsByRuleTypeHttpOutputReference interface {
	cdktf.ComplexObject
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
	InternalValue() *TeamsAccountLoggingSettingsByRuleTypeHttp
	SetInternalValue(val *TeamsAccountLoggingSettingsByRuleTypeHttp)
	LogAll() any
	SetLogAll(val any)
	LogAllInput() any
	LogBlocks() any
	SetLogBlocks(val any)
	LogBlocksInput() any
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TeamsAccountLoggingSettingsByRuleTypeHttpOutputReference
type jsiiProxy_TeamsAccountLoggingSettingsByRuleTypeHttpOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_TeamsAccountLoggingSettingsByRuleTypeHttpOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamsAccountLoggingSettingsByRuleTypeHttpOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamsAccountLoggingSettingsByRuleTypeHttpOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamsAccountLoggingSettingsByRuleTypeHttpOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamsAccountLoggingSettingsByRuleTypeHttpOutputReference) InternalValue() *TeamsAccountLoggingSettingsByRuleTypeHttp {
	var returns *TeamsAccountLoggingSettingsByRuleTypeHttp
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamsAccountLoggingSettingsByRuleTypeHttpOutputReference) LogAll() any {
	var returns any
	_jsii_.Get(
		j,
		"logAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamsAccountLoggingSettingsByRuleTypeHttpOutputReference) LogAllInput() any {
	var returns any
	_jsii_.Get(
		j,
		"logAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamsAccountLoggingSettingsByRuleTypeHttpOutputReference) LogBlocks() any {
	var returns any
	_jsii_.Get(
		j,
		"logBlocks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamsAccountLoggingSettingsByRuleTypeHttpOutputReference) LogBlocksInput() any {
	var returns any
	_jsii_.Get(
		j,
		"logBlocksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamsAccountLoggingSettingsByRuleTypeHttpOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamsAccountLoggingSettingsByRuleTypeHttpOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewTeamsAccountLoggingSettingsByRuleTypeHttpOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) TeamsAccountLoggingSettingsByRuleTypeHttpOutputReference {
	_init_.Initialize()

	if err := validateNewTeamsAccountLoggingSettingsByRuleTypeHttpOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TeamsAccountLoggingSettingsByRuleTypeHttpOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.teamsAccount.TeamsAccountLoggingSettingsByRuleTypeHttpOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewTeamsAccountLoggingSettingsByRuleTypeHttpOutputReference_Override(t TeamsAccountLoggingSettingsByRuleTypeHttpOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.teamsAccount.TeamsAccountLoggingSettingsByRuleTypeHttpOutputReference",
		[]any{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TeamsAccountLoggingSettingsByRuleTypeHttpOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TeamsAccountLoggingSettingsByRuleTypeHttpOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TeamsAccountLoggingSettingsByRuleTypeHttpOutputReference) SetInternalValue(val *TeamsAccountLoggingSettingsByRuleTypeHttp) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TeamsAccountLoggingSettingsByRuleTypeHttpOutputReference) SetLogAll(val any) {
	if err := j.validateSetLogAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logAll",
		val,
	)
}

func (j *jsiiProxy_TeamsAccountLoggingSettingsByRuleTypeHttpOutputReference) SetLogBlocks(val any) {
	if err := j.validateSetLogBlocksParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logBlocks",
		val,
	)
}

func (j *jsiiProxy_TeamsAccountLoggingSettingsByRuleTypeHttpOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TeamsAccountLoggingSettingsByRuleTypeHttpOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TeamsAccountLoggingSettingsByRuleTypeHttpOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TeamsAccountLoggingSettingsByRuleTypeHttpOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
	if err := t.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]any

	_jsii_.Invoke(
		t,
		"getAnyMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TeamsAccountLoggingSettingsByRuleTypeHttpOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := t.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		t,
		"getBooleanAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TeamsAccountLoggingSettingsByRuleTypeHttpOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := t.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		t,
		"getBooleanMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TeamsAccountLoggingSettingsByRuleTypeHttpOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := t.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		t,
		"getListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TeamsAccountLoggingSettingsByRuleTypeHttpOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := t.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		t,
		"getNumberAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TeamsAccountLoggingSettingsByRuleTypeHttpOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := t.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		t,
		"getNumberListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TeamsAccountLoggingSettingsByRuleTypeHttpOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := t.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		t,
		"getNumberMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TeamsAccountLoggingSettingsByRuleTypeHttpOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := t.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		t,
		"getStringAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TeamsAccountLoggingSettingsByRuleTypeHttpOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := t.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		t,
		"getStringMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TeamsAccountLoggingSettingsByRuleTypeHttpOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TeamsAccountLoggingSettingsByRuleTypeHttpOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := t.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationForAttribute",
		[]any{property},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TeamsAccountLoggingSettingsByRuleTypeHttpOutputReference) Resolve(_context cdktf.IResolveContext) any {
	if err := t.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns any

	_jsii_.Invoke(
		t,
		"resolve",
		[]any{_context},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TeamsAccountLoggingSettingsByRuleTypeHttpOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
