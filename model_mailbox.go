/*
 * Copyright (C) 2022 Quintex Software Solutions Pvt. Ltd. - All Rights Reserved.
 *
 * You may use, distribute and modify this code under the terms of the Apache
 * License Version 2.0. You should have received a copy of the license with this file.
 * If not, please write to : opensource@quintexsoftware.com
 *
 */

package asterisk_ari_go

// Represents the state of a mailbox.
type Mailbox struct {
	// Name of the mailbox.
	Name string `json:"name"`
	// Count of new messages in the mailbox.
	NewMessages int32 `json:"new_messages"`
	// Count of old messages in the mailbox.
	OldMessages int32 `json:"old_messages"`
}
