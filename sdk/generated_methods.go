package sdk

// THIS IS A GENERATED FILE.  DO NOT MODIFY

import "context"

// You can use the ListClusterPairs method to list all the clusters that a cluster is paired with. This method returns information about active and pending cluster pairings, such as statistics about the current pairing as well as the connectivity and latency (in milliseconds) of the cluster pairing.
func (sfClient *SFClient) ListClusterPairs(ctx context.Context) (*ListClusterPairsResult, *SdkError) {
	var res ListClusterPairsResult
	_, err := sfClient.MakeSFCall(ctx, "ListClusterPairs", 1, nil, &res)
	return &res, err
}

// The SolidFire Element OS web UI uses the ListSnapMirrorPolicies method to list all SnapMirror policies on a remote ONTAP system.
func (sfClient *SFClient) ListSnapMirrorPolicies(ctx context.Context, req *ListSnapMirrorPoliciesRequest) (*ListSnapMirrorPoliciesResult, *SdkError) {
	var res ListSnapMirrorPoliciesResult
	_, err := sfClient.MakeSFCall(ctx, "ListSnapMirrorPolicies", 1, req, &res)
	return &res, err
}

// RollbackVirtualVolume is used to restore a VMware Virtual Volume snapshot.
func (sfClient *SFClient) RollbackVirtualVolume(ctx context.Context, req *RollbackVirtualVolumeRequest) (*VirtualVolumeAsyncResult, *SdkError) {
	var res VirtualVolumeAsyncResult
	_, err := sfClient.MakeSFCall(ctx, "RollbackVirtualVolume", 1, req, &res)
	return &res, err
}

// GetVolumeStats enables  you to retrieve high-level activity measurements for a single volume. Values are cumulative from the creation of the volume.
func (sfClient *SFClient) GetVolumeStats(ctx context.Context, req *GetVolumeStatsRequest) (*GetVolumeStatsResult, *SdkError) {
	var res GetVolumeStatsResult
	_, err := sfClient.MakeSFCall(ctx, "GetVolumeStats", 1, req, &res)
	return &res, err
}

// ListDeletedVolumes enables you to retrieve the list of volumes that have been marked for deletion and purged from the system.
func (sfClient *SFClient) ListDeletedVolumes(ctx context.Context, req *ListDeletedVolumesRequest) (*ListDeletedVolumesResult, *SdkError) {
	var res ListDeletedVolumesResult
	_, err := sfClient.MakeSFCall(ctx, "ListDeletedVolumes", 1, req, &res)
	return &res, err
}

// GetDriveHardwareInfo returns all the hardware information for the given drive. This generally includes details about manufacturers, vendors, versions, and
// other associated hardware identification information.
func (sfClient *SFClient) GetDriveHardwareInfo(ctx context.Context, req *GetDriveHardwareInfoRequest) (*GetDriveHardwareInfoResult, *SdkError) {
	var res GetDriveHardwareInfoResult
	_, err := sfClient.MakeSFCall(ctx, "GetDriveHardwareInfo", 1, req, &res)
	return &res, err
}

// Enables retrieval of the number of virtual volumes currently in the system.
func (sfClient *SFClient) GetVirtualVolumeCount(ctx context.Context) (*GetVirtualVolumeCountResult, *SdkError) {
	var res GetVirtualVolumeCountResult
	_, err := sfClient.MakeSFCall(ctx, "GetVirtualVolumeCount", 1, nil, &res)
	return &res, err
}

// ModifyVolumePair enables you to pause or restart replication between a pair of volumes.
func (sfClient *SFClient) ModifyVolumePair(ctx context.Context, req *ModifyVolumePairRequest) (*ModifyVolumePairResult, *SdkError) {
	var res ModifyVolumePairResult
	_, err := sfClient.MakeSFCall(ctx, "ModifyVolumePair", 1, req, &res)
	return &res, err
}

// GetStorageContainerEfficiency enables you to retrieve efficiency information about a virtual volume storage container.
func (sfClient *SFClient) GetStorageContainerEfficiency(ctx context.Context, req *GetStorageContainerEfficiencyRequest) (*GetStorageContainerEfficiencyResult, *SdkError) {
	var res GetStorageContainerEfficiencyResult
	_, err := sfClient.MakeSFCall(ctx, "GetStorageContainerEfficiency", 1, req, &res)
	return &res, err
}

// CopyDiffsToVirtualVolume is a three-way merge function.
func (sfClient *SFClient) CopyDiffsToVirtualVolume(ctx context.Context, req *CopyDiffsToVirtualVolumeRequest) (*VirtualVolumeAsyncResult, *SdkError) {
	var res VirtualVolumeAsyncResult
	_, err := sfClient.MakeSFCall(ctx, "CopyDiffsToVirtualVolume", 1, req, &res)
	return &res, err
}

// StartBulkVolumeWrite enables you to initialize a bulk volume write session on a specified volume. Only two bulk volume processes can run simultaneously on a volume. When you initialize the write session, data is written to a SolidFire storage volume from an external backup source. The external data is accessed by a web server running on an SF-series node. Communications and server
// interaction information for external data access is passed by a script running on the storage system.
func (sfClient *SFClient) StartBulkVolumeWrite(ctx context.Context, req *StartBulkVolumeWriteRequest) (*StartBulkVolumeWriteResult, *SdkError) {
	var res StartBulkVolumeWriteResult
	_, err := sfClient.MakeSFCall(ctx, "StartBulkVolumeWrite", 1, req, &res)
	return &res, err
}

// You can use AddClusterAdmin to add a new cluster admin account. A cluster ddmin can manage the cluster using the API and management tools. Cluster admins are completely separate and unrelated to standard tenant accounts.
// Each cluster admin can be restricted to a subset of the API. NetApp recommends using multiple cluster admin accounts for different users and applications. You should give each cluster admin the minimal permissions necessary; this reduces the potential impact of credential compromise.
// You must accept the End User License Agreement (EULA) by setting the acceptEula parameter to true to add a cluster administrator account to the system.
func (sfClient *SFClient) AddClusterAdmin(ctx context.Context, req *AddClusterAdminRequest) (*AddClusterAdminResult, *SdkError) {
	var res AddClusterAdminResult
	_, err := sfClient.MakeSFCall(ctx, "AddClusterAdmin", 1, req, &res)
	return &res, err
}

// The SetConfig API method enables you to set all the configuration information for the node. This includes the same information available via calls to SetClusterConfig and SetNetworkConfig in one API method.
// Note: This method is available only through the per-node API endpoint 5.0 or later.
// Caution: Changing the "bond-mode" on a node can cause a temporary loss of network connectivity. Exercise caution when using this method.
func (sfClient *SFClient) SetConfig(ctx context.Context, req *SetConfigRequest) (*SetConfigResult, *SdkError) {
	var res SetConfigResult
	_, err := sfClient.MakeSFCall(ctx, "SetConfig", 1, req, &res)
	return &res, err
}

// The RestartNetworking API method enables you to restart the networking services on a node.
// Warning: This method restarts all networking services on a node, causing temporary loss of networking connectivity.
// Exercise caution when using this method.
// Note: This method is available only through the per-node API endpoint 5.0 or later.
func (sfClient *SFClient) RestartNetworking(ctx context.Context, req *RestartNetworkingRequest) (*interface{}, *SdkError) {
	var res interface{}
	_, err := sfClient.MakeSFCall(ctx, "RestartNetworking", 1, req, &res)
	return &res, err
}

// ModifyStorageContainer enables you to make changes to an existing virtual volume storage container.
func (sfClient *SFClient) ModifyStorageContainer(ctx context.Context, req *ModifyStorageContainerRequest) (*ModifyStorageContainerResult, *SdkError) {
	var res ModifyStorageContainerResult
	_, err := sfClient.MakeSFCall(ctx, "ModifyStorageContainer", 1, req, &res)
	return &res, err
}

// You can use the GetAPI method to return a list of all the API methods and supported API endpoints that can be used in the system.
func (sfClient *SFClient) GetAPI(ctx context.Context, req *GetAPIRequest) (*GetAPIResult, *SdkError) {
	var res GetAPIResult
	_, err := sfClient.MakeSFCall(ctx, "GetAPI", 1, req, &res)
	return &res, err
}

// Returns the list of KMIP (Key Management Interoperability Protocol) Key Providers which have been created via CreateKeyProviderKmip.  The list can optionally be filtered by specifying additional parameters.
func (sfClient *SFClient) ListKeyProvidersKmip(ctx context.Context, req *ListKeyProvidersKmipRequest) (*ListKeyProvidersKmipResult, *SdkError) {
	var res ListKeyProvidersKmipResult
	_, err := sfClient.MakeSFCall(ctx, "ListKeyProvidersKmip", 1, req, &res)
	return &res, err
}

// Creates a new cluster preference and stores it on the storage cluster. The ClusterInterfacePreference
// related APIs can be used by internal interfaces to the storage cluster such as HCI and UI to store arbitrary
// information in the cluster. Since the API calls in the UI are visible to customers, these APIs are made public.
func (sfClient *SFClient) CreateClusterInterfacePreference(ctx context.Context, req *CreateClusterInterfacePreferenceRequest) (*CreateClusterInterfacePreferenceResult, *SdkError) {
	var res CreateClusterInterfacePreferenceResult
	_, err := sfClient.MakeSFCall(ctx, "CreateClusterInterfacePreference", 1, req, &res)
	return &res, err
}

// Returns the specified KMIP (Key Management Interoperability Protocol) Key Provider object.
func (sfClient *SFClient) GetKeyProviderKmip(ctx context.Context, req *GetKeyProviderKmipRequest) (*GetKeyProviderKmipResult, *SdkError) {
	var res GetKeyProviderKmipResult
	_, err := sfClient.MakeSFCall(ctx, "GetKeyProviderKmip", 1, req, &res)
	return &res, err
}

// ListVirtualNetworks enables you to list all configured virtual networks for the cluster. You can use this method to verify the virtual
// network settings in the cluster.
// There are no required parameters for this method. However, to filter the results, you can pass one or more VirtualNetworkID or
// VirtualNetworkTag values.
func (sfClient *SFClient) ListVirtualNetworks(ctx context.Context, req *ListVirtualNetworksRequest) (*ListVirtualNetworksResult, *SdkError) {
	var res ListVirtualNetworksResult
	_, err := sfClient.MakeSFCall(ctx, "ListVirtualNetworks", 1, req, &res)
	return &res, err
}

// ListProtocolEndpoints enables you to retrieve information about all protocol endpoints in the cluster. Protocol endpoints govern
// access to their associated virtual volume storage containers.
func (sfClient *SFClient) ListProtocolEndpoints(ctx context.Context, req *ListProtocolEndpointsRequest) (*ListProtocolEndpointsResult, *SdkError) {
	var res ListProtocolEndpointsResult
	_, err := sfClient.MakeSFCall(ctx, "ListProtocolEndpoints", 1, req, &res)
	return &res, err
}

// CloneMultipleVolumes enables you to create a clone of a group of specified volumes. You can assign a consistent set of characteristics
// to a group of multiple volumes when they are cloned together.
// Before using groupSnapshotID to clone the volumes in a group snapshot, you must create the group snapshot by using the
// CreateGroupSnapshot API method or the Element OS Web UI. Using groupSnapshotID is optional when cloning multiple volumes.
// Note: Cloning multiple volumes is allowed if cluster fullness is at stage 2 or 3. Clones are not created when cluster fullness is
// at stage 4 or 5.
func (sfClient *SFClient) CloneMultipleVolumes(ctx context.Context, req *CloneMultipleVolumesRequest) (*CloneMultipleVolumesResult, *SdkError) {
	var res CloneMultipleVolumesResult
	_, err := sfClient.MakeSFCall(ctx, "CloneMultipleVolumes", 1, req, &res)
	return &res, err
}

// You can use the GetClusterHardwareInfo method to retrieve the hardware status and information for all Fibre Channel nodes, iSCSI
// nodes and drives in the cluster. This generally includes details about manufacturers, vendors, versions, and other associated hardware
// identification information.
func (sfClient *SFClient) GetClusterHardwareInfo(ctx context.Context, req *GetClusterHardwareInfoRequest) (*GetClusterHardwareInfoResult, *SdkError) {
	var res GetClusterHardwareInfoResult
	_, err := sfClient.MakeSFCall(ctx, "GetClusterHardwareInfo", 1, req, &res)
	return &res, err
}

//
func (sfClient *SFClient) ListSnapMirrorObjectAttributes(ctx context.Context) (*NeedsWorkListSnapMirrorObjectAttributesResult, *SdkError) {
	var res NeedsWorkListSnapMirrorObjectAttributesResult
	_, err := sfClient.MakeSFCall(ctx, "ListSnapMirrorObjectAttributes", 1, nil, &res)
	return &res, err
}

// AddInitiatorsToVolumeAccessGroup enables you
// to add initiators to a specified volume access group.
func (sfClient *SFClient) AddInitiatorsToVolumeAccessGroup(ctx context.Context, req *AddInitiatorsToVolumeAccessGroupRequest) (*ModifyVolumeAccessGroupResult, *SdkError) {
	var res ModifyVolumeAccessGroupResult
	_, err := sfClient.MakeSFCall(ctx, "AddInitiatorsToVolumeAccessGroup", 1, req, &res)
	return &res, err
}

//
func (sfClient *SFClient) ReadSliceLBAChecksum(ctx context.Context) (*NeedsWorkReadSliceLBAChecksumResult, *SdkError) {
	var res NeedsWorkReadSliceLBAChecksumResult
	_, err := sfClient.MakeSFCall(ctx, "ReadSliceLBAChecksum", 1, nil, &res)
	return &res, err
}

// You can use the SetNodeSSLCertificate method to set a user SSL certificate and private key for the management node.
func (sfClient *SFClient) SetNodeSSLCertificate(ctx context.Context, req *SetNodeSSLCertificateRequest) (*SetNodeSSLCertificateResult, *SdkError) {
	var res SetNodeSSLCertificateResult
	_, err := sfClient.MakeSFCall(ctx, "SetNodeSSLCertificate", 1, req, &res)
	return &res, err
}

// SnmpSendTestTraps enables you to test SNMP functionality for a cluster. This method instructs the cluster to send test SNMP traps to the currently configured SNMP manager.
func (sfClient *SFClient) SnmpSendTestTraps(ctx context.Context) (*SnmpSendTestTrapsResult, *SdkError) {
	var res SnmpSendTestTrapsResult
	_, err := sfClient.MakeSFCall(ctx, "SnmpSendTestTraps", 1, nil, &res)
	return &res, err
}

// GetVirtualVolumeAllocatedBitmap scans a VVol segment and returns the number of
// chunks not shared between two volumes. This call will return results in less
// than 30 seconds. If the specified VVol and the base VVil are not related, an
// error is returned. If the offset/length combination is invalid or out of range
// an error is returned.
func (sfClient *SFClient) GetVirtualVolumeUnsharedChunks(ctx context.Context, req *GetVirtualVolumeUnsharedChunksRequest) (*VirtualVolumeUnsharedChunkResult, *SdkError) {
	var res VirtualVolumeUnsharedChunkResult
	_, err := sfClient.MakeSFCall(ctx, "GetVirtualVolumeUnsharedChunks", 1, req, &res)
	return &res, err
}

// CreateVolume enables you to create a new (empty) volume on the cluster. As soon as the volume creation is complete, the volume is
// available for connection via iSCSI.
func (sfClient *SFClient) CreateVolume(ctx context.Context, req *CreateVolumeRequest) (*CreateVolumeResult, *SdkError) {
	var res CreateVolumeResult
	_, err := sfClient.MakeSFCall(ctx, "CreateVolume", 1, req, &res)
	return &res, err
}

// Modifies an existing cluster interface preference.
func (sfClient *SFClient) ModifyClusterInterfacePreference(ctx context.Context, req *ModifyClusterInterfacePreferenceRequest) (*ModifyClusterInterfacePreferenceResult, *SdkError) {
	var res ModifyClusterInterfacePreferenceResult
	_, err := sfClient.MakeSFCall(ctx, "ModifyClusterInterfacePreference", 1, req, &res)
	return &res, err
}

// Creates a new auth auth session for a user.
// Returns a AuthSessionInfo.
// Intended to be used by the element-auth container.
func (sfClient *SFClient) CreateAuthSession(ctx context.Context, req *CreateAuthSessionRequest) (*CreateAuthSessionResult, *SdkError) {
	var res CreateAuthSessionResult
	_, err := sfClient.MakeSFCall(ctx, "CreateAuthSession", 1, req, &res)
	return &res, err
}

// ListEvents returns events detected on the cluster, sorted from oldest to newest.
func (sfClient *SFClient) ListEvents(ctx context.Context, req *ListEventsRequest) (*ListEventsResult, *SdkError) {
	var res ListEventsResult
	_, err := sfClient.MakeSFCall(ctx, "ListEvents", 1, req, &res)
	return &res, err
}

// ListBulkVolumeJobs enables you to retrieve information about each bulk volume read or write operation that is occurring in the
// system.
func (sfClient *SFClient) ListBulkVolumeJobs(ctx context.Context) (*ListBulkVolumeJobsResult, *SdkError) {
	var res ListBulkVolumeJobsResult
	_, err := sfClient.MakeSFCall(ctx, "ListBulkVolumeJobs", 1, nil, &res)
	return &res, err
}

//
func (sfClient *SFClient) SetConstants(ctx context.Context) (*NeedsWorkSetConstantsResult, *SdkError) {
	var res NeedsWorkSetConstantsResult
	_, err := sfClient.MakeSFCall(ctx, "SetConstants", 1, nil, &res)
	return &res, err
}

// Get details about the Encryption At Rest feature.
func (sfClient *SFClient) GetEncryptionAtRestInfo(ctx context.Context) (*GetEncryptionAtRestInfoResult, *SdkError) {
	var res GetEncryptionAtRestInfoResult
	_, err := sfClient.MakeSFCall(ctx, "GetEncryptionAtRestInfo", 1, nil, &res)
	return &res, err
}

// The AbortRecoverDeadVolumes method aborts a previously-initiated RecoverDeadVolumes operation.
//
// The cluster will attempt to cancel operations on all provided volume IDs. It will return a warning for
// any volume ID that was invalid.
//
// This command requires either volumeIDs or sliceIDs but not both.
func (sfClient *SFClient) AbortRecoverDeadVolumes(ctx context.Context, req *AbortRecoverDeadVolumesRequest) (*AbortRecoverDeadVolumesResult, *SdkError) {
	var res AbortRecoverDeadVolumesResult
	_, err := sfClient.MakeSFCall(ctx, "AbortRecoverDeadVolumes", 1, req, &res)
	return &res, err
}

// You can use the GetActiveTlsCiphers method to get a list of the TLS ciphers that are currently accepted on the cluster.
func (sfClient *SFClient) GetActiveTlsCiphers(ctx context.Context) (*GetActiveTlsCiphersResult, *SdkError) {
	var res GetActiveTlsCiphersResult
	_, err := sfClient.MakeSFCall(ctx, "GetActiveTlsCiphers", 1, nil, &res)
	return &res, err
}

//
func (sfClient *SFClient) GetLocalStats(ctx context.Context) (*NeedsWorkGetLocalStatsResult, *SdkError) {
	var res NeedsWorkGetLocalStatsResult
	_, err := sfClient.MakeSFCall(ctx, "GetLocalStats", 1, nil, &res)
	return &res, err
}

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
func (sfClient *SFClient) RecoverDeadVolumes(ctx context.Context, req *RecoverDeadVolumesRequest) (*RecoverDeadVolumesResult, *SdkError) {
	var res RecoverDeadVolumesResult
	_, err := sfClient.MakeSFCall(ctx, "RecoverDeadVolumes", 1, req, &res)
	return &res, err
}

// Used to assign Nodes to user-defined Protection Domains. This information must be provided for all
// Active Nodes in the cluster, and no information may be provided for Nodes that are not Active. All Nodes
// in a given Chassis must be assigned to the same user-defined Protection Domain. The same
// ProtectionDomainType must be supplied for all nodes. ProtectionDomainTypes that are not user-defined
// such as Node and Chassis, must not be included. If any of these are not true, the Custom Protection
// Domains will be ignored, and an appropriate error will be returned.
func (sfClient *SFClient) SetProtectionDomainLayout(ctx context.Context, req *SetProtectionDomainLayoutRequest) (*SetProtectionDomainLayoutResult, *SdkError) {
	var res SetProtectionDomainLayoutResult
	_, err := sfClient.MakeSFCall(ctx, "SetProtectionDomainLayout", 1, req, &res)
	return &res, err
}

//
func (sfClient *SFClient) IsUpgradeInProgress(ctx context.Context) (*NeedsWorkIsUpgradeInProgressResult, *SdkError) {
	var res NeedsWorkIsUpgradeInProgressResult
	_, err := sfClient.MakeSFCall(ctx, "IsUpgradeInProgress", 1, nil, &res)
	return &res, err
}

// You can use the ClearClusterFaults method to clear information about both current and previously detected faults. Both resolved
// and unresolved faults can be cleared.
func (sfClient *SFClient) ClearClusterFaults(ctx context.Context, req *ClearClusterFaultsRequest) (*ClearClusterFaultsResult, *SdkError) {
	var res ClearClusterFaultsResult
	_, err := sfClient.MakeSFCall(ctx, "ClearClusterFaults", 1, req, &res)
	return &res, err
}

// GetOrigin enables you to retrieve the origination certificate for where the node was built. This method might return null if there is no origination certification.
func (sfClient *SFClient) GetOrigin(ctx context.Context, req *GetOriginRequest) (*GetOriginResult, *SdkError) {
	var res GetOriginResult
	_, err := sfClient.MakeSFCall(ctx, "GetOrigin", 1, req, &res)
	return &res, err
}

//
func (sfClient *SFClient) ForceWholeFileSync(ctx context.Context) (*NeedsWorkForceWholeFileSyncResult, *SdkError) {
	var res NeedsWorkForceWholeFileSyncResult
	_, err := sfClient.MakeSFCall(ctx, "ForceWholeFileSync", 1, nil, &res)
	return &res, err
}

// RollbackToSnapshot enables you to make an existing snapshot of the "active" volume image. This method creates a new snapshot
// from an existing snapshot. The new snapshot becomes "active" and the existing snapshot is preserved until you delete it.
// The previously "active" snapshot is deleted unless you set the parameter saveCurrentState to true.
// Note: Creating a snapshot is allowed if cluster fullness is at stage 2 or 3. Snapshots are not created when cluster fullness is
// at stage 4 or 5.
func (sfClient *SFClient) RollbackToSnapshot(ctx context.Context, req *RollbackToSnapshotRequest) (*RollbackToSnapshotResult, *SdkError) {
	var res RollbackToSnapshotResult
	_, err := sfClient.MakeSFCall(ctx, "RollbackToSnapshot", 1, req, &res)
	return &res, err
}

// You can use the TestPing API method to validate the
// connection to all the nodes in a cluster on both 1G and 10G interfaces by using ICMP packets. The test uses the appropriate MTU sizes for each packet based on the MTU settings in the network configuration.
// Note: This method is available only through the per-node API endpoint 5.0 or later.
func (sfClient *SFClient) TestPing(ctx context.Context, req *TestPingRequest) (*TestPingResult, *SdkError) {
	var res TestPingResult
	_, err := sfClient.MakeSFCall(ctx, "TestPing", 1, req, &res)
	return &res, err
}

// ListFibreChannelSessions enables you to retrieve information about the active Fibre Channel sessions on a cluster.
func (sfClient *SFClient) ListFibreChannelSessions(ctx context.Context) (*ListFibreChannelSessionsResult, *SdkError) {
	var res ListFibreChannelSessionsResult
	_, err := sfClient.MakeSFCall(ctx, "ListFibreChannelSessions", 1, nil, &res)
	return &res, err
}

// Test whether the specified KMIP (Key Management Interoperability Protocol) Key Server is functioning normally.
func (sfClient *SFClient) TestKeyServerKmip(ctx context.Context, req *TestKeyServerKmipRequest) (*TestKeyServerKmipResult, *SdkError) {
	var res TestKeyServerKmipResult
	_, err := sfClient.MakeSFCall(ctx, "TestKeyServerKmip", 1, req, &res)
	return &res, err
}

// The ModifyVolumeAccessGroupLunAssignments
// method enables you to define custom LUN assignments
// for specific volumes. This method changes only LUN
// values set on the lunAssignments parameter in the
// volume access group. All other LUN assignments remain
// unchanged. LUN assignment values must be unique for volumes in a volume access group. You cannot define duplicate LUN values within a volume access group. However, you can use the same LUN values again in different volume access groups.
// Note: Correct LUN values are 0 through 16383. The system generates an exception if you pass a LUN value outside of this range. None of the specified LUN assignments are modified if there is an exception.
// Caution: If you change a LUN assignment for a volume with active I/O, the I/O can be disrupted. You might need to change the server configuration before changing volume LUN assignments.
func (sfClient *SFClient) ModifyVolumeAccessGroupLunAssignments(ctx context.Context, req *ModifyVolumeAccessGroupLunAssignmentsRequest) (*ModifyVolumeAccessGroupLunAssignmentsResult, *SdkError) {
	var res ModifyVolumeAccessGroupLunAssignmentsResult
	_, err := sfClient.MakeSFCall(ctx, "ModifyVolumeAccessGroupLunAssignments", 1, req, &res)
	return &res, err
}

// CreateSupportBundle enables you to create a support bundle file under the node's directory. After creation, the bundle is stored on the node as a tar.gz file.
func (sfClient *SFClient) CreateSupportBundle(ctx context.Context, req *CreateSupportBundleRequest) (*CreateSupportBundleResult, *SdkError) {
	var res CreateSupportBundleResult
	_, err := sfClient.MakeSFCall(ctx, "CreateSupportBundle", 1, req, &res)
	return &res, err
}

//
func (sfClient *SFClient) GetDebugOptions(ctx context.Context) (*NeedsWorkGetDebugOptionsResult, *SdkError) {
	var res NeedsWorkGetDebugOptionsResult
	_, err := sfClient.MakeSFCall(ctx, "GetDebugOptions", 1, nil, &res)
	return &res, err
}

// ListNetworkInterfaces enables you to retrieve information about each network interface on a node. The API method is intended for use on individual nodes; userid and password authentication is required for access to individual nodes.
func (sfClient *SFClient) ListNetworkInterfaces(ctx context.Context) (*ListNetworkInterfacesResult, *SdkError) {
	var res ListNetworkInterfacesResult
	_, err := sfClient.MakeSFCall(ctx, "ListNetworkInterfaces", 1, nil, &res)
	return &res, err
}

