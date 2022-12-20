# Replicated MongoDB with Arbiter using Docker Swarm

## How to use

1. Create a Docker Swarm cluster

2. Add the following labels to the nodes you want to use for MongoDB

    1. mongo.replica=1
    2. mongo.replica=2
    3. mongo.arbiter=1

3. Run the following command to deploy the stack
  
    ``$ docker stack deploy -c mongo-stack.yml mongors``

4. Copy file *initrs* to primary node

5. Apply replica set init script on primary node(assume, mongo1)
  
    ``$ docker exec -it $(docker ps -qf name=mongors_mongo1) mongosh --eval "$(cat initrs)"``

6. Check the status of the replica set

    ``$ docker exec -it $(docker ps -qf name=mongors_mongo1) mongosh --eval "rs.status()"``
