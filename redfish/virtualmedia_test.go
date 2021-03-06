//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const virtualMediaBody = `{
	  "@odata.id": "/redfish/v1/Managers/1/VirtualMedia/EXT1",
	  "@odata.context": "/redfish/v1/$metadata#VirtualMedia.VirtualMedia",
	  "@odata.etag": "5fb9f3ba323469f34cf349a889ff49cf",
	  "@odata.type": "#VirtualMedia.v1_3_0.VirtualMedia",
	  "Id": "EXT1",
	  "Name": "VirtualMedia",
	  "Description": "This resource shall be used to represent a virtual media service for a Redfish implementation.",
	  "ConnectedVia": "URI",
	  "Image": "http://192.168.1.2/Core-current.iso",
	  "ImageName": "Core-current.iso",
	  "WriteProtected": true,
	  "Inserted": true,
	  "MediaTypes": [
		"CD",
		"DVD"
	  ]
	}`

// TestVirtualMediaCollection tests the parsing of VirtualMediaCollection objects.
func TestVirtualMedia(t *testing.T) {
	var result VirtualMedia
	err := json.NewDecoder(strings.NewReader(virtualMediaBody)).Decode(&result)

	assert.Nil(t, err, err)
	assert.Equalf(t, result.ODataID, "/redfish/v1/Managers/1/VirtualMedia/EXT1", "Received invalid ODataID: %s", result.ODataID)
	assert.Equalf(t, result.ODataContext, "/redfish/v1/$metadata#VirtualMedia.VirtualMedia", "Received invalid ODataContext: %s", result.ODataContext)
	assert.Equalf(t, result.ODataEtag, "5fb9f3ba323469f34cf349a889ff49cf", "Received invalid ODataEtag: %s", result.ODataEtag)
	assert.Equalf(t, result.ODataType, "#VirtualMedia.v1_3_0.VirtualMedia", "Received invalid ODataType: %s", result.ODataType)
	assert.Equalf(t, result.ID, "EXT1", "Received invalid ID: %s", result.ID)
	assert.Equalf(t, result.Name, "VirtualMedia", "Received invalid Name: %s", result.Name)
	assert.Equalf(t, result.Description, "This resource shall be used to represent a virtual media service for a Redfish implementation.", "Received invalid Description: %s", result.Description)
	assert.Equalf(t, result.ConnectedVia, VirtualMediaConnectedMethod("URI"), "Received invalid ConnectedVia: %s", result.ConnectedVia)
	assert.Equalf(t, result.Image, "http://192.168.1.2/Core-current.iso", "Received invalid Image: %s", result.Image)
	assert.Equalf(t, result.ImageName, "Core-current.iso", "Received invalid ImageName: %s", result.ImageName)
	assert.Equalf(t, result.WriteProtected, true, "Received invalid WriteProtected: %b", result.WriteProtected)
	assert.Equalf(t, result.Inserted, true, "Received invalid Inserted: %b", result.Inserted)
	assert.Equalf(t, len(result.SupportedMediaTypes), 2, "Received invalid SupportedMediaTypes: %d", len(result.SupportedMediaTypes))
	assert.Equalf(t, len(result.rawData) > 0, true, "Raw data not equal: %s", result.rawData)
}
