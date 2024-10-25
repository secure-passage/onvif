// -*- Mode: Go; indent-tabs-mode: t -*-
//
// Copyright (C) 2022 Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by gen_commands.py DO NOT EDIT.

package media2

type AddConfigurationFunction struct{}

func (_ *AddConfigurationFunction) Request() any {
	return &AddConfiguration{}
}
func (_ *AddConfigurationFunction) Response() any {
	return &AddConfigurationResponse{}
}

type GetAnalyticsConfigurationsFunction struct{}

func (_ *GetAnalyticsConfigurationsFunction) Request() any {
	return &GetAnalyticsConfigurations{}
}
func (_ *GetAnalyticsConfigurationsFunction) Response() any {
	return &GetAnalyticsConfigurationsResponse{}
}

type GetProfilesFunction struct{}

func (_ *GetProfilesFunction) Request() any {
	return &GetProfiles{}
}
func (_ *GetProfilesFunction) Response() any {
	return &GetProfilesResponse{}
}

type RemoveConfigurationFunction struct{}

func (_ *RemoveConfigurationFunction) Request() any {
	return &RemoveConfiguration{}
}
func (_ *RemoveConfigurationFunction) Response() any {
	return &RemoveConfigurationResponse{}
}