// The SolidFire Element OS web UI uses the ListSnapMirrorVservers method to list all SnapMirror Vservers available on a remote ONTAP system.
func (sfClient *SFClient) ListSnapMirrorVservers(ctx context.Context, req *ListSnapMirrorVserversRequest) (*ListSnapMirrorVserversResult, *SdkError) {
	var res ListSnapMirrorVserversResult
	_, err := sfClient.MakeSFCall(ctx, "ListSnapMirrorVservers", 1, req, &res)
	return &res, err
}

// You can use the CloneVolumeFromDead method to create a new volume from a secondary replica
// that may not have all data acknowledged to the client. This should only be used in circumstances
// where the primary slice service is permanently inaccessible and you want to recover as much user
// data as possible.
func (sfClient *SFClient) CloneVolumeFromDead(ctx context.Context, req *CloneVolumeFromDeadRequest) (*CloneVolumeResult, *SdkError) {
	var res CloneVolumeResult
	_, err := sfClient.MakeSFCall(ctx, "CloneVolumeFromDead", 1, req, &res)
	return &res, err
}

// Enables all of the provided protection schemes.
func (sfClient *SFClient) EnableProtectionSchemes(ctx context.Context, req *EnableProtectionSchemesRequest) (*EnableProtectionSchemesResult, *SdkError) {
	var res EnableProtectionSchemesResult
	_, err := sfClient.MakeSFCall(ctx, "EnableProtectionSchemes", 1, req, &res)
	return &res, err
}

// The SetClusterConfig API method enables you to set the configuration this node uses to communicate with the cluster it is associated with. To see the states in which these objects can be modified, see Cluster Object Attributes. To display the current cluster
// interface settings for a node, run the GetClusterConfig API method.
// Note: This method is available only through the per-node API endpoint 5.0 or later.
func (sfClient *SFClient) SetClusterConfig(ctx context.Context, req *SetClusterConfigRequest) (*SetClusterConfigResult, *SdkError) {
	var res SetClusterConfigResult
	_, err := sfClient.MakeSFCall(ctx, "SetClusterConfig", 1, req, &res)
	return &res, err
}

//
func (sfClient *SFClient) FillSlice(ctx context.Context) (*NeedsWorkFillSliceResult, *SdkError) {
	var res NeedsWorkFillSliceResult
	_, err := sfClient.MakeSFCall(ctx, "FillSlice", 1, nil, &res)
	return &res, err
}

// The GetNetworkConfig API method enables you to display the network configuration information for a node.
// Note: This method is available only through the per-node API endpoint 5.0 or later.
func (sfClient *SFClient) GetNetworkConfig(ctx context.Context, req *GetNetworkConfigRequest) (*GetNetworkConfigResult, *SdkError) {
	var res GetNetworkConfigResult
	_, err := sfClient.MakeSFCall(ctx, "GetNetworkConfig", 1, req, &res)
	return &res, err
}

// SetSnmpACL enables you to configure SNMP access permissions on the cluster nodes. The values you set with this interface apply to all
// nodes in the cluster, and the values that are passed replace, in whole, all values set in any previous call to SetSnmpACL. Also note
// that the values set with this interface replace all network or usmUsers values set with the older SetSnmpInfo.
func (sfClient *SFClient) SetSnmpACL(ctx context.Context, req *SetSnmpACLRequest) (*SetSnmpACLResult, *SdkError) {
	var res SetSnmpACLResult
	_, err := sfClient.MakeSFCall(ctx, "SetSnmpACL", 1, req, &res)
	return &res, err
}

// GetNvramInfo enables you to retrieve information from each node about the NVRAM card.
func (sfClient *SFClient) GetNvramInfo(ctx context.Context, req *GetNvramInfoRequest) (*GetNvramInfoResult, *SdkError) {
	var res GetNvramInfoResult
	_, err := sfClient.MakeSFCall(ctx, "GetNvramInfo", 1, req, &res)
	return &res, err
}

//
func (sfClient *SFClient) CreateEvent(ctx context.Context) (*NeedsWorkCreateEventResult, *SdkError) {
	var res NeedsWorkCreateEventResult
	_, err := sfClient.MakeSFCall(ctx, "CreateEvent", 1, nil, &res)
	return &res, err
}

// Returns the specified KMIP (Key Management Interoperability Protocol) Key Server object.
func (sfClient *SFClient) GetKeyServerKmip(ctx context.Context, req *GetKeyServerKmipRequest) (*GetKeyServerKmipResult, *SdkError) {
	var res GetKeyServerKmipResult
	_, err := sfClient.MakeSFCall(ctx, "GetKeyServerKmip", 1, req, &res)
	return &res, err
}

// UnbindGetVirtualVolume removes the VVol <-> Host binding.
func (sfClient *SFClient) UnbindVirtualVolumes(ctx context.Context, req *UnbindVirtualVolumesRequest) (*VirtualVolumeUnbindResult, *SdkError) {
	var res VirtualVolumeUnbindResult
	_, err := sfClient.MakeSFCall(ctx, "UnbindVirtualVolumes", 1, req, &res)
	return &res, err
}

// ListClusterAdmins returns the list of all cluster administrators for the cluster. There can be
// several cluster administrator accounts with different levels of permissions.  There can be only
// one primary cluster administrator in the system. The primary Cluster Admin is the
// administrator that was created when the cluster was created.
func (sfClient *SFClient) ListClusterAdmins(ctx context.Context, req *ListClusterAdminsRequest) (*ListClusterAdminsResult, *SdkError) {
	var res ListClusterAdminsResult
	_, err := sfClient.MakeSFCall(ctx, "ListClusterAdmins", 1, req, &res)
	return &res, err
}

// The GetConfig API method enables you to retrieve all configuration information for a node. This method includes the same information available in both the GetClusterConfig and GetNetworkConfig API methods.
// Note: This method is available only through the per-node API endpoint 5.0 or later.
func (sfClient *SFClient) GetConfig(ctx context.Context, req *GetConfigRequest) (*GetConfigResult, *SdkError) {
	var res GetConfigResult
	_, err := sfClient.MakeSFCall(ctx, "GetConfig", 1, req, &res)
	return &res, err
}

// QueryVirtualVolumeMetadata returns a list of VVols matching a metadata query.
func (sfClient *SFClient) QueryVirtualVolumeMetadata(ctx context.Context, req *QueryVirtualVolumeMetadataRequest) (*QueryVirtualVolumeMetadataResult, *SdkError) {
	var res QueryVirtualVolumeMetadataResult
	_, err := sfClient.MakeSFCall(ctx, "QueryVirtualVolumeMetadata", 1, req, &res)
	return &res, err
}

// ListVolumeStatsByVolumeAccessGroup enables you to get total activity measurements for all of the volumes that are a member of the
// specified volume access group(s).
func (sfClient *SFClient) ListVolumeStatsByVolumeAccessGroup(ctx context.Context, req *ListVolumeStatsByVolumeAccessGroupRequest) (*ListVolumeStatsByVolumeAccessGroupResult, *SdkError) {
	var res ListVolumeStatsByVolumeAccessGroupResult
	_, err := sfClient.MakeSFCall(ctx, "ListVolumeStatsByVolumeAccessGroup", 1, req, &res)
	return &res, err
}

// GetNodeHardwareInfo enables you to return all the hardware information and status for the node specified. This generally includes details about
// manufacturers, vendors, versions, and other associated hardware identification information.
func (sfClient *SFClient) GetNodeHardwareInfo(ctx context.Context, req *GetNodeHardwareInfoRequest) (*GetNodeHardwareInfoResult, *SdkError) {
	var res GetNodeHardwareInfoResult
	_, err := sfClient.MakeSFCall(ctx, "GetNodeHardwareInfo", 1, req, &res)
	return &res, err
}

// You can use the StartClusterPairing method to create an encoded key from a cluster that is used to pair with another cluster. The key created from this API method is used in the CompleteClusterPairing API method to establish a cluster pairing. You can pair a cluster with a maximum of four other clusters.
func (sfClient *SFClient) StartClusterPairing(ctx context.Context) (*StartClusterPairingResult, *SdkError) {
	var res StartClusterPairingResult
	_, err := sfClient.MakeSFCall(ctx, "StartClusterPairing", 1, nil, &res)
	return &res, err
}

//
func (sfClient *SFClient) ProtocolVersionUpgrade(ctx context.Context) (*NeedsWorkProtocolVersionUpgradeResult, *SdkError) {
	var res NeedsWorkProtocolVersionUpgradeResult
	_, err := sfClient.MakeSFCall(ctx, "ProtocolVersionUpgrade", 1, nil, &res)
	return &res, err
}

// ListActiveNodes returns the list of currently active nodes that are in the cluster.
func (sfClient *SFClient) ListActiveNodes(ctx context.Context) (*ListActiveNodesResult, *SdkError) {
	var res ListActiveNodesResult
	_, err := sfClient.MakeSFCall(ctx, "ListActiveNodes", 1, nil, &res)
	return &res, err
}

// StartBulkVolumeRead enables you to initialize a bulk volume read session on a specified volume. Only two bulk volume processes
// can run simultaneously on a volume. When you initialize the session, data is read from a SolidFire storage volume for the purposes
// of storing the data on an external backup source. The external data is accessed by a web server running on an SF-series node.
// Communications and server interaction information for external data access is passed by a script running on the storage system.
// At the start of a bulk volume read operation, a snapshot of the volume is made and the snapshot is deleted when the read is complete. You can also read a snapshot of the volume by entering the ID of the snapshot as a parameter. When you read a
// previous snapshot, the system does not create a new snapshot of the volume or delete the previous snapshot when the
// read completes.
// Note: This process creates a new snapshot if the ID of an existing snapshot is not provided. Snapshots can be created if cluster fullness is at stage 2 or 3. Snapshots are not created when cluster fullness is at stage 4 or 5.
func (sfClient *SFClient) StartBulkVolumeRead(ctx context.Context, req *StartBulkVolumeReadRequest) (*StartBulkVolumeReadResult, *SdkError) {
	var res StartBulkVolumeReadResult
	_, err := sfClient.MakeSFCall(ctx, "StartBulkVolumeRead", 1, req, &res)
	return &res, err
}

// GetNtpInfo enables you to return the current network time protocol (NTP) configuration information.
func (sfClient *SFClient) GetNtpInfo(ctx context.Context) (*GetNtpInfoResult, *SdkError) {
	var res GetNtpInfoResult
	_, err := sfClient.MakeSFCall(ctx, "GetNtpInfo", 1, nil, &res)
	return &res, err
}

//
func (sfClient *SFClient) GetClusterStructure(ctx context.Context) (*NeedsWorkGetClusterStructureResult, *SdkError) {
	var res NeedsWorkGetClusterStructureResult
	_, err := sfClient.MakeSFCall(ctx, "GetClusterStructure", 1, nil, &res)
	return &res, err
}

// The SolidFire Element OS web UI uses the InitializeSnapMirrorRelationship method to initialize the destination volume in a SnapMirror relationship by performing an initial baseline transfer between clusters.
func (sfClient *SFClient) InitializeSnapMirrorRelationship(ctx context.Context, req *InitializeSnapMirrorRelationshipRequest) (*InitializeSnapMirrorRelationshipResult, *SdkError) {
	var res InitializeSnapMirrorRelationshipResult
	_, err := sfClient.MakeSFCall(ctx, "InitializeSnapMirrorRelationship", 1, req, &res)
	return &res, err
}

// GetIpmiConfig enables you to retrieve hardware sensor information from sensors that are in your node.
func (sfClient *SFClient) GetIpmiConfig(ctx context.Context, req *GetIpmiConfigRequest) (*GetIpmiConfigResult, *SdkError) {
	var res GetIpmiConfigResult
	_, err := sfClient.MakeSFCall(ctx, "GetIpmiConfig", 1, req, &res)
	return &res, err
}

// SecureEraseDrives enables you to remove any residual data from drives that have a status of "available." You might want to use this method when replacing a drive nearing the end of its service life that contained sensitive data. This method uses a Security Erase Unit command to write a predetermined pattern to the drive and resets the encryption key on the drive. This asynchronous method might take up to two minutes to complete. You can use GetAsyncResult to check on the status of the secure erase operation.
// You can use the ListDrives method to obtain the driveIDs for the drives you want to secure erase.
func (sfClient *SFClient) SecureEraseDrives(ctx context.Context, req *SecureEraseDrivesRequest) (*AsyncHandleResult, *SdkError) {
	var res AsyncHandleResult
	_, err := sfClient.MakeSFCall(ctx, "SecureEraseDrives", 1, req, &res)
	return &res, err
}

//
func (sfClient *SFClient) ModifySliceReserveUsedThresholdPct(ctx context.Context) (*NeedsWorkModifySliceReserveUsedThresholdPctResult, *SdkError) {
	var res NeedsWorkModifySliceReserveUsedThresholdPctResult
	_, err := sfClient.MakeSFCall(ctx, "ModifySliceReserveUsedThresholdPct", 1, nil, &res)
	return &res, err
}

// CreateSchedule enables you to schedule an automatic snapshot of a volume at a defined interval.
// You can use the created snapshot later as a backup or rollback to ensure the data on a volume or group of volumes is consistent for
// the point in time in which the snapshot was created.
// If you schedule a snapshot to run at a time period that is not divisible by 5 minutes, the snapshot runs at the next time period
// that is divisible by 5 minutes. For example, if you schedule a snapshot to run at 12:42:00 UTC, it runs at 12:45:00 UTC.
// Note: You can create snapshots if cluster fullness is at stage 1, 2 or 3. You cannot create snapshots after cluster fullness reaches stage 4 or 5.
func (sfClient *SFClient) CreateSchedule(ctx context.Context, req *CreateScheduleRequest) (*CreateScheduleResult, *SdkError) {
	var res CreateScheduleResult
	_, err := sfClient.MakeSFCall(ctx, "CreateSchedule", 1, req, &res)
	return &res, err
}

// The DisableStandbySliceAssignments method instructs the slice balancer to unassign all slice standbys.
// Callers should monitor the slices report to determine when all standbys are removed.
// The cluster will enable IOPS-based slice balancing when standby assignments are disabled.
func (sfClient *SFClient) DisableStandbySliceAssignments(ctx context.Context) (*DisableStandbySliceAssignmentsResult, *SdkError) {
	var res DisableStandbySliceAssignmentsResult
	_, err := sfClient.MakeSFCall(ctx, "DisableStandbySliceAssignments", 1, nil, &res)
	return &res, err
}

// You can use the ModifyQoSPolicy method to modify an existing QoSPolicy on the system.
func (sfClient *SFClient) ModifyQoSPolicy(ctx context.Context, req *ModifyQoSPolicyRequest) (*ModifyQoSPolicyResult, *SdkError) {
	var res ModifyQoSPolicyResult
	_, err := sfClient.MakeSFCall(ctx, "ModifyQoSPolicy", 1, req, &res)
	return &res, err
}

// ModifyBackupTarget enables you to change attributes of a backup target.
func (sfClient *SFClient) ModifyBackupTarget(ctx context.Context, req *ModifyBackupTargetRequest) (*ModifyBackupTargetResult, *SdkError) {
	var res ModifyBackupTargetResult
	_, err := sfClient.MakeSFCall(ctx, "ModifyBackupTarget", 1, req, &res)
	return &res, err
}

// GetClusterVersionInfo enables you to retrieve information about the Element software version running on each node in the cluster.
// This method also returns information about nodes that are currently in the process of upgrading software.
func (sfClient *SFClient) GetClusterVersionInfo(ctx context.Context) (*GetClusterVersionInfoResult, *SdkError) {
	var res GetClusterVersionInfoResult
	_, err := sfClient.MakeSFCall(ctx, "GetClusterVersionInfo", 1, nil, &res)
	return &res, err
}

//
func (sfClient *SFClient) EnableClusterSsh(ctx context.Context) (*NeedsWorkEnableClusterSshResult, *SdkError) {
	var res NeedsWorkEnableClusterSshResult
	_, err := sfClient.MakeSFCall(ctx, "EnableClusterSsh", 1, nil, &res)
	return &res, err
}

// Returns the cluster constants for a given node.
// Node cluster constants are saved in a per-node file, and take priority over cluster-wide constants (which are saved in the database).
func (sfClient *SFClient) GetNodeConstants(ctx context.Context, req *GetNodeConstantsRequest) (*interface{}, *SdkError) {
	var res interface{}
	_, err := sfClient.MakeSFCall(ctx, "GetNodeConstants", 1, req, &res)
	return &res, err
}

// The TestConnectSvip API method enables you to test the storage connection to the cluster. The test pings the SVIP using ICMP packets, and when successful, connects as an iSCSI initiator.
// Note: This method is available only through the per-node API endpoint 5.0 or later.
func (sfClient *SFClient) TestConnectSvip(ctx context.Context, req *TestConnectSvipRequest) (*TestConnectSvipResult, *SdkError) {
	var res TestConnectSvipResult
	_, err := sfClient.MakeSFCall(ctx, "TestConnectSvip", 1, req, &res)
	return &res, err
}

// You can use ListBackupTargets to retrieve information about all backup targets that have been created.
func (sfClient *SFClient) ListBackupTargets(ctx context.Context) (*ListBackupTargetsResult, *SdkError) {
	var res ListBackupTargetsResult
	_, err := sfClient.MakeSFCall(ctx, "ListBackupTargets", 1, nil, &res)
	return &res, err
}

// ModifyInitiators enables you to change the attributes of one or more existing initiators. You cannot change the name of an existing
// initiator. If you need to change the name of an initiator, delete it first with DeleteInitiators and create a new one with
// CreateInitiators.
// If ModifyInitiators fails to change one of the initiators provided in the parameter, the method returns an error and does not modify
// any initiators (no partial completion is possible).
func (sfClient *SFClient) ModifyInitiators(ctx context.Context, req *ModifyInitiatorsRequest) (*ModifyInitiatorsResult, *SdkError) {
	var res ModifyInitiatorsResult
	_, err := sfClient.MakeSFCall(ctx, "ModifyInitiators", 1, req, &res)
	return &res, err
}

// ListDriveStats enables you to retrieve high-level activity measurements for multiple drives in the cluster. By default, this method returns statistics for all drives in the cluster, and these measurements are cumulative from the addition of the drive to the cluster. Some values this method returns are specific to block drives, and some are specific to metadata drives.
func (sfClient *SFClient) ListDriveStats(ctx context.Context, req *ListDriveStatsRequest) (*ListDriveStatsResult, *SdkError) {
	var res ListDriveStatsResult
	_, err := sfClient.MakeSFCall(ctx, "ListDriveStats", 1, req, &res)
	return &res, err
}

// ListActiveVolumes enables you to return the list of active volumes currently in the system. The list of volumes is returned sorted in
// VolumeID order and can be returned in multiple parts (pages).
func (sfClient *SFClient) ListActiveVolumes(ctx context.Context, req *ListActiveVolumesRequest) (*ListActiveVolumesResult, *SdkError) {
	var res ListActiveVolumesResult
	_, err := sfClient.MakeSFCall(ctx, "ListActiveVolumes", 1, req, &res)
	return &res, err
}

// You can use the GetNodeActiveTlsCiphers method to get a list of the TLS ciphers that are currently accepted on this node.
// You can use this method on both management and storage nodes.
func (sfClient *SFClient) GetNodeActiveTlsCiphers(ctx context.Context) (*GetNodeActiveTlsCiphersResult, *SdkError) {
	var res GetNodeActiveTlsCiphersResult
	_, err := sfClient.MakeSFCall(ctx, "GetNodeActiveTlsCiphers", 1, nil, &res)
	return &res, err
}

// The SolidFire Element OS web UI uses the CreateSnapMirrorVolume method to create a volume on the remote ONTAP system.
func (sfClient *SFClient) CreateSnapMirrorVolume(ctx context.Context, req *CreateSnapMirrorVolumeRequest) (*CreateSnapMirrorVolumeResult, *SdkError) {
	var res CreateSnapMirrorVolumeResult
	_, err := sfClient.MakeSFCall(ctx, "CreateSnapMirrorVolume", 1, req, &res)
	return &res, err
}

// The SolidFire Element OS web UI uses the DeleteSnapMirrorRelationships method to remove one or more SnapMirror relationships between a source and destination endpoint.
func (sfClient *SFClient) DeleteSnapMirrorRelationships(ctx context.Context, req *DeleteSnapMirrorRelationshipsRequest) (*DeleteSnapMirrorRelationshipsResult, *SdkError) {
	var res DeleteSnapMirrorRelationshipsResult
	_, err := sfClient.MakeSFCall(ctx, "DeleteSnapMirrorRelationships", 1, req, &res)
	return &res, err
}

// ModifyVolume enables you to modify settings on an existing volume. You can make modifications to one volume at a time and
// changes take place immediately. If you do not specify QoS values when you modify a volume, they remain the same as before the modification. You can retrieve
// default QoS values for a newly created volume by running the GetDefaultQoS method.
// When you need to increase the size of a volume that is being replicated, do so in the following order to prevent replication errors:
// 1. Increase the size of the "Replication Target" volume.
// 2. Increase the size of the source or "Read / Write" volume.
// NetApp recommends that both the target and source volumes are the same size.
// Note: If you change the "access" status to locked or target, all existing iSCSI connections are terminated.
func (sfClient *SFClient) ModifyVolume(ctx context.Context, req *ModifyVolumeRequest) (*ModifyVolumeResult, *SdkError) {
	var res ModifyVolumeResult
	_, err := sfClient.MakeSFCall(ctx, "ModifyVolume", 1, req, &res)
	return &res, err
}

// You can use the GetClusterCapacity method to return the high-level capacity measurements for an entire cluster. You can use the fields returned from this method to calculate the efficiency rates that are displayed in the Element OS Web UI. You can use the following calculations in scripts to return the efficiency rates for thin provisioning, deduplication, compression, and overall efficiency.
func (sfClient *SFClient) GetClusterCapacity(ctx context.Context) (*GetClusterCapacityResult, *SdkError) {
	var res GetClusterCapacityResult
	_, err := sfClient.MakeSFCall(ctx, "GetClusterCapacity", 1, nil, &res)
	return &res, err
}

// List configurations for third party Identity Provider(s) (IdP), optionally providing an IdP metadata URL to query a specific IdP configuration information.
func (sfClient *SFClient) ListIdpConfigurations(ctx context.Context, req *ListIdpConfigurationsRequest) (*ListIdpConfigurationsResult, *SdkError) {
	var res ListIdpConfigurationsResult
	_, err := sfClient.MakeSFCall(ctx, "ListIdpConfigurations", 1, req, &res)
	return &res, err
}

// ListPendingNodes returns a list of the currently pending nodes in the system. Pending nodes are nodes that are running and configured to join the cluster, but have not yet been added via the AddNodes API method.
func (sfClient *SFClient) ListPendingNodes(ctx context.Context) (*ListPendingNodesResult, *SdkError) {
	var res ListPendingNodesResult
	_, err := sfClient.MakeSFCall(ctx, "ListPendingNodes", 1, nil, &res)
	return &res, err
}

// DeleteStorageContainers enables you to remove up to 2000 Virtual Volume (VVol) storage containers from the system at one time.
// The storage containers you remove must not contain any VVols.
func (sfClient *SFClient) DeleteStorageContainers(ctx context.Context, req *DeleteStorageContainersRequest) (*DeleteStorageContainerResult, *SdkError) {
	var res DeleteStorageContainerResult
	_, err := sfClient.MakeSFCall(ctx, "DeleteStorageContainers", 1, req, &res)
	return &res, err
}

// ListVirtualVolumeTasks returns a list of virtual volume tasks in the system.
func (sfClient *SFClient) ListVirtualVolumeTasks(ctx context.Context, req *ListVirtualVolumeTasksRequest) (*ListVirtualVolumeTasksResult, *SdkError) {
	var res ListVirtualVolumeTasksResult
	_, err := sfClient.MakeSFCall(ctx, "ListVirtualVolumeTasks", 1, req, &res)
	return &res, err
}

// Creates an authentication session based on a third party Identity Provider (IdP).
// SAML attribute statements within the SAML assertion matching multiple IdP cluster
// admin accounts will have the combined access level of those matching IdP cluster
// admin accounts.
// Returns an AuthSessionInfo.
// Intended to be used by the element-auth container.
func (sfClient *SFClient) CreateIdpAuthSession(ctx context.Context, req *CreateIdpAuthSessionRequest) (*CreateAuthSessionResult, *SdkError) {
	var res CreateAuthSessionResult
	_, err := sfClient.MakeSFCall(ctx, "CreateIdpAuthSession", 1, req, &res)
	return &res, err
}

//
func (sfClient *SFClient) GetGCStatus(ctx context.Context) (*NeedsWorkGetGCStatusResult, *SdkError) {
	var res NeedsWorkGetGCStatusResult
	_, err := sfClient.MakeSFCall(ctx, "GetGCStatus", 1, nil, &res)
	return &res, err
}

// You can use the RemoveNodeSSLCertificate method to remove the user SSL certificate and private key for the management node.
// After the certificate and private key are removed, the management node is configured to use the default certificate and private key..
func (sfClient *SFClient) RemoveNodeSSLCertificate(ctx context.Context) (*RemoveNodeSSLCertificateResult, *SdkError) {
	var res RemoveNodeSSLCertificateResult
	_, err := sfClient.MakeSFCall(ctx, "RemoveNodeSSLCertificate", 1, nil, &res)
	return &res, err
}

// You can use the ReassignSlicesForZoneTolerance method to request a tick of the zone tolerance optimizer (balancer).
func (sfClient *SFClient) ReassignSlicesForZoneTolerance(ctx context.Context) (*ReassignSlicesForZoneToleranceResult, *SdkError) {
	var res ReassignSlicesForZoneToleranceResult
	_, err := sfClient.MakeSFCall(ctx, "ReassignSlicesForZoneTolerance", 1, nil, &res)
	return &res, err
}

//
func (sfClient *SFClient) GetBinAssignments(ctx context.Context) (*NeedsWorkGetBinAssignmentsResult, *SdkError) {
	var res NeedsWorkGetBinAssignmentsResult
	_, err := sfClient.MakeSFCall(ctx, "GetBinAssignments", 1, nil, &res)
	return &res, err
}

