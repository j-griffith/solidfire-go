package cloudops

import (
	"context"
	"fmt"
	"github.com/j-griffith/solidfire-go/sdk"
	"github.com/kubernetes-csi/csi-lib-iscsi/iscsi"
	"gopkg.in/yaml.v2"
	"log"
	"os"
	"strings"
)

const GiB = 1073741824

type Client struct {
	SFClient          *sdk.SFClient
	Endpoint          string
	URL               string
	Login             string
	Password          string
	Version           string
	SVIP              string
	DefaultVolumeSize string
	InitiatorIface    string
	TargetSecret      string
	InitiatorSecret   string
	TenantName        string
	AccountID         int64
}

func parseEndpointString(ep string, c *Client) error {
	items := strings.Split(ep, "/")
	creds := strings.Split(items[2], "@")
	login := strings.Split(creds[0], ":")

	c.URL = creds[1]
	c.Login = login[0]
	c.Password = login[1]
	c.Version = items[4]
	return nil

}

func NewClient(c string) (*Client, error) {
	var client Client

	err := yaml.Unmarshal([]byte(c), &client)
	if err != nil {
		log.Printf("failure parsing supplied config yaml: %v\n", err)
		return &client, err
	}
	if parseEndpointString(client.Endpoint, &client) != nil {
		log.Printf("failure parsing endpoint string: %v\n", err)
		os.Exit(1)
	}

	var sf sdk.SFClient
	ctx := context.Background()
	sf.Connect(ctx, client.URL, client.Version, client.Login, client.Password)

	// We want to persist the connection info we created above, otherwise ever call is prefaced with
	// this connect routine (blek)
	client.SFClient = &sf

	if err != nil {
		log.Printf("failure verifying endpoint config while conducting initial client connection: %v\n", err)
		os.Exit(1)
	}
	// Verify specified user tenant/account and ge it's ID, if it doesn't exist
	// create it
	req := sdk.GetAccountByNameRequest{}
	req.Username = client.TenantName
	result, sdkErr := client.SFClient.GetAccountByName(ctx, &req)
	if sdkErr != nil {
		if sdkErr.Detail == "500:xUnknownAccount" {
			req := sdk.AddAccountRequest{}
			req.Username = client.TenantName
			result, sdkErr := client.SFClient.AddAccount(ctx, &req)
			if sdkErr != nil {
				log.Printf("failed to create default account: %+v\n", sdkErr)
				os.Exit(1)
			}
			client.AccountID = result.Account.AccountID

		}
	}
	client.AccountID = result.Account.AccountID
	client.InitiatorSecret = result.Account.InitiatorSecret
	client.TargetSecret = result.Account.TargetSecret

	if client.InitiatorIface == "" {
		client.InitiatorIface = "default"
	}

	return &client, nil
}

func (c *Client) CreateVolume(req sdk.CreateVolumeRequest) (*sdk.Volume, error) {
	ctx := context.Background()
	v, sdkErr := c.SFClient.CreateVolume(ctx, &req)
	if sdkErr != nil {
		log.Printf("failed to create volume: %+v\n", sdkErr)
		return &sdk.Volume{}, sdkErr
	}
	return &v.Volume, nil
}

func (c *Client) DeleteVolume(volumeID int64) error {
	req := sdk.DeleteVolumeRequest{}
	req.VolumeID = volumeID
	ctx := context.Background()
	_, err := c.SFClient.DeleteVolume(ctx, &req)
	if err != nil {
		dneString := fmt.Sprintf("500:Volume %d does not exist.", req.VolumeID)
		if err.Detail == dneString {
			err = nil
		}
	}
	return err
}

func (c *Client) ExpandVolume(volumeID, newSize int64) error {
	req := sdk.ModifyVolumeRequest{}
	req.VolumeID = volumeID
	req.TotalSize = newSize * GiB

	ctx := context.Background()
	_, err := c.SFClient.ModifyVolume(ctx, &req)
	return err
}

func (c *Client) ModifyQoS(volumeID int64, qos *sdk.QoS) error {
	req := sdk.ModifyVolumeRequest{}
	req.VolumeID = volumeID
	req.Qos = *qos

	ctx := context.Background()
	_, err := c.SFClient.ModifyVolume(ctx, &req)
	return err
}

func (c *Client) GetVolume(volumeID int64) (*sdk.Volume, error) {
	req := sdk.ListVolumesForAccountRequest{}
	req.AccountID = c.AccountID
	req.StartVolumeID = volumeID
	req.Limit = 1

	ctx := context.Background()
	response, err := c.SFClient.ListVolumesForAccount(ctx, &req)
	if len(response.Volumes) > 0 {
		return &response.Volumes[0], err
	}
	return &sdk.Volume{}, err
}

func (c *Client) ListVolumes() ([]sdk.Volume, error) {
	req := sdk.ListVolumesForAccountRequest{}
	req.AccountID = c.AccountID

	ctx := context.Background()
	response, err := c.SFClient.ListVolumesForAccount(ctx, &req)
	return response.Volumes, err
}

func (c *Client) buildConnector(volumeID int64) (*iscsi.Connector, error) {
	sfVolume, err := c.GetVolume(volumeID)
	if err != nil {
	}

	chapSecret := iscsi.Secrets{
		SecretsType: "chap",
		UserName:    c.TenantName,
		Password:    c.TargetSecret,
		UserNameIn:  c.TenantName,
		PasswordIn:  c.InitiatorSecret,
	}

	tgtInfo := iscsi.TargetInfo{
		Iqn:    sfVolume.Iqn,
		Portal: c.SVIP,
	}
	connector := &iscsi.Connector{
		AuthType:         "chap",
		Targets:          []iscsi.TargetInfo{tgtInfo},
		DoCHAPDiscovery:  true,
		DiscoverySecrets: chapSecret,
		SessionSecrets:   chapSecret,
		Interface:        c.InitiatorIface,
	}
	return connector, nil
}

func (c *Client) ConnectVolume(volumeID int64) (string, error) {
	connector, _ := c.buildConnector(volumeID)
	path, err := iscsi.Connect(*connector)
	if err != nil {
	}

	return path, nil
}
