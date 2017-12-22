/*-----------------------------------------------------------------------------
 **
 ** - WomBag - your own read it later service -
 **
 ** Copyright 2017 by
 ** SwordLord - the coding crew - http://www.swordlord.com
 ** and contributing authors
 **
 -----------------------------------------------------------------------------*/

const qs = require('querystring');
const moment = require("moment");
var url = require('url');

const log = require('../libs/log').log;
const config = require('../../../config').config;

const ENTRY = require('../libs/db').ENTRY;

// Exporting.
module.exports = {
    put: put,
    getQuery: getQuery,
    getOne: getOne,
    delete: del,
    patch: patch
};

/**
 *
 * @param {ENTRY} e - instance of an ENTRY object
 * @returns {string}
 */
function formatEntry(e)
{
    var entry = "";

    entry += "{";
    entry += "    \"_links\": {";
    entry += "      \"self\": {";
    entry += "        \"href\": \"/api/entries/" + e.id + "\"";
    entry += "      }";
    entry += "    },";
    entry += "    \"annotations\": [ ],";
    entry += "    \"content\": \"" + e.getContent4JSON() + "\",";
    entry += "    \"created_at\": \"" + moment(e.createdAt).format() + "\",";
    entry += "    \"domain_name\": \"" + e.domain + "\",";
    entry += "    \"id\": " + e.id + ",";
    entry += "    \"is_archived\": " + (e.archived ? "1" : "0") + ",";
    entry += "    \"is_starred\": " + (e.starred ? "1" : "0") + ",";
    entry += "    \"language\": \"en\",";
    entry += "    \"mimetype\": \"text/html\",";
    //entry += "    \"preview_picture\": \"http_to_pic\",";
    entry += "    \"reading_time\": 1,";
    entry += "    \"tags\": [ ],";
    entry += "    \"title\": \"" + e.getTitle4JSON() + "\",";
    entry += "    \"updated_at\": \"" + moment(e.updatedAt).format() + "\",";
    entry += "    \"url\": \"" + e.url + "\",";
    entry += "    \"user_email\": \"lordeidi@lordei.di\",";
    entry += "    \"user_id\": 1,";
    entry += "    \"user_name\": \"lordeidi\"";

    entry += "}";

    return entry;
}

/**
 * Client starts query asking for entries
 * @param {comm} c - comm object
 * @param {query} q - query object
 */
function getQuery(c, q)
{
    log.debug("entry.getQuery called");

    c.setHeader("Content-Type", "application/json");
    c.setResponseCode(200);
    
    log.debug(JSON.stringify(q));

    var page = 1;
    var offset = 0;
    var pageSize = 30;

    if(q.hasOwnProperty('page'))
    {
        page = q.page;
    }

    if(q.hasOwnProperty('perPage'))
    {
        pageSize = q.perPage;
    }

    offset = (pageSize * page) - pageSize;

    ENTRY.findAndCountAll(
        {
            limit: pageSize,
            offset: offset
        }
    ).then(function(result)
    {
        c.appendResBody("{");
        c.appendResBody("    \"_embedded\": {");
        c.appendResBody("        \"items\": [ ");

        // we dont use result.count == total row number
        // but result.rows.length for this is the current page size
        for (var j=0; j < result.rows.length; ++j) {

            if(j > 0)
            {
                c.appendResBody(",");
            }
            var entry = result.rows[j];
            c.appendResBody(formatEntry(entry, j));
        }

        c.appendResBody("    ] },");
        c.appendResBody("    \"_links\": {");
        c.appendResBody("        \"first\": {");
        c.appendResBody("            \"href\": \"http://" + config.ip + ":" + config.port + "/api/entries?page=1&perPage=30\"");
        c.appendResBody("        },");
        c.appendResBody("        \"last\": {");
        c.appendResBody("            \"href\": \"http://" + config.ip + ":" + config.port + "/api/entries?page=1&perPage=30\"");
        c.appendResBody("        },");
        c.appendResBody("        \"self\": {");
        c.appendResBody("            \"href\": \"http://" + config.ip + ":" + config.port + "/api/entries?page=1&perPage=30\"");
        c.appendResBody("        }");
        c.appendResBody("    },");
        c.appendResBody("    \"limit\": 30,");
        c.appendResBody("    \"page\": 1,");
        c.appendResBody("    \"pages\": 1,");
        c.appendResBody("    \"total\": " + result.count );
        c.appendResBody("}");

        c.flushResponse();
    });

}

