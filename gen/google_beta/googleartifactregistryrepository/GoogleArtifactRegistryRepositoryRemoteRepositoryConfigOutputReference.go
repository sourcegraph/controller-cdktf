package googleartifactregistryrepository

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googleartifactregistryrepository/internal"
)

type GoogleArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference interface {
	cdktf.ComplexObject
	AptRepository() GoogleArtifactRegistryRepositoryRemoteRepositoryConfigAptRepositoryOutputReference
	AptRepositoryInput() *GoogleArtifactRegistryRepositoryRemoteRepositoryConfigAptRepository
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
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	DisableUpstreamValidation() interface{}
	SetDisableUpstreamValidation(val interface{})
	DisableUpstreamValidationInput() interface{}
	DockerRepository() GoogleArtifactRegistryRepositoryRemoteRepositoryConfigDockerRepositoryOutputReference
	DockerRepositoryInput() *GoogleArtifactRegistryRepositoryRemoteRepositoryConfigDockerRepository
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleArtifactRegistryRepositoryRemoteRepositoryConfig
	SetInternalValue(val *GoogleArtifactRegistryRepositoryRemoteRepositoryConfig)
	MavenRepository() GoogleArtifactRegistryRepositoryRemoteRepositoryConfigMavenRepositoryOutputReference
	MavenRepositoryInput() *GoogleArtifactRegistryRepositoryRemoteRepositoryConfigMavenRepository
	NpmRepository() GoogleArtifactRegistryRepositoryRemoteRepositoryConfigNpmRepositoryOutputReference
	NpmRepositoryInput() *GoogleArtifactRegistryRepositoryRemoteRepositoryConfigNpmRepository
	PythonRepository() GoogleArtifactRegistryRepositoryRemoteRepositoryConfigPythonRepositoryOutputReference
	PythonRepositoryInput() *GoogleArtifactRegistryRepositoryRemoteRepositoryConfigPythonRepository
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UpstreamCredentials() GoogleArtifactRegistryRepositoryRemoteRepositoryConfigUpstreamCredentialsOutputReference
	UpstreamCredentialsInput() *GoogleArtifactRegistryRepositoryRemoteRepositoryConfigUpstreamCredentials
	YumRepository() GoogleArtifactRegistryRepositoryRemoteRepositoryConfigYumRepositoryOutputReference
	YumRepositoryInput() *GoogleArtifactRegistryRepositoryRemoteRepositoryConfigYumRepository
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
	PutAptRepository(value *GoogleArtifactRegistryRepositoryRemoteRepositoryConfigAptRepository)
	PutDockerRepository(value *GoogleArtifactRegistryRepositoryRemoteRepositoryConfigDockerRepository)
	PutMavenRepository(value *GoogleArtifactRegistryRepositoryRemoteRepositoryConfigMavenRepository)
	PutNpmRepository(value *GoogleArtifactRegistryRepositoryRemoteRepositoryConfigNpmRepository)
	PutPythonRepository(value *GoogleArtifactRegistryRepositoryRemoteRepositoryConfigPythonRepository)
	PutUpstreamCredentials(value *GoogleArtifactRegistryRepositoryRemoteRepositoryConfigUpstreamCredentials)
	PutYumRepository(value *GoogleArtifactRegistryRepositoryRemoteRepositoryConfigYumRepository)
	ResetAptRepository()
	ResetDescription()
	ResetDisableUpstreamValidation()
	ResetDockerRepository()
	ResetMavenRepository()
	ResetNpmRepository()
	ResetPythonRepository()
	ResetUpstreamCredentials()
	ResetYumRepository()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference
type jsiiProxy_GoogleArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference) AptRepository() GoogleArtifactRegistryRepositoryRemoteRepositoryConfigAptRepositoryOutputReference {
	var returns GoogleArtifactRegistryRepositoryRemoteRepositoryConfigAptRepositoryOutputReference
	_jsii_.Get(
		j,
		"aptRepository",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference) AptRepositoryInput() *GoogleArtifactRegistryRepositoryRemoteRepositoryConfigAptRepository {
	var returns *GoogleArtifactRegistryRepositoryRemoteRepositoryConfigAptRepository
	_jsii_.Get(
		j,
		"aptRepositoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference) DisableUpstreamValidation() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableUpstreamValidation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference) DisableUpstreamValidationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableUpstreamValidationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference) DockerRepository() GoogleArtifactRegistryRepositoryRemoteRepositoryConfigDockerRepositoryOutputReference {
	var returns GoogleArtifactRegistryRepositoryRemoteRepositoryConfigDockerRepositoryOutputReference
	_jsii_.Get(
		j,
		"dockerRepository",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference) DockerRepositoryInput() *GoogleArtifactRegistryRepositoryRemoteRepositoryConfigDockerRepository {
	var returns *GoogleArtifactRegistryRepositoryRemoteRepositoryConfigDockerRepository
	_jsii_.Get(
		j,
		"dockerRepositoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference) InternalValue() *GoogleArtifactRegistryRepositoryRemoteRepositoryConfig {
	var returns *GoogleArtifactRegistryRepositoryRemoteRepositoryConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference) MavenRepository() GoogleArtifactRegistryRepositoryRemoteRepositoryConfigMavenRepositoryOutputReference {
	var returns GoogleArtifactRegistryRepositoryRemoteRepositoryConfigMavenRepositoryOutputReference
	_jsii_.Get(
		j,
		"mavenRepository",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference) MavenRepositoryInput() *GoogleArtifactRegistryRepositoryRemoteRepositoryConfigMavenRepository {
	var returns *GoogleArtifactRegistryRepositoryRemoteRepositoryConfigMavenRepository
	_jsii_.Get(
		j,
		"mavenRepositoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference) NpmRepository() GoogleArtifactRegistryRepositoryRemoteRepositoryConfigNpmRepositoryOutputReference {
	var returns GoogleArtifactRegistryRepositoryRemoteRepositoryConfigNpmRepositoryOutputReference
	_jsii_.Get(
		j,
		"npmRepository",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference) NpmRepositoryInput() *GoogleArtifactRegistryRepositoryRemoteRepositoryConfigNpmRepository {
	var returns *GoogleArtifactRegistryRepositoryRemoteRepositoryConfigNpmRepository
	_jsii_.Get(
		j,
		"npmRepositoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference) PythonRepository() GoogleArtifactRegistryRepositoryRemoteRepositoryConfigPythonRepositoryOutputReference {
	var returns GoogleArtifactRegistryRepositoryRemoteRepositoryConfigPythonRepositoryOutputReference
	_jsii_.Get(
		j,
		"pythonRepository",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference) PythonRepositoryInput() *GoogleArtifactRegistryRepositoryRemoteRepositoryConfigPythonRepository {
	var returns *GoogleArtifactRegistryRepositoryRemoteRepositoryConfigPythonRepository
	_jsii_.Get(
		j,
		"pythonRepositoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference) UpstreamCredentials() GoogleArtifactRegistryRepositoryRemoteRepositoryConfigUpstreamCredentialsOutputReference {
	var returns GoogleArtifactRegistryRepositoryRemoteRepositoryConfigUpstreamCredentialsOutputReference
	_jsii_.Get(
		j,
		"upstreamCredentials",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference) UpstreamCredentialsInput() *GoogleArtifactRegistryRepositoryRemoteRepositoryConfigUpstreamCredentials {
	var returns *GoogleArtifactRegistryRepositoryRemoteRepositoryConfigUpstreamCredentials
	_jsii_.Get(
		j,
		"upstreamCredentialsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference) YumRepository() GoogleArtifactRegistryRepositoryRemoteRepositoryConfigYumRepositoryOutputReference {
	var returns GoogleArtifactRegistryRepositoryRemoteRepositoryConfigYumRepositoryOutputReference
	_jsii_.Get(
		j,
		"yumRepository",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference) YumRepositoryInput() *GoogleArtifactRegistryRepositoryRemoteRepositoryConfigYumRepository {
	var returns *GoogleArtifactRegistryRepositoryRemoteRepositoryConfigYumRepository
	_jsii_.Get(
		j,
		"yumRepositoryInput",
		&returns,
	)
	return returns
}


func NewGoogleArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleArtifactRegistryRepositoryRemoteRepositoryConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleArtifactRegistryRepository.GoogleArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference_Override(g GoogleArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleArtifactRegistryRepository.GoogleArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_GoogleArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference)SetDisableUpstreamValidation(val interface{}) {
	if err := j.validateSetDisableUpstreamValidationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableUpstreamValidation",
		val,
	)
}

func (j *jsiiProxy_GoogleArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference)SetInternalValue(val *GoogleArtifactRegistryRepositoryRemoteRepositoryConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference) PutAptRepository(value *GoogleArtifactRegistryRepositoryRemoteRepositoryConfigAptRepository) {
	if err := g.validatePutAptRepositoryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAptRepository",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference) PutDockerRepository(value *GoogleArtifactRegistryRepositoryRemoteRepositoryConfigDockerRepository) {
	if err := g.validatePutDockerRepositoryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDockerRepository",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference) PutMavenRepository(value *GoogleArtifactRegistryRepositoryRemoteRepositoryConfigMavenRepository) {
	if err := g.validatePutMavenRepositoryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putMavenRepository",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference) PutNpmRepository(value *GoogleArtifactRegistryRepositoryRemoteRepositoryConfigNpmRepository) {
	if err := g.validatePutNpmRepositoryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putNpmRepository",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference) PutPythonRepository(value *GoogleArtifactRegistryRepositoryRemoteRepositoryConfigPythonRepository) {
	if err := g.validatePutPythonRepositoryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putPythonRepository",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference) PutUpstreamCredentials(value *GoogleArtifactRegistryRepositoryRemoteRepositoryConfigUpstreamCredentials) {
	if err := g.validatePutUpstreamCredentialsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putUpstreamCredentials",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference) PutYumRepository(value *GoogleArtifactRegistryRepositoryRemoteRepositoryConfigYumRepository) {
	if err := g.validatePutYumRepositoryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putYumRepository",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference) ResetAptRepository() {
	_jsii_.InvokeVoid(
		g,
		"resetAptRepository",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference) ResetDescription() {
	_jsii_.InvokeVoid(
		g,
		"resetDescription",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference) ResetDisableUpstreamValidation() {
	_jsii_.InvokeVoid(
		g,
		"resetDisableUpstreamValidation",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference) ResetDockerRepository() {
	_jsii_.InvokeVoid(
		g,
		"resetDockerRepository",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference) ResetMavenRepository() {
	_jsii_.InvokeVoid(
		g,
		"resetMavenRepository",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference) ResetNpmRepository() {
	_jsii_.InvokeVoid(
		g,
		"resetNpmRepository",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference) ResetPythonRepository() {
	_jsii_.InvokeVoid(
		g,
		"resetPythonRepository",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference) ResetUpstreamCredentials() {
	_jsii_.InvokeVoid(
		g,
		"resetUpstreamCredentials",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference) ResetYumRepository() {
	_jsii_.InvokeVoid(
		g,
		"resetYumRepository",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

