package genhelpers_test

import (
	"context"
	"encoding/hex"
	"fmt"
	"math/rand"
	"path/filepath"
	"testing"

	sdkmath "cosmossdk.io/math"
	"github.com/babylonchain/babylon/cmd/babylond/cmd/genhelpers"
	"github.com/babylonchain/babylon/testutil/datagen"
	"github.com/babylonchain/babylon/testutil/helper"
	bbn "github.com/babylonchain/babylon/types"
	btcctypes "github.com/babylonchain/babylon/x/btccheckpoint/types"
	btcstktypes "github.com/babylonchain/babylon/x/btcstaking/types"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/cometbft/cometbft/libs/tempfile"
	"github.com/cosmos/cosmos-sdk/client"
	genutiltest "github.com/cosmos/cosmos-sdk/x/genutil/client/testutil"
	genutiltypes "github.com/cosmos/cosmos-sdk/x/genutil/types"
	"github.com/stretchr/testify/require"
)

func FuzzCmdSetBtcDels(f *testing.F) {
	datagen.AddRandomSeedsToFuzzer(f, 10)

	f.Fuzz(func(t *testing.T, seed int64) {
		r := rand.New(rand.NewSource(seed))

		qntBtcDels := int(datagen.RandomInt(r, 10)) + 1
		btcDelsToAdd := make([]*btcstktypes.BTCDelegation, qntBtcDels)

		fp, err := datagen.GenRandomFinalityProvider(r)
		require.NoError(t, err)
		covenantSKs, _, covenantQuorum := datagen.GenCovenantCommittee(r)
		slashingAddress, err := datagen.GenRandomBTCAddress(r, &chaincfg.SimNetParams)
		require.NoError(t, err)

		startHeight := datagen.RandomInt(r, 100) + 1
		endHeight := datagen.RandomInt(r, 1000) + startHeight + btcctypes.DefaultParams().CheckpointFinalizationTimeout + 1
		slashingRate := sdkmath.LegacyNewDecWithPrec(int64(datagen.RandomInt(r, 41)+10), 2)
		slashingChangeLockTime := uint16(101)

		for i := 0; i < qntBtcDels; i++ {
			delSK, _, err := datagen.GenRandomBTCKeyPair(r)
			require.NoError(t, err)

			del, err := datagen.GenRandomBTCDelegation(
				r,
				t,
				[]bbn.BIP340PubKey{*fp.BtcPk},
				delSK,
				covenantSKs,
				covenantQuorum,
				slashingAddress.EncodeAddress(),
				startHeight, endHeight, 10000,
				slashingRate,
				slashingChangeLockTime,
			)
			require.NoError(t, err)
			btcDelsToAdd[i] = del
		}

		home := t.TempDir()
		h := helper.NewHelper(t)

		app := h.App
		cdc := app.AppCodec()

		err = genutiltest.ExecInitCmd(app.BasicModuleManager, home, cdc)
		require.NoError(t, err)

		clientCtx := client.Context{}.
			WithCodec(app.AppCodec()).
			WithHomeDir(home).
			WithTxConfig(app.TxConfig())
		ctx := context.WithValue(context.Background(), client.ClientContextKey, &clientCtx)

		jsonBytes, err := cdc.MarshalJSON(&btcstktypes.GenesisState{
			BtcDelegations: btcDelsToAdd,
		})
		require.NoError(t, err)

		btcDelsToAddFilePath := filepath.Join(home, "delsToAdd.json")
		err = tempfile.WriteFileAtomic(btcDelsToAddFilePath, jsonBytes, 0600)
		require.NoError(t, err)

		cmdSetBtcDel := genhelpers.CmdSetBtcDels()
		cmdSetBtcDel.SetArgs([]string{btcDelsToAddFilePath})
		cmdSetBtcDel.SetContext(ctx)

		// Runs the cmd to write into the genesis
		err = cmdSetBtcDel.Execute()
		require.NoError(t, err)

		cmtcfg, err := genutiltest.CreateDefaultCometConfig(home)
		require.NoError(t, err)

		// Verifies that the new genesis were created
		appState, _, err := genutiltypes.GenesisStateFromGenFile(cmtcfg.GenesisFile())
		require.NoError(t, err)

		btcstkGenState := btcstktypes.GenesisStateFromAppState(cdc, appState)
		// make sure the same quantity of BTC delegations were created.
		require.Equal(t, qntBtcDels, len(btcstkGenState.BtcDelegations))

		for i := 0; i < qntBtcDels; i++ {
			bzAdd, err := btcDelsToAdd[i].Marshal()
			require.NoError(t, err)

			bzGen, err := btcstkGenState.BtcDelegations[i].Marshal()
			require.NoError(t, err)

			require.Equal(t, hex.EncodeToString(bzAdd), hex.EncodeToString(bzGen))
		}

		// tries to add again, it should error out
		btcDel := btcDelsToAdd[0]

		hash, err := btcDel.GetStakingTxHash()
		require.NoError(t, err)
		key := hash.String()

		err = cmdSetBtcDel.Execute()
		require.EqualError(t, err, fmt.Errorf("error: btc delegation: %+v\nwas already set on genesis, or contains the same staking tx hash %s than another btc delegation", btcDel, key).Error())
	})
}
