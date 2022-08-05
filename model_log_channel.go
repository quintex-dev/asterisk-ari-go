/*
 * Copyright (C) 2022 Quintex Software Solutions Pvt. Ltd. - All Rights Reserved.
 *
 * You may use, distribute and modify this code under the terms of the Apache
 * License Version 2.0. You should have received a copy of the license with this file.
 * If not, please write to : opensource@quintexsoftware.com
 *
 */

package asterisk_ari_go

// Details of an Asterisk log channel
type LogChannel struct {
	// The log channel path
	Channel string `json:"channel"`
	// The various log levels
	Configuration string `json:"configuration"`
	// Whether or not a log type is enabled
	Status string `json:"status"`
	// Types of logs for the log channel
	Type_ string `json:"type"`
}
