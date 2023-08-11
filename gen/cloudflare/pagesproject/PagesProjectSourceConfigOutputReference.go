package pagesproject

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/cloudflare/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/cloudflare/pagesproject/internal"
)

type PagesProjectSourceConfigOutputReference interface {
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
	DeploymentsEnabled() interface{}
	SetDeploymentsEnabled(val interface{})
	DeploymentsEnabledInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *PagesProjectSourceConfig
	SetInternalValue(val *PagesProjectSourceConfig)
	Owner() *string
	SetOwner(val *string)
	OwnerInput() *string
	PrCommentsEnabled() interface{}
	SetPrCommentsEnabled(val interface{})
	PrCommentsEnabledInput() interface{}
	PreviewBranchExcludes() *[]*string
	SetPreviewBranchExcludes(val *[]*string)
	PreviewBranchExcludesInput() *[]*string
	PreviewBranchIncludes() *[]*string
	SetPreviewBranchIncludes(val *[]*string)
	PreviewBranchIncludesInput() *[]*string
	PreviewDeploymentSetting() *string
	SetPreviewDeploymentSetting(val *string)
	PreviewDeploymentSettingInput() *string
	ProductionBranch() *string
	SetProductionBranch(val *string)
	ProductionBranchInput() *string
	ProductionDeploymentEnabled() interface{}
	SetProductionDeploymentEnabled(val interface{})
	ProductionDeploymentEnabledInput() interface{}
	RepoName() *string
	SetRepoName(val *string)
	RepoNameInput() *string
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
	ResetDeploymentsEnabled()
	ResetOwner()
	ResetPrCommentsEnabled()
	ResetPreviewBranchExcludes()
	ResetPreviewBranchIncludes()
	ResetPreviewDeploymentSetting()
	ResetProductionDeploymentEnabled()
	ResetRepoName()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PagesProjectSourceConfigOutputReference
type jsiiProxy_PagesProjectSourceConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_PagesProjectSourceConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagesProjectSourceConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagesProjectSourceConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagesProjectSourceConfigOutputReference) DeploymentsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deploymentsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagesProjectSourceConfigOutputReference) DeploymentsEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deploymentsEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagesProjectSourceConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagesProjectSourceConfigOutputReference) InternalValue() *PagesProjectSourceConfig {
	var returns *PagesProjectSourceConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagesProjectSourceConfigOutputReference) Owner() *string {
	var returns *string
	_jsii_.Get(
		j,
		"owner",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagesProjectSourceConfigOutputReference) OwnerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ownerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagesProjectSourceConfigOutputReference) PrCommentsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"prCommentsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagesProjectSourceConfigOutputReference) PrCommentsEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"prCommentsEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagesProjectSourceConfigOutputReference) PreviewBranchExcludes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"previewBranchExcludes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagesProjectSourceConfigOutputReference) PreviewBranchExcludesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"previewBranchExcludesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagesProjectSourceConfigOutputReference) PreviewBranchIncludes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"previewBranchIncludes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagesProjectSourceConfigOutputReference) PreviewBranchIncludesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"previewBranchIncludesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagesProjectSourceConfigOutputReference) PreviewDeploymentSetting() *string {
	var returns *string
	_jsii_.Get(
		j,
		"previewDeploymentSetting",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagesProjectSourceConfigOutputReference) PreviewDeploymentSettingInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"previewDeploymentSettingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagesProjectSourceConfigOutputReference) ProductionBranch() *string {
	var returns *string
	_jsii_.Get(
		j,
		"productionBranch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagesProjectSourceConfigOutputReference) ProductionBranchInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"productionBranchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagesProjectSourceConfigOutputReference) ProductionDeploymentEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"productionDeploymentEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagesProjectSourceConfigOutputReference) ProductionDeploymentEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"productionDeploymentEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagesProjectSourceConfigOutputReference) RepoName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repoName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagesProjectSourceConfigOutputReference) RepoNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repoNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagesProjectSourceConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagesProjectSourceConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewPagesProjectSourceConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) PagesProjectSourceConfigOutputReference {
	_init_.Initialize()

	if err := validateNewPagesProjectSourceConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PagesProjectSourceConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.pagesProject.PagesProjectSourceConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPagesProjectSourceConfigOutputReference_Override(p PagesProjectSourceConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.pagesProject.PagesProjectSourceConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PagesProjectSourceConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PagesProjectSourceConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PagesProjectSourceConfigOutputReference)SetDeploymentsEnabled(val interface{}) {
	if err := j.validateSetDeploymentsEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deploymentsEnabled",
		val,
	)
}

