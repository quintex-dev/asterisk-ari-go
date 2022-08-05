/*
 * Copyright (C) 2022 Quintex Software Solutions Pvt. Ltd. - All Rights Reserved.
 *
 * You may use, distribute and modify this code under the terms of the Apache
 * License Version 2.0. You should have received a copy of the license with this file.
 * If not, please write to : opensource@quintexsoftware.com
 *
 */

package asterisk_ari_go

// Info about Asterisk status
type StatusInfo struct {
	// Time when Asterisk was last reloaded.
	LastReloadTime string `json:"last_reload_time"`
	// Time when Asterisk was started.
	StartupTime string `json:"startup_time"`
}
