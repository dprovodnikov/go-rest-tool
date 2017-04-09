const express = require("express")
express()
  .use(express.static("."))
  .listen(8080)

