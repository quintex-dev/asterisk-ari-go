/*
 * Copyright (C) 2022 Quintex Software Solutions Pvt. Ltd. - All Rights Reserved.
 *
 * You may use, distribute and modify this code under the terms of the Apache
 * License Version 2.0. You should have received a copy of the license with this file.
 * If not, please write to : opensource@quintexsoftware.com
 *
 */

package asterisk_ari_go

// Asterisk system information
type AsteriskInfo struct {
	// Info about how Asterisk was built
	Build *BuildInfo `json:"build,omitempty"`
	// Info about Asterisk configuration
	Config *ConfigInfo `json:"config,omitempty"`
	// Info about Asterisk status
	Status *StatusInfo `json:"status,omitempty"`
	// Info about the system running Asterisk
	System *SystemInfo `json:"system,omitempty"`
}
