// Copyright (c) 2015-2021 MinIO, Inc.
//
// This file is part of MinIO Object Storage stack
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package cmd

import (
	"crypto/x509"
	"fmt"
	"net"
	"runtime"
	"strings"

	humanize "github.com/dustin/go-humanize"
	"github.com/minio/minio/cmd/config"
	"github.com/minio/minio/cmd/logger"
	color "github.com/minio/minio/pkg/color"
	"github.com/minio/minio/pkg/madmin"
	xnet "github.com/minio/minio/pkg/net"
)

// Documentation links, these are part of message printing code.
const (
	mcQuickStartGuide     = "https://docs.min.io/docs/minio-client-quickstart-guide"
	goQuickStartGuide     = "https://docs.min.io/docs/golang-client-quickstart-guide"
	jsQuickStartGuide     = "https://docs.min.io/docs/javascript-client-quickstart-guide"
	javaQuickStartGuide   = "https://docs.min.io/docs/java-client-quickstart-guide"
	pyQuickStartGuide     = "https://docs.min.io/docs/python-client-quickstart-guide"
	dotnetQuickStartGuide = "https://docs.min.io/docs/dotnet-client-quickstart-guide"
)

// generates format string depending on the string length and padding.
func getFormatStr(strLen int, padding int) string {
	formatStr := fmt.Sprintf("%ds", strLen+padding)
	return "%" + formatStr
}

func mustGetStorageInfo(objAPI ObjectLayer) StorageInfo {
	storageInfo, _ := objAPI.StorageInfo(GlobalContext)
	return storageInfo
}

// Prints the formatted startup message.
func printStartupMessage(apiEndpoints []string, err error) {
	if err != nil {
		logStartupMessage(color.RedBold("Server startup failed with '%v'", err))
		logStartupMessage(color.RedBold("Not all features may be available on this server"))
		logStartupMessage(color.RedBold("Please use 'mc admin' commands to further investigate this issue"))
	}

	strippedAPIEndpoints := stripStandardPorts(apiEndpoints)
	// If cache layer is enabled, print cache capacity.
	cachedObjAPI := newCachedObjectLayerFn()
	if cachedObjAPI != nil {
		printCacheStorageInfo(cachedObjAPI.StorageInfo(GlobalContext))
	}

	// Object layer is initialized then print StorageInfo.
	objAPI := newObjectLayerFn()
	if objAPI != nil {
		printStorageInfo(mustGetStorageInfo(objAPI))
	}

	// Prints credential, region and browser access.
	printServerCommonMsg(strippedAPIEndpoints)

	// Prints `mc` cli configuration message chooses
	// first endpoint as default.
	printCLIAccessMsg(strippedAPIEndpoints[0], "myminio")

	// Prints documentation message.
	printObjectAPIMsg()

	// SSL is configured reads certification chain, prints
	// authority and expiry.
	if color.IsTerminal() && !globalCLIContext.Anonymous {
		if globalIsTLS {
			printCertificateMsg(globalPublicCerts)
		}
	}
}

// Returns true if input is not IPv4, false if it is.
func isNotIPv4(host string) bool {
	h, _, err := net.SplitHostPort(host)
	if err != nil {
		h = host
	}
	ip := net.ParseIP(h)
	ok := ip.To4() != nil // This is always true of IP is IPv4

	// Returns true if input is not IPv4.
	return !ok
}

// strip api endpoints list with standard ports such as
// port "80" and "443" before displaying on the startup
// banner.  Returns a new list of API endpoints.
func stripStandardPorts(apiEndpoints []string) (newAPIEndpoints []string) {
	newAPIEndpoints = make([]string, len(apiEndpoints))
	// Check all API endpoints for standard ports and strip them.
	for i, apiEndpoint := range apiEndpoints {
		u, err := xnet.ParseHTTPURL(apiEndpoint)
		if err != nil {
			continue
		}
		if globalMinioHost == "" && isNotIPv4(u.Host) {
			// Skip all non-IPv4 endpoints when we bind to all interfaces.
			continue
		}
		newAPIEndpoints[i] = u.String()
	}
	return newAPIEndpoints
}

// Prints common server startup message. Prints credential, region and browser access.
func printServerCommonMsg(apiEndpoints []string) {
	// Get saved credentials.
	cred := globalActiveCred

	// Get saved region.
	region := globalServerRegion

	apiEndpointStr := strings.Join(apiEndpoints, "  ")

	// Colorize the message and print.
	logStartupMessage(color.Blue("Endpoint: ") + color.Bold(fmt.Sprintf("%s ", apiEndpointStr)))
	if color.IsTerminal() && !globalCLIContext.Anonymous {
		logStartupMessage(color.Blue("RootUser: ") + color.Bold(fmt.Sprintf("%s ", cred.AccessKey)))
		logStartupMessage(color.Blue("RootPass: ") + color.Bold(fmt.Sprintf("%s ", cred.SecretKey)))
		if region != "" {
			logStartupMessage(color.Blue("Region: ") + color.Bold(fmt.Sprintf(getFormatStr(len(region), 2), region)))
		}
	}
	printEventNotifiers()
}

