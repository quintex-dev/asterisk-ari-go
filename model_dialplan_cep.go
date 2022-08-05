/*
 * Copyright (C) 2022 Quintex Software Solutions Pvt. Ltd. - All Rights Reserved.
 *
 * You may use, distribute and modify this code under the terms of the Apache
 * License Version 2.0. You should have received a copy of the license with this file.
 * If not, please write to : opensource@quintexsoftware.com
 *
 */

package asterisk_ari_go

// Dialplan location (context/extension/priority)
type DialplanCep struct {
	// Parameter of current dialplan application
	AppData string `json:"app_data"`
	// Name of current dialplan application
	AppName string `json:"app_name"`
	// Context in the dialplan
	Context string `json:"context"`
	// Extension in the dialplan
	Exten string `json:"exten"`
	// Priority in the dialplan
	Priority int64 `json:"priority"`
}
