package googledialogflowcxagent

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googledialogflowcxagent/internal"
)

type GoogleDialogflowCxAgentGitIntegrationSettingsGithubSettingsOutputReference interface {
	cdktf.ComplexObject
	AccessToken() *string
	SetAccessToken(val *string)
	AccessTokenInput() *string
	Branches() *[]*string
	SetBranches(val *[]*string)
	BranchesInput() *[]*string
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
	DisplayName() *string
	SetDisplayName(val *string)
	DisplayNameInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleDialogflowCxAgentGitIntegrationSettingsGithubSettings
	SetInternalValue(val *GoogleDialogflowCxAgentGitIntegrationSettingsGithubSettings)
	RepositoryUri() *string
	SetRepositoryUri(val *string)
	RepositoryUriInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TrackingBranch() *string
	SetTrackingBranch(val *string)
	TrackingBranchInput() *string
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
	ResetAccessToken()
	ResetBranches()
	ResetDisplayName()
	ResetRepositoryUri()
	ResetTrackingBranch()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleDialogflowCxAgentGitIntegrationSettingsGithubSettingsOutputReference
type jsiiProxy_GoogleDialogflowCxAgentGitIntegrationSettingsGithubSettingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleDialogflowCxAgentGitIntegrationSettingsGithubSettingsOutputReference) AccessToken() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxAgentGitIntegrationSettingsGithubSettingsOutputReference) AccessTokenInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxAgentGitIntegrationSettingsGithubSettingsOutputReference) Branches() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"branches",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxAgentGitIntegrationSettingsGithubSettingsOutputReference) BranchesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"branchesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxAgentGitIntegrationSettingsGithubSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxAgentGitIntegrationSettingsGithubSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxAgentGitIntegrationSettingsGithubSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxAgentGitIntegrationSettingsGithubSettingsOutputReference) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxAgentGitIntegrationSettingsGithubSettingsOutputReference) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxAgentGitIntegrationSettingsGithubSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxAgentGitIntegrationSettingsGithubSettingsOutputReference) InternalValue() *GoogleDialogflowCxAgentGitIntegrationSettingsGithubSettings {
	var returns *GoogleDialogflowCxAgentGitIntegrationSettingsGithubSettings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxAgentGitIntegrationSettingsGithubSettingsOutputReference) RepositoryUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repositoryUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxAgentGitIntegrationSettingsGithubSettingsOutputReference) RepositoryUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repositoryUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxAgentGitIntegrationSettingsGithubSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxAgentGitIntegrationSettingsGithubSettingsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxAgentGitIntegrationSettingsGithubSettingsOutputReference) TrackingBranch() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trackingBranch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxAgentGitIntegrationSettingsGithubSettingsOutputReference) TrackingBranchInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trackingBranchInput",
		&returns,
	)
	return returns
}


func NewGoogleDialogflowCxAgentGitIntegrationSettingsGithubSettingsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleDialogflowCxAgentGitIntegrationSettingsGithubSettingsOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleDialogflowCxAgentGitIntegrationSettingsGithubSettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleDialogflowCxAgentGitIntegrationSettingsGithubSettingsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleDialogflowCxAgent.GoogleDialogflowCxAgentGitIntegrationSettingsGithubSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleDialogflowCxAgentGitIntegrationSettingsGithubSettingsOutputReference_Override(g GoogleDialogflowCxAgentGitIntegrationSettingsGithubSettingsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleDialogflowCxAgent.GoogleDialogflowCxAgentGitIntegrationSettingsGithubSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxAgentGitIntegrationSettingsGithubSettingsOutputReference)SetAccessToken(val *string) {
	if err := j.validateSetAccessTokenParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accessToken",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxAgentGitIntegrationSettingsGithubSettingsOutputReference)SetBranches(val *[]*string) {
	if err := j.validateSetBranchesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"branches",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxAgentGitIntegrationSettingsGithubSettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxAgentGitIntegrationSettingsGithubSettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxAgentGitIntegrationSettingsGithubSettingsOutputReference)SetDisplayName(val *string) {
	if err := j.validateSetDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxAgentGitIntegrationSettingsGithubSettingsOutputReference)SetInternalValue(val *GoogleDialogflowCxAgentGitIntegrationSettingsGithubSettings) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxAgentGitIntegrationSettingsGithubSettingsOutputReference)SetRepositoryUri(val *string) {
	if err := j.validateSetRepositoryUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"repositoryUri",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxAgentGitIntegrationSettingsGithubSettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxAgentGitIntegrationSettingsGithubSettingsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxAgentGitIntegrationSettingsGithubSettingsOutputReference)SetTrackingBranch(val *string) {
	if err := j.validateSetTrackingBranchParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"trackingBranch",
		val,
	)
}

func (g *jsiiProxy_GoogleDialogflowCxAgentGitIntegrationSettingsGithubSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDialogflowCxAgentGitIntegrationSettingsGithubSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleDialogflowCxAgentGitIntegrationSettingsGithubSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDialogflowCxAgentGitIntegrationSettingsGithubSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleDialogflowCxAgentGitIntegrationSettingsGithubSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleDialogflowCxAgentGitIntegrationSettingsGithubSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleDialogflowCxAgentGitIntegrationSettingsGithubSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleDialogflowCxAgentGitIntegrationSettingsGithubSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleDialogflowCxAgentGitIntegrationSettingsGithubSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleDialogflowCxAgentGitIntegrationSettingsGithubSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleDialogflowCxAgentGitIntegrationSettingsGithubSettingsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDialogflowCxAgentGitIntegrationSettingsGithubSettingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDialogflowCxAgentGitIntegrationSettingsGithubSettingsOutputReference) ResetAccessToken() {
	_jsii_.InvokeVoid(
		g,
		"resetAccessToken",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowCxAgentGitIntegrationSettingsGithubSettingsOutputReference) ResetBranches() {
	_jsii_.InvokeVoid(
		g,
		"resetBranches",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowCxAgentGitIntegrationSettingsGithubSettingsOutputReference) ResetDisplayName() {
	_jsii_.InvokeVoid(
		g,
		"resetDisplayName",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowCxAgentGitIntegrationSettingsGithubSettingsOutputReference) ResetRepositoryUri() {
	_jsii_.InvokeVoid(
		g,
		"resetRepositoryUri",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowCxAgentGitIntegrationSettingsGithubSettingsOutputReference) ResetTrackingBranch() {
	_jsii_.InvokeVoid(
		g,
		"resetTrackingBranch",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowCxAgentGitIntegrationSettingsGithubSettingsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleDialogflowCxAgentGitIntegrationSettingsGithubSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