// Prints bucket notification configurations.
func printEventNotifiers() {
	if globalNotificationSys == nil {
		return
	}

	arns := globalNotificationSys.GetARNList(true)
	if len(arns) == 0 {
		return
	}

	arnMsg := color.Blue("SQS ARNs: ")
	for _, arn := range arns {
		arnMsg += color.Bold(fmt.Sprintf("%s ", arn))
	}

	logStartupMessage(arnMsg)
}

// Prints startup message for command line access. Prints link to our documentation
// and custom platform specific message.
func printCLIAccessMsg(endPoint string, alias string) {
	// Get saved credentials.
	cred := globalActiveCred

	// Configure 'mc', following block prints platform specific information for minio client.
	if color.IsTerminal() && !globalCLIContext.Anonymous {
		logStartupMessage(color.Blue("\nCommand-line Access: ") + mcQuickStartGuide)
		if runtime.GOOS == globalWindowsOSName {
			mcMessage := fmt.Sprintf("$ mc.exe alias set %s %s %s %s", alias,
				endPoint, cred.AccessKey, cred.SecretKey)
			logStartupMessage(fmt.Sprintf(getFormatStr(len(mcMessage), 3), mcMessage))
		} else {
			mcMessage := fmt.Sprintf("$ mc alias set %s %s %s %s", alias,
				endPoint, cred.AccessKey, cred.SecretKey)
			logStartupMessage(fmt.Sprintf(getFormatStr(len(mcMessage), 3), mcMessage))
		}
	}
}

// Prints startup message for Object API acces, prints link to our SDK documentation.
func printObjectAPIMsg() {
	logStartupMessage(color.Blue("\nObject API (Amazon S3 compatible):"))
	logStartupMessage(color.Blue("   Go: ") + fmt.Sprintf(getFormatStr(len(goQuickStartGuide), 8), goQuickStartGuide))
	logStartupMessage(color.Blue("   Java: ") + fmt.Sprintf(getFormatStr(len(javaQuickStartGuide), 6), javaQuickStartGuide))
	logStartupMessage(color.Blue("   Python: ") + fmt.Sprintf(getFormatStr(len(pyQuickStartGuide), 4), pyQuickStartGuide))
	logStartupMessage(color.Blue("   JavaScript: ") + jsQuickStartGuide)
	logStartupMessage(color.Blue("   .NET: ") + fmt.Sprintf(getFormatStr(len(dotnetQuickStartGuide), 6), dotnetQuickStartGuide))
}

// Get formatted disk/storage info message.
func getStorageInfoMsg(storageInfo StorageInfo) string {
	var msg string
	var mcMessage string
	onlineDisks, offlineDisks := getOnlineOfflineDisksStats(storageInfo.Disks)
	if storageInfo.Backend.Type == madmin.Erasure {
		if offlineDisks.Sum() > 0 {
			mcMessage = "Use `mc admin info` to look for latest server/disk info\n"
		}

		diskInfo := fmt.Sprintf(" %d Online, %d Offline. ", onlineDisks.Sum(), offlineDisks.Sum())
		msg += color.Blue("Status:") + fmt.Sprintf(getFormatStr(len(diskInfo), 8), diskInfo)
		if len(mcMessage) > 0 {
			msg = fmt.Sprintf("%s %s", mcMessage, msg)
		}
	}
	return msg
}

// Prints startup message of storage capacity and erasure information.
func printStorageInfo(storageInfo StorageInfo) {
	if msg := getStorageInfoMsg(storageInfo); msg != "" {
		if globalCLIContext.Quiet {
			logger.Info(msg)
		}
		logStartupMessage(msg)
	}
}

func printCacheStorageInfo(storageInfo CacheStorageInfo) {
	msg := fmt.Sprintf("%s %s Free, %s Total", color.Blue("Cache Capacity:"),
		humanize.IBytes(storageInfo.Free),
		humanize.IBytes(storageInfo.Total))
	logStartupMessage(msg)
}

// Prints the certificate expiry message.
func printCertificateMsg(certs []*x509.Certificate) {
	for _, cert := range certs {
		logStartupMessage(config.CertificateText(cert))
	}
}
