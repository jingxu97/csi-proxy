package integrationtests

import (
	"context"
	"fmt"
	"math/rand"
	"os/exec"
	"strings"
	"testing"
	"time"

	diskv1beta3 "github.com/kubernetes-csi/csi-proxy/client/api/disk/v1beta3"
	"github.com/kubernetes-csi/csi-proxy/client/api/volume/v1beta3"
	diskv1beta3client "github.com/kubernetes-csi/csi-proxy/client/groups/disk/v1beta3"
	v1beta3client "github.com/kubernetes-csi/csi-proxy/client/groups/volume/v1beta3"
)

func runPowershellCmd(cmd string) (string, error) {
	result, err := exec.Command("powershell", "/c", cmd).CombinedOutput()
	return string(result), err
}

func diskCleanup(t *testing.T, vhdxPath, mountPath, testPluginPath string) {
	if t.Failed() {
		t.Logf("Test failed. Skipping cleanup!")
		t.Logf("Mount path located at %s", mountPath)
		t.Logf("VHDx path located at %s", vhdxPath)
		return
	}
	var cmd, out string
	var err error

	cmd = fmt.Sprintf("Dismount-VHD -Path %s", vhdxPath)
	if out, err = runPowershellCmd(cmd); err != nil {
		t.Errorf("Error: %v. Command: %s. Out: %s", err, cmd, out)
	}
	cmd = fmt.Sprintf("rm %s", vhdxPath)
	if out, err = runPowershellCmd(cmd); err != nil {
		t.Errorf("Error: %v. Command: %s. Out: %s", err, cmd, out)
	}
	cmd = fmt.Sprintf("rmdir %s", mountPath)
	if out, err = runPowershellCmd(cmd); err != nil {
		t.Errorf("Error: %v. Command: %s. Out: %s", err, cmd, out)
	}
	if testPluginPath != "" {
		cmd = fmt.Sprintf("rmdir %s", testPluginPath)
		if out, err = runPowershellCmd(cmd); err != nil {
			t.Errorf("Error: %v. Command: %s. Out: %s", err, cmd, out)
		}
	}
}

func diskInit(t *testing.T, vhdxPath, mountPath, testPluginPath string) string {
	var cmd, out string
	var err error
	const initialSize = 5 * 1024 * 1024 * 1024
	const partitionStyle = "GPT"

	cmd = fmt.Sprintf("mkdir %s", mountPath)
	if out, err = runPowershellCmd(cmd); err != nil {
		t.Fatalf("Error: %v. Command: %s. Out: %s", err, cmd, out)
	}

	// Initialize the tests, using powershell directly.
	// Create the new vhdx
	cmd = fmt.Sprintf("New-VHD -Path %s -SizeBytes %d", vhdxPath, initialSize)
	if out, err = runPowershellCmd(cmd); err != nil {
		t.Fatalf("Error: %v. Command: %s. Out: %s.", err, cmd, out)
	}

	// Mount the vhdx as a disk
	cmd = fmt.Sprintf("Mount-VHD -Path %s", vhdxPath)
	if out, err = runPowershellCmd(cmd); err != nil {
		t.Fatalf("Error: %v. Command: %s. Out: %s", err, cmd, out)
	}

	var diskNum string
	cmd = fmt.Sprintf("(Get-VHD -Path %s).DiskNumber", vhdxPath)
	if diskNum, err = runPowershellCmd(cmd); err != nil {
		t.Fatalf("Error: %v. Command: %s", err, cmd)
	}
	diskNum = strings.TrimRight(diskNum, "\r\n")

	cmd = fmt.Sprintf("Initialize-Disk -Number %s -PartitionStyle %s", diskNum, partitionStyle)
	if _, err = runPowershellCmd(cmd); err != nil {
		t.Fatalf("Error: %v. Command: %s", err, cmd)
	}

	cmd = fmt.Sprintf("New-Partition -DiskNumber %s -UseMaximumSize", diskNum)
	if _, err = runPowershellCmd(cmd); err != nil {
		t.Fatalf("Error: %v. Command: %s", err, cmd)
	}
	return diskNum
}

