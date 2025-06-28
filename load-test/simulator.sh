#!/bin/bash

# CONFIGURABLE ENDPOINTS
LOGIN_URL="http://localhost:8081/login"
BOOK_URL="http://localhost:8082/book"

# Function to simulate login and booking
login_and_book() {
  user=$1

  # Get token from login API
  token=$(curl -s -X POST $LOGIN_URL \
    -H "Content-Type: application/json" \
    -d "{\"username\":\"user$user\",\"password\":\"pass\"}" | jq -r .token)

  # Use token to book a ticket
  curl -s -X POST $BOOK_URL \
    -H "Content-Type: application/json" \
    -d "{\"token\":\"$token\", \"trainId\":\"1001\", \"passengerName\":\"User$user\"}" > /dev/null
}

# Export function so xargs can call it
export -f login_and_book
export LOGIN_URL BOOK_URL

# Simulate 10 million users (change this for testing)
seq 1 10000000 | xargs -n1 -P1000 -I{} bash -c 'login_and_book {}'
