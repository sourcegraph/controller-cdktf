package teamsaccount

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/cloudflare/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/cloudflare/teamsaccount/internal"
)

type TeamsAccountAntivirusOutputReference interface {
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
	EnabledDownloadPhase() interface{}
	SetEnabledDownloadPhase(val interface{})
	EnabledDownloadPhaseInput() interface{}
	EnabledUploadPhase() interface{}
	SetEnabledUploadPhase(val interface{})
	EnabledUploadPhaseInput() interface{}
	FailClosed() interface{}
	SetFailClosed(val interface{})
	FailClosedInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *TeamsAccountAntivirus
	SetInternalValue(val *TeamsAccountAntivirus)
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TeamsAccountAntivirusOutputReference
type jsiiProxy_TeamsAccountAntivirusOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_TeamsAccountAntivirusOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamsAccountAntivirusOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamsAccountAntivirusOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamsAccountAntivirusOutputReference) EnabledDownloadPhase() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledDownloadPhase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamsAccountAntivirusOutputReference) EnabledDownloadPhaseInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledDownloadPhaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamsAccountAntivirusOutputReference) EnabledUploadPhase() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledUploadPhase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamsAccountAntivirusOutputReference) EnabledUploadPhaseInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledUploadPhaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamsAccountAntivirusOutputReference) FailClosed() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"failClosed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamsAccountAntivirusOutputReference) FailClosedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"failClosedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamsAccountAntivirusOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamsAccountAntivirusOutputReference) InternalValue() *TeamsAccountAntivirus {
	var returns *TeamsAccountAntivirus
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamsAccountAntivirusOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamsAccountAntivirusOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewTeamsAccountAntivirusOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) TeamsAccountAntivirusOutputReference {
	_init_.Initialize()

	if err := validateNewTeamsAccountAntivirusOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TeamsAccountAntivirusOutputReference{}

	_jsii_.Create(
		"cloudflare.teamsAccount.TeamsAccountAntivirusOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewTeamsAccountAntivirusOutputReference_Override(t TeamsAccountAntivirusOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"cloudflare.teamsAccount.TeamsAccountAntivirusOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TeamsAccountAntivirusOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TeamsAccountAntivirusOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TeamsAccountAntivirusOutputReference)SetEnabledDownloadPhase(val interface{}) {
	if err := j.validateSetEnabledDownloadPhaseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabledDownloadPhase",
		val,
	)
}

func (j *jsiiProxy_TeamsAccountAntivirusOutputReference)SetEnabledUploadPhase(val interface{}) {
	if err := j.validateSetEnabledUploadPhaseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabledUploadPhase",
		val,
	)
}

func (j *jsiiProxy_TeamsAccountAntivirusOutputReference)SetFailClosed(val interface{}) {
	if err := j.validateSetFailClosedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"failClosed",
		val,
	)
}

func (j *jsiiProxy_TeamsAccountAntivirusOutputReference)SetInternalValue(val *TeamsAccountAntivirus) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TeamsAccountAntivirusOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TeamsAccountAntivirusOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TeamsAccountAntivirusOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TeamsAccountAntivirusOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := t.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TeamsAccountAntivirusOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := t.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		t,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TeamsAccountAntivirusOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := t.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		t,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TeamsAccountAntivirusOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := t.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		t,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TeamsAccountAntivirusOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := t.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		t,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TeamsAccountAntivirusOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := t.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		t,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TeamsAccountAntivirusOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := t.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		t,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TeamsAccountAntivirusOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := t.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		t,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TeamsAccountAntivirusOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := t.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		t,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TeamsAccountAntivirusOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TeamsAccountAntivirusOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := t.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TeamsAccountAntivirusOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := t.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		t,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TeamsAccountAntivirusOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

