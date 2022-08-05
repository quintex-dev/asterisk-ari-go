/*
 * Copyright (C) 2022 Quintex Software Solutions Pvt. Ltd. - All Rights Reserved.
 *
 * You may use, distribute and modify this code under the terms of the Apache
 * License Version 2.0. You should have received a copy of the license with this file.
 * If not, please write to : opensource@quintexsoftware.com
 *
 */

package asterisk_ari_go

// Object representing the playback of media to a channel
type Playback struct {
	// ID for this playback operation
	Id string `json:"id"`
	// For media types that support multiple languages, the language requested for playback.
	Language string `json:"language,omitempty"`
	// The URI for the media currently being played back.
	MediaUri string `json:"media_uri"`
	// If a list of URIs is being played, the next media URI to be played back.
	NextMediaUri string `json:"next_media_uri,omitempty"`
	// Current state of the playback operation.
	State string `json:"state"`
	// URI for the channel or bridge to play the media on
	TargetUri string `json:"target_uri"`
}
