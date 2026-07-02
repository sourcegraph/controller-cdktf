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
	InternalValue() *TeamOrganizationAccess
	SetInternalValue(val *TeamOrganizationAccess)
	ManageModules() any
	SetManageModules(val any)
	ManageModulesInput() any
	ManagePolicies() any
	SetManagePolicies(val any)
	ManagePoliciesInput() any
	ManagePolicyOverrides() any
	SetManagePolicyOverrides(val any)
	ManagePolicyOverridesInput() any
	ManageProjects() any
	SetManageProjects(val any)
	ManageProjectsInput() any
	ManageProviders() any
	SetManageProviders(val any)
	ManageProvidersInput() any
	ManageRunTasks() any
	SetManageRunTasks(val any)
	ManageRunTasksInput() any
	ManageVcsSettings() any
	SetManageVcsSettings(val any)
	ManageVcsSettingsInput() any
	ManageWorkspaces() any
	SetManageWorkspaces(val any)
	ManageWorkspacesInput() any
	ReadProjects() any
	SetReadProjects(val any)
	ReadProjectsInput() any
	ReadWorkspaces() any
	SetReadWorkspaces(val any)
	ReadWorkspacesInput() any
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
	ResetManageModules()
	ResetManagePolicies()
	ResetManagePolicyOverrides()
	ResetManageProjects()
	ResetManageProviders()
	ResetManageRunTasks()
	ResetManageVcsSettings()
	ResetManageWorkspaces()
	ResetReadProjects()
	ResetReadWorkspaces()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
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

