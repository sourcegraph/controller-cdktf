package googleoracledatabasecloudvmcluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googleoracledatabasecloudvmcluster/internal"
)

type GoogleOracleDatabaseCloudVmClusterPropertiesOutputReference interface {
	cdktf.ComplexObject
	ClusterName() *string
	SetClusterName(val *string)
	ClusterNameInput() *string
	CompartmentId() *string
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
	CpuCoreCount() *float64
	SetCpuCoreCount(val *float64)
	CpuCoreCountInput() *float64
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DataStorageSizeTb() *float64
	SetDataStorageSizeTb(val *float64)
	DataStorageSizeTbInput() *float64
	DbNodeStorageSizeGb() *float64
	SetDbNodeStorageSizeGb(val *float64)
	DbNodeStorageSizeGbInput() *float64
	DbServerOcids() *[]*string
	SetDbServerOcids(val *[]*string)
	DbServerOcidsInput() *[]*string
	DiagnosticsDataCollectionOptions() GoogleOracleDatabaseCloudVmClusterPropertiesDiagnosticsDataCollectionOptionsOutputReference
	DiagnosticsDataCollectionOptionsInput() *GoogleOracleDatabaseCloudVmClusterPropertiesDiagnosticsDataCollectionOptions
	DiskRedundancy() *string
	SetDiskRedundancy(val *string)
	DiskRedundancyInput() *string
	DnsListenerIp() *string
	Domain() *string
	// Experimental.
	Fqn() *string
	GiVersion() *string
	SetGiVersion(val *string)
	GiVersionInput() *string
	Hostname() *string
	HostnamePrefix() *string
	SetHostnamePrefix(val *string)
	HostnamePrefixInput() *string
	InternalValue() *GoogleOracleDatabaseCloudVmClusterProperties
	SetInternalValue(val *GoogleOracleDatabaseCloudVmClusterProperties)
	LicenseType() *string
	SetLicenseType(val *string)
	LicenseTypeInput() *string
	LocalBackupEnabled() interface{}
	SetLocalBackupEnabled(val interface{})
	LocalBackupEnabledInput() interface{}
	MemorySizeGb() *float64
	SetMemorySizeGb(val *float64)
	MemorySizeGbInput() *float64
	NodeCount() *float64
	SetNodeCount(val *float64)
	NodeCountInput() *float64
	Ocid() *string
	OciUrl() *string
	OcpuCount() *float64
	SetOcpuCount(val *float64)
	OcpuCountInput() *float64
	ScanDns() *string
	ScanDnsRecordId() *string
	ScanIpIds() *[]*string
	ScanListenerPortTcp() *float64
	ScanListenerPortTcpSsl() *float64
	Shape() *string
	SparseDiskgroupEnabled() interface{}
	SetSparseDiskgroupEnabled(val interface{})
	SparseDiskgroupEnabledInput() interface{}
	SshPublicKeys() *[]*string
	SetSshPublicKeys(val *[]*string)
	SshPublicKeysInput() *[]*string
	State() *string
	StorageSizeGb() *float64
	SystemVersion() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TimeZone() GoogleOracleDatabaseCloudVmClusterPropertiesTimeZoneOutputReference
	TimeZoneInput() *GoogleOracleDatabaseCloudVmClusterPropertiesTimeZone
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
	PutDiagnosticsDataCollectionOptions(value *GoogleOracleDatabaseCloudVmClusterPropertiesDiagnosticsDataCollectionOptions)
	PutTimeZone(value *GoogleOracleDatabaseCloudVmClusterPropertiesTimeZone)
	ResetClusterName()
	ResetDataStorageSizeTb()
	ResetDbNodeStorageSizeGb()
	ResetDbServerOcids()
	ResetDiagnosticsDataCollectionOptions()
	ResetDiskRedundancy()
	ResetGiVersion()
	ResetHostnamePrefix()
	ResetLocalBackupEnabled()
	ResetMemorySizeGb()
	ResetNodeCount()
	ResetOcpuCount()
	ResetSparseDiskgroupEnabled()
	ResetSshPublicKeys()
	ResetTimeZone()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleOracleDatabaseCloudVmClusterPropertiesOutputReference
type jsiiProxy_GoogleOracleDatabaseCloudVmClusterPropertiesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleOracleDatabaseCloudVmClusterPropertiesOutputReference) ClusterName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseCloudVmClusterPropertiesOutputReference) ClusterNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseCloudVmClusterPropertiesOutputReference) CompartmentId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"compartmentId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseCloudVmClusterPropertiesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseCloudVmClusterPropertiesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseCloudVmClusterPropertiesOutputReference) CpuCoreCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpuCoreCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseCloudVmClusterPropertiesOutputReference) CpuCoreCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpuCoreCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseCloudVmClusterPropertiesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseCloudVmClusterPropertiesOutputReference) DataStorageSizeTb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dataStorageSizeTb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseCloudVmClusterPropertiesOutputReference) DataStorageSizeTbInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dataStorageSizeTbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseCloudVmClusterPropertiesOutputReference) DbNodeStorageSizeGb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dbNodeStorageSizeGb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseCloudVmClusterPropertiesOutputReference) DbNodeStorageSizeGbInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dbNodeStorageSizeGbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseCloudVmClusterPropertiesOutputReference) DbServerOcids() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dbServerOcids",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseCloudVmClusterPropertiesOutputReference) DbServerOcidsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dbServerOcidsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseCloudVmClusterPropertiesOutputReference) DiagnosticsDataCollectionOptions() GoogleOracleDatabaseCloudVmClusterPropertiesDiagnosticsDataCollectionOptionsOutputReference {
	var returns GoogleOracleDatabaseCloudVmClusterPropertiesDiagnosticsDataCollectionOptionsOutputReference
	_jsii_.Get(
		j,
		"diagnosticsDataCollectionOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseCloudVmClusterPropertiesOutputReference) DiagnosticsDataCollectionOptionsInput() *GoogleOracleDatabaseCloudVmClusterPropertiesDiagnosticsDataCollectionOptions {
	var returns *GoogleOracleDatabaseCloudVmClusterPropertiesDiagnosticsDataCollectionOptions
	_jsii_.Get(
		j,
		"diagnosticsDataCollectionOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseCloudVmClusterPropertiesOutputReference) DiskRedundancy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"diskRedundancy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseCloudVmClusterPropertiesOutputReference) DiskRedundancyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"diskRedundancyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseCloudVmClusterPropertiesOutputReference) DnsListenerIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsListenerIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseCloudVmClusterPropertiesOutputReference) Domain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseCloudVmClusterPropertiesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseCloudVmClusterPropertiesOutputReference) GiVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"giVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseCloudVmClusterPropertiesOutputReference) GiVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"giVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseCloudVmClusterPropertiesOutputReference) Hostname() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostname",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseCloudVmClusterPropertiesOutputReference) HostnamePrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostnamePrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseCloudVmClusterPropertiesOutputReference) HostnamePrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostnamePrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseCloudVmClusterPropertiesOutputReference) InternalValue() *GoogleOracleDatabaseCloudVmClusterProperties {
	var returns *GoogleOracleDatabaseCloudVmClusterProperties
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseCloudVmClusterPropertiesOutputReference) LicenseType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"licenseType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseCloudVmClusterPropertiesOutputReference) LicenseTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"licenseTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseCloudVmClusterPropertiesOutputReference) LocalBackupEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"localBackupEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseCloudVmClusterPropertiesOutputReference) LocalBackupEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"localBackupEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseCloudVmClusterPropertiesOutputReference) MemorySizeGb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memorySizeGb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseCloudVmClusterPropertiesOutputReference) MemorySizeGbInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memorySizeGbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseCloudVmClusterPropertiesOutputReference) NodeCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"nodeCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseCloudVmClusterPropertiesOutputReference) NodeCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"nodeCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseCloudVmClusterPropertiesOutputReference) Ocid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ocid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseCloudVmClusterPropertiesOutputReference) OciUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ociUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseCloudVmClusterPropertiesOutputReference) OcpuCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ocpuCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseCloudVmClusterPropertiesOutputReference) OcpuCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ocpuCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseCloudVmClusterPropertiesOutputReference) ScanDns() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scanDns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseCloudVmClusterPropertiesOutputReference) ScanDnsRecordId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scanDnsRecordId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseCloudVmClusterPropertiesOutputReference) ScanIpIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"scanIpIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseCloudVmClusterPropertiesOutputReference) ScanListenerPortTcp() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"scanListenerPortTcp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseCloudVmClusterPropertiesOutputReference) ScanListenerPortTcpSsl() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"scanListenerPortTcpSsl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseCloudVmClusterPropertiesOutputReference) Shape() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shape",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseCloudVmClusterPropertiesOutputReference) SparseDiskgroupEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sparseDiskgroupEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseCloudVmClusterPropertiesOutputReference) SparseDiskgroupEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sparseDiskgroupEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseCloudVmClusterPropertiesOutputReference) SshPublicKeys() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sshPublicKeys",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseCloudVmClusterPropertiesOutputReference) SshPublicKeysInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sshPublicKeysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseCloudVmClusterPropertiesOutputReference) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseCloudVmClusterPropertiesOutputReference) StorageSizeGb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"storageSizeGb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseCloudVmClusterPropertiesOutputReference) SystemVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"systemVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseCloudVmClusterPropertiesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseCloudVmClusterPropertiesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseCloudVmClusterPropertiesOutputReference) TimeZone() GoogleOracleDatabaseCloudVmClusterPropertiesTimeZoneOutputReference {
	var returns GoogleOracleDatabaseCloudVmClusterPropertiesTimeZoneOutputReference
	_jsii_.Get(
		j,
		"timeZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseCloudVmClusterPropertiesOutputReference) TimeZoneInput() *GoogleOracleDatabaseCloudVmClusterPropertiesTimeZone {
	var returns *GoogleOracleDatabaseCloudVmClusterPropertiesTimeZone
	_jsii_.Get(
		j,
		"timeZoneInput",
		&returns,
	)
	return returns
}