//
func (sfClient *SFClient) ListRepositories(ctx context.Context) (*NeedsWorkListRepositoriesResult, *SdkError) {
	var res NeedsWorkListRepositoriesResult
	_, err := sfClient.MakeSFCall(ctx, "ListRepositories", 1, nil, &res)
	return &res, err
}

//
func (sfClient *SFClient) MovePrimariesAwayFromNode(ctx context.Context) (*NeedsWorkMovePrimariesAwayFromNodeResult, *SdkError) {
	var res NeedsWorkMovePrimariesAwayFromNodeResult
	_, err := sfClient.MakeSFCall(ctx, "MovePrimariesAwayFromNode", 1, nil, &res)
	return &res, err
}

// GetNodeStats enables you to retrieve the high-level activity measurements for a single node.
func (sfClient *SFClient) GetNodeStats(ctx context.Context, req *GetNodeStatsRequest) (*GetNodeStatsResult, *SdkError) {
	var res GetNodeStatsResult
	_, err := sfClient.MakeSFCall(ctx, "GetNodeStats", 1, req, &res)
	return &res, err
}

// ListVolumeQoSHistograms returns histograms detailing volume performance relative to QOS settings.
// It may take up to 5 seconds for newly created volumes to have accurate histogram data available.
func (sfClient *SFClient) ListVolumeQoSHistograms(ctx context.Context, req *ListVolumeQoSHistogramsRequest) (*ListVolumeQoSHistogramsResult, *SdkError) {
	var res ListVolumeQoSHistogramsResult
	_, err := sfClient.MakeSFCall(ctx, "ListVolumeQoSHistograms", 1, req, &res)
	return &res, err
}

//
func (sfClient *SFClient) RegenerateClusterUuid(ctx context.Context) (*NeedsWorkRegenerateClusterUuidResult, *SdkError) {
	var res NeedsWorkRegenerateClusterUuidResult
	_, err := sfClient.MakeSFCall(ctx, "RegenerateClusterUuid", 1, nil, &res)
	return &res, err
}

// The ResetNode API method enables you to reset a node to the factory settings. All data, packages (software upgrades, and so on),
// configurations, and log files are deleted from the node when you call this method. However, network settings for the node are
// preserved during this operation. Nodes that are participating in a cluster cannot be reset to the factory settings.
// The ResetNode API can only be used on nodes that are in an "Available" state. It cannot be used on nodes that are "Active" in a
// cluster, or in a "Pending" state.
// Caution: This method clears any data that is on the node. Exercise caution when using this method.
// Note: This method is available only through the per-node API endpoint 5.0 or later.
func (sfClient *SFClient) ResetNode(ctx context.Context, req *ResetNodeRequest) (*ResetNodeResult, *SdkError) {
	var res ResetNodeResult
	_, err := sfClient.MakeSFCall(ctx, "ResetNode", 1, req, &res)
	return &res, err
}

// ListVirtualVolumes enables you to list the virtual volumes currently in the system. You can use this method to list all virtual volumes,
// or only list a subset.
func (sfClient *SFClient) ListVirtualVolumes(ctx context.Context, req *ListVirtualVolumesRequest) (*ListVirtualVolumesResult, *SdkError) {
	var res ListVirtualVolumesResult
	_, err := sfClient.MakeSFCall(ctx, "ListVirtualVolumes", 1, req, &res)
	return &res, err
}

// GetProtectionDomainLayout returns all of the Protection Domain information for the cluster.
func (sfClient *SFClient) GetProtectionDomainLayout(ctx context.Context) (*GetProtectionDomainLayoutResult, *SdkError) {
	var res GetProtectionDomainLayoutResult
	_, err := sfClient.MakeSFCall(ctx, "GetProtectionDomainLayout", 1, nil, &res)
	return &res, err
}

// GetAccountEfficiency enables you to retrieve efficiency statistics about a volume account. This method returns efficiency information
// only for the account you specify as a parameter.
func (sfClient *SFClient) GetAccountEfficiency(ctx context.Context, req *GetAccountEfficiencyRequest) (*GetEfficiencyResult, *SdkError) {
	var res GetEfficiencyResult
	_, err := sfClient.MakeSFCall(ctx, "GetAccountEfficiency", 1, req, &res)
	return &res, err
}

// Disables all of the provided protection schemes.
func (sfClient *SFClient) DisableProtectionSchemes(ctx context.Context, req *DisableProtectionSchemesRequest) (*DisableProtectionSchemesResult, *SdkError) {
	var res DisableProtectionSchemesResult
	_, err := sfClient.MakeSFCall(ctx, "DisableProtectionSchemes", 1, req, &res)
	return &res, err
}

// RemoveAccount enables you to remove an existing account. You must delete and purge all volumes associated with the account
// using DeleteVolume before you can remove the account. If volumes on the account are still pending deletion, you cannot use
// RemoveAccount to remove the account.
func (sfClient *SFClient) RemoveAccount(ctx context.Context, req *RemoveAccountRequest) (*RemoveAccountResult, *SdkError) {
	var res RemoveAccountResult
	_, err := sfClient.MakeSFCall(ctx, "RemoveAccount", 1, req, &res)
	return &res, err
}

// GetClusterInfo enables you to return configuration information about the cluster.
func (sfClient *SFClient) GetClusterInfo(ctx context.Context) (*GetClusterInfoResult, *SdkError) {
	var res GetClusterInfoResult
	_, err := sfClient.MakeSFCall(ctx, "GetClusterInfo", 1, nil, &res)
	return &res, err
}

// You can use EnableFeature to enable cluster features that are disabled by default.
func (sfClient *SFClient) EnableFeature(ctx context.Context, req *EnableFeatureRequest) (*EnableFeatureResult, *SdkError) {
	var res EnableFeatureResult
	_, err := sfClient.MakeSFCall(ctx, "EnableFeature", 1, req, &res)
	return &res, err
}

// The SolidFire Element OS web UI uses the ListSnapMirrorLuns method to list the LUN information for the SnapMirror relationship from the remote ONTAP cluster.
func (sfClient *SFClient) ListSnapMirrorLuns(ctx context.Context, req *ListSnapMirrorLunsRequest) (*ListSnapMirrorLunsResult, *SdkError) {
	var res ListSnapMirrorLunsResult
	_, err := sfClient.MakeSFCall(ctx, "ListSnapMirrorLuns", 1, req, &res)
	return &res, err
}

// CreateVirtualVolumeHost creates a new ESX host.
func (sfClient *SFClient) CreateVirtualVolumeHost(ctx context.Context, req *CreateVirtualVolumeHostRequest) (*VirtualVolumeNullResult, *SdkError) {
	var res VirtualVolumeNullResult
	_, err := sfClient.MakeSFCall(ctx, "CreateVirtualVolumeHost", 1, req, &res)
	return &res, err
}

// The SolidFire Element OS web UI uses DeleteSnapMirrorEndpoints to delete one or more SnapMirror endpoints from the system.
func (sfClient *SFClient) DeleteSnapMirrorEndpoints(ctx context.Context, req *DeleteSnapMirrorEndpointsRequest) (*DeleteSnapMirrorEndpointsResult, *SdkError) {
	var res DeleteSnapMirrorEndpointsResult
	_, err := sfClient.MakeSFCall(ctx, "DeleteSnapMirrorEndpoints", 1, req, &res)
	return &res, err
}

// ComputeVolumeChecksum calculates the checksum of a volume or snapshot.
func (sfClient *SFClient) ComputeVolumeChecksum(ctx context.Context, req *ComputeVolumeChecksumRequest) (*ComputeVolumeChecksumResult, *SdkError) {
	var res ComputeVolumeChecksumResult
	_, err := sfClient.MakeSFCall(ctx, "ComputeVolumeChecksum", 1, req, &res)
	return &res, err
}

// RestoreDeletedVolume marks a deleted volume as active again. This action makes the volume immediately available for iSCSI connection.
func (sfClient *SFClient) RestoreDeletedVolume(ctx context.Context, req *RestoreDeletedVolumeRequest) (*RestoreDeletedVolumeResult, *SdkError) {
	var res RestoreDeletedVolumeResult
	_, err := sfClient.MakeSFCall(ctx, "RestoreDeletedVolume", 1, req, &res)
	return &res, err
}

// The GetLdapConfiguration method enables you to get the currently active LDAP configuration on the cluster.
func (sfClient *SFClient) GetLdapConfiguration(ctx context.Context) (*GetLdapConfigurationResult, *SdkError) {
	var res GetLdapConfigurationResult
	_, err := sfClient.MakeSFCall(ctx, "GetLdapConfiguration", 1, nil, &res)
	return &res, err
}

//
func (sfClient *SFClient) CreateSnapMirrorEndpointUnmanaged(ctx context.Context) (*NeedsWorkCreateSnapMirrorEndpointUnmanagedResult, *SdkError) {
	var res NeedsWorkCreateSnapMirrorEndpointUnmanagedResult
	_, err := sfClient.MakeSFCall(ctx, "CreateSnapMirrorEndpointUnmanaged", 1, nil, &res)
	return &res, err
}

// GetBootstrapConfig returns cluster and node information from the bootstrap configuration file. Use this API method on an individual node before it has been joined with a cluster. You can use the information this method returns in the cluster configuration interface when you create a cluster.
// If a cluster has already been created, this can be used to obtain the MVIP and SVIP addresses of the cluster.
func (sfClient *SFClient) GetBootstrapConfig(ctx context.Context) (*GetBootstrapConfigResult, *SdkError) {
	var res GetBootstrapConfigResult
	_, err := sfClient.MakeSFCall(ctx, "GetBootstrapConfig", 1, nil, &res)
	return &res, err
}

// ListVolumeStatsByAccount returns high-level activity measurements for every account. Values are summed from all the volumes owned by the account.
func (sfClient *SFClient) ListVolumeStatsByAccount(ctx context.Context, req *ListVolumeStatsByAccountRequest) (*ListVolumeStatsByAccountResult, *SdkError) {
	var res ListVolumeStatsByAccountResult
	_, err := sfClient.MakeSFCall(ctx, "ListVolumeStatsByAccount", 1, req, &res)
	return &res, err
}

//
func (sfClient *SFClient) RebalanceSlices(ctx context.Context) (*NeedsWorkRebalanceSlicesResult, *SdkError) {
	var res NeedsWorkRebalanceSlicesResult
	_, err := sfClient.MakeSFCall(ctx, "RebalanceSlices", 1, nil, &res)
	return &res, err
}

// GetVolumeCount enables you to retrieve the number of volumes currently in the system.
func (sfClient *SFClient) GetVolumeCount(ctx context.Context) (*GetVolumeCountResult, *SdkError) {
	var res GetVolumeCountResult
	_, err := sfClient.MakeSFCall(ctx, "GetVolumeCount", 1, nil, &res)
	return &res, err
}

// You can use GetAsyncResult to retrieve the result of asynchronous method calls. Some method calls require some time to run, and
// might not be finished when the system sends the initial response. To obtain the status or result of the method call, use
// GetAsyncResult to poll the asyncHandle value returned by the method.
// GetAsyncResult returns the overall status of the operation (in progress, completed, or error) in a standard fashion, but the actual
// data returned for the operation depends on the original method call and the return data is documented with each method.
func (sfClient *SFClient) GetAsyncResult(ctx context.Context, req *GetAsyncResultRequest) (*interface{}, *SdkError) {
	var res interface{}
	_, err := sfClient.MakeSFCall(ctx, "GetAsyncResult", 1, req, &res)
	return &res, err
}

// Initiate the process of setting a password on self-encrypting drives (SEDs) within the cluster.  This feature is not enabled by default but can be toggled on and off as needed.
// If a password is set on a SED which is removed from the cluster, the password will remain set and the drive is not secure erased.  Data can be secure erased using the SecureEraseDrives API method.
// Note: This does not affect performance or efficiency.
// If no parameters are specified, the password will be generated internally and at random (the only option for endpoints prior to 12.0).  This generated password will be distributed across the nodes using Shamir's Secret Sharing Algorithm such that at least two nodes are required to reconstruct the password.  The complete password to unlock the drives is not stored on any single node and is never sent across the network in its entirety.  This protects against the theft of any number of drives or a single node.
// If a keyProviderID is specified then the password will be generated/retrieved as appropriate per the type of provider.  Commonly this would be via a KMIP (Key Management Interoperability Protocol) Key Server in the case of a KMIP Key Provider (see CreateKeyProviderKmip).  After this operation the specified provider will be considered 'active' and will not be able to be deleted until DisableEncryptionAtRest is called.
func (sfClient *SFClient) EnableEncryptionAtRest(ctx context.Context, req *EnableEncryptionAtRestRequest) (*EnableEncryptionAtRestResult, *SdkError) {
	var res EnableEncryptionAtRestResult
	_, err := sfClient.MakeSFCall(ctx, "EnableEncryptionAtRest", 1, req, &res)
	return &res, err
}

// SetDefaultQoS enables you to configure the default Quality of Service (QoS) values (measured in inputs and outputs per second, or
// IOPS) for a volume. For more information about QoS in a SolidFire cluster, see the User Guide.
func (sfClient *SFClient) SetDefaultQoS(ctx context.Context, req *SetDefaultQoSRequest) (*SetDefaultQoSResult, *SdkError) {
	var res SetDefaultQoSResult
	_, err := sfClient.MakeSFCall(ctx, "SetDefaultQoS", 1, req, &res)
	return &res, err
}

// GetSystemStatus enables you to return whether a reboot ir required or not.
func (sfClient *SFClient) GetSystemStatus(ctx context.Context) (*GetSystemStatusResult, *SdkError) {
	var res GetSystemStatusResult
	_, err := sfClient.MakeSFCall(ctx, "GetSystemStatus", 1, nil, &res)
	return &res, err
}

//
func (sfClient *SFClient) StartUpgrade(ctx context.Context) (*NeedsWorkStartUpgradeResult, *SdkError) {
	var res NeedsWorkStartUpgradeResult
	_, err := sfClient.MakeSFCall(ctx, "StartUpgrade", 1, nil, &res)
	return &res, err
}

// ListNodeStats enables you to view the high-level activity measurements for all nodes in a cluster.
func (sfClient *SFClient) ListNodeStats(ctx context.Context) (*ListNodeStatsResult, *SdkError) {
	var res ListNodeStatsResult
	_, err := sfClient.MakeSFCall(ctx, "ListNodeStats", 1, nil, &res)
	return &res, err
}

// The ListVolumes method enables you to retrieve a list of volumes that are in a cluster. You can specify the volumes you want to
// return in the list by using the available parameters.
func (sfClient *SFClient) ListVolumes(ctx context.Context, req *ListVolumesRequest) (*ListVolumesResult, *SdkError) {
	var res ListVolumesResult
	_, err := sfClient.MakeSFCall(ctx, "ListVolumes", 1, req, &res)
	return &res, err
}

// ResetDrives enables you to proactively initialize drives and remove all data currently residing on a drive. The drive can then be reused
// in an existing node or used in an upgraded node. This method requires the force parameter to be included in the method call.
func (sfClient *SFClient) ResetDrives(ctx context.Context, req *ResetDrivesRequest) (*ResetDrivesResult, *SdkError) {
	var res ResetDrivesResult
	_, err := sfClient.MakeSFCall(ctx, "ResetDrives", 1, req, &res)
	return &res, err
}

//
func (sfClient *SFClient) SetAptSourceLines(ctx context.Context) (*NeedsWorkSetAptSourceLinesResult, *SdkError) {
	var res NeedsWorkSetAptSourceLinesResult
	_, err := sfClient.MakeSFCall(ctx, "SetAptSourceLines", 1, nil, &res)
	return &res, err
}

// You can use the GetSupportedTlsCiphers method to get a list of the supported TLS ciphers on this node.
// You can use this method on both management and storage nodes.
func (sfClient *SFClient) GetNodeSupportedTlsCiphers(ctx context.Context) (*GetNodeSupportedTlsCiphersResult, *SdkError) {
	var res GetNodeSupportedTlsCiphersResult
	_, err := sfClient.MakeSFCall(ctx, "GetNodeSupportedTlsCiphers", 1, nil, &res)
	return &res, err
}

// You can use the ResetNodeSupplementalTlsCiphers method to restore the supplemental ciphers to their defaults.
// You can use this command on management nodes.
func (sfClient *SFClient) ResetNodeSupplementalTlsCiphers(ctx context.Context) (*ResetNodeSupplementalTlsCiphersResult, *SdkError) {
	var res ResetNodeSupplementalTlsCiphersResult
	_, err := sfClient.MakeSFCall(ctx, "ResetNodeSupplementalTlsCiphers", 1, nil, &res)
	return &res, err
}

// StartVolumePairing enables you to create an encoded key from a volume that is used to pair with another volume. The key that this
// method creates is used in the CompleteVolumePairing API method to establish a volume pairing.
func (sfClient *SFClient) StartVolumePairing(ctx context.Context, req *StartVolumePairingRequest) (*StartVolumePairingResult, *SdkError) {
	var res StartVolumePairingResult
	_, err := sfClient.MakeSFCall(ctx, "StartVolumePairing", 1, req, &res)
	return &res, err
}

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
func (sfClient *SFClient) RemoveClusterAdmin(ctx context.Context, req *RemoveClusterAdminRequest) (*RemoveClusterAdminResult, *SdkError) {
	var res RemoveClusterAdminResult
	_, err := sfClient.MakeSFCall(ctx, "RemoveClusterAdmin", 1, req, &res)
	return &res, err
}

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
func (sfClient *SFClient) DeleteVolumes(ctx context.Context, req *DeleteVolumesRequest) (*DeleteVolumesResult, *SdkError) {
	var res DeleteVolumesResult
	_, err := sfClient.MakeSFCall(ctx, "DeleteVolumes", 1, req, &res)
	return &res, err
}

// The SolidFire Element OS web UI uses the ListSnapMirrorRelationships method to list one or all SnapMirror relationships on a SolidFire cluster
func (sfClient *SFClient) ListSnapMirrorRelationships(ctx context.Context, req *ListSnapMirrorRelationshipsRequest) (*ListSnapMirrorRelationshipsResult, *SdkError) {
	var res ListSnapMirrorRelationshipsResult
	_, err := sfClient.MakeSFCall(ctx, "ListSnapMirrorRelationships", 1, req, &res)
	return &res, err
}

// RemoveVirtualNetwork enables you to remove a previously added virtual network.
// Note: This method requires either the virtualNetworkID or the virtualNetworkTag as a parameter, but not both.
func (sfClient *SFClient) RemoveVirtualNetwork(ctx context.Context, req *RemoveVirtualNetworkRequest) (*RemoveVirtualNetworkResult, *SdkError) {
	var res RemoveVirtualNetworkResult
	_, err := sfClient.MakeSFCall(ctx, "RemoveVirtualNetwork", 1, req, &res)
	return &res, err
}

// The DeleteStandbySlices method deactivates any active or failed standby for the given slices and deletes
// the associated metadata. The command is a no-op if the standby is already inactive.
func (sfClient *SFClient) DeleteStandbySlices(ctx context.Context, req *DeleteStandbySlicesRequest) (*DeleteStandbySlicesResult, *SdkError) {
	var res DeleteStandbySlicesResult
	_, err := sfClient.MakeSFCall(ctx, "DeleteStandbySlices", 1, req, &res)
	return &res, err
}

// GetDefaultQoS enables you to retrieve the default QoS values for a newly created volume.
func (sfClient *SFClient) GetDefaultQoS(ctx context.Context) (*VolumeQOS, *SdkError) {
	var res VolumeQOS
	_, err := sfClient.MakeSFCall(ctx, "GetDefaultQoS", 1, nil, &res)
	return &res, err
}

// Lists all expired auth sessions.
// This is only callable by a user with Administrative access rights.
func (sfClient *SFClient) ListExpiredAuthSessions(ctx context.Context) (*ListAuthSessionsResult, *SdkError) {
	var res ListAuthSessionsResult
	_, err := sfClient.MakeSFCall(ctx, "ListExpiredAuthSessions", 1, nil, &res)
	return &res, err
}

//
func (sfClient *SFClient) GetSliceReserveUsedThresholdPct(ctx context.Context) (*NeedsWorkGetSliceReserveUsedThresholdPctResult, *SdkError) {
	var res NeedsWorkGetSliceReserveUsedThresholdPctResult
	_, err := sfClient.MakeSFCall(ctx, "GetSliceReserveUsedThresholdPct", 1, nil, &res)
	return &res, err
}

// Return information regarding the state of authentication using third party Identity Providers
func (sfClient *SFClient) GetIdpAuthenticationState(ctx context.Context) (*GetIdpAuthenticationStateResult, *SdkError) {
	var res GetIdpAuthenticationStateResult
	_, err := sfClient.MakeSFCall(ctx, "GetIdpAuthenticationState", 1, nil, &res)
	return &res, err
}

//
func (sfClient *SFClient) ModifySnapMirrorEndpointUnmanaged(ctx context.Context) (*NeedsWorkModifySnapMirrorEndpointUnmanagedResult, *SdkError) {
	var res NeedsWorkModifySnapMirrorEndpointUnmanagedResult
	_, err := sfClient.MakeSFCall(ctx, "ModifySnapMirrorEndpointUnmanaged", 1, nil, &res)
	return &res, err
}

// You can use the ListTests API method to return the tests that are available to run on a node.
// Note: This method is available only through the per-node API endpoint 5.0 or later.
func (sfClient *SFClient) ListTests(ctx context.Context, req *ListTestsRequest) (*ListTestsResult, *SdkError) {
	var res ListTestsResult
	_, err := sfClient.MakeSFCall(ctx, "ListTests", 1, req, &res)
	return &res, err
}

// GetVirtualVolumeAllocatedBitmap returns a b64-encoded block of data
// representing a bitmap where non-zero bits indicate that data is not the same
// between two volumes for a common segment (LBA range) of the volumes.
func (sfClient *SFClient) GetVirtualVolumeUnsharedBitmap(ctx context.Context, req *GetVirtualVolumeUnsharedBitmapRequest) (*VirtualVolumeBitmapResult, *SdkError) {
	var res VirtualVolumeBitmapResult
	_, err := sfClient.MakeSFCall(ctx, "GetVirtualVolumeUnsharedBitmap", 1, req, &res)
	return &res, err
}

// ListCurrentClusterAdmins returns information for all the cluster admins the calling user is assigned to.
// Intended to be called by element-auth container.
func (sfClient *SFClient) ListCurrentClusterAdmins(ctx context.Context) (*ListCurrentClusterAdminsResult, *SdkError) {
	var res ListCurrentClusterAdminsResult
	_, err := sfClient.MakeSFCall(ctx, "ListCurrentClusterAdmins", 1, nil, &res)
	return &res, err
}

// You can use ListServices to return the services information for nodes, drives, current software, and other services that are running on the cluster.
func (sfClient *SFClient) ListServices(ctx context.Context) (*ListServicesResult, *SdkError) {
	var res ListServicesResult
	_, err := sfClient.MakeSFCall(ctx, "ListServices", 1, nil, &res)
	return &res, err
}

// RollbackToGroupSnapshot enables you to roll back all individual volumes in a snapshot group to each volume's individual snapshot.
// Note: Rolling back to a group snapshot creates a temporary snapshot of each volume within the group snapshot.
// Snapshots are allowed if cluster fullness is at stage 2 or 3. Snapshots are not created when cluster fullness is at stage 4 or 5.
func (sfClient *SFClient) RollbackToGroupSnapshot(ctx context.Context, req *RollbackToGroupSnapshotRequest) (*RollbackToGroupSnapshotResult, *SdkError) {
	var res RollbackToGroupSnapshotResult
	_, err := sfClient.MakeSFCall(ctx, "RollbackToGroupSnapshot", 1, req, &res)
	return &res, err
}

// You can use the ResetSupplementalTlsCiphers method to restore the supplemental ciphers to their defaults.
func (sfClient *SFClient) ResetSupplementalTlsCiphers(ctx context.Context) (*ResetSupplementalTlsCiphersResult, *SdkError) {
	var res ResetSupplementalTlsCiphersResult
	_, err := sfClient.MakeSFCall(ctx, "ResetSupplementalTlsCiphers", 1, nil, &res)
	return &res, err
}

// Update the Vasa Provider info
func (sfClient *SFClient) ModifyVasaProviderInfo(ctx context.Context, req *ModifyVasaProviderInfoRequest) (*VirtualVolumeNullResult, *SdkError) {
	var res VirtualVolumeNullResult
	_, err := sfClient.MakeSFCall(ctx, "ModifyVasaProviderInfo", 1, req, &res)
	return &res, err
}

// The EnableStandbySliceAssignments method instructs the slice balancer to assign a standby to each slice.
// Callers should call CheckVolumeStandbysAssigned or monitor the slices report to determine when
// the assignments are complete.
// The cluster will disable IOPS-based slice balancing while standby assignments are enabled.
func (sfClient *SFClient) EnableStandbySliceAssignments(ctx context.Context) (*EnableStandbySliceAssignmentsResult, *SdkError) {
	var res EnableStandbySliceAssignmentsResult
	_, err := sfClient.MakeSFCall(ctx, "EnableStandbySliceAssignments", 1, nil, &res)
	return &res, err
}

// The SolidFire Element OS web UI uses the BreakSnapMirrorRelationship method to break a SnapMirror relationship. When a SnapMirror relationship is broken, the destination volume is made read-write and independent, and can then diverge from the source. You can reestablish the relationship with the ResyncSnapMirrorRelationship API method. This method requires the ONTAP cluster to be available.
func (sfClient *SFClient) BreakSnapMirrorRelationship(ctx context.Context, req *BreakSnapMirrorRelationshipRequest) (*BreakSnapMirrorRelationshipResult, *SdkError) {
	var res BreakSnapMirrorRelationshipResult
	_, err := sfClient.MakeSFCall(ctx, "BreakSnapMirrorRelationship", 1, req, &res)
	return &res, err
}

// Adds (assigns) the specified KMIP (Key Management Interoperability Protocol) Key Server to the specified Key Provider.  This will result in contacting the server to verify it's functional, as well as to synchronize keys in the event that there are multiple key servers assigned to the provider.  This synchronization may result in conflicts which could cause this to fail.  If the specified KMIP Key Server is already assigned to the specified Key Provider, this is a no-op and no error will be returned.  The assignment can be removed (unassigned) using RemoveKeyServerFromProviderKmip.
func (sfClient *SFClient) AddKeyServerToProviderKmip(ctx context.Context, req *AddKeyServerToProviderKmipRequest) (*AddKeyServerToProviderKmipResult, *SdkError) {
	var res AddKeyServerToProviderKmipResult
	_, err := sfClient.MakeSFCall(ctx, "AddKeyServerToProviderKmip", 1, req, &res)
	return &res, err
}

