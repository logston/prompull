# prompull

A very simple tool to pull metrics out of prometheus.

## Installation

```
go get -u github.com/logston/prompull
```

## Usage

```
prompull '(avg(mystat) by (mylabel))'
```
