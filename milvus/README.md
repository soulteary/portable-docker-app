# Embed Milvus

Use this method if you want to get a quick Milvus learning environment, or learn how to use a vector database.

Start an All in One Milvus instance with a simple command:

```bash
docker run --rm -it --name=milvus soulteary/milvus:embed-2.1.0
```

After the command is executed, we will get the following log:

```bash
---Milvus Proxy successfully initialized and ready to serve!---
```

Next, open **another** terminal window and enter the following command to verify that Milvus is available:

```bash
docker exec -it milvus python hello-world.py
```

We will get results similar to the following, Milvus will randomly create some data and demonstrate the basic data manipulation capabilities.

```bash
=== start connecting to Milvus     ===

Does collection hello_milvus exist in Milvus: False

=== Create collection `hello_milvus` ===


=== Start inserting entities       ===

Number of entities in Milvus: 3000

=== Start Creating index IVF_FLAT  ===


=== Start loading                  ===


=== Start searching based on vector similarity ===

hit: (distance: 0.0, id: 2998), random field: -13.0
hit: (distance: 0.16608965396881104, id: 1059), random field: -12.0
hit: (distance: 0.17376846075057983, id: 976), random field: -18.0
hit: (distance: 0.0, id: 2999), random field: -15.0
hit: (distance: 0.09368982911109924, id: 760), random field: -15.0
hit: (distance: 0.1466047316789627, id: 263), random field: -12.0
search latency = 0.2489s

=== Start querying with `random > -14` ===

query result:
-{'pk': 16, 'random': -11.0, 'embeddings': [0.73126, 0.492735, 0.017334, 0.219567, 0.872692, 0.050455, 0.072988, 0.322331]}
search latency = 0.2517s

=== Start hybrid searching with `random > -12` ===

hit: (distance: 0.2523949146270752, id: 2281), random field: -11.0
hit: (distance: 0.34272605180740356, id: 2146), random field: -11.0
hit: (distance: 0.34762364625930786, id: 467), random field: -11.0
hit: (distance: 0.249067485332489, id: 819), random field: -11.0
hit: (distance: 0.3513686954975128, id: 2786), random field: -11.0
hit: (distance: 0.4150209426879883, id: 1486), random field: -11.0
search latency = 0.3481s

=== Start deleting with expr `pk in [0, 1]` ===

query before delete by expr=`pk in [0, 1]` -> result: 
-{'pk': 0, 'random': -17.0, 'embeddings': [0.336848, 0.42659, 0.09273, 0.587923, 0.559363, 0.854384, 0.348711, 0.16934]}
-{'pk': 1, 'random': -20.0, 'embeddings': [0.273309, 0.255625, 0.597282, 0.181107, 0.198166, 0.758371, 0.186847, 0.120927]}

query after delete by expr=`pk in [0, 1]` -> result: []


=== Drop collection `hello_milvus` ===
```

If you want to play freely, you can try to enter the following command to get an interactive Python environment, and then refer to the Milvus documentation to start your vector database journey.


```bash
docker exec -it milvus python
```

