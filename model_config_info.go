/*
 * Copyright (C) 2022 Quintex Software Solutions Pvt. Ltd. - All Rights Reserved.
 *
 * You may use, distribute and modify this code under the terms of the Apache
 * License Version 2.0. You should have received a copy of the license with this file.
 * If not, please write to : opensource@quintexsoftware.com
 *
 */

package asterisk_ari_go

// Info about Asterisk configuration
type ConfigInfo struct {
	// Default language for media playback.
	DefaultLanguage string `json:"default_language"`
	// Maximum number of simultaneous channels.
	MaxChannels int32 `json:"max_channels,omitempty"`
	// Maximum load avg on system.
	MaxLoad float64 `json:"max_load,omitempty"`
	// Maximum number of open file handles (files, sockets).
	MaxOpenFiles int32 `json:"max_open_files,omitempty"`
	// Asterisk system name.
	Name string `json:"name"`
	// Effective user/group id for running Asterisk.
	Setid *SetId `json:"setid"`
}