func NewGoogleOracleDatabaseCloudVmClusterPropertiesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleOracleDatabaseCloudVmClusterPropertiesOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleOracleDatabaseCloudVmClusterPropertiesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleOracleDatabaseCloudVmClusterPropertiesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleOracleDatabaseCloudVmCluster.GoogleOracleDatabaseCloudVmClusterPropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleOracleDatabaseCloudVmClusterPropertiesOutputReference_Override(g GoogleOracleDatabaseCloudVmClusterPropertiesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleOracleDatabaseCloudVmCluster.GoogleOracleDatabaseCloudVmClusterPropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseCloudVmClusterPropertiesOutputReference)SetClusterName(val *string) {
	if err := j.validateSetClusterNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterName",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseCloudVmClusterPropertiesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseCloudVmClusterPropertiesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseCloudVmClusterPropertiesOutputReference)SetCpuCoreCount(val *float64) {
	if err := j.validateSetCpuCoreCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cpuCoreCount",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseCloudVmClusterPropertiesOutputReference)SetDataStorageSizeTb(val *float64) {
	if err := j.validateSetDataStorageSizeTbParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataStorageSizeTb",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseCloudVmClusterPropertiesOutputReference)SetDbNodeStorageSizeGb(val *float64) {
	if err := j.validateSetDbNodeStorageSizeGbParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dbNodeStorageSizeGb",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseCloudVmClusterPropertiesOutputReference)SetDbServerOcids(val *[]*string) {
	if err := j.validateSetDbServerOcidsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dbServerOcids",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseCloudVmClusterPropertiesOutputReference)SetDiskRedundancy(val *string) {
	if err := j.validateSetDiskRedundancyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"diskRedundancy",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseCloudVmClusterPropertiesOutputReference)SetGiVersion(val *string) {
	if err := j.validateSetGiVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"giVersion",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseCloudVmClusterPropertiesOutputReference)SetHostnamePrefix(val *string) {
	if err := j.validateSetHostnamePrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostnamePrefix",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseCloudVmClusterPropertiesOutputReference)SetInternalValue(val *GoogleOracleDatabaseCloudVmClusterProperties) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseCloudVmClusterPropertiesOutputReference)SetLicenseType(val *string) {
	if err := j.validateSetLicenseTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"licenseType",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseCloudVmClusterPropertiesOutputReference)SetLocalBackupEnabled(val interface{}) {
	if err := j.validateSetLocalBackupEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"localBackupEnabled",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseCloudVmClusterPropertiesOutputReference)SetMemorySizeGb(val *float64) {
	if err := j.validateSetMemorySizeGbParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"memorySizeGb",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseCloudVmClusterPropertiesOutputReference)SetNodeCount(val *float64) {
	if err := j.validateSetNodeCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nodeCount",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseCloudVmClusterPropertiesOutputReference)SetOcpuCount(val *float64) {
	if err := j.validateSetOcpuCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ocpuCount",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseCloudVmClusterPropertiesOutputReference)SetSparseDiskgroupEnabled(val interface{}) {
	if err := j.validateSetSparseDiskgroupEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sparseDiskgroupEnabled",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseCloudVmClusterPropertiesOutputReference)SetSshPublicKeys(val *[]*string) {
	if err := j.validateSetSshPublicKeysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sshPublicKeys",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseCloudVmClusterPropertiesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseCloudVmClusterPropertiesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseCloudVmClusterPropertiesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleOracleDatabaseCloudVmClusterPropertiesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleOracleDatabaseCloudVmClusterPropertiesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleOracleDatabaseCloudVmClusterPropertiesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleOracleDatabaseCloudVmClusterPropertiesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleOracleDatabaseCloudVmClusterPropertiesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleOracleDatabaseCloudVmClusterPropertiesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleOracleDatabaseCloudVmClusterPropertiesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleOracleDatabaseCloudVmClusterPropertiesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleOracleDatabaseCloudVmClusterPropertiesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleOracleDatabaseCloudVmClusterPropertiesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleOracleDatabaseCloudVmClusterPropertiesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleOracleDatabaseCloudVmClusterPropertiesOutputReference) PutDiagnosticsDataCollectionOptions(value *GoogleOracleDatabaseCloudVmClusterPropertiesDiagnosticsDataCollectionOptions) {
	if err := g.validatePutDiagnosticsDataCollectionOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDiagnosticsDataCollectionOptions",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseCloudVmClusterPropertiesOutputReference) PutTimeZone(value *GoogleOracleDatabaseCloudVmClusterPropertiesTimeZone) {
	if err := g.validatePutTimeZoneParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTimeZone",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseCloudVmClusterPropertiesOutputReference) ResetClusterName() {
	_jsii_.InvokeVoid(
		g,
		"resetClusterName",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseCloudVmClusterPropertiesOutputReference) ResetDataStorageSizeTb() {
	_jsii_.InvokeVoid(
		g,
		"resetDataStorageSizeTb",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseCloudVmClusterPropertiesOutputReference) ResetDbNodeStorageSizeGb() {
	_jsii_.InvokeVoid(
		g,
		"resetDbNodeStorageSizeGb",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseCloudVmClusterPropertiesOutputReference) ResetDbServerOcids() {
	_jsii_.InvokeVoid(
		g,
		"resetDbServerOcids",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseCloudVmClusterPropertiesOutputReference) ResetDiagnosticsDataCollectionOptions() {
	_jsii_.InvokeVoid(
		g,
		"resetDiagnosticsDataCollectionOptions",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseCloudVmClusterPropertiesOutputReference) ResetDiskRedundancy() {
	_jsii_.InvokeVoid(
		g,
		"resetDiskRedundancy",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseCloudVmClusterPropertiesOutputReference) ResetGiVersion() {
	_jsii_.InvokeVoid(
		g,
		"resetGiVersion",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseCloudVmClusterPropertiesOutputReference) ResetHostnamePrefix() {
	_jsii_.InvokeVoid(
		g,
		"resetHostnamePrefix",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseCloudVmClusterPropertiesOutputReference) ResetLocalBackupEnabled() {
	_jsii_.InvokeVoid(
		g,
		"resetLocalBackupEnabled",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseCloudVmClusterPropertiesOutputReference) ResetMemorySizeGb() {
	_jsii_.InvokeVoid(
		g,
		"resetMemorySizeGb",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseCloudVmClusterPropertiesOutputReference) ResetNodeCount() {
	_jsii_.InvokeVoid(
		g,
		"resetNodeCount",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseCloudVmClusterPropertiesOutputReference) ResetOcpuCount() {
	_jsii_.InvokeVoid(
		g,
		"resetOcpuCount",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseCloudVmClusterPropertiesOutputReference) ResetSparseDiskgroupEnabled() {
	_jsii_.InvokeVoid(
		g,
		"resetSparseDiskgroupEnabled",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseCloudVmClusterPropertiesOutputReference) ResetSshPublicKeys() {
	_jsii_.InvokeVoid(
		g,
		"resetSshPublicKeys",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseCloudVmClusterPropertiesOutputReference) ResetTimeZone() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeZone",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseCloudVmClusterPropertiesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleOracleDatabaseCloudVmClusterPropertiesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

