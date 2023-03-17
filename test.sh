#!/bin/bash

# Encode the WAV file as Base64
openssl base64 -in track.wav -out track.base64

# Send the PUT request
curl -X PUT -d @track.base64 http://localhost:3000/tracks

# Clean up the temporary file
rm track.base64
