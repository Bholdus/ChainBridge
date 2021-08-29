// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package utils

import (
	events "github.com/ChainSafe/chainbridge-substrate-events"
	"github.com/centrifuge/go-substrate-rpc-client/v3/types"
)

type EventErc721Minted struct {
	Phase   types.Phase
	Owner   types.AccountID
	TokenId types.U256
	Topics  []types.Hash
}

type EventErc721Transferred struct {
	Phase   types.Phase
	From    types.AccountID
	To      types.AccountID
	TokenId types.U256
	Topics  []types.Hash
}

type EventErc721Burned struct {
	Phase   types.Phase
	TokenId types.AccountID
	Topics  []types.Hash
}

type EventExampleRemark struct {
	Phase  types.Phase
	Hash   types.Hash
	Topics []types.Hash
}

// EventNFTDeposited is emitted when NFT is ready to be deposited to other chain.
type EventNFTDeposited struct {
	Phase  types.Phase
	Asset  types.Hash
	Topics []types.Hash
}

// EventFeeChanged is emitted when a fee for a given key is changed.
type EventFeeChanged struct {
	Phase    types.Phase
	Key      types.Hash
	NewPrice types.U128
	Topics   []types.Hash
}

// EventNewMultiAccount is emitted when a multi account has been created.
// First param is the account that created it, second is the multisig account.
type EventNewMultiAccount struct {
	Phase   types.Phase
	Who, ID types.AccountID
	Topics  []types.Hash
}

// EventMultiAccountUpdated is emitted when a multi account has been updated. First param is the multisig account.
type EventMultiAccountUpdated struct {
	Phase  types.Phase
	Who    types.AccountID
	Topics []types.Hash
}

// EventMultiAccountRemoved is emitted when a multi account has been removed. First param is the multisig account.
type EventMultiAccountRemoved struct {
	Phase  types.Phase
	Who    types.AccountID
	Topics []types.Hash
}

// EventNewMultisig is emitted when a new multisig operation has begun.
// First param is the account that is approving, second is the multisig account.
type EventNewMultisig struct {
	Phase   types.Phase
	Who, ID types.AccountID
	Topics  []types.Hash
}

// TimePoint contains height and index
type TimePoint struct {
	Height types.U32
	Index  types.U32
}

// EventMultisigApproval is emitted when a multisig operation has been approved by someone.
// First param is the account that is approving, third is the multisig account.
type EventMultisigApproval struct {
	Phase     types.Phase
	Who       types.AccountID
	TimePoint TimePoint
	ID        types.AccountID
	Topics    []types.Hash
}

// EventMultisigExecuted is emitted when a multisig operation has been executed by someone.
// First param is the account that is approving, third is the multisig account.
type EventMultisigExecuted struct {
	Phase     types.Phase
	Who       types.AccountID
	TimePoint TimePoint
	ID        types.AccountID
	Result    types.DispatchResult
	Topics    []types.Hash
}

// EventMultisigCancelled is emitted when a multisig operation has been cancelled by someone.
// First param is the account that is approving, third is the multisig account.
type EventMultisigCancelled struct {
	Phase     types.Phase
	Who       types.AccountID
	TimePoint TimePoint
	ID        types.AccountID
	Topics    []types.Hash
}

type EventTreasuryMinting struct {
	Phase  types.Phase
	Who    types.AccountID
	Topics []types.Hash
}

// EventRadClaimsClaimed is emitted when RAD Tokens have been claimed
type EventRadClaimsClaimed struct {
	Phase  types.Phase
	Who    types.AccountID
	Value  types.U128
	Topics []types.Hash
}

// EventRadClaimsRootHashStored is emitted when RootHash has been stored for the correspondent RAD Claims batch
type EventRadClaimsRootHashStored struct {
	Phase    types.Phase
	RootHash types.Hash
	Topics   []types.Hash
}

// EventNftTransferred is emitted when the ownership of the asset has been transferred to the account
type EventNftTransferred struct {
	Phase      types.Phase
	RegistryId RegistryId
	AssetId    AssetId
	Who        types.AccountID
	Topics     []types.Hash
}

