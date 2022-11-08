package keeper_test

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"

	"github.com/evmos/ethermint/tests"
	"github.com/kynnoenterprise/core/x/rewardshare/types"
)

func (suite *KeeperTestSuite) TestRewardshares() {
	var (
		req    *types.QueryRewardsharesRequest
		expRes *types.QueryRewardsharesResponse
	)

	testCases := []struct {
		name     string
		malleate func()
		expPass  bool
	}{
		{
			"no fee infos registered",
			func() {
				req = &types.QueryRewardsharesRequest{}
				expRes = &types.QueryRewardsharesResponse{Pagination: &query.PageResponse{}}
			},
			true,
		},
		{
			"1 fee infos registered w/pagination",
			func() {
				req = &types.QueryRewardsharesRequest{
					Pagination: &query.PageRequest{Limit: 10, CountTotal: true},
				}
				rewardshare := types.NewRewardshare(contract, deployer, withdraw)
				suite.app.RewardshareKeeper.SetRewardshare(suite.ctx, rewardshare)

				expRes = &types.QueryRewardsharesResponse{
					Pagination: &query.PageResponse{Total: 1},
					Rewardshares: []types.Rewardshare{
						{
							ContractAddress:   contract.Hex(),
							DeployerAddress:   deployer.String(),
							WithdrawerAddress: withdraw.String(),
						},
					},
				}
			},
			true,
		},
		{
			"2 fee infos registered wo/pagination",
			func() {
				req = &types.QueryRewardsharesRequest{}
				contract2 := tests.GenerateAddress()
				rewardshare := types.NewRewardshare(contract, deployer, withdraw)
				feeSplit2 := types.NewRewardshare(contract2, deployer, nil)
				suite.app.RewardshareKeeper.SetRewardshare(suite.ctx, rewardshare)
				suite.app.RewardshareKeeper.SetRewardshare(suite.ctx, feeSplit2)

				expRes = &types.QueryRewardsharesResponse{
					Pagination: &query.PageResponse{Total: 2},
					Rewardshares: []types.Rewardshare{
						{
							ContractAddress:   contract.Hex(),
							DeployerAddress:   deployer.String(),
							WithdrawerAddress: withdraw.String(),
						},
						{
							ContractAddress: contract2.Hex(),
							DeployerAddress: deployer.String(),
						},
					},
				}
			},
			true,
		},
	}
	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("Case %s", tc.name), func() {
			suite.SetupTest() // reset

			ctx := sdk.WrapSDKContext(suite.ctx)
			tc.malleate()

			res, err := suite.queryClient.Rewardshares(ctx, req)
			if tc.expPass {
				suite.Require().NoError(err)
				suite.Require().Equal(expRes.Pagination, res.Pagination)
				suite.Require().ElementsMatch(expRes.Rewardshares, res.Rewardshares)
			} else {
				suite.Require().Error(err)
			}
		})
	}
}

// Cases that cannot be tested in TestFees
func (suite *KeeperTestSuite) TestRewardshareKeeper() {
	suite.SetupTest()
	ctx := sdk.WrapSDKContext(suite.ctx)
	res, err := suite.app.RewardshareKeeper.Rewardshares(ctx, nil)
	suite.Require().Error(err)
	suite.Require().Nil(res)
}

