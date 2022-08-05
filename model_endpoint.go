/*
 * Copyright (C) 2022 Quintex Software Solutions Pvt. Ltd. - All Rights Reserved.
 *
 * You may use, distribute and modify this code under the terms of the Apache
 * License Version 2.0. You should have received a copy of the license with this file.
 * If not, please write to : opensource@quintexsoftware.com
 *
 */

package asterisk_ari_go

// An external device that may offer/accept calls to/from Asterisk.  Unlike most resources, which have a single unique identifier, an endpoint is uniquely identified by the technology/resource pair.
type Endpoint struct {
	// Id's of channels associated with this endpoint
	ChannelIds []string `json:"channel_ids"`
	// Identifier of the endpoint, specific to the given technology.
	Resource string `json:"resource"`
	// Endpoint's state
	State string `json:"state,omitempty"`
	// Technology of the endpoint
	Technology string `json:"technology"`
}