func (j *jsiiProxy_TeamOrganizationAccessOutputReference) ComplexObjectIndex() any {
	var returns any
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

func (j *jsiiProxy_TeamOrganizationAccessOutputReference) ManageModules() any {
	var returns any
	_jsii_.Get(
		j,
		"manageModules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamOrganizationAccessOutputReference) ManageModulesInput() any {
	var returns any
	_jsii_.Get(
		j,
		"manageModulesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamOrganizationAccessOutputReference) ManagePolicies() any {
	var returns any
	_jsii_.Get(
		j,
		"managePolicies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamOrganizationAccessOutputReference) ManagePoliciesInput() any {
	var returns any
	_jsii_.Get(
		j,
		"managePoliciesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamOrganizationAccessOutputReference) ManagePolicyOverrides() any {
	var returns any
	_jsii_.Get(
		j,
		"managePolicyOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamOrganizationAccessOutputReference) ManagePolicyOverridesInput() any {
	var returns any
	_jsii_.Get(
		j,
		"managePolicyOverridesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamOrganizationAccessOutputReference) ManageProjects() any {
	var returns any
	_jsii_.Get(
		j,
		"manageProjects",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamOrganizationAccessOutputReference) ManageProjectsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"manageProjectsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamOrganizationAccessOutputReference) ManageProviders() any {
	var returns any
	_jsii_.Get(
		j,
		"manageProviders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamOrganizationAccessOutputReference) ManageProvidersInput() any {
	var returns any
	_jsii_.Get(
		j,
		"manageProvidersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamOrganizationAccessOutputReference) ManageRunTasks() any {
	var returns any
	_jsii_.Get(
		j,
		"manageRunTasks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamOrganizationAccessOutputReference) ManageRunTasksInput() any {
	var returns any
	_jsii_.Get(
		j,
		"manageRunTasksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamOrganizationAccessOutputReference) ManageVcsSettings() any {
	var returns any
	_jsii_.Get(
		j,
		"manageVcsSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamOrganizationAccessOutputReference) ManageVcsSettingsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"manageVcsSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamOrganizationAccessOutputReference) ManageWorkspaces() any {
	var returns any
	_jsii_.Get(
		j,
		"manageWorkspaces",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamOrganizationAccessOutputReference) ManageWorkspacesInput() any {
	var returns any
	_jsii_.Get(
		j,
		"manageWorkspacesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamOrganizationAccessOutputReference) ReadProjects() any {
	var returns any
	_jsii_.Get(
		j,
		"readProjects",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamOrganizationAccessOutputReference) ReadProjectsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"readProjectsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamOrganizationAccessOutputReference) ReadWorkspaces() any {
	var returns any
	_jsii_.Get(
		j,
		"readWorkspaces",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamOrganizationAccessOutputReference) ReadWorkspacesInput() any {
	var returns any
	_jsii_.Get(
		j,
		"readWorkspacesInput",
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
		"@cdktf/provider-tfe.team.TeamOrganizationAccessOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewTeamOrganizationAccessOutputReference_Override(t TeamOrganizationAccessOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-tfe.team.TeamOrganizationAccessOutputReference",
		[]any{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TeamOrganizationAccessOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TeamOrganizationAccessOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TeamOrganizationAccessOutputReference) SetInternalValue(val *TeamOrganizationAccess) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TeamOrganizationAccessOutputReference) SetManageModules(val any) {
	if err := j.validateSetManageModulesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"manageModules",
		val,
	)
}

func (j *jsiiProxy_TeamOrganizationAccessOutputReference) SetManagePolicies(val any) {
	if err := j.validateSetManagePoliciesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"managePolicies",
		val,
	)
}

func (j *jsiiProxy_TeamOrganizationAccessOutputReference) SetManagePolicyOverrides(val any) {
	if err := j.validateSetManagePolicyOverridesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"managePolicyOverrides",
		val,
	)
}

func (j *jsiiProxy_TeamOrganizationAccessOutputReference) SetManageProjects(val any) {
	if err := j.validateSetManageProjectsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"manageProjects",
		val,
	)
}

func (j *jsiiProxy_TeamOrganizationAccessOutputReference) SetManageProviders(val any) {
	if err := j.validateSetManageProvidersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"manageProviders",
		val,
	)
}

func (j *jsiiProxy_TeamOrganizationAccessOutputReference) SetManageRunTasks(val any) {
	if err := j.validateSetManageRunTasksParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"manageRunTasks",
		val,
	)
}

func (j *jsiiProxy_TeamOrganizationAccessOutputReference) SetManageVcsSettings(val any) {
	if err := j.validateSetManageVcsSettingsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"manageVcsSettings",
		val,
	)
}

func (j *jsiiProxy_TeamOrganizationAccessOutputReference) SetManageWorkspaces(val any) {
	if err := j.validateSetManageWorkspacesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"manageWorkspaces",
		val,
	)
}

func (j *jsiiProxy_TeamOrganizationAccessOutputReference) SetReadProjects(val any) {
	if err := j.validateSetReadProjectsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"readProjects",
		val,
	)
}

func (j *jsiiProxy_TeamOrganizationAccessOutputReference) SetReadWorkspaces(val any) {
	if err := j.validateSetReadWorkspacesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"readWorkspaces",
		val,
	)
}

func (j *jsiiProxy_TeamOrganizationAccessOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TeamOrganizationAccessOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
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

func (t *jsiiProxy_TeamOrganizationAccessOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
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

func (t *jsiiProxy_TeamOrganizationAccessOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (t *jsiiProxy_TeamOrganizationAccessOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TeamOrganizationAccessOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TeamOrganizationAccessOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TeamOrganizationAccessOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TeamOrganizationAccessOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TeamOrganizationAccessOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TeamOrganizationAccessOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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
		[]any{property},
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

func (t *jsiiProxy_TeamOrganizationAccessOutputReference) ResetReadProjects() {
	_jsii_.InvokeVoid(
		t,
		"resetReadProjects",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TeamOrganizationAccessOutputReference) ResetReadWorkspaces() {
	_jsii_.InvokeVoid(
		t,
		"resetReadWorkspaces",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TeamOrganizationAccessOutputReference) Resolve(_context cdktf.IResolveContext) any {
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
