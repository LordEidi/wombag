/*-----------------------------------------------------------------------------
 **
 ** - WomBag - your own read it later service -
 **
 ** Copyright 2017 by
 ** SwordLord - the coding crew - http://www.swordlord.com
 ** and contributing authors
 **
 -----------------------------------------------------------------------------*/

var log4js = require('log4js');
var log = log4js.getLogger("wombagapp");

function initialise()
{
    log.debug("logger loaded");
}

// Exporting.
module.exports = {
    log: log
};

initialise();
