// -*- Mode: Go; indent-tabs-mode: t -*-
//
// Copyright (C) 2022 Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by gen_commands.py DO NOT EDIT.

package analytics

type CreateAnalyticsModulesFunction struct{}

func (_ *CreateAnalyticsModulesFunction) Request() any {
	return &CreateAnalyticsModules{}
}
func (_ *CreateAnalyticsModulesFunction) Response() any {
	return &CreateAnalyticsModulesResponse{}
}

type CreateRulesFunction struct{}

func (_ *CreateRulesFunction) Request() any {
	return &CreateRules{}
}
func (_ *CreateRulesFunction) Response() any {
	return &CreateRulesResponse{}
}

type DeleteAnalyticsModulesFunction struct{}

func (_ *DeleteAnalyticsModulesFunction) Request() any {
	return &DeleteAnalyticsModules{}
}
func (_ *DeleteAnalyticsModulesFunction) Response() any {
	return &DeleteAnalyticsModulesResponse{}
}

type DeleteRulesFunction struct{}

func (_ *DeleteRulesFunction) Request() any {
	return &DeleteRules{}
}
func (_ *DeleteRulesFunction) Response() any {
	return &DeleteRulesResponse{}
}

type GetAnalyticsModuleOptionsFunction struct{}

func (_ *GetAnalyticsModuleOptionsFunction) Request() any {
	return &GetAnalyticsModuleOptions{}
}
func (_ *GetAnalyticsModuleOptionsFunction) Response() any {
	return &GetAnalyticsModuleOptionsResponse{}
}

type GetAnalyticsModulesFunction struct{}

func (_ *GetAnalyticsModulesFunction) Request() any {
	return &GetAnalyticsModules{}
}
func (_ *GetAnalyticsModulesFunction) Response() any {
	return &GetAnalyticsModulesResponse{}
}

type GetRuleOptionsFunction struct{}

func (_ *GetRuleOptionsFunction) Request() any {
	return &GetRuleOptions{}
}
func (_ *GetRuleOptionsFunction) Response() any {
	return &GetRuleOptionsResponse{}
}

type GetRulesFunction struct{}

func (_ *GetRulesFunction) Request() any {
	return &GetRules{}
}
func (_ *GetRulesFunction) Response() any {
	return &GetRulesResponse{}
}

type GetSupportedAnalyticsModulesFunction struct{}

func (_ *GetSupportedAnalyticsModulesFunction) Request() any {
	return &GetSupportedAnalyticsModules{}
}
func (_ *GetSupportedAnalyticsModulesFunction) Response() any {
	return &GetSupportedAnalyticsModulesResponse{}
}

type GetSupportedRulesFunction struct{}

func (_ *GetSupportedRulesFunction) Request() any {
	return &GetSupportedRules{}
}
func (_ *GetSupportedRulesFunction) Response() any {
	return &GetSupportedRulesResponse{}
}

type ModifyAnalyticsModulesFunction struct{}

func (_ *ModifyAnalyticsModulesFunction) Request() any {
	return &ModifyAnalyticsModules{}
}
func (_ *ModifyAnalyticsModulesFunction) Response() any {
	return &ModifyAnalyticsModulesResponse{}
}

type ModifyRulesFunction struct{}

func (_ *ModifyRulesFunction) Request() any {
	return &ModifyRules{}
}
func (_ *ModifyRulesFunction) Response() any {
	return &ModifyRulesResponse{}
}
