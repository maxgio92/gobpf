package cpupossible

import (
	"io/ioutil"

	"github.com/maxgio92/gobpf/pkg/cpurange"
)

const cpuPossible = "/sys/devices/system/cpu/possible"

// Get returns a slice with the online CPUs, for example `[0, 2, 3]`
func Get() ([]uint, error) {
	buf, err := ioutil.ReadFile(cpuPossible)
	if err != nil {
		return nil, err
	}
	return cpurange.ReadCPURange(string(buf))
}
