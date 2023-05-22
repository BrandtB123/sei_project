# sei_project
A simple indexer in Golang, ran on AWS with a Postgres RDS backend
Domain: http://ec2-18-218-245-88.us-east-2.compute.amazonaws.com/

# APIs

## Endpoints
### Get Proposer Blocks
URL: /api/blocksByProposer<br>
Method: GET<br>
Description: Retrieves the blocks proposed by a specific proposer.<br>
Query Parameters: proposer (required): The ID of the proposer.<br>
Response:<br>
Status Code 200: Returns the block heights proposed by the specified proposer.<br>
Status Code 400: Indicates a missing or invalid query parameter.<br>
### Get Transactions in Past N Blocks
URL: /totalTxs<br>
Method: GET<br>
Description: Retrieves the total number of transactions in the past N blocks.<br>
Query Parameters: n (required): The number of blocks to consider.<br>
Response:<br>
Status Code 200: Returns the total number of transactions in the past N blocks.<br>
Status Code 400: Indicates a missing or invalid query parameter.<br>
### Get N Peers Over N Blocks
URL: /topNPeers<br>
Method: GET<br>
Description: Retrieves the top N peers based on a scoring mechanism over a specified number of blocks.<br>
Query Parameters: n (required): The number of peers to retrieve.<br>
Response:<br>
Status Code 200: Returns the top N peers based on the scoring mechanism.<br>
Status Code 400: Indicates a missing or invalid query parameter.<br>
