package cli_test

import (
	gocontext "context"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"

	sdk "github.com/cosmos/cosmos-sdk/types"

	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"

	"github.com/merlins-labs/merlin/v16/app/apptesting"
	"github.com/merlins-labs/merlin/v16/x/superfluid/types"
)

type QueryTestSuite struct {
	apptesting.KeeperTestHelper
	queryClient types.QueryClient
	val         sdk.ValAddress
}

func (s *QueryTestSuite) SetupSuite() {
	s.Setup()
	s.queryClient = types.NewQueryClient(s.QueryHelper)

	// set up durations
	s.App.IncentivesKeeper.SetLockableDurations(s.Ctx, []time.Duration{
		time.Hour * 24 * 14,
		time.Hour,
		time.Hour * 3,
		time.Hour * 7,
		time.Hour * 24 * 21,
	})
	// set up pool
	s.SetupGammPoolsWithBondDenomMultiplier([]sdk.Dec{sdk.NewDec(20), sdk.NewDec(20)})
	// set up lock with id = 1
	s.LockTokens(s.TestAccs[0], sdk.Coins{sdk.NewCoin("gamm/pool/1", sdk.NewInt(1000000))}, time.Hour*24*21)
	// set up validator
	s.val = s.SetupValidator(stakingtypes.Bonded)
	// set up sfs asset
	err := s.App.SuperfluidKeeper.AddNewSuperfluidAsset(s.Ctx, types.SuperfluidAsset{
		Denom:     "gamm/pool/1",
		AssetType: types.SuperfluidAssetTypeLPShare,
	})
	s.Require().NoError(err)
	// set up sfs delegation
	err = s.App.SuperfluidKeeper.SuperfluidDelegate(s.Ctx, s.TestAccs[0].String(), 1, s.val.String())
	s.Require().NoError(err)

	s.Commit()
}

func (s *QueryTestSuite) TestQueriesNeverAlterState() {
	s.SetupSuite()
	testCases := []struct {
		name   string
		query  string
		input  interface{}
		output interface{}
	}{
		{
			"Query all superfluild assets",
			"/merlin.superfluid.Query/AllAssets",
			&types.AllAssetsRequest{},
			&types.AllAssetsResponse{},
		},
		{
			"Query all intermediary accounts",
			"/merlin.superfluid.Query/AllIntermediaryAccounts",
			&types.AllIntermediaryAccountsRequest{},
			&types.AllIntermediaryAccountsResponse{},
		},
		{
			"Query fury equivalent multiplier of an asset",
			"/merlin.superfluid.Query/AssetMultiplier",
			&types.AssetMultiplierRequest{Denom: "gamm/pool/1"},
			&types.AssetMultiplierResponse{},
		},
		{
			"Query asset type",
			"/merlin.superfluid.Query/AssetType",
			&types.AssetTypeRequest{Denom: "gamm/pool/1"},
			&types.AssetTypeResponse{},
		},
		{
			"Query connected intermediary account",
			"/merlin.superfluid.Query/ConnectedIntermediaryAccount",
			&types.ConnectedIntermediaryAccountRequest{LockId: 1},
			&types.ConnectedIntermediaryAccountResponse{},
		},
		// need to adapt s.val.String() to have an intermediate account,
		// else the response is nil and theres a panic internally.
		// {
		// 	"Query estimate sfs delegated amount by validator & denom",
		// 	"/merlin.superfluid.Query/EstimateSuperfluidDelegatedAmountByValidatorDenom",
		// 	&types.EstimateSuperfluidDelegatedAmountByValidatorDenomRequest{ValidatorAddress: s.val.String(), Denom: "gamm/pool/1"},
		// 	&types.EstimateSuperfluidDelegatedAmountByValidatorDenomResponse{},
		// },
		{
			"Query params",
			"/merlin.superfluid.Query/Params",
			&types.QueryParamsRequest{},
			&types.QueryParamsResponse{},
		},
		{
			"Query sfs delegation amount",
			"/merlin.superfluid.Query/SuperfluidDelegationAmount",
			&types.SuperfluidDelegationAmountRequest{ValidatorAddress: s.val.String(), Denom: "gamm/pool/1", DelegatorAddress: s.TestAccs[0].String()},
			&types.SuperfluidDelegationAmountResponse{},
		},
		{
			"Query sfs delegation by delegator",
			"/merlin.superfluid.Query/SuperfluidDelegationsByDelegator",
			&types.SuperfluidDelegationsByDelegatorRequest{DelegatorAddress: s.TestAccs[0].String()},
			&types.SuperfluidDelegationsByDelegatorResponse{},
		},
		{
			"Query sfs delegation by validator & denom",
			"/merlin.superfluid.Query/SuperfluidDelegationsByValidatorDenom",
			&types.SuperfluidDelegationsByValidatorDenomRequest{ValidatorAddress: s.val.String(), Denom: "gamm/pool/1"},
			&types.SuperfluidDelegationsByValidatorDenomResponse{},
		},
		{
			"Query sfs undelegation by delegator",
			"/merlin.superfluid.Query/SuperfluidUndelegationsByDelegator",
			&types.SuperfluidUndelegationsByDelegatorRequest{DelegatorAddress: s.TestAccs[0].String(), Denom: "gamm/pool/1"},
			&types.SuperfluidUndelegationsByDelegatorResponse{},
		},
		{
			"Query total sfs delegation by delegator",
			"/merlin.superfluid.Query/TotalDelegationByDelegator",
			&types.QueryTotalDelegationByDelegatorRequest{DelegatorAddress: s.TestAccs[0].String()},
			&types.QueryTotalDelegationByDelegatorResponse{},
		},
		{
			"Query total sfs delegations",
			"/merlin.superfluid.Query/TotalSuperfluidDelegations",
			&types.TotalSuperfluidDelegationsRequest{},
			&types.TotalSuperfluidDelegationsResponse{},
		},
	}

	for _, tc := range testCases {
		tc := tc

		s.Run(tc.name, func() {
			s.SetupSuite()
			err := s.QueryHelper.Invoke(gocontext.Background(), tc.query, tc.input, tc.output)
			s.Require().NoError(err)
			s.StateNotAltered()
		})
	}
}

func TestQueryTestSuite(t *testing.T) {
	suite.Run(t, new(QueryTestSuite))
}
