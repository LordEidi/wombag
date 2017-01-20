/*-----------------------------------------------------------------------------
 **
 ** - WomBag - your own read it later service -
 **
 ** Copyright 2017 by
 ** SwordLord - the coding crew - http://www.swordlord.com
 ** and contributing authors
 **
 -----------------------------------------------------------------------------*/

var log = require('../libs/log').log;
var config = require('../config').config;

// Exporting.
/**
 *
 * @type {comm}
 */
module.exports = comm;

/**
 * generates a new request object, constructor like
 * @param req
 * @param res
 * @param reqBody
 * @returns {request}
 */
function comm(req, res, reqBody)
{
    // request object as well as body
    this.req = req;
    this.reqBody = reqBody;

    // response object as well as body we gonna write ourselfs
    this.res = res;
    this.resBody = ""; // response body

    // add oauth headers here
    //    this.user = new userLib.user(username);

    return this;
}

/**
 * Sets the response code given
 * @param responseCode
 */
comm.prototype.setResponseCode = function(responseCode)
{
    log.info("Setting response code: " + responseCode);
    this.res.writeHead(responseCode);
};

/**
 * Writes out the body and sends a response.end
 */
comm.prototype.flushResponse = function()
{
    // prettify XML when we have XML in the body
    var response = this.resBody;

    if(response.substr(0, 5) === "<?xml")
    {
        response = pd.xml(this.resBody);
    }

    log.debug("Returning response: " + response);
    this.res.write(response);
    this.res.end();
};

/**
 * Adds given string to the response body
 * @param str
 */
comm.prototype.appendResBody = function(str)
{
    this.resBody += str;
};

/**
 *
 * @param key
 * @param value
 */
comm.prototype.setHeader = function(key, value)
{
    this.res.setHeader(key, value);
};

/**
 *
 * @param header
 * @returns {*}
 */
comm.prototype.hasHeader = function(header)
{
    return (this.getHeader(header));
};

/**
 *
 * @param header
 * @returns {*}
 */
comm.prototype.getHeader = function(header)
{
    return this.req.headers[header.toLowerCase()];
};
