// Copyright (c) 2017-2020, AT&T Intellectual Property.
// All rights reserved.
//
// Copyright (c) 2014-2017 by Brocade Communications Systems, Inc.
// All rights reserved.
//
// SPDX-License-Identifier: MPL-2.0

package parse

import "errors"

type Cardinality struct {
	Start, End rune
}

type NodeCardinality func(NodeType) map[NodeType]Cardinality

func yangCardinality(ntype NodeType) map[NodeType]Cardinality {
	return cardinalities[ntype]
}

//A table of cardinaliies from the rfc.
var cardinalities = map[NodeType]map[NodeType]Cardinality{
	NodeModule: {
		NodeAnyxml:       {'0', 'n'},
		NodeAugment:      {'0', 'n'},
		NodeChoice:       {'0', 'n'},
		NodeContact:      {'0', '1'},
		NodeContainer:    {'0', 'n'},
		NodeDescription:  {'0', '1'},
		NodeDeviation:    {'0', 'n'},
		NodeExtension:    {'0', 'n'},
		NodeFeature:      {'0', 'n'},
		NodeGrouping:     {'0', 'n'},
		NodeIdentity:     {'0', 'n'},
		NodeImport:       {'0', 'n'},
		NodeInclude:      {'0', 'n'},
		NodeLeaf:         {'0', 'n'},
		NodeLeafList:     {'0', 'n'},
		NodeList:         {'0', 'n'},
		NodeNamespace:    {'1', '1'},
		NodeNotification: {'0', 'n'},
		NodeOrganization: {'0', '1'},
		NodePrefix:       {'1', '1'},
		NodeReference:    {'0', '1'},
		NodeRevision:     {'0', 'n'},
		NodeRpc:          {'0', 'n'},
		NodeTypedef:      {'0', 'n'},
		NodeUses:         {'0', 'n'},
		NodeYangVersion:  {'0', '1'},
		NodeOpdAugment:   {'0', 'n'},
		NodeOpdCommand:   {'0', 'n'},
		NodeOpdOption:    {'0', 'n'},
	},
	NodeImport: {
		NodePrefix:       {'1', '1'},
		NodeRevisionDate: {'0', '1'},
	},
	NodeInclude: {
		NodeRevisionDate: {'0', '1'},
	},
	NodeRevision: {
		NodeDescription: {'0', '1'},
		NodeReference:   {'0', '1'},
	},
	NodeSubmodule: {
		NodeAnyxml:       {'0', 'n'},
		NodeAugment:      {'0', 'n'},
		NodeBelongsTo:    {'1', '1'},
		NodeChoice:       {'0', 'n'},
		NodeContact:      {'0', '1'},
		NodeContainer:    {'0', 'n'},
		NodeDescription:  {'0', '1'},
		NodeDeviation:    {'0', 'n'},
		NodeExtension:    {'0', 'n'},
		NodeFeature:      {'0', 'n'},
		NodeGrouping:     {'0', 'n'},
		NodeIdentity:     {'0', 'n'},
		NodeImport:       {'0', 'n'},
		NodeInclude:      {'0', 'n'},
		NodeLeaf:         {'0', 'n'},
		NodeLeafList:     {'0', 'n'},
		NodeList:         {'0', 'n'},
		NodeNotification: {'0', 'n'},
		NodeOrganization: {'0', 'n'},
		NodeReference:    {'0', '1'},
		NodeRevision:     {'0', 'n'},
		NodeRpc:          {'0', 'n'},
		NodeTypedef:      {'0', 'n'},
		NodeUses:         {'0', 'n'},
		NodeYangVersion:  {'0', '1'},
	},
	NodeBelongsTo: {
		NodePrefix: {'1', '1'},
	},
	NodeTypedef: {
		NodeDefault:     {'0', '1'},
		NodeDescription: {'0', '1'},
		NodeReference:   {'0', '1'},
		NodeStatus:      {'0', '1'},
		NodeTyp:         {'1', '1'},
		NodeUnits:       {'0', '1'},
	},
	NodeTyp: {
		NodeBase:            {'0', '1'},
		NodeBit:             {'0', 'n'},
		NodeEnum:            {'0', 'n'},
		NodeFractionDigits:  {'0', '1'},
		NodeLength:          {'0', '1'},
		NodePath:            {'0', '1'},
		NodePattern:         {'0', 'n'},
		NodeRange:           {'0', '1'},
		NodeRequireInstance: {'0', '1'},
		NodeTyp:             {'0', 'n'},
	},
	NodeContainer: {
		NodeAnyxml:      {'0', 'n'},
		NodeChoice:      {'0', 'n'},
		NodeConfig:      {'0', '1'},
		NodeContainer:   {'0', 'n'},
		NodeDescription: {'0', '1'},
		NodeGrouping:    {'0', 'n'},
		NodeIfFeature:   {'0', 'n'},
		NodeLeaf:        {'0', 'n'},
		NodeLeafList:    {'0', 'n'},
		NodeList:        {'0', 'n'},
		NodeMust:        {'0', 'n'},
		NodePresence:    {'0', '1'},
		NodeReference:   {'0', '1'},
		NodeStatus:      {'0', '1'},
		NodeTypedef:     {'0', 'n'},
		NodeUses:        {'0', 'n'},
		NodeWhen:        {'0', '1'},
	},
	NodeOpdCommand: {
		NodeAnyxml:         {'0', 'n'},
		NodeDescription:    {'0', '1'},
		NodeIfFeature:      {'0', 'n'},
		NodeMust:           {'0', 'n'},
		NodeReference:      {'0', '1'},
		NodeStatus:         {'0', '1'},
		NodeUses:           {'0', 'n'},
		NodeWhen:           {'0', '1'},
		NodeOpdCommand:     {'0', 'n'},
		NodeOpdOption:      {'0', 'n'},
		NodeOpdArgument:    {'0', '1'},
		NodeOpdOnEnter:     {'0', '1'},
		NodeOpdInherit:     {'0', '1'},
		NodeOpdRepeatable:  {'0', '1'},
		NodeOpdPassOpcArgs: {'0', '1'},
		NodeOpdPrivileged:  {'0', '1'},
		NodeOpdLocal:       {'0', '1'},
		NodeOpdSecret:      {'0', '1'},
		NodeOpdHelp:        {'0', '1'},
	},
	NodeMust: {
		NodeDescription:  {'0', '1'},
		NodeErrorAppTag:  {'0', '1'},
		NodeErrorMessage: {'0', '1'},
		NodeReference:    {'0', '1'},
	},
	NodeLeaf: {
		NodeConfig:      {'0', '1'},
		NodeDefault:     {'0', '1'},
		NodeDescription: {'0', '1'},
		NodeIfFeature:   {'0', 'n'},
		NodeMandatory:   {'0', '1'},
		NodeMust:        {'0', 'n'},
		NodeReference:   {'0', '1'},
		NodeStatus:      {'0', '1'},
		NodeTyp:         {'1', '1'},
		NodeUnits:       {'0', '1'},
		NodeWhen:        {'0', '1'},
	},
	NodeOpdOption: {
		NodeDescription:    {'0', '1'},
		NodeIfFeature:      {'0', 'n'},
		NodeMandatory:      {'0', '1'},
		NodeMust:           {'0', 'n'},
		NodeReference:      {'0', '1'},
		NodeStatus:         {'0', '1'},
		NodeTyp:            {'1', '1'},
		NodeUnits:          {'0', '1'},
		NodeWhen:           {'0', '1'},
		NodeUses:           {'0', 'n'},
		NodeOpdCommand:     {'0', 'n'},
		NodeOpdOption:      {'0', 'n'},
		NodeOpdArgument:    {'0', '1'},
		NodeOpdOnEnter:     {'0', '1'},
		NodeOpdInherit:     {'0', '1'},
		NodeOpdRepeatable:  {'0', '1'},
		NodeOpdPassOpcArgs: {'0', '1'},
		NodeOpdPrivileged:  {'0', '1'},
		NodeOpdLocal:       {'0', '1'},
		NodeOpdSecret:      {'0', '1'},
		NodeOpdAllowed:     {'0', '1'},
		NodeOpdHelp:        {'0', '1'},
	},
	NodeOpdArgument: {
		NodeDescription:    {'0', '1'},
		NodeIfFeature:      {'0', 'n'},
		NodeMandatory:      {'0', '1'},
		NodeMust:           {'0', 'n'},
		NodeReference:      {'0', '1'},
		NodeStatus:         {'0', '1'},
		NodeTyp:            {'1', '1'},
		NodeUnits:          {'0', '1'},
		NodeWhen:           {'0', '1'},
		NodeUses:           {'0', 'n'},
		NodeOpdCommand:     {'0', 'n'},
		NodeOpdOption:      {'0', 'n'},
		NodeOpdArgument:    {'0', '1'},
		NodeOpdOnEnter:     {'0', '1'},
		NodeOpdInherit:     {'0', '1'},
		NodeOpdRepeatable:  {'0', '1'},
		NodeOpdPassOpcArgs: {'0', '1'},
		NodeOpdPrivileged:  {'0', '1'},
		NodeOpdLocal:       {'0', '1'},
		NodeOpdSecret:      {'0', '1'},
		NodeOpdAllowed:     {'0', '1'},
		NodeOpdHelp:        {'0', '1'},
	},
	NodeOpdInherit: {
		NodeOpdOnEnter:     {'0', '1'},
		NodeOpdPassOpcArgs: {'0', '1'},
		NodeOpdPrivileged:  {'0', '1'},
	},
	NodeOpdAllowed:     {},
	NodeOpdHelp:        {},
	NodeOpdPatternHelp: {},
	NodeLeafList: {
		NodeConfig:      {'0', '1'},
		NodeDescription: {'0', '1'},
		NodeIfFeature:   {'0', 'n'},
		NodeMaxElements: {'0', '1'},
		NodeMinElements: {'0', '1'},
		NodeMust:        {'0', 'n'},
		NodeOrderedBy:   {'0', '1'},
		NodeReference:   {'0', '1'},
		NodeStatus:      {'0', '1'},
		NodeTyp:         {'1', '1'},
		NodeUnits:       {'0', '1'},
		NodeWhen:        {'0', '1'},
	},
	NodeList: {
		NodeAnyxml:      {'0', 'n'},
		NodeChoice:      {'0', 'n'},
		NodeConfig:      {'0', '1'},
		NodeContainer:   {'0', 'n'},
		NodeDescription: {'0', '1'},
		NodeGrouping:    {'0', 'n'},
		NodeIfFeature:   {'0', 'n'},
		NodeKey:         {'1', '1'}, // 0..1 for config=false
		NodeLeaf:        {'0', 'n'},
		NodeLeafList:    {'0', 'n'},
		NodeList:        {'0', 'n'},
		NodeMaxElements: {'0', '1'},
		NodeMinElements: {'0', '1'},
		NodeMust:        {'0', 'n'},
		NodeOrderedBy:   {'0', '1'},
		NodeReference:   {'0', '1'},
		NodeStatus:      {'0', '1'},
		NodeTypedef:     {'0', 'n'},
		NodeUnique:      {'0', 'n'},
		NodeUses:        {'0', 'n'},
		NodeWhen:        {'0', '1'},
		NodeDataDef:     {'1', 'n'},
	},
	NodeChoice: {
		NodeAnyxml:      {'0', 'n'},
		NodeCase:        {'0', 'n'},
		NodeConfig:      {'0', '1'},
		NodeContainer:   {'0', 'n'},
		NodeDefault:     {'0', '1'},
		NodeDescription: {'0', '1'},
		NodeIfFeature:   {'0', 'n'},
		NodeLeaf:        {'0', 'n'},
		NodeLeafList:    {'0', 'n'},
		NodeList:        {'0', 'n'},
		NodeMandatory:   {'0', '1'},
		NodeReference:   {'0', '1'},
		NodeStatus:      {'0', '1'},
		NodeWhen:        {'0', '1'},
	},
	NodeCase: {
		NodeAnyxml:      {'0', 'n'},
		NodeChoice:      {'0', 'n'},
		NodeContainer:   {'0', 'n'},
		NodeDescription: {'0', '1'},
		NodeIfFeature:   {'0', 'n'},
		NodeLeaf:        {'0', 'n'},
		NodeLeafList:    {'0', 'n'},
		NodeList:        {'0', 'n'},
		NodeReference:   {'0', '1'},
		NodeStatus:      {'0', '1'},
		NodeUses:        {'0', 'n'},
		NodeWhen:        {'0', '1'},
	},
	NodeAnyxml: {
		NodeConfig:      {'0', '1'},
		NodeDescription: {'0', '1'},
		NodeIfFeature:   {'0', 'n'},
		NodeMandatory:   {'0', '1'},
		NodeMust:        {'0', 'n'},
		NodeReference:   {'0', '1'},
		NodeStatus:      {'0', '1'},
		NodeWhen:        {'0', '1'},
	},
	NodeGrouping: {
		NodeAnyxml:      {'0', 'n'},
		NodeChoice:      {'0', 'n'},
		NodeContainer:   {'0', 'n'},
		NodeDescription: {'0', '1'},
		NodeGrouping:    {'0', 'n'},
		NodeLeaf:        {'0', 'n'},
		NodeLeafList:    {'0', 'n'},
		NodeList:        {'0', 'n'},
		NodeReference:   {'0', '1'},
		NodeStatus:      {'0', '1'},
		NodeTypedef:     {'0', 'n'},
		NodeUses:        {'0', 'n'},
		NodeOpdCommand:  {'0', 'n'},
		NodeOpdOption:   {'0', 'n'},
		NodeOpdArgument: {'0', 'n'},
	},
	NodeUses: {
		NodeAugment:     {'0', 'n'},
		NodeDescription: {'0', '1'},
		NodeIfFeature:   {'0', 'n'},
		NodeRefine:      {'0', 'n'},
		NodeReference:   {'0', '1'},
		NodeStatus:      {'0', '1'},
		NodeWhen:        {'0', '1'},
	},
	NodeRpc: {
		NodeDescription: {'0', '1'},
		NodeGrouping:    {'0', 'n'},
		NodeIfFeature:   {'0', 'n'},
		NodeInput:       {'0', '1'},
		NodeOutput:      {'0', '1'},
		NodeReference:   {'0', '1'},
		NodeStatus:      {'0', '1'},
		NodeTypedef:     {'0', 'n'},
	},
	NodeInput: {
		NodeAnyxml:    {'0', 'n'},
		NodeChoice:    {'0', 'n'},
		NodeContainer: {'0', 'n'},
		NodeGrouping:  {'0', 'n'},
		NodeLeaf:      {'0', 'n'},
		NodeLeafList:  {'0', 'n'},
		NodeList:      {'0', 'n'},
		NodeTypedef:   {'0', 'n'},
		NodeUses:      {'0', 'n'},
		NodeDataDef:   {'0', 'n'},
	},
	NodeOutput: {
		NodeAnyxml:    {'0', 'n'},
		NodeChoice:    {'0', 'n'},
		NodeContainer: {'0', 'n'},
		NodeGrouping:  {'0', 'n'},
		NodeLeaf:      {'0', 'n'},
		NodeLeafList:  {'0', 'n'},
		NodeList:      {'0', 'n'},
		NodeTypedef:   {'0', 'n'},
		NodeUses:      {'0', 'n'},
		NodeDataDef:   {'0', 'n'},
	},
	NodeNotification: {
		NodeAnyxml:      {'0', 'n'},
		NodeChoice:      {'0', 'n'},
		NodeContainer:   {'0', 'n'},
		NodeDescription: {'0', '1'},
		NodeGrouping:    {'0', 'n'},
		NodeIfFeature:   {'0', 'n'},
		NodeLeaf:        {'0', 'n'},
		NodeLeafList:    {'0', 'n'},
		NodeList:        {'0', 'n'},
		NodeReference:   {'0', '1'},
		NodeStatus:      {'0', '1'},
		NodeTypedef:     {'0', 'n'},
		NodeUses:        {'0', 'n'},
	},
	NodeAugment: {
		NodeAnyxml:      {'0', 'n'},
		NodeCase:        {'0', 'n'},
		NodeChoice:      {'0', 'n'},
		NodeContainer:   {'0', 'n'},
		NodeDescription: {'0', '1'},
		NodeIfFeature:   {'0', 'n'},
		NodeLeaf:        {'0', 'n'},
		NodeLeafList:    {'0', 'n'},
		NodeList:        {'0', 'n'},
		NodeReference:   {'0', '1'},
		NodeStatus:      {'0', '1'},
		NodeUses:        {'0', 'n'},
		NodeWhen:        {'0', '1'},
	},
	NodeOpdAugment: {
		NodeAnyxml:      {'0', 'n'},
		NodeDescription: {'0', '1'},
		NodeIfFeature:   {'0', 'n'},
		NodeReference:   {'0', '1'},
		NodeStatus:      {'0', '1'},
		NodeUses:        {'0', 'n'},
		NodeWhen:        {'0', '1'},
		NodeOpdCommand:  {'0', 'n'},
		NodeOpdOption:   {'0', 'n'},
		NodeOpdArgument: {'0', '1'},
	},
	NodeIdentity: {
		NodeBase:        {'0', '1'},
		NodeDescription: {'0', '1'},
		NodeReference:   {'0', '1'},
		NodeStatus:      {'0', '1'},
	},
	NodeExtension: {
		NodeArgument:    {'0', '1'},
		NodeDescription: {'0', '1'},
		NodeReference:   {'0', '1'},
		NodeStatus:      {'0', '1'},
	},
	NodeArgument: {
		NodeYinElement: {'0', '1'},
	},
	NodeFeature: {
		NodeDescription: {'0', '1'},
		NodeIfFeature:   {'0', 'n'},
		NodeStatus:      {'0', '1'},
		NodeReference:   {'0', '1'},
	},
	NodeDeviation: {
		NodeDescription:         {'0', '1'},
		NodeDeviate:             {'0', 'n'},
		NodeDeviateAdd:          {'0', 'n'},
		NodeDeviateDelete:       {'0', 'n'},
		NodeDeviateNotSupported: {'0', 'n'},
		NodeDeviateReplace:      {'0', 'n'},
		NodeReference:           {'0', '1'},
	},
	NodeDeviate: {
		NodeConfig:      {'0', '1'},
		NodeDefault:     {'0', '1'},
		NodeMandatory:   {'0', '1'},
		NodeMaxElements: {'0', '1'},
		NodeMinElements: {'0', '1'},
		NodeMust:        {'0', 'n'},
		NodeTyp:         {'0', '1'},
		NodeUnique:      {'0', 'n'},
		NodeUnits:       {'0', '1'},
	},
	NodeDeviateNotSupported: {},
	NodeDeviateAdd: {
		NodeUnits:       {'0', '1'},
		NodeMust:        {'0', 'n'},
		NodeUnique:      {'0', 'n'},
		NodeDefault:     {'0', '1'},
		NodeConfig:      {'0', '1'},
		NodeMandatory:   {'0', '1'},
		NodeMinElements: {'0', '1'},
		NodeMaxElements: {'0', '1'},
	},
	NodeDeviateDelete: {
		NodeUnits:   {'0', '1'},
		NodeMust:    {'0', 'n'},
		NodeUnique:  {'0', 'n'},
		NodeDefault: {'0', '1'},
	},
	NodeDeviateReplace: {
		NodeTyp:         {'0', '1'},
		NodeUnits:       {'0', '1'},
		NodeDefault:     {'0', '1'},
		NodeConfig:      {'0', '1'},
		NodeMandatory:   {'0', '1'},
		NodeMinElements: {'0', '1'},
		NodeMaxElements: {'0', '1'},
	},
	NodeRange: {
		NodeDescription:  {'0', '1'},
		NodeErrorAppTag:  {'0', '1'},
		NodeErrorMessage: {'0', '1'},
		NodeReference:    {'0', '1'},
	},
	NodeLength: {
		NodeDescription:  {'0', '1'},
		NodeErrorAppTag:  {'0', '1'},
		NodeErrorMessage: {'0', '1'},
		NodeReference:    {'0', '1'},
	},
	NodePattern: {
		NodeDescription:  {'0', '1'},
		NodeErrorAppTag:  {'0', '1'},
		NodeErrorMessage: {'0', '1'},
		NodeReference:    {'0', '1'},
	},
	NodeEnum: {
		NodeDescription: {'0', '1'},
		NodeReference:   {'0', '1'},
		NodeStatus:      {'0', '1'},
		NodeValue:       {'0', '1'},
	},
	NodeBit: {
		NodeDescription: {'0', '1'},
		NodeReference:   {'0', '1'},
		NodeStatus:      {'0', '1'},
		NodePosition:    {'0', '1'},
	},
	NodeWhen: {
		NodeDescription: {'0', '1'},
		NodeReference:   {'0', '1'},
	},
}

var ErrCard = errors.New("cardinality mismatch")
