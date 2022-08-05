/*
 * Copyright (C) 2022 Quintex Software Solutions Pvt. Ltd. - All Rights Reserved.
 *
 * You may use, distribute and modify this code under the terms of the Apache
 * License Version 2.0. You should have received a copy of the license with this file.
 * If not, please write to : opensource@quintexsoftware.com
 *
 */

package asterisk_ari_go

// A statistics of a RTP.
type RtPstat struct {
	// The Asterisk channel's unique ID that owns this instance.
	ChannelUniqueid string `json:"channel_uniqueid"`
	// Maximum jitter on local side.
	LocalMaxjitter float64 `json:"local_maxjitter,omitempty"`
	// Maximum number of packets lost on local side.
	LocalMaxrxploss float64 `json:"local_maxrxploss,omitempty"`
	// Minimum jitter on local side.
	LocalMinjitter float64 `json:"local_minjitter,omitempty"`
	// Minimum number of packets lost on local side.
	LocalMinrxploss float64 `json:"local_minrxploss,omitempty"`
	// Average jitter on local side.
	LocalNormdevjitter float64 `json:"local_normdevjitter,omitempty"`
	// Average number of packets lost on local side.
	LocalNormdevrxploss float64 `json:"local_normdevrxploss,omitempty"`
	// Our SSRC.
	LocalSsrc int32 `json:"local_ssrc"`
	// Standard deviation jitter on local side.
	LocalStdevjitter float64 `json:"local_stdevjitter,omitempty"`
	// Standard deviation packets lost on local side.
	LocalStdevrxploss float64 `json:"local_stdevrxploss,omitempty"`
	// Maximum round trip time.
	Maxrtt float64 `json:"maxrtt,omitempty"`
	// Minimum round trip time.
	Minrtt float64 `json:"minrtt,omitempty"`
	// Average round trip time.
	Normdevrtt float64 `json:"normdevrtt,omitempty"`
	// Maximum jitter on remote side.
	RemoteMaxjitter float64 `json:"remote_maxjitter,omitempty"`
	// Maximum number of packets lost on remote side.
	RemoteMaxrxploss float64 `json:"remote_maxrxploss,omitempty"`
	// Minimum jitter on remote side.
	RemoteMinjitter float64 `json:"remote_minjitter,omitempty"`
	// Minimum number of packets lost on remote side.
	RemoteMinrxploss float64 `json:"remote_minrxploss,omitempty"`
	// Average jitter on remote side.
	RemoteNormdevjitter float64 `json:"remote_normdevjitter,omitempty"`
	// Average number of packets lost on remote side.
	RemoteNormdevrxploss float64 `json:"remote_normdevrxploss,omitempty"`
	// Their SSRC.
	RemoteSsrc int32 `json:"remote_ssrc"`
	// Standard deviation jitter on remote side.
	RemoteStdevjitter float64 `json:"remote_stdevjitter,omitempty"`
	// Standard deviation packets lost on remote side.
	RemoteStdevrxploss float64 `json:"remote_stdevrxploss,omitempty"`
	// Total round trip time.
	Rtt float64 `json:"rtt,omitempty"`
	// Number of packets received.
	Rxcount int32 `json:"rxcount"`
	// Jitter on received packets.
	Rxjitter float64 `json:"rxjitter,omitempty"`
	// Number of octets received.
	Rxoctetcount int32 `json:"rxoctetcount"`
	// Number of received packets lost.
	Rxploss int32 `json:"rxploss"`
	// Standard deviation round trip time.
	Stdevrtt float64 `json:"stdevrtt,omitempty"`
	// Number of packets transmitted.
	Txcount int32 `json:"txcount"`
	// Jitter on transmitted packets.
	Txjitter float64 `json:"txjitter,omitempty"`
	// Number of octets transmitted.
	Txoctetcount int32 `json:"txoctetcount"`
	// Number of transmitted packets lost.
	Txploss int32 `json:"txploss"`
}