// EventRegistryMint is emitted when successfully minting an NFT
type EventRegistryMint struct {
	Phase      types.Phase
	RegistryId RegistryId
	TokenId    TokenId
	Topics     []types.Hash
}

// EventRegistryRegistryCreated is emitted when successfully creating a NFT registry
type EventRegistryRegistryCreated struct {
	Phase      types.Phase
	RegistryId RegistryId
	Topics     []types.Hash
}

type OptionElectionCompute struct {
	types.OptionBytes
}

// EventRegistryTmp is emitted only for testing
type EventRegistryTmp struct {
	Phase  types.Phase
	Hash   types.Hash
	Topics []types.Hash
}

// Election Event
type EventElectionSignedPhaseStarted struct {
	Phase  types.Phase
	Index  types.U32
	Topics []types.Hash
}

type EventElectionUnsignedPhaseStarted struct {
	Phase  types.Phase
	Index  types.U32
	Topics []types.Hash
}

type EventElectionElectionFinalized struct {
	Phase   types.Phase
	Compute OptionElectionCompute
	Topics  []types.Hash
}

type EventStakingStakingElection struct {
	Phase  types.Phase
	Topics []types.Hash
}

type EventTokensIdentitySet struct {
	Phase   types.Phase
	AssetId CurrencyId
	Topics  []types.Hash
}
type EventTokensCreated struct {
	Phase   types.Phase
	AssetId CurrencyId
	Creator types.AccountID
	Owner   types.AccountID
	Topics  []types.Hash
}
type EventTokensIssued struct {
	Phase   types.Phase
	AssetId CurrencyId
	Owner   types.AccountID
	Balance types.U128
	Topics  []types.Hash
}
type EventTokensTransferred struct {
	Phase   types.Phase
	AssetId CurrencyId
	From    types.AccountID
	To      types.AccountID
	Balance types.U128
	Topics  []types.Hash
}
type EventTokensBurned struct {
	Phase   types.Phase
	AssetId CurrencyId
	Owner   types.AccountID
	Balance types.U128
	Topics  []types.Hash
}
type EventTokensFrozen struct {
	Phase   types.Phase
	AssetId CurrencyId
	Who     types.AccountID
	Topics  []types.Hash
}
type EventTokensThawed struct {
	Phase   types.Phase
	AssetId CurrencyId
	Who     types.AccountID
	Topics  []types.Hash
}
type EventTokensEndowed struct {
	Phase   types.Phase
	AssetId CurrencyId
	Who     types.AccountID
	Amount  types.U128
	Topics  []types.Hash
}
type EventTokensAssetFrozen struct {
	Phase   types.Phase
	AssetId CurrencyId
	Topics  []types.Hash
}
type EventTokensAssetThawed struct {
	Phase   types.Phase
	AssetId CurrencyId
	Topics  []types.Hash
}
type EventTokensAssetVerified struct {
	Phase   types.Phase
	AssetId CurrencyId
	Topics  []types.Hash
}
type EventTokensDestroyed struct {
	Phase   types.Phase
	AssetId CurrencyId
	Topics  []types.Hash
}
type EventTokensForceCreated struct {
	Phase   types.Phase
	AssetId CurrencyId
	Owner   types.AccountID
	Topics  []types.Hash
}
type EventTokensMetadataSet struct {
	Phase    types.Phase
	AssetId  CurrencyId
	Name     types.Bytes
	Symbol   types.Bytes
	Decimals types.U8
	IsFrozen bool
	Topics   []types.Hash
}
type EventTokensMetadataCleared struct {
	Phase   types.Phase
	AssetId CurrencyId
	Topics  []types.Hash
}
type EventTokensProfileSet struct {
	Phase      types.Phase
	AssetId    CurrencyId
	Name       types.Bytes
	IsVerified bool
	Topics     []types.Hash
}

