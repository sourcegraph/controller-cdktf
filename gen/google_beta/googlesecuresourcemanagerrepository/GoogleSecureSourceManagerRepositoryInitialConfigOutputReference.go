package googlesecuresourcemanagerrepository

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlesecuresourcemanagerrepository/internal"
)

type GoogleSecureSourceManagerRepositoryInitialConfigOutputReference interface {
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
	DefaultBranch() *string
	SetDefaultBranch(val *string)
	DefaultBranchInput() *string
	// Experimental.
	Fqn() *string
	Gitignores() *[]*string
	SetGitignores(val *[]*string)
	GitignoresInput() *[]*string
	InternalValue() *GoogleSecureSourceManagerRepositoryInitialConfig
	SetInternalValue(val *GoogleSecureSourceManagerRepositoryInitialConfig)
	License() *string
	SetLicense(val *string)
	LicenseInput() *string
	Readme() *string
	SetReadme(val *string)
	ReadmeInput() *string
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
	ResetDefaultBranch()
	ResetGitignores()
	ResetLicense()
	ResetReadme()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleSecureSourceManagerRepositoryInitialConfigOutputReference
type jsiiProxy_GoogleSecureSourceManagerRepositoryInitialConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleSecureSourceManagerRepositoryInitialConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSecureSourceManagerRepositoryInitialConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSecureSourceManagerRepositoryInitialConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSecureSourceManagerRepositoryInitialConfigOutputReference) DefaultBranch() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultBranch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSecureSourceManagerRepositoryInitialConfigOutputReference) DefaultBranchInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultBranchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSecureSourceManagerRepositoryInitialConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSecureSourceManagerRepositoryInitialConfigOutputReference) Gitignores() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"gitignores",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSecureSourceManagerRepositoryInitialConfigOutputReference) GitignoresInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"gitignoresInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSecureSourceManagerRepositoryInitialConfigOutputReference) InternalValue() *GoogleSecureSourceManagerRepositoryInitialConfig {
	var returns *GoogleSecureSourceManagerRepositoryInitialConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSecureSourceManagerRepositoryInitialConfigOutputReference) License() *string {
	var returns *string
	_jsii_.Get(
		j,
		"license",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSecureSourceManagerRepositoryInitialConfigOutputReference) LicenseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"licenseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSecureSourceManagerRepositoryInitialConfigOutputReference) Readme() *string {
	var returns *string
	_jsii_.Get(
		j,
		"readme",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSecureSourceManagerRepositoryInitialConfigOutputReference) ReadmeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"readmeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSecureSourceManagerRepositoryInitialConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSecureSourceManagerRepositoryInitialConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleSecureSourceManagerRepositoryInitialConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleSecureSourceManagerRepositoryInitialConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleSecureSourceManagerRepositoryInitialConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleSecureSourceManagerRepositoryInitialConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleSecureSourceManagerRepository.GoogleSecureSourceManagerRepositoryInitialConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleSecureSourceManagerRepositoryInitialConfigOutputReference_Override(g GoogleSecureSourceManagerRepositoryInitialConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleSecureSourceManagerRepository.GoogleSecureSourceManagerRepositoryInitialConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleSecureSourceManagerRepositoryInitialConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleSecureSourceManagerRepositoryInitialConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleSecureSourceManagerRepositoryInitialConfigOutputReference)SetDefaultBranch(val *string) {
	if err := j.validateSetDefaultBranchParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultBranch",
		val,
	)
}

func (j *jsiiProxy_GoogleSecureSourceManagerRepositoryInitialConfigOutputReference)SetGitignores(val *[]*string) {
	if err := j.validateSetGitignoresParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gitignores",
		val,
	)
}

func (j *jsiiProxy_GoogleSecureSourceManagerRepositoryInitialConfigOutputReference)SetInternalValue(val *GoogleSecureSourceManagerRepositoryInitialConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleSecureSourceManagerRepositoryInitialConfigOutputReference)SetLicense(val *string) {
	if err := j.validateSetLicenseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"license",
		val,
	)
}

func (j *jsiiProxy_GoogleSecureSourceManagerRepositoryInitialConfigOutputReference)SetReadme(val *string) {
	if err := j.validateSetReadmeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"readme",
		val,
	)
}

func (j *jsiiProxy_GoogleSecureSourceManagerRepositoryInitialConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleSecureSourceManagerRepositoryInitialConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleSecureSourceManagerRepositoryInitialConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleSecureSourceManagerRepositoryInitialConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleSecureSourceManagerRepositoryInitialConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleSecureSourceManagerRepositoryInitialConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleSecureSourceManagerRepositoryInitialConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleSecureSourceManagerRepositoryInitialConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleSecureSourceManagerRepositoryInitialConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleSecureSourceManagerRepositoryInitialConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleSecureSourceManagerRepositoryInitialConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleSecureSourceManagerRepositoryInitialConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleSecureSourceManagerRepositoryInitialConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleSecureSourceManagerRepositoryInitialConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleSecureSourceManagerRepositoryInitialConfigOutputReference) ResetDefaultBranch() {
	_jsii_.InvokeVoid(
		g,
		"resetDefaultBranch",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleSecureSourceManagerRepositoryInitialConfigOutputReference) ResetGitignores() {
	_jsii_.InvokeVoid(
		g,
		"resetGitignores",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleSecureSourceManagerRepositoryInitialConfigOutputReference) ResetLicense() {
	_jsii_.InvokeVoid(
		g,
		"resetLicense",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleSecureSourceManagerRepositoryInitialConfigOutputReference) ResetReadme() {
	_jsii_.InvokeVoid(
		g,
		"resetReadme",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleSecureSourceManagerRepositoryInitialConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleSecureSourceManagerRepositoryInitialConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

