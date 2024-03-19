package services

import (
	"encoding/asn1"
)

type ContentInfo struct {
	ContentType asn1.ObjectIdentifier
	Content     SignedData `asn1:"tag:0"`
}

type SignedData struct {
	Sequence Sequence
}

type Sequence struct {
	Version          int
	DigestAlgorithms interface{}
	EncapContentInfo EncapsulatedContentInfo
	Certificates     interface{}
	Crls             interface{}
	SignerInfos      interface{}
}

type SignerInfo struct {
	Version int
}

type EncapsulatedContentInfo struct {
	ContentType asn1.ObjectIdentifier
	Content     struct {
		Content []byte
	} `asn1:"tag:0"`
}

func (this *ContentInfo) GetContent() (*MobileProvision, error) {
	return NewMobileProvision(this.Content.Sequence.EncapContentInfo.Content.Content)
}

func NewContentInfo(buffer []byte) (*ContentInfo, error) {
	var info ContentInfo
	asn1.Unmarshal(buffer, &info)
	return &info, nil
}
