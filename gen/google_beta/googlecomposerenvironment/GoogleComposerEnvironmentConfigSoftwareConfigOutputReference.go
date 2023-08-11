package googlecomposerenvironment

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecomposerenvironment/internal"
)

type GoogleComposerEnvironmentConfigSoftwareConfigOutputReference interface {
	cdktf.ComplexObject
	AirflowConfigOverrides() *map[string]*string
	SetAirflowConfigOverrides(val *map[string]*string)
	AirflowConfigOverridesInput() *map[string]*string
	CloudDataLineageIntegration() GoogleComposerEnvironmentConfigSoftwareConfigCloudDataLineageIntegrationOutputReference
	CloudDataLineageIntegrationInput() *GoogleComposerEnvironmentConfigSoftwareConfigCloudDataLineageIntegration
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
	EnvVariables() *map[string]*string
	SetEnvVariables(val *map[string]*string)
	EnvVariablesInput() *map[string]*string
	// Experimental.
	Fqn() *string
	ImageVersion() *string
	SetImageVersion(val *string)
	ImageVersionInput() *string
	InternalValue() *GoogleComposerEnvironmentConfigSoftwareConfig
	SetInternalValue(val *GoogleComposerEnvironmentConfigSoftwareConfig)
	PypiPackages() *map[string]*string
	SetPypiPackages(val *map[string]*string)
	PypiPackagesInput() *map[string]*string
	PythonVersion() *string
	SetPythonVersion(val *string)
	PythonVersionInput() *string
	SchedulerCount() *float64
	SetSchedulerCount(val *float64)
	SchedulerCountInput() *float64
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
	PutCloudDataLineageIntegration(value *GoogleComposerEnvironmentConfigSoftwareConfigCloudDataLineageIntegration)
	ResetAirflowConfigOverrides()
	ResetCloudDataLineageIntegration()
	ResetEnvVariables()
	ResetImageVersion()
	ResetPypiPackages()
	ResetPythonVersion()
	ResetSchedulerCount()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleComposerEnvironmentConfigSoftwareConfigOutputReference
type jsiiProxy_GoogleComposerEnvironmentConfigSoftwareConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleComposerEnvironmentConfigSoftwareConfigOutputReference) AirflowConfigOverrides() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"airflowConfigOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComposerEnvironmentConfigSoftwareConfigOutputReference) AirflowConfigOverridesInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"airflowConfigOverridesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComposerEnvironmentConfigSoftwareConfigOutputReference) CloudDataLineageIntegration() GoogleComposerEnvironmentConfigSoftwareConfigCloudDataLineageIntegrationOutputReference {
	var returns GoogleComposerEnvironmentConfigSoftwareConfigCloudDataLineageIntegrationOutputReference
	_jsii_.Get(
		j,
		"cloudDataLineageIntegration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComposerEnvironmentConfigSoftwareConfigOutputReference) CloudDataLineageIntegrationInput() *GoogleComposerEnvironmentConfigSoftwareConfigCloudDataLineageIntegration {
	var returns *GoogleComposerEnvironmentConfigSoftwareConfigCloudDataLineageIntegration
	_jsii_.Get(
		j,
		"cloudDataLineageIntegrationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComposerEnvironmentConfigSoftwareConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComposerEnvironmentConfigSoftwareConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComposerEnvironmentConfigSoftwareConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComposerEnvironmentConfigSoftwareConfigOutputReference) EnvVariables() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"envVariables",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComposerEnvironmentConfigSoftwareConfigOutputReference) EnvVariablesInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"envVariablesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComposerEnvironmentConfigSoftwareConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComposerEnvironmentConfigSoftwareConfigOutputReference) ImageVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComposerEnvironmentConfigSoftwareConfigOutputReference) ImageVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComposerEnvironmentConfigSoftwareConfigOutputReference) InternalValue() *GoogleComposerEnvironmentConfigSoftwareConfig {
	var returns *GoogleComposerEnvironmentConfigSoftwareConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComposerEnvironmentConfigSoftwareConfigOutputReference) PypiPackages() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"pypiPackages",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComposerEnvironmentConfigSoftwareConfigOutputReference) PypiPackagesInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"pypiPackagesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComposerEnvironmentConfigSoftwareConfigOutputReference) PythonVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pythonVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComposerEnvironmentConfigSoftwareConfigOutputReference) PythonVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pythonVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComposerEnvironmentConfigSoftwareConfigOutputReference) SchedulerCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"schedulerCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComposerEnvironmentConfigSoftwareConfigOutputReference) SchedulerCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"schedulerCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComposerEnvironmentConfigSoftwareConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComposerEnvironmentConfigSoftwareConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleComposerEnvironmentConfigSoftwareConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleComposerEnvironmentConfigSoftwareConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleComposerEnvironmentConfigSoftwareConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleComposerEnvironmentConfigSoftwareConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleComposerEnvironment.GoogleComposerEnvironmentConfigSoftwareConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleComposerEnvironmentConfigSoftwareConfigOutputReference_Override(g GoogleComposerEnvironmentConfigSoftwareConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleComposerEnvironment.GoogleComposerEnvironmentConfigSoftwareConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleComposerEnvironmentConfigSoftwareConfigOutputReference)SetAirflowConfigOverrides(val *map[string]*string) {
	if err := j.validateSetAirflowConfigOverridesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"airflowConfigOverrides",
		val,
	)
}

