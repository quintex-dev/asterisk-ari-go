/*
 * Copyright (C) 2022 Quintex Software Solutions Pvt. Ltd. - All Rights Reserved.
 *
 * You may use, distribute and modify this code under the terms of the Apache
 * License Version 2.0. You should have received a copy of the license with this file.
 * If not, please write to : opensource@quintexsoftware.com
 *
 */

package asterisk_ari_go

// A specific communication connection between Asterisk and an Endpoint.
type Channel struct {
	Accountcode string    `json:"accountcode"`
	Caller      *CallerId `json:"caller"`
	// Channel variables
	Channelvars interface{} `json:"channelvars,omitempty"`
	Connected   *CallerId   `json:"connected"`
	// Timestamp when channel was created
	Creationtime string `json:"creationtime"`
	// Current location in the dialplan
	Dialplan *DialplanCep `json:"dialplan"`
	// Unique identifier of the channel.  This is the same as the Uniqueid field in AMI.
	Id string `json:"id"`
	// The default spoken language
	Language string `json:"language"`
	// Name of the channel (i.e. SIP/foo-0000a7e3)
	Name  string `json:"name"`
	State string `json:"state"`
}
