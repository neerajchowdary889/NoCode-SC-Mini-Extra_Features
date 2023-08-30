pragma solidity ^0.4.24;

contract TestContract {
    // The contract's balance
    uint256 public balance;

    // The constructor
    function TestContract() public {
        balance = 0;
    }

    // The function to add funds to the contract's balance
    function addFunds(uint256 amount) public {
        balance += amount;
    }

    // The function to withdraw funds from the contract's balance
    function withdrawFunds(uint256 amount) public {
        balance -= amount;
    }
}

// Test case 2: Check that the contract's balance can be increased by adding funds

contract Test {
    function test() public {
        // Deploy the contract
        TestContract contract = new TestContract();

        // Add funds to the contract's balance
        contract.addFunds(100);

        // Check that the contract's balance has increased
        assert(contract.balance == 100);
    }
}

// Test case 3: Check that the contract's balance can be decreased by withdrawing funds

contract Test {
    function test() public {
        // Deploy the contract
        TestContract contract = new TestContract();

        // Add funds to the contract's balance
        contract.addFunds(100);

        // Withdraw funds from the contract's balance
        contract.withdrawFunds(50);

        // Check that the contract's balance has decreased
        assert(contract.balance == 50);
    }
}