package googlecolabschedule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecolabschedule/internal"
)

type GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference interface {
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
	DataformRepositorySource() GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobDataformRepositorySourceOutputReference
	DataformRepositorySourceInput() *GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobDataformRepositorySource
	DisplayName() *string
	SetDisplayName(val *string)
	DisplayNameInput() *string
	ExecutionTimeout() *string
	SetExecutionTimeout(val *string)
	ExecutionTimeoutInput() *string
	ExecutionUser() *string
	SetExecutionUser(val *string)
	ExecutionUserInput() *string
	// Experimental.
	Fqn() *string
	GcsNotebookSource() GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobGcsNotebookSourceOutputReference
	GcsNotebookSourceInput() *GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobGcsNotebookSource
	GcsOutputUri() *string
	SetGcsOutputUri(val *string)
	GcsOutputUriInput() *string
	InternalValue() *GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJob
	SetInternalValue(val *GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJob)
	NotebookRuntimeTemplateResourceName() *string
	SetNotebookRuntimeTemplateResourceName(val *string)
	NotebookRuntimeTemplateResourceNameInput() *string
	ServiceAccount() *string
	SetServiceAccount(val *string)
	ServiceAccountInput() *string
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
	PutDataformRepositorySource(value *GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobDataformRepositorySource)
	PutGcsNotebookSource(value *GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobGcsNotebookSource)
	ResetDataformRepositorySource()
	ResetExecutionTimeout()
	ResetExecutionUser()
	ResetGcsNotebookSource()
	ResetServiceAccount()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference
type jsiiProxy_GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) DataformRepositorySource() GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobDataformRepositorySourceOutputReference {
	var returns GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobDataformRepositorySourceOutputReference
	_jsii_.Get(
		j,
		"dataformRepositorySource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) DataformRepositorySourceInput() *GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobDataformRepositorySource {
	var returns *GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobDataformRepositorySource
	_jsii_.Get(
		j,
		"dataformRepositorySourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) ExecutionTimeout() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) ExecutionTimeoutInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) ExecutionUser() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionUser",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) ExecutionUserInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionUserInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) GcsNotebookSource() GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobGcsNotebookSourceOutputReference {
	var returns GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobGcsNotebookSourceOutputReference
	_jsii_.Get(
		j,
		"gcsNotebookSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) GcsNotebookSourceInput() *GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobGcsNotebookSource {
	var returns *GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobGcsNotebookSource
	_jsii_.Get(
		j,
		"gcsNotebookSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) GcsOutputUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gcsOutputUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) GcsOutputUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gcsOutputUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) InternalValue() *GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJob {
	var returns *GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJob
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) NotebookRuntimeTemplateResourceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notebookRuntimeTemplateResourceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) NotebookRuntimeTemplateResourceNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notebookRuntimeTemplateResourceNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) ServiceAccount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) ServiceAccountInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleColabSchedule.GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference_Override(g GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleColabSchedule.GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference)SetDisplayName(val *string) {
	if err := j.validateSetDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference)SetExecutionTimeout(val *string) {
	if err := j.validateSetExecutionTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"executionTimeout",
		val,
	)
}

func (j *jsiiProxy_GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference)SetExecutionUser(val *string) {
	if err := j.validateSetExecutionUserParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"executionUser",
		val,
	)
}

func (j *jsiiProxy_GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference)SetGcsOutputUri(val *string) {
	if err := j.validateSetGcsOutputUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gcsOutputUri",
		val,
	)
}

func (j *jsiiProxy_GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference)SetInternalValue(val *GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJob) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference)SetNotebookRuntimeTemplateResourceName(val *string) {
	if err := j.validateSetNotebookRuntimeTemplateResourceNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"notebookRuntimeTemplateResourceName",
		val,
	)
}

func (j *jsiiProxy_GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference)SetServiceAccount(val *string) {
	if err := j.validateSetServiceAccountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceAccount",
		val,
	)
}

func (j *jsiiProxy_GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) PutDataformRepositorySource(value *GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobDataformRepositorySource) {
	if err := g.validatePutDataformRepositorySourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDataformRepositorySource",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) PutGcsNotebookSource(value *GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobGcsNotebookSource) {
	if err := g.validatePutGcsNotebookSourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putGcsNotebookSource",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) ResetDataformRepositorySource() {
	_jsii_.InvokeVoid(
		g,
		"resetDataformRepositorySource",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) ResetExecutionTimeout() {
	_jsii_.InvokeVoid(
		g,
		"resetExecutionTimeout",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) ResetExecutionUser() {
	_jsii_.InvokeVoid(
		g,
		"resetExecutionUser",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) ResetGcsNotebookSource() {
	_jsii_.InvokeVoid(
		g,
		"resetGcsNotebookSource",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) ResetServiceAccount() {
	_jsii_.InvokeVoid(
		g,
		"resetServiceAccount",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

