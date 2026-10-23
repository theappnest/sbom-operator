package oci

import (
	"testing"

	"github.com/anchore/stereoscope/pkg/image"
	liboci "github.com/ckotzbauer/libk8soci/pkg/oci"
	"github.com/stretchr/testify/assert"
)

func TestResolveCredentialsDockerHubURLKey(t *testing.T) {
	// what `kubectl create secret docker-registry` writes for Docker Hub
	secret := `{"auths":{"https://index.docker.io/v1/":{"username":"u","password":"p"}}}`
	img := &liboci.RegistryImage{
		ImageID:     "docker.io/org/app@sha256:ee96667981a3b57ff78992369861dd014dd1942c236a8bd88696dfb7a9bbe1c0",
		PullSecrets: []*liboci.KubeCreds{{SecretName: "s", SecretCredsData: []byte(secret)}},
	}

	creds := ResolveCredentials(img, nil)

	assert.Equal(t, "index.docker.io", creds[0].Authority)
	assert.NotNil(t, image.RegistryOptions{Credentials: creds}.Authenticator("index.docker.io"))
}
