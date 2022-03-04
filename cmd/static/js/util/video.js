// load the Youtube JS api and get going
var j = document.createElement("script"),
    f = document.getElementsByTagName("script")[0];
    j.src = "https://www.youtube.com/iframe_api";
    j.async = true;
    f.parentNode.insertBefore(j, f);

// report the % played if it matches 0%, 25%, 50%, 75% or completed
function onPlayerPercent(e) {
    videoTime = e["getCurrentTime"]();
    videoTime = Math.ceil(videoTime);

    if(youtubeActionList.length === 0 ){
        return
    }

    let currentTimeActionId;
    let currentTimeActionInfo = {};
    let pass = false;
    for(let i = 0; i !== youtubeActionList.length; i += 1) {
        currentTimeActionInfo = youtubeActionList[i];
        currentTimeActionId = i;

        if(currentTimeActionInfo.hasOwnProperty('time') === false) {
            currentTimeActionInfo.time = 0;
        }

        if(currentTimeActionInfo.hasOwnProperty('id') === false) {
            currentTimeActionInfo.id = 0;
        }

        if(currentTimeActionInfo.time === videoTime) {
            pass = true
            break
        }
    }
    if(pass === false) {
        return;
    }
    /*
    let obj = {
                'id': 0,
                'action': 'clear',
                'time': videoTime,
                'startLineNumber': selectionData.selection.startLineNumber,
                'startColumn': selectionData.selection.startColumn,
                'endLineNumber': selectionData.selection.endLineNumber,
                'endColumn': selectionData.selection.endColumn,
                'youtubeCode': youtubeCode
            }
     */
    if(currentTimeActionInfo.action === 'selection') {
        oldDecorations = editor.deltaDecorations(oldDecorations, [{
            range: new monaco.Range(
                currentTimeActionInfo.startLineNumber,
                currentTimeActionInfo.startColumn,
                currentTimeActionInfo.endLineNumber,
                currentTimeActionInfo.endColumn
            ),
            options: { inlineClassName: 'decorBold' }
        },]);
    }

    if(currentTimeActionInfo.action === 'clear') {
        oldDecorations = editor.deltaDecorations(oldDecorations, [{
            range: new monaco.Range(
                currentTimeActionInfo.startLineNumber,
                currentTimeActionInfo.startColumn,
                currentTimeActionInfo.endLineNumber,
                currentTimeActionInfo.endColumn
            ),
            options: { inlineClassName: 'decorNone' }
        },]);
    }
}

// Crossbrowser onbeforeunload hack/proxy
// https://developer.mozilla.org/en-US/docs/WindowEventHandlers.onbeforeunload
window.onbeforeunload = function (e) {
    var e = e || window.event;
    // For IE and Firefox prior to version 4
    if(e)
        e.returnValue = 'na';
    // For Safari
    return 'na';
};
window.onbeforeunload = trackYTUnload;

function trackYTUnload() {
    for (var i = 0; i < gtmYTplayers.length; i++)
        if (gtmYTlisteners[i].getPlayerState() === 1) { // playing
            var video_data = gtmYTlisteners[i]['getVideoData'](),
                label = video_data.video_id+':'+video_data.title;
            dataLayer.push({
                event: 'youtube',
                action: 'exit',
                label: label
            });
        }
}

// This function waits video element to be ready for use
function verifyHtmlTagVideoReadyForUse(){
    let id = document.getElementById("video");
    if(id === null || id === undefined) {
        setTimeout(verifyHtmlTagVideoReadyForUse, 300)
    } else {
        onYouTubeIframeAPIReady();
    }
}

// attach our YT listener once the API is loaded
function onYouTubeIframeAPIReady() {
    for (let e = document.getElementsByTagName("iframe"), x = e.length; x--;) {
        if (/youtube.*\/embed/.test(e[x].src)) {
            gtmYTListeners.push(new YT.Player(e[x], {
                events: {
                    onStateChange: onPlayerStateChange,
                    onError: onPlayerError
                }
            }));
            YT.gtmLastAction = "p";
        }
    }
}

function appendToYoutubeActionList(obj) {
    youtubeActionList.push(obj);
}

function appendYoutubeVideo(youtubeCode) {
    let newElement = document.createElement('iframe');
    newElement.id = "video"
    newElement.className = "video"
    newElement.src = "https://www.youtube.com/embed/" + youtubeCode
    newElement.width = "333"
    newElement.height = "187"
    newElement.allowFullscreen = false

    document.body.appendChild(newElement);
}

function sendVideoData(dataToSend) {
    let xhr = new XMLHttpRequest();
    xhr.open("POST", videoDataUrl, true);
    xhr.setRequestHeader('Content-Type', 'application/json');
    xhr.send(
        JSON.stringify(dataToSend)
    );
}

// listen for play/pause, other states such as rewind and end could also be added
// also report % played every second

function onPlayerStateChange(e) {
    //e["data"] == YT.PlayerState.PLAYING && setTimeout(onPlayerPercent, 1000, e["target"]);
    var video_data = e.target["getVideoData"](),
        label = video_data.video_id+':'+video_data.title;
    if (e["data"] == YT.PlayerState.PLAYING){
        timeOutVideo = setInterval(onPlayerPercent, 300, e["target"]);
    }
    if (e["data"] == YT.PlayerState.PLAYING && YT.gtmLastAction == "p") {


        /*dataLayer.push({
            event: "youtube",
            action: "play",
            label: label
        });*/
        YT.gtmLastAction = "";
    }
    if (e["data"] == YT.PlayerState.PAUSED) {
        //clearTimeout(timeOutVideo);
        /*dataLayer.push({
            event: "youtube",
            action: "pause",
            label: label
        });*/
        YT.gtmLastAction = "p";
    }
}

// catch all to report errors through the GTM data layer
// once the error is exposed to GTM, it can be tracked in UA as an event!
// refer to https://developers.google.com/youtube/js_api_reference#Events onError
function onPlayerError(e) {
    /*dataLayer.push({
        event: "error",
        action: "GTM",
        label: "youtube:" + e
    })*/
}

var dataLayer = [];


var gtmYTListeners = []; // support multiple players on the same page
function mountYoutubeIFrame() {
    // OPTIONAL: Enable JSAPI if it's not already on the URL
    // note: this will cause the Youtube player to "flash" on the page when reloading to enable the JS API
    for (var e = document.getElementsByTagName("iframe"), x = e.length; x--;)
        if (/youtube.*\/embed/.test(e[x].src))
            if(e[x].src.indexOf('enablejsapi=') === -1)
                e[x].src += (e[x].src.indexOf('?') ===-1 ? '?':'&') + 'enablejsapi=1';
}