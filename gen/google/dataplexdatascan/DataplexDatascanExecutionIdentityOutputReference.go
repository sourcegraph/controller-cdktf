package dataplexdatascan

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google/dataplexdatascan/internal"
)

type DataplexDatascanExecutionIdentityOutputReference interface {
	cdktn.ComplexObject
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
	DataplexServiceAgent() DataplexDatascanExecutionIdentityDataplexServiceAgentOutputReference
	DataplexServiceAgentInput() *DataplexDatascanExecutionIdentityDataplexServiceAgent
	// Experimental.
	Fqn() *string
	InternalValue() *DataplexDatascanExecutionIdentity
	SetInternalValue(val *DataplexDatascanExecutionIdentity)
	ServiceAccount() DataplexDatascanExecutionIdentityServiceAccountOutputReference
	ServiceAccountInput() *DataplexDatascanExecutionIdentityServiceAccount
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	UserCredential() DataplexDatascanExecutionIdentityUserCredentialOutputReference
	UserCredentialInput() *DataplexDatascanExecutionIdentityUserCredential
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable
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
	InterpolationAsList() cdktn.IResolvable
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable
	PutDataplexServiceAgent(value *DataplexDatascanExecutionIdentityDataplexServiceAgent)
	PutServiceAccount(value *DataplexDatascanExecutionIdentityServiceAccount)
	PutUserCredential(value *DataplexDatascanExecutionIdentityUserCredential)
	ResetDataplexServiceAgent()
	ResetServiceAccount()
	ResetUserCredential()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataplexDatascanExecutionIdentityOutputReference
type jsiiProxy_DataplexDatascanExecutionIdentityOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataplexDatascanExecutionIdentityOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexDatascanExecutionIdentityOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexDatascanExecutionIdentityOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexDatascanExecutionIdentityOutputReference) DataplexServiceAgent() DataplexDatascanExecutionIdentityDataplexServiceAgentOutputReference {
	var returns DataplexDatascanExecutionIdentityDataplexServiceAgentOutputReference
	_jsii_.Get(
		j,
		"dataplexServiceAgent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexDatascanExecutionIdentityOutputReference) DataplexServiceAgentInput() *DataplexDatascanExecutionIdentityDataplexServiceAgent {
	var returns *DataplexDatascanExecutionIdentityDataplexServiceAgent
	_jsii_.Get(
		j,
		"dataplexServiceAgentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexDatascanExecutionIdentityOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexDatascanExecutionIdentityOutputReference) InternalValue() *DataplexDatascanExecutionIdentity {
	var returns *DataplexDatascanExecutionIdentity
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexDatascanExecutionIdentityOutputReference) ServiceAccount() DataplexDatascanExecutionIdentityServiceAccountOutputReference {
	var returns DataplexDatascanExecutionIdentityServiceAccountOutputReference
	_jsii_.Get(
		j,
		"serviceAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexDatascanExecutionIdentityOutputReference) ServiceAccountInput() *DataplexDatascanExecutionIdentityServiceAccount {
	var returns *DataplexDatascanExecutionIdentityServiceAccount
	_jsii_.Get(
		j,
		"serviceAccountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexDatascanExecutionIdentityOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexDatascanExecutionIdentityOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexDatascanExecutionIdentityOutputReference) UserCredential() DataplexDatascanExecutionIdentityUserCredentialOutputReference {
	var returns DataplexDatascanExecutionIdentityUserCredentialOutputReference
	_jsii_.Get(
		j,
		"userCredential",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexDatascanExecutionIdentityOutputReference) UserCredentialInput() *DataplexDatascanExecutionIdentityUserCredential {
	var returns *DataplexDatascanExecutionIdentityUserCredential
	_jsii_.Get(
		j,
		"userCredentialInput",
		&returns,
	)
	return returns
}


func NewDataplexDatascanExecutionIdentityOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DataplexDatascanExecutionIdentityOutputReference {
	_init_.Initialize()

	if err := validateNewDataplexDatascanExecutionIdentityOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataplexDatascanExecutionIdentityOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.dataplexDatascan.DataplexDatascanExecutionIdentityOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataplexDatascanExecutionIdentityOutputReference_Override(d DataplexDatascanExecutionIdentityOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.dataplexDatascan.DataplexDatascanExecutionIdentityOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataplexDatascanExecutionIdentityOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataplexDatascanExecutionIdentityOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataplexDatascanExecutionIdentityOutputReference)SetInternalValue(val *DataplexDatascanExecutionIdentity) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataplexDatascanExecutionIdentityOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataplexDatascanExecutionIdentityOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataplexDatascanExecutionIdentityOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataplexDatascanExecutionIdentityOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataplexDatascanExecutionIdentityOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataplexDatascanExecutionIdentityOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataplexDatascanExecutionIdentityOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataplexDatascanExecutionIdentityOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataplexDatascanExecutionIdentityOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataplexDatascanExecutionIdentityOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataplexDatascanExecutionIdentityOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataplexDatascanExecutionIdentityOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataplexDatascanExecutionIdentityOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataplexDatascanExecutionIdentityOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataplexDatascanExecutionIdentityOutputReference) PutDataplexServiceAgent(value *DataplexDatascanExecutionIdentityDataplexServiceAgent) {
	if err := d.validatePutDataplexServiceAgentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putDataplexServiceAgent",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataplexDatascanExecutionIdentityOutputReference) PutServiceAccount(value *DataplexDatascanExecutionIdentityServiceAccount) {
	if err := d.validatePutServiceAccountParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putServiceAccount",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataplexDatascanExecutionIdentityOutputReference) PutUserCredential(value *DataplexDatascanExecutionIdentityUserCredential) {
	if err := d.validatePutUserCredentialParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putUserCredential",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataplexDatascanExecutionIdentityOutputReference) ResetDataplexServiceAgent() {
	_jsii_.InvokeVoid(
		d,
		"resetDataplexServiceAgent",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataplexDatascanExecutionIdentityOutputReference) ResetServiceAccount() {
	_jsii_.InvokeVoid(
		d,
		"resetServiceAccount",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataplexDatascanExecutionIdentityOutputReference) ResetUserCredential() {
	_jsii_.InvokeVoid(
		d,
		"resetUserCredential",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataplexDatascanExecutionIdentityOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := d.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataplexDatascanExecutionIdentityOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