type EventCurrenciesTransferred struct {
	Phase      types.Phase
	CurrencyId CurrencyId
	From       types.AccountID
	To         types.AccountID
	Amount     types.U128
	Topics     []types.Hash
}
type EventCurrenciesBalanceUpdated struct {
	Phase      types.Phase
	CurrencyId CurrencyId
	Who        types.AccountID
	Amount     types.I128
	Topics     []types.Hash
}
type EventCurrenciesDeposited struct {
	Phase      types.Phase
	CurrencyId CurrencyId
	Who        types.AccountID
	Amount     types.U128
	Topics     []types.Hash
}
type EventCurrenciesWithdrawn struct {
	Phase      types.Phase
	CurrencyId CurrencyId
	Who        types.AccountID
	Amount     types.U128
	Topics     []types.Hash
}

type EventDexSwap struct {
	Phase        types.Phase
	Who          types.AccountID
	TradingPath  []CurrencyId
	SupplyAmount types.U128
	TargetAmount types.U128
	Topics       []types.Hash
}
type EventDexTradingPairEnabledFromProvisioning struct {
	Phase       types.Phase
	CurrencyId0 CurrencyId
	Pool0       types.U128
	CurrencyId1 CurrencyId
	Pool1       types.U128
	Topics      []types.Hash
}
type EventDexTradingPairDisabled struct {
	Phase       types.Phase
	TradingPair TradingPair
	Topics      []types.Hash
}
type EventDexTradingPairEnabled struct {
	Phase       types.Phase
	TradingPair TradingPair
	Topics      []types.Hash
}
type EventDexTradingPairProvisioning struct {
	Phase       types.Phase
	TradingPair TradingPair
	Topics      []types.Hash
}
type EventDexAddProvision struct {
	Phase         types.Phase
	Who           types.AccountID
	Currency0     CurrencyId
	Contribution0 types.U128
	Currency1     CurrencyId
	Contribution1 types.U128
	Topics        []types.Hash
}
type EventDexAddLiquidity struct {
	Phase          types.Phase
	Who            types.AccountID
	CurrencyId0    CurrencyId
	PoolIncrement0 types.U128
	CurrencyId1    CurrencyId
	PoolIncrement1 types.U128
	ShareIncrement types.U128
	Topics         []types.Hash
}
type EventDexRemoveLiquidity struct {
	Phase          types.Phase
	Who            types.AccountID
	CurrencyId0    CurrencyId
	PoolDecrement0 types.U128
	CurrencyId1    CurrencyId
	PoolDecrement1 types.U128
	ShareDecrement types.U128
	Topics         []types.Hash
}

type EventChainBridgeTransferResourceIdRegistered struct {
	Phase      types.Phase
	ResourceId types.Bytes32
	CurrencyId CurrencyId
	Topics     []types.Hash
}
type EventChainBridgeTransferResourceIdUnregistered struct {
	Phase      types.Phase
	ResourceId types.Bytes32
	CurrencyId CurrencyId
	Topics     []types.Hash
}

