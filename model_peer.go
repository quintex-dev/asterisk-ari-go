/*
 * Copyright (C) 2022 Quintex Software Solutions Pvt. Ltd. - All Rights Reserved.
 *
 * You may use, distribute and modify this code under the terms of the Apache
 * License Version 2.0. You should have received a copy of the license with this file.
 * If not, please write to : opensource@quintexsoftware.com
 *
 */

package asterisk_ari_go

// Detailed information about a remote peer that communicates with Asterisk.
type Peer struct {
	// The IP address of the peer.
	Address string `json:"address,omitempty"`
	// An optional reason associated with the change in peer_status.
	Cause string `json:"cause,omitempty"`
	// The current state of the peer. Note that the values of the status are dependent on the underlying peer technology.
	PeerStatus string `json:"peer_status"`
	// The port of the peer.
	Port string `json:"port,omitempty"`
	// The last known time the peer was contacted.
	Time string `json:"time,omitempty"`
}
