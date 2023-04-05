# dinner-party
The app from the Google commercial which Addison Rae takes part in.

This app consists of 3 microservices. These Are:

1-Tracks: This microservice puts and gets track data from/to the tracks.json.
2-Search: This microservice gets a base64 encoded wav file and sends it to audd.io and returns the name of the track.
3-CoolTown: This microservice gets a small fragment of a music (base64 encoded) and gets you the wav track (base64 encoded).

Tracks microservice This microservice has the access to the .json file. 
You can change this in the code by replacing all "tracks.json" with the path to your .json file.
It may not work if the .json file can not be read or written from the microservice, please give the full path.

In Audd.io, I put my own key, it can be found under api_token in search microservice.
It may be the unsafest way to include my api, but feel free to change it to your own token from Audd.io :) 

Thanks for reading, have a great dinner party!
