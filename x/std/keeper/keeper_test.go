package keeper_test

import (
	"coslms/x/std/keeper"
	"coslms/x/std/types"
	"fmt"
	"sync"
	"testing"

	"github.com/cosmos/cosmos-sdk/simapp"
	"github.com/cosmos/cosmos-sdk/store"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	"github.com/cosmos/cosmos-sdk/testutil"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	dbm "github.com/tendermint/tm-db"
)

type TestSuite struct {
	suite.Suite
	ctx     sdk.Context
	skeeper keeper.Keeper
	*assert.Assertions
	mu      sync.RWMutex
	require *require.Assertions
	t       *testing.T
}

func (s *TestSuite) SetupTest() {
	db := dbm.NewMemDB()
	cms := store.NewCommitMultiStore(db)
	encCfg := simapp.MakeTestEncodingConfig()
	lmsKey := sdk.NewKVStoreKey(types.StoreKey)
	ctx := testutil.DefaultContext(lmsKey, sdk.NewTransientStoreKey("transient_test"))
	keeper := keeper.NewKeeper(lmsKey, encCfg.Codec)
	cms.MountStoreWithDB(lmsKey, storetypes.StoreTypeIAVL, db)
	s.Require().NoError(cms.LoadLatestVersion())
	s.skeeper = keeper
	s.ctx = ctx
}

// T retrieves the current *testing.T context.
func (suite *TestSuite) T() *testing.T {
	suite.mu.RLock()
	defer suite.mu.RUnlock()
	return suite.t
}

// SetT sets the current *testing.T context.
func (suite *TestSuite) SetT(t *testing.T) {
	suite.mu.Lock()
	defer suite.mu.Unlock()
	suite.t = t
	suite.Assertions = assert.New(t)
	suite.require = require.New(t)
}

// Require returns a require context for suite.
func (suite *TestSuite) Require() *require.Assertions {
	suite.mu.Lock()
	defer suite.mu.Unlock()
	if suite.require == nil {
		suite.require = require.New(suite.T())
	}
	return suite.require
}
func TestTestSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

func (s *TestSuite) TestRegisterAdmin() {
	type registerAdminTest struct {
		arg1     types.RegisterAdminRequest
		expected error
	}

	var registerAdminTests = []registerAdminTest{
		{
			arg1:     types.RegisterAdminRequest{Name: "ruth", Address: sdk.AccAddress("sabab").String()},
			expected: nil,
		},
		{
			arg1:     types.RegisterAdminRequest{Name: "ruth1", Address: sdk.AccAddress("sabak").String()},
			expected: nil,
		},
		{
			arg1:     types.RegisterAdminRequest{Name: "ruth2", Address: sdk.AccAddress("sabal").String()},
			expected: nil,
		},
	}
	require := s.Require()
	for _, test := range registerAdminTests {

		if output := s.skeeper.RegisterAdmin(s.ctx, &test.arg1); output != test.expected {
			require.Equal(test.expected, output)
		}
	}

}
func (s *TestSuite) TestAddStudent() {
	students := []*types.Student{
		{
			Address: sdk.AccAddress("a1").String(), Name: "r1", Id: "1",
		},
		{
			Address: sdk.AccAddress("a2").String(), Name: "r2", Id: "2",
		},
		{
			Address: sdk.AccAddress("a3").String(), Name: "r3", Id: "2",
		},
	}
	req := types.AddStudentRequest{
		Admin:    "ruth",
		Students: students,
	}
	s.skeeper.AddStudent(s.ctx, &req)

}
func (s *TestSuite) TestApplyLeave() {
	type applyLeaveTest struct {
		arg1     types.ApplyLeaveRequest
		expected string
	}
	var applyLeaveTests = []applyLeaveTest{
		{
			arg1: types.ApplyLeaveRequest{
				Address: sdk.AccAddress("r1").String(),
				Reason:  "reason 1",
			},
			expected: "Leave Applied Successfully",
		},
		{
			arg1: types.ApplyLeaveRequest{
				Address: sdk.AccAddress("r2").String(),
				Reason:  "reason2",
			},
			expected: "Leave Applied Successfully",
		},
	}
	require := s.Require()
	for _, test := range applyLeaveTests {
		if output := s.skeeper.ApplyLeave(s.ctx, &test.arg1); output != test.expected {
			require.Equal(test.expected, output)
		}
	}

}
func (s *TestSuite) TestAcceptLeave() {
	req := types.AcceptLeaveRequest{
		Admin:   sdk.AccAddress("abcdef").String(),
		LeaveId: "1",
		Status:  types.LeaveStatus_STATUS_ACCEPTED,
	}
	res := s.skeeper.AcceptLeave(s.ctx, &req)
	fmt.Println(res)
}
func (s *TestSuite) TestGetStudent() {
	s.TestAddStudent()

	s.skeeper.GetStudent(s.ctx, &types.GetStudentsRequest{})
}
func (s *TestSuite) TestGetLeaveRequests() {
	s.TestApplyLeave()
	s.skeeper.GetLeaveRequests(s.ctx, &types.GetLeaveRequestsRequest{})
}

func (s *TestSuite) TestGetLeaveApprovedRequests() {

	s.skeeper.GetLeaveApprovedRequests(s.ctx, &types.GetLeaveApprovedRequestsRequest{})
}
