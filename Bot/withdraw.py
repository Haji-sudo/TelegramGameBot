from block_io import BlockIo
import sys
import json
version = 2 
token = sys.argv[1]
pin = sys.argv[2]
amount_to_send = sys.argv[3]
address = sys.argv[4]
block_io = BlockIo(token, pin, version)
prepared_transaction = block_io.prepare_transaction(amounts=amount_to_send,priority="low",to_addresses=address)
created_transaction_and_signatures = block_io.create_and_sign_transaction(prepared_transaction)
response = block_io.submit_transaction(transaction_data=created_transaction_and_signatures)
print(json.dumps(response))