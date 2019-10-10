package main

import "errors"

var errGFBadResponse error = errors.New("google fonts bad response, please try later")
var errFlutterInvalidPubSpec error = errors.New("your flutter pubspec are invalid")
