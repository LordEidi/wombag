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

// Place all your configuration options here

var config =
{
    version_nr: '0.1.0',

    // Server specific configuration
    // Please use a proxy in front of WomBag to support TLS.
    // We suggest you use nginx as the TLS endpoint
    port: 8888,
    //port: 80,
    ip: '127.0.0.1',
    //ip: '0.0.0.0',

    // db specific configuration. you can use whatever sequelize supports.
    db_name: 'wombag',
    db_uid: 'uid',
    db_pwd: 'pwd',
    db_dialect: 'sqlite',
    db_logging: true,
    db_storage: 'wombag.sqlite3'
};

// Exporting.
module.exports = {
    config: config
};