// The SolidFire Element OS web UI uses the ListSnapMirrorNetworkInterfaces method to list all available SnapMirror interfaces on a remote ONTAP system
func (sfClient *SFClient) ListSnapMirrorNetworkInterfaces(ctx context.Context, req *ListSnapMirrorNetworkInterfacesRequest) (*ListSnapMirrorNetworkInterfacesResult, *SdkError) {
	var res ListSnapMirrorNetworkInterfacesResult
	_, err := sfClient.MakeSFCall(ctx, "ListSnapMirrorNetworkInterfaces", 1, req, &res)
	return &res, err
}

//
func (sfClient *SFClient) DeleteVirtualVolumeHosts(ctx context.Context) (*NeedsWorkDeleteVirtualVolumeHostsResult, *SdkError) {
	var res NeedsWorkDeleteVirtualVolumeHostsResult
	_, err := sfClient.MakeSFCall(ctx, "DeleteVirtualVolumeHosts", 1, nil, &res)
	return &res, err
}

// ListVolumeAccessGroups enables you to return
// information about the volume access groups that are
// currently in the system.
func (sfClient *SFClient) ListVolumeAccessGroups(ctx context.Context, req *ListVolumeAccessGroupsRequest) (*ListVolumeAccessGroupsResult, *SdkError) {
	var res ListVolumeAccessGroupsResult
	_, err := sfClient.MakeSFCall(ctx, "ListVolumeAccessGroups", 1, req, &res)
	return &res, err
}

// GetVolumeEfficiency enables you to retrieve information about a volume. Only the volume you give as a parameter in this API method is used to compute the capacity.
func (sfClient *SFClient) GetVolumeEfficiency(ctx context.Context, req *GetVolumeEfficiencyRequest) (*GetVolumeEfficiencyResult, *SdkError) {
	var res GetVolumeEfficiencyResult
	_, err := sfClient.MakeSFCall(ctx, "GetVolumeEfficiency", 1, req, &res)
	return &res, err
}

// PurgeDeletedVolumes immediately and permanently purges volumes that have been deleted.
// You can use this method to purge up to 500 volumes at one time.
// You must delete volumes using DeleteVolumes before they can be purged.
// Volumes are purged by the system automatically after a period of time, so usage of this method is not typically required.
func (sfClient *SFClient) PurgeDeletedVolumes(ctx context.Context, req *PurgeDeletedVolumesRequest) (*PurgeDeletedVolumesResult, *SdkError) {
	var res PurgeDeletedVolumesResult
	_, err := sfClient.MakeSFCall(ctx, "PurgeDeletedVolumes", 1, req, &res)
	return &res, err
}

// Deletes all auth sessions associated with the specified ClusterAdminID.
// If the specified ClusterAdminID maps to a group of users, all auth sessions for all members of that group will be deleted.
// To see the list of sessions that could be deleted, use ListAuthSessionsByClusterAdmin with the same parameter.
func (sfClient *SFClient) DeleteAuthSessionsByClusterAdmin(ctx context.Context, req *DeleteAuthSessionsByClusterAdminRequest) (*DeleteAuthSessionsResult, *SdkError) {
	var res DeleteAuthSessionsResult
	_, err := sfClient.MakeSFCall(ctx, "DeleteAuthSessionsByClusterAdmin", 1, req, &res)
	return &res, err
}

// The SolidFire Element OS web UI uses the UpdateSnapMirrorRelationship method to make the destination volume in a SnapMirror relationship an up-to-date mirror of the source volume.
func (sfClient *SFClient) UpdateSnapMirrorRelationship(ctx context.Context, req *UpdateSnapMirrorRelationshipRequest) (*UpdateSnapMirrorRelationshipResult, *SdkError) {
	var res UpdateSnapMirrorRelationshipResult
	_, err := sfClient.MakeSFCall(ctx, "UpdateSnapMirrorRelationship", 1, req, &res)
	return &res, err
}

//
func (sfClient *SFClient) SetInitiatorID(ctx context.Context) (*NeedsWorkSetInitiatorIDResult, *SdkError) {
	var res NeedsWorkSetInitiatorIDResult
	_, err := sfClient.MakeSFCall(ctx, "SetInitiatorID", 1, nil, &res)
	return &res, err
}

// The SolidFire Element OS web UI uses the BreakSnapMirrorVolume method to break the SnapMirror relationship between an ONTAP source container and SolidFire target volume. Breaking a SolidFire SnapMirror volume is useful if an ONTAP system becomes unavailable while replicating data to a SolidFire volume. This feature enables a storage administrator to take control of a SolidFire SnapMirror volume, break its relationship with the remote ONTAP system, and revert the volume to a previous snapshot.
func (sfClient *SFClient) BreakSnapMirrorVolume(ctx context.Context, req *BreakSnapMirrorVolumeRequest) (*BreakSnapMirrorVolumeResult, *SdkError) {
	var res BreakSnapMirrorVolumeResult
	_, err := sfClient.MakeSFCall(ctx, "BreakSnapMirrorVolume", 1, req, &res)
	return &res, err
}

// Returns the cluster constants for the cluster.
func (sfClient *SFClient) GetConstants(ctx context.Context, req *GetConstantsRequest) (*interface{}, *SdkError) {
	var res interface{}
	_, err := sfClient.MakeSFCall(ctx, "GetConstants", 1, req, &res)
	return &res, err
}

//
func (sfClient *SFClient) ReassignSlice(ctx context.Context) (*NeedsWorkReassignSliceResult, *SdkError) {
	var res NeedsWorkReassignSliceResult
	_, err := sfClient.MakeSFCall(ctx, "ReassignSlice", 1, nil, &res)
	return &res, err
}

// RemoveInitiatorsFromVolumeAccessGroup enables
// you to remove initiators from a specified volume access
// group.
func (sfClient *SFClient) RemoveInitiatorsFromVolumeAccessGroup(ctx context.Context, req *RemoveInitiatorsFromVolumeAccessGroupRequest) (*ModifyVolumeAccessGroupResult, *SdkError) {
	var res ModifyVolumeAccessGroupResult
	_, err := sfClient.MakeSFCall(ctx, "RemoveInitiatorsFromVolumeAccessGroup", 1, req, &res)
	return &res, err
}

// CreateMultipleVolumes enables you to create new (empty) volumes on the cluster.
// As soon as creation is complete, the volumes are available for connection via iSCSI.
func (sfClient *SFClient) CreateMultipleVolumes(ctx context.Context, req *CreateMultipleVolumesRequest) (*CreateMultipleVolumesResult, *SdkError) {
	var res CreateMultipleVolumesResult
	_, err := sfClient.MakeSFCall(ctx, "CreateMultipleVolumes", 1, req, &res)
	return &res, err
}

// Sets LLDP configuration options. If an option isn't set in the request, then it is unchanged from the previous value.
func (sfClient *SFClient) SetLldpConfig(ctx context.Context, req *SetLldpConfigRequest) (*GetLldpConfigResult, *SdkError) {
	var res GetLldpConfigResult
	_, err := sfClient.MakeSFCall(ctx, "SetLldpConfig", 1, req, &res)
	return &res, err
}

//
func (sfClient *SFClient) GetThreadBacktraces(ctx context.Context) (*NeedsWorkGetThreadBacktracesResult, *SdkError) {
	var res NeedsWorkGetThreadBacktracesResult
	_, err := sfClient.MakeSFCall(ctx, "GetThreadBacktraces", 1, nil, &res)
	return &res, err
}

// CheckProposedCluster validates that creating a cluster from a given set of nodes is likely to succeed.  Any problems with the proposed cluster are returned as errors with a human-readable description and unique error code.
func (sfClient *SFClient) CheckProposedCluster(ctx context.Context, req *CheckProposedClusterRequest) (*CheckProposedResult, *SdkError) {
	var res CheckProposedResult
	_, err := sfClient.MakeSFCall(ctx, "CheckProposedCluster", 1, req, &res)
	return &res, err
}

// You can use the RemoveClusterPair method to close the open connections between two paired clusters.
// Note: Before you remove a cluster pair, you must first remove all volume pairing to the clusters with the "RemoveVolumePair" API method.
func (sfClient *SFClient) RemoveClusterPair(ctx context.Context, req *RemoveClusterPairRequest) (*RemoveClusterPairResult, *SdkError) {
	var res RemoveClusterPairResult
	_, err := sfClient.MakeSFCall(ctx, "RemoveClusterPair", 1, req, &res)
	return &res, err
}

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
func (sfClient *SFClient) DeleteVirtualVolumes(ctx context.Context, req *DeleteVirtualVolumesRequest) (*VirtualVolumeNullResult, *SdkError) {
	var res VirtualVolumeNullResult
	_, err := sfClient.MakeSFCall(ctx, "DeleteVirtualVolumes", 1, req, &res)
	return &res, err
}

// ListClusterFaults enables you to retrieve information about any faults detected on the cluster. With this method, you can retrieve both current faults as well as faults that have been resolved. The system caches faults every 30 seconds.
func (sfClient *SFClient) ListClusterFaults(ctx context.Context, req *ListClusterFaultsRequest) (*ListClusterFaultsResult, *SdkError) {
	var res ListClusterFaultsResult
	_, err := sfClient.MakeSFCall(ctx, "ListClusterFaults", 1, req, &res)
	return &res, err
}

// ListStorageContainers enables you to retrieve information about all virtual volume storage containers known to the system.
func (sfClient *SFClient) ListStorageContainers(ctx context.Context, req *ListStorageContainersRequest) (*ListStorageContainersResult, *SdkError) {
	var res ListStorageContainersResult
	_, err := sfClient.MakeSFCall(ctx, "ListStorageContainers", 1, req, &res)
	return &res, err
}

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
func (sfClient *SFClient) DeleteVolume(ctx context.Context, req *DeleteVolumeRequest) (*DeleteVolumeResult, *SdkError) {
	var res DeleteVolumeResult
	_, err := sfClient.MakeSFCall(ctx, "DeleteVolume", 1, req, &res)
	return &res, err
}

// GetVolumeAccessGroupEfficiency enables you to
// retrieve efficiency information about a volume access
// group. Only the volume access group you provide as the
// parameter in this API method is used to compute the
// capacity.
func (sfClient *SFClient) GetVolumeAccessGroupEfficiency(ctx context.Context, req *GetVolumeAccessGroupEfficiencyRequest) (*GetEfficiencyResult, *SdkError) {
	var res GetEfficiencyResult
	_, err := sfClient.MakeSFCall(ctx, "GetVolumeAccessGroupEfficiency", 1, req, &res)
	return &res, err
}

// RemoveNodes can be used to remove one or more nodes from the cluster. Before removing a node, you must remove all drives from the node using the RemoveDrives method. You cannot remove a node until the RemoveDrives process has completed and all data has been migrated off of the node's drives.
// After removing a node, the removed node registers itself as a pending node. You can add the pending node again or shut it down (shutting the node down removes it from the Pending Node list).
//
// RemoveNodes can fail with xEnsembleInvalidSize if removing the nodes would violate ensemble size restrictions.
// RemoveNodes can fail with xEnsembleQuorumLoss if removing the nodes would cause a loss of quorum.
// RemoveNodes can fail with xEnsembleToleranceChange if there are enabled data protection schemes that can tolerate multiple node failures and removing the nodes would decrease the ensemble's node failure tolerance. The optional ignoreEnsembleToleranceChange parameter can be set true to disable the ensemble tolerance check.
func (sfClient *SFClient) RemoveNodes(ctx context.Context, req *RemoveNodesRequest) (*RemoveNodesResult, *SdkError) {
	var res RemoveNodesResult
	_, err := sfClient.MakeSFCall(ctx, "RemoveNodes", 1, req, &res)
	return &res, err
}

// The SolidFire Element OS web UI uses the ResyncSnapMirrorRelationship method to establish or reestablish a mirror relationship between a source and destination endpoint. When you resync a relationship, the system removes snapshots on the destination volume that are newer than the common snapshot copy, and then mounts the destination volume as a data protection volume with the common snapshot copy as the exported snapshot copy.
func (sfClient *SFClient) ResyncSnapMirrorRelationship(ctx context.Context, req *ResyncSnapMirrorRelationshipRequest) (*ResyncSnapMirrorRelationshipResult, *SdkError) {
	var res ResyncSnapMirrorRelationshipResult
	_, err := sfClient.MakeSFCall(ctx, "ResyncSnapMirrorRelationship", 1, req, &res)
	return &res, err
}

// Used to assign Nodes to user-defined Chassis Protection Domains for test purposes.  Overrides existing
// ChassisNames. This information must be provided for all Active Nodes in the cluster, and no information may
// be provided for Nodes that are not Active. The same ProtectionDomainType must be supplied for all nodes.
// ProtectionDomainTypes other than Chassis must not be included. If any of these are not true, the Chassis
// Protection Domains will be ignored, and an appropriate error will be returned.
func (sfClient *SFClient) SetProtectionDomainLayoutChassisOverride(ctx context.Context, req *SetProtectionDomainLayoutChassisOverrideRequest) (*SetProtectionDomainLayoutChassisOverrideResult, *SdkError) {
	var res SetProtectionDomainLayoutChassisOverrideResult
	_, err := sfClient.MakeSFCall(ctx, "SetProtectionDomainLayoutChassisOverride", 1, req, &res)
	return &res, err
}

//
func (sfClient *SFClient) CreateClusterFault(ctx context.Context) (*NeedsWorkCreateClusterFaultResult, *SdkError) {
	var res NeedsWorkCreateClusterFaultResult
	_, err := sfClient.MakeSFCall(ctx, "CreateClusterFault", 1, nil, &res)
	return &res, err
}

// ModifyGroupSnapshot enables you to change the attributes of a group of snapshots. You can also use this method to enable snapshots created on the Read/Write (source) volume to be remotely replicated to a target SolidFire storage system.
func (sfClient *SFClient) ModifyGroupSnapshot(ctx context.Context, req *ModifyGroupSnapshotRequest) (*ModifyGroupSnapshotResult, *SdkError) {
	var res ModifyGroupSnapshotResult
	_, err := sfClient.MakeSFCall(ctx, "ModifyGroupSnapshot", 1, req, &res)
	return &res, err
}

// The GetClusterState API method enables you to indicate if a node is part of a cluster or not. The three states are:
// Available: Node has not been configured with a cluster name.
// Pending: Node is pending for a specific named cluster and can be added.
// Active: Node is an active member of a cluster and may not be added to another
// cluster.
// Note: This method is available only through the per-node API endpoint 5.0 or later.
func (sfClient *SFClient) GetClusterState(ctx context.Context, req *GetClusterStateRequest) (*GetClusterStateResult, *SdkError) {
	var res GetClusterStateResult
	_, err := sfClient.MakeSFCall(ctx, "GetClusterState", 1, req, &res)
	return &res, err
}

//
func (sfClient *SFClient) GetVolumeSetEfficiency(ctx context.Context) (*NeedsWorkGetVolumeSetEfficiencyResult, *SdkError) {
	var res NeedsWorkGetVolumeSetEfficiencyResult
	_, err := sfClient.MakeSFCall(ctx, "GetVolumeSetEfficiency", 1, nil, &res)
	return &res, err
}

//
func (sfClient *SFClient) ListCloneJobs(ctx context.Context) (*NeedsWorkListCloneJobsResult, *SdkError) {
	var res NeedsWorkListCloneJobsResult
	_, err := sfClient.MakeSFCall(ctx, "ListCloneJobs", 1, nil, &res)
	return &res, err
}

//
func (sfClient *SFClient) SetRsyslogInfo(ctx context.Context) (*NeedsWorkSetRsyslogInfoResult, *SdkError) {
	var res NeedsWorkSetRsyslogInfoResult
	_, err := sfClient.MakeSFCall(ctx, "SetRsyslogInfo", 1, nil, &res)
	return &res, err
}

// Promotes a new node to be cluster master.
// The original cluster master will return success if it initiates the change. Callers should confirm that the
// change occurs by calling the API command GetClusterMasterNodeID after a delay of a few seconds.
func (sfClient *SFClient) PromoteClusterMaster(ctx context.Context, req *PromoteClusterMasterRequest) (*PromoteClusterMasterResult, *SdkError) {
	var res PromoteClusterMasterResult
	_, err := sfClient.MakeSFCall(ctx, "PromoteClusterMaster", 1, req, &res)
	return &res, err
}

//
func (sfClient *SFClient) CreateDatabaseEntry(ctx context.Context) (*NeedsWorkCreateDatabaseEntryResult, *SdkError) {
	var res NeedsWorkCreateDatabaseEntryResult
	_, err := sfClient.MakeSFCall(ctx, "CreateDatabaseEntry", 1, nil, &res)
	return &res, err
}

//
func (sfClient *SFClient) GetReport(ctx context.Context) (*NeedsWorkGetReportResult, *SdkError) {
	var res NeedsWorkGetReportResult
	_, err := sfClient.MakeSFCall(ctx, "GetReport", 1, nil, &res)
	return &res, err
}

// CancelVirtualVolumeTask attempts to cancel the VVol Async Task.
func (sfClient *SFClient) CancelVirtualVolumeTask(ctx context.Context, req *CancelVirtualVolumeTaskRequest) (*VirtualVolumeNullResult, *SdkError) {
	var res VirtualVolumeNullResult
	_, err := sfClient.MakeSFCall(ctx, "CancelVirtualVolumeTask", 1, req, &res)
	return &res, err
}

// Update an existing configuration with a third party Identity Provider (IdP) for the cluster.
func (sfClient *SFClient) UpdateIdpConfiguration(ctx context.Context, req *UpdateIdpConfigurationRequest) (*UpdateIdpConfigurationResult, *SdkError) {
	var res UpdateIdpConfigurationResult
	_, err := sfClient.MakeSFCall(ctx, "UpdateIdpConfiguration", 1, req, &res)
	return &res, err
}

// You can use the GetQoSPolicy method to get details about a specific QoSPolicy from the system.
func (sfClient *SFClient) GetQoSPolicy(ctx context.Context, req *GetQoSPolicyRequest) (*GetQoSPolicyResult, *SdkError) {
	var res GetQoSPolicyResult
	_, err := sfClient.MakeSFCall(ctx, "GetQoSPolicy", 1, req, &res)
	return &res, err
}

// NetApp engineering uses the GetRawStats API method to troubleshoot new features. The data returned from GetRawStats is not documented, changes frequently, and is not guaranteed to be accurate. NetApp does not recommend using GetCompleteStats for collecting performance data or any other
// management integration with a SolidFire cluster.
func (sfClient *SFClient) GetRawStats(ctx context.Context) (*interface{}, *SdkError) {
	var res interface{}
	_, err := sfClient.MakeSFCall(ctx, "GetRawStats", 1, nil, &res)
	return &res, err
}

// Creates SSL public and private keys. These keys can be used to generate Certificate Sign Requests.
// There can be only one key pair in use for the cluster. To replace the existing keys, make sure that they are not being used by any providers before invoking this API.
func (sfClient *SFClient) CreatePublicPrivateKeyPair(ctx context.Context, req *CreatePublicPrivateKeyPairRequest) (*CreatePublicPrivateKeyPairResult, *SdkError) {
	var res CreatePublicPrivateKeyPairResult
	_, err := sfClient.MakeSFCall(ctx, "CreatePublicPrivateKeyPair", 1, req, &res)
	return &res, err
}

// The TestConnectMvip API method enables you to test the
// management connection to the cluster. The test pings the MVIP and executes a simple API method to verify connectivity.
// Note: This method is available only through the per-node API endpoint 5.0 or later.
func (sfClient *SFClient) TestConnectMvip(ctx context.Context, req *TestConnectMvipRequest) (*TestConnectMvipResult, *SdkError) {
	var res TestConnectMvipResult
	_, err := sfClient.MakeSFCall(ctx, "TestConnectMvip", 1, req, &res)
	return &res, err
}

// GetVirtualVolumeAllocatedBitmap returns a b64-encoded block of data
// representing a bitmap where non-zero bits indicate the allocation of a
// segment (LBA range) of the volume.
func (sfClient *SFClient) GetVirtualVolumeAllocatedBitmap(ctx context.Context, req *GetVirtualVolumeAllocatedBitmapRequest) (*VirtualVolumeBitmapResult, *SdkError) {
	var res VirtualVolumeBitmapResult
	_, err := sfClient.MakeSFCall(ctx, "GetVirtualVolumeAllocatedBitmap", 1, req, &res)
	return &res, err
}

//
func (sfClient *SFClient) PairCluster(ctx context.Context) (*NeedsWorkPairClusterResult, *SdkError) {
	var res NeedsWorkPairClusterResult
	_, err := sfClient.MakeSFCall(ctx, "PairCluster", 1, nil, &res)
	return &res, err
}

//
func (sfClient *SFClient) PurgeAllDeletedVolumes(ctx context.Context) (*NeedsWorkPurgeAllDeletedVolumesResult, *SdkError) {
	var res NeedsWorkPurgeAllDeletedVolumesResult
	_, err := sfClient.MakeSFCall(ctx, "PurgeAllDeletedVolumes", 1, nil, &res)
	return &res, err
}

// Returns the list of KMIP (Key Management Interoperability Protocol) Key Servers which have been created via CreateKeyServerKmip.  The list can optionally be filtered by specifying additional parameters.
func (sfClient *SFClient) ListKeyServersKmip(ctx context.Context, req *ListKeyServersKmipRequest) (*ListKeyServersKmipResult, *SdkError) {
	var res ListKeyServersKmipResult
	_, err := sfClient.MakeSFCall(ctx, "ListKeyServersKmip", 1, req, &res)
	return &res, err
}

// The SolidFire Element OS web UI uses the QuiesceSnapMirrorRelationship method to disable future data transfers for a SnapMirror relationship. If a transfer is in progress, the relationship status becomes "quiescing" until the transfer is complete. If the current transfer is aborted, it will not restart. You can reenable data transfers for the relationship using the ResumeSnapMirrorRelationship API method.
func (sfClient *SFClient) QuiesceSnapMirrorRelationship(ctx context.Context, req *QuiesceSnapMirrorRelationshipRequest) (*QuiesceSnapMirrorRelationshipResult, *SdkError) {
	var res QuiesceSnapMirrorRelationshipResult
	_, err := sfClient.MakeSFCall(ctx, "QuiesceSnapMirrorRelationship", 1, req, &res)
	return &res, err
}

// NetApp engineering uses the GetCompleteStats API method to troubleshoot new features. The data returned from GetCompleteStats is not documented, changes frequently, and is not guaranteed to be accurate. NetApp does not recommend using GetCompleteStats for collecting performance data or any other
// management integration with a SolidFire cluster.
func (sfClient *SFClient) GetCompleteStats(ctx context.Context) (*interface{}, *SdkError) {
	var res interface{}
	_, err := sfClient.MakeSFCall(ctx, "GetCompleteStats", 1, nil, &res)
	return &res, err
}

// ListGroupSnapshots enables you to get information about all group snapshots that have been created.
func (sfClient *SFClient) ListGroupSnapshots(ctx context.Context, req *ListGroupSnapshotsRequest) (*ListGroupSnapshotsResult, *SdkError) {
	var res ListGroupSnapshotsResult
	_, err := sfClient.MakeSFCall(ctx, "ListGroupSnapshots", 1, req, &res)
	return &res, err
}

//
func (sfClient *SFClient) GetSliceFileSizeReport(ctx context.Context) (*NeedsWorkGetSliceFileSizeReportResult, *SdkError) {
	var res NeedsWorkGetSliceFileSizeReportResult
	_, err := sfClient.MakeSFCall(ctx, "GetSliceFileSizeReport", 1, nil, &res)
	return &res, err
}

// GetSnmpInfo enables you to retrieve the current simple network management protocol (SNMP) configuration information.
// Note: GetSnmpInfo is available for Element OS 8 and prior releases. It is deprecated for versions later than Element OS 8.
// NetApp recommends that you migrate to the GetSnmpState and SetSnmpACL methods. See details in the Element API Reference Guide
// for their descriptions and usage.
func (sfClient *SFClient) GetSnmpInfo(ctx context.Context) (*GetSnmpInfoResult, *SdkError) {
	var res GetSnmpInfoResult
	_, err := sfClient.MakeSFCall(ctx, "GetSnmpInfo", 1, nil, &res)
	return &res, err
}

// SetSnmpInfo enables you to configure SNMP version 2 and version 3 on cluster nodes. The values you set with this interface apply to
// all nodes in the cluster, and the values that are passed replace, in whole, all values set in any previous call to SetSnmpInfo.
// Note: SetSnmpInfo is deprecated. Use the EnableSnmp and SetSnmpACL methods instead.
func (sfClient *SFClient) SetSnmpInfo(ctx context.Context, req *SetSnmpInfoRequest) (*SetSnmpInfoResult, *SdkError) {
	var res SetSnmpInfoResult
	_, err := sfClient.MakeSFCall(ctx, "SetSnmpInfo", 1, req, &res)
	return &res, err
}

// GetRemoteLoggingHosts enables you to retrieve the current list of log servers.
func (sfClient *SFClient) GetRemoteLoggingHosts(ctx context.Context) (*GetRemoteLoggingHostsResult, *SdkError) {
	var res GetRemoteLoggingHostsResult
	_, err := sfClient.MakeSFCall(ctx, "GetRemoteLoggingHosts", 1, nil, &res)
	return &res, err
}

// ListProtectionDomainLevels returns the Tolerance and Resiliency of the cluster from the perspective
// of each of the supported ProtectionDomainTypes.
func (sfClient *SFClient) ListProtectionDomainLevels(ctx context.Context) (*ListProtectionDomainLevelsResult, *SdkError) {
	var res ListProtectionDomainLevelsResult
	_, err := sfClient.MakeSFCall(ctx, "ListProtectionDomainLevels", 1, nil, &res)
	return &res, err
}

//
func (sfClient *SFClient) ListAptSourceLines(ctx context.Context) (*NeedsWorkListAptSourceLinesResult, *SdkError) {
	var res NeedsWorkListAptSourceLinesResult
	_, err := sfClient.MakeSFCall(ctx, "ListAptSourceLines", 1, nil, &res)
	return &res, err
}

//
func (sfClient *SFClient) PairVolume(ctx context.Context) (*NeedsWorkPairVolumeResult, *SdkError) {
	var res NeedsWorkPairVolumeResult
	_, err := sfClient.MakeSFCall(ctx, "PairVolume", 1, nil, &res)
	return &res, err
}

