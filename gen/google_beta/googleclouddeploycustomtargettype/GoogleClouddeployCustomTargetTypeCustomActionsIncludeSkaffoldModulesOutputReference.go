package googleclouddeploycustomtargettype

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googleclouddeploycustomtargettype/internal"
)

type GoogleClouddeployCustomTargetTypeCustomActionsIncludeSkaffoldModulesOutputReference interface {
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
	Configs() *[]*string
	SetConfigs(val *[]*string)
	ConfigsInput() *[]*string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	Git() GoogleClouddeployCustomTargetTypeCustomActionsIncludeSkaffoldModulesGitOutputReference
	GitInput() *GoogleClouddeployCustomTargetTypeCustomActionsIncludeSkaffoldModulesGit
	GoogleCloudBuildRepo() GoogleClouddeployCustomTargetTypeCustomActionsIncludeSkaffoldModulesGoogleCloudBuildRepoOutputReference
	GoogleCloudBuildRepoInput() *GoogleClouddeployCustomTargetTypeCustomActionsIncludeSkaffoldModulesGoogleCloudBuildRepo
	GoogleCloudStorage() GoogleClouddeployCustomTargetTypeCustomActionsIncludeSkaffoldModulesGoogleCloudStorageOutputReference
	GoogleCloudStorageInput() *GoogleClouddeployCustomTargetTypeCustomActionsIncludeSkaffoldModulesGoogleCloudStorage
	InternalValue() interface{}
	SetInternalValue(val interface{})
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
	PutGit(value *GoogleClouddeployCustomTargetTypeCustomActionsIncludeSkaffoldModulesGit)
	PutGoogleCloudBuildRepo(value *GoogleClouddeployCustomTargetTypeCustomActionsIncludeSkaffoldModulesGoogleCloudBuildRepo)
	PutGoogleCloudStorage(value *GoogleClouddeployCustomTargetTypeCustomActionsIncludeSkaffoldModulesGoogleCloudStorage)
	ResetConfigs()
	ResetGit()
	ResetGoogleCloudBuildRepo()
	ResetGoogleCloudStorage()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleClouddeployCustomTargetTypeCustomActionsIncludeSkaffoldModulesOutputReference
type jsiiProxy_GoogleClouddeployCustomTargetTypeCustomActionsIncludeSkaffoldModulesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleClouddeployCustomTargetTypeCustomActionsIncludeSkaffoldModulesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployCustomTargetTypeCustomActionsIncludeSkaffoldModulesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployCustomTargetTypeCustomActionsIncludeSkaffoldModulesOutputReference) Configs() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"configs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployCustomTargetTypeCustomActionsIncludeSkaffoldModulesOutputReference) ConfigsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"configsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployCustomTargetTypeCustomActionsIncludeSkaffoldModulesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployCustomTargetTypeCustomActionsIncludeSkaffoldModulesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployCustomTargetTypeCustomActionsIncludeSkaffoldModulesOutputReference) Git() GoogleClouddeployCustomTargetTypeCustomActionsIncludeSkaffoldModulesGitOutputReference {
	var returns GoogleClouddeployCustomTargetTypeCustomActionsIncludeSkaffoldModulesGitOutputReference
	_jsii_.Get(
		j,
		"git",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployCustomTargetTypeCustomActionsIncludeSkaffoldModulesOutputReference) GitInput() *GoogleClouddeployCustomTargetTypeCustomActionsIncludeSkaffoldModulesGit {
	var returns *GoogleClouddeployCustomTargetTypeCustomActionsIncludeSkaffoldModulesGit
	_jsii_.Get(
		j,
		"gitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployCustomTargetTypeCustomActionsIncludeSkaffoldModulesOutputReference) GoogleCloudBuildRepo() GoogleClouddeployCustomTargetTypeCustomActionsIncludeSkaffoldModulesGoogleCloudBuildRepoOutputReference {
	var returns GoogleClouddeployCustomTargetTypeCustomActionsIncludeSkaffoldModulesGoogleCloudBuildRepoOutputReference
	_jsii_.Get(
		j,
		"googleCloudBuildRepo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployCustomTargetTypeCustomActionsIncludeSkaffoldModulesOutputReference) GoogleCloudBuildRepoInput() *GoogleClouddeployCustomTargetTypeCustomActionsIncludeSkaffoldModulesGoogleCloudBuildRepo {
	var returns *GoogleClouddeployCustomTargetTypeCustomActionsIncludeSkaffoldModulesGoogleCloudBuildRepo
	_jsii_.Get(
		j,
		"googleCloudBuildRepoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployCustomTargetTypeCustomActionsIncludeSkaffoldModulesOutputReference) GoogleCloudStorage() GoogleClouddeployCustomTargetTypeCustomActionsIncludeSkaffoldModulesGoogleCloudStorageOutputReference {
	var returns GoogleClouddeployCustomTargetTypeCustomActionsIncludeSkaffoldModulesGoogleCloudStorageOutputReference
	_jsii_.Get(
		j,
		"googleCloudStorage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployCustomTargetTypeCustomActionsIncludeSkaffoldModulesOutputReference) GoogleCloudStorageInput() *GoogleClouddeployCustomTargetTypeCustomActionsIncludeSkaffoldModulesGoogleCloudStorage {
	var returns *GoogleClouddeployCustomTargetTypeCustomActionsIncludeSkaffoldModulesGoogleCloudStorage
	_jsii_.Get(
		j,
		"googleCloudStorageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployCustomTargetTypeCustomActionsIncludeSkaffoldModulesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployCustomTargetTypeCustomActionsIncludeSkaffoldModulesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployCustomTargetTypeCustomActionsIncludeSkaffoldModulesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleClouddeployCustomTargetTypeCustomActionsIncludeSkaffoldModulesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) GoogleClouddeployCustomTargetTypeCustomActionsIncludeSkaffoldModulesOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleClouddeployCustomTargetTypeCustomActionsIncludeSkaffoldModulesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleClouddeployCustomTargetTypeCustomActionsIncludeSkaffoldModulesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleClouddeployCustomTargetType.GoogleClouddeployCustomTargetTypeCustomActionsIncludeSkaffoldModulesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewGoogleClouddeployCustomTargetTypeCustomActionsIncludeSkaffoldModulesOutputReference_Override(g GoogleClouddeployCustomTargetTypeCustomActionsIncludeSkaffoldModulesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleClouddeployCustomTargetType.GoogleClouddeployCustomTargetTypeCustomActionsIncludeSkaffoldModulesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		g,
	)
}

func (j *jsiiProxy_GoogleClouddeployCustomTargetTypeCustomActionsIncludeSkaffoldModulesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleClouddeployCustomTargetTypeCustomActionsIncludeSkaffoldModulesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleClouddeployCustomTargetTypeCustomActionsIncludeSkaffoldModulesOutputReference)SetConfigs(val *[]*string) {
	if err := j.validateSetConfigsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"configs",
		val,
	)
}

func (j *jsiiProxy_GoogleClouddeployCustomTargetTypeCustomActionsIncludeSkaffoldModulesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleClouddeployCustomTargetTypeCustomActionsIncludeSkaffoldModulesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleClouddeployCustomTargetTypeCustomActionsIncludeSkaffoldModulesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleClouddeployCustomTargetTypeCustomActionsIncludeSkaffoldModulesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleClouddeployCustomTargetTypeCustomActionsIncludeSkaffoldModulesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleClouddeployCustomTargetTypeCustomActionsIncludeSkaffoldModulesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleClouddeployCustomTargetTypeCustomActionsIncludeSkaffoldModulesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleClouddeployCustomTargetTypeCustomActionsIncludeSkaffoldModulesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleClouddeployCustomTargetTypeCustomActionsIncludeSkaffoldModulesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleClouddeployCustomTargetTypeCustomActionsIncludeSkaffoldModulesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleClouddeployCustomTargetTypeCustomActionsIncludeSkaffoldModulesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleClouddeployCustomTargetTypeCustomActionsIncludeSkaffoldModulesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleClouddeployCustomTargetTypeCustomActionsIncludeSkaffoldModulesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleClouddeployCustomTargetTypeCustomActionsIncludeSkaffoldModulesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleClouddeployCustomTargetTypeCustomActionsIncludeSkaffoldModulesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleClouddeployCustomTargetTypeCustomActionsIncludeSkaffoldModulesOutputReference) PutGit(value *GoogleClouddeployCustomTargetTypeCustomActionsIncludeSkaffoldModulesGit) {
	if err := g.validatePutGitParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putGit",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleClouddeployCustomTargetTypeCustomActionsIncludeSkaffoldModulesOutputReference) PutGoogleCloudBuildRepo(value *GoogleClouddeployCustomTargetTypeCustomActionsIncludeSkaffoldModulesGoogleCloudBuildRepo) {
	if err := g.validatePutGoogleCloudBuildRepoParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putGoogleCloudBuildRepo",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleClouddeployCustomTargetTypeCustomActionsIncludeSkaffoldModulesOutputReference) PutGoogleCloudStorage(value *GoogleClouddeployCustomTargetTypeCustomActionsIncludeSkaffoldModulesGoogleCloudStorage) {
	if err := g.validatePutGoogleCloudStorageParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putGoogleCloudStorage",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleClouddeployCustomTargetTypeCustomActionsIncludeSkaffoldModulesOutputReference) ResetConfigs() {
	_jsii_.InvokeVoid(
		g,
		"resetConfigs",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleClouddeployCustomTargetTypeCustomActionsIncludeSkaffoldModulesOutputReference) ResetGit() {
	_jsii_.InvokeVoid(
		g,
		"resetGit",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleClouddeployCustomTargetTypeCustomActionsIncludeSkaffoldModulesOutputReference) ResetGoogleCloudBuildRepo() {
	_jsii_.InvokeVoid(
		g,
		"resetGoogleCloudBuildRepo",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleClouddeployCustomTargetTypeCustomActionsIncludeSkaffoldModulesOutputReference) ResetGoogleCloudStorage() {
	_jsii_.InvokeVoid(
		g,
		"resetGoogleCloudStorage",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleClouddeployCustomTargetTypeCustomActionsIncludeSkaffoldModulesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleClouddeployCustomTargetTypeCustomActionsIncludeSkaffoldModulesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

