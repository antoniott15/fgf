package main

import "errors"

var errGFBadResponse error = errors.New("google fonts bad response, please try later")
var errFlutterInvalidPubSpec error = errors.New("your flutter pubspec are invalid")
var errInvalidCommand error = errors.New("invalid command, please use a valid command")
var errNumberOfParamsNotValid error = errors.New("invalid params, please use a correct number of params")
var errFamilyFontNotFound error = errors.New("family font not exist or name are incorrect")
var errNotImplemented error = errors.New("action not implemented yet")
