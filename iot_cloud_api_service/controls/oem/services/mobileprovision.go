package services

import (
	"bytes"
	"crypto/x509"
	"time"

	"howett.net/plist"
)

type MobileProvision struct {
	AppIDName                   string    `plist:"AppIDName"`
	ApplicationIdentifierPrefix []string  `plist:"ApplicationIdentifierPrefix"`
	CreationDate                time.Time `plist:"CreationDate"`
	DeveloperCertificates       [][]byte  `plist:"DeveloperCertificates"`
	Entitlements                struct {
		KeychainAccessGroups  []string `plist:"keychain-access-groups"`
		GetTaskAllow          bool     `plist:"get-task-allow"`
		ApplicationIDentifier string   `plist:"application-identifier"`
		TeamIdentifier        string   `plist:"com.apple.developer.team-identifier"`
	} `plist:"Entitlements"`
	ExpirationDate       time.Time `plist:"ExpirationDate"`
	Name                 string    `plist:"Name"`
	ProvisionsAllDevices bool      `plist:"ProvisionsAllDevices"`
	ProvisionedDevices   []string  `plist:"ProvisionedDevices"`
	TeamIdentifier       []string  `plist:"TeamIdentifier"`
	TeamName             string    `plist:"TeamName"`
	TimeToLive           int       `plist:"TimeToLive"`
	UUID                 string    `plist:"UUID"`
	Version              int       `plist:"Version"`
}

func (this *MobileProvision) GetDeveloperCertificates() ([]*x509.Certificate, error) {
	certificates := make([]*x509.Certificate, len(this.DeveloperCertificates))

	for i := range this.DeveloperCertificates {
		certificate, err := x509.ParseCertificate(this.DeveloperCertificates[i])

		if err != nil && certificate == nil {
			return nil, err
		}

		certificates[i] = certificate
	}

	return certificates, nil
}

func (this *MobileProvision) IsProvisionedDevice(udid string) bool {
	return this.ProvisionsAllDevices || func(udid string) bool {
		for _, id := range this.ProvisionedDevices {
			if id == udid {
				return true
			}
		}

		return false
	}(udid)
}

func (this *MobileProvision) IsExpired(time time.Time) bool {
	valid := time.After(this.CreationDate) && time.Before(this.ExpirationDate)
	return !valid
}

func NewMobileProvision(content []byte) (*MobileProvision, error) {
	buf := bytes.NewReader(content)
	var data MobileProvision
	decoder := plist.NewDecoder(buf)
	err := decoder.Decode(&data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}
