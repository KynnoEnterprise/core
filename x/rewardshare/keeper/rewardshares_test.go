package keeper_test

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/evmos/ethermint/tests"
	"github.com/kynnoenterprise/core/x/rewardshare/types"
)

func (suite *KeeperTestSuite) TestGetFees() {
	var expRes []types.Rewardshare

	testCases := []struct {
		name     string
		malleate func()
	}{
		{
			"no rewardshares registered",
			func() { expRes = []types.Rewardshare{} },
		},
		{
			"one rewardshare registered with withdraw address",
			func() {
				rewardshare := types.NewRewardshare(contract, deployer, withdraw)
				suite.app.RewardshareKeeper.SetRewardshare(suite.ctx, rewardshare)
				expRes = []types.Rewardshare{rewardshare}
			},
		},
		{
			"one rewardshare registered with no withdraw address",
			func() {
				rewardshare := types.NewRewardshare(contract, deployer, nil)
				suite.app.RewardshareKeeper.SetRewardshare(suite.ctx, rewardshare)
				expRes = []types.Rewardshare{rewardshare}
			},
		},
		{
			"multiple rewardshares registered",
			func() {
				deployer2 := sdk.AccAddress(tests.GenerateAddress().Bytes())
				contract2 := tests.GenerateAddress()
				contract3 := tests.GenerateAddress()
				rewardshare := types.NewRewardshare(contract, deployer, withdraw)
				feeSplit2 := types.NewRewardshare(contract2, deployer, nil)
				feeSplit3 := types.NewRewardshare(contract3, deployer2, nil)
				suite.app.RewardshareKeeper.SetRewardshare(suite.ctx, rewardshare)
				suite.app.RewardshareKeeper.SetRewardshare(suite.ctx, feeSplit2)
				suite.app.RewardshareKeeper.SetRewardshare(suite.ctx, feeSplit3)
				expRes = []types.Rewardshare{rewardshare, feeSplit2, feeSplit3}
			},
		},
	}
	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("Case %s", tc.name), func() {
			suite.SetupTest() // reset
			tc.malleate()

			res := suite.app.RewardshareKeeper.GetRewardshares(suite.ctx)
			suite.Require().ElementsMatch(expRes, res, tc.name)
		})
	}
}

func (suite *KeeperTestSuite) TestIterateFees() {
	var expRes []types.Rewardshare

	testCases := []struct {
		name     string
		malleate func()
	}{
		{
			"no rewardshares registered",
			func() { expRes = []types.Rewardshare{} },
		},
		{
			"one rewardshare registered with withdraw address",
			func() {
				rewardshare := types.NewRewardshare(contract, deployer, withdraw)
				suite.app.RewardshareKeeper.SetRewardshare(suite.ctx, rewardshare)
				expRes = []types.Rewardshare{
					types.NewRewardshare(contract, deployer, withdraw),
				}
			},
		},
		{
			"one rewardshare registered with no withdraw address",
			func() {
				rewardshare := types.NewRewardshare(contract, deployer, nil)
				suite.app.RewardshareKeeper.SetRewardshare(suite.ctx, rewardshare)
				expRes = []types.Rewardshare{
					types.NewRewardshare(contract, deployer, nil),
				}
			},
		},
		{
			"multiple rewardshares registered",
			func() {
				deployer2 := sdk.AccAddress(tests.GenerateAddress().Bytes())
				contract2 := tests.GenerateAddress()
				contract3 := tests.GenerateAddress()
				rewardshare := types.NewRewardshare(contract, deployer, withdraw)
				feeSplit2 := types.NewRewardshare(contract2, deployer, nil)
				feeSplit3 := types.NewRewardshare(contract3, deployer2, nil)
				suite.app.RewardshareKeeper.SetRewardshare(suite.ctx, rewardshare)
				suite.app.RewardshareKeeper.SetRewardshare(suite.ctx, feeSplit2)
				suite.app.RewardshareKeeper.SetRewardshare(suite.ctx, feeSplit3)
				expRes = []types.Rewardshare{rewardshare, feeSplit2, feeSplit3}
			},
		},
	}
	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("Case %s", tc.name), func() {
			suite.SetupTest() // reset
			tc.malleate()

			suite.app.RewardshareKeeper.IterateRewardshares(suite.ctx, func(rewardshare types.Rewardshare) (stop bool) {
				suite.Require().Contains(expRes, rewardshare, tc.name)
				return false
			})
		})
	}
}

