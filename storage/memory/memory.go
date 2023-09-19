package memory

import (
	"fmt"
	"runtime"

	"github.com/rancher/dynamiclistener"
	"github.com/sirupsen/logrus"
	v1 "k8s.io/api/core/v1"
)

func New() dynamiclistener.TLSStorage {
	return &memory{}
}

func NewBacked(storage dynamiclistener.TLSStorage) dynamiclistener.TLSStorage {
	return &memory{storage: storage}
}

type memory struct {
	storage dynamiclistener.TLSStorage
	secret  *v1.Secret
}

func (m *memory) Get() (*v1.Secret, error) {
	if m.secret == nil && m.storage != nil {
		secret, err := m.storage.Get()
		if err != nil {
			return nil, err
		}
		m.secret = secret
	}

	return m.secret, nil
}

func (m *memory) Update(secret *v1.Secret) error {
	_, file, no, ok := runtime.Caller(1)
	if ok {
		fmt.Printf("called from %s#%d\n", file, no)
	}

	if m.secret != nil {
		logrus.Errorf("FELIPE - DEBUG memoryUpdate called, m.secret.ResourceVersion: %s  secret.ResourceVersion %s name: %s, namespace: %s", m.secret.ResourceVersion, secret.ResourceVersion, secret.Name, secret.Namespace)

	} else {
		logrus.Error("FELIPE - DEBUG memoryUpdate called, secret nill")

	}
	if m.secret != nil {
		logrus.Errorf("FELIPE - DEBUG  m.secret data  name: %s, namespace %s", m.secret.Name, m.secret.Namespace)
	} else {
		logrus.Error("FELIPE - DEBUG  m.secret is nill")
	}

	if m.secret == nil || m.secret.ResourceVersion == "" || m.secret.ResourceVersion != secret.ResourceVersion {
		if m.storage != nil {
			logrus.Errorf("FELIPE - DEBUG  - Calling storage update")
			if err := m.storage.Update(secret); err != nil {
				return err
			}
		}

		logrus.Infof("Active TLS secret %s/%s (ver=%s) (count %d): %v", secret.Namespace, secret.Name, secret.ResourceVersion, len(secret.Annotations)-1, secret.Annotations)
		m.secret = secret
	}
	return nil
}
