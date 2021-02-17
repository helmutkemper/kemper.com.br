function youtubePause(editor, monaco) {
    editor.addAction({
        id: "pauseVideo",
        label: "Pause Youtube",
        keybindings: [monaco.KeyMod.CtrlCmd | monaco.KeyCode.KEY_3],
        contextMenuGroupId: "0codeVideoGroup",
        contextMenuOrder: 4.0,
        run: function() {
            let stopAllYouTubeVideos = () => {
                let iframes = document.querySelectorAll('iframe');
                Array.prototype.forEach.call(iframes, iframe => {
                    iframe.contentWindow.postMessage(JSON.stringify({ event: 'command',
                        func: 'pauseVideo' }), '*');
                });
            }
            stopAllYouTubeVideos();
        }
    });
}

function youtubePlay(editor, monaco) {
    editor.addAction({
        id: "playVideo",
        label: "Play Youtube",
        keybindings: [monaco.KeyMod.CtrlCmd | monaco.KeyCode.KEY_2],
        contextMenuGroupId: "0codeVideoGroup",
        contextMenuOrder: 3.0,
        run: function() {
            var stopAllYouTubeVideos = () => {
                var iframes = document.querySelectorAll('iframe');
                Array.prototype.forEach.call(iframes, iframe => {
                    iframe.contentWindow.postMessage(JSON.stringify({ event: 'command',
                        func: 'playVideo' }), '*');
                });
            }
            stopAllYouTubeVideos();
        }
    });
}

function youtubeSetVideoCode(editor, monaco) {
    editor.addAction({
        id: "videoCode",
        label: "set youtube video code",
        keybindings: [monaco.KeyMod.CtrlCmd | monaco.KeyCode.KEY_9],
        contextMenuGroupId: "0youtubeConfig",
        contextMenuOrder: 1.0,
        run: function() {
            youtubeCode = prompt("Qual o código do vídeo do Youtube?");

            if(youtubeCode === ''){
                youtubeCode = '5oYfO5qPv44'
            }

            appendYoutubeVideo(youtubeCode);
            mountYoutubeIFrame();


            verifyHtmlTagVideoReadyForUse();
        }
    });
}