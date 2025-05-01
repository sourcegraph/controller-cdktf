package googlestoragetransferjob

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlestoragetransferjob/internal"
)

type GoogleStorageTransferJobReplicationSpecOutputReference interface {
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
	GcsDataSink() GoogleStorageTransferJobReplicationSpecGcsDataSinkOutputReference
	GcsDataSinkInput() *GoogleStorageTransferJobReplicationSpecGcsDataSink
	GcsDataSource() GoogleStorageTransferJobReplicationSpecGcsDataSourceOutputReference
	GcsDataSourceInput() *GoogleStorageTransferJobReplicationSpecGcsDataSource
	InternalValue() *GoogleStorageTransferJobReplicationSpec
	SetInternalValue(val *GoogleStorageTransferJobReplicationSpec)
	ObjectConditions() GoogleStorageTransferJobReplicationSpecObjectConditionsOutputReference
	ObjectConditionsInput() *GoogleStorageTransferJobReplicationSpecObjectConditions
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TransferOptions() GoogleStorageTransferJobReplicationSpecTransferOptionsOutputReference
	TransferOptionsInput() *GoogleStorageTransferJobReplicationSpecTransferOptions
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
	PutGcsDataSink(value *GoogleStorageTransferJobReplicationSpecGcsDataSink)
	PutGcsDataSource(value *GoogleStorageTransferJobReplicationSpecGcsDataSource)
	PutObjectConditions(value *GoogleStorageTransferJobReplicationSpecObjectConditions)
	PutTransferOptions(value *GoogleStorageTransferJobReplicationSpecTransferOptions)
	ResetGcsDataSink()
	ResetGcsDataSource()
	ResetObjectConditions()
	ResetTransferOptions()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleStorageTransferJobReplicationSpecOutputReference
type jsiiProxy_GoogleStorageTransferJobReplicationSpecOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleStorageTransferJobReplicationSpecOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageTransferJobReplicationSpecOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageTransferJobReplicationSpecOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageTransferJobReplicationSpecOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageTransferJobReplicationSpecOutputReference) GcsDataSink() GoogleStorageTransferJobReplicationSpecGcsDataSinkOutputReference {
	var returns GoogleStorageTransferJobReplicationSpecGcsDataSinkOutputReference
	_jsii_.Get(
		j,
		"gcsDataSink",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageTransferJobReplicationSpecOutputReference) GcsDataSinkInput() *GoogleStorageTransferJobReplicationSpecGcsDataSink {
	var returns *GoogleStorageTransferJobReplicationSpecGcsDataSink
	_jsii_.Get(
		j,
		"gcsDataSinkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageTransferJobReplicationSpecOutputReference) GcsDataSource() GoogleStorageTransferJobReplicationSpecGcsDataSourceOutputReference {
	var returns GoogleStorageTransferJobReplicationSpecGcsDataSourceOutputReference
	_jsii_.Get(
		j,
		"gcsDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageTransferJobReplicationSpecOutputReference) GcsDataSourceInput() *GoogleStorageTransferJobReplicationSpecGcsDataSource {
	var returns *GoogleStorageTransferJobReplicationSpecGcsDataSource
	_jsii_.Get(
		j,
		"gcsDataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageTransferJobReplicationSpecOutputReference) InternalValue() *GoogleStorageTransferJobReplicationSpec {
	var returns *GoogleStorageTransferJobReplicationSpec
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageTransferJobReplicationSpecOutputReference) ObjectConditions() GoogleStorageTransferJobReplicationSpecObjectConditionsOutputReference {
	var returns GoogleStorageTransferJobReplicationSpecObjectConditionsOutputReference
	_jsii_.Get(
		j,
		"objectConditions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageTransferJobReplicationSpecOutputReference) ObjectConditionsInput() *GoogleStorageTransferJobReplicationSpecObjectConditions {
	var returns *GoogleStorageTransferJobReplicationSpecObjectConditions
	_jsii_.Get(
		j,
		"objectConditionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageTransferJobReplicationSpecOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageTransferJobReplicationSpecOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageTransferJobReplicationSpecOutputReference) TransferOptions() GoogleStorageTransferJobReplicationSpecTransferOptionsOutputReference {
	var returns GoogleStorageTransferJobReplicationSpecTransferOptionsOutputReference
	_jsii_.Get(
		j,
		"transferOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageTransferJobReplicationSpecOutputReference) TransferOptionsInput() *GoogleStorageTransferJobReplicationSpecTransferOptions {
	var returns *GoogleStorageTransferJobReplicationSpecTransferOptions
	_jsii_.Get(
		j,
		"transferOptionsInput",
		&returns,
	)
	return returns
}


func NewGoogleStorageTransferJobReplicationSpecOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleStorageTransferJobReplicationSpecOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleStorageTransferJobReplicationSpecOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleStorageTransferJobReplicationSpecOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleStorageTransferJob.GoogleStorageTransferJobReplicationSpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleStorageTransferJobReplicationSpecOutputReference_Override(g GoogleStorageTransferJobReplicationSpecOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleStorageTransferJob.GoogleStorageTransferJobReplicationSpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleStorageTransferJobReplicationSpecOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleStorageTransferJobReplicationSpecOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleStorageTransferJobReplicationSpecOutputReference)SetInternalValue(val *GoogleStorageTransferJobReplicationSpec) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleStorageTransferJobReplicationSpecOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleStorageTransferJobReplicationSpecOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleStorageTransferJobReplicationSpecOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleStorageTransferJobReplicationSpecOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleStorageTransferJobReplicationSpecOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleStorageTransferJobReplicationSpecOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleStorageTransferJobReplicationSpecOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleStorageTransferJobReplicationSpecOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleStorageTransferJobReplicationSpecOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleStorageTransferJobReplicationSpecOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleStorageTransferJobReplicationSpecOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleStorageTransferJobReplicationSpecOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleStorageTransferJobReplicationSpecOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleStorageTransferJobReplicationSpecOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleStorageTransferJobReplicationSpecOutputReference) PutGcsDataSink(value *GoogleStorageTransferJobReplicationSpecGcsDataSink) {
	if err := g.validatePutGcsDataSinkParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putGcsDataSink",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleStorageTransferJobReplicationSpecOutputReference) PutGcsDataSource(value *GoogleStorageTransferJobReplicationSpecGcsDataSource) {
	if err := g.validatePutGcsDataSourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putGcsDataSource",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleStorageTransferJobReplicationSpecOutputReference) PutObjectConditions(value *GoogleStorageTransferJobReplicationSpecObjectConditions) {
	if err := g.validatePutObjectConditionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putObjectConditions",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleStorageTransferJobReplicationSpecOutputReference) PutTransferOptions(value *GoogleStorageTransferJobReplicationSpecTransferOptions) {
	if err := g.validatePutTransferOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTransferOptions",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleStorageTransferJobReplicationSpecOutputReference) ResetGcsDataSink() {
	_jsii_.InvokeVoid(
		g,
		"resetGcsDataSink",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleStorageTransferJobReplicationSpecOutputReference) ResetGcsDataSource() {
	_jsii_.InvokeVoid(
		g,
		"resetGcsDataSource",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleStorageTransferJobReplicationSpecOutputReference) ResetObjectConditions() {
	_jsii_.InvokeVoid(
		g,
		"resetObjectConditions",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleStorageTransferJobReplicationSpecOutputReference) ResetTransferOptions() {
	_jsii_.InvokeVoid(
		g,
		"resetTransferOptions",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleStorageTransferJobReplicationSpecOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleStorageTransferJobReplicationSpecOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

