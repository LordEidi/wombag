/*-----------------------------------------------------------------------------
 **
 ** - WomBag - your own read it later service -
 **
 ** Copyright 2017 by
 ** SwordLord - the coding crew - http://www.swordlord.com
 ** and contributing authors
 **
 -----------------------------------------------------------------------------*/

var test = require('tape');
var request = require('request');
var moment = require('moment');
var uuid = require('uuid');

var config = require('../../../../config').config;

test('Calling POST on entry', function (t) {

    t.plan(1);

    var payload = "url=http%3A%2F%2Fwww.theregister.co.uk%2F2017%2F01%2F18%2Fdarpa_funds_smart_artillery_shell%2F\r\n";
    //payload += "access_token: ZGJmNTA2MDdmYTdmNWFiZjcxOWY3MWYyYzkyZDdlNWIzOTU4NWY3NTU1MDFjOTdhMTk2MGI3YjY1ZmI2NzM5MA\r\n";

    var options = {
        method: 'POST',
        uri: "http://" + config.ip + ":" + config.port + "/api/entries/",
        body: payload,
        followRedirect: false
    }

    request(options, function (error, response, body) {

        if (!error) {
            t.equal(response.statusCode, 200, "StatusCode matches");
        }
        else {
            t.fail(error);
        }
    });
});