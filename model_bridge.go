/*
 * Copyright (C) 2022 Quintex Software Solutions Pvt. Ltd. - All Rights Reserved.
 *
 * You may use, distribute and modify this code under the terms of the Apache
 * License Version 2.0. You should have received a copy of the license with this file.
 * If not, please write to : opensource@quintexsoftware.com
 *
 */

package asterisk_ari_go

// The merging of media from one or more channels.  Everyone on the bridge receives the same audio.
type Bridge struct {
	// Bridging class
	BridgeClass string `json:"bridge_class"`
	// Type of bridge technology
	BridgeType string `json:"bridge_type"`
	// Ids of channels participating in this bridge
	Channels []string `json:"channels"`
	// Timestamp when bridge was created
	Creationtime string `json:"creationtime"`
	// Entity that created the bridge
	Creator string `json:"creator"`
	// Unique identifier for this bridge
	Id string `json:"id"`
	// Name the creator gave the bridge
	Name string `json:"name"`
	// Name of the current bridging technology
	Technology string `json:"technology"`
	// The video mode the bridge is using. One of 'none', 'talker', or 'single'.
	VideoMode string `json:"video_mode,omitempty"`
	// The ID of the channel that is the source of video in this bridge, if one exists.
	VideoSourceId string `json:"video_source_id,omitempty"`
}
