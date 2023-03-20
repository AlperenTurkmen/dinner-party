#!/bin/bash

# Encode the WAV file as Base64
openssl base64 -in track.wav -out track.base64

# Send the PUT request
curl -X PUT -d @track.wav http://localhost:3000/tracks

# Clean up the temporary file
