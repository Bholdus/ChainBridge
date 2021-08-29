// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package utils

// An available method on the substrate chain
type Method string

var AddRelayerMethod Method = BridgePalletName + ".add_relayer"
var SetResourceMethod Method = BridgePalletName + ".set_resource"
var SetThresholdMethod Method = BridgePalletName + ".set_threshold"
var WhitelistChainMethod Method = BridgePalletName + ".whitelist_chain"
var ChainBridgeTransferToBridgeMethod Method = "ChainBridgeTransfer.transfer_to_bridge"
var ChainBridgeTransferNativeToBridgeMethod Method = "ChainBridgeTransfer.transfer_native_to_bridge"
var ChainBridgeRegisterResourceIdMethod Method = "ChainBridgeTransfer.register_resource_id"
var ChainBridgeTransferFromBridgeMethod Method = "ChainBridgeTransfer.transfer_from_bridge"
var Erc721MintMethod Method = "Erc721.mint"
var SudoMethod Method = "Sudo.sudo"
