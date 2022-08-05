/*
 * Copyright (C) 2022 Quintex Software Solutions Pvt. Ltd. - All Rights Reserved.
 *
 * You may use, distribute and modify this code under the terms of the Apache
 * License Version 2.0. You should have received a copy of the license with this file.
 * If not, please write to : opensource@quintexsoftware.com
 *
 */

package asterisk_ari_go

// Info about how Asterisk was built
type BuildInfo struct {
	// Date and time when Asterisk was built.
	Date string `json:"date"`
	// Kernel version Asterisk was built on.
	Kernel string `json:"kernel"`
	// Machine architecture (x86_64, i686, ppc, etc.)
	Machine string `json:"machine"`
	// Compile time options, or empty string if default.
	Options string `json:"options"`
	// OS Asterisk was built on.
	Os string `json:"os"`
	// Username that build Asterisk
	User string `json:"user"`
}
