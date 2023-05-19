package googledatastreamstream

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googledatastreamstream/internal"
)

type GoogleDatastreamStreamSourceConfigOracleSourceConfigOutputReference interface {
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
	DropLargeObjects() GoogleDatastreamStreamSourceConfigOracleSourceConfigDropLargeObjectsOutputReference
	DropLargeObjectsInput() *GoogleDatastreamStreamSourceConfigOracleSourceConfigDropLargeObjects
	ExcludeObjects() GoogleDatastreamStreamSourceConfigOracleSourceConfigExcludeObjectsOutputReference
	ExcludeObjectsInput() *GoogleDatastreamStreamSourceConfigOracleSourceConfigExcludeObjects
	// Experimental.
	Fqn() *string
	IncludeObjects() GoogleDatastreamStreamSourceConfigOracleSourceConfigIncludeObjectsOutputReference
	IncludeObjectsInput() *GoogleDatastreamStreamSourceConfigOracleSourceConfigIncludeObjects
	InternalValue() *GoogleDatastreamStreamSourceConfigOracleSourceConfig
	SetInternalValue(val *GoogleDatastreamStreamSourceConfigOracleSourceConfig)
	MaxConcurrentBackfillTasks() *float64
	SetMaxConcurrentBackfillTasks(val *float64)
	MaxConcurrentBackfillTasksInput() *float64
	MaxConcurrentCdcTasks() *float64
	SetMaxConcurrentCdcTasks(val *float64)
	MaxConcurrentCdcTasksInput() *float64
	StreamLargeObjects() GoogleDatastreamStreamSourceConfigOracleSourceConfigStreamLargeObjectsOutputReference
	StreamLargeObjectsInput() *GoogleDatastreamStreamSourceConfigOracleSourceConfigStreamLargeObjects
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
	PutDropLargeObjects(value *GoogleDatastreamStreamSourceConfigOracleSourceConfigDropLargeObjects)
	PutExcludeObjects(value *GoogleDatastreamStreamSourceConfigOracleSourceConfigExcludeObjects)
	PutIncludeObjects(value *GoogleDatastreamStreamSourceConfigOracleSourceConfigIncludeObjects)
	PutStreamLargeObjects(value *GoogleDatastreamStreamSourceConfigOracleSourceConfigStreamLargeObjects)
	ResetDropLargeObjects()
	ResetExcludeObjects()
	ResetIncludeObjects()
	ResetMaxConcurrentBackfillTasks()
	ResetMaxConcurrentCdcTasks()
	ResetStreamLargeObjects()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleDatastreamStreamSourceConfigOracleSourceConfigOutputReference
type jsiiProxy_GoogleDatastreamStreamSourceConfigOracleSourceConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleDatastreamStreamSourceConfigOracleSourceConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamStreamSourceConfigOracleSourceConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamStreamSourceConfigOracleSourceConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamStreamSourceConfigOracleSourceConfigOutputReference) DropLargeObjects() GoogleDatastreamStreamSourceConfigOracleSourceConfigDropLargeObjectsOutputReference {
	var returns GoogleDatastreamStreamSourceConfigOracleSourceConfigDropLargeObjectsOutputReference
	_jsii_.Get(
		j,
		"dropLargeObjects",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamStreamSourceConfigOracleSourceConfigOutputReference) DropLargeObjectsInput() *GoogleDatastreamStreamSourceConfigOracleSourceConfigDropLargeObjects {
	var returns *GoogleDatastreamStreamSourceConfigOracleSourceConfigDropLargeObjects
	_jsii_.Get(
		j,
		"dropLargeObjectsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamStreamSourceConfigOracleSourceConfigOutputReference) ExcludeObjects() GoogleDatastreamStreamSourceConfigOracleSourceConfigExcludeObjectsOutputReference {
	var returns GoogleDatastreamStreamSourceConfigOracleSourceConfigExcludeObjectsOutputReference
	_jsii_.Get(
		j,
		"excludeObjects",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamStreamSourceConfigOracleSourceConfigOutputReference) ExcludeObjectsInput() *GoogleDatastreamStreamSourceConfigOracleSourceConfigExcludeObjects {
	var returns *GoogleDatastreamStreamSourceConfigOracleSourceConfigExcludeObjects
	_jsii_.Get(
		j,
		"excludeObjectsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamStreamSourceConfigOracleSourceConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamStreamSourceConfigOracleSourceConfigOutputReference) IncludeObjects() GoogleDatastreamStreamSourceConfigOracleSourceConfigIncludeObjectsOutputReference {
	var returns GoogleDatastreamStreamSourceConfigOracleSourceConfigIncludeObjectsOutputReference
	_jsii_.Get(
		j,
		"includeObjects",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamStreamSourceConfigOracleSourceConfigOutputReference) IncludeObjectsInput() *GoogleDatastreamStreamSourceConfigOracleSourceConfigIncludeObjects {
	var returns *GoogleDatastreamStreamSourceConfigOracleSourceConfigIncludeObjects
	_jsii_.Get(
		j,
		"includeObjectsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamStreamSourceConfigOracleSourceConfigOutputReference) InternalValue() *GoogleDatastreamStreamSourceConfigOracleSourceConfig {
	var returns *GoogleDatastreamStreamSourceConfigOracleSourceConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamStreamSourceConfigOracleSourceConfigOutputReference) MaxConcurrentBackfillTasks() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxConcurrentBackfillTasks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamStreamSourceConfigOracleSourceConfigOutputReference) MaxConcurrentBackfillTasksInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxConcurrentBackfillTasksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamStreamSourceConfigOracleSourceConfigOutputReference) MaxConcurrentCdcTasks() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxConcurrentCdcTasks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamStreamSourceConfigOracleSourceConfigOutputReference) MaxConcurrentCdcTasksInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxConcurrentCdcTasksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamStreamSourceConfigOracleSourceConfigOutputReference) StreamLargeObjects() GoogleDatastreamStreamSourceConfigOracleSourceConfigStreamLargeObjectsOutputReference {
	var returns GoogleDatastreamStreamSourceConfigOracleSourceConfigStreamLargeObjectsOutputReference
	_jsii_.Get(
		j,
		"streamLargeObjects",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamStreamSourceConfigOracleSourceConfigOutputReference) StreamLargeObjectsInput() *GoogleDatastreamStreamSourceConfigOracleSourceConfigStreamLargeObjects {
	var returns *GoogleDatastreamStreamSourceConfigOracleSourceConfigStreamLargeObjects
	_jsii_.Get(
		j,
		"streamLargeObjectsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamStreamSourceConfigOracleSourceConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamStreamSourceConfigOracleSourceConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleDatastreamStreamSourceConfigOracleSourceConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleDatastreamStreamSourceConfigOracleSourceConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleDatastreamStreamSourceConfigOracleSourceConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleDatastreamStreamSourceConfigOracleSourceConfigOutputReference{}

	_jsii_.Create(
		"google-beta.googleDatastreamStream.GoogleDatastreamStreamSourceConfigOracleSourceConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleDatastreamStreamSourceConfigOracleSourceConfigOutputReference_Override(g GoogleDatastreamStreamSourceConfigOracleSourceConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"google-beta.googleDatastreamStream.GoogleDatastreamStreamSourceConfigOracleSourceConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleDatastreamStreamSourceConfigOracleSourceConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleDatastreamStreamSourceConfigOracleSourceConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleDatastreamStreamSourceConfigOracleSourceConfigOutputReference)SetInternalValue(val *GoogleDatastreamStreamSourceConfigOracleSourceConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleDatastreamStreamSourceConfigOracleSourceConfigOutputReference)SetMaxConcurrentBackfillTasks(val *float64) {
	if err := j.validateSetMaxConcurrentBackfillTasksParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxConcurrentBackfillTasks",
		val,
	)
}

func (j *jsiiProxy_GoogleDatastreamStreamSourceConfigOracleSourceConfigOutputReference)SetMaxConcurrentCdcTasks(val *float64) {
	if err := j.validateSetMaxConcurrentCdcTasksParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxConcurrentCdcTasks",
		val,
	)
}

func (j *jsiiProxy_GoogleDatastreamStreamSourceConfigOracleSourceConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleDatastreamStreamSourceConfigOracleSourceConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleDatastreamStreamSourceConfigOracleSourceConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDatastreamStreamSourceConfigOracleSourceConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleDatastreamStreamSourceConfigOracleSourceConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDatastreamStreamSourceConfigOracleSourceConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleDatastreamStreamSourceConfigOracleSourceConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleDatastreamStreamSourceConfigOracleSourceConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleDatastreamStreamSourceConfigOracleSourceConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleDatastreamStreamSourceConfigOracleSourceConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleDatastreamStreamSourceConfigOracleSourceConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleDatastreamStreamSourceConfigOracleSourceConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleDatastreamStreamSourceConfigOracleSourceConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDatastreamStreamSourceConfigOracleSourceConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDatastreamStreamSourceConfigOracleSourceConfigOutputReference) PutDropLargeObjects(value *GoogleDatastreamStreamSourceConfigOracleSourceConfigDropLargeObjects) {
	if err := g.validatePutDropLargeObjectsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDropLargeObjects",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDatastreamStreamSourceConfigOracleSourceConfigOutputReference) PutExcludeObjects(value *GoogleDatastreamStreamSourceConfigOracleSourceConfigExcludeObjects) {
	if err := g.validatePutExcludeObjectsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putExcludeObjects",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDatastreamStreamSourceConfigOracleSourceConfigOutputReference) PutIncludeObjects(value *GoogleDatastreamStreamSourceConfigOracleSourceConfigIncludeObjects) {
	if err := g.validatePutIncludeObjectsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putIncludeObjects",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDatastreamStreamSourceConfigOracleSourceConfigOutputReference) PutStreamLargeObjects(value *GoogleDatastreamStreamSourceConfigOracleSourceConfigStreamLargeObjects) {
	if err := g.validatePutStreamLargeObjectsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putStreamLargeObjects",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDatastreamStreamSourceConfigOracleSourceConfigOutputReference) ResetDropLargeObjects() {
	_jsii_.InvokeVoid(
		g,
		"resetDropLargeObjects",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDatastreamStreamSourceConfigOracleSourceConfigOutputReference) ResetExcludeObjects() {
	_jsii_.InvokeVoid(
		g,
		"resetExcludeObjects",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDatastreamStreamSourceConfigOracleSourceConfigOutputReference) ResetIncludeObjects() {
	_jsii_.InvokeVoid(
		g,
		"resetIncludeObjects",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDatastreamStreamSourceConfigOracleSourceConfigOutputReference) ResetMaxConcurrentBackfillTasks() {
	_jsii_.InvokeVoid(
		g,
		"resetMaxConcurrentBackfillTasks",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDatastreamStreamSourceConfigOracleSourceConfigOutputReference) ResetMaxConcurrentCdcTasks() {
	_jsii_.InvokeVoid(
		g,
		"resetMaxConcurrentCdcTasks",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDatastreamStreamSourceConfigOracleSourceConfigOutputReference) ResetStreamLargeObjects() {
	_jsii_.InvokeVoid(
		g,
		"resetStreamLargeObjects",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDatastreamStreamSourceConfigOracleSourceConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleDatastreamStreamSourceConfigOracleSourceConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