func runNegativeListVolumeRequest(t *testing.T, client *v1beta3client.Client, diskNum string) {
	listRequest := &v1beta3.ListVolumesOnDiskRequest{
		DiskId: diskNum,
	}
	_, err := client.ListVolumesOnDisk(context.TODO(), listRequest)
	if err == nil {
		t.Fatalf("Empty error. Expected error for disknum:%s", diskNum)
	}
}

func runNegativeIsVolumeFormattedRequest(t *testing.T, client *v1beta3client.Client, volumeID string) {
	isVolumeFormattedRequest := &v1beta3.IsVolumeFormattedRequest{
		VolumeId: volumeID,
	}
	_, err := client.IsVolumeFormatted(context.TODO(), isVolumeFormattedRequest)
	if err == nil {
		t.Fatalf("Empty error. Expected error for volumeID: %s", volumeID)
	}
}

func runNegativeFormatVolumeRequest(t *testing.T, client *v1beta3client.Client, volumeID string) {
	formatVolumeRequest := &v1beta3.FormatVolumeRequest{
		VolumeId: volumeID,
	}
	_, err := client.FormatVolume(context.TODO(), formatVolumeRequest)
	if err == nil {
		t.Fatalf("Empty error. Expected error for volume id: %s", volumeID)
	}
}

func runNegativeResizeVolumeRequest(t *testing.T, client *v1beta3client.Client, volumeID string, size int64) {
	resizeVolumeRequest := &v1beta3.ResizeVolumeRequest{
		VolumeId: volumeID,
		Size:     size,
	}
	_, err := client.ResizeVolume(context.TODO(), resizeVolumeRequest)
	if err == nil {
		t.Fatalf("Error empty for volume resize. Volume: %s, Size: %d", volumeID, size)
	}
}

func runNegativeMountVolumeRequest(t *testing.T, client *v1beta3client.Client, volumeID, mountPath string) {
	// Mount the volume
	mountVolumeRequest := &v1beta3.MountVolumeRequest{
		VolumeId: volumeID,
		Path:     mountPath,
	}

	_, err := client.MountVolume(context.TODO(), mountVolumeRequest)
	if err == nil {
		t.Fatalf("Error empty for volume(%s) mount to path %s.", volumeID, mountPath)
	}
}

func runNegativeDismountVolumeRequest(t *testing.T, client *v1beta3client.Client, volumeID, mountPath string) {
	// Dismount the volume
	dismountVolumeRequest := &v1beta3.DismountVolumeRequest{
		VolumeId: volumeID,
		Path:     mountPath,
	}
	_, err := client.DismountVolume(context.TODO(), dismountVolumeRequest)
	if err == nil {
		t.Fatalf("Empty error. Volume id %s dismount from path %s ", volumeID, mountPath)
	}
}

func runNegativeVolumeStatsRequest(t *testing.T, client *v1beta3client.Client, volumeID string) {
	// Get VolumeStats
	volumeStatsRequest := &v1beta3.VolumeStatsRequest{
		VolumeId: volumeID,
	}
	_, err := client.VolumeStats(context.TODO(), volumeStatsRequest)
	if err == nil {
		t.Errorf("Empty error. VolumeStats for id %s", volumeID)
	}
}

