package googlesqldatabaseinstance

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlesqldatabaseinstance/internal"
)

type GoogleSqlDatabaseInstanceReplicaConfigurationOutputReference interface {
	cdktf.ComplexObject
	CaCertificate() *string
	SetCaCertificate(val *string)
	CaCertificateInput() *string
	ClientCertificate() *string
	SetClientCertificate(val *string)
	ClientCertificateInput() *string
	ClientKey() *string
	SetClientKey(val *string)
	ClientKeyInput() *string
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
	ConnectRetryInterval() *float64
	SetConnectRetryInterval(val *float64)
	ConnectRetryIntervalInput() *float64
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DumpFilePath() *string
	SetDumpFilePath(val *string)
	DumpFilePathInput() *string
	FailoverTarget() interface{}
	SetFailoverTarget(val interface{})
	FailoverTargetInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleSqlDatabaseInstanceReplicaConfiguration
	SetInternalValue(val *GoogleSqlDatabaseInstanceReplicaConfiguration)
	MasterHeartbeatPeriod() *float64
	SetMasterHeartbeatPeriod(val *float64)
	MasterHeartbeatPeriodInput() *float64
	Password() *string
	SetPassword(val *string)
	PasswordInput() *string
	SslCipher() *string
	SetSslCipher(val *string)
	SslCipherInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Username() *string
	SetUsername(val *string)
	UsernameInput() *string
	VerifyServerCertificate() interface{}
	SetVerifyServerCertificate(val interface{})
	VerifyServerCertificateInput() interface{}
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
	ResetCaCertificate()
	ResetClientCertificate()
	ResetClientKey()
	ResetConnectRetryInterval()
	ResetDumpFilePath()
	ResetFailoverTarget()
	ResetMasterHeartbeatPeriod()
	ResetPassword()
	ResetSslCipher()
	ResetUsername()
	ResetVerifyServerCertificate()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleSqlDatabaseInstanceReplicaConfigurationOutputReference
type jsiiProxy_GoogleSqlDatabaseInstanceReplicaConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceReplicaConfigurationOutputReference) CaCertificate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"caCertificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceReplicaConfigurationOutputReference) CaCertificateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"caCertificateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceReplicaConfigurationOutputReference) ClientCertificate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientCertificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceReplicaConfigurationOutputReference) ClientCertificateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientCertificateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceReplicaConfigurationOutputReference) ClientKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceReplicaConfigurationOutputReference) ClientKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceReplicaConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceReplicaConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceReplicaConfigurationOutputReference) ConnectRetryInterval() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"connectRetryInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceReplicaConfigurationOutputReference) ConnectRetryIntervalInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"connectRetryIntervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceReplicaConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceReplicaConfigurationOutputReference) DumpFilePath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dumpFilePath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceReplicaConfigurationOutputReference) DumpFilePathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dumpFilePathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceReplicaConfigurationOutputReference) FailoverTarget() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"failoverTarget",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceReplicaConfigurationOutputReference) FailoverTargetInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"failoverTargetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceReplicaConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceReplicaConfigurationOutputReference) InternalValue() *GoogleSqlDatabaseInstanceReplicaConfiguration {
	var returns *GoogleSqlDatabaseInstanceReplicaConfiguration
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceReplicaConfigurationOutputReference) MasterHeartbeatPeriod() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"masterHeartbeatPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceReplicaConfigurationOutputReference) MasterHeartbeatPeriodInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"masterHeartbeatPeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceReplicaConfigurationOutputReference) Password() *string {
	var returns *string
	_jsii_.Get(
		j,
		"password",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceReplicaConfigurationOutputReference) PasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceReplicaConfigurationOutputReference) SslCipher() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslCipher",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceReplicaConfigurationOutputReference) SslCipherInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslCipherInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceReplicaConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceReplicaConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceReplicaConfigurationOutputReference) Username() *string {
	var returns *string
	_jsii_.Get(
		j,
		"username",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceReplicaConfigurationOutputReference) UsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usernameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceReplicaConfigurationOutputReference) VerifyServerCertificate() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"verifyServerCertificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceReplicaConfigurationOutputReference) VerifyServerCertificateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"verifyServerCertificateInput",
		&returns,
	)
	return returns
}


