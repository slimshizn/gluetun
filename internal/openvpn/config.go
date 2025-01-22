package openvpn

import (
	"os"
	"strings"
)

func (c *Configurator) WriteConfig(lines []string) error {
	const permission = 0o644
	file, err := os.OpenFile(c.configPath, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, permission)
	if err != nil {
		return err
	}
	_, err = file.WriteString(strings.Join(lines, "\n"))
	if err != nil {
		_ = file.Close()
		return err
	}

	err = file.Chown(c.puid, c.pgid)
	if err != nil {
		_ = file.Close()
		return err
	}

	return file.Close()
}