// The ListNodeFibreChannelPortInfo API method enables you to retrieve information about the Fibre Channel ports on a node. The API method is intended for use on individual nodes; userid and password authentication is required for access to individual Fibre Channel nodes.
func (sfClient *SFClient) ListNodeFibreChannelPortInfo(ctx context.Context) (*ListNodeFibreChannelPortInfoResult, *SdkError) {
	var res ListNodeFibreChannelPortInfoResult
	_, err := sfClient.MakeSFCall(ctx, "ListNodeFibreChannelPortInfo", 1, nil, &res)
	return &res, err
}

// ModifyVirtualVolumeMetadata is used to selectively modify the VVol metadata.
func (sfClient *SFClient) ModifyVirtualVolumeMetadata(ctx context.Context, req *ModifyVirtualVolumeMetadataRequest) (*VirtualVolumeNullResult, *SdkError) {
	var res VirtualVolumeNullResult
	_, err := sfClient.MakeSFCall(ctx, "ModifyVirtualVolumeMetadata", 1, req, &res)
	return &res, err
}

// The CreateCluster method enables you to initialize the node in a cluster that has ownership of the "mvip" and "svip" addresses. Each new cluster is initialized using the management IP (MIP) of the first node in the cluster. This method also automatically adds all the nodes being configured into the cluster. You only need to use this method once each time a new cluster is initialized.
// Note: You need to log in to the node that is used as the master node for the cluster. After you log in, run the GetBootstrapConfig method on the node to get the IP addresses for the rest of the nodes that you want to include in the
// cluster. Then, run the CreateCluster method.
func (sfClient *SFClient) CreateCluster(ctx context.Context, req *CreateClusterRequest) (*CreateClusterResult, *SdkError) {
	var res CreateClusterResult
	_, err := sfClient.MakeSFCall(ctx, "CreateCluster", 1, req, &res)
	return &res, err
}

// The GetClusterConfig API method enables you to return information about the cluster configuration this node uses to communicate with the cluster that it is a part of.
func (sfClient *SFClient) GetClusterConfig(ctx context.Context, req *GetClusterConfigRequest) (*GetClusterConfigResult, *SdkError) {
	var res GetClusterConfigResult
	_, err := sfClient.MakeSFCall(ctx, "GetClusterConfig", 1, req, &res)
	return &res, err
}

// You can use the DeleteQoSPolicy method to delete a QoS policy from the system.
// The QoS settings for all volumes created of modified with this policy are unaffected.
func (sfClient *SFClient) DeleteQoSPolicy(ctx context.Context, req *DeleteQoSPolicyRequest) (*DeleteQoSPolicyResult, *SdkError) {
	var res DeleteQoSPolicyResult
	_, err := sfClient.MakeSFCall(ctx, "DeleteQoSPolicy", 1, req, &res)
	return &res, err
}

// The GetVolumeAccessGroupLunAssignments
// method enables you to retrieve details on LUN mappings
// of a specified volume access group.
func (sfClient *SFClient) GetVolumeAccessGroupLunAssignments(ctx context.Context, req *GetVolumeAccessGroupLunAssignmentsRequest) (*GetVolumeAccessGroupLunAssignmentsResult, *SdkError) {
	var res GetVolumeAccessGroupLunAssignmentsResult
	_, err := sfClient.MakeSFCall(ctx, "GetVolumeAccessGroupLunAssignments", 1, req, &res)
	return &res, err
}

//
func (sfClient *SFClient) DisableBmcColdReset(ctx context.Context) (*NeedsWorkDisableBmcColdResetResult, *SdkError) {
	var res NeedsWorkDisableBmcColdResetResult
	_, err := sfClient.MakeSFCall(ctx, "DisableBmcColdReset", 1, nil, &res)
	return &res, err
}

// CreateGroupSnapshot enables you to create a point-in-time copy of a group of volumes. You can use this snapshot later as a backup or rollback to ensure the data on the group of volumes is consistent for the point in time that you created the snapshot.
// Note: Creating a group snapshot is allowed if cluster fullness is at stage 2 or 3. Snapshots are not created when cluster fullness is at stage 4 or 5.
func (sfClient *SFClient) CreateGroupSnapshot(ctx context.Context, req *CreateGroupSnapshotRequest) (*CreateGroupSnapshotResult, *SdkError) {
	var res CreateGroupSnapshotResult
	_, err := sfClient.MakeSFCall(ctx, "CreateGroupSnapshot", 1, req, &res)
	return &res, err
}

// SetRemoteLoggingHosts enables you to configure remote logging from the nodes in the storage cluster to a centralized log server or servers. Remote logging is performed over TCP using the default port 514. This API does not add to the existing logging hosts. Rather, it replaces what currently exists with new values specified by this API method. You can use GetRemoteLoggingHosts to determine what the current logging hosts are, and then use SetRemoteLoggingHosts to set the desired list of current and new logging hosts.
func (sfClient *SFClient) SetRemoteLoggingHosts(ctx context.Context, req *SetRemoteLoggingHostsRequest) (*SetRemoteLoggingHostsResult, *SdkError) {
	var res SetRemoteLoggingHostsResult
	_, err := sfClient.MakeSFCall(ctx, "SetRemoteLoggingHosts", 1, req, &res)
	return &res, err
}

// You can use the AddBlockToBS method to store a block and associated blockID to the block services.
func (sfClient *SFClient) AddBlockToBS(ctx context.Context, req *AddBlockToBSRequest) (*AddBlockToBSResult, *SdkError) {
	var res AddBlockToBSResult
	_, err := sfClient.MakeSFCall(ctx, "AddBlockToBS", 1, req, &res)
	return &res, err
}

// CreateInitiators enables you to create multiple new initiator IQNs or World Wide Port Names (WWPNs) and optionally assign them
// aliases and attributes. When you use CreateInitiators to create new initiators, you can also add them to volume access groups.
// If CreateInitiators fails to create one of the initiators provided in the parameter, the method returns an error and does not create
// any initiators (no partial completion is possible).
func (sfClient *SFClient) CreateInitiators(ctx context.Context, req *CreateInitiatorsRequest) (*CreateInitiatorsResult, *SdkError) {
	var res CreateInitiatorsResult
	_, err := sfClient.MakeSFCall(ctx, "CreateInitiators", 1, req, &res)
	return &res, err
}

//
func (sfClient *SFClient) ReconfigureEnsemble(ctx context.Context) (*NeedsWorkReconfigureEnsembleResult, *SdkError) {
	var res NeedsWorkReconfigureEnsembleResult
	_, err := sfClient.MakeSFCall(ctx, "ReconfigureEnsemble", 1, nil, &res)
	return &res, err
}

// The CheckVolumeStandbysAssigned method enables a maintenance application to determine when standbys have been
// assigned to all volumes managed by a node. This should be sent after standby assignments have been enabled
// but before the node is taken offline, such as for an upgrade.
func (sfClient *SFClient) CheckVolumeStandbysAssigned(ctx context.Context, req *CheckVolumeStandbysAssignedRequest) (*CheckVolumeStandbysAssignedResult, *SdkError) {
	var res CheckVolumeStandbysAssignedResult
	_, err := sfClient.MakeSFCall(ctx, "CheckVolumeStandbysAssigned", 1, req, &res)
	return &res, err
}

// List all auth sessions associated with the specified ClusterAdminID.
// If the specified ClusterAdminID maps to a group of users, all auth sessions for all members of that group will be listed.
func (sfClient *SFClient) ListAuthSessionsByClusterAdmin(ctx context.Context, req *ListAuthSessionsByClusterAdminRequest) (*ListAuthSessionsResult, *SdkError) {
	var res ListAuthSessionsResult
	_, err := sfClient.MakeSFCall(ctx, "ListAuthSessionsByClusterAdmin", 1, req, &res)
	return &res, err
}

// This allows the auth container to store configuration data in the cluster
// database.  The configuration data itself is expected to be in text form.
// Subsequent calls to this method will replace data stored by earlier calls.
//
// In order to guarantee that the auth container has the current revision of the configuration,
// a version is required to be included which should correspond to the current configuration version
// in the database.  On success, the version in the database will be the version provided plus one.
// GetAuthConfiguration may be used to retrieve the current configuration and its version.
func (sfClient *SFClient) SetAuthConfiguration(ctx context.Context, req *SetAuthConfigurationRequest) (*SetAuthConfigurationResult, *SdkError) {
	var res SetAuthConfigurationResult
	_, err := sfClient.MakeSFCall(ctx, "SetAuthConfiguration", 1, req, &res)
	return &res, err
}

// You can use the GetSupportedTlsCiphers method to get a list of the supported TLS ciphers.
func (sfClient *SFClient) GetSupportedTlsCiphers(ctx context.Context) (*GetSupportedTlsCiphersResult, *SdkError) {
	var res GetSupportedTlsCiphersResult
	_, err := sfClient.MakeSFCall(ctx, "GetSupportedTlsCiphers", 1, nil, &res)
	return &res, err
}

// GetDriveConfig enables you to display drive information for expected slice and block drive counts as well as the number of slices
// and block drives that are currently connected to the node.
// Note: This method is available only through the per-node API endpoint 5.0 or later.
func (sfClient *SFClient) GetDriveConfig(ctx context.Context, req *GetDriveConfigRequest) (*GetDriveConfigResult, *SdkError) {
	var res GetDriveConfigResult
	_, err := sfClient.MakeSFCall(ctx, "GetDriveConfig", 1, req, &res)
	return &res, err
}

//
func (sfClient *SFClient) SetUpgradeNodeId(ctx context.Context) (*NeedsWorkSetUpgradeNodeIdResult, *SdkError) {
	var res NeedsWorkSetUpgradeNodeIdResult
	_, err := sfClient.MakeSFCall(ctx, "SetUpgradeNodeId", 1, nil, &res)
	return &res, err
}

// ListVirtualVolumeBindings returns a list of all virtual volumes in the cluster that are bound to protocol endpoints.
func (sfClient *SFClient) ListVirtualVolumeBindings(ctx context.Context, req *ListVirtualVolumeBindingsRequest) (*ListVirtualVolumeBindingsResult, *SdkError) {
	var res ListVirtualVolumeBindingsResult
	_, err := sfClient.MakeSFCall(ctx, "ListVirtualVolumeBindings", 1, req, &res)
	return &res, err
}

// You can use the SetSSLCertificate method to set a user SSL certificate and a private key for the cluster.
func (sfClient *SFClient) SetSSLCertificate(ctx context.Context, req *SetSSLCertificateRequest) (*SetSSLCertificateResult, *SdkError) {
	var res SetSSLCertificateResult
	_, err := sfClient.MakeSFCall(ctx, "SetSSLCertificate", 1, req, &res)
	return &res, err
}

// Create a potential trust relationship for authentication using a third party Identity Provider (IdP) for the cluster.
// A SAML Service Provider certificate is required for IdP communication, which will be generated as necessary.
func (sfClient *SFClient) CreateIdpConfiguration(ctx context.Context, req *CreateIdpConfigurationRequest) (*CreateIdpConfigurationResult, *SdkError) {
	var res CreateIdpConfigurationResult
	_, err := sfClient.MakeSFCall(ctx, "CreateIdpConfiguration", 1, req, &res)
	return &res, err
}

// You can use the CompleteClusterPairing method with the encoded key received from the  StartClusterPairing method to complete the cluster pairing process. The CompleteClusterPairing method is the second step in the cluster pairing process.
func (sfClient *SFClient) CompleteClusterPairing(ctx context.Context, req *CompleteClusterPairingRequest) (*CompleteClusterPairingResult, *SdkError) {
	var res CompleteClusterPairingResult
	_, err := sfClient.MakeSFCall(ctx, "CompleteClusterPairing", 1, req, &res)
	return &res, err
}

// You can use the CreateQoSPolicy method to create a QoSPolicy object that you can later apply to a volume upon creation or modification. A QoS policy has a unique ID, a name, and QoS settings.
func (sfClient *SFClient) CreateQoSPolicy(ctx context.Context, req *CreateQoSPolicyRequest) (*CreateQoSPolicyResult, *SdkError) {
	var res CreateQoSPolicyResult
	_, err := sfClient.MakeSFCall(ctx, "CreateQoSPolicy", 1, req, &res)
	return &res, err
}

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
func (sfClient *SFClient) ModifyVolumes(ctx context.Context, req *ModifyVolumesRequest) (*ModifyVolumesResult, *SdkError) {
	var res ModifyVolumesResult
	_, err := sfClient.MakeSFCall(ctx, "ModifyVolumes", 1, req, &res)
	return &res, err
}

//
func (sfClient *SFClient) DeleteSnapMirrorObjectAttributes(ctx context.Context) (*NeedsWorkDeleteSnapMirrorObjectAttributesResult, *SdkError) {
	var res NeedsWorkDeleteSnapMirrorObjectAttributesResult
	_, err := sfClient.MakeSFCall(ctx, "DeleteSnapMirrorObjectAttributes", 1, nil, &res)
	return &res, err
}

//
func (sfClient *SFClient) SetClusterSettings(ctx context.Context) (*NeedsWorkSetClusterSettingsResult, *SdkError) {
	var res NeedsWorkSetClusterSettingsResult
	_, err := sfClient.MakeSFCall(ctx, "SetClusterSettings", 1, nil, &res)
	return &res, err
}

//
func (sfClient *SFClient) DisconnectISCSISessions(ctx context.Context) (*NeedsWorkDisconnectISCSISessionsResult, *SdkError) {
	var res NeedsWorkDisconnectISCSISessionsResult
	_, err := sfClient.MakeSFCall(ctx, "DisconnectISCSISessions", 1, nil, &res)
	return &res, err
}

// ListVolumeStats returns high-level activity measurements for a single volume, list of volumes, or all volumes (if you omit the volumeIDs parameter). Measurement values are cumulative from the creation of the volume.
func (sfClient *SFClient) ListVolumeStats(ctx context.Context, req *ListVolumeStatsRequest) (*ListVolumeStatsResult, *SdkError) {
	var res ListVolumeStatsResult
	_, err := sfClient.MakeSFCall(ctx, "ListVolumeStats", 1, req, &res)
	return &res, err
}

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
func (sfClient *SFClient) RemoveDrives(ctx context.Context, req *RemoveDrivesRequest) (*AsyncHandleResult, *SdkError) {
	var res AsyncHandleResult
	_, err := sfClient.MakeSFCall(ctx, "RemoveDrives", 1, req, &res)
	return &res, err
}

// GetHardwareConfig enables you to display the hardware configuration information for a node.
// Note: This method is available only through the per-node API endpoint 5.0 or later.
func (sfClient *SFClient) GetHardwareConfig(ctx context.Context, req *GetHardwareConfigRequest) (*GetHardwareConfigResult, *SdkError) {
	var res GetHardwareConfigResult
	_, err := sfClient.MakeSFCall(ctx, "GetHardwareConfig", 1, req, &res)
	return &res, err
}

// The SolidFire Element OS web UI uses GetSnapMirrorClusterIdentity to get identity information about the ONTAP cluster.
func (sfClient *SFClient) GetSnapMirrorClusterIdentity(ctx context.Context, req *GetSnapMirrorClusterIdentityRequest) (*GetSnapMirrorClusterIdentityResult, *SdkError) {
	var res GetSnapMirrorClusterIdentityResult
	_, err := sfClient.MakeSFCall(ctx, "GetSnapMirrorClusterIdentity", 1, req, &res)
	return &res, err
}

// You can use the GetLoginBanner method to get the currently active Terms of Use banner that users see when they log on to the web interface.
func (sfClient *SFClient) GetLoginBanner(ctx context.Context) (*GetLoginBannerResult, *SdkError) {
	var res GetLoginBannerResult
	_, err := sfClient.MakeSFCall(ctx, "GetLoginBanner", 1, nil, &res)
	return &res, err
}

// This method dumps information about the chassis, local interfaces, and
// neighbors from the LLDP daemon
func (sfClient *SFClient) GetLldpInfo(ctx context.Context) (*GetLldpInfoResult, *SdkError) {
	var res GetLldpInfoResult
	_, err := sfClient.MakeSFCall(ctx, "GetLldpInfo", 1, nil, &res)
	return &res, err
}

// Lists all active auth sessions.
// This is only callable by a user with Administrative access rights.
func (sfClient *SFClient) ListActiveAuthSessions(ctx context.Context) (*ListAuthSessionsResult, *SdkError) {
	var res ListAuthSessionsResult
	_, err := sfClient.MakeSFCall(ctx, "ListActiveAuthSessions", 1, nil, &res)
	return &res, err
}

// Delete an existing configuration with a third party Identity Provider (IdP) for the cluster.
// Deleting the last IdP Configuration will remove the SAML Service Provider certificate from the cluster.
func (sfClient *SFClient) DeleteIdpConfiguration(ctx context.Context, req *DeleteIdpConfigurationRequest) (*DeleteIdpConfigurationResult, *SdkError) {
	var res DeleteIdpConfigurationResult
	_, err := sfClient.MakeSFCall(ctx, "DeleteIdpConfiguration", 1, req, &res)
	return &res, err
}

//
func (sfClient *SFClient) FinishUpgrade(ctx context.Context) (*NeedsWorkFinishUpgradeResult, *SdkError) {
	var res NeedsWorkFinishUpgradeResult
	_, err := sfClient.MakeSFCall(ctx, "FinishUpgrade", 1, nil, &res)
	return &res, err
}

// ListAccounts returns the entire list of accounts, with optional paging support.
func (sfClient *SFClient) ListAccounts(ctx context.Context, req *ListAccountsRequest) (*ListAccountsResult, *SdkError) {
	var res ListAccountsResult
	_, err := sfClient.MakeSFCall(ctx, "ListAccounts", 1, req, &res)
	return &res, err
}

//
func (sfClient *SFClient) DeleteDatabaseEntry(ctx context.Context) (*NeedsWorkDeleteDatabaseEntryResult, *SdkError) {
	var res NeedsWorkDeleteDatabaseEntryResult
	_, err := sfClient.MakeSFCall(ctx, "DeleteDatabaseEntry", 1, nil, &res)
	return &res, err
}

// GetBackupTarget enables you to return information about a specific backup target that you have created.
func (sfClient *SFClient) GetBackupTarget(ctx context.Context, req *GetBackupTargetRequest) (*GetBackupTargetResult, *SdkError) {
	var res GetBackupTargetResult
	_, err := sfClient.MakeSFCall(ctx, "GetBackupTarget", 1, req, &res)
	return &res, err
}

// GetLimits enables you to retrieve the limit values set by the API. These values might change between releases of Element OS, but do not change without an update to the system. Knowing the limit values set by the API can be useful when writing API scripts for user-facing tools.
// Note: The GetLimits method returns the limits for the current software version regardless of the API endpoint version used to pass the method.
func (sfClient *SFClient) GetLimits(ctx context.Context) (*GetLimitsResult, *SdkError) {
	var res GetLimitsResult
	_, err := sfClient.MakeSFCall(ctx, "GetLimits", 1, nil, &res)
	return &res, err
}

//
func (sfClient *SFClient) DeleteSnapshotSnapMirrorObjectAttributes(ctx context.Context) (*NeedsWorkDeleteSnapshotSnapMirrorObjectAttributesResult, *SdkError) {
	var res NeedsWorkDeleteSnapshotSnapMirrorObjectAttributesResult
	_, err := sfClient.MakeSFCall(ctx, "DeleteSnapshotSnapMirrorObjectAttributes", 1, nil, &res)
	return &res, err
}

//
func (sfClient *SFClient) GetClusterSettings(ctx context.Context) (*NeedsWorkGetClusterSettingsResult, *SdkError) {
	var res NeedsWorkGetClusterSettingsResult
	_, err := sfClient.MakeSFCall(ctx, "GetClusterSettings", 1, nil, &res)
	return &res, err
}

// The Shutdown API method enables you to restart or shutdown a node that has not yet been added to a cluster. To use this method,
// log in to the MIP for the pending node, and enter the "shutdown" method with either the "restart" or "halt" options.
func (sfClient *SFClient) Shutdown(ctx context.Context, req *ShutdownRequest) (*ShutdownResult, *SdkError) {
	var res ShutdownResult
	_, err := sfClient.MakeSFCall(ctx, "Shutdown", 1, req, &res)
	return &res, err
}

// Demotes a node from being cluster master. The cluster will promote a new node.
// The original cluster master will return success if it initiates the change. Callers should confirm that the
// change occurs by calling the API command GetClusterMasterNodeID after a delay of a few seconds.
func (sfClient *SFClient) DemoteClusterMaster(ctx context.Context, req *DemoteClusterMasterRequest) (*DemoteClusterMasterResult, *SdkError) {
	var res DemoteClusterMasterResult
	_, err := sfClient.MakeSFCall(ctx, "DemoteClusterMaster", 1, req, &res)
	return &res, err
}

// You can use GetClusterFullThreshold to view the stages set for cluster fullness levels. This method returns all fullness metrics for the
// cluster.
// Note: When a cluster reaches the Error stage of block cluster fullness, the maximum IOPS on all volumes are reduced linearly to the volume's minimum IOPS as the cluster approaches the Critical stage. This helps prevent the cluster from
// reaching the Critical stage of block cluster fullness.
func (sfClient *SFClient) GetClusterFullThreshold(ctx context.Context) (*GetClusterFullThresholdResult, *SdkError) {
	var res GetClusterFullThresholdResult
	_, err := sfClient.MakeSFCall(ctx, "GetClusterFullThreshold", 1, nil, &res)
	return &res, err
}

// You can use GetSnmpState to return the current state of the SNMP feature.
func (sfClient *SFClient) GetSnmpState(ctx context.Context) (*GetSnmpStateResult, *SdkError) {
	var res GetSnmpStateResult
	_, err := sfClient.MakeSFCall(ctx, "GetSnmpState", 1, nil, &res)
	return &res, err
}

// The TestLdapAuthentication method enables you to validate the currently enabled LDAP authentication settings. If the configuration is
// correct, the API call returns the group membership of the tested user.
func (sfClient *SFClient) TestLdapAuthentication(ctx context.Context, req *TestLdapAuthenticationRequest) (*TestLdapAuthenticationResult, *SdkError) {
	var res TestLdapAuthenticationResult
	_, err := sfClient.MakeSFCall(ctx, "TestLdapAuthentication", 1, req, &res)
	return &res, err
}

// The SolidFire Element OS web UI uses the ListSnapMirrorAggregates method to list all SnapMirror aggregates that are available on the remote ONTAP system. An aggregate describes a set of physical storage resources.
func (sfClient *SFClient) ListSnapMirrorAggregates(ctx context.Context, req *ListSnapMirrorAggregatesRequest) (*ListSnapMirrorAggregatesResult, *SdkError) {
	var res ListSnapMirrorAggregatesResult
	_, err := sfClient.MakeSFCall(ctx, "ListSnapMirrorAggregates", 1, req, &res)
	return &res, err
}

// Creates a KMIP (Key Management Interoperability Protocol) Key Server with the specified attributes. The server
// will not be contacted as part of this operation so it need not exist or be configured prior.  See TestKeyServerKmip.
// For clustered Key Server configurations, the hostnames or IP Addresses, of all server nodes, must be provided in
// the kmipKeyServerHostnames parameter.
func (sfClient *SFClient) CreateKeyServerKmip(ctx context.Context, req *CreateKeyServerKmipRequest) (*CreateKeyServerKmipResult, *SdkError) {
	var res CreateKeyServerKmipResult
	_, err := sfClient.MakeSFCall(ctx, "CreateKeyServerKmip", 1, req, &res)
	return &res, err
}

// Modifies a KMIP (Key Management Interoperability Protocol) Key Server to the specified attributes. The only required
// parameter is the keyServerID. A request which contains only the keyServerID will be a no-op and no error will be
// returned. Any other parameters which are specified will replace the existing values on the KMIP Key Server with
// the specified keyServerID. Because this server might be part of an active provider this will result in contacting
// the server to verify it's functional. Multiple hostnames or IP addresses must only be provided to the kmipKeyServerHostnames
// parameter if the key servers are in a clustered configuration.
func (sfClient *SFClient) ModifyKeyServerKmip(ctx context.Context, req *ModifyKeyServerKmipRequest) (*ModifyKeyServerKmipResult, *SdkError) {
	var res ModifyKeyServerKmipResult
	_, err := sfClient.MakeSFCall(ctx, "ModifyKeyServerKmip", 1, req, &res)
	return &res, err
}

// Enable support for authentication using a third party Identity Provider (IdP) for the cluster.
// Once IdP authentication is enabled, cluster and Ldap admins will no longer be able to access the cluster via supported UIs and any active authenticated sessions will be invalidated/logged out.
// Only third party IdP authenticated users will be able to access the cluster via the supported UIs.
func (sfClient *SFClient) EnableIdpAuthentication(ctx context.Context, req *EnableIdpAuthenticationRequest) (*EnableIdpAuthenticationResult, *SdkError) {
	var res EnableIdpAuthenticationResult
	_, err := sfClient.MakeSFCall(ctx, "EnableIdpAuthentication", 1, req, &res)
	return &res, err
}

//
func (sfClient *SFClient) GetRsyslogInfo(ctx context.Context) (*NeedsWorkGetRsyslogInfoResult, *SdkError) {
	var res NeedsWorkGetRsyslogInfoResult
	_, err := sfClient.MakeSFCall(ctx, "GetRsyslogInfo", 1, nil, &res)
	return &res, err
}

// CreateVirtualVolume is used to create a new (empty) Virtual Volume on the cluster.
// When the volume is created successfully it is available for connection via PE.
func (sfClient *SFClient) CreateVirtualVolume(ctx context.Context, req *CreateVirtualVolumeRequest) (*VirtualVolumeSyncResult, *SdkError) {
	var res VirtualVolumeSyncResult
	_, err := sfClient.MakeSFCall(ctx, "CreateVirtualVolume", 1, req, &res)
	return &res, err
}

// AddVolumesToVolumeAccessGroup enables you to add
// volumes to a specified volume access group.
func (sfClient *SFClient) AddVolumesToVolumeAccessGroup(ctx context.Context, req *AddVolumesToVolumeAccessGroupRequest) (*ModifyVolumeAccessGroupResult, *SdkError) {
	var res ModifyVolumeAccessGroupResult
	_, err := sfClient.MakeSFCall(ctx, "AddVolumesToVolumeAccessGroup", 1, req, &res)
	return &res, err
}

