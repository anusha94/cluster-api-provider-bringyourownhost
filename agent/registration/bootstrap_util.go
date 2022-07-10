// Copyright 2022 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package registration

import (
	"fmt"

	"github.com/go-logr/logr"
)

func HandleBootstrapFlow(logger logr.Logger, bootstrapKubeconfig, hostName string, certExpiryDuration int64) error {
	logger.Info("initiated bootstrap kubeconfig flow")
	bootstrapClientConfig, err := LoadRESTClientConfig(bootstrapKubeconfig)
	if err != nil {
		return fmt.Errorf("bootstrap client config load failed: %v", err)
	}
	byohCSR, err := NewByohCSR(bootstrapClientConfig, logger, certExpiryDuration)
	if err != nil {
		return fmt.Errorf("ByohCSR intialization failed: %v", err)
	}
	err = byohCSR.BootstrapKubeconfig(hostName)
	if err != nil {
		return fmt.Errorf("kubeconfig generation failed: %v", err)
	}
	return nil
}