func (suite *KeeperTestSuite) TestGetRewardshare() {
	testCases := []struct {
		name        string
		contract    common.Address
		deployer    sdk.AccAddress
		withdraw    sdk.AccAddress
		found       bool
		expWithdraw bool
	}{
		{
			"fee with no withdraw address",
			contract,
			deployer,
			nil,
			true,
			false,
		},
		{
			"fee with withdraw address same as deployer",
			contract,
			deployer,
			deployer,
			true,
			false,
		},
		{
			"fee with withdraw address same as contract",
			contract,
			deployer,
			sdk.AccAddress(contract.Bytes()),
			true,
			true,
		},
		{
			"fee with withdraw address different than deployer",
			contract,
			deployer,
			withdraw,
			true,
			true,
		},
		{
			"no fee",
			common.Address{},
			nil,
			nil,
			false,
			false,
		},
	}
	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("Case %s", tc.name), func() {
			suite.SetupTest() // reset

			if tc.found {
				rewardshare := types.NewRewardshare(tc.contract, tc.deployer, tc.withdraw)
				if tc.deployer.Equals(tc.withdraw) {
					rewardshare.WithdrawerAddress = ""
				}

				suite.app.RewardshareKeeper.SetRewardshare(suite.ctx, rewardshare)
				suite.app.RewardshareKeeper.SetDeployerMap(suite.ctx, tc.deployer, tc.contract)
			}

			if tc.expWithdraw {
				suite.app.RewardshareKeeper.SetWithdrawerMap(suite.ctx, tc.withdraw, tc.contract)
			}

			rewardshare, found := suite.app.RewardshareKeeper.GetRewardshare(suite.ctx, tc.contract)
			foundD := suite.app.RewardshareKeeper.IsDeployerMapSet(suite.ctx, tc.deployer, tc.contract)
			foundW := suite.app.RewardshareKeeper.IsWithdrawerMapSet(suite.ctx, tc.withdraw, tc.contract)

			if tc.found {
				suite.Require().True(found, tc.name)
				suite.Require().Equal(tc.deployer.String(), rewardshare.DeployerAddress, tc.name)
				suite.Require().Equal(tc.contract.Hex(), rewardshare.ContractAddress, tc.name)

				suite.Require().True(foundD, tc.name)

				if tc.expWithdraw {
					suite.Require().Equal(tc.withdraw.String(), rewardshare.WithdrawerAddress, tc.name)
					suite.Require().True(foundW, tc.name)
				} else {
					suite.Require().Equal("", rewardshare.WithdrawerAddress, tc.name)
					suite.Require().False(foundW, tc.name)
				}
			} else {
				suite.Require().False(found, tc.name)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestDeleteRewardshare() {
	rewardshare := types.NewRewardshare(contract, deployer, withdraw)
	suite.app.RewardshareKeeper.SetRewardshare(suite.ctx, rewardshare)

	initialFee, found := suite.app.RewardshareKeeper.GetRewardshare(suite.ctx, contract)
	suite.Require().True(found)

	testCases := []struct {
		name     string
		malleate func()
		ok       bool
	}{
		{"existing rewardshare", func() {}, true},
		{
			"deleted rewardshare",
			func() {
				suite.app.RewardshareKeeper.DeleteRewardshare(suite.ctx, rewardshare)
			},
			false,
		},
	}
	for _, tc := range testCases {
		tc.malleate()
		rewardshare, found := suite.app.RewardshareKeeper.GetRewardshare(suite.ctx, contract)
		if tc.ok {
			suite.Require().True(found, tc.name)
			suite.Require().Equal(initialFee, rewardshare, tc.name)
		} else {
			suite.Require().False(found, tc.name)
			suite.Require().Equal(types.Rewardshare{}, rewardshare, tc.name)
		}
	}
}

func (suite *KeeperTestSuite) TestDeleteDeployerMap() {
	suite.app.RewardshareKeeper.SetDeployerMap(suite.ctx, deployer, contract)
	found := suite.app.RewardshareKeeper.IsDeployerMapSet(suite.ctx, deployer, contract)
	suite.Require().True(found)

	testCases := []struct {
		name     string
		malleate func()
		ok       bool
	}{
		{"existing deployer", func() {}, true},
		{
			"deleted deployer",
			func() {
				suite.app.RewardshareKeeper.DeleteDeployerMap(suite.ctx, deployer, contract)
			},
			false,
		},
	}
	for _, tc := range testCases {
		tc.malleate()
		found := suite.app.RewardshareKeeper.IsDeployerMapSet(suite.ctx, deployer, contract)
		if tc.ok {
			suite.Require().True(found, tc.name)
		} else {
			suite.Require().False(found, tc.name)
		}
	}
}

func (suite *KeeperTestSuite) TestDeleteWithdrawMap() {
	suite.app.RewardshareKeeper.SetWithdrawerMap(suite.ctx, withdraw, contract)
	found := suite.app.RewardshareKeeper.IsWithdrawerMapSet(suite.ctx, withdraw, contract)
	suite.Require().True(found)

	testCases := []struct {
		name     string
		malleate func()
		ok       bool
	}{
		{"existing withdraw", func() {}, true},
		{
			"deleted withdraw",
			func() {
				suite.app.RewardshareKeeper.DeleteWithdrawerMap(suite.ctx, withdraw, contract)
			},
			false,
		},
	}
	for _, tc := range testCases {
		tc.malleate()
		found := suite.app.RewardshareKeeper.IsWithdrawerMapSet(suite.ctx, withdraw, contract)
		if tc.ok {
			suite.Require().True(found, tc.name)
		} else {
			suite.Require().False(found, tc.name)
		}
	}
}

func (suite *KeeperTestSuite) TestIsRewardshareRegistered() {
	rewardshare := types.NewRewardshare(contract, deployer, withdraw)
	suite.app.RewardshareKeeper.SetRewardshare(suite.ctx, rewardshare)
	_, found := suite.app.RewardshareKeeper.GetRewardshare(suite.ctx, contract)
	suite.Require().True(found)

	testCases := []struct {
		name     string
		contract common.Address
		ok       bool
	}{
		{"registered rewardshare", contract, true},
		{"rewardshare not registered", common.Address{}, false},
		{"rewardshare not registered", tests.GenerateAddress(), false},
	}
	for _, tc := range testCases {
		found := suite.app.RewardshareKeeper.IsRewardshareRegistered(suite.ctx, tc.contract)
		if tc.ok {
			suite.Require().True(found, tc.name)
		} else {
			suite.Require().False(found, tc.name)
		}
	}
}
