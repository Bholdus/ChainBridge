// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package utils

import (
	"github.com/centrifuge/go-substrate-rpc-client/v3/scale"
	"github.com/centrifuge/go-substrate-rpc-client/v3/types"
)

const BridgePalletName = "ChainBridge"
const BridgeStoragePrefix = "ChainBridge"

type Erc721Token struct {
	Id       types.U256
	Metadata types.Bytes
}

type RegistryId types.H160
type TokenId types.U256

type AssetId struct {
	RegistryId RegistryId
	TokenId    TokenId
}

type CurrencyId struct {
	IsToken    bool
	AsToken    TokenSymbol
	IsDexShare bool
	AsDexShare struct {
		DexShare1 DexShare
		DexShare2 DexShare
	}
}

func (m *CurrencyId) Decode(decoder scale.Decoder) error {
	b, err := decoder.ReadOneByte()

	if err != nil {
		return err
	}

	if b == 0 {
		m.IsToken = true
		err = decoder.Decode(&m.AsToken)
	} else if b == 1 {
		m.IsDexShare = true
		err = decoder.Decode(&m.AsDexShare)
	}

	if err != nil {
		return err
	}

	return nil
}

func (m CurrencyId) Encode(encoder scale.Encoder) error {
	var err1, err2 error
	if m.IsToken {
		err1 = encoder.PushByte(0)
		err2 = encoder.Encode(m.AsToken)
	} else if m.IsDexShare {
		err1 = encoder.PushByte(1)
		err2 = encoder.Encode(m.AsDexShare)
	}

	if err1 != nil {
		return err1
	}
	if err2 != nil {
		return err2
	}

	return nil
}

type DexShare struct {
	IsToken bool
	AsToken TokenSymbol
}

func (m *DexShare) Decode(decoder scale.Decoder) error {
	b, err := decoder.ReadOneByte()

	if err != nil {
		return err
	}

	if b == 0 {
		m.IsToken = true
		err = decoder.Decode(&m.AsToken)
	}

	if err != nil {
		return err
	}

	return nil
}

func (m DexShare) Encode(encoder scale.Encoder) error {
	var err1, err2 error
	if m.IsToken {
		err1 = encoder.PushByte(0)
		err2 = encoder.Encode(m.AsToken)
	}

	if err1 != nil {
		return err1
	}
	if err2 != nil {
		return err2
	}

	return nil
}

type TokenInfo struct {
	ID types.U32
}

type TokenSymbol struct {
	IsNative bool
	IsToken  bool
	AsToken  TokenInfo
}

func (m *TokenSymbol) Decode(decoder scale.Decoder) error {
	b, err := decoder.ReadOneByte()

	if err != nil {
		return err
	}

	if b == 0 {
		m.IsNative = true
	} else if b == 1 {
		m.IsToken = true
		err = decoder.Decode(&m.AsToken)
	}

	if err != nil {
		return err
	}

	return nil
}

func (m TokenSymbol) Encode(encoder scale.Encoder) error {
	var err1, err2 error
	if m.IsNative {
		err1 = encoder.PushByte(0)
	} else if m.IsToken {
		err1 = encoder.PushByte(1)
		err2 = encoder.Encode(m.AsToken)
	}

	if err1 != nil {
		return err1
	}
	if err2 != nil {
		return err2
	}

	return nil
}

type TradingPair struct {
	CurrencyId0 CurrencyId
	CurrencyId1 CurrencyId
}
