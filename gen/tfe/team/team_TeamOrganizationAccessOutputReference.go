package team

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/tfe/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/tfe/team/internal"
)

type TeamOrganizationAccessOutputReference interface {
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
	InternalValue() *TeamOrganizationAccess
	SetInternalValue(val *TeamOrganizationAccess)
	ManageModules() interface{}
	SetManageModules(val interface{})
	ManageModulesInput() interface{}
	ManagePolicies() interface{}
	SetManagePolicies(val interface{})
	ManagePoliciesInput() interface{}
	ManagePolicyOverrides() interface{}
	SetManagePolicyOverrides(val interface{})
	ManagePolicyOverridesInput() interface{}
	ManageProjects() interface{}
	SetManageProjects(val interface{})
	ManageProjectsInput() interface{}
	ManageProviders() interface{}
	SetManageProviders(val interface{})
	ManageProvidersInput() interface{}
	ManageRunTasks() interface{}
	SetManageRunTasks(val interface{})
	ManageRunTasksInput() interface{}
	ManageVcsSettings() interface{}
	SetManageVcsSettings(val interface{})
	ManageVcsSettingsInput() interface{}
	ManageWorkspaces() interface{}
	SetManageWorkspaces(val interface{})
	ManageWorkspacesInput() interface{}
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
	ResetManageModules()
	ResetManagePolicies()
	ResetManagePolicyOverrides()
	ResetManageProjects()
	ResetManageProviders()
	ResetManageRunTasks()
	ResetManageVcsSettings()
	ResetManageWorkspaces()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TeamOrganizationAccessOutputReference
type jsiiProxy_TeamOrganizationAccessOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_TeamOrganizationAccessOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamOrganizationAccessOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamOrganizationAccessOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamOrganizationAccessOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamOrganizationAccessOutputReference) InternalValue() *TeamOrganizationAccess {
	var returns *TeamOrganizationAccess
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamOrganizationAccessOutputReference) ManageModules() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"manageModules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamOrganizationAccessOutputReference) ManageModulesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"manageModulesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamOrganizationAccessOutputReference) ManagePolicies() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"managePolicies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamOrganizationAccessOutputReference) ManagePoliciesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"managePoliciesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamOrganizationAccessOutputReference) ManagePolicyOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"managePolicyOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamOrganizationAccessOutputReference) ManagePolicyOverridesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"managePolicyOverridesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamOrganizationAccessOutputReference) ManageProjects() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"manageProjects",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamOrganizationAccessOutputReference) ManageProjectsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"manageProjectsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamOrganizationAccessOutputReference) ManageProviders() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"manageProviders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamOrganizationAccessOutputReference) ManageProvidersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"manageProvidersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamOrganizationAccessOutputReference) ManageRunTasks() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"manageRunTasks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamOrganizationAccessOutputReference) ManageRunTasksInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"manageRunTasksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamOrganizationAccessOutputReference) ManageVcsSettings() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"manageVcsSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamOrganizationAccessOutputReference) ManageVcsSettingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"manageVcsSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamOrganizationAccessOutputReference) ManageWorkspaces() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"manageWorkspaces",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamOrganizationAccessOutputReference) ManageWorkspacesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"manageWorkspacesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamOrganizationAccessOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamOrganizationAccessOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewTeamOrganizationAccessOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) TeamOrganizationAccessOutputReference {
	_init_.Initialize()

	if err := validateNewTeamOrganizationAccessOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TeamOrganizationAccessOutputReference{}

	_jsii_.Create(
		"tfe.team.TeamOrganizationAccessOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewTeamOrganizationAccessOutputReference_Override(t TeamOrganizationAccessOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"tfe.team.TeamOrganizationAccessOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TeamOrganizationAccessOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TeamOrganizationAccessOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TeamOrganizationAccessOutputReference)SetInternalValue(val *TeamOrganizationAccess) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TeamOrganizationAccessOutputReference)SetManageModules(val interface{}) {
	if err := j.validateSetManageModulesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"manageModules",
		val,
	)
}

func (j *jsiiProxy_TeamOrganizationAccessOutputReference)SetManagePolicies(val interface{}) {
	if err := j.validateSetManagePoliciesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"managePolicies",
		val,
	)
}

func (j *jsiiProxy_TeamOrganizationAccessOutputReference)SetManagePolicyOverrides(val interface{}) {
	if err := j.validateSetManagePolicyOverridesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"managePolicyOverrides",
		val,
	)
}

func (j *jsiiProxy_TeamOrganizationAccessOutputReference)SetManageProjects(val interface{}) {
	if err := j.validateSetManageProjectsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"manageProjects",
		val,
	)
}

func (j *jsiiProxy_TeamOrganizationAccessOutputReference)SetManageProviders(val interface{}) {
	if err := j.validateSetManageProvidersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"manageProviders",
		val,
	)
}

func (j *jsiiProxy_TeamOrganizationAccessOutputReference)SetManageRunTasks(val interface{}) {
	if err := j.validateSetManageRunTasksParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"manageRunTasks",
		val,
	)
}

func (j *jsiiProxy_TeamOrganizationAccessOutputReference)SetManageVcsSettings(val interface{}) {
	if err := j.validateSetManageVcsSettingsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"manageVcsSettings",
		val,
	)
}

func (j *jsiiProxy_TeamOrganizationAccessOutputReference)SetManageWorkspaces(val interface{}) {
	if err := j.validateSetManageWorkspacesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"manageWorkspaces",
		val,
	)
}

func (j *jsiiProxy_TeamOrganizationAccessOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TeamOrganizationAccessOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TeamOrganizationAccessOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TeamOrganizationAccessOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TeamOrganizationAccessOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (t *jsiiProxy_TeamOrganizationAccessOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TeamOrganizationAccessOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TeamOrganizationAccessOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TeamOrganizationAccessOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TeamOrganizationAccessOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TeamOrganizationAccessOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TeamOrganizationAccessOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TeamOrganizationAccessOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TeamOrganizationAccessOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (t *jsiiProxy_TeamOrganizationAccessOutputReference) ResetManageModules() {
	_jsii_.InvokeVoid(
		t,
		"resetManageModules",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TeamOrganizationAccessOutputReference) ResetManagePolicies() {
	_jsii_.InvokeVoid(
		t,
		"resetManagePolicies",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TeamOrganizationAccessOutputReference) ResetManagePolicyOverrides() {
	_jsii_.InvokeVoid(
		t,
		"resetManagePolicyOverrides",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TeamOrganizationAccessOutputReference) ResetManageProjects() {
	_jsii_.InvokeVoid(
		t,
		"resetManageProjects",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TeamOrganizationAccessOutputReference) ResetManageProviders() {
	_jsii_.InvokeVoid(
		t,
		"resetManageProviders",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TeamOrganizationAccessOutputReference) ResetManageRunTasks() {
	_jsii_.InvokeVoid(
		t,
		"resetManageRunTasks",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TeamOrganizationAccessOutputReference) ResetManageVcsSettings() {
	_jsii_.InvokeVoid(
		t,
		"resetManageVcsSettings",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TeamOrganizationAccessOutputReference) ResetManageWorkspaces() {
	_jsii_.InvokeVoid(
		t,
		"resetManageWorkspaces",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TeamOrganizationAccessOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (t *jsiiProxy_TeamOrganizationAccessOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

