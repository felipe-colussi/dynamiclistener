package cert

import (
	"github.com/sirupsen/logrus"
	v1 "k8s.io/api/core/v1"
)

func IsValidTLSSecret(secret *v1.Secret) bool {
	if secret == nil {
		logrus.Errorf("IsValidTLSSecret nil")
		return false
	}
	if _, ok := secret.Data[v1.TLSCertKey]; !ok {
		logrus.Errorf("IsValidTLSSecret TLSCertKey")

		return false
	}
	if _, ok := secret.Data[v1.TLSPrivateKeyKey]; !ok {
		logrus.Errorf("IsValidTLSSecret TLSPrivateKeyKey")

		return false
	}
	logrus.Errorf("IS OK")
	return true
}
