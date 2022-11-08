package keeper_test

import "github.com/kynnoenterprise/core/x/rewardshare/types"

func (suite *KeeperTestSuite) TestParams() {
	params := suite.app.RewardshareKeeper.GetParams(suite.ctx)
	params.EnableRewardshare = true
	suite.Require().Equal(types.DefaultParams(), params)
	params.EnableRewardshare = false
	suite.app.RewardshareKeeper.SetParams(suite.ctx, params)
	newParams := suite.app.RewardshareKeeper.GetParams(suite.ctx)
	suite.Require().Equal(newParams, params)
}
