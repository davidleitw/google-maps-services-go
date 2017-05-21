// Copyright 2015 Google Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// More information about Google Geolocation API is available on
// https://developers.google.com/maps/documentation/geolocation

package maps

import (
	"fmt"
	"golang.org/x/net/context"
)

var geolocationAPI = &apiConfig{
	host:            "https://www.googleapis.com",
	path:            "/geolocation/v1/geolocate",
	acceptsClientID: true,
}

// Geolocate makes a Geolocation API request
func (c *Client) Geolocate(ctx context.Context, r *GeolocationRequest) (*GeolocationResult, error) {
	var response struct {
		GeolocationResult
		Error GeolocationError
	}
	if err := c.postJSON(ctx, geolocationAPI, r, &response); err != nil {
		return nil, err
	}
	// TODO: much more error detail available here, what do?
	if response.Error.Code != 0 || len(response.Error.Errors) > 0 {
		return nil, fmt.Errorf("%s", response.Error.Message)
	}
	return &response.GeolocationResult, nil
}

// RadioType defines mobile radio types
type RadioType string

// Allowed radio types
const (
	RadioTypeLTE   RadioType = "lte"
	RadioTypeGSM   RadioType = "gsm"
	RadioTypeCDMA  RadioType = "cdma"
	RadioTypeWCDMA RadioType = "wcdma"
)

// CellTower is a cell tower object for localisation requests
type CellTower struct {
	// CellID Unique identifier of the cell
	CellID string `json:"cellId,omitempty"`
	// LocationAreaCode is the Location Area Code (LAC) for GSM and WCDMAnetworks. The Network ID (NID) for CDMA networks.
	LocationAreaCode string `json:"locatonAreaCode,omitempty"`
	// MobileCountryCode is the cell tower's Mobile Country Code (MCC).
	MobileCountryCode string `json:"mobileCountryCode,omitempty"`
	// MobileNetworkCode is the cell tower's Mobile Network Code. This is the MNC for GSM and WCDMA; CDMA uses the System ID (SID).
	MobileNetworkCode string `json:"mobileNetworkCode,omitempty"`
}

// WiFiAccessPoint is a WiFi access point object for localisation requests
type WiFiAccessPoint struct {
	// MacAddress is the MAC address of the WiFi node. Separators must be : (colon).
	MACAddress string `json:"macAddress,omitempty"`
	// SignalStrength is the current signal strength measured in dBm.
	SignalStrength float64 `json:"signalStrength,omitempty"`
	// Age is the number of milliseconds since this access point was detected.
	Age uint64 `json:"age,omitempty"`
	// Channel is the channel over which the client is communicating with the access point.
	Channel int `json:"channel,omitempty"`
	// SignalToNoiseRatio is the current signal to noise ratio measured in dB.
	SignalToNoiseRatio float64 `json:"signalToNoiseRatio,omitempty"`
}

// GeolocationRequest is the request structure for Geolocation API
// All fields are optional
type GeolocationRequest struct {
	// HomeMobileCountryCode is the mobile country code (MCC) for the device's home network.
	HomeMobileCountryCode string `json:"homeMobileCountryCode,omitempty"`
	// HomeMobileNetworkCode is the mobile network code (MNC) for the device's home network.
	HomeMobileNetworkCode string `json:"homeMobileNetworkCode,omitempty"`
	// RadioType is the mobile radio type, this is optional but should be included if available
	RadioType RadioType `json:"radioType,omitempty"`
	// Carrier is the carrier name
	Carrier string `json:"carrier,omitempty"`
	// ConsiderIP Specifies whether to fall back to IP geolocation if wifi and cell tower signals are not available.
	ConsiderIP bool `json:"considerIp"`
	// CellTowers is an array of CellTower objects.
	CellTowers []CellTower `json:"cellTowers,omitempty"`
	// WifiAccessPoints is an array of WifiAccessPoint objects.
	WiFiAccessPoints []WiFiAccessPoint `json:"wifiAccessPoints,omitempty"`
}

// GeolocationResult is an approximate location and accuracy
type GeolocationResult struct {
	// Location is the predicted location
	Location LatLng
	// Accuracy is the accuracy of the provided location in meters
	Accuracy float64
}

// GeolocationError is an error object reporting a request error
type GeolocationError struct {
	Errors []struct {
		Domain  string
		Reason  string
		Message string
	}
	Code    int
	Message string
}
