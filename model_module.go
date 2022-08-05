/*
 * Copyright (C) 2022 Quintex Software Solutions Pvt. Ltd. - All Rights Reserved.
 *
 * You may use, distribute and modify this code under the terms of the Apache
 * License Version 2.0. You should have received a copy of the license with this file.
 * If not, please write to : opensource@quintexsoftware.com
 *
 */

package asterisk_ari_go

// Details of an Asterisk module
type Module struct {
	// The description of this module
	Description string `json:"description"`
	// The name of this module
	Name string `json:"name"`
	// The running status of this module
	Status string `json:"status"`
	// The support state of this module
	SupportLevel string `json:"support_level"`
	// The number of times this module is being used
	UseCount int32 `json:"use_count"`
}
