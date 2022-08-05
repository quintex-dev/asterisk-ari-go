/*
 * Copyright (C) 2022 Quintex Software Solutions Pvt. Ltd. - All Rights Reserved.
 *
 * You may use, distribute and modify this code under the terms of the Apache
 * License Version 2.0. You should have received a copy of the license with this file.
 * If not, please write to : opensource@quintexsoftware.com
 *
 */

package asterisk_ari_go

// A media file that may be played back.
type Sound struct {
	// The formats and languages in which this sound is available.
	Formats []FormatLangPair `json:"formats"`
	// Sound's identifier.
	Id string `json:"id"`
	// Text description of the sound, usually the words spoken.
	Text string `json:"text,omitempty"`
}
