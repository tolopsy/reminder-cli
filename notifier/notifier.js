const express = require("express");
const bodyParser = require("body-parser");
const app = express();
const port = process.env.PORT || 5000;

app.use(bodyParser.json());
app.get("/health", (req, res) => res.status(200).send());
app.post()
app.listen(port, () => console.log(`Notifier server is now active at port ${port}`));