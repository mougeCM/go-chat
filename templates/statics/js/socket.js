/**
 * Created by chenmin on 16/8/22.
 */
$(function() {
    var conn;
    var msg = $("#msg");
    var log = $("#log");

    function appendLog(msg) {
        var d = log[0];
        var doScroll = d.scrollTop == d.scrollHeight - d.clientHeight;
        var dt = $("<dl class='me m1'><dt><a>敏哥</a><span>2016-08-22</span></dt><dd>" + msg + "</dd></dl>");
        log.append(dt);
        if (doScroll) {
            d.scrollTop = d.scrollHeight - d.clientHeight;
        }
    }

    $("#form").submit(function() {
        if (!conn) {
            return false;
        }
        if (!msg.val()) {
            return false;
        }
        conn.send(msg.val());
        msg.val("");
        return false
    });

    if (window["WebSocket"]) {
        var host = $(".host").val();
        conn = new WebSocket("ws://"+ host +"/ws");
        conn.onclose = function() {
            appendLog($("<div><b>Connection closed.</b></div>"))
        }
        conn.onmessage = function(evt) {
            appendLog(evt.data);
        }
    } else {
        appendLog($("<div><b>Your browser does not support WebSockets.</b></div>"))
    }
});
