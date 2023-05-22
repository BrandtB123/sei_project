# sei_project
A simple indexer in Golang, ran on AWS with a Postgres RDS backend

# APIs

## Endpoints
### Get Proposer Blocks
URL: /api/blocksByProposer
Method: GET
Description: Retrieves the blocks proposed by a specific proposer.
Query Parameters:
proposer (required): The ID of the proposer.
Response:
Status Code 200: Returns the block heights proposed by the specified proposer.
Status Code 400: Indicates a missing or invalid query parameter.
### Get Transactions in Past N Blocks
URL: /totalTxs
Method: GET
Description: Retrieves the total number of transactions in the past N blocks.
Query Parameters:
n (required): The number of blocks to consider.
Response:
Status Code 200: Returns the total number of transactions in the past N blocks.
Status Code 400: Indicates a missing or invalid query parameter.
### Get N Peers Over N Blocks
URL: /topNPeers
Method: GET
Description: Retrieves the top N peers based on a scoring mechanism over a specified number of blocks.
Query Parameters:
n (required): The number of peers to retrieve.
Response:
Status Code 200: Returns the top N peers based on the scoring mechanism.
Status Code 400: Indicates a missing or invalid query parameter.
