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
	eKeeper keeper.Keeper
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
	s.eKeeper = keeper
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

func (s *TestSuite) TestAddStudent() {
	students := []*types.Student{
		{
			Address: sdk.AccAddress("abcd").String(), Name: "name1", Id: "12345",
		},
	}
	req := types.AddStudentRequest{
		Admin:    "admin1",
		Students: students,
	}
	res := s.eKeeper.AddStudent(s.ctx, &req)
	fmt.Println(res)
}

type registerAdminTest struct {
	arg1     types.RegisterAdminRequest
	expected string
}

var registerAdminTests = []registerAdminTest{
	{
		arg1: types.RegisterAdminRequest{
			Name:    "admin1",
			Address: sdk.AccAddress("abcd").String(),
		},
		expected: "Admin Registered",
	},
}

func (s *TestSuite) TestRegisterAdmin() {
	require := s.Require()
	for _, test := range registerAdminTests {

		if output := s.eKeeper.RegisterAdmin(s.ctx, &test.arg1); output != test.expected {
			require.Equal(test.expected, output)
		}
		s.eKeeper.GetAdmin(s.ctx, sdk.AccAddress("sakjhfdd").String())
	}

}
