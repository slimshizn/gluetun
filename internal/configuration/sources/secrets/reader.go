package secrets

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/qdm12/gluetun/internal/configuration/sources/files"
)

type Source struct {
	rootDirectory string
	environ       map[string]string
	warner        Warner
	cached        struct {
		wireguardLoaded bool
		wireguardConf   files.WireguardConfig
	}
}

func New(warner Warner) (source *Source) {
	const rootDirectory = "/run/secrets"
	osEnviron := os.Environ()
	environ := make(map[string]string, len(osEnviron))
	for _, pair := range osEnviron {
		const maxSplit = 2
		split := strings.SplitN(pair, "=", maxSplit)
		environ[split[0]] = split[1]
	}

	return &Source{
		rootDirectory: rootDirectory,
		environ:       environ,
		warner:        warner,
	}
}

func (s *Source) String() string { return "secret files" }

func (s *Source) Get(key string) (value string, isSet bool) {
	if key == "" {
		return "", false
	}
	// TODO v4 custom environment variable to set the secrets parent directory
	// and not to set each secret file to a specific path
	envKey := strings.ToUpper(key)
	envKey = strings.ReplaceAll(envKey, "-", "_")
	envKey += "_SECRETFILE" // TODO v4 change _SECRETFILE to _FILE
	path := s.environ[envKey]
	if path == "" {
		path = filepath.Join(s.rootDirectory, key)
	}

	// Special file parsing
	switch key {
	// TODO timezone from /etc/localtime
	case "openvpn_clientcrt", "openvpn_clientkey", "openvpn_encrypted_key":
		value, isSet, err := files.ReadPEMFile(path)
		if err != nil {
			s.warner.Warnf("skipping %s: parsing PEM: %s", path, err)
		}
		return value, isSet
	case "wireguard_private_key":
		privateKey := s.lazyLoadWireguardConf().PrivateKey
		if privateKey != nil {
			return *privateKey, true
		} // else continue to read from individual secret file
	case "wireguard_preshared_key":
		preSharedKey := s.lazyLoadWireguardConf().PreSharedKey
		if preSharedKey != nil {
			return *preSharedKey, true
		} // else continue to read from individual secret file
	case "wireguard_addresses":
		addresses := s.lazyLoadWireguardConf().Addresses
		if addresses != nil {
			return *addresses, true
		} // else continue to read from individual secret file
	case "wireguard_public_key":
		return strPtrToStringIsSet(s.lazyLoadWireguardConf().PublicKey)
	case "wireguard_endpoint_ip":
		return strPtrToStringIsSet(s.lazyLoadWireguardConf().EndpointIP)
	case "wireguard_endpoint_port":
		return strPtrToStringIsSet(s.lazyLoadWireguardConf().EndpointPort)
	}

	value, isSet, err := files.ReadFromFile(path)
	if err != nil {
		s.warner.Warnf("skipping %s: reading file: %s", path, err)
	}
	return value, isSet
}

func (s *Source) KeyTransform(key string) string {
	switch key {
	// TODO v4 remove these irregular cases
	case "OPENVPN_KEY":
		return "openvpn_clientkey"
	case "OPENVPN_CERT":
		return "openvpn_clientcrt"
	case "OPENVPN_ENCRYPTED_KEY":
		return "openvpn_encrypted_key"
	default:
		key = strings.ToLower(key) // HTTPROXY_USER -> httpproxy_user
		return key
	}
}