func (suite *KeeperTestSuite) TestFee() {
	var (
		req    *types.QueryRewardshareRequest
		expRes *types.QueryRewardshareResponse
	)

	testCases := []struct {
		name     string
		malleate func()
		expPass  bool
	}{
		{
			"empty contract address",
			func() {
				req = &types.QueryRewardshareRequest{}
				expRes = &types.QueryRewardshareResponse{}
			},
			false,
		},
		{
			"invalid contract address",
			func() {
				req = &types.QueryRewardshareRequest{
					ContractAddress: "1234",
				}
				expRes = &types.QueryRewardshareResponse{}
			},
			false,
		},
		{
			"fee info not found",
			func() {
				req = &types.QueryRewardshareRequest{
					ContractAddress: contract.String(),
				}
				expRes = &types.QueryRewardshareResponse{}
			},
			false,
		},
		{
			"fee info found",
			func() {
				rewardshare := types.NewRewardshare(contract, deployer, withdraw)
				suite.app.RewardshareKeeper.SetRewardshare(suite.ctx, rewardshare)

				req = &types.QueryRewardshareRequest{
					ContractAddress: contract.Hex(),
				}
				expRes = &types.QueryRewardshareResponse{Rewardshare: types.Rewardshare{
					ContractAddress:   contract.Hex(),
					DeployerAddress:   deployer.String(),
					WithdrawerAddress: withdraw.String(),
				}}
			},
			true,
		},
	}
	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("Case %s", tc.name), func() {
			suite.SetupTest() // reset

			ctx := sdk.WrapSDKContext(suite.ctx)
			tc.malleate()

			res, err := suite.queryClient.Rewardshare(ctx, req)
			if tc.expPass {
				suite.Require().NoError(err)
				suite.Require().Equal(expRes, res)
			} else {
				suite.Require().Error(err)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestDeployerFees() {
	var (
		req    *types.QueryDeployerRewardsharesRequest
		expRes *types.QueryDeployerRewardsharesResponse
	)

	testCases := []struct {
		name     string
		malleate func()
		expPass  bool
	}{
		{
			"no contract registered",
			func() {
				req = &types.QueryDeployerRewardsharesRequest{}
				expRes = &types.QueryDeployerRewardsharesResponse{Pagination: &query.PageResponse{}}
			},
			false,
		},
		{
			"invalid deployer address",
			func() {
				req = &types.QueryDeployerRewardsharesRequest{
					DeployerAddress: "123",
				}
				expRes = &types.QueryDeployerRewardsharesResponse{Pagination: &query.PageResponse{}}
			},
			false,
		},
		{
			"1 fee registered w/pagination",
			func() {
				req = &types.QueryDeployerRewardsharesRequest{
					Pagination:      &query.PageRequest{Limit: 10, CountTotal: true},
					DeployerAddress: deployer.String(),
				}

				rewardshare := types.NewRewardshare(contract, deployer, withdraw)
				suite.app.RewardshareKeeper.SetRewardshare(suite.ctx, rewardshare)
				suite.app.RewardshareKeeper.SetDeployerMap(suite.ctx, deployer, contract)
				suite.app.RewardshareKeeper.SetWithdrawerMap(suite.ctx, withdraw, contract)

				expRes = &types.QueryDeployerRewardsharesResponse{
					Pagination: &query.PageResponse{Total: 1},
					ContractAddresses: []string{
						contract.Hex(),
					},
				}
			},
			true,
		},
		{
			"2 fee infos registered for one contract wo/pagination",
			func() {
				req = &types.QueryDeployerRewardsharesRequest{
					DeployerAddress: deployer.String(),
				}
				contract2 := tests.GenerateAddress()
				rewardshare := types.NewRewardshare(contract, deployer, withdraw)
				suite.app.RewardshareKeeper.SetRewardshare(suite.ctx, rewardshare)
				suite.app.RewardshareKeeper.SetDeployerMap(suite.ctx, deployer, contract)
				suite.app.RewardshareKeeper.SetWithdrawerMap(suite.ctx, withdraw, contract)

				feeSplit2 := types.NewRewardshare(contract2, deployer, nil)
				suite.app.RewardshareKeeper.SetRewardshare(suite.ctx, feeSplit2)
				suite.app.RewardshareKeeper.SetDeployerMap(suite.ctx, deployer, contract2)

				expRes = &types.QueryDeployerRewardsharesResponse{
					Pagination: &query.PageResponse{Total: 2},
					ContractAddresses: []string{
						contract.Hex(),
						contract2.Hex(),
					},
				}
			},
			true,
		},
	}
	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("Case %s", tc.name), func() {
			suite.SetupTest() // reset

			ctx := sdk.WrapSDKContext(suite.ctx)
			tc.malleate()

			res, err := suite.queryClient.DeployerRewardshares(ctx, req)
			if tc.expPass {
				suite.Require().NoError(err)
				suite.Require().Equal(expRes.Pagination, res.Pagination)
				suite.Require().ElementsMatch(expRes.ContractAddresses, res.ContractAddresses)
			} else {
				suite.Require().Error(err)
			}
		})
	}
}

// Cases that cannot be tested in TestDeployerFees
func (suite *KeeperTestSuite) TestDeployerRewardshareKeeper() {
	suite.SetupTest()
	ctx := sdk.WrapSDKContext(suite.ctx)
	res, err := suite.app.RewardshareKeeper.DeployerRewardshares(ctx, nil)
	suite.Require().Error(err)
	suite.Require().Nil(res)
}

func (suite *KeeperTestSuite) TestWithdrawerRewardshares() {
	var (
		req    *types.QueryWithdrawerRewardsharesRequest
		expRes *types.QueryWithdrawerRewardsharesResponse
	)

	testCases := []struct {
		name     string
		malleate func()
		expPass  bool
	}{
		{
			"no contract registered",
			func() {
				req = &types.QueryWithdrawerRewardsharesRequest{}
				expRes = &types.QueryWithdrawerRewardsharesResponse{Pagination: &query.PageResponse{}}
			},
			false,
		},
		{
			"invalid withdraw address",
			func() {
				req = &types.QueryWithdrawerRewardsharesRequest{
					WithdrawerAddress: "123",
				}
				expRes = &types.QueryWithdrawerRewardsharesResponse{Pagination: &query.PageResponse{}}
			},
			false,
		},
		{
			"1 fee registered w/pagination",
			func() {
				req = &types.QueryWithdrawerRewardsharesRequest{
					Pagination:        &query.PageRequest{Limit: 10, CountTotal: true},
					WithdrawerAddress: withdraw.String(),
				}

				rewardshare := types.NewRewardshare(contract, deployer, withdraw)
				suite.app.RewardshareKeeper.SetRewardshare(suite.ctx, rewardshare)
				suite.app.RewardshareKeeper.SetDeployerMap(suite.ctx, deployer, contract)
				suite.app.RewardshareKeeper.SetWithdrawerMap(suite.ctx, withdraw, contract)

				expRes = &types.QueryWithdrawerRewardsharesResponse{
					Pagination: &query.PageResponse{Total: 1},
					ContractAddresses: []string{
						contract.Hex(),
					},
				}
			},
			true,
		},
		{
			"2 fees registered for one withdraw address wo/pagination",
			func() {
				req = &types.QueryWithdrawerRewardsharesRequest{
					WithdrawerAddress: withdraw.String(),
				}
				contract2 := tests.GenerateAddress()
				deployer2 := sdk.AccAddress(tests.GenerateAddress().Bytes())

				rewardshare := types.NewRewardshare(contract, deployer, withdraw)
				suite.app.RewardshareKeeper.SetRewardshare(suite.ctx, rewardshare)
				suite.app.RewardshareKeeper.SetDeployerMap(suite.ctx, deployer, contract)
				suite.app.RewardshareKeeper.SetWithdrawerMap(suite.ctx, withdraw, contract)

				feeSplit2 := types.NewRewardshare(contract2, deployer2, withdraw)
				suite.app.RewardshareKeeper.SetRewardshare(suite.ctx, feeSplit2)
				suite.app.RewardshareKeeper.SetDeployerMap(suite.ctx, deployer2, contract2)
				suite.app.RewardshareKeeper.SetWithdrawerMap(suite.ctx, withdraw, contract2)

				expRes = &types.QueryWithdrawerRewardsharesResponse{
					Pagination: &query.PageResponse{Total: 2},
					ContractAddresses: []string{
						contract.Hex(),
						contract2.Hex(),
					},
				}
			},
			true,
		},
	}
	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("Case %s", tc.name), func() {
			suite.SetupTest() // reset

			ctx := sdk.WrapSDKContext(suite.ctx)
			tc.malleate()

			res, err := suite.queryClient.WithdrawerRewardshares(ctx, req)
			if tc.expPass {
				suite.Require().NoError(err)
				suite.Require().Equal(expRes.Pagination, res.Pagination)
				suite.Require().ElementsMatch(expRes.ContractAddresses, res.ContractAddresses)
			} else {
				suite.Require().Error(err)
			}
		})
	}
}

// Cases that cannot be tested in TestWithdrawerFees
func (suite *KeeperTestSuite) TestWithdrawerRewardshareKeeper() {
	suite.SetupTest()
	ctx := sdk.WrapSDKContext(suite.ctx)
	res, err := suite.app.RewardshareKeeper.WithdrawerRewardshares(ctx, nil)
	suite.Require().Error(err)
	suite.Require().Nil(res)
}

func (suite *KeeperTestSuite) TestQueryParams() {
	ctx := sdk.WrapSDKContext(suite.ctx)
	expParams := types.DefaultParams()
	expParams.EnableRewardshare = true

	res, err := suite.queryClient.Params(ctx, &types.QueryParamsRequest{})
	suite.Require().NoError(err)
	suite.Require().Equal(expParams, res.Params)
}
