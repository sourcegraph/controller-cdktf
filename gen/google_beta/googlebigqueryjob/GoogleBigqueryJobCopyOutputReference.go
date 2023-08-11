package googlebigqueryjob

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlebigqueryjob/internal"
)

type GoogleBigqueryJobCopyOutputReference interface {
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
	CreateDisposition() *string
	SetCreateDisposition(val *string)
	CreateDispositionInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DestinationEncryptionConfiguration() GoogleBigqueryJobCopyDestinationEncryptionConfigurationOutputReference
	DestinationEncryptionConfigurationInput() *GoogleBigqueryJobCopyDestinationEncryptionConfiguration
	DestinationTable() GoogleBigqueryJobCopyDestinationTableOutputReference
	DestinationTableInput() *GoogleBigqueryJobCopyDestinationTable
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleBigqueryJobCopy
	SetInternalValue(val *GoogleBigqueryJobCopy)
	SourceTables() GoogleBigqueryJobCopySourceTablesList
	SourceTablesInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	WriteDisposition() *string
	SetWriteDisposition(val *string)
	WriteDispositionInput() *string
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
	PutDestinationEncryptionConfiguration(value *GoogleBigqueryJobCopyDestinationEncryptionConfiguration)
	PutDestinationTable(value *GoogleBigqueryJobCopyDestinationTable)
	PutSourceTables(value interface{})
	ResetCreateDisposition()
	ResetDestinationEncryptionConfiguration()
	ResetDestinationTable()
	ResetWriteDisposition()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleBigqueryJobCopyOutputReference
type jsiiProxy_GoogleBigqueryJobCopyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleBigqueryJobCopyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryJobCopyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryJobCopyOutputReference) CreateDisposition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createDisposition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryJobCopyOutputReference) CreateDispositionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createDispositionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryJobCopyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryJobCopyOutputReference) DestinationEncryptionConfiguration() GoogleBigqueryJobCopyDestinationEncryptionConfigurationOutputReference {
	var returns GoogleBigqueryJobCopyDestinationEncryptionConfigurationOutputReference
	_jsii_.Get(
		j,
		"destinationEncryptionConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryJobCopyOutputReference) DestinationEncryptionConfigurationInput() *GoogleBigqueryJobCopyDestinationEncryptionConfiguration {
	var returns *GoogleBigqueryJobCopyDestinationEncryptionConfiguration
	_jsii_.Get(
		j,
		"destinationEncryptionConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryJobCopyOutputReference) DestinationTable() GoogleBigqueryJobCopyDestinationTableOutputReference {
	var returns GoogleBigqueryJobCopyDestinationTableOutputReference
	_jsii_.Get(
		j,
		"destinationTable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryJobCopyOutputReference) DestinationTableInput() *GoogleBigqueryJobCopyDestinationTable {
	var returns *GoogleBigqueryJobCopyDestinationTable
	_jsii_.Get(
		j,
		"destinationTableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryJobCopyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryJobCopyOutputReference) InternalValue() *GoogleBigqueryJobCopy {
	var returns *GoogleBigqueryJobCopy
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryJobCopyOutputReference) SourceTables() GoogleBigqueryJobCopySourceTablesList {
	var returns GoogleBigqueryJobCopySourceTablesList
	_jsii_.Get(
		j,
		"sourceTables",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryJobCopyOutputReference) SourceTablesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sourceTablesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryJobCopyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryJobCopyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryJobCopyOutputReference) WriteDisposition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"writeDisposition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryJobCopyOutputReference) WriteDispositionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"writeDispositionInput",
		&returns,
	)
	return returns
}


func NewGoogleBigqueryJobCopyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleBigqueryJobCopyOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleBigqueryJobCopyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleBigqueryJobCopyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleBigqueryJob.GoogleBigqueryJobCopyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleBigqueryJobCopyOutputReference_Override(g GoogleBigqueryJobCopyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleBigqueryJob.GoogleBigqueryJobCopyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleBigqueryJobCopyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryJobCopyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryJobCopyOutputReference)SetCreateDisposition(val *string) {
	if err := j.validateSetCreateDispositionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createDisposition",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryJobCopyOutputReference)SetInternalValue(val *GoogleBigqueryJobCopy) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryJobCopyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryJobCopyOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryJobCopyOutputReference)SetWriteDisposition(val *string) {
	if err := j.validateSetWriteDispositionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"writeDisposition",
		val,
	)
}

func (g *jsiiProxy_GoogleBigqueryJobCopyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBigqueryJobCopyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleBigqueryJobCopyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleBigqueryJobCopyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleBigqueryJobCopyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleBigqueryJobCopyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleBigqueryJobCopyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleBigqueryJobCopyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleBigqueryJobCopyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleBigqueryJobCopyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleBigqueryJobCopyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBigqueryJobCopyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleBigqueryJobCopyOutputReference) PutDestinationEncryptionConfiguration(value *GoogleBigqueryJobCopyDestinationEncryptionConfiguration) {
	if err := g.validatePutDestinationEncryptionConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDestinationEncryptionConfiguration",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleBigqueryJobCopyOutputReference) PutDestinationTable(value *GoogleBigqueryJobCopyDestinationTable) {
	if err := g.validatePutDestinationTableParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDestinationTable",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleBigqueryJobCopyOutputReference) PutSourceTables(value interface{}) {
	if err := g.validatePutSourceTablesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSourceTables",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleBigqueryJobCopyOutputReference) ResetCreateDisposition() {
	_jsii_.InvokeVoid(
		g,
		"resetCreateDisposition",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryJobCopyOutputReference) ResetDestinationEncryptionConfiguration() {
	_jsii_.InvokeVoid(
		g,
		"resetDestinationEncryptionConfiguration",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryJobCopyOutputReference) ResetDestinationTable() {
	_jsii_.InvokeVoid(
		g,
		"resetDestinationTable",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryJobCopyOutputReference) ResetWriteDisposition() {
	_jsii_.InvokeVoid(
		g,
		"resetWriteDisposition",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryJobCopyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleBigqueryJobCopyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

