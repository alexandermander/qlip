#!/usr/bin/env bash

inputFileName="$1"

upload_file() {

	if [ -z "$inputFileName" ]; then
		echo "No input file provided. Please provide a file containing the OTP password."
		exit 1
	fi
	bace64_encoded_file=$(base64 "$inputFileName")

	read -p "OTP password: " input_text

	result=$(curl "https://qlip.alexandermander.dk/getotp?$input_text")


	lineTwo=$(echo "$result" | sed -n '2p')
	echo "$lineTwo"


	qlip_value=$(echo "$lineTwo" | grep -oP 'otp=\K[^&]+')

	b64=$(base64 -w0 "$inputFileName" 2>/dev/null || base64 "$inputFileName" | tr -d '\r\n')
printf '{"otp":"%s","qlip":"%s"}' "$qlip_value" "$b64" | \
	curl -s -X POST https://qlip.alexandermander.dk/testpost \
			 -H "Content-Type: application/json" \
			 --data-binary @-
}

download_file() {
	curl -s "https://qlip.alexandermander.dk/" | base64 -d 
}

if [[ " $* " == *" -u "* ]]; then
    upload_file
elif [[ " $* " == *" -d "* ]]; then
    download_file
else
    echo "No -u or -d provided."
fi

