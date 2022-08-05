/*
 * Copyright (C) 2022 Quintex Software Solutions Pvt. Ltd. - All Rights Reserved.
 *
 * You may use, distribute and modify this code under the terms of the Apache
 * License Version 2.0. You should have received a copy of the license with this file.
 * If not, please write to : opensource@quintexsoftware.com
 *
 */

package asterisk_ari_go

// Details of a Stasis application
type Application struct {
	// Id's for bridges subscribed to.
	BridgeIds []string `json:"bridge_ids"`
	// Id's for channels subscribed to.
	ChannelIds []string `json:"channel_ids"`
	// Names of the devices subscribed to.
	DeviceNames []string `json:"device_names"`
	// {tech}/{resource} for endpoints subscribed to.
	EndpointIds []string `json:"endpoint_ids"`
	// Event types sent to the application.
	EventsAllowed []interface{} `json:"events_allowed"`
	// Event types not sent to the application.
	EventsDisallowed []interface{} `json:"events_disallowed"`
	// Name of this application
	Name string `json:"name"`
}
