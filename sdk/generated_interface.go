package sdk

// THIS IS A GENERATED FILE.  DO NOT MODIFY

import "context"

type SFApi interface {
	// The SolidFire Element OS web UI uses the ListSnapMirrorPolicies method to list all SnapMirror policies on a remote ONTAP system.
	ListSnapMirrorPolicies(ctx context.Context, req *ListSnapMirrorPoliciesRequest) (*ListSnapMirrorPoliciesResult, *SdkError)

	// RollbackVirtualVolume is used to restore a VMware Virtual Volume snapshot.
	RollbackVirtualVolume(ctx context.Context, req *RollbackVirtualVolumeRequest) (*VirtualVolumeAsyncResult, *SdkError)

	// GetVolumeStats enables  you to retrieve high-level activity measurements for a single volume. Values are cumulative from the creation of the volume.
	GetVolumeStats(ctx context.Context, req *GetVolumeStatsRequest) (*GetVolumeStatsResult, *SdkError)

	// ListDeletedVolumes enables you to retrieve the list of volumes that have been marked for deletion and purged from the system.
	ListDeletedVolumes(ctx context.Context, req *ListDeletedVolumesRequest) (*ListDeletedVolumesResult, *SdkError)

	// GetDriveHardwareInfo returns all the hardware information for the given drive. This generally includes details about manufacturers, vendors, versions, and
	// other associated hardware identification information.
	GetDriveHardwareInfo(ctx context.Context, req *GetDriveHardwareInfoRequest) (*GetDriveHardwareInfoResult, *SdkError)

	// You can use the ListClusterPairs method to list all the clusters that a cluster is paired with. This method returns information about active and pending cluster pairings, such as statistics about the current pairing as well as the connectivity and latency (in milliseconds) of the cluster pairing.
	ListClusterPairs(ctx context.Context) (*ListClusterPairsResult, *SdkError)

	// Enables retrieval of the number of virtual volumes currently in the system.
	GetVirtualVolumeCount(ctx context.Context) (*GetVirtualVolumeCountResult, *SdkError)

	// CopyDiffsToVirtualVolume is a three-way merge function.
	CopyDiffsToVirtualVolume(ctx context.Context, req *CopyDiffsToVirtualVolumeRequest) (*VirtualVolumeAsyncResult, *SdkError)

	// StartBulkVolumeWrite enables you to initialize a bulk volume write session on a specified volume. Only two bulk volume processes can run simultaneously on a volume. When you initialize the write session, data is written to a SolidFire storage volume from an external backup source. The external data is accessed by a web server running on an SF-series node. Communications and server
	// interaction information for external data access is passed by a script running on the storage system.
	StartBulkVolumeWrite(ctx context.Context, req *StartBulkVolumeWriteRequest) (*StartBulkVolumeWriteResult, *SdkError)

	// You can use AddClusterAdmin to add a new cluster admin account. A cluster ddmin can manage the cluster using the API and management tools. Cluster admins are completely separate and unrelated to standard tenant accounts.
	// Each cluster admin can be restricted to a subset of the API. NetApp recommends using multiple cluster admin accounts for different users and applications. You should give each cluster admin the minimal permissions necessary; this reduces the potential impact of credential compromise.
	// You must accept the End User License Agreement (EULA) by setting the acceptEula parameter to true to add a cluster administrator account to the system.
	AddClusterAdmin(ctx context.Context, req *AddClusterAdminRequest) (*AddClusterAdminResult, *SdkError)

	// The SetConfig API method enables you to set all the configuration information for the node. This includes the same information available via calls to SetClusterConfig and SetNetworkConfig in one API method.
	// Note: This method is available only through the per-node API endpoint 5.0 or later.
	// Caution: Changing the "bond-mode" on a node can cause a temporary loss of network connectivity. Exercise caution when using this method.
	SetConfig(ctx context.Context, req *SetConfigRequest) (*SetConfigResult, *SdkError)

	// ModifyVolumePair enables you to pause or restart replication between a pair of volumes.
	ModifyVolumePair(ctx context.Context, req *ModifyVolumePairRequest) (*ModifyVolumePairResult, *SdkError)

	// GetStorageContainerEfficiency enables you to retrieve efficiency information about a virtual volume storage container.
	GetStorageContainerEfficiency(ctx context.Context, req *GetStorageContainerEfficiencyRequest) (*GetStorageContainerEfficiencyResult, *SdkError)

	// The RestartNetworking API method enables you to restart the networking services on a node.
	// Warning: This method restarts all networking services on a node, causing temporary loss of networking connectivity.
	// Exercise caution when using this method.
	// Note: This method is available only through the per-node API endpoint 5.0 or later.
	RestartNetworking(ctx context.Context, req *RestartNetworkingRequest) (*interface{}, *SdkError)

	// You can use the GetAPI method to return a list of all the API methods and supported API endpoints that can be used in the system.
	GetAPI(ctx context.Context, req *GetAPIRequest) (*GetAPIResult, *SdkError)

	// Returns the list of KMIP (Key Management Interoperability Protocol) Key Providers which have been created via CreateKeyProviderKmip.  The list can optionally be filtered by specifying additional parameters.
	ListKeyProvidersKmip(ctx context.Context, req *ListKeyProvidersKmipRequest) (*ListKeyProvidersKmipResult, *SdkError)

	// ModifyStorageContainer enables you to make changes to an existing virtual volume storage container.
	ModifyStorageContainer(ctx context.Context, req *ModifyStorageContainerRequest) (*ModifyStorageContainerResult, *SdkError)

	// Creates a new cluster preference and stores it on the storage cluster. The ClusterInterfacePreference
	// related APIs can be used by internal interfaces to the storage cluster such as HCI and UI to store arbitrary
	// information in the cluster. Since the API calls in the UI are visible to customers, these APIs are made public.
	CreateClusterInterfacePreference(ctx context.Context, req *CreateClusterInterfacePreferenceRequest) (*CreateClusterInterfacePreferenceResult, *SdkError)

	// ListProtocolEndpoints enables you to retrieve information about all protocol endpoints in the cluster. Protocol endpoints govern
	// access to their associated virtual volume storage containers.
	ListProtocolEndpoints(ctx context.Context, req *ListProtocolEndpointsRequest) (*ListProtocolEndpointsResult, *SdkError)

	// CloneMultipleVolumes enables you to create a clone of a group of specified volumes. You can assign a consistent set of characteristics
	// to a group of multiple volumes when they are cloned together.
	// Before using groupSnapshotID to clone the volumes in a group snapshot, you must create the group snapshot by using the
	// CreateGroupSnapshot API method or the Element OS Web UI. Using groupSnapshotID is optional when cloning multiple volumes.
	// Note: Cloning multiple volumes is allowed if cluster fullness is at stage 2 or 3. Clones are not created when cluster fullness is
	// at stage 4 or 5.
	CloneMultipleVolumes(ctx context.Context, req *CloneMultipleVolumesRequest) (*CloneMultipleVolumesResult, *SdkError)

	// You can use the GetClusterHardwareInfo method to retrieve the hardware status and information for all Fibre Channel nodes, iSCSI
	// nodes and drives in the cluster. This generally includes details about manufacturers, vendors, versions, and other associated hardware
	// identification information.
	GetClusterHardwareInfo(ctx context.Context, req *GetClusterHardwareInfoRequest) (*GetClusterHardwareInfoResult, *SdkError)

	//
	ListSnapMirrorObjectAttributes(ctx context.Context) (*NeedsWorkListSnapMirrorObjectAttributesResult, *SdkError)

	// Returns the specified KMIP (Key Management Interoperability Protocol) Key Provider object.
	GetKeyProviderKmip(ctx context.Context, req *GetKeyProviderKmipRequest) (*GetKeyProviderKmipResult, *SdkError)

	// ListVirtualNetworks enables you to list all configured virtual networks for the cluster. You can use this method to verify the virtual
	// network settings in the cluster.
	// There are no required parameters for this method. However, to filter the results, you can pass one or more VirtualNetworkID or
	// VirtualNetworkTag values.
	ListVirtualNetworks(ctx context.Context, req *ListVirtualNetworksRequest) (*ListVirtualNetworksResult, *SdkError)

	// You can use the SetNodeSSLCertificate method to set a user SSL certificate and private key for the management node.
	SetNodeSSLCertificate(ctx context.Context, req *SetNodeSSLCertificateRequest) (*SetNodeSSLCertificateResult, *SdkError)

	// SnmpSendTestTraps enables you to test SNMP functionality for a cluster. This method instructs the cluster to send test SNMP traps to the currently configured SNMP manager.
	SnmpSendTestTraps(ctx context.Context) (*SnmpSendTestTrapsResult, *SdkError)

	// GetVirtualVolumeAllocatedBitmap scans a VVol segment and returns the number of
	// chunks not shared between two volumes. This call will return results in less
	// than 30 seconds. If the specified VVol and the base VVil are not related, an
	// error is returned. If the offset/length combination is invalid or out of range
	// an error is returned.
	GetVirtualVolumeUnsharedChunks(ctx context.Context, req *GetVirtualVolumeUnsharedChunksRequest) (*VirtualVolumeUnsharedChunkResult, *SdkError)

	// CreateVolume enables you to create a new (empty) volume on the cluster. As soon as the volume creation is complete, the volume is
	// available for connection via iSCSI.
	CreateVolume(ctx context.Context, req *CreateVolumeRequest) (*CreateVolumeResult, *SdkError)

	// Modifies an existing cluster interface preference.
	ModifyClusterInterfacePreference(ctx context.Context, req *ModifyClusterInterfacePreferenceRequest) (*ModifyClusterInterfacePreferenceResult, *SdkError)

	// Creates a new auth auth session for a user.
	// Returns a AuthSessionInfo.
	// Intended to be used by the element-auth container.
	CreateAuthSession(ctx context.Context, req *CreateAuthSessionRequest) (*CreateAuthSessionResult, *SdkError)

	// AddInitiatorsToVolumeAccessGroup enables you
	// to add initiators to a specified volume access group.
	AddInitiatorsToVolumeAccessGroup(ctx context.Context, req *AddInitiatorsToVolumeAccessGroupRequest) (*ModifyVolumeAccessGroupResult, *SdkError)

	//
	ReadSliceLBAChecksum(ctx context.Context) (*NeedsWorkReadSliceLBAChecksumResult, *SdkError)

	// ListEvents returns events detected on the cluster, sorted from oldest to newest.
	ListEvents(ctx context.Context, req *ListEventsRequest) (*ListEventsResult, *SdkError)

	// ListBulkVolumeJobs enables you to retrieve information about each bulk volume read or write operation that is occurring in the
	// system.
	ListBulkVolumeJobs(ctx context.Context) (*ListBulkVolumeJobsResult, *SdkError)

	// The AbortRecoverDeadVolumes method aborts a previously-initiated RecoverDeadVolumes operation.
	//
	// The cluster will attempt to cancel operations on all provided volume IDs. It will return a warning for
	// any volume ID that was invalid.
	//
	// This command requires either volumeIDs or sliceIDs but not both.
	AbortRecoverDeadVolumes(ctx context.Context, req *AbortRecoverDeadVolumesRequest) (*AbortRecoverDeadVolumesResult, *SdkError)

	// You can use the GetActiveTlsCiphers method to get a list of the TLS ciphers that are currently accepted on the cluster.
	GetActiveTlsCiphers(ctx context.Context) (*GetActiveTlsCiphersResult, *SdkError)

	//
	GetLocalStats(ctx context.Context) (*NeedsWorkGetLocalStatsResult, *SdkError)

	//
	SetConstants(ctx context.Context) (*NeedsWorkSetConstantsResult, *SdkError)

	// Get details about the Encryption At Rest feature.
	GetEncryptionAtRestInfo(ctx context.Context) (*GetEncryptionAtRestInfoResult, *SdkError)

	// Used to assign Nodes to user-defined Protection Domains. This information must be provided for all
	// Active Nodes in the cluster, and no information may be provided for Nodes that are not Active. All Nodes
	// in a given Chassis must be assigned to the same user-defined Protection Domain. The same
	// ProtectionDomainType must be supplied for all nodes. ProtectionDomainTypes that are not user-defined
	// such as Node and Chassis, must not be included. If any of these are not true, the Custom Protection
	// Domains will be ignored, and an appropriate error will be returned.
	SetProtectionDomainLayout(ctx context.Context, req *SetProtectionDomainLayoutRequest) (*SetProtectionDomainLayoutResult, *SdkError)

	//
	IsUpgradeInProgress(ctx context.Context) (*NeedsWorkIsUpgradeInProgressResult, *SdkError)

	// The RecoverDeadVolumes method recovers a volume from a standby slice. This should only be used in
	// circumstances where the primary slice service became permanently inaccessible while standby assignments
	// were enabled.
	//
	// The cluster will return a response immediately if it initiates recovery on the volumes. Users should
	// monitor the slices report for status. The cluster will generate events when each recover operation
	// completes. The cluster will make the standby Inactive after completion and delete the standby slice file
	// (even if the standby is Failed).
	//
	// The cluster will not start recovery on any volume if it cannot start operations on all of them.
	//
	// If a standby is Active then, upon successful recovery, the cluster will automatically promote the dead
	// secondary to live and then it will takeover as primary. If the standby is Failed then the cluster will
	// leave the secondary as dead. Users must then call CloneVolumeFrom Dead and/or PromoteDeadSecondary to
	// continue recovery.
	//
	// This command requires either volumeIDs or sliceIDs but not both.
	RecoverDeadVolumes(ctx context.Context, req *RecoverDeadVolumesRequest) (*RecoverDeadVolumesResult, *SdkError)

	// You can use the ClearClusterFaults method to clear information about both current and previously detected faults. Both resolved
	// and unresolved faults can be cleared.
	ClearClusterFaults(ctx context.Context, req *ClearClusterFaultsRequest) (*ClearClusterFaultsResult, *SdkError)

	// GetOrigin enables you to retrieve the origination certificate for where the node was built. This method might return null if there is no origination certification.
	GetOrigin(ctx context.Context, req *GetOriginRequest) (*GetOriginResult, *SdkError)

	//
	ForceWholeFileSync(ctx context.Context) (*NeedsWorkForceWholeFileSyncResult, *SdkError)

	// RollbackToSnapshot enables you to make an existing snapshot of the "active" volume image. This method creates a new snapshot
	// from an existing snapshot. The new snapshot becomes "active" and the existing snapshot is preserved until you delete it.
	// The previously "active" snapshot is deleted unless you set the parameter saveCurrentState to true.
	// Note: Creating a snapshot is allowed if cluster fullness is at stage 2 or 3. Snapshots are not created when cluster fullness is
	// at stage 4 or 5.
	RollbackToSnapshot(ctx context.Context, req *RollbackToSnapshotRequest) (*RollbackToSnapshotResult, *SdkError)

	// You can use the TestPing API method to validate the
	// connection to all the nodes in a cluster on both 1G and 10G interfaces by using ICMP packets. The test uses the appropriate MTU sizes for each packet based on the MTU settings in the network configuration.
	// Note: This method is available only through the per-node API endpoint 5.0 or later.
	TestPing(ctx context.Context, req *TestPingRequest) (*TestPingResult, *SdkError)

	// The ModifyVolumeAccessGroupLunAssignments
	// method enables you to define custom LUN assignments
	// for specific volumes. This method changes only LUN
	// values set on the lunAssignments parameter in the
	// volume access group. All other LUN assignments remain
	// unchanged. LUN assignment values must be unique for volumes in a volume access group. You cannot define duplicate LUN values within a volume access group. However, you can use the same LUN values again in different volume access groups.
	// Note: Correct LUN values are 0 through 16383. The system generates an exception if you pass a LUN value outside of this range. None of the specified LUN assignments are modified if there is an exception.
	// Caution: If you change a LUN assignment for a volume with active I/O, the I/O can be disrupted. You might need to change the server configuration before changing volume LUN assignments.
	ModifyVolumeAccessGroupLunAssignments(ctx context.Context, req *ModifyVolumeAccessGroupLunAssignmentsRequest) (*ModifyVolumeAccessGroupLunAssignmentsResult, *SdkError)

	// CreateSupportBundle enables you to create a support bundle file under the node's directory. After creation, the bundle is stored on the node as a tar.gz file.
	CreateSupportBundle(ctx context.Context, req *CreateSupportBundleRequest) (*CreateSupportBundleResult, *SdkError)

	//
	GetDebugOptions(ctx context.Context) (*NeedsWorkGetDebugOptionsResult, *SdkError)

	// ListFibreChannelSessions enables you to retrieve information about the active Fibre Channel sessions on a cluster.
	ListFibreChannelSessions(ctx context.Context) (*ListFibreChannelSessionsResult, *SdkError)

	// Test whether the specified KMIP (Key Management Interoperability Protocol) Key Server is functioning normally.
	TestKeyServerKmip(ctx context.Context, req *TestKeyServerKmipRequest) (*TestKeyServerKmipResult, *SdkError)

	// You can use the CloneVolumeFromDead method to create a new volume from a secondary replica
	// that may not have all data acknowledged to the client. This should only be used in circumstances
	// where the primary slice service is permanently inaccessible and you want to recover as much user
	// data as possible.
	CloneVolumeFromDead(ctx context.Context, req *CloneVolumeFromDeadRequest) (*CloneVolumeResult, *SdkError)

	// Enables all of the provided protection schemes.
	EnableProtectionSchemes(ctx context.Context, req *EnableProtectionSchemesRequest) (*EnableProtectionSchemesResult, *SdkError)

	// The SetClusterConfig API method enables you to set the configuration this node uses to communicate with the cluster it is associated with. To see the states in which these objects can be modified, see Cluster Object Attributes. To display the current cluster
	// interface settings for a node, run the GetClusterConfig API method.
	// Note: This method is available only through the per-node API endpoint 5.0 or later.
	SetClusterConfig(ctx context.Context, req *SetClusterConfigRequest) (*SetClusterConfigResult, *SdkError)

	// ListNetworkInterfaces enables you to retrieve information about each network interface on a node. The API method is intended for use on individual nodes; userid and password authentication is required for access to individual nodes.
	ListNetworkInterfaces(ctx context.Context) (*ListNetworkInterfacesResult, *SdkError)

	// The SolidFire Element OS web UI uses the ListSnapMirrorVservers method to list all SnapMirror Vservers available on a remote ONTAP system.
	ListSnapMirrorVservers(ctx context.Context, req *ListSnapMirrorVserversRequest) (*ListSnapMirrorVserversResult, *SdkError)

	// SetSnmpACL enables you to configure SNMP access permissions on the cluster nodes. The values you set with this interface apply to all
	// nodes in the cluster, and the values that are passed replace, in whole, all values set in any previous call to SetSnmpACL. Also note
	// that the values set with this interface replace all network or usmUsers values set with the older SetSnmpInfo.
	SetSnmpACL(ctx context.Context, req *SetSnmpACLRequest) (*SetSnmpACLResult, *SdkError)

	// GetNvramInfo enables you to retrieve information from each node about the NVRAM card.
	GetNvramInfo(ctx context.Context, req *GetNvramInfoRequest) (*GetNvramInfoResult, *SdkError)

	//
	CreateEvent(ctx context.Context) (*NeedsWorkCreateEventResult, *SdkError)

	//
	FillSlice(ctx context.Context) (*NeedsWorkFillSliceResult, *SdkError)

	// The GetNetworkConfig API method enables you to display the network configuration information for a node.
	// Note: This method is available only through the per-node API endpoint 5.0 or later.
	GetNetworkConfig(ctx context.Context, req *GetNetworkConfigRequest) (*GetNetworkConfigResult, *SdkError)

	// ListClusterAdmins returns the list of all cluster administrators for the cluster. There can be
	// several cluster administrator accounts with different levels of permissions.  There can be only
	// one primary cluster administrator in the system. The primary Cluster Admin is the
	// administrator that was created when the cluster was created.
	ListClusterAdmins(ctx context.Context, req *ListClusterAdminsRequest) (*ListClusterAdminsResult, *SdkError)

	// The GetConfig API method enables you to retrieve all configuration information for a node. This method includes the same information available in both the GetClusterConfig and GetNetworkConfig API methods.
	// Note: This method is available only through the per-node API endpoint 5.0 or later.
	GetConfig(ctx context.Context, req *GetConfigRequest) (*GetConfigResult, *SdkError)

	// Returns the specified KMIP (Key Management Interoperability Protocol) Key Server object.
	GetKeyServerKmip(ctx context.Context, req *GetKeyServerKmipRequest) (*GetKeyServerKmipResult, *SdkError)

	// UnbindGetVirtualVolume removes the VVol <-> Host binding.
	UnbindVirtualVolumes(ctx context.Context, req *UnbindVirtualVolumesRequest) (*VirtualVolumeUnbindResult, *SdkError)

	// GetNodeHardwareInfo enables you to return all the hardware information and status for the node specified. This generally includes details about
	// manufacturers, vendors, versions, and other associated hardware identification information.
	GetNodeHardwareInfo(ctx context.Context, req *GetNodeHardwareInfoRequest) (*GetNodeHardwareInfoResult, *SdkError)

	// You can use the StartClusterPairing method to create an encoded key from a cluster that is used to pair with another cluster. The key created from this API method is used in the CompleteClusterPairing API method to establish a cluster pairing. You can pair a cluster with a maximum of four other clusters.
	StartClusterPairing(ctx context.Context) (*StartClusterPairingResult, *SdkError)

	// QueryVirtualVolumeMetadata returns a list of VVols matching a metadata query.
	QueryVirtualVolumeMetadata(ctx context.Context, req *QueryVirtualVolumeMetadataRequest) (*QueryVirtualVolumeMetadataResult, *SdkError)

	// ListVolumeStatsByVolumeAccessGroup enables you to get total activity measurements for all of the volumes that are a member of the
	// specified volume access group(s).
	ListVolumeStatsByVolumeAccessGroup(ctx context.Context, req *ListVolumeStatsByVolumeAccessGroupRequest) (*ListVolumeStatsByVolumeAccessGroupResult, *SdkError)

	// StartBulkVolumeRead enables you to initialize a bulk volume read session on a specified volume. Only two bulk volume processes
	// can run simultaneously on a volume. When you initialize the session, data is read from a SolidFire storage volume for the purposes
	// of storing the data on an external backup source. The external data is accessed by a web server running on an SF-series node.
	// Communications and server interaction information for external data access is passed by a script running on the storage system.
	// At the start of a bulk volume read operation, a snapshot of the volume is made and the snapshot is deleted when the read is complete. You can also read a snapshot of the volume by entering the ID of the snapshot as a parameter. When you read a
	// previous snapshot, the system does not create a new snapshot of the volume or delete the previous snapshot when the
	// read completes.
	// Note: This process creates a new snapshot if the ID of an existing snapshot is not provided. Snapshots can be created if cluster fullness is at stage 2 or 3. Snapshots are not created when cluster fullness is at stage 4 or 5.
	StartBulkVolumeRead(ctx context.Context, req *StartBulkVolumeReadRequest) (*StartBulkVolumeReadResult, *SdkError)

	// GetNtpInfo enables you to return the current network time protocol (NTP) configuration information.
	GetNtpInfo(ctx context.Context) (*GetNtpInfoResult, *SdkError)

	//
	GetClusterStructure(ctx context.Context) (*NeedsWorkGetClusterStructureResult, *SdkError)

	//
	ProtocolVersionUpgrade(ctx context.Context) (*NeedsWorkProtocolVersionUpgradeResult, *SdkError)

	// ListActiveNodes returns the list of currently active nodes that are in the cluster.
	ListActiveNodes(ctx context.Context) (*ListActiveNodesResult, *SdkError)

	// The SolidFire Element OS web UI uses the InitializeSnapMirrorRelationship method to initialize the destination volume in a SnapMirror relationship by performing an initial baseline transfer between clusters.
	InitializeSnapMirrorRelationship(ctx context.Context, req *InitializeSnapMirrorRelationshipRequest) (*InitializeSnapMirrorRelationshipResult, *SdkError)

	// SecureEraseDrives enables you to remove any residual data from drives that have a status of "available." You might want to use this method when replacing a drive nearing the end of its service life that contained sensitive data. This method uses a Security Erase Unit command to write a predetermined pattern to the drive and resets the encryption key on the drive. This asynchronous method might take up to two minutes to complete. You can use GetAsyncResult to check on the status of the secure erase operation.
	// You can use the ListDrives method to obtain the driveIDs for the drives you want to secure erase.
	SecureEraseDrives(ctx context.Context, req *SecureEraseDrivesRequest) (*AsyncHandleResult, *SdkError)

	//
	ModifySliceReserveUsedThresholdPct(ctx context.Context) (*NeedsWorkModifySliceReserveUsedThresholdPctResult, *SdkError)

	// GetIpmiConfig enables you to retrieve hardware sensor information from sensors that are in your node.
	GetIpmiConfig(ctx context.Context, req *GetIpmiConfigRequest) (*GetIpmiConfigResult, *SdkError)

	// CreateSchedule enables you to schedule an automatic snapshot of a volume at a defined interval.
	// You can use the created snapshot later as a backup or rollback to ensure the data on a volume or group of volumes is consistent for
	// the point in time in which the snapshot was created.
	// If you schedule a snapshot to run at a time period that is not divisible by 5 minutes, the snapshot runs at the next time period
	// that is divisible by 5 minutes. For example, if you schedule a snapshot to run at 12:42:00 UTC, it runs at 12:45:00 UTC.
	// Note: You can create snapshots if cluster fullness is at stage 1, 2 or 3. You cannot create snapshots after cluster fullness reaches stage 4 or 5.
	CreateSchedule(ctx context.Context, req *CreateScheduleRequest) (*CreateScheduleResult, *SdkError)

	// ModifyBackupTarget enables you to change attributes of a backup target.
	ModifyBackupTarget(ctx context.Context, req *ModifyBackupTargetRequest) (*ModifyBackupTargetResult, *SdkError)

	// GetClusterVersionInfo enables you to retrieve information about the Element software version running on each node in the cluster.
	// This method also returns information about nodes that are currently in the process of upgrading software.
	GetClusterVersionInfo(ctx context.Context) (*GetClusterVersionInfoResult, *SdkError)

	// The DisableStandbySliceAssignments method instructs the slice balancer to unassign all slice standbys.
	// Callers should monitor the slices report to determine when all standbys are removed.
	// The cluster will enable IOPS-based slice balancing when standby assignments are disabled.
	DisableStandbySliceAssignments(ctx context.Context) (*DisableStandbySliceAssignmentsResult, *SdkError)

	// You can use the ModifyQoSPolicy method to modify an existing QoSPolicy on the system.
	ModifyQoSPolicy(ctx context.Context, req *ModifyQoSPolicyRequest) (*ModifyQoSPolicyResult, *SdkError)

	// The TestConnectSvip API method enables you to test the storage connection to the cluster. The test pings the SVIP using ICMP packets, and when successful, connects as an iSCSI initiator.
	// Note: This method is available only through the per-node API endpoint 5.0 or later.
	TestConnectSvip(ctx context.Context, req *TestConnectSvipRequest) (*TestConnectSvipResult, *SdkError)

	// You can use ListBackupTargets to retrieve information about all backup targets that have been created.
	ListBackupTargets(ctx context.Context) (*ListBackupTargetsResult, *SdkError)

	// ModifyInitiators enables you to change the attributes of one or more existing initiators. You cannot change the name of an existing
	// initiator. If you need to change the name of an initiator, delete it first with DeleteInitiators and create a new one with
	// CreateInitiators.
	// If ModifyInitiators fails to change one of the initiators provided in the parameter, the method returns an error and does not modify
	// any initiators (no partial completion is possible).
	ModifyInitiators(ctx context.Context, req *ModifyInitiatorsRequest) (*ModifyInitiatorsResult, *SdkError)

	//
	EnableClusterSsh(ctx context.Context) (*NeedsWorkEnableClusterSshResult, *SdkError)

	// Returns the cluster constants for a given node.
	// Node cluster constants are saved in a per-node file, and take priority over cluster-wide constants (which are saved in the database).
	GetNodeConstants(ctx context.Context, req *GetNodeConstantsRequest) (*interface{}, *SdkError)

	// ListDriveStats enables you to retrieve high-level activity measurements for multiple drives in the cluster. By default, this method returns statistics for all drives in the cluster, and these measurements are cumulative from the addition of the drive to the cluster. Some values this method returns are specific to block drives, and some are specific to metadata drives.
	ListDriveStats(ctx context.Context, req *ListDriveStatsRequest) (*ListDriveStatsResult, *SdkError)

	// ListActiveVolumes enables you to return the list of active volumes currently in the system. The list of volumes is returned sorted in
	// VolumeID order and can be returned in multiple parts (pages).
	ListActiveVolumes(ctx context.Context, req *ListActiveVolumesRequest) (*ListActiveVolumesResult, *SdkError)

	// The SolidFire Element OS web UI uses the DeleteSnapMirrorRelationships method to remove one or more SnapMirror relationships between a source and destination endpoint.
	DeleteSnapMirrorRelationships(ctx context.Context, req *DeleteSnapMirrorRelationshipsRequest) (*DeleteSnapMirrorRelationshipsResult, *SdkError)

	// ModifyVolume enables you to modify settings on an existing volume. You can make modifications to one volume at a time and
	// changes take place immediately. If you do not specify QoS values when you modify a volume, they remain the same as before the modification. You can retrieve
	// default QoS values for a newly created volume by running the GetDefaultQoS method.
	// When you need to increase the size of a volume that is being replicated, do so in the following order to prevent replication errors:
	// 1. Increase the size of the "Replication Target" volume.
	// 2. Increase the size of the source or "Read / Write" volume.
	// NetApp recommends that both the target and source volumes are the same size.
	// Note: If you change the "access" status to locked or target, all existing iSCSI connections are terminated.
	ModifyVolume(ctx context.Context, req *ModifyVolumeRequest) (*ModifyVolumeResult, *SdkError)

	// You can use the GetClusterCapacity method to return the high-level capacity measurements for an entire cluster. You can use the fields returned from this method to calculate the efficiency rates that are displayed in the Element OS Web UI. You can use the following calculations in scripts to return the efficiency rates for thin provisioning, deduplication, compression, and overall efficiency.
	GetClusterCapacity(ctx context.Context) (*GetClusterCapacityResult, *SdkError)

	// List configurations for third party Identity Provider(s) (IdP), optionally providing an IdP metadata URL to query a specific IdP configuration information.
	ListIdpConfigurations(ctx context.Context, req *ListIdpConfigurationsRequest) (*ListIdpConfigurationsResult, *SdkError)

	// You can use the GetNodeActiveTlsCiphers method to get a list of the TLS ciphers that are currently accepted on this node.
	// You can use this method on both management and storage nodes.
	GetNodeActiveTlsCiphers(ctx context.Context) (*GetNodeActiveTlsCiphersResult, *SdkError)

	// The SolidFire Element OS web UI uses the CreateSnapMirrorVolume method to create a volume on the remote ONTAP system.
	CreateSnapMirrorVolume(ctx context.Context, req *CreateSnapMirrorVolumeRequest) (*CreateSnapMirrorVolumeResult, *SdkError)

	// ListVirtualVolumeTasks returns a list of virtual volume tasks in the system.
	ListVirtualVolumeTasks(ctx context.Context, req *ListVirtualVolumeTasksRequest) (*ListVirtualVolumeTasksResult, *SdkError)

	// Creates an authentication session based on a third party Identity Provider (IdP).
	// SAML attribute statements within the SAML assertion matching multiple IdP cluster
	// admin accounts will have the combined access level of those matching IdP cluster
	// admin accounts.
	// Returns an AuthSessionInfo.
	// Intended to be used by the element-auth container.
	CreateIdpAuthSession(ctx context.Context, req *CreateIdpAuthSessionRequest) (*CreateAuthSessionResult, *SdkError)

	//
	GetGCStatus(ctx context.Context) (*NeedsWorkGetGCStatusResult, *SdkError)

	// ListPendingNodes returns a list of the currently pending nodes in the system. Pending nodes are nodes that are running and configured to join the cluster, but have not yet been added via the AddNodes API method.
	ListPendingNodes(ctx context.Context) (*ListPendingNodesResult, *SdkError)

	// DeleteStorageContainers enables you to remove up to 2000 Virtual Volume (VVol) storage containers from the system at one time.
	// The storage containers you remove must not contain any VVols.
	DeleteStorageContainers(ctx context.Context, req *DeleteStorageContainersRequest) (*DeleteStorageContainerResult, *SdkError)

	// You can use the RemoveNodeSSLCertificate method to remove the user SSL certificate and private key for the management node.
	// After the certificate and private key are removed, the management node is configured to use the default certificate and private key..
	RemoveNodeSSLCertificate(ctx context.Context) (*RemoveNodeSSLCertificateResult, *SdkError)

	// You can use the ReassignSlicesForZoneTolerance method to request a tick of the zone tolerance optimizer (balancer).
	ReassignSlicesForZoneTolerance(ctx context.Context) (*ReassignSlicesForZoneToleranceResult, *SdkError)

	//
	ListRepositories(ctx context.Context) (*NeedsWorkListRepositoriesResult, *SdkError)

	//
	MovePrimariesAwayFromNode(ctx context.Context) (*NeedsWorkMovePrimariesAwayFromNodeResult, *SdkError)

	// GetNodeStats enables you to retrieve the high-level activity measurements for a single node.
	GetNodeStats(ctx context.Context, req *GetNodeStatsRequest) (*GetNodeStatsResult, *SdkError)

	//
	GetBinAssignments(ctx context.Context) (*NeedsWorkGetBinAssignmentsResult, *SdkError)

	// ListVolumeQoSHistograms returns histograms detailing volume performance relative to QOS settings.
	// It may take up to 5 seconds for newly created volumes to have accurate histogram data available.
	ListVolumeQoSHistograms(ctx context.Context, req *ListVolumeQoSHistogramsRequest) (*ListVolumeQoSHistogramsResult, *SdkError)

	// ListVirtualVolumes enables you to list the virtual volumes currently in the system. You can use this method to list all virtual volumes,
	// or only list a subset.
	ListVirtualVolumes(ctx context.Context, req *ListVirtualVolumesRequest) (*ListVirtualVolumesResult, *SdkError)

	// GetProtectionDomainLayout returns all of the Protection Domain information for the cluster.
	GetProtectionDomainLayout(ctx context.Context) (*GetProtectionDomainLayoutResult, *SdkError)

	//
	RegenerateClusterUuid(ctx context.Context) (*NeedsWorkRegenerateClusterUuidResult, *SdkError)

	// The ResetNode API method enables you to reset a node to the factory settings. All data, packages (software upgrades, and so on),
	// configurations, and log files are deleted from the node when you call this method. However, network settings for the node are
	// preserved during this operation. Nodes that are participating in a cluster cannot be reset to the factory settings.
	// The ResetNode API can only be used on nodes that are in an "Available" state. It cannot be used on nodes that are "Active" in a
	// cluster, or in a "Pending" state.
	// Caution: This method clears any data that is on the node. Exercise caution when using this method.
	// Note: This method is available only through the per-node API endpoint 5.0 or later.
	ResetNode(ctx context.Context, req *ResetNodeRequest) (*ResetNodeResult, *SdkError)

	// GetAccountEfficiency enables you to retrieve efficiency statistics about a volume account. This method returns efficiency information
	// only for the account you specify as a parameter.
	GetAccountEfficiency(ctx context.Context, req *GetAccountEfficiencyRequest) (*GetEfficiencyResult, *SdkError)

	// Disables all of the provided protection schemes.
	DisableProtectionSchemes(ctx context.Context, req *DisableProtectionSchemesRequest) (*DisableProtectionSchemesResult, *SdkError)

	// RemoveAccount enables you to remove an existing account. You must delete and purge all volumes associated with the account
	// using DeleteVolume before you can remove the account. If volumes on the account are still pending deletion, you cannot use
	// RemoveAccount to remove the account.
	RemoveAccount(ctx context.Context, req *RemoveAccountRequest) (*RemoveAccountResult, *SdkError)

	// GetClusterInfo enables you to return configuration information about the cluster.
	GetClusterInfo(ctx context.Context) (*GetClusterInfoResult, *SdkError)

	// The SolidFire Element OS web UI uses the ListSnapMirrorLuns method to list the LUN information for the SnapMirror relationship from the remote ONTAP cluster.
	ListSnapMirrorLuns(ctx context.Context, req *ListSnapMirrorLunsRequest) (*ListSnapMirrorLunsResult, *SdkError)

	// CreateVirtualVolumeHost creates a new ESX host.
	CreateVirtualVolumeHost(ctx context.Context, req *CreateVirtualVolumeHostRequest) (*VirtualVolumeNullResult, *SdkError)

	// You can use EnableFeature to enable cluster features that are disabled by default.
	EnableFeature(ctx context.Context, req *EnableFeatureRequest) (*EnableFeatureResult, *SdkError)

	// RestoreDeletedVolume marks a deleted volume as active again. This action makes the volume immediately available for iSCSI connection.
	RestoreDeletedVolume(ctx context.Context, req *RestoreDeletedVolumeRequest) (*RestoreDeletedVolumeResult, *SdkError)

	// The GetLdapConfiguration method enables you to get the currently active LDAP configuration on the cluster.
	GetLdapConfiguration(ctx context.Context) (*GetLdapConfigurationResult, *SdkError)

	//
	CreateSnapMirrorEndpointUnmanaged(ctx context.Context) (*NeedsWorkCreateSnapMirrorEndpointUnmanagedResult, *SdkError)

	// The SolidFire Element OS web UI uses DeleteSnapMirrorEndpoints to delete one or more SnapMirror endpoints from the system.
	DeleteSnapMirrorEndpoints(ctx context.Context, req *DeleteSnapMirrorEndpointsRequest) (*DeleteSnapMirrorEndpointsResult, *SdkError)

	// ComputeVolumeChecksum calculates the checksum of a volume or snapshot.
	ComputeVolumeChecksum(ctx context.Context, req *ComputeVolumeChecksumRequest) (*ComputeVolumeChecksumResult, *SdkError)

	// GetBootstrapConfig returns cluster and node information from the bootstrap configuration file. Use this API method on an individual node before it has been joined with a cluster. You can use the information this method returns in the cluster configuration interface when you create a cluster.
	// If a cluster has already been created, this can be used to obtain the MVIP and SVIP addresses of the cluster.
	GetBootstrapConfig(ctx context.Context) (*GetBootstrapConfigResult, *SdkError)

	// ListVolumeStatsByAccount returns high-level activity measurements for every account. Values are summed from all the volumes owned by the account.
	ListVolumeStatsByAccount(ctx context.Context, req *ListVolumeStatsByAccountRequest) (*ListVolumeStatsByAccountResult, *SdkError)

	// GetVolumeCount enables you to retrieve the number of volumes currently in the system.
	GetVolumeCount(ctx context.Context) (*GetVolumeCountResult, *SdkError)

	// You can use GetAsyncResult to retrieve the result of asynchronous method calls. Some method calls require some time to run, and
	// might not be finished when the system sends the initial response. To obtain the status or result of the method call, use
	// GetAsyncResult to poll the asyncHandle value returned by the method.
	// GetAsyncResult returns the overall status of the operation (in progress, completed, or error) in a standard fashion, but the actual
	// data returned for the operation depends on the original method call and the return data is documented with each method.
	GetAsyncResult(ctx context.Context, req *GetAsyncResultRequest) (*interface{}, *SdkError)

	//
	RebalanceSlices(ctx context.Context) (*NeedsWorkRebalanceSlicesResult, *SdkError)

	// GetSystemStatus enables you to return whether a reboot ir required or not.
	GetSystemStatus(ctx context.Context) (*GetSystemStatusResult, *SdkError)

	//
	StartUpgrade(ctx context.Context) (*NeedsWorkStartUpgradeResult, *SdkError)

	// Initiate the process of setting a password on self-encrypting drives (SEDs) within the cluster.  This feature is not enabled by default but can be toggled on and off as needed.
	// If a password is set on a SED which is removed from the cluster, the password will remain set and the drive is not secure erased.  Data can be secure erased using the SecureEraseDrives API method.
	// Note: This does not affect performance or efficiency.
	// If no parameters are specified, the password will be generated internally and at random (the only option for endpoints prior to 12.0).  This generated password will be distributed across the nodes using Shamir's Secret Sharing Algorithm such that at least two nodes are required to reconstruct the password.  The complete password to unlock the drives is not stored on any single node and is never sent across the network in its entirety.  This protects against the theft of any number of drives or a single node.
	// If a keyProviderID is specified then the password will be generated/retrieved as appropriate per the type of provider.  Commonly this would be via a KMIP (Key Management Interoperability Protocol) Key Server in the case of a KMIP Key Provider (see CreateKeyProviderKmip).  After this operation the specified provider will be considered 'active' and will not be able to be deleted until DisableEncryptionAtRest is called.
	EnableEncryptionAtRest(ctx context.Context, req *EnableEncryptionAtRestRequest) (*EnableEncryptionAtRestResult, *SdkError)

	// SetDefaultQoS enables you to configure the default Quality of Service (QoS) values (measured in inputs and outputs per second, or
	// IOPS) for a volume. For more information about QoS in a SolidFire cluster, see the User Guide.
	SetDefaultQoS(ctx context.Context, req *SetDefaultQoSRequest) (*SetDefaultQoSResult, *SdkError)

	// ListNodeStats enables you to view the high-level activity measurements for all nodes in a cluster.
	ListNodeStats(ctx context.Context) (*ListNodeStatsResult, *SdkError)

	// The ListVolumes method enables you to retrieve a list of volumes that are in a cluster. You can specify the volumes you want to
	// return in the list by using the available parameters.
	ListVolumes(ctx context.Context, req *ListVolumesRequest) (*ListVolumesResult, *SdkError)

	// You can use the GetSupportedTlsCiphers method to get a list of the supported TLS ciphers on this node.
	// You can use this method on both management and storage nodes.
	GetNodeSupportedTlsCiphers(ctx context.Context) (*GetNodeSupportedTlsCiphersResult, *SdkError)

	// You can use the ResetNodeSupplementalTlsCiphers method to restore the supplemental ciphers to their defaults.
	// You can use this command on management nodes.
	ResetNodeSupplementalTlsCiphers(ctx context.Context) (*ResetNodeSupplementalTlsCiphersResult, *SdkError)

	// StartVolumePairing enables you to create an encoded key from a volume that is used to pair with another volume. The key that this
	// method creates is used in the CompleteVolumePairing API method to establish a volume pairing.
	StartVolumePairing(ctx context.Context, req *StartVolumePairingRequest) (*StartVolumePairingResult, *SdkError)

	// One can use this API to remove a local cluster admin, an LDAP cluster admin, or a third
	// party Identity Provider (IdP) cluster admin.
	// One cannot remove the administrator cluster admin account.
	// When an IdP Admin is removed that has authenticated sessions associated with a third party
	// Identity Provider (IdP), those sessions will either logout or possibly experience a loss of
	// access rights within their current session.  The access rights loss will depend on whether the
	// removed IdP cluster admin matched one of multiple IdP cluster admins from a given user's
	// SAML Attributes and the remaining set of matching IdP cluster admins results in a reduced
	// set of aggregate access rights.
	// Other cluster admin user types will be logged out upon their cluster admin removal.
	RemoveClusterAdmin(ctx context.Context, req *RemoveClusterAdminRequest) (*RemoveClusterAdminResult, *SdkError)

	// ResetDrives enables you to proactively initialize drives and remove all data currently residing on a drive. The drive can then be reused
	// in an existing node or used in an upgraded node. This method requires the force parameter to be included in the method call.
	ResetDrives(ctx context.Context, req *ResetDrivesRequest) (*ResetDrivesResult, *SdkError)

	//
	SetAptSourceLines(ctx context.Context) (*NeedsWorkSetAptSourceLinesResult, *SdkError)

	// The SolidFire Element OS web UI uses the ListSnapMirrorRelationships method to list one or all SnapMirror relationships on a SolidFire cluster
	ListSnapMirrorRelationships(ctx context.Context, req *ListSnapMirrorRelationshipsRequest) (*ListSnapMirrorRelationshipsResult, *SdkError)

	// DeleteVolumes marks multiple (up to 500) active volumes for deletion.
	// Once marked, the volumes are purged (permanently deleted) after the cleanup interval elapses.
	// The cleanup interval can be set in the SetClusterSettings method.
	// For more information on using this method, see SetClusterSettings on page 1.
	// After making a request to delete volumes, any active iSCSI connections to the volumes are immediately terminated
	// and no further connections are allowed while the volumes are in this state.
	// A marked volume is not returned in target discovery requests.
	// Any snapshots of a volume that has been marked for deletion are not affected.
	// Snapshots are kept until the volume is purged from the system.
	// If a volume is marked for deletion and has a bulk volume read or bulk volume write operation in progress,
	// the bulk volume read or write operation is stopped.
	// If the volumes you delete are paired with a volume, replication between the paired volumes is suspended
	// and no data is transferred to them or from them while in a deleted state.
	// The remote volumes the deleted volumes were paired with enter into a PausedMisconfigured state
	// and data is no longer sent to them or from the deleted volumes.
	// Until the deleted volumes are purged, they can be restored and data transfers resume.
	// If the deleted volumes are purged from the system, the volumes they were paired with enter into a
	// StoppedMisconfigured state and the volume pairing status is removed.
	// The purged volumes become permanently unavailable.
	DeleteVolumes(ctx context.Context, req *DeleteVolumesRequest) (*DeleteVolumesResult, *SdkError)

	// GetDefaultQoS enables you to retrieve the default QoS values for a newly created volume.
	GetDefaultQoS(ctx context.Context) (*VolumeQOS, *SdkError)

	// Lists all expired auth sessions.
	// This is only callable by a user with Administrative access rights.
	ListExpiredAuthSessions(ctx context.Context) (*ListAuthSessionsResult, *SdkError)

	//
	GetSliceReserveUsedThresholdPct(ctx context.Context) (*NeedsWorkGetSliceReserveUsedThresholdPctResult, *SdkError)

	// RemoveVirtualNetwork enables you to remove a previously added virtual network.
	// Note: This method requires either the virtualNetworkID or the virtualNetworkTag as a parameter, but not both.
	RemoveVirtualNetwork(ctx context.Context, req *RemoveVirtualNetworkRequest) (*RemoveVirtualNetworkResult, *SdkError)

	// The DeleteStandbySlices method deactivates any active or failed standby for the given slices and deletes
	// the associated metadata. The command is a no-op if the standby is already inactive.
	DeleteStandbySlices(ctx context.Context, req *DeleteStandbySlicesRequest) (*DeleteStandbySlicesResult, *SdkError)

	//
	ModifySnapMirrorEndpointUnmanaged(ctx context.Context) (*NeedsWorkModifySnapMirrorEndpointUnmanagedResult, *SdkError)

	// You can use the ListTests API method to return the tests that are available to run on a node.
	// Note: This method is available only through the per-node API endpoint 5.0 or later.
	ListTests(ctx context.Context, req *ListTestsRequest) (*ListTestsResult, *SdkError)

	// GetVirtualVolumeAllocatedBitmap returns a b64-encoded block of data
	// representing a bitmap where non-zero bits indicate that data is not the same
	// between two volumes for a common segment (LBA range) of the volumes.
	GetVirtualVolumeUnsharedBitmap(ctx context.Context, req *GetVirtualVolumeUnsharedBitmapRequest) (*VirtualVolumeBitmapResult, *SdkError)

	// Return information regarding the state of authentication using third party Identity Providers
	GetIdpAuthenticationState(ctx context.Context) (*GetIdpAuthenticationStateResult, *SdkError)

	// You can use ListServices to return the services information for nodes, drives, current software, and other services that are running on the cluster.
	ListServices(ctx context.Context) (*ListServicesResult, *SdkError)

	// RollbackToGroupSnapshot enables you to roll back all individual volumes in a snapshot group to each volume's individual snapshot.
	// Note: Rolling back to a group snapshot creates a temporary snapshot of each volume within the group snapshot.
	// Snapshots are allowed if cluster fullness is at stage 2 or 3. Snapshots are not created when cluster fullness is at stage 4 or 5.
	RollbackToGroupSnapshot(ctx context.Context, req *RollbackToGroupSnapshotRequest) (*RollbackToGroupSnapshotResult, *SdkError)

	// ListCurrentClusterAdmins returns information for all the cluster admins the calling user is assigned to.
	// Intended to be called by element-auth container.
	ListCurrentClusterAdmins(ctx context.Context) (*ListCurrentClusterAdminsResult, *SdkError)

	// You can use the ResetSupplementalTlsCiphers method to restore the supplemental ciphers to their defaults.
	ResetSupplementalTlsCiphers(ctx context.Context) (*ResetSupplementalTlsCiphersResult, *SdkError)

	// The EnableStandbySliceAssignments method instructs the slice balancer to assign a standby to each slice.
	// Callers should call CheckVolumeStandbysAssigned or monitor the slices report to determine when
	// the assignments are complete.
	// The cluster will disable IOPS-based slice balancing while standby assignments are enabled.
	EnableStandbySliceAssignments(ctx context.Context) (*EnableStandbySliceAssignmentsResult, *SdkError)

	// The SolidFire Element OS web UI uses the BreakSnapMirrorRelationship method to break a SnapMirror relationship. When a SnapMirror relationship is broken, the destination volume is made read-write and independent, and can then diverge from the source. You can reestablish the relationship with the ResyncSnapMirrorRelationship API method. This method requires the ONTAP cluster to be available.
	BreakSnapMirrorRelationship(ctx context.Context, req *BreakSnapMirrorRelationshipRequest) (*BreakSnapMirrorRelationshipResult, *SdkError)

	// Update the Vasa Provider info
	ModifyVasaProviderInfo(ctx context.Context, req *ModifyVasaProviderInfoRequest) (*VirtualVolumeNullResult, *SdkError)

	//
	DeleteVirtualVolumeHosts(ctx context.Context) (*NeedsWorkDeleteVirtualVolumeHostsResult, *SdkError)

	// Adds (assigns) the specified KMIP (Key Management Interoperability Protocol) Key Server to the specified Key Provider.  This will result in contacting the server to verify it's functional, as well as to synchronize keys in the event that there are multiple key servers assigned to the provider.  This synchronization may result in conflicts which could cause this to fail.  If the specified KMIP Key Server is already assigned to the specified Key Provider, this is a no-op and no error will be returned.  The assignment can be removed (unassigned) using RemoveKeyServerFromProviderKmip.
	AddKeyServerToProviderKmip(ctx context.Context, req *AddKeyServerToProviderKmipRequest) (*AddKeyServerToProviderKmipResult, *SdkError)

	// The SolidFire Element OS web UI uses the ListSnapMirrorNetworkInterfaces method to list all available SnapMirror interfaces on a remote ONTAP system
	ListSnapMirrorNetworkInterfaces(ctx context.Context, req *ListSnapMirrorNetworkInterfacesRequest) (*ListSnapMirrorNetworkInterfacesResult, *SdkError)

	// ListVolumeAccessGroups enables you to return
	// information about the volume access groups that are
	// currently in the system.
	ListVolumeAccessGroups(ctx context.Context, req *ListVolumeAccessGroupsRequest) (*ListVolumeAccessGroupsResult, *SdkError)

	// Deletes all auth sessions associated with the specified ClusterAdminID.
	// If the specified ClusterAdminID maps to a group of users, all auth sessions for all members of that group will be deleted.
	// To see the list of sessions that could be deleted, use ListAuthSessionsByClusterAdmin with the same parameter.
	DeleteAuthSessionsByClusterAdmin(ctx context.Context, req *DeleteAuthSessionsByClusterAdminRequest) (*DeleteAuthSessionsResult, *SdkError)

	// The SolidFire Element OS web UI uses the UpdateSnapMirrorRelationship method to make the destination volume in a SnapMirror relationship an up-to-date mirror of the source volume.
	UpdateSnapMirrorRelationship(ctx context.Context, req *UpdateSnapMirrorRelationshipRequest) (*UpdateSnapMirrorRelationshipResult, *SdkError)

	// GetVolumeEfficiency enables you to retrieve information about a volume. Only the volume you give as a parameter in this API method is used to compute the capacity.
	GetVolumeEfficiency(ctx context.Context, req *GetVolumeEfficiencyRequest) (*GetVolumeEfficiencyResult, *SdkError)

	// PurgeDeletedVolumes immediately and permanently purges volumes that have been deleted.
	// You can use this method to purge up to 500 volumes at one time.
	// You must delete volumes using DeleteVolumes before they can be purged.
	// Volumes are purged by the system automatically after a period of time, so usage of this method is not typically required.
	PurgeDeletedVolumes(ctx context.Context, req *PurgeDeletedVolumesRequest) (*PurgeDeletedVolumesResult, *SdkError)

	// Returns the cluster constants for the cluster.
	GetConstants(ctx context.Context, req *GetConstantsRequest) (*interface{}, *SdkError)

	//
	ReassignSlice(ctx context.Context) (*NeedsWorkReassignSliceResult, *SdkError)

	//
	SetInitiatorID(ctx context.Context) (*NeedsWorkSetInitiatorIDResult, *SdkError)

	// The SolidFire Element OS web UI uses the BreakSnapMirrorVolume method to break the SnapMirror relationship between an ONTAP source container and SolidFire target volume. Breaking a SolidFire SnapMirror volume is useful if an ONTAP system becomes unavailable while replicating data to a SolidFire volume. This feature enables a storage administrator to take control of a SolidFire SnapMirror volume, break its relationship with the remote ONTAP system, and revert the volume to a previous snapshot.
	BreakSnapMirrorVolume(ctx context.Context, req *BreakSnapMirrorVolumeRequest) (*BreakSnapMirrorVolumeResult, *SdkError)

	// RemoveInitiatorsFromVolumeAccessGroup enables
	// you to remove initiators from a specified volume access
	// group.
	RemoveInitiatorsFromVolumeAccessGroup(ctx context.Context, req *RemoveInitiatorsFromVolumeAccessGroupRequest) (*ModifyVolumeAccessGroupResult, *SdkError)

	// CreateMultipleVolumes enables you to create new (empty) volumes on the cluster.
	// As soon as creation is complete, the volumes are available for connection via iSCSI.
	CreateMultipleVolumes(ctx context.Context, req *CreateMultipleVolumesRequest) (*CreateMultipleVolumesResult, *SdkError)

	// CheckProposedCluster validates that creating a cluster from a given set of nodes is likely to succeed.  Any problems with the proposed cluster are returned as errors with a human-readable description and unique error code.
	CheckProposedCluster(ctx context.Context, req *CheckProposedClusterRequest) (*CheckProposedResult, *SdkError)

	// You can use the RemoveClusterPair method to close the open connections between two paired clusters.
	// Note: Before you remove a cluster pair, you must first remove all volume pairing to the clusters with the "RemoveVolumePair" API method.
	RemoveClusterPair(ctx context.Context, req *RemoveClusterPairRequest) (*RemoveClusterPairResult, *SdkError)

	// DeleteVirtualVolume marks an active volume for deletion.
	// It is purged (permanently deleted) after the cleanup interval elapses.
	// After making a request to delete a volume, any active iSCSI connections to the volume is immediately terminated and no further connections are allowed while the volume is in this state.
	// It is not returned in target discovery requests.
	//
	// Any snapshots of a volume that has been marked to delete are not affected.
	// Snapshots are kept until the volume is purged from the system.
	//
	// If a volume is marked for deletion, and it has a bulk volume read or bulk volume write operation in progress, the bulk volume operation is stopped.
	//
	// If the volume you delete is paired with a volume, replication between the paired volumes is suspended and no data is transferred to it or from it while in a deleted state.
	// The remote volume the deleted volume was paired with enters into a PausedMisconfigured state and data is no integerer sent to it or from the deleted volume.
	// Until the deleted volume is purged, it can be restored and data transfers resumes.
	// If the deleted volume gets purged from the system, the volume it was paired with enters into a StoppedMisconfigured state and the volume pairing status is removed.
	// The purged volume becomes permanently unavailable.
	DeleteVirtualVolumes(ctx context.Context, req *DeleteVirtualVolumesRequest) (*VirtualVolumeNullResult, *SdkError)

	// ListClusterFaults enables you to retrieve information about any faults detected on the cluster. With this method, you can retrieve both current faults as well as faults that have been resolved. The system caches faults every 30 seconds.
	ListClusterFaults(ctx context.Context, req *ListClusterFaultsRequest) (*ListClusterFaultsResult, *SdkError)

	// Sets LLDP configuration options. If an option isn't set in the request, then it is unchanged from the previous value.
	SetLldpConfig(ctx context.Context, req *SetLldpConfigRequest) (*GetLldpConfigResult, *SdkError)

	//
	GetThreadBacktraces(ctx context.Context) (*NeedsWorkGetThreadBacktracesResult, *SdkError)

	// GetVolumeAccessGroupEfficiency enables you to
	// retrieve efficiency information about a volume access
	// group. Only the volume access group you provide as the
	// parameter in this API method is used to compute the
	// capacity.
	GetVolumeAccessGroupEfficiency(ctx context.Context, req *GetVolumeAccessGroupEfficiencyRequest) (*GetEfficiencyResult, *SdkError)

	// RemoveNodes can be used to remove one or more nodes from the cluster. Before removing a node, you must remove all drives from the node using the RemoveDrives method. You cannot remove a node until the RemoveDrives process has completed and all data has been migrated off of the node's drives.
	// After removing a node, the removed node registers itself as a pending node. You can add the pending node again or shut it down (shutting the node down removes it from the Pending Node list).
	//
	// RemoveNodes can fail with xEnsembleInvalidSize if removing the nodes would violate ensemble size restrictions.
	// RemoveNodes can fail with xEnsembleQuorumLoss if removing the nodes would cause a loss of quorum.
	// RemoveNodes can fail with xEnsembleToleranceChange if there are enabled data protection schemes that can tolerate multiple node failures and removing the nodes would decrease the ensemble's node failure tolerance. The optional ignoreEnsembleToleranceChange parameter can be set true to disable the ensemble tolerance check.
	RemoveNodes(ctx context.Context, req *RemoveNodesRequest) (*RemoveNodesResult, *SdkError)

	// The SolidFire Element OS web UI uses the ResyncSnapMirrorRelationship method to establish or reestablish a mirror relationship between a source and destination endpoint. When you resync a relationship, the system removes snapshots on the destination volume that are newer than the common snapshot copy, and then mounts the destination volume as a data protection volume with the common snapshot copy as the exported snapshot copy.
	ResyncSnapMirrorRelationship(ctx context.Context, req *ResyncSnapMirrorRelationshipRequest) (*ResyncSnapMirrorRelationshipResult, *SdkError)

	// ListStorageContainers enables you to retrieve information about all virtual volume storage containers known to the system.
	ListStorageContainers(ctx context.Context, req *ListStorageContainersRequest) (*ListStorageContainersResult, *SdkError)

	// DeleteVolume marks an active volume for deletion. When marked, the volume is purged (permanently deleted) after the cleanup
	// interval elapses. After making a request to delete a volume, any active iSCSI connections to the volume are immediately terminated
	// and no further connections are allowed while the volume is in this state. A marked volume is not returned in target discovery
	// requests.
	// Any snapshots of a volume that has been marked for deletion are not affected. Snapshots are kept until the volume is purged from
	// the system.
	// If a volume is marked for deletion and has a bulk volume read or bulk volume write operation in progress, the bulk volume read or
	// write operation is stopped.
	// If the volume you delete is paired with a volume, replication between the paired volumes is suspended and no data is transferred
	// to it or from it while in a deleted state. The remote volume that the deleted volume was paired with enters into a PausedMisconfigured state and data is no longer sent to it or from the deleted volume. Until the deleted volume is purged, it can be restored and data transfers resume. If the deleted volume gets purged from the system, the volume it was paired with enters into a StoppedMisconfigured state and the volume pairing status is removed. The purged volume becomes permanently unavailable.
	DeleteVolume(ctx context.Context, req *DeleteVolumeRequest) (*DeleteVolumeResult, *SdkError)

	// Used to assign Nodes to user-defined Chassis Protection Domains for test purposes.  Overrides existing
	// ChassisNames. This information must be provided for all Active Nodes in the cluster, and no information may
	// be provided for Nodes that are not Active. The same ProtectionDomainType must be supplied for all nodes.
	// ProtectionDomainTypes other than Chassis must not be included. If any of these are not true, the Chassis
	// Protection Domains will be ignored, and an appropriate error will be returned.
	SetProtectionDomainLayoutChassisOverride(ctx context.Context, req *SetProtectionDomainLayoutChassisOverrideRequest) (*SetProtectionDomainLayoutChassisOverrideResult, *SdkError)

	//
	CreateClusterFault(ctx context.Context) (*NeedsWorkCreateClusterFaultResult, *SdkError)

	// ModifyGroupSnapshot enables you to change the attributes of a group of snapshots. You can also use this method to enable snapshots created on the Read/Write (source) volume to be remotely replicated to a target SolidFire storage system.
	ModifyGroupSnapshot(ctx context.Context, req *ModifyGroupSnapshotRequest) (*ModifyGroupSnapshotResult, *SdkError)

	// The GetClusterState API method enables you to indicate if a node is part of a cluster or not. The three states are:
	// Available: Node has not been configured with a cluster name.
	// Pending: Node is pending for a specific named cluster and can be added.
	// Active: Node is an active member of a cluster and may not be added to another
	// cluster.
	// Note: This method is available only through the per-node API endpoint 5.0 or later.
	GetClusterState(ctx context.Context, req *GetClusterStateRequest) (*GetClusterStateResult, *SdkError)

	//
	GetVolumeSetEfficiency(ctx context.Context) (*NeedsWorkGetVolumeSetEfficiencyResult, *SdkError)

	//
	ListCloneJobs(ctx context.Context) (*NeedsWorkListCloneJobsResult, *SdkError)

	//
	SetRsyslogInfo(ctx context.Context) (*NeedsWorkSetRsyslogInfoResult, *SdkError)

	// Promotes a new node to be cluster master.
	// The original cluster master will return success if it initiates the change. Callers should confirm that the
	// change occurs by calling the API command GetClusterMasterNodeID after a delay of a few seconds.
	PromoteClusterMaster(ctx context.Context, req *PromoteClusterMasterRequest) (*PromoteClusterMasterResult, *SdkError)

	//
	CreateDatabaseEntry(ctx context.Context) (*NeedsWorkCreateDatabaseEntryResult, *SdkError)

	// CancelVirtualVolumeTask attempts to cancel the VVol Async Task.
	CancelVirtualVolumeTask(ctx context.Context, req *CancelVirtualVolumeTaskRequest) (*VirtualVolumeNullResult, *SdkError)

	//
	GetReport(ctx context.Context) (*NeedsWorkGetReportResult, *SdkError)

	// You can use the GetQoSPolicy method to get details about a specific QoSPolicy from the system.
	GetQoSPolicy(ctx context.Context, req *GetQoSPolicyRequest) (*GetQoSPolicyResult, *SdkError)

	// NetApp engineering uses the GetRawStats API method to troubleshoot new features. The data returned from GetRawStats is not documented, changes frequently, and is not guaranteed to be accurate. NetApp does not recommend using GetCompleteStats for collecting performance data or any other
	// management integration with a SolidFire cluster.
	GetRawStats(ctx context.Context) (*interface{}, *SdkError)

	// Update an existing configuration with a third party Identity Provider (IdP) for the cluster.
	UpdateIdpConfiguration(ctx context.Context, req *UpdateIdpConfigurationRequest) (*UpdateIdpConfigurationResult, *SdkError)

	// The TestConnectMvip API method enables you to test the
	// management connection to the cluster. The test pings the MVIP and executes a simple API method to verify connectivity.
	// Note: This method is available only through the per-node API endpoint 5.0 or later.
	TestConnectMvip(ctx context.Context, req *TestConnectMvipRequest) (*TestConnectMvipResult, *SdkError)

	// GetVirtualVolumeAllocatedBitmap returns a b64-encoded block of data
	// representing a bitmap where non-zero bits indicate the allocation of a
	// segment (LBA range) of the volume.
	GetVirtualVolumeAllocatedBitmap(ctx context.Context, req *GetVirtualVolumeAllocatedBitmapRequest) (*VirtualVolumeBitmapResult, *SdkError)

	// Creates SSL public and private keys. These keys can be used to generate Certificate Sign Requests.
	// There can be only one key pair in use for the cluster. To replace the existing keys, make sure that they are not being used by any providers before invoking this API.
	CreatePublicPrivateKeyPair(ctx context.Context, req *CreatePublicPrivateKeyPairRequest) (*CreatePublicPrivateKeyPairResult, *SdkError)

	//
	PurgeAllDeletedVolumes(ctx context.Context) (*NeedsWorkPurgeAllDeletedVolumesResult, *SdkError)

	// Returns the list of KMIP (Key Management Interoperability Protocol) Key Servers which have been created via CreateKeyServerKmip.  The list can optionally be filtered by specifying additional parameters.
	ListKeyServersKmip(ctx context.Context, req *ListKeyServersKmipRequest) (*ListKeyServersKmipResult, *SdkError)

	// The SolidFire Element OS web UI uses the QuiesceSnapMirrorRelationship method to disable future data transfers for a SnapMirror relationship. If a transfer is in progress, the relationship status becomes "quiescing" until the transfer is complete. If the current transfer is aborted, it will not restart. You can reenable data transfers for the relationship using the ResumeSnapMirrorRelationship API method.
	QuiesceSnapMirrorRelationship(ctx context.Context, req *QuiesceSnapMirrorRelationshipRequest) (*QuiesceSnapMirrorRelationshipResult, *SdkError)

	// NetApp engineering uses the GetCompleteStats API method to troubleshoot new features. The data returned from GetCompleteStats is not documented, changes frequently, and is not guaranteed to be accurate. NetApp does not recommend using GetCompleteStats for collecting performance data or any other
	// management integration with a SolidFire cluster.
	GetCompleteStats(ctx context.Context) (*interface{}, *SdkError)

	//
	PairCluster(ctx context.Context) (*NeedsWorkPairClusterResult, *SdkError)

	// ListGroupSnapshots enables you to get information about all group snapshots that have been created.
	ListGroupSnapshots(ctx context.Context, req *ListGroupSnapshotsRequest) (*ListGroupSnapshotsResult, *SdkError)

	// SetSnmpInfo enables you to configure SNMP version 2 and version 3 on cluster nodes. The values you set with this interface apply to
	// all nodes in the cluster, and the values that are passed replace, in whole, all values set in any previous call to SetSnmpInfo.
	// Note: SetSnmpInfo is deprecated. Use the EnableSnmp and SetSnmpACL methods instead.
	SetSnmpInfo(ctx context.Context, req *SetSnmpInfoRequest) (*SetSnmpInfoResult, *SdkError)

	// GetRemoteLoggingHosts enables you to retrieve the current list of log servers.
	GetRemoteLoggingHosts(ctx context.Context) (*GetRemoteLoggingHostsResult, *SdkError)

	//
	GetSliceFileSizeReport(ctx context.Context) (*NeedsWorkGetSliceFileSizeReportResult, *SdkError)

	// GetSnmpInfo enables you to retrieve the current simple network management protocol (SNMP) configuration information.
	// Note: GetSnmpInfo is available for Element OS 8 and prior releases. It is deprecated for versions later than Element OS 8.
	// NetApp recommends that you migrate to the GetSnmpState and SetSnmpACL methods. See details in the Element API Reference Guide
	// for their descriptions and usage.
	GetSnmpInfo(ctx context.Context) (*GetSnmpInfoResult, *SdkError)

	//
	ListAptSourceLines(ctx context.Context) (*NeedsWorkListAptSourceLinesResult, *SdkError)

	//
	PairVolume(ctx context.Context) (*NeedsWorkPairVolumeResult, *SdkError)

	// The ListNodeFibreChannelPortInfo API method enables you to retrieve information about the Fibre Channel ports on a node. The API method is intended for use on individual nodes; userid and password authentication is required for access to individual Fibre Channel nodes.
	ListNodeFibreChannelPortInfo(ctx context.Context) (*ListNodeFibreChannelPortInfoResult, *SdkError)

	// ModifyVirtualVolumeMetadata is used to selectively modify the VVol metadata.
	ModifyVirtualVolumeMetadata(ctx context.Context, req *ModifyVirtualVolumeMetadataRequest) (*VirtualVolumeNullResult, *SdkError)

	// The CreateCluster method enables you to initialize the node in a cluster that has ownership of the "mvip" and "svip" addresses. Each new cluster is initialized using the management IP (MIP) of the first node in the cluster. This method also automatically adds all the nodes being configured into the cluster. You only need to use this method once each time a new cluster is initialized.
	// Note: You need to log in to the node that is used as the master node for the cluster. After you log in, run the GetBootstrapConfig method on the node to get the IP addresses for the rest of the nodes that you want to include in the
	// cluster. Then, run the CreateCluster method.
	CreateCluster(ctx context.Context, req *CreateClusterRequest) (*CreateClusterResult, *SdkError)

	// The GetClusterConfig API method enables you to return information about the cluster configuration this node uses to communicate with the cluster that it is a part of.
	GetClusterConfig(ctx context.Context, req *GetClusterConfigRequest) (*GetClusterConfigResult, *SdkError)

	// ListProtectionDomainLevels returns the Tolerance and Resiliency of the cluster from the perspective
	// of each of the supported ProtectionDomainTypes.
	ListProtectionDomainLevels(ctx context.Context) (*ListProtectionDomainLevelsResult, *SdkError)

	// You can use the DeleteQoSPolicy method to delete a QoS policy from the system.
	// The QoS settings for all volumes created of modified with this policy are unaffected.
	DeleteQoSPolicy(ctx context.Context, req *DeleteQoSPolicyRequest) (*DeleteQoSPolicyResult, *SdkError)

	//
	DisableBmcColdReset(ctx context.Context) (*NeedsWorkDisableBmcColdResetResult, *SdkError)

	// CreateGroupSnapshot enables you to create a point-in-time copy of a group of volumes. You can use this snapshot later as a backup or rollback to ensure the data on the group of volumes is consistent for the point in time that you created the snapshot.
	// Note: Creating a group snapshot is allowed if cluster fullness is at stage 2 or 3. Snapshots are not created when cluster fullness is at stage 4 or 5.
	CreateGroupSnapshot(ctx context.Context, req *CreateGroupSnapshotRequest) (*CreateGroupSnapshotResult, *SdkError)

	// The GetVolumeAccessGroupLunAssignments
	// method enables you to retrieve details on LUN mappings
	// of a specified volume access group.
	GetVolumeAccessGroupLunAssignments(ctx context.Context, req *GetVolumeAccessGroupLunAssignmentsRequest) (*GetVolumeAccessGroupLunAssignmentsResult, *SdkError)

	// You can use the AddBlockToBS method to store a block and associated blockID to the block services.
	AddBlockToBS(ctx context.Context, req *AddBlockToBSRequest) (*AddBlockToBSResult, *SdkError)

	// SetRemoteLoggingHosts enables you to configure remote logging from the nodes in the storage cluster to a centralized log server or servers. Remote logging is performed over TCP using the default port 514. This API does not add to the existing logging hosts. Rather, it replaces what currently exists with new values specified by this API method. You can use GetRemoteLoggingHosts to determine what the current logging hosts are, and then use SetRemoteLoggingHosts to set the desired list of current and new logging hosts.
	SetRemoteLoggingHosts(ctx context.Context, req *SetRemoteLoggingHostsRequest) (*SetRemoteLoggingHostsResult, *SdkError)

	// The CheckVolumeStandbysAssigned method enables a maintenance application to determine when standbys have been
	// assigned to all volumes managed by a node. This should be sent after standby assignments have been enabled
	// but before the node is taken offline, such as for an upgrade.
	CheckVolumeStandbysAssigned(ctx context.Context, req *CheckVolumeStandbysAssignedRequest) (*CheckVolumeStandbysAssignedResult, *SdkError)

	// List all auth sessions associated with the specified ClusterAdminID.
	// If the specified ClusterAdminID maps to a group of users, all auth sessions for all members of that group will be listed.
	ListAuthSessionsByClusterAdmin(ctx context.Context, req *ListAuthSessionsByClusterAdminRequest) (*ListAuthSessionsResult, *SdkError)

	// This allows the auth container to store configuration data in the cluster
	// database.  The configuration data itself is expected to be in text form.
	// Subsequent calls to this method will replace data stored by earlier calls.
	//
	// In order to guarantee that the auth container has the current revision of the configuration,
	// a version is required to be included which should correspond to the current configuration version
	// in the database.  On success, the version in the database will be the version provided plus one.
	// GetAuthConfiguration may be used to retrieve the current configuration and its version.
	SetAuthConfiguration(ctx context.Context, req *SetAuthConfigurationRequest) (*SetAuthConfigurationResult, *SdkError)

	// CreateInitiators enables you to create multiple new initiator IQNs or World Wide Port Names (WWPNs) and optionally assign them
	// aliases and attributes. When you use CreateInitiators to create new initiators, you can also add them to volume access groups.
	// If CreateInitiators fails to create one of the initiators provided in the parameter, the method returns an error and does not create
	// any initiators (no partial completion is possible).
	CreateInitiators(ctx context.Context, req *CreateInitiatorsRequest) (*CreateInitiatorsResult, *SdkError)

	//
	ReconfigureEnsemble(ctx context.Context) (*NeedsWorkReconfigureEnsembleResult, *SdkError)

	// You can use the GetSupportedTlsCiphers method to get a list of the supported TLS ciphers.
	GetSupportedTlsCiphers(ctx context.Context) (*GetSupportedTlsCiphersResult, *SdkError)

	// GetDriveConfig enables you to display drive information for expected slice and block drive counts as well as the number of slices
	// and block drives that are currently connected to the node.
	// Note: This method is available only through the per-node API endpoint 5.0 or later.
	GetDriveConfig(ctx context.Context, req *GetDriveConfigRequest) (*GetDriveConfigResult, *SdkError)

	// You can use the SetSSLCertificate method to set a user SSL certificate and a private key for the cluster.
	SetSSLCertificate(ctx context.Context, req *SetSSLCertificateRequest) (*SetSSLCertificateResult, *SdkError)

	// Create a potential trust relationship for authentication using a third party Identity Provider (IdP) for the cluster.
	// A SAML Service Provider certificate is required for IdP communication, which will be generated as necessary.
	CreateIdpConfiguration(ctx context.Context, req *CreateIdpConfigurationRequest) (*CreateIdpConfigurationResult, *SdkError)

	//
	SetUpgradeNodeId(ctx context.Context) (*NeedsWorkSetUpgradeNodeIdResult, *SdkError)

	// ListVirtualVolumeBindings returns a list of all virtual volumes in the cluster that are bound to protocol endpoints.
	ListVirtualVolumeBindings(ctx context.Context, req *ListVirtualVolumeBindingsRequest) (*ListVirtualVolumeBindingsResult, *SdkError)

	// ModifyVolumes allows you to configure up to 500 existing volumes at one time. Changes take place immediately.
	// If ModifyVolumes fails to modify any of the specified volumes, none of the specified volumes are changed.
	// If you do not specify QoS values when you modify volumes, the QoS values for each volume remain unchanged.
	// You can retrieve default QoS values for a newly created volume by running the GetDefaultQoS method.
	// When you need to increase the size of volumes that are being replicated, do so in the following order
	// to prevent replication errors:
	//    Increase the size of the "Replication Target" volume.
	//    Increase the size of the source or "Read / Write" volume.
	// Recommend that both the target and source volumes be the same size.
	// NOTE: If you change access status to locked or replicationTarget all existing iSCSI connections are terminated.
	ModifyVolumes(ctx context.Context, req *ModifyVolumesRequest) (*ModifyVolumesResult, *SdkError)

	//
	DeleteSnapMirrorObjectAttributes(ctx context.Context) (*NeedsWorkDeleteSnapMirrorObjectAttributesResult, *SdkError)

	//
	SetClusterSettings(ctx context.Context) (*NeedsWorkSetClusterSettingsResult, *SdkError)

	// You can use the CompleteClusterPairing method with the encoded key received from the  StartClusterPairing method to complete the cluster pairing process. The CompleteClusterPairing method is the second step in the cluster pairing process.
	CompleteClusterPairing(ctx context.Context, req *CompleteClusterPairingRequest) (*CompleteClusterPairingResult, *SdkError)

	// You can use the CreateQoSPolicy method to create a QoSPolicy object that you can later apply to a volume upon creation or modification. A QoS policy has a unique ID, a name, and QoS settings.
	CreateQoSPolicy(ctx context.Context, req *CreateQoSPolicyRequest) (*CreateQoSPolicyResult, *SdkError)

	// You can use RemoveDrives to proactively remove drives that are part of the cluster. You might want to use this method when
	// reducing cluster capacity or preparing to replace drives nearing the end of their service life. Any data on the drives is removed and
	// migrated to other drives in the cluster before the drive is removed from the cluster. This is an asynchronous method. Depending on
	// the total capacity of the drives being removed, it might take several minutes to migrate all of the data. Use the GetAsyncResult
	// method to check the status of the remove operation.
	// When removing multiple drives, use a single RemoveDrives method call rather than multiple individual methods with a single drive
	// each. This reduces the amount of data balancing that must occur to even stabilize the storage load on the cluster.
	// You can also remove drives with a "failed" status using RemoveDrives. When you remove a drive with a "failed" status it is not
	// returned to an "available" or active status. The drive is unavailable for use in the cluster.
	// Use the ListDrives method to obtain the driveIDs for the drives you want to remove.
	RemoveDrives(ctx context.Context, req *RemoveDrivesRequest) (*AsyncHandleResult, *SdkError)

	// GetHardwareConfig enables you to display the hardware configuration information for a node.
	// Note: This method is available only through the per-node API endpoint 5.0 or later.
	GetHardwareConfig(ctx context.Context, req *GetHardwareConfigRequest) (*GetHardwareConfigResult, *SdkError)

	//
	DisconnectISCSISessions(ctx context.Context) (*NeedsWorkDisconnectISCSISessionsResult, *SdkError)

	// ListVolumeStats returns high-level activity measurements for a single volume, list of volumes, or all volumes (if you omit the volumeIDs parameter). Measurement values are cumulative from the creation of the volume.
	ListVolumeStats(ctx context.Context, req *ListVolumeStatsRequest) (*ListVolumeStatsResult, *SdkError)

	// You can use the GetLoginBanner method to get the currently active Terms of Use banner that users see when they log on to the web interface.
	GetLoginBanner(ctx context.Context) (*GetLoginBannerResult, *SdkError)

	// This method dumps information about the chassis, local interfaces, and
	// neighbors from the LLDP daemon
	GetLldpInfo(ctx context.Context) (*GetLldpInfoResult, *SdkError)

	// The SolidFire Element OS web UI uses GetSnapMirrorClusterIdentity to get identity information about the ONTAP cluster.
	GetSnapMirrorClusterIdentity(ctx context.Context, req *GetSnapMirrorClusterIdentityRequest) (*GetSnapMirrorClusterIdentityResult, *SdkError)

	// Lists all active auth sessions.
	// This is only callable by a user with Administrative access rights.
	ListActiveAuthSessions(ctx context.Context) (*ListAuthSessionsResult, *SdkError)

	// Delete an existing configuration with a third party Identity Provider (IdP) for the cluster.
	// Deleting the last IdP Configuration will remove the SAML Service Provider certificate from the cluster.
	DeleteIdpConfiguration(ctx context.Context, req *DeleteIdpConfigurationRequest) (*DeleteIdpConfigurationResult, *SdkError)

	// ListAccounts returns the entire list of accounts, with optional paging support.
	ListAccounts(ctx context.Context, req *ListAccountsRequest) (*ListAccountsResult, *SdkError)

	//
	DeleteDatabaseEntry(ctx context.Context) (*NeedsWorkDeleteDatabaseEntryResult, *SdkError)

	//
	FinishUpgrade(ctx context.Context) (*NeedsWorkFinishUpgradeResult, *SdkError)

	// GetBackupTarget enables you to return information about a specific backup target that you have created.
	GetBackupTarget(ctx context.Context, req *GetBackupTargetRequest) (*GetBackupTargetResult, *SdkError)

	//
	GetClusterSettings(ctx context.Context) (*NeedsWorkGetClusterSettingsResult, *SdkError)

	// The Shutdown API method enables you to restart or shutdown a node that has not yet been added to a cluster. To use this method,
	// log in to the MIP for the pending node, and enter the "shutdown" method with either the "restart" or "halt" options.
	Shutdown(ctx context.Context, req *ShutdownRequest) (*ShutdownResult, *SdkError)

	// Demotes a node from being cluster master. The cluster will promote a new node.
	// The original cluster master will return success if it initiates the change. Callers should confirm that the
	// change occurs by calling the API command GetClusterMasterNodeID after a delay of a few seconds.
	DemoteClusterMaster(ctx context.Context, req *DemoteClusterMasterRequest) (*DemoteClusterMasterResult, *SdkError)

	// You can use GetClusterFullThreshold to view the stages set for cluster fullness levels. This method returns all fullness metrics for the
	// cluster.
	// Note: When a cluster reaches the Error stage of block cluster fullness, the maximum IOPS on all volumes are reduced linearly to the volume's minimum IOPS as the cluster approaches the Critical stage. This helps prevent the cluster from
	// reaching the Critical stage of block cluster fullness.
	GetClusterFullThreshold(ctx context.Context) (*GetClusterFullThresholdResult, *SdkError)

	// GetLimits enables you to retrieve the limit values set by the API. These values might change between releases of Element OS, but do not change without an update to the system. Knowing the limit values set by the API can be useful when writing API scripts for user-facing tools.
	// Note: The GetLimits method returns the limits for the current software version regardless of the API endpoint version used to pass the method.
	GetLimits(ctx context.Context) (*GetLimitsResult, *SdkError)

	//
	DeleteSnapshotSnapMirrorObjectAttributes(ctx context.Context) (*NeedsWorkDeleteSnapshotSnapMirrorObjectAttributesResult, *SdkError)

	// The TestLdapAuthentication method enables you to validate the currently enabled LDAP authentication settings. If the configuration is
	// correct, the API call returns the group membership of the tested user.
	TestLdapAuthentication(ctx context.Context, req *TestLdapAuthenticationRequest) (*TestLdapAuthenticationResult, *SdkError)

	// The SolidFire Element OS web UI uses the ListSnapMirrorAggregates method to list all SnapMirror aggregates that are available on the remote ONTAP system. An aggregate describes a set of physical storage resources.
	ListSnapMirrorAggregates(ctx context.Context, req *ListSnapMirrorAggregatesRequest) (*ListSnapMirrorAggregatesResult, *SdkError)

	// You can use GetSnmpState to return the current state of the SNMP feature.
	GetSnmpState(ctx context.Context) (*GetSnmpStateResult, *SdkError)

	// Enable support for authentication using a third party Identity Provider (IdP) for the cluster.
	// Once IdP authentication is enabled, cluster and Ldap admins will no longer be able to access the cluster via supported UIs and any active authenticated sessions will be invalidated/logged out.
	// Only third party IdP authenticated users will be able to access the cluster via the supported UIs.
	EnableIdpAuthentication(ctx context.Context, req *EnableIdpAuthenticationRequest) (*EnableIdpAuthenticationResult, *SdkError)

	//
	GetRsyslogInfo(ctx context.Context) (*NeedsWorkGetRsyslogInfoResult, *SdkError)

	// Creates a KMIP (Key Management Interoperability Protocol) Key Server with the specified attributes. The server
	// will not be contacted as part of this operation so it need not exist or be configured prior.  See TestKeyServerKmip.
	// For clustered Key Server configurations, the hostnames or IP Addresses, of all server nodes, must be provided in
	// the kmipKeyServerHostnames parameter.
	CreateKeyServerKmip(ctx context.Context, req *CreateKeyServerKmipRequest) (*CreateKeyServerKmipResult, *SdkError)

	// Modifies a KMIP (Key Management Interoperability Protocol) Key Server to the specified attributes. The only required
	// parameter is the keyServerID. A request which contains only the keyServerID will be a no-op and no error will be
	// returned. Any other parameters which are specified will replace the existing values on the KMIP Key Server with
	// the specified keyServerID. Because this server might be part of an active provider this will result in contacting
	// the server to verify it's functional. Multiple hostnames or IP addresses must only be provided to the kmipKeyServerHostnames
	// parameter if the key servers are in a clustered configuration.
	ModifyKeyServerKmip(ctx context.Context, req *ModifyKeyServerKmipRequest) (*ModifyKeyServerKmipResult, *SdkError)

	// CreateVirtualVolume is used to create a new (empty) Virtual Volume on the cluster.
	// When the volume is created successfully it is available for connection via PE.
	CreateVirtualVolume(ctx context.Context, req *CreateVirtualVolumeRequest) (*VirtualVolumeSyncResult, *SdkError)

	// AddVolumesToVolumeAccessGroup enables you to add
	// volumes to a specified volume access group.
	AddVolumesToVolumeAccessGroup(ctx context.Context, req *AddVolumesToVolumeAccessGroupRequest) (*ModifyVolumeAccessGroupResult, *SdkError)

	// Adds a cluster administrator user authenticated by a third party Identity Provider (IdP).
	// IdP cluster admin accounts are configured based on SAML attribute-value information provided
	// within the IdP's SAML assertion associated with the user.  If a user successfully
	// authenticates with the IdP and has SAML attribute statements within the SAML assertion
	// matching multiple IdP cluster admin accounts, the user will have the combined access level
	// of those matching IdP cluster admin accounts.
	AddIdpClusterAdmin(ctx context.Context, req *AddIdpClusterAdminRequest) (*AddClusterAdminResult, *SdkError)

	//
	StartGC(ctx context.Context) (*NeedsWorkStartGCResult, *SdkError)

	// AddDrives enables you to add one or more available drives to the cluster, enabling the drives to host a portion of the cluster's data.
	// When you add a node to the cluster or install new drives in an existing node, the new drives are marked as "available" and must be
	// added via AddDrives before they can be utilized. Use the ListDrives method to display drives that are "available" to be added. When
	// you add multiple drives, it is more efficient to add them in a single AddDrives method call rather than multiple individual methods
	// with a single drive each. This reduces the amount of data balancing that must occur to stabilize the storage load on the cluster.
	// When you add a drive, the system automatically determines the "type" of drive it should be.
	// The method is asynchronous and returns immediately. However, it can take some time for the data in the cluster to be rebalanced
	// using the newly added drives. As the new drives are syncing on the system, you can use the ListSyncJobs method to see how the
	// drives are being rebalanced and the progress of adding the new drive. You can also use the GetAsyncResult method to query the
	// method's returned asyncHandle.
	//
	// For testing, the following API Constants can be modified to reduce the usable capacity of the drives.
	// * cUsableDriveCapacityBytes:        usable capacity in bytes, ignored when set to 0 (default).
	// * cUsableDriveCapacityPercent:      usable capacity as a percent of max usable capacity, ignored when set to 0.0 (default).
	// If both constants are set, cUsableDriveCapacityPercent is used and cUsableDriveCapacityBytes is ignored.
	// Both constants are ignored if the usable capacity is less than the maximum usable capacity prior to calling this API.
	AddDrives(ctx context.Context, req *AddDrivesRequest) (*AddDrivesResult, *SdkError)

	// ModifySnapshot enables you to change the attributes currently assigned to a snapshot. You can use this method to enable snapshots created on
	// the Read/Write (source) volume to be remotely replicated to a target SolidFire storage system.
	ModifySnapshot(ctx context.Context, req *ModifySnapshotRequest) (*ModifySnapshotResult, *SdkError)

	// You can use ModifyVirtualNetwork to change the attributes of an existing virtual network. This method enables you to add or remove
	// address blocks, change the netmask, or modify the name or description of the virtual network. You can also use it to enable or
	// disable namespaces, as well as add or remove a gateway if namespaces are enabled on the virtual network.
	// Note: This method requires either the VirtualNetworkID or the VirtualNetworkTag as a parameter, but not both.
	// Caution: Enabling or disabling the Routable Storage VLANs functionality for an existing virtual network by changing the
	// "namespace" parameter disrupts any traffic handled by the virtual network. NetApp strongly recommends changing the
	// "namespace" parameter only during a scheduled maintenance window.
	ModifyVirtualNetwork(ctx context.Context, req *ModifyVirtualNetworkRequest) (*AddVirtualNetworkResult, *SdkError)

	// You can use AddAccount to add a new account to the system. You can create new volumes under the new account. The CHAP settings you specify for the account apply to all volumes owned by the account.
	AddAccount(ctx context.Context, req *AddAccountRequest) (*AddAccountResult, *SdkError)

	//
	EnableBmcColdReset(ctx context.Context) (*NeedsWorkEnableBmcColdResetResult, *SdkError)

	//
	SetClusterFullThresholds(ctx context.Context) (*NeedsWorkSetClusterFullThresholdsResult, *SdkError)

	// You can use ModifySnapMirrorRelationship to change the intervals at which a scheduled snapshot occurs. You can also delete or pause a schedule by using this method.
	ModifySnapMirrorRelationship(ctx context.Context, req *ModifySnapMirrorRelationshipRequest) (*ModifySnapMirrorRelationshipResult, *SdkError)

	// GetSnmpACL enables you to return the current SNMP access permissions on the cluster nodes.
	GetSnmpACL(ctx context.Context) (*GetSnmpACLResult, *SdkError)

	// You can use GetSnmpTrapInfo to return current SNMP trap configuration information.
	GetSnmpTrapInfo(ctx context.Context) (*GetSnmpTrapInfoResult, *SdkError)

	// ListInitiators enables you to list initiator IQNs or World Wide Port Names (WWPNs).
	ListInitiators(ctx context.Context, req *ListInitiatorsRequest) (*ListInitiatorsResult, *SdkError)

	//
	SetRepositories(ctx context.Context) (*NeedsWorkSetRepositoriesResult, *SdkError)

	// Remove (unassign) the specified KMIP (Key Management Interoperability Protocol) Key Server from the provider it was assigned to via AddKeyServerToProviderKmip (if any).  A KMIP Key Server can be unassigned from its provider unless it's the last one and that provider is active (providing keys which are currently in use).  If the specified KMIP Key Server is not assigned to a provider, this is a no-op and no error will be returned.
	RemoveKeyServerFromProviderKmip(ctx context.Context, req *RemoveKeyServerFromProviderKmipRequest) (*RemoveKeyServerFromProviderKmipResult, *SdkError)

	// You can use DisableSnmp to disable SNMP on the cluster nodes.
	DisableSnmp(ctx context.Context) (*DisableSnmpResult, *SdkError)

	//
	GetSliceInfo(ctx context.Context) (*NeedsWorkGetSliceInfoResult, *SdkError)

	// ListDrives enables you to retrieve the list of the drives that exist in the cluster's active nodes. This method returns drives that have
	// been added as volume metadata or block drives as well as drives that have not been added and are available.
	ListDrives(ctx context.Context) (*ListDrivesResult, *SdkError)

	// The SolidFire Element OS web UI uses the ModifySnapMirrorEndpoint method to change the name and management attributes for a SnapMirror endpoint.
	ModifySnapMirrorEndpoint(ctx context.Context, req *ModifySnapMirrorEndpointRequest) (*ModifySnapMirrorEndpointResult, *SdkError)

	// ListSnapshots enables you to return the attributes of each snapshot taken on the volume. Information about snapshots that reside on the target cluster is displayed on the source cluster when this method is called from the source cluster.
	ListSnapshots(ctx context.Context, req *ListSnapshotsRequest) (*ListSnapshotsResult, *SdkError)

	// You can use the ResurrectDeadVirtualVolume method to recover a virtual volume from a secondary replica
	// that may not have all data acknowledged to the client. This should only be used in circumstances
	// where the primary slice service is permanently inaccessible and you want to recover as much user
	// data as possible.
	ResurrectDeadVirtualVolume(ctx context.Context, req *ResurrectDeadVirtualVolumeRequest) (*ResurrectDeadVirtualVolumeResult, *SdkError)

	//
	ListSliceBranchesByService(ctx context.Context) (*NeedsWorkListSliceBranchesByServiceResult, *SdkError)

	// DeleteGroupSnapshot enables you to delete a group snapshot. You can use the saveMembers parameter to preserve all the snapshots that were made for the volumes in the group, but the group association is removed.
	DeleteGroupSnapshot(ctx context.Context, req *DeleteGroupSnapshotRequest) (*DeleteGroupSnapshotResult, *SdkError)

	// The GetImmutableValues API returns immutable constants that affect cluster behavior.
	GetImmutableValues(ctx context.Context) (*GetImmutableValuesResult, *SdkError)

	// Deletes all auth sessions for the given user.
	// A caller not in AccessGroup ClusterAdmins / Administrator may only delete their own sessions.
	// A caller with ClusterAdmins / Administrator privileges may delete sessions belonging to any user.
	// To see the list of sessions that could be deleted, use ListAuthSessionsByUsername with the same parameters.
	DeleteAuthSessionsByUsername(ctx context.Context, req *DeleteAuthSessionsByUsernameRequest) (*DeleteAuthSessionsResult, *SdkError)

	// Refreshes an auth session.
	// Intended to be used by the element-auth container.
	UpdateAuthSession(ctx context.Context, req *UpdateAuthSessionRequest) (*UpdateAuthSessionResult, *SdkError)

	// You can use SetLoginSessionInfo to set the period of time that a session's login authentication is valid. After the log in period elapses without activity on the system, the authentication expires. New login credentials are required for continued access to the cluster after the timeout period has elapsed.
	SetLoginSessionInfo(ctx context.Context, req *SetLoginSessionInfoRequest) (*SetLoginSessionInfoResult, *SdkError)

	// CheckProposedNodeAdditions validates that adding a node (or nodes) to an existing cluster is likely to succeed.  Any problems with the proposed new cluster are returned as errors with a human-readable description and unique error code.
	CheckProposedNodeAdditions(ctx context.Context, req *CheckProposedNodeAdditionsRequest) (*CheckProposedResult, *SdkError)

	// CreateSnapshot enables you to create a point-in-time copy of a volume. You can create a snapshot from any volume or from an existing snapshot. If you do not provide a SnapshotID with this API method, a snapshot is created from the volume's active branch.
	// If the volume from which the snapshot is created is being replicated to a remote cluster, the snapshot can also be replicated to the same target. Use the enableRemoteReplication parameter to enable snapshot replication.
	// Note: Creating a snapshot is allowed if cluster fullness is at stage 2 or 3. Snapshots are not created when cluster fullness is at stage 4 or 5.
	CreateSnapshot(ctx context.Context, req *CreateSnapshotRequest) (*CreateSnapshotResult, *SdkError)

	// The EnableLdapAuthentication method enables you to configure an LDAP directory connection to use for LDAP authentication to a cluster. Users that are members of the LDAP directory can then log in to the storage system using their LDAP credentials.
	EnableLdapAuthentication(ctx context.Context, req *EnableLdapAuthenticationRequest) (*EnableLdapAuthenticationResult, *SdkError)

	//
	StopClusterBSCheck(ctx context.Context) (*NeedsWorkStopClusterBSCheckResult, *SdkError)

	// PrepareVirtualSnapshot is used to set up VMware Virtual Volume snapshot.
	PrepareVirtualSnapshot(ctx context.Context, req *PrepareVirtualSnapshotRequest) (*PrepareVirtualSnapshotResult, *SdkError)

	// Lists all the existing cluster interface preferences.
	ListClusterInterfacePreferences(ctx context.Context) (*ListClusterInterfacePreferencesResult, *SdkError)

	//
	GetDatabaseEntry(ctx context.Context) (*NeedsWorkGetDatabaseEntryResult, *SdkError)

	// Delete the specified KMIP (Key Management Interoperability Protocol) Key Server.  A KMIP Key Server can be deleted unless it's the last one assigned to its provider, and that provider is active (providing keys which are currently in use).
	DeleteKeyServerKmip(ctx context.Context, req *DeleteKeyServerKmipRequest) (*DeleteKeyServerKmipResult, *SdkError)

	// The TestConnectEnsemble API method enables you to verify connectivity with a specified database ensemble. By default, it uses the ensemble for the cluster that the node is associated with. Alternatively, you can provide a different ensemble to test connectivity with.
	// Note: This method is available only through the per-node API endpoint 5.0 or later.
	TestConnectEnsemble(ctx context.Context, req *TestConnectEnsembleRequest) (*TestConnectEnsembleResult, *SdkError)

	// Creates a KMIP (Key Management Interoperability Protocol) Key Provider with the specified name.  A Key Provider defines a mechanism and location to retrieve authentication keys.  A KMIP Key Provider represents a collection of one or more KMIP Key Servers.  A newly created KMIP Key Provider will not have any KMIP Key Servers assigned to it.  To create a KMIP Key Server see CreateKeyServerKmip and to assign it to a provider created via this method see AddKeyServerToProviderKmip.
	CreateKeyProviderKmip(ctx context.Context, req *CreateKeyProviderKmipRequest) (*CreateKeyProviderKmipResult, *SdkError)

	// The RestartServices API method enables you to restart the services on a node.
	// Caution: This method causes temporary node services interruption. Exercise caution when using this method.
	// Note: This method is available only through the per-node API endpoint 5.0 or later.
	RestartServices(ctx context.Context, req *RestartServicesRequest) (*interface{}, *SdkError)

	// The SolidFire Element OS web UI uses the CreateSnapMirrorEndpoint method to create a relationship with a remote SnapMirror endpoint.
	CreateSnapMirrorEndpoint(ctx context.Context, req *CreateSnapMirrorEndpointRequest) (*CreateSnapMirrorEndpointResult, *SdkError)

	// AddLdapClusterAdmin enables you to add a new LDAP cluster administrator user. An LDAP cluster administrator can manage the
	// cluster via the API and management tools. LDAP cluster admin accounts are completely separate and unrelated to standard tenant
	// accounts.
	// You can also use this method to add an LDAP group that has been defined in Active Directory. The access level that is given to the group is passed to the individual users in the LDAP group.
	AddLdapClusterAdmin(ctx context.Context, req *AddLdapClusterAdminRequest) (*AddLdapClusterAdminResult, *SdkError)

	//
	APIWriteBlock(ctx context.Context) (*NeedsWorkAPIWriteBlockResult, *SdkError)

	//
	ForceRecycle(ctx context.Context) (*NeedsWorkForceRecycleResult, *SdkError)

	//
	ModifyClusterFault(ctx context.Context) (*NeedsWorkModifyClusterFaultResult, *SdkError)

	// CreateBackupTarget enables you to create and store backup target information so that you do not need to re-enter it each time a backup is created.
	CreateBackupTarget(ctx context.Context, req *CreateBackupTargetRequest) (*CreateBackupTargetResult, *SdkError)

	//
	SetDaemonStatus(ctx context.Context) (*NeedsWorkSetDaemonStatusResult, *SdkError)

	// ListSchedule enables you to retrieve information about all scheduled snapshots that have been created.
	ListSchedules(ctx context.Context) (*ListSchedulesResult, *SdkError)

	// CopyVolume enables you to overwrite the data contents of an existing volume with the data contents of another volume (or
	// snapshot). Attributes of the destination volume such as IQN, QoS settings, size, account, and volume access group membership are
	// not changed. The destination volume must already exist and must be the same size as the source volume.
	// NetApp strongly recommends that clients unmount the destination volume before the CopyVolume operation begins. If the
	// destination volume is modified during the copy operation, the changes will be lost.
	// This method is asynchronous and may take a variable amount of time to complete. You can use the GetAsyncResult method to
	// determine when the process has finished, and ListSyncJobs to see the progress of the copy.
	CopyVolume(ctx context.Context, req *CopyVolumeRequest) (*CopyVolumeResult, *SdkError)

	//
	ReadSliceLBA(ctx context.Context) (*NeedsWorkReadSliceLBAResult, *SdkError)

	// The SetNetworkConfig API method enables you to set the network configuration for a node. To display the current network settings for a node, run the GetNetworkConfig API method.
	// Note: This method is available only through the per-node API endpoint 5.0 or later.
	// Changing the "bond-mode" on a node can cause a temporary loss of network connectivity. Exercise caution when using this method.
	SetNetworkConfig(ctx context.Context, req *SetNetworkConfigRequest) (*SetNetworkConfigResult, *SdkError)

	// GetFeatureStatus enables you to retrieve the status of a cluster feature.
	GetFeatureStatus(ctx context.Context, req *GetFeatureStatusRequest) (*GetFeatureStatusResult, *SdkError)

	// ListVirtualVolumeHosts returns a list of all virtual volume hosts known to the cluster. A virtual volume host is a VMware ESX host
	// that has initiated a session with the VASA API provider.
	ListVirtualVolumeHosts(ctx context.Context, req *ListVirtualVolumeHostsRequest) (*ListVirtualVolumeHostsResult, *SdkError)

	// You can use ModifyVolumeAccessGroup to update initiators and add or remove volumes from a volume access group. If a specified initiator or volume is a duplicate of what currently exists, the volume access group is left as-is. If you do not specify a value for volumes or initiators, the current list of initiators and volumes is not changed.
	ModifyVolumeAccessGroup(ctx context.Context, req *ModifyVolumeAccessGroupRequest) (*ModifyVolumeAccessGroupResult, *SdkError)

	// DeleteAllSupportBundles enables you to delete all support bundles generated with the CreateSupportBundle API method.
	DeleteAllSupportBundles(ctx context.Context) (*DeleteAllSupportBundlesResult, *SdkError)

	//
	GetFibreChannelVolumeAccessInfo(ctx context.Context) (*NeedsWorkGetFibreChannelVolumeAccessInfoResult, *SdkError)

	// You can use the GetNodeSSLCertificate method to retrieve the SSL certificate that is currently active on the cluster.
	// You can use this method on both management and storage nodes.
	GetNodeSSLCertificate(ctx context.Context) (*GetNodeSSLCertificateResult, *SdkError)

	// BindVirtualVolume binds a VVol with a Host.
	BindVirtualVolumes(ctx context.Context, req *BindVirtualVolumesRequest) (*VirtualVolumeBindingListResult, *SdkError)

	// GetFipsReport enables you to retrieve FIPS compliance status on a per node basis.
	GetFipsReport(ctx context.Context) (*GetFipsReportResult, *SdkError)

	// Deletes an existing cluster interface preference.
	DeleteClusterInterfacePreference(ctx context.Context, req *DeleteClusterInterfacePreferenceRequest) (*DeleteClusterInterfacePreferenceResult, *SdkError)

	// You can use the RemoveBlockFromBS method to remove a block from the block services.
	RemoveBlockFromBS(ctx context.Context, req *RemoveBlockFromBSRequest) (*RemoveBlockFromBSResult, *SdkError)

	// GetLldpConfig returns the current LLDP configuration for this node.
	GetLldpConfig(ctx context.Context) (*GetLldpConfigResult, *SdkError)

	// The GetHardwareInfo API method enables you to return hardware information and status for a single node. This generally includes details about manufacturers, vendors, versions, drives, and other associated hardware identification information.
	GetHardwareInfo(ctx context.Context) (*GetHardwareInfoResult, *SdkError)

	// The SolidFire Element OS web UI uses the ListSnapMirrorNodes method to get a list of nodes in a remote ONTAP cluster.
	ListSnapMirrorNodes(ctx context.Context, req *ListSnapMirrorNodesRequest) (*ListSnapMirrorNodesResult, *SdkError)

	// ListVolumeStatsByVirtualVolume enables you to list volume statistics for any volumes in the system that are associated with virtual volumes. Statistics are cumulative from the creation of the volume.
	ListVolumeStatsByVirtualVolume(ctx context.Context, req *ListVolumeStatsByVirtualVolumeRequest) (*ListVolumeStatsByVirtualVolumeResult, *SdkError)

	// ListDriveHardware returns all the drives connected to a node. Use this method on individual nodes to return drive hardware
	// information or use this method on the cluster master node MVIP to see information for all the drives on all nodes.
	// Note: The "securitySupported": true line of the method response does not imply that the drives are capable of
	// encryption; only that the security status can be queried. If you have a node type with a model number ending in "-NE",
	// commands to enable security features on these drives will fail. See the EnableEncryptionAtRest method for more information.
	ListDriveHardware(ctx context.Context, req *ListDriveHardwareRequest) (*ListDriveHardwareResult, *SdkError)

	//
	BDTPipeline(ctx context.Context) (*NeedsWorkBDTPipelineResult, *SdkError)

	// EnableSnmp enables you to enable SNMP on cluster nodes. When you enable SNMP, the action applies to all nodes in the cluster, and
	// the values that are passed replace, in whole, all values set in any previous call to EnableSnmp.
	EnableSnmp(ctx context.Context, req *EnableSnmpRequest) (*EnableSnmpResult, *SdkError)

	// You can use ModifyClusterFullThreshold to change the level at which the system generates an event when the storage cluster approaches a certain capacity utilization. You can use the threshold settings to indicate the acceptable amount of utilized block storage or metadata storage before the system generates a warning. For example, if you want to be alerted when the system reaches 3% below the "Error" level block storage utilization, enter a value of "3" for the stage3BlockThresholdPercent parameter. If this level is reached, the system sends an alert to the Event Log in the Cluster Management Console.
	ModifyClusterFullThreshold(ctx context.Context, req *ModifyClusterFullThresholdRequest) (*ModifyClusterFullThresholdResult, *SdkError)

	// This method returns a string containing element specific configuration data set by the auth container.
	GetAuthConfiguration(ctx context.Context, req *GetAuthConfigurationRequest) (*GetAuthConfigurationResult, *SdkError)

	// GetCurrentClusterAdmin returns information about the calling ClusterAdmin.
	// If the authMethod in the return value is Ldap or Idp, then other fields in the return value may contain data aggregated from multiple LdapAdmins or IdpAdmins, respectively.
	GetCurrentClusterAdmin(ctx context.Context) (*GetCurrentClusterAdminResult, *SdkError)

	//
	ListClusterCapacityHistory(ctx context.Context) (*NeedsWorkListClusterCapacityHistoryResult, *SdkError)

	// ListPendingActiveNodes returns the list of nodes in the cluster that are currently in the PendingActive state, between the pending and active states. These are nodes that are currently being returned to the factory image.
	ListPendingActiveNodes(ctx context.Context) (*ListPendingActiveNodesResult, *SdkError)

	// You can use the CompleteVolumePairing method to complete the pairing of two volumes.
	CompleteVolumePairing(ctx context.Context, req *CompleteVolumePairingRequest) (*CompleteVolumePairingResult, *SdkError)

	// ListActivePairedVolumes enables you to list all the active volumes paired with a volume. This method returns information about volumes with active and pending pairings.
	ListActivePairedVolumes(ctx context.Context, req *ListActivePairedVolumesRequest) (*ListActivePairedVolumesResult, *SdkError)

	// You can use the GetSSLCertificate method to retrieve the SSL certificate that is currently active on the cluster.
	GetSSLCertificate(ctx context.Context) (*GetSSLCertificateResult, *SdkError)

	//
	GetConnectivityReport(ctx context.Context) (*NeedsWorkGetConnectivityReportResult, *SdkError)

	//
	SetClusterStructure(ctx context.Context) (*NeedsWorkSetClusterStructureResult, *SdkError)

	// GetIpmiInfo enables you to display a detailed reporting of sensors (objects) for node fans, intake and exhaust temperatures, and power supplies that are monitored by the system.
	GetIpmiInfo(ctx context.Context, req *GetIpmiInfoRequest) (*GetIpmiInfoResult, *SdkError)

	// The RemoveVolumeFromVolumeAccessGroup method enables you to remove volumes from a volume access group.
	RemoveVolumesFromVolumeAccessGroup(ctx context.Context, req *RemoveVolumesFromVolumeAccessGroupRequest) (*ModifyVolumeAccessGroupResult, *SdkError)

	//
	ProtocolVersionUpgradeAvailable(ctx context.Context) (*NeedsWorkProtocolVersionUpgradeAvailableResult, *SdkError)

	// ListFibreChannelPortInfo enables you to retrieve information about the Fibre Channel ports on a node.  The API method is intended for use on individual nodes; userid and password authentication is required for access to individual Fibre Channel nodes.
	ListFibreChannelPortInfo(ctx context.Context) (*ListFibreChannelPortInfoResult, *SdkError)

	//
	ReleaseFreeMemory(ctx context.Context) (*NeedsWorkReleaseFreeMemoryResult, *SdkError)

	// Gets the Vasa Provider info
	GetVasaProviderInfo(ctx context.Context) (*VasaProviderInfoResult, *SdkError)

	// RemoveBackupTarget allows you to delete backup targets.
	RemoveBackupTarget(ctx context.Context, req *RemoveBackupTargetRequest) (*RemoveBackupTargetResult, *SdkError)

	// The DisableLdapAuthentication method enables you to disable LDAP authentication and remove all LDAP configuration settings. This method does not remove any configured cluster admin accounts (user or group). However, those cluster admin accounts will no longer be able to log in.
	DisableLdapAuthentication(ctx context.Context) (*DisableLdapAuthenticationResult, *SdkError)

	// DeleteSnapshot enables you to delete a snapshot. A snapshot that is currently the "active" snapshot cannot be deleted. You must
	// rollback and make another snapshot "active" before the current snapshot can be deleted. For more details on rolling back snapshots, see RollbackToSnapshot.
	DeleteSnapshot(ctx context.Context, req *DeleteSnapshotRequest) (*DeleteSnapshotResult, *SdkError)

	// ModifyAccount enables you to modify an existing account.
	// When you lock an account, any existing connections from that account are immediately terminated. When you change an account's
	// CHAP settings, any existing connections remain active, and the new CHAP settings are used on subsequent connections or
	// reconnections.
	// To clear an account's attributes, specify {} for the attributes parameter.
	ModifyAccount(ctx context.Context, req *ModifyAccountRequest) (*ModifyAccountResult, *SdkError)

	// ListSyncJobs enables you to return information about synchronization jobs that are running on a SolidFire cluster. The type of
	// synchronization jobs that are returned with this method are slice, clone, and remote.
	ListSyncJobs(ctx context.Context) (*ListSyncJobsResult, *SdkError)

	// RemoveVolumePair enables you to remove the remote pairing between two volumes. Use this method on both the source and target volumes that are paired together. When you remove the volume pairing information, data is no longer replicated to or from the volume.
	RemoveVolumePair(ctx context.Context, req *RemoveVolumePairRequest) (*RemoveVolumePairResult, *SdkError)

	// ModifySchedule enables you to change the intervals at which a scheduled snapshot occurs. This allows for adjustment to the snapshot frequency and retention.
	ModifySchedule(ctx context.Context, req *ModifyScheduleRequest) (*ModifyScheduleResult, *SdkError)

	// CancelClone enables you to stop an ongoing CloneVolume or CopyVolume process. When you cancel a group clone operation, the
	// system completes and removes the operation's associated asyncHandle.
	CancelClone(ctx context.Context, req *CancelCloneRequest) (*CancelCloneResult, *SdkError)

	// You can use the RemoveSSLCertificate method to remove the user SSL certificate and private key for the cluster.
	// After the certificate and private key are removed, the cluster is configured to use the default certificate and private key.
	RemoveSSLCertificate(ctx context.Context) (*RemoveSSLCertificateResult, *SdkError)

	//
	SetDebugOptions(ctx context.Context) (*NeedsWorkSetDebugOptionsResult, *SdkError)

	//
	TestClientConnectivity(ctx context.Context) (*NeedsWorkTestClientConnectivityResult, *SdkError)

	// ListAllNodes enables you to retrieve a list of active and pending nodes in the cluster.
	ListAllNodes(ctx context.Context) (*ListAllNodesResult, *SdkError)

	// The SolidFire Element OS web UI uses the ListSnapMirrorEndpoints method to list all SnapMirror endpoints that the SolidFire cluster is communicating with.
	ListSnapMirrorEndpoints(ctx context.Context, req *ListSnapMirrorEndpointsRequest) (*ListSnapMirrorEndpointsResult, *SdkError)

	// You can use the SetLoginBanner method to set the active Terms of Use banner users see when they log on to the web interface.
	SetLoginBanner(ctx context.Context, req *SetLoginBannerRequest) (*SetLoginBannerResult, *SdkError)

	//
	NotifyIntentToRestart(ctx context.Context) (*NeedsWorkNotifyIntentToRestartResult, *SdkError)

	// PurgeDeletedVolume immediately and permanently purges a volume that has been deleted. You must delete a volume using
	// DeleteVolume before it can be purged. Volumes are purged automatically after a period of time, so usage of this method is not
	// typically required.
	PurgeDeletedVolume(ctx context.Context, req *PurgeDeletedVolumeRequest) (*PurgeDeletedVolumeResult, *SdkError)

	//
	ListDatabaseChildrenData(ctx context.Context) (*NeedsWorkListDatabaseChildrenDataResult, *SdkError)

	// GetAccountByName enables you to retrieve details about a specific account, given its username.
	GetAccountByName(ctx context.Context, req *GetAccountByNameRequest) (*GetAccountResult, *SdkError)

	// GetClusterStats enables you to retrieve high-level activity measurements for the cluster. Values returned are cumulative from the
	// creation of the cluster.
	GetClusterStats(ctx context.Context) (*GetClusterStatsResult, *SdkError)

	// DeleteInitiators enables you to delete one or more initiators from the system (and from any associated volumes or volume access
	// groups).
	// If DeleteInitiators fails to delete one of the initiators provided in the parameter, the system returns an error and does not delete any
	// initiators (no partial completion is possible).
	DeleteInitiators(ctx context.Context, req *DeleteInitiatorsRequest) (*DeleteInitiatorsResult, *SdkError)

	// ResetConstants is used to reset constants that have been modified with SetConstants.
	// ResetConstants returns a constant/value mapping for all constants that are still set to non-default values after the command completes.
	ResetConstants(ctx context.Context, req *ResetConstantsRequest) (*interface{}, *SdkError)

	// The SolidFire Element OS web UI uses the ResumeSnapMirrorRelationship method to enable future transfers for a quiesced SnapMirror relationship.
	ResumeSnapMirrorRelationship(ctx context.Context, req *ResumeSnapMirrorRelationshipRequest) (*ResumeSnapMirrorRelationshipResult, *SdkError)

	// You can use ListAsyncResults to list the results of all currently running and completed asynchronous methods on the system.
	// Querying asynchronous results with ListAsyncResults does not cause completed asyncHandles to expire; you can use GetAsyncResult
	// to query any of the asyncHandles returned by ListAsyncResults.
	ListAsyncResults(ctx context.Context, req *ListAsyncResultsRequest) (*ListAsyncResultsResult, *SdkError)

	// Retrieves an existing cluster interface preference.
	GetClusterInterfacePreference(ctx context.Context, req *GetClusterInterfacePreferenceRequest) (*GetClusterInterfacePreferenceResult, *SdkError)

	//
	DeleteVolumeSnapMirrorObjectAttributes(ctx context.Context) (*NeedsWorkDeleteVolumeSnapMirrorObjectAttributesResult, *SdkError)

	// You can use the ListQoSPolicies method to list all the settings of all QoS policies on the system.
	ListQoSPolicies(ctx context.Context) (*ListQoSPoliciesResult, *SdkError)

	// The CheckPingOnVlan API method provides the ability to test IP address(s) reachability on a not-yet-configured VLAN.
	// This API creates a temporary vlan interface, uses it to ping the provided list of host IP addresses, removes the VLAN interface.
	// If the vlan interface already exits TestPing should be used.
	CheckPingOnVlan(ctx context.Context, req *CheckPingOnVlanRequest) (*CheckPingOnVlanResult, *SdkError)

	// You can use GetPendingOperation to detect an operation on a node that is currently in progress. You can also use this method to report back when an operation has completed.
	// Note: method is available only through the per-node API endpoint 5.0 or later.
	GetPendingOperation(ctx context.Context) (*GetPendingOperationResult, *SdkError)

	// Delete the specified inactive Key Provider.
	DeleteKeyProviderKmip(ctx context.Context, req *DeleteKeyProviderKmipRequest) (*DeleteKeyProviderKmipResult, *SdkError)

	// Generates a Certificate Sign Request which can be signed by a Certificate Authority to generate a client certificate for the cluster.  This is part of establishing a trust relationship for interacting with external services.
	GetClientCertificateSignRequest(ctx context.Context) (*GetClientCertificateSignRequestResult, *SdkError)

	// GetAccountByID enables you to return details about a specific account, given its accountID.
	GetAccountByID(ctx context.Context, req *GetAccountByIDRequest) (*GetAccountResult, *SdkError)

	//
	CopyVolumeToSnapshot(ctx context.Context) (*NeedsWorkCopyVolumeToSnapshotResult, *SdkError)

	//
	DisableClusterSsh(ctx context.Context) (*NeedsWorkDisableClusterSshResult, *SdkError)

	// ListVolumeStatsByVolume returns high-level activity measurements for every volume, by volume. Values are cumulative from the
	// creation of the volume.
	ListVolumeStatsByVolume(ctx context.Context, req *ListVolumeStatsByVolumeRequest) (*ListVolumeStatsByVolumeResult, *SdkError)

	// You can use the AddVirtualNetwork method to add a new virtual network to a cluster configuration. When you add a virtual network,
	// an interface for each node is created and each interface will require a virtual network IP address. The number of IP addresses you
	// specify as a parameter for this API method must be equal to or greater than the number of nodes in the cluster. The system bulk
	// provisions virtual network addresses and assigns them to individual nodes automatically. You do not need to assign virtual
	// network addresses to nodes manually.
	// Note: You can use AddVirtualNetwork only to create a new virtual network. If you want to make changes to an
	// existing virtual network, use ModifyVirtualNetwork.
	// Note: Virtual network parameters must be unique to each virtual network when setting the namespace parameter to false.
	AddVirtualNetwork(ctx context.Context, req *AddVirtualNetworkRequest) (*AddVirtualNetworkResult, *SdkError)

	// CloneVolume enables you to create a copy of a volume. This method is asynchronous and might take a variable amount of time to complete. The cloning process begins immediately when you make the CloneVolume request and is representative of the state of the volume when the API method is issued. You can use the GetAsyncResult method to determine when the cloning process is complete and the new volume is available for connections. You can use ListSyncJobs to see the progress of creating the clone.
	// Note: The initial attributes and QoS settings for the volume are inherited from the volume being cloned. You can change these settings with ModifyVolume.
	// Note: Cloned volumes do not inherit volume access group memberships from the source volume.
	CloneVolume(ctx context.Context, req *CloneVolumeRequest) (*CloneVolumeResult, *SdkError)

	// GetVirtualVolumeTaskUpdate checks the status of a VVol Async Task.
	GetVirtualVolumeTaskUpdate(ctx context.Context, req *GetVirtualVolumeTaskUpdateRequest) (*VirtualVolumeTaskResult, *SdkError)

	// Deletes an individual auth session
	// If the calling user is not in the ClusterAdmins / Administrator AccessGroup, only auth session belonging
	// to the calling user can be deleted.
	DeleteAuthSession(ctx context.Context, req *DeleteAuthSessionRequest) (*DeleteAuthSessionResult, *SdkError)

	// Retrieve the protection schemes supported by the node.
	GetProtectionSchemes(ctx context.Context, req *GetProtectionSchemesRequest) (*GetProtectionSchemesResult, *SdkError)

	// ModifyVirtualVolumeHost changes an existing ESX host.
	ModifyVirtualVolumeHost(ctx context.Context, req *ModifyVirtualVolumeHostRequest) (*VirtualVolumeNullResult, *SdkError)

	// GetClusterMasterNodeID enables you to retrieve the ID of the node that can perform cluster-wide administration tasks and holds the
	// storage virtual IP address (SVIP) and management virtual IP address (MVIP).
	GetClusterMasterNodeID(ctx context.Context) (*GetClusterMasterNodeIDResult, *SdkError)

	// You can use ResizeDrives to modify the usable capacity for a slice or block drive. It is an internal API method mainly for
	// speeding up the execution of the tests by shrinking the drive size. At least one drive ID and the new usable size in bytes
	// fields shall be provided for this API method.  The method updates the usable capacity field in driveInfo in database for each
	// given drive ID. Drives can only have their size modified if they are in the "available" status. The actual returned new
	// usable capacity might be adjusted based on the following rules.
	//
	// - Usable drive size must be a multiple of the segment file size of the drive.
	// - Usable drive size cannot be increased past the raw size of the drive.
	// - Usable drive size cannot be decreased lower than the minimum usable capacity below.
	//
	// The minimum usable capacity requires at least one segment file size for recycling, as well as at least
	// 30 additional segment files. Additionally the drive must always be at least 25GB.  Expressed as a formula:
	// minUsableCapacityBytes = max(cMinUsableDriveCapacityBytes, (segmentFileSizeBytes / min(cBSRecycleAtPct, cSSRecycleAtPct) + (segmentFileSizeBytes * cResizePaddingSegmentFiles))
	//
	// * cBSRecycleAtPct:        block service recycle threshold, 5% of drive usable capacity by default.
	// * cSSRecycleAtPct:        slice service recycle threshold, 10% of drive usable capacity by default.
	// * cResizePaddingSegmentFiles: extra number of segment files reserved for usable capacity, 30 by default.
	// * cMinUsableDriveCapacityBytes: absolute minimum size of a drive, 25GB by default.
	//
	// A drive that has a size smaller than the maximum usable size is not eligible to be resized
	// via the :ref:`AddDrives` API using the cUsableDriveCapacityBytes or cUsableDriveCapacityPercent API constants.
	//
	// When resizing multiple drives, all the database operations for all the given drives will be implemented atomically. It shall
	// not introduce any partial database changes due to single drive operation failure. The API method returns when all the database
	// changes complete, or when error happens.
	ResizeDrives(ctx context.Context, req *ResizeDrivesRequest) (*ResizeDrivesResult, *SdkError)

	// GetLoginSessionInfo enables you to return the period of time a log in authentication session is valid for both log in shells and the TUI.
	GetLoginSessionInfo(ctx context.Context) (*GetLoginSessionInfoResult, *SdkError)

	//
	GetDaemonStatus(ctx context.Context) (*NeedsWorkGetDaemonStatusResult, *SdkError)

	// Initiate the process of removing the password from self-encrypting drives (SEDs) within the cluster.
	DisableEncryptionAtRest(ctx context.Context, req *DisableEncryptionAtRestRequest) (*DisableEncryptionAtRestResult, *SdkError)

	// CreateStorageContainer enables you to create a Virtual Volume (VVol) storage container. Storage containers are associated with a SolidFire storage system account, and are used for reporting and resource allocation. Storage containers can only be associated with virtual volumes. You need at least one storage container to use the Virtual Volumes feature.
	CreateStorageContainer(ctx context.Context, req *CreateStorageContainerRequest) (*CreateStorageContainerResult, *SdkError)

	//
	GetEnsemble(ctx context.Context) (*NeedsWorkGetEnsembleResult, *SdkError)

	// FastCloneVirtualVolume is used to execute a VMware Virtual Volume fast clone.
	FastCloneVirtualVolume(ctx context.Context, req *FastCloneVirtualVolumeRequest) (*VirtualVolumeAsyncResult, *SdkError)

	// ListVolumesForAccount returns the list of active and (pending) deleted volumes for an account.
	ListVolumesForAccount(ctx context.Context, req *ListVolumesForAccountRequest) (*ListVolumesForAccountResult, *SdkError)

	// DeleteVolumeAccessGroup enables you to delete a
	// volume access group.
	DeleteVolumeAccessGroup(ctx context.Context, req *DeleteVolumeAccessGroupRequest) (*DeleteVolumeAccessGroupResult, *SdkError)

	// UnbindAllVirtualVolumesFromHost removes all VVol <-> Host binding.
	UnbindAllVirtualVolumesFromHost(ctx context.Context, req *UnbindAllVirtualVolumesFromHostRequest) (*UnbindAllVirtualVolumesFromHostResult, *SdkError)

	// You can use UpdateBulkVolumeStatus in a script to update the status of a bulk volume job that you started with the
	// StartBulkVolumeRead or StartBulkVolumeWrite methods.
	UpdateBulkVolumeStatus(ctx context.Context, req *UpdateBulkVolumeStatusRequest) (*UpdateBulkVolumeStatusResult, *SdkError)

	// Disable support for authentication using third party Identity Providers (IdP) for the cluster.
	// Once disabled, users authenticated by third party IdPs will no longer be able to access the cluster and any active authenticated sessions will be invalidated/logged out.
	// Ldap and cluster admins will be able to access the cluster via supported UIs.
	DisableIdpAuthentication(ctx context.Context) (*DisableIdpAuthenticationResult, *SdkError)

	//
	CheckBlock(ctx context.Context) (*NeedsWorkCheckBlockResult, *SdkError)

	//
	CheckClusterPair(ctx context.Context) (*NeedsWorkCheckClusterPairResult, *SdkError)

	//
	ListDatabaseChildren(ctx context.Context) (*NeedsWorkListDatabaseChildrenResult, *SdkError)

	// AddNodes enables you to add one or more new nodes to a cluster. When a node that is not configured starts up for the first time, you are prompted to configure the node. After you configure the node, it is registered as a "pending node" with the cluster.
	// Note: It might take several seconds after adding a new node for it to start up and register its drives as available.
	AddNodes(ctx context.Context, req *AddNodesRequest) (*AddNodesResult, *SdkError)

	// You can use the TestDrives API method to run a hardware validation on all drives on the node. This method detects hardware
	// failures on the drives (if present) and reports them in the results of the validation tests.
	// You can only use the TestDrives method on nodes that are not "active" in a cluster.
	// Note: This test takes approximately 10 minutes.
	// Note: This method is available only through the per-node API endpoint 5.0 or later.
	TestDrives(ctx context.Context, req *TestDrivesRequest) (*TestDrivesResult, *SdkError)

	// SetNtpInfo enables you to configure NTP on cluster nodes. The values you set with this interface apply to all nodes in the cluster. If an NTP broadcast server periodically broadcasts time information on your network, you can optionally configure nodes as broadcast clients.
	// Note: NetApp recommends using NTP servers that are internal to your network, rather than the installation defaults.
	SetNtpInfo(ctx context.Context, req *SetNtpInfoRequest) (*SetNtpInfoResult, *SdkError)

	//
	SetDatabaseEntry(ctx context.Context) (*NeedsWorkSetDatabaseEntryResult, *SdkError)

	// You can use ListISCSISessions to return iSCSI information for volumes in the cluster.
	ListISCSISessions(ctx context.Context) (*ListISCSISessionsResult, *SdkError)

	// The SolidFire Element OS web UI uses GetOntapVersionInfo to get information about API version support from the ONTAP cluster in a SnapMirror relationship.
	GetOntapVersionInfo(ctx context.Context, req *GetOntapVersionInfoRequest) (*GetOntapVersionInfoResult, *SdkError)

	// The GetNodeFipsDrivesReport reports the FipsDrives capability of a node. This is a per-node API.
	GetNodeFipsDrivesReport(ctx context.Context) (*GetNodeFipsDrivesReportResult, *SdkError)

	// You can use the GetSchedule method to retrieve information about a scheduled snapshot. You can see information about a specific
	// schedule if there are many snapshot schedules in the system. You also retrieve information about more than one schedule with this
	// method by specifying additional scheduleIDs in the parameter.
	GetSchedule(ctx context.Context, req *GetScheduleRequest) (*GetScheduleResult, *SdkError)

	// The SolidFire Element OS web UI uses the CreateSnapMirrorRelationship method to create a SnapMirror extended data protection relationship between a source and destination endpoint.
	CreateSnapMirrorRelationship(ctx context.Context, req *CreateSnapMirrorRelationshipRequest) (*CreateSnapMirrorRelationshipResult, *SdkError)

	// CloneVirtualVolume is used to execute a VMware Virtual Volume clone.
	CloneVirtualVolume(ctx context.Context, req *CloneVirtualVolumeRequest) (*VirtualVolumeAsyncResult, *SdkError)

	// You can use SetSnmpTrapInfo to enable and disable the generation of cluster SNMP notifications (traps) and to specify the set of network host computers that receive the notifications. The values you pass with each SetSnmpTrapInfo method call replace all values set in any previous call to SetSnmpTrapInfo.
	SetSnmpTrapInfo(ctx context.Context, req *SetSnmpTrapInfoRequest) (*SetSnmpTrapInfoResult, *SdkError)

	//
	GetClusterSshInfo(ctx context.Context) (*NeedsWorkGetClusterSshInfoResult, *SdkError)

	//
	GetServiceStatus(ctx context.Context) (*NeedsWorkGetServiceStatusResult, *SdkError)

	// ModifyVirtualVolume is used to modify settings on an existing virtual volume.
	ModifyVirtualVolume(ctx context.Context, req *ModifyVirtualVolumeRequest) (*VirtualVolumeNullResult, *SdkError)

	// Sets the default protection scheme stored in the cluster info.
	SetDefaultProtectionScheme(ctx context.Context, req *SetDefaultProtectionSchemeRequest) (*SetDefaultProtectionSchemeResult, *SdkError)

	//
	FindMismatchedSliceLBAs(ctx context.Context) (*NeedsWorkFindMismatchedSliceLBAsResult, *SdkError)

	// Test whether the specified Key Provider is functioning normally.
	TestKeyProviderKmip(ctx context.Context, req *TestKeyProviderKmipRequest) (*TestKeyProviderKmipResult, *SdkError)

	// Lists all auth sessions for the given user.
	// A caller not in AccessGroup ClusterAdmins / Administrator privileges may only list their own sessions.
	// A caller with ClusterAdmins / Administrator privileges may list sessions belonging to any user.
	ListAuthSessionsByUsername(ctx context.Context, req *ListAuthSessionsByUsernameRequest) (*ListAuthSessionsResult, *SdkError)

	// The SolidFire Element OS web UI uses the AbortSnapMirrorRelationship method to stop SnapMirror transfers that have started but are not yet complete.
	AbortSnapMirrorRelationship(ctx context.Context, req *AbortSnapMirrorRelationshipRequest) (*AbortSnapMirrorRelationshipResult, *SdkError)

	// SnapshotVirtualVolume is used to take a VMware Virtual Volume snapshot.
	SnapshotVirtualVolume(ctx context.Context, req *SnapshotVirtualVolumeRequest) (*SnapshotVirtualVolumeResult, *SdkError)

	// You can use the SetSupplementalTlsCiphers method to specify the list of supplemental TLS ciphers.
	SetSupplementalTlsCiphers(ctx context.Context, req *SetSupplementalTlsCiphersRequest) (*SetSupplementalTlsCiphersResult, *SdkError)

	// You can use the SetNodeSupplementalTlsCiphers method to specify the list of supplemental TLS ciphers for this node.
	// You can use this command on management nodes.
	SetNodeSupplementalTlsCiphers(ctx context.Context, req *SetNodeSupplementalTlsCiphersRequest) (*SetNodeSupplementalTlsCiphersResult, *SdkError)

	// The SolidFire Element OS web UI uses the ListSnapMirrorSchedules method to get a list of schedules that are available on a remote ONTAP cluster.
	ListSnapMirrorSchedules(ctx context.Context, req *ListSnapMirrorSchedulesRequest) (*ListSnapMirrorSchedulesResult, *SdkError)

	// This will invoke any API method supported by the SolidFire API for the version and port the connection is using.
	// Returns a nested hashtable of key/value pairs that contain the result of the invoked method.
	InvokeSFApi(ctx context.Context, req *InvokeSFApiRequest) (*interface{}, *SdkError)

	// You can use ModifyClusterAdmin to change the settings for a cluster admin, LDAP cluster admin,
	// or third party Identity Provider (IdP) cluster admin.  You cannot change access for the
	// administrator cluster admin account.
	ModifyClusterAdmin(ctx context.Context, req *ModifyClusterAdminRequest) (*ModifyClusterAdminResult, *SdkError)

	// GetDriveStats returns high-level activity measurements for a single drive. Values are cumulative from the addition of the drive to the
	// cluster. Some values are specific to block drives. You might not obtain statistical data for both block and metadata drives when you
	// run this method.
	GetDriveStats(ctx context.Context, req *GetDriveStatsRequest) (*GetDriveStatsResult, *SdkError)

	// CancelGroupClone enables you to stop an ongoing CloneMultipleVolumes process occurring on a group of volumes. When you cancel
	// a group clone operation, the system completes and removes the operation's associated asyncHandle.
	CancelGroupClone(ctx context.Context, req *CancelGroupCloneRequest) (*CancelGroupCloneResult, *SdkError)

	//
	GetCodeTimings(ctx context.Context) (*NeedsWorkGetCodeTimingsResult, *SdkError)

	// The StartClusterBSCheck API method is used to tell all slice services to go through all of their
	// slices and read all the block IDs and read them from all the block services that should have that
	// block ID. This API method finds xUnknownBlockID and reports them in the event report. The fields
	// volumeID, sliceID, volumeIDs, and sliceIDs are mutually exclusive and if all fields they will be
	// handled in the order of volumeID, sliceID, volumeIDSs, then sliceIDs.
	StartClusterBSCheck(ctx context.Context, req *StartClusterBSCheckRequest) (*StartClusterBSCheckResult, *SdkError)

	// The SolidFire Element OS web UI uses the ListSnapMirrorVolumes method to list all SnapMirror volumes available on a remote ONTAP system.
	ListSnapMirrorVolumes(ctx context.Context, req *ListSnapMirrorVolumesRequest) (*ListSnapMirrorVolumesResult, *SdkError)

	// You can use the ListUtilities API method to return the operations that are available to run on a node.
	// Note: This method is available only through the per-node API endpoint 5.0 or later.
	ListUtilities(ctx context.Context, req *ListUtilitiesRequest) (*ListUtilitiesResult, *SdkError)

	// You can use CreateVolumeAccessGroup to create a new volume access group. When you create the volume access group, you need to give it a name, and you can optionally enter initiators and volumes. After you create the group, you can add volumes and initiator IQNs. Any initiator IQN that you add to the volume access group is able to access any volume in the group without CHAP authentication.
	CreateVolumeAccessGroup(ctx context.Context, req *CreateVolumeAccessGroupRequest) (*CreateVolumeAccessGroupResult, *SdkError)
}
