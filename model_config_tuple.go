/*
 * Copyright (C) 2022 Quintex Software Solutions Pvt. Ltd. - All Rights Reserved.
 *
 * You may use, distribute and modify this code under the terms of the Apache
 * License Version 2.0. You should have received a copy of the license with this file.
 * If not, please write to : opensource@quintexsoftware.com
 *
 */

package asterisk_ari_go

// A key/value pair that makes up part of a configuration object.
type ConfigTuple struct {
	// A configuration object attribute.
	Attribute string `json:"attribute"`
	// The value for the attribute.
	Value string `json:"value"`
}
