## Express App readME
npm install --save request   
npm install --save express3-handlebars      

docker network create --driver bridge clipper       
docker run --name mongodb --network clipper -p 27017:27017 -d mongo:3.7     


docker run --network clipper --name rides-goapi-shivani -p 3000:3000 -td --env MONGO_SERVER=mongodb --env MONGO_DB=ridesDB  clipper

<b>GoAPI Port Mapping</b>       

| Container name   | Port          | 
|------------------|---------------|
| mongo-rides      | 27017         | 
| mongo-card       | 27018         | 
| mongo-payment    | 27019         | 
| mongo-user       | 27020         | 
| goapi-rides      | 3000          | 
| goapi-card       | 3001          | 
| goapi-payment    | 3002          | 
| goapi-user       | 3003          | 

npm start 

<b> Note: </b> In future, node will also be in a container      

Check localhost:5000/

   