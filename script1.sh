#!/bin/sh
ID="track"
ESCAPED=`perl -e "use URI::Escape; print uri_escape(\"$ID\")"`
AUDIO=`base64 -i "$ID".wav`
echo "AUDIO: $AUDIO"  # Add this line to print the audio variable
RESOURCE=localhost:3000/tracks/$ESCAPED
echo "{ \"Id\":\"$ID\", \"Audio\":\"$AUDIO\" }" > input
curl -v -X PUT -d @input $RESOURCE
