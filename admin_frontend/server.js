let express = require('express');
let serveStatic = require('serve-static');
let route = require('./routes/route');
let logger = require('morgan');
let bodyParser = require('body-parser');

app = express();

app.use(serveStatic(__dirname + "/dist"));
let port = process.env.PORT || 8217;
app.listen(port);

app.set('route',__dirname+'/routes');

app.use(bodyParser.json());
app.use('/',route);
app.use(logger('dev'));

console.log('server started '+ port);
module.exports=app;