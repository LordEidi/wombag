/*-----------------------------------------------------------------------------
 **
 ** - WomBag - your own read it later service -
 **
 ** Copyright 2017 by
 ** SwordLord - the coding crew - http://www.swordlord.com
 ** and contributing authors
 **
 ** This program is free software; you can redistribute it and/or modify it
 ** under the terms of the GNU General Public License as published by the Free
 ** Software Foundation, either version 3 of the License, or (at your option)
 ** any later version.
 **
 ** This program is distributed in the hope that it will be useful, but WITHOUT
 ** ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or
 ** FITNESS FOR A PARTICULAR PURPOSE.  See the GNU General Public License for
 ** more details.
 **
 ** You should have received a copy of the GNU General Public License along
 ** with this program. If not, see <http://www.gnu.org/licenses/>.
 **
 **-----------------------------------------------------------------------------
 **
 ** Original Authors:
 ** LordEidi@swordlord.com
 **
 ** $Id:
 **
-----------------------------------------------------------------------------*/
var config = require('./config').config;
var http = require('http');
var url = require('url');
var log = require('./libs/log').log;

var hE = require('./handler/entry');

var crossroads = require('crossroads');
crossroads.ignoreState = true;

var comm = require('./libs/comm');

/**
 * Called when the URL is not matched against any known/defined pattern
 * @param comm
 * @param path
 */
function onBypass(c, path)
{
    log.info('URL unknown: ' + c.req.url);

    c.res.writeHead(200);
    c.res.write(c.req.url + " is not known");
    c.res.end();
}

/**
 * Gets called when the / URL is hit
 * @param comm
 */
function onHitRoot(c)
{
    log.debug("Called the root");
    c.res.writeHead(200);
    c.res.write("Root");
    c.res.end();
}

function onRemoveAnnotation(c)
{
   c.res.writeHead(200);
   c.res.write("Ping");
   c.res.end();
}
function onUpdateAnnotation(c)
{
    c.res.writeHead(200);
    c.res.write("Ping");
    c.res.end();
}
function onRetrieveAnnotation(c)
{
    c.res.writeHead(200);
    c.res.write("Ping");
    c.res.end();
}
function onCreateNewAnnotation(c)
{
    c.res.writeHead(200);
    c.res.write("Ping");
    c.res.end();
}
/**
 * Client calls /api/entries:?query:'
 * @param c
 */
function onCreateOrRetrieveEntries(c, query)
{
    switch(c.req.method)
    {
        case "GET":
            log.info("get entry");
            log.debug(c.reqBody);
            hE.getQuery(c, query);
            break;
        case "POST":
            log.info("post entry");
            log.debug(c.reqBody);
            hE.put(c);
            break;
        default:
            log.error("method not supported");
            c.setResponseCode(200);
            c.appendResBody("Ping");
            c.flushResponse();
            break;
    }

}
function onDeleteGetSingleChangeEntry(c, entry, format, query)
{
    switch(c.req.method)
    {
        case "DELETE":
            log.info("delete entry");
            log.debug(c.reqBody);
            hE.delete(c, entry, query);
            break;
        case "PATCH":
            log.info("patch entry");
            log.debug(c.reqBody);
            hE.patch(c, entry, query);
            break;
        default:
            log.error("method not supported");
            c.setResponseCode(500);
            c.appendResBody("Ping");
            c.flushResponse();
            break;
    }
}
function onGetSingleEntryFormatted(c)
{
    c.res.writeHead(200);
    c.res.write("Ping");
    c.res.end();
}
function onReloadEntry(c)
{
    c.res.writeHead(200);
    c.res.write("Ping");
    c.res.end();
}
function onRetrieveTagsForEntry(c)
{
    c.res.writeHead(200);
    c.res.write("Ping");
    c.res.end();
}
function onAddTagsToEntry(c)
{
    c.res.writeHead(200);
    c.res.write("Ping");
    c.res.end();
}
function onDeleteTagsOnEntry(c)
{
    c.res.writeHead(200);
    c.res.write("Ping");
    c.res.end();
}
function onDeleteTagOnEntry(c)
{
    c.res.writeHead(200);
    c.res.write("Ping");
    c.res.end();
}
function onRetrieveAllTags(c)
{
    c.res.writeHead(200);
    c.res.write("Ping");
    c.res.end();
}
function onRemoveTagsFromEveryEntry(c)
{
    c.res.writeHead(200);
    c.res.write("Ping");
    c.res.end();
}
function onRemoveTagFromEveryEntry(c)
{
    c.res.writeHead(200);
    c.res.write("Ping");
    c.res.end();
}
function onRetrieveVersionNumber(c)
{
    c.res.writeHead(200);
    c.res.write("Ping");
    c.res.end();
}