func NewGoogleSqlDatabaseInstanceReplicaConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleSqlDatabaseInstanceReplicaConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleSqlDatabaseInstanceReplicaConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleSqlDatabaseInstanceReplicaConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleSqlDatabaseInstance.GoogleSqlDatabaseInstanceReplicaConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleSqlDatabaseInstanceReplicaConfigurationOutputReference_Override(g GoogleSqlDatabaseInstanceReplicaConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleSqlDatabaseInstance.GoogleSqlDatabaseInstanceReplicaConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceReplicaConfigurationOutputReference)SetCaCertificate(val *string) {
	if err := j.validateSetCaCertificateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"caCertificate",
		val,
	)
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceReplicaConfigurationOutputReference)SetClientCertificate(val *string) {
	if err := j.validateSetClientCertificateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientCertificate",
		val,
	)
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceReplicaConfigurationOutputReference)SetClientKey(val *string) {
	if err := j.validateSetClientKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientKey",
		val,
	)
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceReplicaConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceReplicaConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceReplicaConfigurationOutputReference)SetConnectRetryInterval(val *float64) {
	if err := j.validateSetConnectRetryIntervalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connectRetryInterval",
		val,
	)
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceReplicaConfigurationOutputReference)SetDumpFilePath(val *string) {
	if err := j.validateSetDumpFilePathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dumpFilePath",
		val,
	)
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceReplicaConfigurationOutputReference)SetFailoverTarget(val interface{}) {
	if err := j.validateSetFailoverTargetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"failoverTarget",
		val,
	)
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceReplicaConfigurationOutputReference)SetInternalValue(val *GoogleSqlDatabaseInstanceReplicaConfiguration) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceReplicaConfigurationOutputReference)SetMasterHeartbeatPeriod(val *float64) {
	if err := j.validateSetMasterHeartbeatPeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"masterHeartbeatPeriod",
		val,
	)
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceReplicaConfigurationOutputReference)SetPassword(val *string) {
	if err := j.validateSetPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"password",
		val,
	)
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceReplicaConfigurationOutputReference)SetSslCipher(val *string) {
	if err := j.validateSetSslCipherParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sslCipher",
		val,
	)
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceReplicaConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceReplicaConfigurationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceReplicaConfigurationOutputReference)SetUsername(val *string) {
	if err := j.validateSetUsernameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"username",
		val,
	)
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceReplicaConfigurationOutputReference)SetVerifyServerCertificate(val interface{}) {
	if err := j.validateSetVerifyServerCertificateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"verifyServerCertificate",
		val,
	)
}

func (g *jsiiProxy_GoogleSqlDatabaseInstanceReplicaConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleSqlDatabaseInstanceReplicaConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleSqlDatabaseInstanceReplicaConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleSqlDatabaseInstanceReplicaConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleSqlDatabaseInstanceReplicaConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleSqlDatabaseInstanceReplicaConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleSqlDatabaseInstanceReplicaConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleSqlDatabaseInstanceReplicaConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleSqlDatabaseInstanceReplicaConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleSqlDatabaseInstanceReplicaConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleSqlDatabaseInstanceReplicaConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleSqlDatabaseInstanceReplicaConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleSqlDatabaseInstanceReplicaConfigurationOutputReference) ResetCaCertificate() {
	_jsii_.InvokeVoid(
		g,
		"resetCaCertificate",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleSqlDatabaseInstanceReplicaConfigurationOutputReference) ResetClientCertificate() {
	_jsii_.InvokeVoid(
		g,
		"resetClientCertificate",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleSqlDatabaseInstanceReplicaConfigurationOutputReference) ResetClientKey() {
	_jsii_.InvokeVoid(
		g,
		"resetClientKey",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleSqlDatabaseInstanceReplicaConfigurationOutputReference) ResetConnectRetryInterval() {
	_jsii_.InvokeVoid(
		g,
		"resetConnectRetryInterval",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleSqlDatabaseInstanceReplicaConfigurationOutputReference) ResetDumpFilePath() {
	_jsii_.InvokeVoid(
		g,
		"resetDumpFilePath",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleSqlDatabaseInstanceReplicaConfigurationOutputReference) ResetFailoverTarget() {
	_jsii_.InvokeVoid(
		g,
		"resetFailoverTarget",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleSqlDatabaseInstanceReplicaConfigurationOutputReference) ResetMasterHeartbeatPeriod() {
	_jsii_.InvokeVoid(
		g,
		"resetMasterHeartbeatPeriod",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleSqlDatabaseInstanceReplicaConfigurationOutputReference) ResetPassword() {
	_jsii_.InvokeVoid(
		g,
		"resetPassword",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleSqlDatabaseInstanceReplicaConfigurationOutputReference) ResetSslCipher() {
	_jsii_.InvokeVoid(
		g,
		"resetSslCipher",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleSqlDatabaseInstanceReplicaConfigurationOutputReference) ResetUsername() {
	_jsii_.InvokeVoid(
		g,
		"resetUsername",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleSqlDatabaseInstanceReplicaConfigurationOutputReference) ResetVerifyServerCertificate() {
	_jsii_.InvokeVoid(
		g,
		"resetVerifyServerCertificate",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleSqlDatabaseInstanceReplicaConfigurationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleSqlDatabaseInstanceReplicaConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

