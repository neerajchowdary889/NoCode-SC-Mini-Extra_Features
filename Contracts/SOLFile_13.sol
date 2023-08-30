pragma solidity ^0.4.24;

contract SimpleNumberGame {

    // The current number being guessed
    uint256 private currentNumber;

    // The address of the player who guessed the correct number
    address private winningPlayer;

    // Event emitted when a player guesses the correct number
    event Won(address winner);

    // Constructor that sets the initial number to be guessed
    constructor() public {
        currentNumber = 10;
    }

    // Function to guess the current number
    function guess(uint256 guess) public {
        // If the guess is correct, set the winning player and emit the Won event
        if (guess == currentNumber) {
            winningPlayer = msg.sender;
            emit Won(winningPlayer);
        }
    }

    // Function to get the current number being guessed
    function getCurrentNumber() public view returns (uint256) {
        return currentNumber;
    }

    // Function to get the address of the player who guessed the correct number
    function getWinningPlayer() public view returns (address) {
        return winningPlayer;
    }
}