/**
 * Ah well, you get a hard coded token, but since there is no authentication yet :)
 * @param c
 */
// TODO: Clean up...
function onOAuth(c)
{
    c.setHeader("Content-Type", "application/json");
    c.setResponseCode(200);

    var response = "{";
    response += "\"access_token\": \"ZGJmNTA2MDdmYTdmNWFiZjcxOWY3MWYyYzkyZDdlNWIzOTU4NWY3NTU1MDFjOTdhMTk2MGI3YjY1ZmI2NzM5MA\",";
    response += "\"expires_in\": 3600,";
    response += "\"refresh_token\": \"OTNlZGE5OTJjNWQwYzc2NDI5ZGE5MDg3ZTNjNmNkYTY0ZWZhZDVhNDBkZTc1ZTNiMmQ0MjQ0OThlNTFjNTQyMQ\",";
    response += "\"scope\": null,";
    response += "\"token_type\": \"bearer\" }";

    c.appendResBody(response);
    c.flushResponse();
}

crossroads.addRoute('/api/annotations/{annotation}.{format}{?query}', onRemoveAnnotation); // DELETE
crossroads.addRoute('/api/annotations/{annotation}.{format}{?query}', onUpdateAnnotation); // PUT
crossroads.addRoute('/api/annotations/{annotation_id}.{format}{?query}', onRetrieveAnnotation); // GET
crossroads.addRoute('/api/annotations/{entry}.{_format}{?query}', onCreateNewAnnotation); // POST
crossroads.addRoute('/api/entries/:?query:', onCreateOrRetrieveEntries); // POST or GET
crossroads.addRoute('/api/entries/{entry}:format::?query:', onDeleteGetSingleChangeEntry); // DELETE or GET or PATCH
crossroads.addRoute('/api/entries/{entry}/export.{_format}{?query}', onGetSingleEntryFormatted); // GET
crossroads.addRoute('/api/entries/{entry}/reload.{_format}{?query}', onReloadEntry); // PATCH
crossroads.addRoute('/api/entries/{entry}/tags.{_format}{?query}', onRetrieveTagsForEntry); // GET
crossroads.addRoute('/api/entries/{entry}/tags.{_format}{?query}', onAddTagsToEntry); // POST
crossroads.addRoute('/api/entries/{entry}/tags/{tag}.{_format}{?query}', onDeleteTagsOnEntry); // DELETE
crossroads.addRoute('/api/tag/label.{_format}{?query}', onDeleteTagOnEntry); // DELETE
crossroads.addRoute('/api/tags.{_format}{?query}', onRetrieveAllTags); // GET
crossroads.addRoute('/api/tags/label.{_format}{?query}', onRemoveTagsFromEveryEntry); // DELETE
crossroads.addRoute('/api/tags/{tag}.{_format}{?query}', onRemoveTagFromEveryEntry); // DELETE
crossroads.addRoute('/api/version.{_format}', onRetrieveVersionNumber); // GET

crossroads.addRoute('oauth/v2/token', onOAuth);

crossroads.addRoute('/', onHitRoot);
crossroads.bypassed.add(onBypass);

// start the server and process requests
var server = http.createServer(function (req, res)
{
    log.debug("Method: " + req.method + ", URL: " + req.url + ", Query: " + req.query);

    // will contain the whole body submitted
	var reqBody = "";

    req.on('data', function (data)
    {
        reqBody += data.toString();
    });

    req.on('end',function()
    {
        // headers only used sometimes
        var access_token_header = req.headers['access_token'];
        //var access_token_url = req.query['access_token'] ? req.query['access_token'] : '';

        var c = new comm(req, res, reqBody);

        var sUrl = url.parse(req.url).href;
        log.debug("Request body: " + reqBody);
        crossroads.parse(sUrl, [c]);
    });
});

server.listen(config.port);

server.on('error', function (e)
{
    log.warn('Caught error: ' + e.message);
    log.debug(e.stack);
});

process.on('uncaughtException', function(err)
{
    log.warn('Caught exception: ' + err.message);
    log.debug(err.stack);
});

// Put a friendly message on the terminal
log.info("WomBag server running at http://" + config.ip + ":" + config.port + "/");
