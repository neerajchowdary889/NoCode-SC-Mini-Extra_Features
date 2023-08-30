package main

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

// Ballot represents a single vote in the system.
type Ballot struct {
	// Voter is the address of the voter who cast this ballot.
	Voter common.Address

	// Choice is the option that the voter selected.
	Choice int
}

// Approver represents an account that is authorized to approve ballots.
type Approver struct {
	// Address is the address of the approver.
	Address common.Address

	// Signature is the signature of the approver's approval.
	Signature []byte
}

// VotingContract is a smart contract that implements a simple voting system.
type VotingContract struct {
	// BallotCount is the number of ballots that have been cast.
	BallotCount int

	// ApprovedBallots is a list of the ballots that have been approved.
	ApprovedBallots []Ballot

	// Approvers is a list of the accounts that are authorized to approve ballots.
	Approvers []Approver

	// ContractAddress is the address of the contract.
	ContractAddress common.Address
}

// NewVotingContract creates a new VotingContract instance.
func NewVotingContract(contractAddress common.Address) *VotingContract {
	return &VotingContract{
		BallotCount: 0,
		ApprovedBallots: []Ballot{},
		Approvers: []Approver{},
		ContractAddress: contractAddress,
	}
}

// ApproveBallot approves a ballot.
func (v *VotingContract) ApproveBallot(ballot Ballot, approver Approver) error {
	// Check that the approver is authorized to approve ballots.
	if !v.isApprover(approver) {
		return fmt.Errorf("approver is not authorized")
	}

	// Check that the ballot has not already been approved.
	if v.isBallotApproved(ballot) {
		return fmt.Errorf("ballot has already been approved")
	}

	// Add the ballot to the list of approved ballots.
	v.ApprovedBallots = append(v.ApprovedBallots, ballot)

	// Increment the ballot count.
	v.BallotCount++

	return nil
}

// IsBallotApproved returns true if the given ballot has been approved.
func (v *VotingContract) IsBallotApproved(ballot Ballot) bool {
	for _, approvedBallot := range v.ApprovedBallots {
		if approvedBallot.Voter == ballot.Voter && approvedBallot.Choice == ballot.Choice {
			return true
		}
	}

	return false
}

// IsApprover returns true if the given account is authorized to approve ballots.
func (v *VotingContract) IsApprover(approver Approver) bool {
	for _, approvedApprover := range v.Approvers {
		if approvedApprover.Address == approver.Address &&
			crypto.VerifySignature(v.ContractAddress.Bytes(), approver.Signature, approvedApprover.Address.Bytes()) {
			return true
		}
	}

	return false
}

// ABI is the ABI for the VotingContract contract.
var ABI = abi.ABI{
	Methods: []abi.Method{
		{
			Name: "ApproveBallot",
			Inputs: []abi.Argument{
				{Type: "address", Name: "voter"},
				{Type: "int", Name: "choice"},
				{Type: "address", Name: "approver"},
			},
			Outputs: []abi.Argument{
				{Type: "bool", Name: "success"},
			},
		},
		{
			Name: "IsBallotApproved",
			Inputs: []abi.Argument{
				{Type: "address", Name: "voter"},
				{Type: "int", Name: "choice"},
			}