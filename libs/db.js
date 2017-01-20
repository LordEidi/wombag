/*-----------------------------------------------------------------------------
 **
 ** - WomBag - your own read it later service -
 **
 ** Copyright 2017 by
 ** SwordLord - the coding crew - http://www.swordlord.com
 ** and contributing authors
 **
 -----------------------------------------------------------------------------*/
const uuidV4 = require('uuid/v4');

var log = require('../libs/log').log;
var config = require('../config').config;

var Sequelize = require('sequelize');

var sequelize = new Sequelize(config.db_name, config.db_uid, config.db_pwd, {
    dialect: config.db_dialect,
    logging: config.db_logging,
    storage: config.db_storage,
    omitNull: true
});

/**
 * Entry object
 * @type {Model}
 */
// todo: the id as pkey is legacy, pkey should become the primary key, eventually
var ENTRY = sequelize.define('entry', {
        pkey: { type: Sequelize.STRING, allowNull: false, unique: true},
        id: { type: Sequelize.INTEGER, allowNull: false, unique: true, autoIncrement: true, primaryKey: true},
        url: { type: Sequelize.STRING, allowNull: false},
        preview_pic: { type: Sequelize.STRING, allowNull: true},
        domain: { type: Sequelize.TEXT, allowNull: false},
        title: { type: Sequelize.TEXT, allowNull: true},
        content: { type: Sequelize.TEXT, allowNull: true},
        tags: { type: Sequelize.TEXT, allowNull: true},
        starred: { type: Sequelize.BOOLEAN, allowNull: false, defaultValue: false },
        archived: { type: Sequelize.BOOLEAN, allowNull: false, defaultValue: false }
    },
    {
        freezeTableName: true,
        hooks: {
            beforeValidate: function(entry, options) {
                // make sure that the pkey field is set before validation happens
                if(entry.pkey === null || entry.pkey === undefined)
                {
                    entry.pkey = uuidV4();
                }
            }
        },
        instanceMethods: {
            getContent4JSON: function() {
                return JSON.stringify(this.content).slice(1, -1);
            },
            getTitle4JSON: function() {
                return JSON.stringify(this.title).slice(1, -1);
            }
        }
    });

sequelize.sync().then(function()
    {
        log.info("Database structure updated");
    }).error(function(error)
    {
        log.error("Database structure update crashed: " + error);
    }
);

// Exporting.
module.exports = {
    ENTRY: ENTRY,
    sequelize: sequelize
};