func negativeVolumeTests(t *testing.T) {
	var client *v1beta3client.Client
	var err error

	if client, err = v1beta3client.NewClient(); err != nil {
		t.Fatalf("Client new error: %v", err)
	}
	defer client.Close()

	// Empty volume test
	runNegativeIsVolumeFormattedRequest(t, client, "")
	// -ve volume id
	runNegativeIsVolumeFormattedRequest(t, client, "-1")

	// Format volume negative tests
	runNegativeFormatVolumeRequest(t, client, "")
	runNegativeFormatVolumeRequest(t, client, "-1")

	// Resize volume negative tests
	runNegativeResizeVolumeRequest(t, client, "", 2*1024*1024)
	runNegativeResizeVolumeRequest(t, client, "-1", 2*1024*1024)

	// Mount volume negative tests
	runNegativeMountVolumeRequest(t, client, "", "")
	runNegativeMountVolumeRequest(t, client, "-1", "")

	// Dismount volume negative tests
	runNegativeDismountVolumeRequest(t, client, "", "")
	runNegativeDismountVolumeRequest(t, client, "-1", "")

	runNegativeVolumeStatsRequest(t, client, "")
	runNegativeVolumeStatsRequest(t, client, "-1")
}

func negativeDiskTests(t *testing.T) {
	var client *v1beta3client.Client
	var err error

	if client, err = v1beta3client.NewClient(); err != nil {
		t.Fatalf("Client new error: %v", err)
	}
	defer client.Close()

	// Empty disk Id test
	runNegativeListVolumeRequest(t, client, "")
	// Negative disk id test
	runNegativeListVolumeRequest(t, client, "-1")
}

