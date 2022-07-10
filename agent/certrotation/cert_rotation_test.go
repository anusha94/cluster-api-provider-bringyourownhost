// Copyright 2022 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package certrotation_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/vmware-tanzu/cluster-api-provider-bringyourownhost/agent/certrotation/certrotationfakes"
)

var _ = Describe("Certificate Rotation", func() {

	BeforeEach(func() {
		var (
			fakeCertRotation *certrotationfakes.FakeCertificateRotation
		)
	})
})
