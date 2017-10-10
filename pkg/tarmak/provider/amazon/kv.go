// Copyright Jetstack Ltd. See LICENSE for details.
package amazon

import (
	"fmt"

	"github.com/jetstack-experimental/vault-unsealer/pkg/kv"
	"github.com/jetstack-experimental/vault-unsealer/pkg/kv/aws_kms"
	"github.com/jetstack-experimental/vault-unsealer/pkg/kv/aws_ssm"
)

func (a *Amazon) secretsKMSKeyID() (string, error) {
	tf := a.tarmak.Terraform()
	output, err := tf.Output(a.tarmak.Cluster().Environment().StateStack())
	if err != nil {
		return "", fmt.Errorf("error getting state stack output: %s", err)
	}

	key := "secrets_kms_arn"

	secretIDIntf, ok := output[key]
	if !ok {
		return "", fmt.Errorf("error could not find '%s' in terraform state output", key)
	}

	secretID, ok := secretIDIntf.(string)
	if !ok {
		return "", fmt.Errorf("error unexpected type for '%s': %T", key, secretIDIntf)
	}

	return secretID, nil

}

func (a *Amazon) vaultUnsealKeyName() (string, error) {
	key := "vault_unseal_key_name"

	keyNameIntf, ok := a.tarmak.Cluster().Environment().VaultStack().Output()[key]
	if !ok {
		return "", fmt.Errorf("error could not find '%s' in terraform vault output", key)
	}

	keyName, ok := keyNameIntf.(string)
	if !ok {
		return "", fmt.Errorf("error unexpected type for '%s': %T", key, keyNameIntf)
	}

	return keyName, nil

}

func (a *Amazon) VaultKV() (kv.Service, error) {
	session, err := a.Session()
	if err != nil {
		return nil, err
	}

	kmsKeyID, err := a.secretsKMSKeyID()
	if err != nil {
		return nil, err
	}

	unsealKeyName, err := a.vaultUnsealKeyName()
	if err != nil {
		return nil, err
	}

	ssm, err := aws_ssm.NewWithSession(session, unsealKeyName)
	if err != nil {
		return nil, fmt.Errorf("error creating Amazon SSM kv store: %s", err.Error())
	}

	kms, err := aws_kms.NewWithSession(session, ssm, kmsKeyID)
	if err != nil {
		return nil, fmt.Errorf("error creating Amazon KMS ID kv store: %s", err.Error())
	}

	return kms, nil
}