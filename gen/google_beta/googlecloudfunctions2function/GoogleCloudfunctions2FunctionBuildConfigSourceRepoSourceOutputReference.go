package googlecloudfunctions2function

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecloudfunctions2function/internal"
)

type GoogleCloudfunctions2FunctionBuildConfigSourceRepoSourceOutputReference interface {
	cdktf.ComplexObject
	BranchName() *string
	SetBranchName(val *string)
	BranchNameInput() *string
	CommitSha() *string
	SetCommitSha(val *string)
	CommitShaInput() *string
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
	Dir() *string
	SetDir(val *string)
	DirInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleCloudfunctions2FunctionBuildConfigSourceRepoSource
	SetInternalValue(val *GoogleCloudfunctions2FunctionBuildConfigSourceRepoSource)
	InvertRegex() interface{}
	SetInvertRegex(val interface{})
	InvertRegexInput() interface{}
	ProjectId() *string
	SetProjectId(val *string)
	ProjectIdInput() *string
	RepoName() *string
	SetRepoName(val *string)
	RepoNameInput() *string
	TagName() *string
	SetTagName(val *string)
	TagNameInput() *string
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
	ResetBranchName()
	ResetCommitSha()
	ResetDir()
	ResetInvertRegex()
	ResetProjectId()
	ResetRepoName()
	ResetTagName()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleCloudfunctions2FunctionBuildConfigSourceRepoSourceOutputReference
type jsiiProxy_GoogleCloudfunctions2FunctionBuildConfigSourceRepoSourceOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleCloudfunctions2FunctionBuildConfigSourceRepoSourceOutputReference) BranchName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"branchName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctions2FunctionBuildConfigSourceRepoSourceOutputReference) BranchNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"branchNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctions2FunctionBuildConfigSourceRepoSourceOutputReference) CommitSha() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commitSha",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctions2FunctionBuildConfigSourceRepoSourceOutputReference) CommitShaInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commitShaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctions2FunctionBuildConfigSourceRepoSourceOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctions2FunctionBuildConfigSourceRepoSourceOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctions2FunctionBuildConfigSourceRepoSourceOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctions2FunctionBuildConfigSourceRepoSourceOutputReference) Dir() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dir",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctions2FunctionBuildConfigSourceRepoSourceOutputReference) DirInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dirInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctions2FunctionBuildConfigSourceRepoSourceOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctions2FunctionBuildConfigSourceRepoSourceOutputReference) InternalValue() *GoogleCloudfunctions2FunctionBuildConfigSourceRepoSource {
	var returns *GoogleCloudfunctions2FunctionBuildConfigSourceRepoSource
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctions2FunctionBuildConfigSourceRepoSourceOutputReference) InvertRegex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"invertRegex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctions2FunctionBuildConfigSourceRepoSourceOutputReference) InvertRegexInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"invertRegexInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctions2FunctionBuildConfigSourceRepoSourceOutputReference) ProjectId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctions2FunctionBuildConfigSourceRepoSourceOutputReference) ProjectIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctions2FunctionBuildConfigSourceRepoSourceOutputReference) RepoName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repoName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctions2FunctionBuildConfigSourceRepoSourceOutputReference) RepoNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repoNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctions2FunctionBuildConfigSourceRepoSourceOutputReference) TagName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tagName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctions2FunctionBuildConfigSourceRepoSourceOutputReference) TagNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tagNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctions2FunctionBuildConfigSourceRepoSourceOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctions2FunctionBuildConfigSourceRepoSourceOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleCloudfunctions2FunctionBuildConfigSourceRepoSourceOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleCloudfunctions2FunctionBuildConfigSourceRepoSourceOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleCloudfunctions2FunctionBuildConfigSourceRepoSourceOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleCloudfunctions2FunctionBuildConfigSourceRepoSourceOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleCloudfunctions2Function.GoogleCloudfunctions2FunctionBuildConfigSourceRepoSourceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleCloudfunctions2FunctionBuildConfigSourceRepoSourceOutputReference_Override(g GoogleCloudfunctions2FunctionBuildConfigSourceRepoSourceOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleCloudfunctions2Function.GoogleCloudfunctions2FunctionBuildConfigSourceRepoSourceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleCloudfunctions2FunctionBuildConfigSourceRepoSourceOutputReference)SetBranchName(val *string) {
	if err := j.validateSetBranchNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"branchName",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudfunctions2FunctionBuildConfigSourceRepoSourceOutputReference)SetCommitSha(val *string) {
	if err := j.validateSetCommitShaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"commitSha",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudfunctions2FunctionBuildConfigSourceRepoSourceOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudfunctions2FunctionBuildConfigSourceRepoSourceOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudfunctions2FunctionBuildConfigSourceRepoSourceOutputReference)SetDir(val *string) {
	if err := j.validateSetDirParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dir",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudfunctions2FunctionBuildConfigSourceRepoSourceOutputReference)SetInternalValue(val *GoogleCloudfunctions2FunctionBuildConfigSourceRepoSource) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudfunctions2FunctionBuildConfigSourceRepoSourceOutputReference)SetInvertRegex(val interface{}) {
	if err := j.validateSetInvertRegexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"invertRegex",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudfunctions2FunctionBuildConfigSourceRepoSourceOutputReference)SetProjectId(val *string) {
	if err := j.validateSetProjectIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"projectId",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudfunctions2FunctionBuildConfigSourceRepoSourceOutputReference)SetRepoName(val *string) {
	if err := j.validateSetRepoNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"repoName",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudfunctions2FunctionBuildConfigSourceRepoSourceOutputReference)SetTagName(val *string) {
	if err := j.validateSetTagNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagName",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudfunctions2FunctionBuildConfigSourceRepoSourceOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudfunctions2FunctionBuildConfigSourceRepoSourceOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleCloudfunctions2FunctionBuildConfigSourceRepoSourceOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCloudfunctions2FunctionBuildConfigSourceRepoSourceOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleCloudfunctions2FunctionBuildConfigSourceRepoSourceOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleCloudfunctions2FunctionBuildConfigSourceRepoSourceOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleCloudfunctions2FunctionBuildConfigSourceRepoSourceOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleCloudfunctions2FunctionBuildConfigSourceRepoSourceOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleCloudfunctions2FunctionBuildConfigSourceRepoSourceOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleCloudfunctions2FunctionBuildConfigSourceRepoSourceOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleCloudfunctions2FunctionBuildConfigSourceRepoSourceOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleCloudfunctions2FunctionBuildConfigSourceRepoSourceOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleCloudfunctions2FunctionBuildConfigSourceRepoSourceOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCloudfunctions2FunctionBuildConfigSourceRepoSourceOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleCloudfunctions2FunctionBuildConfigSourceRepoSourceOutputReference) ResetBranchName() {
	_jsii_.InvokeVoid(
		g,
		"resetBranchName",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudfunctions2FunctionBuildConfigSourceRepoSourceOutputReference) ResetCommitSha() {
	_jsii_.InvokeVoid(
		g,
		"resetCommitSha",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudfunctions2FunctionBuildConfigSourceRepoSourceOutputReference) ResetDir() {
	_jsii_.InvokeVoid(
		g,
		"resetDir",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudfunctions2FunctionBuildConfigSourceRepoSourceOutputReference) ResetInvertRegex() {
	_jsii_.InvokeVoid(
		g,
		"resetInvertRegex",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudfunctions2FunctionBuildConfigSourceRepoSourceOutputReference) ResetProjectId() {
	_jsii_.InvokeVoid(
		g,
		"resetProjectId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudfunctions2FunctionBuildConfigSourceRepoSourceOutputReference) ResetRepoName() {
	_jsii_.InvokeVoid(
		g,
		"resetRepoName",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudfunctions2FunctionBuildConfigSourceRepoSourceOutputReference) ResetTagName() {
	_jsii_.InvokeVoid(
		g,
		"resetTagName",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudfunctions2FunctionBuildConfigSourceRepoSourceOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleCloudfunctions2FunctionBuildConfigSourceRepoSourceOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

