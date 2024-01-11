# Go-Blockchain

Overview

1. This is an implementation of a blockchain in Go. A blockchain is a decentralized, distributed ledger that records transactions across a network of computers. Each block in the chain contains a hash of the previous block, creating a chain of blocks.

2. Components
2.1 Block
**Fields:**

-Index: The position of the block in the blockchain.
-Timestamp: The time the block was created.
-Data: The information stored in the block (e.g., transaction data).
-PreviousHash: The hash of the previous block in the chain.
-Hash: The hash of the current block.
**Functions:**

-CalculateHash(block Block) string: Calculates the hash of a block.
-GenerateBlock(previousBlock Block, data string) Block: Creates a new block in the blockchain.

2.2 Blockchain
**Fields:**

-Chain: A slice of blocks representing the entire blockchain.
**Functions:**

-AddBlock(data string): Adds a new block to the blockchain.
-IsValid() bool: Checks if the blockchain is valid by verifying hashes and previous block references.

3. Running the program
   -go run blockchain.go

4. Output
-The program prints information for each block in the blockchain, including index, timestamp, data, previous hash, and hash.
-In the last line of the output indicates whether the blockchain is valid or not.

5. Important Notes
-This is a simplified blockchain implementation that can be expanded upon.
-In a real-world scenario, it would require additional features such as network communication, consensus algorithms, and security measures.
