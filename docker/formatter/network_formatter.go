package formatter

import (
	"strconv"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/moncho/dry/docker"
)

const (
	networkIDHeader    = "NETWORK ID"
	name               = "NAME"
	driver             = "DRIVER"
	numberOfContainers = "NUMBER OF CONTAINERS"
	scope              = "SCOPE"
)

//NetworkFormatter knows how to pretty-print the information of an network
type NetworkFormatter struct {
	trunc   bool
	header  []string
	network types.NetworkResource
}

//NewNetworkFormatter creates an network formatter
func NewNetworkFormatter(network types.NetworkResource, trunc bool) *NetworkFormatter {
	return &NetworkFormatter{trunc: trunc, network: network}
}

func (formatter *NetworkFormatter) addHeader(header string) {
	if formatter.header == nil {
		formatter.header = []string{}
	}
	formatter.header = append(formatter.header, strings.ToUpper(header))
}

//ID prettifies the id
func (formatter *NetworkFormatter) ID() string {
	formatter.addHeader(networkIDHeader)
	if formatter.trunc {
		return docker.TruncateID(docker.ImageID(formatter.network.ID))
	}
	return docker.ImageID(formatter.network.ID)
}

//Name prettifies the network name
func (formatter *NetworkFormatter) Name() string {
	formatter.addHeader(name)
	return formatter.network.Name
}

//Driver prettifies the network driver
func (formatter *NetworkFormatter) Driver() string {
	formatter.addHeader(driver)
	return formatter.network.Driver
}

//Containers prettifies the number of containers using the network
func (formatter *NetworkFormatter) Containers() string {
	formatter.addHeader(numberOfContainers)
	if formatter.network.Containers != nil {
		return strconv.Itoa(len(formatter.network.Containers))
	}
	return "0"
}

//Scope prettifies the network scope
func (formatter *NetworkFormatter) Scope() string {
	formatter.addHeader(scope)
	return formatter.network.Scope
}
