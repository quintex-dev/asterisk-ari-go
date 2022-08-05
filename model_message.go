/*
 * Copyright (C) 2022 Quintex Software Solutions Pvt. Ltd. - All Rights Reserved.
 *
 * You may use, distribute and modify this code under the terms of the Apache
 * License Version 2.0. You should have received a copy of the license with this file.
 * If not, please write to : opensource@quintexsoftware.com
 *
 */

package asterisk_ari_go

// Base type for errors and events
type Message struct {
	// The unique ID for the Asterisk instance that raised this event.
	AsteriskId string `json:"asterisk_id,omitempty"`
	// Indicates the type of this message.
	Type_ string `json:"type"`
}
