#!/bin/bash

# Closeboc73
# This script is to get weather dataw from openweathermap.com in the form of a json file
# so that conky will still display the weather when offline even though it doesn't up to date

# you can use this or replace with yours
api_key=f90d99975cf2748d40dd7a5ad7df132d
# get your city id at https://openweathermap.org/find and replace
lat=49.1952
lon=16.608

url="api.openweathermap.org/data/2.5/weather?lat=${lat}&lon=${lon}&appid=${api_key}&units=metric&lang=cz"
curl ${url} -s -o ./weather.json
