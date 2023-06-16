package teamaccess

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/tfe/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/tfe/teamaccess/internal"
)

type TeamAccessPermissionsOutputReference interface {
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
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Runs() *string
	SetRuns(val *string)
	RunsInput() *string
	RunTasks() interface{}
	SetRunTasks(val interface{})
	RunTasksInput() interface{}
	SentinelMocks() *string
	SetSentinelMocks(val *string)
	SentinelMocksInput() *string
	StateVersions() *string
	SetStateVersions(val *string)
	StateVersionsInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Variables() *string
	SetVariables(val *string)
	VariablesInput() *string
	WorkspaceLocking() interface{}
	SetWorkspaceLocking(val interface{})
	WorkspaceLockingInput() interface{}
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

// The jsii proxy struct for TeamAccessPermissionsOutputReference
type jsiiProxy_TeamAccessPermissionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_TeamAccessPermissionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamAccessPermissionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamAccessPermissionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamAccessPermissionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamAccessPermissionsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamAccessPermissionsOutputReference) Runs() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamAccessPermissionsOutputReference) RunsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamAccessPermissionsOutputReference) RunTasks() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"runTasks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamAccessPermissionsOutputReference) RunTasksInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"runTasksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamAccessPermissionsOutputReference) SentinelMocks() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sentinelMocks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamAccessPermissionsOutputReference) SentinelMocksInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sentinelMocksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamAccessPermissionsOutputReference) StateVersions() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateVersions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamAccessPermissionsOutputReference) StateVersionsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateVersionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamAccessPermissionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamAccessPermissionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamAccessPermissionsOutputReference) Variables() *string {
	var returns *string
	_jsii_.Get(
		j,
		"variables",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamAccessPermissionsOutputReference) VariablesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"variablesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamAccessPermissionsOutputReference) WorkspaceLocking() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"workspaceLocking",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamAccessPermissionsOutputReference) WorkspaceLockingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"workspaceLockingInput",
		&returns,
	)
	return returns
}


func NewTeamAccessPermissionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TeamAccessPermissionsOutputReference {
	_init_.Initialize()

	if err := validateNewTeamAccessPermissionsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TeamAccessPermissionsOutputReference{}

	_jsii_.Create(
		"tfe.teamAccess.TeamAccessPermissionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewTeamAccessPermissionsOutputReference_Override(t TeamAccessPermissionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"tfe.teamAccess.TeamAccessPermissionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TeamAccessPermissionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TeamAccessPermissionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TeamAccessPermissionsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TeamAccessPermissionsOutputReference)SetRuns(val *string) {
	if err := j.validateSetRunsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runs",
		val,
	)
}

func (j *jsiiProxy_TeamAccessPermissionsOutputReference)SetRunTasks(val interface{}) {
	if err := j.validateSetRunTasksParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runTasks",
		val,
	)
}

func (j *jsiiProxy_TeamAccessPermissionsOutputReference)SetSentinelMocks(val *string) {
	if err := j.validateSetSentinelMocksParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sentinelMocks",
		val,
	)
}

func (j *jsiiProxy_TeamAccessPermissionsOutputReference)SetStateVersions(val *string) {
	if err := j.validateSetStateVersionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stateVersions",
		val,
	)
}

func (j *jsiiProxy_TeamAccessPermissionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TeamAccessPermissionsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_TeamAccessPermissionsOutputReference)SetVariables(val *string) {
	if err := j.validateSetVariablesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"variables",
		val,
	)
}

func (j *jsiiProxy_TeamAccessPermissionsOutputReference)SetWorkspaceLocking(val interface{}) {
	if err := j.validateSetWorkspaceLockingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"workspaceLocking",
		val,
	)
}

func (t *jsiiProxy_TeamAccessPermissionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TeamAccessPermissionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TeamAccessPermissionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (t *jsiiProxy_TeamAccessPermissionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TeamAccessPermissionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TeamAccessPermissionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TeamAccessPermissionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TeamAccessPermissionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TeamAccessPermissionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TeamAccessPermissionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TeamAccessPermissionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TeamAccessPermissionsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (t *jsiiProxy_TeamAccessPermissionsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (t *jsiiProxy_TeamAccessPermissionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

