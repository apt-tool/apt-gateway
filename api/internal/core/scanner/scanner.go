package scanner

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
)

const (
	reportsDir = "./reports/"
)

var (
	template = "./scanner --url=%s --blockRegex true --reportFormat json --reportName %s --skipWAFIdentification true --skipWAFBlockCheck true --noEmailReport true"
)

// Scan a host by using apt-scanner
func Scan(host, name string) ([]string, error) {
	// create command
	command := fmt.Sprintf(template, host, name)

	// execute command
	cmd := exec.Command(command)
	if err := cmd.Start(); err != nil {
		return report{}.vulnerabilities, err
	}

	// read report file
	context, err := os.ReadFile(reportsDir + name + ".json")
	if err != nil {
		return report{}.vulnerabilities, err
	}

	r := new(report)
	if er := json.Unmarshal(context, r); er != nil {
		return report{}.vulnerabilities, err
	}

	return r.vulnerabilities, nil
}
