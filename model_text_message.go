/*
 * Copyright (C) 2022 Quintex Software Solutions Pvt. Ltd. - All Rights Reserved.
 *
 * You may use, distribute and modify this code under the terms of the Apache
 * License Version 2.0. You should have received a copy of the license with this file.
 * If not, please write to : opensource@quintexsoftware.com
 *
 */

package asterisk_ari_go

// A text message.
type TextMessage struct {
	// The text of the message.
	Body string `json:"body"`
	// A technology specific URI specifying the source of the message. For sip and pjsip technologies, any SIP URI can be specified. For xmpp, the URI must correspond to the client connection being used to send the message.
	From string `json:"from"`
	// A technology specific URI specifying the destination of the message. Valid technologies include sip, pjsip, and xmp. The destination of a message should be an endpoint.
	To string `json:"to"`
	// Technology specific key/value pairs (JSON object) associated with the message.
	Variables interface{} `json:"variables,omitempty"`
}
