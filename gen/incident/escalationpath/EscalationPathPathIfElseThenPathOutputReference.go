package escalationpath

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/incident/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/incident/escalationpath/internal"
)

type EscalationPathPathIfElseThenPathOutputReference interface {
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
	// Experimental.
	Fqn() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	IfElse() EscalationPathPathIfElseOutputReference
	IfElseInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Level() EscalationPathPathIfElseThenPathLevelOutputReference
	LevelInput() interface{}
	NotifyChannel() EscalationPathPathIfElseThenPathNotifyChannelOutputReference
	NotifyChannelInput() interface{}
	Repeat() EscalationPathPathIfElseThenPathRepeatOutputReference
	RepeatInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Type() *string
	SetType(val *string)
	TypeInput() *string
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
	PutIfElse(value *EscalationPathPathIfElse)
	PutLevel(value *EscalationPathPathIfElseThenPathLevel)
	PutNotifyChannel(value *EscalationPathPathIfElseThenPathNotifyChannel)
	PutRepeat(value *EscalationPathPathIfElseThenPathRepeat)
	ResetId()
	ResetIfElse()
	ResetLevel()
	ResetNotifyChannel()
	ResetRepeat()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EscalationPathPathIfElseThenPathOutputReference
type jsiiProxy_EscalationPathPathIfElseThenPathOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_EscalationPathPathIfElseThenPathOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EscalationPathPathIfElseThenPathOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EscalationPathPathIfElseThenPathOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EscalationPathPathIfElseThenPathOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EscalationPathPathIfElseThenPathOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EscalationPathPathIfElseThenPathOutputReference) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EscalationPathPathIfElseThenPathOutputReference) IfElse() EscalationPathPathIfElseOutputReference {
	var returns EscalationPathPathIfElseOutputReference
	_jsii_.Get(
		j,
		"ifElse",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EscalationPathPathIfElseThenPathOutputReference) IfElseInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ifElseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EscalationPathPathIfElseThenPathOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EscalationPathPathIfElseThenPathOutputReference) Level() EscalationPathPathIfElseThenPathLevelOutputReference {
	var returns EscalationPathPathIfElseThenPathLevelOutputReference
	_jsii_.Get(
		j,
		"level",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EscalationPathPathIfElseThenPathOutputReference) LevelInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"levelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EscalationPathPathIfElseThenPathOutputReference) NotifyChannel() EscalationPathPathIfElseThenPathNotifyChannelOutputReference {
	var returns EscalationPathPathIfElseThenPathNotifyChannelOutputReference
	_jsii_.Get(
		j,
		"notifyChannel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EscalationPathPathIfElseThenPathOutputReference) NotifyChannelInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"notifyChannelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EscalationPathPathIfElseThenPathOutputReference) Repeat() EscalationPathPathIfElseThenPathRepeatOutputReference {
	var returns EscalationPathPathIfElseThenPathRepeatOutputReference
	_jsii_.Get(
		j,
		"repeat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EscalationPathPathIfElseThenPathOutputReference) RepeatInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"repeatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EscalationPathPathIfElseThenPathOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EscalationPathPathIfElseThenPathOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EscalationPathPathIfElseThenPathOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EscalationPathPathIfElseThenPathOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


func NewEscalationPathPathIfElseThenPathOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) EscalationPathPathIfElseThenPathOutputReference {
	_init_.Initialize()

	if err := validateNewEscalationPathPathIfElseThenPathOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_EscalationPathPathIfElseThenPathOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-incident.escalationPath.EscalationPathPathIfElseThenPathOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewEscalationPathPathIfElseThenPathOutputReference_Override(e EscalationPathPathIfElseThenPathOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-incident.escalationPath.EscalationPathPathIfElseThenPathOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		e,
	)
}

func (j *jsiiProxy_EscalationPathPathIfElseThenPathOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_EscalationPathPathIfElseThenPathOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_EscalationPathPathIfElseThenPathOutputReference)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_EscalationPathPathIfElseThenPathOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EscalationPathPathIfElseThenPathOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EscalationPathPathIfElseThenPathOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_EscalationPathPathIfElseThenPathOutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (e *jsiiProxy_EscalationPathPathIfElseThenPathOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EscalationPathPathIfElseThenPathOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := e.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EscalationPathPathIfElseThenPathOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := e.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EscalationPathPathIfElseThenPathOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := e.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EscalationPathPathIfElseThenPathOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := e.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EscalationPathPathIfElseThenPathOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := e.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EscalationPathPathIfElseThenPathOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := e.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EscalationPathPathIfElseThenPathOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := e.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EscalationPathPathIfElseThenPathOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := e.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EscalationPathPathIfElseThenPathOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := e.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EscalationPathPathIfElseThenPathOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EscalationPathPathIfElseThenPathOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := e.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EscalationPathPathIfElseThenPathOutputReference) PutIfElse(value *EscalationPathPathIfElse) {
	if err := e.validatePutIfElseParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putIfElse",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EscalationPathPathIfElseThenPathOutputReference) PutLevel(value *EscalationPathPathIfElseThenPathLevel) {
	if err := e.validatePutLevelParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putLevel",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EscalationPathPathIfElseThenPathOutputReference) PutNotifyChannel(value *EscalationPathPathIfElseThenPathNotifyChannel) {
	if err := e.validatePutNotifyChannelParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putNotifyChannel",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EscalationPathPathIfElseThenPathOutputReference) PutRepeat(value *EscalationPathPathIfElseThenPathRepeat) {
	if err := e.validatePutRepeatParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putRepeat",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EscalationPathPathIfElseThenPathOutputReference) ResetId() {
	_jsii_.InvokeVoid(
		e,
		"resetId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EscalationPathPathIfElseThenPathOutputReference) ResetIfElse() {
	_jsii_.InvokeVoid(
		e,
		"resetIfElse",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EscalationPathPathIfElseThenPathOutputReference) ResetLevel() {
	_jsii_.InvokeVoid(
		e,
		"resetLevel",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EscalationPathPathIfElseThenPathOutputReference) ResetNotifyChannel() {
	_jsii_.InvokeVoid(
		e,
		"resetNotifyChannel",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EscalationPathPathIfElseThenPathOutputReference) ResetRepeat() {
	_jsii_.InvokeVoid(
		e,
		"resetRepeat",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EscalationPathPathIfElseThenPathOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := e.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EscalationPathPathIfElseThenPathOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

