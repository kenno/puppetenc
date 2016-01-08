# puppetenc
Puppet External Node Classifier

## Installation

  go get -u github.com/kenno/puppetenc

## Usage & Example

    ./puppetenc matht123.example.com.au
    ---
    name: matht123.example.com.au
    classes:
      role::staff_desktop: {}
    parameters:
      organization: maths
      usage: staff
      printer: MTH-LC00-64-REDCCM022
      owner: z1234567
      room: RC-M022