/**
 * Clients sends new record to be stored
 * @param c
 */
function put(c)
{
    log.debug("entry.put called");

    var body = c.reqBody;

    console.log(body);

    var post = qs.parse(body);

    log.debug("Putting URL: " + post.url);

    var u = url.parse(post.url);

    ENTRY
        .create({ url: u.href, domain : u.host })
        .then(function(entry) {

            const ra = require('read-art');
            ra(u.href, function(err, art, options, resp){
                if(err){
                    log.error(err);
                    c.setResponseCode(500);
                    c.flushResponse();
                }
                else {
                    entry.title = art.title;
                    entry.content = art.content;

                    entry.save().then(function()
                    {
                        log.info('entry ' + entry.pkey + 'updated with title and content');

                        c.setResponseCode(200);
                        c.appendResBody(formatEntry(entry));
                        c.flushResponse();
                    });
                }
            });

        }).catch(function (err) {

            log.error("insert failed: " + err.detail);

            if(err.original != null && err.original != undefined)
            {
                log.error("detail: " + err.original.message + " - " + err.original.sql);
            }

            c.setResponseCode(500);
            c.appendResBody("Put");
            c.flushResponse();
            return;
        });
}

/**
 * Clients requests entry to be deleted
 * @param {comm} c
 * @param {ENTRY} e
 * @param {query} q
 */
function del(c, entryId, q)
{
    log.debug("entry.delete called");

    c.setHeader("Content-Type", "application/json");
    c.setHeader("Server", "WomBag");
    c.setResponseCode(200);

    ENTRY.find({ where: {id: entryId} }).then(function(entry)
        {
            if(entry === null)
            {
                log.warn('err: could not find entry');
            }
            else
            {
                entry.destroy().then(function()
                {
                    log.debug('entry deleted');
                })
            }

            c.flushResponse();

        }).catch(function (err) {

            log.error("delete failed: " + err.detail);

            if(err.original != null && err.original != undefined)
            {
                log.error("detail: " + err.original.message + " - " + err.original.sql);
            }

            c.setResponseCode(500);
            c.appendResBody("Delete");
            c.flushResponse();
            return;
        });
}

/**
 * Clients asks for one specific entry
 * @param c
 */
function getOne(c)
{

}

/**
 * Changes one entry
 * @param c
 */
function patch(c, entryId, query)
{
    log.debug("entry.patch called");

    c.setHeader("Content-Type", "application/json");
    c.setHeader("Server", "WomBag");

    ENTRY.find({ where: {id: entryId} }).then(function(entry)
    {
        if(entry === null)
        {
            log.warn('err: could not find entry');

            c.setResponseCode(500);
            c.flushResponse();
        }
        else
        {
            var post = qs.parse(c.reqBody);

            if(post.starred !== undefined)
            {
                entry.starred = !!parseInt(post.starred);
            }
            if(post.archive !== undefined)
            {
                entry.archived = !!parseInt(post.archive);
            }
            entry.save().then(function()
            {
                log.debug('entry patched');

                c.setResponseCode(200);
                c.appendResBody(formatEntry(entry));
                c.flushResponse();
            })
        }

    }).catch(function (err) {

        log.error("patch failed: " + err.detail);

        if(err.original != null && err.original != undefined)
        {
            log.error("detail: " + err.original.message + " - " + err.original.sql);
        }

        c.setResponseCode(500);
        c.appendResBody("Patch");
        c.flushResponse();
        return;
    });
}
//EOF