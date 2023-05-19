package googledataprocworkflowtemplate

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googledataprocworkflowtemplate/internal"
)

type GoogleDataprocWorkflowTemplateJobsSparkSqlJobOutputReference interface {
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
	InternalValue() *GoogleDataprocWorkflowTemplateJobsSparkSqlJob
	SetInternalValue(val *GoogleDataprocWorkflowTemplateJobsSparkSqlJob)
	JarFileUris() *[]*string
	SetJarFileUris(val *[]*string)
	JarFileUrisInput() *[]*string
	LoggingConfig() GoogleDataprocWorkflowTemplateJobsSparkSqlJobLoggingConfigOutputReference
	LoggingConfigInput() *GoogleDataprocWorkflowTemplateJobsSparkSqlJobLoggingConfig
	Properties() *map[string]*string
	SetProperties(val *map[string]*string)
	PropertiesInput() *map[string]*string
	QueryFileUri() *string
	SetQueryFileUri(val *string)
	QueryFileUriInput() *string
	QueryList() GoogleDataprocWorkflowTemplateJobsSparkSqlJobQueryListOutputReference
	QueryListInput() *GoogleDataprocWorkflowTemplateJobsSparkSqlJobQueryList
	ScriptVariables() *map[string]*string
	SetScriptVariables(val *map[string]*string)
	ScriptVariablesInput() *map[string]*string
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
	PutLoggingConfig(value *GoogleDataprocWorkflowTemplateJobsSparkSqlJobLoggingConfig)
	PutQueryList(value *GoogleDataprocWorkflowTemplateJobsSparkSqlJobQueryList)
	ResetJarFileUris()
	ResetLoggingConfig()
	ResetProperties()
	ResetQueryFileUri()
	ResetQueryList()
	ResetScriptVariables()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleDataprocWorkflowTemplateJobsSparkSqlJobOutputReference
type jsiiProxy_GoogleDataprocWorkflowTemplateJobsSparkSqlJobOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplateJobsSparkSqlJobOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplateJobsSparkSqlJobOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplateJobsSparkSqlJobOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplateJobsSparkSqlJobOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplateJobsSparkSqlJobOutputReference) InternalValue() *GoogleDataprocWorkflowTemplateJobsSparkSqlJob {
	var returns *GoogleDataprocWorkflowTemplateJobsSparkSqlJob
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplateJobsSparkSqlJobOutputReference) JarFileUris() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"jarFileUris",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplateJobsSparkSqlJobOutputReference) JarFileUrisInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"jarFileUrisInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplateJobsSparkSqlJobOutputReference) LoggingConfig() GoogleDataprocWorkflowTemplateJobsSparkSqlJobLoggingConfigOutputReference {
	var returns GoogleDataprocWorkflowTemplateJobsSparkSqlJobLoggingConfigOutputReference
	_jsii_.Get(
		j,
		"loggingConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplateJobsSparkSqlJobOutputReference) LoggingConfigInput() *GoogleDataprocWorkflowTemplateJobsSparkSqlJobLoggingConfig {
	var returns *GoogleDataprocWorkflowTemplateJobsSparkSqlJobLoggingConfig
	_jsii_.Get(
		j,
		"loggingConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplateJobsSparkSqlJobOutputReference) Properties() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"properties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplateJobsSparkSqlJobOutputReference) PropertiesInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"propertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplateJobsSparkSqlJobOutputReference) QueryFileUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queryFileUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplateJobsSparkSqlJobOutputReference) QueryFileUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queryFileUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplateJobsSparkSqlJobOutputReference) QueryList() GoogleDataprocWorkflowTemplateJobsSparkSqlJobQueryListOutputReference {
	var returns GoogleDataprocWorkflowTemplateJobsSparkSqlJobQueryListOutputReference
	_jsii_.Get(
		j,
		"queryList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplateJobsSparkSqlJobOutputReference) QueryListInput() *GoogleDataprocWorkflowTemplateJobsSparkSqlJobQueryList {
	var returns *GoogleDataprocWorkflowTemplateJobsSparkSqlJobQueryList
	_jsii_.Get(
		j,
		"queryListInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplateJobsSparkSqlJobOutputReference) ScriptVariables() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"scriptVariables",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplateJobsSparkSqlJobOutputReference) ScriptVariablesInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"scriptVariablesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplateJobsSparkSqlJobOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplateJobsSparkSqlJobOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleDataprocWorkflowTemplateJobsSparkSqlJobOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleDataprocWorkflowTemplateJobsSparkSqlJobOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleDataprocWorkflowTemplateJobsSparkSqlJobOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleDataprocWorkflowTemplateJobsSparkSqlJobOutputReference{}

	_jsii_.Create(
		"google-beta.googleDataprocWorkflowTemplate.GoogleDataprocWorkflowTemplateJobsSparkSqlJobOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleDataprocWorkflowTemplateJobsSparkSqlJobOutputReference_Override(g GoogleDataprocWorkflowTemplateJobsSparkSqlJobOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"google-beta.googleDataprocWorkflowTemplate.GoogleDataprocWorkflowTemplateJobsSparkSqlJobOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplateJobsSparkSqlJobOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplateJobsSparkSqlJobOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplateJobsSparkSqlJobOutputReference)SetInternalValue(val *GoogleDataprocWorkflowTemplateJobsSparkSqlJob) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplateJobsSparkSqlJobOutputReference)SetJarFileUris(val *[]*string) {
	if err := j.validateSetJarFileUrisParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jarFileUris",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplateJobsSparkSqlJobOutputReference)SetProperties(val *map[string]*string) {
	if err := j.validateSetPropertiesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"properties",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplateJobsSparkSqlJobOutputReference)SetQueryFileUri(val *string) {
	if err := j.validateSetQueryFileUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"queryFileUri",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplateJobsSparkSqlJobOutputReference)SetScriptVariables(val *map[string]*string) {
	if err := j.validateSetScriptVariablesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scriptVariables",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplateJobsSparkSqlJobOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplateJobsSparkSqlJobOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleDataprocWorkflowTemplateJobsSparkSqlJobOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataprocWorkflowTemplateJobsSparkSqlJobOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleDataprocWorkflowTemplateJobsSparkSqlJobOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDataprocWorkflowTemplateJobsSparkSqlJobOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleDataprocWorkflowTemplateJobsSparkSqlJobOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleDataprocWorkflowTemplateJobsSparkSqlJobOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleDataprocWorkflowTemplateJobsSparkSqlJobOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleDataprocWorkflowTemplateJobsSparkSqlJobOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleDataprocWorkflowTemplateJobsSparkSqlJobOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleDataprocWorkflowTemplateJobsSparkSqlJobOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleDataprocWorkflowTemplateJobsSparkSqlJobOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataprocWorkflowTemplateJobsSparkSqlJobOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDataprocWorkflowTemplateJobsSparkSqlJobOutputReference) PutLoggingConfig(value *GoogleDataprocWorkflowTemplateJobsSparkSqlJobLoggingConfig) {
	if err := g.validatePutLoggingConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putLoggingConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataprocWorkflowTemplateJobsSparkSqlJobOutputReference) PutQueryList(value *GoogleDataprocWorkflowTemplateJobsSparkSqlJobQueryList) {
	if err := g.validatePutQueryListParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putQueryList",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataprocWorkflowTemplateJobsSparkSqlJobOutputReference) ResetJarFileUris() {
	_jsii_.InvokeVoid(
		g,
		"resetJarFileUris",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocWorkflowTemplateJobsSparkSqlJobOutputReference) ResetLoggingConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetLoggingConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocWorkflowTemplateJobsSparkSqlJobOutputReference) ResetProperties() {
	_jsii_.InvokeVoid(
		g,
		"resetProperties",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocWorkflowTemplateJobsSparkSqlJobOutputReference) ResetQueryFileUri() {
	_jsii_.InvokeVoid(
		g,
		"resetQueryFileUri",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocWorkflowTemplateJobsSparkSqlJobOutputReference) ResetQueryList() {
	_jsii_.InvokeVoid(
		g,
		"resetQueryList",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocWorkflowTemplateJobsSparkSqlJobOutputReference) ResetScriptVariables() {
	_jsii_.InvokeVoid(
		g,
		"resetScriptVariables",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocWorkflowTemplateJobsSparkSqlJobOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleDataprocWorkflowTemplateJobsSparkSqlJobOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

