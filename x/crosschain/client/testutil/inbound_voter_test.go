//go:build TESTNET
// +build TESTNET

package testutil

import (
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	authcli "github.com/cosmos/cosmos-sdk/x/auth/client/cli"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	crosschainCli "github.com/zeta-chain/zetacore/x/crosschain/client/cli"
	crosschaintypes "github.com/zeta-chain/zetacore/x/crosschain/types"
	observerCli "github.com/zeta-chain/zetacore/x/observer/client/cli"
	observerTypes "github.com/zeta-chain/zetacore/x/observer/types"
)

func (s *IntegrationTestSuite) TestCCTXInboundVoter() {
	tt := []struct {
		name        string
		votes       map[string]observerTypes.VoteType
		finalresult observerTypes.BallotStatus
	}{
		{
			name: "All observers voted success",
			votes: map[string]observerTypes.VoteType{
				"zeta13c7p3xrhd6q2rx3h235jpt8pjdwvacyw6twpax": observerTypes.VoteType_SuccessObservation,
				"zeta1f203dypqg5jh9hqfx0gfkmmnkdfuat3jr45ep2": observerTypes.VoteType_SuccessObservation,
				"zeta1szrskhdeleyt6wmn0nfxvcvt2l6f4fn06uaga4": observerTypes.VoteType_SuccessObservation,
				"zeta16h3y7s7030l4chcznwq3n6uz2m9wvmzu5vwt7c": observerTypes.VoteType_SuccessObservation,
				"zeta1xl2rfsrmx8nxryty3lsjuxwdxs59cn2q65e5ca": observerTypes.VoteType_SuccessObservation,
				"zeta1ktmprjdvc72jq0mpu8tn8sqx9xwj685qx0q6kt": observerTypes.VoteType_SuccessObservation,
				"zeta1ygeyr8pqfjvclxay5234gulnjzv2mkz6lph9y4": observerTypes.VoteType_SuccessObservation,
				"zeta1zegyenj7xg5nck04ykkzndm2qxdzc6v83mklsy": observerTypes.VoteType_SuccessObservation,
				"zeta1us2qpqdcctk6q7qv2c9d9jvjxlv88jscf68kav": observerTypes.VoteType_SuccessObservation,
				"zeta1e9fyaulgntkrnqnl0es4nyxghp3petpn2ntu3t": observerTypes.VoteType_SuccessObservation,
			},
			finalresult: observerTypes.BallotStatus_BallotFinalized_SuccessObservation,
		},
		//{
		//	name: "5 votes only ballot does not get finalized",
		//	votes: map[string]observerTypes.VoteType{
		//		"zeta13c7p3xrhd6q2rx3h235jpt8pjdwvacyw6twpax": observerTypes.VoteType_SuccessObservation,
		//		"zeta1f203dypqg5jh9hqfx0gfkmmnkdfuat3jr45ep2": observerTypes.VoteType_SuccessObservation,
		//		"zeta1szrskhdeleyt6wmn0nfxvcvt2l6f4fn06uaga4": observerTypes.VoteType_SuccessObservation,
		//		"zeta16h3y7s7030l4chcznwq3n6uz2m9wvmzu5vwt7c": observerTypes.VoteType_SuccessObservation,
		//		"zeta1xl2rfsrmx8nxryty3lsjuxwdxs59cn2q65e5ca": observerTypes.VoteType_SuccessObservation,
		//		"zeta1ktmprjdvc72jq0mpu8tn8sqx9xwj685qx0q6kt": observerTypes.VoteType_NotYetVoted,
		//		"zeta1ygeyr8pqfjvclxay5234gulnjzv2mkz6lph9y4": observerTypes.VoteType_NotYetVoted,
		//		"zeta1zegyenj7xg5nck04ykkzndm2qxdzc6v83mklsy": observerTypes.VoteType_NotYetVoted,
		//		"zeta1us2qpqdcctk6q7qv2c9d9jvjxlv88jscf68kav": observerTypes.VoteType_NotYetVoted,
		//		"zeta1e9fyaulgntkrnqnl0es4nyxghp3petpn2ntu3t": observerTypes.VoteType_NotYetVoted,
		//	},
		//	finalresult: observerTypes.BallotStatus_BallotInProgress,
		//},
		//{
		//	name: "1 byzantine vote but correct ballot is still finalized",
		//	votes: map[string]observerTypes.VoteType{
		//		"zeta13c7p3xrhd6q2rx3h235jpt8pjdwvacyw6twpax": observerTypes.VoteType_SuccessObservation,
		//		"zeta1f203dypqg5jh9hqfx0gfkmmnkdfuat3jr45ep2": observerTypes.VoteType_SuccessObservation,
		//		"zeta1szrskhdeleyt6wmn0nfxvcvt2l6f4fn06uaga4": observerTypes.VoteType_SuccessObservation,
		//		"zeta16h3y7s7030l4chcznwq3n6uz2m9wvmzu5vwt7c": observerTypes.VoteType_SuccessObservation,
		//		"zeta1xl2rfsrmx8nxryty3lsjuxwdxs59cn2q65e5ca": observerTypes.VoteType_SuccessObservation,
		//		"zeta1ktmprjdvc72jq0mpu8tn8sqx9xwj685qx0q6kt": observerTypes.VoteType_SuccessObservation,
		//		"zeta1ygeyr8pqfjvclxay5234gulnjzv2mkz6lph9y4": observerTypes.VoteType_SuccessObservation,
		//		"zeta1zegyenj7xg5nck04ykkzndm2qxdzc6v83mklsy": observerTypes.VoteType_FailureObservation,
		//		"zeta1us2qpqdcctk6q7qv2c9d9jvjxlv88jscf68kav": observerTypes.VoteType_SuccessObservation,
		//		"zeta1e9fyaulgntkrnqnl0es4nyxghp3petpn2ntu3t": observerTypes.VoteType_NotYetVoted,
		//	},
		//	finalresult: observerTypes.BallotStatus_BallotFinalized_SuccessObservation,
		//},
		//{
		//	name: "2 ballots with 5 votes each no ballot gets finalized",
		//	votes: map[string]observerTypes.VoteType{
		//		"zeta13c7p3xrhd6q2rx3h235jpt8pjdwvacyw6twpax": observerTypes.VoteType_SuccessObservation,
		//		"zeta1f203dypqg5jh9hqfx0gfkmmnkdfuat3jr45ep2": observerTypes.VoteType_SuccessObservation,
		//		"zeta1szrskhdeleyt6wmn0nfxvcvt2l6f4fn06uaga4": observerTypes.VoteType_SuccessObservation,
		//		"zeta16h3y7s7030l4chcznwq3n6uz2m9wvmzu5vwt7c": observerTypes.VoteType_SuccessObservation,
		//		"zeta1xl2rfsrmx8nxryty3lsjuxwdxs59cn2q65e5ca": observerTypes.VoteType_SuccessObservation,
		//		"zeta1ktmprjdvc72jq0mpu8tn8sqx9xwj685qx0q6kt": observerTypes.VoteType_FailureObservation,
		//		"zeta1ygeyr8pqfjvclxay5234gulnjzv2mkz6lph9y4": observerTypes.VoteType_FailureObservation,
		//		"zeta1zegyenj7xg5nck04ykkzndm2qxdzc6v83mklsy": observerTypes.VoteType_FailureObservation,
		//		"zeta1us2qpqdcctk6q7qv2c9d9jvjxlv88jscf68kav": observerTypes.VoteType_FailureObservation,
		//		"zeta1e9fyaulgntkrnqnl0es4nyxghp3petpn2ntu3t": observerTypes.VoteType_FailureObservation,
		//	},
		//	finalresult: observerTypes.BallotStatus_BallotInProgress,
		//},
		//{
		//	name: "false ballot finalized majority byzantine votes",
		//	votes: map[string]observerTypes.VoteType{
		//		"zeta13c7p3xrhd6q2rx3h235jpt8pjdwvacyw6twpax": observerTypes.VoteType_SuccessObservation,
		//		"zeta1f203dypqg5jh9hqfx0gfkmmnkdfuat3jr45ep2": observerTypes.VoteType_SuccessObservation,
		//		"zeta1szrskhdeleyt6wmn0nfxvcvt2l6f4fn06uaga4": observerTypes.VoteType_SuccessObservation,
		//		"zeta16h3y7s7030l4chcznwq3n6uz2m9wvmzu5vwt7c": observerTypes.VoteType_FailureObservation,
		//		"zeta1xl2rfsrmx8nxryty3lsjuxwdxs59cn2q65e5ca": observerTypes.VoteType_FailureObservation,
		//		"zeta1ktmprjdvc72jq0mpu8tn8sqx9xwj685qx0q6kt": observerTypes.VoteType_FailureObservation,
		//		"zeta1ygeyr8pqfjvclxay5234gulnjzv2mkz6lph9y4": observerTypes.VoteType_FailureObservation,
		//		"zeta1zegyenj7xg5nck04ykkzndm2qxdzc6v83mklsy": observerTypes.VoteType_FailureObservation,
		//		"zeta1us2qpqdcctk6q7qv2c9d9jvjxlv88jscf68kav": observerTypes.VoteType_FailureObservation,
		//		"zeta1e9fyaulgntkrnqnl0es4nyxghp3petpn2ntu3t": observerTypes.VoteType_FailureObservation,
		//	},
		//	finalresult: observerTypes.BallotStatus_BallotInProgress,
		//},
		//{
		//	name: "7 votes only just crossed threshold",
		//	votes: map[string]observerTypes.VoteType{
		//		"zeta13c7p3xrhd6q2rx3h235jpt8pjdwvacyw6twpax": observerTypes.VoteType_SuccessObservation,
		//		"zeta1f203dypqg5jh9hqfx0gfkmmnkdfuat3jr45ep2": observerTypes.VoteType_SuccessObservation,
		//		"zeta1szrskhdeleyt6wmn0nfxvcvt2l6f4fn06uaga4": observerTypes.VoteType_SuccessObservation,
		//		"zeta16h3y7s7030l4chcznwq3n6uz2m9wvmzu5vwt7c": observerTypes.VoteType_SuccessObservation,
		//		"zeta1xl2rfsrmx8nxryty3lsjuxwdxs59cn2q65e5ca": observerTypes.VoteType_SuccessObservation,
		//		"zeta1ktmprjdvc72jq0mpu8tn8sqx9xwj685qx0q6kt": observerTypes.VoteType_NotYetVoted,
		//		"zeta1ygeyr8pqfjvclxay5234gulnjzv2mkz6lph9y4": observerTypes.VoteType_SuccessObservation,
		//		"zeta1zegyenj7xg5nck04ykkzndm2qxdzc6v83mklsy": observerTypes.VoteType_NotYetVoted,
		//		"zeta1us2qpqdcctk6q7qv2c9d9jvjxlv88jscf68kav": observerTypes.VoteType_NotYetVoted,
		//		"zeta1e9fyaulgntkrnqnl0es4nyxghp3petpn2ntu3t": observerTypes.VoteType_SuccessObservation,
		//	},
		//	finalresult: observerTypes.BallotStatus_BallotFinalized_SuccessObservation,
		//},
	}
	for _, test := range tt {
		test := test
		s.Run(test.name, func() {
			broadcaster := s.network.Validators[0]
			for _, val := range s.network.Validators {
				out, err := clitestutil.ExecTestCLICmd(broadcaster.ClientCtx, authcli.GetAccountCmd(), []string{val.Address.String(), "--output", "json"})
				var account authtypes.AccountI
				s.NoError(val.ClientCtx.Codec.UnmarshalInterfaceJSON(out.Bytes(), &account))
				signedTx := BuildSignedGasPriceVote(s.T(), val, s.cfg.BondDenom, account)
				_, err = clitestutil.ExecTestCLICmd(broadcaster.ClientCtx, authcli.GetBroadcastCommand(), []string{signedTx.Name(), "--broadcast-mode", "sync"})
				s.Require().NoError(err)
			}
			for _, val := range s.network.Validators {
				out, err := clitestutil.ExecTestCLICmd(broadcaster.ClientCtx, authcli.GetAccountCmd(), []string{val.Address.String(), "--output", "json"})
				var account authtypes.AccountI
				s.NoError(val.ClientCtx.Codec.UnmarshalInterfaceJSON(out.Bytes(), &account))
				signedTx := BuildSignedTssVote(s.T(), val, s.cfg.BondDenom, account)
				out, err = clitestutil.ExecTestCLICmd(broadcaster.ClientCtx, authcli.GetBroadcastCommand(), []string{signedTx.Name(), "--broadcast-mode", "sync"})
				s.Require().NoError(err)
			}
			for _, val := range s.network.Validators {
				vote := test.votes[val.Address.String()]
				if vote == observerTypes.VoteType_NotYetVoted {
					continue
				}
				out, err := clitestutil.ExecTestCLICmd(broadcaster.ClientCtx, authcli.GetAccountCmd(), []string{val.Address.String(), "--output", "json"})
				var account authtypes.AccountI
				s.NoError(val.ClientCtx.Codec.UnmarshalInterfaceJSON(out.Bytes(), &account))

				message := test.name
				if vote == observerTypes.VoteType_FailureObservation {
					message = message + "falseVote"
				}
				signedTx := BuildSignedInboundVote(s.T(), val, s.cfg.BondDenom, account, message)
				out, err = clitestutil.ExecTestCLICmd(broadcaster.ClientCtx, authcli.GetBroadcastCommand(), []string{signedTx.Name(), "--broadcast-mode", "sync"})
				s.Require().NoError(err)
			}

			s.Require().NoError(s.network.WaitForNBlocks(2))
			ballotIdentifier := GetBallotIdentifier(test.name)
			out, err := clitestutil.ExecTestCLICmd(broadcaster.ClientCtx, observerCli.CmdBallotByIdentifier(), []string{ballotIdentifier, "--output", "json"})
			s.Require().NoError(err)
			ballot := observerTypes.QueryBallotByIdentifierResponse{}
			s.NoError(broadcaster.ClientCtx.Codec.UnmarshalJSON(out.Bytes(), &ballot))
			s.Assert().Equal(len(test.votes), len(ballot.Voters))
			for _, vote := range ballot.Voters {
				if test.votes[vote.VoterAddress] == observerTypes.VoteType_FailureObservation {
					s.Assert().Equal(observerTypes.VoteType_NotYetVoted, vote.VoteType)
					continue
				}
				s.Assert().Equal(test.votes[vote.VoterAddress], vote.VoteType)
			}
			s.Assert().Equal(test.finalresult, ballot.BallotStatus)

			if test.finalresult == observerTypes.BallotStatus_BallotFinalized_SuccessObservation {
				out, err = clitestutil.ExecTestCLICmd(broadcaster.ClientCtx, crosschainCli.CmdShowSend(), []string{ballotIdentifier, "--output", "json"})
				cctx := crosschaintypes.QueryGetCctxResponse{}
				s.NoError(broadcaster.ClientCtx.Codec.UnmarshalJSON(out.Bytes(), &cctx))

			}
		})
	}

}