func simpleE2e(t *testing.T) {
	var volumeClient *v1beta3client.Client
	var diskClient *diskv1beta3client.Client
	var err error

	if volumeClient, err = v1beta3client.NewClient(); err != nil {
		t.Fatalf("Client new error: %v", err)
	}
	defer volumeClient.Close()

	if diskClient, err = diskv1beta3client.NewClient(); err != nil {
		t.Fatalf("DiskClient new error: %v", err)
	}
	defer diskClient.Close()

	s1 := rand.NewSource(time.Now().UTC().UnixNano())
	r1 := rand.New(s1)

	testPluginPath := fmt.Sprintf("C:\\var\\lib\\kubelet\\plugins\\testplugin-%d.csi.io\\", r1.Intn(100))
	mountPath := fmt.Sprintf("%smount-%d", testPluginPath, r1.Intn(100))
	vhdxPath := fmt.Sprintf("%sdisk-%d.vhdx", testPluginPath, r1.Intn(100))

	defer diskCleanup(t, vhdxPath, mountPath, testPluginPath)
	diskNum := diskInit(t, vhdxPath, mountPath, testPluginPath)

	listRequest := &v1beta3.ListVolumesOnDiskRequest{
		DiskId: diskNum,
	}
	listResponse, err := volumeClient.ListVolumesOnDisk(context.TODO(), listRequest)
	if err != nil {
		t.Fatalf("List response: %v", err)
	}

	volumeIDsLen := len(listResponse.VolumeIds)
	if volumeIDsLen != 1 {
		t.Fatalf("Number of volumes not equal to 1: %d", volumeIDsLen)
	}
	volumeID := listResponse.VolumeIds[0]

	isVolumeFormattedRequest := &v1beta3.IsVolumeFormattedRequest{
		VolumeId: volumeID,
	}
	isVolumeFormattedResponse, err := volumeClient.IsVolumeFormatted(context.TODO(), isVolumeFormattedRequest)
	if err != nil {
		t.Fatalf("Is volume formatted request error: %v", err)
	}
	if isVolumeFormattedResponse.Formatted {
		t.Fatal("Volume formatted. Unexpected !!")
	}

	formatVolumeRequest := &v1beta3.FormatVolumeRequest{
		VolumeId: volumeID,
	}
	_, err = volumeClient.FormatVolume(context.TODO(), formatVolumeRequest)
	if err != nil {
		t.Fatalf("Volume format failed. Error: %v", err)
	}

	isVolumeFormattedResponse, err = volumeClient.IsVolumeFormatted(context.TODO(), isVolumeFormattedRequest)
	if err != nil {
		t.Fatalf("Is volume formatted request error: %v", err)
	}
	if !isVolumeFormattedResponse.Formatted {
		t.Fatal("Volume should be formatted. Unexpected !!")
	}

	t.Logf("VolumeId %v", volumeID)
	volumeStatsRequest := &v1beta3.VolumeStatsRequest{
		VolumeId: volumeID,
	}

	volumeStatsResponse, err := volumeClient.VolumeStats(context.TODO(), volumeStatsRequest)
	if err != nil {
		t.Fatalf("VolumeStats request error: %v", err)
	}

	if volumeStatsResponse.VolumeSize == -1 {
		t.Fatalf("VolumeSize reported is not valid, it is %v", volumeStatsResponse.VolumeSize)
	}

	oldSize := volumeStatsResponse.VolumeSize

	resizeVolumeRequest := &v1beta3.ResizeVolumeRequest{
		VolumeId: volumeID,
		// Resize from 5G to 2G
		Size: 2 * 1024 * 1024 * 1024,
	}

	_, err = volumeClient.ResizeVolume(context.TODO(), resizeVolumeRequest)
	if err != nil {
		t.Fatalf("Volume resize request failed. Error: %v", err)
	}

	volumeStatsResponse, err = volumeClient.VolumeStats(context.TODO(), volumeStatsRequest)
	if err != nil {
		t.Fatalf("VolumeStats request after resize error: %v", err)
	}

	if volumeStatsResponse.VolumeSize >= oldSize {
		t.Fatalf("VolumeSize reported is not smaller after resize, it is %v", volumeStatsResponse.VolumeSize)
	}

	volumeDiskNumberRequest := &v1beta3.VolumeDiskNumberRequest{
		VolumeId: volumeID,
	}

	volumeDiskNumberResponse, err := volumeClient.GetVolumeDiskNumber(context.TODO(), volumeDiskNumberRequest)
	if err != nil {
		t.Fatalf("GetVolumeDiskNumber failed: %v", err)
	}

	diskNumberString := fmt.Sprintf("%d", volumeDiskNumberResponse.DiskNumber)

	diskStatsRequest := &diskv1beta3.DiskStatsRequest{
		DiskID: diskNumberString,
	}

	diskStatsResponse, err := diskClient.DiskStats(context.TODO(), diskStatsRequest)
	if err != nil {
		t.Fatalf("DiskStats request error: %v", err)
	}

	if diskStatsResponse.DiskSize < 0 {
		t.Fatalf("Invalid disk size was returned %v", diskStatsResponse.DiskSize)
	}

	// Mount the volume
	mountVolumeRequest := &v1beta3.MountVolumeRequest{
		VolumeId: volumeID,
		Path:     mountPath,
	}
	_, err = volumeClient.MountVolume(context.TODO(), mountVolumeRequest)
	if err != nil {
		t.Fatalf("Volume id %s mount to path %s failed. Error: %v", volumeID, mountPath, err)
	}

	// Dismount the volume
	dismountVolumeRequest := &v1beta3.DismountVolumeRequest{
		VolumeId: volumeID,
		Path:     mountPath,
	}
	_, err = volumeClient.DismountVolume(context.TODO(), dismountVolumeRequest)
	if err != nil {
		t.Fatalf("Volume id %s mount to path %s failed. Error: %v", volumeID, mountPath, err)
	}
}

func TestVolumeAPIs(t *testing.T) {
	// todo: This test will fail on Github Actions because Hyper-V needs to be enabled
	// Skip on GH actions till we find a better solution
	t.Run("SimpleE2E", func(t *testing.T) {
		skipTestOnCondition(t, isRunningOnGhActions())
		simpleE2e(t)
	})

	t.Run("NegativeDiskTests", func(t *testing.T) {
		negativeDiskTests(t)
	})
	t.Run("NegativeVolumeTests", func(t *testing.T) {
		negativeVolumeTests(t)
	})
}