// Adds a cluster administrator user authenticated by a third party Identity Provider (IdP).
// IdP cluster admin accounts are configured based on SAML attribute-value information provided
// within the IdP's SAML assertion associated with the user.  If a user successfully
// authenticates with the IdP and has SAML attribute statements within the SAML assertion
// matching multiple IdP cluster admin accounts, the user will have the combined access level
// of those matching IdP cluster admin accounts.
func (sfClient *SFClient) AddIdpClusterAdmin(ctx context.Context, req *AddIdpClusterAdminRequest) (*AddClusterAdminResult, *SdkError) {
	var res AddClusterAdminResult
	_, err := sfClient.MakeSFCall(ctx, "AddIdpClusterAdmin", 1, req, &res)
	return &res, err
}

//
func (sfClient *SFClient) StartGC(ctx context.Context) (*NeedsWorkStartGCResult, *SdkError) {
	var res NeedsWorkStartGCResult
	_, err := sfClient.MakeSFCall(ctx, "StartGC", 1, nil, &res)
	return &res, err
}

// You can use ModifyVirtualNetwork to change the attributes of an existing virtual network. This method enables you to add or remove
// address blocks, change the netmask, or modify the name or description of the virtual network. You can also use it to enable or
// disable namespaces, as well as add or remove a gateway if namespaces are enabled on the virtual network.
// Note: This method requires either the VirtualNetworkID or the VirtualNetworkTag as a parameter, but not both.
// Caution: Enabling or disabling the Routable Storage VLANs functionality for an existing virtual network by changing the
// "namespace" parameter disrupts any traffic handled by the virtual network. NetApp strongly recommends changing the
// "namespace" parameter only during a scheduled maintenance window.
func (sfClient *SFClient) ModifyVirtualNetwork(ctx context.Context, req *ModifyVirtualNetworkRequest) (*AddVirtualNetworkResult, *SdkError) {
	var res AddVirtualNetworkResult
	_, err := sfClient.MakeSFCall(ctx, "ModifyVirtualNetwork", 1, req, &res)
	return &res, err
}

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
func (sfClient *SFClient) AddDrives(ctx context.Context, req *AddDrivesRequest) (*AddDrivesResult, *SdkError) {
	var res AddDrivesResult
	_, err := sfClient.MakeSFCall(ctx, "AddDrives", 1, req, &res)
	return &res, err
}

// ModifySnapshot enables you to change the attributes currently assigned to a snapshot. You can use this method to enable snapshots created on
// the Read/Write (source) volume to be remotely replicated to a target SolidFire storage system.
func (sfClient *SFClient) ModifySnapshot(ctx context.Context, req *ModifySnapshotRequest) (*ModifySnapshotResult, *SdkError) {
	var res ModifySnapshotResult
	_, err := sfClient.MakeSFCall(ctx, "ModifySnapshot", 1, req, &res)
	return &res, err
}

//
func (sfClient *SFClient) SetClusterFullThresholds(ctx context.Context) (*NeedsWorkSetClusterFullThresholdsResult, *SdkError) {
	var res NeedsWorkSetClusterFullThresholdsResult
	_, err := sfClient.MakeSFCall(ctx, "SetClusterFullThresholds", 1, nil, &res)
	return &res, err
}

// You can use ModifySnapMirrorRelationship to change the intervals at which a scheduled snapshot occurs. You can also delete or pause a schedule by using this method.
func (sfClient *SFClient) ModifySnapMirrorRelationship(ctx context.Context, req *ModifySnapMirrorRelationshipRequest) (*ModifySnapMirrorRelationshipResult, *SdkError) {
	var res ModifySnapMirrorRelationshipResult
	_, err := sfClient.MakeSFCall(ctx, "ModifySnapMirrorRelationship", 1, req, &res)
	return &res, err
}

// You can use AddAccount to add a new account to the system. You can create new volumes under the new account. The CHAP settings you specify for the account apply to all volumes owned by the account.
func (sfClient *SFClient) AddAccount(ctx context.Context, req *AddAccountRequest) (*AddAccountResult, *SdkError) {
	var res AddAccountResult
	_, err := sfClient.MakeSFCall(ctx, "AddAccount", 1, req, &res)
	return &res, err
}

//
func (sfClient *SFClient) EnableBmcColdReset(ctx context.Context) (*NeedsWorkEnableBmcColdResetResult, *SdkError) {
	var res NeedsWorkEnableBmcColdResetResult
	_, err := sfClient.MakeSFCall(ctx, "EnableBmcColdReset", 1, nil, &res)
	return &res, err
}

// Remove (unassign) the specified KMIP (Key Management Interoperability Protocol) Key Server from the provider it was assigned to via AddKeyServerToProviderKmip (if any).  A KMIP Key Server can be unassigned from its provider unless it's the last one and that provider is active (providing keys which are currently in use).  If the specified KMIP Key Server is not assigned to a provider, this is a no-op and no error will be returned.
func (sfClient *SFClient) RemoveKeyServerFromProviderKmip(ctx context.Context, req *RemoveKeyServerFromProviderKmipRequest) (*RemoveKeyServerFromProviderKmipResult, *SdkError) {
	var res RemoveKeyServerFromProviderKmipResult
	_, err := sfClient.MakeSFCall(ctx, "RemoveKeyServerFromProviderKmip", 1, req, &res)
	return &res, err
}

// You can use DisableSnmp to disable SNMP on the cluster nodes.
func (sfClient *SFClient) DisableSnmp(ctx context.Context) (*DisableSnmpResult, *SdkError) {
	var res DisableSnmpResult
	_, err := sfClient.MakeSFCall(ctx, "DisableSnmp", 1, nil, &res)
	return &res, err
}

// GetSnmpACL enables you to return the current SNMP access permissions on the cluster nodes.
func (sfClient *SFClient) GetSnmpACL(ctx context.Context) (*GetSnmpACLResult, *SdkError) {
	var res GetSnmpACLResult
	_, err := sfClient.MakeSFCall(ctx, "GetSnmpACL", 1, nil, &res)
	return &res, err
}

// You can use GetSnmpTrapInfo to return current SNMP trap configuration information.
func (sfClient *SFClient) GetSnmpTrapInfo(ctx context.Context) (*GetSnmpTrapInfoResult, *SdkError) {
	var res GetSnmpTrapInfoResult
	_, err := sfClient.MakeSFCall(ctx, "GetSnmpTrapInfo", 1, nil, &res)
	return &res, err
}

// ListInitiators enables you to list initiator IQNs or World Wide Port Names (WWPNs).
func (sfClient *SFClient) ListInitiators(ctx context.Context, req *ListInitiatorsRequest) (*ListInitiatorsResult, *SdkError) {
	var res ListInitiatorsResult
	_, err := sfClient.MakeSFCall(ctx, "ListInitiators", 1, req, &res)
	return &res, err
}

//
func (sfClient *SFClient) SetRepositories(ctx context.Context) (*NeedsWorkSetRepositoriesResult, *SdkError) {
	var res NeedsWorkSetRepositoriesResult
	_, err := sfClient.MakeSFCall(ctx, "SetRepositories", 1, nil, &res)
	return &res, err
}

//
func (sfClient *SFClient) GetSliceInfo(ctx context.Context) (*NeedsWorkGetSliceInfoResult, *SdkError) {
	var res NeedsWorkGetSliceInfoResult
	_, err := sfClient.MakeSFCall(ctx, "GetSliceInfo", 1, nil, &res)
	return &res, err
}

// ListSnapshots enables you to return the attributes of each snapshot taken on the volume. Information about snapshots that reside on the target cluster is displayed on the source cluster when this method is called from the source cluster.
func (sfClient *SFClient) ListSnapshots(ctx context.Context, req *ListSnapshotsRequest) (*ListSnapshotsResult, *SdkError) {
	var res ListSnapshotsResult
	_, err := sfClient.MakeSFCall(ctx, "ListSnapshots", 1, req, &res)
	return &res, err
}

// You can use the ResurrectDeadVirtualVolume method to recover a virtual volume from a secondary replica
// that may not have all data acknowledged to the client. This should only be used in circumstances
// where the primary slice service is permanently inaccessible and you want to recover as much user
// data as possible.
func (sfClient *SFClient) ResurrectDeadVirtualVolume(ctx context.Context, req *ResurrectDeadVirtualVolumeRequest) (*ResurrectDeadVirtualVolumeResult, *SdkError) {
	var res ResurrectDeadVirtualVolumeResult
	_, err := sfClient.MakeSFCall(ctx, "ResurrectDeadVirtualVolume", 1, req, &res)
	return &res, err
}

// ListDrives enables you to retrieve the list of the drives that exist in the cluster's active nodes. This method returns drives that have
// been added as volume metadata or block drives as well as drives that have not been added and are available.
func (sfClient *SFClient) ListDrives(ctx context.Context) (*ListDrivesResult, *SdkError) {
	var res ListDrivesResult
	_, err := sfClient.MakeSFCall(ctx, "ListDrives", 1, nil, &res)
	return &res, err
}

// The SolidFire Element OS web UI uses the ModifySnapMirrorEndpoint method to change the name and management attributes for a SnapMirror endpoint.
func (sfClient *SFClient) ModifySnapMirrorEndpoint(ctx context.Context, req *ModifySnapMirrorEndpointRequest) (*ModifySnapMirrorEndpointResult, *SdkError) {
	var res ModifySnapMirrorEndpointResult
	_, err := sfClient.MakeSFCall(ctx, "ModifySnapMirrorEndpoint", 1, req, &res)
	return &res, err
}

// Refreshes an auth session.
// Intended to be used by the element-auth container.
func (sfClient *SFClient) UpdateAuthSession(ctx context.Context, req *UpdateAuthSessionRequest) (*UpdateAuthSessionResult, *SdkError) {
	var res UpdateAuthSessionResult
	_, err := sfClient.MakeSFCall(ctx, "UpdateAuthSession", 1, req, &res)
	return &res, err
}

// You can use SetLoginSessionInfo to set the period of time that a session's login authentication is valid. After the log in period elapses without activity on the system, the authentication expires. New login credentials are required for continued access to the cluster after the timeout period has elapsed.
func (sfClient *SFClient) SetLoginSessionInfo(ctx context.Context, req *SetLoginSessionInfoRequest) (*SetLoginSessionInfoResult, *SdkError) {
	var res SetLoginSessionInfoResult
	_, err := sfClient.MakeSFCall(ctx, "SetLoginSessionInfo", 1, req, &res)
	return &res, err
}

//
func (sfClient *SFClient) ListSliceBranchesByService(ctx context.Context) (*NeedsWorkListSliceBranchesByServiceResult, *SdkError) {
	var res NeedsWorkListSliceBranchesByServiceResult
	_, err := sfClient.MakeSFCall(ctx, "ListSliceBranchesByService", 1, nil, &res)
	return &res, err
}

// DeleteGroupSnapshot enables you to delete a group snapshot. You can use the saveMembers parameter to preserve all the snapshots that were made for the volumes in the group, but the group association is removed.
func (sfClient *SFClient) DeleteGroupSnapshot(ctx context.Context, req *DeleteGroupSnapshotRequest) (*DeleteGroupSnapshotResult, *SdkError) {
	var res DeleteGroupSnapshotResult
	_, err := sfClient.MakeSFCall(ctx, "DeleteGroupSnapshot", 1, req, &res)
	return &res, err
}

// The GetImmutableValues API returns immutable constants that affect cluster behavior.
func (sfClient *SFClient) GetImmutableValues(ctx context.Context) (*GetImmutableValuesResult, *SdkError) {
	var res GetImmutableValuesResult
	_, err := sfClient.MakeSFCall(ctx, "GetImmutableValues", 1, nil, &res)
	return &res, err
}

// Deletes all auth sessions for the given user.
// A caller not in AccessGroup ClusterAdmins / Administrator may only delete their own sessions.
// A caller with ClusterAdmins / Administrator privileges may delete sessions belonging to any user.
// To see the list of sessions that could be deleted, use ListAuthSessionsByUsername with the same parameters.
func (sfClient *SFClient) DeleteAuthSessionsByUsername(ctx context.Context, req *DeleteAuthSessionsByUsernameRequest) (*DeleteAuthSessionsResult, *SdkError) {
	var res DeleteAuthSessionsResult
	_, err := sfClient.MakeSFCall(ctx, "DeleteAuthSessionsByUsername", 1, req, &res)
	return &res, err
}

// CheckProposedNodeAdditions validates that adding a node (or nodes) to an existing cluster is likely to succeed.  Any problems with the proposed new cluster are returned as errors with a human-readable description and unique error code.
func (sfClient *SFClient) CheckProposedNodeAdditions(ctx context.Context, req *CheckProposedNodeAdditionsRequest) (*CheckProposedResult, *SdkError) {
	var res CheckProposedResult
	_, err := sfClient.MakeSFCall(ctx, "CheckProposedNodeAdditions", 1, req, &res)
	return &res, err
}

// CreateSnapshot enables you to create a point-in-time copy of a volume. You can create a snapshot from any volume or from an existing snapshot. If you do not provide a SnapshotID with this API method, a snapshot is created from the volume's active branch.
// If the volume from which the snapshot is created is being replicated to a remote cluster, the snapshot can also be replicated to the same target. Use the enableRemoteReplication parameter to enable snapshot replication.
// Note: Creating a snapshot is allowed if cluster fullness is at stage 2 or 3. Snapshots are not created when cluster fullness is at stage 4 or 5.
func (sfClient *SFClient) CreateSnapshot(ctx context.Context, req *CreateSnapshotRequest) (*CreateSnapshotResult, *SdkError) {
	var res CreateSnapshotResult
	_, err := sfClient.MakeSFCall(ctx, "CreateSnapshot", 1, req, &res)
	return &res, err
}

// PrepareVirtualSnapshot is used to set up VMware Virtual Volume snapshot.
func (sfClient *SFClient) PrepareVirtualSnapshot(ctx context.Context, req *PrepareVirtualSnapshotRequest) (*PrepareVirtualSnapshotResult, *SdkError) {
	var res PrepareVirtualSnapshotResult
	_, err := sfClient.MakeSFCall(ctx, "PrepareVirtualSnapshot", 1, req, &res)
	return &res, err
}

// The EnableLdapAuthentication method enables you to configure an LDAP directory connection to use for LDAP authentication to a cluster. Users that are members of the LDAP directory can then log in to the storage system using their LDAP credentials.
func (sfClient *SFClient) EnableLdapAuthentication(ctx context.Context, req *EnableLdapAuthenticationRequest) (*EnableLdapAuthenticationResult, *SdkError) {
	var res EnableLdapAuthenticationResult
	_, err := sfClient.MakeSFCall(ctx, "EnableLdapAuthentication", 1, req, &res)
	return &res, err
}

//
func (sfClient *SFClient) StopClusterBSCheck(ctx context.Context) (*NeedsWorkStopClusterBSCheckResult, *SdkError) {
	var res NeedsWorkStopClusterBSCheckResult
	_, err := sfClient.MakeSFCall(ctx, "StopClusterBSCheck", 1, nil, &res)
	return &res, err
}

// Delete the specified KMIP (Key Management Interoperability Protocol) Key Server.  A KMIP Key Server can be deleted unless it's the last one assigned to its provider, and that provider is active (providing keys which are currently in use).
func (sfClient *SFClient) DeleteKeyServerKmip(ctx context.Context, req *DeleteKeyServerKmipRequest) (*DeleteKeyServerKmipResult, *SdkError) {
	var res DeleteKeyServerKmipResult
	_, err := sfClient.MakeSFCall(ctx, "DeleteKeyServerKmip", 1, req, &res)
	return &res, err
}

// The TestConnectEnsemble API method enables you to verify connectivity with a specified database ensemble. By default, it uses the ensemble for the cluster that the node is associated with. Alternatively, you can provide a different ensemble to test connectivity with.
// Note: This method is available only through the per-node API endpoint 5.0 or later.
func (sfClient *SFClient) TestConnectEnsemble(ctx context.Context, req *TestConnectEnsembleRequest) (*TestConnectEnsembleResult, *SdkError) {
	var res TestConnectEnsembleResult
	_, err := sfClient.MakeSFCall(ctx, "TestConnectEnsemble", 1, req, &res)
	return &res, err
}

// Lists all the existing cluster interface preferences.
func (sfClient *SFClient) ListClusterInterfacePreferences(ctx context.Context) (*ListClusterInterfacePreferencesResult, *SdkError) {
	var res ListClusterInterfacePreferencesResult
	_, err := sfClient.MakeSFCall(ctx, "ListClusterInterfacePreferences", 1, nil, &res)
	return &res, err
}

//
func (sfClient *SFClient) GetDatabaseEntry(ctx context.Context) (*NeedsWorkGetDatabaseEntryResult, *SdkError) {
	var res NeedsWorkGetDatabaseEntryResult
	_, err := sfClient.MakeSFCall(ctx, "GetDatabaseEntry", 1, nil, &res)
	return &res, err
}

//
func (sfClient *SFClient) ForceRecycle(ctx context.Context) (*NeedsWorkForceRecycleResult, *SdkError) {
	var res NeedsWorkForceRecycleResult
	_, err := sfClient.MakeSFCall(ctx, "ForceRecycle", 1, nil, &res)
	return &res, err
}

//
func (sfClient *SFClient) ModifyClusterFault(ctx context.Context) (*NeedsWorkModifyClusterFaultResult, *SdkError) {
	var res NeedsWorkModifyClusterFaultResult
	_, err := sfClient.MakeSFCall(ctx, "ModifyClusterFault", 1, nil, &res)
	return &res, err
}

// Creates a KMIP (Key Management Interoperability Protocol) Key Provider with the specified name.  A Key Provider defines a mechanism and location to retrieve authentication keys.  A KMIP Key Provider represents a collection of one or more KMIP Key Servers.  A newly created KMIP Key Provider will not have any KMIP Key Servers assigned to it.  To create a KMIP Key Server see CreateKeyServerKmip and to assign it to a provider created via this method see AddKeyServerToProviderKmip.
func (sfClient *SFClient) CreateKeyProviderKmip(ctx context.Context, req *CreateKeyProviderKmipRequest) (*CreateKeyProviderKmipResult, *SdkError) {
	var res CreateKeyProviderKmipResult
	_, err := sfClient.MakeSFCall(ctx, "CreateKeyProviderKmip", 1, req, &res)
	return &res, err
}

// The RestartServices API method enables you to restart the services on a node.
// Caution: This method causes temporary node services interruption. Exercise caution when using this method.
// Note: This method is available only through the per-node API endpoint 5.0 or later.
func (sfClient *SFClient) RestartServices(ctx context.Context, req *RestartServicesRequest) (*interface{}, *SdkError) {
	var res interface{}
	_, err := sfClient.MakeSFCall(ctx, "RestartServices", 1, req, &res)
	return &res, err
}

// The SolidFire Element OS web UI uses the CreateSnapMirrorEndpoint method to create a relationship with a remote SnapMirror endpoint.
func (sfClient *SFClient) CreateSnapMirrorEndpoint(ctx context.Context, req *CreateSnapMirrorEndpointRequest) (*CreateSnapMirrorEndpointResult, *SdkError) {
	var res CreateSnapMirrorEndpointResult
	_, err := sfClient.MakeSFCall(ctx, "CreateSnapMirrorEndpoint", 1, req, &res)
	return &res, err
}

// AddLdapClusterAdmin enables you to add a new LDAP cluster administrator user. An LDAP cluster administrator can manage the
// cluster via the API and management tools. LDAP cluster admin accounts are completely separate and unrelated to standard tenant
// accounts.
// You can also use this method to add an LDAP group that has been defined in Active Directory. The access level that is given to the group is passed to the individual users in the LDAP group.
func (sfClient *SFClient) AddLdapClusterAdmin(ctx context.Context, req *AddLdapClusterAdminRequest) (*AddLdapClusterAdminResult, *SdkError) {
	var res AddLdapClusterAdminResult
	_, err := sfClient.MakeSFCall(ctx, "AddLdapClusterAdmin", 1, req, &res)
	return &res, err
}

//
func (sfClient *SFClient) APIWriteBlock(ctx context.Context) (*NeedsWorkAPIWriteBlockResult, *SdkError) {
	var res NeedsWorkAPIWriteBlockResult
	_, err := sfClient.MakeSFCall(ctx, "APIWriteBlock", 1, nil, &res)
	return &res, err
}

// ListSchedule enables you to retrieve information about all scheduled snapshots that have been created.
func (sfClient *SFClient) ListSchedules(ctx context.Context) (*ListSchedulesResult, *SdkError) {
	var res ListSchedulesResult
	_, err := sfClient.MakeSFCall(ctx, "ListSchedules", 1, nil, &res)
	return &res, err
}

// CopyVolume enables you to overwrite the data contents of an existing volume with the data contents of another volume (or
// snapshot). Attributes of the destination volume such as IQN, QoS settings, size, account, and volume access group membership are
// not changed. The destination volume must already exist and must be the same size as the source volume.
// NetApp strongly recommends that clients unmount the destination volume before the CopyVolume operation begins. If the
// destination volume is modified during the copy operation, the changes will be lost.
// This method is asynchronous and may take a variable amount of time to complete. You can use the GetAsyncResult method to
// determine when the process has finished, and ListSyncJobs to see the progress of the copy.
func (sfClient *SFClient) CopyVolume(ctx context.Context, req *CopyVolumeRequest) (*CopyVolumeResult, *SdkError) {
	var res CopyVolumeResult
	_, err := sfClient.MakeSFCall(ctx, "CopyVolume", 1, req, &res)
	return &res, err
}

// CreateBackupTarget enables you to create and store backup target information so that you do not need to re-enter it each time a backup is created.
func (sfClient *SFClient) CreateBackupTarget(ctx context.Context, req *CreateBackupTargetRequest) (*CreateBackupTargetResult, *SdkError) {
	var res CreateBackupTargetResult
	_, err := sfClient.MakeSFCall(ctx, "CreateBackupTarget", 1, req, &res)
	return &res, err
}

//
func (sfClient *SFClient) SetDaemonStatus(ctx context.Context) (*NeedsWorkSetDaemonStatusResult, *SdkError) {
	var res NeedsWorkSetDaemonStatusResult
	_, err := sfClient.MakeSFCall(ctx, "SetDaemonStatus", 1, nil, &res)
	return &res, err
}

// GetFeatureStatus enables you to retrieve the status of a cluster feature.
func (sfClient *SFClient) GetFeatureStatus(ctx context.Context, req *GetFeatureStatusRequest) (*GetFeatureStatusResult, *SdkError) {
	var res GetFeatureStatusResult
	_, err := sfClient.MakeSFCall(ctx, "GetFeatureStatus", 1, req, &res)
	return &res, err
}

//
func (sfClient *SFClient) ReadSliceLBA(ctx context.Context) (*NeedsWorkReadSliceLBAResult, *SdkError) {
	var res NeedsWorkReadSliceLBAResult
	_, err := sfClient.MakeSFCall(ctx, "ReadSliceLBA", 1, nil, &res)
	return &res, err
}

// The SetNetworkConfig API method enables you to set the network configuration for a node. To display the current network settings for a node, run the GetNetworkConfig API method.
// Note: This method is available only through the per-node API endpoint 5.0 or later.
// Changing the "bond-mode" on a node can cause a temporary loss of network connectivity. Exercise caution when using this method.
func (sfClient *SFClient) SetNetworkConfig(ctx context.Context, req *SetNetworkConfigRequest) (*SetNetworkConfigResult, *SdkError) {
	var res SetNetworkConfigResult
	_, err := sfClient.MakeSFCall(ctx, "SetNetworkConfig", 1, req, &res)
	return &res, err
}

// You can use the GetNodeSSLCertificate method to retrieve the SSL certificate that is currently active on the cluster.
// You can use this method on both management and storage nodes.
func (sfClient *SFClient) GetNodeSSLCertificate(ctx context.Context) (*GetNodeSSLCertificateResult, *SdkError) {
	var res GetNodeSSLCertificateResult
	_, err := sfClient.MakeSFCall(ctx, "GetNodeSSLCertificate", 1, nil, &res)
	return &res, err
}

// BindVirtualVolume binds a VVol with a Host.
func (sfClient *SFClient) BindVirtualVolumes(ctx context.Context, req *BindVirtualVolumesRequest) (*VirtualVolumeBindingListResult, *SdkError) {
	var res VirtualVolumeBindingListResult
	_, err := sfClient.MakeSFCall(ctx, "BindVirtualVolumes", 1, req, &res)
	return &res, err
}

// ListVirtualVolumeHosts returns a list of all virtual volume hosts known to the cluster. A virtual volume host is a VMware ESX host
// that has initiated a session with the VASA API provider.
func (sfClient *SFClient) ListVirtualVolumeHosts(ctx context.Context, req *ListVirtualVolumeHostsRequest) (*ListVirtualVolumeHostsResult, *SdkError) {
	var res ListVirtualVolumeHostsResult
	_, err := sfClient.MakeSFCall(ctx, "ListVirtualVolumeHosts", 1, req, &res)
	return &res, err
}

// You can use ModifyVolumeAccessGroup to update initiators and add or remove volumes from a volume access group. If a specified initiator or volume is a duplicate of what currently exists, the volume access group is left as-is. If you do not specify a value for volumes or initiators, the current list of initiators and volumes is not changed.
func (sfClient *SFClient) ModifyVolumeAccessGroup(ctx context.Context, req *ModifyVolumeAccessGroupRequest) (*ModifyVolumeAccessGroupResult, *SdkError) {
	var res ModifyVolumeAccessGroupResult
	_, err := sfClient.MakeSFCall(ctx, "ModifyVolumeAccessGroup", 1, req, &res)
	return &res, err
}

// DeleteAllSupportBundles enables you to delete all support bundles generated with the CreateSupportBundle API method.
func (sfClient *SFClient) DeleteAllSupportBundles(ctx context.Context) (*DeleteAllSupportBundlesResult, *SdkError) {
	var res DeleteAllSupportBundlesResult
	_, err := sfClient.MakeSFCall(ctx, "DeleteAllSupportBundles", 1, nil, &res)
	return &res, err
}

//
func (sfClient *SFClient) GetFibreChannelVolumeAccessInfo(ctx context.Context) (*NeedsWorkGetFibreChannelVolumeAccessInfoResult, *SdkError) {
	var res NeedsWorkGetFibreChannelVolumeAccessInfoResult
	_, err := sfClient.MakeSFCall(ctx, "GetFibreChannelVolumeAccessInfo", 1, nil, &res)
	return &res, err
}

