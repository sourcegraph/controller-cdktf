package certificatepack

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/cloudflare/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/cloudflare/certificatepack/internal"
)

type CertificatePackValidationRecordsOutputReference interface {
	cdktf.ComplexObject
	CnameName() *string
	SetCnameName(val *string)
	CnameNameInput() *string
	CnameTarget() *string
	SetCnameTarget(val *string)
	CnameTargetInput() *string
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
	Emails() *[]*string
	SetEmails(val *[]*string)
	EmailsInput() *[]*string
	// Experimental.
	Fqn() *string
	HttpBody() *string
	SetHttpBody(val *string)
	HttpBodyInput() *string
	HttpUrl() *string
	SetHttpUrl(val *string)
	HttpUrlInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TxtName() *string
	SetTxtName(val *string)
	TxtNameInput() *string
	TxtValue() *string
	SetTxtValue(val *string)
	TxtValueInput() *string
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
	ResetCnameName()
	ResetCnameTarget()
	ResetEmails()
	ResetHttpBody()
	ResetHttpUrl()
	ResetTxtName()
	ResetTxtValue()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CertificatePackValidationRecordsOutputReference
type jsiiProxy_CertificatePackValidationRecordsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CertificatePackValidationRecordsOutputReference) CnameName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cnameName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CertificatePackValidationRecordsOutputReference) CnameNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cnameNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CertificatePackValidationRecordsOutputReference) CnameTarget() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cnameTarget",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CertificatePackValidationRecordsOutputReference) CnameTargetInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cnameTargetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CertificatePackValidationRecordsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CertificatePackValidationRecordsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CertificatePackValidationRecordsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CertificatePackValidationRecordsOutputReference) Emails() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"emails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CertificatePackValidationRecordsOutputReference) EmailsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"emailsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CertificatePackValidationRecordsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CertificatePackValidationRecordsOutputReference) HttpBody() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpBody",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CertificatePackValidationRecordsOutputReference) HttpBodyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpBodyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CertificatePackValidationRecordsOutputReference) HttpUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CertificatePackValidationRecordsOutputReference) HttpUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CertificatePackValidationRecordsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CertificatePackValidationRecordsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CertificatePackValidationRecordsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CertificatePackValidationRecordsOutputReference) TxtName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"txtName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CertificatePackValidationRecordsOutputReference) TxtNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"txtNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CertificatePackValidationRecordsOutputReference) TxtValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"txtValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CertificatePackValidationRecordsOutputReference) TxtValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"txtValueInput",
		&returns,
	)
	return returns
}


func NewCertificatePackValidationRecordsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) CertificatePackValidationRecordsOutputReference {
	_init_.Initialize()

	if err := validateNewCertificatePackValidationRecordsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_CertificatePackValidationRecordsOutputReference{}

	_jsii_.Create(
		"cloudflare.certificatePack.CertificatePackValidationRecordsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewCertificatePackValidationRecordsOutputReference_Override(c CertificatePackValidationRecordsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"cloudflare.certificatePack.CertificatePackValidationRecordsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		c,
	)
}

func (j *jsiiProxy_CertificatePackValidationRecordsOutputReference)SetCnameName(val *string) {
	if err := j.validateSetCnameNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cnameName",
		val,
	)
}

func (j *jsiiProxy_CertificatePackValidationRecordsOutputReference)SetCnameTarget(val *string) {
	if err := j.validateSetCnameTargetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cnameTarget",
		val,
	)
}

func (j *jsiiProxy_CertificatePackValidationRecordsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CertificatePackValidationRecordsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CertificatePackValidationRecordsOutputReference)SetEmails(val *[]*string) {
	if err := j.validateSetEmailsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"emails",
		val,
	)
}

func (j *jsiiProxy_CertificatePackValidationRecordsOutputReference)SetHttpBody(val *string) {
	if err := j.validateSetHttpBodyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpBody",
		val,
	)
}

func (j *jsiiProxy_CertificatePackValidationRecordsOutputReference)SetHttpUrl(val *string) {
	if err := j.validateSetHttpUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpUrl",
		val,
	)
}

func (j *jsiiProxy_CertificatePackValidationRecordsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CertificatePackValidationRecordsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CertificatePackValidationRecordsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_CertificatePackValidationRecordsOutputReference)SetTxtName(val *string) {
	if err := j.validateSetTxtNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"txtName",
		val,
	)
}

func (j *jsiiProxy_CertificatePackValidationRecordsOutputReference)SetTxtValue(val *string) {
	if err := j.validateSetTxtValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"txtValue",
		val,
	)
}

func (c *jsiiProxy_CertificatePackValidationRecordsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CertificatePackValidationRecordsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CertificatePackValidationRecordsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CertificatePackValidationRecordsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CertificatePackValidationRecordsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CertificatePackValidationRecordsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CertificatePackValidationRecordsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CertificatePackValidationRecordsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CertificatePackValidationRecordsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CertificatePackValidationRecordsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CertificatePackValidationRecordsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CertificatePackValidationRecordsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CertificatePackValidationRecordsOutputReference) ResetCnameName() {
	_jsii_.InvokeVoid(
		c,
		"resetCnameName",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CertificatePackValidationRecordsOutputReference) ResetCnameTarget() {
	_jsii_.InvokeVoid(
		c,
		"resetCnameTarget",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CertificatePackValidationRecordsOutputReference) ResetEmails() {
	_jsii_.InvokeVoid(
		c,
		"resetEmails",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CertificatePackValidationRecordsOutputReference) ResetHttpBody() {
	_jsii_.InvokeVoid(
		c,
		"resetHttpBody",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CertificatePackValidationRecordsOutputReference) ResetHttpUrl() {
	_jsii_.InvokeVoid(
		c,
		"resetHttpUrl",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CertificatePackValidationRecordsOutputReference) ResetTxtName() {
	_jsii_.InvokeVoid(
		c,
		"resetTxtName",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CertificatePackValidationRecordsOutputReference) ResetTxtValue() {
	_jsii_.InvokeVoid(
		c,
		"resetTxtValue",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CertificatePackValidationRecordsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := c.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CertificatePackValidationRecordsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

