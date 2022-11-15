Blockchain :- acts as a medium for sending and receiving cryptocurrencies

In traditional network, banks are involved in transactions. if a person wants to send money to some other person using an ATM machine
or any web browser, then this transaction is performed through the bank. The transfer request is made through the bank and then the bank 
sends the payment to the recipient. This is the process involved in sending money and all the transaction data is present in the bank's
central system. 

In Blockchain, we don't have any one server managing the entire transaction system. Instead, we have many servers that participate 
in the blockchain network as a group. Each transaction record is managed by each and every server in the system. The transaction details 
are stored in a block like manner where each block contains following details :- prev hash, timestamp, nonce and transactions.
A hash is generated corresponding to each block and each block is linked to another block with the help of this hash. Each block contains
the hash of the previous block. Creation of a new block and it's linking to prev block is only performed when the validity of the
transaction is confirmed by the other user to whom the transaction is being performed. At the end we get a chain of blocks and hence
it is known as blockchain!!

The transactional data is stored in each server and hence to hack a blockchain system one needs to hack all the servers in the system
which is impossible. If one tries to hack the data of one server and modifies the block content, then firstly it won't match the data
stored in other servers and secondly the block modified will generate a diff hash key and hence the next block in the chain will get disconnected.
Therefore, hacking one server won't benefit the hacker. 

The process of finding the value of nonce by running an algo on your respective server, bringing the blocks together and getting 
reward for the same is known as mining.