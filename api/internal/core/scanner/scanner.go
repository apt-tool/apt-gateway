package scanner

import (
	"fmt"
	"os/exec"
)

var (
	template = "./scanner --url=%s --blockRegex true --reportFormat json --reportName %s --skipWAFIdentification true --skipWAFBlockCheck true --noEmailReport true"
)

// Scan a host by using apt-scanner
func Scan(host, name string) (Report, error) {
	// create command
	command := fmt.Sprintf(template, host, name)

	// execute command
	cmd := exec.Command(command)
	if err := cmd.Start(); err != nil {
		return Report{}, err
	}

	return Report{}, nil
}
