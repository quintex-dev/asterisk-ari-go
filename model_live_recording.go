/*
 * Copyright (C) 2022 Quintex Software Solutions Pvt. Ltd. - All Rights Reserved.
 *
 * You may use, distribute and modify this code under the terms of the Apache
 * License Version 2.0. You should have received a copy of the license with this file.
 * If not, please write to : opensource@quintexsoftware.com
 *
 */

package asterisk_ari_go

// A recording that is in progress
type LiveRecording struct {
	// Cause for recording failure if failed
	Cause string `json:"cause,omitempty"`
	// Duration in seconds of the recording
	Duration int32 `json:"duration,omitempty"`
	// Recording format (wav, gsm, etc.)
	Format string `json:"format"`
	// Base name for the recording
	Name string `json:"name"`
	// Duration of silence, in seconds, detected in the recording. This is only available if the recording was initiated with a non-zero maxSilenceSeconds.
	SilenceDuration int32  `json:"silence_duration,omitempty"`
	State           string `json:"state"`
	// Duration of talking, in seconds, detected in the recording. This is only available if the recording was initiated with a non-zero maxSilenceSeconds.
	TalkingDuration int32 `json:"talking_duration,omitempty"`
	// URI for the channel or bridge being recorded
	TargetUri string `json:"target_uri"`
}
