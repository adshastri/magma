/*
 * Copyright (c) Facebook, Inc. and its affiliates.
 * All rights reserved.
 *
 * This source code is licensed under the BSD-style license found in the
 * LICENSE file in the root directory of this source tree.
 */

package models

import (
	"crypto/x509"
	"errors"
	"fmt"

	"github.com/go-openapi/strfmt"
)

const echoKeyType = "ECHO"
const ecdsaKeyType = "SOFTWARE_ECDSA_SHA256"

func (m *NetworkDNSConfig) ValidateModel() error {
	return m.Validate(strfmt.Default)
}

func (m *NetworkFeatures) ValidateModel() error {
	return m.Validate(strfmt.Default)
}

func (m NetworkDNSRecords) ValidateModel() error {
	return m.Validate(strfmt.Default)
}

func (m *MagmadGateway) ValidateModel() error {
	return m.Validate(strfmt.Default)
}

func (m *GatewayDevice) ValidateModel() error {
	if err := m.Key.ValidateModel(); err != nil {
		return err
	}
	return m.Validate(strfmt.Default)
}

func (m *ChallengeKey) ValidateModel() error {
	switch m.KeyType {
	case echoKeyType:
		if m.Key != nil {
			return errors.New("ECHO mode should not have key value")
		}
		return nil
	case ecdsaKeyType:
		if m.Key == nil {
			return fmt.Errorf("No key supplied")
		}
		_, err := x509.ParsePKIXPublicKey(*m.Key)
		if err != nil {
			return fmt.Errorf("Failed to parse key: %s", err)
		}
		return nil
	default:
		return fmt.Errorf("Unknown key type %s", m.KeyType)
	}
}

func (m *MagmadGatewayConfigs) ValidateModel() error {
	if err := m.Validate(strfmt.Default); err != nil {
		return err
	}
	return nil
}