// GetLldpConfig returns the current LLDP configuration for this node.
func (sfClient *SFClient) GetLldpConfig(ctx context.Context) (*GetLldpConfigResult, *SdkError) {
	var res GetLldpConfigResult
	_, err := sfClient.MakeSFCall(ctx, "GetLldpConfig", 1, nil, &res)
	return &res, err
}

// GetFipsReport enables you to retrieve FIPS compliance status on a per node basis.
func (sfClient *SFClient) GetFipsReport(ctx context.Context) (*GetFipsReportResult, *SdkError) {
	var res GetFipsReportResult
	_, err := sfClient.MakeSFCall(ctx, "GetFipsReport", 1, nil, &res)
	return &res, err
}

// Deletes an existing cluster interface preference.
func (sfClient *SFClient) DeleteClusterInterfacePreference(ctx context.Context, req *DeleteClusterInterfacePreferenceRequest) (*DeleteClusterInterfacePreferenceResult, *SdkError) {
	var res DeleteClusterInterfacePreferenceResult
	_, err := sfClient.MakeSFCall(ctx, "DeleteClusterInterfacePreference", 1, req, &res)
	return &res, err
}

// You can use the RemoveBlockFromBS method to remove a block from the block services.
func (sfClient *SFClient) RemoveBlockFromBS(ctx context.Context, req *RemoveBlockFromBSRequest) (*RemoveBlockFromBSResult, *SdkError) {
	var res RemoveBlockFromBSResult
	_, err := sfClient.MakeSFCall(ctx, "RemoveBlockFromBS", 1, req, &res)
	return &res, err
}

// The GetHardwareInfo API method enables you to return hardware information and status for a single node. This generally includes details about manufacturers, vendors, versions, drives, and other associated hardware identification information.
func (sfClient *SFClient) GetHardwareInfo(ctx context.Context) (*GetHardwareInfoResult, *SdkError) {
	var res GetHardwareInfoResult
	_, err := sfClient.MakeSFCall(ctx, "GetHardwareInfo", 1, nil, &res)
	return &res, err
}

// ListVolumeStatsByVirtualVolume enables you to list volume statistics for any volumes in the system that are associated with virtual volumes. Statistics are cumulative from the creation of the volume.
func (sfClient *SFClient) ListVolumeStatsByVirtualVolume(ctx context.Context, req *ListVolumeStatsByVirtualVolumeRequest) (*ListVolumeStatsByVirtualVolumeResult, *SdkError) {
	var res ListVolumeStatsByVirtualVolumeResult
	_, err := sfClient.MakeSFCall(ctx, "ListVolumeStatsByVirtualVolume", 1, req, &res)
	return &res, err
}

// The SolidFire Element OS web UI uses the ListSnapMirrorNodes method to get a list of nodes in a remote ONTAP cluster.
func (sfClient *SFClient) ListSnapMirrorNodes(ctx context.Context, req *ListSnapMirrorNodesRequest) (*ListSnapMirrorNodesResult, *SdkError) {
	var res ListSnapMirrorNodesResult
	_, err := sfClient.MakeSFCall(ctx, "ListSnapMirrorNodes", 1, req, &res)
	return &res, err
}

// This method returns a string containing element specific configuration data set by the auth container.
func (sfClient *SFClient) GetAuthConfiguration(ctx context.Context, req *GetAuthConfigurationRequest) (*GetAuthConfigurationResult, *SdkError) {
	var res GetAuthConfigurationResult
	_, err := sfClient.MakeSFCall(ctx, "GetAuthConfiguration", 1, req, &res)
	return &res, err
}

// GetCurrentClusterAdmin returns information about the calling ClusterAdmin.
// If the authMethod in the return value is Ldap or Idp, then other fields in the return value may contain data aggregated from multiple LdapAdmins or IdpAdmins, respectively.
func (sfClient *SFClient) GetCurrentClusterAdmin(ctx context.Context) (*GetCurrentClusterAdminResult, *SdkError) {
	var res GetCurrentClusterAdminResult
	_, err := sfClient.MakeSFCall(ctx, "GetCurrentClusterAdmin", 1, nil, &res)
	return &res, err
}

// ListDriveHardware returns all the drives connected to a node. Use this method on individual nodes to return drive hardware
// information or use this method on the cluster master node MVIP to see information for all the drives on all nodes.
// Note: The "securitySupported": true line of the method response does not imply that the drives are capable of
// encryption; only that the security status can be queried. If you have a node type with a model number ending in "-NE",
// commands to enable security features on these drives will fail. See the EnableEncryptionAtRest method for more information.
func (sfClient *SFClient) ListDriveHardware(ctx context.Context, req *ListDriveHardwareRequest) (*ListDriveHardwareResult, *SdkError) {
	var res ListDriveHardwareResult
	_, err := sfClient.MakeSFCall(ctx, "ListDriveHardware", 1, req, &res)
	return &res, err
}

//
func (sfClient *SFClient) BDTPipeline(ctx context.Context) (*NeedsWorkBDTPipelineResult, *SdkError) {
	var res NeedsWorkBDTPipelineResult
	_, err := sfClient.MakeSFCall(ctx, "BDTPipeline", 1, nil, &res)
	return &res, err
}

// EnableSnmp enables you to enable SNMP on cluster nodes. When you enable SNMP, the action applies to all nodes in the cluster, and
// the values that are passed replace, in whole, all values set in any previous call to EnableSnmp.
func (sfClient *SFClient) EnableSnmp(ctx context.Context, req *EnableSnmpRequest) (*EnableSnmpResult, *SdkError) {
	var res EnableSnmpResult
	_, err := sfClient.MakeSFCall(ctx, "EnableSnmp", 1, req, &res)
	return &res, err
}

// You can use ModifyClusterFullThreshold to change the level at which the system generates an event when the storage cluster approaches a certain capacity utilization. You can use the threshold settings to indicate the acceptable amount of utilized block storage or metadata storage before the system generates a warning. For example, if you want to be alerted when the system reaches 3% below the "Error" level block storage utilization, enter a value of "3" for the stage3BlockThresholdPercent parameter. If this level is reached, the system sends an alert to the Event Log in the Cluster Management Console.
func (sfClient *SFClient) ModifyClusterFullThreshold(ctx context.Context, req *ModifyClusterFullThresholdRequest) (*ModifyClusterFullThresholdResult, *SdkError) {
	var res ModifyClusterFullThresholdResult
	_, err := sfClient.MakeSFCall(ctx, "ModifyClusterFullThreshold", 1, req, &res)
	return &res, err
}

// You can use the GetSSLCertificate method to retrieve the SSL certificate that is currently active on the cluster.
func (sfClient *SFClient) GetSSLCertificate(ctx context.Context) (*GetSSLCertificateResult, *SdkError) {
	var res GetSSLCertificateResult
	_, err := sfClient.MakeSFCall(ctx, "GetSSLCertificate", 1, nil, &res)
	return &res, err
}

//
func (sfClient *SFClient) GetConnectivityReport(ctx context.Context) (*NeedsWorkGetConnectivityReportResult, *SdkError) {
	var res NeedsWorkGetConnectivityReportResult
	_, err := sfClient.MakeSFCall(ctx, "GetConnectivityReport", 1, nil, &res)
	return &res, err
}

//
func (sfClient *SFClient) ListClusterCapacityHistory(ctx context.Context) (*NeedsWorkListClusterCapacityHistoryResult, *SdkError) {
	var res NeedsWorkListClusterCapacityHistoryResult
	_, err := sfClient.MakeSFCall(ctx, "ListClusterCapacityHistory", 1, nil, &res)
	return &res, err
}

// ListPendingActiveNodes returns the list of nodes in the cluster that are currently in the PendingActive state, between the pending and active states. These are nodes that are currently being returned to the factory image.
func (sfClient *SFClient) ListPendingActiveNodes(ctx context.Context) (*ListPendingActiveNodesResult, *SdkError) {
	var res ListPendingActiveNodesResult
	_, err := sfClient.MakeSFCall(ctx, "ListPendingActiveNodes", 1, nil, &res)
	return &res, err
}

// You can use the CompleteVolumePairing method to complete the pairing of two volumes.
func (sfClient *SFClient) CompleteVolumePairing(ctx context.Context, req *CompleteVolumePairingRequest) (*CompleteVolumePairingResult, *SdkError) {
	var res CompleteVolumePairingResult
	_, err := sfClient.MakeSFCall(ctx, "CompleteVolumePairing", 1, req, &res)
	return &res, err
}

// ListActivePairedVolumes enables you to list all the active volumes paired with a volume. This method returns information about volumes with active and pending pairings.
func (sfClient *SFClient) ListActivePairedVolumes(ctx context.Context, req *ListActivePairedVolumesRequest) (*ListActivePairedVolumesResult, *SdkError) {
	var res ListActivePairedVolumesResult
	_, err := sfClient.MakeSFCall(ctx, "ListActivePairedVolumes", 1, req, &res)
	return &res, err
}

// The RemoveVolumeFromVolumeAccessGroup method enables you to remove volumes from a volume access group.
func (sfClient *SFClient) RemoveVolumesFromVolumeAccessGroup(ctx context.Context, req *RemoveVolumesFromVolumeAccessGroupRequest) (*ModifyVolumeAccessGroupResult, *SdkError) {
	var res ModifyVolumeAccessGroupResult
	_, err := sfClient.MakeSFCall(ctx, "RemoveVolumesFromVolumeAccessGroup", 1, req, &res)
	return &res, err
}

//
func (sfClient *SFClient) SetClusterStructure(ctx context.Context) (*NeedsWorkSetClusterStructureResult, *SdkError) {
	var res NeedsWorkSetClusterStructureResult
	_, err := sfClient.MakeSFCall(ctx, "SetClusterStructure", 1, nil, &res)
	return &res, err
}

// GetIpmiInfo enables you to display a detailed reporting of sensors (objects) for node fans, intake and exhaust temperatures, and power supplies that are monitored by the system.
func (sfClient *SFClient) GetIpmiInfo(ctx context.Context, req *GetIpmiInfoRequest) (*GetIpmiInfoResult, *SdkError) {
	var res GetIpmiInfoResult
	_, err := sfClient.MakeSFCall(ctx, "GetIpmiInfo", 1, req, &res)
	return &res, err
}

//
func (sfClient *SFClient) ProtocolVersionUpgradeAvailable(ctx context.Context) (*NeedsWorkProtocolVersionUpgradeAvailableResult, *SdkError) {
	var res NeedsWorkProtocolVersionUpgradeAvailableResult
	_, err := sfClient.MakeSFCall(ctx, "ProtocolVersionUpgradeAvailable", 1, nil, &res)
	return &res, err
}

// ListFibreChannelPortInfo enables you to retrieve information about the Fibre Channel ports on a node.  The API method is intended for use on individual nodes; userid and password authentication is required for access to individual Fibre Channel nodes.
func (sfClient *SFClient) ListFibreChannelPortInfo(ctx context.Context) (*ListFibreChannelPortInfoResult, *SdkError) {
	var res ListFibreChannelPortInfoResult
	_, err := sfClient.MakeSFCall(ctx, "ListFibreChannelPortInfo", 1, nil, &res)
	return &res, err
}

// The DisableLdapAuthentication method enables you to disable LDAP authentication and remove all LDAP configuration settings. This method does not remove any configured cluster admin accounts (user or group). However, those cluster admin accounts will no longer be able to log in.
func (sfClient *SFClient) DisableLdapAuthentication(ctx context.Context) (*DisableLdapAuthenticationResult, *SdkError) {
	var res DisableLdapAuthenticationResult
	_, err := sfClient.MakeSFCall(ctx, "DisableLdapAuthentication", 1, nil, &res)
	return &res, err
}

//
func (sfClient *SFClient) ReleaseFreeMemory(ctx context.Context) (*NeedsWorkReleaseFreeMemoryResult, *SdkError) {
	var res NeedsWorkReleaseFreeMemoryResult
	_, err := sfClient.MakeSFCall(ctx, "ReleaseFreeMemory", 1, nil, &res)
	return &res, err
}

// Gets the Vasa Provider info
func (sfClient *SFClient) GetVasaProviderInfo(ctx context.Context) (*VasaProviderInfoResult, *SdkError) {
	var res VasaProviderInfoResult
	_, err := sfClient.MakeSFCall(ctx, "GetVasaProviderInfo", 1, nil, &res)
	return &res, err
}

// RemoveBackupTarget allows you to delete backup targets.
func (sfClient *SFClient) RemoveBackupTarget(ctx context.Context, req *RemoveBackupTargetRequest) (*RemoveBackupTargetResult, *SdkError) {
	var res RemoveBackupTargetResult
	_, err := sfClient.MakeSFCall(ctx, "RemoveBackupTarget", 1, req, &res)
	return &res, err
}

// RemoveVolumePair enables you to remove the remote pairing between two volumes. Use this method on both the source and target volumes that are paired together. When you remove the volume pairing information, data is no longer replicated to or from the volume.
func (sfClient *SFClient) RemoveVolumePair(ctx context.Context, req *RemoveVolumePairRequest) (*RemoveVolumePairResult, *SdkError) {
	var res RemoveVolumePairResult
	_, err := sfClient.MakeSFCall(ctx, "RemoveVolumePair", 1, req, &res)
	return &res, err
}

// ModifySchedule enables you to change the intervals at which a scheduled snapshot occurs. This allows for adjustment to the snapshot frequency and retention.
func (sfClient *SFClient) ModifySchedule(ctx context.Context, req *ModifyScheduleRequest) (*ModifyScheduleResult, *SdkError) {
	var res ModifyScheduleResult
	_, err := sfClient.MakeSFCall(ctx, "ModifySchedule", 1, req, &res)
	return &res, err
}

// DeleteSnapshot enables you to delete a snapshot. A snapshot that is currently the "active" snapshot cannot be deleted. You must
// rollback and make another snapshot "active" before the current snapshot can be deleted. For more details on rolling back snapshots, see RollbackToSnapshot.
func (sfClient *SFClient) DeleteSnapshot(ctx context.Context, req *DeleteSnapshotRequest) (*DeleteSnapshotResult, *SdkError) {
	var res DeleteSnapshotResult
	_, err := sfClient.MakeSFCall(ctx, "DeleteSnapshot", 1, req, &res)
	return &res, err
}

// ModifyAccount enables you to modify an existing account.
// When you lock an account, any existing connections from that account are immediately terminated. When you change an account's
// CHAP settings, any existing connections remain active, and the new CHAP settings are used on subsequent connections or
// reconnections.
// To clear an account's attributes, specify {} for the attributes parameter.
func (sfClient *SFClient) ModifyAccount(ctx context.Context, req *ModifyAccountRequest) (*ModifyAccountResult, *SdkError) {
	var res ModifyAccountResult
	_, err := sfClient.MakeSFCall(ctx, "ModifyAccount", 1, req, &res)
	return &res, err
}

// ListSyncJobs enables you to return information about synchronization jobs that are running on a SolidFire cluster. The type of
// synchronization jobs that are returned with this method are slice, clone, and remote.
func (sfClient *SFClient) ListSyncJobs(ctx context.Context) (*ListSyncJobsResult, *SdkError) {
	var res ListSyncJobsResult
	_, err := sfClient.MakeSFCall(ctx, "ListSyncJobs", 1, nil, &res)
	return &res, err
}

//
func (sfClient *SFClient) TestClientConnectivity(ctx context.Context) (*NeedsWorkTestClientConnectivityResult, *SdkError) {
	var res NeedsWorkTestClientConnectivityResult
	_, err := sfClient.MakeSFCall(ctx, "TestClientConnectivity", 1, nil, &res)
	return &res, err
}

// ListAllNodes enables you to retrieve a list of active and pending nodes in the cluster.
func (sfClient *SFClient) ListAllNodes(ctx context.Context) (*ListAllNodesResult, *SdkError) {
	var res ListAllNodesResult
	_, err := sfClient.MakeSFCall(ctx, "ListAllNodes", 1, nil, &res)
	return &res, err
}

// CancelClone enables you to stop an ongoing CloneVolume or CopyVolume process. When you cancel a group clone operation, the
// system completes and removes the operation's associated asyncHandle.
func (sfClient *SFClient) CancelClone(ctx context.Context, req *CancelCloneRequest) (*CancelCloneResult, *SdkError) {
	var res CancelCloneResult
	_, err := sfClient.MakeSFCall(ctx, "CancelClone", 1, req, &res)
	return &res, err
}

// You can use the RemoveSSLCertificate method to remove the user SSL certificate and private key for the cluster.
// After the certificate and private key are removed, the cluster is configured to use the default certificate and private key.
func (sfClient *SFClient) RemoveSSLCertificate(ctx context.Context) (*RemoveSSLCertificateResult, *SdkError) {
	var res RemoveSSLCertificateResult
	_, err := sfClient.MakeSFCall(ctx, "RemoveSSLCertificate", 1, nil, &res)
	return &res, err
}

//
func (sfClient *SFClient) SetDebugOptions(ctx context.Context) (*NeedsWorkSetDebugOptionsResult, *SdkError) {
	var res NeedsWorkSetDebugOptionsResult
	_, err := sfClient.MakeSFCall(ctx, "SetDebugOptions", 1, nil, &res)
	return &res, err
}

// The SolidFire Element OS web UI uses the ListSnapMirrorEndpoints method to list all SnapMirror endpoints that the SolidFire cluster is communicating with.
func (sfClient *SFClient) ListSnapMirrorEndpoints(ctx context.Context, req *ListSnapMirrorEndpointsRequest) (*ListSnapMirrorEndpointsResult, *SdkError) {
	var res ListSnapMirrorEndpointsResult
	_, err := sfClient.MakeSFCall(ctx, "ListSnapMirrorEndpoints", 1, req, &res)
	return &res, err
}

// You can use the SetLoginBanner method to set the active Terms of Use banner users see when they log on to the web interface.
func (sfClient *SFClient) SetLoginBanner(ctx context.Context, req *SetLoginBannerRequest) (*SetLoginBannerResult, *SdkError) {
	var res SetLoginBannerResult
	_, err := sfClient.MakeSFCall(ctx, "SetLoginBanner", 1, req, &res)
	return &res, err
}

//
func (sfClient *SFClient) NotifyIntentToRestart(ctx context.Context) (*NeedsWorkNotifyIntentToRestartResult, *SdkError) {
	var res NeedsWorkNotifyIntentToRestartResult
	_, err := sfClient.MakeSFCall(ctx, "NotifyIntentToRestart", 1, nil, &res)
	return &res, err
}

// PurgeDeletedVolume immediately and permanently purges a volume that has been deleted. You must delete a volume using
// DeleteVolume before it can be purged. Volumes are purged automatically after a period of time, so usage of this method is not
// typically required.
func (sfClient *SFClient) PurgeDeletedVolume(ctx context.Context, req *PurgeDeletedVolumeRequest) (*PurgeDeletedVolumeResult, *SdkError) {
	var res PurgeDeletedVolumeResult
	_, err := sfClient.MakeSFCall(ctx, "PurgeDeletedVolume", 1, req, &res)
	return &res, err
}

// DeleteInitiators enables you to delete one or more initiators from the system (and from any associated volumes or volume access
// groups).
// If DeleteInitiators fails to delete one of the initiators provided in the parameter, the system returns an error and does not delete any
// initiators (no partial completion is possible).
func (sfClient *SFClient) DeleteInitiators(ctx context.Context, req *DeleteInitiatorsRequest) (*DeleteInitiatorsResult, *SdkError) {
	var res DeleteInitiatorsResult
	_, err := sfClient.MakeSFCall(ctx, "DeleteInitiators", 1, req, &res)
	return &res, err
}

//
func (sfClient *SFClient) ListDatabaseChildrenData(ctx context.Context) (*NeedsWorkListDatabaseChildrenDataResult, *SdkError) {
	var res NeedsWorkListDatabaseChildrenDataResult
	_, err := sfClient.MakeSFCall(ctx, "ListDatabaseChildrenData", 1, nil, &res)
	return &res, err
}

// GetAccountByName enables you to retrieve details about a specific account, given its username.
func (sfClient *SFClient) GetAccountByName(ctx context.Context, req *GetAccountByNameRequest) (*GetAccountResult, *SdkError) {
	var res GetAccountResult
	_, err := sfClient.MakeSFCall(ctx, "GetAccountByName", 1, req, &res)
	return &res, err
}

// GetClusterStats enables you to retrieve high-level activity measurements for the cluster. Values returned are cumulative from the
// creation of the cluster.
func (sfClient *SFClient) GetClusterStats(ctx context.Context) (*GetClusterStatsResult, *SdkError) {
	var res GetClusterStatsResult
	_, err := sfClient.MakeSFCall(ctx, "GetClusterStats", 1, nil, &res)
	return &res, err
}

//
func (sfClient *SFClient) DeleteVolumeSnapMirrorObjectAttributes(ctx context.Context) (*NeedsWorkDeleteVolumeSnapMirrorObjectAttributesResult, *SdkError) {
	var res NeedsWorkDeleteVolumeSnapMirrorObjectAttributesResult
	_, err := sfClient.MakeSFCall(ctx, "DeleteVolumeSnapMirrorObjectAttributes", 1, nil, &res)
	return &res, err
}

// ResetConstants is used to reset constants that have been modified with SetConstants.
// ResetConstants returns a constant/value mapping for all constants that are still set to non-default values after the command completes.
func (sfClient *SFClient) ResetConstants(ctx context.Context, req *ResetConstantsRequest) (*interface{}, *SdkError) {
	var res interface{}
	_, err := sfClient.MakeSFCall(ctx, "ResetConstants", 1, req, &res)
	return &res, err
}

// The SolidFire Element OS web UI uses the ResumeSnapMirrorRelationship method to enable future transfers for a quiesced SnapMirror relationship.
func (sfClient *SFClient) ResumeSnapMirrorRelationship(ctx context.Context, req *ResumeSnapMirrorRelationshipRequest) (*ResumeSnapMirrorRelationshipResult, *SdkError) {
	var res ResumeSnapMirrorRelationshipResult
	_, err := sfClient.MakeSFCall(ctx, "ResumeSnapMirrorRelationship", 1, req, &res)
	return &res, err
}

// You can use ListAsyncResults to list the results of all currently running and completed asynchronous methods on the system.
// Querying asynchronous results with ListAsyncResults does not cause completed asyncHandles to expire; you can use GetAsyncResult
// to query any of the asyncHandles returned by ListAsyncResults.
func (sfClient *SFClient) ListAsyncResults(ctx context.Context, req *ListAsyncResultsRequest) (*ListAsyncResultsResult, *SdkError) {
	var res ListAsyncResultsResult
	_, err := sfClient.MakeSFCall(ctx, "ListAsyncResults", 1, req, &res)
	return &res, err
}

// Retrieves an existing cluster interface preference.
func (sfClient *SFClient) GetClusterInterfacePreference(ctx context.Context, req *GetClusterInterfacePreferenceRequest) (*GetClusterInterfacePreferenceResult, *SdkError) {
	var res GetClusterInterfacePreferenceResult
	_, err := sfClient.MakeSFCall(ctx, "GetClusterInterfacePreference", 1, req, &res)
	return &res, err
}

// You can use the ListQoSPolicies method to list all the settings of all QoS policies on the system.
func (sfClient *SFClient) ListQoSPolicies(ctx context.Context) (*ListQoSPoliciesResult, *SdkError) {
	var res ListQoSPoliciesResult
	_, err := sfClient.MakeSFCall(ctx, "ListQoSPolicies", 1, nil, &res)
	return &res, err
}

//
func (sfClient *SFClient) CopyVolumeToSnapshot(ctx context.Context) (*NeedsWorkCopyVolumeToSnapshotResult, *SdkError) {
	var res NeedsWorkCopyVolumeToSnapshotResult
	_, err := sfClient.MakeSFCall(ctx, "CopyVolumeToSnapshot", 1, nil, &res)
	return &res, err
}

//
func (sfClient *SFClient) DisableClusterSsh(ctx context.Context) (*NeedsWorkDisableClusterSshResult, *SdkError) {
	var res NeedsWorkDisableClusterSshResult
	_, err := sfClient.MakeSFCall(ctx, "DisableClusterSsh", 1, nil, &res)
	return &res, err
}

// The CheckPingOnVlan API method provides the ability to test IP address(s) reachability on a not-yet-configured VLAN.
// This API creates a temporary vlan interface, uses it to ping the provided list of host IP addresses, removes the VLAN interface.
// If the vlan interface already exits TestPing should be used.
func (sfClient *SFClient) CheckPingOnVlan(ctx context.Context, req *CheckPingOnVlanRequest) (*CheckPingOnVlanResult, *SdkError) {
	var res CheckPingOnVlanResult
	_, err := sfClient.MakeSFCall(ctx, "CheckPingOnVlan", 1, req, &res)
	return &res, err
}

// You can use GetPendingOperation to detect an operation on a node that is currently in progress. You can also use this method to report back when an operation has completed.
// Note: method is available only through the per-node API endpoint 5.0 or later.
func (sfClient *SFClient) GetPendingOperation(ctx context.Context) (*GetPendingOperationResult, *SdkError) {
	var res GetPendingOperationResult
	_, err := sfClient.MakeSFCall(ctx, "GetPendingOperation", 1, nil, &res)
	return &res, err
}

// Delete the specified inactive Key Provider.
func (sfClient *SFClient) DeleteKeyProviderKmip(ctx context.Context, req *DeleteKeyProviderKmipRequest) (*DeleteKeyProviderKmipResult, *SdkError) {
	var res DeleteKeyProviderKmipResult
	_, err := sfClient.MakeSFCall(ctx, "DeleteKeyProviderKmip", 1, req, &res)
	return &res, err
}

// Generates a Certificate Sign Request which can be signed by a Certificate Authority to generate a client certificate for the cluster.  This is part of establishing a trust relationship for interacting with external services.
func (sfClient *SFClient) GetClientCertificateSignRequest(ctx context.Context) (*GetClientCertificateSignRequestResult, *SdkError) {
	var res GetClientCertificateSignRequestResult
	_, err := sfClient.MakeSFCall(ctx, "GetClientCertificateSignRequest", 1, nil, &res)
	return &res, err
}

// GetAccountByID enables you to return details about a specific account, given its accountID.
func (sfClient *SFClient) GetAccountByID(ctx context.Context, req *GetAccountByIDRequest) (*GetAccountResult, *SdkError) {
	var res GetAccountResult
	_, err := sfClient.MakeSFCall(ctx, "GetAccountByID", 1, req, &res)
	return &res, err
}