func (j *jsiiProxy_GoogleComposerEnvironmentConfigSoftwareConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleComposerEnvironmentConfigSoftwareConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleComposerEnvironmentConfigSoftwareConfigOutputReference)SetEnvVariables(val *map[string]*string) {
	if err := j.validateSetEnvVariablesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"envVariables",
		val,
	)
}

func (j *jsiiProxy_GoogleComposerEnvironmentConfigSoftwareConfigOutputReference)SetImageVersion(val *string) {
	if err := j.validateSetImageVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"imageVersion",
		val,
	)
}

func (j *jsiiProxy_GoogleComposerEnvironmentConfigSoftwareConfigOutputReference)SetInternalValue(val *GoogleComposerEnvironmentConfigSoftwareConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleComposerEnvironmentConfigSoftwareConfigOutputReference)SetPypiPackages(val *map[string]*string) {
	if err := j.validateSetPypiPackagesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pypiPackages",
		val,
	)
}

func (j *jsiiProxy_GoogleComposerEnvironmentConfigSoftwareConfigOutputReference)SetPythonVersion(val *string) {
	if err := j.validateSetPythonVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pythonVersion",
		val,
	)
}

func (j *jsiiProxy_GoogleComposerEnvironmentConfigSoftwareConfigOutputReference)SetSchedulerCount(val *float64) {
	if err := j.validateSetSchedulerCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"schedulerCount",
		val,
	)
}

func (j *jsiiProxy_GoogleComposerEnvironmentConfigSoftwareConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleComposerEnvironmentConfigSoftwareConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleComposerEnvironmentConfigSoftwareConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComposerEnvironmentConfigSoftwareConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleComposerEnvironmentConfigSoftwareConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleComposerEnvironmentConfigSoftwareConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleComposerEnvironmentConfigSoftwareConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleComposerEnvironmentConfigSoftwareConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleComposerEnvironmentConfigSoftwareConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleComposerEnvironmentConfigSoftwareConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleComposerEnvironmentConfigSoftwareConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleComposerEnvironmentConfigSoftwareConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleComposerEnvironmentConfigSoftwareConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComposerEnvironmentConfigSoftwareConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleComposerEnvironmentConfigSoftwareConfigOutputReference) PutCloudDataLineageIntegration(value *GoogleComposerEnvironmentConfigSoftwareConfigCloudDataLineageIntegration) {
	if err := g.validatePutCloudDataLineageIntegrationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putCloudDataLineageIntegration",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComposerEnvironmentConfigSoftwareConfigOutputReference) ResetAirflowConfigOverrides() {
	_jsii_.InvokeVoid(
		g,
		"resetAirflowConfigOverrides",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComposerEnvironmentConfigSoftwareConfigOutputReference) ResetCloudDataLineageIntegration() {
	_jsii_.InvokeVoid(
		g,
		"resetCloudDataLineageIntegration",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComposerEnvironmentConfigSoftwareConfigOutputReference) ResetEnvVariables() {
	_jsii_.InvokeVoid(
		g,
		"resetEnvVariables",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComposerEnvironmentConfigSoftwareConfigOutputReference) ResetImageVersion() {
	_jsii_.InvokeVoid(
		g,
		"resetImageVersion",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComposerEnvironmentConfigSoftwareConfigOutputReference) ResetPypiPackages() {
	_jsii_.InvokeVoid(
		g,
		"resetPypiPackages",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComposerEnvironmentConfigSoftwareConfigOutputReference) ResetPythonVersion() {
	_jsii_.InvokeVoid(
		g,
		"resetPythonVersion",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComposerEnvironmentConfigSoftwareConfigOutputReference) ResetSchedulerCount() {
	_jsii_.InvokeVoid(
		g,
		"resetSchedulerCount",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComposerEnvironmentConfigSoftwareConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleComposerEnvironmentConfigSoftwareConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

