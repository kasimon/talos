// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package v1alpha1

import (
	"github.com/siderolabs/go-pointer"

	"github.com/siderolabs/talos/pkg/machinery/config/machine"
)

// NetworkConfigOption generates NetworkConfig.
type NetworkConfigOption func(machine.Type, *NetworkConfig) error

// WithNetworkConfig sets whole network config structure, overwrites any previous options.
func WithNetworkConfig(c *NetworkConfig) NetworkConfigOption {
	return func(_ machine.Type, cfg *NetworkConfig) error {
		*cfg = *c

		return nil
	}
}

// WithNetworkNameservers sets global nameservers list.
func WithNetworkNameservers(nameservers ...string) NetworkConfigOption {
	return func(_ machine.Type, cfg *NetworkConfig) error {
		cfg.NameServers = append(cfg.NameServers, nameservers...)

		return nil
	}
}

// WithNetworkInterfaceIgnore marks interface as ignored.
func WithNetworkInterfaceIgnore(iface string) NetworkConfigOption {
	return func(_ machine.Type, cfg *NetworkConfig) error {
		cfg.getDevice(iface).DeviceIgnore = pointer.To(true)

		return nil
	}
}

// WithNetworkInterfaceDHCP enables DHCP for the interface.
func WithNetworkInterfaceDHCP(iface string, enable bool) NetworkConfigOption {
	return func(_ machine.Type, cfg *NetworkConfig) error {
		cfg.getDevice(iface).DeviceDHCP = pointer.To(true)

		return nil
	}
}

// WithNetworkInterfaceDHCPv4 enables DHCPv4 for the interface.
func WithNetworkInterfaceDHCPv4(iface string, enable bool) NetworkConfigOption {
	return func(_ machine.Type, cfg *NetworkConfig) error {
		dev := cfg.getDevice(iface)

		if dev.DeviceDHCPOptions == nil {
			dev.DeviceDHCPOptions = &DHCPOptions{}
		}

		dev.DeviceDHCPOptions.DHCPIPv4 = pointer.To(enable)

		return nil
	}
}

// WithNetworkInterfaceDHCPv6 enables DHCPv6 for the interface.
func WithNetworkInterfaceDHCPv6(iface string, enable bool) NetworkConfigOption {
	return func(_ machine.Type, cfg *NetworkConfig) error {
		dev := cfg.getDevice(iface)

		if dev.DeviceDHCPOptions == nil {
			dev.DeviceDHCPOptions = &DHCPOptions{}
		}

		dev.DeviceDHCPOptions.DHCPIPv6 = pointer.To(enable)

		return nil
	}
}

// WithNetworkInterfaceCIDR configures interface for static addressing.
func WithNetworkInterfaceCIDR(iface, cidr string) NetworkConfigOption {
	return func(_ machine.Type, cfg *NetworkConfig) error {
		cfg.getDevice(iface).DeviceAddresses = append(cfg.getDevice(iface).DeviceAddresses, cidr)

		return nil
	}
}

// WithNetworkInterfaceMTU configures interface MTU.
func WithNetworkInterfaceMTU(iface string, mtu int) NetworkConfigOption {
	return func(_ machine.Type, cfg *NetworkConfig) error {
		cfg.getDevice(iface).DeviceMTU = mtu

		return nil
	}
}

// WithNetworkInterfaceWireguard configures interface for Wireguard.
func WithNetworkInterfaceWireguard(iface string, wireguardConfig *DeviceWireguardConfig) NetworkConfigOption {
	return func(_ machine.Type, cfg *NetworkConfig) error {
		cfg.getDevice(iface).DeviceWireguardConfig = wireguardConfig

		return nil
	}
}

// WithNetworkInterfaceVirtualIP configures interface for Virtual IP.
func WithNetworkInterfaceVirtualIP(iface, cidr string) NetworkConfigOption {
	return func(machineType machine.Type, cfg *NetworkConfig) error {
		if machineType == machine.TypeWorker {
			return nil
		}

		cfg.getDevice(iface).DeviceVIPConfig = &DeviceVIPConfig{
			SharedIP: cidr,
		}

		return nil
	}
}

// WithKubeSpan configures a KubeSpan interface.
func WithKubeSpan() NetworkConfigOption {
	return func(_ machine.Type, cfg *NetworkConfig) error {
		if cfg.NetworkKubeSpan == nil {
			cfg.NetworkKubeSpan = &NetworkKubeSpan{}
		}

		cfg.NetworkKubeSpan.KubeSpanEnabled = pointer.To(true)

		return nil
	}
}
