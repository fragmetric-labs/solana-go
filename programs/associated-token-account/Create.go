// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package associated_token_account

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// Create is the `create` instruction.
type Create struct {

	// [0] = [WRITE, SIGNER] payer
	//
	// [1] = [WRITE] associated_token_account
	//
	// [2] = [] wallet
	//
	// [3] = [] mint
	//
	// [4] = [] system_program
	//
	// [5] = [] token_program
	//
	// [6] = [] sysvar_rent
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewCreateInstructionBuilder creates a new `Create` instruction builder.
func NewCreateInstructionBuilder() *Create {
	nd := &Create{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 7),
	}
	return nd
}

// SetPayerAccount sets the "payer" account.
func (inst *Create) SetPayerAccount(payer ag_solanago.PublicKey) *Create {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(payer).WRITE().SIGNER()
	return inst
}

// GetPayerAccount gets the "payer" account.
func (inst *Create) GetPayerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetAssociatedTokenAccount sets the "associated_token_account" account.
func (inst *Create) SetAssociatedTokenAccount(associatedTokenAccount ag_solanago.PublicKey) *Create {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(associatedTokenAccount).WRITE()
	return inst
}

// GetAssociatedTokenAccount gets the "associated_token_account" account.
func (inst *Create) GetAssociatedTokenAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetWalletAccount sets the "wallet" account.
func (inst *Create) SetWalletAccount(wallet ag_solanago.PublicKey) *Create {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(wallet)
	return inst
}

// GetWalletAccount gets the "wallet" account.
func (inst *Create) GetWalletAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetMintAccount sets the "mint" account.
func (inst *Create) SetMintAccount(mint ag_solanago.PublicKey) *Create {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(mint)
	return inst
}

// GetMintAccount gets the "mint" account.
func (inst *Create) GetMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetSystemProgramAccount sets the "system_program" account.
func (inst *Create) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *Create {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "system_program" account.
func (inst *Create) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetTokenProgramAccount sets the "token_program" account.
func (inst *Create) SetTokenProgramAccount(tokenProgram ag_solanago.PublicKey) *Create {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(tokenProgram)
	return inst
}

// GetTokenProgramAccount gets the "token_program" account.
func (inst *Create) GetTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetSysvarRentAccount sets the "sysvar_rent" account.
func (inst *Create) SetSysvarRentAccount(sysvarRent ag_solanago.PublicKey) *Create {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(sysvarRent)
	return inst
}

// GetSysvarRentAccount gets the "sysvar_rent" account.
func (inst *Create) GetSysvarRentAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

func (inst Create) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: ag_binary.TypeIDFromUint8(Instruction_Create),
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst Create) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *Create) Validate() error {
	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Payer is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.AssociatedTokenAccount is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.Wallet is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.Mint is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.TokenProgram is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.SysvarRent is not set")
		}
	}
	return nil
}

func (inst *Create) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("Create")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=0]").ParentFunc(func(paramsBranch ag_treeout.Branches) {})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=7]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("            payer", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("associated_token_", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("           wallet", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("             mint", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("   system_program", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("    token_program", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("      sysvar_rent", inst.AccountMetaSlice.Get(6)))
					})
				})
		})
}

func (obj Create) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	return nil
}
func (obj *Create) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	return nil
}

// NewCreateInstruction declares a new Create instruction with the provided parameters and accounts.
func NewCreateInstruction(
	// Accounts:
	payer ag_solanago.PublicKey,
	associatedTokenAccount ag_solanago.PublicKey,
	wallet ag_solanago.PublicKey,
	mint ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey,
	tokenProgram ag_solanago.PublicKey,
	sysvarRent ag_solanago.PublicKey) *Create {
	return NewCreateInstructionBuilder().
		SetPayerAccount(payer).
		SetAssociatedTokenAccount(associatedTokenAccount).
		SetWalletAccount(wallet).
		SetMintAccount(mint).
		SetSystemProgramAccount(systemProgram).
		SetTokenProgramAccount(tokenProgram).
		SetSysvarRentAccount(sysvarRent)
}

// NewCreateInstruction declares a new Create instruction with the provided parameters and accounts.
func NewCreateInstructionWithDerivation(
// Accounts:
	payer ag_solanago.PublicKey,
	wallet ag_solanago.PublicKey,
	mint ag_solanago.PublicKey,
	tokenProgram ag_solanago.PublicKey) *Create {
	associatedTokenAccount, _, _ := ag_solanago.FindAssociatedTokenAddress(
		wallet,
		mint,
		tokenProgram,
	)
	return NewCreateInstructionBuilder().
		SetPayerAccount(payer).
		SetAssociatedTokenAccount(associatedTokenAccount).
		SetWalletAccount(wallet).
		SetMintAccount(mint).
		SetSystemProgramAccount(ag_solanago.SystemProgramID).
		SetTokenProgramAccount(tokenProgram).
		SetSysvarRentAccount(ag_solanago.SysVarRentPubkey)
}
