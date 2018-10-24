// Copyright 2018 gopcua authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package services

import (
	"fmt"
)

// SignedSoftwareCertificate represents a SignedSoftwareCertificate.
//
// Specification: Part 4, 7.33
type SignedSoftwareCertificate struct {
	CertificateData []byte
	Signature       []byte
}

// NewSignedSoftwareCertificate creates a new SignedSoftwareCertificate.
func NewSignedSoftwareCertificate(cert, signature []byte) *SignedSoftwareCertificate {
	return &SignedSoftwareCertificate{
		CertificateData: cert,
		Signature:       signature,
	}
}

// datatypes.ByteString returns SignedSoftwareCertificate in string.
func (s *SignedSoftwareCertificate) String() string {
	return fmt.Sprintf("%x, %x", s.CertificateData, s.Signature)
}
