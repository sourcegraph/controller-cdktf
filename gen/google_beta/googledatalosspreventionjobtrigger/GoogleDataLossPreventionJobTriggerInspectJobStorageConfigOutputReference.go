package googledatalosspreventionjobtrigger

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googledatalosspreventionjobtrigger/internal"
)

type GoogleDataLossPreventionJobTriggerInspectJobStorageConfigOutputReference interface {
	cdktf.ComplexObject
	BigQueryOptions() GoogleDataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsOutputReference
	BigQueryOptionsInput() *GoogleDataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptions
	CloudStorageOptions() GoogleDataLossPreventionJobTriggerInspectJobStorageConfigCloudStorageOptionsOutputReference
	CloudStorageOptionsInput() *GoogleDataLossPreventionJobTriggerInspectJobStorageConfigCloudStorageOptions
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
	DatastoreOptions() GoogleDataLossPreventionJobTriggerInspectJobStorageConfigDatastoreOptionsOutputReference
	DatastoreOptionsInput() *GoogleDataLossPreventionJobTriggerInspectJobStorageConfigDatastoreOptions
	// Experimental.
	Fqn() *string
	HybridOptions() GoogleDataLossPreventionJobTriggerInspectJobStorageConfigHybridOptionsOutputReference
	HybridOptionsInput() *GoogleDataLossPreventionJobTriggerInspectJobStorageConfigHybridOptions
	InternalValue() *GoogleDataLossPreventionJobTriggerInspectJobStorageConfig
	SetInternalValue(val *GoogleDataLossPreventionJobTriggerInspectJobStorageConfig)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TimespanConfig() GoogleDataLossPreventionJobTriggerInspectJobStorageConfigTimespanConfigOutputReference
	TimespanConfigInput() *GoogleDataLossPreventionJobTriggerInspectJobStorageConfigTimespanConfig
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
	PutBigQueryOptions(value *GoogleDataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptions)
	PutCloudStorageOptions(value *GoogleDataLossPreventionJobTriggerInspectJobStorageConfigCloudStorageOptions)
	PutDatastoreOptions(value *GoogleDataLossPreventionJobTriggerInspectJobStorageConfigDatastoreOptions)
	PutHybridOptions(value *GoogleDataLossPreventionJobTriggerInspectJobStorageConfigHybridOptions)
	PutTimespanConfig(value *GoogleDataLossPreventionJobTriggerInspectJobStorageConfigTimespanConfig)
	ResetBigQueryOptions()
	ResetCloudStorageOptions()
	ResetDatastoreOptions()
	ResetHybridOptions()
	ResetTimespanConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleDataLossPreventionJobTriggerInspectJobStorageConfigOutputReference
type jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobStorageConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobStorageConfigOutputReference) BigQueryOptions() GoogleDataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsOutputReference {
	var returns GoogleDataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsOutputReference
	_jsii_.Get(
		j,
		"bigQueryOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobStorageConfigOutputReference) BigQueryOptionsInput() *GoogleDataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptions {
	var returns *GoogleDataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptions
	_jsii_.Get(
		j,
		"bigQueryOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobStorageConfigOutputReference) CloudStorageOptions() GoogleDataLossPreventionJobTriggerInspectJobStorageConfigCloudStorageOptionsOutputReference {
	var returns GoogleDataLossPreventionJobTriggerInspectJobStorageConfigCloudStorageOptionsOutputReference
	_jsii_.Get(
		j,
		"cloudStorageOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobStorageConfigOutputReference) CloudStorageOptionsInput() *GoogleDataLossPreventionJobTriggerInspectJobStorageConfigCloudStorageOptions {
	var returns *GoogleDataLossPreventionJobTriggerInspectJobStorageConfigCloudStorageOptions
	_jsii_.Get(
		j,
		"cloudStorageOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobStorageConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobStorageConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobStorageConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobStorageConfigOutputReference) DatastoreOptions() GoogleDataLossPreventionJobTriggerInspectJobStorageConfigDatastoreOptionsOutputReference {
	var returns GoogleDataLossPreventionJobTriggerInspectJobStorageConfigDatastoreOptionsOutputReference
	_jsii_.Get(
		j,
		"datastoreOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobStorageConfigOutputReference) DatastoreOptionsInput() *GoogleDataLossPreventionJobTriggerInspectJobStorageConfigDatastoreOptions {
	var returns *GoogleDataLossPreventionJobTriggerInspectJobStorageConfigDatastoreOptions
	_jsii_.Get(
		j,
		"datastoreOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobStorageConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobStorageConfigOutputReference) HybridOptions() GoogleDataLossPreventionJobTriggerInspectJobStorageConfigHybridOptionsOutputReference {
	var returns GoogleDataLossPreventionJobTriggerInspectJobStorageConfigHybridOptionsOutputReference
	_jsii_.Get(
		j,
		"hybridOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobStorageConfigOutputReference) HybridOptionsInput() *GoogleDataLossPreventionJobTriggerInspectJobStorageConfigHybridOptions {
	var returns *GoogleDataLossPreventionJobTriggerInspectJobStorageConfigHybridOptions
	_jsii_.Get(
		j,
		"hybridOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobStorageConfigOutputReference) InternalValue() *GoogleDataLossPreventionJobTriggerInspectJobStorageConfig {
	var returns *GoogleDataLossPreventionJobTriggerInspectJobStorageConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobStorageConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobStorageConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobStorageConfigOutputReference) TimespanConfig() GoogleDataLossPreventionJobTriggerInspectJobStorageConfigTimespanConfigOutputReference {
	var returns GoogleDataLossPreventionJobTriggerInspectJobStorageConfigTimespanConfigOutputReference
	_jsii_.Get(
		j,
		"timespanConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobStorageConfigOutputReference) TimespanConfigInput() *GoogleDataLossPreventionJobTriggerInspectJobStorageConfigTimespanConfig {
	var returns *GoogleDataLossPreventionJobTriggerInspectJobStorageConfigTimespanConfig
	_jsii_.Get(
		j,
		"timespanConfigInput",
		&returns,
	)
	return returns
}


func NewGoogleDataLossPreventionJobTriggerInspectJobStorageConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleDataLossPreventionJobTriggerInspectJobStorageConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleDataLossPreventionJobTriggerInspectJobStorageConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobStorageConfigOutputReference{}

	_jsii_.Create(
		"google-beta.googleDataLossPreventionJobTrigger.GoogleDataLossPreventionJobTriggerInspectJobStorageConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleDataLossPreventionJobTriggerInspectJobStorageConfigOutputReference_Override(g GoogleDataLossPreventionJobTriggerInspectJobStorageConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"google-beta.googleDataLossPreventionJobTrigger.GoogleDataLossPreventionJobTriggerInspectJobStorageConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobStorageConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobStorageConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobStorageConfigOutputReference)SetInternalValue(val *GoogleDataLossPreventionJobTriggerInspectJobStorageConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobStorageConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobStorageConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobStorageConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobStorageConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobStorageConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobStorageConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobStorageConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobStorageConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobStorageConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobStorageConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobStorageConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobStorageConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobStorageConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobStorageConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobStorageConfigOutputReference) PutBigQueryOptions(value *GoogleDataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptions) {
	if err := g.validatePutBigQueryOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putBigQueryOptions",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobStorageConfigOutputReference) PutCloudStorageOptions(value *GoogleDataLossPreventionJobTriggerInspectJobStorageConfigCloudStorageOptions) {
	if err := g.validatePutCloudStorageOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putCloudStorageOptions",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobStorageConfigOutputReference) PutDatastoreOptions(value *GoogleDataLossPreventionJobTriggerInspectJobStorageConfigDatastoreOptions) {
	if err := g.validatePutDatastoreOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDatastoreOptions",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobStorageConfigOutputReference) PutHybridOptions(value *GoogleDataLossPreventionJobTriggerInspectJobStorageConfigHybridOptions) {
	if err := g.validatePutHybridOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putHybridOptions",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobStorageConfigOutputReference) PutTimespanConfig(value *GoogleDataLossPreventionJobTriggerInspectJobStorageConfigTimespanConfig) {
	if err := g.validatePutTimespanConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTimespanConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobStorageConfigOutputReference) ResetBigQueryOptions() {
	_jsii_.InvokeVoid(
		g,
		"resetBigQueryOptions",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobStorageConfigOutputReference) ResetCloudStorageOptions() {
	_jsii_.InvokeVoid(
		g,
		"resetCloudStorageOptions",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobStorageConfigOutputReference) ResetDatastoreOptions() {
	_jsii_.InvokeVoid(
		g,
		"resetDatastoreOptions",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobStorageConfigOutputReference) ResetHybridOptions() {
	_jsii_.InvokeVoid(
		g,
		"resetHybridOptions",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobStorageConfigOutputReference) ResetTimespanConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetTimespanConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobStorageConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobStorageConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