// CloneVolume enables you to create a copy of a volume. This method is asynchronous and might take a variable amount of time to complete. The cloning process begins immediately when you make the CloneVolume request and is representative of the state of the volume when the API method is issued. You can use the GetAsyncResult method to determine when the cloning process is complete and the new volume is available for connections. You can use ListSyncJobs to see the progress of creating the clone.
// Note: The initial attributes and QoS settings for the volume are inherited from the volume being cloned. You can change these settings with ModifyVolume.
// Note: Cloned volumes do not inherit volume access group memberships from the source volume.
func (sfClient *SFClient) CloneVolume(ctx context.Context, req *CloneVolumeRequest) (*CloneVolumeResult, *SdkError) {
	var res CloneVolumeResult
	_, err := sfClient.MakeSFCall(ctx, "CloneVolume", 1, req, &res)
	return &res, err
}

// ListVolumeStatsByVolume returns high-level activity measurements for every volume, by volume. Values are cumulative from the
// creation of the volume.
func (sfClient *SFClient) ListVolumeStatsByVolume(ctx context.Context, req *ListVolumeStatsByVolumeRequest) (*ListVolumeStatsByVolumeResult, *SdkError) {
	var res ListVolumeStatsByVolumeResult
	_, err := sfClient.MakeSFCall(ctx, "ListVolumeStatsByVolume", 1, req, &res)
	return &res, err
}

// You can use the AddVirtualNetwork method to add a new virtual network to a cluster configuration. When you add a virtual network,
// an interface for each node is created and each interface will require a virtual network IP address. The number of IP addresses you
// specify as a parameter for this API method must be equal to or greater than the number of nodes in the cluster. The system bulk
// provisions virtual network addresses and assigns them to individual nodes automatically. You do not need to assign virtual
// network addresses to nodes manually.
// Note: You can use AddVirtualNetwork only to create a new virtual network. If you want to make changes to an
// existing virtual network, use ModifyVirtualNetwork.
// Note: Virtual network parameters must be unique to each virtual network when setting the namespace parameter to false.
func (sfClient *SFClient) AddVirtualNetwork(ctx context.Context, req *AddVirtualNetworkRequest) (*AddVirtualNetworkResult, *SdkError) {
	var res AddVirtualNetworkResult
	_, err := sfClient.MakeSFCall(ctx, "AddVirtualNetwork", 1, req, &res)
	return &res, err
}

// GetVirtualVolumeTaskUpdate checks the status of a VVol Async Task.
func (sfClient *SFClient) GetVirtualVolumeTaskUpdate(ctx context.Context, req *GetVirtualVolumeTaskUpdateRequest) (*VirtualVolumeTaskResult, *SdkError) {
	var res VirtualVolumeTaskResult
	_, err := sfClient.MakeSFCall(ctx, "GetVirtualVolumeTaskUpdate", 1, req, &res)
	return &res, err
}

// Deletes an individual auth session
// If the calling user is not in the ClusterAdmins / Administrator AccessGroup, only auth session belonging
// to the calling user can be deleted.
func (sfClient *SFClient) DeleteAuthSession(ctx context.Context, req *DeleteAuthSessionRequest) (*DeleteAuthSessionResult, *SdkError) {
	var res DeleteAuthSessionResult
	_, err := sfClient.MakeSFCall(ctx, "DeleteAuthSession", 1, req, &res)
	return &res, err
}

// GetLoginSessionInfo enables you to return the period of time a log in authentication session is valid for both log in shells and the TUI.
func (sfClient *SFClient) GetLoginSessionInfo(ctx context.Context) (*GetLoginSessionInfoResult, *SdkError) {
	var res GetLoginSessionInfoResult
	_, err := sfClient.MakeSFCall(ctx, "GetLoginSessionInfo", 1, nil, &res)
	return &res, err
}

//
func (sfClient *SFClient) GetDaemonStatus(ctx context.Context) (*NeedsWorkGetDaemonStatusResult, *SdkError) {
	var res NeedsWorkGetDaemonStatusResult
	_, err := sfClient.MakeSFCall(ctx, "GetDaemonStatus", 1, nil, &res)
	return &res, err
}

// Retrieve the protection schemes supported by the node.
func (sfClient *SFClient) GetProtectionSchemes(ctx context.Context, req *GetProtectionSchemesRequest) (*GetProtectionSchemesResult, *SdkError) {
	var res GetProtectionSchemesResult
	_, err := sfClient.MakeSFCall(ctx, "GetProtectionSchemes", 1, req, &res)
	return &res, err
}

// ModifyVirtualVolumeHost changes an existing ESX host.
func (sfClient *SFClient) ModifyVirtualVolumeHost(ctx context.Context, req *ModifyVirtualVolumeHostRequest) (*VirtualVolumeNullResult, *SdkError) {
	var res VirtualVolumeNullResult
	_, err := sfClient.MakeSFCall(ctx, "ModifyVirtualVolumeHost", 1, req, &res)
	return &res, err
}

// GetClusterMasterNodeID enables you to retrieve the ID of the node that can perform cluster-wide administration tasks and holds the
// storage virtual IP address (SVIP) and management virtual IP address (MVIP).
func (sfClient *SFClient) GetClusterMasterNodeID(ctx context.Context) (*GetClusterMasterNodeIDResult, *SdkError) {
	var res GetClusterMasterNodeIDResult
	_, err := sfClient.MakeSFCall(ctx, "GetClusterMasterNodeID", 1, nil, &res)
	return &res, err
}

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
func (sfClient *SFClient) ResizeDrives(ctx context.Context, req *ResizeDrivesRequest) (*ResizeDrivesResult, *SdkError) {
	var res ResizeDrivesResult
	_, err := sfClient.MakeSFCall(ctx, "ResizeDrives", 1, req, &res)
	return &res, err
}

// Initiate the process of removing the password from self-encrypting drives (SEDs) within the cluster.
func (sfClient *SFClient) DisableEncryptionAtRest(ctx context.Context, req *DisableEncryptionAtRestRequest) (*DisableEncryptionAtRestResult, *SdkError) {
	var res DisableEncryptionAtRestResult
	_, err := sfClient.MakeSFCall(ctx, "DisableEncryptionAtRest", 1, req, &res)
	return &res, err
}

// CreateStorageContainer enables you to create a Virtual Volume (VVol) storage container. Storage containers are associated with a SolidFire storage system account, and are used for reporting and resource allocation. Storage containers can only be associated with virtual volumes. You need at least one storage container to use the Virtual Volumes feature.
func (sfClient *SFClient) CreateStorageContainer(ctx context.Context, req *CreateStorageContainerRequest) (*CreateStorageContainerResult, *SdkError) {
	var res CreateStorageContainerResult
	_, err := sfClient.MakeSFCall(ctx, "CreateStorageContainer", 1, req, &res)
	return &res, err
}

// ListVolumesForAccount returns the list of active and (pending) deleted volumes for an account.
func (sfClient *SFClient) ListVolumesForAccount(ctx context.Context, req *ListVolumesForAccountRequest) (*ListVolumesForAccountResult, *SdkError) {
	var res ListVolumesForAccountResult
	_, err := sfClient.MakeSFCall(ctx, "ListVolumesForAccount", 1, req, &res)
	return &res, err
}

// DeleteVolumeAccessGroup enables you to delete a
// volume access group.
func (sfClient *SFClient) DeleteVolumeAccessGroup(ctx context.Context, req *DeleteVolumeAccessGroupRequest) (*DeleteVolumeAccessGroupResult, *SdkError) {
	var res DeleteVolumeAccessGroupResult
	_, err := sfClient.MakeSFCall(ctx, "DeleteVolumeAccessGroup", 1, req, &res)
	return &res, err
}

//
func (sfClient *SFClient) GetEnsemble(ctx context.Context) (*NeedsWorkGetEnsembleResult, *SdkError) {
	var res NeedsWorkGetEnsembleResult
	_, err := sfClient.MakeSFCall(ctx, "GetEnsemble", 1, nil, &res)
	return &res, err
}

// FastCloneVirtualVolume is used to execute a VMware Virtual Volume fast clone.
func (sfClient *SFClient) FastCloneVirtualVolume(ctx context.Context, req *FastCloneVirtualVolumeRequest) (*VirtualVolumeAsyncResult, *SdkError) {
	var res VirtualVolumeAsyncResult
	_, err := sfClient.MakeSFCall(ctx, "FastCloneVirtualVolume", 1, req, &res)
	return &res, err
}

//
func (sfClient *SFClient) CheckClusterPair(ctx context.Context) (*NeedsWorkCheckClusterPairResult, *SdkError) {
	var res NeedsWorkCheckClusterPairResult
	_, err := sfClient.MakeSFCall(ctx, "CheckClusterPair", 1, nil, &res)
	return &res, err
}

//
func (sfClient *SFClient) ListDatabaseChildren(ctx context.Context) (*NeedsWorkListDatabaseChildrenResult, *SdkError) {
	var res NeedsWorkListDatabaseChildrenResult
	_, err := sfClient.MakeSFCall(ctx, "ListDatabaseChildren", 1, nil, &res)
	return &res, err
}

// UnbindAllVirtualVolumesFromHost removes all VVol <-> Host binding.
func (sfClient *SFClient) UnbindAllVirtualVolumesFromHost(ctx context.Context, req *UnbindAllVirtualVolumesFromHostRequest) (*UnbindAllVirtualVolumesFromHostResult, *SdkError) {
	var res UnbindAllVirtualVolumesFromHostResult
	_, err := sfClient.MakeSFCall(ctx, "UnbindAllVirtualVolumesFromHost", 1, req, &res)
	return &res, err
}

// You can use UpdateBulkVolumeStatus in a script to update the status of a bulk volume job that you started with the
// StartBulkVolumeRead or StartBulkVolumeWrite methods.
func (sfClient *SFClient) UpdateBulkVolumeStatus(ctx context.Context, req *UpdateBulkVolumeStatusRequest) (*UpdateBulkVolumeStatusResult, *SdkError) {
	var res UpdateBulkVolumeStatusResult
	_, err := sfClient.MakeSFCall(ctx, "UpdateBulkVolumeStatus", 1, req, &res)
	return &res, err
}

// Disable support for authentication using third party Identity Providers (IdP) for the cluster.
// Once disabled, users authenticated by third party IdPs will no longer be able to access the cluster and any active authenticated sessions will be invalidated/logged out.
// Ldap and cluster admins will be able to access the cluster via supported UIs.
func (sfClient *SFClient) DisableIdpAuthentication(ctx context.Context) (*DisableIdpAuthenticationResult, *SdkError) {
	var res DisableIdpAuthenticationResult
	_, err := sfClient.MakeSFCall(ctx, "DisableIdpAuthentication", 1, nil, &res)
	return &res, err
}

//
func (sfClient *SFClient) CheckBlock(ctx context.Context) (*NeedsWorkCheckBlockResult, *SdkError) {
	var res NeedsWorkCheckBlockResult
	_, err := sfClient.MakeSFCall(ctx, "CheckBlock", 1, nil, &res)
	return &res, err
}

// AddNodes enables you to add one or more new nodes to a cluster. When a node that is not configured starts up for the first time, you are prompted to configure the node. After you configure the node, it is registered as a "pending node" with the cluster.
// Note: It might take several seconds after adding a new node for it to start up and register its drives as available.
func (sfClient *SFClient) AddNodes(ctx context.Context, req *AddNodesRequest) (*AddNodesResult, *SdkError) {
	var res AddNodesResult
	_, err := sfClient.MakeSFCall(ctx, "AddNodes", 1, req, &res)
	return &res, err
}

// You can use the TestDrives API method to run a hardware validation on all drives on the node. This method detects hardware
// failures on the drives (if present) and reports them in the results of the validation tests.
// You can only use the TestDrives method on nodes that are not "active" in a cluster.
// Note: This test takes approximately 10 minutes.
// Note: This method is available only through the per-node API endpoint 5.0 or later.
func (sfClient *SFClient) TestDrives(ctx context.Context, req *TestDrivesRequest) (*TestDrivesResult, *SdkError) {
	var res TestDrivesResult
	_, err := sfClient.MakeSFCall(ctx, "TestDrives", 1, req, &res)
	return &res, err
}

// You can use ListISCSISessions to return iSCSI information for volumes in the cluster.
func (sfClient *SFClient) ListISCSISessions(ctx context.Context) (*ListISCSISessionsResult, *SdkError) {
	var res ListISCSISessionsResult
	_, err := sfClient.MakeSFCall(ctx, "ListISCSISessions", 1, nil, &res)
	return &res, err
}

// SetNtpInfo enables you to configure NTP on cluster nodes. The values you set with this interface apply to all nodes in the cluster. If an NTP broadcast server periodically broadcasts time information on your network, you can optionally configure nodes as broadcast clients.
// Note: NetApp recommends using NTP servers that are internal to your network, rather than the installation defaults.
func (sfClient *SFClient) SetNtpInfo(ctx context.Context, req *SetNtpInfoRequest) (*SetNtpInfoResult, *SdkError) {
	var res SetNtpInfoResult
	_, err := sfClient.MakeSFCall(ctx, "SetNtpInfo", 1, req, &res)
	return &res, err
}

//
func (sfClient *SFClient) SetDatabaseEntry(ctx context.Context) (*NeedsWorkSetDatabaseEntryResult, *SdkError) {
	var res NeedsWorkSetDatabaseEntryResult
	_, err := sfClient.MakeSFCall(ctx, "SetDatabaseEntry", 1, nil, &res)
	return &res, err
}

// The SolidFire Element OS web UI uses GetOntapVersionInfo to get information about API version support from the ONTAP cluster in a SnapMirror relationship.
func (sfClient *SFClient) GetOntapVersionInfo(ctx context.Context, req *GetOntapVersionInfoRequest) (*GetOntapVersionInfoResult, *SdkError) {
	var res GetOntapVersionInfoResult
	_, err := sfClient.MakeSFCall(ctx, "GetOntapVersionInfo", 1, req, &res)
	return &res, err
}

// You can use the GetSchedule method to retrieve information about a scheduled snapshot. You can see information about a specific
// schedule if there are many snapshot schedules in the system. You also retrieve information about more than one schedule with this
// method by specifying additional scheduleIDs in the parameter.
func (sfClient *SFClient) GetSchedule(ctx context.Context, req *GetScheduleRequest) (*GetScheduleResult, *SdkError) {
	var res GetScheduleResult
	_, err := sfClient.MakeSFCall(ctx, "GetSchedule", 1, req, &res)
	return &res, err
}

// The SolidFire Element OS web UI uses the CreateSnapMirrorRelationship method to create a SnapMirror extended data protection relationship between a source and destination endpoint.
func (sfClient *SFClient) CreateSnapMirrorRelationship(ctx context.Context, req *CreateSnapMirrorRelationshipRequest) (*CreateSnapMirrorRelationshipResult, *SdkError) {
	var res CreateSnapMirrorRelationshipResult
	_, err := sfClient.MakeSFCall(ctx, "CreateSnapMirrorRelationship", 1, req, &res)
	return &res, err
}

// The GetNodeFipsDrivesReport reports the FipsDrives capability of a node. This is a per-node API.
func (sfClient *SFClient) GetNodeFipsDrivesReport(ctx context.Context) (*GetNodeFipsDrivesReportResult, *SdkError) {
	var res GetNodeFipsDrivesReportResult
	_, err := sfClient.MakeSFCall(ctx, "GetNodeFipsDrivesReport", 1, nil, &res)
	return &res, err
}

// CloneVirtualVolume is used to execute a VMware Virtual Volume clone.
func (sfClient *SFClient) CloneVirtualVolume(ctx context.Context, req *CloneVirtualVolumeRequest) (*VirtualVolumeAsyncResult, *SdkError) {
	var res VirtualVolumeAsyncResult
	_, err := sfClient.MakeSFCall(ctx, "CloneVirtualVolume", 1, req, &res)
	return &res, err
}

//
func (sfClient *SFClient) GetClusterSshInfo(ctx context.Context) (*NeedsWorkGetClusterSshInfoResult, *SdkError) {
	var res NeedsWorkGetClusterSshInfoResult
	_, err := sfClient.MakeSFCall(ctx, "GetClusterSshInfo", 1, nil, &res)
	return &res, err
}

//
func (sfClient *SFClient) GetServiceStatus(ctx context.Context) (*NeedsWorkGetServiceStatusResult, *SdkError) {
	var res NeedsWorkGetServiceStatusResult
	_, err := sfClient.MakeSFCall(ctx, "GetServiceStatus", 1, nil, &res)
	return &res, err
}

// You can use SetSnmpTrapInfo to enable and disable the generation of cluster SNMP notifications (traps) and to specify the set of network host computers that receive the notifications. The values you pass with each SetSnmpTrapInfo method call replace all values set in any previous call to SetSnmpTrapInfo.
func (sfClient *SFClient) SetSnmpTrapInfo(ctx context.Context, req *SetSnmpTrapInfoRequest) (*SetSnmpTrapInfoResult, *SdkError) {
	var res SetSnmpTrapInfoResult
	_, err := sfClient.MakeSFCall(ctx, "SetSnmpTrapInfo", 1, req, &res)
	return &res, err
}

// Test whether the specified Key Provider is functioning normally.
func (sfClient *SFClient) TestKeyProviderKmip(ctx context.Context, req *TestKeyProviderKmipRequest) (*TestKeyProviderKmipResult, *SdkError) {
	var res TestKeyProviderKmipResult
	_, err := sfClient.MakeSFCall(ctx, "TestKeyProviderKmip", 1, req, &res)
	return &res, err
}

// ModifyVirtualVolume is used to modify settings on an existing virtual volume.
func (sfClient *SFClient) ModifyVirtualVolume(ctx context.Context, req *ModifyVirtualVolumeRequest) (*VirtualVolumeNullResult, *SdkError) {
	var res VirtualVolumeNullResult
	_, err := sfClient.MakeSFCall(ctx, "ModifyVirtualVolume", 1, req, &res)
	return &res, err
}

// Sets the default protection scheme stored in the cluster info.
func (sfClient *SFClient) SetDefaultProtectionScheme(ctx context.Context, req *SetDefaultProtectionSchemeRequest) (*SetDefaultProtectionSchemeResult, *SdkError) {
	var res SetDefaultProtectionSchemeResult
	_, err := sfClient.MakeSFCall(ctx, "SetDefaultProtectionScheme", 1, req, &res)
	return &res, err
}

//
func (sfClient *SFClient) FindMismatchedSliceLBAs(ctx context.Context) (*NeedsWorkFindMismatchedSliceLBAsResult, *SdkError) {
	var res NeedsWorkFindMismatchedSliceLBAsResult
	_, err := sfClient.MakeSFCall(ctx, "FindMismatchedSliceLBAs", 1, nil, &res)
	return &res, err
}

// The SolidFire Element OS web UI uses the AbortSnapMirrorRelationship method to stop SnapMirror transfers that have started but are not yet complete.
func (sfClient *SFClient) AbortSnapMirrorRelationship(ctx context.Context, req *AbortSnapMirrorRelationshipRequest) (*AbortSnapMirrorRelationshipResult, *SdkError) {
	var res AbortSnapMirrorRelationshipResult
	_, err := sfClient.MakeSFCall(ctx, "AbortSnapMirrorRelationship", 1, req, &res)
	return &res, err
}

// Lists all auth sessions for the given user.
// A caller not in AccessGroup ClusterAdmins / Administrator privileges may only list their own sessions.
// A caller with ClusterAdmins / Administrator privileges may list sessions belonging to any user.
func (sfClient *SFClient) ListAuthSessionsByUsername(ctx context.Context, req *ListAuthSessionsByUsernameRequest) (*ListAuthSessionsResult, *SdkError) {
	var res ListAuthSessionsResult
	_, err := sfClient.MakeSFCall(ctx, "ListAuthSessionsByUsername", 1, req, &res)
	return &res, err
}

// The SolidFire Element OS web UI uses the ListSnapMirrorSchedules method to get a list of schedules that are available on a remote ONTAP cluster.
func (sfClient *SFClient) ListSnapMirrorSchedules(ctx context.Context, req *ListSnapMirrorSchedulesRequest) (*ListSnapMirrorSchedulesResult, *SdkError) {
	var res ListSnapMirrorSchedulesResult
	_, err := sfClient.MakeSFCall(ctx, "ListSnapMirrorSchedules", 1, req, &res)
	return &res, err
}

// SnapshotVirtualVolume is used to take a VMware Virtual Volume snapshot.
func (sfClient *SFClient) SnapshotVirtualVolume(ctx context.Context, req *SnapshotVirtualVolumeRequest) (*SnapshotVirtualVolumeResult, *SdkError) {
	var res SnapshotVirtualVolumeResult
	_, err := sfClient.MakeSFCall(ctx, "SnapshotVirtualVolume", 1, req, &res)
	return &res, err
}

// You can use the SetSupplementalTlsCiphers method to specify the list of supplemental TLS ciphers.
func (sfClient *SFClient) SetSupplementalTlsCiphers(ctx context.Context, req *SetSupplementalTlsCiphersRequest) (*SetSupplementalTlsCiphersResult, *SdkError) {
	var res SetSupplementalTlsCiphersResult
	_, err := sfClient.MakeSFCall(ctx, "SetSupplementalTlsCiphers", 1, req, &res)
	return &res, err
}

// You can use the SetNodeSupplementalTlsCiphers method to specify the list of supplemental TLS ciphers for this node.
// You can use this command on management nodes.
func (sfClient *SFClient) SetNodeSupplementalTlsCiphers(ctx context.Context, req *SetNodeSupplementalTlsCiphersRequest) (*SetNodeSupplementalTlsCiphersResult, *SdkError) {
	var res SetNodeSupplementalTlsCiphersResult
	_, err := sfClient.MakeSFCall(ctx, "SetNodeSupplementalTlsCiphers", 1, req, &res)
	return &res, err
}

// GetDriveStats returns high-level activity measurements for a single drive. Values are cumulative from the addition of the drive to the
// cluster. Some values are specific to block drives. You might not obtain statistical data for both block and metadata drives when you
// run this method.
func (sfClient *SFClient) GetDriveStats(ctx context.Context, req *GetDriveStatsRequest) (*GetDriveStatsResult, *SdkError) {
	var res GetDriveStatsResult
	_, err := sfClient.MakeSFCall(ctx, "GetDriveStats", 1, req, &res)
	return &res, err
}

// This will invoke any API method supported by the SolidFire API for the version and port the connection is using.
// Returns a nested hashtable of key/value pairs that contain the result of the invoked method.
func (sfClient *SFClient) InvokeSFApi(ctx context.Context, req *InvokeSFApiRequest) (*interface{}, *SdkError) {
	var res interface{}
	_, err := sfClient.MakeSFCall(ctx, "InvokeSFApi", 1, req, &res)
	return &res, err
}

// You can use ModifyClusterAdmin to change the settings for a cluster admin, LDAP cluster admin,
// or third party Identity Provider (IdP) cluster admin.  You cannot change access for the
// administrator cluster admin account.
func (sfClient *SFClient) ModifyClusterAdmin(ctx context.Context, req *ModifyClusterAdminRequest) (*ModifyClusterAdminResult, *SdkError) {
	var res ModifyClusterAdminResult
	_, err := sfClient.MakeSFCall(ctx, "ModifyClusterAdmin", 1, req, &res)
	return &res, err
}

// CancelGroupClone enables you to stop an ongoing CloneMultipleVolumes process occurring on a group of volumes. When you cancel
// a group clone operation, the system completes and removes the operation's associated asyncHandle.
func (sfClient *SFClient) CancelGroupClone(ctx context.Context, req *CancelGroupCloneRequest) (*CancelGroupCloneResult, *SdkError) {
	var res CancelGroupCloneResult
	_, err := sfClient.MakeSFCall(ctx, "CancelGroupClone", 1, req, &res)
	return &res, err
}

//
func (sfClient *SFClient) GetCodeTimings(ctx context.Context) (*NeedsWorkGetCodeTimingsResult, *SdkError) {
	var res NeedsWorkGetCodeTimingsResult
	_, err := sfClient.MakeSFCall(ctx, "GetCodeTimings", 1, nil, &res)
	return &res, err
}

// The StartClusterBSCheck API method is used to tell all slice services to go through all of their
// slices and read all the block IDs and read them from all the block services that should have that
// block ID. This API method finds xUnknownBlockID and reports them in the event report. The fields
// volumeID, sliceID, volumeIDs, and sliceIDs are mutually exclusive and if all fields they will be
// handled in the order of volumeID, sliceID, volumeIDSs, then sliceIDs.
func (sfClient *SFClient) StartClusterBSCheck(ctx context.Context, req *StartClusterBSCheckRequest) (*StartClusterBSCheckResult, *SdkError) {
	var res StartClusterBSCheckResult
	_, err := sfClient.MakeSFCall(ctx, "StartClusterBSCheck", 1, req, &res)
	return &res, err
}

// You can use CreateVolumeAccessGroup to create a new volume access group. When you create the volume access group, you need to give it a name, and you can optionally enter initiators and volumes. After you create the group, you can add volumes and initiator IQNs. Any initiator IQN that you add to the volume access group is able to access any volume in the group without CHAP authentication.
func (sfClient *SFClient) CreateVolumeAccessGroup(ctx context.Context, req *CreateVolumeAccessGroupRequest) (*CreateVolumeAccessGroupResult, *SdkError) {
	var res CreateVolumeAccessGroupResult
	_, err := sfClient.MakeSFCall(ctx, "CreateVolumeAccessGroup", 1, req, &res)
	return &res, err
}

// The SolidFire Element OS web UI uses the ListSnapMirrorVolumes method to list all SnapMirror volumes available on a remote ONTAP system.
func (sfClient *SFClient) ListSnapMirrorVolumes(ctx context.Context, req *ListSnapMirrorVolumesRequest) (*ListSnapMirrorVolumesResult, *SdkError) {
	var res ListSnapMirrorVolumesResult
	_, err := sfClient.MakeSFCall(ctx, "ListSnapMirrorVolumes", 1, req, &res)
	return &res, err
}

// You can use the ListUtilities API method to return the operations that are available to run on a node.
// Note: This method is available only through the per-node API endpoint 5.0 or later.
func (sfClient *SFClient) ListUtilities(ctx context.Context, req *ListUtilitiesRequest) (*ListUtilitiesResult, *SdkError) {
	var res ListUtilitiesResult
	_, err := sfClient.MakeSFCall(ctx, "ListUtilities", 1, req, &res)
	return &res, err
}
