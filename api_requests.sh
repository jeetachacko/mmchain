#Create_contract takes as input recordingId - integer, userId - integer and contractType - integer
curl --request POST  --url http://localhost:3000/invoke  --header 'content-type: application/x-www-form-urlencoded'   --data =   --data channelid=mychannel  --data chaincodeid=basic  --data function=Create_contract  --data args=333  --data args=6  --data args=1

#Get_contracts takes as input userId - integer
curl --request GET  --url 'http://localhost:3000/query?channelid=mychannel&chaincodeid=basic&function=Get_contracts&args=3'

#Get_allcontracts needs no arguments as input
curl --request GET --url 'http://localhost:3000/query?channelid=mychannel&chaincodeid=basic&function=Get_allcontracts'
