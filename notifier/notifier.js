const express = require("express");
const bodyParser = require("body-parser");
const notifier = require("node-notifier");
const app = express();
const port = process.env.PORT || 5000;
const path = require("path");

app.use(bodyParser.json());

app.get("/health", (req, res) => res.status(200).send());
app.post("/notify", (req, res) => {
    notify(req.body, callbackFn = (reply) => res.send(reply))
});

app.listen(port, () => console.log(`Notifier server is now active at port ${port}`));

const notify = ({title, message}, callbackFn) => {
    notifier.notify(
        {
            title: title || "No title",
            message: message || "No message",
            icon: path.join(__dirname, "icon.jpg"),
            sound: true,
            wait: true,
            reply: true,
            appId: "Local Man",
            closeLabel: "Completed?",
            timeout: 15,

        },
        (err, response, reply) => {
            callbackFn(reply)
        }
    );
};