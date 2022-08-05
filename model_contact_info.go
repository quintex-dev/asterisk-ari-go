/*
 * Copyright (C) 2022 Quintex Software Solutions Pvt. Ltd. - All Rights Reserved.
 *
 * You may use, distribute and modify this code under the terms of the Apache
 * License Version 2.0. You should have received a copy of the license with this file.
 * If not, please write to : opensource@quintexsoftware.com
 *
 */

package asterisk_ari_go

// Detailed information about a contact on an endpoint.
type ContactInfo struct {
	// The Address of Record this contact belongs to.
	Aor string `json:"aor"`
	// The current status of the contact.
	ContactStatus string `json:"contact_status"`
	// Current round trip time, in microseconds, for the contact.
	RoundtripUsec string `json:"roundtrip_usec,omitempty"`
	// The location of the contact.
	Uri string `json:"uri"`
}
