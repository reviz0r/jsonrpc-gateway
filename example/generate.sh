protoc \
  --go_out=plugins=grpc:.   \
  --jsonrpc-gateway_out=.   \
  --doc_out=.               \
  --doc_opt=html,index.html \
  service/*.proto
