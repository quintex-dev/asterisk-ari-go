/*
 * Copyright (C) 2022 Quintex Software Solutions Pvt. Ltd. - All Rights Reserved.
 *
 * You may use, distribute and modify this code under the terms of the Apache
 * License Version 2.0. You should have received a copy of the license with this file.
 * If not, please write to : opensource@quintexsoftware.com
 *
 */

package asterisk_ari_go

// Asterisk ping information
type AsteriskPing struct {
	// Asterisk id info
	AsteriskId string `json:"asterisk_id"`
	// Always string value is pong
	Ping string `json:"ping"`
	// The timestamp string of request received time
	Timestamp string `json:"timestamp"`
}