func (j *jsiiProxy_PagesProjectSourceConfigOutputReference)SetInternalValue(val *PagesProjectSourceConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PagesProjectSourceConfigOutputReference)SetOwner(val *string) {
	if err := j.validateSetOwnerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"owner",
		val,
	)
}

func (j *jsiiProxy_PagesProjectSourceConfigOutputReference)SetPrCommentsEnabled(val interface{}) {
	if err := j.validateSetPrCommentsEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"prCommentsEnabled",
		val,
	)
}

func (j *jsiiProxy_PagesProjectSourceConfigOutputReference)SetPreviewBranchExcludes(val *[]*string) {
	if err := j.validateSetPreviewBranchExcludesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"previewBranchExcludes",
		val,
	)
}

func (j *jsiiProxy_PagesProjectSourceConfigOutputReference)SetPreviewBranchIncludes(val *[]*string) {
	if err := j.validateSetPreviewBranchIncludesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"previewBranchIncludes",
		val,
	)
}

func (j *jsiiProxy_PagesProjectSourceConfigOutputReference)SetPreviewDeploymentSetting(val *string) {
	if err := j.validateSetPreviewDeploymentSettingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"previewDeploymentSetting",
		val,
	)
}

func (j *jsiiProxy_PagesProjectSourceConfigOutputReference)SetProductionBranch(val *string) {
	if err := j.validateSetProductionBranchParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"productionBranch",
		val,
	)
}

func (j *jsiiProxy_PagesProjectSourceConfigOutputReference)SetProductionDeploymentEnabled(val interface{}) {
	if err := j.validateSetProductionDeploymentEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"productionDeploymentEnabled",
		val,
	)
}

func (j *jsiiProxy_PagesProjectSourceConfigOutputReference)SetRepoName(val *string) {
	if err := j.validateSetRepoNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"repoName",
		val,
	)
}

func (j *jsiiProxy_PagesProjectSourceConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PagesProjectSourceConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PagesProjectSourceConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PagesProjectSourceConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := p.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PagesProjectSourceConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := p.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PagesProjectSourceConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := p.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		p,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PagesProjectSourceConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := p.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PagesProjectSourceConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := p.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		p,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PagesProjectSourceConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := p.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		p,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PagesProjectSourceConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := p.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		p,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PagesProjectSourceConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := p.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		p,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PagesProjectSourceConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := p.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		p,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PagesProjectSourceConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PagesProjectSourceConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := p.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PagesProjectSourceConfigOutputReference) ResetDeploymentsEnabled() {
	_jsii_.InvokeVoid(
		p,
		"resetDeploymentsEnabled",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PagesProjectSourceConfigOutputReference) ResetOwner() {
	_jsii_.InvokeVoid(
		p,
		"resetOwner",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PagesProjectSourceConfigOutputReference) ResetPrCommentsEnabled() {
	_jsii_.InvokeVoid(
		p,
		"resetPrCommentsEnabled",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PagesProjectSourceConfigOutputReference) ResetPreviewBranchExcludes() {
	_jsii_.InvokeVoid(
		p,
		"resetPreviewBranchExcludes",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PagesProjectSourceConfigOutputReference) ResetPreviewBranchIncludes() {
	_jsii_.InvokeVoid(
		p,
		"resetPreviewBranchIncludes",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PagesProjectSourceConfigOutputReference) ResetPreviewDeploymentSetting() {
	_jsii_.InvokeVoid(
		p,
		"resetPreviewDeploymentSetting",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PagesProjectSourceConfigOutputReference) ResetProductionDeploymentEnabled() {
	_jsii_.InvokeVoid(
		p,
		"resetProductionDeploymentEnabled",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PagesProjectSourceConfigOutputReference) ResetRepoName() {
	_jsii_.InvokeVoid(
		p,
		"resetRepoName",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PagesProjectSourceConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := p.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		p,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PagesProjectSourceConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

