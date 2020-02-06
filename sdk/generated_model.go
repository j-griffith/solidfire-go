package sdk

// THIS IS A GENERATED FILE.  DO NOT MODIFY

//A Protection Domain is a set of one or more components whose simultaneous failure is protected
//from causing data unavailability or loss. This specifies one of the types of Protection Domains
//recognized by this cluster.
type ProtectionDomainType string

//The method of protecting data on the cluster
type ProtectionScheme string

//This specifies error code for a proposed node addition.
type ProposedNodeErrorCode string

//This type qualifies a ClusterAdmin with its authentication method.
type AuthMethod string

//This specifies a drive's encryption capability.
type DriveEncryptionCapabilityType string

//This specifies a node's FIPS 140-2 compliance status.
type FipsDrivesStatusType string

//The category of the protection scheme.
type ProtectionSchemeCategory string

//The public visibility of the protection scheme.
type ProtectionSchemeVisibility string

//Status of the remote snapshot on the target cluster as seen on the source cluster
type RemoteClusterSnapshotStatus string

//Describes host access for a volume.
type VolumeAccess string

//This type indicates the configuration data which will be accessed or modified by the element auth container.
type AuthConfigType string
type NeedsWorkEnableClusterSshResult struct {
}

type VirtualVolumeHostResult struct {
	//
	Host VirtualVolumeHost `json:"host,"`
}

type VirtualVolumeSyncResult struct {
	//VolumeID for the newly created volume.
	VolumeID int64 `json:"volumeID,"`
	//virtualVolumeID for the newly created volume.
	VirtualVolumeID string `json:"virtualVolumeID,"`
}

type GetClusterInterfacePreferenceResult struct {
	//The cluster interface preference for the given name.
	Preference ClusterInterfacePreference `json:"preference,"`
}

type ListActiveNodesResult struct {
	//
	Nodes []Node `json:"nodes,"`
}

type AddAccountResult struct {
	//AccountID for the newly created Account.
	AccountID int64 `json:"accountID,"`
	//The full account object
	Account Account `json:"account,omitempty"`
}

type TestPingResult struct {
	//Result of the ping test.
	Result string `json:"result,"`
	//The total duration of the ping test.
	Duration string `json:"duration,"`
	//List of each IP the node was able to communicate with.
	Details interface{} `json:"details,"`
}

type ListSnapMirrorLunsResult struct {
	//A list of objects containing information about SnapMirror LUNs.
	SnapMirrorLunInfos []SnapMirrorLunInfo `json:"snapMirrorLunInfos,"`
}

type NeedsWorkProtocolVersionUpgradeResult struct {
}

type VirtualNetworkAddress struct {
	//SolidFire unique identifier for a virtual network.
	VirtualNetworkID int64 `json:"virtualNetworkID,"`
	//Virtual Network Address.
	Address string `json:"address,"`
}

type FibreChannelPortList struct {
	//List of all physical Fibre Channel ports.
	FibreChannelPorts []FibreChannelPortInfo `json:"fibreChannelPorts,"`
}

type GetIpmiInfoResult struct {
	//Detailed information from each sensor within a node.
	Nodes []GetIpmiInfoNodesResult `json:"nodes,"`
}

type EnableFeatureResult struct {
}

type ListVolumeQoSHistogramsResult struct {
	//List of VolumeQoSHistograms.
	QosHistograms []VolumeQoSHistograms `json:"qosHistograms,"`
}

type GetEncryptionAtRestInfoResult struct {
	//Details about the Encryption At Rest feature.
	EncryptionAtRestInfo EncryptionAtRestInfo `json:"encryptionAtRestInfo,"`
}

type ProtectionDomainResiliency struct {
	//The predicted number of simultaneous failures which may occur without losing
	//the ability to automatically heal to where the Ensemble Quorum has Node
	//Tolerance.
	SustainableFailuresForEnsemble int64 `json:"sustainableFailuresForEnsemble,"`
	//The maximum number of bytes that can be stored on the cluster before losing
	//the ability to automatically heal to where the data has Node Tolerance.
	SingleFailureThresholdBytesForBlockData int64 `json:"singleFailureThresholdBytesForBlockData,"`
	//List of objects detailing failure resiliency information for the associated
	//ProtectionDomainType, one for each Protection Scheme.
	ProtectionSchemeResiliencies []ProtectionSchemeResiliency `json:"protectionSchemeResiliencies,"`
}

type PurgeDeletedVolumesResult struct {
}

type SetSnmpInfoResult struct {
}

type SnmpNetwork struct {
	//ro: read-only access.
	//rw: for read-write access.
	//rosys: for read-only access to a restricted set of system information
	//SolidFire recommends that all networks other than the default "localhost" be set to "ro" access, because all SolidFire MIB objects are read-only.
	Access string `json:"access,"`
	//A CIDR network mask. This network mask must be an integer greater than or equal to 0, and less than or equal to 32. It must also not be equal to 31.
	Cidr int64 `json:"cidr,"`
	//SNMP community string.
	Community string `json:"community,"`
	//This parameter ainteger with the cidr variable is used to control which network the access and community string apply to. The special value of "default" is used to specify an entry that applies to all networks. The cidr mask is ignored when network value is either a host name or default.
	Network string `json:"network,"`
}

type DeleteKeyProviderKmipResult struct {
}

type DisableLdapAuthenticationResult struct {
}

type DrivesConfigInfo struct {
	//
	Drives []DriveConfigInfo `json:"drives,"`
	//
	NumBlockActual int64 `json:"numBlockActual,"`
	//
	NumBlockExpected int64 `json:"numBlockExpected,"`
	//
	NumSliceActual int64 `json:"numSliceActual,"`
	//
	NumSliceExpected int64 `json:"numSliceExpected,"`
	//
	NumTotalActual int64 `json:"numTotalActual,"`
	//
	NumTotalExpected int64 `json:"numTotalExpected,"`
}

type ResetDriveDetails struct {
	//Drive name
	Drive string `json:"drive,"`
	//
	ReturnCode int64 `json:"returnCode,"`
	//
	Stderr string `json:"stderr,"`
	//
	Stdout string `json:"stdout,"`
}

type SnmpV3UsmUser struct {
	//rouser: read-only access.
	//rwuser: for read-write access.
	//rosys: for read-only access to a restricted set of system information
	//SolidFire recommends that all USM users be set to "rouser" access, because all SolidFire MIB objects are read-only.
	Access string `json:"access,"`
	//The name of the user. Must contain at least one character, but no more than 32 characters. Blank spaces are not allowed.
	Name string `json:"name,"`
	//The password of the user. Must be between 8 and 255 characters integer (inclusive). Blank spaces are not allowed. Required if "secLevel" is "auth" or "priv."
	Password string `json:"password,"`
	//The passphrase of the user. Must be between 8 and 255 characters integer (inclusive). Blank spaces are not allowed. Required if "secLevel" is "priv."
	Passphrase string `json:"passphrase,"`
	//noauth: No password or passphrase is required.
	//auth: A password is required for user access.
	//priv: A password and passphrase is required for user access.
	SecLevel string `json:"secLevel,"`
}

type GetClusterMasterNodeIDResult struct {
	//ID of the master node.
	NodeID int64 `json:"nodeID,"`
}

type ListPendingActiveNodesResult struct {
	//List of objects detailing information about all PendingActive nodes in the system.
	PendingActiveNodes []PendingActiveNode `json:"pendingActiveNodes,"`
}

type GetStorageContainerEfficiencyResult struct {
	//
	Compression float64 `json:"compression,"`
	//
	Deduplication float64 `json:"deduplication,"`
	//The volumes that could not be queried for efficiency data. Missing volumes can be caused by the Garbage Collection (GC) cycle being less than an hour old, temporary loss of network connectivity, or restarted services since the GC cycle.
	MissingVolumes []int64 `json:"missingVolumes,"`
	//
	ThinProvisioning float64 `json:"thinProvisioning,"`
	//The last time efficiency data was collected after Garbage Collection (GC).
	Timestamp string `json:"timestamp,"`
}

type ListVirtualVolumesResult struct {
	//
	VirtualVolumes []VirtualVolumeInfo `json:"virtualVolumes,"`
	//
	NextVirtualVolumeID string `json:"nextVirtualVolumeID,omitempty"`
}

type ModifySnapMirrorRelationshipResult struct {
	//An object containg the modified SnapMirror relationship attributes.
	SnapMirrorRelationship []SnapMirrorRelationship `json:"snapMirrorRelationship,"`
}

type ModifyVolumeResult struct {
	//Object containing information about the newly modified volume.
	Volume Volume `json:"volume,omitempty"`
}

type NetworkParams struct {
	//Name of the storage node network interface used for management traffic.
	Bond1G NetworkConfigParams `json:"Bond1G,omitempty"`
	//Name of the storage node network interface used for storage and cluster traffic.
	Bond10G NetworkConfigParams `json:"Bond10G,omitempty"`
	//Name of the witness node network interface used for management traffic.
	Net0 NetworkConfigParams `json:"net0,omitempty"`
	//Name of the witness node network interface used for storage and cluster traffic.
	Net1 NetworkConfigParams `json:"net1,omitempty"`
	//
	Eth0 NetworkConfigParams `json:"eth0,omitempty"`
	//
	Eth1 NetworkConfigParams `json:"eth1,omitempty"`
	//
	Eth2 NetworkConfigParams `json:"eth2,omitempty"`
	//
	Eth3 NetworkConfigParams `json:"eth3,omitempty"`
	//
	Lo NetworkConfigParams `json:"lo,omitempty"`
}

type VolumeQOS struct {
	//Desired minimum 4KB IOPS to guarantee.
	//The allowed IOPS will only drop below this level if all volumes have been capped
	//at their min IOPS value and there is still insufficient performance capacity.
	MinIOPS int64 `json:"minIOPS,"`
	//Desired maximum 4KB IOPS allowed over an extended period of time.
	MaxIOPS int64 `json:"maxIOPS,"`
	//Maximum "peak" 4KB IOPS allowed for short periods of time.
	//Allows for bursts of I/O activity over the normal max IOPS value.
	BurstIOPS int64 `json:"burstIOPS,"`
	//The length of time burst IOPS is allowed.
	//The value returned is represented in time units of seconds.
	//Note: this value is calculated by the system based on IOPS set for QoS.
	BurstTime int64 `json:"burstTime,"`
	//The curve is a set of key-value pairs.
	//The keys are I/O sizes in bytes.
	//The values represent the cost of performing an IOP at a specific I/O size.
	//The curve is calculated relative to a 4096 byte operation set at 100 IOPS.
	Curve int64 `json:"curve,"`
}

type BlockSizeHistogram struct {
	//Number of block size samples between 512 and 4095 bytes
	Bucket512To4095 int64 `json:"Bucket512To4095,"`
	//Number of block size samples between 4096 and 8191 bytes
	Bucket4096to8191 int64 `json:"bucket4096to8191,"`
	//Number of block size samples between 8192 and 16383 bytes
	Bucket8192To16383 int64 `json:"Bucket8192To16383,"`
	//Number of block size samples between 16384 and 32767 bytes
	Bucket16384To32767 int64 `json:"Bucket16384To32767,"`
	//Number of block size samples between 32768 and 65535 bytes
	Bucket32768To65535 int64 `json:"Bucket32768To65535,"`
	//Number of block size samples between 65536 and 131071 bytes
	Bucket65536To131071 int64 `json:"Bucket65536To131071,"`
	//Number of block size samples greater than or equal to 131072 bytes
	Bucket131072Plus int64 `json:"Bucket131072Plus,"`
}

type DriveDetails struct {
	ReturnCode int64  `json:"returnCode,"`
	Drive      string `json:"drive,"`
	Stderr     string `json:"stderr,"`
	Stdout     string `json:"stdout,"`
}

type NeedsWorkGetGCStatusResult struct {
}

type NeedsWorkReadSliceLBAChecksumResult struct {
}

type PendingNode struct {
	//
	PendingNodeID int64 `json:"pendingNodeID,"`
	//
	AssignedNodeID int64 `json:"assignedNodeID,"`
	//The host name for this node.
	Name string `json:"name,"`
	//Indicates whether the pending node's software version is compatible with the cluster.
	Compatible bool `json:"compatible,"`
	//Information about the node's hardware.
	PlatformInfo Platform `json:"platformInfo,"`
	//The node's role in the cluster. Possible values are Management, Storage, Compute, and Witness.
	Role string `json:"role,"`
	//IP address used for both intra-cluster and inter-cluster communication.
	Cip string `json:"cip,"`
	//The machine's name for the "cip" interface.
	Cipi string `json:"cipi,"`
	//IP address used for the per-node API and UI.
	Mip string `json:"mip,"`
	//The machine's name for the "mip" interface.
	Mipi string `json:"mipi,"`
	//IP address used for iSCSI traffic.
	Sip string `json:"sip,"`
	//The machine's name for the "sip" interface.
	Sipi string `json:"sipi,"`
	//The version of SolidFire software currently running on this node.
	SoftwareVersion string `json:"softwareVersion,"`
	//UUID of node.
	Uuid string `json:"uuid,"`
	//
	NodeSlot string `json:"nodeSlot,omitempty"`
	//Uniquely identifies a chassis, and identical for all nodes in a given chassis.
	ChassisName string `json:"chassisName,"`
	//Uniquely identifies a custom protection domain, identical for all nodes within all chassis in a given custom protection domain.
	CustomProtectionDomainName string `json:"customProtectionDomainName,"`
}

type ClusterCapacity struct {
	//
	ThinProvisioningPercent int64 `json:"thinProvisioningPercent,"`
	//
	DeDuplicationPercent int64 `json:"deDuplicationPercent,"`
	//
	CompressionPercent int64 `json:"compressionPercent,"`
	//
	EfficiencyPercent int64 `json:"efficiencyPercent,"`
	//The amount of space on the block drives.
	//This includes additional information such as metadata entries and space which can be cleaned up.
	ActiveBlockSpace int64 `json:"activeBlockSpace,"`
	//Number of active iSCSI sessions communicating with the cluster
	ActiveSessions int64 `json:"activeSessions,"`
	//Average IPS for the cluster since midnight Coordinated Universal Time (UTC).
	AverageIOPS int64 `json:"averageIOPS,"`
	//The average size of IOPS to all volumes in the cluster.
	ClusterRecentIOSize int64 `json:"clusterRecentIOSize,"`
	//Average IOPS for all volumes in the cluster over the last 5 seconds.
	CurrentIOPS int64 `json:"currentIOPS,"`
	//Estimated maximum IOPS capability of the current cluster.
	MaxIOPS int64 `json:"maxIOPS,"`
	//The maximum amount of provisionable space.
	//This is a computed value.
	//You cannot create new volumes if the current provisioned space plus the new volume size would exceed this number:
	//maxOverProvisionableSpace = maxProvisionedSpace * GetClusterFull
	MaxOverProvisionableSpace int64 `json:"maxOverProvisionableSpace,"`
	//The total amount of provisionable space if all volumes are 100% filled (no thin provisioned metadata).
	MaxProvisionedSpace int64 `json:"maxProvisionedSpace,"`
	//The amount of bytes on volume drives used to store metadata.
	MaxUsedMetadataSpace int64 `json:"maxUsedMetadataSpace,"`
	//The total amount of space on all active block drives.
	MaxUsedSpace int64 `json:"maxUsedSpace,"`
	//Total number of 4KiB blocks with data after the last garbage collection operation has completed.
	NonZeroBlocks int64 `json:"nonZeroBlocks,"`
	//Peak number of iSCSI connections since midnight UTC.
	PeakActiveSessions int64 `json:"peakActiveSessions,"`
	//The highest value for currentIOPS since midnight UTC.
	PeakIOPS int64 `json:"peakIOPS,"`
	//Total amount of space provisioned in all volumes on the cluster.
	ProvisionedSpace int64 `json:"provisionedSpace,"`
	//Total number of 4KiB blocks in snapshots with data.
	SnapshotNonZeroBlocks int64 `json:"snapshotNonZeroBlocks,"`
	//The date and time this cluster capacity sample was taken.
	Timestamp string `json:"timestamp,"`
	//The total number of I/O operations performed throughout the lifetime of the cluster
	TotalOps int64 `json:"totalOps,"`
	//The total number of blocks stored on the block drives.
	//The value includes replicated blocks.
	UniqueBlocks int64 `json:"uniqueBlocks,"`
	//The total amount of data the uniqueBlocks take up on the block drives.
	//This number is always consistent with the uniqueBlocks value.
	UniqueBlocksUsedSpace int64 `json:"uniqueBlocksUsedSpace,"`
	//The total amount of bytes on volume drives used to store metadata
	UsedMetadataSpace int64 `json:"usedMetadataSpace,"`
	//The amount of bytes on volume drives used for storing unique data in snapshots.
	//This number provides an estimate of how much metadata space would be regained by deleting all snapshots on the system.
	UsedMetadataSpaceInSnapshots int64 `json:"usedMetadataSpaceInSnapshots,"`
	//Total amount of space used by all block drives in the system.
	UsedSpace int64 `json:"usedSpace,"`
	//Total number of 4KiB blocks without data after the last round of garabage collection operation has completed.
	ZeroBlocks int64 `json:"zeroBlocks,"`
}

type EnableLdapAuthenticationResult struct {
}

type SnmpSendTestTrapsResult struct {
	//
	Status string `json:"status,"`
}

type StartVolumePairingResult struct {
	//A string of characters that is used by the "CompleteVolumePairing" API method.
	VolumePairingKey string `json:"volumePairingKey,"`
}

type GetSnmpTrapInfoResult struct {
	//List of hosts that are to receive the traps generated by the cluster.
	TrapRecipients []SnmpTrapRecipient `json:"trapRecipients,"`
	//If "true", when a cluster fault is logged a corresponding solidFireClusterFaultNotification is sent to the configured list of trap recipients.
	ClusterFaultTrapsEnabled bool `json:"clusterFaultTrapsEnabled,"`
	//If "true", when a cluster fault is logged a corresponding solidFireClusterFaultResolvedNotification is sent to the configured list of trap recipients.
	ClusterFaultResolvedTrapsEnabled bool `json:"clusterFaultResolvedTrapsEnabled,"`
	//If "true", when a cluster fault is logged a corresponding solidFireClusterEventNotification is sent to the configured list of trap recipients.
	ClusterEventTrapsEnabled bool `json:"clusterEventTrapsEnabled,"`
}

type ListAsyncResultsResult struct {
	//An array of serialized asynchronous method results.
	AsyncHandles []AsyncHandle `json:"asyncHandles,"`
}

type NeedsWorkListClusterCapacityHistoryResult struct {
}

type CloneVolumeResult struct {
	//The resulting volume
	Volume Volume `json:"volume,omitempty"`
	//The ID of the newly-created clone.
	CloneID int64 `json:"cloneID,"`
	//The snapshot ID of the newly-created clone.
	SnapshotID int64 `json:"snapshotID,"`
	//The volume ID of the newly-created clone.
	VolumeID int64 `json:"volumeID,"`
	//The curve is a set of key-value pairs.
	//The keys are I/O sizes in bytes.
	//The values represent the cost of performing an IOP at a specific I/O size.
	//The curve is calculated relative to a 4096 byte operation set at 100 IOPS.
	Curve int64 `json:"curve,"`
	//Handle value used to track the progress of the clone.
	AsyncHandle int64 `json:"asyncHandle,"`
}

type ListVolumeStatsByVolumeAccessGroupResult struct {
	//List of account activity information.
	//Note: The volumeID member is 0 for each entry, as the values represent the summation of all volumes owned by the account.
	VolumeStats []VolumeStats `json:"volumeStats,"`
}

type NeedsWorkListDatabaseChildrenResult struct {
}

type NetworkConfigParams struct {
	//
	Default bool `json:"#default,omitempty"`
	//
	Bond_master string `json:"bond-master,omitempty"`
	//
	VirtualNetworkTag string `json:"virtualNetworkTag,omitempty"`
	//
	Address string `json:"address,omitempty"`
	//
	Auto bool `json:"auto,omitempty"`
	//
	Bond_downdelay string `json:"bond-downdelay,omitempty"`
	//
	Bond_fail_over_mac string `json:"bond-fail_over_mac,omitempty"`
	//
	Bond_primary_reselect string `json:"bond-primary_reselect,omitempty"`
	//
	Bond_lacp_rate string `json:"bond-lacp_rate,omitempty"`
	//
	Bond_miimon string `json:"bond-miimon,omitempty"`
	//
	Bond_mode string `json:"bond-mode,omitempty"`
	//
	Bond_slaves string `json:"bond-slaves,omitempty"`
	//
	Bond_updelay string `json:"bond-updelay,omitempty"`
	//
	Dns_nameservers string `json:"dns-nameservers,omitempty"`
	//
	Dns_search string `json:"dns-search,omitempty"`
	//
	Family string `json:"family,omitempty"`
	//
	Gateway string `json:"gateway,omitempty"`
	//
	MacAddress string `json:"macAddress,omitempty"`
	//
	MacAddressPermanent string `json:"macAddressPermanent,omitempty"`
	//
	Method string `json:"method,omitempty"`
	//
	Mtu string `json:"mtu,omitempty"`
	//
	Netmask string `json:"netmask,omitempty"`
	//
	Network string `json:"network,omitempty"`
	//
	Physical PhysicalAdapter `json:"physical,omitempty"`
	//
	Routes []interface{} `json:"routes,omitempty"`
	//
	Status string `json:"status,omitempty"`
	//
	SymmetricRouteRules []string `json:"symmetricRouteRules,omitempty"`
	//
	UpAndRunning bool `json:"upAndRunning,omitempty"`
}

type SetSSLCertificateResult struct {
}

type KeyProviderKmip struct {
	//The name of the KMIP Key Provider.
	KeyProviderName string `json:"keyProviderName,"`
	//The ID of the KMIP Key Provider.  This is a unique value assigned by the cluster during CreateKeyProviderKmip which cannot be changed.
	KeyProviderID int64 `json:"keyProviderID,"`
	//True if the KMIP Key Provider is active.  A provider is considered active if are still outstanding keys which were created but not yet deleted and therefore assumed to still be in use.
	KeyProviderIsActive bool `json:"keyProviderIsActive,"`
	//A list of keyServerIDs which are grouped together within this provider.  At least one server must be added via AddKeyServerToProviderKmip before this provider can become active.  The last server cannot be removed via RemoveKeyServerFromProviderKmip or DeleteKeyServerKmip while this provider is active.
	KeyServerIDs []int64 `json:"keyServerIDs,"`
	//The capabilities of this KMIP Key Provider including details about the underlying library, FIPS compliance, SSL provider, etc.
	KmipCapabilities string `json:"kmipCapabilities,"`
}

type ListProtocolEndpointsResult struct {
	//
	ProtocolEndpoints []ProtocolEndpoint `json:"protocolEndpoints,"`
}

type SnapMirrorRelationship struct {
	//The ID of the destination ONTAP system.
	SnapMirrorEndpointID int64 `json:"snapMirrorEndpointID,"`
	//The name of the destination ONTAP cluster.
	ClusterName string `json:"clusterName,"`
	//The unique identifier for each snapMirrorRelationship object
	//in an array as would be returned in ListSnapMirrorRelationships.
	//This UUID is created and returned from the ONTAP system.
	SnapMirrorRelationshipID string `json:"snapMirrorRelationshipID,"`
	//An object describing the source volume.
	SourceVolume SnapMirrorVolumeInfo `json:"sourceVolume,"`
	//An object describing the destination volume.
	DestinationVolume SnapMirrorVolumeInfo `json:"destinationVolume,"`
	//The current maximum transfer rate between the
	//source and destination volumes, in kilobytes per second.
	CurrentMaxTransferRate int64 `json:"currentMaxTransferRate,"`
	//Whether the relationship is healthy or not.
	//Possible values:
	//true:  The relationship is healthy.
	//false: The relationship is not healthy. This can be caused by a manual or
	//       scheduled update failing or being aborted, or by the last scheduled
	//       update being delayed.
	IsHealthy bool `json:"isHealthy,"`
	//The amount of time in seconds by which the data on the
	//destination volume lags behind the data on the source volume.
	Lagtime int64 `json:"lagtime,"`
	//The amount of time in seconds it took for the last transfer to complete.
	LastTransferDuration int64 `json:"lastTransferDuration,"`
	//A message describing the cause of the last transfer failure.
	LastTransferError string `json:"lastTransferError,"`
	//The total number of bytes transferred during the last transfer.
	LastTransferSize int64 `json:"lastTransferSize,"`
	//The timestamp of the end of the last transfer.
	LastTransferEndTimestamp string `json:"lastTransferEndTimestamp,"`
	//The type of the previous transfer in the relationship.
	LastTransferType string `json:"lastTransferType,"`
	//Specifies the maximum data transfer rate
	//between the volumes in kilobytes per second.
	//The default value, 0, is unlimited and permits the SnapMirror
	//relationship to fully utilize the available network bandwidth.
	MaxTransferRate int64 `json:"maxTransferRate,"`
	//The mirror state of the SnapMirror relationship.
	//Possible values:
	//uninitialized: The destination volume has not been initialized.
	//snapmirrored:  The destination volume has been initialized and is ready to recieve SnapMirror updates.
	//broken-off:    The destination volume is read-write and snapshots are present.
	MirrorState string `json:"mirrorState,"`
	//The name of the newest Snapshot copy on the destination volume.
	NewestSnapshot string `json:"newestSnapshot,"`
	//Specifies the name of the ONTAP SnapMirror policy for the relationship.
	//A list of available policies can be retrieved with ListSnapMirrorPolicies.
	//Example values are "MirrorLatest" and "MirrorAndVault".
	PolicyName string `json:"policyName,"`
	//The type of the ONTAP SnapMirror policy for the relationship.
	//See ListSnapMirrorPolicies. Examples are: "async_mirror" or "mirror_vault"
	PolicyType string `json:"policyType,"`
	//The status of the SnapMirror relationship.
	//Possible values:
	//idle
	//transferring
	//checking
	//quiescing
	//quiesced
	//queued
	//preparing
	//finalizing
	//aborting
	//breaking
	RelationshipStatus string `json:"relationshipStatus,"`
	//The type of the SnapMirror relationship.
	//On SolidFire systems, this value is always "extended_data_protection".
	ReleationshipType string `json:"releationshipType,"`
	//The name of the pre-existing cron schedule on the ONTAP
	//system that is used to update the SnapMirror relationship.
	//A list of available schedules can be retrieved with ListSnapMirrorSchedules.
	ScheduleName string `json:"scheduleName,"`
	//The reason the relationship is not healthy.
	UnhealthyReason string `json:"unhealthyReason,"`
}

type SnapMirrorVserverAggregateInfo struct {
	//The name of the aggregate assigned to a Vserver.
	AggrName string `json:"aggrName,"`
	//The assigned aggregate's available size.
	AggrAvailSize int64 `json:"aggrAvailSize,"`
}

type ListVirtualNetworksResult struct {
	//Object containing virtual network IP addresses.
	VirtualNetworks []VirtualNetwork `json:"virtualNetworks,"`
}

type QuiesceSnapMirrorRelationshipResult struct {
	//An object containg information about the quiesced SnapMirror relationship.
	SnapMirrorRelationship SnapMirrorRelationship `json:"snapMirrorRelationship,"`
}

type RemoveNodeSSLCertificateResult struct {
}

type SnapMirrorPolicy struct {
	//The ID of the destination ONTAP system.
	SnapMirrorEndpointID int64 `json:"snapMirrorEndpointID,"`
	//The unique name assigned to the policy.
	PolicyName string `json:"policyName,"`
	//The type of policy.
	//Possible values:
	//async_mirror
	//mirror_vault
	PolicyType string `json:"policyType,"`
	//A human-readable description associated with the SnapMirror policy.
	Comment string `json:"comment,"`
	//The priority at which a SnapMirror transfer runs.
	//Possible values:
	//normal: The default priority. These transfers are
	//        scheduled before most low priority transfers.
	//low:    These transfers have the lowest priority and
	//        are scheduled after most normal priority transfers.
	TransferPriority string `json:"transferPriority,"`
	//A list of objects describing the policy rules.
	PolicyRules []SnapMirrorPolicyRule `json:"policyRules,"`
	//The total retention count for all rules in the policy.
	TotalKeepCount int64 `json:"totalKeepCount,"`
	//The total number of rules in the policy.
	TotalRules int64 `json:"totalRules,"`
	//The name of the Vserver for the SnapMirror policy.
	VserverName string `json:"vserverName,"`
}

type Origin struct {
	//
	Signature Signature `json:"<signature>,"`
	//
	Contract_date string `json:"contract-date,"`
	//
	Contract_name string `json:"contract-name,"`
	//
	Contract_quantity int64 `json:"contract-quantity,"`
	//
	Contract_type string `json:"contract-type,"`
	//
	Integrator string `json:"integrator,"`
	//
	Location string `json:"location,"`
	//
	Organization string `json:"organization,"`
	//
	Type string `json:"type,"`
}

type PrepareVirtualSnapshotResult struct {
	//The ID of the clone task.
	VirtualVolumeTaskID string `json:"virtualVolumeTaskID,"`
	//The volume ID of the newly-created clone.
	VolumeID int64 `json:"volumeID,"`
	//snapshotID for the prepared VVol snapshot.
	SnapshotID int64 `json:"snapshotID,"`
	//virtualVolumeID for the newly created clone.
	VirtualVolumeID string `json:"virtualVolumeID,"`
}

type GetNodeActiveTlsCiphersResult struct {
	//List of mandatory TLS cipher suites for the node.
	MandatoryCiphers []string `json:"mandatoryCiphers,"`
	//List of supplemental TLS cipher suites for the node.
	SupplementalCiphers []string `json:"supplementalCiphers,"`
}

type ListSnapMirrorNodesResult struct {
	//A list of the nodes on the ONTAP cluster.
	SnapMirrorNodes []SnapMirrorNode `json:"snapMirrorNodes,"`
}

type LoggingServer struct {
	//Hostname or IP address of the log server.
	Host string `json:"host,"`
	//Port number that the log server is listening on.
	Port int64 `json:"port,"`
}

type NeedsWorkEnableBmcColdResetResult struct {
}

type RemoveClusterPairResult struct {
}

type TestLdapAuthenticationResult struct {
	//List of LDAP groups that the tested user is a member of.
	Groups []string `json:"groups,"`
	//The tested user's full LDAP distinguished name.
	UserDN string `json:"userDN,"`
}

type AuthSessionInfo struct {
	//Cluster AdminID(s) associated with this session.  For sessions related to LDAP or
	//a third party Identity Provider (IdP), this will be an aggregate list of matching
	//Cluster AdminIDs associated with this session.
	ClusterAdminIDs []int64 `json:"clusterAdminIDs,"`
	//Username associated with this session.  For sessions related to LDAP
	//this will be the user's LDAP DN.  For sessions related to a third party
	//Identity Provider (IdP), this will be an arbitrary name-value pair that will be
	//used for auditing operations within the session.  It will not necessarily match
	//a cluster admin name on the cluster.  For example, a SAML Subject NameID,
	//but this will be dictated by the configuration of the IdP and the resultant content
	//of the SAML assertion.
	Username string `json:"username,"`
	//UUID for this session.
	SessionID string `json:"sessionID,"`
	//Time at which the session was created.
	SessionCreationTime string `json:"sessionCreationTime,"`
	//Time at which the session becomes invalid.
	//This is set when the session is created and cannot be changed.
	FinalTimeout string `json:"finalTimeout,"`
	//Time at which the session becomes invalid due to inactivity.
	//It is set to a new value when the session is accessed for use,
	//up to the time where the session becomes invalid due to finalTimeout being reached.
	LastAccessTimeout string `json:"lastAccessTimeout,"`
	//List of access groups for the user.
	AccessGroupList []string `json:"accessGroupList,"`
	//Method in which the cluster admin was authenticated.
	AuthMethod AuthMethod `json:"authMethod,"`
}

type ClusterVersionInfo struct {
	//
	NodeID int64 `json:"nodeID,"`
	//
	NodeVersion string `json:"nodeVersion,"`
	//
	NodeInternalRevision string `json:"nodeInternalRevision,"`
}

type VirtualVolumeAsyncResult struct {
	//The ID of the clone task.
	VirtualVolumeTaskID string `json:"virtualVolumeTaskID,"`
	//The volume ID of the newly-created clone.
	VolumeID int64 `json:"volumeID,"`
	//virtualVolumeID for the newly created clone.
	VirtualVolumeID string `json:"virtualVolumeID,"`
}

type GetClusterInfoResult struct {
	//
	ClusterInfo ClusterInfo `json:"clusterInfo,"`
}

type GetHardwareInfoResult struct {
	//Hardware information for this node.
	HardwareInfo interface{} `json:"hardwareInfo,"`
}

type FastCloneVirtualVolumeRequest struct {
	//The ID of the Virtual Volume to clone.
	VirtualVolumeID string `json:"virtualVolumeID,"`
	//The name for the newly-created volume.
	Name string `json:"name,omitempty"`
	//New quality of service settings for this volume.
	Qos QoS `json:"qos,omitempty"`
	//
	Metadata interface{} `json:"metadata,omitempty"`
	//
	CallingVirtualVolumeHostID string `json:"callingVirtualVolumeHostID,omitempty"`
}

type FeatureObject struct {
	//True if the feature is enabled, otherwise false.
	Enabled bool `json:"enabled,"`
	//The name of the feature.
	Feature string `json:"feature,"`
}

type DeleteAuthSessionResult struct {
	//SessionInfo for the auth session deleted.
	Session AuthSessionInfo `json:"session,"`
}

type GetVirtualVolumeTaskUpdateRequest struct {
	//The UUID of the VVol Task.
	VirtualVolumeTaskID string `json:"virtualVolumeTaskID,"`
	//
	CallingVirtualVolumeHostID string `json:"callingVirtualVolumeHostID,omitempty"`
}

type ListKeyProvidersKmipResult struct {
	//A list of KMIP (Key Management Interoperability Protocol) Key Providers which have been created via CreateKeyProviderKmip.
	KmipKeyProviders []KeyProviderKmip `json:"kmipKeyProviders,"`
}

type ProtocolEndpoint struct {
	//
	ProtocolEndpointID string `json:"protocolEndpointID,"`
	//
	ProtocolEndpointState string `json:"protocolEndpointState,"`
	//
	ProviderType string `json:"providerType,"`
	//
	PrimaryProviderID int64 `json:"primaryProviderID,"`
	//
	SecondaryProviderID int64 `json:"secondaryProviderID,"`
	//
	ScsiNAADeviceID string `json:"scsiNAADeviceID,"`
}

type SnapMirrorAggregate struct {
	//The ID of the destination ONTAP system.
	SnapMirrorEndpointID int64 `json:"snapMirrorEndpointID,"`
	//The name of the aggregate.
	AggregateName string `json:"aggregateName,"`
	//The name of the ONTAP node that owns this aggregate.
	NodeName string `json:"nodeName,"`
	//The number of available bytes remaining in the aggregate.
	SizeAvailable int64 `json:"sizeAvailable,"`
	//The total size (int bytes) of the aggregate.
	SizeTotal int64 `json:"sizeTotal,"`
	//The percentage of disk space currently in use.
	PercentUsedCapacity int64 `json:"percentUsedCapacity,"`
	//The number of volumes in the aggregate.
	VolumeCount int64 `json:"volumeCount,"`
}

type CancelVirtualVolumeTaskRequest struct {
	//The UUID of the VVol Task to cancel.
	VirtualVolumeTaskID string `json:"virtualVolumeTaskID,"`
	//
	CallingVirtualVolumeHostID string `json:"callingVirtualVolumeHostID,omitempty"`
}

type CreateClusterResult struct {
}

type GetVolumeStatsResult struct {
	//Volume activity information.
	VolumeStats VolumeStats `json:"volumeStats,"`
}

type NeedsWorkPairVolumeResult struct {
}

type SnapMirrorClusterIdentity struct {
	//The ID of the destination ONTAP system.
	SnapMirrorEndpointID int64 `json:"snapMirrorEndpointID,"`
	//The name of the destination ONTAP cluster.
	ClusterName string `json:"clusterName,"`
	//The 128-bit universally-unique identifier of the destination ONTAP cluster.
	ClusterUUID string `json:"clusterUUID,"`
	//The serial number of the destination ONTAP cluster.
	ClusterSerialNumber string `json:"clusterSerialNumber,"`
}

type VirtualVolumeBinding struct {
	//The unique ID of the protocol endpoint.
	ProtocolEndpointID string `json:"protocolEndpointID,"`
	//The scsiNAADeviceID of the protocol endpoint. For more information, see protocolEndpoint.
	ProtocolEndpointInBandID string `json:"protocolEndpointInBandID,"`
	//The type of protocol endpoint. SCSI is the only value returned for the protocol endpoint type.
	ProtocolEndpointType string `json:"protocolEndpointType,"`
	//The unique ID of the virtual volume binding object.
	VirtualVolumeBindingID int64 `json:"virtualVolumeBindingID,"`
	//The unique ID of the virtual volume host.
	VirtualVolumeHostID string `json:"virtualVolumeHostID,"`
	//The unique ID of the virtual volume.
	VirtualVolumeID string `json:"virtualVolumeID,"`
	//The secondary ID of the virtual volume.
	VirtualVolumeSecondaryID string `json:"virtualVolumeSecondaryID,"`
}

type ClusterInterfacePreference struct {
	//Name of the cluster interface preference
	Name string `json:"name,"`
	//Value of the cluster interface preference
	Value string `json:"value,"`
}

type CompleteClusterPairingResult struct {
	//Unique identifier for the cluster pair.
	ClusterPairID int64 `json:"clusterPairID,"`
}

type QueryVirtualVolumeMetadataResult struct {
	//
	MatchingVirtualVolumeIDs []string `json:"matchingVirtualVolumeIDs,"`
}

type RollbackToGroupSnapshotResult struct {
	//
	GroupSnapshot GroupSnapshot `json:"groupSnapshot,omitempty"`
	//Unique ID of the new group snapshot.
	GroupSnapshotID int64 `json:"groupSnapshotID,omitempty"`
	//List of checksum, volumeIDs and snapshotIDs for each member of the group.
	Members []GroupSnapshotMembers `json:"members,omitempty"`
}

type RollbackToSnapshotResult struct {
	//
	Snapshot Snapshot `json:"snapshot,omitempty"`
	//ID of the newly-created snapshot.
	SnapshotID int64 `json:"snapshotID,omitempty"`
	//A string that represents the correct digits in the stored snapshot.
	//This checksum can be used later to compare other snapshots to detect errors in the data.
	Checksum string `json:"checksum,omitempty"`
}

type GetClusterConfigResult struct {
	//Cluster configuration information the node uses to communicate with the cluster.
	Cluster ClusterConfig `json:"cluster,"`
}

type NodeStatsNodes struct {
	//Node activity information for a single node.
	Nodes []NodeStatsInfo `json:"nodes,"`
}

type GetLoginSessionInfoResult struct {
	//The authentication expiration period. Formatted in H:mm:ss. For example: 1:30:00, 20:00, 5:00. All leading zeros and colons are removed regardless of the format the timeout was entered.
	//Objects returned are:
	//timeout - The time, in minutes, when this session will timeout and expire.
	LoginSessionInfo LoginSessionInfo `json:"loginSessionInfo,"`
}

type GetNvramInfoResult struct {
	//Arrays of events and errors detected on the NVRAM card.
	NvramInfo interface{} `json:"nvramInfo,"`
}

type ListSnapMirrorVolumesResult struct {
	//A list of the SnapMirror volumes available on the ONTAP storage system.
	SnapMirrorVolumes []SnapMirrorVolume `json:"snapMirrorVolumes,"`
}

type ModifyQoSPolicyResult struct {
	//Details of the newly modified QoSPolicy object.
	QosPolicy QoSPolicy `json:"qosPolicy,"`
}

type NeedsWorkCheckBlockResult struct {
}

type NeedsWorkGetDaemonStatusResult struct {
}

type AbortSnapMirrorRelationshipResult struct {
	//An object containing information about the aborted SnapMirror relationship.
	SnapMirrorRelationship SnapMirrorRelationship `json:"snapMirrorRelationship,"`
}

type GetLimitsResult struct {
	//
	AccountCountMax int64 `json:"accountCountMax,"`
	//
	AccountNameLengthMax int64 `json:"accountNameLengthMax,"`
	//
	AccountNameLengthMin int64 `json:"accountNameLengthMin,"`
	//
	BulkVolumeJobsPerNodeMax int64 `json:"bulkVolumeJobsPerNodeMax,"`
	//
	BulkVolumeJobsPerVolumeMax int64 `json:"bulkVolumeJobsPerVolumeMax,"`
	//
	CloneJobsPerVolumeMax int64 `json:"cloneJobsPerVolumeMax,"`
	//
	ClusterPairsCountMax int64 `json:"clusterPairsCountMax,"`
	//
	InitiatorNameLengthMax int64 `json:"initiatorNameLengthMax,"`
	//
	InitiatorCountMax int64 `json:"initiatorCountMax,"`
	//
	InitiatorsPerVolumeAccessGroupCountMax int64 `json:"initiatorsPerVolumeAccessGroupCountMax,"`
	//
	IscsiSessionsFromFibreChannelNodesMax int64 `json:"iscsiSessionsFromFibreChannelNodesMax,"`
	//
	QosPolicyCountMax int64 `json:"qosPolicyCountMax,"`
	//
	SecretLengthMax int64 `json:"secretLengthMax,"`
	//
	ScheduleNameLengthMax int64 `json:"scheduleNameLengthMax,"`
	//
	SecretLengthMin int64 `json:"secretLengthMin,"`
	//
	SnapshotNameLengthMax int64 `json:"snapshotNameLengthMax,"`
	//
	SnapshotsPerVolumeMax int64 `json:"snapshotsPerVolumeMax,"`
	//
	VolumeAccessGroupCountMax int64 `json:"volumeAccessGroupCountMax,"`
	//
	VolumeAccessGroupLunMax int64 `json:"volumeAccessGroupLunMax,"`
	//
	VolumeAccessGroupNameLengthMax int64 `json:"volumeAccessGroupNameLengthMax,"`
	//
	VolumeAccessGroupNameLengthMin int64 `json:"volumeAccessGroupNameLengthMin,"`
	//
	VolumeAccessGroupsPerInitiatorCountMax int64 `json:"volumeAccessGroupsPerInitiatorCountMax,"`
	//
	VolumeAccessGroupsPerVolumeCountMax int64 `json:"volumeAccessGroupsPerVolumeCountMax,"`
	//
	InitiatorAliasLengthMax int64 `json:"initiatorAliasLengthMax,"`
	//
	VolumeBurstIOPSMax int64 `json:"volumeBurstIOPSMax,"`
	//
	VolumeBurstIOPSMin int64 `json:"volumeBurstIOPSMin,"`
	//
	VolumeCountMax int64 `json:"volumeCountMax,"`
	//
	VolumeMaxIOPSMax int64 `json:"volumeMaxIOPSMax,"`
	//
	VolumeMaxIOPSMin int64 `json:"volumeMaxIOPSMin,"`
	//
	VolumeMinIOPSMax int64 `json:"volumeMinIOPSMax,"`
	//
	VolumeMinIOPSMin int64 `json:"volumeMinIOPSMin,"`
	//
	VolumeNameLengthMax int64 `json:"volumeNameLengthMax,"`
	//
	VolumeNameLengthMin int64 `json:"volumeNameLengthMin,"`
	//
	VolumeSizeMax int64 `json:"volumeSizeMax,"`
	//
	VolumeSizeMin int64 `json:"volumeSizeMin,"`
	//
	VolumesPerAccountCountMax int64 `json:"volumesPerAccountCountMax,"`
	//
	VolumesPerGroupSnapshotMax int64 `json:"volumesPerGroupSnapshotMax,"`
	//
	VolumesPerVolumeAccessGroupCountMax int64 `json:"volumesPerVolumeAccessGroupCountMax,"`
	//
	ClusterAdminAccountMax int64 `json:"clusterAdminAccountMax,omitempty"`
	//
	FibreChannelVolumeAccessMax int64 `json:"fibreChannelVolumeAccessMax,omitempty"`
	//
	VirtualVolumesPerAccountCountMax int64 `json:"virtualVolumesPerAccountCountMax,omitempty"`
	//
	VirtualVolumeCountMax int64 `json:"virtualVolumeCountMax,omitempty"`
}

type SetNodeSupplementalTlsCiphersResult struct {
	//List of mandatory TLS cipher suites for the node.
	MandatoryCiphers []string `json:"mandatoryCiphers,"`
	//List of supplemental TLS cipher suites for the node.
	SupplementalCiphers []string `json:"supplementalCiphers,"`
}

type GetRemoteLoggingHostsResult struct {
	//List of hosts to forward logging information to.
	RemoteHosts []LoggingServer `json:"remoteHosts,"`
}

type PhysicalAdapter struct {
	//
	Address string `json:"address,omitempty"`
	//
	MacAddress string `json:"macAddress,omitempty"`
	//
	MacAddressPermanent string `json:"macAddressPermanent,omitempty"`
	//
	Mtu string `json:"mtu,omitempty"`
	//
	Netmask string `json:"netmask,omitempty"`
	//
	Network string `json:"network,omitempty"`
	//
	UpAndRunning bool `json:"upAndRunning,omitempty"`
}

type ResetDrivesDetails struct {
	//Details of a single drive that is being reset.
	Drives []ResetDriveDetails `json:"drives,"`
}

type UpdateIdpConfigurationResult struct {
	//Information around the third party Identity Provider (IdP) configuration.
	IdpConfigInfo IdpConfigInfo `json:"idpConfigInfo,"`
}

type DisableSnmpResult struct {
}

type DrivesHardware struct {
	//
	DriveHardware []DriveHardware `json:"driveHardware,"`
}

type ProposedNodeErrors struct {
	//error code associated with a proposed node
	ErrorCode ProposedNodeErrorCode `json:"errorCode,"`
	//description of the error code
	Description string `json:"description,"`
	//CIP of node(s) encountering the error code
	NodeIPs []string `json:"nodeIPs,"`
}

type RemoveClusterAdminResult struct {
}

type AsyncResult struct {
	//The return value for the original method call if the call was completed successfully.
	Message string `json:"message,"`
}

type NeedsWorkAPIWriteBlockResult struct {
}

type NeedsWorkReleaseFreeMemoryResult struct {
}

type SnapMirrorEndpoint struct {
	//The unique identifier for the object in the local cluster.
	SnapMirrorEndpointID int64 `json:"snapMirrorEndpointID,"`
	//The cluster management IP address of the endpoint.
	ManagementIP string `json:"managementIP,"`
	//The ONTAP cluster name. This value is automatically populated with
	//the value of "clusterName" from the snapMirrorClusterIdentity object.
	ClusterName string `json:"clusterName,"`
	//The management username for the ONTAP system.
	Username string `json:"username,"`
	//List of the inter-cluster storage IP addresses for all nodes in the cluster.
	//You can get these IP addresses with the ListSnapMirrorNetworkInterfaces method.
	IpAddresses []string `json:"ipAddresses,"`
	//The connectivity status of the control link to the ONTAP cluster.
	IsConnected bool `json:"isConnected,"`
}

type VirtualVolumeBindingListResult struct {
	//Describes the VVol <-> Host binding.
	Bindings []VirtualVolumeBinding `json:"bindings,"`
}

type CreateClusterSupportBundleResult struct {
	//The details of the support bundle. Values returned in 'details':
	//bundleName- The name specified in the 'CreateSupportBundle API method. If no name was specified, 'supportbundle' will be used.
	//extraArgs- The arguments passed with this method.
	//files- A list of the support bundle files that were created.
	//output- The command line output from the script that creates the support bundle.
	//timeoutSec- The timeout specified for the support bundle creation process.
	//url- URL to the support bundle created.
	Details interface{} `json:"details,"`
	//The amount of time required to create the support bundle in the format HH:MM:SS.ssssss
	Duration string `json:"duration,"`
	//Whether the support bundle creation passed of failed.
	Result string `json:"result,"`
}

type ModifyVolumeAccessGroupResult struct {
	//An object containing information about the newly modified volume access group.
	VolumeAccessGroup VolumeAccessGroup `json:"volumeAccessGroup,"`
}

type InitializeSnapMirrorRelationshipResult struct {
	//Information about the initialized SnapMirror relationship.
	SnapMirrorRelationship SnapMirrorRelationship `json:"snapMirrorRelationship,"`
}

type NeedsWorkGetClusterSettingsResult struct {
}

type NeedsWorkIsUpgradeInProgressResult struct {
}

type TestConnectMvipDetails struct {
	//Details of the ping tests with 56 Bytes and 1500 Bytes.
	PingBytes interface{} `json:"pingBytes,"`
	//The MVIP tested against.
	Mvip string `json:"mvip,"`
	//Whether the test could connect to the MVIP.
	Connected bool `json:"connected,"`
}

type Drive struct {
	//A unique identifier for this drive.
	DriveID int64 `json:"driveID,"`
	//The node this drive is located.
	//If the drive has been physically removed from the node, this is where it was last seen.
	NodeID int64 `json:"nodeID,"`
	//If this drive is hosting a service, the identifier for that service.
	AssignedService int64 `json:"assignedService,omitempty"`
	//The list of asynchronous jobs currently running on the drive (for example: a secure erase job).
	AsyncResultIDs []int64 `json:"asyncResultIDs,"`
	//Total Raw capacity of the drive, in bytes.
	Capacity int64 `json:"capacity,"`
	//Total Usable capacity of the drive, in bytes.
	UsableCapacity int64 `json:"usableCapacity,"`
	//Segment File Size of the drive, in bytes.
	SegmentFileSize int64 `json:"segmentFileSize,"`
	//The manufacturer's serial number for this drive.
	Serial string `json:"serial,"`
	//Slot number in the server chassis where this drive is located.
	//If the drive has been physically removed from the node, this will not have a value.
	Slot int64 `json:"slot,omitempty"`
	//The current status of this drive.
	DriveStatus string `json:"driveStatus,"`
	//If a drive's status is 'Failed', this field provides more detail on why the drive was marked failed.
	DriveFailureDetail string `json:"driveFailureDetail,omitempty"`
	//If enabling or disabling drive security failed, this is the reason why it failed.
	//If the value is 'none', there was no failure.
	DriveSecurityFaultReason string `json:"driveSecurityFaultReason,omitempty"`
	//Identifies the provider of the authentication key for unlocking this drive.
	KeyProviderID int64 `json:"keyProviderID,omitempty"`
	//The keyID used by the key provider to acquire the authentication key for unlocking this drive.
	KeyID string `json:"keyID,omitempty"`
	//The type of this drive.
	DriveType string `json:"driveType,"`
	//
	ReservedSliceFileCapacity int64 `json:"reservedSliceFileCapacity,omitempty"`
	//
	CustomerSliceFileCapacity int64 `json:"customerSliceFileCapacity,omitempty"`
	SmartSsdWriteCapable      bool  `json:"smartSsdWriteCapable,omitempty"`
	//List of Name/Value pairs in JSON object format.
	Attributes interface{} `json:"attributes,"`
}

type GetBootstrapConfigResult struct {
	//Name of the cluster.
	ClusterName string `json:"clusterName,"`
	//Name of the node.
	NodeName string `json:"nodeName,"`
	//List of descriptions for each node that is actively waiting to join this cluster: compatible - Indicates if the listed node is compatible with the node the API call was executed against. name - IP address of each node. version - version of SolidFire Element software currently installed on the node.
	Nodes []NodeWaitingToJoin `json:"nodes,"`
	//Version of the SolidFire Element software currently installed.
	Version string `json:"version,"`
	//Cluster MVIP address.  This will be null if the node is not in a cluster.
	Mvip string `json:"mvip,"`
	//Cluster SVIP address.  This will be null if the node is not in a cluster.
	Svip string `json:"svip,"`
}

type ResyncSnapMirrorRelationshipResult struct {
	//An object containing information about the resynced SnapMirror relationship.
	SnapMirrorRelationship SnapMirrorRelationship `json:"snapMirrorRelationship,"`
}

type AddressBlockParams struct {
	//Start of the IP address range.
	Start string `json:"start,"`
	//Number of IP addresses to include in the block.
	Size int64 `json:"size,"`
	//dynamic bitset
	Available string `json:"available,"`
}

type ResurrectDeadVirtualVolumeResult struct {
}

type GroupSnapshotRemoteStatus struct {
	//Current status of the remote group snapshot on the target cluster as seen on the source cluster
	RemoteStatus RemoteClusterSnapshotStatus `json:"remoteStatus,"`
}

type CreateSnapMirrorVolumeResult struct {
	//Information about a SnapMirror volume.
	SnapMirrorVolume SnapMirrorVolume `json:"snapMirrorVolume,"`
}

type CreateVirtualVolumeRequest struct {
	//Name of the Virtual Volume.
	//Not required to be unique, but it is recommended.
	//May be 1 to 64 characters in length.
	Name string `json:"name,"`
	//UUID for the Storage Container of this volume.
	StorageContainerID string `json:"storageContainerID,"`
	//VMW_TYPE value for this volume.
	VirtualVolumeType string `json:"virtualVolumeType,"`
	//Total size of the volume, in bytes. Size is rounded up to the nearest 1MB size.
	TotalSize int64 `json:"totalSize,"`
	//Initial quality of service settings for this volume.
	//
	//Volumes created without specified QoS values are created with the default values for QoS.
	//Default values for a volume can be found by running the GetDefaultQoS method.
	Qos QoS `json:"qos,omitempty"`
	//List of name/value pairs to save in the volume's metadata.
	Metadata interface{} `json:"metadata,omitempty"`
	//
	CallingVirtualVolumeHostID string `json:"callingVirtualVolumeHostID,omitempty"`
}

type NeedsWorkTestClientConnectivityResult struct {
}

type VolumeQoSHistograms struct {
	//VolumeID for this volume.
	VolumeID []int64 `json:"volumeID,"`
	//The time and date that the histograms were returned.
	Timestamp []int64 `json:"timestamp,"`
	//Shows the distribution of samples where IO sent to the volume was below its minimum IOP setting.
	BelowMinIopsPercentages []QuintileHistogram `json:"belowMinIopsPercentages,"`
	//Shows the distribution of samples where IO sent to the volume was above its minimum IOP setting.
	//Burst is shown in the histogram's Bucket101Plus entry.
	MinToMaxIopsPercentages []QuintileHistogram `json:"minToMaxIopsPercentages,"`
	//Shows the volume's overall utilization.
	TargetUtilizationPercentages []QuintileHistogram `json:"targetUtilizationPercentages,"`
	//Shows how often and how severely the volume was being throttled.
	ThrottlePercentages []QuintileHistogram `json:"throttlePercentages,"`
	//Shows the distribution of block sizes for read requests
	ReadBlockSizes []BlockSizeHistogram `json:"readBlockSizes,"`
	//Shows the distribution of block sizes for write requests
	WriteBlockSizes []BlockSizeHistogram `json:"writeBlockSizes,"`
}

type GetAccountResult struct {
	//Account details.
	Account Account `json:"account,"`
}

type GetDriveHardwareInfoResult struct {
	//
	DriveHardwareInfo DriveHardwareInfo `json:"driveHardwareInfo,"`
}

type VolumeAccessGroup struct {
	//A list of deleted volumes that have yet to be purged from the VAG.
	DeletedVolumes []int64 `json:"deletedVolumes,"`
	//Unique ID for this volume access group.
	VolumeAccessGroupID int64 `json:"volumeAccessGroupID,"`
	//Name of the volume access group.
	Name string `json:"name,"`
	//A list of IDs of initiators that are mapped to the VAG.
	InitiatorIDs []int64 `json:"initiatorIDs,"`
	//List of unique initiator names beintegering to the volume access group.
	Initiators []string `json:"initiators,"`
	//List of volumes beintegering to the volume access group.
	Volumes []int64 `json:"volumes,"`
	//List of name/value pairs
	Attributes interface{} `json:"attributes,"`
}

type CheckVolumeStandbysAssignedResult struct {
	//Set to true if standbys are assigned on all volumes hosted by this node.
	AllStandbysAssigned bool `json:"allStandbysAssigned,"`
	//Volumes that do not have a standby slice service assigned.
	VolumeIDsWithUnassignedStandbys []int64 `json:"volumeIDsWithUnassignedStandbys,"`
}

type CreatePublicPrivateKeyPairResult struct {
}

type BackupTarget struct {
	//Name for the backup target.
	Name string `json:"name,"`
	//Unique identifier assigned to the backup target.
	BackupTargetID int64 `json:"backupTargetID,"`
	//List of Name/Value pairs in JSON object format.
	Attributes interface{} `json:"attributes,omitempty"`
}

type NeedsWorkGetSliceInfoResult struct {
}

type GetVolumeEfficiencyResult struct {
	//The amount of space being saved by compressing data on a single volume.
	//Stated as a ratio where "1" means data has been stored without being compressed.
	Compression float64 `json:"compression,omitempty"`
	//The amount of space being saved on a single volume by not duplicating data.
	//Stated as a ratio.
	Deduplication float64 `json:"deduplication,"`
	//The volumes that could not be queried for efficiency data.
	//Missing volumes can be caused by GC being less than hour old, temporary network loss or restarted services since the GC cycle.
	MissingVolumes []int64 `json:"missingVolumes,"`
	//The ratio of space used to the amount of space allocated for storing data.
	//Stated as a ratio.
	ThinProvisioning float64 `json:"thinProvisioning,"`
	//The last time efficiency data was collected after Garbage Collection (GC).
	Timestamp string `json:"timestamp,"`
}

type ListClusterInterfacePreferencesResult struct {
	//The cluster interface preferences.
	Preferences []ClusterInterfacePreference `json:"preferences,"`
}

type NodeNetworkInterface struct {
	//The ID of the node.
	NodeID int64 `json:"nodeID,"`
	//
	Result NetworkInterfaces `json:"result,"`
}

type RemoveAccountResult struct {
}

type SnapMirrorLunInfo struct {
	//The ID of the destination ONTAP system.
	SnapMirrorEndpointID int64 `json:"snapMirrorEndpointID,"`
	//The creation time of the LUN.
	CreationTimestamp string `json:"creationTimestamp,"`
	//The name of the LUN.
	LunName string `json:"lunName,"`
	//The path of the LUN.
	Path string `json:"path,"`
	//The size of the LUN in bytes.
	Size int64 `json:"size,"`
	//The number of bytes used by the LUN.
	SizeUsed int64 `json:"sizeUsed,"`
	//The current access state of the LUN.
	//Possible values:
	//online
	//offline
	//foreign_lun_error
	//nvfail
	//space_error
	State string `json:"state,"`
	//The name of the volume that contains the LUN.
	Volume string `json:"volume,"`
	//The Vserver that contains the LUN.
	Vserver string `json:"vserver,"`
}

type CreateKeyProviderKmipResult struct {
	//The KMIP (Key Management Interoperability Protocol) Key Provider which has been created.
	KmipKeyProvider KeyProviderKmip `json:"kmipKeyProvider,"`
}

type GetBackupTargetResult struct {
	//Object returned for backup target.
	BackupTarget BackupTarget `json:"backupTarget,"`
}

type ListAllNodesResult struct {
	//
	Nodes []Node `json:"nodes,"`
	//
	PendingNodes []PendingNode `json:"pendingNodes,"`
	//List of objects detailing information about all PendingActive nodes in the system.
	PendingActiveNodes []PendingActiveNode `json:"pendingActiveNodes,omitempty"`
}

type NeedsWorkDisableClusterSshResult struct {
}

type CertificateDetails struct {
	//The name of the issuer.
	Issuer string `json:"issuer,"`
	//The modulus of the public.
	Modulus string `json:"modulus,"`
	//The expiry date of the certificate.
	NotAfter string `json:"notAfter,"`
	//The start date of the certificate.
	NotBefore string `json:"notBefore,"`
	//The certificate serial number.
	Serial string `json:"serial,"`
	//The digest of the DER-encoded version of the certificate.
	Sha1Fingerprint string `json:"sha1Fingerprint,"`
	//The subject name.
	Subject string `json:"subject,"`
}

type CreateBackupTargetResult struct {
	//Unique identifier assigned to the backup target.
	BackupTargetID int64 `json:"backupTargetID,"`
}

type BreakSnapMirrorRelationshipResult struct {
	//An object containing information about the broken SnapMirror relationship.
	SnapMirrorRelationship SnapMirrorRelationship `json:"snapMirrorRelationship,"`
}

type NodeStateResult struct {
	//ID of the node.
	NodeID int64 `json:"nodeID,"`
	//NodeStateInfo object.
	Result NodeStateInfo `json:"result,omitempty"`
}

type ProtectionDomainTolerance struct {
	//The number of simultaneous failures of the associated ProtectionDomainType which
	//can occur without losing the ensemble quorum.
	SustainableFailuresForEnsemble int64 `json:"sustainableFailuresForEnsemble,"`
	//List of objects detailing failure tolerance information for the associated
	//ProtectionDomainType, one for each Protection Scheme.
	ProtectionSchemeTolerances []ProtectionSchemeTolerance `json:"protectionSchemeTolerances,"`
}

type ResetNodeSupplementalTlsCiphersResult struct {
}

type GetOriginNode struct {
	//
	NodeID int64 `json:"nodeID,"`
	//
	Result GetOriginNodeResult `json:"result,"`
}

type Initiator struct {
	//The friendly name assigned to this initiator.
	Alias string `json:"alias,"`
	//The numeric ID of the initiator that has been created.
	InitiatorID int64 `json:"initiatorID,"`
	//The name of the initiator that has been created.
	InitiatorName string `json:"initiatorName,"`
	//A list of volumeAccessGroupIDs to which this initiator belongs.
	VolumeAccessGroups []int64 `json:"volumeAccessGroups,"`
	//A set of JSON attributes assigned to this initiator.
	Attributes interface{} `json:"attributes,"`
	//True if CHAP is required for this Initiator.
	RequireChap bool `json:"requireChap,"`
	//The unique CHAP username for this initiator.
	ChapUsername string `json:"chapUsername,omitempty"`
	//The CHAP secret used to authenticate the initiator.
	InitiatorSecret string `json:"initiatorSecret,omitempty"`
	//The CHAP secret used to authenticate the target (mutual CHAP authentication).
	TargetSecret string `json:"targetSecret,omitempty"`
}

type PendingActiveNode struct {
	//
	ActiveNodeKey string `json:"activeNodeKey,"`
	//
	PendingActiveNodeID int64 `json:"pendingActiveNodeID,"`
	//
	PendingNodeID int64 `json:"pendingNodeID,"`
	//
	AssignedNodeID int64 `json:"assignedNodeID,"`
	//
	AsyncHandle int64 `json:"asyncHandle,"`
	//IP address used for both intra-cluster and inter-cluster communication.
	Cip string `json:"cip,"`
	//IP address used for the per-node API and UI.
	Mip string `json:"mip,"`
	//IP address used for iSCSI traffic.
	Sip string `json:"sip,"`
	//Information about the node's hardware.
	PlatformInfo Platform `json:"platformInfo,"`
	//The node's role in the cluster. Possible values are Management, Storage, Compute, and Witness.
	Role string `json:"role,"`
	//The version of SolidFire software currently running on this node.
	SoftwareVersion string `json:"softwareVersion,"`
}

type DriveStats struct {
	//
	ActiveSessions int64 `json:"activeSessions,omitempty"`
	//
	DriveID int64 `json:"driveID,omitempty"`
	//
	FailedDieCount int64 `json:"failedDieCount,"`
	//
	LifeRemainingPercent int64 `json:"lifeRemainingPercent,"`
	//
	LifetimeReadBytes int64 `json:"lifetimeReadBytes,"`
	//
	LifetimeWriteBytes int64 `json:"lifetimeWriteBytes,"`
	//
	PowerOnHours int64 `json:"powerOnHours,"`
	//
	ReadBytes int64 `json:"readBytes,omitempty"`
	//
	ReadOps int64 `json:"readOps,omitempty"`
	//
	ReallocatedSectors int64 `json:"reallocatedSectors,"`
	//
	ReserveCapacityPercent int64 `json:"reserveCapacityPercent,"`
	//
	Timestamp string `json:"timestamp,"`
	//
	TotalCapacity int64 `json:"totalCapacity,"`
	//
	UsedCapacity int64 `json:"usedCapacity,omitempty"`
	//
	UsedMemory int64 `json:"usedMemory,"`
	//
	WriteBytes int64 `json:"writeBytes,omitempty"`
	//
	WriteOps int64 `json:"writeOps,omitempty"`
}

type GetDefaultQoSResult struct {
	//Default QoS values.
	Qos VolumeQOS `json:"qos,"`
}

type SetRemoteLoggingHostsResult struct {
}

type ClearClusterFaultsResult struct {
}

type ListSyncJobsResult struct {
	//
	SyncJobs []SyncJob `json:"syncJobs,"`
}

type Node struct {
	//The unique identifier for this node.
	NodeID int64 `json:"nodeID,"`
	//The master service responsible for controlling other services on this node.
	AssociatedMasterServiceID int64 `json:"associatedMasterServiceID,"`
	//
	AssociatedFServiceID int64 `json:"associatedFServiceID,"`
	//
	FibreChannelTargetPortGroup int64 `json:"fibreChannelTargetPortGroup,omitempty"`
	//
	Name string `json:"name,"`
	//Information about the node's hardware.
	PlatformInfo Platform `json:"platformInfo,"`
	//The node's role in the cluster. Possible values are Management, Storage, Compute, and Witness.
	Role string `json:"role,"`
	//The version of SolidFire software currently running on this node.
	SoftwareVersion string `json:"softwareVersion,"`
	//IP address used for both intra-cluster and inter-cluster communication.
	Cip string `json:"cip,"`
	//The machine's name for the "cip" interface.
	Cipi string `json:"cipi,"`
	//IP address used for the per-node API and UI.
	Mip string `json:"mip,"`
	//The machine's name for the "mip" interface.
	Mipi string `json:"mipi,"`
	//IP address used for iSCSI traffic.
	Sip string `json:"sip,"`
	//The machine's name for the "sip" interface.
	Sipi string `json:"sipi,"`
	//UUID of node.
	Uuid string `json:"uuid,"`
	//
	VirtualNetworks []VirtualNetworkAddress `json:"virtualNetworks,"`
	//
	Attributes interface{} `json:"attributes,"`
	//
	NodeSlot string `json:"nodeSlot,omitempty"`
	//Uniquely identifies a chassis, and identical for all nodes in a given chassis.
	ChassisName string `json:"chassisName,"`
	//Uniquely identifies a custom protection domain, identical for all nodes within all chassis in a given custom protection domain.
	CustomProtectionDomainName string `json:"customProtectionDomainName,"`
}

type ModifySnapMirrorEndpointResult struct {
	//Information about the modified SnapMirror endpoint.
	SnapMirrorVolumes SnapMirrorEndpoint `json:"snapMirrorVolumes,"`
}

type ModifyStorageContainerResult struct {
	//
	StorageContainer StorageContainer `json:"storageContainer,"`
}

type ResetSupplementalTlsCiphersResult struct {
}

type RtfiInfo struct {
	//
	Mipi string `json:"mipi,omitempty"`
	//
	Generation string `json:"generation,"`
	//
	StatusUrlLogfile string `json:"statusUrlLogfile,omitempty"`
	//
	Build string `json:"build,"`
	//
	StatusUrlAll string `json:"statusUrlAll,"`
	//
	GenerationNext int64 `json:"generationNext,omitempty"`
	//
	Mip string `json:"mip,omitempty"`
	//
	StatusUrlCurrent string `json:"statusUrlCurrent,"`
	//
	Options interface{} `json:"options,omitempty"`
}

type VirtualVolumeHost struct {
	//
	VirtualVolumeHostID string `json:"virtualVolumeHostID,"`
	//
	ClusterID string `json:"clusterID,omitempty"`
	//
	VisibleProtocolEndpointIDs []string `json:"visibleProtocolEndpointIDs,"`
	//
	Bindings []int64 `json:"bindings,"`
	//
	InitiatorNames []string `json:"initiatorNames,"`
	//
	HostAddress string `json:"hostAddress,omitempty"`
}

type CloneMultipleVolumesResult struct {
	//A value returned from an asynchronous method call.
	AsyncHandle int64 `json:"asyncHandle,"`
	//Unique ID of the new group clone.
	GroupCloneID int64 `json:"groupCloneID,"`
	//List of volumeIDs for the source and destination volume pairs.
	Members []GroupCloneVolumeMember `json:"members,"`
}

type CreateQoSPolicyResult struct {
	//The newly created QoSPolicy object.
	QosPolicy QoSPolicy `json:"qosPolicy,"`
}

type NeedsWorkCopyVolumeToSnapshotResult struct {
}

type RemoveVolumePairResult struct {
}

type GetKeyServerKmipResult struct {
	//A KMIP (Key Management Interoperability Protocol) Key Server which was created previously via CreateKeyServerKmip.
	KmipKeyServer KeyServerKmip `json:"kmipKeyServer,"`
}

type GetNtpInfoResult struct {
	//Indicates whether or not the nodes in the cluster are listening for broadcast NTP messages. Possible values:
	//true
	//false
	Broadcastclient bool `json:"broadcastclient,"`
	//List of NTP servers.
	Servers []string `json:"servers,"`
}

type NeedsWorkSetAptSourceLinesResult struct {
}

type RemoveNodesResult struct {
}

type AsyncHandleResult struct {
	//
	AsyncHandle int64 `json:"asyncHandle,"`
}

type NeedsWorkFindMismatchedSliceLBAsResult struct {
}

type Schedule struct {
	ScheduleName string      `json:"scheduleName,"`
	Monthdays    []int64     `json:"monthdays,omitempty"`
	Weekdays     int64       `json:"weekdays,omitempty"`
	Hours        int64       `json:"hours,"`
	Minutes      int64       `json:"minutes,"`
	ScheduleType string      `json:"scheduleType,"`
	Attributes   interface{} `json:"attributes,"`
	//Indicates whether or not the schedule has errors.
	HasError bool `json:"hasError,omitempty"`
	//Indicates the status of the last scheduled snapshot.
	//Valid values are:
	//Success
	//Failed
	LastRunStatus string `json:"lastRunStatus,omitempty"`
	//Indicates the last time the schedule started n ISO 8601 date string.
	//Valid values are:
	//Success
	//Failed
	LastRunTimeStarted string `json:"lastRunTimeStarted,omitempty"`
	//Indicates whether or not the schedule is paused.
	Paused bool `json:"paused,omitempty"`
	//Indicates whether or not the schedule is recurring.
	Recurring bool `json:"recurring,omitempty"`
	//Indicates whether or not the schedule will run the next time the scheduler is active. When set to "true", the schedule will run the next time the scheduler is active and then reset back to "false".
	RunNextInterval bool `json:"runNextInterval,omitempty"`
	//Unique ID of the schedule
	ScheduleID int64 `json:"scheduleID,omitempty"`
	//Includes the unique name given to the schedule, the retention period for the snapshot that was created, and the volume ID of the volume from which the snapshot was created.
	ScheduleInfo ScheduleInfo `json:"scheduleInfo,"`
	//Indicates the date the first time the schedule began of will begin. Formatted in UTC time.
	StartingDate string `json:"startingDate,omitempty"`
	//Indicates if the schedule is marked for deletion.
	ToBeDeleted bool `json:"toBeDeleted,omitempty"`
	//The snapMirrorLabel to be applied to the created Snapshot or Group Snapshot, contained in the scheduleInfo.
	SnapMirrorLabel string `json:"snapMirrorLabel,omitempty"`
}

type SetDefaultProtectionSchemeResult struct {
	//The default protection scheme for the cluster
	DefaultProtectionScheme ProtectionScheme `json:"defaultProtectionScheme,"`
}

type VolumeAccessGroupLunAssignments struct {
	//Unique volume access group ID for which the LUN assignments will be modified.
	VolumeAccessGroupID int64 `json:"volumeAccessGroupID,"`
	//The volume IDs with assigned LUN values.
	LunAssignments []LunAssignment `json:"lunAssignments,"`
	//The volume IDs with deleted LUN values.
	DeletedLunAssignments []LunAssignment `json:"deletedLunAssignments,"`
}

type GetNetworkConfigResult struct {
	//
	Network Network `json:"network,"`
}

type NeedsWorkDisconnectISCSISessionsResult struct {
}

type GetProtectionSchemesResult struct {
	//The available protection schemes
	ProtectionSchemes ProtectionSchemeInfo `json:"protectionSchemes,"`
}

type ListProtectionDomainLevelsResult struct {
	//A list of the different Protection Domain Levels, where each supplies the
	//cluster's Tolerance and Resiliency information from its own perspective.
	//This will include an element for each of the defined Protection Domain Types.
	ProtectionDomainLevels []ProtectionDomainLevel `json:"protectionDomainLevels,"`
}

type NeedsWorkStopClusterBSCheckResult struct {
}

type RemoveVirtualNetworkResult struct {
}

type ConfigParams struct {
	//
	Cluster ClusterConfig `json:"cluster,"`
	//
	Network NetworkParams `json:"network,"`
}

type DeleteVolumeResult struct {
	//
	Volume Volume `json:"volume,omitempty"`
}

type CreateVolumeResult struct {
	//
	Volume Volume `json:"volume,omitempty"`
	//VolumeID for the newly created volume.
	VolumeID int64 `json:"volumeID,"`
	//The curve is a set of key-value pairs.
	//The keys are I/O sizes in bytes.
	//The values represent the cost of performing an IOP at a specific I/O size.
	//The curve is calculated relative to a 4096 byte operation set at 100 IOPS.
	Curve int64 `json:"curve,"`
}

type DeleteAuthSessionsResult struct {
	//SessionInfos for those auth sessions deleted.
	Sessions []AuthSessionInfo `json:"sessions,"`
}

type FipsErrorNodeReportType struct {
	//Error description
	Error string `json:"error,"`
	//Node ID
	NodeID int64 `json:"nodeID,"`
}

type GetSnmpStateResult struct {
	//If the nodes in the cluster are configured for SNMP.
	Enabled bool `json:"enabled,"`
	//If the node in the cluster is configured for SNMP v3.
	SnmpV3Enabled bool `json:"snmpV3Enabled,"`
}

type ListQoSPoliciesResult struct {
	//A list of details about each QoS policy.
	QosPolicies []QoSPolicy `json:"qosPolicies,"`
}

type AbortRecoverDeadVolumesResult struct {
}

type CreateVolumeAccessGroupResult struct {
	//The ID for the newly-created volume access group.
	VolumeAccessGroupID int64 `json:"volumeAccessGroupID,"`
	//
	VolumeAccessGroup VolumeAccessGroup `json:"volumeAccessGroup,omitempty"`
}

type PairedCluster struct {
	//Name of the other cluster in the pair
	ClusterName string `json:"clusterName,"`
	//Unique ID given to each cluster in the pair.
	ClusterPairID int64 `json:"clusterPairID,"`
	//Universally unique identifier.
	ClusterPairUUID string `json:"clusterPairUUID,"`
	//Number, in milliseconds, of latency between clusters.
	Latency int64 `json:"latency,"`
	//IP of the management connection for paired clusters.
	Mvip string `json:"mvip,"`
	//Can be one of the following:
	//Connected
	//Misconfigured
	//Disconnected
	Status string `json:"status,"`
	//The Element OS version of the other cluster in the pair.
	Version string `json:"version,"`
	//The cluster UUID
	ClusterUUID string `json:"clusterUUID,omitempty"`
}

type GetCurrentClusterAdminResult struct {
	//Information about the calling ClusterAdmin.
	//In case the returned ClusterAdmin object has authMethod value of Ldap or Idp:
	//        The access field may contain data aggregated from multiple LdapAdmins or IdpAdmins.
	//        If this is the case, the clusterAdminId will be set to -1 to indicate that there may not be a unique, 1:1 mapping of the calling ClusterAdmin with a ClusterAdmin on the cluster.
	ClusterAdmin ClusterAdmin `json:"clusterAdmin,"`
}

type ListClusterFaultsResult struct {
	//The list of Cluster Fault objects.
	Faults []ClusterFaultInfo `json:"faults,"`
}

type GetAPIResult struct {
	//
	CurrentVersion string `json:"currentVersion,"`
	//
	SupportedVersions []string `json:"supportedVersions,"`
	V1_0              []string `json:"1.0,"`
	V5_0              []string `json:"5.0,"`
	V7_0              []string `json:"7.0,"`
	V7_1              []string `json:"7.1,"`
	V7_2              []string `json:"7.2,"`
	V7_3              []string `json:"7.3,"`
	V7_4              []string `json:"7.4,"`
	V8_0              []string `json:"8.0,"`
	V8_1              []string `json:"8.1,"`
	V8_2              []string `json:"8.2,"`
	V8_3              []string `json:"8.3,"`
	V8_4              []string `json:"8.4,"`
	V8_5              []string `json:"8.5,"`
	V8_6              []string `json:"8.6,"`
	V8_7              []string `json:"8.7,"`
	V9_0              []string `json:"9.0,"`
	V9_1              []string `json:"9.1,"`
	V9_2              []string `json:"9.2,"`
	V9_3              []string `json:"9.3,"`
	V10_0             []string `json:"10.0,"`
	V10_1             []string `json:"10.1,"`
	V10_2             []string `json:"10.2,"`
}

type GetDriveStatsResult struct {
	//
	DriveStats DriveStats `json:"driveStats,"`
}

type ListVirtualVolumeHostsResult struct {
	//List of known ESX hosts.
	Hosts []VirtualVolumeHost `json:"hosts,"`
}

type NeedsWorkSetDebugOptionsResult struct {
}

type ResetNodeResult struct {
	//
	Details ResetNodeDetails `json:"details,omitempty"`
	//
	Duration string `json:"duration,omitempty"`
	//
	Result string `json:"result,omitempty"`
}

type SetLoginBannerResult struct {
	LoginBanner LoginBanner `json:"loginBanner,"`
}

type AddClusterAdminResult struct {
	//ClusterAdminID for the newly created Cluster Admin.
	ClusterAdminID int64 `json:"clusterAdminID,"`
}

type EventInfo struct {
	//ID of event.
	EventID int64 `json:"eventID,"`
	//Unused
	Severity int64 `json:"severity,"`
	//Event type.
	EventInfoType string `json:"eventInfoType,"`
	//The message associated with the event.
	Message string `json:"message,"`
	//ServiceID associated with the event.
	ServiceID int64 `json:"serviceID,"`
	//NodeID associated with the event.
	NodeID int64 `json:"nodeID,"`
	//Derived from driveIDs field. Either the first item in driveIDs array, or empty.
	DriveID int64 `json:"driveID,"`
	//Drive IDs associated with the event.
	DriveIDs []int64 `json:"driveIDs,"`
	//The time this event was reported.
	TimeOfReport string `json:"timeOfReport,"`
	//The time this event was published into the database.
	TimeOfPublish string `json:"timeOfPublish,"`
	//Data assoicated with the event, such as data report or exception details.
	Details string `json:"details,omitempty"`
}

type NeedsWorkStartUpgradeResult struct {
}

type SupportBundleDetails struct {
	//The name specified in the 'CreateSupportBundle API method. If no name was specified, 'supportbundle' will be used.
	BundleName string `json:"bundleName,"`
	//The arguments passed with this method.
	ExtraArgs string `json:"extraArgs,"`
	//A list of the support bundle files that were created.
	Files []string `json:"files,"`
	//The URL that you can use to download the bundle file(s). Should correspond 1:1 with files list.
	Url []string `json:"url,"`
	//The commands that were run and their output, with newlines replaced by HTML <br> elements.
	Output string `json:"output,"`
	//The timeout specified for the support bundle creation process.
	TimeoutSec int64 `json:"timeoutSec,"`
}

type GetVirtualVolumeUnsharedChunksRequest struct {
	//The ID of the Virtual Volume.
	VirtualVolumeID string `json:"virtualVolumeID,"`
	//The ID of the Virtual Volume to compare against.
	BaseVirtualVolumeID string `json:"baseVirtualVolumeID,"`
	//Start Byte offset.
	SegmentStart int64 `json:"segmentStart,"`
	//Length of the scan segment in bytes.
	SegmentLength int64 `json:"segmentLength,"`
	//Number of bytes represented by one bit in the bitmap.
	ChunkSize int64 `json:"chunkSize,"`
	//
	CallingVirtualVolumeHostID string `json:"callingVirtualVolumeHostID,omitempty"`
}

type ListSnapMirrorVserversResult struct {
	//A list of the SnapMirror Vservers available on the ONTAP storage system.
	SnapMirrorVservers []SnapMirrorVserver `json:"snapMirrorVservers,"`
}

type ListKeyServersKmipResult struct {
	//The complete list of KMIP (Key Management Interoperability Protocol) Key Servers which have been created via CreateKeyServerKmip.
	KmipKeyServers []KeyServerKmip `json:"kmipKeyServers,"`
}

type Snapshot struct {
	//Unique ID for this snapshot.
	SnapshotID int64 `json:"snapshotID,"`
	//The volume this snapshot was taken of.
	VolumeID int64 `json:"volumeID,"`
	//A name for this snapshot.
	//It is not necessarily unique.
	Name string `json:"name,"`
	//A string that represents the correct digits in the stored snapshot.
	//This checksum can be used later to compare other snapshots to detect errors in the data.
	Checksum string `json:"checksum,"`
	//Identifies if snapshot is enabled for remote replication.
	EnableRemoteReplication bool `json:"enableRemoteReplication,"`
	//Indicates how the snapshot expiration was set. Possible values:
	//api: expiration time was set by using the API.
	//none: there is no expiration time set.
	//test: expiration time was set for testing.
	ExpirationReason string `json:"expirationReason,"`
	//The time at which this snapshot will expire and be purged from the cluster.
	ExpirationTime string `json:"expirationTime,omitempty"`
	//Replication status details of the snapshot as seen on the source cluster.
	RemoteStatuses []SnapshotRemoteStatus `json:"remoteStatuses,omitempty"`
	//Current status of the snapshot
	//Preparing: A snapshot that is being prepared for use and is not yet writable.
	//Done: A snapshot that has finished being prepared and is now usable.
	//Active: This snapshot is the active branch.
	Status string `json:"status,"`
	//Universal Unique ID of an existing snapshot.
	SnapshotUUID string `json:"snapshotUUID,"`
	//Total size of this snapshot in bytes.
	//It is equivalent to totalSize of the volume when this snapshot was taken.
	TotalSize int64 `json:"totalSize,"`
	//If present, the ID of the group this snapshot is a part of.
	//If not present, this snapshot is not part of a group.
	GroupID int64 `json:"groupID,omitempty"`
	//Universal Unique ID of the group this snapshot is part of.
	GroupSnapshotUUID string `json:"groupSnapshotUUID,"`
	//The time this snapshot was taken.
	CreateTime string `json:"createTime,"`
	//
	InstanceCreateTime string `json:"instanceCreateTime,"`
	//
	VolumeName string `json:"volumeName,"`
	//
	InstanceSnapshotUUID string `json:"instanceSnapshotUUID,"`
	//The ID of the virtual volume with which the snapshot is associated.
	VirtualVolumeID string `json:"virtualVolumeID,omitempty"`
	//List of Name/Value pairs in JSON object format.
	Attributes interface{} `json:"attributes,"`
	//Label used by SnapMirror software to specify snapshot retention policy on SnapMirror endpoint.
	SnapMirrorLabel string `json:"snapMirrorLabel,omitempty"`
}

type NeedsWorkGetThreadBacktracesResult struct {
}

type SetProtectionDomainLayoutChassisOverrideResult struct {
	//How all of the nodes are grouped into different ProtectionDomains.
	ProtectionDomainLayout ProtectionDomainLayout `json:"protectionDomainLayout,"`
}

type GetNodeSupportedTlsCiphersResult struct {
	//List of mandatory TLS cipher suites for the node.
	//Mandatory ciphers are those ciphers which will always be active on the node.
	MandatoryCiphers []string `json:"mandatoryCiphers,"`
	//List of default supplemental TLS cipher suites for the node.
	//The supplemental ciphers will be restored to this list when the ResetNodeSupplementalTlsCiphers command is run.
	DefaultSupplementalCiphers []string `json:"defaultSupplementalCiphers,"`
	//List of available supplemental TLS cipher suites which can be configured with the SetNodeSupplementalTlsCiphers command.
	SupportedSupplementalCiphers []string `json:"supportedSupplementalCiphers,"`
}

type NeedsWorkDeleteSnapMirrorObjectAttributesResult struct {
}

type NeedsWorkModifySliceReserveUsedThresholdPctResult struct {
}

type AddBlockToBSResult struct {
	//A map from block service id to an exception string
	//for all of the block services that failed to add the block.
	Failed string `json:"failed,omitempty"`
	//A map from block service id to a boolean (always true)
	//for all of the block services that successfully added the block.
	Success bool `json:"success,omitempty"`
}

type CreateKeyServerKmipResult struct {
	//The KMIP (Key Management Interoperability Protocol) Key Server which has been created.
	KmipKeyServer KeyServerKmip `json:"kmipKeyServer,"`
}

type GetConfigResult struct {
	//The details of the cluster. Values returned in "config": cluster- Cluster information that identifies how the node communicates with the cluster it is associated with. (Object) network - Network information for bonding and Ethernet connections. (Object)
	Config Config `json:"config,"`
}

type GetSystemStatusResult struct {
	//
	RebootRequired bool `json:"rebootRequired,"`
}

type SetNetworkConfigResult struct {
	//
	Network Network `json:"network,"`
}

type TestDrivesResult struct {
	//
	Details string `json:"details,"`
	//
	Duration string `json:"duration,"`
	//
	Result string `json:"result,"`
}

type ClusterHardwareInfo struct {
	//
	Drives DriveHardwareInfo `json:"drives,"`
	//
	Nodes interface{} `json:"nodes,"`
}

type CompleteVolumePairingResult struct {
}

type NodeDriveHardware struct {
	//
	NodeID int64 `json:"nodeID,"`
	//
	Result DrivesHardware `json:"result,"`
}

type GetLldpConfigResult struct {
	//Enable the LLDP service
	EnableLldp bool `json:"enableLldp,omitempty"`
	//Enable MED, an extension to LLDP that provides inventory information
	EnableMed bool `json:"enableMed,omitempty"`
	//Enable other discovery protocols: CDP, FDP, EDP, and SONMP.
	EnableOtherProtocols bool `json:"enableOtherProtocols,omitempty"`
}

type NeedsWorkFillSliceResult struct {
}

type SoftwareVersionInfo struct {
	//
	CurrentVersion string `json:"currentVersion,"`
	//
	NodeID int64 `json:"nodeID,"`
	//
	PackageName string `json:"packageName,"`
	//
	PendingVersion string `json:"pendingVersion,"`
	//
	StartTime string `json:"startTime,"`
}

type AddedNode struct {
	//
	NodeID int64 `json:"nodeID,omitempty"`
	//
	PendingNodeID int64 `json:"pendingNodeID,"`
	//
	ActiveNodeKey string `json:"activeNodeKey,omitempty"`
	//
	AssignedNodeID int64 `json:"assignedNodeID,omitempty"`
	//
	AsyncHandle int64 `json:"asyncHandle,omitempty"`
	//
	Cip string `json:"cip,omitempty"`
	//
	Mip string `json:"mip,omitempty"`
	//
	PlatformInfo Platform `json:"platformInfo,omitempty"`
	//
	Sip string `json:"sip,omitempty"`
	//
	SoftwareVersion string `json:"softwareVersion,omitempty"`
}

type NodeProtectionDomains struct {
	//The unique identifier for the node.
	NodeID int64 `json:"nodeID,"`
	//The Protection Domains of which the Node is a member.
	ProtectionDomains []ProtectionDomain `json:"protectionDomains,"`
}

type ListBulkVolumeJobsResult struct {
	//An array of information for each bulk volume job.
	BulkVolumeJobs []BulkVolumeJob `json:"bulkVolumeJobs,"`
}

type PromoteClusterMasterResult struct {
	//The old cluster master's NodeID.
	OldClusterMasterNodeID int64 `json:"oldClusterMasterNodeID,"`
	//The old cluster master's ServiceID.
	OldClusterMasterServiceID int64 `json:"oldClusterMasterServiceID,"`
	//The new cluster master's NodeID.
	RequestedClusterMasterNodeID int64 `json:"requestedClusterMasterNodeID,"`
	//The new cluster master's ServiceID.
	RequestedClusterMasterServiceID int64 `json:"requestedClusterMasterServiceID,"`
	//The new cluster master's sequence number.
	TargetSequenceNumber int64 `json:"targetSequenceNumber,"`
}

type DemoteClusterMasterResult struct {
	//The old cluster master's NodeID.
	NodeID int64 `json:"nodeID,"`
	//The old cluster master's ServiceID.
	ServiceID int64 `json:"serviceID,"`
	//The old cluster master's sequence number.
	SequenceNumber int64 `json:"sequenceNumber,"`
}

type ListNetworkInterfacesResult struct {
	//
	Interfaces []NetworkInterface `json:"interfaces,"`
}

type ListSnapMirrorEndpointsResult struct {
	//A list of existing SnapMirror endpoints.
	SnapMirrorEndpoints []SnapMirrorEndpoint `json:"snapMirrorEndpoints,"`
}

type NeedsWorkBDTPipelineResult struct {
}

type AddLdapClusterAdminResult struct {
	//
	ClusterAdminID int64 `json:"clusterAdminID,omitempty"`
}

type ListAccountsResult struct {
	//List of accounts.
	Accounts []Account `json:"accounts,"`
}

type FipsNodeReportType struct {
	//FIPS https feature status.
	HttpsEnabled bool `json:"httpsEnabled,"`
	//Node's FipsDrives capability status.
	FipsDrives FipsDrivesStatusType `json:"fipsDrives,"`
	//Node ID.
	NodeID int64 `json:"nodeID,"`
}

type NeedsWorkReconfigureEnsembleResult struct {
}

type RecoverDeadVolumesResult struct {
}

type StorageContainer struct {
	//
	Name string `json:"name,"`
	//
	StorageContainerID string `json:"storageContainerID,"`
	//
	AccountID int64 `json:"accountID,"`
	//
	ProtocolEndpointType string `json:"protocolEndpointType,"`
	//
	InitiatorSecret string `json:"initiatorSecret,"`
	//
	TargetSecret string `json:"targetSecret,"`
	//
	Status string `json:"status,"`
	//
	VirtualVolumes []string `json:"virtualVolumes,omitempty"`
}

type ClusterFaultInfo struct {
	//
	DriveIDs []int64 `json:"driveIDs,omitempty"`
	//
	NetworkInterface string `json:"networkInterface,omitempty"`
	//
	Severity string `json:"severity,"`
	//
	Type string `json:"type,"`
	//
	Code string `json:"code,"`
	//
	Details string `json:"details,"`
	//
	NodeHardwareFaultID int64 `json:"nodeHardwareFaultID,"`
	//
	NodeID int64 `json:"nodeID,"`
	//
	ServiceID int64 `json:"serviceID,"`
	//
	DriveID int64 `json:"driveID,"`
	//
	Resolved bool `json:"resolved,"`
	//
	ClusterFaultID int64 `json:"clusterFaultID,"`
	//
	Date string `json:"date,"`
	//
	ResolvedDate string `json:"resolvedDate,"`
	//
	Data interface{} `json:"data,omitempty"`
	//
	ExternalSource string `json:"externalSource,omitempty"`
}

type EnableStandbySliceAssignmentsResult struct {
}

type LunAssignment struct {
	//The volume ID assigned to the Lun.
	VolumeID int64 `json:"volumeID,"`
	//Correct LUN values are 0 - 16383. An exception will be seen if an incorrect LUN value is passed.
	Lun int64 `json:"lun,"`
}

type NeedsWorkGetReportResult struct {
}

type DisableIdpAuthenticationResult struct {
}

type GetClusterFullThresholdResult struct {
	//Current computed level of block fullness of the cluster.
	//Possible values: stage1Happy: No alerts or error conditions. stage2Aware: 3 nodes of capacity available. stage3Low: 2 nodes of capacity available. stage4Critical: 1 node of capacity available. No new volumes or clones can be created. stage5CompletelyConsumed: Completely consumed. Cluster is read-only, iSCSI connection is maintained but all writes are suspended.
	BlockFullness string `json:"blockFullness,"`
	//Reflects the highest level of fullness between "blockFullness" and "metadataFullness".
	Fullness string `json:"fullness,"`
	//A value representative of the number of times metadata space can be over provisioned relative to the amount of space available. For example, if there was enough metadata space to store 100 TiB of volumes and this number was set to 5, then 500 TiB worth of volumes could be created.
	MaxMetadataOverProvisionFactor int64 `json:"maxMetadataOverProvisionFactor,"`
	//Current computed level of metadata fullness of the cluster.
	MetadataFullness string `json:"metadataFullness,"`
	//Error condition; message sent to "Alerts" if the reserved slice utilization is greater than the sliceReserveUsedThresholdPct value returned.
	SliceReserveUsedThresholdPct int64 `json:"sliceReserveUsedThresholdPct,"`
	//Awareness condition: Value that is set for "Stage 2" cluster threshold level.
	Stage2AwareThreshold float64 `json:"stage2AwareThreshold,"`
	//Awareness condition: Value that is set for "Stage 3" cluster threshold level.
	Stage3LowThresholdInternalUseOnly float64 `json:"stage3LowThresholdInternalUseOnly,"`
	//Awareness condition: Value that is set for "Stage 4" cluster threshold level.
	Stage4CriticalThresholdInternalUseOnly float64 `json:"stage4CriticalThresholdInternalUseOnly,"`
	//Number of bytes being used by the cluster at which a stage2 condition will exist.
	Stage2BlockThresholdBytes int64 `json:"stage2BlockThresholdBytes,"`
	//Number of bytes being used by the cluster at which a stage3 condition will exist.
	Stage3BlockThresholdBytes int64 `json:"stage3BlockThresholdBytes,"`
	//The percent value set for stage3 of block fullness. At this percent full, a warning will be posted in the Alerts log.
	Stage3BlockThresholdPercent int64 `json:"stage3BlockThresholdPercent,"`
	//The percent value set for stage3 of metadata fullness. At this percent full, a warning will be posted in the Alerts log.
	Stage3MetadataThresholdPercent int64 `json:"stage3MetadataThresholdPercent,"`
	//Error condition; message sent to "Alerts" that capacity on a cluster is getting low.
	Stage3LowThreshold int64 `json:"stage3LowThreshold,"`
	//Error condition; message sent to "Alerts" that capacity on a cluster is critically low.
	Stage4CriticalThreshold int64 `json:"stage4CriticalThreshold,"`
	//Number of bytes being used by the cluster at which a stage4 condition will exist.
	Stage4BlockThresholdBytes int64 `json:"stage4BlockThresholdBytes,"`
	//Number of bytes being used by the cluster at which a stage5 condition will exist.
	Stage5BlockThresholdBytes int64 `json:"stage5BlockThresholdBytes,"`
	//Physical capacity of the cluster measured in bytes.
	SumTotalClusterBytes int64 `json:"sumTotalClusterBytes,"`
	//Total amount of space that can be used to store metadata.
	SumTotalMetadataClusterBytes int64 `json:"sumTotalMetadataClusterBytes,"`
	//Number of bytes used on the cluster.
	SumUsedClusterBytes int64 `json:"sumUsedClusterBytes,"`
	//Amount of space used on volume drives to store metadata.
	SumUsedMetadataClusterBytes int64 `json:"sumUsedMetadataClusterBytes,"`
	//Number of metadata bytes being used by the cluster at which a stage2 condition will exist.
	Stage2MetadataThresholdBytes int64 `json:"stage2MetadataThresholdBytes,"`
	//Number of metadata bytes being used by the cluster at which a stage3 condition will exist.
	Stage3MetadataThresholdBytes int64 `json:"stage3MetadataThresholdBytes,"`
	//Number of metadata bytes being used by the cluster at which a stage4 condition will exist.
	Stage4MetadataThresholdBytes int64 `json:"stage4MetadataThresholdBytes,"`
	//Number of metadata bytes being used by the cluster at which a stage5 condition will exist.
	Stage5MetadataThresholdBytes int64 `json:"stage5MetadataThresholdBytes,"`
}

type ListSnapMirrorRelationshipsResult struct {
	//A list of objects containing information about SnapMirror relationships.
	SnapMirrorRelationships []SnapMirrorRelationship `json:"snapMirrorRelationships,"`
}

type NeedsWorkSetUpgradeNodeIdResult struct {
}

type VirtualVolumeUnbindResult struct {
	//
	UnbindResults []UnbindResult `json:"unbindResults,"`
}

type GetKeyProviderKmipResult struct {
	//A KMIP (Key Management Interoperability Protocol) Key Provider which was created previously via CreateKeyProviderKmip.
	KmipKeyProvider KeyProviderKmip `json:"kmipKeyProvider,"`
}

type ListClusterPairsResult struct {
	//Information about each paired cluster.
	ClusterPairs []PairedCluster `json:"clusterPairs,"`
}

type GetProtectionDomainLayoutResult struct {
	//How all of the nodes are grouped into different ProtectionDomains.
	ProtectionDomainLayout ProtectionDomainLayout `json:"protectionDomainLayout,"`
}

type GetVirtualVolumeUnsharedBitmapRequest struct {
	//The ID of the Virtual Volume.
	VirtualVolumeID string `json:"virtualVolumeID,"`
	//The ID of the Virtual Volume to compare against.
	BaseVirtualVolumeID string `json:"baseVirtualVolumeID,"`
	//Byte offset.
	SegmentStart int64 `json:"segmentStart,"`
	//Byte length adjusted to end on a chunk boundary.
	SegmentLength int64 `json:"segmentLength,"`
	//Number of bytes represented by one bit in the bitmap.
	ChunkSize int64 `json:"chunkSize,"`
	//
	CallingVirtualVolumeHostID string `json:"callingVirtualVolumeHostID,omitempty"`
}

type ListEventsResult struct {
	//event queue type
	EventQueueType string `json:"eventQueueType,"`
	//list of events
	Events []EventInfo `json:"events,"`
}

type NeedsWorkMovePrimariesAwayFromNodeResult struct {
}

type NodeStatsInfo struct {
	//Bytes in on the cluster interface.
	CBytesIn int64 `json:"cBytesIn,"`
	//Bytes out on the cluster interface.
	CBytesOut int64 `json:"cBytesOut,"`
	//
	Count int64 `json:"count,"`
	//CPU Usage %
	Cpu int64 `json:"cpu,"`
	//CPU Total
	CpuTotal int64 `json:"cpuTotal,"`
	//Bytes in on the management interface.
	MBytesIn int64 `json:"mBytesIn,"`
	//Bytes out on the management interface.
	MBytesOut int64 `json:"mBytesOut,"`
	//Network interface utilization (in %) for the cluster network interface.
	NetworkUtilizationCluster int64 `json:"networkUtilizationCluster,"`
	//Network interface utilization (in %) for the storage network interface.
	NetworkUtilizationStorage int64 `json:"networkUtilizationStorage,"`
	//
	NodeID int64 `json:"nodeID,"`
	//Read Operations.
	ReadOps int64 `json:"readOps,"`
	//
	ReadLatencyUSecTotal int64 `json:"readLatencyUSecTotal,"`
	//Bytes in on the storage interface.
	SBytesIn int64 `json:"sBytesIn,"`
	//Bytes out on the storage interface.
	SBytesOut int64 `json:"sBytesOut,"`
	//A histogram of SS load measurements.
	SsLoadHistogram QuintileHistogram `json:"ssLoadHistogram,"`
	//Current time in UTC format ISO 8691 date string.
	Timestamp string `json:"timestamp,"`
	//Total memory usage in bytes.
	UsedMemory int64 `json:"usedMemory,"`
	//
	WriteLatencyUSecTotal int64 `json:"writeLatencyUSecTotal,"`
	//Write Operations
	WriteOps int64 `json:"writeOps,"`
}

type FibreChannelPortInfoResult struct {
	//Used to return information about the Fibre Channel ports.
	Result FibreChannelPortList `json:"result,"`
}

type GetAsyncResultResult struct {
	//The resulting message for the original method call if the call was completed successfully.
	Result AsyncResult `json:"result,"`
	//Status of the asynchronous method call
	//running: The method is still running.
	//complete: The method is complete and the result or error is available.
	Status string `json:"status,"`
}

type ProtectionSchemeInfo struct {
	//The category of the protection scheme.
	Category ProtectionSchemeCategory `json:"category,"`
	//The total number of replicas used by the protection scheme.
	RepCount int64 `json:"repCount,"`
	//The public visibility of the scheme.
	Visibility ProtectionSchemeVisibility `json:"visibility,"`
}

type NeedsWorkRebalanceSlicesResult struct {
}

type AddKeyServerToProviderKmipResult struct {
}

type NeedsWorkListSliceBranchesByServiceResult struct {
}

type DeleteKeyServerKmipResult struct {
}

type NeedsWorkSetClusterStructureResult struct {
}

type ListVolumeStatsByVirtualVolumeResult struct {
	//
	VolumeStats []VirtualVolumeStats `json:"volumeStats,"`
}

type ModifyClusterAdminResult struct {
}

type NeedsWorkGetCodeTimingsResult struct {
}

type ShutdownResult struct {
	//
	Failed []int64 `json:"failed,"`
	//
	Successful []int64 `json:"successful,"`
}

type CreateSupportBundleResult struct {
	//The details of the support bundle.
	Details SupportBundleDetails `json:"details,"`
	//The amount of time required to create the support bundle in the format HH:MM:SS.ssssss
	Duration string `json:"duration,"`
	//Whether the support bundle creation passed of failed.
	Result string `json:"result,"`
}

type DisableStandbySliceAssignmentsResult struct {
}

type NeedsWorkGetDebugOptionsResult struct {
}

type NeedsWorkGetEnsembleResult struct {
}

type NeedsWorkSetClusterSettingsResult struct {
}

type SnapMirrorVolumeInfo struct {
	//The type of volume.
	//Possible values:
	//solidfire: The volume resides on a SolidFire cluster.
	//ontap:     The volume resides on a remote ONTAP cluster.
	Type string `json:"type,"`
	//The ID of the volume. Only valid if "type" is solidfire.
	VolumeID int64 `json:"volumeID,"`
	//The name of the Vserver that owns this volume. Only valid if "type" is ONTAP.
	Vserver string `json:"vserver,"`
	//The name of the volume.
	Name string `json:"name,"`
}

type CreateSnapMirrorEndpointResult struct {
	//The newly created SnapMirror endpoint.
	SnapMirrorEndpoint SnapMirrorEndpoint `json:"snapMirrorEndpoint,"`
}

type GetActiveTlsCiphersResult struct {
	//List of mandatory TLS cipher suites for the cluster.
	MandatoryCiphers []string `json:"mandatoryCiphers,"`
	//List of supplemental TLS cipher suites for the cluster.
	SupplementalCiphers []string `json:"supplementalCiphers,"`
}

type GetLoginBannerResult struct {
	LoginBanner LoginBanner `json:"loginBanner,"`
}

type PendingOperation struct {
	//true: operation is still in progress.
	//false: operation is no integerer in progress.
	Pending bool `json:"pending,"`
	//Name of operation that is in progress or has completed.
	Operation string `json:"operation,"`
}

type RemoveBackupTargetResult struct {
}

type SnmpTrapRecipient struct {
	//The IP address or host name of the target network management station.
	Host string `json:"host,"`
	//SNMP community string.
	Community string `json:"community,"`
	//The UDP port number on the host where the trap is to be sent. Valid range is 1 - 65535. 0 (zero) is not a valid port number. Default is 162.
	Port int64 `json:"port,"`
}

type ClusterAdmin struct {
	//Method in which the cluster admin can be authenticated.
	AuthMethod string `json:"authMethod,"`
	//Controls which methods this cluster admin can use. For more details, see Access Control in the Element API Reference Guide.
	Access []string `json:"access,"`
	//Unique identifier for the cluster admin
	ClusterAdminID int64 `json:"clusterAdminID,"`
	//Username, LDAP DN, or SAML Attribute for the cluster admin.
	Username string `json:"username,"`
	//List of Name/Value pairs in JSON object format.
	Attributes interface{} `json:"attributes,omitempty"`
}

type GetClusterHardwareInfoResult struct {
	//Hardware information for all nodes and drives in the cluster. Each object in this output is labeled with the nodeID of the given node.
	ClusterHardwareInfo ClusterHardwareInfo `json:"clusterHardwareInfo,"`
}

type UpdateBulkVolumeStatusResult struct {
	//Status of the session requested. Returned status:
	//preparing
	//active
	//done
	//failed
	Status string `json:"status,"`
	//The URL to access the node's web server provided only if the session is still active.
	Url string `json:"url,"`
	//Returns attributes that were specified in the BulkVolumeStatusUpdate method. Values are returned if they have changed or not.
	Attributes interface{} `json:"attributes,"`
}

type AuthConfiguration struct {
	//Version of the data currently stored in the database.  If no configuration has been set, then this will be zero.
	Version int64 `json:"version,"`
	//A string previously set by SetAuthConfiguration.  If no configuration has been set, then this will be empty.
	Configuration string `json:"configuration,"`
}

type CreateStorageContainerResult struct {
	//
	StorageContainer StorageContainer `json:"storageContainer,"`
}

type GetClusterStatsResult struct {
	//
	ClusterStats ClusterStats `json:"clusterStats,"`
}

type GetIpmiConfigNodesResult struct {
	//
	NodeID int64 `json:"nodeID,"`
	//
	Result interface{} `json:"result,"`
}

type NeedsWorkSetInitiatorIDResult struct {
}

type TestKeyProviderKmipResult struct {
}

type UpdateSnapMirrorRelationshipResult struct {
	//An object containg information about the updated SnapMirror relationship.
	SnapMirrorRelationship SnapMirrorRelationship `json:"snapMirrorRelationship,"`
}

type CreateSnapMirrorRelationshipResult struct {
	//Information about the newly created SnapMirror relationship.
	SnapMirrorRelationship SnapMirrorRelationship `json:"snapMirrorRelationship,"`
}

type GetVolumeAccessGroupLunAssignmentsResult struct {
	//List of all physical Fibre Channel ports, or a port for a single node.
	VolumeAccessGroupLunAssignments VolumeAccessGroupLunAssignments `json:"volumeAccessGroupLunAssignments,"`
}

type Platform struct {
	//SolidFire's name for this platform.
	NodeType string `json:"nodeType,"`
	//Name of the chassis (example: "R620").
	ChassisType string `json:"chassisType,"`
	//The model of CPU used on this platform.
	CpuModel string `json:"cpuModel,"`
	//The amount of memory on this platform in GiB.
	NodeMemoryGB int64 `json:"nodeMemoryGB,"`
	//
	PlatformConfigVersion string `json:"platformConfigVersion,omitempty"`
}

type SnapshotVirtualVolumeRequest struct {
	//The ID of the Virtual Volume to clone.
	VirtualVolumeID string `json:"virtualVolumeID,"`
	//Number of seconds to complete or fail.
	Timeout int64 `json:"timeout,"`
	//
	Metadata interface{} `json:"metadata,omitempty"`
	//
	CallingVirtualVolumeHostID string `json:"callingVirtualVolumeHostID,omitempty"`
}

type UnbindResult struct {
	//
	Fault string `json:"fault,"`
}

type ListDrivesResult struct {
	//Information for the drives that are connected to the cluster.
	Drives []DriveInfo `json:"drives,"`
}

type NeedsWorkDisableBmcColdResetResult struct {
}

type ClusterStats struct {
	//The amount of cluster capacity being utilized.
	ClusterUtilization float64 `json:"clusterUtilization,"`
	//
	ClientQueueDepth int64 `json:"clientQueueDepth,"`
	//
	NormalizedIOPS int64 `json:"normalizedIOPS,"`
	//Total bytes read by clients.
	ReadBytes int64 `json:"readBytes,"`
	//
	ReadLatencyUSecTotal int64 `json:"readLatencyUSecTotal,"`
	//Total read operations.
	ReadOps int64 `json:"readOps,"`
	//Services count
	ServicesCount int64 `json:"servicesCount,"`
	//Total services.
	ServicesTotal int64 `json:"servicesTotal,"`
	//Current time in UTC format. ISO 8601 date string.
	Timestamp string `json:"timestamp,"`
	//Total bytes written by clients.
	WriteBytes int64 `json:"writeBytes,"`
	//
	WriteLatencyUSecTotal int64 `json:"writeLatencyUSecTotal,"`
	//Total write operations.
	WriteOps int64 `json:"writeOps,"`
	//
	ActualIOPS int64 `json:"actualIOPS,omitempty"`
	//
	AverageIOPSize int64 `json:"averageIOPSize,omitempty"`
	//
	LatencyUSec int64 `json:"latencyUSec,omitempty"`
	//
	ReadBytesLastSample int64 `json:"readBytesLastSample,omitempty"`
	//
	ReadLatencyUSec int64 `json:"readLatencyUSec,omitempty"`
	//
	ReadOpsLastSample int64 `json:"readOpsLastSample,omitempty"`
	//
	SamplePeriodMsec int64 `json:"samplePeriodMsec,omitempty"`
	//
	UnalignedReads int64 `json:"unalignedReads,omitempty"`
	//
	UnalignedWrites int64 `json:"unalignedWrites,omitempty"`
	//
	WriteBytesLastSample int64 `json:"writeBytesLastSample,omitempty"`
	//
	WriteLatencyUSec int64 `json:"writeLatencyUSec,omitempty"`
	//
	WriteOpsLastSample int64 `json:"writeOpsLastSample,omitempty"`
}

type GetVirtualVolumeCountResult struct {
	//The number of virtual volumes currently in the system.
	Count int64 `json:"count,"`
}

type ListSchedulesResult struct {
	//The list of schedules currently on the cluster.
	Schedules []Schedule `json:"schedules,"`
}

type ModifyVolumePairResult struct {
}

type NeedsWorkDeleteVolumeSnapMirrorObjectAttributesResult struct {
}

type NeedsWorkListSnapMirrorObjectAttributesResult struct {
}

type ResumeSnapMirrorRelationshipResult struct {
	//An object containg information about the resumed SnapMirror relationship.
	SnapMirrorRelationship SnapMirrorRelationship `json:"snapMirrorRelationship,"`
}

type DeleteSnapMirrorRelationshipsResult struct {
	//If the delete action succeeded, this object
	//contains a success message. If the action failed,
	//it contains an error message.
	Result string `json:"result,"`
}

type DriveHardwareInfo struct {
	//
	Description string `json:"description,"`
	//
	Dev string `json:"dev,"`
	//
	Devpath string `json:"devpath,"`
	//
	DriveSecurityAtMaximum bool `json:"driveSecurityAtMaximum,"`
	//
	DriveSecurityFrozen bool `json:"driveSecurityFrozen,"`
	//
	DriveSecurityLocked bool `json:"driveSecurityLocked,"`
	//
	Logicalname string `json:"logicalname,"`
	//
	Product string `json:"product,"`
	//
	ScsiCompatID string `json:"scsiCompatID,"`
	//
	SecurityFeatureEnabled bool `json:"securityFeatureEnabled,"`
	//
	SecurityFeatureSupported bool `json:"securityFeatureSupported,"`
	//
	Serial string `json:"serial,"`
	//
	Size int64 `json:"size,"`
	//
	Uuid string `json:"uuid,"`
	//
	Version string `json:"version,"`
}

type ListSnapMirrorNetworkInterfacesResult struct {
	//A list of the SnapMirror network interfaces available on the remote ONTAP storage system.
	SnapMirrorNetworkInterfaces []SnapMirrorNetworkInterface `json:"snapMirrorNetworkInterfaces,"`
}

type ModifyInitiator struct {
	//The numeric ID of the initiator to modify.
	InitiatorID int64 `json:"initiatorID,"`
	//A new friendly name to assign to the initiator.
	Alias string `json:"alias,omitempty"`
	//The ID of the volume access group to which the newly created initiator should be added.
	//If the initiator was previously in a different volume access group, it is removed from the old volume access group.
	//If this key is present but null, the initiator is removed from its current volume access group
	//but not placed in any new volume access group.
	VolumeAccessGroupID int64 `json:"volumeAccessGroupID,omitempty"`
	//A new set of JSON attributes assigned to this initiator.
	Attributes interface{} `json:"attributes,omitempty"`
	//Allows the user to complete the command during an upgrade.
	ForceDuringUpgrade bool `json:"forceDuringUpgrade,omitempty"`
	//"requireChap" determines if the initiator is required to use CHAP during session login. CHAP is optional if "requireChap" is false.
	RequireChap bool `json:"requireChap,omitempty"`
	//The CHAP username for this initiator. Defaults to the initiator name (IQN) if not specified during creation and "requireChap" is true.
	ChapUsername string `json:"chapUsername,omitempty"`
	//The CHAP secret used for authentication of the initiator. Defaults to a randomly generated secret if not specified during creation and "requireChap" is true.
	InitiatorSecret string `json:"initiatorSecret,omitempty"`
	//The CHAP secret used for authentication of the target. Defaults to a randomly generated secret if not specified during creation and "requireChap" is true.
	TargetSecret string `json:"targetSecret,omitempty"`
}

type NeedsWorkSetRsyslogInfoResult struct {
}

type ResetNodeDetails struct {
	//
	RtfiInfo RtfiInfo `json:"rtfiInfo,"`
}

type DeleteClusterInterfacePreferenceResult struct {
}

type GetIpmiConfigResult struct {
	//
	Nodes []GetIpmiConfigNodesResult `json:"nodes,"`
}

type GetEfficiencyResult struct {
	//The amount of space being saved by compressing data on a single volume. Stated as a ratio where "1" means data has been stored without being compressed.
	Compression float64 `json:"compression,omitempty"`
	//The amount of space being saved on a single volume by not duplicating data. Stated as a ratio.
	Deduplication float64 `json:"deduplication,omitempty"`
	//The ratio of space used to the amount of space allocated for storing data. Stated as a ratio.
	ThinProvisioning float64 `json:"thinProvisioning,omitempty"`
	//The last time efficiency data was collected after Garbage Collection (GC). ISO 8601 data string.
	Timestamp string `json:"timestamp,"`
	//The volumes that could not be queried for efficiency data. Missing volumes can be caused by GC being less than hour old, temporary network loss or restarted services since the GC cycle.
	MissingVolumes []int64 `json:"missingVolumes,"`
}

type GetOriginResult struct {
	//
	Nodes []GetOriginNode `json:"nodes,"`
}

type NeedsWorkGetBinAssignmentsResult struct {
}

type NeedsWorkGetFibreChannelVolumeAccessInfoResult struct {
}

type NeedsWorkPairClusterResult struct {
}

type NetworkInterfaces struct {
	//
	Interfaces []NetworkInterface `json:"interfaces,"`
}

type CopyDiffsToVirtualVolumeRequest struct {
	//The ID of the snapshot Virtual Volume.
	VirtualVolumeID string `json:"virtualVolumeID,"`
	//The ID of the base Virtual Volume.
	BaseVirtualVolumeID string `json:"baseVirtualVolumeID,"`
	//The ID of the Virtual Volume to be overwritten.
	DstVirtualVolumeID string `json:"dstVirtualVolumeID,"`
	//
	CallingVirtualVolumeHostID string `json:"callingVirtualVolumeHostID,omitempty"`
}

type DeleteVolumesResult struct {
	//Information about the newly deleted volume.
	Volumes []Volume `json:"volumes,"`
}

type SetSnmpACLResult struct {
}

type NodeStateInfo struct {
	//Name of the cluster.
	Cluster string `json:"cluster,"`
	//<strong>Available:</strong> Node has not been configured with a cluster name.<br><strong>Pending:</strong> Node is pending for a specific named cluster and can be added.<br><strong>Active:</strong> Node is active and a member of a cluster and may not be added to another cluster.
	State string `json:"state,"`
}

type RemoveBlockFromBSResult struct {
	//A map from block service id to an exception string
	//for all of the block services that failed to remove the block.
	Failed string `json:"failed,omitempty"`
	//A map from block service id to a boolean (always true)
	//for all of the block services that successfully removed the block.
	Success bool `json:"success,omitempty"`
}

type TestConnectSvipResult struct {
	//Information about the test operation
	Details TestConnectSvipDetails `json:"details,"`
	//The length of time required to run the test.
	Duration string `json:"duration,"`
	//The results of the entire test
	Result string `json:"result,"`
}

type NeedsWorkDeleteSnapshotSnapMirrorObjectAttributesResult struct {
}

type TestConnectEnsembleResult struct {
	//
	Details TestConnectEnsembleDetails `json:"details,"`
	//The length of time required to run the test.
	Duration string `json:"duration,"`
	//The results of the entire test
	Result string `json:"result,"`
}

type NeedsWorkGetDatabaseEntryResult struct {
}

type DeleteSnapshotResult struct {
}

type DriveConfigInfo struct {
	//
	CanonicalName string `json:"canonicalName,"`
	//
	Connected bool `json:"connected,"`
	//
	Dev int64 `json:"dev,"`
	//
	DevPath string `json:"devPath,"`
	//
	DriveType string `json:"driveType,"`
	//
	Product string `json:"product,"`
	//
	Name string `json:"name,"`
	//
	Path string `json:"path,"`
	//
	PathLink string `json:"pathLink,"`
	//
	ScsiCompatId string `json:"scsiCompatId,"`
	//
	SmartSsdWriteCapable bool `json:"smartSsdWriteCapable,omitempty"`
	//
	SecurityEnabled bool `json:"securityEnabled,"`
	//
	SecurityFrozen bool `json:"securityFrozen,"`
	//
	SecurityLocked bool `json:"securityLocked,"`
	//
	SecuritySupported bool `json:"securitySupported,"`
	//
	Size int64 `json:"size,"`
	//
	Slot int64 `json:"slot,"`
	//
	Uuid string `json:"uuid,"`
	//
	Vendor string `json:"vendor,"`
	//
	Version string `json:"version,"`
	//
	SecurityAtMaximum bool `json:"securityAtMaximum,"`
	//
	Serial string `json:"serial,"`
	//
	ScsiState string `json:"scsiState,"`
}

type Frequency struct {
}

type NeedsWorkGetRsyslogInfoResult struct {
}

type RemoveSSLCertificateResult struct {
}

type AddDrivesResult struct {
	//
	AsyncHandle int64 `json:"asyncHandle,omitempty"`
}

type DisableProtectionSchemesResult struct {
	//The protection schemes that are enabled on the cluster
	EnabledProtectionSchemes []ProtectionScheme `json:"enabledProtectionSchemes,"`
}

type VirtualVolumeUnsharedChunkResult struct {
	//Number of allocated/unshared chunks.
	Chunks int64 `json:"chunks,"`
	//Number of chunks scanned.
	ScannedChunks int64 `json:"scannedChunks,"`
	//Size of each chunk.
	ChunkSize int64 `json:"chunkSize,"`
}

type AddressBlock struct {
	//Start of the IP address range.
	Start string `json:"start,"`
	//Number of IP addresses to include in the block.
	Size int64 `json:"size,"`
	//Nuber of available blocks
	Available string `json:"available,"`
}

type SnapMirrorNode struct {
	//The ID of the destination ONTAP system.
	SnapMirrorEndpointID int64 `json:"snapMirrorEndpointID,"`
	//The name of the ONTAP node.
	Name string `json:"name,"`
	//The model of the ONTAP node.
	Model string `json:"model,"`
	//The serial number of the ONTAP node.
	SerialNumber string `json:"serialNumber,"`
	//The ONTAP product version.
	ProductVersion string `json:"productVersion,"`
	//The health of a node in the ONTAP cluster.
	//Possible values:
	//true
	//false
	IsNodeHealthy string `json:"isNodeHealthy,"`
	//Whether or not the node is eligible to participate in a ONTAP cluster.
	//Possible values:
	//true
	//false
	IsNodeEligible string `json:"isNodeEligible,"`
}

type DriveInfo struct {
	//Total Raw capacity of the drive, in bytes.
	Capacity int64 `json:"capacity,"`
	//Total Usable capacity of the drive, in bytes.
	UsableCapacity int64 `json:"usableCapacity,"`
	//Segment File Size of the drive, in bytes.
	SegmentFileSize int64 `json:"segmentFileSize,"`
	//DriveID for this drive.
	DriveID int64 `json:"driveID,"`
	//NodeID where this drive is located.
	NodeID int64 `json:"nodeID,"`
	//Drive serial number.
	Serial string `json:"serial,"`
	//For HCI platforms, this value is the node letter and slot number in the server chassis where this drive is located.
	//For legacy platforms, the slot number is a string representation of the 'slot' integer.
	ChassisSlot string `json:"chassisSlot,"`
	//Slot number in the server chassis where this drive is located, or -1 if SATADimm used for internal metadata drive.
	Slot int64 `json:"slot,"`
	//
	Status string `json:"status,"`
	//If a drive's status is 'Failed', this field provides more detail on why the drive was marked failed.
	DriveFailureDetail string `json:"driveFailureDetail,omitempty"`
	//If enabling or disabling drive security failed, this is the reason why it failed.
	//If the value is 'none', there was no failure.
	DriveSecurityFaultReason string `json:"driveSecurityFaultReason,omitempty"`
	//Identifies the provider of the authentication key for unlocking this drive.
	KeyProviderID int64 `json:"keyProviderID,omitempty"`
	//The keyID used by the key provider to acquire the authentication key for unlocking this drive.
	KeyID string `json:"keyID,omitempty"`
	//
	Type string `json:"type,"`
	//List of Name/Value pairs in JSON object format.
	Attributes interface{} `json:"attributes,"`
}

type FibreChannelSession struct {
	//The WWPN of the initiator which is logged into the target port.
	InitiatorWWPN string `json:"initiatorWWPN,"`
	//The node owning the Fibre Channel session.
	NodeID int64 `json:"nodeID,"`
	//The service ID of the FService owning this Fibre Channel session
	ServiceID int64 `json:"serviceID,"`
	//The WWPN of the target port involved in this session.
	TargetWWPN string `json:"targetWWPN,"`
	//The ID of the volume access group to which the initiatorWWPN beintegers. If not in a volume access group, the value will be null.
	VolumeAccessGroupID int64 `json:"volumeAccessGroupID,omitempty"`
}

type GetVirtualVolumeAllocatedBitmapRequest struct {
	//The ID of the Virtual Volume.
	VirtualVolumeID string `json:"virtualVolumeID,"`
	//Byte offset.
	SegmentStart int64 `json:"segmentStart,"`
	//Byte length adjusted to end on a chunk boundary.
	SegmentLength int64 `json:"segmentLength,"`
	//Number of bytes represented by one bit in the bitmap.
	ChunkSize int64 `json:"chunkSize,"`
	//
	CallingVirtualVolumeHostID string `json:"callingVirtualVolumeHostID,omitempty"`
}

type ModifyVolumeAccessGroupLunAssignmentsResult struct {
	//
	VolumeAccessGroupLunAssignments VolumeAccessGroupLunAssignments `json:"volumeAccessGroupLunAssignments,"`
}

type NeedsWorkCreateSnapMirrorEndpointUnmanagedResult struct {
}

type NeedsWorkListCloneJobsResult struct {
}

type CheckProposedResult struct {
	//True if there were no errors found with the proposed cluster, false otherwise
	ProposedClusterValid bool `json:"proposedClusterValid,"`
	//The errors associated with the proposed cluster.
	ProposedClusterErrors []ProposedClusterError `json:"proposedClusterErrors,"`
}

type DeleteInitiatorsResult struct {
}

type VirtualVolumeTaskResult struct {
	//Returns the state of a VVol Async Task.
	Task VirtualVolumeTask `json:"task,"`
}

type NeedsWorkPurgeAllDeletedVolumesResult struct {
}

type NetworkInterface struct {
	//IP address of the network.
	Address string `json:"address,"`
	//Address of NTP broadcast.
	Broadcast string `json:"broadcast,"`
	//Address used to configure the interface.
	MacAddress string `json:"macAddress,"`
	//Packet size on the network interface.
	Mtu int64 `json:"mtu,"`
	//Name of the network interface.
	Name string `json:"name,"`
	//Netmask for the network interface.
	Netmask string `json:"netmask,"`
	//Status of the network.
	Status string `json:"status,"`
	//The type of network, ie, BondMaster.
	Type string `json:"type,"`
	//Virtual Network Tag if on virtual network.
	VirtualNetworkTag int64 `json:"virtualNetworkTag,"`
	//
	Namespace bool `json:"namespace,omitempty"`
}

type NodeWaitingToJoin struct {
	//
	Name string `json:"name,omitempty"`
	//
	Version string `json:"version,"`
	//
	NodeID int64 `json:"nodeID,omitempty"`
	//
	PendingNodeID int64 `json:"pendingNodeID,omitempty"`
	//
	Mip string `json:"mip,omitempty"`
	//
	Cip string `json:"cip,omitempty"`
	//
	Sip string `json:"sip,omitempty"`
	//
	Compatible bool `json:"compatible,"`
	//
	ChassisType string `json:"chassisType,omitempty"`
	//
	Hostname string `json:"hostname,omitempty"`
	//
	NodeType string `json:"nodeType,omitempty"`
}

type SetProtectionDomainLayoutResult struct {
	//How all of the nodes are grouped into different ProtectionDomains.
	ProtectionDomainLayout ProtectionDomainLayout `json:"protectionDomainLayout,"`
}

type CreateInitiatorsResult struct {
	//List of objects containing details about the newly created initiators.
	Initiators []Initiator `json:"initiators,"`
}

type DeleteIdpConfigurationResult struct {
}

type GetFipsReportResult struct {
	//Array of nodes containing FIPS information.
	Nodes []FipsNodeReportType `json:"nodes,"`
	//Array of nodes that failed to gather FIPS information.
	ErrorNodes []FipsErrorNodeReportType `json:"errorNodes,"`
}

type KernelCrashDump struct {
	//
	KernelCrashDumpMinFreeGb int64 `json:"kernelCrashDumpMinFreeGb,"`
	//
	KernelCrashDumpDirectory string `json:"kernelCrashDumpDirectory,"`
	//
	KernelCrashDumpKernelOptions string `json:"kernelCrashDumpKernelOptions,"`
	//
	KernelCrashDumpMakedumpfileLevel int64 `json:"kernelCrashDumpMakedumpfileLevel,"`
	//
	KernelCrashDumpDefaultState string `json:"kernelCrashDumpDefaultState,"`
}

type ListAuthSessionsResult struct {
	//List of AuthSessionInfos.
	Sessions []AuthSessionInfo `json:"sessions,"`
}

type ListUtilitiesResult struct {
	//List of utilities currently available to run on the node.
	Utilities string `json:"utilities,"`
}

type CopyVolumeResult struct {
	//
	CloneID int64 `json:"cloneID,"`
	//Handle value used to track the progress of the volume copy.
	AsyncHandle int64 `json:"asyncHandle,"`
}

type CreateGroupSnapshotResult struct {
	//
	GroupSnapshot GroupSnapshot `json:"groupSnapshot,"`
	//Unique ID of the new group snapshot.
	GroupSnapshotID int64 `json:"groupSnapshotID,"`
	//List of checksum, volumeIDs and snapshotIDs for each member of the group.
	Members []GroupSnapshotMembers `json:"members,"`
}

type Service struct {
	//Unique identifier for this service.
	ServiceID int64 `json:"serviceID,"`
	//
	ServiceType string `json:"serviceType,"`
	//The node this service resides on.
	NodeID int64 `json:"nodeID,"`
	//This service's associated bulk volume service.
	//This will only be set if the service type is a slice service.
	AssociatedBV int64 `json:"associatedBV,omitempty"`
	//This service's associated transport service.
	//This will only be set if the service type is a slice service.
	AssociatedTS int64 `json:"associatedTS,omitempty"`
	//This service's associated volume service.
	//This will only be set if the service type is a slice service.
	AssociatedVS int64 `json:"associatedVS,omitempty"`
	//The list of asynchronous jobs currently running for this service.
	AsyncResultIDs []int64 `json:"asyncResultIDs,"`
	//If this service resides on a drive, the ID of that drive.
	DriveID int64 `json:"driveID,omitempty"`
	//Has this service started successfully?
	//When a new drive is added to the system, the created service will initially have a value of false here.
	//After the service has started, this value will be set to true.
	//This can be used to check if the service has ever started.
	FirstTimeStartup bool `json:"firstTimeStartup,"`
	//The port used for intra-cluster communication.
	//This will be in the 4000-4100 range.
	IpcPort int64 `json:"ipcPort,"`
	//The port used for iSCSI traffic.
	//This will only be set if the service type is a transport service.
	IscsiPort int64 `json:"iscsiPort,"`
	//
	Status string `json:"status,"`
	//
	StartedDriveIDs []int64 `json:"startedDriveIDs,"`
	//
	DriveIDs []int64 `json:"driveIDs,"`
	//
	SmartSsdWriteEnabled bool `json:"smartSsdWriteEnabled,omitempty"`
}

type DeleteQoSPolicyResult struct {
}

type GetIpmiInfoNodesResult struct {
	//
	NodeID int64 `json:"nodeID,"`
	//
	Result GetIpmiInfoNodesResultObject `json:"result,"`
}

type GetSnmpInfoResult struct {
	//List of networks and access types enabled for SNMP.
	//
	//Note: "networks" will only be present if SNMP V3 is disabled.
	Networks []SnmpNetwork `json:"networks,omitempty"`
	//If the nodes in the cluster are configured for SNMP.
	Enabled bool `json:"enabled,"`
	//If the nodes in the cluster are configured for SNMP v3.
	SnmpV3Enabled bool `json:"snmpV3Enabled,"`
	//If SNMP v3 is enabled, the values returned is a list of user access parameters for SNMP information from the cluster. This will be returned instead of the "networks" parameter.
	UsmUsers []SnmpV3UsmUser `json:"usmUsers,omitempty"`
}

type ListActivePairedVolumesResult struct {
	//Volume information for the paired volumes.
	Volumes []Volume `json:"volumes,"`
}

type SetLoginSessionInfoResult struct {
}

type CreateSnapshotResult struct {
	//
	Snapshot Snapshot `json:"snapshot,"`
	//ID of the newly-created snapshot.
	SnapshotID int64 `json:"snapshotID,"`
	//A string that represents the correct digits in the stored snapshot.
	//This checksum can be used later to compare other snapshots to detect errors in the data.
	Checksum string `json:"checksum,"`
}

type DeleteGroupSnapshotResult struct {
}

type GetFeatureStatusResult struct {
	//An array of feature objects indicating the feature name and its status.
	Features []FeatureObject `json:"features,"`
}

type NeedsWorkListRepositoriesResult struct {
}

type NeedsWorkStartGCResult struct {
}

type CHAPSecret struct {
}

type FibreChannelPortInfo struct {
	//The version of the firmware installed on the Fibre Channel port.
	Firmware string `json:"firmware,"`
	//The ID of the individual HBA port.
	HbaPort int64 `json:"hbaPort,"`
	//Model of the HBA on the port.
	Model string `json:"model,"`
	//Unique SolidFire port node ID.
	NPortID string `json:"nPortID,"`
	//Slot in which the pci card resides on the Fibre Channel node hardware.
	PciSlot int64 `json:"pciSlot,"`
	//Serial number on the Fibre Channel port.
	Serial string `json:"serial,"`
	//Speed of the HBA on the port.
	Speed string `json:"speed,"`
	//Possible values:
	//
	//<strong>UnknownNotPresentOnlineOfflineBlockedBypassedDiagnosticsLinkdownErrorLoopbackDeleted</strong>
	State string `json:"state,"`
	//The World Wide Name of the Fibre Channel switch port.
	SwitchWwn string `json:"switchWwn,"`
	//World Wide Node Name of the HBA node.
	Wwnn string `json:"wwnn,"`
	//World Wide Port Name assigned to the physical port of the HBA.
	Wwpn string `json:"wwpn,"`
}

type GetIdpAuthenticationStateResult struct {
	//Whether third party Identity Provider Authentication is enabled.
	Enabled bool `json:"enabled,"`
}

type NeedsWorkCreateClusterFaultResult struct {
}

type QueryVirtualVolumeMetadataRequest struct {
	//
	QueryConstraints interface{} `json:"queryConstraints,omitempty"`
	//
	WildcardConstraints []string `json:"wildcardConstraints,omitempty"`
	//
	CallingVirtualVolumeHostID string `json:"callingVirtualVolumeHostID,omitempty"`
}

type CreateScheduleResult struct {
	//
	ScheduleID int64 `json:"scheduleID,"`
	//
	Schedule Schedule `json:"schedule,omitempty"`
}

type EnableIdpAuthenticationResult struct {
}

type NeedsWorkModifySnapMirrorEndpointUnmanagedResult struct {
}

type NeedsWorkNotifyIntentToRestartResult struct {
}

type ListVirtualVolumeTasksResult struct {
	//List of VVol Async Tasks.
	Tasks []VirtualVolumeTask `json:"tasks,"`
}

type ModifyClusterFullThresholdResult struct {
	//Current computed level of block fullness of the cluster.
	//Possible values: stage1Happy: No alerts or error conditions. stage2Aware: 3 nodes of capacity available. stage3Low: 2 nodes of capacity available. stage4Critical: 1 node of capacity available. No new volumes or clones can be created. stage5CompletelyConsumed: Completely consumed. Cluster is read-only, iSCSI connection is maintained but all writes are suspended.
	BlockFullness string `json:"blockFullness,"`
	//Reflects the highest level of fullness between "blockFullness" and "metadataFullness".
	Fullness string `json:"fullness,"`
	//A value representative of the number of times metadata space can be over provisioned relative to the amount of space available. For example, if there was enough metadata space to store 100 TiB of volumes and this number was set to 5, then 500 TiB worth of volumes could be created.
	MaxMetadataOverProvisionFactor int64 `json:"maxMetadataOverProvisionFactor,"`
	//Current computed level of metadata fullness of the cluster.
	MetadataFullness string `json:"metadataFullness,"`
	//Error condition; message sent to "Alerts" if the reserved slice utilization is greater than the sliceReserveUsedThresholdPct value returned.
	SliceReserveUsedThresholdPct int64 `json:"sliceReserveUsedThresholdPct,"`
	//Awareness condition: Value that is set for "Stage 2" cluster threshold level.
	Stage2AwareThreshold int64 `json:"stage2AwareThreshold,"`
	//Number of bytes being used by the cluster at which a stage2 condition will exist.
	Stage2BlockThresholdBytes int64 `json:"stage2BlockThresholdBytes,"`
	//Number of bytes being used by the cluster at which a stage3 condition will exist.
	Stage3BlockThresholdBytes int64 `json:"stage3BlockThresholdBytes,"`
	//The percent value set for stage3 of block fullness. At this percent full, a warning will be posted in the Alerts log.
	Stage3BlockThresholdPercent int64 `json:"stage3BlockThresholdPercent,"`
	//The percent value set for stage3 of metadata fullness. At this percent full, a warning will be posted in the Alerts log.
	Stage3MetadataThresholdPercent int64 `json:"stage3MetadataThresholdPercent,"`
	//Error condition; message sent to "Alerts" that capacity on a cluster is getting low.
	Stage3LowThreshold int64 `json:"stage3LowThreshold,"`
	//Error condition; message sent to "Alerts" that capacity on a cluster is critically low.
	Stage4CriticalThreshold int64 `json:"stage4CriticalThreshold,"`
	//Number of bytes being used by the cluster at which a stage4 condition will exist.
	Stage4BlockThresholdBytes int64 `json:"stage4BlockThresholdBytes,"`
	//Number of bytes being used by the cluster at which a stage5 condition will exist.
	Stage5BlockThresholdBytes int64 `json:"stage5BlockThresholdBytes,"`
	//Physical capacity of the cluster measured in bytes.
	SumTotalClusterBytes int64 `json:"sumTotalClusterBytes,"`
	//Total amount of space that can be used to store metadata.
	SumTotalMetadataClusterBytes int64 `json:"sumTotalMetadataClusterBytes,"`
	//Number of bytes used on the cluster.
	SumUsedClusterBytes int64 `json:"sumUsedClusterBytes,"`
	//Amount of space used on volume drives to store metadata.
	SumUsedMetadataClusterBytes int64 `json:"sumUsedMetadataClusterBytes,"`
	//Number of metadata bytes being used by the cluster at which a stage2 condition will exist.
	Stage2MetadataThresholdBytes int64 `json:"stage2MetadataThresholdBytes,"`
	//Number of metadata bytes being used by the cluster at which a stage3 condition will exist.
	Stage3MetadataThresholdBytes int64 `json:"stage3MetadataThresholdBytes,"`
	//Number of metadata bytes being used by the cluster at which a stage4 condition will exist.
	Stage4MetadataThresholdBytes int64 `json:"stage4MetadataThresholdBytes,"`
	//Number of metadata bytes being used by the cluster at which a stage5 condition will exist.
	Stage5MetadataThresholdBytes int64 `json:"stage5MetadataThresholdBytes,"`
}

type UnbindArguments struct {
	//
	VirtualVolumeID string `json:"virtualVolumeID,"`
	//
	ProtocolEndpointType string `json:"protocolEndpointType,"`
	//
	ProtocolEndpointInBandID string `json:"protocolEndpointInBandID,"`
	//
	VirtualVolumeSecondaryID string `json:"virtualVolumeSecondaryID,"`
}

type CreateInitiator struct {
	//The name of the initiator (IQN or WWPN) to create.
	Name string `json:"name,"`
	//The friendly name to assign to this initiator.
	Alias string `json:"alias,omitempty"`
	//The ID of the volume access group to which this newly created initiator will be added.
	VolumeAccessGroupID int64 `json:"volumeAccessGroupID,omitempty"`
	//A set of JSON attributes assigned to this initiator. (JSON Object)
	Attributes interface{} `json:"attributes,omitempty"`
	//"requireChap" determines if the initiator is required to use CHAP during session login. CHAP is optional if "requireChap" is false.
	RequireChap bool `json:"requireChap,omitempty"`
	//The CHAP username for this initiator. Defaults to the initiator name (IQN) if not specified during creation and "requireChap" is true.
	ChapUsername string `json:"chapUsername,omitempty"`
	//The CHAP secret used for authentication of the initiator. Defaults to a randomly generated secret if not specified during creation and "requireChap" is true.
	InitiatorSecret string `json:"initiatorSecret,omitempty"`
	//The CHAP secret used for authentication of the target. Defaults to a randomly generated secret if not specified during creation and "requireChap" is true.
	TargetSecret string `json:"targetSecret,omitempty"`
}

type ProtocolEndpointResult struct {
	//
	ProtocolEndpoint ProtocolEndpoint `json:"protocolEndpoint,"`
}

type ListInitiatorsResult struct {
	//List of the initiator information.
	Initiators []Initiator `json:"initiators,"`
}

type ModifyVolumesResult struct {
	//
	Volumes []Volume `json:"volumes,"`
	//
	Qos QoS `json:"qos,omitempty"`
}

type NeedsWorkCreateEventResult struct {
}

type NeedsWorkGetSliceFileSizeReportResult struct {
}

type NeedsWorkRegenerateClusterUuidResult struct {
}

type PurgeDeletedVolumeResult struct {
}

type GroupSnapshot struct {
	//Unique ID of the new group snapshot.
	GroupSnapshotID int64 `json:"groupSnapshotID,"`
	//UUID of the group snapshot.
	GroupSnapshotUUID string `json:"groupSnapshotUUID,"`
	//List of snapshots that are members of the group.
	Members []Snapshot `json:"members,"`
	//Name of the group snapshot, or, if none was given, the UTC formatted day and time on which the snapshot was created.
	Name string `json:"name,"`
	//The UTC formatted day and time on which the snapshot was created.
	CreateTime string `json:"createTime,"`
	//Status of the snapshot.
	//Possible values:
	//Preparing: A snapshot that is being prepared for use and is not yet writable.
	//Done: A snapshot that has finished being prepared and is now usable
	Status string `json:"status,"`
	//Identifies if group snapshot is enabled for remote replication.
	EnableRemoteReplication bool `json:"enableRemoteReplication,"`
	//Replication status of the group snapshot as seen on the source cluster.
	//Shows if the group snapshot replication is currently in progress, or
	//has successfully completed.
	RemoteStatuses []GroupSnapshotRemoteStatus `json:"remoteStatuses,omitempty"`
	//List of Name/Value pairs in JSON object format.
	Attributes interface{} `json:"attributes,"`
}

type IpmiInfo struct {
	//
	Sensors []interface{} `json:"sensors,"`
}

type SolidfireDefaults struct {
	//
	SliceFileLogFileCapacity int64 `json:"sliceFileLogFileCapacity,"`
	//
	PostCallbackThreadCount int64 `json:"postCallbackThreadCount,"`
	//
	CpuDmaLatency int64 `json:"cpuDmaLatency,"`
	//
	BufferCacheGB int64 `json:"bufferCacheGB,"`
	//
	MaxIncomingSliceSyncs int64 `json:"maxIncomingSliceSyncs,"`
	//
	ConfiguredIops int64 `json:"configuredIops,"`
	//
	SCacheFileCapacity int64 `json:"sCacheFileCapacity,"`
	//
	MaxDriveWriteThroughputMBPerSec int64 `json:"maxDriveWriteThroughputMBPerSec,"`
	//
	DriveWriteThroughputMBPerSleep int64 `json:"driveWriteThroughputMBPerSleep,"`
}

type StartClusterPairingResult struct {
	//A string of characters that is used by the "CompleteClusterPairing" API method.
	ClusterPairingKey string `json:"clusterPairingKey,"`
	//Unique identifier for the cluster pair.
	ClusterPairID int64 `json:"clusterPairID,"`
}

type SetSnmpTrapInfoResult struct {
}

type Signature struct {
	//
	Data string `json:"data,"`
	//
	Pubkey string `json:"pubkey,"`
	//
	Version int64 `json:"version,"`
}

type GetImmutableValuesResult struct {
	//The number of additional bytes that must be persisted as metadata for each block stored on the block services.
	PerBlockOverheadBytes int64 `json:"perBlockOverheadBytes,"`
}

type NeedsWorkGetSliceReserveUsedThresholdPctResult struct {
}

type SnapMirrorPolicyRule struct {
	//The ID of the destination ONTAP system.
	SnapMirrorEndpointID int64 `json:"snapMirrorEndpointID,"`
	//The snapshot copy label, used for snapshot copy
	//selection in extended data protection relationships.
	SnapMirrorLabel string `json:"snapMirrorLabel,"`
	//Specifies the maximum number of snapshot copies that are
	//retained on the SnapMirror destination volume for a rule.
	KeepCount string `json:"keepCount,"`
}

type CheckPingOnVlanResult struct {
	//Result of the ping on vlan test.
	Result string `json:"result,"`
	//Details of the test for each destination host
	//Name/Value pairs: Name corresponding to ip address of the host
	Details interface{} `json:"details,"`
	//The total duration of the test.
	Duration string `json:"duration,"`
}

type EnableProtectionSchemesResult struct {
	//The protection schemes that are enabled on the cluster
	EnabledProtectionSchemes []ProtectionScheme `json:"enabledProtectionSchemes,"`
}

type SetAuthConfigurationResult struct {
}

type ListBackupTargetsResult struct {
	//Objects returned for each backup target.
	BackupTargets []BackupTarget `json:"backupTargets,"`
}

type ModifyVasaProviderInfoRequest struct {
	//Signed SSL certificate for the Vasa Provider
	Keystore string `json:"keystore,omitempty"`
	//UUID identifying the vasa provider
	VasaProviderID string `json:"vasaProviderID,omitempty"`
	//
	Options interface{} `json:"options,omitempty"`
}

type DisableEncryptionAtRestResult struct {
}

type DriveHardware struct {
	//
	CanonicalName string `json:"canonicalName,"`
	//
	Connected bool `json:"connected,"`
	//
	Dev int64 `json:"dev,"`
	//
	DevPath string `json:"devPath,"`
	//
	DriveType string `json:"driveType,"`
	//
	DriveEncryptionCapability DriveEncryptionCapabilityType `json:"driveEncryptionCapability,"`
	//
	LifeRemainingPercent int64 `json:"lifeRemainingPercent,"`
	//
	LifetimeReadBytes int64 `json:"lifetimeReadBytes,"`
	//
	LifetimeWriteBytes int64 `json:"lifetimeWriteBytes,"`
	//
	Name string `json:"name,"`
	//
	Path string `json:"path,"`
	//
	PathLink string `json:"pathLink,"`
	//
	PowerOnHours int64 `json:"powerOnHours,"`
	//
	Product string `json:"product,"`
	//
	ReallocatedSectors int64 `json:"reallocatedSectors,"`
	//
	ReserveCapacityPercent int64 `json:"reserveCapacityPercent,"`
	//
	ScsiCompatId string `json:"scsiCompatId,"`
	//
	ScsiState string `json:"scsiState,"`
	//
	SecurityAtMaximum bool `json:"securityAtMaximum,"`
	//
	SecurityEnabled bool `json:"securityEnabled,"`
	//
	SecurityFrozen bool `json:"securityFrozen,"`
	//
	SecurityLocked bool `json:"securityLocked,"`
	//
	SecuritySupported bool `json:"securitySupported,"`
	//
	Serial string `json:"serial,"`
	//
	Size int64 `json:"size,"`
	//
	Slot int64 `json:"slot,"`
	//
	SmartSsdWriteCapable bool `json:"smartSsdWriteCapable,omitempty"`
	//
	Uuid string `json:"uuid,"`
	//
	Vendor string `json:"vendor,"`
	//
	Version string `json:"version,"`
}

type GetNodeStatsResult struct {
	//Node activity information.
	NodeStats NodeStatsInfo `json:"nodeStats,"`
}

type ListVolumeAccessGroupsResult struct {
	//A list of objects describing each volume access group.
	VolumeAccessGroups []VolumeAccessGroup `json:"volumeAccessGroups,"`
	//A list of volume access groups not found by the system. Present if you used the "volumeAccessGroups" parameter and the system was unable to find one or more volume access groups that you specified.
	VolumeAccessGroupsNotFound []int64 `json:"volumeAccessGroupsNotFound,omitempty"`
}

type NeedsWorkGetServiceStatusResult struct {
}

type NeedsWorkModifyClusterFaultResult struct {
}

type CloneVirtualVolumeRequest struct {
	//The ID of the Virtual Volume to clone.
	VirtualVolumeID string `json:"virtualVolumeID,"`
	//The name for the newly-created volume.
	Name string `json:"name,omitempty"`
	//New quality of service settings for this volume.
	Qos QoS `json:"qos,omitempty"`
	//
	Metadata interface{} `json:"metadata,omitempty"`
	//
	NewContainerID string `json:"newContainerID,omitempty"`
	//
	CallingVirtualVolumeHostID string `json:"callingVirtualVolumeHostID,omitempty"`
}

type ClusterConfig struct {
	//Network interface used for cluster communication.
	Cipi string `json:"cipi,omitempty"`
	//Unique cluster name.
	Cluster string `json:"cluster,omitempty"`
	//Nodes that are participating in the cluster.
	Ensemble []string `json:"ensemble,omitempty"`
	//Network interface used for node management.
	Mipi string `json:"mipi,omitempty"`
	//Unique cluster name.
	Name string `json:"name,omitempty"`
	//
	NodeID int64 `json:"nodeID,omitempty"`
	//
	PendingNodeID int64 `json:"pendingNodeID,omitempty"`
	//Identifies the role of the node
	Role string `json:"role,omitempty"`
	//Network interface used for storage.
	Sipi string `json:"sipi,omitempty"`
	//
	State string `json:"state,omitempty"`
	//This field indicates whether the node supports encryption.
	EncryptionCapable bool `json:"encryptionCapable,omitempty"`
	//This field indicates whether the node supports FIPS 140-2 certified drives.
	FipsDriveConfiguration bool `json:"fipsDriveConfiguration,omitempty"`
	//
	HasLocalAdmin bool `json:"hasLocalAdmin,omitempty"`
	//
	Version string `json:"version,omitempty"`
}

type SnapMirrorJobScheduleCronInfo struct {
	//The ID of the destination ONTAP system.
	SnapMirrorEndpointID int64 `json:"snapMirrorEndpointID,"`
	//The name of the job schedule.
	JobScheduleName string `json:"jobScheduleName,"`
	//An automatically-generated human-readable summary of the schedule.
	JobScheduleDescription string `json:"jobScheduleDescription,"`
}

type VasaProviderInfo struct {
	//
	Keystore string `json:"keystore,"`
	//
	ControlPort int64 `json:"controlPort,"`
	//
	VasaProviderID string `json:"vasaProviderID,"`
	//
	Options interface{} `json:"options,"`
}

type ProtectionDomainLevel struct {
	//The type of the Protection Domain which has the associated Tolerance and Resiliency.
	ProtectionDomainType ProtectionDomainType `json:"protectionDomainType,"`
	//The current Tolerance of this cluster from the perspective of this Protection Domain Type.
	Tolerance ProtectionDomainTolerance `json:"tolerance,"`
	//The current Resiliency of this cluster from the perspective of this Protection Domain Type.
	Resiliency ProtectionDomainResiliency `json:"resiliency,"`
}

type SetNodeSSLCertificateResult struct {
}

type GetLldpInfoResult struct {
	//Contains information about the local chassis
	LldpChassis string `json:"lldpChassis,"`
	//Contains information about the local interfaces
	LldpInterfaces string `json:"lldpInterfaces,"`
	//Contains information about known LLDP neighbors
	LldpNeighbors string `json:"lldpNeighbors,"`
}

type SyncJob struct {
	//
	BytesPerSecond float64 `json:"bytesPerSecond,"`
	//
	CurrentBytes int64 `json:"currentBytes,"`
	//
	DstServiceID int64 `json:"dstServiceID,"`
	//
	ElapsedTime float64 `json:"elapsedTime,"`
	//
	PercentComplete float64 `json:"percentComplete,"`
	//
	RemainingTime int64 `json:"remainingTime,omitempty"`
	//
	SliceID int64 `json:"sliceID,"`
	//
	SrcServiceID int64 `json:"srcServiceID,"`
	//
	TotalBytes int64 `json:"totalBytes,"`
	//
	Type string `json:"type,"`
	//
	CloneID int64 `json:"cloneID,"`
	//
	DstVolumeID int64 `json:"dstVolumeID,"`
	//
	NodeID int64 `json:"nodeID,"`
	//
	SnapshotID int64 `json:"snapshotID,"`
	//
	SrcVolumeID int64 `json:"srcVolumeID,"`
	//
	BlocksPerSecond float64 `json:"blocksPerSecond,"`
	//
	Stage string `json:"stage,"`
	//
	GroupCloneID int64 `json:"groupCloneID,omitempty"`
}

type DeleteStandbySlicesResult struct {
}

type GetDriveConfigResult struct {
	//Configuration information for the drives that are connected to the cluster
	DriveConfig DrivesConfigInfo `json:"driveConfig,"`
}

type GetSnmpACLResult struct {
	//List of networks and what type of access they have to the SNMP servers running on the cluster nodes. Present if SNMP v3 is disabled.
	Networks []SnmpNetwork `json:"networks,omitempty"`
	//List of users and the type of access they have to the SNMP servers running on the cluster nodes. Present if SNMP v3 is enabled.
	UsmUsers []SnmpV3UsmUser `json:"usmUsers,omitempty"`
}

type ListIdpConfigurationsResult struct {
	//Information around the third party Identity Provider (IdP) configuration(s).
	IdpConfigInfos []IdpConfigInfo `json:"idpConfigInfos,"`
}

type ListPendingNodesResult struct {
	//
	PendingNodes []PendingNode `json:"pendingNodes,"`
}

type Network struct {
	//Name of the storage node network interface used for management traffic.
	Bond1G NetworkConfig `json:"Bond1G,omitempty"`
	//Name of the storage node network interface used for storage and cluster traffic.
	Bond10G NetworkConfig `json:"Bond10G,omitempty"`
	//Name of the witness node network interface used for management traffic.
	Net0 NetworkConfig `json:"net0,omitempty"`
	//Name of the witness node network interface used for storage and cluster traffic.
	Net1 NetworkConfig `json:"net1,omitempty"`
	//
	Eth0 NetworkConfig `json:"eth0,omitempty"`
	//
	Eth1 NetworkConfig `json:"eth1,omitempty"`
	//
	Eth2 NetworkConfig `json:"eth2,omitempty"`
	//
	Eth3 NetworkConfig `json:"eth3,omitempty"`
	//
	Eth4 NetworkConfig `json:"eth4,omitempty"`
	//
	Eth5 NetworkConfig `json:"eth5,omitempty"`
	//
	Lo NetworkConfig `json:"lo,omitempty"`
}

type SnapMirrorNetworkInterface struct {
	//The ID of the destination ONTAP system.
	SnapMirrorEndpointID int64 `json:"snapMirrorEndpointID,"`
	//The logical interface (LIF) name.
	InterfaceName string `json:"interfaceName,"`
	//The IP address of the LIF.
	NetworkAddress string `json:"networkAddress,"`
	//The network mask of the LIF.
	NetworkMask string `json:"networkMask,"`
	//The role of the LIF.
	//Possible values:
	//undef
	//cluster
	//data
	//node_mgmt
	//intercluster
	//cluster_mgmt
	InterfaceRole string `json:"interfaceRole,"`
	//Specifies the operational status of the LIF.
	//Possible values:
	//up
	//down
	OperationalStatus string `json:"operationalStatus,"`
	//Specifies the administrative status of the LIF. The administrative status can differ from the operational status.
	//For instance, if you specify the status as up but a network problem prevents the interface from functioning, the operational status remains as down.
	//Possible values:
	//up
	//down
	AdministrativeStatus string `json:"administrativeStatus,"`
	//The name of the Vserver.
	VserverName string `json:"vserverName,"`
}

type StartBulkVolumeWriteResult struct {
	//ID of the async process to be checked for completion.
	AsyncHandle int64 `json:"asyncHandle,"`
	//Opaque key uniquely identifying the session.
	Key string `json:"key,"`
	//URL to access the node's web server
	Url string `json:"url,"`
}

type CancelCloneResult struct {
}

type CloneMultipleVolumeParams struct {
	//Required parameter for "volumes" array: volumeID.
	VolumeID int64 `json:"volumeID,"`
	//Access settings for the new volume.
	//readOnly: Only read operations are allowed.
	//readWrite: Reads and writes are allowed.
	//locked: No reads or writes are allowed.
	//replicationTarget: Identify a volume as the target volume for a paired set of volumes. If the volume is not paired, the access status is locked.
	//
	//If unspecified, the access settings of the clone will be the same as the source.
	Access string `json:"access,omitempty"`
	//New name for the clone.
	Name string `json:"name,omitempty"`
	//Account ID for the new volume.
	NewAccountID int64 `json:"newAccountID,omitempty"`
	//New size Total size of the volume, in bytes. Size is rounded up to the nearest 1MB size.
	NewSize int64 `json:"newSize,omitempty"`
	//List of Name/Value pairs in JSON object format.
	Attributes interface{} `json:"attributes,omitempty"`
}

type ListSnapMirrorSchedulesResult struct {
	//A list of the SnapMirror schedules on the remote ONTAP cluster.
	SnapMirrorSchedules []SnapMirrorJobScheduleCronInfo `json:"snapMirrorSchedules,"`
}

type ListTestsResult struct {
	//List of tests that can be performed on the node.
	Tests string `json:"tests,"`
}

type ModifyGroupSnapshotResult struct {
	//
	GroupSnapshot GroupSnapshot `json:"groupSnapshot,"`
}

type NeedsWorkFinishUpgradeResult struct {
}

type NeedsWorkGetLocalStatsResult struct {
}

type GetSSLCertificateResult struct {
	//The full PEM-encoded test of the certificate.
	Certificate string `json:"certificate,"`
	//The decoded information of the certificate.
	Details interface{} `json:"details,"`
}

type ListSnapMirrorAggregatesResult struct {
	//A list of the aggregates available on the ONTAP storage system.
	SnapMirrorAggregates []SnapMirrorAggregate `json:"snapMirrorAggregates,"`
}

type ResizeDrivesResult struct {
	//Information for the drives that have been resized.
	Drives []DriveInfo `json:"drives,"`
}

type UnbindAllVirtualVolumesFromHostResult struct {
}

type AddVirtualNetworkResult struct {
	//The virtual network ID of the new virtual network.
	VirtualNetworkID int64 `json:"virtualNetworkID,omitempty"`
}

type CheckProposedNodeAdditionsResult struct {
	//This field indicates the validity of proposed cluster.
	ProposedClusterValid bool `json:"proposedClusterValid,"`
	//This field is a list of errors if proposed cluster is invalid.
	ProposedClusterErrors []ProposedNodeErrors `json:"proposedClusterErrors,"`
}

type ListVolumeStatsByVolumeResult struct {
	//List of account activity information.
	VolumeStats []VolumeStats `json:"volumeStats,"`
}

type NeedsWorkSetDaemonStatusResult struct {
}

type TestConnectEnsembleDetails struct {
	//A list of each ensemble node in the test and the results of the tests.
	Nodes interface{} `json:"nodes,"`
}

type ComputeVolumeChecksumResult struct {
	//Map of checksums for the primary and each live secondary slice service.
	//Keys are service IDs and values are their checksums. Both are strings.
	Checksums interface{} `json:"checksums,"`
}

type GetSnapMirrorClusterIdentityResult struct {
	//A list of cluster identities of SnapMirror endpoints.
	SnapMirrorClusterIdentity []SnapMirrorClusterIdentity `json:"snapMirrorClusterIdentity,"`
}

type LdapConfiguration struct {
	//Identifies which user authentcation method will be used.
	//Valid values:
	//DirectBind
	//SearchAndBind
	AuthType string `json:"authType,"`
	//Identifies whether or not the system is enabled for LDAP.
	//Valid values:
	//true
	//false
	Enabled bool `json:"enabled,"`
	//The base DN of the tree to start the group search (will do a subtree search from here).
	GroupSearchBaseDN string `json:"groupSearchBaseDN,"`
	//The custom search filter used.
	GroupSearchCustomFilter string `json:"groupSearchCustomFilter,"`
	//Controls the default group search filter used, can be one of the following:
	//NoGroups: No group support.
	//ActiveDirectory: Nested membership of all of a user's AD groups.
	//MemberDN: MemberDN style groups (single-level).
	GroupSearchType string `json:"groupSearchType,"`
	//A fully qualified DN to log in with to perform an LDAP search for the user (needs read access to the LDAP directory).
	SearchBindDN string `json:"searchBindDN,"`
	//A comma-separated list of LDAP server URIs (examples: "ldap://1.2.3.4" and ldaps://1.2.3.4:123")
	ServerURIs []string `json:"serverURIs,"`
	//A string that is used to form a fully qualified user DN.
	UserDNTemplate string `json:"userDNTemplate,"`
	//The base DN of the tree used to start the search (will do a subtree search from here).
	UserSearchBaseDN string `json:"userSearchBaseDN,"`
	//The LDAP filter used.
	UserSearchFilter string `json:"userSearchFilter,"`
}

type LoginBanner struct {
	//The current text of the Terms of Use banner.
	//This value can contain text even when the banner is disabled.
	Banner string `json:"banner,"`
	//The status of the Terms of Use banner. Possible values:
	//true: The Terms of Use banner is displayed upon web interface login.
	//false: The Terms of Use banner is not displayed upon web interface login.
	Enabled bool `json:"enabled,"`
}

type MetadataHosts struct {
	//Secondary metadata (slice) services that are in a dead state.
	DeadSecondaries []int64 `json:"deadSecondaries,"`
	//Secondary metadata (slice) services that are currently in a "live" state.
	LiveSecondaries []int64 `json:"liveSecondaries,"`
	//The primary metadata (slice) services hosting the volume.
	Primary int64 `json:"primary,"`
}

type ModifyScheduleResult struct {
	//
	Schedule Schedule `json:"schedule,omitempty"`
}

type NeedsWorkGetClusterSshInfoResult struct {
}

type SnapMirrorVserver struct {
	//The ID of the destination ONTAP system.
	SnapMirrorEndpointID int64 `json:"snapMirrorEndpointID,"`
	//The name of the Vserver.
	VserverName string `json:"vserverName,"`
	//The type of the Vserver.
	//Possible values:
	//data
	//admin
	//system
	//node
	VserverType string `json:"vserverType,"`
	//The subtype of the Vserver.
	//Possible values:
	//default
	//dp_destination
	//data
	//sync_source
	//sync_destination
	VserverSubtype string `json:"vserverSubtype,"`
	//The root volume of the Vserver.
	RootVolume string `json:"rootVolume,"`
	//The aggregate on which the root volume will be created.
	RootVolumeAggregate string `json:"rootVolumeAggregate,"`
	//An array of SnapMirrorVserverAggregateInfo objects.
	VserverAggregateInfo []SnapMirrorVserverAggregateInfo `json:"vserverAggregateInfo,"`
	//The detailed administrative state of the Vserver.
	//Possible values:
	//running
	//stopped
	//starting
	//stopping
	//initialized
	//deleting
	AdminState string `json:"adminState,"`
	//The basic operational state of the Vserver.
	//Possible values:
	//running
	//stopped
	OperationalState string `json:"operationalState,"`
}

type CreateAuthSessionResult struct {
	//The newly created AuthSessionInfo.
	Session AuthSessionInfo `json:"session,"`
}

type DeleteAllSupportBundlesResult struct {
	//
	Duration string `json:"duration,"`
	//
	Details interface{} `json:"details,"`
	//
	Result string `json:"result,"`
}

type TestKeyServerKmipResult struct {
}

type EnableSnmpResult struct {
}

type RemoteReplication struct {
	//Volume replication mode.
	//Possible values:
	//Async: Writes are acknowledged when they complete locally. The cluster does not wait for writes to be replicated to the target cluster.
	//Sync: Source acknowledges write when the data is stored locally and on the remote cluster.
	//SnapshotsOnly: Only snapshots created on the source cluster will be replicated. Active writes from the source volume will not be replicated.
	Mode string `json:"mode,"`
	//The number of occurring write ops before auto-pausing, on a per volume pair level.
	PauseLimit int64 `json:"pauseLimit,"`
	//The remote slice service ID.
	RemoteServiceID int64 `json:"remoteServiceID,"`
	//Reserved for future use.
	ResumeDetails string `json:"resumeDetails,"`
	//The details of snapshot replication.
	SnapshotReplication SnapshotReplication `json:"snapshotReplication,"`
	//The state of the volume replication.
	State string `json:"state,"`
	//Reserved for future use.
	StateDetails string `json:"stateDetails,"`
}

type GetClientCertificateSignRequestResult struct {
	//A PEM format Base64 encoded PKCS#10 X.509 client certificate sign request.
	ClientCertificateSignRequest string `json:"clientCertificateSignRequest,"`
}

type GetOriginNodeResult struct {
	//
	Origin Origin `json:"origin,omitempty"`
}

type QoSPolicy struct {
	//A unique integer identifier for the QoSPolicy auto-assigned by the SolidFire cluster.
	QosPolicyID int64 `json:"qosPolicyID,"`
	//The name of the QoS policy. For example: gold, platinum, or silver.
	Name string `json:"name,"`
	//A list of volumes associated with this policy.
	VolumeIDs []int64 `json:"volumeIDs,"`
	//Quality of service settings for this volume.
	Qos VolumeQOS `json:"qos,"`
}

type Config struct {
	//
	Cluster ClusterConfig `json:"cluster,"`
	//
	Network Network `json:"network,"`
}

type CreateVirtualVolumeHostRequest struct {
	//The GUID of the ESX host.
	VirtualVolumeHostID string `json:"virtualVolumeHostID,"`
	//The GUID of the ESX Cluster.
	ClusterID string `json:"clusterID,"`
	//
	InitiatorNames []string `json:"initiatorNames,omitempty"`
	//A list of PEs the host is aware of.
	VisibleProtocolEndpointIDs []string `json:"visibleProtocolEndpointIDs,omitempty"`
	//IP or DNS name for the host.
	HostAddress string `json:"hostAddress,omitempty"`
	//
	CallingVirtualVolumeHostID string `json:"callingVirtualVolumeHostID,omitempty"`
}

type ProtectionSchemeTolerance struct {
	//The Protection Scheme.
	ProtectionScheme ProtectionScheme `json:"protectionScheme,"`
	//The number of simultaneous failures which can occur without losing block data
	//availability for the Protection Scheme.
	SustainableFailuresForBlockData int64 `json:"sustainableFailuresForBlockData,"`
	//The number of simultaneous failures which can occur without losing metadata
	//or Vvol availability for the Protection Scheme.
	SustainableFailuresForMetadata int64 `json:"sustainableFailuresForMetadata,"`
}

type SetSupplementalTlsCiphersResult struct {
	//List of mandatory TLS cipher suites for the cluster.
	MandatoryCiphers []string `json:"mandatoryCiphers,"`
	//List of supplemental TLS cipher suites for the cluster.
	SupplementalCiphers []string `json:"supplementalCiphers,"`
}

type VolumePair struct {
	//The remote cluster a volume is paired with.
	ClusterPairID int64 `json:"clusterPairID,"`
	//The VolumeID on the remote cluster a volume is paired with.
	RemoteVolumeID int64 `json:"remoteVolumeID,"`
	//The SliceID on the remote cluster a volume is paired with.
	RemoteSliceID int64 `json:"remoteSliceID,"`
	//The last-observed name of the volume on the remote cluster a volume is paired with.
	RemoteVolumeName string `json:"remoteVolumeName,"`
	//A UUID in canonical form.
	VolumePairUUID string `json:"volumePairUUID,"`
	//Details about the replication configuration for this volume pair.
	RemoteReplication RemoteReplication `json:"remoteReplication,"`
}

type BreakSnapMirrorVolumeResult struct {
}

type ISCSISession struct {
	//
	DriveIDs []int64 `json:"driveIDs,omitempty"`
	//
	AccountID int64 `json:"accountID,"`
	//
	Initiator Initiator `json:"initiator,omitempty"`
	//
	AccountName string `json:"accountName,"`
	//
	DriveID int64 `json:"driveID,"`
	//
	InitiatorIP string `json:"initiatorIP,"`
	//
	InitiatorPortName string `json:"initiatorPortName,"`
	//
	TargetPortName string `json:"targetPortName,"`
	//
	InitiatorName string `json:"initiatorName,"`
	//
	NodeID int64 `json:"nodeID,"`
	//
	ServiceID int64 `json:"serviceID,"`
	//
	SessionID int64 `json:"sessionID,"`
	//
	TargetName string `json:"targetName,"`
	//
	TargetIP string `json:"targetIP,"`
	//
	VirtualNetworkID int64 `json:"virtualNetworkID,"`
	//
	VolumeID int64 `json:"volumeID,"`
	//
	CreateTime string `json:"createTime,"`
	//
	VolumeInstance int64 `json:"volumeInstance,"`
	//
	InitiatorSessionID int64 `json:"initiatorSessionID,"`
	//
	MsSinceLastScsiCommand int64 `json:"msSinceLastScsiCommand,omitempty"`
	//
	MsSinceLastIscsiPDU int64 `json:"msSinceLastIscsiPDU,omitempty"`
}

type ListDriveHardwareResult struct {
	//
	Nodes []NodeDriveHardware `json:"nodes,"`
}

type ListVolumesResult struct {
	//List of volumes.
	Volumes []Volume `json:"volumes,"`
}

type SetNtpInfoResult struct {
}

type VirtualVolumeInfo struct {
	//
	VirtualVolumeID string `json:"virtualVolumeID,"`
	//
	ParentVirtualVolumeID string `json:"parentVirtualVolumeID,"`
	//
	StorageContainer StorageContainer `json:"storageContainer,"`
	//
	VolumeID int64 `json:"volumeID,"`
	//
	SnapshotID int64 `json:"snapshotID,"`
	//
	VirtualVolumeType string `json:"virtualVolumeType,"`
	//
	Status string `json:"status,"`
	//
	Bindings []int64 `json:"bindings,"`
	//
	Children []string `json:"children,"`
	//
	Metadata interface{} `json:"metadata,"`
	//
	SnapshotInfo Snapshot `json:"snapshotInfo,omitempty"`
	//
	VolumeInfo Volume `json:"volumeInfo,omitempty"`
	//
	Descendants []int64 `json:"descendants,omitempty"`
}

type GetVolumeCountResult struct {
	//The number of volumes currently in the system.
	Count int64 `json:"count,"`
}

type ListClusterAdminsResult struct {
	//Information about the cluster admin.
	ClusterAdmins []ClusterAdmin `json:"clusterAdmins,"`
}

type ListISCSISessionsResult struct {
	//
	Sessions []ISCSISession `json:"sessions,"`
}

type ListVolumesForAccountResult struct {
	//List of volumes.
	Volumes []Volume `json:"volumes,"`
}

type NeedsWorkCreateDatabaseEntryResult struct {
}

type NetworkConfig struct {
	//
	Default bool `json:"#default,omitempty"`
	//
	Bond_master string `json:"bond-master,omitempty"`
	//
	VirtualNetworkTag string `json:"virtualNetworkTag,omitempty"`
	//
	Address string `json:"address,omitempty"`
	//
	Auto bool `json:"auto,omitempty"`
	//
	Bond_downdelay string `json:"bond-downdelay,omitempty"`
	//
	Bond_fail_over_mac string `json:"bond-fail_over_mac,omitempty"`
	//
	Bond_primary_reselect string `json:"bond-primary_reselect,omitempty"`
	//
	Bond_lacp_rate string `json:"bond-lacp_rate,omitempty"`
	//
	Bond_miimon string `json:"bond-miimon,omitempty"`
	//
	Bond_mode string `json:"bond-mode,omitempty"`
	//
	Bond_slaves string `json:"bond-slaves,omitempty"`
	//
	Bond_updelay string `json:"bond-updelay,omitempty"`
	//
	Dns_nameservers string `json:"dns-nameservers,omitempty"`
	//
	Dns_search string `json:"dns-search,omitempty"`
	//
	Family string `json:"family,omitempty"`
	//
	Gateway string `json:"gateway,omitempty"`
	//
	MacAddress string `json:"macAddress,omitempty"`
	//
	MacAddressPermanent string `json:"macAddressPermanent,omitempty"`
	//
	Method string `json:"method,omitempty"`
	//
	Mtu string `json:"mtu,omitempty"`
	//
	Netmask string `json:"netmask,omitempty"`
	//
	Network string `json:"network,omitempty"`
	//
	Physical PhysicalAdapter `json:"physical,omitempty"`
	//
	Routes []interface{} `json:"routes,omitempty"`
	//
	Status string `json:"status,omitempty"`
	//
	SymmetricRouteRules []string `json:"symmetricRouteRules,omitempty"`
	//
	UpAndRunning bool `json:"upAndRunning,omitempty"`
	//
	Bond_xmit_hash_policy string `json:"bond-xmit_hash_policy,omitempty"`
	//
	Bond_ad_num_ports string `json:"bond-ad_num_ports,omitempty"`
}

type RestoreDeletedVolumeResult struct {
}

type SnapMirrorVolume struct {
	//The ID of the destination ONTAP system.
	SnapMirrorEndpointID int64 `json:"snapMirrorEndpointID,"`
	//The name of the volume.
	Name string `json:"name,"`
	//The type of the volume.
	//Possible values:
	//rw: Read-write volume
	//ls: Loadsharing-volume
	//dp: Data protection volume
	Type string `json:"type,"`
	//The name of the Vserver that owns this volume.
	Vserver string `json:"vserver,"`
	//The containing aggregate name.
	AggrName string `json:"aggrName,"`
	//The state of volume.
	//Possible values:
	//online
	//restricted
	//offline
	//mixed
	State string `json:"state,"`
	//The total filesystem size (in bytes) of the volume.
	Size string `json:"size,"`
	//The size (in bytes) of the available space in the volume.
	AvailSize string `json:"availSize,"`
}

type DeleteStorageContainerResult struct {
}

type ListActiveVolumesResult struct {
	//List of active volumes.
	Volumes []Volume `json:"volumes,"`
}

type TestConnectSvipDetails struct {
	//Details of the ping tests with 56 Bytes and 1500 Bytes.
	PingBytes interface{} `json:"pingBytes,"`
	//The SVIP tested against.
	Svip string `json:"svip,"`
	//Whether the test could connect to the MVIP.
	Connected bool `json:"connected,"`
}

type NeedsWorkSetClusterFullThresholdsResult struct {
}

type OntapVersionInfo struct {
	//The ID of the destination ONTAP system.
	SnapMirrorEndpointID int64 `json:"snapMirrorEndpointID,"`
	//The ONTAP API major version in use by the SolidFire API client.
	ClientAPIMajorVersion string `json:"clientAPIMajorVersion,"`
	//The ONTAP API minor version in use by the SolidFire API client.
	ClientAPIMinorVersion string `json:"clientAPIMinorVersion,"`
	//The current API major version supported by the ONTAP system.
	OntapAPIMajorVersion string `json:"ontapAPIMajorVersion,"`
	//The current API minor version supported by the ONTAP system.
	OntapAPIMinorVersion string `json:"ontapAPIMinorVersion,"`
	//The current software version running on the ONTAP cluster.
	OntapVersion string `json:"ontapVersion,"`
}

type DeleteVirtualVolumesRequest struct {
	//The UUID of the volume to delete.
	VirtualVolumes []string `json:"virtualVolumes,"`
	//
	CallingVirtualVolumeHostID string `json:"callingVirtualVolumeHostID,omitempty"`
}

type ListStorageContainersResult struct {
	//
	StorageContainers []StorageContainer `json:"storageContainers,"`
}

type NeedsWorkDeleteVirtualVolumeHostsResult struct {
}

type NeedsWorkGetConnectivityReportResult struct {
}

type ReassignSlicesForZoneToleranceResult struct {
}

type CancelGroupCloneResult struct {
}

type DeleteSnapMirrorEndpointsResult struct {
}

type StartClusterBSCheckResult struct {
}

type VirtualVolumeBindResult struct {
	//Describes the VVol <-> Host binding.
	Binding VirtualVolumeBinding `json:"binding,"`
}

type SnapshotReplication struct {
	//The state of the snapshot replication.
	State string `json:"state,"`
	//Reserved for future use.
	StateDetails string `json:"stateDetails,"`
}

type SnapshotVirtualVolumeResult struct {
	//The volume ID of the newly-created clone.
	VolumeID int64 `json:"volumeID,"`
	//snapshotID for the prepared VVol snapshot.
	SnapshotID int64 `json:"snapshotID,"`
	//virtualVolumeID for the newly created clone.
	VirtualVolumeID string `json:"virtualVolumeID,"`
}

type LoginSessionInfo struct {
	//The time, in minutes, when this session will timeout and expire.
	//Formatted in H:mm:ss. For example: 1:30:00, 20:00, 5:00.
	//All leading zeros and colons are removed regardless of the format the timeout was entered.
	Timeout string `json:"timeout,"`
}

type NeedsWorkCheckClusterPairResult struct {
}

type PendingNodeToNode struct {
	//
	NodeID int64 `json:"nodeID,"`
	//
	PendingNodeID int64 `json:"pendingNodeID,"`
}

type GetScheduleResult struct {
	//The schedule attributes.
	Schedule Schedule `json:"schedule,"`
}

type GroupSnapshotMembers struct {
	//The source volume ID for the snapshot.
	VolumeID int64 `json:"volumeID,"`
	//Unique ID of a snapshot from which the new snapshot is made.
	//The snapshotID passed must be a snapshot on the given volume.
	SnapshotID int64 `json:"snapshotID,"`
	//A string that represents the correct digits in the stored snapshot.
	//This checksum can be used later to compare other snapshots to detect errors in the data.
	Checksum string `json:"checksum,"`
}

type GetClusterStateResult struct {
	//Array of NodeStateResult objects for each node in the cluster.
	Nodes []NodeStateResult `json:"nodes,omitempty"`
	//
	Cluster string `json:"cluster,omitempty"`
	//
	State string `json:"state,omitempty"`
}

type GetPendingOperationResult struct {
	//
	PendingOperation PendingOperation `json:"pendingOperation,"`
}

type PrepareVirtualSnapshotRequest struct {
	//The ID of the Virtual Volume to clone.
	VirtualVolumeID string `json:"virtualVolumeID,"`
	//The name for the newly-created volume.
	Name string `json:"name,omitempty"`
	//Will the snapshot be writable?
	WritableSnapshot bool `json:"writableSnapshot,omitempty"`
	//
	CallingVirtualVolumeHostID string `json:"callingVirtualVolumeHostID,omitempty"`
}

type UpdateAuthSessionResult struct {
	//SessionInfo for the updated auth session.
	Session AuthSessionInfo `json:"session,"`
}

type CreateClusterInterfacePreferenceResult struct {
}

type NewDrive struct {
	//A unique identifier for this drive.
	DriveID int64 `json:"driveID,"`
	//block or slice
	Type string `json:"type,omitempty"`
}

type NeedsWorkForceRecycleResult struct {
}

type BulkVolumeJob struct {
	//The internal bulk volume job ID.
	BulkVolumeID int64 `json:"bulkVolumeID,"`
	//Timestamp created for the bulk volume job.
	CreateTime string `json:"createTime,"`
	//The number of seconds since the job began.
	ElapsedTime int64 `json:"elapsedTime,"`
	//Format is either "compressed" or "native".
	Format string `json:"format,"`
	//The unique key created by the bulk volume session.
	Key string `json:"key,"`
	//The completed percentage reported by the operation.
	PercentComplete int64 `json:"percentComplete,"`
	//The estimated time remaining in seconds.
	RemainingTime int64 `json:"remainingTime,"`
	//The source volume ID.
	SrcVolumeID int64 `json:"srcVolumeID,"`
	//Can be one of the following:
	//preparing
	//active
	//done
	//failed
	Status string `json:"status,"`
	//The name of the script if one is provided.
	Script string `json:"script,omitempty"`
	//ID of the snapshot if a snapshot is in the source of the bulk volume job.
	SnapshotID int64 `json:"snapshotID,omitempty"`
	//Can be one of the following:
	//read
	//write
	Type string `json:"type,"`
	//JSON attributes on the bulk volume job.
	Attributes interface{} `json:"attributes,"`
}

type GetClusterVersionInfoResult struct {
	//
	ClusterAPIVersion string `json:"clusterAPIVersion,"`
	//
	ClusterVersion string `json:"clusterVersion,"`
	//
	ClusterVersionInfo []ClusterVersionInfo `json:"clusterVersionInfo,"`
	//
	SoftwareVersionInfo SoftwareVersionInfo `json:"softwareVersionInfo,"`
}

type NeedsWorkReassignSliceResult struct {
}

type QoS struct {
	//Desired minimum 4KB IOPS to guarantee.
	//The allowed IOPS will only drop below this level if all volumes have been capped
	//at their minimum IOPS value and there is still insufficient performance capacity.
	MinIOPS int64 `json:"minIOPS,omitempty"`
	//Desired maximum 4KB IOPS allowed over an extended period of time.
	MaxIOPS int64 `json:"maxIOPS,omitempty"`
	//Maximum "peak" 4KB IOPS allowed for short periods of time.
	//Allows for bursts of I/O activity over the normal max IOPS value.
	BurstIOPS int64 `json:"burstIOPS,omitempty"`
	//The length of time burst IOPS is allowed.
	//The value returned is represented in time units of seconds.
	//Note: this value is calculated by the system based on IOPS set for QoS.
	BurstTime int64 `json:"burstTime,omitempty"`
	//The curve is a set of key-value pairs.
	//The keys are I/O sizes in bytes.
	//The values represent the cost of performing an IOP at a specific I/O size.
	//The curve is calculated relative to a 4096 byte operation set at 100 IOPS.
	Curve interface{} `json:"curve,omitempty"`
}

type CreateMultipleVolumesResult struct {
	//Volume information for the newly created volumes.
	Volumes []Volume `json:"volumes,omitempty"`
	//VolumeIDs for the newly created volumes.
	VolumeIDs []int64 `json:"volumeIDs,"`
}

type ListServicesResult struct {
	//
	Services []DetailedService `json:"services,"`
}

type NeedsWorkDeleteDatabaseEntryResult struct {
}

type VasaProviderInfoResult struct {
	//
	VasaProviderInfo VasaProviderInfo `json:"vasaProviderInfo,"`
}

type NeedsWorkForceWholeFileSyncResult struct {
}

type NeedsWorkGetClusterStructureResult struct {
}

type NeedsWorkSetDatabaseEntryResult struct {
}

type ProposedClusterError struct {
	//Human readable description of the error.
	Description string `json:"description,"`
	//Unique error code for the error.
	ErrorCode ProposedNodeErrorCode `json:"errorCode,"`
	//The IP addresses of the nodes associated with the error.
	NodeIPs []string `json:"nodeIPs,"`
}

type ProtectionDomainLayout struct {
	//All of the different Protection Domains associated with each Node.
	ProtectionDomainLayout []NodeProtectionDomains `json:"ProtectionDomainLayout,"`
}

type RemoveKeyServerFromProviderKmipResult struct {
}

type VirtualVolumeBitmapResult struct {
	//The b64 bitmap.
	Bitmap string `json:"bitmap,"`
	//Byte length, adjusted to end on a chunk boundary.
	SegmentLength int64 `json:"segmentLength,"`
}

type GetAuthConfigurationResult struct {
	//
	AuthConfiguration AuthConfiguration `json:"authConfiguration,"`
}

type GetNodeFipsDrivesReportResult struct {
	//Node's FipsDrives capability status.
	FipsDrives FipsDrivesStatusType `json:"fipsDrives,"`
}

type ListDeletedVolumesResult struct {
	//List of deleted volumes.
	Volumes []Volume `json:"volumes,"`
}

type ListVirtualVolumeBindingsResult struct {
	//Describes the VVol <-> Host binding.
	Bindings []VirtualVolumeBinding `json:"bindings,"`
}

type ModifySnapshotResult struct {
	//
	Snapshot Snapshot `json:"snapshot,omitempty"`
}

type ModifyVirtualVolumeRequest struct {
	//VvolVolumeID for the volume to be modified.
	VirtualVolumeID string `json:"virtualVolumeID,"`
	//New quality of service settings for this volume.
	Qos QoS `json:"qos,omitempty"`
	//New size of the volume in bytes.
	//Size is rounded up to the nearest 1MiB size.
	//This parameter can only be used to *increase* the size of a volume.
	TotalSize int64 `json:"totalSize,omitempty"`
	//
	CallingVirtualVolumeHostID string `json:"callingVirtualVolumeHostID,omitempty"`
}

type AddNodesResult struct {
	//
	AutoInstall bool `json:"autoInstall,omitempty"`
	//An array of objects mapping the previous "pendingNodeID" to the "nodeID".
	Nodes []AddedNode `json:"nodes,"`
}

type GetQoSPolicyResult struct {
	//Details of the requested QoS policy.
	QosPolicy QoSPolicy `json:"qosPolicy,"`
}

type ModifyKeyServerKmipResult struct {
	//The resulting KMIP (Key Management Interoperability Protocol) Key Server after the modifications have been applied.
	KmipKeyServer KeyServerKmip `json:"kmipKeyServer,"`
}

type SnapshotRemoteStatus struct {
	//Current status of the remote snapshot on the target cluster as seen on the source cluster
	RemoteStatus RemoteClusterSnapshotStatus `json:"remoteStatus,"`
	//Universal identifier of the volume pair
	VolumePairUUID string `json:"volumePairUUID,"`
}

type VirtualVolumeStats struct {
	//AccountID of the volume owner.
	AccountID int64 `json:"accountID,"`
	//Current actual IOPS to the volume in the last 500 milliseconds.
	ActualIOPS int64 `json:"actualIOPS,omitempty"`
	//The length of time since the volume was last synced with the remote cluster.
	//If the volume is not paired, this is null.
	//
	//Note: A target volume in an active replication state always has an async delay of 0 (zero).
	//Target volumes are system-aware during replication and assume async delay is accurate at all times.
	AsyncDelay string `json:"asyncDelay,omitempty"`
	//Average size in bytes of recent I/O to the volume in the last 500 milliseconds.
	AverageIOPSize int64 `json:"averageIOPSize,omitempty"`
	//The total number of IOP credits available to the user.
	//When users are not using up to the max IOPS, credits are accrued.
	BurstIOPSCredit int64 `json:"burstIOPSCredit,omitempty"`
	//The number of outstanding read and write operations to the cluster.
	ClientQueueDepth int64 `json:"clientQueueDepth,omitempty"`
	//The volume services being migrated to if the volume metadata is getting migrated between volume services.
	//A "null" value means the volume is not migrating.
	DesiredMetadataHosts MetadataHosts `json:"desiredMetadataHosts,omitempty"`
	//The observed latency time, in microseconds, to complete operations to a volume.
	//A "0" (zero) value means there is no I/O to the volume.
	LatencyUSec int64 `json:"latencyUSec,omitempty"`
	//The volume services on which the volume metadata resides.
	MetadataHosts MetadataHosts `json:"metadataHosts,omitempty"`
	//The number of 4KiB blocks with data after the last garbage collection operation has completed.
	NonZeroBlocks int64 `json:"nonZeroBlocks,"`
	//Total bytes read by clients.
	ReadBytes int64 `json:"readBytes,"`
	//The average time, in microseconds, to complete read operations.
	ReadLatencyUSec int64 `json:"readLatencyUSec,omitempty"`
	//Total read operations.
	ReadOps int64 `json:"readOps,"`
	//A floating value between 0 and 1 that represents how much the system is throttling clients
	//below their max IOPS because of re-replication of data, transient errors and snapshots taken.
	Throttle float64 `json:"throttle,omitempty"`
	//The current time in UTC.
	Timestamp string `json:"timestamp,"`
	//The average time, in microseconds, to complete read and write operations to a volume.
	TotalLatencyUSec int64 `json:"totalLatencyUSec,omitempty"`
	//For 512e volumes, the number of read operations that were not on a 4k sector boundary.
	//High numbers of unaligned reads may indicate improper partition alignment.
	UnalignedReads int64 `json:"unalignedReads,"`
	//For 512e volumes, the number of write operations that were not on a 4k sector boundary.
	//High numbers of unaligned writes may indicate improper partition alignment.
	UnalignedWrites int64 `json:"unalignedWrites,"`
	//List of volume access group(s) to which a volume beintegers.
	VolumeAccessGroups []int64 `json:"volumeAccessGroups,"`
	//Volume ID of the volume.
	VolumeID int64 `json:"volumeID,"`
	//Total provisioned capacity in bytes.
	VolumeSize int64 `json:"volumeSize,"`
	//A floating value that describes how much the client is using the volume.
	//
	//Values:
	// 0 = Client is not using the volume
	//1 = Client is using their max
	//>1 = Client is using their burst
	VolumeUtilization float64 `json:"volumeUtilization,omitempty"`
	//Total bytes written by clients.
	WriteBytes int64 `json:"writeBytes,"`
	//The average time, in microseconds, to complete write operations.
	WriteLatencyUSec int64 `json:"writeLatencyUSec,omitempty"`
	//Total write operations occurring on the volume.
	WriteOps int64 `json:"writeOps,"`
	//Total number of 4KiB blocks without data after the last round of garbage collection operation has completed.
	ZeroBlocks int64 `json:"zeroBlocks,"`
	//The total number of bytes written to the volume during the last sample period.
	WriteBytesLastSample int64 `json:"writeBytesLastSample,omitempty"`
	//The length of the sample period in milliseconds.
	SamplePeriodMSec int64 `json:"samplePeriodMSec,omitempty"`
	//The total number of bytes read from the volume during the last sample period.
	ReadBytesLastSample int64 `json:"readBytesLastSample,omitempty"`
	//The total number of read operations durin gth elast sample period.
	ReadOpsLastSample int64 `json:"readOpsLastSample,omitempty"`
	//The total number of write operations during the last sample period.
	WriteOpsLastSample int64 `json:"writeOpsLastSample,omitempty"`
	//If the volume of interest is associated with a virtual volume, this is the virtual volume ID.
	VirtualVolumeID string `json:"virtualVolumeID,omitempty"`
}

type VolumeStats struct {
	//AccountID of the volume owner.
	AccountID int64 `json:"accountID,"`
	//Current actual IOPS to the volume in the last 500 milliseconds.
	ActualIOPS int64 `json:"actualIOPS,omitempty"`
	//Average size in bytes of recent I/O to the volume in the last 500 milliseconds.
	AverageIOPSize int64 `json:"averageIOPSize,omitempty"`
	//The total number of IOP credits available to the user.
	//When users are not using up to the max IOPS, credits are accrued.
	BurstIOPSCredit int64 `json:"burstIOPSCredit,omitempty"`
	//The number of outstanding read and write operations to the cluster.
	ClientQueueDepth int64 `json:"clientQueueDepth,omitempty"`
	//The observed latency time, in microseconds, to complete operations to a volume.
	//A "0" (zero) value means there is no I/O to the volume.
	LatencyUSec int64  `json:"latencyUSec,omitempty"`
	AsyncDelay  string `json:"asyncDelay,omitempty"`
	//The volume services on which the volume metadata resides.
	MetadataHosts        MetadataHosts `json:"metadataHosts,omitempty"`
	DesiredMetadataHosts MetadataHosts `json:"desiredMetadataHosts,omitempty"`
	//The number of 4KiB blocks with data after the last garbage collection operation has completed.
	NonZeroBlocks int64 `json:"nonZeroBlocks,"`
	//Total bytes read by clients.
	ReadBytes int64 `json:"readBytes,"`
	//The average time, in microseconds, to complete read operations.
	ReadLatencyUSec int64 `json:"readLatencyUSec,omitempty"`
	//Total read operations.
	ReadOps int64 `json:"readOps,"`
	//A floating value between 0 and 1 that represents how much the system is throttling clients
	//below their max IOPS because of re-replication of data, transient errors and snapshots taken.
	Throttle float64 `json:"throttle,omitempty"`
	//The current time in UTC.
	Timestamp string `json:"timestamp,"`
	//The average time, in microseconds, to complete read and write operations to a volume.
	TotalLatencyUSec int64 `json:"totalLatencyUSec,omitempty"`
	//For 512e volumes, the number of read operations that were not on a 4k sector boundary.
	//High numbers of unaligned reads may indicate improper partition alignment.
	UnalignedReads int64 `json:"unalignedReads,"`
	//For 512e volumes, the number of write operations that were not on a 4k sector boundary.
	//High numbers of unaligned writes may indicate improper partition alignment.
	UnalignedWrites int64 `json:"unalignedWrites,"`
	//List of volume access group(s) to which a volume beintegers.
	VolumeAccessGroups []int64 `json:"volumeAccessGroups,"`
	//Volume ID of the volume.
	VolumeID int64 `json:"volumeID,"`
	//Total provisioned capacity in bytes.
	VolumeSize int64 `json:"volumeSize,"`
	//A floating value that describes how much the client is using the volume.
	//
	//Values:
	// 0 = Client is not using the volume
	//1 = Client is using their max
	//>1 = Client is using their burst
	VolumeUtilization float64 `json:"volumeUtilization,omitempty"`
	//Total bytes written by clients.
	WriteBytes int64 `json:"writeBytes,"`
	//The average time, in microseconds, to complete write operations.
	WriteLatencyUSec int64 `json:"writeLatencyUSec,omitempty"`
	//Total write operations occurring on the volume.
	WriteOps int64 `json:"writeOps,"`
	//Total number of 4KiB blocks without data after the last round of garbage collection operation has completed.
	ZeroBlocks int64 `json:"zeroBlocks,"`
	//The total number of bytes written to the volume during the last sample period.
	WriteBytesLastSample int64 `json:"writeBytesLastSample,omitempty"`
	//The length of the sample period in milliseconds.
	SamplePeriodMSec int64 `json:"samplePeriodMSec,omitempty"`
	//The total number of bytes read from the volume during the last sample period.
	ReadBytesLastSample int64 `json:"readBytesLastSample,omitempty"`
	//The total number of read operations durin gth elast sample period.
	ReadOpsLastSample int64 `json:"readOpsLastSample,omitempty"`
	//The total number of write operations during the last sample period.
	WriteOpsLastSample int64 `json:"writeOpsLastSample,omitempty"`
}

type DetailedService struct {
	//
	Service Service `json:"service,"`
	//
	Node Node `json:"node,"`
	//
	Drive Drive `json:"drive,omitempty"`
	//
	Drives []Drive `json:"drives,"`
}

type KeyServerKmip struct {
	//If this KMIP Key Server is assigned to a provider, this field will contain the ID of the KMIP Key Provider it's assigned to.  Otherwise it will be null.
	KeyProviderID int64 `json:"keyProviderID,omitempty"`
	//If this KMIP Key Server is assigned to a provider (keyProviderID is not null), this field will indicate whether that provider is active (providing keys which are currently in use).  Otherwise it will be null.
	KmipAssignedProviderIsActive bool `json:"kmipAssignedProviderIsActive,omitempty"`
	//The public key certificate of the external key server's root CA.
	//This will be used to verify the certificate presented by external key server in the TLS communication.
	//For key server clusters where individual servers use different CAs, provide a concatenated string containing the root certificates of all the CAs.
	KmipCaCertificate string `json:"kmipCaCertificate,"`
	//A PEM format Base64 encoded PKCS#10 X.509 certificate used by the Solidfire KMIP client.
	KmipClientCertificate string `json:"kmipClientCertificate,"`
	//The hostnames or IP addresses associated with this KMIP Key Server.
	KmipKeyServerHostnames []string `json:"kmipKeyServerHostnames,"`
	//The ID of the KMIP Key Server.  This is a unique value assigned by the cluster during CreateKeyServer which cannot be changedKmip.
	KeyServerID int64 `json:"keyServerID,"`
	//The name of the KMIP Key Server.  This name is only used for display purposes and does not need to be unique.
	KmipKeyServerName string `json:"kmipKeyServerName,"`
	//The port number associated with this KMIP Key Server (typically 5696).
	KmipKeyServerPort int64 `json:"kmipKeyServerPort,"`
}

type ListSnapshotsResult struct {
	//Information about each snapshot for each volume.
	//Snapshots that are in a group will be returned with a "groupID".
	Snapshots []Snapshot `json:"snapshots,"`
}

type ListVolumeStatsResult struct {
	//List of volume activity information.
	VolumeStats []VolumeStats `json:"volumeStats,"`
}

type ModifyClusterInterfacePreferenceResult struct {
}

type ModifyVirtualVolumeMetadataRequest struct {
	//VvolVolumeID for the volume to be modified.
	VirtualVolumeID string `json:"virtualVolumeID,"`
	//
	Metadata interface{} `json:"metadata,"`
	//
	RemoveKeys []string `json:"removeKeys,"`
	//
	CallingVirtualVolumeHostID string `json:"callingVirtualVolumeHostID,omitempty"`
}

type NeedsWorkProtocolVersionUpgradeAvailableResult struct {
}

type ProtectionSchemeResiliency struct {
	//The Protection Scheme.
	ProtectionScheme ProtectionScheme `json:"protectionScheme,"`
	//The predicted number of simultaneous failures which may occur without losing
	//the ability to automatically heal to where the data has Node Tolerance.
	SustainableFailuresForBlockData int64 `json:"sustainableFailuresForBlockData,"`
	//The predicted number of simultaneous failures which may occur without losing
	//the ability to automatically heal to where the Metadata and Vvols have Node
	//Tolerance.
	SustainableFailuresForMetadata int64 `json:"sustainableFailuresForMetadata,"`
}

type CreateIdpConfigurationResult struct {
	//Information around the third party Identity Provider (IdP) configuration.
	IdpConfigInfo IdpConfigInfo `json:"idpConfigInfo,"`
}

type GetLdapConfigurationResult struct {
	//List of the current LDAP configuration settings. This API call will not return the plain text of the search account password.
	//
	//Note: If LDAP authentication is currently disabled, all the returned settings will be empty with the exception of "authType", and "groupSearchType" which are set to "SearchAndBind" and "ActiveDirectory" respectively.
	LdapConfiguration LdapConfiguration `json:"ldapConfiguration,"`
}

type Volume struct {
	//Unique VolumeID for the volume.
	VolumeID int64 `json:"volumeID,"`
	//Name of the volume as provided at creation time.
	Name string `json:"name,"`
	//Unique AccountID for the account.
	AccountID int64 `json:"accountID,"`
	//UTC formatted time the volume was created.
	CreateTime string `json:"createTime,"`
	//
	VolumeConsistencyGroupUUID string `json:"volumeConsistencyGroupUUID,"`
	//
	VolumeUUID string `json:"volumeUUID,"`
	//
	EnableSnapMirrorReplication bool `json:"enableSnapMirrorReplication,"`
	//Current status of the volume
	//init: A volume that is being initialized and is not ready for connections.
	//active: An active volume ready for connections.
	Status string `json:"status,"`
	//Access allowed for the volume
	Access VolumeAccess `json:"access,"`
	//If "true", the volume provides 512 byte sector emulation.
	Enable512e bool `json:"enable512e,"`
	//Volume iSCSI Qualified Name.
	Iqn string `json:"iqn,omitempty"`
	//Globally unique SCSI device identifier for the volume in EUI-64 based 16-byte format.
	ScsiEUIDeviceID string `json:"scsiEUIDeviceID,"`
	//Globally unique SCSI device identifier for the volume in NAA IEEE Registered Extended format.
	ScsiNAADeviceID string `json:"scsiNAADeviceID,"`
	//Quality of service settings for this volume.
	Qos VolumeQOS `json:"qos,"`
	//The QoS policy ID associated with the volume.
	//The value is null if the volume is not associated with a policy.
	QosPolicyID int64 `json:"qosPolicyID,omitempty"`
	//List of volume access groups to which a volume beintegers.
	VolumeAccessGroups []int64 `json:"volumeAccessGroups,"`
	//Information about a paired volume.
	//Available only if a volume is paired.
	//@see VolumePairs for return values.
	VolumePairs []VolumePair `json:"volumePairs,"`
	//The time this volume was deleted.
	//If this has no value, the volume has not yet been deleted.
	DeleteTime string `json:"deleteTime,omitempty"`
	//The time this volume will be purged from the system.
	//If this has no value, the volume has not yet been deleted (and is not scheduled for purging).
	PurgeTime string `json:"purgeTime,omitempty"`
	//The last time any access to this volume occurred.
	//If this has no value, the last access time is not known.
	LastAccessTime string `json:"lastAccessTime,omitempty"`
	//The last time I/O access to this volume occurred.
	//If this has no value, the last I/O access time is not known.
	LastAccessTimeIO string `json:"lastAccessTimeIO,omitempty"`
	//The number of slices backing this volume.
	//In the current software, this value will always be 1.
	SliceCount int64 `json:"sliceCount,"`
	//Total size of this volume in bytes.
	TotalSize int64 `json:"totalSize,"`
	//Size of the blocks on the volume.
	BlockSize int64 `json:"blockSize,"`
	//Virtual volume ID this volume backs.
	VirtualVolumeID string `json:"virtualVolumeID,omitempty"`
	//List of Name/Value pairs in JSON object format.
	Attributes interface{} `json:"attributes,"`
	//Protection scheme that is being used for this volume
	//If a volume is converting from one protection scheme to another,
	//this field will be set to the protection scheme that the volume is converting to.
	CurrentProtectionScheme ProtectionScheme `json:"currentProtectionScheme,"`
	//If a volume is converting from one protection scheme to another,
	//this field will be set to the protection scheme the volume is converting from.
	//This field will not change until another conversion is started.
	//If a volume has never been converted, this field will be null.
	PreviousProtectionScheme ProtectionScheme `json:"previousProtectionScheme,omitempty"`
}

type ScheduleInfo struct {
	Name     string  `json:"name,omitempty"`
	VolumeID int64   `json:"volumeID,omitempty"`
	Volumes  []int64 `json:"volumes,omitempty"`
	//Indicates if the snapshot should be included in remote replication.
	EnableRemoteReplication bool `json:"enableRemoteReplication,omitempty"`
	//The amount of time the snapshot will be retained in HH:mm:ss.
	Retention string `json:"retention,omitempty"`
}

type StartBulkVolumeReadResult struct {
	//ID of the async process to be checked for completion.
	AsyncHandle int64 `json:"asyncHandle,"`
	//Opaque key uniquely identifying the session.
	Key string `json:"key,"`
	//URL to access the node's web server
	Url string `json:"url,"`
}

type ResetDrivesResult struct {
	//Details of drives that are being reset.
	Details  ResetDrivesDetails `json:"details,"`
	Duration string             `json:"duration,omitempty"`
	Result   string             `json:"result,omitempty"`
}

type SetConfigResult struct {
	//The new and current configuration for the node.
	Config Config `json:"config,"`
}

type GetNodeHardwareInfoResult struct {
	//Hardware information for the specified nodeID.
	NodeHardwareInfo interface{} `json:"nodeHardwareInfo,"`
}

type NeedsWorkStartClusterBSCheckResult struct {
}

type TestConnectMvipResult struct {
	//Information about the test operation
	Details TestConnectMvipDetails `json:"details,"`
	//The length of time required to run the test.
	Duration string `json:"duration,"`
	//The results of the entire test
	Result string `json:"result,"`
}

type ClusterInfo struct {
	//
	MvipInterface string `json:"mvipInterface,omitempty"`
	//
	MvipVlanTag string `json:"mvipVlanTag,omitempty"`
	//
	SvipInterface string `json:"svipInterface,omitempty"`
	//
	SvipVlanTag string `json:"svipVlanTag,omitempty"`
	//Encryption at rest state.
	EncryptionAtRestState string `json:"encryptionAtRestState,"`
	//Array of Node IP addresses that are participating in the cluster.
	Ensemble []string `json:"ensemble,"`
	//Management network interface.
	Mvip string `json:"mvip,"`
	//Node holding the master MVIP address
	MvipNodeID int64 `json:"mvipNodeID,"`
	//Unique cluster name.
	Name string `json:"name,"`
	//Number of replicas of each piece of data to store in the cluster.
	RepCount int64 `json:"repCount,"`
	//A list of all of the protection schemes that are supported on this cluster.
	SupportedProtectionSchemes []ProtectionScheme `json:"supportedProtectionSchemes,"`
	//A list of all of the protection schemes that have been enabled on this cluster.
	EnabledProtectionSchemes []ProtectionScheme `json:"enabledProtectionSchemes,"`
	//If a protection scheme is not provided to the CreateVolume call, this protection scheme will be used for the new volume.
	//This protection scheme must always be in the set of enabled protection schemes.
	DefaultProtectionScheme ProtectionScheme `json:"defaultProtectionScheme,"`
	//Storage virtual IP
	Svip string `json:"svip,"`
	//Node holding the master SVIP address.
	SvipNodeID int64 `json:"svipNodeID,"`
	//Unique ID for the cluster.
	UniqueID string `json:"uniqueID,"`
	//
	Uuid string `json:"uuid,"`
	//List of Name/Value pairs in JSON object format.
	Attributes interface{} `json:"attributes,"`
}

type ModifyInitiatorsResult struct {
	//List of objects containing details about the modified initiators.
	Initiators []Initiator `json:"initiators,"`
}

type ListClusterNetworkInterfacesResult struct {
	//
	Nodes []NodeNetworkInterface `json:"nodes,"`
}

type NeedsWorkSetRepositoriesResult struct {
}

type QuintileHistogram struct {
	//Number of samples measured at 0 percent.
	Bucket0 int64 `json:"Bucket0,omitempty"`
	//Number of samples measured between 1 and 19 percent.
	Bucket1To19 int64 `json:"Bucket1To19,"`
	//Number of samples measured between 20 and 39 percent.
	Bucket20To39 int64 `json:"Bucket20To39,"`
	//Number of samples measured between 40 and 59 percent.
	Bucket40To59 int64 `json:"Bucket40To59,"`
	//Number of samples measured between 60 and 79 percent.
	Bucket60To79 int64 `json:"Bucket60To79,"`
	//Number of samples measured between 80 and 100 percent.
	Bucket80To100 int64 `json:"Bucket80To100,"`
	//Number of samples measured at greater than 100% percent.
	Bucket101Plus int64 `json:"Bucket101Plus,omitempty"`
}

type EnableEncryptionAtRestResult struct {
}

type GetIpmiInfoNodesResultObject struct {
	//
	IpmiInfo IpmiInfo `json:"ipmiInfo,"`
}

type NodeFibreChannelPortInfoResult struct {
	//The ID of the Fibre Channel node.
	NodeID int64 `json:"nodeID,"`
	//Contains a list of information about the Fibre Channel ports.
	Result FibreChannelPortList `json:"result,"`
}

type ProtectionDomain struct {
	//The type of the ProtectionDomain.
	ProtectionDomainType ProtectionDomainType `json:"protectionDomainType,"`
	//The name of the ProtectionDomain.
	ProtectionDomainName string `json:"protectionDomainName,"`
}

type RollbackVirtualVolumeRequest struct {
	//The ID of the Virtual Volume snapshot.
	SrcVirtualVolumeID string `json:"srcVirtualVolumeID,"`
	//The ID of the Virtual Volume to restore to.
	DstVirtualVolumeID string `json:"dstVirtualVolumeID,"`
	//
	CallingVirtualVolumeHostID string `json:"callingVirtualVolumeHostID,omitempty"`
}

type VirtualNetwork struct {
	//SolidFire unique identifier for a virtual network.
	VirtualNetworkID int64 `json:"virtualNetworkID,"`
	//VLAN Tag identifier.
	VirtualNetworkTag int64 `json:"virtualNetworkTag,"`
	//Range of address blocks currently assigned to the virtual network.
	//available: Binary string in "1"s and "0"s. 1 equals the IP is available and 0 equals the IP is not available. The string is read from right to left with the digit to the far right being the first IP address in the list of addressBlocks.
	//size: the size of this block of addresses.
	//start: first IP address in the block.
	AddressBlocks []AddressBlock `json:"addressBlocks,"`
	//The name assigned to the virtual network.
	Name string `json:"name,"`
	//IP address of the netmask for the virtual network.
	Netmask string `json:"netmask,"`
	//Storage IP address for the virtual network.
	Svip string `json:"svip,"`
	//
	Gateway string `json:"gateway,omitempty"`
	//
	Namespace bool `json:"namespace,omitempty"`
	//List of Name/Value pairs in JSON object format.
	Attributes interface{} `json:"attributes,omitempty"`
}

type GetClusterCapacityResult struct {
	//
	ClusterCapacity ClusterCapacity `json:"clusterCapacity,"`
}

type GetHardwareConfigResult struct {
	//List of hardware information and current settings.
	HardwareConfig interface{} `json:"hardwareConfig,"`
}

type ModifyAccountResult struct {
	//
	Account Account `json:"account,"`
}

type GetSupportedTlsCiphersResult struct {
	//List of mandatory TLS cipher suites for the cluster.
	//Mandatory ciphers are those ciphers which will always be active on the cluster.
	MandatoryCiphers []string `json:"mandatoryCiphers,"`
	//List of default supplemental TLS cipher suites for the cluster.
	//The supplemental ciphers will be restored to this list when the ResetSupplementalTlsCiphers command is run.
	DefaultSupplementalCiphers []string `json:"defaultSupplementalCiphers,"`
	//List of available supplemental TLS cipher suites which can be configured with the SetSupplementalTlsCiphers command.
	SupportedSupplementalCiphers []string `json:"supportedSupplementalCiphers,"`
}

type ListNodeFibreChannelPortInfoResult struct {
	//List of fibre channel port info results grouped by node.
	Nodes []NodeFibreChannelPortInfoResult `json:"nodes,omitempty"`
	//List of all physical Fibre Channel ports.
	FibreChannelPorts []FibreChannelPortInfo `json:"fibreChannelPorts,"`
}

type ListDriveStatsResult struct {
	//List of drive activity information for each drive.
	DriveStats []DriveStats `json:"driveStats,"`
	//If there are errors retrieving information about a drive, this list contains the driveID and associated error message. Always present, and empty if there are no errors.
	Errors []interface{} `json:"errors,"`
}

type NeedsWorkGetVolumeSetEfficiencyResult struct {
}

type NeedsWorkSetConstantsResult struct {
}

type SetClusterConfigResult struct {
	//Settings for the cluster. All new and current settings are returned.
	Cluster ClusterConfig `json:"cluster,"`
}

type VirtualVolumeNullResult struct {
}

type DeleteVolumeAccessGroupResult struct {
}

type EncryptionAtRestInfo struct {
	//The current generation of the password segment distribution which is in use.
	GenerationNumber int64 `json:"generationNumber,"`
	//The next generation of the password segment distribution which will become active if the distribution to enough nodes is successful.
	NextGenerationToStore int64 `json:"nextGenerationToStore,"`
	//The timestamp of when the password was originally created (regardless of how many times it has been redistributed or when).
	PasswordCreatedTimestamp int64 `json:"passwordCreatedTimestamp,"`
	//The number of pending password segment redistributions which still need to occur due to cluster changes.
	RedistributionCount int64 `json:"redistributionCount,"`
	//The current state of the Encryption At Rest feature, one of: disabled, enabling, enabled, or disabling.
	State string `json:"state,"`
	//If true, all drives on this cluster are required to be secure-enabled.
	SecureEnableRequired bool `json:"secureEnableRequired,omitempty"`
}

type GetOntapVersionInfoResult struct {
	//The software version information of the ONTAP endpoint.
	OntapVersionInfo []OntapVersionInfo `json:"ontapVersionInfo,"`
}

type ListNodeStatsResult struct {
	//Node activity information for all nodes.
	NodeStats NodeStatsNodes `json:"nodeStats,"`
}

type GetNodeSSLCertificateResult struct {
	//The full PEM-encoded test of the certificate.
	Certificate string `json:"certificate,"`
	//The decoded information of the certificate.
	Details interface{} `json:"details,"`
}

type NeedsWorkListAptSourceLinesResult struct {
}

type ListCurrentClusterAdminsResult struct {
	//Name of the calling user. In case of LDAP user, it is the user's DN
	Username string `json:"username,"`
	//List of ClusterAdmin objects
	ClusterAdmins []ClusterAdmin `json:"clusterAdmins,"`
}

type ListFibreChannelPortInfoResult struct {
	//Used to return information about the Fibre Channel ports.
	FibreChannelPortInfo FibreChannelPortInfoResult `json:"fibreChannelPortInfo,"`
}

type ListGroupSnapshotsResult struct {
	//List of Group Snapshots.
	GroupSnapshots []GroupSnapshot `json:"groupSnapshots,"`
}

type ModifyBackupTargetResult struct {
}

type NeedsWorkListDatabaseChildrenDataResult struct {
}

type SetDefaultQoSResult struct {
	//The minimum number of sustained IOPS that are provided by the cluster to a volume.
	MinIOPS int64 `json:"minIOPS,"`
	//The maximum number of sustained IOPS that are provided by the cluster to a volume.
	MaxIOPS int64 `json:"maxIOPS,"`
	//The maximum number of IOPS allowed in a short burst scenario.
	BurstIOPS int64 `json:"burstIOPS,"`
}

type Account struct {
	//Unique AccountID for the account.
	AccountID int64 `json:"accountID,"`
	//User name for the account.
	Username string `json:"username,"`
	//Current status of the account.
	Status string `json:"status,"`
	//List of VolumeIDs for Volumes owned by this account.
	Volumes []int64 `json:"volumes,"`
	//CHAP secret to use for the initiator.
	InitiatorSecret string `json:"initiatorSecret,omitempty"`
	//CHAP secret to use for the target (mutual CHAP authentication).
	TargetSecret string `json:"targetSecret,omitempty"`
	//The id of the storage container associated with the account
	StorageContainerID string `json:"storageContainerID,omitempty"`
	//List of Name/Value pairs in JSON object format.
	Attributes interface{} `json:"attributes,omitempty"`
}

type AsyncHandle struct {
	//The ID of the result.
	AsyncResultID int64 `json:"asyncResultID,"`
	//Returns true if it is completed and false if it isn't.
	Completed bool `json:"completed,"`
	//The time at which the asyncronous result was created
	CreateTime string `json:"createTime,"`
	//Time at which the result was last updated
	LastUpdateTime string `json:"lastUpdateTime,"`
	//The type of result. Could be Clone, DriveAdd, etc.
	ResultType string `json:"resultType,"`
	//Returns whether the result was a success or failure.
	Success bool `json:"success,"`
	//Attributes related to the result
	Data interface{} `json:"data,"`
}

type ListSnapMirrorPoliciesResult struct {
	//A list of the SnapMirror policies on the ONTAP storage system.
	SnapMirrorPolicies []SnapMirrorPolicy `json:"snapMirrorPolicies,"`
}

type GroupCloneVolumeMember struct {
	//The VolumeID of the cloned volume.
	VolumeID int64 `json:"volumeID,"`
	//The VolumeID of the source volume.
	SrcVolumeID int64 `json:"srcVolumeID,"`
}

type IdpConfigInfo struct {
	//UUID for the third party Identity Provider (IdP) Configuration.
	IdpConfigurationID string `json:"idpConfigurationID,"`
	//URL for retrieving IdP Metadata for configuration and integration details for SAML 2.0 single sign-on.
	IdpMetadataUrl string `json:"idpMetadataUrl,"`
	//URL for retrieving Service Provider (SP) Metadata from the Cluster to provide to the IdP for establish a trust relationship.
	SpMetadataUrl string `json:"spMetadataUrl,"`
	//A PEM format Base64 encoded PKCS#10 X.509 certificate to be used for communication with this IDP.
	ServiceProviderCertificate string `json:"serviceProviderCertificate,"`
	//Whether this third party Identity Provider configuration is enabled.
	Enabled bool `json:"enabled,"`
}

type NeedsWorkReadSliceLBAResult struct {
}

type VirtualVolumeTask struct {
	//
	VirtualVolumeTaskID string `json:"virtualVolumeTaskID,"`
	//
	VirtualvolumeID string `json:"virtualvolumeID,"`
	//
	CloneVirtualVolumeID string `json:"cloneVirtualVolumeID,"`
	//
	Status string `json:"status,"`
	//
	Operation string `json:"operation,"`
	//
	VirtualVolumeHostID string `json:"virtualVolumeHostID,"`
	//
	ParentMetadata interface{} `json:"parentMetadata,"`
	//
	ParentTotalSize int64 `json:"parentTotalSize,"`
	//
	ParentUsedSize int64 `json:"parentUsedSize,"`
	//
	Cancelled bool `json:"cancelled,"`
}

type ListFibreChannelSessionsResult struct {
	//A list of FibreChannelSession objects with information about the Fibre Channel session.
	Sessions []FibreChannelSession `json:"sessions,"`
}

type ListVolumeStatsByAccountResult struct {
	//List of account activity information.
	//Note: The volumeID member is 0 for each entry, as the values represent the summation of all volumes owned by the account.
	VolumeStats []VolumeStats `json:"volumeStats,"`
}

type StartClusterBSCheckRequest struct {
	//Handle value used to obtain the operation result.
	AsyncHandle bool `json:"asyncHandle,omitempty"`
	//The slice service on which cluster BS check will be started.
	ServiceID int64 `json:"serviceID,omitempty"`
	//The slice on which cluster BS check will be started.
	SliceID int64 `json:"sliceID,omitempty"`
	//The volume on which cluster BS check will be started.
	VolumeID int64 `json:"volumeID,omitempty"`
	//A list of slices on which cluster BS check will be started.
	SliceIDs []int64 `json:"sliceIDs,omitempty"`
	//A list of volumes on which cluster BS check will be started.
	VolumeIDs []int64 `json:"volumeIDs,omitempty"`
}

type ListSnapMirrorVolumesRequest struct {
	//List only the volumes associated
	//with the specified endpoint ID.
	//If no endpoint ID is provided, the system
	//lists volumes from all known SnapMirror endpoints.
	SnapMirrorEndpointID int64 `json:"snapMirrorEndpointID,omitempty"`
	//List volumes hosted on the specified Vserver.
	//The Vserver must be of type "data".
	Vserver string `json:"vserver,omitempty"`
	//List only ONTAP volumes with the specified name.
	Name string `json:"name,omitempty"`
	//List only ONTAP volumes of the specified type.
	//Possible values:
	//rw: Read-write volumes
	//ls: Load-sharing volumes
	//dp: Data protection volumes
	Type string `json:"type,omitempty"`
}

type ListUtilitiesRequest struct {
	Force        bool `json:"force,omitempty"`
	IncludeHints bool `json:"includeHints,omitempty"`
}

type CreateVolumeAccessGroupRequest struct {
	//The name for this volume access group. Not required to be unique, but recommended.
	Name string `json:"name,"`
	//List of initiators to include in the volume access group. If unspecified, the access group's configured initiators are not modified.
	Initiators []string `json:"initiators,omitempty"`
	//List of volumes to initially include in the volume access group. If unspecified, the access group's volumes are not modified.
	Volumes []int64 `json:"volumes,omitempty"`
	//The ID of the SolidFire virtual network to associate the volume access group with.
	VirtualNetworkID []int64 `json:"virtualNetworkID,omitempty"`
	//The ID of the SolidFire virtual network to associate the volume access group with.
	VirtualNetworkTags []int64 `json:"virtualNetworkTags,omitempty"`
	//List of name-value pairs in JSON object format.
	Attributes interface{} `json:"attributes,omitempty"`
}

type GetDriveHardwareInfoRequest struct {
	//DriveID for the drive information requested. You can get DriveIDs by using the ListDrives method.
	DriveID int64 `json:"driveID,"`
}

type ListSnapMirrorPoliciesRequest struct {
	//List only the policies associated
	//with the specified endpoint ID.
	//If no endpoint ID is provided, the system
	//lists policies from all known SnapMirror endpoints.
	SnapMirrorEndpointID int64 `json:"snapMirrorEndpointID,omitempty"`
}

type GetVolumeStatsRequest struct {
	//Specifies the volume for which statistics are gathered.
	VolumeID int64 `json:"volumeID,"`
}

type ListDeletedVolumesRequest struct {
	//Specifies that virtual volumes are included in the response by default.
	//To exclude virtual volumes, set to false.
	IncludeVirtualVolumes bool `json:"includeVirtualVolumes,omitempty"`
}

type AddClusterAdminRequest struct {
	//Unique username for this cluster admin. Must be between 1 and 1024 characters in length.
	Username string `json:"username,"`
	//Password used to authenticate this cluster admin.
	Password string `json:"password,"`
	//Controls which methods this cluster admin can use. For more details on the levels of access, see Access Control in the Element API Reference Guide.
	Access []string `json:"access,"`
	//Required to indicate your acceptance of the End User License
	//Agreement when creating this cluster. To accept the EULA,
	//set this parameter to true.
	AcceptEula bool `json:"acceptEula,"`
	//List of name-value pairs in JSON object format.
	Attributes interface{} `json:"attributes,omitempty"`
}

type SetConfigRequest struct {
	//Objects that you want changed for the cluster interface settings.
	Config ConfigParams `json:"config,"`
}

type ModifyVolumePairRequest struct {
	//The ID of the volume to be modified.
	VolumeID int64 `json:"volumeID,"`
	//Specifies whether to pause or restart volume replication process. Valid values are:
	//true: Pauses volume replication
	//false: Restarts volume replication
	PausedManual bool `json:"pausedManual,omitempty"`
	//Specifies the volume replication mode. Possible values are:
	//Async: Writes are acknowledged when they complete locally. The cluster does not wait for writes to be replicated to the target cluster.
	//Sync: The source acknowledges the write when the data is stored locally and on the remote cluster.
	//SnapshotsOnly: Only snapshots created on the source cluster are replicated. Active writes from the source volume are not replicated.
	Mode string `json:"mode,omitempty"`
	//Internal use only.
	PauseLimit int64 `json:"pauseLimit,omitempty"`
}

type GetStorageContainerEfficiencyRequest struct {
	//The ID of the storage container for which to retrieve efficiency information.
	StorageContainerID string `json:"storageContainerID,"`
}

type StartBulkVolumeWriteRequest struct {
	//The ID of the volume to be written to.
	VolumeID int64 `json:"volumeID,"`
	//The format of the volume data. It can be either of the following formats:
	//uncompressed: Every byte of the volume is returned without any compression.
	//native: Opaque data is returned that is smaller and more efficiently stored and written on a subsequent bulk
	//volume write.
	Format string `json:"format,"`
	//The executable name of a script. If unspecified,
	//the key and URL are necessary to access SF-series
	//nodes. The script runs on the primary node and the key
	//and URL is returned to the script, so the local web server
	//can be contacted.
	Script string `json:"script,omitempty"`
	//JSON parameters to pass to the script.
	ScriptParameters interface{} `json:"scriptParameters,omitempty"`
	//JSON attributes for the bulk volume job.
	Attributes interface{} `json:"attributes,omitempty"`
}

type RestartNetworkingRequest struct {
	//Required parameter to successfully reset the node.
	Force bool `json:"force,"`
}

type GetAPIRequest struct {
	//This tells the cluster to print out not only the public methods, but also the secret ones.
	V__includeSecretMethods bool `json:"__includeSecretMethods,"`
}

type ListKeyProvidersKmipRequest struct {
	//If omitted, returned KMIP Key Provider objects will not be filtered based on whether they're active.
	//If true, returns only KMIP Key Provider objects which are active (providing keys which are currently in use).
	//If false, returns only KMIP Key Provider objects which are inactive (not providing any keys and able to be deleted).
	KeyProviderIsActive bool `json:"keyProviderIsActive,omitempty"`
	//If omitted, returned KMIP Key Provider objects will not be filtered based on whether they have a KMIP Key Server assigned.
	//If true, returns only KMIP Key Provider objects which have a KMIP Key Server assigned.
	//If false, returns only KMIP Key Provider objects which do not have a KMIP Key Server assigned.
	KmipKeyProviderHasServerAssigned bool `json:"kmipKeyProviderHasServerAssigned,omitempty"`
}

type ModifyStorageContainerRequest struct {
	//The unique ID of the virtual volume storage container to modify.
	StorageContainerID string `json:"storageContainerID,"`
	//The new secret for CHAP authentication for the initiator.
	InitiatorSecret string `json:"initiatorSecret,omitempty"`
	//The new secret for CHAP authentication for the target.
	TargetSecret string `json:"targetSecret,omitempty"`
}

type CreateClusterInterfacePreferenceRequest struct {
	//Name of the cluster interface preference.
	Name string `json:"name,"`
	//Value of the cluster interface preference.
	Value string `json:"value,"`
}

type GetClusterHardwareInfoRequest struct {
	//Includes only a certain type of hardware information in
	//the response. Possible values are:
	//drives: List only drive information in the response.
	//nodes: List only node information in the response.
	//all: Include both drive and node information in the
	//response.
	//If this parameter is omitted, a type of "all" is assumed.
	Type string `json:"type,omitempty"`
}

type GetKeyProviderKmipRequest struct {
	//The ID of the KMIP Key Provider object to return.
	KeyProviderID int64 `json:"keyProviderID,"`
}

type ListVirtualNetworksRequest struct {
	//Network ID to filter the list for a single virtual network.
	VirtualNetworkID int64 `json:"virtualNetworkID,omitempty"`
	//Network tag to filter the list for a single virtual network.
	VirtualNetworkTag int64 `json:"virtualNetworkTag,omitempty"`
	//Network IDs to include in the list.
	VirtualNetworkIDs []int64 `json:"virtualNetworkIDs,omitempty"`
	//Network tag to include in the list.
	VirtualNetworkTags []int64 `json:"virtualNetworkTags,omitempty"`
}

type ListProtocolEndpointsRequest struct {
	//A list of protocol endpoint IDs for which to retrieve information. If you omit this parameter, the method returns information about all protocol endpoints.
	ProtocolEndpointIDs []string `json:"protocolEndpointIDs,omitempty"`
}

type CloneMultipleVolumesRequest struct {
	//Unique ID for each volume to include in the clone. If
	//optional parameters are not specified, the values are inherited from the source volumes.
	//Required parameter for "volumes" array:
	//volumeID
	//Optional parameters for "volumes" array:
	//access: Can be one of readOnly, readWrite,
	//locked, or replicationTarget
	//attributes: List of name-value pairs in JSON object
	//format.
	//name: New name for the clone.
	//newAccountID: Account ID for the new volumes.
	//newSize: New size Total size of the volume, in bytes.
	//Size is rounded up to the nearest 1MB.
	Volumes []CloneMultipleVolumeParams `json:"volumes,"`
	//New default access method for the new volumes if not
	//overridden by information passed in the volume's array.
	Access string `json:"access,omitempty"`
	//ID of the group snapshot to use as a basis for the clone.
	GroupSnapshotID int64 `json:"groupSnapshotID,omitempty"`
	//New account ID for the volumes if not overridden by
	//information passed in the volumes array.
	NewAccountID int64 `json:"newAccountID,omitempty"`
}

type ModifyClusterInterfacePreferenceRequest struct {
	//Name of the cluster interface preference.
	Name string `json:"name,"`
	//Value of the cluster interface preference.
	Value string `json:"value,"`
}

type CreateAuthSessionRequest struct {
	//Name that uniquely identifies the user. In case of an LDAP user it is user's DN.
	Username string `json:"username,"`
	//List of cluster adminIDs the user is assigned to
	//The list should not be empty
	ClusterAdminIDs []int64 `json:"clusterAdminIDs,"`
	//The length of time after which the session will expire.
	//This will only be used if it is less than the system defined max expiration time.
	DesiredExpirationDuration string `json:"desiredExpirationDuration,omitempty"`
}

type AddInitiatorsToVolumeAccessGroupRequest struct {
	//The ID of the volume access group
	//to modify.
	VolumeAccessGroupID int64 `json:"volumeAccessGroupID,"`
	//The list of initiators to add to the volume access
	//group.
	Initiators []string `json:"initiators,"`
}

type SetNodeSSLCertificateRequest struct {
	//The PEM-encoded text version of the certificate.
	Certificate string `json:"certificate,"`
	//The PEM-encoded text version of the private key.
	PrivateKey string `json:"privateKey,"`
}

type CreateVolumeRequest struct {
	//The name of the volume access group (might be user specified).
	//Not required to be unique, but recommended.
	//Might be 1 to 64 characters in length.
	Name string `json:"name,"`
	//AccountID for the owner of this volume.
	AccountID int64 `json:"accountID,"`
	//Total size of the volume, in bytes. Size is rounded up to
	//the nearest 1MB size.
	TotalSize int64 `json:"totalSize,"`
	//Specifies whether 512e emulation is enabled or not. Possible values are:
	//true: The volume provides 512-byte sector emulation.
	//false: 512e emulation is not enabled.
	Enable512e bool `json:"enable512e,omitempty"`
	//Initial quality of service settings for this volume. Default
	//values are used if none are specified. Valid settings are:
	//minIOPS
	//maxIOPS
	//burstIOPS
	//You can get the default values for a volume by using the GetDefaultQoS method.
	Qos QoS `json:"qos,omitempty"`
	//The list of name-value pairs in JSON object format.
	//Total attribute size must be less than 1000B, or 1KB,
	//including JSON formatting characters.
	Attributes interface{} `json:"attributes,omitempty"`
	//Not really used. The system limits sliceCount to a value of 1, which may be due to the currently acceptable number of slice drives in a storage node.
	SliceCount int64 `json:"sliceCount,omitempty"`
	//Associate the volume with the specified QoS policy.
	//Possible values:
	//true: Associate the volume with the QoS policy specified in the QoSPolicyID parameter.
	//false: Do not assosciate the volume with the QoS policy specified in the QoSPolicyID parameter. When false, any existing policy association is removed regardless of whether you specify a QoS policy in the QoSPolicyID parameter.
	AssociateWithQoSPolicy bool `json:"associateWithQoSPolicy,omitempty"`
	//The access mode for the volume. Only snapMirrorTarget is allowed.
	Access string `json:"access,omitempty"`
	//Specifies whether SnapMirror replication is enabled or not.
	EnableSnapMirrorReplication bool `json:"enableSnapMirrorReplication,omitempty"`
	//The ID for the policy whose QoS settings should be applied to the specified volumes.
	//This parameter is mutually exclusive with the qos parameter.
	QosPolicyID int64 `json:"qosPolicyID,omitempty"`
	//Protection scheme that should be used for this volume.
	//The default value is the defaultProtectionScheme stored in the ClusterInfo object.
	ProtectionScheme ProtectionScheme `json:"protectionScheme,omitempty"`
}

type ListEventsRequest struct {
	//Specifies the maximum number of events to return.
	MaxEvents int64 `json:"maxEvents,omitempty"`
	//Specifies the beginning of a range of events to return.
	StartEventID int64 `json:"startEventID,omitempty"`
	//Specifies the end of a range of events to return.
	EndEventID int64 `json:"endEventID,omitempty"`
	//Specifies the type of event queue, can be either standard or vvol.
	EventQueueType string `json:"eventQueueType,omitempty"`
	//Specifies the type of events to return.
	EventType string `json:"eventType,omitempty"`
	//Specifies that only events with this ServiceID will be returned.
	ServiceID int64 `json:"serviceID,omitempty"`
	//Specifies that only events with this NodeID will be returned.
	NodeID int64 `json:"nodeID,omitempty"`
	//Specifies that only events with this DriveID will be returned.
	DriveID int64 `json:"driveID,omitempty"`
	//Specifies that only events reported after this time will be returned.
	StartReportTime string `json:"startReportTime,omitempty"`
	//Specifies that only events reported earlier than this time will be returned.
	EndReportTime string `json:"endReportTime,omitempty"`
	//Specifies that only events published after this time will be returned.
	StartPublishTime string `json:"startPublishTime,omitempty"`
	//Specifies that only events published earlier than this time will be returned.
	EndPublishTime string `json:"endPublishTime,omitempty"`
}

type AbortRecoverDeadVolumesRequest struct {
	//Volumes with ongoing recovery operations that should be aborted.
	VolumeIDs []int64 `json:"volumeIDs,omitempty"`
	//Slices with ongoing recovery operations that should be aborted.
	SliceIDs []int64 `json:"sliceIDs,omitempty"`
}

type SetProtectionDomainLayoutRequest struct {
	//The Protection Domains for each Node.
	ProtectionDomainLayout ProtectionDomainLayout `json:"protectionDomainLayout,"`
	//Allows setting the Protection Domains in a cluster with a witness node.
	ForceForWitnessNode bool `json:"forceForWitnessNode,omitempty"`
	//Allows setting the Protection Domains with a mixture of default and non-default names.
	ForceIncompleteProtectionDomainLayout bool `json:"forceIncompleteProtectionDomainLayout,omitempty"`
}

type RecoverDeadVolumesRequest struct {
	//Volumes to be recovered.
	VolumeIDs []int64 `json:"volumeIDs,omitempty"`
	//Slices to be recovered.
	SliceIDs []int64 `json:"sliceIDs,omitempty"`
}

type ClearClusterFaultsRequest struct {
	//Determines the types of faults cleared. Possible values are:
	//current: Faults that are currently detected and have
	//not been resolved.
	//resolved: (Default) Faults that were previously
	//detected and resolved.
	//all: Both current and resolved faults are cleared. The
	//fault status can be determined by the resolved field of
	//the fault object.
	FaultTypes string `json:"faultTypes,omitempty"`
}

type GetOriginRequest struct {
	//Force this method to run on all nodes. Only needed when issuing the API to a cluster instead of a single node.
	Force bool `json:"force,"`
}

type RollbackToSnapshotRequest struct {
	//VolumeID for the volume.
	VolumeID int64 `json:"volumeID,"`
	//The ID of a previously created snapshot on the given volume.
	SnapshotID int64 `json:"snapshotID,"`
	//Specifies whether to save an active volume image or delete it. Values are:
	//true: The previous active volume image is kept.
	//false: (default) The previous active volume image is deleted.
	SaveCurrentState bool `json:"saveCurrentState,"`
	//Name for the snapshot. If unspecified, the name of the snapshot being rolled back to is used with "- copy" appended to the end of the name.
	Name string `json:"name,omitempty"`
	//List of name-value pairs in JSON object format.
	Attributes interface{} `json:"attributes,omitempty"`
}

type TestPingRequest struct {
	//Specifies the number of times the system
	//should repeat the test ping. The default value is 5.
	Attempts int64 `json:"attempts,omitempty"`
	//Specifies a comma-separated list of addresses or hostnames of devices to ping.
	Hosts string `json:"hosts,omitempty"`
	//Specifies the length of time the ping should wait for a system response before issuing the next ping attempt or ending the process.
	TotalTimeoutSec int64 `json:"totalTimeoutSec,omitempty"`
	//Specifies the number of bytes to send in the ICMP packet that is sent to each IP. The number must be less than the maximum MTU specified in the network configuration.
	PacketSize int64 `json:"packetSize,omitempty"`
	//Specifies the number of milliseconds to wait for each individual ping response. The default value is 500 ms.
	PingTimeoutMsec int64 `json:"pingTimeoutMsec,omitempty"`
	//Specifies that the Do not Fragment (DF) flag is enabled for the ICMP packets.
	ProhibitFragmentation bool `json:"prohibitFragmentation,omitempty"`
	Force                 bool `json:"force,omitempty"`
}

type CreateSupportBundleRequest struct {
	//The unique name for the support bundle. If no name is provided, "supportbundle" and the node name are used as the filename.
	BundleName string `json:"bundleName,omitempty"`
	//Passed to the sf_make_support_bundle script. You should use this parameter only at the request of NetApp SolidFire Support.
	ExtraArgs string `json:"extraArgs,omitempty"`
	//The number of seconds to allow the support bundle script to run before stopping. The default value is 1500 seconds.
	TimeoutSec int64 `json:"timeoutSec,omitempty"`
}

type TestKeyServerKmipRequest struct {
	//The ID of the KMIP Key Server to test.
	KeyServerID int64 `json:"keyServerID,"`
}

type ModifyVolumeAccessGroupLunAssignmentsRequest struct {
	//The ID of the volume access group for which the LUN assignments will be modified.
	VolumeAccessGroupID int64 `json:"volumeAccessGroupID,"`
	//The volume IDs with new assigned LUN values.
	LunAssignments []LunAssignment `json:"lunAssignments,"`
}

type EnableProtectionSchemesRequest struct {
	//The protection schemes that should be enabled
	ProtectionSchemes []ProtectionScheme `json:"protectionSchemes,"`
}

type SetClusterConfigRequest struct {
	//Objects that are changed for the cluster interface settings.
	Cluster ClusterConfig `json:"cluster,"`
	//
	Force              bool `json:"force,omitempty"`
	V__useTheForceLuke bool `json:"__useTheForceLuke,omitempty"`
}

type ListSnapMirrorVserversRequest struct {
	//List only the Vservers associated with the specified endpoint ID.
	//If no endpoint ID is provided, the system lists
	//Vservers from all known SnapMirror endpoints.
	SnapMirrorEndpointID int64 `json:"snapMirrorEndpointID,omitempty"`
	//List only Vservers of the specified type.
	//Possible values:
	//admin
	//data
	//node
	//system
	VserverType string `json:"vserverType,omitempty"`
	//List only Vservers with the specified name.
	VserverName string `json:"vserverName,omitempty"`
}

type CloneVolumeFromDeadRequest struct {
	//VolumeID for the volume to be cloned.
	VolumeID int64 `json:"volumeID,"`
	//service ID of the dead secondary slice service from which to clone
	SrcServiceID int64 `json:"srcServiceID,"`
	//The name of the new cloned volume. Must be 1 to 64 characters in length.
	Name string `json:"name,"`
	//AccountID for the owner of the new volume. If unspecified, the
	//accountID of the owner of the volume being cloned is used.
	NewAccountID int64 `json:"newAccountID,omitempty"`
	//New size of the volume, in bytes. Must be greater or less than the size of
	//the volume being cloned. If unspecified, the volume size is not
	//changed. Size is rounded to the nearest 1MB.
	NewSize int64 `json:"newSize,omitempty"`
	//Specifies the level of access allowed for the new volume. If unspecified, the
	//level of access of the volume being cloned is used. If replicationTarget is
	//is passed and the volume is not paired, the access gets set to locked.
	Access VolumeAccess `json:"access,omitempty"`
	//ID of the snapshot that is used as the source of the clone. If no ID is
	//provided, the current active volume is used.
	SnapshotID int64 `json:"snapshotID,omitempty"`
	//List of name-value pairs in JSON object format.
	Attributes interface{} `json:"attributes,omitempty"`
	//Specifies whether the new volume should use 512-byte sector emulation.
	//If unspecified, the setting of the volume being cloned is used.
	Enable512e bool `json:"enable512e,omitempty"`
}

type GetNvramInfoRequest struct {
	//Required parameter to successfully run on all
	//nodes in the cluster.
	Force bool `json:"force,omitempty"`
}

type GetNetworkConfigRequest struct {
	Force        bool `json:"force,omitempty"`
	IncludeHints bool `json:"includeHints,omitempty"`
}

type SetSnmpACLRequest struct {
	//List of networks and what type of access they have to the SNMP servers running on the cluster nodes. See SNMP
	//Network Object for possible "networks" values. This parameter is required if SNMP v3 is disabled.
	Networks []SnmpNetwork `json:"networks,"`
	//List of users and the type of access they have to the SNMP servers running on the cluster nodes.
	UsmUsers []SnmpV3UsmUser `json:"usmUsers,"`
}

type ListClusterAdminsRequest struct {
	//Shows hidden cluster admin users such as the SNMP admin.
	ShowHidden bool `json:"showHidden,omitempty"`
}

type GetConfigRequest struct {
	Force        bool `json:"force,omitempty"`
	IncludeHints bool `json:"includeHints,omitempty"`
}

type GetKeyServerKmipRequest struct {
	//The ID of the KMIP Key Server object to return.
	KeyServerID int64 `json:"keyServerID,"`
}

type UnbindVirtualVolumesRequest struct {
	//Normal, Start, or End?
	UnbindContext string `json:"unbindContext,"`
	//The UUID of the host making the call to the virtual volume.
	VirtualVolumeHostID string `json:"virtualVolumeHostID,"`
	//A list of objects containing the keys virtualVolumID, protocolEndpoindType, protocolEndpoingInBandID, and virtualVolumeSecondaryID. All of these keys are required in each of the objects.
	UnbindArgs []UnbindArguments `json:"unbindArgs,"`
}

type GetNodeHardwareInfoRequest struct {
	//The ID of the node for which hardware information is being requested. Information
	//about a Fibre Channel node is returned if a Fibre Channel node is specified.
	NodeID int64 `json:"nodeID,"`
}

type ListVolumeStatsByVolumeAccessGroupRequest struct {
	//An array of VolumeAccessGroupIDs for which volume
	//activity is returned. If omitted, statistics for all volume
	//access groups are returned.
	VolumeAccessGroups []int64 `json:"volumeAccessGroups,omitempty"`
	//Specifies that virtual volumes are included in the response by default.
	//To exclude virtual volumes, set to false.
	IncludeVirtualVolumes bool `json:"includeVirtualVolumes,omitempty"`
}

type StartBulkVolumeReadRequest struct {
	//The ID of the volume to be read.
	VolumeID int64 `json:"volumeID,"`
	//The format of the volume data. It can be either of the following formats:
	//uncompressed: Every byte of the volume is returned without any compression.
	//native: Opaque data is returned that is smaller and more efficiently stored and written on a subsequent bulk
	//volume write.
	Format string `json:"format,"`
	//The ID of a previously created snapshot used for bulk volume
	//reads. If no ID is entered, a snapshot of the current active
	//volume image is made.
	SnapshotID int64 `json:"snapshotID,omitempty"`
	//The executable name of a script. If unspecified, the key and URL is necessary to access SF-series nodes. The script is run on the primary node and the key
	//and URL is returned to the script so the local web server
	//can be contacted.
	Script string `json:"script,omitempty"`
	//JSON parameters to pass to the script.
	ScriptParameters interface{} `json:"scriptParameters,omitempty"`
	//JSON attributes for the bulk volume job.
	Attributes interface{} `json:"attributes,omitempty"`
}

type InitializeSnapMirrorRelationshipRequest struct {
	//The ID of the remote ONTAP system.
	SnapMirrorEndpointID int64 `json:"snapMirrorEndpointID,"`
	//The destination volume's name in the SnapMirror relationship.
	DestinationVolume SnapMirrorVolumeInfo `json:"destinationVolume,"`
	//Specifies the maximum data transfer rate
	//between the volumes in kilobytes per second.
	//The default value, 0, is unlimited and permits the SnapMirror
	//relationship to fully utilize the available network bandwidth.
	MaxTransferRate int64 `json:"maxTransferRate,omitempty"`
}

type SecureEraseDrivesRequest struct {
	//List of driveIDs to be secure erased.
	Drives []int64 `json:"drives,"`
}

type GetIpmiConfigRequest struct {
	//Displays information for each node chassis type.
	//Valid values are:
	//all: Returns sensor information for each chassis type.
	//{chassis type}: Returns sensor information for a specified chassis type.
	ChassisType string `json:"chassisType,omitempty"`
	//Force this method to run on all nodes. Only needed when issuing the API to a cluster instead of a single node.
	Force bool `json:"force,"`
}

type CreateScheduleRequest struct {
	//Specifies a unique name for the schedule.
	ScheduleName string `json:"scheduleName,"`
	//Specifies the days of the month that a snapshot will be created. Valid values are 1 through 31.
	Monthdays []int64 `json:"monthdays,omitempty"`
	//Specifies the day of the week the snapshot will be created.
	//Required Values (if used) are:
	//Day: 0 - 6 (Sunday - Saturday).
	//Offset: 1
	Weekdays int64 `json:"weekdays,omitempty"`
	//Specifies the number of hours between snapshots or hour in GMT time that the snapshot will occur in Days of Week or Days of Month mode. Valid values are 0 through 23.
	Hours int64 `json:"hours,omitempty"`
	//Specifies the number of minutes between snapshots or the minute in GMT time that the snapshot will occur in Days of Week or Days of Month mode. Valid values are 0 through 59.
	Minutes int64 `json:"minutes,omitempty"`
	//Indicates the type of schedule to create.
	//Valid value is:
	//snapshot
	ScheduleType string `json:"scheduleType,"`
	//Use the "frequency" string to indicate the frequency of the snapshot.
	//Valid values for "frequency" are:
	//Days of Week
	//Days of Month
	//Time Interval
	Attributes interface{} `json:"attributes,"`
	ScheduleID int64       `json:"scheduleID,"`
	//An object of schedule information about how the snapshot should be created at each scheduled interval. Values are:
	//volumeID: The ID of the volume to be included in the snapshot. (Integer)
	//volumes: A list of volume IDs to be included in the group snapshot. (Array of Integers)
	//name: The snapshot name to be used. (String)
	//enableRemoteReplication:  Indicates if the snapshot should be included in remote
	//replication. (Boolean)
	//retention: The amount of time the snapshot will be retained in HH:mm:ss. (String)
	ScheduleInfo    ScheduleInfo `json:"scheduleInfo,omitempty"`
	LastRunStatus   string       `json:"lastRunStatus,omitempty"`
	RunNextInterval bool         `json:"runNextInterval,omitempty"`
	//Indicates if the schedule should be paused or not.
	//Valid values are:
	//true
	//false
	Paused bool `json:"paused,omitempty"`
	//Indicates if the schedule recurs or not.
	//Valid values are:
	//true
	//false
	Recurring bool `json:"recurring,omitempty"`
	HasError  bool `json:"hasError,omitempty"`
	//Specifies the UTC time after which the schedule will be run. If unspecified, the schedule starts immediately.
	StartingDate string `json:"startingDate,omitempty"`
	ToBeDeleted  bool   `json:"toBeDeleted,omitempty"`
	//Label used by SnapMirror software to specify snapshot retention policy on SnapMirror endpoint.
	SnapMirrorLabel string `json:"snapMirrorLabel,omitempty"`
}

type ModifyBackupTargetRequest struct {
	//The unique target ID for the target to modify.
	BackupTargetID int64 `json:"backupTargetID,"`
	//The new name for the backup target.
	Name string `json:"name,omitempty"`
	//List of name-value pairs in JSON object format.
	Attributes interface{} `json:"attributes,omitempty"`
}

type ModifyQoSPolicyRequest struct {
	//The ID of the policy to be modified.
	QosPolicyID int64 `json:"qosPolicyID,"`
	//If supplied, the name of the QoS Policy (e.g. gold, platinum, silver) is changed to this value.
	Name string `json:"name,omitempty"`
	//If supplied, the QoS settings for this policy are changed to these sttings.
	//You can supply partial QoS values and only change some of the QoS settings.
	Qos QoS `json:"qos,omitempty"`
}

type ModifyInitiatorsRequest struct {
	//A list of objects containing characteristics of each initiator to modify.
	Initiators []ModifyInitiator `json:"initiators,"`
}

type GetNodeConstantsRequest struct {
	//Only return the values for the list of provided constants.
	//If not provided, the value of all constants will be returned.
	Constants []string `json:"constants,omitempty"`
	//If true, only return the set of constants that have changed from their default value.
	ChangedOnly bool `json:"changedOnly,omitempty"`
}

type TestConnectSvipRequest struct {
	//If specified, tests the storage connection of a
	//different SVIP. You do not need to use this value when
	//testing the connection to the target cluster. This parameter is optional.
	Svip  string `json:"svip,omitempty"`
	Force bool   `json:"force,omitempty"`
}

type ListDriveStatsRequest struct {
	//Optional list of DriveIDs for which to return drive
	//statistics. If you omit this parameter, measurements for
	//all drives are returned.
	Drives []int64 `json:"drives,omitempty"`
}

type ListActiveVolumesRequest struct {
	//Starting VolumeID to return. If no volume exists with this
	//VolumeID, the next volume by VolumeID order is used as
	//the start of the list. To page through the list, pass the
	//VolumeID of the last volume in the previous response +
	//1.
	StartVolumeID int64 `json:"startVolumeID,omitempty"`
	//Maximum number of Volume Info objects to return. A value of 0
	//(zero) returns all volumes (unlimited).
	Limit int64 `json:"limit,omitempty"`
	//Specifies that virtual volumes are included in the response by default.
	//To exclude virtual volumes, set to false.
	IncludeVirtualVolumes bool `json:"includeVirtualVolumes,omitempty"`
}

type ListIdpConfigurationsRequest struct {
	//Filters the result to an IdP configuration information for a specific IdP.
	IdpMetadataUrl string `json:"idpMetadataUrl,omitempty"`
}

type CreateSnapMirrorVolumeRequest struct {
	//The endpoint ID of the remote ONTAP storage system communicating with the SolidFire cluster.
	SnapMirrorEndpointID int64 `json:"snapMirrorEndpointID,"`
	//The name of the Vserver.
	Vserver string `json:"vserver,"`
	//The destination ONTAP volume name.
	Name string `json:"name,"`
	//The volume type.
	//Possible values:
	//rw: Read-write volume
	//ls: Load-sharing volume
	//dp: Data protection volume
	//If no type is provided the default type is dp.
	Type string `json:"type,omitempty"`
	//The containing ONTAP aggregate in which to create the volume.
	//You can use ListSnapMirrorAggregates to get
	//information about available ONTAP aggregates.
	Aggregate string `json:"aggregate,"`
	//The size of the volume in bytes.
	Size int64 `json:"size,"`
}

type DeleteSnapMirrorRelationshipsRequest struct {
	//The endpoint ID of the remote ONTAP storage system communicating
	//with the SolidFire cluster.
	SnapMirrorEndpointID int64 `json:"snapMirrorEndpointID,"`
	//The destination volume or volumes in the SnapMirror relationship.
	DestinationVolumes []SnapMirrorVolumeInfo `json:"destinationVolumes,"`
}

type ModifyVolumeRequest struct {
	//VolumeID for the volume to be modified.
	VolumeID int64 `json:"volumeID,"`
	//AccountID to which the volume is reassigned. If unspecified, the previous account name is used.
	AccountID int64 `json:"accountID,omitempty"`
	//Specifies the access allowed for the volume. Possible values are:
	//readOnly: Only read operations are allowed.
	//readWrite: Reads and writes are allowed.
	//locked: No reads or writes are allowed.
	//If not specified, the access value does not change.
	//replicationTarget: Identify a volume as the target volume
	//for a paired set of volumes. If the volume is not paired, the
	//access status is locked.
	//If a value is not specified, the access value does not change.
	Access string `json:"access,omitempty"`
	//New QoS settings for this volume. If not specified,
	//the QoS settings are not changed.
	Qos QoS `json:"qos,omitempty"`
	//New size of the volume in bytes. 1000000000 is equal to 1GB.
	//Size is rounded up to the nearest 1MB. This parameter
	//can only be used to increase the size of a volume.
	TotalSize int64 `json:"totalSize,omitempty"`
	//If set to true, changes the recorded date of volume creation.
	SetCreateTime bool `json:"setCreateTime,omitempty"`
	//An ISO 8601 date string to set as the new volume creation date.
	//Required if "setCreateTime" is set to true.
	CreateTime string `json:"createTime,omitempty"`
	//List of name-value pairs in JSON object format.
	Attributes interface{} `json:"attributes,omitempty"`
	//Associate the volume with the specified QoS policy.
	//Possible values:
	//true: Associate the volume with the QoS policy specified in the QoSPolicyID parameter.
	//false: Do not assosciate the volume with the QoS policy specified in the QoSPolicyID parameter. When false, any existing policy association is removed regardless of whether you specify a QoS policy in the QoSPolicyID parameter.
	AssociateWithQoSPolicy bool `json:"associateWithQoSPolicy,omitempty"`
	//The ID for the policy whose QoS settings should be applied to the specified volumes.
	//The volume will not maintain any association with the policy; this is an alternate way to apply QoS settings to the volume.
	//This parameter and the qos parameter cannot be specified at the same time.
	QosPolicyID int64 `json:"qosPolicyID,omitempty"`
	//Determines whether the volume can be used for replication with SnapMirror endpoints.
	//Possible values:
	//true
	//false
	EnableSnapMirrorReplication bool `json:"enableSnapMirrorReplication,omitempty"`
}

type CreateIdpAuthSessionRequest struct {
	//List of attribute name-value pairs for determining IdP Cluster AdminID(s) for the
	//session to be created.
	SamlAttributeStatements []string `json:"samlAttributeStatements,"`
	//The length of time after which the session will expire.
	//This will only be used if it is less than the system defined max expiration time.
	DesiredExpirationDuration string `json:"desiredExpirationDuration,omitempty"`
	//Arbitrary string to uniquely identify the user and will be used for auditing
	//operations within the session by that user.  For example, a SAML Subject
	//NameID=Value pairing, but this will be dictated by the configuration of the
	//IdP and the resultant content of the SAML assertion.
	Username string `json:"username,"`
}

type DeleteStorageContainersRequest struct {
	//A list of IDs of the storage containers to delete. You can specify up to 2000 IDs in the list.
	StorageContainerIDs []string `json:"storageContainerIDs,"`
}

type ListVirtualVolumeTasksRequest struct {
	//A list of virtual volume task IDs for which to retrieve information. If you omit this parameter, the method returns information about all virtual volume tasks.
	VirtualVolumeTaskIDs []string `json:"virtualVolumeTaskIDs,omitempty"`
}

type GetNodeStatsRequest struct {
	//Specifies the node for which statistics are gathered.
	NodeID int64 `json:"nodeID,"`
}

type ListVolumeQoSHistogramsRequest struct {
	//List of volumes to return data for.
	//If no volumes are specified then information for all volumes will be returned.
	VolumeIDs int64 `json:"volumeIDs,omitempty"`
}

type ResetNodeRequest struct {
	//Specifies the URL to a remote Element software image to which the node will
	//be reset.
	Build string `json:"build,"`
	//Required parameter to successfully reset the node.
	Force bool `json:"force,"`
	//Set to true if you want to reboot the node.
	Reboot bool `json:"reboot,omitempty"`
	//Used to enter specifications for running the reset operation.
	//Available options:
	//'edebug': '',
	//'sf_auto': '0',
	//'sf_bond_mode': 'ActivePassive',
	//'sf_check_hardware': '0',
	//'sf_disable_otpw': '0',
	//'sf_fa_host': '',
	//'sf_hostname': 'SF-FA18',
	//'sf_inplace': '1',
	//'sf_inplace_die_action': 'kexec',
	//'sf_inplace_safe': '0',
	//'sf_keep_cluster_config': '0',
	//'sf_keep_data': '0',
	//'sf_keep_hostname': '0',
	//'sf_keep_network_config': '0',
	//'sf_keep_paths': '/var/log/hardware.xml
	//'sf_max_archives': '5',
	//'sf_nvram_size': '',
	//'sf_oldroot': '',
	//'sf_postinst_erase_root_drive': '0',
	//'sf_root_drive': '',
	//'sf_rtfi_cleanup_state': '',
	//'sf_secure_erase': '1',
	//'sf_secure_erase_retries': '5',
	//'sf_slice_size': '',
	//'sf_ssh_key': '1',
	//'sf_ssh_root': '1',
	//'sf_start_rtfi': '1',
	//'sf_status_httpserver': '1',
	//'sf_status_httpserver_stop_delay': '5m',
	//'sf_status_inject_failure': '',
	//'sf_status_json': '0',
	//'sf_support_host': 'sfsupport.solidfire.com',
	//'sf_test_hardware': '0',
	//'sf_upgrade': '0',
	//'sf_upgrade_firmware': '0',
	//'sf_upload_logs_url': ''
	Options string `json:"options,omitempty"`
	//Blindly force the command, regardless of any sanity checks.
	//Specifically, this can be used to force the reset even if the node is part of a cluster.
	V__useTheForceLuke bool `json:"__useTheForceLuke,omitempty"`
}

type ListVirtualVolumesRequest struct {
	//Specifies the level of detail about each virtual volume that is returned. Possible values are:
	//true: Include more details about each virtual volume in the response.
	//false: Include the standard level of detail about each virtual volume in
	//the response.
	Details bool `json:"details,omitempty"`
	//The maximum number of virtual volumes to list.
	Limit int64 `json:"limit,omitempty"`
	//Specifies whether to include information about the children of each virtual volume in the response.
	//Possible values are:
	//true: Include information about the children of each virtual volume in
	//the response.
	//false: Do not include information about the children of each
	//virtual volume in the response.
	Recursive bool `json:"recursive,omitempty"`
	//The ID of the virtual volume at which to begin the list.
	StartVirtualVolumeID string `json:"startVirtualVolumeID,omitempty"`
	//A list of virtual volume IDs for which to retrieve information. If
	//you specify this parameter, the method returns information
	//about only these virtual volumes.
	VirtualVolumeIDs []string `json:"virtualVolumeIDs,omitempty"`
}

type GetAccountEfficiencyRequest struct {
	//Specifies the volume account for which efficiency
	//statistics are returned.
	AccountID int64 `json:"accountID,"`
}

type DisableProtectionSchemesRequest struct {
	//The protection schemes that should be disabled
	ProtectionSchemes []ProtectionScheme `json:"protectionSchemes,"`
}

type RemoveAccountRequest struct {
	//Specifies the AccountID for the account to be removed.
	AccountID int64 `json:"accountID,"`
}

type ListSnapMirrorLunsRequest struct {
	//List only the LUN information associated with the specified endpoint ID.
	SnapMirrorEndpointID int64 `json:"snapMirrorEndpointID,"`
	//The destination volume in the SnapMirror relationship.
	DestinationVolume SnapMirrorVolumeInfo `json:"destinationVolume,"`
}

type EnableFeatureRequest struct {
	//Indicates which feature to enable. Valid values are:
	//vvols: Enable the NetApp SolidFire VVols cluster feature.
	//FipsDrives: Enable the NetApp SolidFire cluster FIPS 140-2 drive support.
	Feature string `json:"feature,"`
}

type DeleteSnapMirrorEndpointsRequest struct {
	//An array of IDs of SnapMirror endpoints to delete.
	SnapMirrorEndpointIDs []int64 `json:"snapMirrorEndpointIDs,"`
}

type ComputeVolumeChecksumRequest struct {
	//Volume on which to compute checksum.
	VolumeID int `json:"volumeID,"`
	//Snapshot on which to compute checksum. If not provided, compute checksum on volume.
	SnapshotID int `json:"snapshotID,omitempty"`
	//Read all metadata from the slice drive and recompute the checksum.
	//
	//WARNING:  This paramter has significant limitations when executed on large volumes.
	//1. The RPC message timeout from the cluster master to the primary slice service is 30 seconds.
	//Receiving checksums will be impossible if it takes longer than this time to read the metadata.
	//Regardless of that timeout, the primary and secondary slice services will run to completion.
	//2. The slice service acquires a per-slice read lock when computing the checksums. This will block
	//any thread that attempts to acquire the write lock and could cause a HangDetector trip.
	//3. The checksum code aquires an exclusive lock. If this command is issued back-to-back then the
	//second thread will be blocked until the first finishes. This could easily cause a HangDetector
	//trip on the second thread.
	ForceReadFromDiskAndRecompute bool `json:"forceReadFromDiskAndRecompute,omitempty"`
}

type RestoreDeletedVolumeRequest struct {
	//VolumeID of the deleted volume to be restored.
	VolumeID int64 `json:"volumeID,"`
}

type ListVolumeStatsByAccountRequest struct {
	//One or more account ids by which to filter the result.
	Accounts []int64 `json:"accounts,omitempty"`
	//Includes virtual volumes in the response by default. To exclude virtual volumes, set to false.
	IncludeVirtualVolumes bool `json:"includeVirtualVolumes,omitempty"`
}

type GetAsyncResultRequest struct {
	//A value that was returned from the original
	//asynchronous method call.
	AsyncHandle int64 `json:"asyncHandle,"`
	//If true, GetAsyncResult does not remove the
	//asynchronous result upon returning it, enabling future
	//queries to that asyncHandle.
	KeepResult bool `json:"keepResult,omitempty"`
}

type EnableEncryptionAtRestRequest struct {
	//The ID of a Key Provider to use.  This is a unique value returned as part of one of the CreateKeyProvider* methods.
	KeyProviderID int64 `json:"keyProviderID,omitempty"`
	//If true, allow enable encryption at rest even when bin sync is in progress.
	ForceDuringBinSync bool `json:"forceDuringBinSync,omitempty"`
}

type SetDefaultQoSRequest struct {
	//The minimum number of sustained IOPS provided by the cluster to a
	//volume.
	MinIOPS int64 `json:"minIOPS,omitempty"`
	//The maximum number of sustained IOPS provided by the cluster to a
	//volume.
	MaxIOPS int64 `json:"maxIOPS,omitempty"`
	//The maximum number of IOPS allowed in a short burst scenario.
	BurstIOPS int64 `json:"burstIOPS,omitempty"`
}

type ListVolumesRequest struct {
	//Only volumes with an ID greater than or equal to this
	//value are returned. Mutually exclusive with the
	//volumeIDs parameter.
	StartVolumeID int64 `json:"startVolumeID,omitempty"`
	//Specifies the maximum number of volume
	//results that are returned. Mutually exclusive with the
	//volumeIDs parameter.
	Limit int64 `json:"limit,omitempty"`
	//Only volumes with a status equal to the status value are
	//returned.
	//Possible values are:
	//creating
	//snapshotting
	//active
	//deleted
	VolumeStatus string `json:"volumeStatus,omitempty"`
	//Returns only the volumes owned by the accounts you specify here. Mutually exclusive with the volumeIDs parameter.
	Accounts []int64 `json:"accounts,omitempty"`
	//Returns volumes that are paired or not paired.
	//Possible values are:
	//true: Returns all paired volumes.
	//false: Returns all volumes that are not paired.
	IsPaired bool `json:"isPaired,omitempty"`
	//A list of volume IDs. If you supply this parameter, other
	//parameters operate only on this set of volumes. Mutually
	//exclusive with the accounts, startVolumeID, and limit
	//parameters.
	VolumeIDs []int64 `json:"volumeIDs,omitempty"`
	//Only volume object information matching the volume
	//name is returned.
	VolumeName string `json:"volumeName,omitempty"`
	//Specifies that virtual volumes are included in the response by default.
	//To exclude virtual volumes, set to false.
	IncludeVirtualVolumes bool `json:"includeVirtualVolumes,omitempty"`
	//Only volumes that are using one of the protection schemes in this set are returned.
	ProtectionSchemes []ProtectionScheme `json:"protectionSchemes,omitempty"`
}

type RemoveClusterAdminRequest struct {
	//ClusterAdminID for the cluster admin to remove.
	ClusterAdminID int64 `json:"clusterAdminID,"`
}

type ResetDrivesRequest struct {
	//List of device names (not driveIDs) to reset.
	Drives string `json:"drives,"`
	//Required parameter to successfully reset a drive.
	Force bool `json:"force,"`
}

type StartVolumePairingRequest struct {
	//The ID of the volume on which to start the pairing process.
	VolumeID int64 `json:"volumeID,"`
	//The mode of the volume on which to start the pairing process. The mode can only be set if the volume is the source volume. Possible values are: Async: (default if no mode parameter specified) Writes are acknowledged when they complete locally. The cluster does not wait for writes to be replicated to the target cluster. Sync: Source acknowledges write when the data is stored locally and on the remote cluster. SnapshotsOnly: Only snapshots created on the source cluster will be replicated. Active writes from the source volume are not replicated.
	Mode string `json:"mode,omitempty"`
}

type ListSnapMirrorRelationshipsRequest struct {
	//List only the relationships associated
	//with the specified endpoint ID.
	//If no endpoint ID is provided, the system
	//lists relationships from all known SnapMirror endpoints.
	SnapMirrorEndpointID int64 `json:"snapMirrorEndpointID,omitempty"`
	//List relationships associated with the specified destination volume.
	DestinationVolume SnapMirrorVolumeInfo `json:"destinationVolume,omitempty"`
	//List relationships associated with the specified source volume.
	SourceVolume SnapMirrorVolumeInfo `json:"sourceVolume,omitempty"`
	//List relationships on the specified Vserver.
	Vserver string `json:"vserver,omitempty"`
	//List relationships associated with the specified relationshipID.
	RelationshipID string `json:"relationshipID,omitempty"`
}

type DeleteVolumesRequest struct {
	//A list of account IDs. All volumes from these accounts are deleted from the system.
	AccountIDs []int64 `json:"accountIDs,omitempty"`
	//A list of volume access group IDs. All of the volumes from all of the volume access groups you specify in this list are deleted from the system.
	VolumeAccessGroupIDs []int64 `json:"volumeAccessGroupIDs,omitempty"`
	//The list of IDs of the volumes to delete from the system.
	VolumeIDs []int64 `json:"volumeIDs,omitempty"`
}

type RemoveVirtualNetworkRequest struct {
	//Network ID that identifies the virtual network to remove.
	VirtualNetworkID int64 `json:"virtualNetworkID,omitempty"`
	//Network tag that identifies the virtual network to remove.
	VirtualNetworkTag int64 `json:"virtualNetworkTag,omitempty"`
}

type DeleteStandbySlicesRequest struct {
	//Slices on which to delete standbys.
	SliceIDs []int64 `json:"sliceIDs,"`
}

type ListTestsRequest struct {
	IncludeHints bool `json:"includeHints,"`
	Force        bool `json:"force,"`
}

type RollbackToGroupSnapshotRequest struct {
	//Specifies the unique ID of the group snapshot.
	GroupSnapshotID int64 `json:"groupSnapshotID,"`
	//Specifies whether to save an active volume image or delete it. Values are:
	//true: The previous active volume image is kept.
	//false: (default) The previous active volume image is deleted.
	SaveCurrentState bool `json:"saveCurrentState,"`
	//Name for the group snapshot of the volume's
	//current state that is created if "saveCurrentState" is
	//set to true. If you do not give a name, the
	//name of the snapshots (group and individual
	//volume) are set to a timestamp of the time that
	//the rollback occurred.
	Name string `json:"name,omitempty"`
	//List of name-value pairs in JSON object format.
	Attributes interface{} `json:"attributes,omitempty"`
}

type BreakSnapMirrorRelationshipRequest struct {
	//The endpoint ID of the remote ONTAP storage system communicating with the SolidFire cluster.
	SnapMirrorEndpointID int64 `json:"snapMirrorEndpointID,"`
	//The destination volume in the SnapMirror relationship.
	DestinationVolume SnapMirrorVolumeInfo `json:"destinationVolume,"`
}

type AddKeyServerToProviderKmipRequest struct {
	//The ID of the Key Provider to assign the KMIP Key Server to.
	KeyProviderID int64 `json:"keyProviderID,"`
	//The ID of the KMIP Key Server to assign.
	KeyServerID int64 `json:"keyServerID,"`
}

type ListSnapMirrorNetworkInterfacesRequest struct {
	//Return only the network interfaces associated with the specified endpoint ID.
	//If no endpoint ID is provided, the system lists interfaces
	//from all known SnapMirror endpoints.
	SnapMirrorEndpointID int64 `json:"snapMirrorEndpointID,omitempty"`
	//List only the network interface serving the specified role.
	InterfaceRole string `json:"interfaceRole,omitempty"`
}

type ListVolumeAccessGroupsRequest struct {
	//The volume access group ID at which to begin the listing. If unspecified, there is no lower limit (implicitly 0).
	StartVolumeAccessGroupID int64 `json:"startVolumeAccessGroupID,omitempty"`
	//The maximum number of results to return. This can be
	//useful for paging.
	Limit int64 `json:"limit,omitempty"`
	//The list of ids of the volume access groups you wish to list
	VolumeAccessGroups []int64 `json:"volumeAccessGroups,omitempty"`
}

type DeleteAuthSessionsByClusterAdminRequest struct {
	//ID that identifies a clusterAdmin.
	ClusterAdminID int64 `json:"clusterAdminID,"`
}

type UpdateSnapMirrorRelationshipRequest struct {
	//The endpoint ID of the remote ONTAP storage system communicating with the SolidFire cluster.
	SnapMirrorEndpointID int64 `json:"snapMirrorEndpointID,"`
	//The destination volume in the SnapMirror relationship.
	DestinationVolume SnapMirrorVolumeInfo `json:"destinationVolume,"`
	//Specifies the maximum data transfer rate
	//between the volumes in kilobytes per second.
	//The default value, 0, is unlimited and permits the SnapMirror
	//relationship to fully utilize the available network bandwidth.
	MaxTransferRate int64 `json:"maxTransferRate,omitempty"`
}

type GetVolumeEfficiencyRequest struct {
	//Specifies the volume for which capacity is computed.
	VolumeID int64 `json:"volumeID,"`
}

type PurgeDeletedVolumesRequest struct {
	//A list of volumeIDs of volumes to be purged from the system.
	VolumeIDs []int64 `json:"volumeIDs,omitempty"`
	//A list of accountIDs. All of the volumes from all of the specified accounts are purged from the system.
	AccountIDs []int64 `json:"accountIDs,omitempty"`
	//A list of volumeAccessGroupIDs. All of the volumes from all of the specified Volume Access Groups are purged from the system.
	VolumeAccessGroupIDs []int64 `json:"volumeAccessGroupIDs,omitempty"`
}

type GetConstantsRequest struct {
	//Only return the values for the list of provided constants.
	//If not provided, the value of all constants will be returned.
	Constants []string `json:"constants,omitempty"`
	//If true, only return the set of constants that have changed from their default value.
	ChangedOnly bool `json:"changedOnly,omitempty"`
}

type BreakSnapMirrorVolumeRequest struct {
	//The volume on which to perform the break operation.
	//The volume access mode must be snapMirrorTarget.
	VolumeID int64 `json:"volumeID,"`
	//Roll back the volume to the snapshot identified by this ID.
	//The default behavior is to roll back to the most recent snapshot.
	SnapshotID int64 `json:"snapshotID,omitempty"`
	//Preserve any snapshots newer than the snapshot identified by snapshotID.
	//Possible values:
	//true: Preserve snapshots newer than snapshotID.
	//false: Do not preserve snapshots newer than snapshotID.
	//If false, any snapshots newer than snapshotID are deleted.
	Preserve bool `json:"preserve,omitempty"`
	//Resulting volume access mode.
	//Possible values:
	//readWrite
	//readOnly
	//locked
	Access string `json:"access,omitempty"`
}

type RemoveInitiatorsFromVolumeAccessGroupRequest struct {
	//The ID of the volume access group
	//from which the initiators are removed.
	VolumeAccessGroupID int64 `json:"volumeAccessGroupID,"`
	//The list of initiators to remove from the volume
	//access group.
	Initiators []string `json:"initiators,"`
	//true: Default. Delete initiator objects after they are removed from a volume access group.
	//false: Do not delete initiator objects after they are removed from a volume access group.
	DeleteOrphanInitiators bool `json:"deleteOrphanInitiators,omitempty"`
}

type CreateMultipleVolumesRequest struct {
	//The name of the volume access group (might be user specified).
	//Not required to be unique, but recommended.
	//Might be 1 to 64 characters in length.
	Names []string `json:"names,"`
	//AccountID for the owner of the volumes.
	AccountID int64 `json:"accountID,"`
	//Total size of each volume, in bytes. Size is rounded up to
	//the nearest 1MB size.
	TotalSize int64 `json:"totalSize,"`
	//Specifies whether 512e emulation is enabled or not. Possible values are:
	//true: The volumes provides 512-byte sector emulation.
	//false: 512e emulation is not enabled.
	Enable512e bool `json:"enable512e,"`
	//Initial quality of service settings for this volumes. Default
	//values are used if none are specified. Valid settings are:
	//minIOPS
	//maxIOPS
	//burstIOPS
	//You can get the default values for a volumes by using the GetDefaultQoS method.
	Qos QoS `json:"qos,omitempty"`
	//The list of name-value pairs in JSON object format.
	//Total attribute size must be less than 1000B, or 1KB,
	//including JSON formatting characters.
	Attributes interface{} `json:"attributes,omitempty"`
	//Associate the volumes with the specified QoS policy.
	//Possible values:
	//true: Associate the volumes with the QoS policy specified in the QoSPolicyID parameter.
	//false: Do not assosciate the volumes with the QoS policy specified in the QoSPolicyID parameter. When false, any existing policy association is removed regardless of whether you specify a QoS policy in the QoSPolicyID parameter.
	AssociateWithQoSPolicy bool `json:"associateWithQoSPolicy,omitempty"`
	//The ID for the policy whose QoS settings should be applied to the volumes.
	//This parameter is mutually exclusive with the qos parameter.
	QosPolicyID int64 `json:"qosPolicyID,omitempty"`
	//Access settings for the new volumes.
	Access string `json:"access,omitempty"`
	//volumeIDs that should be used for the volumes
	VolumeIDs bool `json:"volumeIDs,omitempty"`
	//Allow volume creation even if the cluster fullness check fails
	V__skipFullnessCheck bool `json:"__skipFullnessCheck,omitempty"`
	//Determines whether the volumes can be used for replication with SnapMirror endpoints.
	//Possible values:
	//true
	//false
	EnableSnapMirrorReplication bool `json:"enableSnapMirrorReplication,omitempty"`
	//Protection scheme that should be used for the volumes.
	//The default value is the defaultProtectionScheme stored in the ClusterInfo object.
	ProtectionScheme ProtectionScheme `json:"protectionScheme,omitempty"`
}

type ListClusterFaultsRequest struct {
	//Shows more detail about the exception that caused the cluster fault.
	Exceptions bool `json:"exceptions,omitempty"`
	//Specifies whether to include faults triggered by suboptimal system configuration.
	//Possible values are:
	//true
	//false
	BestPractices bool `json:"bestPractices,omitempty"`
	//Refreshes the latest fault information as the method call is made.
	Update bool `json:"update,omitempty"`
	//Determines the types of faults returned. Possible values are:
	//current: List active, unresolved faults.
	//resolved: List faults that were previously detected and
	//resolved.
	//all: (Default) List both current and resolved faults. You can
	//see the fault status in the resolved field of the Cluster Fault
	//object.
	FaultTypes string `json:"faultTypes,omitempty"`
}

type SetLldpConfigRequest struct {
	//Enable/disable LLDP entirely
	EnableLldp bool `json:"enableLldp,omitempty"`
	//Enable or disable MED, an extension to LLDP that provides inventory information for NetApp HCI products
	EnableMed bool `json:"enableMed,omitempty"`
	//Enable other discovery protocols: CDP, FDP, EDP, and SONMP.
	EnableOtherProtocols bool `json:"enableOtherProtocols,omitempty"`
}

type CheckProposedClusterRequest struct {
	//List of node IPs for the nodes in the new cluster.
	Nodes []string `json:"nodes,"`
}

type RemoveClusterPairRequest struct {
	//Unique identifier used to pair two clusters.
	ClusterPairID int64 `json:"clusterPairID,"`
}

type RemoveNodesRequest struct {
	//List of NodeIDs for the nodes to be removed.
	Nodes []int64 `json:"nodes,"`
	//Ignore changes to the ensemble's node failure tolerance when removing nodes.
	IgnoreEnsembleToleranceChange bool `json:"ignoreEnsembleToleranceChange,omitempty"`
}

type ResyncSnapMirrorRelationshipRequest struct {
	//The endpoint ID of the remote ONTAP storage system communicating with the SolidFire cluster.
	SnapMirrorEndpointID int64 `json:"snapMirrorEndpointID,"`
	//The destination volume in the SnapMirror relationship.
	DestinationVolume SnapMirrorVolumeInfo `json:"destinationVolume,"`
	//Specifies the maximum data transfer rate
	//between the volumes in kilobytes per second.
	//The default value, 0, is unlimited and permits the SnapMirror
	//relationship to fully utilize the available network bandwidth.
	MaxTransferRate int64 `json:"maxTransferRate,omitempty"`
	//The source volume in the SnapMirror relationship.
	SourceVolume SnapMirrorVolumeInfo `json:"sourceVolume,omitempty"`
}

type ListStorageContainersRequest struct {
	//A list of storage container IDs for which to retrieve information. If you omit this parameter, the method returns information about all storage containers in the system.
	StorageContainerIDs []string `json:"storageContainerIDs,omitempty"`
}

type DeleteVolumeRequest struct {
	//The ID of the volume to be deleted.
	VolumeID int64 `json:"volumeID,"`
}

type GetVolumeAccessGroupEfficiencyRequest struct {
	//The volume access group for which
	//capacity is computed.
	VolumeAccessGroupID int64 `json:"volumeAccessGroupID,"`
}

type SetProtectionDomainLayoutChassisOverrideRequest struct {
	//The Protection Domains information for all Nodes in the Cluster.
	ProtectionDomainLayout ProtectionDomainLayout `json:"protectionDomainLayout,"`
	//Allows setting the Chassis Protection Domains in a cluster with a witness node.
	ForceForWitnessNode bool `json:"forceForWitnessNode,omitempty"`
}

type GetClusterStateRequest struct {
	//To run this command, the force parameter must be set to true.
	Force        bool `json:"force,"`
	IncludeHints bool `json:"includeHints,omitempty"`
}

type ModifyGroupSnapshotRequest struct {
	//Specifies the ID of the group of snapshots.
	GroupSnapshotID int64 `json:"groupSnapshotID,"`
	//Sets the time when the snapshot should be
	//removed. If unspecified, the current time is used.
	ExpirationTime string `json:"expirationTime,omitempty"`
	//Replicates the snapshot created to a remote cluster.
	//Possible values are:
	//true: The snapshot is replicated to remote storage.
	//false: Default. The snapshot is not replicated.
	EnableRemoteReplication bool `json:"enableRemoteReplication,omitempty"`
	//Label used by SnapMirror software to specify snapshot retention policy on SnapMirror endpoint.
	SnapMirrorLabel string `json:"snapMirrorLabel,omitempty"`
}

type PromoteClusterMasterRequest struct {
	//Node ID to promote to cluster master.
	NodeID int64 `json:"nodeID,omitempty"`
	//Master service ID to promote to cluster master.
	ServiceID int64 `json:"serviceID,omitempty"`
}

type UpdateIdpConfigurationRequest struct {
	//UUID for the third party Identity Provider (IdP) Configuration.
	IdpConfigurationID string `json:"idpConfigurationID,"`
	//URL for retrieving IdP Metadata for configuration and integration details for SAML 2.0 single sign-on.
	//One must disable the IdP Configuration before the URL can be updated.
	IdpMetadataUrl string `json:"idpMetadataUrl,omitempty"`
	//If not provided or false, the expiration will remain unchanged.
	//If true, update the certificate used for IdP communication to the system defined max expiration time.
	RefreshCertificateExpirationTime bool `json:"refreshCertificateExpirationTime,omitempty"`
}

type GetQoSPolicyRequest struct {
	//The ID of the policy to be retrieved.
	QosPolicyID int64 `json:"qosPolicyID,"`
}

type CreatePublicPrivateKeyPairRequest struct {
	//This is the X.509 distinguished name Common Name field (CN).
	CommonName string `json:"commonName,omitempty"`
	//This is the X.509 distinguished name Organization Name field (O).
	Organization string `json:"organization,omitempty"`
	//This is the X.509 distinguished name Organizational Unit Name field (OU).
	OrganizationalUnit string `json:"organizationalUnit,omitempty"`
	//This is the X.509 distinguished name Locality Name field (L).
	Locality string `json:"locality,omitempty"`
	//This is the X.509 distinguished name State or Province Name field (ST or SP or S).
	State string `json:"state,omitempty"`
	//This is the X.509 distinguished name Country field (C).
	Country string `json:"country,omitempty"`
	//This is the X.509 distinguished name Email Address field (MAIL).
	EmailAddress string `json:"emailAddress,omitempty"`
}

type TestConnectMvipRequest struct {
	//If specified, tests the management connection of a different MVIP. You do not need to use this value when testing the connection to the target cluster. This parameter is optional.
	Mvip string `json:"mvip,omitempty"`
	//Force is required to run this on a cluster connection.
	Force bool `json:"force,omitempty"`
}

type ListKeyServersKmipRequest struct {
	//If omitted, returned KMIP Key Server objects will not be filtered based on whether they're assigned to the specified KMIP Key Provider.
	//If specified, returned KMIP Key Server objects will be filtered to those assigned to the specified KMIP Key Provider.
	KeyProviderID int64 `json:"keyProviderID,omitempty"`
	//If omitted, returned KMIP Key Server objects will not be filtered based on whether they're active.
	//If true, returns only KMIP Key Server objects which are active (providing keys which are currently in use).
	//If false, returns only KMIP Key Server objects which are inactive (not providing any keys and able to be deleted).
	KmipAssignedProviderIsActive bool `json:"kmipAssignedProviderIsActive,omitempty"`
	//If omitted, returned KMIP Key Server objects will not be filtered based on whether they have a KMIP Key Provider assigned.
	//If true, returns only KMIP Key Server objects which have a KMIP Key Provider assigned.
	//If false, returns only KMIP Key Server objects which do not have a KMIP Key Provider assigned.
	KmipHasProviderAssigned bool `json:"kmipHasProviderAssigned,omitempty"`
}

type QuiesceSnapMirrorRelationshipRequest struct {
	//The endpoint ID of the remote ONTAP storage system communicating with the SolidFire cluster.
	SnapMirrorEndpointID int64 `json:"snapMirrorEndpointID,"`
	//The destination volume in the SnapMirror relationship.
	DestinationVolume SnapMirrorVolumeInfo `json:"destinationVolume,"`
}

type ListGroupSnapshotsRequest struct {
	//An array of unique volume IDs to query. If you do not
	//specify this parameter, all group snapshots on the cluster
	//are included.
	Volumes []int64 `json:"volumes,omitempty"`
	//Retrieves information for a specific group snapshot ID.
	GroupSnapshotID int64 `json:"groupSnapshotID,omitempty"`
}

type SetSnmpInfoRequest struct {
	//List of networks and what type of access they have to the
	//SNMP servers running on the cluster nodes. See the SNMP
	//Network Object for possible "networks" values. This parameter is required only for SNMP v2.
	Networks []SnmpNetwork `json:"networks,omitempty"`
	//If set to true, SNMP is enabled on each node in the cluster.
	Enabled bool `json:"enabled,omitempty"`
	//If set to true, SNMP v3 is enabled on each node in the cluster.
	SnmpV3Enabled bool `json:"snmpV3Enabled,omitempty"`
	//If SNMP v3 is enabled, this value must be passed in place of the networks parameter. This parameter is required only for SNMP v3.
	UsmUsers []SnmpV3UsmUser `json:"usmUsers,omitempty"`
}

type CreateClusterRequest struct {
	//Required to indicate your acceptance of the End User License
	//Agreement when creating this cluster. To accept the EULA,
	//set this parameter to true.
	AcceptEula bool `json:"acceptEula,omitempty"`
	//Required to allow creation of a 2 node ensemble using 2 storage nodes.
	//To allow a 2 node ensemble, set this parameter to true.
	Allow2NodeEnsemble bool `json:"allow2NodeEnsemble,omitempty"`
	//Floating (virtual) IP address for the cluster on the management network.
	Mvip string `json:"mvip,"`
	//Floating (virtual) IP address for the cluster on the storage (iSCSI) network.
	Svip string `json:"svip,"`
	//The set of testOnly schemes that should be added to the cluster's set of supported protection schemes.
	//By default, no testOnly schemes are supported.
	//If any of the schemes are unrecognized, xUnknownEnumType is returned.
	SupportedTestOnlyProtectionSchemes []ProtectionScheme `json:"supportedTestOnlyProtectionSchemes,omitempty"`
	//Username for the cluster admin.
	Username string `json:"username,"`
	//Initial password for the cluster admin account.
	Password string `json:"password,"`
	//CIP/SIP addresses of the initial set of nodes making up the cluster. This node's IP must be in the list.
	Nodes []string `json:"nodes,"`
	//List of name-value pairs in JSON object format.
	Attributes interface{} `json:"attributes,omitempty"`
	//Enable this flag to skip node type checks for a two node cluster.
	SkipNodeTypeChecks bool `json:"skipNodeTypeChecks,omitempty"`
}

type GetClusterConfigRequest struct {
	IncludeHints bool `json:"includeHints,omitempty"`
	Force        bool `json:"force,omitempty"`
}

type DeleteQoSPolicyRequest struct {
	//The ID of the QoS policy to be deleted.
	QosPolicyID int64 `json:"qosPolicyID,"`
}

type CreateGroupSnapshotRequest struct {
	//Unique ID of the volume image from which to copy.
	Volumes []int64 `json:"volumes,"`
	//Name for the group snapshot. If unspecified, the date and time the group snapshot was taken is used.
	Name string `json:"name,omitempty"`
	//Replicates the snapshot created to remote storage.
	//Possible values are:
	//true: The snapshot is replicated to remote storage.
	//false: Default. The snapshot is not replicated.
	EnableRemoteReplication bool `json:"enableRemoteReplication,omitempty"`
	//Specifies the amount of time for which the snapshots are retained. The format is HH:mm:ss.
	Retention string `json:"retention,omitempty"`
	//List of name-value pairs in JSON object format.
	Attributes interface{} `json:"attributes,omitempty"`
	//Label used by SnapMirror software to specify snapshot retention policy on SnapMirror endpoint.
	SnapMirrorLabel string `json:"snapMirrorLabel,omitempty"`
}

type GetVolumeAccessGroupLunAssignmentsRequest struct {
	//The unique volume access group ID used to return information.
	VolumeAccessGroupID int64 `json:"volumeAccessGroupID,"`
}

type AddBlockToBSRequest struct {
	//Block data to store.
	//Must be a string of hex characters that is between 10 and 4096 bytes.
	HashData string `json:"hashData,"`
	//The blockID for the data.
	//Must be the correct blockID for the data, unless validateData=false.
	Hash string `json:"hash,"`
	//Whether to validate that the blockID matches the data.
	//Default is true.
	ValidateData bool `json:"validateData,omitempty"`
	//The protection schemes for which the block needs to be added.
	//The default value is the list of enabledProtectionSchemes.
	ProtectionSchemes []ProtectionScheme `json:"protectionSchemes,omitempty"`
}

type SetRemoteLoggingHostsRequest struct {
	//A list of hosts to send log messages to.
	RemoteHosts []LoggingServer `json:"remoteHosts,"`
}

type ListAuthSessionsByClusterAdminRequest struct {
	//ID that identifies a clusterAdmin.
	ClusterAdminID int64 `json:"clusterAdminID,"`
}

type SetAuthConfigurationRequest struct {
	//Indicates the configuration data that will be set.
	AuthType AuthConfigType `json:"authType,"`
	//Version of the data currently stored in the database.  If there is no data
	//currently set, then this should be zero.
	Version int64 `json:"version,"`
	//A string which will be stored in the cluster database up to 20KB in size.
	Configuration string `json:"configuration,"`
}

type CreateInitiatorsRequest struct {
	//A list of objects containing characteristics of each new initiator.
	Initiators []CreateInitiator `json:"initiators,"`
}

type CheckVolumeStandbysAssignedRequest struct {
	//The node to check.
	NodeID int64 `json:"nodeID,"`
}

type GetDriveConfigRequest struct {
	Force bool `json:"force,omitempty"`
}

type SetSSLCertificateRequest struct {
	//The PEM-encoded text version of the certificate.
	Certificate string `json:"certificate,"`
	//The PEM-encoded text version of the private key.
	PrivateKey string `json:"privateKey,"`
}

type CreateIdpConfigurationRequest struct {
	//URL for retrieving IdP Metadata for configuration and integration details for SAML 2.0 single sign-on.
	IdpMetadataUrl string `json:"idpMetadataUrl,"`
}

type ListVirtualVolumeBindingsRequest struct {
	//A list of virtual volume binding IDs for which to retrieve information. If you omit this parameter, the method returns information about all virtual volume bindings.
	VirtualVolumeBindingIDs []int64 `json:"virtualVolumeBindingIDs,omitempty"`
}

type CompleteClusterPairingRequest struct {
	//A string of characters that is returned from the "StartClusterPairing" API method.
	ClusterPairingKey string `json:"clusterPairingKey,"`
}

type CreateQoSPolicyRequest struct {
	//The name of the QoS policy; for example, gold, platinum, or silver.
	Name string `json:"name,"`
	//The QoS settings that this policy represents.
	Qos QoS `json:"qos,"`
}

type ModifyVolumesRequest struct {
	//A list of volumeIDs for the volumes to be modified.
	VolumeIDs []int64 `json:"volumeIDs,"`
	//AccountID to which the volume is reassigned. If none is specified, the previous account name is used.
	AccountID int64 `json:"accountID,omitempty"`
	//Access allowed for the volume. Possible values:readOnly: Only read operations are allowed.readWrite: Reads and writes are allowed.locked: No reads or writes are allowed.If not specified, the access value does not change.replicationTarget: Identify a volume as the target volume for a paired set of volumes. If the volume is not paired, the access status is locked.If a value is not specified, the access value does not change.
	Access string `json:"access,omitempty"`
	//New quality of service settings for this volume.If not specified, the QoS settings are not changed.
	Qos QoS `json:"qos,omitempty"`
	//New size of the volume in bytes. 1000000000 is equal to 1GB. Size is rounded up to the nearest 1MB in size. This parameter can only be used to increase the size of a volume.
	TotalSize int64 `json:"totalSize,omitempty"`
	//Associate the volume with the specified QoS policy.
	//Possible values:
	//true: Associate the volume with the QoS policy specified in the QoSPolicyID parameter.
	//false: Do not assosciate the volume with the QoS policy specified in the QoSPolicyID parameter. When false, any existing policy association is removed regardless of whether you specify a QoS policy in the QoSPolicyID parameter.
	AssociateWithQoSPolicy bool `json:"associateWithQoSPolicy,omitempty"`
	//The ID for the policy whose QoS settings should be applied to the specified volumes.
	//This parameter is mutually exclusive with the qos parameter.
	QosPolicyID int64 `json:"qosPolicyID,omitempty"`
	//List of name/value pairs in JSON object format.
	Attributes interface{} `json:"attributes,omitempty"`
	//Determines whether the volume can be used for replication with SnapMirror endpoints.
	//Possible values:
	//true
	//false
	EnableSnapMirrorReplication bool `json:"enableSnapMirrorReplication,omitempty"`
}

type RemoveDrivesRequest struct {
	//List of driveIDs to remove from the cluster.
	Drives []int64 `json:"drives,"`
	//If you want to remove a drive during upgrade, this must be set to true.
	ForceDuringUpgrade bool `json:"forceDuringUpgrade,omitempty"`
	//Allows the user to force the removal of drives if it would result in an unbalanced cluster.
	IgnoreUnbalancedClusterCheck bool `json:"ignoreUnbalancedClusterCheck,omitempty"`
}

type GetHardwareConfigRequest struct {
	Force bool `json:"force,omitempty"`
}

type ListVolumeStatsRequest struct {
	//A list of volume IDs of volumes from which to retrieve activity information.
	VolumeIDs []int64 `json:"volumeIDs,omitempty"`
}

type GetSnapMirrorClusterIdentityRequest struct {
	//If provided, the system lists the cluster identity of the endpoint with the specified snapMirrorEndpointID.
	//If not provided, the system lists the cluster identity of all known SnapMirror endpoints.
	SnapMirrorEndpointID int64 `json:"snapMirrorEndpointID,omitempty"`
}

type DeleteIdpConfigurationRequest struct {
	//UUID for the third party Identity Provider (IdP) Configuration.
	IdpConfigurationID string `json:"idpConfigurationID,"`
}

type ListAccountsRequest struct {
	//Starting AccountID to return. If no account exists with this
	//AccountID, the next account by AccountID order is used as
	//the start of the list. To page through the list, pass the
	//AccountID of the last account in the previous response +
	//1.
	StartAccountID int64 `json:"startAccountID,omitempty"`
	//Maximum number of AccountInfo objects to return.
	Limit int64 `json:"limit,omitempty"`
	//Includes storage containers in the response by
	//default. To exclude storage containers, set to false.
	IncludeStorageContainers bool `json:"includeStorageContainers,omitempty"`
}

type GetBackupTargetRequest struct {
	//The unique identifier assigned to the backup target.
	BackupTargetID int64 `json:"backupTargetID,"`
}

type DemoteClusterMasterRequest struct {
	//Node ID to demote. If not provided the cluster will demote the current cluster master.
	//If the provided node ID is not cluster master then it will be moved to the end of the
	//list of possible cluster masters.
	NodeID int64 `json:"nodeID,omitempty"`
	//Master service ID to demote. If not provided the cluster will demote whichever service is cluster
	//master. If the provided service ID is not cluster master then it will be moved to the end of the
	//list of possible cluster masters.
	ServiceID int64 `json:"serviceID,omitempty"`
}

type ShutdownRequest struct {
	//List of NodeIDs for the nodes to be shutdown.
	Nodes []int64 `json:"nodes,"`
	//Specifies the action to take for the node shutdown. Possible values are:
	//restart: Restarts the node.
	//halt: Shuts down the node.
	Option string `json:"option,omitempty"`
}

type TestLdapAuthenticationRequest struct {
	//The username to be tested.
	Username string `json:"username,"`
	//The password for the username to be tested.
	Password string `json:"password,"`
	//An ldapConfiguration object to be tested. If specified, the API call tests the provided
	//configuration even if LDAP authentication is disabled.
	LdapConfiguration LdapConfiguration `json:"ldapConfiguration,omitempty"`
}

type ListSnapMirrorAggregatesRequest struct {
	//Return only the aggregates associated with the specified endpoint ID.
	//If no endpoint ID is provided, the system lists aggregates
	//from all known SnapMirror endpoints.
	SnapMirrorEndpointID int64 `json:"snapMirrorEndpointID,omitempty"`
}

type EnableIdpAuthenticationRequest struct {
	//UUID for the third party Identity Provider (IdP) Configuration.
	//If only one IdP Configuration exists, then we will default to enabling that configuration.
	IdpConfigurationID string `json:"idpConfigurationID,omitempty"`
}

type CreateKeyServerKmipRequest struct {
	//The public key certificate of the external key server's root CA.
	//This will be used to verify the certificate presented by external key server in the TLS communication.
	//For key server clusters where individual servers use different CAs, provide a concatenated string containing the root certificates of all the CAs.
	KmipCaCertificate string `json:"kmipCaCertificate,"`
	//A PEM format Base64 encoded PKCS#10 X.509 certificate used by the Solidfire KMIP client.
	KmipClientCertificate string `json:"kmipClientCertificate,"`
	//Array of the hostnames or IP addresses associated with this KMIP Key Server. Multiple hostnames or IP addresses
	//must only be provided if the key servers are in a clustered configuration.
	KmipKeyServerHostnames []string `json:"kmipKeyServerHostnames,"`
	//The name of the KMIP Key Server.  This name is only used for display purposes and does not need to be unique.
	KmipKeyServerName string `json:"kmipKeyServerName,"`
	//The port number associated with this KMIP Key Server (typically 5696).
	KmipKeyServerPort int64 `json:"kmipKeyServerPort,omitempty"`
}

type ModifyKeyServerKmipRequest struct {
	//The public key certificate of the external key server's root CA.
	//This will be used to verify the certificate presented by external key server in the TLS communication.
	//For key server clusters where individual servers use different CAs, provide a concatenated string containing the root certificates of all the CAs.
	KmipCaCertificate string `json:"kmipCaCertificate,omitempty"`
	//A PEM format Base64 encoded PKCS#10 X.509 certificate used by the Solidfire KMIP client.
	KmipClientCertificate string `json:"kmipClientCertificate,omitempty"`
	//Array of the hostnames or IP addresses associated with this KMIP Key Server. Multiple hostnames or IP addresses
	//must only be provided if the key servers are in a clustered configuration.
	KmipKeyServerHostnames []string `json:"kmipKeyServerHostnames,omitempty"`
	//The ID of the KMIP Key Server to modify.
	KeyServerID int64 `json:"keyServerID,"`
	//The name of the KMIP Key Server.  This name is only used for display purposes and does not need to be unique.
	KmipKeyServerName string `json:"kmipKeyServerName,omitempty"`
	//The port number associated with this KMIP Key Server (typically 5696).
	KmipKeyServerPort int64 `json:"kmipKeyServerPort,omitempty"`
}

type AddVolumesToVolumeAccessGroupRequest struct {
	//The ID of the volume access group to which volumes are added.
	VolumeAccessGroupID int64 `json:"volumeAccessGroupID,"`
	//The list of volumes to add to the volume access
	//group.
	Volumes []int64 `json:"volumes,"`
}

type AddIdpClusterAdminRequest struct {
	//A SAML attribute-value mapping to a IdP cluster admin (e.g. email=test@example.com).
	//This could be defined using a specific SAML subject using NameID, or an entry in the
	//SAML attribute statement such as eduPersonAffiliation.
	Username string `json:"username,"`
	//Controls which methods this IdP Cluster Admin can use. For more details on the levels of access,
	//see the Access Control appendix in the SolidFire API Reference.
	Access []string `json:"access,"`
	//Accept the End User License Agreement. Set to true to add a cluster administrator account to the system.
	//If omitted or set to false, the method call fails.
	AcceptEula bool `json:"acceptEula,omitempty"`
	//List of name-value pairs in JSON object format.
	Attributes interface{} `json:"attributes,omitempty"`
}

type AddDrivesRequest struct {
	//Returns information about each drive to be added to the
	//cluster. Possible values are:
	//driveID: The ID of the drive to
	//add. (Integer)
	//type: (Optional) The type of drive
	//to add. Valid values are "slice" or "block". If
	//omitted, the system assigns the correct
	//type. (String)
	Drives []NewDrive `json:"drives,"`
	//Allows the user to force the addition of drives during an upgrade.
	ForceDuringUpgrade bool `json:"forceDuringUpgrade,omitempty"`
	//Allows addition of non-Fips drives to cluster while FipsDrives feature is enabled.
	ForceNonFips bool `json:"forceNonFips,omitempty"`
	//Allows the user to force the addition of drives during a bin sync operation.
	ForceDuringBinSync bool `json:"forceDuringBinSync,omitempty"`
	//Allows the user to force the addition of drives if it would result in an unbalanced cluster.
	IgnoreUnbalancedClusterCheck bool `json:"ignoreUnbalancedClusterCheck,omitempty"`
	//Allows the user to force the addition of block drives without validating block drive usable capacity.
	IgnoreUsableCapacityCheck bool `json:"ignoreUsableCapacityCheck,omitempty"`
	//Allows the user to force the addition of drives into unknown custom Protection Domains.
	IgnoreProtectionDomainLayoutCheck bool `json:"ignoreProtectionDomainLayoutCheck,omitempty"`
}

type ModifySnapshotRequest struct {
	//Specifies the ID of the snapshot.
	SnapshotID int64 `json:"snapshotID,"`
	//Sets the time when the snapshot should be
	//removed.
	ExpirationTime string `json:"expirationTime,omitempty"`
	//Replicates the snapshot created to a remote cluster.
	//Possible values are:
	//true: The snapshot is replicated to remote storage.
	//false: Default. The snapshot is not replicated.
	EnableRemoteReplication bool `json:"enableRemoteReplication,omitempty"`
	//Label used by SnapMirror software to specify snapshot retention policy on SnapMirror endpoint.
	SnapMirrorLabel string `json:"snapMirrorLabel,omitempty"`
}

type ModifyVirtualNetworkRequest struct {
	//The unique identifier of the virtual network to modify. This is the virtual
	//network ID assigned by the cluster.
	//Note: This parameter is optional
	//but either virtualNetworkID or virtualNetworkTag must be specified
	//with this API method.
	VirtualNetworkID int64 `json:"virtualNetworkID,omitempty"`
	//The network tag that identifies the virtual network to modify.
	//Note: This parameter is optional
	//but either virtualNetworkID or virtualNetworkTag must be specified
	//with this API method.
	VirtualNetworkTag int64 `json:"virtualNetworkTag,omitempty"`
	//The new name for the virtual network.
	Name string `json:"name,omitempty"`
	//The new addressBlock to set for this virtual network. This might contain
	//new address blocks to add to the existing object or omit
	//unused address blocks that need to be removed. Alternatively, you
	//can extend or reduce the size of existing address blocks. You can only
	//increase the size of the starting addressBlocks for a virtual network
	//object; you can never decrease it.
	//Attributes for this parameter are:
	//start: The start of the IP address range. (String)
	//size: The number of IP addresses to include in the block. (Integer)
	AddressBlocks []AddressBlockParams `json:"addressBlocks,omitempty"`
	//New network mask for this virtual network.
	Netmask string `json:"netmask,omitempty"`
	//The storage virtual IP address for this virtual network. The svip for a
	//virtual network cannot be changed. You must create a new virtual
	//network to use a different svip address.
	Svip string `json:"svip,omitempty"`
	//The IP address of a gateway of the virtual network. This parameter is valid only if the namespace parameter is set to true (meaning VRF is enabled).
	Gateway string `json:"gateway,omitempty"`
	//When set to true, enables the Routable Storage VLANs functionality by recreating the virtual network and configuring a namespace to contain it.
	//When set to false, disables the VRF functionality for the virtual network.
	//Changing this value disrupts traffic running through this virtual network.
	Namespace bool `json:"namespace,omitempty"`
	//A new list of name-value pairs in JSON object format.
	Attributes interface{} `json:"attributes,omitempty"`
}

type AddAccountRequest struct {
	//Specifies the username for this account. (Might be 1 to 64 characters in length).
	Username string `json:"username,"`
	//The CHAP secret to use for the initiator.
	//If unspecified, a random secret is created.
	InitiatorSecret string `json:"initiatorSecret,omitempty"`
	//The CHAP secret to use for the target (mutual CHAP authentication).
	//If unspecified, a random secret is created.
	TargetSecret string `json:"targetSecret,omitempty"`
	//List of name-value pairs in JSON object format.
	Attributes interface{} `json:"attributes,omitempty"`
}

type ModifySnapMirrorRelationshipRequest struct {
	//The destination volume in the SnapMirror relationship.
	DestinationVolume SnapMirrorVolumeInfo `json:"destinationVolume,"`
	//Specifies the maximum data transfer rate
	//between the volumes in kilobytes per second.
	//The default value, 0, is unlimited and permits the SnapMirror
	//relationship to fully utilize the available network bandwidth.
	MaxTransferRate int64 `json:"maxTransferRate,omitempty"`
	//Specifies the name of the ONTAP SnapMirror policy for the relationship.
	PolicyName string `json:"policyName,omitempty"`
	//The name of the pre-existing cron schedule on the ONTAP system that is used to update the SnapMirror relationship.
	ScheduleName string `json:"scheduleName,omitempty"`
	//The endpoint ID of the remote ONTAP storage system communicating with the SolidFire cluster.
	SnapMirrorEndpointID int64 `json:"snapMirrorEndpointID,"`
}

type ListInitiatorsRequest struct {
	//The initiator ID at which to begin the listing. You can supply this
	//parameter or the "initiators" parameter, but not both.
	StartInitiatorID int64 `json:"startInitiatorID,omitempty"`
	//The maximum number of initiator objects to return.
	Limit int64 `json:"limit,omitempty"`
	//A list of initiator IDs to retrieve. You can provide a value for this parameter or
	//the "startInitiatorID" parameter, but not both.
	Initiators []int64 `json:"initiators,omitempty"`
}

type RemoveKeyServerFromProviderKmipRequest struct {
	//The ID of the KMIP Key Server to unassign.
	KeyServerID int64 `json:"keyServerID,"`
}

type ModifySnapMirrorEndpointRequest struct {
	//The SnapMirror endpoint to modify.
	SnapMirrorEndpointID int64 `json:"snapMirrorEndpointID,"`
	//The new management IP Address for the ONTAP system.
	ManagementIP string `json:"managementIP,omitempty"`
	//The new management username for the ONTAP system.
	Username string `json:"username,omitempty"`
	//The new management password for the ONTAP system.
	Password string `json:"password,omitempty"`
}

type ListSnapshotsRequest struct {
	//Retrieves snapshots for a volume. If volumeID is not provided,
	//all snapshots for all volumes are returned.
	VolumeID int64 `json:"volumeID,omitempty"`
	//Retrieves information for a specific snapshot ID.
	SnapshotID int64 `json:"snapshotID,omitempty"`
	//List the active branch that each volume has; related to internal snapshots managed by the system.
	Internal bool `json:"internal,omitempty"`
}

type ResurrectDeadVirtualVolumeRequest struct {
	//VolumeID of the volume to be recovered.
	VolumeID int64 `json:"volumeID,"`
	//Service ID of the dead secondary slice service to recover.
	DeadServiceID int64 `json:"deadServiceID,"`
}

type DeleteAuthSessionsByUsernameRequest struct {
	//Name that uniquely identifies the user.
	//When authMethod is Cluster, this specifies the ClusterAdmin username.
	//When authMethod is Ldap, this specifies the user's LDAP DN.
	//When authMethod is Idp, this may specify the user's IdP uid or NameID. If the IdP is not configured to return either, this specifies a random UUID issued when the session was created.
	//Only a caller in the ClusterAdmins / Administrator AccessGroup can provide this parameter.
	Username string `json:"username,omitempty"`
	//Authentication method of the user sessions to be deleted.
	//Only a caller in the ClusterAdmins / Administrator AccessGroup can provide this parameter.
	AuthMethod AuthMethod `json:"authMethod,omitempty"`
}

type UpdateAuthSessionRequest struct {
	//UUID for the auth session to be refreshed.
	SessionID string `json:"sessionID,"`
}

type SetLoginSessionInfoRequest struct {
	//Cluster authentication expiration period. Formatted in
	//HH:mm:ss. For example, 01:30:00, 00:90:00, and 00:00:5400 can
	//be used to equal a 90 minute timeout period. The default value is 30 minutes.
	//The minimum value is 1 minute.
	Timeout string `json:"timeout,omitempty"`
}

type DeleteGroupSnapshotRequest struct {
	//Specifies the unique ID of the group snapshot.
	GroupSnapshotID int64 `json:"groupSnapshotID,"`
	//Specifies whether to preserve snapshots or delete them. Valid values are:
	//true: Snapshots are preserved, but group association is removed.
	//false: The group and snapshots are deleted.
	SaveMembers bool `json:"saveMembers,"`
}

type CheckProposedNodeAdditionsRequest struct {
	//List of node IPs for the nodes that will be added to the cluster.
	Nodes []string `json:"nodes,"`
}

type CreateSnapshotRequest struct {
	//Specifies the unique ID of the volume image from which to copy.
	VolumeID int64 `json:"volumeID,"`
	//Specifies the unique ID of a snapshot from which the new snapshot
	//is made. The snapshotID passed must be a snapshot on the given volume.
	SnapshotID int64 `json:"snapshotID,omitempty"`
	//Specifies a name for the snapshot. If unspecified, the date and time the snapshot was taken is used.
	Name string `json:"name,omitempty"`
	//Replicates the snapshot created to a remote cluster.
	//Possible values are:
	//true: The snapshot is replicated to remote storage.
	//false: Default. The snapshot is not replicated.
	EnableRemoteReplication bool `json:"enableRemoteReplication,omitempty"`
	//Specifies the amount of time for which the snapshot is retained. The format is HH:mm:ss.
	Retention string `json:"retention,omitempty"`
	//List of name-value pairs in JSON object format.
	Attributes interface{} `json:"attributes,omitempty"`
	//Label used by SnapMirror software to specify snapshot retention policy on SnapMirror endpoint.
	SnapMirrorLabel string `json:"snapMirrorLabel,omitempty"`
}

type EnableLdapAuthenticationRequest struct {
	//Identifies which user authentication method to use. Must be one of the following:
	//DirectBind
	//SearchAndBind
	AuthType string `json:"authType,omitempty"`
	//The base DN of the tree to start the group search (will do a subtree search from here).
	GroupSearchBaseDN string `json:"groupSearchBaseDN,omitempty"`
	//For use with the CustomFilter search type, an LDAP filter to use to return the DNs of a users groups. The string can have placeholder text of %USERNAME% and %USERDN% to be replaced with their username and full userDN as needed.
	GroupSearchCustomFilter string `json:"groupSearchCustomFilter,omitempty"`
	//Controls the default group search filter used, and must be one of the following:
	//NoGroups: No group support.
	//ActiveDirectory: Nested membership of all of a users AD groups.
	//MemberDN: MemberDN style groups (single level).
	GroupSearchType string `json:"groupSearchType,omitempty"`
	//A fully qualified DN to log in with to perform an LDAP search for the user (needs read access to the LDAP directory).
	SearchBindDN string `json:"searchBindDN,omitempty"`
	//The password for the searchBindDN account used for searching.
	SearchBindPassword string `json:"searchBindPassword,omitempty"`
	//A comma-separated list of LDAP server URIs (examples: "ldap://1.2.3.4" and ldaps://1.2.3.4:123")
	ServerURIs []string `json:"serverURIs,"`
	//A string that is used to form a fully qualified user DN. The string should have the placeholder text %USERNAME%, which is replaced with the username of the authenticating user.
	UserDNTemplate string `json:"userDNTemplate,omitempty"`
	//The base DN of the tree to start the search (will do a subtree search from here).
	UserSearchBaseDN string `json:"userSearchBaseDN,omitempty"`
	//The LDAP filter to use. The string should have the placeholder text %USERNAME% which is replaced with the username of the authenticating user. Example: (&(objectClass=person)(sAMAccountName=%USERNAME%)) will use the sAMAccountName field in Active Directory to match the username entered at cluster login.
	UserSearchFilter string `json:"userSearchFilter,omitempty"`
}

type DeleteKeyServerKmipRequest struct {
	//The ID of the KMIP Key Server to delete.
	KeyServerID int64 `json:"keyServerID,"`
}

type TestConnectEnsembleRequest struct {
	//Uses a comma-separated list of ensemble node cluster IP addresses to test connectivity. This parameter is optional.
	Ensemble string `json:"ensemble,omitempty"`
	Force    bool   `json:"force,omitempty"`
}

type AddLdapClusterAdminRequest struct {
	//The distinguished user name for the new LDAP cluster admin.
	Username string `json:"username,"`
	//Controls which methods this Cluster Admin can use. For more details on the levels of access, see the Access Control appendix in the SolidFire API Reference.
	Access []string `json:"access,"`
	//Accept the End User License Agreement. Set to true to add a cluster administrator account to the system. If omitted or set to false, the method call fails.
	AcceptEula bool `json:"acceptEula,omitempty"`
	//List of name-value pairs in JSON object format.
	Attributes interface{} `json:"attributes,omitempty"`
}

type CreateKeyProviderKmipRequest struct {
	//The name to associate with the created KMIP Key Provider.  This name is only used for display purposes and does not need to be unique.
	KeyProviderName string `json:"keyProviderName,"`
}

type RestartServicesRequest struct {
	//Required parameter to successfully restart services on a node.
	Force bool `json:"force,"`
	//Service name to be restarted.
	Service string `json:"service,omitempty"`
	//Action to perform on the service (start, stop, restart).
	Action string `json:"action,omitempty"`
}

type CreateSnapMirrorEndpointRequest struct {
	//The management IP address of the remote SnapMirror endpoint.
	ManagementIP string `json:"managementIP,"`
	//The management username for the ONTAP system.
	Username string `json:"username,"`
	//The management password for the ONTAP system.
	Password string `json:"password,"`
}

type CreateBackupTargetRequest struct {
	//The name of the backup target.
	Name string `json:"name,"`
	//List of name-value pairs in JSON object format.
	Attributes interface{} `json:"attributes,"`
}

type CopyVolumeRequest struct {
	//VolumeID of the volume to be read from.
	VolumeID int64 `json:"volumeID,"`
	//VolumeID of the volume to be overwritten.
	DstVolumeID int64 `json:"dstVolumeID,"`
	//ID of the snapshot that is used as the source of the clone.
	//If no ID is provided, the current active volume is used.
	SnapshotID int64 `json:"snapshotID,omitempty"`
}

type SetNetworkConfigRequest struct {
	//An object containing node network settings to modify.
	Network            NetworkParams `json:"network,"`
	Force              bool          `json:"force,omitempty"`
	V__useTheForceLuke bool          `json:"__useTheForceLuke,omitempty"`
}

type GetFeatureStatusRequest struct {
	//Specifies the feature for which the status is returned. Valid values are:
	//vvols: Retrieve status for the NetApp SolidFire VVols
	//cluster feature.
	//FipsDrives: Enable the NetApp SolidFire cluster FIPS 140-2 drive support.
	Feature string `json:"feature,omitempty"`
}

type BindVirtualVolumesRequest struct {
	//The UUID of the VVol to bind.
	VirtualVolumeIDs []string `json:"virtualVolumeIDs,"`
	//The UUID of the ESX host.
	VirtualVolumeHostID string `json:"virtualVolumeHostID,"`
	//Normal or Start?
	BindContext string `json:"bindContext,"`
}

type ListVirtualVolumeHostsRequest struct {
	//A list of virtual volume host IDs for which to retrieve information. If you omit this parameter, the method returns information about all virtual volume hosts.
	VirtualVolumeHostIDs []string `json:"virtualVolumeHostIDs,omitempty"`
}

type ModifyVolumeAccessGroupRequest struct {
	//The ID of the volume access group to modify.
	VolumeAccessGroupID int64 `json:"volumeAccessGroupID,"`
	//The ID of the SolidFire virtual network to associate the volume access group with.
	VirtualNetworkID []int64 `json:"virtualNetworkID,omitempty"`
	//The ID of the SolidFire virtual network to associate the volume access group with.
	VirtualNetworkTags []int64 `json:"virtualNetworkTags,omitempty"`
	//The new name for this volume access group. Not required to be unique, but recommended.
	Name string `json:"name,omitempty"`
	//List of initiators to include in the volume access group. If unspecified, the access group's configured initiators are not modified.
	Initiators []string `json:"initiators,omitempty"`
	//List of volumes to initially include in the volume access group. If unspecified, the access group's volumes are not modified.
	Volumes []int64 `json:"volumes,omitempty"`
	//true: Default. Delete initiator objects after they are removed from a volume access group.
	//false: Do not delete initiator objects after they are removed from a volume access group.
	DeleteOrphanInitiators bool `json:"deleteOrphanInitiators,omitempty"`
	//List of name-value pairs in JSON object format.
	Attributes interface{} `json:"attributes,omitempty"`
}

type DeleteClusterInterfacePreferenceRequest struct {
	//Name of the cluster interface preference.
	Name string `json:"name,"`
}

type RemoveBlockFromBSRequest struct {
	//The blockID of the block that will be removed.
	Hash string `json:"hash,"`
	//The protection schemes for which the block needs to be removed.
	//The default value is the list of enabledProtectionSchemes.
	ProtectionSchemes []ProtectionScheme `json:"protectionSchemes,omitempty"`
}

type ListSnapMirrorNodesRequest struct {
	//If provided, the system lists the nodes of
	//the endpoint with the specified snapMirrorEndpointID.
	//If not provided, the system lists the
	//nodes of all known SnapMirror endpoints.
	SnapMirrorEndpointID int64 `json:"snapMirrorEndpointID,omitempty"`
}

type ListVolumeStatsByVirtualVolumeRequest struct {
	//A list of one or more virtual volume IDs for which to retrieve information. If you specify this parameter, the method returns information about only these virtual volumes.
	VirtualVolumeIDs []string `json:"virtualVolumeIDs,omitempty"`
}

type ModifyClusterFullThresholdRequest struct {
	//The number of nodes of capacity remaining in the cluster before the system triggers a
	//capacity notification.
	Stage2AwareThreshold int64 `json:"stage2AwareThreshold,omitempty"`
	//The percentage of block storage utilization below the "Error" threshold that causes the
	//system to trigger a cluster "Warning" alert.
	Stage3BlockThresholdPercent int64 `json:"stage3BlockThresholdPercent,omitempty"`
	//The percentage of metadata storage utilization below the "Error" threshold that causes the
	//system to trigger a cluster "Warning" alert.
	Stage3MetadataThresholdPercent int64 `json:"stage3MetadataThresholdPercent,omitempty"`
	//A value representative of the number of times metadata space can be overprovisioned relative to the amount of space available. For example, if there was enough metadata space to store 100 TiB of volumes and this number was set to 5, then 500 TiB worth of volumes can be created.
	MaxMetadataOverProvisionFactor int64 `json:"maxMetadataOverProvisionFactor,omitempty"`
}

type GetAuthConfigurationRequest struct {
	//Indicates the configuration data that will be retrieved.
	AuthType AuthConfigType `json:"authType,"`
}

type ListDriveHardwareRequest struct {
	//To run this command, the force parameter must be set to true.
	Force bool `json:"force,"`
}

type EnableSnmpRequest struct {
	//If set to "true", then SNMP v3 is enabled on each node in the
	//cluster. If set to "false", then SNMP v2 is enabled.
	SnmpV3Enabled bool `json:"snmpV3Enabled,"`
}

type CompleteVolumePairingRequest struct {
	//The key returned from the StartVolumePairing method.
	VolumePairingKey string `json:"volumePairingKey,"`
	//The ID of the volume on which to complete the pairing process.
	VolumeID int64 `json:"volumeID,"`
}

type ListActivePairedVolumesRequest struct {
	//The beginning of the range of active paired volumes to return.
	StartVolumeID int64 `json:"startVolumeID,omitempty"`
	//Maximum number of active paired volumes to return.
	Limit int64 `json:"limit,omitempty"`
}

type GetIpmiInfoRequest struct {
	//Force this method to run on all nodes. Only needed when issuing the API to a cluster instead of a single node.
	Force bool `json:"force,"`
}

type RemoveVolumesFromVolumeAccessGroupRequest struct {
	//The ID of the volume access group to remove volumes from.
	VolumeAccessGroupID int64 `json:"volumeAccessGroupID,"`
	//The ID of the volume access group to remove volumes from.
	Volumes []int64 `json:"volumes,"`
}

type RemoveBackupTargetRequest struct {
	//The unique target ID of the target to remove.
	BackupTargetID int64 `json:"backupTargetID,"`
}

type ModifyAccountRequest struct {
	//Specifies the AccountID for the account to be modified.
	AccountID int64 `json:"accountID,"`
	//Specifies the username associated with the
	//account. (Might be 1 to 64 characters in length).
	Username string `json:"username,omitempty"`
	//Sets the status for the account. Possible values are:
	//active: The account is active and connections are allowed.
	//locked: The account is locked and connections are refused.
	Status string `json:"status,omitempty"`
	//The CHAP secret to use for the initiator.
	InitiatorSecret string `json:"initiatorSecret,omitempty"`
	//The CHAP secret to use for the target (mutual CHAP authentication).
	TargetSecret string `json:"targetSecret,omitempty"`
	//List of name-value pairs in JSON object format.
	Attributes interface{} `json:"attributes,omitempty"`
}

type RemoveVolumePairRequest struct {
	//The ID of the volume on which to stop the replication process.
	VolumeID int64 `json:"volumeID,"`
}

type ModifyScheduleRequest struct {
	//The unique name assigned to the schedule.
	ScheduleName string `json:"scheduleName,"`
	//Indicates the days of the month that a snapshot will be made.
	//Valid values are 1 through 31.
	Monthdays []int64 `json:"monthdays,omitempty"`
	//Day of the week the snapshot is to be created. The day of the
	//week starts at Sunday with the value of "0" and an offset of
	//"1." Monday has a value of "1" with an offset of "1", and so on.
	Weekdays int64 `json:"weekdays,omitempty"`
	//Specifies the number of hours between snapshots or hour in GMT time that the snapshot will occur in Days of Week or Days of Month mode. Valid values are 0 through 24.
	Hours int64 `json:"hours,"`
	//Specifies the number of minutes between snapshots or the minute in GMT time that the snapshot will occur in Days of Week or Days of Month mode. Valid values are 0 through 59.
	Minutes int64 `json:"minutes,"`
	//The type of schedule. Only "snapshot" is supported at this time.
	ScheduleType string `json:"scheduleType,"`
	//Use the "frequency" string to indicate the frequency of the snapshot.
	//Valid values for "frequency" are:
	//Days of Week
	//Days of Month
	//Time Interval
	Attributes interface{} `json:"attributes,"`
	//The unique ID of the schedule.
	ScheduleID int64 `json:"scheduleID,"`
	//The unique name given to the schedule, the retention period
	//for the snapshot that was created, and the volume ID of the
	//volume from which the snapshot was created.
	//Valid values are:
	//volumeID: The ID of the volume to be included in the snapshot. (Integer)
	//volumes: A list of volume IDs to be included in the group snapshot. (Array of integers)
	//name: The snapshot name to be used. (String)
	//enableRemoteReplication: Indicates if the snapshot should be included in remote replication. (Boolean)
	//retention: The amount of time for which the snapshot is retained in HH:mm:ss. (String)
	ScheduleInfo  ScheduleInfo `json:"scheduleInfo,omitempty"`
	LastRunStatus string       `json:"lastRunStatus,omitempty"`
	//Specifies whether or not to run the snapshot the next
	//time the scheduler is active. When set to true, the scheduled
	//snapshot runs the next time the scheduler is active and resets back to false.
	//Valid values are:
	//true
	//false
	RunNextInterval bool `json:"runNextInterval,omitempty"`
	//Indicates if the schedule should be paused or not.
	//Valid values are:
	//true
	//false
	Paused bool `json:"paused,omitempty"`
	//Indicates if the schedule recurs or not.
	//Valid values are:
	//true
	//false
	Recurring bool `json:"recurring,omitempty"`
	HasError  bool `json:"hasError,omitempty"`
	//Indicates the date the first time the schedule began or will
	//begin.
	StartingDate string `json:"startingDate,omitempty"`
	//Indicates if the schedule is marked for deletion.
	//Valid values are:
	//true
	//false
	ToBeDeleted bool `json:"toBeDeleted,omitempty"`
	//Label used by SnapMirror software to specify snapshot retention policy on SnapMirror endpoint.
	SnapMirrorLabel string `json:"snapMirrorLabel,omitempty"`
}

type DeleteSnapshotRequest struct {
	//The ID of the snapshot to be deleted.
	SnapshotID int64 `json:"snapshotID,"`
}

type CancelCloneRequest struct {
	//The cloneID for the ongoing clone process.
	CloneID int64 `json:"cloneID,"`
}

type SetLoginBannerRequest struct {
	//The desired text of the Terms of Use banner.
	Banner string `json:"banner,omitempty"`
	//The status of the Terms of Use banner. Possible values:
	//true: The Terms of Use banner is displayed upon web interface login.
	//false: The Terms of Use banner is not displayed upon web interface login.
	Enabled bool `json:"enabled,omitempty"`
}

type ListSnapMirrorEndpointsRequest struct {
	//Return only the objects associated with these IDs.
	//If no IDs are provided or the array is empty,
	//the method returns all SnapMirror endpoint IDs.
	SnapMirrorEndpointIDs []int64 `json:"snapMirrorEndpointIDs,omitempty"`
}

type PurgeDeletedVolumeRequest struct {
	//The ID of the volume to be purged.
	VolumeID int64 `json:"volumeID,"`
}

type GetAccountByNameRequest struct {
	//Username for the account.
	Username string `json:"username,"`
}

type DeleteInitiatorsRequest struct {
	//An array of IDs of initiators to delete.
	Initiators []int64 `json:"initiators,"`
}

type ListAsyncResultsRequest struct {
	//An optional list of types of results. You can use this list to restrict the
	//results to only these types of operations. Possible values are:
	//BulkVolume: Copy operations between volumes, such as backups or
	//restores.
	//Clone: Volume cloning operations.
	//DriveRemoval: Operations involving the system copying data from a
	//drive in preparation to remove it from the cluster.
	//RtfiPendingNode: Operations involving the system installing
	//compatible software on a node before adding it to the cluster
	AsyncResultTypes []string `json:"asyncResultTypes,omitempty"`
}

type GetClusterInterfacePreferenceRequest struct {
	//Name of the cluster interface preference.
	Name string `json:"name,"`
}

type ResetConstantsRequest struct {
	//List of cluster constants to reset to their default values.
	Constants []string `json:"constants,omitempty"`
	//Reset all cluster constants to their default values.
	ResetAll bool `json:"resetAll,omitempty"`
}

type ResumeSnapMirrorRelationshipRequest struct {
	//The endpoint ID of the remote ONTAP storage system communicating with the SolidFire cluster.
	SnapMirrorEndpointID int64 `json:"snapMirrorEndpointID,"`
	//The destination volume in the SnapMirror relationship.
	DestinationVolume SnapMirrorVolumeInfo `json:"destinationVolume,"`
}

type GetAccountByIDRequest struct {
	//Specifies the account for which details are gathered.
	AccountID int64 `json:"accountID,"`
}

type CheckPingOnVlanRequest struct {
	//The ipv4 source address to be used in the ICMP ping packets
	//sourceAddressV4 or sourceAddressV6 is required
	SourceAddressV4 string `json:"sourceAddressV4,omitempty"`
	//The ipv6 source address to be used in the ICMP ping packets
	//sourceAddressV4 or sourceAddressV6 is required
	SourceAddressV6 string `json:"sourceAddressV6,omitempty"`
	//Existing interface on which the temporary vlan interface is created
	Interface string `json:"interface,"`
	//VLAN on which host addresses reachability needs to be tested
	//The temporary vlan interface is created with this tag
	VirtualNetworkTag int64 `json:"virtualNetworkTag,"`
	//String containing comma separated IP addresses whose reachability needs to be tested.
	Hosts string `json:"hosts,"`
}

type DeleteKeyProviderKmipRequest struct {
	//The ID of the Key Provider to delete.
	KeyProviderID int64 `json:"keyProviderID,"`
}

type AddVirtualNetworkRequest struct {
	//A unique virtual network (VLAN) tag. Supported values are 1 through 4094.The number zero (0) is not supported.
	VirtualNetworkTag int64 `json:"virtualNetworkTag,"`
	//A user-defined name for the new virtual network.
	Name string `json:"name,"`
	//Unique range of IP addresses to include in the virtual network.
	//Attributes for this parameter are:
	//start: The start of the IP address range. (String)
	//size: The number of IP addresses to include in the block. (Integer)
	AddressBlocks []AddressBlockParams `json:"addressBlocks,"`
	//Unique network mask for the virtual network being created.
	Netmask string `json:"netmask,"`
	//Unique storage IP address for the virtual network being created.
	Svip string `json:"svip,"`
	//The IP address of a gateway of the virtual network. This parameter is valid only if the namespace parameter is set to true (meaning VRF is enabled).
	Gateway string `json:"gateway,omitempty"`
	//When set to true, enables the Routable Storage VLANs functionality by recreating the virtual network and configuring a namespace to contain it.
	//When set to false, disables the VRF functionality for the virtual network.
	//Changing this value disrupts traffic running through this virtual network.
	Namespace bool `json:"namespace,omitempty"`
	//List of name-value pairs in JSON object format.
	Attributes interface{} `json:"attributes,omitempty"`
}

type CloneVolumeRequest struct {
	//VolumeID for the volume to be cloned.
	VolumeID int64 `json:"volumeID,"`
	//The name of the new cloned volume. Must be 1 to 64 characters in length.
	Name string `json:"name,"`
	//AccountID for the owner of the new volume. If unspecified, the
	//accountID of the owner of the volume being cloned is used.
	NewAccountID int64 `json:"newAccountID,omitempty"`
	//New size of the volume, in bytes. Must be greater or less than the size of
	//the volume being cloned. If unspecified, the volume size is not
	//changed. Size is rounded to the nearest 1MB.
	NewSize int64 `json:"newSize,omitempty"`
	//Specifies the level of access allowed for the new volume. If unspecified, the
	//level of access of the volume being cloned is used. If replicationTarget is
	//is passed and the volume is not paired, the access gets set to locked.
	Access VolumeAccess `json:"access,omitempty"`
	//ID of the snapshot that is used as the source of the clone. If no ID is
	//provided, the current active volume is used.
	SnapshotID int64 `json:"snapshotID,omitempty"`
	//List of name-value pairs in JSON object format.
	Attributes interface{} `json:"attributes,omitempty"`
	//Specifies whether the new volume should use 512-byte sector emulation.
	//If unspecified, the setting of the volume being cloned is used.
	Enable512e bool `json:"enable512e,omitempty"`
	//Specifies whether SnapMirror replication is enabled or not. Defaults to false.
	EnableSnapMirrorReplication bool `json:"enableSnapMirrorReplication,omitempty"`
}

type ListVolumeStatsByVolumeRequest struct {
	//Specifies that virtual volumes are included in the response by default.
	//To exclude virtual volumes, set to false.
	IncludeVirtualVolumes bool `json:"includeVirtualVolumes,omitempty"`
}

type DeleteAuthSessionRequest struct {
	//UUID for the auth session to be deleted.
	SessionID string `json:"sessionID,"`
}

type ResizeDrivesRequest struct {
	//List of driveIDs to resize.
	Drives []int64 `json:"drives,"`
	//The new usable capacity in bytes for the given drive(s).
	NewUsableSize int64 `json:"newUsableSize,"`
}

type GetProtectionSchemesRequest struct {
	//Force this method to run on all nodes. Only needed when issuing the API to a cluster instead of a single node.
	Force bool `json:"force,omitempty"`
}

type ModifyVirtualVolumeHostRequest struct {
	//The GUID of the ESX host.
	VirtualVolumeHostID string `json:"virtualVolumeHostID,"`
	//The GUID of the ESX Cluster.
	ClusterID string `json:"clusterID,omitempty"`
	//A list of PEs the host is aware of.
	VisibleProtocolEndpointIDs []string `json:"visibleProtocolEndpointIDs,omitempty"`
	//List of iSCSI initiator IQNs for the host.
	InitiatorNames []string `json:"initiatorNames,omitempty"`
	//IP or DNS name for the host.
	HostAddress string `json:"hostAddress,omitempty"`
	//The UUID of the host making the call to the virtual volume.
	CallingVirtualVolumeHostID string `json:"callingVirtualVolumeHostID,omitempty"`
}

type DisableEncryptionAtRestRequest struct {
	//If true, allow disable encryption at rest even when bin sync is in progress.
	ForceDuringBinSync bool `json:"forceDuringBinSync,omitempty"`
}

type CreateStorageContainerRequest struct {
	//The name of the storage container. Follows SolidFire account
	//naming restrictions.
	Name string `json:"name,"`
	//The secret for CHAP authentication for the initiator.
	InitiatorSecret string `json:"initiatorSecret,omitempty"`
	//The secret for CHAP authentication for the target.
	TargetSecret string `json:"targetSecret,omitempty"`
	//Non-storage container account that will become a
	//storage container.
	AccountID int64 `json:"accountID,omitempty"`
}

type ListVolumesForAccountRequest struct {
	//Returns all volumes owned by this AccountID.
	AccountID int64 `json:"accountID,"`
	//The ID of the first volume to list.
	//This can be useful for paging results.
	//By default, this starts at the lowest VolumeID.
	StartVolumeID int64 `json:"startVolumeID,omitempty"`
	//The maximum number of volumes to return from the API.
	Limit int64 `json:"limit,omitempty"`
	//Specifies that virtual volumes are included in the response by default.
	//To exclude virtual volumes, set to false.
	IncludeVirtualVolumes bool `json:"includeVirtualVolumes,omitempty"`
}

type DeleteVolumeAccessGroupRequest struct {
	//The ID of the volume access group
	//to be deleted.
	VolumeAccessGroupID int64 `json:"volumeAccessGroupID,"`
	//true: Default. Delete initiator objects after they are removed from a volume access group.
	//false: Do not delete initiator objects after they are removed from a volume access group.
	DeleteOrphanInitiators bool `json:"deleteOrphanInitiators,omitempty"`
	//Adding this flag will force the volume access group to be deleted even though it has a Virtual Network ID or Tag.
	//true: Volume access group will be deleted.
	//false: Default. Do not delete the volume access group if it has a Virtual Network ID or Tag.
	Force bool `json:"force,omitempty"`
}

type UnbindAllVirtualVolumesFromHostRequest struct {
	//The UUID of the host making the call to the virtual volume.
	VirtualVolumeHostID string `json:"virtualVolumeHostID,"`
}

type UpdateBulkVolumeStatusRequest struct {
	//The key assigned during initialization of a
	//StartBulkVolumeRead or StartBulkVolumeWrite session.
	Key string `json:"key,"`
	//The status of the given bulk volume job. The system sets the status. Possible values are:
	//running: Jobs that are still active.
	//complete: Jobs that are done.
	//failed: Jobs that failed.
	Status string `json:"status,"`
	//The completed progress of the bulk volume job as a
	//percentage value.
	PercentComplete string `json:"percentComplete,omitempty"`
	//The message returned indicating the status of the bulk volume job after the job is complete.
	Message string `json:"message,omitempty"`
	//JSON attributes; updates what is on the bulk volume job.
	Attributes interface{} `json:"attributes,omitempty"`
}

type TestDrivesRequest struct {
	//Specifies the number of minutes to run the test.
	Minutes int64 `json:"minutes,omitempty"`
	//Required parameter to successfully test the drives on the node.
	Force bool `json:"force,omitempty"`
}

type AddNodesRequest struct {
	// List of pending NodeIDs for the nodes to be added. You can  obtain the list of pending nodes using the ListPendingNodes method.
	PendingNodes []int64 `json:"pendingNodes,"`
	//If true, RTFI will be performed on the nodes.  The default behavior is to perform RTFI.
	AutoInstall bool `json:"autoInstall,omitempty"`
	//Allows nodes to be added during an upgrade
	ForceDuringUpgrade bool `json:"forceDuringUpgrade,omitempty"`
	//Allows addition of non-Fips nodes to cluster while FipsDrives feature is enabled.
	ForceNonFips bool `json:"forceNonFips,omitempty"`
	//Allows addition of non-encryption capable nodes to cluster while EAR is enabled.
	ForceNE bool `json:"forceNE,omitempty"`
	//Skips node type checks.
	//This skips any checks that validate whether the node type is deprecated.
	SkipNodeTypeChecks bool `json:"skipNodeTypeChecks,omitempty"`
}

type SetNtpInfoRequest struct {
	//List of NTP servers to add to each nodes NTP configuration.
	Servers []string `json:"servers,"`
	//Enables every node in the cluster as a broadcast client.
	Broadcastclient bool `json:"broadcastclient,omitempty"`
}

type GetOntapVersionInfoRequest struct {
	//If provided, the system lists the version information from the
	//endpoint with the specified snapMirrorEndpointID.
	//If not provided, the system lists the version
	//information of all known SnapMirror endpoints.
	SnapMirrorEndpointID int64 `json:"snapMirrorEndpointID,omitempty"`
}

type GetScheduleRequest struct {
	//Specifies the unique ID of the schedule or multiple schedules to display.
	ScheduleID int64 `json:"scheduleID,"`
}

type CreateSnapMirrorRelationshipRequest struct {
	//The endpoint ID of the remote ONTAP storage system communicating with the SolidFire cluster.
	SnapMirrorEndpointID int64 `json:"snapMirrorEndpointID,"`
	//The source volume in the relationship.
	SourceVolume SnapMirrorVolumeInfo `json:"sourceVolume,"`
	//The destination volume in the relationship.
	DestinationVolume SnapMirrorVolumeInfo `json:"destinationVolume,"`
	//The type of relationship.
	//On SolidFire systems, this value is always "extended_data_protection".
	RelationshipType string `json:"relationshipType,omitempty"`
	//Specifies the name of the ONTAP SnapMirror policy for the relationship.
	//If not specified, the default policy name is MirrorLatest.
	PolicyName string `json:"policyName,omitempty"`
	//The name of the preexisting cron schedule on the ONTAP system
	//that is used to update the SnapMirror relationship.
	//If no schedule is designated, snapMirror updates are not scheduled
	//and must be updated manually.
	ScheduleName string `json:"scheduleName,omitempty"`
	//Specifies the maximum data transfer rate
	//between the volumes in kilobytes per second.
	//The default value, 0, is unlimited and permits the SnapMirror
	//relationship to fully utilize the available network bandwidth.
	MaxTransferRate int64 `json:"maxTransferRate,omitempty"`
}

type SetSnmpTrapInfoRequest struct {
	//List of hosts that are to receive the traps generated by the Cluster Master. At least one object is required if any one of the trap types is enabled.
	TrapRecipients []SnmpTrapRecipient `json:"trapRecipients,omitempty"`
	//If the value is set to true, a corresponding solidFireClusterFaultNotification is sent to the configured list of trap recipients when a cluster fault is logged. The default value is false.
	ClusterFaultTrapsEnabled bool `json:"clusterFaultTrapsEnabled,"`
	//If the value is set to true, a corresponding solidFireClusterFaultResolvedNotification is sent to the configured list of trap recipients when a cluster fault is resolved. The default value is false.
	ClusterFaultResolvedTrapsEnabled bool `json:"clusterFaultResolvedTrapsEnabled,"`
	//If the value is set to true, a corresponding solidFireClusterEventNotification is sent to the configured list of trap recipients when a cluster event is logged. The default value is false.
	ClusterEventTrapsEnabled bool `json:"clusterEventTrapsEnabled,"`
}

type SetDefaultProtectionSchemeRequest struct {
	//The new default scheme
	DefaultProtectionScheme ProtectionScheme `json:"defaultProtectionScheme,"`
}

type TestKeyProviderKmipRequest struct {
	//The ID of the Key Provider to test.
	KeyProviderID int64 `json:"keyProviderID,"`
}

type ListAuthSessionsByUsernameRequest struct {
	//Name that uniquely identifies the user.
	//When authMethod is Cluster, this specifies the ClusterAdmin username.
	//When authMethod is Ldap, this specifies the user's LDAP DN.
	//When authMethod is Idp, this may specify the user's IdP uid or NameID. If the IdP is not configured to return either, this specifies a random UUID issued when the session was created.
	//Only a caller in the ClusterAdmins / Administrator AccessGroup can provide this parameter.
	Username string `json:"username,omitempty"`
	//Authentication method of the user sessions to be listed.
	//Only a caller in the ClusterAdmins / Administrator AccessGroup can provide this parameter.
	AuthMethod AuthMethod `json:"authMethod,omitempty"`
}

type AbortSnapMirrorRelationshipRequest struct {
	//The endpoint ID of the remote ONTAP storage system communicating with the SolidFire cluster.
	SnapMirrorEndpointID int64 `json:"snapMirrorEndpointID,"`
	//The destination volume in the SnapMirror relationship.
	DestinationVolume SnapMirrorVolumeInfo `json:"destinationVolume,"`
	//Determines whether or not to clear the restart checkpoint.
	ClearCheckpoint bool `json:"clearCheckpoint,omitempty"`
}

type SetSupplementalTlsCiphersRequest struct {
	//The supplemental cipher suite names using the OpenSSL naming scheme. Use of cipher suite names is case-insensitive.
	SupplementalCiphers []string `json:"supplementalCiphers,"`
}

type SetNodeSupplementalTlsCiphersRequest struct {
	//The supplemental cipher suite names using the OpenSSL naming scheme. Use of cipher suite names is case-insensitive.
	SupplementalCiphers []string `json:"supplementalCiphers,"`
}

type ListSnapMirrorSchedulesRequest struct {
	//If provided, the system lists the schedules of the
	//endpoint with the specified SnapMirror endpoint ID.
	//If not provided, the system lists the schedules
	//of all known SnapMirror endpoints.
	SnapMirrorEndpointID int64 `json:"snapMirrorEndpointID,omitempty"`
}

type ModifyClusterAdminRequest struct {
	//ClusterAdminID for the cluster admin, LDAP cluster admin, or IdP cluster admin to modify.
	ClusterAdminID int64 `json:"clusterAdminID,"`
	//Password used to authenticate this cluster admin.
	//This parameter does not apply for an LDAP or IdP cluster admin.
	Password string `json:"password,omitempty"`
	//Controls which methods this cluster admin can use. For more details, see Access Control in the Element API Reference Guide.
	Access []string `json:"access,omitempty"`
	//List of name-value pairs in JSON object format.
	Attributes interface{} `json:"attributes,omitempty"`
}

type GetDriveStatsRequest struct {
	//Specifies the drive for which statistics are gathered.
	DriveID int64 `json:"driveID,"`
}

type InvokeSFApiRequest struct {
	//The name of the method to invoke. This is case sensitive.
	Method string `json:"method,"`
	//An object, normally a dictionary or hashtable of the key/value pairs, to be passed as the params for the method being invoked.
	Parameters string `json:"parameters,omitempty"`
}

type CancelGroupCloneRequest struct {
	//The cloneID for the ongoing clone process.
	GroupCloneID int64 `json:"groupCloneID,"`
}
