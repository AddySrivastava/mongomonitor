package mongo

import (
	"fmt"
	"mongomonitor/config"
	"mongomonitor/utils"
	"strings"
)

type AtlasClient struct {
	PublicKey  string
	PrivateKey string
}

func (atlasClient *AtlasClient) GetLogsWithRange(projectId string, hostname string, logName string, startTime int32, endTime int32) []byte {

	// Declare the atlasLogURI and headers
	var atlasLogURI strings.Builder

	headers := make(map[string]string)

	// Populate the values of the varibales in the relative log uri
	populatedLogRelativeUrl := fmt.Sprintf(config.ATLAS_API_RELATIVE_LOG_URI, projectId, hostname, logName)

	// Concat the root uri with the relative URI
	atlasLogURI.WriteString(config.ATLAS_API_ROOT_URI)
	atlasLogURI.WriteString(populatedLogRelativeUrl)

	//Prepare the headers
	headers["Content-Type"] = "application/gzip"
	headers["Accept-Encoding"] = "gzip"

	// Create a HTTP client
	httpClient := utils.HTTPDigest{URI: atlasLogURI.String(), METHOD: "GET", USERNAME: atlasClient.PublicKey, PASSWORD: atlasClient.PrivateKey, BODY: []byte{}, HEADERS: headers}

	//Make HTTP digest call to fetch the logs
	respBody, _, _ := httpClient.MakeRequest()

	return respBody
}

func (atlasClient *AtlasClient) GetPrimaryHostByProjects(projectId string) []byte {

	var clusterInfoURI strings.Builder

	populatedClusterInfoRelativeUrl := fmt.Sprintf(config.ATLAS_CLUSTER_INFO_URI, projectId)

	clusterInfoURI.WriteString(config.ATLAS_API_ROOT_URI)
	clusterInfoURI.WriteString(populatedClusterInfoRelativeUrl)

	httpClient := utils.HTTPDigest{URI: clusterInfoURI.String(), METHOD: "GET", USERNAME: atlasClient.PublicKey, PASSWORD: atlasClient.PrivateKey, BODY: []byte{}}

	respBody, _, _ := httpClient.MakeRequest()

	return respBody

}
