// Copyright 2022 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package certrotation

import (
	"crypto/x509"
	"encoding/pem"
	"time"

	"github.com/go-logr/logr"
	"github.com/vmware-tanzu/cluster-api-provider-bringyourownhost/agent/registration"
	"k8s.io/client-go/rest"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

//counterfeiter:generate . CertificateRotation
type CertificateRotation interface {
	RotateCertificate() error
}

type BYOHCertificateRotation struct {
	Logger                    logr.Logger
	HostName                  string
	BYOHConfig                *rest.Config
	BootstrapKubeconfigPath   string
	CertificateExpiryDuration int64
}

func (bcr *BYOHCertificateRotation) RotateCertificate() error {
	for {
		block, _ := pem.Decode(bcr.BYOHConfig.CertData)
		if block == nil || block.Type != "CERTIFICATE" {
			bcr.Logger.Info("failed to decode PEM block containing certificate")
			return nil
		}

		cert, err := x509.ParseCertificate(block.Bytes)
		if err != nil {
			bcr.Logger.Error(err, "Certifcate parse failed")
			return err
		}

		totalTimeCert := cert.NotAfter.Sub(cert.NotBefore)

		// if less than 20% time left, renew the certs
		if time.Now().After(cert.NotAfter.Add(totalTimeCert / -5)) {
			bcr.Logger.Info("certificate expiration time left is less than 20%, renewing")
			if err = registration.HandleBootstrapFlow(bcr.Logger, bcr.BootstrapKubeconfigPath, bcr.HostName, bcr.CertificateExpiryDuration); err != nil {
				bcr.Logger.Error(err, "bootstrap flow failed")
			}
		} else {
			bcr.Logger.Info("certificate are valid", "will be renewed after", cert.NotAfter.Add(totalTimeCert/-5))
		}

		// Poll after every 4 seconds
		time.Sleep(4 * time.Second) //nolint
	}
}
