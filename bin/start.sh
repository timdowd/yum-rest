export ENV="dev"

go build -o server.exe \
&& \
./server.exe both \
--servicename="yum-rest" \
--server-ip="localhost" \
--rest-port="6000" \
--rpc-port="6001" \
--rpc-address="localhost" \
--traceserviceaccountfile="${HOME}/Documents/serviceAccountFiles/trace-serviceaccount.json" \
--projectid="phdigidev"
