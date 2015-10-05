package wireless

// #cgo LDFLAGS: -liw
// #include <iwlib.h>
// #include <stdlib.h>
import "C"

import (
	"fmt"
	"io/ioutil"
	"strings"
	"unsafe"
)

// NetworkInterface is the name of the wireless interface you wish to perform operations on
type NetworkInterface string

// IwParam is the iwparam struct from wireless.22.h
type IwParam struct {
	Value    int
	Fixed    bool
	Disabled bool
	Flags    uint16
}

// Config is the wireless_config struct from iwlib.h
type Config struct {
	Name      string
	HasNwid   bool
	Nwid      IwParam
	HasFreq   bool
	Freq      float64
	FreqFlags int
	HasKey    bool
	Key       string
	KeySize   int
	KeyFlags  int
	HasEssid  bool
	EssidOn   bool
	Essid     string
	HasMode   bool
	Mode      int
}

func toConfig(config C.struct_wireless_config) Config {
	c := Config{
		Name:     C.GoStringN(&config.name[0], C.IFNAMSIZ),
		HasNwid:  int(config.has_nwid) == 1,
		HasFreq:  int(config.has_freq) == 1,
		Freq:     float64(config.freq),
		HasKey:   int(config.has_key) == 1,
		Key:      C.GoStringN((*C.char)(unsafe.Pointer(&config.key[0])), C.IW_ENCODING_TOKEN_MAX),
		KeySize:  int(config.key_size),
		KeyFlags: int(config.key_flags),
		HasEssid: int(config.has_essid) == 1,
		EssidOn:  int(config.essid_on) == 1,
		Essid:    C.GoStringN(&config.essid[0], C.IW_ESSID_MAX_SIZE),
		HasMode:  int(config.has_mode) == 1,
		Mode:     int(config.mode),
	}

	if c.HasNwid {
		c.Nwid = IwParam{
			Value:    int(config.nwid.value),
			Fixed:    int(config.nwid.fixed) == 1,
			Disabled: int(config.nwid.disabled) == 1,
			Flags:    uint16(config.nwid.flags),
		}
	}

	return c
}

// BasicConfig retrieves the wireless configuration for the given interface name.
func (i NetworkInterface) BasicConfig() (Config, error) {
	sock := C.iw_sockets_open()
	defer C.iw_sockets_close(sock)

	cIface := C.CString(string(i))
	defer C.free(unsafe.Pointer(cIface))

	var wirelessConfig C.struct_wireless_config

	ok := C.iw_get_basic_config(sock, cIface, &wirelessConfig)
	if ok != 0 {
		return Config{}, fmt.Errorf("No wireless extensions")
	}

	return toConfig(wirelessConfig), nil
}

// NetworkInterfaces enumerates /proc/net/wireless to discover all of the available wireless interfaces
func NetworkInterfaces() ([]NetworkInterface, error) {
	b, err := ioutil.ReadFile("/proc/net/wireless")
	if err != nil {
		return nil, err
	}

	var interfaces []NetworkInterface
	for _, line := range strings.Split(string(b), "\n")[2:] {
		parts := strings.Split(line, ":")
		i := strings.TrimSpace(parts[0])
		if i != "" {
			interfaces = append(interfaces, NetworkInterface(i))
		}
	}

	return interfaces, nil
}
