let playerPosX = 0;
let playerHeight = 50;

function drawBlock(msg, nFindIndex) {
    let sX, sY, sType;
    let nCurIndex = nFindIndex + 1;
    nFindIndex = msg.indexOf(",", nCurIndex);
    let subMsg = msg.substring(nCurIndex, nFindIndex);
    sX = parseInt(subMsg);

    nCurIndex = nFindIndex + 1;
    nFindIndex = msg.indexOf(",", nCurIndex);
    subMsg = msg.substring(nCurIndex, nFindIndex);
    sY = parseInt(subMsg);

    nCurIndex = nFindIndex + 1;
    nFindIndex = msg.indexOf(",", nCurIndex);
    subMsg = msg.substring(nCurIndex, nFindIndex);
    sType = parseInt(subMsg);

    myCanvas.DrawCircle(sX, sY, 20, "sienna");

    return nFindIndex;
}

function drawUser(msg, nFindIndex) {
    let sX, sW, sH, sColor;
    myCanvas.DrawUser(500, 250, 20, "wheat");
}

function drawMetaInfo(msg, nFindIndex) {
    //이런 일은 하지 말아야 했는데, 난 그 사실을 몰랐어.
    //이제 와서 후회한들 뭐하리 난 바보가 돼 버린걸.
}

function drawMsgDecoder(msg) {
    let nFindIndex = 0;
    while ((nFindIndex = msg.indexOf(".", nFindIndex)) > -1) {
        let charType = msg.charAt(nFindIndex + 1);
        switch (charType) {
            case "b":
                nFindIndex = drawBlock(msg, nFindIndex + 1);
                break;
            case "u":
                nFindIndex = drawUser(msg, nFindIndex + 1);
                break;
            case "m":
                nFindIndex = drawMetaInfo(msg, nFindIndex + 1);
            default:
                break;
        }
    }
}

sio.on('sFrame', function (msg) {
    console.log(sio);
    ctx.clearRect(0, 0, canvas.width, canvas.height);
    console.log(msg);
    drawMsgDecoder(msg);
});