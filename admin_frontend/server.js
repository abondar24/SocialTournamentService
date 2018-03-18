var express = require('express');
var serveStatic = require('serve-static');
var route = require('./routes/route');
var bodyParser = require('body-parser');
var logger = require('morgan');

app = express();

app.use(serveStatic(__dirname + "/dist"));
var port = process.env.PORT || 8217;
app.listen(port);

app.set('route',__dirname+'/routes');

app.use('/',route);
app.use(bodyParser.json());
app.use(logger('dev'));

console.log('server started '+ port);
module.exports=app;