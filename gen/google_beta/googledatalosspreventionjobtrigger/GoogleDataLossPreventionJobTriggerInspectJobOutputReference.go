package googledatalosspreventionjobtrigger

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googledatalosspreventionjobtrigger/internal"
)

type GoogleDataLossPreventionJobTriggerInspectJobOutputReference interface {
	cdktf.ComplexObject
	Actions() GoogleDataLossPreventionJobTriggerInspectJobActionsList
	ActionsInput() interface{}
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
	InspectConfig() GoogleDataLossPreventionJobTriggerInspectJobInspectConfigOutputReference
	InspectConfigInput() *GoogleDataLossPreventionJobTriggerInspectJobInspectConfig
	InspectTemplateName() *string
	SetInspectTemplateName(val *string)
	InspectTemplateNameInput() *string
	InternalValue() *GoogleDataLossPreventionJobTriggerInspectJob
	SetInternalValue(val *GoogleDataLossPreventionJobTriggerInspectJob)
	StorageConfig() GoogleDataLossPreventionJobTriggerInspectJobStorageConfigOutputReference
	StorageConfigInput() *GoogleDataLossPreventionJobTriggerInspectJobStorageConfig
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
	PutActions(value interface{})
	PutInspectConfig(value *GoogleDataLossPreventionJobTriggerInspectJobInspectConfig)
	PutStorageConfig(value *GoogleDataLossPreventionJobTriggerInspectJobStorageConfig)
	ResetActions()
	ResetInspectConfig()
	ResetInspectTemplateName()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleDataLossPreventionJobTriggerInspectJobOutputReference
type jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobOutputReference) Actions() GoogleDataLossPreventionJobTriggerInspectJobActionsList {
	var returns GoogleDataLossPreventionJobTriggerInspectJobActionsList
	_jsii_.Get(
		j,
		"actions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobOutputReference) ActionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"actionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobOutputReference) InspectConfig() GoogleDataLossPreventionJobTriggerInspectJobInspectConfigOutputReference {
	var returns GoogleDataLossPreventionJobTriggerInspectJobInspectConfigOutputReference
	_jsii_.Get(
		j,
		"inspectConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobOutputReference) InspectConfigInput() *GoogleDataLossPreventionJobTriggerInspectJobInspectConfig {
	var returns *GoogleDataLossPreventionJobTriggerInspectJobInspectConfig
	_jsii_.Get(
		j,
		"inspectConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobOutputReference) InspectTemplateName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inspectTemplateName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobOutputReference) InspectTemplateNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inspectTemplateNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobOutputReference) InternalValue() *GoogleDataLossPreventionJobTriggerInspectJob {
	var returns *GoogleDataLossPreventionJobTriggerInspectJob
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobOutputReference) StorageConfig() GoogleDataLossPreventionJobTriggerInspectJobStorageConfigOutputReference {
	var returns GoogleDataLossPreventionJobTriggerInspectJobStorageConfigOutputReference
	_jsii_.Get(
		j,
		"storageConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobOutputReference) StorageConfigInput() *GoogleDataLossPreventionJobTriggerInspectJobStorageConfig {
	var returns *GoogleDataLossPreventionJobTriggerInspectJobStorageConfig
	_jsii_.Get(
		j,
		"storageConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleDataLossPreventionJobTriggerInspectJobOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleDataLossPreventionJobTriggerInspectJobOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleDataLossPreventionJobTriggerInspectJobOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleDataLossPreventionJobTrigger.GoogleDataLossPreventionJobTriggerInspectJobOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleDataLossPreventionJobTriggerInspectJobOutputReference_Override(g GoogleDataLossPreventionJobTriggerInspectJobOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleDataLossPreventionJobTrigger.GoogleDataLossPreventionJobTriggerInspectJobOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobOutputReference)SetInspectTemplateName(val *string) {
	if err := j.validateSetInspectTemplateNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inspectTemplateName",
		val,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobOutputReference)SetInternalValue(val *GoogleDataLossPreventionJobTriggerInspectJob) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobOutputReference) PutActions(value interface{}) {
	if err := g.validatePutActionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putActions",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobOutputReference) PutInspectConfig(value *GoogleDataLossPreventionJobTriggerInspectJobInspectConfig) {
	if err := g.validatePutInspectConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putInspectConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobOutputReference) PutStorageConfig(value *GoogleDataLossPreventionJobTriggerInspectJobStorageConfig) {
	if err := g.validatePutStorageConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putStorageConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobOutputReference) ResetActions() {
	_jsii_.InvokeVoid(
		g,
		"resetActions",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobOutputReference) ResetInspectConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetInspectConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobOutputReference) ResetInspectTemplateName() {
	_jsii_.InvokeVoid(
		g,
		"resetInspectTemplateName",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

