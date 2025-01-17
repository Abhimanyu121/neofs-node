package morph

import (
	"fmt"

	nns "github.com/nspcc-dev/neo-go/examples/nft-nd-nns"
	"github.com/nspcc-dev/neo-go/pkg/core/native"
	"github.com/nspcc-dev/neo-go/pkg/core/state"
	"github.com/nspcc-dev/neo-go/pkg/io"
	"github.com/nspcc-dev/neo-go/pkg/smartcontract"
	"github.com/nspcc-dev/neo-go/pkg/smartcontract/callflag"
	"github.com/nspcc-dev/neo-go/pkg/util"
	"github.com/nspcc-dev/neo-go/pkg/vm"
	"github.com/nspcc-dev/neo-go/pkg/vm/emit"
	"github.com/nspcc-dev/neo-go/pkg/vm/opcode"
)

const defaultNameServiceDomainPrice = 10_0000_0000
const defaultNameServiceSysfee = 4000_0000
const defaultRegisterSysfee = 10_0000_0000 + defaultNameServiceDomainPrice

func (c *initializeContext) setNNS() error {
	ctrPath, err := c.Command.Flags().GetString(contractsInitFlag)
	if err != nil {
		return fmt.Errorf("missing contracts path: %w", err)
	}

	cs, err := c.readContract(ctrPath, nnsContract)
	if err != nil {
		return err
	}

	h := state.CreateContractHash(c.CommitteeAcc.Contract.ScriptHash(), cs.NEF.Checksum, cs.Manifest.Name)
	ok, err := c.nnsRootRegistered(h)
	if err != nil {
		return err
	} else if !ok {
		bw := io.NewBufBinWriter()
		emit.AppCall(bw.BinWriter, h, "addRoot", callflag.All, "neofs")
		if err := c.sendCommitteeTx(bw.Bytes(), -1); err != nil {
			return fmt.Errorf("can't add domain root to NNS: %w", err)
		}
		if err := c.awaitTx(); err != nil {
			return err
		}
	}

	for _, ctrName := range contractList {
		cs, err := c.readContract(ctrPath, ctrName)
		if err != nil {
			return err
		}
		cs.Hash = state.CreateContractHash(c.CommitteeAcc.Contract.ScriptHash(), cs.NEF.Checksum, cs.Manifest.Name)

		domain := ctrName + ".neofs"
		if ok, err := c.nnsDomainAvailable(h, domain); err != nil {
			return err
		} else if !ok {
			continue
		}

		bw := io.NewBufBinWriter()
		emit.AppCall(bw.BinWriter, h, "register", callflag.All,
			domain, c.CommitteeAcc.Contract.ScriptHash())
		emit.Opcodes(bw.BinWriter, opcode.ASSERT)
		emit.AppCall(bw.BinWriter, h, "setRecord", callflag.All,
			domain, int64(nns.TXT), cs.Hash.StringLE())

		sysFee := int64(defaultRegisterSysfee + native.GASFactor)
		if err := c.sendCommitteeTx(bw.Bytes(), sysFee); err != nil {
			return err
		}
	}

	return c.awaitTx()
}

func (c *initializeContext) nnsRootRegistered(nnsHash util.Uint160) (bool, error) {
	params := []smartcontract.Parameter{{Type: smartcontract.StringType, Value: "name.neofs"}}
	res, err := c.Client.InvokeFunction(nnsHash, "isAvailable", params, nil)
	if err != nil {
		return false, err
	}
	return res.State == vm.HaltState.String(), nil
}

func (c *initializeContext) nnsDomainAvailable(nnsHash util.Uint160, domain string) (bool, error) {
	params := []smartcontract.Parameter{{Type: smartcontract.StringType, Value: domain}}
	res, err := c.Client.InvokeFunction(nnsHash, "isAvailable", params, nil)
	if err != nil {
		return false, err
	}
	if len(res.Stack) == 0 {
		return true, nil
	}
	if ok, err := res.Stack[0].TryBool(); err == nil {
		return ok, nil
	}
	return true, nil
}