type Events struct {
	types.EventRecords
	events.Events
	Erc721_Minted                    []EventErc721Minted                   //nolint:stylecheck,golint
	Erc721_Transferred               []EventErc721Transferred              //nolint:stylecheck,golint
	Erc721_Burned                    []EventErc721Burned                   //nolint:stylecheck,golint
	Example_Remark                   []EventExampleRemark                  //nolint:stylecheck,golint
	Nfts_DepositAsset                []EventNFTDeposited                   //nolint:stylecheck,golint
	Council_Proposed                 []types.EventCollectiveProposed       //nolint:stylecheck,golint
	Council_Voted                    []types.EventCollectiveVoted          //nolint:stylecheck,golint
	Council_Approved                 []types.EventCollectiveApproved       //nolint:stylecheck,golint
	Council_Disapproved              []types.EventCollectiveDisapproved    //nolint:stylecheck,golint
	Council_Executed                 []types.EventCollectiveExecuted       //nolint:stylecheck,golint
	Council_MemberExecuted           []types.EventCollectiveMemberExecuted //nolint:stylecheck,golint
	Council_Closed                   []types.EventCollectiveClosed         //nolint:stylecheck,golint
	Fees_FeeChanged                  []EventFeeChanged                     //nolint:stylecheck,golint
	MultiAccount_NewMultiAccount     []EventNewMultiAccount                //nolint:stylecheck,golint
	MultiAccount_MultiAccountUpdated []EventMultiAccountUpdated            //nolint:stylecheck,golint
	MultiAccount_MultiAccountRemoved []EventMultiAccountRemoved            //nolint:stylecheck,golint
	MultiAccount_NewMultisig         []EventNewMultisig                    //nolint:stylecheck,golint
	MultiAccount_MultisigApproval    []EventMultisigApproval               //nolint:stylecheck,golint
	MultiAccount_MultisigExecuted    []EventMultisigExecuted               //nolint:stylecheck,golint
	MultiAccount_MultisigCancelled   []EventMultisigCancelled              //nolint:stylecheck,golint
	TreasuryReward_TreasuryMinting   []EventTreasuryMinting                //nolint:stylecheck,golint
	Nft_Transferred                  []EventNftTransferred                 //nolint:stylecheck,golint
	RadClaims_Claimed                []EventRadClaimsClaimed               //nolint:stylecheck,golint
	RadClaims_RootHashStored         []EventRadClaimsRootHashStored        //nolint:stylecheck,golint
	Registry_Mint                    []EventRegistryMint                   //nolint:stylecheck,golint
	Registry_RegistryCreated         []EventRegistryRegistryCreated        //nolint:stylecheck,golint
	Registry_RegistryTmp             []EventRegistryTmp                    //nolint:stylecheck,golint

	ElectionProviderMultiPhase_SignedPhaseStarted   []EventElectionSignedPhaseStarted
	ElectionProviderMultiPhase_UnsignedPhaseStarted []EventElectionUnsignedPhaseStarted
	ElectionProviderMultiPhase_ElectionFinalized    []EventElectionElectionFinalized

	Staking_StakingElection []EventStakingStakingElection

	Tokens_IdentitySet     []EventTokensIdentitySet
	Tokens_Created         []EventTokensCreated
	Tokens_Issued          []EventTokensIssued
	Tokens_Transferred     []EventTokensTransferred
	Tokens_Burned          []EventTokensBurned
	Tokens_Frozen          []EventTokensFrozen
	Tokens_Thawed          []EventTokensThawed
	Tokens_Endowed         []EventTokensEndowed
	Tokens_AssetFrozen     []EventTokensFrozen
	Tokens_AssetThawed     []EventTokensAssetThawed
	Tokens_AssetVerified   []EventTokensAssetVerified
	Tokens_Destroyed       []EventTokensDestroyed
	Tokens_ForceCreated    []EventTokensForceCreated
	Tokens_MetadataSet     []EventTokensMetadataSet
	Tokens_MetadataCleared []EventTokensMetadataCleared
	Tokens_ProfileSet      []EventTokensProfileSet

	Currencies_Transferred    []EventCurrenciesTransferred
	Currencies_BalanceUpdated []EventCurrenciesBalanceUpdated
	Currencies_Deposited      []EventCurrenciesBalanceUpdated
	Currencies_Withdrawn      []EventCurrenciesWithdrawn

	Dex_Swap                               []EventDexSwap
	Dex_TradingPairEnabledFromProvisioning []EventDexTradingPairEnabledFromProvisioning
	Dex_TradingPairDisabled                []EventDexTradingPairDisabled
	Dex_TradingPairEnabled                 []EventDexTradingPairEnabled
	Dex_TradingPairProvisioning            []EventDexTradingPairProvisioning
	Dex_AddProvision                       []EventDexAddProvision
	Dex_AddLiquidity                       []EventDexAddLiquidity
	Dex_RemoveLiquidity                    []EventDexRemoveLiquidity

	ChainBridgeTransfer_ResourceIdRegistered   []EventChainBridgeTransferResourceIdRegistered
	ChainBridgeTransfer_ResourceIdUnregistered []EventChainBridgeTransferResourceIdUnregistered
}
